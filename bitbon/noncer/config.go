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

// Config represents the configuration state of a bitbon agent service.
type Config struct {
	Enabled       bool
	RedisConnStrs []string
	Key           string
	DialTimeout   uint64
	MaxRetries    uint64
}