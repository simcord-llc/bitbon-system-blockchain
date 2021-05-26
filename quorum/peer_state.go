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
	"sync/atomic"
	"time"
)

type peerState struct {
	mu sync.RWMutex // Mutex protects all fields

	lastPing *int64 // last ping duration in nanoseconds

	inPingList       bool // indicator if this peer in our ping list (ping duration satisfies our limit)
	isMinerCandidate bool // indicator if this peer wants to mine AT THE MOMENT (not depending to epoch)
}

type PeerStateDTO struct {
	LastPing         time.Duration
	LastPingStr      string
	InPingList       bool
	IsMinerCandidate bool
	PeerRate         int
}

func (p *peerState) ToDTO() *PeerStateDTO {
	p.mu.RLock()
	defer p.mu.RUnlock()

	ping := time.Duration(atomic.LoadInt64(p.lastPing))
	return &PeerStateDTO{
		LastPing:         ping,
		LastPingStr:      ping.String(),
		InPingList:       p.inPingList,
		IsMinerCandidate: p.isMinerCandidate,
	}
}

// ---------------------- Data methods

// markInPingList marks peerState inPingList filed
func (p *peerState) markInPingList(res bool) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.inPingList = res
}

// InPingList returns if this peerState in our ping list
func (p *peerState) InPingList() bool {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.inPingList
}

// MarkIsMinerCandidate marks peer as miner candidate AT THE MOMENT
func (p *peerState) MarkIsMinerCandidate(res bool) {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.isMinerCandidate = res
}

// IsMinerCandidate returns if this peer wants to mine AT THE MOMENT (not depending to epoch)
func (p *peerState) IsMinerCandidate() bool {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return p.isMinerCandidate
}

// IsQuorumCandidate returns if this peerState is our quorum candidate now
func (p *peerState) IsQuorumCandidate() bool {
	return p.InPingList()
}

func (p *Peer) setLastPing(d time.Duration) {
	atomic.StoreInt64(p.lastPing, d.Nanoseconds())
}
