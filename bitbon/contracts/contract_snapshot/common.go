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

type ContractType string

const (
	ContractTypeAssetboxStorageImpl      ContractType = "AssetboxInfoStorageImpl"
	ContractTypeBitbonImpl               ContractType = "BitbonImpl"
	ContractTypeTransferImpl             ContractType = "TransferImpl"
	ContractTypeSafeTransferStorageImpl  ContractType = "SafeTransferStorageImpl"
	ContractTypeBitbonSupportImpl        ContractType = "BitbonSupportImpl"
	ContractTypeBitbonStorageImpl        ContractType = "BitbonStorageImpl"
	ContractTypeAdminStorageImpl         ContractType = "AdminStorageImpl"
	ContractTypeReservedAliasStorageImpl ContractType = "ReservedAliasStorageImpl"
	ContractTypeNewContractStorage       ContractType = "NewContractStorage"
	ContractTypeAssetboxImpl             ContractType = "AssetboxImpl"
	ContractTypeAccessStorageImpl        ContractType = "AccessStorageImpl"
	ContractTypeDistributionStorageImpl  ContractType = "DistributionStorageImpl"
	ContractTypeMiningAgentStorageImpl   ContractType = "MiningAgentStorageImpl"
	ContractTypeContractStorageImpl      ContractType = "ContractStorageImpl"
)

var contractTypeMap = map[ContractType]struct{}{
	ContractTypeAssetboxStorageImpl:      {},
	ContractTypeBitbonImpl:               {},
	ContractTypeTransferImpl:             {},
	ContractTypeSafeTransferStorageImpl:  {},
	ContractTypeBitbonSupportImpl:        {},
	ContractTypeBitbonStorageImpl:        {},
	ContractTypeAdminStorageImpl:         {},
	ContractTypeReservedAliasStorageImpl: {},
	ContractTypeNewContractStorage:       {},
	ContractTypeAssetboxImpl:             {},
	ContractTypeAccessStorageImpl:        {},
	ContractTypeDistributionStorageImpl:  {},
	ContractTypeMiningAgentStorageImpl:   {},
	ContractTypeContractStorageImpl:      {},
}

func ContractTypeExistsString(contractType string) bool {
	_, ok := contractTypeMap[ContractType(contractType)]

	return ok
}

func ContractTypeExists(contractType ContractType) bool {
	_, ok := contractTypeMap[contractType]

	return ok
}
