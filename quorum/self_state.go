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
	"sync"
)

type SelfState struct {
	mu               sync.RWMutex
	isMinerCandidate bool // indicator if self node wants to be a miner AT THE MOMENT (not depending to epoch)
}

func NewSelfState() *SelfState {
	return &SelfState{}
}

// IsMinerCandidate - returns if self node is miner candidate AT THE MOMENT (not depending do epoch)
func (ss *SelfState) IsMinerCandidate() bool {
	ss.mu.RLock()
	defer ss.mu.RUnlock()

	return ss.isMinerCandidate
}

// MarkIsMinerCandidate - marks self node is miner candidate AT THE MOMENT
func (ss *SelfState) MarkIsMinerCandidate(in bool) {
	ss.mu.Lock()
	defer ss.mu.Unlock()

	ss.isMinerCandidate = in
}
