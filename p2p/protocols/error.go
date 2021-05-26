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

package protocols

// HandlerError wraps standard error
// This error is handled specially by protocol.Run
// It causes the protocol to return with ErrHandler(err)
type breakError struct {
	err error
}

// Break wraps error and creates a special error that is treated specially in the protocol.Run event loop
// It causes protocol.Run event loop to be exit and drop the peer.
func Break(err error) error {
	return &breakError{
		err: err,
	}
}

// Unwrap returns an underlying error
func (e *breakError) Unwrap() error { return e.err }

// Error implements function of the standard go error interface
func (e *breakError) Error() string {
	return e.err.Error()
}
