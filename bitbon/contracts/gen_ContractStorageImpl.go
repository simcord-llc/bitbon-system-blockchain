// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"math/big"
	"strings"

	ethereum "github.com/simcord-llc/bitbon-system-blockchain"
	"github.com/simcord-llc/bitbon-system-blockchain/accounts/abi"
	"github.com/simcord-llc/bitbon-system-blockchain/accounts/abi/bind"
	"github.com/simcord-llc/bitbon-system-blockchain/common"
	"github.com/simcord-llc/bitbon-system-blockchain/core/types"
	"github.com/simcord-llc/bitbon-system-blockchain/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ContractStorageImplABI is the input ABI used to generate the binding from.
const ContractStorageImplABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"contractIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"ContractDeployed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"ContractStorageMoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"NewContractStorageDeployed\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ACCESS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ACCESS_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ADMINABLE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ADMIN_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ASSETBOX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ASSETBOX_INFO\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ASSETBOX_INFO_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ASSETBOX_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_BALANCE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_BITBON\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_BITBON_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_BITBON_SUPPORT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_BONBALANCE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_BONTRANSFER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_DEX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_DISTRIBUTION_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_EXCHANGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_EXCHANGE_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_FEE_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_GENERATOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MINING_AGENT_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MSBON\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MSBON_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_ADD_ADDRESS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_ADD_ADMIN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_ADD_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_ADD_ROLE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_DROP_PERMISSION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_EDIT_ADMIN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_EDIT_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_EDIT_ROLE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REMOVE_ADMIN_BY_ADDRESS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REMOVE_ADMIN_BY_INDEX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REMOVE_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REMOVE_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REMOVE_ROLE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REPLENISH_CAPITALIZATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REPLENISH_CERTIFICATE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REPLENISH_EMISSION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_SET_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_SET_PERMISSION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_OTC\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_OTC_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_REPLENISH_CAPITALIZATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_RESERVED_ALIAS_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ROLE_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_SAFE_TRANSFER_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_TRANSFER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_TRANSFER_FROM_CAPITALIZATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_ACCESS_RESTORATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_ADMIN_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_CONTRACTS_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_CONTRACT_DEPLOY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_EMISSION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_FEES_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_MULTISIG_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_ROLES_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_ACCESS_VERIFIER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_BITBON_ISSUE_VERIFIER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_COMMISSION_VERIFIER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_DEPLOY_ADMIN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_PERMISSION_ADMIN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"contractDeployers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"contracts\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"contractsTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"newContactStorage\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"previousContactStorage\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"contractTypes\",\"type\":\"uint256[]\"}],\"name\":\"getContractAddresses\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"contractType\",\"type\":\"uint256\"}],\"name\":\"getContractAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"contractType\",\"type\":\"uint256\"}],\"name\":\"getContractDeployTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"contractType\",\"type\":\"uint256\"}],\"name\":\"getContractDeployer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getPreviousContractStorageAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setPreviousContractStorageAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"contractType\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setContractAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getMultisigStorage\",\"outputs\":[{\"internalType\":\"contractMultisigStorage\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setMultisigStorage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getFeeStorage\",\"outputs\":[{\"internalType\":\"contractFeeStorage\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setFeeStorage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getBitbon\",\"outputs\":[{\"internalType\":\"contractBitbon\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setBitbon\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getBitbonStorage\",\"outputs\":[{\"internalType\":\"contractBitbonStorage\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setBitbonStorage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getMsbonStorage\",\"outputs\":[{\"internalType\":\"contractMsbonStorage\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setMsbonStorage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getBitbonSupport\",\"outputs\":[{\"internalType\":\"contractBitbonSupport\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setBitbonSupport\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getThisContractIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAssetbox\",\"outputs\":[{\"internalType\":\"contractAssetbox\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setAssetbox\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAssetboxInfo\",\"outputs\":[{\"internalType\":\"contractAssetboxInfo\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setAssetboxInfo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAssetboxStorage\",\"outputs\":[{\"internalType\":\"contractAssetboxStorage\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setAssetboxStorage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getReservedAliasStorage\",\"outputs\":[{\"internalType\":\"contractReservedAliasStorage\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setReservedAliasStorage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getSafeTransferStorage\",\"outputs\":[{\"internalType\":\"contractSafeTransferStorage\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setSafeTransferStorage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getBalanceGate\",\"outputs\":[{\"internalType\":\"contractBalance\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setBalanceGate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTransferGate\",\"outputs\":[{\"internalType\":\"contractTransfer\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setTransferGate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getExchange\",\"outputs\":[{\"internalType\":\"contractExchange\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setExchange\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getExchangeStorage\",\"outputs\":[{\"internalType\":\"contractExchangeStorage\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setExchangeStorage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOtc\",\"outputs\":[{\"internalType\":\"contractOtc\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setOtc\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOtcStorage\",\"outputs\":[{\"internalType\":\"contractOtcStorage\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setOtcStorage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAccess\",\"outputs\":[{\"internalType\":\"contractAccess\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setAccess\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAccessStorage\",\"outputs\":[{\"internalType\":\"contractAccessStorage\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setAccessStorage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newContractStorageAddress\",\"type\":\"address\"}],\"name\":\"notifyRedeploy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"notifyFirstDeploy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNewContactStorageAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAdminable\",\"outputs\":[{\"internalType\":\"contractAdminable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"adminableContractAddress\",\"type\":\"address\"}],\"name\":\"setAdminable\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getDistributionStorage\",\"outputs\":[{\"internalType\":\"contractDistributionStorage\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"distributionStorageContractAddress\",\"type\":\"address\"}],\"name\":\"setDistributionStorage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getMiningAgentStorage\",\"outputs\":[{\"internalType\":\"contractMiningAgentStorage\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"miningAgentStorageContractAddress\",\"type\":\"address\"}],\"name\":\"setMiningAgentStorage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getMsbon\",\"outputs\":[{\"internalType\":\"contractMsbon\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setMsbon\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getBonBalanceGate\",\"outputs\":[{\"internalType\":\"contractBonBalance\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setBonBalanceGate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getBonTransferGate\",\"outputs\":[{\"internalType\":\"contractBonTransfer\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setBonTransferGate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ContractStorageImpl is an auto generated Go binding around an Ethereum contract.
type ContractStorageImpl struct {
	ContractStorageImplCaller     // Read-only binding to the contract
	ContractStorageImplTransactor // Write-only binding to the contract
	ContractStorageImplFilterer   // Log filterer for contract events
}

// ContractStorageImplCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractStorageImplCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractStorageImplTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractStorageImplTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractStorageImplFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractStorageImplFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractStorageImplSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractStorageImplSession struct {
	Contract     *ContractStorageImpl // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ContractStorageImplCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractStorageImplCallerSession struct {
	Contract *ContractStorageImplCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// ContractStorageImplTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractStorageImplTransactorSession struct {
	Contract     *ContractStorageImplTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// ContractStorageImplRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractStorageImplRaw struct {
	Contract *ContractStorageImpl // Generic contract binding to access the raw methods on
}

// ContractStorageImplCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractStorageImplCallerRaw struct {
	Contract *ContractStorageImplCaller // Generic read-only contract binding to access the raw methods on
}

// ContractStorageImplTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractStorageImplTransactorRaw struct {
	Contract *ContractStorageImplTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractStorageImpl creates a new instance of ContractStorageImpl, bound to a specific deployed contract.
func NewContractStorageImpl(address common.Address, backend bind.ContractBackend) (*ContractStorageImpl, error) {
	contract, err := bindContractStorageImpl(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractStorageImpl{ContractStorageImplCaller: ContractStorageImplCaller{contract: contract}, ContractStorageImplTransactor: ContractStorageImplTransactor{contract: contract}, ContractStorageImplFilterer: ContractStorageImplFilterer{contract: contract}}, nil
}

// NewContractStorageImplCaller creates a new read-only instance of ContractStorageImpl, bound to a specific deployed contract.
func NewContractStorageImplCaller(address common.Address, caller bind.ContractCaller) (*ContractStorageImplCaller, error) {
	contract, err := bindContractStorageImpl(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractStorageImplCaller{contract: contract}, nil
}

// NewContractStorageImplTransactor creates a new write-only instance of ContractStorageImpl, bound to a specific deployed contract.
func NewContractStorageImplTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractStorageImplTransactor, error) {
	contract, err := bindContractStorageImpl(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractStorageImplTransactor{contract: contract}, nil
}

// NewContractStorageImplFilterer creates a new log filterer instance of ContractStorageImpl, bound to a specific deployed contract.
func NewContractStorageImplFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractStorageImplFilterer, error) {
	contract, err := bindContractStorageImpl(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractStorageImplFilterer{contract: contract}, nil
}

// bindContractStorageImpl binds a generic wrapper to an already deployed contract.
func bindContractStorageImpl(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractStorageImplABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractStorageImpl *ContractStorageImplRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ContractStorageImpl.Contract.ContractStorageImplCaller.contract.Call(opts, result, method, params...)
}

// Return Abi for ContractStorageImplCaller
func (_ContractStorageImpl *ContractStorageImplCaller) ABI() abi.ABI {
	return _ContractStorageImpl.contract.Abi
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractStorageImpl *ContractStorageImplRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractStorageImpl.Contract.ContractStorageImplTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractStorageImpl *ContractStorageImplRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractStorageImpl.Contract.ContractStorageImplTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractStorageImpl *ContractStorageImplCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ContractStorageImpl.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractStorageImpl *ContractStorageImplTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractStorageImpl.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractStorageImpl *ContractStorageImplTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractStorageImpl.Contract.contract.Transact(opts, method, params...)
}

// CONTRACTACCESS is a free data retrieval call binding the contract method 0x0ac62b0c.
//
// Solidity: function CONTRACT_ACCESS() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCaller) CONTRACTACCESS(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ContractStorageImpl.contract.Call(opts, out, "CONTRACT_ACCESS")
	return *ret0, err
}

// CONTRACTACCESS is a free data retrieval call binding the contract method 0x0ac62b0c.
//
// Solidity: function CONTRACT_ACCESS() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplSession) CONTRACTACCESS() (*big.Int, error) {
	return _ContractStorageImpl.Contract.CONTRACTACCESS(&_ContractStorageImpl.CallOpts)
}

// CONTRACTACCESS is a free data retrieval call binding the contract method 0x0ac62b0c.
//
// Solidity: function CONTRACT_ACCESS() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCallerSession) CONTRACTACCESS() (*big.Int, error) {
	return _ContractStorageImpl.Contract.CONTRACTACCESS(&_ContractStorageImpl.CallOpts)
}

// CONTRACTACCESSSTORAGE is a free data retrieval call binding the contract method 0xceac9bd5.
//
// Solidity: function CONTRACT_ACCESS_STORAGE() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCaller) CONTRACTACCESSSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ContractStorageImpl.contract.Call(opts, out, "CONTRACT_ACCESS_STORAGE")
	return *ret0, err
}

// CONTRACTACCESSSTORAGE is a free data retrieval call binding the contract method 0xceac9bd5.
//
// Solidity: function CONTRACT_ACCESS_STORAGE() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplSession) CONTRACTACCESSSTORAGE() (*big.Int, error) {
	return _ContractStorageImpl.Contract.CONTRACTACCESSSTORAGE(&_ContractStorageImpl.CallOpts)
}

// CONTRACTACCESSSTORAGE is a free data retrieval call binding the contract method 0xceac9bd5.
//
// Solidity: function CONTRACT_ACCESS_STORAGE() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCallerSession) CONTRACTACCESSSTORAGE() (*big.Int, error) {
	return _ContractStorageImpl.Contract.CONTRACTACCESSSTORAGE(&_ContractStorageImpl.CallOpts)
}

// CONTRACTADMINABLE is a free data retrieval call binding the contract method 0xbbe3a78f.
//
// Solidity: function CONTRACT_ADMINABLE() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCaller) CONTRACTADMINABLE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ContractStorageImpl.contract.Call(opts, out, "CONTRACT_ADMINABLE")
	return *ret0, err
}

// CONTRACTADMINABLE is a free data retrieval call binding the contract method 0xbbe3a78f.
//
// Solidity: function CONTRACT_ADMINABLE() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplSession) CONTRACTADMINABLE() (*big.Int, error) {
	return _ContractStorageImpl.Contract.CONTRACTADMINABLE(&_ContractStorageImpl.CallOpts)
}

// CONTRACTADMINABLE is a free data retrieval call binding the contract method 0xbbe3a78f.
//
// Solidity: function CONTRACT_ADMINABLE() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCallerSession) CONTRACTADMINABLE() (*big.Int, error) {
	return _ContractStorageImpl.Contract.CONTRACTADMINABLE(&_ContractStorageImpl.CallOpts)
}

// CONTRACTADMINSTORAGE is a free data retrieval call binding the contract method 0xc9aebae9.
//
// Solidity: function CONTRACT_ADMIN_STORAGE() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCaller) CONTRACTADMINSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ContractStorageImpl.contract.Call(opts, out, "CONTRACT_ADMIN_STORAGE")
	return *ret0, err
}

// CONTRACTADMINSTORAGE is a free data retrieval call binding the contract method 0xc9aebae9.
//
// Solidity: function CONTRACT_ADMIN_STORAGE() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplSession) CONTRACTADMINSTORAGE() (*big.Int, error) {
	return _ContractStorageImpl.Contract.CONTRACTADMINSTORAGE(&_ContractStorageImpl.CallOpts)
}

// CONTRACTADMINSTORAGE is a free data retrieval call binding the contract method 0xc9aebae9.
//
// Solidity: function CONTRACT_ADMIN_STORAGE() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCallerSession) CONTRACTADMINSTORAGE() (*big.Int, error) {
	return _ContractStorageImpl.Contract.CONTRACTADMINSTORAGE(&_ContractStorageImpl.CallOpts)
}

// CONTRACTASSETBOX is a free data retrieval call binding the contract method 0x161aa17c.
//
// Solidity: function CONTRACT_ASSETBOX() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCaller) CONTRACTASSETBOX(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ContractStorageImpl.contract.Call(opts, out, "CONTRACT_ASSETBOX")
	return *ret0, err
}

// CONTRACTASSETBOX is a free data retrieval call binding the contract method 0x161aa17c.
//
// Solidity: function CONTRACT_ASSETBOX() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplSession) CONTRACTASSETBOX() (*big.Int, error) {
	return _ContractStorageImpl.Contract.CONTRACTASSETBOX(&_ContractStorageImpl.CallOpts)
}

// CONTRACTASSETBOX is a free data retrieval call binding the contract method 0x161aa17c.
//
// Solidity: function CONTRACT_ASSETBOX() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCallerSession) CONTRACTASSETBOX() (*big.Int, error) {
	return _ContractStorageImpl.Contract.CONTRACTASSETBOX(&_ContractStorageImpl.CallOpts)
}

// CONTRACTASSETBOXINFO is a free data retrieval call binding the contract method 0xb572fd3d.
//
// Solidity: function CONTRACT_ASSETBOX_INFO() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCaller) CONTRACTASSETBOXINFO(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ContractStorageImpl.contract.Call(opts, out, "CONTRACT_ASSETBOX_INFO")
	return *ret0, err
}

// CONTRACTASSETBOXINFO is a free data retrieval call binding the contract method 0xb572fd3d.
//
// Solidity: function CONTRACT_ASSETBOX_INFO() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplSession) CONTRACTASSETBOXINFO() (*big.Int, error) {
	return _ContractStorageImpl.Contract.CONTRACTASSETBOXINFO(&_ContractStorageImpl.CallOpts)
}

// CONTRACTASSETBOXINFO is a free data retrieval call binding the contract method 0xb572fd3d.
//
// Solidity: function CONTRACT_ASSETBOX_INFO() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCallerSession) CONTRACTASSETBOXINFO() (*big.Int, error) {
	return _ContractStorageImpl.Contract.CONTRACTASSETBOXINFO(&_ContractStorageImpl.CallOpts)
}

// CONTRACTASSETBOXINFOSTORAGE is a free data retrieval call binding the contract method 0xf32a115d.
//
// Solidity: function CONTRACT_ASSETBOX_INFO_STORAGE() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCaller) CONTRACTASSETBOXINFOSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ContractStorageImpl.contract.Call(opts, out, "CONTRACT_ASSETBOX_INFO_STORAGE")
	return *ret0, err
}

// CONTRACTASSETBOXINFOSTORAGE is a free data retrieval call binding the contract method 0xf32a115d.
//
// Solidity: function CONTRACT_ASSETBOX_INFO_STORAGE() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplSession) CONTRACTASSETBOXINFOSTORAGE() (*big.Int, error) {
	return _ContractStorageImpl.Contract.CONTRACTASSETBOXINFOSTORAGE(&_ContractStorageImpl.CallOpts)
}

// CONTRACTASSETBOXINFOSTORAGE is a free data retrieval call binding the contract method 0xf32a115d.
//
// Solidity: function CONTRACT_ASSETBOX_INFO_STORAGE() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCallerSession) CONTRACTASSETBOXINFOSTORAGE() (*big.Int, error) {
	return _ContractStorageImpl.Contract.CONTRACTASSETBOXINFOSTORAGE(&_ContractStorageImpl.CallOpts)
}

// CONTRACTASSETBOXSTORAGE is a free data retrieval call binding the contract method 0xfa5bdafc.
//
// Solidity: function CONTRACT_ASSETBOX_STORAGE() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCaller) CONTRACTASSETBOXSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ContractStorageImpl.contract.Call(opts, out, "CONTRACT_ASSETBOX_STORAGE")
	return *ret0, err
}

// CONTRACTASSETBOXSTORAGE is a free data retrieval call binding the contract method 0xfa5bdafc.
//
// Solidity: function CONTRACT_ASSETBOX_STORAGE() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplSession) CONTRACTASSETBOXSTORAGE() (*big.Int, error) {
	return _ContractStorageImpl.Contract.CONTRACTASSETBOXSTORAGE(&_ContractStorageImpl.CallOpts)
}

// CONTRACTASSETBOXSTORAGE is a free data retrieval call binding the contract method 0xfa5bdafc.
//
// Solidity: function CONTRACT_ASSETBOX_STORAGE() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCallerSession) CONTRACTASSETBOXSTORAGE() (*big.Int, error) {
	return _ContractStorageImpl.Contract.CONTRACTASSETBOXSTORAGE(&_ContractStorageImpl.CallOpts)
}

// CONTRACTBALANCE is a free data retrieval call binding the contract method 0xd59045f5.
//
// Solidity: function CONTRACT_BALANCE() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCaller) CONTRACTBALANCE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ContractStorageImpl.contract.Call(opts, out, "CONTRACT_BALANCE")
	return *ret0, err
}

// CONTRACTBALANCE is a free data retrieval call binding the contract method 0xd59045f5.
//
// Solidity: function CONTRACT_BALANCE() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplSession) CONTRACTBALANCE() (*big.Int, error) {
	return _ContractStorageImpl.Contract.CONTRACTBALANCE(&_ContractStorageImpl.CallOpts)
}

// CONTRACTBALANCE is a free data retrieval call binding the contract method 0xd59045f5.
//
// Solidity: function CONTRACT_BALANCE() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCallerSession) CONTRACTBALANCE() (*big.Int, error) {
	return _ContractStorageImpl.Contract.CONTRACTBALANCE(&_ContractStorageImpl.CallOpts)
}

// CONTRACTBITBON is a free data retrieval call binding the contract method 0xb6b433ab.
//
// Solidity: function CONTRACT_BITBON() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCaller) CONTRACTBITBON(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ContractStorageImpl.contract.Call(opts, out, "CONTRACT_BITBON")
	return *ret0, err
}

// CONTRACTBITBON is a free data retrieval call binding the contract method 0xb6b433ab.
//
// Solidity: function CONTRACT_BITBON() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplSession) CONTRACTBITBON() (*big.Int, error) {
	return _ContractStorageImpl.Contract.CONTRACTBITBON(&_ContractStorageImpl.CallOpts)
}

// CONTRACTBITBON is a free data retrieval call binding the contract method 0xb6b433ab.
//
// Solidity: function CONTRACT_BITBON() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCallerSession) CONTRACTBITBON() (*big.Int, error) {
	return _ContractStorageImpl.Contract.CONTRACTBITBON(&_ContractStorageImpl.CallOpts)
}

// CONTRACTBITBONSTORAGE is a free data retrieval call binding the contract method 0xe7f6c30e.
//
// Solidity: function CONTRACT_BITBON_STORAGE() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCaller) CONTRACTBITBONSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ContractStorageImpl.contract.Call(opts, out, "CONTRACT_BITBON_STORAGE")
	return *ret0, err
}

