// Copyright Simcord LLC
// This file is part of Bitbon System Blockchain Node.
//
// Bitbon System Blockchain Node is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// Bitbon System Blockchain Node is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with Bitbon System Blockchain Node. If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"fmt"
	"strings"

	"github.com/simcord-llc/bitbon-system-blockchain/common"
	"github.com/simcord-llc/bitbon-system-blockchain/crypto"
)

func main() {
	key := "99db6ae33e9bb611771d8df0452fffff6aaf7a30daa57fcb822b8aab37400000"

	prKey, _ := crypto.HexToECDSA(key)

	b := crypto.FromECDSAPub(&prKey.PublicKey)[1:]
	addr := common.BytesToAddress(b)
	fmt.Println(strings.ToLower(addr.Hex()))
}
