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

package noncer

const (
	ProtocolVersionStr = "1.0.0"         // The same, as a string
	ProtocolName       = "bitbon_noncer" // Nickname of the protocol in geth

	// Default parameters for redis cluster connect configuration

	DefaultDialTimeout = 5000
	DefaultMaxRetries  = 20
)

func IsAllowed() bool {
	return true
}
