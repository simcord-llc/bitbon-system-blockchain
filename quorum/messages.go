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
	"math/big"

	"github.com/simcord-llc/bitbon-system-blockchain/common"
)

type NodeCandidateMsg struct {
	Address          common.Address
	IsMinerCandidate bool
}

func NewNodeCandidateMsg(address common.Address, isMinerCandidate bool) *NodeCandidateMsg {
	return &NodeCandidateMsg{Address: address, IsMinerCandidate: isMinerCandidate}
}

type NodeCandidateListMsg struct {
	List  []*NodeCandidateMsg
	Epoch *big.Int
}

func (m *NodeCandidateListMsg) ToMap() map[common.Address]bool {
	list := make(map[common.Address]bool, len(m.List))
	for _, nc := range m.List {
		list[nc.Address] = nc.IsMinerCandidate
	}

	return list
}

func NewNodeCandidateListMsg(nodes map[common.Address]bool, epoch uint64) *NodeCandidateListMsg {
	list := make([]*NodeCandidateMsg, 0, len(nodes))
	for id, isMinerCandidate := range nodes {
		list = append(list, NewNodeCandidateMsg(id, isMinerCandidate))
	}
	return &NodeCandidateListMsg{List: list, Epoch: new(big.Int).SetUint64(epoch)}
}

// VoteMinerAddMsg is sent to ask remote peers to add self node to miner nodes
type VoteMinerAddMsg struct {
}

func NewVoteMinerAddMsg() *VoteMinerAddMsg {
	return &VoteMinerAddMsg{}
}

// VoteMinerDelMsg is sent to ask remote peers to delete self node from miner nodes
type VoteMinerDelMsg struct {
}

func NewVoteMinerDelMsg() *VoteMinerDelMsg {
	return &VoteMinerDelMsg{}
}

type MinersMsg struct {
	Epoch  *big.Int
	Hash   common.Hash
	Miners []common.Address
}

func NewMinersMsg(epoch uint64, miners []common.Address, hash common.Hash) *MinersMsg {
	return &MinersMsg{Epoch: new(big.Int).SetUint64(epoch), Hash: hash, Miners: miners}
}