// CONTRACTBITBONSTORAGE is a free data retrieval call binding the contract method 0xe7f6c30e.
//
// Solidity: function CONTRACT_BITBON_STORAGE() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplSession) CONTRACTBITBONSTORAGE() (*big.Int, error) {
	return _ContractStorageImpl.Contract.CONTRACTBITBONSTORAGE(&_ContractStorageImpl.CallOpts)
}

// CONTRACTBITBONSTORAGE is a free data retrieval call binding the contract method 0xe7f6c30e.
//
// Solidity: function CONTRACT_BITBON_STORAGE() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCallerSession) CONTRACTBITBONSTORAGE() (*big.Int, error) {
	return _ContractStorageImpl.Contract.CONTRACTBITBONSTORAGE(&_ContractStorageImpl.CallOpts)
}

// CONTRACTBITBONSUPPORT is a free data retrieval call binding the contract method 0x20139c74.
//
// Solidity: function CONTRACT_BITBON_SUPPORT() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCaller) CONTRACTBITBONSUPPORT(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ContractStorageImpl.contract.Call(opts, out, "CONTRACT_BITBON_SUPPORT")
	return *ret0, err
}

// CONTRACTBITBONSUPPORT is a free data retrieval call binding the contract method 0x20139c74.
//
// Solidity: function CONTRACT_BITBON_SUPPORT() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplSession) CONTRACTBITBONSUPPORT() (*big.Int, error) {
	return _ContractStorageImpl.Contract.CONTRACTBITBONSUPPORT(&_ContractStorageImpl.CallOpts)
}

// CONTRACTBITBONSUPPORT is a free data retrieval call binding the contract method 0x20139c74.
//
// Solidity: function CONTRACT_BITBON_SUPPORT() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCallerSession) CONTRACTBITBONSUPPORT() (*big.Int, error) {
	return _ContractStorageImpl.Contract.CONTRACTBITBONSUPPORT(&_ContractStorageImpl.CallOpts)
}

// CONTRACTBONBALANCE is a free data retrieval call binding the contract method 0x40e02b7c.
//
// Solidity: function CONTRACT_BONBALANCE() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCaller) CONTRACTBONBALANCE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ContractStorageImpl.contract.Call(opts, out, "CONTRACT_BONBALANCE")
	return *ret0, err
}

// CONTRACTBONBALANCE is a free data retrieval call binding the contract method 0x40e02b7c.
//
// Solidity: function CONTRACT_BONBALANCE() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplSession) CONTRACTBONBALANCE() (*big.Int, error) {
	return _ContractStorageImpl.Contract.CONTRACTBONBALANCE(&_ContractStorageImpl.CallOpts)
}

// CONTRACTBONBALANCE is a free data retrieval call binding the contract method 0x40e02b7c.
//
// Solidity: function CONTRACT_BONBALANCE() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCallerSession) CONTRACTBONBALANCE() (*big.Int, error) {
	return _ContractStorageImpl.Contract.CONTRACTBONBALANCE(&_ContractStorageImpl.CallOpts)
}

// CONTRACTBONTRANSFER is a free data retrieval call binding the contract method 0xc18f4a8a.
//
// Solidity: function CONTRACT_BONTRANSFER() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCaller) CONTRACTBONTRANSFER(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ContractStorageImpl.contract.Call(opts, out, "CONTRACT_BONTRANSFER")
	return *ret0, err
}

// CONTRACTBONTRANSFER is a free data retrieval call binding the contract method 0xc18f4a8a.
//
// Solidity: function CONTRACT_BONTRANSFER() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplSession) CONTRACTBONTRANSFER() (*big.Int, error) {
	return _ContractStorageImpl.Contract.CONTRACTBONTRANSFER(&_ContractStorageImpl.CallOpts)
}

// CONTRACTBONTRANSFER is a free data retrieval call binding the contract method 0xc18f4a8a.
//
// Solidity: function CONTRACT_BONTRANSFER() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCallerSession) CONTRACTBONTRANSFER() (*big.Int, error) {
	return _ContractStorageImpl.Contract.CONTRACTBONTRANSFER(&_ContractStorageImpl.CallOpts)
}

// CONTRACTDEX is a free data retrieval call binding the contract method 0x52c1f844.
//
// Solidity: function CONTRACT_DEX() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCaller) CONTRACTDEX(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ContractStorageImpl.contract.Call(opts, out, "CONTRACT_DEX")
	return *ret0, err
}

// CONTRACTDEX is a free data retrieval call binding the contract method 0x52c1f844.
//
// Solidity: function CONTRACT_DEX() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplSession) CONTRACTDEX() (*big.Int, error) {
	return _ContractStorageImpl.Contract.CONTRACTDEX(&_ContractStorageImpl.CallOpts)
}

// CONTRACTDEX is a free data retrieval call binding the contract method 0x52c1f844.
//
// Solidity: function CONTRACT_DEX() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCallerSession) CONTRACTDEX() (*big.Int, error) {
	return _ContractStorageImpl.Contract.CONTRACTDEX(&_ContractStorageImpl.CallOpts)
}

// CONTRACTDISTRIBUTIONSTORAGE is a free data retrieval call binding the contract method 0x08097327.
//
// Solidity: function CONTRACT_DISTRIBUTION_STORAGE() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCaller) CONTRACTDISTRIBUTIONSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ContractStorageImpl.contract.Call(opts, out, "CONTRACT_DISTRIBUTION_STORAGE")
	return *ret0, err
}

// CONTRACTDISTRIBUTIONSTORAGE is a free data retrieval call binding the contract method 0x08097327.
//
// Solidity: function CONTRACT_DISTRIBUTION_STORAGE() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplSession) CONTRACTDISTRIBUTIONSTORAGE() (*big.Int, error) {
	return _ContractStorageImpl.Contract.CONTRACTDISTRIBUTIONSTORAGE(&_ContractStorageImpl.CallOpts)
}

// CONTRACTDISTRIBUTIONSTORAGE is a free data retrieval call binding the contract method 0x08097327.
//
// Solidity: function CONTRACT_DISTRIBUTION_STORAGE() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCallerSession) CONTRACTDISTRIBUTIONSTORAGE() (*big.Int, error) {
	return _ContractStorageImpl.Contract.CONTRACTDISTRIBUTIONSTORAGE(&_ContractStorageImpl.CallOpts)
}

// CONTRACTEXCHANGE is a free data retrieval call binding the contract method 0xa0e3a10e.
//
// Solidity: function CONTRACT_EXCHANGE() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCaller) CONTRACTEXCHANGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ContractStorageImpl.contract.Call(opts, out, "CONTRACT_EXCHANGE")
	return *ret0, err
}

// CONTRACTEXCHANGE is a free data retrieval call binding the contract method 0xa0e3a10e.
//
// Solidity: function CONTRACT_EXCHANGE() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplSession) CONTRACTEXCHANGE() (*big.Int, error) {
	return _ContractStorageImpl.Contract.CONTRACTEXCHANGE(&_ContractStorageImpl.CallOpts)
}

// CONTRACTEXCHANGE is a free data retrieval call binding the contract method 0xa0e3a10e.
//
// Solidity: function CONTRACT_EXCHANGE() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCallerSession) CONTRACTEXCHANGE() (*big.Int, error) {
	return _ContractStorageImpl.Contract.CONTRACTEXCHANGE(&_ContractStorageImpl.CallOpts)
}

// CONTRACTEXCHANGESTORAGE is a free data retrieval call binding the contract method 0x001570f2.
//
// Solidity: function CONTRACT_EXCHANGE_STORAGE() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCaller) CONTRACTEXCHANGESTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ContractStorageImpl.contract.Call(opts, out, "CONTRACT_EXCHANGE_STORAGE")
	return *ret0, err
}

// CONTRACTEXCHANGESTORAGE is a free data retrieval call binding the contract method 0x001570f2.
//
// Solidity: function CONTRACT_EXCHANGE_STORAGE() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplSession) CONTRACTEXCHANGESTORAGE() (*big.Int, error) {
	return _ContractStorageImpl.Contract.CONTRACTEXCHANGESTORAGE(&_ContractStorageImpl.CallOpts)
}

// CONTRACTEXCHANGESTORAGE is a free data retrieval call binding the contract method 0x001570f2.
//
// Solidity: function CONTRACT_EXCHANGE_STORAGE() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCallerSession) CONTRACTEXCHANGESTORAGE() (*big.Int, error) {
	return _ContractStorageImpl.Contract.CONTRACTEXCHANGESTORAGE(&_ContractStorageImpl.CallOpts)
}

// CONTRACTFEESTORAGE is a free data retrieval call binding the contract method 0xcd5e65a4.
//
// Solidity: function CONTRACT_FEE_STORAGE() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCaller) CONTRACTFEESTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ContractStorageImpl.contract.Call(opts, out, "CONTRACT_FEE_STORAGE")
	return *ret0, err
}

// CONTRACTFEESTORAGE is a free data retrieval call binding the contract method 0xcd5e65a4.
//
// Solidity: function CONTRACT_FEE_STORAGE() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplSession) CONTRACTFEESTORAGE() (*big.Int, error) {
	return _ContractStorageImpl.Contract.CONTRACTFEESTORAGE(&_ContractStorageImpl.CallOpts)
}

// CONTRACTFEESTORAGE is a free data retrieval call binding the contract method 0xcd5e65a4.
//
// Solidity: function CONTRACT_FEE_STORAGE() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCallerSession) CONTRACTFEESTORAGE() (*big.Int, error) {
	return _ContractStorageImpl.Contract.CONTRACTFEESTORAGE(&_ContractStorageImpl.CallOpts)
}

// CONTRACTGENERATOR is a free data retrieval call binding the contract method 0xccea2b46.
//
// Solidity: function CONTRACT_GENERATOR() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCaller) CONTRACTGENERATOR(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ContractStorageImpl.contract.Call(opts, out, "CONTRACT_GENERATOR")
	return *ret0, err
}

// CONTRACTGENERATOR is a free data retrieval call binding the contract method 0xccea2b46.
//
// Solidity: function CONTRACT_GENERATOR() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplSession) CONTRACTGENERATOR() (*big.Int, error) {
	return _ContractStorageImpl.Contract.CONTRACTGENERATOR(&_ContractStorageImpl.CallOpts)
}

// CONTRACTGENERATOR is a free data retrieval call binding the contract method 0xccea2b46.
//
// Solidity: function CONTRACT_GENERATOR() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCallerSession) CONTRACTGENERATOR() (*big.Int, error) {
	return _ContractStorageImpl.Contract.CONTRACTGENERATOR(&_ContractStorageImpl.CallOpts)
}

// CONTRACTMININGAGENTSTORAGE is a free data retrieval call binding the contract method 0x83cc588d.
//
// Solidity: function CONTRACT_MINING_AGENT_STORAGE() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCaller) CONTRACTMININGAGENTSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ContractStorageImpl.contract.Call(opts, out, "CONTRACT_MINING_AGENT_STORAGE")
	return *ret0, err
}

// CONTRACTMININGAGENTSTORAGE is a free data retrieval call binding the contract method 0x83cc588d.
//
// Solidity: function CONTRACT_MINING_AGENT_STORAGE() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplSession) CONTRACTMININGAGENTSTORAGE() (*big.Int, error) {
	return _ContractStorageImpl.Contract.CONTRACTMININGAGENTSTORAGE(&_ContractStorageImpl.CallOpts)
}

// CONTRACTMININGAGENTSTORAGE is a free data retrieval call binding the contract method 0x83cc588d.
//
// Solidity: function CONTRACT_MINING_AGENT_STORAGE() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCallerSession) CONTRACTMININGAGENTSTORAGE() (*big.Int, error) {
	return _ContractStorageImpl.Contract.CONTRACTMININGAGENTSTORAGE(&_ContractStorageImpl.CallOpts)
}

// CONTRACTMSBON is a free data retrieval call binding the contract method 0x28b3a976.
//
// Solidity: function CONTRACT_MSBON() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCaller) CONTRACTMSBON(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ContractStorageImpl.contract.Call(opts, out, "CONTRACT_MSBON")
	return *ret0, err
}

// CONTRACTMSBON is a free data retrieval call binding the contract method 0x28b3a976.
//
// Solidity: function CONTRACT_MSBON() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplSession) CONTRACTMSBON() (*big.Int, error) {
	return _ContractStorageImpl.Contract.CONTRACTMSBON(&_ContractStorageImpl.CallOpts)
}

// CONTRACTMSBON is a free data retrieval call binding the contract method 0x28b3a976.
//
// Solidity: function CONTRACT_MSBON() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCallerSession) CONTRACTMSBON() (*big.Int, error) {
	return _ContractStorageImpl.Contract.CONTRACTMSBON(&_ContractStorageImpl.CallOpts)
}

// CONTRACTMSBONSTORAGE is a free data retrieval call binding the contract method 0xe04e6061.
//
// Solidity: function CONTRACT_MSBON_STORAGE() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCaller) CONTRACTMSBONSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ContractStorageImpl.contract.Call(opts, out, "CONTRACT_MSBON_STORAGE")
	return *ret0, err
}

// CONTRACTMSBONSTORAGE is a free data retrieval call binding the contract method 0xe04e6061.
//
// Solidity: function CONTRACT_MSBON_STORAGE() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplSession) CONTRACTMSBONSTORAGE() (*big.Int, error) {
	return _ContractStorageImpl.Contract.CONTRACTMSBONSTORAGE(&_ContractStorageImpl.CallOpts)
}

// CONTRACTMSBONSTORAGE is a free data retrieval call binding the contract method 0xe04e6061.
//
// Solidity: function CONTRACT_MSBON_STORAGE() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCallerSession) CONTRACTMSBONSTORAGE() (*big.Int, error) {
	return _ContractStorageImpl.Contract.CONTRACTMSBONSTORAGE(&_ContractStorageImpl.CallOpts)
}

// CONTRACTMULTISIGADDADDRESS is a free data retrieval call binding the contract method 0xf4f65a5a.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ADDRESS() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCaller) CONTRACTMULTISIGADDADDRESS(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ContractStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_ADD_ADDRESS")
	return *ret0, err
}

// CONTRACTMULTISIGADDADDRESS is a free data retrieval call binding the contract method 0xf4f65a5a.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ADDRESS() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplSession) CONTRACTMULTISIGADDADDRESS() (*big.Int, error) {
	return _ContractStorageImpl.Contract.CONTRACTMULTISIGADDADDRESS(&_ContractStorageImpl.CallOpts)
}

// CONTRACTMULTISIGADDADDRESS is a free data retrieval call binding the contract method 0xf4f65a5a.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ADDRESS() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCallerSession) CONTRACTMULTISIGADDADDRESS() (*big.Int, error) {
	return _ContractStorageImpl.Contract.CONTRACTMULTISIGADDADDRESS(&_ContractStorageImpl.CallOpts)
}

// CONTRACTMULTISIGADDADMIN is a free data retrieval call binding the contract method 0x83888c68.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ADMIN() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCaller) CONTRACTMULTISIGADDADMIN(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ContractStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_ADD_ADMIN")
	return *ret0, err
}

// CONTRACTMULTISIGADDADMIN is a free data retrieval call binding the contract method 0x83888c68.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ADMIN() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplSession) CONTRACTMULTISIGADDADMIN() (*big.Int, error) {
	return _ContractStorageImpl.Contract.CONTRACTMULTISIGADDADMIN(&_ContractStorageImpl.CallOpts)
}

// CONTRACTMULTISIGADDADMIN is a free data retrieval call binding the contract method 0x83888c68.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ADMIN() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCallerSession) CONTRACTMULTISIGADDADMIN() (*big.Int, error) {
	return _ContractStorageImpl.Contract.CONTRACTMULTISIGADDADMIN(&_ContractStorageImpl.CallOpts)
}

// CONTRACTMULTISIGADDFEE is a free data retrieval call binding the contract method 0xc43e8cab.
//
// Solidity: function CONTRACT_MULTISIG_ADD_FEE() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCaller) CONTRACTMULTISIGADDFEE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ContractStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_ADD_FEE")
	return *ret0, err
}

// CONTRACTMULTISIGADDFEE is a free data retrieval call binding the contract method 0xc43e8cab.
//
// Solidity: function CONTRACT_MULTISIG_ADD_FEE() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplSession) CONTRACTMULTISIGADDFEE() (*big.Int, error) {
	return _ContractStorageImpl.Contract.CONTRACTMULTISIGADDFEE(&_ContractStorageImpl.CallOpts)
}

// CONTRACTMULTISIGADDFEE is a free data retrieval call binding the contract method 0xc43e8cab.
//
// Solidity: function CONTRACT_MULTISIG_ADD_FEE() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCallerSession) CONTRACTMULTISIGADDFEE() (*big.Int, error) {
	return _ContractStorageImpl.Contract.CONTRACTMULTISIGADDFEE(&_ContractStorageImpl.CallOpts)
}

// CONTRACTMULTISIGADDROLE is a free data retrieval call binding the contract method 0x20a0914b.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ROLE() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCaller) CONTRACTMULTISIGADDROLE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ContractStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_ADD_ROLE")
	return *ret0, err
}

// CONTRACTMULTISIGADDROLE is a free data retrieval call binding the contract method 0x20a0914b.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ROLE() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplSession) CONTRACTMULTISIGADDROLE() (*big.Int, error) {
	return _ContractStorageImpl.Contract.CONTRACTMULTISIGADDROLE(&_ContractStorageImpl.CallOpts)
}

// CONTRACTMULTISIGADDROLE is a free data retrieval call binding the contract method 0x20a0914b.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ROLE() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCallerSession) CONTRACTMULTISIGADDROLE() (*big.Int, error) {
	return _ContractStorageImpl.Contract.CONTRACTMULTISIGADDROLE(&_ContractStorageImpl.CallOpts)
}

// CONTRACTMULTISIGDROPPERMISSION is a free data retrieval call binding the contract method 0x3f3fbc39.
//
// Solidity: function CONTRACT_MULTISIG_DROP_PERMISSION() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCaller) CONTRACTMULTISIGDROPPERMISSION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ContractStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_DROP_PERMISSION")
	return *ret0, err
}

// CONTRACTMULTISIGDROPPERMISSION is a free data retrieval call binding the contract method 0x3f3fbc39.
//
// Solidity: function CONTRACT_MULTISIG_DROP_PERMISSION() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplSession) CONTRACTMULTISIGDROPPERMISSION() (*big.Int, error) {
	return _ContractStorageImpl.Contract.CONTRACTMULTISIGDROPPERMISSION(&_ContractStorageImpl.CallOpts)
}

// CONTRACTMULTISIGDROPPERMISSION is a free data retrieval call binding the contract method 0x3f3fbc39.
//
// Solidity: function CONTRACT_MULTISIG_DROP_PERMISSION() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCallerSession) CONTRACTMULTISIGDROPPERMISSION() (*big.Int, error) {
	return _ContractStorageImpl.Contract.CONTRACTMULTISIGDROPPERMISSION(&_ContractStorageImpl.CallOpts)
}

// CONTRACTMULTISIGEDITADMIN is a free data retrieval call binding the contract method 0x91b0f051.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_ADMIN() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCaller) CONTRACTMULTISIGEDITADMIN(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ContractStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_EDIT_ADMIN")
	return *ret0, err
}

// CONTRACTMULTISIGEDITADMIN is a free data retrieval call binding the contract method 0x91b0f051.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_ADMIN() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplSession) CONTRACTMULTISIGEDITADMIN() (*big.Int, error) {
	return _ContractStorageImpl.Contract.CONTRACTMULTISIGEDITADMIN(&_ContractStorageImpl.CallOpts)
}

// CONTRACTMULTISIGEDITADMIN is a free data retrieval call binding the contract method 0x91b0f051.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_ADMIN() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCallerSession) CONTRACTMULTISIGEDITADMIN() (*big.Int, error) {
	return _ContractStorageImpl.Contract.CONTRACTMULTISIGEDITADMIN(&_ContractStorageImpl.CallOpts)
}

// CONTRACTMULTISIGEDITFEE is a free data retrieval call binding the contract method 0x6b2abe17.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_FEE() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCaller) CONTRACTMULTISIGEDITFEE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ContractStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_EDIT_FEE")
	return *ret0, err
}

// CONTRACTMULTISIGEDITFEE is a free data retrieval call binding the contract method 0x6b2abe17.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_FEE() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplSession) CONTRACTMULTISIGEDITFEE() (*big.Int, error) {
	return _ContractStorageImpl.Contract.CONTRACTMULTISIGEDITFEE(&_ContractStorageImpl.CallOpts)
}

