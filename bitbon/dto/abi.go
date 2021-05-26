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

package dto

import (
	"strings"

	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/contracts/contract_snapshot"
	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/storage/models"
	"github.com/simcord-llc/bitbon-system-blockchain/common"
)

type AbiInfo struct {
	Address      common.Address                 `json:"address"`
	AbiJSON      string                         `json:"abiJson"`
	Version      ContractVersion                `json:"version"`
	ContractType contract_snapshot.ContractType `json:"contractType"`
}

func (a *AbiInfo) ToStorageAbiInfo() *models.AbiInfo {
	ai := &models.AbiInfo{
		Address:      strings.ToLower(a.Address.Hex()),
		Abi:          a.AbiJSON,
		Version:      string(a.Version),
		ContractType: string(a.ContractType),
	}
	return ai
}

func AbiInfoFromStorage(sa *models.AbiInfo) *AbiInfo {
	a := &AbiInfo{
		Address:      common.HexToAddress(sa.Address),
		AbiJSON:      sa.Abi,
		Version:      ContractVersion(sa.Version),
		ContractType: contract_snapshot.ContractType(sa.ContractType),
	}
	return a
}

func AbiInfoAddressToStorage(addr common.Address) string {
	return strings.ToLower(addr.Hex())
}

func StorageToAbiInfoAddress(addr string) common.Address {
	return common.HexToAddress(addr)
}
