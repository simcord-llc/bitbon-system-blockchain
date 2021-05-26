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

package downloader

import "fmt"

// SyncMode represents the synchronisation mode of the downloader.
type SyncMode int

const (
	FullSync  SyncMode = iota // Synchronise the entire blockchain history from full blocks
	FastSync                  // Quickly download the headers, full sync only at the chain head
	LightSync                 // Download only the headers and terminate afterwards
)

const (
	lightVal = "light"
	fastVal  = "fast"
	fullVal  = "full"
)

func (mode SyncMode) IsValid() bool {
	return mode >= FullSync && mode <= LightSync
}

// String implements the stringer interface.
func (mode SyncMode) String() string {
	switch mode {
	case FullSync:
		return fullVal
	case FastSync:
		return fastVal
	case LightSync:
		return lightVal
	default:
		return "unknown"
	}
}

func (mode SyncMode) MarshalText() ([]byte, error) {
	switch mode {
	case FullSync:
		return []byte(fullVal), nil
	case FastSync:
		return []byte(fastVal), nil
	case LightSync:
		return []byte(lightVal), nil
	default:
		return nil, fmt.Errorf("unknown sync mode %d", mode)
	}
}

func (mode *SyncMode) UnmarshalText(text []byte) error {
	switch string(text) {
	case fullVal:
		*mode = FullSync
	case fastVal:
		*mode = FastSync
	case lightVal:
		*mode = LightSync
	default:
		return fmt.Errorf(`unknown sync mode %q, want "full", "fast" or "light"`, text)
	}
	return nil
}