// CONTRACTMULTISIGEDITFEE is a free data retrieval call binding the contract method 0x6b2abe17.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_FEE() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCallerSession) CONTRACTMULTISIGEDITFEE() (*big.Int, error) {
	return _ContractStorageImpl.Contract.CONTRACTMULTISIGEDITFEE(&_ContractStorageImpl.CallOpts)
}

// CONTRACTMULTISIGEDITROLE is a free data retrieval call binding the contract method 0xc5f66412.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_ROLE() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCaller) CONTRACTMULTISIGEDITROLE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ContractStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_EDIT_ROLE")
	return *ret0, err
}

// CONTRACTMULTISIGEDITROLE is a free data retrieval call binding the contract method 0xc5f66412.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_ROLE() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplSession) CONTRACTMULTISIGEDITROLE() (*big.Int, error) {
	return _ContractStorageImpl.Contract.CONTRACTMULTISIGEDITROLE(&_ContractStorageImpl.CallOpts)
}

// CONTRACTMULTISIGEDITROLE is a free data retrieval call binding the contract method 0xc5f66412.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_ROLE() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCallerSession) CONTRACTMULTISIGEDITROLE() (*big.Int, error) {
	return _ContractStorageImpl.Contract.CONTRACTMULTISIGEDITROLE(&_ContractStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEADMINBYADDRESS is a free data retrieval call binding the contract method 0xc816b091.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_BY_ADDRESS() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCaller) CONTRACTMULTISIGREMOVEADMINBYADDRESS(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ContractStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REMOVE_ADMIN_BY_ADDRESS")
	return *ret0, err
}

// CONTRACTMULTISIGREMOVEADMINBYADDRESS is a free data retrieval call binding the contract method 0xc816b091.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_BY_ADDRESS() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplSession) CONTRACTMULTISIGREMOVEADMINBYADDRESS() (*big.Int, error) {
	return _ContractStorageImpl.Contract.CONTRACTMULTISIGREMOVEADMINBYADDRESS(&_ContractStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEADMINBYADDRESS is a free data retrieval call binding the contract method 0xc816b091.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_BY_ADDRESS() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCallerSession) CONTRACTMULTISIGREMOVEADMINBYADDRESS() (*big.Int, error) {
	return _ContractStorageImpl.Contract.CONTRACTMULTISIGREMOVEADMINBYADDRESS(&_ContractStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEADMINBYINDEX is a free data retrieval call binding the contract method 0x751fbea2.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_BY_INDEX() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCaller) CONTRACTMULTISIGREMOVEADMINBYINDEX(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ContractStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REMOVE_ADMIN_BY_INDEX")
	return *ret0, err
}

// CONTRACTMULTISIGREMOVEADMINBYINDEX is a free data retrieval call binding the contract method 0x751fbea2.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_BY_INDEX() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplSession) CONTRACTMULTISIGREMOVEADMINBYINDEX() (*big.Int, error) {
	return _ContractStorageImpl.Contract.CONTRACTMULTISIGREMOVEADMINBYINDEX(&_ContractStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEADMINBYINDEX is a free data retrieval call binding the contract method 0x751fbea2.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_BY_INDEX() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCallerSession) CONTRACTMULTISIGREMOVEADMINBYINDEX() (*big.Int, error) {
	return _ContractStorageImpl.Contract.CONTRACTMULTISIGREMOVEADMINBYINDEX(&_ContractStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEADMINROLE is a free data retrieval call binding the contract method 0x53c00200.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_ROLE() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCaller) CONTRACTMULTISIGREMOVEADMINROLE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ContractStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REMOVE_ADMIN_ROLE")
	return *ret0, err
}

// CONTRACTMULTISIGREMOVEADMINROLE is a free data retrieval call binding the contract method 0x53c00200.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_ROLE() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplSession) CONTRACTMULTISIGREMOVEADMINROLE() (*big.Int, error) {
	return _ContractStorageImpl.Contract.CONTRACTMULTISIGREMOVEADMINROLE(&_ContractStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEADMINROLE is a free data retrieval call binding the contract method 0x53c00200.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_ROLE() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCallerSession) CONTRACTMULTISIGREMOVEADMINROLE() (*big.Int, error) {
	return _ContractStorageImpl.Contract.CONTRACTMULTISIGREMOVEADMINROLE(&_ContractStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEFEE is a free data retrieval call binding the contract method 0xdb614d25.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_FEE() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCaller) CONTRACTMULTISIGREMOVEFEE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ContractStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REMOVE_FEE")
	return *ret0, err
}

// CONTRACTMULTISIGREMOVEFEE is a free data retrieval call binding the contract method 0xdb614d25.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_FEE() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplSession) CONTRACTMULTISIGREMOVEFEE() (*big.Int, error) {
	return _ContractStorageImpl.Contract.CONTRACTMULTISIGREMOVEFEE(&_ContractStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEFEE is a free data retrieval call binding the contract method 0xdb614d25.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_FEE() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCallerSession) CONTRACTMULTISIGREMOVEFEE() (*big.Int, error) {
	return _ContractStorageImpl.Contract.CONTRACTMULTISIGREMOVEFEE(&_ContractStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEROLE is a free data retrieval call binding the contract method 0x23648496.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ROLE() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCaller) CONTRACTMULTISIGREMOVEROLE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ContractStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REMOVE_ROLE")
	return *ret0, err
}

// CONTRACTMULTISIGREMOVEROLE is a free data retrieval call binding the contract method 0x23648496.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ROLE() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplSession) CONTRACTMULTISIGREMOVEROLE() (*big.Int, error) {
	return _ContractStorageImpl.Contract.CONTRACTMULTISIGREMOVEROLE(&_ContractStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEROLE is a free data retrieval call binding the contract method 0x23648496.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ROLE() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCallerSession) CONTRACTMULTISIGREMOVEROLE() (*big.Int, error) {
	return _ContractStorageImpl.Contract.CONTRACTMULTISIGREMOVEROLE(&_ContractStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREPLENISHCAPITALIZATION is a free data retrieval call binding the contract method 0x62d5d523.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_CAPITALIZATION() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCaller) CONTRACTMULTISIGREPLENISHCAPITALIZATION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ContractStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REPLENISH_CAPITALIZATION")
	return *ret0, err
}

// CONTRACTMULTISIGREPLENISHCAPITALIZATION is a free data retrieval call binding the contract method 0x62d5d523.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_CAPITALIZATION() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplSession) CONTRACTMULTISIGREPLENISHCAPITALIZATION() (*big.Int, error) {
	return _ContractStorageImpl.Contract.CONTRACTMULTISIGREPLENISHCAPITALIZATION(&_ContractStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREPLENISHCAPITALIZATION is a free data retrieval call binding the contract method 0x62d5d523.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_CAPITALIZATION() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCallerSession) CONTRACTMULTISIGREPLENISHCAPITALIZATION() (*big.Int, error) {
	return _ContractStorageImpl.Contract.CONTRACTMULTISIGREPLENISHCAPITALIZATION(&_ContractStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREPLENISHCERTIFICATE is a free data retrieval call binding the contract method 0xfd19b7e1.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_CERTIFICATE() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCaller) CONTRACTMULTISIGREPLENISHCERTIFICATE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ContractStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REPLENISH_CERTIFICATE")
	return *ret0, err
}

// CONTRACTMULTISIGREPLENISHCERTIFICATE is a free data retrieval call binding the contract method 0xfd19b7e1.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_CERTIFICATE() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplSession) CONTRACTMULTISIGREPLENISHCERTIFICATE() (*big.Int, error) {
	return _ContractStorageImpl.Contract.CONTRACTMULTISIGREPLENISHCERTIFICATE(&_ContractStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREPLENISHCERTIFICATE is a free data retrieval call binding the contract method 0xfd19b7e1.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_CERTIFICATE() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCallerSession) CONTRACTMULTISIGREPLENISHCERTIFICATE() (*big.Int, error) {
	return _ContractStorageImpl.Contract.CONTRACTMULTISIGREPLENISHCERTIFICATE(&_ContractStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREPLENISHEMISSION is a free data retrieval call binding the contract method 0x0cb59fde.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_EMISSION() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCaller) CONTRACTMULTISIGREPLENISHEMISSION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ContractStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REPLENISH_EMISSION")
	return *ret0, err
}

// CONTRACTMULTISIGREPLENISHEMISSION is a free data retrieval call binding the contract method 0x0cb59fde.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_EMISSION() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplSession) CONTRACTMULTISIGREPLENISHEMISSION() (*big.Int, error) {
	return _ContractStorageImpl.Contract.CONTRACTMULTISIGREPLENISHEMISSION(&_ContractStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREPLENISHEMISSION is a free data retrieval call binding the contract method 0x0cb59fde.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_EMISSION() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCallerSession) CONTRACTMULTISIGREPLENISHEMISSION() (*big.Int, error) {
	return _ContractStorageImpl.Contract.CONTRACTMULTISIGREPLENISHEMISSION(&_ContractStorageImpl.CallOpts)
}

// CONTRACTMULTISIGSETADMINROLE is a free data retrieval call binding the contract method 0x8f209ab7.
//
// Solidity: function CONTRACT_MULTISIG_SET_ADMIN_ROLE() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCaller) CONTRACTMULTISIGSETADMINROLE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ContractStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_SET_ADMIN_ROLE")
	return *ret0, err
}

// CONTRACTMULTISIGSETADMINROLE is a free data retrieval call binding the contract method 0x8f209ab7.
//
// Solidity: function CONTRACT_MULTISIG_SET_ADMIN_ROLE() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplSession) CONTRACTMULTISIGSETADMINROLE() (*big.Int, error) {
	return _ContractStorageImpl.Contract.CONTRACTMULTISIGSETADMINROLE(&_ContractStorageImpl.CallOpts)
}

// CONTRACTMULTISIGSETADMINROLE is a free data retrieval call binding the contract method 0x8f209ab7.
//
// Solidity: function CONTRACT_MULTISIG_SET_ADMIN_ROLE() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCallerSession) CONTRACTMULTISIGSETADMINROLE() (*big.Int, error) {
	return _ContractStorageImpl.Contract.CONTRACTMULTISIGSETADMINROLE(&_ContractStorageImpl.CallOpts)
}

// CONTRACTMULTISIGSETPERMISSION is a free data retrieval call binding the contract method 0xb0e91914.
//
// Solidity: function CONTRACT_MULTISIG_SET_PERMISSION() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCaller) CONTRACTMULTISIGSETPERMISSION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ContractStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_SET_PERMISSION")
	return *ret0, err
}

// CONTRACTMULTISIGSETPERMISSION is a free data retrieval call binding the contract method 0xb0e91914.
//
// Solidity: function CONTRACT_MULTISIG_SET_PERMISSION() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplSession) CONTRACTMULTISIGSETPERMISSION() (*big.Int, error) {
	return _ContractStorageImpl.Contract.CONTRACTMULTISIGSETPERMISSION(&_ContractStorageImpl.CallOpts)
}

// CONTRACTMULTISIGSETPERMISSION is a free data retrieval call binding the contract method 0xb0e91914.
//
// Solidity: function CONTRACT_MULTISIG_SET_PERMISSION() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCallerSession) CONTRACTMULTISIGSETPERMISSION() (*big.Int, error) {
	return _ContractStorageImpl.Contract.CONTRACTMULTISIGSETPERMISSION(&_ContractStorageImpl.CallOpts)
}

// CONTRACTMULTISIGSTORAGE is a free data retrieval call binding the contract method 0xe7e1a303.
//
// Solidity: function CONTRACT_MULTISIG_STORAGE() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCaller) CONTRACTMULTISIGSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ContractStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_STORAGE")
	return *ret0, err
}

// CONTRACTMULTISIGSTORAGE is a free data retrieval call binding the contract method 0xe7e1a303.
//
// Solidity: function CONTRACT_MULTISIG_STORAGE() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplSession) CONTRACTMULTISIGSTORAGE() (*big.Int, error) {
	return _ContractStorageImpl.Contract.CONTRACTMULTISIGSTORAGE(&_ContractStorageImpl.CallOpts)
}

// CONTRACTMULTISIGSTORAGE is a free data retrieval call binding the contract method 0xe7e1a303.
//
// Solidity: function CONTRACT_MULTISIG_STORAGE() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCallerSession) CONTRACTMULTISIGSTORAGE() (*big.Int, error) {
	return _ContractStorageImpl.Contract.CONTRACTMULTISIGSTORAGE(&_ContractStorageImpl.CallOpts)
}

// CONTRACTOTC is a free data retrieval call binding the contract method 0x5eceac70.
//
// Solidity: function CONTRACT_OTC() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCaller) CONTRACTOTC(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ContractStorageImpl.contract.Call(opts, out, "CONTRACT_OTC")
	return *ret0, err
}

// CONTRACTOTC is a free data retrieval call binding the contract method 0x5eceac70.
//
// Solidity: function CONTRACT_OTC() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplSession) CONTRACTOTC() (*big.Int, error) {
	return _ContractStorageImpl.Contract.CONTRACTOTC(&_ContractStorageImpl.CallOpts)
}

// CONTRACTOTC is a free data retrieval call binding the contract method 0x5eceac70.
//
// Solidity: function CONTRACT_OTC() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCallerSession) CONTRACTOTC() (*big.Int, error) {
	return _ContractStorageImpl.Contract.CONTRACTOTC(&_ContractStorageImpl.CallOpts)
}

// CONTRACTOTCSTORAGE is a free data retrieval call binding the contract method 0x825b476f.
//
// Solidity: function CONTRACT_OTC_STORAGE() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCaller) CONTRACTOTCSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ContractStorageImpl.contract.Call(opts, out, "CONTRACT_OTC_STORAGE")
	return *ret0, err
}

// CONTRACTOTCSTORAGE is a free data retrieval call binding the contract method 0x825b476f.
//
// Solidity: function CONTRACT_OTC_STORAGE() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplSession) CONTRACTOTCSTORAGE() (*big.Int, error) {
	return _ContractStorageImpl.Contract.CONTRACTOTCSTORAGE(&_ContractStorageImpl.CallOpts)
}

// CONTRACTOTCSTORAGE is a free data retrieval call binding the contract method 0x825b476f.
//
// Solidity: function CONTRACT_OTC_STORAGE() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCallerSession) CONTRACTOTCSTORAGE() (*big.Int, error) {
	return _ContractStorageImpl.Contract.CONTRACTOTCSTORAGE(&_ContractStorageImpl.CallOpts)
}

// CONTRACTREPLENISHCAPITALIZATION is a free data retrieval call binding the contract method 0xff5e09d8.
//
// Solidity: function CONTRACT_REPLENISH_CAPITALIZATION() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCaller) CONTRACTREPLENISHCAPITALIZATION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ContractStorageImpl.contract.Call(opts, out, "CONTRACT_REPLENISH_CAPITALIZATION")
	return *ret0, err
}

// CONTRACTREPLENISHCAPITALIZATION is a free data retrieval call binding the contract method 0xff5e09d8.
//
// Solidity: function CONTRACT_REPLENISH_CAPITALIZATION() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplSession) CONTRACTREPLENISHCAPITALIZATION() (*big.Int, error) {
	return _ContractStorageImpl.Contract.CONTRACTREPLENISHCAPITALIZATION(&_ContractStorageImpl.CallOpts)
}

// CONTRACTREPLENISHCAPITALIZATION is a free data retrieval call binding the contract method 0xff5e09d8.
//
// Solidity: function CONTRACT_REPLENISH_CAPITALIZATION() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCallerSession) CONTRACTREPLENISHCAPITALIZATION() (*big.Int, error) {
	return _ContractStorageImpl.Contract.CONTRACTREPLENISHCAPITALIZATION(&_ContractStorageImpl.CallOpts)
}

// CONTRACTRESERVEDALIASSTORAGE is a free data retrieval call binding the contract method 0x70a94aa7.
//
// Solidity: function CONTRACT_RESERVED_ALIAS_STORAGE() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCaller) CONTRACTRESERVEDALIASSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ContractStorageImpl.contract.Call(opts, out, "CONTRACT_RESERVED_ALIAS_STORAGE")
	return *ret0, err
}

// CONTRACTRESERVEDALIASSTORAGE is a free data retrieval call binding the contract method 0x70a94aa7.
//
// Solidity: function CONTRACT_RESERVED_ALIAS_STORAGE() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplSession) CONTRACTRESERVEDALIASSTORAGE() (*big.Int, error) {
	return _ContractStorageImpl.Contract.CONTRACTRESERVEDALIASSTORAGE(&_ContractStorageImpl.CallOpts)
}

// CONTRACTRESERVEDALIASSTORAGE is a free data retrieval call binding the contract method 0x70a94aa7.
//
// Solidity: function CONTRACT_RESERVED_ALIAS_STORAGE() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCallerSession) CONTRACTRESERVEDALIASSTORAGE() (*big.Int, error) {
	return _ContractStorageImpl.Contract.CONTRACTRESERVEDALIASSTORAGE(&_ContractStorageImpl.CallOpts)
}

// CONTRACTROLESTORAGE is a free data retrieval call binding the contract method 0xb938235b.
//
// Solidity: function CONTRACT_ROLE_STORAGE() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCaller) CONTRACTROLESTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ContractStorageImpl.contract.Call(opts, out, "CONTRACT_ROLE_STORAGE")
	return *ret0, err
}

// CONTRACTROLESTORAGE is a free data retrieval call binding the contract method 0xb938235b.
//
// Solidity: function CONTRACT_ROLE_STORAGE() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplSession) CONTRACTROLESTORAGE() (*big.Int, error) {
	return _ContractStorageImpl.Contract.CONTRACTROLESTORAGE(&_ContractStorageImpl.CallOpts)
}

// CONTRACTROLESTORAGE is a free data retrieval call binding the contract method 0xb938235b.
//
// Solidity: function CONTRACT_ROLE_STORAGE() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCallerSession) CONTRACTROLESTORAGE() (*big.Int, error) {
	return _ContractStorageImpl.Contract.CONTRACTROLESTORAGE(&_ContractStorageImpl.CallOpts)
}

// CONTRACTSAFETRANSFERSTORAGE is a free data retrieval call binding the contract method 0xe7e45392.
//
// Solidity: function CONTRACT_SAFE_TRANSFER_STORAGE() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCaller) CONTRACTSAFETRANSFERSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ContractStorageImpl.contract.Call(opts, out, "CONTRACT_SAFE_TRANSFER_STORAGE")
	return *ret0, err
}

// CONTRACTSAFETRANSFERSTORAGE is a free data retrieval call binding the contract method 0xe7e45392.
//
// Solidity: function CONTRACT_SAFE_TRANSFER_STORAGE() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplSession) CONTRACTSAFETRANSFERSTORAGE() (*big.Int, error) {
	return _ContractStorageImpl.Contract.CONTRACTSAFETRANSFERSTORAGE(&_ContractStorageImpl.CallOpts)
}

// CONTRACTSAFETRANSFERSTORAGE is a free data retrieval call binding the contract method 0xe7e45392.
//
// Solidity: function CONTRACT_SAFE_TRANSFER_STORAGE() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCallerSession) CONTRACTSAFETRANSFERSTORAGE() (*big.Int, error) {
	return _ContractStorageImpl.Contract.CONTRACTSAFETRANSFERSTORAGE(&_ContractStorageImpl.CallOpts)
}

// CONTRACTTRANSFER is a free data retrieval call binding the contract method 0xef94fe38.
//
// Solidity: function CONTRACT_TRANSFER() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCaller) CONTRACTTRANSFER(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ContractStorageImpl.contract.Call(opts, out, "CONTRACT_TRANSFER")
	return *ret0, err
}

// CONTRACTTRANSFER is a free data retrieval call binding the contract method 0xef94fe38.
//
// Solidity: function CONTRACT_TRANSFER() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplSession) CONTRACTTRANSFER() (*big.Int, error) {
	return _ContractStorageImpl.Contract.CONTRACTTRANSFER(&_ContractStorageImpl.CallOpts)
}

// CONTRACTTRANSFER is a free data retrieval call binding the contract method 0xef94fe38.
//
// Solidity: function CONTRACT_TRANSFER() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCallerSession) CONTRACTTRANSFER() (*big.Int, error) {
	return _ContractStorageImpl.Contract.CONTRACTTRANSFER(&_ContractStorageImpl.CallOpts)
}

