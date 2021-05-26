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

package ping

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/simcord-llc/bitbon-system-blockchain/p2p/enode"

	"github.com/simcord-llc/bitbon-system-blockchain/log"
	"github.com/simcord-llc/bitbon-system-blockchain/p2p"
	"github.com/simcord-llc/bitbon-system-blockchain/p2p/protocols"
)

type Msg struct {
	Created uint64
	Pong    bool // set if message is pong reply
}

type HandshakeMsg struct {
	Version uint64
}

// Spec is the spec of the ping protocol
var Spec = &protocols.Spec{
	Name:       "qrm-png",
	Version:    1,
	MaxMsgSize: 10 * 1024 * 1024,
	Messages: []interface{}{
		HandshakeMsg{},
		Msg{},
	},
}

func SpecString(spec *protocols.Spec) string {
	return fmt.Sprintf("%s-v.%d", spec.Name, spec.Version)
}

type Protocol struct {
	stateMu sync.Mutex
	states  map[enode.ID]*State

	logger log.Logger

	peerMu sync.Mutex
	peers  map[enode.ID]*Peer

	handshakeTimeout time.Duration
}

// NewProtocol is the Ping protocol constructor
func NewProtocol(handshakeTimeout time.Duration) *Protocol {
	return &Protocol{
		states: make(map[enode.ID]*State),
		peers:  make(map[enode.ID]*Peer),

		logger:           log.New("protocol", SpecString(Spec)),
		handshakeTimeout: handshakeTimeout,
	}
}

func (png *Protocol) Protocols() []p2p.Protocol {
	return []p2p.Protocol{
		{
			Name:    Spec.Name,
			Version: Spec.Version,
			Length:  Spec.Length(),
			Run:     png.runProtocol,
		},
	}
}

// runProtocol is the p2p protocol run function for the quorum base protocol
// that negotiates the quorum handshake and ping
func (png *Protocol) runProtocol(p *p2p.Peer, rw p2p.MsgReadWriter) (err error) {
	state, _ := png.GetState(p.ID())
	if !<-state.InitCh() {
		return fmt.Errorf("%s: handshake already started on peer %s", SpecString(Spec), p.ID().String())
	}
	state.Init()

	defer png.RemoveState(p.ID())

	pp := NewPeer(protocols.NewPeer(p, rw, Spec))
	defer pp.Close()

	png.addPeer(pp)
	defer png.removePeer(pp.ID())

	err = png.start(pp, state)
	if err != nil {
		return err
	}

	go pp.sendLoop()

	return pp.Run(pp.pingHandler)
}

// start logic of ping protocol
func (png *Protocol) start(pp *Peer, state *State) error {
	defer state.Done()

	err := png.performHandshake(pp)
	if err != nil {
		err = fmt.Errorf("%s: handshake failed with remote peer %s: %v", SpecString(Spec), pp.ID().String(), err)
		log.Warn(err.Error())
		state.SetError(err)
		return err
	}
	return nil
}

// removeState removes sate for peer with peerID
// from the sate store
func (png *Protocol) RemoveState(peerID enode.ID) {
	png.stateMu.Lock()
	defer png.stateMu.Unlock()
	delete(png.states, peerID)
}

// getState returns the state with the remote peer with peerID
func (png *Protocol) GetState(peerID enode.ID) (*State, bool) {
	png.stateMu.Lock()
	defer png.stateMu.Unlock()
	state, found := png.states[peerID]
	if !found {
		state = NewState()
		png.states[peerID] = state
	}

	return state, found
}

// removePeer removes peer with peerID
// from the peer store
func (png *Protocol) removePeer(peerID enode.ID) {
	png.peerMu.Lock()
	defer png.peerMu.Unlock()
	delete(png.peers, peerID)
}

// addPeer adds the peer with peerID
func (png *Protocol) addPeer(p *Peer) {
	png.peerMu.Lock()
	defer png.peerMu.Unlock()
	png.peers[p.ID()] = p
}

// GetPeer returns the peer with peerID
func (png *Protocol) GetPeer(peerID enode.ID) *Peer {
	png.peerMu.Lock()
	defer png.peerMu.Unlock()
	return png.peers[peerID]
}

// performHandshake implements the negotiation of the quorum handshake
// shared among quorum subprotocols
func (png *Protocol) performHandshake(p *Peer) error {
	ctx, cancel := context.WithTimeout(context.Background(), png.handshakeTimeout)
	defer cancel()
	_, err := p.Handshake(ctx, &HandshakeMsg{Version: uint64(Spec.Version)}, png.checkHandshake)
	return err
}

// Perform initiates the handshake and validates the remote handshake message
func (png *Protocol) checkHandshake(hs interface{}) error {
	rhs, ok := hs.(*HandshakeMsg)
	if !ok {
		return fmt.Errorf("wanted (*HandshakeMsg) got (%T)", hs)
	}
	if rhs.Version != uint64(Spec.Version) {
		return fmt.Errorf("version mismatch %d (!= %d)", rhs.Version, Spec.Version)
	}
	return nil
}
