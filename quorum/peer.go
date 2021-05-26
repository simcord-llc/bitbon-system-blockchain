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
	"time"

	"github.com/simcord-llc/bitbon-system-blockchain/common"
	"github.com/simcord-llc/bitbon-system-blockchain/crypto"

	"github.com/simcord-llc/bitbon-system-blockchain/p2p/enode"

	"github.com/simcord-llc/bitbon-system-blockchain/log"
	"github.com/simcord-llc/bitbon-system-blockchain/p2p/protocols"
)

// Peer is a representation of quorum service peer
type Peer struct {
	*protocols.Peer                 // represents the connection for online peer
	PingPeer        PingPeerService // represents the connection for online ping peer
	logger          log.Logger
	quorumHandler   Handlers

	address common.Address
	*peerState

	closeCh chan struct{}

	fakePingDelay time.Duration // fake ping delay for testing purposes
}

type PeerDTO struct {
	ID    enode.ID
	Enode string
	*PeerStateDTO
}

func (p *Peer) ToDTO() *PeerDTO {
	return &PeerDTO{
		ID:           p.ID(),
		Enode:        p.Node().String(),
		PeerStateDTO: p.peerState.ToDTO(),
	}
}

// NewPeer creates peer with given quorum protocol peer
func NewPeer(p *protocols.Peer, pingPeer PingPeerService, quorum Handlers) *Peer {
	peer := &Peer{
		Peer:          p,
		PingPeer:      pingPeer,
		logger:        log.New("type", "quorumPeer", "peer", p),
		quorumHandler: quorum,
		peerState:     &peerState{lastPing: new(int64)},
		address:       crypto.PubkeyToAddress(*(p.Node().Pubkey())),
		closeCh:       make(chan struct{}),
	}
	return peer
}

func (p *Peer) Close() error {
	if !p.Closed() {
		close(p.closeCh)
	}

	return nil
}

func (p *Peer) Closed() bool {
	select {
	case <-p.closeCh:
		return true
	default:
		return false
	}
}

func (p *Peer) pingWatch(pingFrequency, quorumMaxPing time.Duration) {
	// init ticker with smallest value for first call (do not wait for main pingFrequency)
	ticker := time.NewTicker(time.Nanosecond)
	for {
		select {
		case <-p.closeCh:
			return
		default:
		}

		select {
		case <-ticker.C:
			ticker.Stop()
			// measure ping duration via calling ping-pong

			var (
				pingDuration time.Duration
				err          error
			)

			if p.fakePingDelay == 0 {
				if pingDuration, err = p.PingPeer.PingPong(quorumMaxPing); err != nil {
					log.Debug(fmt.Sprintf("%s: protocol ping peer error on peer %s: %v", SpecString(Spec), p.ID().String(), err))
				}
			} else {
				pingDuration = p.fakePingDelay
			}

			p.setLastPing(pingDuration)
			if pingDuration > quorumMaxPing || pingDuration == 0 {
				log.Debug("PEER MARKED NOT IN PING LIST", "peer", p, "time took", pingDuration, "maxPing", quorumMaxPing)
				p.markInPingList(false)
			} else {
				p.markInPingList(true)
			}

			ticker = time.NewTicker(pingFrequency)
		case <-p.closeCh:
			return
		}
	}
}

func (p *Peer) GetState() *peerState { // nolint
	return p.peerState
}

func (p *Peer) Address() common.Address {
	return p.address
}

func (p *Peer) setState(state *peerState) {
	p.peerState = state
}

func (p *Peer) String() string {
	return fmt.Sprintf("Quorum peer: (%s), ping peer: (%s)", p.Peer.String(), p.PingPeer.String())
}

// HandleMsg is the message handler that delegates incoming messages
func (p *Peer) HandleMsg(_ context.Context, msg interface{}) error {
	switch msg := msg.(type) {
	case *VoteMinerAddMsg:
		go p.quorumHandler.HandleVoteMinerAddMsg(p)
		return nil

	case *VoteMinerDelMsg:
		go p.quorumHandler.HandleVoteMinerDelMsg(p)
		return nil

	case *NodeCandidateListMsg:
		go p.quorumHandler.HandleNodeCandidateListMsg(msg, p)
		return nil

	case *MinersMsg:
		go p.quorumHandler.HandleMinersMsg(msg, p)
		return nil

	default:
		return fmt.Errorf("unknown message type: %T", msg)
	}
}

// ---------------------- Send methods

// SendVoteMinerAddMsg sends VoteMinerAddMsg
func (p *Peer) SendVoteMinerAddMsg(ctx context.Context) error {
	msg := NewVoteMinerAddMsg()
	err := p.Send(ctx, msg)
	return err
}

// SendVoteMinerDelMsg sends NewVoteMinerDelMsg
func (p *Peer) SendVoteMinerDelMsg(ctx context.Context) error {
	msg := NewVoteMinerDelMsg()
	err := p.Send(ctx, msg)
	return err
}

// SendNodeCandidateList sends MinerCandidateListMsg
func (p *Peer) SendNodeCandidateList(ctx context.Context, msg *NodeCandidateListMsg) error {
	err := p.Send(ctx, msg)
	return err
}

func (p *Peer) SendMiners(ctx context.Context, msg *MinersMsg) error {
	err := p.Send(ctx, msg)
	return err
}