// CONTRACTTRANSFERFROMCAPITALIZATION is a free data retrieval call binding the contract method 0xc8f8decb.
//
// Solidity: function CONTRACT_TRANSFER_FROM_CAPITALIZATION() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCaller) CONTRACTTRANSFERFROMCAPITALIZATION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ContractStorageImpl.contract.Call(opts, out, "CONTRACT_TRANSFER_FROM_CAPITALIZATION")
	return *ret0, err
}

// CONTRACTTRANSFERFROMCAPITALIZATION is a free data retrieval call binding the contract method 0xc8f8decb.
//
// Solidity: function CONTRACT_TRANSFER_FROM_CAPITALIZATION() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplSession) CONTRACTTRANSFERFROMCAPITALIZATION() (*big.Int, error) {
	return _ContractStorageImpl.Contract.CONTRACTTRANSFERFROMCAPITALIZATION(&_ContractStorageImpl.CallOpts)
}

// CONTRACTTRANSFERFROMCAPITALIZATION is a free data retrieval call binding the contract method 0xc8f8decb.
//
// Solidity: function CONTRACT_TRANSFER_FROM_CAPITALIZATION() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCallerSession) CONTRACTTRANSFERFROMCAPITALIZATION() (*big.Int, error) {
	return _ContractStorageImpl.Contract.CONTRACTTRANSFERFROMCAPITALIZATION(&_ContractStorageImpl.CallOpts)
}

// PERMISSIONACCESSRESTORATION is a free data retrieval call binding the contract method 0x5c3053bf.
//
// Solidity: function PERMISSION_ACCESS_RESTORATION() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCaller) PERMISSIONACCESSRESTORATION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ContractStorageImpl.contract.Call(opts, out, "PERMISSION_ACCESS_RESTORATION")
	return *ret0, err
}

// PERMISSIONACCESSRESTORATION is a free data retrieval call binding the contract method 0x5c3053bf.
//
// Solidity: function PERMISSION_ACCESS_RESTORATION() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplSession) PERMISSIONACCESSRESTORATION() (*big.Int, error) {
	return _ContractStorageImpl.Contract.PERMISSIONACCESSRESTORATION(&_ContractStorageImpl.CallOpts)
}

// PERMISSIONACCESSRESTORATION is a free data retrieval call binding the contract method 0x5c3053bf.
//
// Solidity: function PERMISSION_ACCESS_RESTORATION() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCallerSession) PERMISSIONACCESSRESTORATION() (*big.Int, error) {
	return _ContractStorageImpl.Contract.PERMISSIONACCESSRESTORATION(&_ContractStorageImpl.CallOpts)
}

// PERMISSIONADMINSTORAGE is a free data retrieval call binding the contract method 0x0c9fc0db.
//
// Solidity: function PERMISSION_ADMIN_STORAGE() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCaller) PERMISSIONADMINSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ContractStorageImpl.contract.Call(opts, out, "PERMISSION_ADMIN_STORAGE")
	return *ret0, err
}

// PERMISSIONADMINSTORAGE is a free data retrieval call binding the contract method 0x0c9fc0db.
//
// Solidity: function PERMISSION_ADMIN_STORAGE() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplSession) PERMISSIONADMINSTORAGE() (*big.Int, error) {
	return _ContractStorageImpl.Contract.PERMISSIONADMINSTORAGE(&_ContractStorageImpl.CallOpts)
}

// PERMISSIONADMINSTORAGE is a free data retrieval call binding the contract method 0x0c9fc0db.
//
// Solidity: function PERMISSION_ADMIN_STORAGE() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCallerSession) PERMISSIONADMINSTORAGE() (*big.Int, error) {
	return _ContractStorageImpl.Contract.PERMISSIONADMINSTORAGE(&_ContractStorageImpl.CallOpts)
}

// PERMISSIONCONTRACTSSTORAGE is a free data retrieval call binding the contract method 0x1516e83b.
//
// Solidity: function PERMISSION_CONTRACTS_STORAGE() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCaller) PERMISSIONCONTRACTSSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ContractStorageImpl.contract.Call(opts, out, "PERMISSION_CONTRACTS_STORAGE")
	return *ret0, err
}

// PERMISSIONCONTRACTSSTORAGE is a free data retrieval call binding the contract method 0x1516e83b.
//
// Solidity: function PERMISSION_CONTRACTS_STORAGE() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplSession) PERMISSIONCONTRACTSSTORAGE() (*big.Int, error) {
	return _ContractStorageImpl.Contract.PERMISSIONCONTRACTSSTORAGE(&_ContractStorageImpl.CallOpts)
}

// PERMISSIONCONTRACTSSTORAGE is a free data retrieval call binding the contract method 0x1516e83b.
//
// Solidity: function PERMISSION_CONTRACTS_STORAGE() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCallerSession) PERMISSIONCONTRACTSSTORAGE() (*big.Int, error) {
	return _ContractStorageImpl.Contract.PERMISSIONCONTRACTSSTORAGE(&_ContractStorageImpl.CallOpts)
}

// PERMISSIONCONTRACTDEPLOY is a free data retrieval call binding the contract method 0x257a4a04.
//
// Solidity: function PERMISSION_CONTRACT_DEPLOY() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCaller) PERMISSIONCONTRACTDEPLOY(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ContractStorageImpl.contract.Call(opts, out, "PERMISSION_CONTRACT_DEPLOY")
	return *ret0, err
}

// PERMISSIONCONTRACTDEPLOY is a free data retrieval call binding the contract method 0x257a4a04.
//
// Solidity: function PERMISSION_CONTRACT_DEPLOY() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplSession) PERMISSIONCONTRACTDEPLOY() (*big.Int, error) {
	return _ContractStorageImpl.Contract.PERMISSIONCONTRACTDEPLOY(&_ContractStorageImpl.CallOpts)
}

// PERMISSIONCONTRACTDEPLOY is a free data retrieval call binding the contract method 0x257a4a04.
//
// Solidity: function PERMISSION_CONTRACT_DEPLOY() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCallerSession) PERMISSIONCONTRACTDEPLOY() (*big.Int, error) {
	return _ContractStorageImpl.Contract.PERMISSIONCONTRACTDEPLOY(&_ContractStorageImpl.CallOpts)
}

// PERMISSIONEMISSION is a free data retrieval call binding the contract method 0x4ccd92d1.
//
// Solidity: function PERMISSION_EMISSION() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCaller) PERMISSIONEMISSION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ContractStorageImpl.contract.Call(opts, out, "PERMISSION_EMISSION")
	return *ret0, err
}

// PERMISSIONEMISSION is a free data retrieval call binding the contract method 0x4ccd92d1.
//
// Solidity: function PERMISSION_EMISSION() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplSession) PERMISSIONEMISSION() (*big.Int, error) {
	return _ContractStorageImpl.Contract.PERMISSIONEMISSION(&_ContractStorageImpl.CallOpts)
}

// PERMISSIONEMISSION is a free data retrieval call binding the contract method 0x4ccd92d1.
//
// Solidity: function PERMISSION_EMISSION() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCallerSession) PERMISSIONEMISSION() (*big.Int, error) {
	return _ContractStorageImpl.Contract.PERMISSIONEMISSION(&_ContractStorageImpl.CallOpts)
}

// PERMISSIONFEESSTORAGE is a free data retrieval call binding the contract method 0xe837eb87.
//
// Solidity: function PERMISSION_FEES_STORAGE() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCaller) PERMISSIONFEESSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ContractStorageImpl.contract.Call(opts, out, "PERMISSION_FEES_STORAGE")
	return *ret0, err
}

// PERMISSIONFEESSTORAGE is a free data retrieval call binding the contract method 0xe837eb87.
//
// Solidity: function PERMISSION_FEES_STORAGE() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplSession) PERMISSIONFEESSTORAGE() (*big.Int, error) {
	return _ContractStorageImpl.Contract.PERMISSIONFEESSTORAGE(&_ContractStorageImpl.CallOpts)
}

// PERMISSIONFEESSTORAGE is a free data retrieval call binding the contract method 0xe837eb87.
//
// Solidity: function PERMISSION_FEES_STORAGE() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCallerSession) PERMISSIONFEESSTORAGE() (*big.Int, error) {
	return _ContractStorageImpl.Contract.PERMISSIONFEESSTORAGE(&_ContractStorageImpl.CallOpts)
}

// PERMISSIONMULTISIGSTORAGE is a free data retrieval call binding the contract method 0xf1bb7aad.
//
// Solidity: function PERMISSION_MULTISIG_STORAGE() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCaller) PERMISSIONMULTISIGSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ContractStorageImpl.contract.Call(opts, out, "PERMISSION_MULTISIG_STORAGE")
	return *ret0, err
}

// PERMISSIONMULTISIGSTORAGE is a free data retrieval call binding the contract method 0xf1bb7aad.
//
// Solidity: function PERMISSION_MULTISIG_STORAGE() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplSession) PERMISSIONMULTISIGSTORAGE() (*big.Int, error) {
	return _ContractStorageImpl.Contract.PERMISSIONMULTISIGSTORAGE(&_ContractStorageImpl.CallOpts)
}

// PERMISSIONMULTISIGSTORAGE is a free data retrieval call binding the contract method 0xf1bb7aad.
//
// Solidity: function PERMISSION_MULTISIG_STORAGE() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCallerSession) PERMISSIONMULTISIGSTORAGE() (*big.Int, error) {
	return _ContractStorageImpl.Contract.PERMISSIONMULTISIGSTORAGE(&_ContractStorageImpl.CallOpts)
}

// PERMISSIONROLESSTORAGE is a free data retrieval call binding the contract method 0x12a16032.
//
// Solidity: function PERMISSION_ROLES_STORAGE() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCaller) PERMISSIONROLESSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ContractStorageImpl.contract.Call(opts, out, "PERMISSION_ROLES_STORAGE")
	return *ret0, err
}

// PERMISSIONROLESSTORAGE is a free data retrieval call binding the contract method 0x12a16032.
//
// Solidity: function PERMISSION_ROLES_STORAGE() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplSession) PERMISSIONROLESSTORAGE() (*big.Int, error) {
	return _ContractStorageImpl.Contract.PERMISSIONROLESSTORAGE(&_ContractStorageImpl.CallOpts)
}

// PERMISSIONROLESSTORAGE is a free data retrieval call binding the contract method 0x12a16032.
//
// Solidity: function PERMISSION_ROLES_STORAGE() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCallerSession) PERMISSIONROLESSTORAGE() (*big.Int, error) {
	return _ContractStorageImpl.Contract.PERMISSIONROLESSTORAGE(&_ContractStorageImpl.CallOpts)
}

// ROLEACCESSVERIFIER is a free data retrieval call binding the contract method 0xc5405f92.
//
// Solidity: function ROLE_ACCESS_VERIFIER() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCaller) ROLEACCESSVERIFIER(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ContractStorageImpl.contract.Call(opts, out, "ROLE_ACCESS_VERIFIER")
	return *ret0, err
}

// ROLEACCESSVERIFIER is a free data retrieval call binding the contract method 0xc5405f92.
//
// Solidity: function ROLE_ACCESS_VERIFIER() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplSession) ROLEACCESSVERIFIER() (*big.Int, error) {
	return _ContractStorageImpl.Contract.ROLEACCESSVERIFIER(&_ContractStorageImpl.CallOpts)
}

// ROLEACCESSVERIFIER is a free data retrieval call binding the contract method 0xc5405f92.
//
// Solidity: function ROLE_ACCESS_VERIFIER() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCallerSession) ROLEACCESSVERIFIER() (*big.Int, error) {
	return _ContractStorageImpl.Contract.ROLEACCESSVERIFIER(&_ContractStorageImpl.CallOpts)
}

// ROLEBITBONISSUEVERIFIER is a free data retrieval call binding the contract method 0xf0a012a6.
//
// Solidity: function ROLE_BITBON_ISSUE_VERIFIER() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCaller) ROLEBITBONISSUEVERIFIER(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ContractStorageImpl.contract.Call(opts, out, "ROLE_BITBON_ISSUE_VERIFIER")
	return *ret0, err
}

// ROLEBITBONISSUEVERIFIER is a free data retrieval call binding the contract method 0xf0a012a6.
//
// Solidity: function ROLE_BITBON_ISSUE_VERIFIER() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplSession) ROLEBITBONISSUEVERIFIER() (*big.Int, error) {
	return _ContractStorageImpl.Contract.ROLEBITBONISSUEVERIFIER(&_ContractStorageImpl.CallOpts)
}

// ROLEBITBONISSUEVERIFIER is a free data retrieval call binding the contract method 0xf0a012a6.
//
// Solidity: function ROLE_BITBON_ISSUE_VERIFIER() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCallerSession) ROLEBITBONISSUEVERIFIER() (*big.Int, error) {
	return _ContractStorageImpl.Contract.ROLEBITBONISSUEVERIFIER(&_ContractStorageImpl.CallOpts)
}

// ROLECOMMISSIONVERIFIER is a free data retrieval call binding the contract method 0x8ab3de8c.
//
// Solidity: function ROLE_COMMISSION_VERIFIER() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCaller) ROLECOMMISSIONVERIFIER(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ContractStorageImpl.contract.Call(opts, out, "ROLE_COMMISSION_VERIFIER")
	return *ret0, err
}

// ROLECOMMISSIONVERIFIER is a free data retrieval call binding the contract method 0x8ab3de8c.
//
// Solidity: function ROLE_COMMISSION_VERIFIER() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplSession) ROLECOMMISSIONVERIFIER() (*big.Int, error) {
	return _ContractStorageImpl.Contract.ROLECOMMISSIONVERIFIER(&_ContractStorageImpl.CallOpts)
}

// ROLECOMMISSIONVERIFIER is a free data retrieval call binding the contract method 0x8ab3de8c.
//
// Solidity: function ROLE_COMMISSION_VERIFIER() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCallerSession) ROLECOMMISSIONVERIFIER() (*big.Int, error) {
	return _ContractStorageImpl.Contract.ROLECOMMISSIONVERIFIER(&_ContractStorageImpl.CallOpts)
}

// ROLEDEPLOYADMIN is a free data retrieval call binding the contract method 0x667bf6e5.
//
// Solidity: function ROLE_DEPLOY_ADMIN() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCaller) ROLEDEPLOYADMIN(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ContractStorageImpl.contract.Call(opts, out, "ROLE_DEPLOY_ADMIN")
	return *ret0, err
}

// ROLEDEPLOYADMIN is a free data retrieval call binding the contract method 0x667bf6e5.
//
// Solidity: function ROLE_DEPLOY_ADMIN() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplSession) ROLEDEPLOYADMIN() (*big.Int, error) {
	return _ContractStorageImpl.Contract.ROLEDEPLOYADMIN(&_ContractStorageImpl.CallOpts)
}

// ROLEDEPLOYADMIN is a free data retrieval call binding the contract method 0x667bf6e5.
//
// Solidity: function ROLE_DEPLOY_ADMIN() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCallerSession) ROLEDEPLOYADMIN() (*big.Int, error) {
	return _ContractStorageImpl.Contract.ROLEDEPLOYADMIN(&_ContractStorageImpl.CallOpts)
}

// ROLEPERMISSIONADMIN is a free data retrieval call binding the contract method 0x1b239996.
//
// Solidity: function ROLE_PERMISSION_ADMIN() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCaller) ROLEPERMISSIONADMIN(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ContractStorageImpl.contract.Call(opts, out, "ROLE_PERMISSION_ADMIN")
	return *ret0, err
}

// ROLEPERMISSIONADMIN is a free data retrieval call binding the contract method 0x1b239996.
//
// Solidity: function ROLE_PERMISSION_ADMIN() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplSession) ROLEPERMISSIONADMIN() (*big.Int, error) {
	return _ContractStorageImpl.Contract.ROLEPERMISSIONADMIN(&_ContractStorageImpl.CallOpts)
}

// ROLEPERMISSIONADMIN is a free data retrieval call binding the contract method 0x1b239996.
//
// Solidity: function ROLE_PERMISSION_ADMIN() view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCallerSession) ROLEPERMISSIONADMIN() (*big.Int, error) {
	return _ContractStorageImpl.Contract.ROLEPERMISSIONADMIN(&_ContractStorageImpl.CallOpts)
}

// ContractDeployers is a free data retrieval call binding the contract method 0x1491fa69.
//
// Solidity: function contractDeployers(uint256 ) view returns(address)
func (_ContractStorageImpl *ContractStorageImplCaller) ContractDeployers(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ContractStorageImpl.contract.Call(opts, out, "contractDeployers", arg0)
	return *ret0, err
}

// ContractDeployers is a free data retrieval call binding the contract method 0x1491fa69.
//
// Solidity: function contractDeployers(uint256 ) view returns(address)
func (_ContractStorageImpl *ContractStorageImplSession) ContractDeployers(arg0 *big.Int) (common.Address, error) {
	return _ContractStorageImpl.Contract.ContractDeployers(&_ContractStorageImpl.CallOpts, arg0)
}

// ContractDeployers is a free data retrieval call binding the contract method 0x1491fa69.
//
// Solidity: function contractDeployers(uint256 ) view returns(address)
func (_ContractStorageImpl *ContractStorageImplCallerSession) ContractDeployers(arg0 *big.Int) (common.Address, error) {
	return _ContractStorageImpl.Contract.ContractDeployers(&_ContractStorageImpl.CallOpts, arg0)
}

// Contracts is a free data retrieval call binding the contract method 0x474da79a.
//
// Solidity: function contracts(uint256 ) view returns(address)
func (_ContractStorageImpl *ContractStorageImplCaller) Contracts(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ContractStorageImpl.contract.Call(opts, out, "contracts", arg0)
	return *ret0, err
}

// Contracts is a free data retrieval call binding the contract method 0x474da79a.
//
// Solidity: function contracts(uint256 ) view returns(address)
func (_ContractStorageImpl *ContractStorageImplSession) Contracts(arg0 *big.Int) (common.Address, error) {
	return _ContractStorageImpl.Contract.Contracts(&_ContractStorageImpl.CallOpts, arg0)
}

// Contracts is a free data retrieval call binding the contract method 0x474da79a.
//
// Solidity: function contracts(uint256 ) view returns(address)
func (_ContractStorageImpl *ContractStorageImplCallerSession) Contracts(arg0 *big.Int) (common.Address, error) {
	return _ContractStorageImpl.Contract.Contracts(&_ContractStorageImpl.CallOpts, arg0)
}

// ContractsTime is a free data retrieval call binding the contract method 0x566e4314.
//
// Solidity: function contractsTime(uint256 ) view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCaller) ContractsTime(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ContractStorageImpl.contract.Call(opts, out, "contractsTime", arg0)
	return *ret0, err
}

// ContractsTime is a free data retrieval call binding the contract method 0x566e4314.
//
// Solidity: function contractsTime(uint256 ) view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplSession) ContractsTime(arg0 *big.Int) (*big.Int, error) {
	return _ContractStorageImpl.Contract.ContractsTime(&_ContractStorageImpl.CallOpts, arg0)
}

// ContractsTime is a free data retrieval call binding the contract method 0x566e4314.
//
// Solidity: function contractsTime(uint256 ) view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCallerSession) ContractsTime(arg0 *big.Int) (*big.Int, error) {
	return _ContractStorageImpl.Contract.ContractsTime(&_ContractStorageImpl.CallOpts, arg0)
}

// GetAccess is a free data retrieval call binding the contract method 0xf7ca8431.
//
// Solidity: function getAccess() view returns(address)
func (_ContractStorageImpl *ContractStorageImplCaller) GetAccess(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ContractStorageImpl.contract.Call(opts, out, "getAccess")
	return *ret0, err
}

// GetAccess is a free data retrieval call binding the contract method 0xf7ca8431.
//
// Solidity: function getAccess() view returns(address)
func (_ContractStorageImpl *ContractStorageImplSession) GetAccess() (common.Address, error) {
	return _ContractStorageImpl.Contract.GetAccess(&_ContractStorageImpl.CallOpts)
}

// GetAccess is a free data retrieval call binding the contract method 0xf7ca8431.
//
// Solidity: function getAccess() view returns(address)
func (_ContractStorageImpl *ContractStorageImplCallerSession) GetAccess() (common.Address, error) {
	return _ContractStorageImpl.Contract.GetAccess(&_ContractStorageImpl.CallOpts)
}

// GetAccessStorage is a free data retrieval call binding the contract method 0x0d047185.
//
// Solidity: function getAccessStorage() view returns(address)
func (_ContractStorageImpl *ContractStorageImplCaller) GetAccessStorage(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ContractStorageImpl.contract.Call(opts, out, "getAccessStorage")
	return *ret0, err
}

