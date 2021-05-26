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

package snapshot

import (
	"sync"

	"github.com/simcord-llc/bitbon-system-blockchain/log"

	"github.com/simcord-llc/bitbon-system-blockchain/common"

	"github.com/simcord-llc/bitbon-system-blockchain/p2p/enode"
)

// Node could be only quorum candidate node!!!
type Node struct {
	mu               sync.RWMutex
	id               enode.ID // node ID
	address          common.Address
	isMinerCandidate bool // indicator if this peer wants to mine AT THE EPOCH (depending to epoch)
	// remoteNodeCandidates -> node candidate list of this remote peer (sent by p2p).
	// map of quorum candidate to miner flag
	remoteNodeCandidates map[common.Address]bool
	remoteMiners         remoteMiners // miners (list and checksum) of this remote peer (sent by p2p).
}

type remoteMiners struct {
	hash   common.Hash
	miners []common.Address
}

type NodeDTO struct {
	ID                   enode.ID
	Address              common.Address
	IsMinerCandidate     bool
	RemoteNodeCandidates map[common.Address]bool
	RemoteMiners         RemoteMinersDTO
}

type RemoteMinersDTO struct {
	Hash   common.Hash
	Miners []common.Address
}

func (n *Node) ToDTO() *NodeDTO {
	n.mu.RLock()
	defer n.mu.RUnlock()

	return &NodeDTO{
		ID:                   n.id,
		Address:              n.address,
		IsMinerCandidate:     n.isMinerCandidate,
		RemoteNodeCandidates: n.remoteNodeCandidates,
		RemoteMiners: RemoteMinersDTO{
			Hash:   n.remoteMiners.hash,
			Miners: n.remoteMiners.miners,
		},
	}
}

func NewNode(id enode.ID, address common.Address, isMinerCandidate bool,
	remoteNodeCandidates map[common.Address]bool) *Node {
	if remoteNodeCandidates == nil {
		remoteNodeCandidates = make(map[common.Address]bool)
	}

	return &Node{
		id:                   id,
		address:              address,
		isMinerCandidate:     isMinerCandidate,
		remoteNodeCandidates: remoteNodeCandidates,
	}
}

// ID return node candidate id.
func (n *Node) getID() enode.ID {
	n.mu.RLock()
	defer n.mu.RUnlock()
	return n.id
}

// Address return node candidate address.
func (n *Node) getAddress() common.Address {
	n.mu.RLock()
	defer n.mu.RUnlock()
	return n.address
}

// IsMinerCandidate return miner candidate status in self miner candidate list.
func (n *Node) getIsMinerCandidate() bool {
	n.mu.RLock()
	defer n.mu.RUnlock()
	return n.isMinerCandidate
}

// IsSetRemoteNodeCandidates returns if node candidate sent its own node candidates
func (n *Node) isSetRemoteNodeCandidates() bool {
	n.mu.RLock()
	defer n.mu.RUnlock()
	return len(n.remoteNodeCandidates) > 0
}

// IsSetRemoteNodeCandidates returns if node candidate sent its own node candidates
func (n *Node) hasRemoteMinerCandidate(address common.Address, logger log.Logger) bool {
	if logger == nil {
		logger = log.New()
	}

	n.mu.RLock()
	defer n.mu.RUnlock()

	isMiner, ok := n.remoteNodeCandidates[address]
	// node is not exists in remoteNodeCandidates
	if !ok {
		logger.Error("Node does not have miner in its remoteNodeCandidates",
			"Node", n.id.String(), "wantedMiner", address.String())
		return false
	}
	if !isMiner {
		logger.Error("Node has miner in its remoteNodeCandidates, but it is not miner for it",
			"Node", n.id.String(), "minerCandidate", address.String())
		return false
	}
	return true
}

// GetRemoteQuorumCandidates returns quorum candidate list of remote peers
func (n *Node) getRemoteMinerCandidates() []common.Address {
	n.mu.RLock()
	defer n.mu.RUnlock()
	var minerCandidates []common.Address
	for nodeAddress := range n.remoteNodeCandidates {
		if n.remoteNodeCandidates[nodeAddress] {
			minerCandidates = append(minerCandidates, nodeAddress)
		}
	}
	return minerCandidates
}

// EqualRemoteMinersHash returns if remoteMinersHash of node equals given
func (n *Node) equalRemoteMinersHash(hash common.Hash) bool {
	n.mu.RLock()
	defer n.mu.RUnlock()
	return n.remoteMiners.hash == hash
}

// SetRemoteNodeCandidates sets node candidate list of remote peers
func (n *Node) setRemoteNodeCandidates(nodeCandidates map[common.Address]bool) {
	n.mu.Lock()
	defer n.mu.Unlock()
	n.remoteNodeCandidates = nodeCandidates
}

// SetRemoteMinersHash sets miners hash of remote peers
func (n *Node) setRemoteMiners(miners []common.Address, hash common.Hash) {
	n.mu.Lock()
	defer n.mu.Unlock()
	n.remoteMiners.miners = miners
	n.remoteMiners.hash = hash
}

// getRemoteMiners returns miners addresses of remote peers
func (n *Node) getRemoteMiners() ([]common.Address, common.Hash) {
	n.mu.RLock()
	defer n.mu.RUnlock()
	return n.remoteMiners.miners, n.remoteMiners.hash
}
