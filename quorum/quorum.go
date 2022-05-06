// Copyright Simcord LLC
// This file is part of the Bitbon System Blockchain Node library.
//
// The Bitbon System Blockchain Node library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The Bitbon System Blockchain Node library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the Bitbon System Blockchain Node library. If not, see <http://www.gnu.org/licenses/>.

package quorum

import (
	"context"
	"fmt"
	"math/big"
	"sync"
	"sync/atomic"
	"time"

	"github.com/simcord-llc/bitbon-system-blockchain/consensus"

	"github.com/davecgh/go-spew/spew"

	"github.com/simcord-llc/bitbon-system-blockchain/quorum/snapshot"

	"github.com/simcord-llc/bitbon-system-blockchain/consensus/quorum"
	"github.com/simcord-llc/bitbon-system-blockchain/rpc"

	"github.com/simcord-llc/bitbon-system-blockchain/common"
	"github.com/simcord-llc/bitbon-system-blockchain/crypto"
	"github.com/simcord-llc/bitbon-system-blockchain/p2p/protocols"

	"github.com/simcord-llc/bitbon-system-blockchain/quorum/ping"

	_ "github.com/simcord-llc/bitbon-system-blockchain" // nolint
	"github.com/simcord-llc/bitbon-system-blockchain/log"
	"github.com/simcord-llc/bitbon-system-blockchain/node"
	"github.com/simcord-llc/bitbon-system-blockchain/p2p"
	"github.com/simcord-llc/bitbon-system-blockchain/p2p/enode"
)

type ContractServiceWith func(ctx context.Context) (distribution map[common.Address]uint64, err error)

func (c ContractServiceWith) GetCurrentDistribution(ctx context.Context) (distribution map[common.Address]uint64,
	err error) {
	return c(ctx)
}

// Spec (QuorumSpec) is the spec of the quorum protocol ...
var Spec = &protocols.Spec{
	Name:       "qrm",
	Version:    1,
	MaxMsgSize: 10 * 1024 * 1024,
	Messages: []interface{}{
		VoteMinerAddMsg{},
		VoteMinerDelMsg{},
		NodeCandidateListMsg{},
		MinersMsg{},
	},
}

type Quorum struct {
	ping *ping.Protocol

	cfg *Config

	srv P2PServer

	logger log.Logger

	engine    consensus.Engine
	contracts ContractService // Contract Service for decoupled interactions with Bitbon contracts

	// allNodes set once - so for the moment do not protect it via mutex
	allNodes       map[enode.ID]*enode.Node // list of all known nodes from static and trusted nodes (self node excluded)
	peerList       PeerListService          // list of quorum service peers
	snapshotSystem SnapshotService          // snapshot system for gathering miner candidates per epoch

	addressesEnodeIDMux sync.RWMutex
	addressesEnodeID    map[common.Address]enode.ID // store for every miner address it enode.ID

	done          chan struct{}
	started       chan struct{} // channel to indicate if Quorum service started (channel will be closed)
	peersGathered *int32        // indicator if majority of peers connected

	peerHook func(p *Peer) // hook for testing purposes that is being called from runProtocol
}

func New(_ *node.ServiceContext, cfg *Config) (q *Quorum, err error) {
	q = &Quorum{
		ping:             ping.NewProtocol(cfg.HandshakeTimeout),
		cfg:              cfg,
		logger:           log.New("type", "quorumService"),
		allNodes:         make(map[enode.ID]*enode.Node),
		done:             make(chan struct{}),
		started:          make(chan struct{}),
		peersGathered:    new(int32),
		addressesEnodeID: make(map[common.Address]enode.ID),
	}

	q.logger.Info("Quorum service created")

	return q, nil
}

func (q *Quorum) Start(srvr *p2p.Server) error {
	peerList := newPeerList(srvr.Self(), q.hasMajority)
	snapshotSystem := snapshot.NewSystem(q.CreateQuorumSnapshot, q.logger.New("subComponent", "SnapshotSystem"))
	return q.start(srvr, peerList, snapshotSystem)
}

func (q *Quorum) start(srvr *p2p.Server, peerList PeerListService, snapshotSystem SnapshotService) error {
	q.logger.Info("Quorum service started")

	q.srv = srvr

	q.setAllNodes(append(srvr.StaticNodes, srvr.TrustedNodes...))
	q.peerList = peerList
	q.snapshotSystem = snapshotSystem

	go q.waitForPeers()

	close(q.started)
	return nil
}

