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

import "time"

// Config represents the configuration state of a quorum service.
type Config struct {
	PrepareMinersTimeout time.Duration
	PingFrequency        time.Duration
	QuorumMaxPing        time.Duration
	HandshakeTimeout     time.Duration

	Test bool
}

func (c *Config) PingProtocolTimeout() time.Duration {
	return c.HandshakeTimeout + c.QuorumMaxPing + 1*time.Second
}

// DefaultConfig represents (shocker!) the default configuration.
var DefaultConfig = Config{
	PrepareMinersTimeout: DefaultPrepareMinersTimeout,
	PingFrequency:        DefaultPingFrequency,
	QuorumMaxPing:        DefaultQuorumMaxPing,
	HandshakeTimeout:     DefaultHandshakeTimeout,
	Test:                 false,
}
