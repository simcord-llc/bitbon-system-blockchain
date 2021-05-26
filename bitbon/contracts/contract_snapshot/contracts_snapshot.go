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

package contract_snapshot

import (
	"errors"
	"sync"

	"github.com/simcord-llc/bitbon-system-blockchain/common"
)

type ContractsSnapshot struct {
	mux               sync.RWMutex                    // protect contractStorageAddress, contractAddresses
	contractAddresses map[ContractType]common.Address // map of deployed contract addresses in blockchain
}

func New(contractStorageAddress common.Address, contractAddresses map[ContractType]common.Address) *ContractsSnapshot {
	contractAddresses[ContractTypeContractStorageImpl] = contractStorageAddress

	cs := &ContractsSnapshot{
		contractAddresses: contractAddresses,
	}

	return cs
}

func (cs *ContractsSnapshot) GetContractAddress(contractType ContractType) (common.Address, error) {
	cs.mux.RLock()
	defer cs.mux.RUnlock()

	if !ContractTypeExists(contractType) {
		return common.Address{}, errors.New("unknown contract type")
	}

	addr, ok := cs.contractAddresses[contractType]
	if !ok {
		return common.Address{}, errors.New("no address matches the provided contract type")
	}

	return addr, nil
}

func (cs *ContractsSnapshot) SetContractAddress(contractType ContractType, address common.Address) error {
	cs.mux.Lock()
	defer cs.mux.Unlock()

	if !ContractTypeExists(contractType) {
		return errors.New("unknown contract type")
	}

	cs.contractAddresses[contractType] = address

	return nil
}

func (cs *ContractsSnapshot) SetContractStorageAddress(address common.Address) error {
	return cs.SetContractAddress(ContractTypeContractStorageImpl, address)
}

func (cs *ContractsSnapshot) GetContractStorageAddress() (common.Address, error) {
	return cs.GetContractAddress(ContractTypeContractStorageImpl)
}

func (cs *ContractsSnapshot) GetContractAddresses() map[ContractType]common.Address {
	cs.mux.RLock()
	defer cs.mux.RUnlock()

	return cs.contractAddresses
}
