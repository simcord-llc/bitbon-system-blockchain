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

	"github.com/simcord-llc/bitbon-system-blockchain/log"
)

// sendNodeCandidateList send NodeCandidateListMsg to all given peers
func (q *Quorum) sendNodeCandidateList(peers []*Peer, msg *NodeCandidateListMsg) {
	log.Trace("sendNodeCandidateList")

	for _, peer := range peers {
		if err := peer.SendNodeCandidateList(context.TODO(), msg); err != nil {
			q.logger.Error("failed to notify NodeCandidateListMsg", "peer", peer, "msg", msg, "error", err)
		}
	}
}

// notifyVoteMinerAdd notifies all given peers that self node wants to be a miner
func (q *Quorum) notifyVoteMinerAdd(peers []*Peer) {
	log.Trace("notifyVoteMinerAdd")

	for _, peer := range peers {
		if err := peer.SendVoteMinerAddMsg(context.TODO()); err != nil {
			q.logger.Error("failed to notify VoteMinerAdd", "peer", peer, "error", err)
		}
	}
}

// notifyVoteMinerDel notifies all given peers that self node DOES NOT want to be a miner
func (q *Quorum) notifyVoteMinerDel(peers []*Peer) {
	log.Trace("notifyVoteMinerDel")

	for _, peer := range peers {
		if err := peer.SendVoteMinerDelMsg(context.TODO()); err != nil {
			q.logger.Error("failed to notify VoteMinerDel", "peer", peer, "error", err)
		}
	}
}

// sendMiners send MinersHashMsg to all given peers
func (q *Quorum) sendMiners(peers []*Peer, msg *MinersMsg) {
	log.Trace("sendMiners")

	for _, peer := range peers {
		if err := peer.SendMiners(context.TODO(), msg); err != nil {
			q.logger.Error("failed to send MinersMsg", "peer", peer, "msg", msg, "error", err)
		}
	}
}
