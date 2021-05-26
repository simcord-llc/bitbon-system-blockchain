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

	"github.com/simcord-llc/bitbon-system-blockchain/common"

	"github.com/simcord-llc/bitbon-system-blockchain/log"
	"github.com/simcord-llc/bitbon-system-blockchain/p2p/enode"
)

type MinerList struct {
	mu     sync.RWMutex
	miners []common.Address
	hash   common.Hash
}

type MinerListDTO struct {
	Miners []common.Address
	Hash   common.Hash
}

func (ml *MinerList) ToDTO() *MinerListDTO {
	ml.mu.RLock()
	defer ml.mu.RUnlock()

	return &MinerListDTO{
		Miners: ml.miners,
		Hash:   ml.hash,
	}
}

func (ml *MinerList) Hash() common.Hash {
	ml.mu.RLock()
	defer ml.mu.RUnlock()
	return ml.hash
}

func (ml *MinerList) Miners() []common.Address {
	ml.mu.RLock()
	defer ml.mu.RUnlock()
	return ml.miners
}

type majorityChecker func(nodes int, addSelf bool) bool

// ----------------------------------------

type Snapshot struct {
	mu sync.RWMutex

	epoch     uint64
	selfID    enode.ID
	nodes     []*Node
	minerList *MinerList

	hasMajority majorityChecker
	logger      log.Logger
}

type SnapDTO struct {
	Epoch     uint64
	SelfID    enode.ID
	Nodes     []*NodeDTO
	MinerList *MinerListDTO
}

func NewSnapshot(epoch uint64, selfID enode.ID, nodes []*Node,
	hasMajority majorityChecker, logger log.Logger) *Snapshot {
	s := &Snapshot{
		epoch:       epoch,
		selfID:      selfID,
		nodes:       nodes,
		hasMajority: hasMajority,
		logger:      logger,
	}

	return s
}

func (s *Snapshot) ToDTO() *SnapDTO {
	s.mu.RLock()
	defer s.mu.RUnlock()

	nodes := make([]*NodeDTO, len(s.nodes))
	for i, n := range s.nodes {
		nodes[i] = n.ToDTO()
	}

	var minerListDTO *MinerListDTO
	if s.minerList != nil {
		minerListDTO = s.minerList.ToDTO()
	}

	return &SnapDTO{
		Epoch:     s.epoch,
		SelfID:    s.selfID,
		Nodes:     nodes,
		MinerList: minerListDTO,
	}
}

func (s *Snapshot) ToMap() map[common.Address]bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	dto := make(map[common.Address]bool, len(s.nodes))
	for _, n := range s.nodes {
		dto[n.getAddress()] = n.getIsMinerCandidate()
	}

	return dto
}

// MinerList returns the slice of peers enode.Node of miner list at snapshot if is gathered
func (s *Snapshot) MinerList() *MinerList {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.minerList
}

// TryGetSelfMinerList tries to return the slice of peers enode.Node of miner list at snapshot
func (s *Snapshot) TryGetSelfMinerList() *MinerList {
	s.mu.Lock()
	defer s.mu.Unlock()

	if minerList := s.getRemotePeerMiners(s.selfID, s.logger); minerList != nil {
		return minerList
	}

	// base check - there are should be more then one node in list (self node)
	if len(s.nodes) <= 1 {
		s.logger.Warn("Not enough nodes in snapshot", "len(s.nodes)", len(s.nodes))
		return nil
	}

	// nodesWithRemotes includes self node data
	nodesWithRemotes := 0
	for _, node := range s.nodes {
		if node.isSetRemoteNodeCandidates() {
			nodesWithRemotes++
		}
	}
	// base check - there should be a majority of nodes who sent there RemoteNodeCandidates
	if !s.hasMajority(nodesWithRemotes, false) {
		s.logger.Warn("There is no majority of nodes in snapshot who sent there RemoteNodeCandidates")
		return nil
	}

	var miners []common.Address

	// at first - gather all possible miner candidates
	allMiners := make(map[common.Address]struct{})
	for idx := range s.nodes {
		allAddresses := s.nodes[idx].getRemoteMinerCandidates()
		for idx := range allAddresses {
			allMiners[allAddresses[idx]] = struct{}{}
		}
	}

	for address := range allMiners {
		minerInRemotes := 0
		// check miner candidate to be present in RemoteMinerCandidates
		for idx := range s.nodes {
			if s.nodes[idx].hasRemoteMinerCandidate(address, s.logger) {
				minerInRemotes++
			}
		}
		// if there is majority  - add miner to miner list
		if s.hasMajority(minerInRemotes, false) {
			miners = append(miners, address)
		}
	}

	if len(miners) == 0 {
		s.logger.Warn("Can not gather miners using algorithm")
		return nil
	}

	hash := common.Keccak256Addresses(miners...)
	minerList := &MinerList{
		miners: miners,
		hash:   hash,
	}

	// set miners hash for self node (it won`t be sent via p2p)
	s.setRemotePeerMiners(s.selfID, miners, hash, s.logger)

	return minerList
}

