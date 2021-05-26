// Copyright Simcord LLC
// This file is part of the Bitbon System Blockchain Node library.
// This file has been modified, you can find the original version by following the link
// <https://github.com/ethereum/go-ethereum>
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

package discover

import (
	"github.com/hashicorp/golang-lru/simplelru"
	"github.com/simcord-llc/bitbon-system-blockchain/common/mclock"
	"github.com/simcord-llc/bitbon-system-blockchain/p2p/enode"
)

// The sessionCache keeps negotiated encryption keys and
// state for in-progress handshakes in the Discovery v5 wire protocol.
type sessionCache struct {
	sessions   *simplelru.LRU
	handshakes map[sessionID]*whoareyouV5
	clock      mclock.Clock
}

// sessionID identifies a session or handshake.
type sessionID struct {
	id   enode.ID
	addr string
}

// session contains session information
type session struct {
	writeKey []byte
	readKey  []byte
}

func newSessionCache(maxItems int, clock mclock.Clock) *sessionCache {
	cache, err := simplelru.NewLRU(maxItems, nil)
	if err != nil {
		panic("can't create session cache")
	}
	return &sessionCache{
		sessions:   cache,
		handshakes: make(map[sessionID]*whoareyouV5),
		clock:      clock,
	}
}

// session returns the current session for the given node, if any.
func (sc *sessionCache) session(id enode.ID, addr string) *session {
	item, ok := sc.sessions.Get(sessionID{id, addr})
	if !ok {
		return nil
	}
	return item.(*session)
}

// readKey returns the current read key for the given node.
func (sc *sessionCache) readKey(id enode.ID, addr string) []byte {
	if s := sc.session(id, addr); s != nil {
		return s.readKey
	}
	return nil
}

// writeKey returns the current read key for the given node.
func (sc *sessionCache) writeKey(id enode.ID, addr string) []byte {
	if s := sc.session(id, addr); s != nil {
		return s.writeKey
	}
	return nil
}

// storeNewSession stores new encryption keys in the cache.
func (sc *sessionCache) storeNewSession(id enode.ID, addr string, r, w []byte) {
	sc.sessions.Add(sessionID{id, addr}, &session{
		readKey: r, writeKey: w,
	})
}

// getHandshake gets the handshake challenge we previously sent to the given remote node.
func (sc *sessionCache) getHandshake(id enode.ID, addr string) *whoareyouV5 {
	return sc.handshakes[sessionID{id, addr}]
}

// storeSentHandshake stores the handshake challenge sent to the given remote node.
func (sc *sessionCache) storeSentHandshake(id enode.ID, addr string, challenge *whoareyouV5) {
	challenge.sent = sc.clock.Now()
	sc.handshakes[sessionID{id, addr}] = challenge
}

// deleteHandshake deletes handshake data for the given node.
func (sc *sessionCache) deleteHandshake(id enode.ID, addr string) {
	delete(sc.handshakes, sessionID{id, addr})
}

// handshakeGC deletes timed-out handshakes.
func (sc *sessionCache) handshakeGC() {
	deadline := sc.clock.Now().Add(-handshakeTimeout)
	for key, challenge := range sc.handshakes {
		if challenge.sent < deadline {
			delete(sc.handshakes, key)
		}
	}
}