// GetAccessStorage is a free data retrieval call binding the contract method 0x0d047185.
//
// Solidity: function getAccessStorage() view returns(address)
func (_ContractStorageImpl *ContractStorageImplSession) GetAccessStorage() (common.Address, error) {
	return _ContractStorageImpl.Contract.GetAccessStorage(&_ContractStorageImpl.CallOpts)
}

// GetAccessStorage is a free data retrieval call binding the contract method 0x0d047185.
//
// Solidity: function getAccessStorage() view returns(address)
func (_ContractStorageImpl *ContractStorageImplCallerSession) GetAccessStorage() (common.Address, error) {
	return _ContractStorageImpl.Contract.GetAccessStorage(&_ContractStorageImpl.CallOpts)
}

// GetAdminable is a free data retrieval call binding the contract method 0x23886180.
//
// Solidity: function getAdminable() view returns(address)
func (_ContractStorageImpl *ContractStorageImplCaller) GetAdminable(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ContractStorageImpl.contract.Call(opts, out, "getAdminable")
	return *ret0, err
}

// GetAdminable is a free data retrieval call binding the contract method 0x23886180.
//
// Solidity: function getAdminable() view returns(address)
func (_ContractStorageImpl *ContractStorageImplSession) GetAdminable() (common.Address, error) {
	return _ContractStorageImpl.Contract.GetAdminable(&_ContractStorageImpl.CallOpts)
}

// GetAdminable is a free data retrieval call binding the contract method 0x23886180.
//
// Solidity: function getAdminable() view returns(address)
func (_ContractStorageImpl *ContractStorageImplCallerSession) GetAdminable() (common.Address, error) {
	return _ContractStorageImpl.Contract.GetAdminable(&_ContractStorageImpl.CallOpts)
}

// GetAssetbox is a free data retrieval call binding the contract method 0x5285aa82.
//
// Solidity: function getAssetbox() view returns(address)
func (_ContractStorageImpl *ContractStorageImplCaller) GetAssetbox(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ContractStorageImpl.contract.Call(opts, out, "getAssetbox")
	return *ret0, err
}

// GetAssetbox is a free data retrieval call binding the contract method 0x5285aa82.
//
// Solidity: function getAssetbox() view returns(address)
func (_ContractStorageImpl *ContractStorageImplSession) GetAssetbox() (common.Address, error) {
	return _ContractStorageImpl.Contract.GetAssetbox(&_ContractStorageImpl.CallOpts)
}

// GetAssetbox is a free data retrieval call binding the contract method 0x5285aa82.
//
// Solidity: function getAssetbox() view returns(address)
func (_ContractStorageImpl *ContractStorageImplCallerSession) GetAssetbox() (common.Address, error) {
	return _ContractStorageImpl.Contract.GetAssetbox(&_ContractStorageImpl.CallOpts)
}

// GetAssetboxInfo is a free data retrieval call binding the contract method 0x33e1a11f.
//
// Solidity: function getAssetboxInfo() view returns(address)
func (_ContractStorageImpl *ContractStorageImplCaller) GetAssetboxInfo(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ContractStorageImpl.contract.Call(opts, out, "getAssetboxInfo")
	return *ret0, err
}

// GetAssetboxInfo is a free data retrieval call binding the contract method 0x33e1a11f.
//
// Solidity: function getAssetboxInfo() view returns(address)
func (_ContractStorageImpl *ContractStorageImplSession) GetAssetboxInfo() (common.Address, error) {
	return _ContractStorageImpl.Contract.GetAssetboxInfo(&_ContractStorageImpl.CallOpts)
}

// GetAssetboxInfo is a free data retrieval call binding the contract method 0x33e1a11f.
//
// Solidity: function getAssetboxInfo() view returns(address)
func (_ContractStorageImpl *ContractStorageImplCallerSession) GetAssetboxInfo() (common.Address, error) {
	return _ContractStorageImpl.Contract.GetAssetboxInfo(&_ContractStorageImpl.CallOpts)
}

// GetAssetboxStorage is a free data retrieval call binding the contract method 0x4b488eb0.
//
// Solidity: function getAssetboxStorage() view returns(address)
func (_ContractStorageImpl *ContractStorageImplCaller) GetAssetboxStorage(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ContractStorageImpl.contract.Call(opts, out, "getAssetboxStorage")
	return *ret0, err
}

// GetAssetboxStorage is a free data retrieval call binding the contract method 0x4b488eb0.
//
// Solidity: function getAssetboxStorage() view returns(address)
func (_ContractStorageImpl *ContractStorageImplSession) GetAssetboxStorage() (common.Address, error) {
	return _ContractStorageImpl.Contract.GetAssetboxStorage(&_ContractStorageImpl.CallOpts)
}

// GetAssetboxStorage is a free data retrieval call binding the contract method 0x4b488eb0.
//
// Solidity: function getAssetboxStorage() view returns(address)
func (_ContractStorageImpl *ContractStorageImplCallerSession) GetAssetboxStorage() (common.Address, error) {
	return _ContractStorageImpl.Contract.GetAssetboxStorage(&_ContractStorageImpl.CallOpts)
}

// GetBalanceGate is a free data retrieval call binding the contract method 0x4267dd3b.
//
// Solidity: function getBalanceGate() view returns(address)
func (_ContractStorageImpl *ContractStorageImplCaller) GetBalanceGate(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ContractStorageImpl.contract.Call(opts, out, "getBalanceGate")
	return *ret0, err
}

// GetBalanceGate is a free data retrieval call binding the contract method 0x4267dd3b.
//
// Solidity: function getBalanceGate() view returns(address)
func (_ContractStorageImpl *ContractStorageImplSession) GetBalanceGate() (common.Address, error) {
	return _ContractStorageImpl.Contract.GetBalanceGate(&_ContractStorageImpl.CallOpts)
}

// GetBalanceGate is a free data retrieval call binding the contract method 0x4267dd3b.
//
// Solidity: function getBalanceGate() view returns(address)
func (_ContractStorageImpl *ContractStorageImplCallerSession) GetBalanceGate() (common.Address, error) {
	return _ContractStorageImpl.Contract.GetBalanceGate(&_ContractStorageImpl.CallOpts)
}

// GetBitbon is a free data retrieval call binding the contract method 0xc17e0252.
//
// Solidity: function getBitbon() view returns(address)
func (_ContractStorageImpl *ContractStorageImplCaller) GetBitbon(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ContractStorageImpl.contract.Call(opts, out, "getBitbon")
	return *ret0, err
}

// GetBitbon is a free data retrieval call binding the contract method 0xc17e0252.
//
// Solidity: function getBitbon() view returns(address)
func (_ContractStorageImpl *ContractStorageImplSession) GetBitbon() (common.Address, error) {
	return _ContractStorageImpl.Contract.GetBitbon(&_ContractStorageImpl.CallOpts)
}

// GetBitbon is a free data retrieval call binding the contract method 0xc17e0252.
//
// Solidity: function getBitbon() view returns(address)
func (_ContractStorageImpl *ContractStorageImplCallerSession) GetBitbon() (common.Address, error) {
	return _ContractStorageImpl.Contract.GetBitbon(&_ContractStorageImpl.CallOpts)
}

// GetBitbonStorage is a free data retrieval call binding the contract method 0xf46b7e96.
//
// Solidity: function getBitbonStorage() view returns(address)
func (_ContractStorageImpl *ContractStorageImplCaller) GetBitbonStorage(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ContractStorageImpl.contract.Call(opts, out, "getBitbonStorage")
	return *ret0, err
}

// GetBitbonStorage is a free data retrieval call binding the contract method 0xf46b7e96.
//
// Solidity: function getBitbonStorage() view returns(address)
func (_ContractStorageImpl *ContractStorageImplSession) GetBitbonStorage() (common.Address, error) {
	return _ContractStorageImpl.Contract.GetBitbonStorage(&_ContractStorageImpl.CallOpts)
}

// GetBitbonStorage is a free data retrieval call binding the contract method 0xf46b7e96.
//
// Solidity: function getBitbonStorage() view returns(address)
func (_ContractStorageImpl *ContractStorageImplCallerSession) GetBitbonStorage() (common.Address, error) {
	return _ContractStorageImpl.Contract.GetBitbonStorage(&_ContractStorageImpl.CallOpts)
}

// GetBitbonSupport is a free data retrieval call binding the contract method 0xfecf10bb.
//
// Solidity: function getBitbonSupport() view returns(address)
func (_ContractStorageImpl *ContractStorageImplCaller) GetBitbonSupport(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ContractStorageImpl.contract.Call(opts, out, "getBitbonSupport")
	return *ret0, err
}

// GetBitbonSupport is a free data retrieval call binding the contract method 0xfecf10bb.
//
// Solidity: function getBitbonSupport() view returns(address)
func (_ContractStorageImpl *ContractStorageImplSession) GetBitbonSupport() (common.Address, error) {
	return _ContractStorageImpl.Contract.GetBitbonSupport(&_ContractStorageImpl.CallOpts)
}

// GetBitbonSupport is a free data retrieval call binding the contract method 0xfecf10bb.
//
// Solidity: function getBitbonSupport() view returns(address)
func (_ContractStorageImpl *ContractStorageImplCallerSession) GetBitbonSupport() (common.Address, error) {
	return _ContractStorageImpl.Contract.GetBitbonSupport(&_ContractStorageImpl.CallOpts)
}

// GetBonBalanceGate is a free data retrieval call binding the contract method 0x0835e907.
//
// Solidity: function getBonBalanceGate() view returns(address)
func (_ContractStorageImpl *ContractStorageImplCaller) GetBonBalanceGate(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ContractStorageImpl.contract.Call(opts, out, "getBonBalanceGate")
	return *ret0, err
}

// GetBonBalanceGate is a free data retrieval call binding the contract method 0x0835e907.
//
// Solidity: function getBonBalanceGate() view returns(address)
func (_ContractStorageImpl *ContractStorageImplSession) GetBonBalanceGate() (common.Address, error) {
	return _ContractStorageImpl.Contract.GetBonBalanceGate(&_ContractStorageImpl.CallOpts)
}

// GetBonBalanceGate is a free data retrieval call binding the contract method 0x0835e907.
//
// Solidity: function getBonBalanceGate() view returns(address)
func (_ContractStorageImpl *ContractStorageImplCallerSession) GetBonBalanceGate() (common.Address, error) {
	return _ContractStorageImpl.Contract.GetBonBalanceGate(&_ContractStorageImpl.CallOpts)
}

// GetBonTransferGate is a free data retrieval call binding the contract method 0x132592a2.
//
// Solidity: function getBonTransferGate() view returns(address)
func (_ContractStorageImpl *ContractStorageImplCaller) GetBonTransferGate(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ContractStorageImpl.contract.Call(opts, out, "getBonTransferGate")
	return *ret0, err
}

// GetBonTransferGate is a free data retrieval call binding the contract method 0x132592a2.
//
// Solidity: function getBonTransferGate() view returns(address)
func (_ContractStorageImpl *ContractStorageImplSession) GetBonTransferGate() (common.Address, error) {
	return _ContractStorageImpl.Contract.GetBonTransferGate(&_ContractStorageImpl.CallOpts)
}

// GetBonTransferGate is a free data retrieval call binding the contract method 0x132592a2.
//
// Solidity: function getBonTransferGate() view returns(address)
func (_ContractStorageImpl *ContractStorageImplCallerSession) GetBonTransferGate() (common.Address, error) {
	return _ContractStorageImpl.Contract.GetBonTransferGate(&_ContractStorageImpl.CallOpts)
}

// GetContractAddress is a free data retrieval call binding the contract method 0xaefa7d98.
//
// Solidity: function getContractAddress(uint256 contractType) view returns(address)
func (_ContractStorageImpl *ContractStorageImplCaller) GetContractAddress(opts *bind.CallOpts, contractType *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ContractStorageImpl.contract.Call(opts, out, "getContractAddress", contractType)
	return *ret0, err
}

// GetContractAddress is a free data retrieval call binding the contract method 0xaefa7d98.
//
// Solidity: function getContractAddress(uint256 contractType) view returns(address)
func (_ContractStorageImpl *ContractStorageImplSession) GetContractAddress(contractType *big.Int) (common.Address, error) {
	return _ContractStorageImpl.Contract.GetContractAddress(&_ContractStorageImpl.CallOpts, contractType)
}

// GetContractAddress is a free data retrieval call binding the contract method 0xaefa7d98.
//
// Solidity: function getContractAddress(uint256 contractType) view returns(address)
func (_ContractStorageImpl *ContractStorageImplCallerSession) GetContractAddress(contractType *big.Int) (common.Address, error) {
	return _ContractStorageImpl.Contract.GetContractAddress(&_ContractStorageImpl.CallOpts, contractType)
}

// GetContractAddresses is a free data retrieval call binding the contract method 0x8e04500f.
//
// Solidity: function getContractAddresses(uint256[] contractTypes) view returns(address[])
func (_ContractStorageImpl *ContractStorageImplCaller) GetContractAddresses(opts *bind.CallOpts, contractTypes []*big.Int) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _ContractStorageImpl.contract.Call(opts, out, "getContractAddresses", contractTypes)
	return *ret0, err
}

// GetContractAddresses is a free data retrieval call binding the contract method 0x8e04500f.
//
// Solidity: function getContractAddresses(uint256[] contractTypes) view returns(address[])
func (_ContractStorageImpl *ContractStorageImplSession) GetContractAddresses(contractTypes []*big.Int) ([]common.Address, error) {
	return _ContractStorageImpl.Contract.GetContractAddresses(&_ContractStorageImpl.CallOpts, contractTypes)
}

// GetContractAddresses is a free data retrieval call binding the contract method 0x8e04500f.
//
// Solidity: function getContractAddresses(uint256[] contractTypes) view returns(address[])
func (_ContractStorageImpl *ContractStorageImplCallerSession) GetContractAddresses(contractTypes []*big.Int) ([]common.Address, error) {
	return _ContractStorageImpl.Contract.GetContractAddresses(&_ContractStorageImpl.CallOpts, contractTypes)
}

// GetContractDeployTime is a free data retrieval call binding the contract method 0x012753ac.
//
// Solidity: function getContractDeployTime(uint256 contractType) view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCaller) GetContractDeployTime(opts *bind.CallOpts, contractType *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ContractStorageImpl.contract.Call(opts, out, "getContractDeployTime", contractType)
	return *ret0, err
}

// GetContractDeployTime is a free data retrieval call binding the contract method 0x012753ac.
//
// Solidity: function getContractDeployTime(uint256 contractType) view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplSession) GetContractDeployTime(contractType *big.Int) (*big.Int, error) {
	return _ContractStorageImpl.Contract.GetContractDeployTime(&_ContractStorageImpl.CallOpts, contractType)
}

// GetContractDeployTime is a free data retrieval call binding the contract method 0x012753ac.
//
// Solidity: function getContractDeployTime(uint256 contractType) view returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCallerSession) GetContractDeployTime(contractType *big.Int) (*big.Int, error) {
	return _ContractStorageImpl.Contract.GetContractDeployTime(&_ContractStorageImpl.CallOpts, contractType)
}

// GetContractDeployer is a free data retrieval call binding the contract method 0xf9bede86.
//
// Solidity: function getContractDeployer(uint256 contractType) view returns(address)
func (_ContractStorageImpl *ContractStorageImplCaller) GetContractDeployer(opts *bind.CallOpts, contractType *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ContractStorageImpl.contract.Call(opts, out, "getContractDeployer", contractType)
	return *ret0, err
}

// GetContractDeployer is a free data retrieval call binding the contract method 0xf9bede86.
//
// Solidity: function getContractDeployer(uint256 contractType) view returns(address)
func (_ContractStorageImpl *ContractStorageImplSession) GetContractDeployer(contractType *big.Int) (common.Address, error) {
	return _ContractStorageImpl.Contract.GetContractDeployer(&_ContractStorageImpl.CallOpts, contractType)
}

// GetContractDeployer is a free data retrieval call binding the contract method 0xf9bede86.
//
// Solidity: function getContractDeployer(uint256 contractType) view returns(address)
func (_ContractStorageImpl *ContractStorageImplCallerSession) GetContractDeployer(contractType *big.Int) (common.Address, error) {
	return _ContractStorageImpl.Contract.GetContractDeployer(&_ContractStorageImpl.CallOpts, contractType)
}

// GetDistributionStorage is a free data retrieval call binding the contract method 0x16ccb1ac.
//
// Solidity: function getDistributionStorage() view returns(address)
func (_ContractStorageImpl *ContractStorageImplCaller) GetDistributionStorage(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ContractStorageImpl.contract.Call(opts, out, "getDistributionStorage")
	return *ret0, err
}

// GetDistributionStorage is a free data retrieval call binding the contract method 0x16ccb1ac.
//
// Solidity: function getDistributionStorage() view returns(address)
func (_ContractStorageImpl *ContractStorageImplSession) GetDistributionStorage() (common.Address, error) {
	return _ContractStorageImpl.Contract.GetDistributionStorage(&_ContractStorageImpl.CallOpts)
}

// GetDistributionStorage is a free data retrieval call binding the contract method 0x16ccb1ac.
//
// Solidity: function getDistributionStorage() view returns(address)
func (_ContractStorageImpl *ContractStorageImplCallerSession) GetDistributionStorage() (common.Address, error) {
	return _ContractStorageImpl.Contract.GetDistributionStorage(&_ContractStorageImpl.CallOpts)
}

// GetExchange is a free data retrieval call binding the contract method 0xf807cd22.
//
// Solidity: function getExchange() view returns(address)
func (_ContractStorageImpl *ContractStorageImplCaller) GetExchange(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ContractStorageImpl.contract.Call(opts, out, "getExchange")
	return *ret0, err
}

// GetExchange is a free data retrieval call binding the contract method 0xf807cd22.
//
// Solidity: function getExchange() view returns(address)
func (_ContractStorageImpl *ContractStorageImplSession) GetExchange() (common.Address, error) {
	return _ContractStorageImpl.Contract.GetExchange(&_ContractStorageImpl.CallOpts)
}

// GetExchange is a free data retrieval call binding the contract method 0xf807cd22.
//
// Solidity: function getExchange() view returns(address)
func (_ContractStorageImpl *ContractStorageImplCallerSession) GetExchange() (common.Address, error) {
	return _ContractStorageImpl.Contract.GetExchange(&_ContractStorageImpl.CallOpts)
}

// GetExchangeStorage is a free data retrieval call binding the contract method 0xa765f617.
//
// Solidity: function getExchangeStorage() view returns(address)
func (_ContractStorageImpl *ContractStorageImplCaller) GetExchangeStorage(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ContractStorageImpl.contract.Call(opts, out, "getExchangeStorage")
	return *ret0, err
}

// GetExchangeStorage is a free data retrieval call binding the contract method 0xa765f617.
//
// Solidity: function getExchangeStorage() view returns(address)
func (_ContractStorageImpl *ContractStorageImplSession) GetExchangeStorage() (common.Address, error) {
	return _ContractStorageImpl.Contract.GetExchangeStorage(&_ContractStorageImpl.CallOpts)
}

// GetExchangeStorage is a free data retrieval call binding the contract method 0xa765f617.
//
// Solidity: function getExchangeStorage() view returns(address)
func (_ContractStorageImpl *ContractStorageImplCallerSession) GetExchangeStorage() (common.Address, error) {
	return _ContractStorageImpl.Contract.GetExchangeStorage(&_ContractStorageImpl.CallOpts)
}

// GetFeeStorage is a free data retrieval call binding the contract method 0x343ddaa3.
//
// Solidity: function getFeeStorage() view returns(address)
func (_ContractStorageImpl *ContractStorageImplCaller) GetFeeStorage(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ContractStorageImpl.contract.Call(opts, out, "getFeeStorage")
	return *ret0, err
}

// GetFeeStorage is a free data retrieval call binding the contract method 0x343ddaa3.
//
// Solidity: function getFeeStorage() view returns(address)
func (_ContractStorageImpl *ContractStorageImplSession) GetFeeStorage() (common.Address, error) {
	return _ContractStorageImpl.Contract.GetFeeStorage(&_ContractStorageImpl.CallOpts)
}

// GetFeeStorage is a free data retrieval call binding the contract method 0x343ddaa3.
//
// Solidity: function getFeeStorage() view returns(address)
func (_ContractStorageImpl *ContractStorageImplCallerSession) GetFeeStorage() (common.Address, error) {
	return _ContractStorageImpl.Contract.GetFeeStorage(&_ContractStorageImpl.CallOpts)
}

