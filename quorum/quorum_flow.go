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
	"time"
)

func (q *Quorum) wait(period time.Duration) bool {
	timer := time.NewTimer(period)
	defer timer.Stop()

	select {
	case <-timer.C:
		return true
	case <-q.done:
		return false
	}
}

// waitForPeers method waits for majority of ping list gathered.
// Than it marks quorum service that peers are gathered and terminates
func (q *Quorum) waitForPeers() {
	checkPeriod := 1 * time.Second
	for {
		// wait before each iteration
		if !q.wait(checkPeriod) {
			return
		}

		// check majority of ping list gathered. If it has already gathered - return.
		// If not - check and mark (if majority exists)
		if q.hasMajority(len(q.peerList.PingList()), true) {
			q.markPeersGathered(true)

			// terminate function when peers connected
			if q.SelfIsMinerCandidate() {
				q.notifyVoteMinerAdd(q.peerList.Peers())
			}
			return
		}
	}
}
