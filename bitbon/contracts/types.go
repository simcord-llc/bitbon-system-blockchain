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

package contracts

import (
	"math/big"

	"github.com/simcord-llc/bitbon-system-blockchain/common"
)

type Assetbox struct {
	Address   common.Address
	Alias     string
	ServiceID []byte
	IsPublic  bool
	ExtraInfo string
	IsMining  bool
}

type Balance struct {
	Address      common.Address
	Balance      *big.Int
	AssetboxType *big.Int
}

type Transfer struct {
	To         common.Address
	Value      *big.Int
	Proof      []byte
	VK         []byte
	TransferID []byte
	Retries    *big.Int
	ExpiresAt  *big.Int
	ExtraData  []byte
}

type TransferExpired struct {
	TransferID []byte
}

type ReceiptTransfer struct {
	State          uint8
	TransferIdHash [32]byte
	Source         common.Address
	Dest           common.Address
	Amount         *big.Int
	Proof          []byte
	Vk             []byte
	CreatedAt      *big.Int
	ExpiresAt      *big.Int
	CurrentRetries *big.Int
	MaxRetries     *big.Int
	Canceller      common.Address
}