// GetMiningAgentStorage is a free data retrieval call binding the contract method 0x4c9a68b4.
//
// Solidity: function getMiningAgentStorage() view returns(address)
func (_ContractStorageImpl *ContractStorageImplCaller) GetMiningAgentStorage(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ContractStorageImpl.contract.Call(opts, out, "getMiningAgentStorage")
	return *ret0, err
}

// GetMiningAgentStorage is a free data retrieval call binding the contract method 0x4c9a68b4.
//
// Solidity: function getMiningAgentStorage() view returns(address)
func (_ContractStorageImpl *ContractStorageImplSession) GetMiningAgentStorage() (common.Address, error) {
	return _ContractStorageImpl.Contract.GetMiningAgentStorage(&_ContractStorageImpl.CallOpts)
}

// GetMiningAgentStorage is a free data retrieval call binding the contract method 0x4c9a68b4.
//
// Solidity: function getMiningAgentStorage() view returns(address)
func (_ContractStorageImpl *ContractStorageImplCallerSession) GetMiningAgentStorage() (common.Address, error) {
	return _ContractStorageImpl.Contract.GetMiningAgentStorage(&_ContractStorageImpl.CallOpts)
}

// GetMsbon is a free data retrieval call binding the contract method 0xc8ddbddc.
//
// Solidity: function getMsbon() view returns(address)
func (_ContractStorageImpl *ContractStorageImplCaller) GetMsbon(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ContractStorageImpl.contract.Call(opts, out, "getMsbon")
	return *ret0, err
}

// GetMsbon is a free data retrieval call binding the contract method 0xc8ddbddc.
//
// Solidity: function getMsbon() view returns(address)
func (_ContractStorageImpl *ContractStorageImplSession) GetMsbon() (common.Address, error) {
	return _ContractStorageImpl.Contract.GetMsbon(&_ContractStorageImpl.CallOpts)
}

// GetMsbon is a free data retrieval call binding the contract method 0xc8ddbddc.
//
// Solidity: function getMsbon() view returns(address)
func (_ContractStorageImpl *ContractStorageImplCallerSession) GetMsbon() (common.Address, error) {
	return _ContractStorageImpl.Contract.GetMsbon(&_ContractStorageImpl.CallOpts)
}

// GetMsbonStorage is a free data retrieval call binding the contract method 0x3bc37b68.
//
// Solidity: function getMsbonStorage() view returns(address)
func (_ContractStorageImpl *ContractStorageImplCaller) GetMsbonStorage(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ContractStorageImpl.contract.Call(opts, out, "getMsbonStorage")
	return *ret0, err
}

// GetMsbonStorage is a free data retrieval call binding the contract method 0x3bc37b68.
//
// Solidity: function getMsbonStorage() view returns(address)
func (_ContractStorageImpl *ContractStorageImplSession) GetMsbonStorage() (common.Address, error) {
	return _ContractStorageImpl.Contract.GetMsbonStorage(&_ContractStorageImpl.CallOpts)
}

// GetMsbonStorage is a free data retrieval call binding the contract method 0x3bc37b68.
//
// Solidity: function getMsbonStorage() view returns(address)
func (_ContractStorageImpl *ContractStorageImplCallerSession) GetMsbonStorage() (common.Address, error) {
	return _ContractStorageImpl.Contract.GetMsbonStorage(&_ContractStorageImpl.CallOpts)
}

// GetMultisigStorage is a free data retrieval call binding the contract method 0x3c3ddbd1.
//
// Solidity: function getMultisigStorage() view returns(address)
func (_ContractStorageImpl *ContractStorageImplCaller) GetMultisigStorage(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ContractStorageImpl.contract.Call(opts, out, "getMultisigStorage")
	return *ret0, err
}

// GetMultisigStorage is a free data retrieval call binding the contract method 0x3c3ddbd1.
//
// Solidity: function getMultisigStorage() view returns(address)
func (_ContractStorageImpl *ContractStorageImplSession) GetMultisigStorage() (common.Address, error) {
	return _ContractStorageImpl.Contract.GetMultisigStorage(&_ContractStorageImpl.CallOpts)
}

// GetMultisigStorage is a free data retrieval call binding the contract method 0x3c3ddbd1.
//
// Solidity: function getMultisigStorage() view returns(address)
func (_ContractStorageImpl *ContractStorageImplCallerSession) GetMultisigStorage() (common.Address, error) {
	return _ContractStorageImpl.Contract.GetMultisigStorage(&_ContractStorageImpl.CallOpts)
}

// GetNewContactStorageAddress is a free data retrieval call binding the contract method 0x3dcf05c0.
//
// Solidity: function getNewContactStorageAddress() view returns(address)
func (_ContractStorageImpl *ContractStorageImplCaller) GetNewContactStorageAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ContractStorageImpl.contract.Call(opts, out, "getNewContactStorageAddress")
	return *ret0, err
}

// GetNewContactStorageAddress is a free data retrieval call binding the contract method 0x3dcf05c0.
//
// Solidity: function getNewContactStorageAddress() view returns(address)
func (_ContractStorageImpl *ContractStorageImplSession) GetNewContactStorageAddress() (common.Address, error) {
	return _ContractStorageImpl.Contract.GetNewContactStorageAddress(&_ContractStorageImpl.CallOpts)
}

// GetNewContactStorageAddress is a free data retrieval call binding the contract method 0x3dcf05c0.
//
// Solidity: function getNewContactStorageAddress() view returns(address)
func (_ContractStorageImpl *ContractStorageImplCallerSession) GetNewContactStorageAddress() (common.Address, error) {
	return _ContractStorageImpl.Contract.GetNewContactStorageAddress(&_ContractStorageImpl.CallOpts)
}

// GetOtc is a free data retrieval call binding the contract method 0xb89ed48f.
//
// Solidity: function getOtc() view returns(address)
func (_ContractStorageImpl *ContractStorageImplCaller) GetOtc(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ContractStorageImpl.contract.Call(opts, out, "getOtc")
	return *ret0, err
}

// GetOtc is a free data retrieval call binding the contract method 0xb89ed48f.
//
// Solidity: function getOtc() view returns(address)
func (_ContractStorageImpl *ContractStorageImplSession) GetOtc() (common.Address, error) {
	return _ContractStorageImpl.Contract.GetOtc(&_ContractStorageImpl.CallOpts)
}

// GetOtc is a free data retrieval call binding the contract method 0xb89ed48f.
//
// Solidity: function getOtc() view returns(address)
func (_ContractStorageImpl *ContractStorageImplCallerSession) GetOtc() (common.Address, error) {
	return _ContractStorageImpl.Contract.GetOtc(&_ContractStorageImpl.CallOpts)
}

// GetOtcStorage is a free data retrieval call binding the contract method 0xb8101948.
//
// Solidity: function getOtcStorage() view returns(address)
func (_ContractStorageImpl *ContractStorageImplCaller) GetOtcStorage(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ContractStorageImpl.contract.Call(opts, out, "getOtcStorage")
	return *ret0, err
}

// GetOtcStorage is a free data retrieval call binding the contract method 0xb8101948.
//
// Solidity: function getOtcStorage() view returns(address)
func (_ContractStorageImpl *ContractStorageImplSession) GetOtcStorage() (common.Address, error) {
	return _ContractStorageImpl.Contract.GetOtcStorage(&_ContractStorageImpl.CallOpts)
}

// GetOtcStorage is a free data retrieval call binding the contract method 0xb8101948.
//
// Solidity: function getOtcStorage() view returns(address)
func (_ContractStorageImpl *ContractStorageImplCallerSession) GetOtcStorage() (common.Address, error) {
	return _ContractStorageImpl.Contract.GetOtcStorage(&_ContractStorageImpl.CallOpts)
}

// GetPreviousContractStorageAddress is a free data retrieval call binding the contract method 0xbd5f8e9b.
//
// Solidity: function getPreviousContractStorageAddress() view returns(address)
func (_ContractStorageImpl *ContractStorageImplCaller) GetPreviousContractStorageAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ContractStorageImpl.contract.Call(opts, out, "getPreviousContractStorageAddress")
	return *ret0, err
}

// GetPreviousContractStorageAddress is a free data retrieval call binding the contract method 0xbd5f8e9b.
//
// Solidity: function getPreviousContractStorageAddress() view returns(address)
func (_ContractStorageImpl *ContractStorageImplSession) GetPreviousContractStorageAddress() (common.Address, error) {
	return _ContractStorageImpl.Contract.GetPreviousContractStorageAddress(&_ContractStorageImpl.CallOpts)
}

// GetPreviousContractStorageAddress is a free data retrieval call binding the contract method 0xbd5f8e9b.
//
// Solidity: function getPreviousContractStorageAddress() view returns(address)
func (_ContractStorageImpl *ContractStorageImplCallerSession) GetPreviousContractStorageAddress() (common.Address, error) {
	return _ContractStorageImpl.Contract.GetPreviousContractStorageAddress(&_ContractStorageImpl.CallOpts)
}

// GetReservedAliasStorage is a free data retrieval call binding the contract method 0x6c62e96f.
//
// Solidity: function getReservedAliasStorage() view returns(address)
func (_ContractStorageImpl *ContractStorageImplCaller) GetReservedAliasStorage(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ContractStorageImpl.contract.Call(opts, out, "getReservedAliasStorage")
	return *ret0, err
}

// GetReservedAliasStorage is a free data retrieval call binding the contract method 0x6c62e96f.
//
// Solidity: function getReservedAliasStorage() view returns(address)
func (_ContractStorageImpl *ContractStorageImplSession) GetReservedAliasStorage() (common.Address, error) {
	return _ContractStorageImpl.Contract.GetReservedAliasStorage(&_ContractStorageImpl.CallOpts)
}

// GetReservedAliasStorage is a free data retrieval call binding the contract method 0x6c62e96f.
//
// Solidity: function getReservedAliasStorage() view returns(address)
func (_ContractStorageImpl *ContractStorageImplCallerSession) GetReservedAliasStorage() (common.Address, error) {
	return _ContractStorageImpl.Contract.GetReservedAliasStorage(&_ContractStorageImpl.CallOpts)
}

// GetSafeTransferStorage is a free data retrieval call binding the contract method 0xc31fdffe.
//
// Solidity: function getSafeTransferStorage() view returns(address)
func (_ContractStorageImpl *ContractStorageImplCaller) GetSafeTransferStorage(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ContractStorageImpl.contract.Call(opts, out, "getSafeTransferStorage")
	return *ret0, err
}

// GetSafeTransferStorage is a free data retrieval call binding the contract method 0xc31fdffe.
//
// Solidity: function getSafeTransferStorage() view returns(address)
func (_ContractStorageImpl *ContractStorageImplSession) GetSafeTransferStorage() (common.Address, error) {
	return _ContractStorageImpl.Contract.GetSafeTransferStorage(&_ContractStorageImpl.CallOpts)
}

// GetSafeTransferStorage is a free data retrieval call binding the contract method 0xc31fdffe.
//
// Solidity: function getSafeTransferStorage() view returns(address)
func (_ContractStorageImpl *ContractStorageImplCallerSession) GetSafeTransferStorage() (common.Address, error) {
	return _ContractStorageImpl.Contract.GetSafeTransferStorage(&_ContractStorageImpl.CallOpts)
}

// GetThisContractIndex is a free data retrieval call binding the contract method 0x70c6d367.
//
// Solidity: function getThisContractIndex() pure returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCaller) GetThisContractIndex(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ContractStorageImpl.contract.Call(opts, out, "getThisContractIndex")
	return *ret0, err
}

// GetThisContractIndex is a free data retrieval call binding the contract method 0x70c6d367.
//
// Solidity: function getThisContractIndex() pure returns(uint256)
func (_ContractStorageImpl *ContractStorageImplSession) GetThisContractIndex() (*big.Int, error) {
	return _ContractStorageImpl.Contract.GetThisContractIndex(&_ContractStorageImpl.CallOpts)
}

// GetThisContractIndex is a free data retrieval call binding the contract method 0x70c6d367.
//
// Solidity: function getThisContractIndex() pure returns(uint256)
func (_ContractStorageImpl *ContractStorageImplCallerSession) GetThisContractIndex() (*big.Int, error) {
	return _ContractStorageImpl.Contract.GetThisContractIndex(&_ContractStorageImpl.CallOpts)
}

// GetTransferGate is a free data retrieval call binding the contract method 0xe81dba9d.
//
// Solidity: function getTransferGate() view returns(address)
func (_ContractStorageImpl *ContractStorageImplCaller) GetTransferGate(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ContractStorageImpl.contract.Call(opts, out, "getTransferGate")
	return *ret0, err
}

// GetTransferGate is a free data retrieval call binding the contract method 0xe81dba9d.
//
// Solidity: function getTransferGate() view returns(address)
func (_ContractStorageImpl *ContractStorageImplSession) GetTransferGate() (common.Address, error) {
	return _ContractStorageImpl.Contract.GetTransferGate(&_ContractStorageImpl.CallOpts)
}

// GetTransferGate is a free data retrieval call binding the contract method 0xe81dba9d.
//
// Solidity: function getTransferGate() view returns(address)
func (_ContractStorageImpl *ContractStorageImplCallerSession) GetTransferGate() (common.Address, error) {
	return _ContractStorageImpl.Contract.GetTransferGate(&_ContractStorageImpl.CallOpts)
}

// NewContactStorage is a free data retrieval call binding the contract method 0xe5934630.
//
// Solidity: function newContactStorage() view returns(address)
func (_ContractStorageImpl *ContractStorageImplCaller) NewContactStorage(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ContractStorageImpl.contract.Call(opts, out, "newContactStorage")
	return *ret0, err
}

// NewContactStorage is a free data retrieval call binding the contract method 0xe5934630.
//
// Solidity: function newContactStorage() view returns(address)
func (_ContractStorageImpl *ContractStorageImplSession) NewContactStorage() (common.Address, error) {
	return _ContractStorageImpl.Contract.NewContactStorage(&_ContractStorageImpl.CallOpts)
}

// NewContactStorage is a free data retrieval call binding the contract method 0xe5934630.
//
// Solidity: function newContactStorage() view returns(address)
func (_ContractStorageImpl *ContractStorageImplCallerSession) NewContactStorage() (common.Address, error) {
	return _ContractStorageImpl.Contract.NewContactStorage(&_ContractStorageImpl.CallOpts)
}

// PreviousContactStorage is a free data retrieval call binding the contract method 0xeac8c753.
//
// Solidity: function previousContactStorage() view returns(address)
func (_ContractStorageImpl *ContractStorageImplCaller) PreviousContactStorage(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ContractStorageImpl.contract.Call(opts, out, "previousContactStorage")
	return *ret0, err
}

// PreviousContactStorage is a free data retrieval call binding the contract method 0xeac8c753.
//
// Solidity: function previousContactStorage() view returns(address)
func (_ContractStorageImpl *ContractStorageImplSession) PreviousContactStorage() (common.Address, error) {
	return _ContractStorageImpl.Contract.PreviousContactStorage(&_ContractStorageImpl.CallOpts)
}

// PreviousContactStorage is a free data retrieval call binding the contract method 0xeac8c753.
//
// Solidity: function previousContactStorage() view returns(address)
func (_ContractStorageImpl *ContractStorageImplCallerSession) PreviousContactStorage() (common.Address, error) {
	return _ContractStorageImpl.Contract.PreviousContactStorage(&_ContractStorageImpl.CallOpts)
}

// NotifyFirstDeploy is a paid mutator transaction binding the contract method 0x4fc6bddb.
//
// Solidity: function notifyFirstDeploy() returns()
func (_ContractStorageImpl *ContractStorageImplTransactor) NotifyFirstDeploy(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractStorageImpl.contract.Transact(opts, "notifyFirstDeploy")
}

// NotifyFirstDeploy is a paid mutator transaction binding the contract method 0x4fc6bddb.
//
// Solidity: function notifyFirstDeploy() returns()
func (_ContractStorageImpl *ContractStorageImplSession) NotifyFirstDeploy() (*types.Transaction, error) {
	return _ContractStorageImpl.Contract.NotifyFirstDeploy(&_ContractStorageImpl.TransactOpts)
}

// NotifyFirstDeploy is a paid mutator transaction binding the contract method 0x4fc6bddb.
//
// Solidity: function notifyFirstDeploy() returns()
func (_ContractStorageImpl *ContractStorageImplTransactorSession) NotifyFirstDeploy() (*types.Transaction, error) {
	return _ContractStorageImpl.Contract.NotifyFirstDeploy(&_ContractStorageImpl.TransactOpts)
}

// NotifyRedeploy is a paid mutator transaction binding the contract method 0x0f989265.
//
// Solidity: function notifyRedeploy(address newContractStorageAddress) returns()
func (_ContractStorageImpl *ContractStorageImplTransactor) NotifyRedeploy(opts *bind.TransactOpts, newContractStorageAddress common.Address) (*types.Transaction, error) {
	return _ContractStorageImpl.contract.Transact(opts, "notifyRedeploy", newContractStorageAddress)
}

// NotifyRedeploy is a paid mutator transaction binding the contract method 0x0f989265.
//
// Solidity: function notifyRedeploy(address newContractStorageAddress) returns()
func (_ContractStorageImpl *ContractStorageImplSession) NotifyRedeploy(newContractStorageAddress common.Address) (*types.Transaction, error) {
	return _ContractStorageImpl.Contract.NotifyRedeploy(&_ContractStorageImpl.TransactOpts, newContractStorageAddress)
}

// NotifyRedeploy is a paid mutator transaction binding the contract method 0x0f989265.
//
// Solidity: function notifyRedeploy(address newContractStorageAddress) returns()
func (_ContractStorageImpl *ContractStorageImplTransactorSession) NotifyRedeploy(newContractStorageAddress common.Address) (*types.Transaction, error) {
	return _ContractStorageImpl.Contract.NotifyRedeploy(&_ContractStorageImpl.TransactOpts, newContractStorageAddress)
}

// SetAccess is a paid mutator transaction binding the contract method 0x7d713ac2.
//
// Solidity: function setAccess(address contractAddress) returns()
func (_ContractStorageImpl *ContractStorageImplTransactor) SetAccess(opts *bind.TransactOpts, contractAddress common.Address) (*types.Transaction, error) {
	return _ContractStorageImpl.contract.Transact(opts, "setAccess", contractAddress)
}

// SetAccess is a paid mutator transaction binding the contract method 0x7d713ac2.
//
// Solidity: function setAccess(address contractAddress) returns()
func (_ContractStorageImpl *ContractStorageImplSession) SetAccess(contractAddress common.Address) (*types.Transaction, error) {
	return _ContractStorageImpl.Contract.SetAccess(&_ContractStorageImpl.TransactOpts, contractAddress)
}

// SetAccess is a paid mutator transaction binding the contract method 0x7d713ac2.
//
// Solidity: function setAccess(address contractAddress) returns()
func (_ContractStorageImpl *ContractStorageImplTransactorSession) SetAccess(contractAddress common.Address) (*types.Transaction, error) {
	return _ContractStorageImpl.Contract.SetAccess(&_ContractStorageImpl.TransactOpts, contractAddress)
}

// SetAccessStorage is a paid mutator transaction binding the contract method 0xdd191535.
//
// Solidity: function setAccessStorage(address contractAddress) returns()
func (_ContractStorageImpl *ContractStorageImplTransactor) SetAccessStorage(opts *bind.TransactOpts, contractAddress common.Address) (*types.Transaction, error) {
	return _ContractStorageImpl.contract.Transact(opts, "setAccessStorage", contractAddress)
}

// SetAccessStorage is a paid mutator transaction binding the contract method 0xdd191535.
//
// Solidity: function setAccessStorage(address contractAddress) returns()
func (_ContractStorageImpl *ContractStorageImplSession) SetAccessStorage(contractAddress common.Address) (*types.Transaction, error) {
	return _ContractStorageImpl.Contract.SetAccessStorage(&_ContractStorageImpl.TransactOpts, contractAddress)
}

// SetAccessStorage is a paid mutator transaction binding the contract method 0xdd191535.
//
// Solidity: function setAccessStorage(address contractAddress) returns()
func (_ContractStorageImpl *ContractStorageImplTransactorSession) SetAccessStorage(contractAddress common.Address) (*types.Transaction, error) {
	return _ContractStorageImpl.Contract.SetAccessStorage(&_ContractStorageImpl.TransactOpts, contractAddress)
}

// SetAdminable is a paid mutator transaction binding the contract method 0x14a35cf2.
//
// Solidity: function setAdminable(address adminableContractAddress) returns()
func (_ContractStorageImpl *ContractStorageImplTransactor) SetAdminable(opts *bind.TransactOpts, adminableContractAddress common.Address) (*types.Transaction, error) {
	return _ContractStorageImpl.contract.Transact(opts, "setAdminable", adminableContractAddress)
}

