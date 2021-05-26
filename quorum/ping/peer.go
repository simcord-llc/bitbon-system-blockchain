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
	"errors"
	"time"

	"github.com/simcord-llc/bitbon-system-blockchain/log"

	"github.com/simcord-llc/bitbon-system-blockchain/p2p/protocols"
)

var ErrPeerClosed = errors.New("(quorum ping) peer closed")
var ErrTimeout = errors.New("(quorum ping) ping-pong time out")

type Peer struct {
	*protocols.Peer // represents the connection for online peers
	logger          log.Logger

	outCh chan bool // trigger ping/ pong to remote peer (outCh - packet to remote peer). "false" if ping, "true" if pong
	// (inCh - packet from remote peer) optional, report back to calling code. "false" if ping, "true" if pong
	inCh    chan bool
	closeCh chan struct{}
}

func NewPeer(p *protocols.Peer) *Peer {
	return &Peer{
		Peer:    p,
		outCh:   make(chan bool),
		inCh:    make(chan bool),
		closeCh: make(chan struct{}),
		logger:  log.New("type", "pingPeer", "peer", p),
	}
}

func (pp *Peer) Close() error {
	if !pp.Closed() {
		close(pp.closeCh)
	}
	return nil
}

func (pp *Peer) Closed() bool {
	select {
	case <-pp.closeCh:
		return true
	default:
		return false
	}
}

func (pp *Peer) PingPong(timeout time.Duration) (dur time.Duration, err error) {
	now := time.Now()
	timer := time.NewTimer(timeout)

	defer func() {
		timer.Stop()
		dur = time.Since(now)
	}()

	select {
	case <-timer.C:
		return 0, ErrTimeout
	case <-pp.closeCh:
		return 0, ErrPeerClosed
	case pp.outCh <- false:
	}

	select {
	case <-timer.C:
		return 0, ErrTimeout
	case <-pp.closeCh:
		return 0, ErrPeerClosed
	case <-pp.inCh:
	}

	return 0, nil
}

func (pp *Peer) sendLoop() {
	for {
		select {
		case ispong := <-pp.outCh:
			msg := &Msg{
				Created: uint64(time.Now().Unix()),
				Pong:    ispong,
			}
			err := pp.Send(context.TODO(), msg)
			pp.logger.Debug("-> sent ping message", "msg", msg, "error", err)

		case <-pp.closeCh:
			return
		}
	}
}

func (pp *Peer) pingHandler(ctx context.Context, msg interface{}) error {
	var pingmsg *Msg
	var ok bool
	if pingmsg, ok = msg.(*Msg); !ok {
		return errors.New("invalid msg")
	}

	// notify code about incoming pong request
	if pingmsg.Pong {
		select {
		case pp.inCh <- pingmsg.Pong:
		default:
		}
	}

	// if it is ping request - response to remote peer
	if !pingmsg.Pong {
		select {
		case <-pp.closeCh:
			return ErrPeerClosed
		case pp.outCh <- true:
		}
	}
	return nil
}
