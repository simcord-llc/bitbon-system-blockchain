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
	"github.com/simcord-llc/bitbon-system-blockchain/log"
)

// HandleVoteMinerAddMsg handles VoteMinerAddMsg from peer
func (q *Quorum) HandleVoteMinerAddMsg(peer *Peer) {
	log.Trace("HandleVoteMinerAddMsg", "peer", peer)
	peer.MarkIsMinerCandidate(true)
}

// HandleVoteMinerDelMsg handles VoteMinerDelMsg from peer
func (q *Quorum) HandleVoteMinerDelMsg(peer *Peer) {
	log.Trace("HandleVoteMinerDelMsg", "peer", peer)
	peer.MarkIsMinerCandidate(false)
}

// HandleNodeCandidateListMsg handles NodeCandidateListMsg from peer
func (q *Quorum) HandleNodeCandidateListMsg(msg *NodeCandidateListMsg, peer *Peer) {
	epoch := msg.Epoch.Uint64()

	log.Trace("HandleNodeCandidateListMsg", "peer", peer, "epoch", epoch)

	// try to get snapshot at epoch
	snapshot := q.snapshotSystem.GetSnapshot(epoch)
	// if it is not set - that means that we haven`t initialized this data structure before - do it
	if snapshot == nil {
		q.PrepareNodeListAtEpoch(epoch)
		if snapshot = q.snapshotSystem.GetSnapshot(epoch); snapshot == nil {
			log.Debug("Unable to get snapshot after PrepareNodeListAtEpoch", "epoch", epoch)
			return
		}
	}

	// set received data to snapshot
	snapshot.SetRemotePeerNodeCandidateList(peer.ID(), msg.ToMap())
}

// HandleMinersMsg handles MinersMsg from peer
func (q *Quorum) HandleMinersMsg(msg *MinersMsg, peer *Peer) {
	epoch := msg.Epoch.Uint64()
	log.Trace("HandleMinersMsg", "peer", peer, "epoch", epoch)

	// try to get snapshot at epoch
	snapshot := q.snapshotSystem.GetSnapshot(epoch)
	// if it is not set - that means that we haven`t initialized this data structure before - do it
	if snapshot == nil {
		log.Debug("(HandleMinersMsg) Unable to get snapshot. Skip message", "epoch", epoch)
		return
	}

	snapshot.SetRemotePeerMiners(peer.ID(), msg.Miners, msg.Hash)
	snapshot.TryApproveMinerList()
}