// SetAdminable is a paid mutator transaction binding the contract method 0x14a35cf2.
//
// Solidity: function setAdminable(address adminableContractAddress) returns()
func (_ContractStorageImpl *ContractStorageImplSession) SetAdminable(adminableContractAddress common.Address) (*types.Transaction, error) {
	return _ContractStorageImpl.Contract.SetAdminable(&_ContractStorageImpl.TransactOpts, adminableContractAddress)
}

// SetAdminable is a paid mutator transaction binding the contract method 0x14a35cf2.
//
// Solidity: function setAdminable(address adminableContractAddress) returns()
func (_ContractStorageImpl *ContractStorageImplTransactorSession) SetAdminable(adminableContractAddress common.Address) (*types.Transaction, error) {
	return _ContractStorageImpl.Contract.SetAdminable(&_ContractStorageImpl.TransactOpts, adminableContractAddress)
}

// SetAssetbox is a paid mutator transaction binding the contract method 0xaf386d64.
//
// Solidity: function setAssetbox(address contractAddress) returns()
func (_ContractStorageImpl *ContractStorageImplTransactor) SetAssetbox(opts *bind.TransactOpts, contractAddress common.Address) (*types.Transaction, error) {
	return _ContractStorageImpl.contract.Transact(opts, "setAssetbox", contractAddress)
}

// SetAssetbox is a paid mutator transaction binding the contract method 0xaf386d64.
//
// Solidity: function setAssetbox(address contractAddress) returns()
func (_ContractStorageImpl *ContractStorageImplSession) SetAssetbox(contractAddress common.Address) (*types.Transaction, error) {
	return _ContractStorageImpl.Contract.SetAssetbox(&_ContractStorageImpl.TransactOpts, contractAddress)
}

// SetAssetbox is a paid mutator transaction binding the contract method 0xaf386d64.
//
// Solidity: function setAssetbox(address contractAddress) returns()
func (_ContractStorageImpl *ContractStorageImplTransactorSession) SetAssetbox(contractAddress common.Address) (*types.Transaction, error) {
	return _ContractStorageImpl.Contract.SetAssetbox(&_ContractStorageImpl.TransactOpts, contractAddress)
}

// SetAssetboxInfo is a paid mutator transaction binding the contract method 0x909c2416.
//
// Solidity: function setAssetboxInfo(address contractAddress) returns()
func (_ContractStorageImpl *ContractStorageImplTransactor) SetAssetboxInfo(opts *bind.TransactOpts, contractAddress common.Address) (*types.Transaction, error) {
	return _ContractStorageImpl.contract.Transact(opts, "setAssetboxInfo", contractAddress)
}

// SetAssetboxInfo is a paid mutator transaction binding the contract method 0x909c2416.
//
// Solidity: function setAssetboxInfo(address contractAddress) returns()
func (_ContractStorageImpl *ContractStorageImplSession) SetAssetboxInfo(contractAddress common.Address) (*types.Transaction, error) {
	return _ContractStorageImpl.Contract.SetAssetboxInfo(&_ContractStorageImpl.TransactOpts, contractAddress)
}

// SetAssetboxInfo is a paid mutator transaction binding the contract method 0x909c2416.
//
// Solidity: function setAssetboxInfo(address contractAddress) returns()
func (_ContractStorageImpl *ContractStorageImplTransactorSession) SetAssetboxInfo(contractAddress common.Address) (*types.Transaction, error) {
	return _ContractStorageImpl.Contract.SetAssetboxInfo(&_ContractStorageImpl.TransactOpts, contractAddress)
}

// SetAssetboxStorage is a paid mutator transaction binding the contract method 0x5472ff6b.
//
// Solidity: function setAssetboxStorage(address contractAddress) returns()
func (_ContractStorageImpl *ContractStorageImplTransactor) SetAssetboxStorage(opts *bind.TransactOpts, contractAddress common.Address) (*types.Transaction, error) {
	return _ContractStorageImpl.contract.Transact(opts, "setAssetboxStorage", contractAddress)
}

// SetAssetboxStorage is a paid mutator transaction binding the contract method 0x5472ff6b.
//
// Solidity: function setAssetboxStorage(address contractAddress) returns()
func (_ContractStorageImpl *ContractStorageImplSession) SetAssetboxStorage(contractAddress common.Address) (*types.Transaction, error) {
	return _ContractStorageImpl.Contract.SetAssetboxStorage(&_ContractStorageImpl.TransactOpts, contractAddress)
}

// SetAssetboxStorage is a paid mutator transaction binding the contract method 0x5472ff6b.
//
// Solidity: function setAssetboxStorage(address contractAddress) returns()
func (_ContractStorageImpl *ContractStorageImplTransactorSession) SetAssetboxStorage(contractAddress common.Address) (*types.Transaction, error) {
	return _ContractStorageImpl.Contract.SetAssetboxStorage(&_ContractStorageImpl.TransactOpts, contractAddress)
}

// SetBalanceGate is a paid mutator transaction binding the contract method 0x286bf9d5.
//
// Solidity: function setBalanceGate(address contractAddress) returns()
func (_ContractStorageImpl *ContractStorageImplTransactor) SetBalanceGate(opts *bind.TransactOpts, contractAddress common.Address) (*types.Transaction, error) {
	return _ContractStorageImpl.contract.Transact(opts, "setBalanceGate", contractAddress)
}

// SetBalanceGate is a paid mutator transaction binding the contract method 0x286bf9d5.
//
// Solidity: function setBalanceGate(address contractAddress) returns()
func (_ContractStorageImpl *ContractStorageImplSession) SetBalanceGate(contractAddress common.Address) (*types.Transaction, error) {
	return _ContractStorageImpl.Contract.SetBalanceGate(&_ContractStorageImpl.TransactOpts, contractAddress)
}

// SetBalanceGate is a paid mutator transaction binding the contract method 0x286bf9d5.
//
// Solidity: function setBalanceGate(address contractAddress) returns()
func (_ContractStorageImpl *ContractStorageImplTransactorSession) SetBalanceGate(contractAddress common.Address) (*types.Transaction, error) {
	return _ContractStorageImpl.Contract.SetBalanceGate(&_ContractStorageImpl.TransactOpts, contractAddress)
}

// SetBitbon is a paid mutator transaction binding the contract method 0x76ebaba3.
//
// Solidity: function setBitbon(address contractAddress) returns()
func (_ContractStorageImpl *ContractStorageImplTransactor) SetBitbon(opts *bind.TransactOpts, contractAddress common.Address) (*types.Transaction, error) {
	return _ContractStorageImpl.contract.Transact(opts, "setBitbon", contractAddress)
}

// SetBitbon is a paid mutator transaction binding the contract method 0x76ebaba3.
//
// Solidity: function setBitbon(address contractAddress) returns()
func (_ContractStorageImpl *ContractStorageImplSession) SetBitbon(contractAddress common.Address) (*types.Transaction, error) {
	return _ContractStorageImpl.Contract.SetBitbon(&_ContractStorageImpl.TransactOpts, contractAddress)
}

// SetBitbon is a paid mutator transaction binding the contract method 0x76ebaba3.
//
// Solidity: function setBitbon(address contractAddress) returns()
func (_ContractStorageImpl *ContractStorageImplTransactorSession) SetBitbon(contractAddress common.Address) (*types.Transaction, error) {
	return _ContractStorageImpl.Contract.SetBitbon(&_ContractStorageImpl.TransactOpts, contractAddress)
}

// SetBitbonStorage is a paid mutator transaction binding the contract method 0x18e02cb3.
//
// Solidity: function setBitbonStorage(address contractAddress) returns()
func (_ContractStorageImpl *ContractStorageImplTransactor) SetBitbonStorage(opts *bind.TransactOpts, contractAddress common.Address) (*types.Transaction, error) {
	return _ContractStorageImpl.contract.Transact(opts, "setBitbonStorage", contractAddress)
}

// SetBitbonStorage is a paid mutator transaction binding the contract method 0x18e02cb3.
//
// Solidity: function setBitbonStorage(address contractAddress) returns()
func (_ContractStorageImpl *ContractStorageImplSession) SetBitbonStorage(contractAddress common.Address) (*types.Transaction, error) {
	return _ContractStorageImpl.Contract.SetBitbonStorage(&_ContractStorageImpl.TransactOpts, contractAddress)
}

// SetBitbonStorage is a paid mutator transaction binding the contract method 0x18e02cb3.
//
// Solidity: function setBitbonStorage(address contractAddress) returns()
func (_ContractStorageImpl *ContractStorageImplTransactorSession) SetBitbonStorage(contractAddress common.Address) (*types.Transaction, error) {
	return _ContractStorageImpl.Contract.SetBitbonStorage(&_ContractStorageImpl.TransactOpts, contractAddress)
}

// SetBitbonSupport is a paid mutator transaction binding the contract method 0xabd5d1c3.
//
// Solidity: function setBitbonSupport(address contractAddress) returns()
func (_ContractStorageImpl *ContractStorageImplTransactor) SetBitbonSupport(opts *bind.TransactOpts, contractAddress common.Address) (*types.Transaction, error) {
	return _ContractStorageImpl.contract.Transact(opts, "setBitbonSupport", contractAddress)
}

// SetBitbonSupport is a paid mutator transaction binding the contract method 0xabd5d1c3.
//
// Solidity: function setBitbonSupport(address contractAddress) returns()
func (_ContractStorageImpl *ContractStorageImplSession) SetBitbonSupport(contractAddress common.Address) (*types.Transaction, error) {
	return _ContractStorageImpl.Contract.SetBitbonSupport(&_ContractStorageImpl.TransactOpts, contractAddress)
}

// SetBitbonSupport is a paid mutator transaction binding the contract method 0xabd5d1c3.
//
// Solidity: function setBitbonSupport(address contractAddress) returns()
func (_ContractStorageImpl *ContractStorageImplTransactorSession) SetBitbonSupport(contractAddress common.Address) (*types.Transaction, error) {
	return _ContractStorageImpl.Contract.SetBitbonSupport(&_ContractStorageImpl.TransactOpts, contractAddress)
}

// SetBonBalanceGate is a paid mutator transaction binding the contract method 0x787f0ff3.
//
// Solidity: function setBonBalanceGate(address contractAddress) returns()
func (_ContractStorageImpl *ContractStorageImplTransactor) SetBonBalanceGate(opts *bind.TransactOpts, contractAddress common.Address) (*types.Transaction, error) {
	return _ContractStorageImpl.contract.Transact(opts, "setBonBalanceGate", contractAddress)
}

// SetBonBalanceGate is a paid mutator transaction binding the contract method 0x787f0ff3.
//
// Solidity: function setBonBalanceGate(address contractAddress) returns()
func (_ContractStorageImpl *ContractStorageImplSession) SetBonBalanceGate(contractAddress common.Address) (*types.Transaction, error) {
	return _ContractStorageImpl.Contract.SetBonBalanceGate(&_ContractStorageImpl.TransactOpts, contractAddress)
}

// SetBonBalanceGate is a paid mutator transaction binding the contract method 0x787f0ff3.
//
// Solidity: function setBonBalanceGate(address contractAddress) returns()
func (_ContractStorageImpl *ContractStorageImplTransactorSession) SetBonBalanceGate(contractAddress common.Address) (*types.Transaction, error) {
	return _ContractStorageImpl.Contract.SetBonBalanceGate(&_ContractStorageImpl.TransactOpts, contractAddress)
}

// SetBonTransferGate is a paid mutator transaction binding the contract method 0x4ed2fb16.
//
// Solidity: function setBonTransferGate(address contractAddress) returns()
func (_ContractStorageImpl *ContractStorageImplTransactor) SetBonTransferGate(opts *bind.TransactOpts, contractAddress common.Address) (*types.Transaction, error) {
	return _ContractStorageImpl.contract.Transact(opts, "setBonTransferGate", contractAddress)
}

// SetBonTransferGate is a paid mutator transaction binding the contract method 0x4ed2fb16.
//
// Solidity: function setBonTransferGate(address contractAddress) returns()
func (_ContractStorageImpl *ContractStorageImplSession) SetBonTransferGate(contractAddress common.Address) (*types.Transaction, error) {
	return _ContractStorageImpl.Contract.SetBonTransferGate(&_ContractStorageImpl.TransactOpts, contractAddress)
}

// SetBonTransferGate is a paid mutator transaction binding the contract method 0x4ed2fb16.
//
// Solidity: function setBonTransferGate(address contractAddress) returns()
func (_ContractStorageImpl *ContractStorageImplTransactorSession) SetBonTransferGate(contractAddress common.Address) (*types.Transaction, error) {
	return _ContractStorageImpl.Contract.SetBonTransferGate(&_ContractStorageImpl.TransactOpts, contractAddress)
}

// SetContractAddress is a paid mutator transaction binding the contract method 0xbb0165b1.
//
// Solidity: function setContractAddress(uint256 contractType, address contractAddress) returns()
func (_ContractStorageImpl *ContractStorageImplTransactor) SetContractAddress(opts *bind.TransactOpts, contractType *big.Int, contractAddress common.Address) (*types.Transaction, error) {
	return _ContractStorageImpl.contract.Transact(opts, "setContractAddress", contractType, contractAddress)
}

// SetContractAddress is a paid mutator transaction binding the contract method 0xbb0165b1.
//
// Solidity: function setContractAddress(uint256 contractType, address contractAddress) returns()
func (_ContractStorageImpl *ContractStorageImplSession) SetContractAddress(contractType *big.Int, contractAddress common.Address) (*types.Transaction, error) {
	return _ContractStorageImpl.Contract.SetContractAddress(&_ContractStorageImpl.TransactOpts, contractType, contractAddress)
}

// SetContractAddress is a paid mutator transaction binding the contract method 0xbb0165b1.
//
// Solidity: function setContractAddress(uint256 contractType, address contractAddress) returns()
func (_ContractStorageImpl *ContractStorageImplTransactorSession) SetContractAddress(contractType *big.Int, contractAddress common.Address) (*types.Transaction, error) {
	return _ContractStorageImpl.Contract.SetContractAddress(&_ContractStorageImpl.TransactOpts, contractType, contractAddress)
}

// SetDistributionStorage is a paid mutator transaction binding the contract method 0x99cc5501.
//
// Solidity: function setDistributionStorage(address distributionStorageContractAddress) returns()
func (_ContractStorageImpl *ContractStorageImplTransactor) SetDistributionStorage(opts *bind.TransactOpts, distributionStorageContractAddress common.Address) (*types.Transaction, error) {
	return _ContractStorageImpl.contract.Transact(opts, "setDistributionStorage", distributionStorageContractAddress)
}

// SetDistributionStorage is a paid mutator transaction binding the contract method 0x99cc5501.
//
// Solidity: function setDistributionStorage(address distributionStorageContractAddress) returns()
func (_ContractStorageImpl *ContractStorageImplSession) SetDistributionStorage(distributionStorageContractAddress common.Address) (*types.Transaction, error) {
	return _ContractStorageImpl.Contract.SetDistributionStorage(&_ContractStorageImpl.TransactOpts, distributionStorageContractAddress)
}

// SetDistributionStorage is a paid mutator transaction binding the contract method 0x99cc5501.
//
// Solidity: function setDistributionStorage(address distributionStorageContractAddress) returns()
func (_ContractStorageImpl *ContractStorageImplTransactorSession) SetDistributionStorage(distributionStorageContractAddress common.Address) (*types.Transaction, error) {
	return _ContractStorageImpl.Contract.SetDistributionStorage(&_ContractStorageImpl.TransactOpts, distributionStorageContractAddress)
}

// SetExchange is a paid mutator transaction binding the contract method 0x67b1f5df.
//
// Solidity: function setExchange(address contractAddress) returns()
func (_ContractStorageImpl *ContractStorageImplTransactor) SetExchange(opts *bind.TransactOpts, contractAddress common.Address) (*types.Transaction, error) {
	return _ContractStorageImpl.contract.Transact(opts, "setExchange", contractAddress)
}

// SetExchange is a paid mutator transaction binding the contract method 0x67b1f5df.
//
// Solidity: function setExchange(address contractAddress) returns()
func (_ContractStorageImpl *ContractStorageImplSession) SetExchange(contractAddress common.Address) (*types.Transaction, error) {
	return _ContractStorageImpl.Contract.SetExchange(&_ContractStorageImpl.TransactOpts, contractAddress)
}

// SetExchange is a paid mutator transaction binding the contract method 0x67b1f5df.
//
// Solidity: function setExchange(address contractAddress) returns()
func (_ContractStorageImpl *ContractStorageImplTransactorSession) SetExchange(contractAddress common.Address) (*types.Transaction, error) {
	return _ContractStorageImpl.Contract.SetExchange(&_ContractStorageImpl.TransactOpts, contractAddress)
}

// SetExchangeStorage is a paid mutator transaction binding the contract method 0xe7b0a826.
//
// Solidity: function setExchangeStorage(address contractAddress) returns()
func (_ContractStorageImpl *ContractStorageImplTransactor) SetExchangeStorage(opts *bind.TransactOpts, contractAddress common.Address) (*types.Transaction, error) {
	return _ContractStorageImpl.contract.Transact(opts, "setExchangeStorage", contractAddress)
}

// SetExchangeStorage is a paid mutator transaction binding the contract method 0xe7b0a826.
//
// Solidity: function setExchangeStorage(address contractAddress) returns()
func (_ContractStorageImpl *ContractStorageImplSession) SetExchangeStorage(contractAddress common.Address) (*types.Transaction, error) {
	return _ContractStorageImpl.Contract.SetExchangeStorage(&_ContractStorageImpl.TransactOpts, contractAddress)
}

// SetExchangeStorage is a paid mutator transaction binding the contract method 0xe7b0a826.
//
// Solidity: function setExchangeStorage(address contractAddress) returns()
func (_ContractStorageImpl *ContractStorageImplTransactorSession) SetExchangeStorage(contractAddress common.Address) (*types.Transaction, error) {
	return _ContractStorageImpl.Contract.SetExchangeStorage(&_ContractStorageImpl.TransactOpts, contractAddress)
}

// SetFeeStorage is a paid mutator transaction binding the contract method 0xde6e3411.
//
// Solidity: function setFeeStorage(address contractAddress) returns()
func (_ContractStorageImpl *ContractStorageImplTransactor) SetFeeStorage(opts *bind.TransactOpts, contractAddress common.Address) (*types.Transaction, error) {
	return _ContractStorageImpl.contract.Transact(opts, "setFeeStorage", contractAddress)
}

// SetFeeStorage is a paid mutator transaction binding the contract method 0xde6e3411.
//
// Solidity: function setFeeStorage(address contractAddress) returns()
func (_ContractStorageImpl *ContractStorageImplSession) SetFeeStorage(contractAddress common.Address) (*types.Transaction, error) {
	return _ContractStorageImpl.Contract.SetFeeStorage(&_ContractStorageImpl.TransactOpts, contractAddress)
}

// SetFeeStorage is a paid mutator transaction binding the contract method 0xde6e3411.
//
// Solidity: function setFeeStorage(address contractAddress) returns()
func (_ContractStorageImpl *ContractStorageImplTransactorSession) SetFeeStorage(contractAddress common.Address) (*types.Transaction, error) {
	return _ContractStorageImpl.Contract.SetFeeStorage(&_ContractStorageImpl.TransactOpts, contractAddress)
}

// SetMiningAgentStorage is a paid mutator transaction binding the contract method 0xe557af04.
//
// Solidity: function setMiningAgentStorage(address miningAgentStorageContractAddress) returns()
func (_ContractStorageImpl *ContractStorageImplTransactor) SetMiningAgentStorage(opts *bind.TransactOpts, miningAgentStorageContractAddress common.Address) (*types.Transaction, error) {
	return _ContractStorageImpl.contract.Transact(opts, "setMiningAgentStorage", miningAgentStorageContractAddress)
}

// SetMiningAgentStorage is a paid mutator transaction binding the contract method 0xe557af04.
//
// Solidity: function setMiningAgentStorage(address miningAgentStorageContractAddress) returns()
func (_ContractStorageImpl *ContractStorageImplSession) SetMiningAgentStorage(miningAgentStorageContractAddress common.Address) (*types.Transaction, error) {
	return _ContractStorageImpl.Contract.SetMiningAgentStorage(&_ContractStorageImpl.TransactOpts, miningAgentStorageContractAddress)
}

