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
	"errors"
	"sync"

	"github.com/simcord-llc/bitbon-system-blockchain/quorum/snapshot"

	"github.com/simcord-llc/bitbon-system-blockchain/crypto"

	"github.com/simcord-llc/bitbon-system-blockchain/common"

	"github.com/simcord-llc/bitbon-system-blockchain/log"

	"github.com/simcord-llc/bitbon-system-blockchain/p2p/enode"
)

var (
	errNoMajorityOfNodes = errors.New("there are no majority of nodes in peer list while getting snapshot")
)

type majorityChecker func(nodes int, addSelf bool) bool

// PeerList is a representation of quorum service peers list with additional helper methods
type PeerList struct {
	mu          sync.RWMutex
	self        *enode.Node
	selfAddress common.Address
	selfState   *SelfState

	peers map[enode.ID]*Peer // map of active peers (connected at the moment)
	// recentPeers is map of recent peers that was dropped from peers map some time ago.
	// Needed to store peer snapshots while short term reconnects and troubles
	recentPeers map[enode.ID]*peerState

	hasMajority majorityChecker
}

type PeerListDTO struct {
	SelfID               enode.ID
	SelfEnode            string
	SelfIsMinerCandidate bool
	Peers                []*PeerDTO
	RecentPeers          map[enode.ID]*PeerStateDTO
}

// newPeerList - PeerList snapshot
func newPeerList(self *enode.Node, hasMajority majorityChecker) *PeerList {
	return &PeerList{
		self:        self,
		selfAddress: crypto.PubkeyToAddress(*(self.Pubkey())),
		selfState:   NewSelfState(),
		peers:       make(map[enode.ID]*Peer),
		recentPeers: make(map[enode.ID]*peerState),
		hasMajority: hasMajority,
	}
}

func (pl *PeerList) ToDTO() *PeerListDTO {
	pl.mu.RLock()
	defer pl.mu.RUnlock()

	res := &PeerListDTO{
		SelfID:               pl.self.ID(),
		SelfEnode:            pl.self.String(),
		SelfIsMinerCandidate: pl.selfState.isMinerCandidate,
		Peers:                make([]*PeerDTO, 0, len(pl.peers)),
		RecentPeers:          make(map[enode.ID]*PeerStateDTO, len(pl.recentPeers)),
	}

	for _, peer := range pl.peers {
		res.Peers = append(res.Peers, peer.ToDTO())
	}

	for id, ps := range pl.recentPeers {
		res.RecentPeers[id] = ps.ToDTO()
	}

	return res
}

func (pl *PeerList) Self() *enode.Node {
	return pl.self
}

func (pl *PeerList) Peers() []*Peer {
	pl.mu.RLock()
	defer pl.mu.RUnlock()

	list := make([]*Peer, 0, len(pl.peers))
	for _, peer := range pl.peers {
		list = append(list, peer)
	}
	return list
}

// RemovePeer removes peer with peerID from the peer store
// actually deleting not happening. it just moves from peers map to recentPeers map
func (pl *PeerList) RemovePeer(peerID enode.ID) {
	pl.mu.Lock()
	defer pl.mu.Unlock()

	peer, ok := pl.peers[peerID]
	if ok {
		pl.recentPeers[peerID] = peer.GetState()
	}

	delete(pl.peers, peerID)
}

// AddPeer adds the peer using quorum protocol peer
// if peer state found in recent peers - restore state from that map
func (pl *PeerList) AddPeer(p *Peer) {
	pl.mu.Lock()
	defer pl.mu.Unlock()

	state, ok := pl.recentPeers[p.ID()]
	if ok {
		log.Trace("Peer`s state has restored from recent peers", "peer", p)
		p.setState(state)
	}

	delete(pl.recentPeers, p.ID())
	pl.peers[p.ID()] = p
}

// GetPeer returns the peer with peerID
func (pl *PeerList) GetPeer(peerID enode.ID) *Peer {
	pl.mu.RLock()
	defer pl.mu.RUnlock()

	return pl.peers[peerID]
}

func (pl *PeerList) Ids() []enode.ID {
	pl.mu.RLock()
	defer pl.mu.RUnlock()

	list := make([]enode.ID, 0, len(pl.peers))
	for id := range pl.peers {
		list = append(list, id)
	}
	return list
}

// Len returns the number of elements in list
func (pl *PeerList) Len() int {
	pl.mu.RLock()
	defer pl.mu.RUnlock()

	return len(pl.peers)
}

//---------------------- Ping list

// PingList returns the slice of enode.ID of ping list (doesn`t include self node)
func (pl *PeerList) PingList() []enode.ID {
	pl.mu.RLock()
	defer pl.mu.RUnlock()

	list := make([]enode.ID, 0, len(pl.peers))
	for id, peer := range pl.peers {
		if peer.InPingList() {
			list = append(list, id)
		}
	}
	return list
}

// PingListPeers returns the slice of peers of ping list (doesn`t include self node)
func (pl *PeerList) PingListPeers() []*Peer {
	pl.mu.RLock()
	defer pl.mu.RUnlock()

	list := make([]*Peer, 0, len(pl.peers))
	for _, peer := range pl.peers {
		if peer.InPingList() {
			list = append(list, peer)
		}
	}
	return list
}

// PingListLen returns the number of peers in ping list (doesn`t include self node)
func (pl *PeerList) PingListLen() (length int) {
	pl.mu.RLock()
	defer pl.mu.RUnlock()

	for _, peer := range pl.peers {
		if peer.InPingList() {
			length++
		}
	}
	return length
}

//---------------------- Node Candidate

// SelfIsMinerCandidate - returns if self node is miner candidate AT THE MOMENT (not depending do epoch)
func (pl *PeerList) SelfIsMinerCandidate() bool {
	return pl.selfState.IsMinerCandidate()
}

// MarkSelfIsMinerCandidate - marks self node is miner candidate AT THE MOMENT
func (pl *PeerList) MarkSelfIsMinerCandidate(in bool) {
	pl.selfState.MarkIsMinerCandidate(in)
}

// GetSnapshotNodes methods creates a snapshot of current peers state and returns slice of snapshot.Node
// Result also contains data about self node.
// Self node remoteNodeCandidates are also feels in while getting snapshot nodes
func (pl *PeerList) GetSnapshotNodes() ([]*snapshot.Node, error) {
	pl.mu.RLock()
	defer pl.mu.RUnlock()

	// base check - if there is no majority in ping list  - do not generate snapshot
	if !pl.hasMajority(len(pl.peers), true) {
		return nil, errNoMajorityOfNodes
	}

	var (
		nodes           []*snapshot.Node
		selfRemoteNodes = make(map[common.Address]bool)
	)
	// check and prepare peers
	for _, peer := range pl.peers {
		// SnapshotNode could be only quorum candidate node!!!
		if peer.IsQuorumCandidate() {
			selfRemoteNodes[peer.Address()] = peer.IsMinerCandidate()
			nodes = append(nodes, snapshot.NewNode(peer.ID(), peer.Address(), peer.IsMinerCandidate(), nil))
		}
	}

	selfRemoteNodes[pl.selfAddress] = pl.SelfIsMinerCandidate()
	nodes = append(nodes, snapshot.NewNode(pl.self.ID(), pl.selfAddress, pl.SelfIsMinerCandidate(), selfRemoteNodes))

	return nodes, nil
}
