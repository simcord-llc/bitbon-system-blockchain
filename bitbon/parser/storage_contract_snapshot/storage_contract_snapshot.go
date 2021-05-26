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

package storage_contract_snapshot // nolint

import (
	"strings"

	"github.com/pkg/errors"
	"github.com/simcord-llc/bitbon-system-blockchain/accounts/abi"
	cs "github.com/simcord-llc/bitbon-system-blockchain/bitbon/contracts/contract_snapshot"
	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/dto"
	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/parser/interfaces"
	"github.com/simcord-llc/bitbon-system-blockchain/common"
)

type StorageContractSnapshot struct {
	*cs.ContractsSnapshot

	abis    map[cs.ContractType]abi.ABI
	Version dto.ContractVersion
	Storage interfaces.BitbonStorage
}

func NewStorageSnapshotByVersion(storage interfaces.BitbonStorage,
	version dto.ContractVersion) (*StorageContractSnapshot, error) {
	abiInfos, err := storage.GetAbiInfosByVersion(nil, version)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get abi infos for snapshot from storage")
	}

	if len(abiInfos) == 0 {
		return nil, errors.Errorf("no contracts for version %v", version)
	}

	contractAddresses := make(map[cs.ContractType]common.Address, len(abiInfos))
	abis := make(map[cs.ContractType]abi.ABI, len(abiInfos))
	for _, ai := range abiInfos {
		contractAddresses[ai.ContractType] = ai.Address
		abiJSON := ai.AbiJSON
		if abiJSON == "" {
			return nil, errors.Errorf("loaded empty abi for contract type %s of version %s", string(ai.ContractType), version)
		}

		a, err := abi.JSON(strings.NewReader(abiJSON))
		if err != nil {
			return nil, errors.Wrap(err, "unable to create ABI object")
		}

		abis[ai.ContractType] = a
	}

	contractStorageAddress, ok := contractAddresses[cs.ContractTypeContractStorageImpl]
	if !ok {
		return nil, errors.New("unable to create snapshot from storage without contract storage address")
	}

	ss := &StorageContractSnapshot{
		ContractsSnapshot: cs.New(contractStorageAddress, contractAddresses),
		abis:              abis,
		Version:           version,
		Storage:           storage,
	}
	return ss, nil
}

func NewStorageSnapshotByContractAddress(storage interfaces.BitbonStorage,
	address common.Address) (*StorageContractSnapshot, error) {
	fullAbiInfo, err := storage.GetFullAbiInfo(nil, address)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get full abi info")
	}

	return NewStorageSnapshotByVersion(storage, fullAbiInfo.Version)
}

func (scs *StorageContractSnapshot) GetAbi(contractType cs.ContractType) (abi.ABI, error) {
	a, ok := scs.abis[contractType]
	if !ok {
		return abi.ABI{}, errors.Errorf("unable to find ABI for version %s and contract type %s", scs.Version, contractType)
	}

	return a, nil
}