// SetRemoteMinerCandidates set miner candidate list of remote peer
func (s *Snapshot) SetRemotePeerNodeCandidateList(id enode.ID, nodes map[common.Address]bool) { // nolint:interfacer
	s.mu.RLock()
	defer s.mu.RUnlock()

	// base check - skip setting if minerList is gathered
	if s.minerList != nil {
		s.logger.Debug("Try to set RemoteNodeCandidates when minerList has already been gathered")
		return
	}

	var snapshotNode *Node

	for _, node := range s.nodes {
		if node.getID() == id {
			snapshotNode = node
			break
		}
	}
	if snapshotNode == nil {
		s.logger.Debug("Try to set RemoteNodeCandidates to peer that is not exists in snapshot.")
		return
	}
	if snapshotNode.isSetRemoteNodeCandidates() {
		s.logger.Debug("Try to set RemoteNodeCandidates to peer that already has RemoteNodeCandidates. Skip..")
		return
	}

	snapshotNode.setRemoteNodeCandidates(nodes)
}

func (s *Snapshot) TryApproveMinerList() {
	s.mu.RLock()
	defer s.mu.RUnlock()

	// base check - minerList set
	if s.minerList != nil {
		s.logger.Debug("Try approve miner list which has already been approved. Skip..")
		return
	}

	// choose the miner hash and miners which are present in majority of nodes
	for _, node := range s.nodes {
		miners, hash := node.getRemoteMiners()

		if miners == nil || hash == (common.Hash{}) {
			continue
		}

		nodesWithHash := 0
		for _, node := range s.nodes {
			if node.equalRemoteMinersHash(hash) {
				nodesWithHash++
			}
		}

		if s.hasMajority(nodesWithHash, false) {
			// there is a majority with this hash
			s.minerList = &MinerList{
				hash:   hash,
				miners: miners,
			}
		}
	}
}

// SetRemoteMinerCandidates set miner candidate list of remote peer
func (s *Snapshot) SetRemotePeerMiners(id enode.ID, miners []common.Address, hash common.Hash) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	// base check - skip setting if minerList is gathered and approved
	if s.minerList != nil {
		s.logger.Debug("Try to set RemoteMinersHash to peer when miner list has already been approved. Skip..")
		return
	}

	s.setRemotePeerMiners(id, miners, hash, s.logger)
}

func (s *Snapshot) getSnapshotNode(id enode.ID) (snapshotNode *Node) {
	for _, node := range s.nodes {
		if node.getID() == id {
			snapshotNode = node
			break
		}
	}
	return
}

// setRemotePeerMiners set miner candidate list of remote peer
func (s *Snapshot) setRemotePeerMiners(id enode.ID, miners []common.Address, hash common.Hash, logger log.Logger) {
	snapshotNode := s.getSnapshotNode(id)

	if snapshotNode == nil {
		logger.Debug("Try to set RemoteMinersHash to peer that is not exists in snapshot.", "id", id.String())
		return
	}

	if !snapshotNode.isSetRemoteNodeCandidates() {
		logger.Debug("Try to set RemoteMinersHash to peer without RemoteNodeCandidates. Skip..", "id", id.String())
		return
	}

	snapshotNode.setRemoteMiners(miners, hash)
}

func (s *Snapshot) getRemotePeerMiners(id enode.ID, logger log.Logger) *MinerList {
	snapshotNode := s.getSnapshotNode(id)

	if snapshotNode == nil {
		logger.Debug("Try to get RemoteMinersHash to peer that is not exists in snapshot.", "id", id.String())
		return nil
	}

	miners, hash := snapshotNode.getRemoteMiners()
	if miners == nil {
		return nil
	}

	return &MinerList{
		miners: miners,
		hash:   hash,
	}
}
