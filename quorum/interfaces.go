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
	"time"

	"github.com/simcord-llc/bitbon-system-blockchain/common"
	"github.com/simcord-llc/bitbon-system-blockchain/p2p/enode"
	"github.com/simcord-llc/bitbon-system-blockchain/quorum/snapshot"
)

type ContractService interface {
	GetCurrentDistribution(ctx context.Context) (distribution map[common.Address]uint64, err error)
}

type P2PServer interface {
	Self() *enode.Node
}

type PeerListService interface {
	PingListLen() (length int)
	PingListPeers() []*Peer
	PingList() []enode.ID
	ToDTO() *PeerListDTO
	AddPeer(p *Peer)
	RemovePeer(peerID enode.ID)
	Peers() []*Peer
	SelfIsMinerCandidate() bool
	MarkSelfIsMinerCandidate(in bool)
	GetSnapshotNodes() ([]*snapshot.Node, error)
}

type SnapshotService interface {
	GetSnapshot(epoch uint64) *snapshot.Snapshot
	ToDTO() *snapshot.SystemDTO
	Epochs() []uint64
	RecentEpoch() uint64
	MinerListAtEpoch(epoch uint64) []common.Address
	PrepareSnapshot(epoch uint64) (*snapshot.Snapshot, error)
}

type Handlers interface {
	HandleVoteMinerAddMsg(peer *Peer)
	HandleVoteMinerDelMsg(peer *Peer)
	HandleNodeCandidateListMsg(msg *NodeCandidateListMsg, peer *Peer)
	HandleMinersMsg(msg *MinersMsg, peer *Peer)
}

type PingPeerService interface {
	PingPong(timeout time.Duration) (dur time.Duration, err error)
	String() string
}
