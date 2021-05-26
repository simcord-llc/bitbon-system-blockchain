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
	"context"
	"time"

	"github.com/pkg/errors"
	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/contracts/contract_snapshot"
	"github.com/simcord-llc/bitbon-system-blockchain/common"

	"github.com/simcord-llc/bitbon-system-blockchain/params"
)

const (
	defaultGasPrice = 5 * params.Wei
	defaultGasLimit = 10000000
	// timeout to wait for transaction receipt
	waitForTxReceiptTimeout = 20 * time.Second
	// amount of operations we give assetbox by one ether emission
	operationsPerEmitEther = 10000
)

type GetterFunc func(ctx context.Context, address common.Address) (addr common.Address, err error)

var addressGetters = map[contract_snapshot.ContractType]func(m *Manager) GetterFunc{

	contract_snapshot.ContractTypeAssetboxStorageImpl: func(m *Manager) GetterFunc {
		return m.GetAssetboxInfoStorageContractAddress
	},
	contract_snapshot.ContractTypeBitbonImpl: func(m *Manager) GetterFunc {
		return m.GetBitbonContractAddress
	},
	contract_snapshot.ContractTypeTransferImpl: func(m *Manager) GetterFunc {
		return m.GetTransferContractAddress
	},
	contract_snapshot.ContractTypeSafeTransferStorageImpl: func(m *Manager) GetterFunc {
		return m.GetSafeTransferStorageContractAddress
	},
	contract_snapshot.ContractTypeBitbonSupportImpl: func(m *Manager) GetterFunc {
		return m.GetBitbonSupportContractAddress
	},
	contract_snapshot.ContractTypeBitbonStorageImpl: func(m *Manager) GetterFunc {
		return m.GetBitbonStorageContractAddress
	},
	contract_snapshot.ContractTypeReservedAliasStorageImpl: func(m *Manager) GetterFunc {
		return m.GetReservedAliasStorageContractAddress
	},
	contract_snapshot.ContractTypeAssetboxImpl: func(m *Manager) GetterFunc {
		return m.GetAssetboxContractAddress
	},
	contract_snapshot.ContractTypeAccessStorageImpl: func(m *Manager) GetterFunc {
		return m.GetAccessStorageContractAddress
	},
	contract_snapshot.ContractTypeDistributionStorageImpl: func(m *Manager) GetterFunc {
		return m.GetDistributionStorageContractAddress
	},
	contract_snapshot.ContractTypeMiningAgentStorageImpl: func(m *Manager) GetterFunc {
		return m.GetMiningAgentStorageContractAddress
	},
}

var contractTypeGenAbi = map[contract_snapshot.ContractType]string{
	contract_snapshot.ContractTypeAssetboxStorageImpl:      AssetboxStorageImplABI,
	contract_snapshot.ContractTypeBitbonImpl:               BitbonImplABI,
	contract_snapshot.ContractTypeTransferImpl:             TransferImplABI,
	contract_snapshot.ContractTypeSafeTransferStorageImpl:  SafeTransferStorageImplABI,
	contract_snapshot.ContractTypeBitbonSupportImpl:        BitbonSupportImplABI,
	contract_snapshot.ContractTypeBitbonStorageImpl:        BitbonStorageImplABI,
	contract_snapshot.ContractTypeReservedAliasStorageImpl: ReservedAliasStorageImplABI,
	contract_snapshot.ContractTypeAssetboxImpl:             AssetboxImplABI,
	contract_snapshot.ContractTypeAccessStorageImpl:        AccessStorageImplABI,
	contract_snapshot.ContractTypeDistributionStorageImpl:  DistributionStorageImplABI,
	contract_snapshot.ContractTypeMiningAgentStorageImpl:   MiningAgentStorageImplABI,
	contract_snapshot.ContractTypeContractStorageImpl:      ContractStorageImplABI,
}

func ContractTypeAbi(contractType contract_snapshot.ContractType) (string, error) {
	a, ok := contractTypeGenAbi[contractType]
	if !ok {
		return "", errors.Errorf("no abi for contract type %s", contractType)
	}

	return a, nil
}