func (q *Quorum) Stop() error {
	if q.Stopped() {
		q.logger.Error("Quorum service already stopped")
	} else {
		close(q.done)
	}

	q.logger.Info("Quorum service stopped")
	return nil
}

// APIs is a meaningless implementation of node.Service.
func (q *Quorum) APIs() []rpc.API {
	return []rpc.API{
		{
			Namespace: ProtocolName,
			Version:   ProtocolVersionStr,
			Service:   NewAPI(q),
			Public:    true,
		},
	}
}

// Protocols is a meaningless implementation of node.Service.
func (q *Quorum) Protocols() []p2p.Protocol {
	return append(q.ping.Protocols(), []p2p.Protocol{
		{
			Name:    Spec.Name,
			Version: Spec.Version,
			Length:  Spec.Length(),
			Run:     q.runProtocol,
		},
	}...)
}

func (q *Quorum) Stopped() bool {
	select {
	case <-q.done:
		return true
	default:
		return false
	}
}

func (q *Quorum) Prepared() bool {
	return atomic.LoadInt32(q.peersGathered) == 1
}

func (q *Quorum) markPeersGathered(res bool) {
	if res {
		atomic.StoreInt32(q.peersGathered, 1)
	} else {
		atomic.StoreInt32(q.peersGathered, 0)
	}
}

func (q *Quorum) Engine() consensus.Engine {
	return q.engine
}

func (q *Quorum) SetEngine(engine *quorum.QuorumEngine) {
	q.engine = engine
}

func (q *Quorum) SetContractService(c ContractService) {
	q.contracts = c
}

// runProtocol is the p2p protocol run function for the qrm base protocol
func (q *Quorum) runProtocol(p *p2p.Peer, rw p2p.MsgReadWriter) (err error) {
	// wait for quorum service started
	// (some times runProtocol on peer starts earlier than quorum service Start method)
	<-q.started
	q.logger.Debug("quorum protocol", "peer", p)
	defer func() {
		q.logger.Debug("quorum protocol terminated", "peer", p, "error", err)
	}() // wait for the ping protocol to perform the handshake
	state, _ := q.ping.GetState(p.ID())
	defer q.ping.RemoveState(p.ID())

	select {
	case <-state.DoneCh():
	case <-time.After(q.cfg.PingProtocolTimeout()):
		return fmt.Errorf("%s: protocol timeout waiting for handshake on peer %s", SpecString(Spec), p.ID().String())
	}
	if state.Error() != nil {
		return fmt.Errorf("%s: protocol closed on peer %s: %v", SpecString(Spec), p.ID().String(), state.Error())
	}

	pingPeer := q.ping.GetPeer(p.ID())
	if pingPeer == nil {
		return fmt.Errorf("%s: protocol can`t get ping peer on %s", SpecString(Spec), p.ID().String())
	}

	// create quorumPeer
	quorumPeer := NewPeer(protocols.NewPeer(p, rw, Spec), pingPeer, q)
	defer quorumPeer.Close()

	if q.peerHook != nil {
		q.peerHook(quorumPeer)
	}

	go quorumPeer.pingWatch(q.cfg.PingFrequency, q.cfg.QuorumMaxPing)

	q.peerList.AddPeer(quorumPeer)
	defer q.peerList.RemovePeer(quorumPeer.ID())

	go q.initPeer(quorumPeer)

	// store miner (coinbase) address map to enode.ID
	address := crypto.PubkeyToAddress(*p.Node().Pubkey())

	q.addressesEnodeIDMux.Lock()
	q.addressesEnodeID[address] = p.ID()
	q.addressesEnodeIDMux.Unlock()

	return quorumPeer.Run(quorumPeer.HandleMsg)
}

// initPeer sends all initial information to peer
func (q *Quorum) initPeer(peer *Peer) {
	if q.SelfIsMinerCandidate() {
		if err := peer.SendVoteMinerAddMsg(context.TODO()); err != nil {
			q.logger.Trace("failed to notify VoteMinerAdd while init peer", "peer", peer, "error", err)
		}
	}
}

func (q *Quorum) setAllNodes(nodes []*enode.Node) {
	for _, n := range nodes {
		q.allNodes[n.ID()] = n
	}
	// there is a chance that self node in given nodes => so delete it
	delete(q.allNodes, q.srv.Self().ID())
}

