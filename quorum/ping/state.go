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

type State struct {
	init chan bool
	done chan struct{}
	err  error
}

func NewState() *State {
	state := &State{
		init: make(chan bool, 1),
		done: make(chan struct{}),
	}
	// when state is first created for a remote peer
	// it is initialized with the init
	state.init <- true

	return state
}

// returns init channel
// this channel will be closed when State is initialized
func (s *State) InitCh() <-chan bool {
	return s.init
}

// Init initializes State
func (s *State) Init() {
	close(s.init)
}

// returns done channel
// this channel will be closed when State is done
func (s *State) DoneCh() <-chan struct{} {
	return s.done
}

// Make State done
func (s *State) Done() {
	close(s.done)
}

// Set State`s error
func (s *State) SetError(err error) {
	s.err = err
}

// returns State`s error
func (s *State) Error() error {
	return s.err
}