// SetMiningAgentStorage is a paid mutator transaction binding the contract method 0xe557af04.
//
// Solidity: function setMiningAgentStorage(address miningAgentStorageContractAddress) returns()
func (_ContractStorageImpl *ContractStorageImplTransactorSession) SetMiningAgentStorage(miningAgentStorageContractAddress common.Address) (*types.Transaction, error) {
	return _ContractStorageImpl.Contract.SetMiningAgentStorage(&_ContractStorageImpl.TransactOpts, miningAgentStorageContractAddress)
}

// SetMsbon is a paid mutator transaction binding the contract method 0x28c2c155.
//
// Solidity: function setMsbon(address contractAddress) returns()
func (_ContractStorageImpl *ContractStorageImplTransactor) SetMsbon(opts *bind.TransactOpts, contractAddress common.Address) (*types.Transaction, error) {
	return _ContractStorageImpl.contract.Transact(opts, "setMsbon", contractAddress)
}

// SetMsbon is a paid mutator transaction binding the contract method 0x28c2c155.
//
// Solidity: function setMsbon(address contractAddress) returns()
func (_ContractStorageImpl *ContractStorageImplSession) SetMsbon(contractAddress common.Address) (*types.Transaction, error) {
	return _ContractStorageImpl.Contract.SetMsbon(&_ContractStorageImpl.TransactOpts, contractAddress)
}

// SetMsbon is a paid mutator transaction binding the contract method 0x28c2c155.
//
// Solidity: function setMsbon(address contractAddress) returns()
func (_ContractStorageImpl *ContractStorageImplTransactorSession) SetMsbon(contractAddress common.Address) (*types.Transaction, error) {
	return _ContractStorageImpl.Contract.SetMsbon(&_ContractStorageImpl.TransactOpts, contractAddress)
}

// SetMsbonStorage is a paid mutator transaction binding the contract method 0x9e529cb4.
//
// Solidity: function setMsbonStorage(address contractAddress) returns()
func (_ContractStorageImpl *ContractStorageImplTransactor) SetMsbonStorage(opts *bind.TransactOpts, contractAddress common.Address) (*types.Transaction, error) {
	return _ContractStorageImpl.contract.Transact(opts, "setMsbonStorage", contractAddress)
}

// SetMsbonStorage is a paid mutator transaction binding the contract method 0x9e529cb4.
//
// Solidity: function setMsbonStorage(address contractAddress) returns()
func (_ContractStorageImpl *ContractStorageImplSession) SetMsbonStorage(contractAddress common.Address) (*types.Transaction, error) {
	return _ContractStorageImpl.Contract.SetMsbonStorage(&_ContractStorageImpl.TransactOpts, contractAddress)
}

// SetMsbonStorage is a paid mutator transaction binding the contract method 0x9e529cb4.
//
// Solidity: function setMsbonStorage(address contractAddress) returns()
func (_ContractStorageImpl *ContractStorageImplTransactorSession) SetMsbonStorage(contractAddress common.Address) (*types.Transaction, error) {
	return _ContractStorageImpl.Contract.SetMsbonStorage(&_ContractStorageImpl.TransactOpts, contractAddress)
}

// SetMultisigStorage is a paid mutator transaction binding the contract method 0xa86baad0.
//
// Solidity: function setMultisigStorage(address contractAddress) returns()
func (_ContractStorageImpl *ContractStorageImplTransactor) SetMultisigStorage(opts *bind.TransactOpts, contractAddress common.Address) (*types.Transaction, error) {
	return _ContractStorageImpl.contract.Transact(opts, "setMultisigStorage", contractAddress)
}

// SetMultisigStorage is a paid mutator transaction binding the contract method 0xa86baad0.
//
// Solidity: function setMultisigStorage(address contractAddress) returns()
func (_ContractStorageImpl *ContractStorageImplSession) SetMultisigStorage(contractAddress common.Address) (*types.Transaction, error) {
	return _ContractStorageImpl.Contract.SetMultisigStorage(&_ContractStorageImpl.TransactOpts, contractAddress)
}

// SetMultisigStorage is a paid mutator transaction binding the contract method 0xa86baad0.
//
// Solidity: function setMultisigStorage(address contractAddress) returns()
func (_ContractStorageImpl *ContractStorageImplTransactorSession) SetMultisigStorage(contractAddress common.Address) (*types.Transaction, error) {
	return _ContractStorageImpl.Contract.SetMultisigStorage(&_ContractStorageImpl.TransactOpts, contractAddress)
}

// SetOtc is a paid mutator transaction binding the contract method 0x53585888.
//
// Solidity: function setOtc(address contractAddress) returns()
func (_ContractStorageImpl *ContractStorageImplTransactor) SetOtc(opts *bind.TransactOpts, contractAddress common.Address) (*types.Transaction, error) {
	return _ContractStorageImpl.contract.Transact(opts, "setOtc", contractAddress)
}

// SetOtc is a paid mutator transaction binding the contract method 0x53585888.
//
// Solidity: function setOtc(address contractAddress) returns()
func (_ContractStorageImpl *ContractStorageImplSession) SetOtc(contractAddress common.Address) (*types.Transaction, error) {
	return _ContractStorageImpl.Contract.SetOtc(&_ContractStorageImpl.TransactOpts, contractAddress)
}

// SetOtc is a paid mutator transaction binding the contract method 0x53585888.
//
// Solidity: function setOtc(address contractAddress) returns()
func (_ContractStorageImpl *ContractStorageImplTransactorSession) SetOtc(contractAddress common.Address) (*types.Transaction, error) {
	return _ContractStorageImpl.Contract.SetOtc(&_ContractStorageImpl.TransactOpts, contractAddress)
}

// SetOtcStorage is a paid mutator transaction binding the contract method 0x0042aa78.
//
// Solidity: function setOtcStorage(address contractAddress) returns()
func (_ContractStorageImpl *ContractStorageImplTransactor) SetOtcStorage(opts *bind.TransactOpts, contractAddress common.Address) (*types.Transaction, error) {
	return _ContractStorageImpl.contract.Transact(opts, "setOtcStorage", contractAddress)
}

// SetOtcStorage is a paid mutator transaction binding the contract method 0x0042aa78.
//
// Solidity: function setOtcStorage(address contractAddress) returns()
func (_ContractStorageImpl *ContractStorageImplSession) SetOtcStorage(contractAddress common.Address) (*types.Transaction, error) {
	return _ContractStorageImpl.Contract.SetOtcStorage(&_ContractStorageImpl.TransactOpts, contractAddress)
}

// SetOtcStorage is a paid mutator transaction binding the contract method 0x0042aa78.
//
// Solidity: function setOtcStorage(address contractAddress) returns()
func (_ContractStorageImpl *ContractStorageImplTransactorSession) SetOtcStorage(contractAddress common.Address) (*types.Transaction, error) {
	return _ContractStorageImpl.Contract.SetOtcStorage(&_ContractStorageImpl.TransactOpts, contractAddress)
}

// SetPreviousContractStorageAddress is a paid mutator transaction binding the contract method 0x896f578f.
//
// Solidity: function setPreviousContractStorageAddress(address contractAddress) returns()
func (_ContractStorageImpl *ContractStorageImplTransactor) SetPreviousContractStorageAddress(opts *bind.TransactOpts, contractAddress common.Address) (*types.Transaction, error) {
	return _ContractStorageImpl.contract.Transact(opts, "setPreviousContractStorageAddress", contractAddress)
}

// SetPreviousContractStorageAddress is a paid mutator transaction binding the contract method 0x896f578f.
//
// Solidity: function setPreviousContractStorageAddress(address contractAddress) returns()
func (_ContractStorageImpl *ContractStorageImplSession) SetPreviousContractStorageAddress(contractAddress common.Address) (*types.Transaction, error) {
	return _ContractStorageImpl.Contract.SetPreviousContractStorageAddress(&_ContractStorageImpl.TransactOpts, contractAddress)
}

// SetPreviousContractStorageAddress is a paid mutator transaction binding the contract method 0x896f578f.
//
// Solidity: function setPreviousContractStorageAddress(address contractAddress) returns()
func (_ContractStorageImpl *ContractStorageImplTransactorSession) SetPreviousContractStorageAddress(contractAddress common.Address) (*types.Transaction, error) {
	return _ContractStorageImpl.Contract.SetPreviousContractStorageAddress(&_ContractStorageImpl.TransactOpts, contractAddress)
}

// SetReservedAliasStorage is a paid mutator transaction binding the contract method 0x3aa6bb19.
//
// Solidity: function setReservedAliasStorage(address contractAddress) returns()
func (_ContractStorageImpl *ContractStorageImplTransactor) SetReservedAliasStorage(opts *bind.TransactOpts, contractAddress common.Address) (*types.Transaction, error) {
	return _ContractStorageImpl.contract.Transact(opts, "setReservedAliasStorage", contractAddress)
}

// SetReservedAliasStorage is a paid mutator transaction binding the contract method 0x3aa6bb19.
//
// Solidity: function setReservedAliasStorage(address contractAddress) returns()
func (_ContractStorageImpl *ContractStorageImplSession) SetReservedAliasStorage(contractAddress common.Address) (*types.Transaction, error) {
	return _ContractStorageImpl.Contract.SetReservedAliasStorage(&_ContractStorageImpl.TransactOpts, contractAddress)
}

// SetReservedAliasStorage is a paid mutator transaction binding the contract method 0x3aa6bb19.
//
// Solidity: function setReservedAliasStorage(address contractAddress) returns()
func (_ContractStorageImpl *ContractStorageImplTransactorSession) SetReservedAliasStorage(contractAddress common.Address) (*types.Transaction, error) {
	return _ContractStorageImpl.Contract.SetReservedAliasStorage(&_ContractStorageImpl.TransactOpts, contractAddress)
}

// SetSafeTransferStorage is a paid mutator transaction binding the contract method 0x933ef0d8.
//
// Solidity: function setSafeTransferStorage(address contractAddress) returns()
func (_ContractStorageImpl *ContractStorageImplTransactor) SetSafeTransferStorage(opts *bind.TransactOpts, contractAddress common.Address) (*types.Transaction, error) {
	return _ContractStorageImpl.contract.Transact(opts, "setSafeTransferStorage", contractAddress)
}

// SetSafeTransferStorage is a paid mutator transaction binding the contract method 0x933ef0d8.
//
// Solidity: function setSafeTransferStorage(address contractAddress) returns()
func (_ContractStorageImpl *ContractStorageImplSession) SetSafeTransferStorage(contractAddress common.Address) (*types.Transaction, error) {
	return _ContractStorageImpl.Contract.SetSafeTransferStorage(&_ContractStorageImpl.TransactOpts, contractAddress)
}

// SetSafeTransferStorage is a paid mutator transaction binding the contract method 0x933ef0d8.
//
// Solidity: function setSafeTransferStorage(address contractAddress) returns()
func (_ContractStorageImpl *ContractStorageImplTransactorSession) SetSafeTransferStorage(contractAddress common.Address) (*types.Transaction, error) {
	return _ContractStorageImpl.Contract.SetSafeTransferStorage(&_ContractStorageImpl.TransactOpts, contractAddress)
}

// SetTransferGate is a paid mutator transaction binding the contract method 0x61d8c817.
//
// Solidity: function setTransferGate(address contractAddress) returns()
func (_ContractStorageImpl *ContractStorageImplTransactor) SetTransferGate(opts *bind.TransactOpts, contractAddress common.Address) (*types.Transaction, error) {
	return _ContractStorageImpl.contract.Transact(opts, "setTransferGate", contractAddress)
}

// SetTransferGate is a paid mutator transaction binding the contract method 0x61d8c817.
//
// Solidity: function setTransferGate(address contractAddress) returns()
func (_ContractStorageImpl *ContractStorageImplSession) SetTransferGate(contractAddress common.Address) (*types.Transaction, error) {
	return _ContractStorageImpl.Contract.SetTransferGate(&_ContractStorageImpl.TransactOpts, contractAddress)
}

// SetTransferGate is a paid mutator transaction binding the contract method 0x61d8c817.
//
// Solidity: function setTransferGate(address contractAddress) returns()
func (_ContractStorageImpl *ContractStorageImplTransactorSession) SetTransferGate(contractAddress common.Address) (*types.Transaction, error) {
	return _ContractStorageImpl.Contract.SetTransferGate(&_ContractStorageImpl.TransactOpts, contractAddress)
}

// ContractStorageImplContractDeployedIterator is returned from FilterContractDeployed and is used to iterate over the raw logs and unpacked data for ContractDeployed events raised by the ContractStorageImpl contract.
type ContractStorageImplContractDeployedIterator struct {
	Event *ContractStorageImplContractDeployed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractStorageImplContractDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractStorageImplContractDeployed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractStorageImplContractDeployed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractStorageImplContractDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractStorageImplContractDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractStorageImplContractDeployed represents a ContractDeployed event raised by the ContractStorageImpl contract.
type ContractStorageImplContractDeployed struct {
	ContractIndex *big.Int
	OldAddress    common.Address
	NewAddress    common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterContractDeployed is a free log retrieval operation binding the contract event 0xd274d16612c69f430bc382e211c9223ab05c6e67d054e0db95d6f9345069613d.
//
// Solidity: event ContractDeployed(uint256 contractIndex, address oldAddress, address newAddress)
func (_ContractStorageImpl *ContractStorageImplFilterer) FilterContractDeployed(opts *bind.FilterOpts) (*ContractStorageImplContractDeployedIterator, error) {

	logs, sub, err := _ContractStorageImpl.contract.FilterLogs(opts, "ContractDeployed")
	if err != nil {
		return nil, err
	}
	return &ContractStorageImplContractDeployedIterator{contract: _ContractStorageImpl.contract, event: "ContractDeployed", logs: logs, sub: sub}, nil
}

// WatchContractDeployed is a free log subscription operation binding the contract event 0xd274d16612c69f430bc382e211c9223ab05c6e67d054e0db95d6f9345069613d.
//
// Solidity: event ContractDeployed(uint256 contractIndex, address oldAddress, address newAddress)
func (_ContractStorageImpl *ContractStorageImplFilterer) WatchContractDeployed(opts *bind.WatchOpts, sink chan<- *ContractStorageImplContractDeployed) (event.Subscription, error) {

	logs, sub, err := _ContractStorageImpl.contract.WatchLogs(opts, "ContractDeployed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractStorageImplContractDeployed)
				if err := _ContractStorageImpl.contract.UnpackLog(event, "ContractDeployed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseContractDeployed is a log parse operation binding the contract event 0xd274d16612c69f430bc382e211c9223ab05c6e67d054e0db95d6f9345069613d.
//
// Solidity: event ContractDeployed(uint256 contractIndex, address oldAddress, address newAddress)
func (_ContractStorageImpl *ContractStorageImplFilterer) ParseContractDeployed(log types.Log) (*ContractStorageImplContractDeployed, error) {
	event := new(ContractStorageImplContractDeployed)
	if err := _ContractStorageImpl.contract.UnpackLog(event, "ContractDeployed", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ContractStorageImplContractStorageMovedIterator is returned from FilterContractStorageMoved and is used to iterate over the raw logs and unpacked data for ContractStorageMoved events raised by the ContractStorageImpl contract.
type ContractStorageImplContractStorageMovedIterator struct {
	Event *ContractStorageImplContractStorageMoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractStorageImplContractStorageMovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractStorageImplContractStorageMoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractStorageImplContractStorageMoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractStorageImplContractStorageMovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractStorageImplContractStorageMovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractStorageImplContractStorageMoved represents a ContractStorageMoved event raised by the ContractStorageImpl contract.
type ContractStorageImplContractStorageMoved struct {
	OldAddress common.Address
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterContractStorageMoved is a free log retrieval operation binding the contract event 0x697834b862967028200fa0b8d24435bf6677d69e49574445af8dd965412d8772.
//
// Solidity: event ContractStorageMoved(address oldAddress, address newAddress)
func (_ContractStorageImpl *ContractStorageImplFilterer) FilterContractStorageMoved(opts *bind.FilterOpts) (*ContractStorageImplContractStorageMovedIterator, error) {

	logs, sub, err := _ContractStorageImpl.contract.FilterLogs(opts, "ContractStorageMoved")
	if err != nil {
		return nil, err
	}
	return &ContractStorageImplContractStorageMovedIterator{contract: _ContractStorageImpl.contract, event: "ContractStorageMoved", logs: logs, sub: sub}, nil
}

// WatchContractStorageMoved is a free log subscription operation binding the contract event 0x697834b862967028200fa0b8d24435bf6677d69e49574445af8dd965412d8772.
//
// Solidity: event ContractStorageMoved(address oldAddress, address newAddress)
func (_ContractStorageImpl *ContractStorageImplFilterer) WatchContractStorageMoved(opts *bind.WatchOpts, sink chan<- *ContractStorageImplContractStorageMoved) (event.Subscription, error) {

	logs, sub, err := _ContractStorageImpl.contract.WatchLogs(opts, "ContractStorageMoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractStorageImplContractStorageMoved)
				if err := _ContractStorageImpl.contract.UnpackLog(event, "ContractStorageMoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseContractStorageMoved is a log parse operation binding the contract event 0x697834b862967028200fa0b8d24435bf6677d69e49574445af8dd965412d8772.
//
// Solidity: event ContractStorageMoved(address oldAddress, address newAddress)
func (_ContractStorageImpl *ContractStorageImplFilterer) ParseContractStorageMoved(log types.Log) (*ContractStorageImplContractStorageMoved, error) {
	event := new(ContractStorageImplContractStorageMoved)
	if err := _ContractStorageImpl.contract.UnpackLog(event, "ContractStorageMoved", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ContractStorageImplNewContractStorageDeployedIterator is returned from FilterNewContractStorageDeployed and is used to iterate over the raw logs and unpacked data for NewContractStorageDeployed events raised by the ContractStorageImpl contract.
type ContractStorageImplNewContractStorageDeployedIterator struct {
	Event *ContractStorageImplNewContractStorageDeployed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractStorageImplNewContractStorageDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractStorageImplNewContractStorageDeployed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractStorageImplNewContractStorageDeployed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractStorageImplNewContractStorageDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractStorageImplNewContractStorageDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractStorageImplNewContractStorageDeployed represents a NewContractStorageDeployed event raised by the ContractStorageImpl contract.
type ContractStorageImplNewContractStorageDeployed struct {
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNewContractStorageDeployed is a free log retrieval operation binding the contract event 0xe328ac2c88b0be7b0d3d83128c921fb5ecc78e12a5cf6a2337d2aef5adf74ec5.
//
// Solidity: event NewContractStorageDeployed(address newAddress)
func (_ContractStorageImpl *ContractStorageImplFilterer) FilterNewContractStorageDeployed(opts *bind.FilterOpts) (*ContractStorageImplNewContractStorageDeployedIterator, error) {

	logs, sub, err := _ContractStorageImpl.contract.FilterLogs(opts, "NewContractStorageDeployed")
	if err != nil {
		return nil, err
	}
	return &ContractStorageImplNewContractStorageDeployedIterator{contract: _ContractStorageImpl.contract, event: "NewContractStorageDeployed", logs: logs, sub: sub}, nil
}

// WatchNewContractStorageDeployed is a free log subscription operation binding the contract event 0xe328ac2c88b0be7b0d3d83128c921fb5ecc78e12a5cf6a2337d2aef5adf74ec5.
//
// Solidity: event NewContractStorageDeployed(address newAddress)
func (_ContractStorageImpl *ContractStorageImplFilterer) WatchNewContractStorageDeployed(opts *bind.WatchOpts, sink chan<- *ContractStorageImplNewContractStorageDeployed) (event.Subscription, error) {

	logs, sub, err := _ContractStorageImpl.contract.WatchLogs(opts, "NewContractStorageDeployed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractStorageImplNewContractStorageDeployed)
				if err := _ContractStorageImpl.contract.UnpackLog(event, "NewContractStorageDeployed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewContractStorageDeployed is a log parse operation binding the contract event 0xe328ac2c88b0be7b0d3d83128c921fb5ecc78e12a5cf6a2337d2aef5adf74ec5.
//
// Solidity: event NewContractStorageDeployed(address newAddress)
func (_ContractStorageImpl *ContractStorageImplFilterer) ParseNewContractStorageDeployed(log types.Log) (*ContractStorageImplNewContractStorageDeployed, error) {
	event := new(ContractStorageImplNewContractStorageDeployed)
	if err := _ContractStorageImpl.contract.UnpackLog(event, "NewContractStorageDeployed", log); err != nil {
		return nil, err
	}
	return event, nil
}
