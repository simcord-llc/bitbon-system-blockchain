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

package interfaces

import (
	"context"

	"github.com/simcord-llc/bitbon-system-blockchain/accounts/abi/bind"
	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/contracts/contract_snapshot"
	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/dto"
	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/storage"
	"github.com/simcord-llc/bitbon-system-blockchain/common"
	"github.com/simcord-llc/bitbon-system-blockchain/core/types"
)

type MethodParser interface {
	Parse(ctx context.Context, tx *types.Transaction) (*dto.BitbonTx, error)
}

type BitbonStorage interface {
	SaveAbiInfo(tx storage.Transaction, abiInfo *dto.AbiInfo) (err error)
	EditAbiInfo(tx storage.Transaction, abiInfo *dto.AbiInfo) (err error)
	DeleteAbiInfoByAddress(tx storage.Transaction, address common.Address) (err error)
	GetAbiJSON(tx storage.Transaction, address common.Address) (abiJSON string, err error)
	GetAbiVersion(tx storage.Transaction, address common.Address) (version dto.ContractVersion, err error)
	GetFullAbiInfo(tx storage.Transaction, address common.Address) (abiInfo *dto.AbiInfo, err error)
	GetFullAbiInfoByVersionAndType(tx storage.Transaction, version dto.ContractVersion,
		contractType contract_snapshot.ContractType) (abiInfo *dto.AbiInfo, err error)
	GetAbiAddressesByVersion(tx storage.Transaction, version dto.ContractVersion) (addresses []common.Address, err error)
	GetAbiInfosByVersion(tx storage.Transaction, version dto.ContractVersion) (abiInfos []*dto.AbiInfo, err error)
}

type Contract interface {
	bind.ContractCaller
	bind.ContractTransactor
	bind.ContractFilterer
	TransactionReceipt(context.Context, common.Hash) (*types.Receipt, error)
}
