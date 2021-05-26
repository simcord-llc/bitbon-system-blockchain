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
	"fmt"
	"math/big"
	"time"

	"github.com/simcord-llc/bitbon-system-blockchain/p2p/protocols"
)

// quorum constants
const (
	ProtocolVersionStr = "1.0.0"  // The same, as a string
	ProtocolName       = "quorum" // Nickname of the protocol in geth

	DefaultPrepareMinersTimeout = 10 * time.Second
	DefaultPingFrequency        = 10 * time.Second
	DefaultQuorumMaxPing        = 1200 * time.Millisecond
	DefaultHandshakeTimeout     = 3000 * time.Millisecond

	coefficientPrecision = 5
)

var (
	// 2/3
	majorityCoefficient = new(big.Float).Quo(big.NewFloat(2), big.NewFloat(3)).SetPrec(coefficientPrecision)
)

func SpecString(spec *protocols.Spec) string {
	return fmt.Sprintf("%s-v.%d", spec.Name, spec.Version)
}