func (q *Quorum) hasMajority(nodes int, addSelf bool) bool {
	// add self to given number of nodes
	totalNodes := len(q.allNodes) + 1
	if addSelf {
		nodes++
	}

	coefficient := new(big.Float).Quo(big.NewFloat(float64(nodes)), big.NewFloat(float64(totalNodes))).
		SetPrec(coefficientPrecision)
	return coefficient.Cmp(majorityCoefficient) >= 0
}

func (q *Quorum) MinerListAtEpoch(epoch uint64) []common.Address {
	minerAddresses := q.snapshotSystem.MinerListAtEpoch(epoch)

	if len(minerAddresses) == 0 {
		return nil
	}

	if q.contracts == nil {
		return minerAddresses
	}

	// In order to take mining agent distributions into account
	// we get current distribution from smart contracts and intersect it with quorum miner list
	// If an error occurs while getting the current distribution from smart contracts
	// or if the distribution is empty, the quorum miner list remains unchanged
	contractMinerAddressesMap, err := q.contracts.GetCurrentDistribution(context.TODO())
	if err != nil || len(contractMinerAddressesMap) == 0 {
		log.Debug("Current distribution from contracts is not relevant",
			"error", err, "contractMinerAddressesMap", contractMinerAddressesMap)
		return minerAddresses
	}

	buf := make([]common.Address, 0, len(minerAddresses))
	for idx := range minerAddresses {
		if _, ok := contractMinerAddressesMap[minerAddresses[idx]]; ok {
			buf = append(buf, minerAddresses[idx])
		}
	}
	minerAddresses = buf

	return minerAddresses
}

func (q *Quorum) MinerListAtRecentEpoch() []common.Address {
	epoch := q.snapshotSystem.RecentEpoch()
	if list := q.snapshotSystem.MinerListAtEpoch(epoch); list != nil {
		return list
	}
	return q.snapshotSystem.MinerListAtEpoch(epoch - 1)
}

func (q *Quorum) PrepareNodeListAtEpoch(epoch uint64) {
	s, err := q.snapshotSystem.PrepareSnapshot(epoch)
	if err != nil {
		if err != snapshot.ErrAlreadyExists {
			q.logger.Debug("PrepareNodeListAtEpoch", "epoch", epoch, "error", err)
		}
		return
	}

	go func() {
		q.sendNodeCandidateList(q.peerList.Peers(), NewNodeCandidateListMsg(s.ToMap(), epoch))

		// wait for PrepareMinersTimeout duration and try to get miner list of snapshot
		// if miner list gathered  - send miner hash to all peers to approve it
		timer := time.NewTimer(q.cfg.PrepareMinersTimeout)
		defer timer.Stop()

		<-timer.C
		minerList := s.TryGetSelfMinerList()
		if minerList != nil {
			q.sendMiners(q.peerList.Peers(), NewMinersMsg(epoch, minerList.Miners(), minerList.Hash()))
			s.TryApproveMinerList()
		} else {
			q.logger.Trace("CAN NOT GetMinerList AFTER TIMEOUT", "epoch",
				epoch, "timeout", q.cfg.PrepareMinersTimeout, "snapshot", spew.Sdump(s.ToDTO()))
		}
	}()
}

// ---------------------- minerCandidate

// SelfIsMinerCandidate returns self miner candidate status
func (q *Quorum) SelfIsMinerCandidate() bool {
	return q.peerList.SelfIsMinerCandidate()
}

// MarkMinerCandidate
// when node enters in quorum, it marks self as miner candidate
func (q *Quorum) MarkMinerCandidate(res bool) {
	q.peerList.MarkSelfIsMinerCandidate(res)
	if res {
		q.notifyVoteMinerAdd(q.peerList.Peers())
	} else {
		q.notifyVoteMinerDel(q.peerList.Peers())
	}
}

// CreateQuorumSnapshot creates a snapshot of the current state of peer list
func (q *Quorum) CreateQuorumSnapshot(epoch uint64, logger log.Logger) (*snapshot.Snapshot, error) {
	snapshotNodes, err := q.peerList.GetSnapshotNodes()
	if err != nil {
		return nil, err
	}
	return snapshot.NewSnapshot(epoch, q.srv.Self().ID(), snapshotNodes, q.hasMajority, logger), nil
}
