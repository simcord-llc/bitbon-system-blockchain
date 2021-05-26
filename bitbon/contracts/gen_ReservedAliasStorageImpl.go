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

// ReservedAliasStorageImplABI is the input ABI used to generate the binding from.
const ReservedAliasStorageImplABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractStorageAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"aliasString\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"aliasHash\",\"type\":\"bytes32\"}],\"name\":\"AliasReserved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"aliasString\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"aliasHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"assetbox\",\"type\":\"address\"}],\"name\":\"ReservedAliasAllowed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"oldAlias\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"oldAliasHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newAlias\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newAliasHash\",\"type\":\"bytes32\"}],\"name\":\"ReservedAliasEdited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"aliasString\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"aliasHash\",\"type\":\"bytes32\"}],\"name\":\"ReservedAliasRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"aliasString\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"aliasHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"assetbox\",\"type\":\"address\"}],\"name\":\"ReservedAliasTaken\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ACCESS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ACCESS_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ADMINABLE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ADMIN_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ASSETBOX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ASSETBOX_INFO\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ASSETBOX_INFO_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ASSETBOX_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_BALANCE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_BITBON\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_BITBON_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_BITBON_SUPPORT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_BONBALANCE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_BONTRANSFER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_DEX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_DISTRIBUTION_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_EXCHANGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_EXCHANGE_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_FEE_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_GENERATOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MINING_AGENT_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MSBON\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MSBON_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_ADD_ADDRESS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_ADD_ADMIN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_ADD_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_ADD_ROLE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_DROP_PERMISSION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_EDIT_ADMIN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_EDIT_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_EDIT_ROLE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REMOVE_ADMIN_BY_ADDRESS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REMOVE_ADMIN_BY_INDEX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REMOVE_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REMOVE_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REMOVE_ROLE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REPLENISH_CAPITALIZATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REPLENISH_CERTIFICATE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REPLENISH_EMISSION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_SET_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_SET_PERMISSION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_OTC\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_OTC_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_REPLENISH_CAPITALIZATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_RESERVED_ALIAS_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ROLE_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_SAFE_TRANSFER_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_TRANSFER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_TRANSFER_FROM_CAPITALIZATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_ACCESS_RESTORATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_ADMIN_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_CONTRACTS_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_CONTRACT_DEPLOY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_EMISSION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_FEES_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_MULTISIG_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_ROLES_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_ACCESS_VERIFIER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_BITBON_ISSUE_VERIFIER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_COMMISSION_VERIFIER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_DEPLOY_ADMIN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_PERMISSION_ADMIN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"aliasAllowance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"aliasIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"aliases\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"contractStorage\",\"outputs\":[{\"internalType\":\"contractContractStorage\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getThisContractIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"length\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"aliasString\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"assetbox\",\"type\":\"address\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"aliasString\",\"type\":\"string\"}],\"name\":\"allowedAssetbox\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"aliasString\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"assetbox\",\"type\":\"address\"}],\"name\":\"isAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getAt\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAllAliases\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAllSanitizedAliases\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"aliasString\",\"type\":\"string\"}],\"name\":\"containsAlias\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"aliasString\",\"type\":\"string\"}],\"name\":\"indexOfAlias\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"aliasString\",\"type\":\"string\"}],\"name\":\"reserveAlias\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"oldAlias\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"newAlias\",\"type\":\"string\"}],\"name\":\"editReservedAlias\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"aliasString\",\"type\":\"string\"}],\"name\":\"removeAliasReservation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"removeAliasReservationAt\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ReservedAliasStorageImpl is an auto generated Go binding around an Ethereum contract.
type ReservedAliasStorageImpl struct {
	ReservedAliasStorageImplCaller     // Read-only binding to the contract
	ReservedAliasStorageImplTransactor // Write-only binding to the contract
	ReservedAliasStorageImplFilterer   // Log filterer for contract events
}

// ReservedAliasStorageImplCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReservedAliasStorageImplCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReservedAliasStorageImplTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReservedAliasStorageImplTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReservedAliasStorageImplFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReservedAliasStorageImplFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReservedAliasStorageImplSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReservedAliasStorageImplSession struct {
	Contract     *ReservedAliasStorageImpl // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ReservedAliasStorageImplCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReservedAliasStorageImplCallerSession struct {
	Contract *ReservedAliasStorageImplCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// ReservedAliasStorageImplTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReservedAliasStorageImplTransactorSession struct {
	Contract     *ReservedAliasStorageImplTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// ReservedAliasStorageImplRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReservedAliasStorageImplRaw struct {
	Contract *ReservedAliasStorageImpl // Generic contract binding to access the raw methods on
}

// ReservedAliasStorageImplCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReservedAliasStorageImplCallerRaw struct {
	Contract *ReservedAliasStorageImplCaller // Generic read-only contract binding to access the raw methods on
}

// ReservedAliasStorageImplTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReservedAliasStorageImplTransactorRaw struct {
	Contract *ReservedAliasStorageImplTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReservedAliasStorageImpl creates a new instance of ReservedAliasStorageImpl, bound to a specific deployed contract.
func NewReservedAliasStorageImpl(address common.Address, backend bind.ContractBackend) (*ReservedAliasStorageImpl, error) {
	contract, err := bindReservedAliasStorageImpl(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ReservedAliasStorageImpl{ReservedAliasStorageImplCaller: ReservedAliasStorageImplCaller{contract: contract}, ReservedAliasStorageImplTransactor: ReservedAliasStorageImplTransactor{contract: contract}, ReservedAliasStorageImplFilterer: ReservedAliasStorageImplFilterer{contract: contract}}, nil
}

// NewReservedAliasStorageImplCaller creates a new read-only instance of ReservedAliasStorageImpl, bound to a specific deployed contract.
func NewReservedAliasStorageImplCaller(address common.Address, caller bind.ContractCaller) (*ReservedAliasStorageImplCaller, error) {
	contract, err := bindReservedAliasStorageImpl(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReservedAliasStorageImplCaller{contract: contract}, nil
}

// NewReservedAliasStorageImplTransactor creates a new write-only instance of ReservedAliasStorageImpl, bound to a specific deployed contract.
func NewReservedAliasStorageImplTransactor(address common.Address, transactor bind.ContractTransactor) (*ReservedAliasStorageImplTransactor, error) {
	contract, err := bindReservedAliasStorageImpl(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReservedAliasStorageImplTransactor{contract: contract}, nil
}

// NewReservedAliasStorageImplFilterer creates a new log filterer instance of ReservedAliasStorageImpl, bound to a specific deployed contract.
func NewReservedAliasStorageImplFilterer(address common.Address, filterer bind.ContractFilterer) (*ReservedAliasStorageImplFilterer, error) {
	contract, err := bindReservedAliasStorageImpl(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReservedAliasStorageImplFilterer{contract: contract}, nil
}

// bindReservedAliasStorageImpl binds a generic wrapper to an already deployed contract.
func bindReservedAliasStorageImpl(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ReservedAliasStorageImplABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ReservedAliasStorageImpl.Contract.ReservedAliasStorageImplCaller.contract.Call(opts, result, method, params...)
}

// Return Abi for ReservedAliasStorageImplCaller
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCaller) ABI() abi.ABI {
	return _ReservedAliasStorageImpl.contract.Abi
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReservedAliasStorageImpl.Contract.ReservedAliasStorageImplTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReservedAliasStorageImpl.Contract.ReservedAliasStorageImplTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ReservedAliasStorageImpl.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReservedAliasStorageImpl.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReservedAliasStorageImpl.Contract.contract.Transact(opts, method, params...)
}

// CONTRACTACCESS is a free data retrieval call binding the contract method 0x0ac62b0c.
//
// Solidity: function CONTRACT_ACCESS() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCaller) CONTRACTACCESS(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReservedAliasStorageImpl.contract.Call(opts, out, "CONTRACT_ACCESS")
	return *ret0, err
}

// CONTRACTACCESS is a free data retrieval call binding the contract method 0x0ac62b0c.
//
// Solidity: function CONTRACT_ACCESS() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplSession) CONTRACTACCESS() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.CONTRACTACCESS(&_ReservedAliasStorageImpl.CallOpts)
}

// CONTRACTACCESS is a free data retrieval call binding the contract method 0x0ac62b0c.
//
// Solidity: function CONTRACT_ACCESS() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCallerSession) CONTRACTACCESS() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.CONTRACTACCESS(&_ReservedAliasStorageImpl.CallOpts)
}

// CONTRACTACCESSSTORAGE is a free data retrieval call binding the contract method 0xceac9bd5.
//
// Solidity: function CONTRACT_ACCESS_STORAGE() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCaller) CONTRACTACCESSSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReservedAliasStorageImpl.contract.Call(opts, out, "CONTRACT_ACCESS_STORAGE")
	return *ret0, err
}

// CONTRACTACCESSSTORAGE is a free data retrieval call binding the contract method 0xceac9bd5.
//
// Solidity: function CONTRACT_ACCESS_STORAGE() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplSession) CONTRACTACCESSSTORAGE() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.CONTRACTACCESSSTORAGE(&_ReservedAliasStorageImpl.CallOpts)
}

// CONTRACTACCESSSTORAGE is a free data retrieval call binding the contract method 0xceac9bd5.
//
// Solidity: function CONTRACT_ACCESS_STORAGE() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCallerSession) CONTRACTACCESSSTORAGE() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.CONTRACTACCESSSTORAGE(&_ReservedAliasStorageImpl.CallOpts)
}

// CONTRACTADMINABLE is a free data retrieval call binding the contract method 0xbbe3a78f.
//
// Solidity: function CONTRACT_ADMINABLE() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCaller) CONTRACTADMINABLE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReservedAliasStorageImpl.contract.Call(opts, out, "CONTRACT_ADMINABLE")
	return *ret0, err
}

// CONTRACTADMINABLE is a free data retrieval call binding the contract method 0xbbe3a78f.
//
// Solidity: function CONTRACT_ADMINABLE() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplSession) CONTRACTADMINABLE() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.CONTRACTADMINABLE(&_ReservedAliasStorageImpl.CallOpts)
}

// CONTRACTADMINABLE is a free data retrieval call binding the contract method 0xbbe3a78f.
//
// Solidity: function CONTRACT_ADMINABLE() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCallerSession) CONTRACTADMINABLE() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.CONTRACTADMINABLE(&_ReservedAliasStorageImpl.CallOpts)
}

// CONTRACTADMINSTORAGE is a free data retrieval call binding the contract method 0xc9aebae9.
//
// Solidity: function CONTRACT_ADMIN_STORAGE() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCaller) CONTRACTADMINSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReservedAliasStorageImpl.contract.Call(opts, out, "CONTRACT_ADMIN_STORAGE")
	return *ret0, err
}

// CONTRACTADMINSTORAGE is a free data retrieval call binding the contract method 0xc9aebae9.
//
// Solidity: function CONTRACT_ADMIN_STORAGE() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplSession) CONTRACTADMINSTORAGE() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.CONTRACTADMINSTORAGE(&_ReservedAliasStorageImpl.CallOpts)
}

// CONTRACTADMINSTORAGE is a free data retrieval call binding the contract method 0xc9aebae9.
//
// Solidity: function CONTRACT_ADMIN_STORAGE() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCallerSession) CONTRACTADMINSTORAGE() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.CONTRACTADMINSTORAGE(&_ReservedAliasStorageImpl.CallOpts)
}

// CONTRACTASSETBOX is a free data retrieval call binding the contract method 0x161aa17c.
//
// Solidity: function CONTRACT_ASSETBOX() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCaller) CONTRACTASSETBOX(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReservedAliasStorageImpl.contract.Call(opts, out, "CONTRACT_ASSETBOX")
	return *ret0, err
}

// CONTRACTASSETBOX is a free data retrieval call binding the contract method 0x161aa17c.
//
// Solidity: function CONTRACT_ASSETBOX() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplSession) CONTRACTASSETBOX() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.CONTRACTASSETBOX(&_ReservedAliasStorageImpl.CallOpts)
}

// CONTRACTASSETBOX is a free data retrieval call binding the contract method 0x161aa17c.
//
// Solidity: function CONTRACT_ASSETBOX() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCallerSession) CONTRACTASSETBOX() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.CONTRACTASSETBOX(&_ReservedAliasStorageImpl.CallOpts)
}

// CONTRACTASSETBOXINFO is a free data retrieval call binding the contract method 0xb572fd3d.
//
// Solidity: function CONTRACT_ASSETBOX_INFO() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCaller) CONTRACTASSETBOXINFO(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReservedAliasStorageImpl.contract.Call(opts, out, "CONTRACT_ASSETBOX_INFO")
	return *ret0, err
}

// CONTRACTASSETBOXINFO is a free data retrieval call binding the contract method 0xb572fd3d.
//
// Solidity: function CONTRACT_ASSETBOX_INFO() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplSession) CONTRACTASSETBOXINFO() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.CONTRACTASSETBOXINFO(&_ReservedAliasStorageImpl.CallOpts)
}

// CONTRACTASSETBOXINFO is a free data retrieval call binding the contract method 0xb572fd3d.
//
// Solidity: function CONTRACT_ASSETBOX_INFO() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCallerSession) CONTRACTASSETBOXINFO() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.CONTRACTASSETBOXINFO(&_ReservedAliasStorageImpl.CallOpts)
}

// CONTRACTASSETBOXINFOSTORAGE is a free data retrieval call binding the contract method 0xf32a115d.
//
// Solidity: function CONTRACT_ASSETBOX_INFO_STORAGE() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCaller) CONTRACTASSETBOXINFOSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReservedAliasStorageImpl.contract.Call(opts, out, "CONTRACT_ASSETBOX_INFO_STORAGE")
	return *ret0, err
}

// CONTRACTASSETBOXINFOSTORAGE is a free data retrieval call binding the contract method 0xf32a115d.
//
// Solidity: function CONTRACT_ASSETBOX_INFO_STORAGE() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplSession) CONTRACTASSETBOXINFOSTORAGE() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.CONTRACTASSETBOXINFOSTORAGE(&_ReservedAliasStorageImpl.CallOpts)
}

// CONTRACTASSETBOXINFOSTORAGE is a free data retrieval call binding the contract method 0xf32a115d.
//
// Solidity: function CONTRACT_ASSETBOX_INFO_STORAGE() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCallerSession) CONTRACTASSETBOXINFOSTORAGE() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.CONTRACTASSETBOXINFOSTORAGE(&_ReservedAliasStorageImpl.CallOpts)
}

// CONTRACTASSETBOXSTORAGE is a free data retrieval call binding the contract method 0xfa5bdafc.
//
// Solidity: function CONTRACT_ASSETBOX_STORAGE() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCaller) CONTRACTASSETBOXSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReservedAliasStorageImpl.contract.Call(opts, out, "CONTRACT_ASSETBOX_STORAGE")
	return *ret0, err
}

// CONTRACTASSETBOXSTORAGE is a free data retrieval call binding the contract method 0xfa5bdafc.
//
// Solidity: function CONTRACT_ASSETBOX_STORAGE() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplSession) CONTRACTASSETBOXSTORAGE() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.CONTRACTASSETBOXSTORAGE(&_ReservedAliasStorageImpl.CallOpts)
}

// CONTRACTASSETBOXSTORAGE is a free data retrieval call binding the contract method 0xfa5bdafc.
//
// Solidity: function CONTRACT_ASSETBOX_STORAGE() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCallerSession) CONTRACTASSETBOXSTORAGE() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.CONTRACTASSETBOXSTORAGE(&_ReservedAliasStorageImpl.CallOpts)
}

// CONTRACTBALANCE is a free data retrieval call binding the contract method 0xd59045f5.
//
// Solidity: function CONTRACT_BALANCE() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCaller) CONTRACTBALANCE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReservedAliasStorageImpl.contract.Call(opts, out, "CONTRACT_BALANCE")
	return *ret0, err
}

// CONTRACTBALANCE is a free data retrieval call binding the contract method 0xd59045f5.
//
// Solidity: function CONTRACT_BALANCE() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplSession) CONTRACTBALANCE() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.CONTRACTBALANCE(&_ReservedAliasStorageImpl.CallOpts)
}

// CONTRACTBALANCE is a free data retrieval call binding the contract method 0xd59045f5.
//
// Solidity: function CONTRACT_BALANCE() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCallerSession) CONTRACTBALANCE() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.CONTRACTBALANCE(&_ReservedAliasStorageImpl.CallOpts)
}

// CONTRACTBITBON is a free data retrieval call binding the contract method 0xb6b433ab.
//
// Solidity: function CONTRACT_BITBON() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCaller) CONTRACTBITBON(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReservedAliasStorageImpl.contract.Call(opts, out, "CONTRACT_BITBON")
	return *ret0, err
}

// CONTRACTBITBON is a free data retrieval call binding the contract method 0xb6b433ab.
//
// Solidity: function CONTRACT_BITBON() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplSession) CONTRACTBITBON() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.CONTRACTBITBON(&_ReservedAliasStorageImpl.CallOpts)
}

// CONTRACTBITBON is a free data retrieval call binding the contract method 0xb6b433ab.
//
// Solidity: function CONTRACT_BITBON() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCallerSession) CONTRACTBITBON() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.CONTRACTBITBON(&_ReservedAliasStorageImpl.CallOpts)
}

// CONTRACTBITBONSTORAGE is a free data retrieval call binding the contract method 0xe7f6c30e.
//
// Solidity: function CONTRACT_BITBON_STORAGE() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCaller) CONTRACTBITBONSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReservedAliasStorageImpl.contract.Call(opts, out, "CONTRACT_BITBON_STORAGE")
	return *ret0, err
}

// CONTRACTBITBONSTORAGE is a free data retrieval call binding the contract method 0xe7f6c30e.
//
// Solidity: function CONTRACT_BITBON_STORAGE() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplSession) CONTRACTBITBONSTORAGE() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.CONTRACTBITBONSTORAGE(&_ReservedAliasStorageImpl.CallOpts)
}

// CONTRACTBITBONSTORAGE is a free data retrieval call binding the contract method 0xe7f6c30e.
//
// Solidity: function CONTRACT_BITBON_STORAGE() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCallerSession) CONTRACTBITBONSTORAGE() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.CONTRACTBITBONSTORAGE(&_ReservedAliasStorageImpl.CallOpts)
}

// CONTRACTBITBONSUPPORT is a free data retrieval call binding the contract method 0x20139c74.
//
// Solidity: function CONTRACT_BITBON_SUPPORT() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCaller) CONTRACTBITBONSUPPORT(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReservedAliasStorageImpl.contract.Call(opts, out, "CONTRACT_BITBON_SUPPORT")
	return *ret0, err
}

// CONTRACTBITBONSUPPORT is a free data retrieval call binding the contract method 0x20139c74.
//
// Solidity: function CONTRACT_BITBON_SUPPORT() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplSession) CONTRACTBITBONSUPPORT() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.CONTRACTBITBONSUPPORT(&_ReservedAliasStorageImpl.CallOpts)
}

// CONTRACTBITBONSUPPORT is a free data retrieval call binding the contract method 0x20139c74.
//
// Solidity: function CONTRACT_BITBON_SUPPORT() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCallerSession) CONTRACTBITBONSUPPORT() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.CONTRACTBITBONSUPPORT(&_ReservedAliasStorageImpl.CallOpts)
}

// CONTRACTBONBALANCE is a free data retrieval call binding the contract method 0x40e02b7c.
//
// Solidity: function CONTRACT_BONBALANCE() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCaller) CONTRACTBONBALANCE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReservedAliasStorageImpl.contract.Call(opts, out, "CONTRACT_BONBALANCE")
	return *ret0, err
}

// CONTRACTBONBALANCE is a free data retrieval call binding the contract method 0x40e02b7c.
//
// Solidity: function CONTRACT_BONBALANCE() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplSession) CONTRACTBONBALANCE() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.CONTRACTBONBALANCE(&_ReservedAliasStorageImpl.CallOpts)
}

// CONTRACTBONBALANCE is a free data retrieval call binding the contract method 0x40e02b7c.
//
// Solidity: function CONTRACT_BONBALANCE() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCallerSession) CONTRACTBONBALANCE() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.CONTRACTBONBALANCE(&_ReservedAliasStorageImpl.CallOpts)
}

// CONTRACTBONTRANSFER is a free data retrieval call binding the contract method 0xc18f4a8a.
//
// Solidity: function CONTRACT_BONTRANSFER() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCaller) CONTRACTBONTRANSFER(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReservedAliasStorageImpl.contract.Call(opts, out, "CONTRACT_BONTRANSFER")
	return *ret0, err
}

// CONTRACTBONTRANSFER is a free data retrieval call binding the contract method 0xc18f4a8a.
//
// Solidity: function CONTRACT_BONTRANSFER() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplSession) CONTRACTBONTRANSFER() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.CONTRACTBONTRANSFER(&_ReservedAliasStorageImpl.CallOpts)
}

// CONTRACTBONTRANSFER is a free data retrieval call binding the contract method 0xc18f4a8a.
//
// Solidity: function CONTRACT_BONTRANSFER() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCallerSession) CONTRACTBONTRANSFER() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.CONTRACTBONTRANSFER(&_ReservedAliasStorageImpl.CallOpts)
}

// CONTRACTDEX is a free data retrieval call binding the contract method 0x52c1f844.
//
// Solidity: function CONTRACT_DEX() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCaller) CONTRACTDEX(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReservedAliasStorageImpl.contract.Call(opts, out, "CONTRACT_DEX")
	return *ret0, err
}

// CONTRACTDEX is a free data retrieval call binding the contract method 0x52c1f844.
//
// Solidity: function CONTRACT_DEX() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplSession) CONTRACTDEX() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.CONTRACTDEX(&_ReservedAliasStorageImpl.CallOpts)
}

// CONTRACTDEX is a free data retrieval call binding the contract method 0x52c1f844.
//
// Solidity: function CONTRACT_DEX() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCallerSession) CONTRACTDEX() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.CONTRACTDEX(&_ReservedAliasStorageImpl.CallOpts)
}

// CONTRACTDISTRIBUTIONSTORAGE is a free data retrieval call binding the contract method 0x08097327.
//
// Solidity: function CONTRACT_DISTRIBUTION_STORAGE() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCaller) CONTRACTDISTRIBUTIONSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReservedAliasStorageImpl.contract.Call(opts, out, "CONTRACT_DISTRIBUTION_STORAGE")
	return *ret0, err
}

// CONTRACTDISTRIBUTIONSTORAGE is a free data retrieval call binding the contract method 0x08097327.
//
// Solidity: function CONTRACT_DISTRIBUTION_STORAGE() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplSession) CONTRACTDISTRIBUTIONSTORAGE() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.CONTRACTDISTRIBUTIONSTORAGE(&_ReservedAliasStorageImpl.CallOpts)
}

// CONTRACTDISTRIBUTIONSTORAGE is a free data retrieval call binding the contract method 0x08097327.
//
// Solidity: function CONTRACT_DISTRIBUTION_STORAGE() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCallerSession) CONTRACTDISTRIBUTIONSTORAGE() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.CONTRACTDISTRIBUTIONSTORAGE(&_ReservedAliasStorageImpl.CallOpts)
}

// CONTRACTEXCHANGE is a free data retrieval call binding the contract method 0xa0e3a10e.
//
// Solidity: function CONTRACT_EXCHANGE() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCaller) CONTRACTEXCHANGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReservedAliasStorageImpl.contract.Call(opts, out, "CONTRACT_EXCHANGE")
	return *ret0, err
}

// CONTRACTEXCHANGE is a free data retrieval call binding the contract method 0xa0e3a10e.
//
// Solidity: function CONTRACT_EXCHANGE() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplSession) CONTRACTEXCHANGE() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.CONTRACTEXCHANGE(&_ReservedAliasStorageImpl.CallOpts)
}

// CONTRACTEXCHANGE is a free data retrieval call binding the contract method 0xa0e3a10e.
//
// Solidity: function CONTRACT_EXCHANGE() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCallerSession) CONTRACTEXCHANGE() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.CONTRACTEXCHANGE(&_ReservedAliasStorageImpl.CallOpts)
}

// CONTRACTEXCHANGESTORAGE is a free data retrieval call binding the contract method 0x001570f2.
//
// Solidity: function CONTRACT_EXCHANGE_STORAGE() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCaller) CONTRACTEXCHANGESTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReservedAliasStorageImpl.contract.Call(opts, out, "CONTRACT_EXCHANGE_STORAGE")
	return *ret0, err
}

// CONTRACTEXCHANGESTORAGE is a free data retrieval call binding the contract method 0x001570f2.
//
// Solidity: function CONTRACT_EXCHANGE_STORAGE() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplSession) CONTRACTEXCHANGESTORAGE() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.CONTRACTEXCHANGESTORAGE(&_ReservedAliasStorageImpl.CallOpts)
}

// CONTRACTEXCHANGESTORAGE is a free data retrieval call binding the contract method 0x001570f2.
//
// Solidity: function CONTRACT_EXCHANGE_STORAGE() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCallerSession) CONTRACTEXCHANGESTORAGE() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.CONTRACTEXCHANGESTORAGE(&_ReservedAliasStorageImpl.CallOpts)
}

// CONTRACTFEESTORAGE is a free data retrieval call binding the contract method 0xcd5e65a4.
//
// Solidity: function CONTRACT_FEE_STORAGE() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCaller) CONTRACTFEESTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReservedAliasStorageImpl.contract.Call(opts, out, "CONTRACT_FEE_STORAGE")
	return *ret0, err
}

// CONTRACTFEESTORAGE is a free data retrieval call binding the contract method 0xcd5e65a4.
//
// Solidity: function CONTRACT_FEE_STORAGE() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplSession) CONTRACTFEESTORAGE() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.CONTRACTFEESTORAGE(&_ReservedAliasStorageImpl.CallOpts)
}

// CONTRACTFEESTORAGE is a free data retrieval call binding the contract method 0xcd5e65a4.
//
// Solidity: function CONTRACT_FEE_STORAGE() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCallerSession) CONTRACTFEESTORAGE() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.CONTRACTFEESTORAGE(&_ReservedAliasStorageImpl.CallOpts)
}

// CONTRACTGENERATOR is a free data retrieval call binding the contract method 0xccea2b46.
//
// Solidity: function CONTRACT_GENERATOR() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCaller) CONTRACTGENERATOR(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReservedAliasStorageImpl.contract.Call(opts, out, "CONTRACT_GENERATOR")
	return *ret0, err
}

// CONTRACTGENERATOR is a free data retrieval call binding the contract method 0xccea2b46.
//
// Solidity: function CONTRACT_GENERATOR() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplSession) CONTRACTGENERATOR() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.CONTRACTGENERATOR(&_ReservedAliasStorageImpl.CallOpts)
}

// CONTRACTGENERATOR is a free data retrieval call binding the contract method 0xccea2b46.
//
// Solidity: function CONTRACT_GENERATOR() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCallerSession) CONTRACTGENERATOR() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.CONTRACTGENERATOR(&_ReservedAliasStorageImpl.CallOpts)
}

// CONTRACTMININGAGENTSTORAGE is a free data retrieval call binding the contract method 0x83cc588d.
//
// Solidity: function CONTRACT_MINING_AGENT_STORAGE() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCaller) CONTRACTMININGAGENTSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReservedAliasStorageImpl.contract.Call(opts, out, "CONTRACT_MINING_AGENT_STORAGE")
	return *ret0, err
}

// CONTRACTMININGAGENTSTORAGE is a free data retrieval call binding the contract method 0x83cc588d.
//
// Solidity: function CONTRACT_MINING_AGENT_STORAGE() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplSession) CONTRACTMININGAGENTSTORAGE() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.CONTRACTMININGAGENTSTORAGE(&_ReservedAliasStorageImpl.CallOpts)
}

// CONTRACTMININGAGENTSTORAGE is a free data retrieval call binding the contract method 0x83cc588d.
//
// Solidity: function CONTRACT_MINING_AGENT_STORAGE() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCallerSession) CONTRACTMININGAGENTSTORAGE() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.CONTRACTMININGAGENTSTORAGE(&_ReservedAliasStorageImpl.CallOpts)
}

// CONTRACTMSBON is a free data retrieval call binding the contract method 0x28b3a976.
//
// Solidity: function CONTRACT_MSBON() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCaller) CONTRACTMSBON(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReservedAliasStorageImpl.contract.Call(opts, out, "CONTRACT_MSBON")
	return *ret0, err
}

// CONTRACTMSBON is a free data retrieval call binding the contract method 0x28b3a976.
//
// Solidity: function CONTRACT_MSBON() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplSession) CONTRACTMSBON() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.CONTRACTMSBON(&_ReservedAliasStorageImpl.CallOpts)
}

// CONTRACTMSBON is a free data retrieval call binding the contract method 0x28b3a976.
//
// Solidity: function CONTRACT_MSBON() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCallerSession) CONTRACTMSBON() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.CONTRACTMSBON(&_ReservedAliasStorageImpl.CallOpts)
}

// CONTRACTMSBONSTORAGE is a free data retrieval call binding the contract method 0xe04e6061.
//
// Solidity: function CONTRACT_MSBON_STORAGE() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCaller) CONTRACTMSBONSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReservedAliasStorageImpl.contract.Call(opts, out, "CONTRACT_MSBON_STORAGE")
	return *ret0, err
}

// CONTRACTMSBONSTORAGE is a free data retrieval call binding the contract method 0xe04e6061.
//
// Solidity: function CONTRACT_MSBON_STORAGE() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplSession) CONTRACTMSBONSTORAGE() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.CONTRACTMSBONSTORAGE(&_ReservedAliasStorageImpl.CallOpts)
}

// CONTRACTMSBONSTORAGE is a free data retrieval call binding the contract method 0xe04e6061.
//
// Solidity: function CONTRACT_MSBON_STORAGE() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCallerSession) CONTRACTMSBONSTORAGE() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.CONTRACTMSBONSTORAGE(&_ReservedAliasStorageImpl.CallOpts)
}

// CONTRACTMULTISIGADDADDRESS is a free data retrieval call binding the contract method 0xf4f65a5a.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ADDRESS() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCaller) CONTRACTMULTISIGADDADDRESS(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReservedAliasStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_ADD_ADDRESS")
	return *ret0, err
}

// CONTRACTMULTISIGADDADDRESS is a free data retrieval call binding the contract method 0xf4f65a5a.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ADDRESS() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplSession) CONTRACTMULTISIGADDADDRESS() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.CONTRACTMULTISIGADDADDRESS(&_ReservedAliasStorageImpl.CallOpts)
}

// CONTRACTMULTISIGADDADDRESS is a free data retrieval call binding the contract method 0xf4f65a5a.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ADDRESS() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCallerSession) CONTRACTMULTISIGADDADDRESS() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.CONTRACTMULTISIGADDADDRESS(&_ReservedAliasStorageImpl.CallOpts)
}

// CONTRACTMULTISIGADDADMIN is a free data retrieval call binding the contract method 0x83888c68.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ADMIN() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCaller) CONTRACTMULTISIGADDADMIN(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReservedAliasStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_ADD_ADMIN")
	return *ret0, err
}

// CONTRACTMULTISIGADDADMIN is a free data retrieval call binding the contract method 0x83888c68.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ADMIN() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplSession) CONTRACTMULTISIGADDADMIN() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.CONTRACTMULTISIGADDADMIN(&_ReservedAliasStorageImpl.CallOpts)
}

// CONTRACTMULTISIGADDADMIN is a free data retrieval call binding the contract method 0x83888c68.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ADMIN() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCallerSession) CONTRACTMULTISIGADDADMIN() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.CONTRACTMULTISIGADDADMIN(&_ReservedAliasStorageImpl.CallOpts)
}

// CONTRACTMULTISIGADDFEE is a free data retrieval call binding the contract method 0xc43e8cab.
//
// Solidity: function CONTRACT_MULTISIG_ADD_FEE() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCaller) CONTRACTMULTISIGADDFEE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReservedAliasStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_ADD_FEE")
	return *ret0, err
}

// CONTRACTMULTISIGADDFEE is a free data retrieval call binding the contract method 0xc43e8cab.
//
// Solidity: function CONTRACT_MULTISIG_ADD_FEE() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplSession) CONTRACTMULTISIGADDFEE() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.CONTRACTMULTISIGADDFEE(&_ReservedAliasStorageImpl.CallOpts)
}

// CONTRACTMULTISIGADDFEE is a free data retrieval call binding the contract method 0xc43e8cab.
//
// Solidity: function CONTRACT_MULTISIG_ADD_FEE() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCallerSession) CONTRACTMULTISIGADDFEE() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.CONTRACTMULTISIGADDFEE(&_ReservedAliasStorageImpl.CallOpts)
}

// CONTRACTMULTISIGADDROLE is a free data retrieval call binding the contract method 0x20a0914b.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ROLE() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCaller) CONTRACTMULTISIGADDROLE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReservedAliasStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_ADD_ROLE")
	return *ret0, err
}

// CONTRACTMULTISIGADDROLE is a free data retrieval call binding the contract method 0x20a0914b.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ROLE() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplSession) CONTRACTMULTISIGADDROLE() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.CONTRACTMULTISIGADDROLE(&_ReservedAliasStorageImpl.CallOpts)
}

// CONTRACTMULTISIGADDROLE is a free data retrieval call binding the contract method 0x20a0914b.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ROLE() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCallerSession) CONTRACTMULTISIGADDROLE() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.CONTRACTMULTISIGADDROLE(&_ReservedAliasStorageImpl.CallOpts)
}

// CONTRACTMULTISIGDROPPERMISSION is a free data retrieval call binding the contract method 0x3f3fbc39.
//
// Solidity: function CONTRACT_MULTISIG_DROP_PERMISSION() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCaller) CONTRACTMULTISIGDROPPERMISSION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReservedAliasStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_DROP_PERMISSION")
	return *ret0, err
}

// CONTRACTMULTISIGDROPPERMISSION is a free data retrieval call binding the contract method 0x3f3fbc39.
//
// Solidity: function CONTRACT_MULTISIG_DROP_PERMISSION() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplSession) CONTRACTMULTISIGDROPPERMISSION() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.CONTRACTMULTISIGDROPPERMISSION(&_ReservedAliasStorageImpl.CallOpts)
}

// CONTRACTMULTISIGDROPPERMISSION is a free data retrieval call binding the contract method 0x3f3fbc39.
//
// Solidity: function CONTRACT_MULTISIG_DROP_PERMISSION() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCallerSession) CONTRACTMULTISIGDROPPERMISSION() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.CONTRACTMULTISIGDROPPERMISSION(&_ReservedAliasStorageImpl.CallOpts)
}

// CONTRACTMULTISIGEDITADMIN is a free data retrieval call binding the contract method 0x91b0f051.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_ADMIN() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCaller) CONTRACTMULTISIGEDITADMIN(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReservedAliasStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_EDIT_ADMIN")
	return *ret0, err
}

// CONTRACTMULTISIGEDITADMIN is a free data retrieval call binding the contract method 0x91b0f051.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_ADMIN() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplSession) CONTRACTMULTISIGEDITADMIN() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.CONTRACTMULTISIGEDITADMIN(&_ReservedAliasStorageImpl.CallOpts)
}

// CONTRACTMULTISIGEDITADMIN is a free data retrieval call binding the contract method 0x91b0f051.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_ADMIN() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCallerSession) CONTRACTMULTISIGEDITADMIN() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.CONTRACTMULTISIGEDITADMIN(&_ReservedAliasStorageImpl.CallOpts)
}

// CONTRACTMULTISIGEDITFEE is a free data retrieval call binding the contract method 0x6b2abe17.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_FEE() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCaller) CONTRACTMULTISIGEDITFEE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReservedAliasStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_EDIT_FEE")
	return *ret0, err
}

// CONTRACTMULTISIGEDITFEE is a free data retrieval call binding the contract method 0x6b2abe17.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_FEE() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplSession) CONTRACTMULTISIGEDITFEE() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.CONTRACTMULTISIGEDITFEE(&_ReservedAliasStorageImpl.CallOpts)
}

// CONTRACTMULTISIGEDITFEE is a free data retrieval call binding the contract method 0x6b2abe17.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_FEE() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCallerSession) CONTRACTMULTISIGEDITFEE() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.CONTRACTMULTISIGEDITFEE(&_ReservedAliasStorageImpl.CallOpts)
}

// CONTRACTMULTISIGEDITROLE is a free data retrieval call binding the contract method 0xc5f66412.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_ROLE() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCaller) CONTRACTMULTISIGEDITROLE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReservedAliasStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_EDIT_ROLE")
	return *ret0, err
}

// CONTRACTMULTISIGEDITROLE is a free data retrieval call binding the contract method 0xc5f66412.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_ROLE() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplSession) CONTRACTMULTISIGEDITROLE() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.CONTRACTMULTISIGEDITROLE(&_ReservedAliasStorageImpl.CallOpts)
}

// CONTRACTMULTISIGEDITROLE is a free data retrieval call binding the contract method 0xc5f66412.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_ROLE() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCallerSession) CONTRACTMULTISIGEDITROLE() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.CONTRACTMULTISIGEDITROLE(&_ReservedAliasStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEADMINBYADDRESS is a free data retrieval call binding the contract method 0xc816b091.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_BY_ADDRESS() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCaller) CONTRACTMULTISIGREMOVEADMINBYADDRESS(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReservedAliasStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REMOVE_ADMIN_BY_ADDRESS")
	return *ret0, err
}

// CONTRACTMULTISIGREMOVEADMINBYADDRESS is a free data retrieval call binding the contract method 0xc816b091.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_BY_ADDRESS() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplSession) CONTRACTMULTISIGREMOVEADMINBYADDRESS() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.CONTRACTMULTISIGREMOVEADMINBYADDRESS(&_ReservedAliasStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEADMINBYADDRESS is a free data retrieval call binding the contract method 0xc816b091.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_BY_ADDRESS() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCallerSession) CONTRACTMULTISIGREMOVEADMINBYADDRESS() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.CONTRACTMULTISIGREMOVEADMINBYADDRESS(&_ReservedAliasStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEADMINBYINDEX is a free data retrieval call binding the contract method 0x751fbea2.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_BY_INDEX() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCaller) CONTRACTMULTISIGREMOVEADMINBYINDEX(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReservedAliasStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REMOVE_ADMIN_BY_INDEX")
	return *ret0, err
}

// CONTRACTMULTISIGREMOVEADMINBYINDEX is a free data retrieval call binding the contract method 0x751fbea2.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_BY_INDEX() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplSession) CONTRACTMULTISIGREMOVEADMINBYINDEX() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.CONTRACTMULTISIGREMOVEADMINBYINDEX(&_ReservedAliasStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEADMINBYINDEX is a free data retrieval call binding the contract method 0x751fbea2.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_BY_INDEX() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCallerSession) CONTRACTMULTISIGREMOVEADMINBYINDEX() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.CONTRACTMULTISIGREMOVEADMINBYINDEX(&_ReservedAliasStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEADMINROLE is a free data retrieval call binding the contract method 0x53c00200.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_ROLE() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCaller) CONTRACTMULTISIGREMOVEADMINROLE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReservedAliasStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REMOVE_ADMIN_ROLE")
	return *ret0, err
}

// CONTRACTMULTISIGREMOVEADMINROLE is a free data retrieval call binding the contract method 0x53c00200.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_ROLE() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplSession) CONTRACTMULTISIGREMOVEADMINROLE() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.CONTRACTMULTISIGREMOVEADMINROLE(&_ReservedAliasStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEADMINROLE is a free data retrieval call binding the contract method 0x53c00200.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_ROLE() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCallerSession) CONTRACTMULTISIGREMOVEADMINROLE() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.CONTRACTMULTISIGREMOVEADMINROLE(&_ReservedAliasStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEFEE is a free data retrieval call binding the contract method 0xdb614d25.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_FEE() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCaller) CONTRACTMULTISIGREMOVEFEE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReservedAliasStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REMOVE_FEE")
	return *ret0, err
}

// CONTRACTMULTISIGREMOVEFEE is a free data retrieval call binding the contract method 0xdb614d25.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_FEE() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplSession) CONTRACTMULTISIGREMOVEFEE() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.CONTRACTMULTISIGREMOVEFEE(&_ReservedAliasStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEFEE is a free data retrieval call binding the contract method 0xdb614d25.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_FEE() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCallerSession) CONTRACTMULTISIGREMOVEFEE() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.CONTRACTMULTISIGREMOVEFEE(&_ReservedAliasStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEROLE is a free data retrieval call binding the contract method 0x23648496.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ROLE() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCaller) CONTRACTMULTISIGREMOVEROLE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReservedAliasStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REMOVE_ROLE")
	return *ret0, err
}

// CONTRACTMULTISIGREMOVEROLE is a free data retrieval call binding the contract method 0x23648496.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ROLE() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplSession) CONTRACTMULTISIGREMOVEROLE() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.CONTRACTMULTISIGREMOVEROLE(&_ReservedAliasStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEROLE is a free data retrieval call binding the contract method 0x23648496.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ROLE() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCallerSession) CONTRACTMULTISIGREMOVEROLE() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.CONTRACTMULTISIGREMOVEROLE(&_ReservedAliasStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREPLENISHCAPITALIZATION is a free data retrieval call binding the contract method 0x62d5d523.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_CAPITALIZATION() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCaller) CONTRACTMULTISIGREPLENISHCAPITALIZATION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReservedAliasStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REPLENISH_CAPITALIZATION")
	return *ret0, err
}

// CONTRACTMULTISIGREPLENISHCAPITALIZATION is a free data retrieval call binding the contract method 0x62d5d523.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_CAPITALIZATION() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplSession) CONTRACTMULTISIGREPLENISHCAPITALIZATION() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.CONTRACTMULTISIGREPLENISHCAPITALIZATION(&_ReservedAliasStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREPLENISHCAPITALIZATION is a free data retrieval call binding the contract method 0x62d5d523.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_CAPITALIZATION() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCallerSession) CONTRACTMULTISIGREPLENISHCAPITALIZATION() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.CONTRACTMULTISIGREPLENISHCAPITALIZATION(&_ReservedAliasStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREPLENISHCERTIFICATE is a free data retrieval call binding the contract method 0xfd19b7e1.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_CERTIFICATE() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCaller) CONTRACTMULTISIGREPLENISHCERTIFICATE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReservedAliasStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REPLENISH_CERTIFICATE")
	return *ret0, err
}

// CONTRACTMULTISIGREPLENISHCERTIFICATE is a free data retrieval call binding the contract method 0xfd19b7e1.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_CERTIFICATE() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplSession) CONTRACTMULTISIGREPLENISHCERTIFICATE() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.CONTRACTMULTISIGREPLENISHCERTIFICATE(&_ReservedAliasStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREPLENISHCERTIFICATE is a free data retrieval call binding the contract method 0xfd19b7e1.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_CERTIFICATE() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCallerSession) CONTRACTMULTISIGREPLENISHCERTIFICATE() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.CONTRACTMULTISIGREPLENISHCERTIFICATE(&_ReservedAliasStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREPLENISHEMISSION is a free data retrieval call binding the contract method 0x0cb59fde.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_EMISSION() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCaller) CONTRACTMULTISIGREPLENISHEMISSION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReservedAliasStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REPLENISH_EMISSION")
	return *ret0, err
}

// CONTRACTMULTISIGREPLENISHEMISSION is a free data retrieval call binding the contract method 0x0cb59fde.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_EMISSION() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplSession) CONTRACTMULTISIGREPLENISHEMISSION() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.CONTRACTMULTISIGREPLENISHEMISSION(&_ReservedAliasStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREPLENISHEMISSION is a free data retrieval call binding the contract method 0x0cb59fde.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_EMISSION() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCallerSession) CONTRACTMULTISIGREPLENISHEMISSION() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.CONTRACTMULTISIGREPLENISHEMISSION(&_ReservedAliasStorageImpl.CallOpts)
}

// CONTRACTMULTISIGSETADMINROLE is a free data retrieval call binding the contract method 0x8f209ab7.
//
// Solidity: function CONTRACT_MULTISIG_SET_ADMIN_ROLE() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCaller) CONTRACTMULTISIGSETADMINROLE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReservedAliasStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_SET_ADMIN_ROLE")
	return *ret0, err
}

// CONTRACTMULTISIGSETADMINROLE is a free data retrieval call binding the contract method 0x8f209ab7.
//
// Solidity: function CONTRACT_MULTISIG_SET_ADMIN_ROLE() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplSession) CONTRACTMULTISIGSETADMINROLE() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.CONTRACTMULTISIGSETADMINROLE(&_ReservedAliasStorageImpl.CallOpts)
}

// CONTRACTMULTISIGSETADMINROLE is a free data retrieval call binding the contract method 0x8f209ab7.
//
// Solidity: function CONTRACT_MULTISIG_SET_ADMIN_ROLE() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCallerSession) CONTRACTMULTISIGSETADMINROLE() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.CONTRACTMULTISIGSETADMINROLE(&_ReservedAliasStorageImpl.CallOpts)
}

// CONTRACTMULTISIGSETPERMISSION is a free data retrieval call binding the contract method 0xb0e91914.
//
// Solidity: function CONTRACT_MULTISIG_SET_PERMISSION() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCaller) CONTRACTMULTISIGSETPERMISSION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReservedAliasStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_SET_PERMISSION")
	return *ret0, err
}

// CONTRACTMULTISIGSETPERMISSION is a free data retrieval call binding the contract method 0xb0e91914.
//
// Solidity: function CONTRACT_MULTISIG_SET_PERMISSION() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplSession) CONTRACTMULTISIGSETPERMISSION() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.CONTRACTMULTISIGSETPERMISSION(&_ReservedAliasStorageImpl.CallOpts)
}

// CONTRACTMULTISIGSETPERMISSION is a free data retrieval call binding the contract method 0xb0e91914.
//
// Solidity: function CONTRACT_MULTISIG_SET_PERMISSION() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCallerSession) CONTRACTMULTISIGSETPERMISSION() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.CONTRACTMULTISIGSETPERMISSION(&_ReservedAliasStorageImpl.CallOpts)
}

// CONTRACTMULTISIGSTORAGE is a free data retrieval call binding the contract method 0xe7e1a303.
//
// Solidity: function CONTRACT_MULTISIG_STORAGE() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCaller) CONTRACTMULTISIGSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReservedAliasStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_STORAGE")
	return *ret0, err
}

// CONTRACTMULTISIGSTORAGE is a free data retrieval call binding the contract method 0xe7e1a303.
//
// Solidity: function CONTRACT_MULTISIG_STORAGE() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplSession) CONTRACTMULTISIGSTORAGE() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.CONTRACTMULTISIGSTORAGE(&_ReservedAliasStorageImpl.CallOpts)
}

// CONTRACTMULTISIGSTORAGE is a free data retrieval call binding the contract method 0xe7e1a303.
//
// Solidity: function CONTRACT_MULTISIG_STORAGE() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCallerSession) CONTRACTMULTISIGSTORAGE() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.CONTRACTMULTISIGSTORAGE(&_ReservedAliasStorageImpl.CallOpts)
}

// CONTRACTOTC is a free data retrieval call binding the contract method 0x5eceac70.
//
// Solidity: function CONTRACT_OTC() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCaller) CONTRACTOTC(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReservedAliasStorageImpl.contract.Call(opts, out, "CONTRACT_OTC")
	return *ret0, err
}

// CONTRACTOTC is a free data retrieval call binding the contract method 0x5eceac70.
//
// Solidity: function CONTRACT_OTC() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplSession) CONTRACTOTC() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.CONTRACTOTC(&_ReservedAliasStorageImpl.CallOpts)
}

// CONTRACTOTC is a free data retrieval call binding the contract method 0x5eceac70.
//
// Solidity: function CONTRACT_OTC() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCallerSession) CONTRACTOTC() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.CONTRACTOTC(&_ReservedAliasStorageImpl.CallOpts)
}

// CONTRACTOTCSTORAGE is a free data retrieval call binding the contract method 0x825b476f.
//
// Solidity: function CONTRACT_OTC_STORAGE() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCaller) CONTRACTOTCSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReservedAliasStorageImpl.contract.Call(opts, out, "CONTRACT_OTC_STORAGE")
	return *ret0, err
}

// CONTRACTOTCSTORAGE is a free data retrieval call binding the contract method 0x825b476f.
//
// Solidity: function CONTRACT_OTC_STORAGE() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplSession) CONTRACTOTCSTORAGE() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.CONTRACTOTCSTORAGE(&_ReservedAliasStorageImpl.CallOpts)
}

// CONTRACTOTCSTORAGE is a free data retrieval call binding the contract method 0x825b476f.
//
// Solidity: function CONTRACT_OTC_STORAGE() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCallerSession) CONTRACTOTCSTORAGE() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.CONTRACTOTCSTORAGE(&_ReservedAliasStorageImpl.CallOpts)
}

// CONTRACTREPLENISHCAPITALIZATION is a free data retrieval call binding the contract method 0xff5e09d8.
//
// Solidity: function CONTRACT_REPLENISH_CAPITALIZATION() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCaller) CONTRACTREPLENISHCAPITALIZATION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReservedAliasStorageImpl.contract.Call(opts, out, "CONTRACT_REPLENISH_CAPITALIZATION")
	return *ret0, err
}

// CONTRACTREPLENISHCAPITALIZATION is a free data retrieval call binding the contract method 0xff5e09d8.
//
// Solidity: function CONTRACT_REPLENISH_CAPITALIZATION() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplSession) CONTRACTREPLENISHCAPITALIZATION() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.CONTRACTREPLENISHCAPITALIZATION(&_ReservedAliasStorageImpl.CallOpts)
}

// CONTRACTREPLENISHCAPITALIZATION is a free data retrieval call binding the contract method 0xff5e09d8.
//
// Solidity: function CONTRACT_REPLENISH_CAPITALIZATION() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCallerSession) CONTRACTREPLENISHCAPITALIZATION() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.CONTRACTREPLENISHCAPITALIZATION(&_ReservedAliasStorageImpl.CallOpts)
}

// CONTRACTRESERVEDALIASSTORAGE is a free data retrieval call binding the contract method 0x70a94aa7.
//
// Solidity: function CONTRACT_RESERVED_ALIAS_STORAGE() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCaller) CONTRACTRESERVEDALIASSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReservedAliasStorageImpl.contract.Call(opts, out, "CONTRACT_RESERVED_ALIAS_STORAGE")
	return *ret0, err
}

// CONTRACTRESERVEDALIASSTORAGE is a free data retrieval call binding the contract method 0x70a94aa7.
//
// Solidity: function CONTRACT_RESERVED_ALIAS_STORAGE() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplSession) CONTRACTRESERVEDALIASSTORAGE() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.CONTRACTRESERVEDALIASSTORAGE(&_ReservedAliasStorageImpl.CallOpts)
}

// CONTRACTRESERVEDALIASSTORAGE is a free data retrieval call binding the contract method 0x70a94aa7.
//
// Solidity: function CONTRACT_RESERVED_ALIAS_STORAGE() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCallerSession) CONTRACTRESERVEDALIASSTORAGE() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.CONTRACTRESERVEDALIASSTORAGE(&_ReservedAliasStorageImpl.CallOpts)
}

// CONTRACTROLESTORAGE is a free data retrieval call binding the contract method 0xb938235b.
//
// Solidity: function CONTRACT_ROLE_STORAGE() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCaller) CONTRACTROLESTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReservedAliasStorageImpl.contract.Call(opts, out, "CONTRACT_ROLE_STORAGE")
	return *ret0, err
}

// CONTRACTROLESTORAGE is a free data retrieval call binding the contract method 0xb938235b.
//
// Solidity: function CONTRACT_ROLE_STORAGE() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplSession) CONTRACTROLESTORAGE() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.CONTRACTROLESTORAGE(&_ReservedAliasStorageImpl.CallOpts)
}

// CONTRACTROLESTORAGE is a free data retrieval call binding the contract method 0xb938235b.
//
// Solidity: function CONTRACT_ROLE_STORAGE() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCallerSession) CONTRACTROLESTORAGE() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.CONTRACTROLESTORAGE(&_ReservedAliasStorageImpl.CallOpts)
}

// CONTRACTSAFETRANSFERSTORAGE is a free data retrieval call binding the contract method 0xe7e45392.
//
// Solidity: function CONTRACT_SAFE_TRANSFER_STORAGE() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCaller) CONTRACTSAFETRANSFERSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReservedAliasStorageImpl.contract.Call(opts, out, "CONTRACT_SAFE_TRANSFER_STORAGE")
	return *ret0, err
}

// CONTRACTSAFETRANSFERSTORAGE is a free data retrieval call binding the contract method 0xe7e45392.
//
// Solidity: function CONTRACT_SAFE_TRANSFER_STORAGE() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplSession) CONTRACTSAFETRANSFERSTORAGE() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.CONTRACTSAFETRANSFERSTORAGE(&_ReservedAliasStorageImpl.CallOpts)
}

// CONTRACTSAFETRANSFERSTORAGE is a free data retrieval call binding the contract method 0xe7e45392.
//
// Solidity: function CONTRACT_SAFE_TRANSFER_STORAGE() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCallerSession) CONTRACTSAFETRANSFERSTORAGE() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.CONTRACTSAFETRANSFERSTORAGE(&_ReservedAliasStorageImpl.CallOpts)
}

// CONTRACTTRANSFER is a free data retrieval call binding the contract method 0xef94fe38.
//
// Solidity: function CONTRACT_TRANSFER() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCaller) CONTRACTTRANSFER(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReservedAliasStorageImpl.contract.Call(opts, out, "CONTRACT_TRANSFER")
	return *ret0, err
}

// CONTRACTTRANSFER is a free data retrieval call binding the contract method 0xef94fe38.
//
// Solidity: function CONTRACT_TRANSFER() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplSession) CONTRACTTRANSFER() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.CONTRACTTRANSFER(&_ReservedAliasStorageImpl.CallOpts)
}

// CONTRACTTRANSFER is a free data retrieval call binding the contract method 0xef94fe38.
//
// Solidity: function CONTRACT_TRANSFER() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCallerSession) CONTRACTTRANSFER() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.CONTRACTTRANSFER(&_ReservedAliasStorageImpl.CallOpts)
}

// CONTRACTTRANSFERFROMCAPITALIZATION is a free data retrieval call binding the contract method 0xc8f8decb.
//
// Solidity: function CONTRACT_TRANSFER_FROM_CAPITALIZATION() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCaller) CONTRACTTRANSFERFROMCAPITALIZATION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReservedAliasStorageImpl.contract.Call(opts, out, "CONTRACT_TRANSFER_FROM_CAPITALIZATION")
	return *ret0, err
}

// CONTRACTTRANSFERFROMCAPITALIZATION is a free data retrieval call binding the contract method 0xc8f8decb.
//
// Solidity: function CONTRACT_TRANSFER_FROM_CAPITALIZATION() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplSession) CONTRACTTRANSFERFROMCAPITALIZATION() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.CONTRACTTRANSFERFROMCAPITALIZATION(&_ReservedAliasStorageImpl.CallOpts)
}

// CONTRACTTRANSFERFROMCAPITALIZATION is a free data retrieval call binding the contract method 0xc8f8decb.
//
// Solidity: function CONTRACT_TRANSFER_FROM_CAPITALIZATION() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCallerSession) CONTRACTTRANSFERFROMCAPITALIZATION() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.CONTRACTTRANSFERFROMCAPITALIZATION(&_ReservedAliasStorageImpl.CallOpts)
}

// PERMISSIONACCESSRESTORATION is a free data retrieval call binding the contract method 0x5c3053bf.
//
// Solidity: function PERMISSION_ACCESS_RESTORATION() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCaller) PERMISSIONACCESSRESTORATION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReservedAliasStorageImpl.contract.Call(opts, out, "PERMISSION_ACCESS_RESTORATION")
	return *ret0, err
}

// PERMISSIONACCESSRESTORATION is a free data retrieval call binding the contract method 0x5c3053bf.
//
// Solidity: function PERMISSION_ACCESS_RESTORATION() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplSession) PERMISSIONACCESSRESTORATION() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.PERMISSIONACCESSRESTORATION(&_ReservedAliasStorageImpl.CallOpts)
}

// PERMISSIONACCESSRESTORATION is a free data retrieval call binding the contract method 0x5c3053bf.
//
// Solidity: function PERMISSION_ACCESS_RESTORATION() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCallerSession) PERMISSIONACCESSRESTORATION() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.PERMISSIONACCESSRESTORATION(&_ReservedAliasStorageImpl.CallOpts)
}

// PERMISSIONADMINSTORAGE is a free data retrieval call binding the contract method 0x0c9fc0db.
//
// Solidity: function PERMISSION_ADMIN_STORAGE() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCaller) PERMISSIONADMINSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReservedAliasStorageImpl.contract.Call(opts, out, "PERMISSION_ADMIN_STORAGE")
	return *ret0, err
}

// PERMISSIONADMINSTORAGE is a free data retrieval call binding the contract method 0x0c9fc0db.
//
// Solidity: function PERMISSION_ADMIN_STORAGE() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplSession) PERMISSIONADMINSTORAGE() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.PERMISSIONADMINSTORAGE(&_ReservedAliasStorageImpl.CallOpts)
}

// PERMISSIONADMINSTORAGE is a free data retrieval call binding the contract method 0x0c9fc0db.
//
// Solidity: function PERMISSION_ADMIN_STORAGE() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCallerSession) PERMISSIONADMINSTORAGE() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.PERMISSIONADMINSTORAGE(&_ReservedAliasStorageImpl.CallOpts)
}

// PERMISSIONCONTRACTSSTORAGE is a free data retrieval call binding the contract method 0x1516e83b.
//
// Solidity: function PERMISSION_CONTRACTS_STORAGE() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCaller) PERMISSIONCONTRACTSSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReservedAliasStorageImpl.contract.Call(opts, out, "PERMISSION_CONTRACTS_STORAGE")
	return *ret0, err
}

// PERMISSIONCONTRACTSSTORAGE is a free data retrieval call binding the contract method 0x1516e83b.
//
// Solidity: function PERMISSION_CONTRACTS_STORAGE() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplSession) PERMISSIONCONTRACTSSTORAGE() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.PERMISSIONCONTRACTSSTORAGE(&_ReservedAliasStorageImpl.CallOpts)
}

// PERMISSIONCONTRACTSSTORAGE is a free data retrieval call binding the contract method 0x1516e83b.
//
// Solidity: function PERMISSION_CONTRACTS_STORAGE() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCallerSession) PERMISSIONCONTRACTSSTORAGE() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.PERMISSIONCONTRACTSSTORAGE(&_ReservedAliasStorageImpl.CallOpts)
}

// PERMISSIONCONTRACTDEPLOY is a free data retrieval call binding the contract method 0x257a4a04.
//
// Solidity: function PERMISSION_CONTRACT_DEPLOY() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCaller) PERMISSIONCONTRACTDEPLOY(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReservedAliasStorageImpl.contract.Call(opts, out, "PERMISSION_CONTRACT_DEPLOY")
	return *ret0, err
}

// PERMISSIONCONTRACTDEPLOY is a free data retrieval call binding the contract method 0x257a4a04.
//
// Solidity: function PERMISSION_CONTRACT_DEPLOY() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplSession) PERMISSIONCONTRACTDEPLOY() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.PERMISSIONCONTRACTDEPLOY(&_ReservedAliasStorageImpl.CallOpts)
}

// PERMISSIONCONTRACTDEPLOY is a free data retrieval call binding the contract method 0x257a4a04.
//
// Solidity: function PERMISSION_CONTRACT_DEPLOY() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCallerSession) PERMISSIONCONTRACTDEPLOY() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.PERMISSIONCONTRACTDEPLOY(&_ReservedAliasStorageImpl.CallOpts)
}

// PERMISSIONEMISSION is a free data retrieval call binding the contract method 0x4ccd92d1.
//
// Solidity: function PERMISSION_EMISSION() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCaller) PERMISSIONEMISSION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReservedAliasStorageImpl.contract.Call(opts, out, "PERMISSION_EMISSION")
	return *ret0, err
}

// PERMISSIONEMISSION is a free data retrieval call binding the contract method 0x4ccd92d1.
//
// Solidity: function PERMISSION_EMISSION() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplSession) PERMISSIONEMISSION() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.PERMISSIONEMISSION(&_ReservedAliasStorageImpl.CallOpts)
}

// PERMISSIONEMISSION is a free data retrieval call binding the contract method 0x4ccd92d1.
//
// Solidity: function PERMISSION_EMISSION() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCallerSession) PERMISSIONEMISSION() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.PERMISSIONEMISSION(&_ReservedAliasStorageImpl.CallOpts)
}

// PERMISSIONFEESSTORAGE is a free data retrieval call binding the contract method 0xe837eb87.
//
// Solidity: function PERMISSION_FEES_STORAGE() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCaller) PERMISSIONFEESSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReservedAliasStorageImpl.contract.Call(opts, out, "PERMISSION_FEES_STORAGE")
	return *ret0, err
}

// PERMISSIONFEESSTORAGE is a free data retrieval call binding the contract method 0xe837eb87.
//
// Solidity: function PERMISSION_FEES_STORAGE() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplSession) PERMISSIONFEESSTORAGE() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.PERMISSIONFEESSTORAGE(&_ReservedAliasStorageImpl.CallOpts)
}

// PERMISSIONFEESSTORAGE is a free data retrieval call binding the contract method 0xe837eb87.
//
// Solidity: function PERMISSION_FEES_STORAGE() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCallerSession) PERMISSIONFEESSTORAGE() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.PERMISSIONFEESSTORAGE(&_ReservedAliasStorageImpl.CallOpts)
}

// PERMISSIONMULTISIGSTORAGE is a free data retrieval call binding the contract method 0xf1bb7aad.
//
// Solidity: function PERMISSION_MULTISIG_STORAGE() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCaller) PERMISSIONMULTISIGSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReservedAliasStorageImpl.contract.Call(opts, out, "PERMISSION_MULTISIG_STORAGE")
	return *ret0, err
}

// PERMISSIONMULTISIGSTORAGE is a free data retrieval call binding the contract method 0xf1bb7aad.
//
// Solidity: function PERMISSION_MULTISIG_STORAGE() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplSession) PERMISSIONMULTISIGSTORAGE() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.PERMISSIONMULTISIGSTORAGE(&_ReservedAliasStorageImpl.CallOpts)
}

// PERMISSIONMULTISIGSTORAGE is a free data retrieval call binding the contract method 0xf1bb7aad.
//
// Solidity: function PERMISSION_MULTISIG_STORAGE() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCallerSession) PERMISSIONMULTISIGSTORAGE() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.PERMISSIONMULTISIGSTORAGE(&_ReservedAliasStorageImpl.CallOpts)
}

// PERMISSIONROLESSTORAGE is a free data retrieval call binding the contract method 0x12a16032.
//
// Solidity: function PERMISSION_ROLES_STORAGE() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCaller) PERMISSIONROLESSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReservedAliasStorageImpl.contract.Call(opts, out, "PERMISSION_ROLES_STORAGE")
	return *ret0, err
}

// PERMISSIONROLESSTORAGE is a free data retrieval call binding the contract method 0x12a16032.
//
// Solidity: function PERMISSION_ROLES_STORAGE() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplSession) PERMISSIONROLESSTORAGE() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.PERMISSIONROLESSTORAGE(&_ReservedAliasStorageImpl.CallOpts)
}

// PERMISSIONROLESSTORAGE is a free data retrieval call binding the contract method 0x12a16032.
//
// Solidity: function PERMISSION_ROLES_STORAGE() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCallerSession) PERMISSIONROLESSTORAGE() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.PERMISSIONROLESSTORAGE(&_ReservedAliasStorageImpl.CallOpts)
}

// ROLEACCESSVERIFIER is a free data retrieval call binding the contract method 0xc5405f92.
//
// Solidity: function ROLE_ACCESS_VERIFIER() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCaller) ROLEACCESSVERIFIER(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReservedAliasStorageImpl.contract.Call(opts, out, "ROLE_ACCESS_VERIFIER")
	return *ret0, err
}

// ROLEACCESSVERIFIER is a free data retrieval call binding the contract method 0xc5405f92.
//
// Solidity: function ROLE_ACCESS_VERIFIER() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplSession) ROLEACCESSVERIFIER() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.ROLEACCESSVERIFIER(&_ReservedAliasStorageImpl.CallOpts)
}

// ROLEACCESSVERIFIER is a free data retrieval call binding the contract method 0xc5405f92.
//
// Solidity: function ROLE_ACCESS_VERIFIER() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCallerSession) ROLEACCESSVERIFIER() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.ROLEACCESSVERIFIER(&_ReservedAliasStorageImpl.CallOpts)
}

// ROLEBITBONISSUEVERIFIER is a free data retrieval call binding the contract method 0xf0a012a6.
//
// Solidity: function ROLE_BITBON_ISSUE_VERIFIER() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCaller) ROLEBITBONISSUEVERIFIER(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReservedAliasStorageImpl.contract.Call(opts, out, "ROLE_BITBON_ISSUE_VERIFIER")
	return *ret0, err
}

// ROLEBITBONISSUEVERIFIER is a free data retrieval call binding the contract method 0xf0a012a6.
//
// Solidity: function ROLE_BITBON_ISSUE_VERIFIER() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplSession) ROLEBITBONISSUEVERIFIER() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.ROLEBITBONISSUEVERIFIER(&_ReservedAliasStorageImpl.CallOpts)
}

// ROLEBITBONISSUEVERIFIER is a free data retrieval call binding the contract method 0xf0a012a6.
//
// Solidity: function ROLE_BITBON_ISSUE_VERIFIER() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCallerSession) ROLEBITBONISSUEVERIFIER() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.ROLEBITBONISSUEVERIFIER(&_ReservedAliasStorageImpl.CallOpts)
}

// ROLECOMMISSIONVERIFIER is a free data retrieval call binding the contract method 0x8ab3de8c.
//
// Solidity: function ROLE_COMMISSION_VERIFIER() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCaller) ROLECOMMISSIONVERIFIER(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReservedAliasStorageImpl.contract.Call(opts, out, "ROLE_COMMISSION_VERIFIER")
	return *ret0, err
}

// ROLECOMMISSIONVERIFIER is a free data retrieval call binding the contract method 0x8ab3de8c.
//
// Solidity: function ROLE_COMMISSION_VERIFIER() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplSession) ROLECOMMISSIONVERIFIER() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.ROLECOMMISSIONVERIFIER(&_ReservedAliasStorageImpl.CallOpts)
}

// ROLECOMMISSIONVERIFIER is a free data retrieval call binding the contract method 0x8ab3de8c.
//
// Solidity: function ROLE_COMMISSION_VERIFIER() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCallerSession) ROLECOMMISSIONVERIFIER() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.ROLECOMMISSIONVERIFIER(&_ReservedAliasStorageImpl.CallOpts)
}

// ROLEDEPLOYADMIN is a free data retrieval call binding the contract method 0x667bf6e5.
//
// Solidity: function ROLE_DEPLOY_ADMIN() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCaller) ROLEDEPLOYADMIN(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReservedAliasStorageImpl.contract.Call(opts, out, "ROLE_DEPLOY_ADMIN")
	return *ret0, err
}

// ROLEDEPLOYADMIN is a free data retrieval call binding the contract method 0x667bf6e5.
//
// Solidity: function ROLE_DEPLOY_ADMIN() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplSession) ROLEDEPLOYADMIN() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.ROLEDEPLOYADMIN(&_ReservedAliasStorageImpl.CallOpts)
}

// ROLEDEPLOYADMIN is a free data retrieval call binding the contract method 0x667bf6e5.
//
// Solidity: function ROLE_DEPLOY_ADMIN() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCallerSession) ROLEDEPLOYADMIN() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.ROLEDEPLOYADMIN(&_ReservedAliasStorageImpl.CallOpts)
}

// ROLEPERMISSIONADMIN is a free data retrieval call binding the contract method 0x1b239996.
//
// Solidity: function ROLE_PERMISSION_ADMIN() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCaller) ROLEPERMISSIONADMIN(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReservedAliasStorageImpl.contract.Call(opts, out, "ROLE_PERMISSION_ADMIN")
	return *ret0, err
}

// ROLEPERMISSIONADMIN is a free data retrieval call binding the contract method 0x1b239996.
//
// Solidity: function ROLE_PERMISSION_ADMIN() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplSession) ROLEPERMISSIONADMIN() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.ROLEPERMISSIONADMIN(&_ReservedAliasStorageImpl.CallOpts)
}

// ROLEPERMISSIONADMIN is a free data retrieval call binding the contract method 0x1b239996.
//
// Solidity: function ROLE_PERMISSION_ADMIN() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCallerSession) ROLEPERMISSIONADMIN() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.ROLEPERMISSIONADMIN(&_ReservedAliasStorageImpl.CallOpts)
}

// AliasAllowance is a free data retrieval call binding the contract method 0x0ca42ce8.
//
// Solidity: function aliasAllowance(bytes32 ) view returns(address)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCaller) AliasAllowance(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ReservedAliasStorageImpl.contract.Call(opts, out, "aliasAllowance", arg0)
	return *ret0, err
}

// AliasAllowance is a free data retrieval call binding the contract method 0x0ca42ce8.
//
// Solidity: function aliasAllowance(bytes32 ) view returns(address)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplSession) AliasAllowance(arg0 [32]byte) (common.Address, error) {
	return _ReservedAliasStorageImpl.Contract.AliasAllowance(&_ReservedAliasStorageImpl.CallOpts, arg0)
}

// AliasAllowance is a free data retrieval call binding the contract method 0x0ca42ce8.
//
// Solidity: function aliasAllowance(bytes32 ) view returns(address)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCallerSession) AliasAllowance(arg0 [32]byte) (common.Address, error) {
	return _ReservedAliasStorageImpl.Contract.AliasAllowance(&_ReservedAliasStorageImpl.CallOpts, arg0)
}

// AliasIndex is a free data retrieval call binding the contract method 0xc52de8c7.
//
// Solidity: function aliasIndex(bytes32 ) view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCaller) AliasIndex(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReservedAliasStorageImpl.contract.Call(opts, out, "aliasIndex", arg0)
	return *ret0, err
}

// AliasIndex is a free data retrieval call binding the contract method 0xc52de8c7.
//
// Solidity: function aliasIndex(bytes32 ) view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplSession) AliasIndex(arg0 [32]byte) (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.AliasIndex(&_ReservedAliasStorageImpl.CallOpts, arg0)
}

// AliasIndex is a free data retrieval call binding the contract method 0xc52de8c7.
//
// Solidity: function aliasIndex(bytes32 ) view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCallerSession) AliasIndex(arg0 [32]byte) (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.AliasIndex(&_ReservedAliasStorageImpl.CallOpts, arg0)
}

// Aliases is a free data retrieval call binding the contract method 0x6a0e97c3.
//
// Solidity: function aliases(uint256 ) view returns(string)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCaller) Aliases(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ReservedAliasStorageImpl.contract.Call(opts, out, "aliases", arg0)
	return *ret0, err
}

// Aliases is a free data retrieval call binding the contract method 0x6a0e97c3.
//
// Solidity: function aliases(uint256 ) view returns(string)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplSession) Aliases(arg0 *big.Int) (string, error) {
	return _ReservedAliasStorageImpl.Contract.Aliases(&_ReservedAliasStorageImpl.CallOpts, arg0)
}

// Aliases is a free data retrieval call binding the contract method 0x6a0e97c3.
//
// Solidity: function aliases(uint256 ) view returns(string)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCallerSession) Aliases(arg0 *big.Int) (string, error) {
	return _ReservedAliasStorageImpl.Contract.Aliases(&_ReservedAliasStorageImpl.CallOpts, arg0)
}

// AllowedAssetbox is a free data retrieval call binding the contract method 0xadd93303.
//
// Solidity: function allowedAssetbox(string aliasString) view returns(address)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCaller) AllowedAssetbox(opts *bind.CallOpts, aliasString string) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ReservedAliasStorageImpl.contract.Call(opts, out, "allowedAssetbox", aliasString)
	return *ret0, err
}

// AllowedAssetbox is a free data retrieval call binding the contract method 0xadd93303.
//
// Solidity: function allowedAssetbox(string aliasString) view returns(address)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplSession) AllowedAssetbox(aliasString string) (common.Address, error) {
	return _ReservedAliasStorageImpl.Contract.AllowedAssetbox(&_ReservedAliasStorageImpl.CallOpts, aliasString)
}

// AllowedAssetbox is a free data retrieval call binding the contract method 0xadd93303.
//
// Solidity: function allowedAssetbox(string aliasString) view returns(address)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCallerSession) AllowedAssetbox(aliasString string) (common.Address, error) {
	return _ReservedAliasStorageImpl.Contract.AllowedAssetbox(&_ReservedAliasStorageImpl.CallOpts, aliasString)
}

// ContainsAlias is a free data retrieval call binding the contract method 0xf9cd856f.
//
// Solidity: function containsAlias(string aliasString) view returns(bool)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCaller) ContainsAlias(opts *bind.CallOpts, aliasString string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ReservedAliasStorageImpl.contract.Call(opts, out, "containsAlias", aliasString)
	return *ret0, err
}

// ContainsAlias is a free data retrieval call binding the contract method 0xf9cd856f.
//
// Solidity: function containsAlias(string aliasString) view returns(bool)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplSession) ContainsAlias(aliasString string) (bool, error) {
	return _ReservedAliasStorageImpl.Contract.ContainsAlias(&_ReservedAliasStorageImpl.CallOpts, aliasString)
}

// ContainsAlias is a free data retrieval call binding the contract method 0xf9cd856f.
//
// Solidity: function containsAlias(string aliasString) view returns(bool)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCallerSession) ContainsAlias(aliasString string) (bool, error) {
	return _ReservedAliasStorageImpl.Contract.ContainsAlias(&_ReservedAliasStorageImpl.CallOpts, aliasString)
}

// ContractStorage is a free data retrieval call binding the contract method 0x0549b27e.
//
// Solidity: function contractStorage() view returns(address)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCaller) ContractStorage(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ReservedAliasStorageImpl.contract.Call(opts, out, "contractStorage")
	return *ret0, err
}

// ContractStorage is a free data retrieval call binding the contract method 0x0549b27e.
//
// Solidity: function contractStorage() view returns(address)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplSession) ContractStorage() (common.Address, error) {
	return _ReservedAliasStorageImpl.Contract.ContractStorage(&_ReservedAliasStorageImpl.CallOpts)
}

// ContractStorage is a free data retrieval call binding the contract method 0x0549b27e.
//
// Solidity: function contractStorage() view returns(address)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCallerSession) ContractStorage() (common.Address, error) {
	return _ReservedAliasStorageImpl.Contract.ContractStorage(&_ReservedAliasStorageImpl.CallOpts)
}

// GetAllAliases is a free data retrieval call binding the contract method 0xb7da057b.
//
// Solidity: function getAllAliases() view returns(string[])
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCaller) GetAllAliases(opts *bind.CallOpts) ([]string, error) {
	var (
		ret0 = new([]string)
	)
	out := ret0
	err := _ReservedAliasStorageImpl.contract.Call(opts, out, "getAllAliases")
	return *ret0, err
}

// GetAllAliases is a free data retrieval call binding the contract method 0xb7da057b.
//
// Solidity: function getAllAliases() view returns(string[])
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplSession) GetAllAliases() ([]string, error) {
	return _ReservedAliasStorageImpl.Contract.GetAllAliases(&_ReservedAliasStorageImpl.CallOpts)
}

// GetAllAliases is a free data retrieval call binding the contract method 0xb7da057b.
//
// Solidity: function getAllAliases() view returns(string[])
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCallerSession) GetAllAliases() ([]string, error) {
	return _ReservedAliasStorageImpl.Contract.GetAllAliases(&_ReservedAliasStorageImpl.CallOpts)
}

// GetAllSanitizedAliases is a free data retrieval call binding the contract method 0x2b633aae.
//
// Solidity: function getAllSanitizedAliases() view returns(bytes32[])
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCaller) GetAllSanitizedAliases(opts *bind.CallOpts) ([][32]byte, error) {
	var (
		ret0 = new([][32]byte)
	)
	out := ret0
	err := _ReservedAliasStorageImpl.contract.Call(opts, out, "getAllSanitizedAliases")
	return *ret0, err
}

// GetAllSanitizedAliases is a free data retrieval call binding the contract method 0x2b633aae.
//
// Solidity: function getAllSanitizedAliases() view returns(bytes32[])
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplSession) GetAllSanitizedAliases() ([][32]byte, error) {
	return _ReservedAliasStorageImpl.Contract.GetAllSanitizedAliases(&_ReservedAliasStorageImpl.CallOpts)
}

// GetAllSanitizedAliases is a free data retrieval call binding the contract method 0x2b633aae.
//
// Solidity: function getAllSanitizedAliases() view returns(bytes32[])
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCallerSession) GetAllSanitizedAliases() ([][32]byte, error) {
	return _ReservedAliasStorageImpl.Contract.GetAllSanitizedAliases(&_ReservedAliasStorageImpl.CallOpts)
}

// GetAt is a free data retrieval call binding the contract method 0x303eaeed.
//
// Solidity: function getAt(uint256 index) view returns(string)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCaller) GetAt(opts *bind.CallOpts, index *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ReservedAliasStorageImpl.contract.Call(opts, out, "getAt", index)
	return *ret0, err
}

// GetAt is a free data retrieval call binding the contract method 0x303eaeed.
//
// Solidity: function getAt(uint256 index) view returns(string)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplSession) GetAt(index *big.Int) (string, error) {
	return _ReservedAliasStorageImpl.Contract.GetAt(&_ReservedAliasStorageImpl.CallOpts, index)
}

// GetAt is a free data retrieval call binding the contract method 0x303eaeed.
//
// Solidity: function getAt(uint256 index) view returns(string)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCallerSession) GetAt(index *big.Int) (string, error) {
	return _ReservedAliasStorageImpl.Contract.GetAt(&_ReservedAliasStorageImpl.CallOpts, index)
}

// GetThisContractIndex is a free data retrieval call binding the contract method 0x70c6d367.
//
// Solidity: function getThisContractIndex() pure returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCaller) GetThisContractIndex(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReservedAliasStorageImpl.contract.Call(opts, out, "getThisContractIndex")
	return *ret0, err
}

// GetThisContractIndex is a free data retrieval call binding the contract method 0x70c6d367.
//
// Solidity: function getThisContractIndex() pure returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplSession) GetThisContractIndex() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.GetThisContractIndex(&_ReservedAliasStorageImpl.CallOpts)
}

// GetThisContractIndex is a free data retrieval call binding the contract method 0x70c6d367.
//
// Solidity: function getThisContractIndex() pure returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCallerSession) GetThisContractIndex() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.GetThisContractIndex(&_ReservedAliasStorageImpl.CallOpts)
}

// IndexOfAlias is a free data retrieval call binding the contract method 0x19a57d64.
//
// Solidity: function indexOfAlias(string aliasString) view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCaller) IndexOfAlias(opts *bind.CallOpts, aliasString string) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReservedAliasStorageImpl.contract.Call(opts, out, "indexOfAlias", aliasString)
	return *ret0, err
}

// IndexOfAlias is a free data retrieval call binding the contract method 0x19a57d64.
//
// Solidity: function indexOfAlias(string aliasString) view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplSession) IndexOfAlias(aliasString string) (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.IndexOfAlias(&_ReservedAliasStorageImpl.CallOpts, aliasString)
}

// IndexOfAlias is a free data retrieval call binding the contract method 0x19a57d64.
//
// Solidity: function indexOfAlias(string aliasString) view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCallerSession) IndexOfAlias(aliasString string) (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.IndexOfAlias(&_ReservedAliasStorageImpl.CallOpts, aliasString)
}

// IsAllowed is a free data retrieval call binding the contract method 0x79f6909e.
//
// Solidity: function isAllowed(string aliasString, address assetbox) view returns(bool)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCaller) IsAllowed(opts *bind.CallOpts, aliasString string, assetbox common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ReservedAliasStorageImpl.contract.Call(opts, out, "isAllowed", aliasString, assetbox)
	return *ret0, err
}

// IsAllowed is a free data retrieval call binding the contract method 0x79f6909e.
//
// Solidity: function isAllowed(string aliasString, address assetbox) view returns(bool)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplSession) IsAllowed(aliasString string, assetbox common.Address) (bool, error) {
	return _ReservedAliasStorageImpl.Contract.IsAllowed(&_ReservedAliasStorageImpl.CallOpts, aliasString, assetbox)
}

// IsAllowed is a free data retrieval call binding the contract method 0x79f6909e.
//
// Solidity: function isAllowed(string aliasString, address assetbox) view returns(bool)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCallerSession) IsAllowed(aliasString string, assetbox common.Address) (bool, error) {
	return _ReservedAliasStorageImpl.Contract.IsAllowed(&_ReservedAliasStorageImpl.CallOpts, aliasString, assetbox)
}

// Length is a free data retrieval call binding the contract method 0x1f7b6d32.
//
// Solidity: function length() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCaller) Length(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReservedAliasStorageImpl.contract.Call(opts, out, "length")
	return *ret0, err
}

// Length is a free data retrieval call binding the contract method 0x1f7b6d32.
//
// Solidity: function length() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplSession) Length() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.Length(&_ReservedAliasStorageImpl.CallOpts)
}

// Length is a free data retrieval call binding the contract method 0x1f7b6d32.
//
// Solidity: function length() view returns(uint256)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplCallerSession) Length() (*big.Int, error) {
	return _ReservedAliasStorageImpl.Contract.Length(&_ReservedAliasStorageImpl.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0xab78f006.
//
// Solidity: function approve(string aliasString, address assetbox) returns(bool)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplTransactor) Approve(opts *bind.TransactOpts, aliasString string, assetbox common.Address) (*types.Transaction, error) {
	return _ReservedAliasStorageImpl.contract.Transact(opts, "approve", aliasString, assetbox)
}

// Approve is a paid mutator transaction binding the contract method 0xab78f006.
//
// Solidity: function approve(string aliasString, address assetbox) returns(bool)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplSession) Approve(aliasString string, assetbox common.Address) (*types.Transaction, error) {
	return _ReservedAliasStorageImpl.Contract.Approve(&_ReservedAliasStorageImpl.TransactOpts, aliasString, assetbox)
}

// Approve is a paid mutator transaction binding the contract method 0xab78f006.
//
// Solidity: function approve(string aliasString, address assetbox) returns(bool)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplTransactorSession) Approve(aliasString string, assetbox common.Address) (*types.Transaction, error) {
	return _ReservedAliasStorageImpl.Contract.Approve(&_ReservedAliasStorageImpl.TransactOpts, aliasString, assetbox)
}

// EditReservedAlias is a paid mutator transaction binding the contract method 0x845d48a4.
//
// Solidity: function editReservedAlias(string oldAlias, string newAlias) returns(bool)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplTransactor) EditReservedAlias(opts *bind.TransactOpts, oldAlias string, newAlias string) (*types.Transaction, error) {
	return _ReservedAliasStorageImpl.contract.Transact(opts, "editReservedAlias", oldAlias, newAlias)
}

// EditReservedAlias is a paid mutator transaction binding the contract method 0x845d48a4.
//
// Solidity: function editReservedAlias(string oldAlias, string newAlias) returns(bool)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplSession) EditReservedAlias(oldAlias string, newAlias string) (*types.Transaction, error) {
	return _ReservedAliasStorageImpl.Contract.EditReservedAlias(&_ReservedAliasStorageImpl.TransactOpts, oldAlias, newAlias)
}

// EditReservedAlias is a paid mutator transaction binding the contract method 0x845d48a4.
//
// Solidity: function editReservedAlias(string oldAlias, string newAlias) returns(bool)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplTransactorSession) EditReservedAlias(oldAlias string, newAlias string) (*types.Transaction, error) {
	return _ReservedAliasStorageImpl.Contract.EditReservedAlias(&_ReservedAliasStorageImpl.TransactOpts, oldAlias, newAlias)
}

// RemoveAliasReservation is a paid mutator transaction binding the contract method 0xef234326.
//
// Solidity: function removeAliasReservation(string aliasString) returns(bool)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplTransactor) RemoveAliasReservation(opts *bind.TransactOpts, aliasString string) (*types.Transaction, error) {
	return _ReservedAliasStorageImpl.contract.Transact(opts, "removeAliasReservation", aliasString)
}

// RemoveAliasReservation is a paid mutator transaction binding the contract method 0xef234326.
//
// Solidity: function removeAliasReservation(string aliasString) returns(bool)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplSession) RemoveAliasReservation(aliasString string) (*types.Transaction, error) {
	return _ReservedAliasStorageImpl.Contract.RemoveAliasReservation(&_ReservedAliasStorageImpl.TransactOpts, aliasString)
}

// RemoveAliasReservation is a paid mutator transaction binding the contract method 0xef234326.
//
// Solidity: function removeAliasReservation(string aliasString) returns(bool)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplTransactorSession) RemoveAliasReservation(aliasString string) (*types.Transaction, error) {
	return _ReservedAliasStorageImpl.Contract.RemoveAliasReservation(&_ReservedAliasStorageImpl.TransactOpts, aliasString)
}

// RemoveAliasReservationAt is a paid mutator transaction binding the contract method 0x01e7b297.
//
// Solidity: function removeAliasReservationAt(uint256 index) returns(bool)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplTransactor) RemoveAliasReservationAt(opts *bind.TransactOpts, index *big.Int) (*types.Transaction, error) {
	return _ReservedAliasStorageImpl.contract.Transact(opts, "removeAliasReservationAt", index)
}

// RemoveAliasReservationAt is a paid mutator transaction binding the contract method 0x01e7b297.
//
// Solidity: function removeAliasReservationAt(uint256 index) returns(bool)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplSession) RemoveAliasReservationAt(index *big.Int) (*types.Transaction, error) {
	return _ReservedAliasStorageImpl.Contract.RemoveAliasReservationAt(&_ReservedAliasStorageImpl.TransactOpts, index)
}

// RemoveAliasReservationAt is a paid mutator transaction binding the contract method 0x01e7b297.
//
// Solidity: function removeAliasReservationAt(uint256 index) returns(bool)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplTransactorSession) RemoveAliasReservationAt(index *big.Int) (*types.Transaction, error) {
	return _ReservedAliasStorageImpl.Contract.RemoveAliasReservationAt(&_ReservedAliasStorageImpl.TransactOpts, index)
}

// ReserveAlias is a paid mutator transaction binding the contract method 0x67f7eaba.
//
// Solidity: function reserveAlias(string aliasString) returns(bool)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplTransactor) ReserveAlias(opts *bind.TransactOpts, aliasString string) (*types.Transaction, error) {
	return _ReservedAliasStorageImpl.contract.Transact(opts, "reserveAlias", aliasString)
}

// ReserveAlias is a paid mutator transaction binding the contract method 0x67f7eaba.
//
// Solidity: function reserveAlias(string aliasString) returns(bool)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplSession) ReserveAlias(aliasString string) (*types.Transaction, error) {
	return _ReservedAliasStorageImpl.Contract.ReserveAlias(&_ReservedAliasStorageImpl.TransactOpts, aliasString)
}

// ReserveAlias is a paid mutator transaction binding the contract method 0x67f7eaba.
//
// Solidity: function reserveAlias(string aliasString) returns(bool)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplTransactorSession) ReserveAlias(aliasString string) (*types.Transaction, error) {
	return _ReservedAliasStorageImpl.Contract.ReserveAlias(&_ReservedAliasStorageImpl.TransactOpts, aliasString)
}

// ReservedAliasStorageImplAliasReservedIterator is returned from FilterAliasReserved and is used to iterate over the raw logs and unpacked data for AliasReserved events raised by the ReservedAliasStorageImpl contract.
type ReservedAliasStorageImplAliasReservedIterator struct {
	Event *ReservedAliasStorageImplAliasReserved // Event containing the contract specifics and raw log

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
func (it *ReservedAliasStorageImplAliasReservedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReservedAliasStorageImplAliasReserved)
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
		it.Event = new(ReservedAliasStorageImplAliasReserved)
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
func (it *ReservedAliasStorageImplAliasReservedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReservedAliasStorageImplAliasReservedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReservedAliasStorageImplAliasReserved represents a AliasReserved event raised by the ReservedAliasStorageImpl contract.
type ReservedAliasStorageImplAliasReserved struct {
	AliasString string
	AliasHash   [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAliasReserved is a free log retrieval operation binding the contract event 0x91e2b3191ca4e97bd18434ab643c7f95cf2481830ea74460bbc8a06361ceaedb.
//
// Solidity: event AliasReserved(string aliasString, bytes32 aliasHash)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplFilterer) FilterAliasReserved(opts *bind.FilterOpts) (*ReservedAliasStorageImplAliasReservedIterator, error) {

	logs, sub, err := _ReservedAliasStorageImpl.contract.FilterLogs(opts, "AliasReserved")
	if err != nil {
		return nil, err
	}
	return &ReservedAliasStorageImplAliasReservedIterator{contract: _ReservedAliasStorageImpl.contract, event: "AliasReserved", logs: logs, sub: sub}, nil
}

// WatchAliasReserved is a free log subscription operation binding the contract event 0x91e2b3191ca4e97bd18434ab643c7f95cf2481830ea74460bbc8a06361ceaedb.
//
// Solidity: event AliasReserved(string aliasString, bytes32 aliasHash)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplFilterer) WatchAliasReserved(opts *bind.WatchOpts, sink chan<- *ReservedAliasStorageImplAliasReserved) (event.Subscription, error) {

	logs, sub, err := _ReservedAliasStorageImpl.contract.WatchLogs(opts, "AliasReserved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReservedAliasStorageImplAliasReserved)
				if err := _ReservedAliasStorageImpl.contract.UnpackLog(event, "AliasReserved", log); err != nil {
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

// ParseAliasReserved is a log parse operation binding the contract event 0x91e2b3191ca4e97bd18434ab643c7f95cf2481830ea74460bbc8a06361ceaedb.
//
// Solidity: event AliasReserved(string aliasString, bytes32 aliasHash)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplFilterer) ParseAliasReserved(log types.Log) (*ReservedAliasStorageImplAliasReserved, error) {
	event := new(ReservedAliasStorageImplAliasReserved)
	if err := _ReservedAliasStorageImpl.contract.UnpackLog(event, "AliasReserved", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ReservedAliasStorageImplReservedAliasAllowedIterator is returned from FilterReservedAliasAllowed and is used to iterate over the raw logs and unpacked data for ReservedAliasAllowed events raised by the ReservedAliasStorageImpl contract.
type ReservedAliasStorageImplReservedAliasAllowedIterator struct {
	Event *ReservedAliasStorageImplReservedAliasAllowed // Event containing the contract specifics and raw log

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
func (it *ReservedAliasStorageImplReservedAliasAllowedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReservedAliasStorageImplReservedAliasAllowed)
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
		it.Event = new(ReservedAliasStorageImplReservedAliasAllowed)
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
func (it *ReservedAliasStorageImplReservedAliasAllowedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReservedAliasStorageImplReservedAliasAllowedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReservedAliasStorageImplReservedAliasAllowed represents a ReservedAliasAllowed event raised by the ReservedAliasStorageImpl contract.
type ReservedAliasStorageImplReservedAliasAllowed struct {
	AliasString string
	AliasHash   [32]byte
	Assetbox    common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterReservedAliasAllowed is a free log retrieval operation binding the contract event 0xadc0ca024f8da9b1e789cd1d49f6f1fda3dfef4c346ec1eab0eededbd3cc32dc.
//
// Solidity: event ReservedAliasAllowed(string aliasString, bytes32 aliasHash, address assetbox)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplFilterer) FilterReservedAliasAllowed(opts *bind.FilterOpts) (*ReservedAliasStorageImplReservedAliasAllowedIterator, error) {

	logs, sub, err := _ReservedAliasStorageImpl.contract.FilterLogs(opts, "ReservedAliasAllowed")
	if err != nil {
		return nil, err
	}
	return &ReservedAliasStorageImplReservedAliasAllowedIterator{contract: _ReservedAliasStorageImpl.contract, event: "ReservedAliasAllowed", logs: logs, sub: sub}, nil
}

// WatchReservedAliasAllowed is a free log subscription operation binding the contract event 0xadc0ca024f8da9b1e789cd1d49f6f1fda3dfef4c346ec1eab0eededbd3cc32dc.
//
// Solidity: event ReservedAliasAllowed(string aliasString, bytes32 aliasHash, address assetbox)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplFilterer) WatchReservedAliasAllowed(opts *bind.WatchOpts, sink chan<- *ReservedAliasStorageImplReservedAliasAllowed) (event.Subscription, error) {

	logs, sub, err := _ReservedAliasStorageImpl.contract.WatchLogs(opts, "ReservedAliasAllowed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReservedAliasStorageImplReservedAliasAllowed)
				if err := _ReservedAliasStorageImpl.contract.UnpackLog(event, "ReservedAliasAllowed", log); err != nil {
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

// ParseReservedAliasAllowed is a log parse operation binding the contract event 0xadc0ca024f8da9b1e789cd1d49f6f1fda3dfef4c346ec1eab0eededbd3cc32dc.
//
// Solidity: event ReservedAliasAllowed(string aliasString, bytes32 aliasHash, address assetbox)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplFilterer) ParseReservedAliasAllowed(log types.Log) (*ReservedAliasStorageImplReservedAliasAllowed, error) {
	event := new(ReservedAliasStorageImplReservedAliasAllowed)
	if err := _ReservedAliasStorageImpl.contract.UnpackLog(event, "ReservedAliasAllowed", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ReservedAliasStorageImplReservedAliasEditedIterator is returned from FilterReservedAliasEdited and is used to iterate over the raw logs and unpacked data for ReservedAliasEdited events raised by the ReservedAliasStorageImpl contract.
type ReservedAliasStorageImplReservedAliasEditedIterator struct {
	Event *ReservedAliasStorageImplReservedAliasEdited // Event containing the contract specifics and raw log

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
func (it *ReservedAliasStorageImplReservedAliasEditedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReservedAliasStorageImplReservedAliasEdited)
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
		it.Event = new(ReservedAliasStorageImplReservedAliasEdited)
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
func (it *ReservedAliasStorageImplReservedAliasEditedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReservedAliasStorageImplReservedAliasEditedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReservedAliasStorageImplReservedAliasEdited represents a ReservedAliasEdited event raised by the ReservedAliasStorageImpl contract.
type ReservedAliasStorageImplReservedAliasEdited struct {
	OldAlias     string
	OldAliasHash [32]byte
	NewAlias     string
	NewAliasHash [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterReservedAliasEdited is a free log retrieval operation binding the contract event 0x45da68949b8f5e6f7e9fed38f940ad8e9ea3b951fa1e2eb8e299bb41c773e377.
//
// Solidity: event ReservedAliasEdited(string oldAlias, bytes32 oldAliasHash, string newAlias, bytes32 newAliasHash)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplFilterer) FilterReservedAliasEdited(opts *bind.FilterOpts) (*ReservedAliasStorageImplReservedAliasEditedIterator, error) {

	logs, sub, err := _ReservedAliasStorageImpl.contract.FilterLogs(opts, "ReservedAliasEdited")
	if err != nil {
		return nil, err
	}
	return &ReservedAliasStorageImplReservedAliasEditedIterator{contract: _ReservedAliasStorageImpl.contract, event: "ReservedAliasEdited", logs: logs, sub: sub}, nil
}

// WatchReservedAliasEdited is a free log subscription operation binding the contract event 0x45da68949b8f5e6f7e9fed38f940ad8e9ea3b951fa1e2eb8e299bb41c773e377.
//
// Solidity: event ReservedAliasEdited(string oldAlias, bytes32 oldAliasHash, string newAlias, bytes32 newAliasHash)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplFilterer) WatchReservedAliasEdited(opts *bind.WatchOpts, sink chan<- *ReservedAliasStorageImplReservedAliasEdited) (event.Subscription, error) {

	logs, sub, err := _ReservedAliasStorageImpl.contract.WatchLogs(opts, "ReservedAliasEdited")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReservedAliasStorageImplReservedAliasEdited)
				if err := _ReservedAliasStorageImpl.contract.UnpackLog(event, "ReservedAliasEdited", log); err != nil {
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

// ParseReservedAliasEdited is a log parse operation binding the contract event 0x45da68949b8f5e6f7e9fed38f940ad8e9ea3b951fa1e2eb8e299bb41c773e377.
//
// Solidity: event ReservedAliasEdited(string oldAlias, bytes32 oldAliasHash, string newAlias, bytes32 newAliasHash)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplFilterer) ParseReservedAliasEdited(log types.Log) (*ReservedAliasStorageImplReservedAliasEdited, error) {
	event := new(ReservedAliasStorageImplReservedAliasEdited)
	if err := _ReservedAliasStorageImpl.contract.UnpackLog(event, "ReservedAliasEdited", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ReservedAliasStorageImplReservedAliasRemovedIterator is returned from FilterReservedAliasRemoved and is used to iterate over the raw logs and unpacked data for ReservedAliasRemoved events raised by the ReservedAliasStorageImpl contract.
type ReservedAliasStorageImplReservedAliasRemovedIterator struct {
	Event *ReservedAliasStorageImplReservedAliasRemoved // Event containing the contract specifics and raw log

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
func (it *ReservedAliasStorageImplReservedAliasRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReservedAliasStorageImplReservedAliasRemoved)
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
		it.Event = new(ReservedAliasStorageImplReservedAliasRemoved)
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
func (it *ReservedAliasStorageImplReservedAliasRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReservedAliasStorageImplReservedAliasRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReservedAliasStorageImplReservedAliasRemoved represents a ReservedAliasRemoved event raised by the ReservedAliasStorageImpl contract.
type ReservedAliasStorageImplReservedAliasRemoved struct {
	AliasString string
	AliasHash   [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterReservedAliasRemoved is a free log retrieval operation binding the contract event 0x22dadbc77b1d3441a8895010c03f267bbad81fe092fbecc2e1b816d7e877ca92.
//
// Solidity: event ReservedAliasRemoved(string aliasString, bytes32 aliasHash)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplFilterer) FilterReservedAliasRemoved(opts *bind.FilterOpts) (*ReservedAliasStorageImplReservedAliasRemovedIterator, error) {

	logs, sub, err := _ReservedAliasStorageImpl.contract.FilterLogs(opts, "ReservedAliasRemoved")
	if err != nil {
		return nil, err
	}
	return &ReservedAliasStorageImplReservedAliasRemovedIterator{contract: _ReservedAliasStorageImpl.contract, event: "ReservedAliasRemoved", logs: logs, sub: sub}, nil
}

// WatchReservedAliasRemoved is a free log subscription operation binding the contract event 0x22dadbc77b1d3441a8895010c03f267bbad81fe092fbecc2e1b816d7e877ca92.
//
// Solidity: event ReservedAliasRemoved(string aliasString, bytes32 aliasHash)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplFilterer) WatchReservedAliasRemoved(opts *bind.WatchOpts, sink chan<- *ReservedAliasStorageImplReservedAliasRemoved) (event.Subscription, error) {

	logs, sub, err := _ReservedAliasStorageImpl.contract.WatchLogs(opts, "ReservedAliasRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReservedAliasStorageImplReservedAliasRemoved)
				if err := _ReservedAliasStorageImpl.contract.UnpackLog(event, "ReservedAliasRemoved", log); err != nil {
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

// ParseReservedAliasRemoved is a log parse operation binding the contract event 0x22dadbc77b1d3441a8895010c03f267bbad81fe092fbecc2e1b816d7e877ca92.
//
// Solidity: event ReservedAliasRemoved(string aliasString, bytes32 aliasHash)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplFilterer) ParseReservedAliasRemoved(log types.Log) (*ReservedAliasStorageImplReservedAliasRemoved, error) {
	event := new(ReservedAliasStorageImplReservedAliasRemoved)
	if err := _ReservedAliasStorageImpl.contract.UnpackLog(event, "ReservedAliasRemoved", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ReservedAliasStorageImplReservedAliasTakenIterator is returned from FilterReservedAliasTaken and is used to iterate over the raw logs and unpacked data for ReservedAliasTaken events raised by the ReservedAliasStorageImpl contract.
type ReservedAliasStorageImplReservedAliasTakenIterator struct {
	Event *ReservedAliasStorageImplReservedAliasTaken // Event containing the contract specifics and raw log

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
func (it *ReservedAliasStorageImplReservedAliasTakenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReservedAliasStorageImplReservedAliasTaken)
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
		it.Event = new(ReservedAliasStorageImplReservedAliasTaken)
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
func (it *ReservedAliasStorageImplReservedAliasTakenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReservedAliasStorageImplReservedAliasTakenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReservedAliasStorageImplReservedAliasTaken represents a ReservedAliasTaken event raised by the ReservedAliasStorageImpl contract.
type ReservedAliasStorageImplReservedAliasTaken struct {
	AliasString string
	AliasHash   [32]byte
	Assetbox    common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterReservedAliasTaken is a free log retrieval operation binding the contract event 0x4919e61240c68702fe637c8d8bb1eff8c8519c0f301b85f83cf31e009fb1e171.
//
// Solidity: event ReservedAliasTaken(string aliasString, bytes32 aliasHash, address assetbox)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplFilterer) FilterReservedAliasTaken(opts *bind.FilterOpts) (*ReservedAliasStorageImplReservedAliasTakenIterator, error) {

	logs, sub, err := _ReservedAliasStorageImpl.contract.FilterLogs(opts, "ReservedAliasTaken")
	if err != nil {
		return nil, err
	}
	return &ReservedAliasStorageImplReservedAliasTakenIterator{contract: _ReservedAliasStorageImpl.contract, event: "ReservedAliasTaken", logs: logs, sub: sub}, nil
}

// WatchReservedAliasTaken is a free log subscription operation binding the contract event 0x4919e61240c68702fe637c8d8bb1eff8c8519c0f301b85f83cf31e009fb1e171.
//
// Solidity: event ReservedAliasTaken(string aliasString, bytes32 aliasHash, address assetbox)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplFilterer) WatchReservedAliasTaken(opts *bind.WatchOpts, sink chan<- *ReservedAliasStorageImplReservedAliasTaken) (event.Subscription, error) {

	logs, sub, err := _ReservedAliasStorageImpl.contract.WatchLogs(opts, "ReservedAliasTaken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReservedAliasStorageImplReservedAliasTaken)
				if err := _ReservedAliasStorageImpl.contract.UnpackLog(event, "ReservedAliasTaken", log); err != nil {
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

// ParseReservedAliasTaken is a log parse operation binding the contract event 0x4919e61240c68702fe637c8d8bb1eff8c8519c0f301b85f83cf31e009fb1e171.
//
// Solidity: event ReservedAliasTaken(string aliasString, bytes32 aliasHash, address assetbox)
func (_ReservedAliasStorageImpl *ReservedAliasStorageImplFilterer) ParseReservedAliasTaken(log types.Log) (*ReservedAliasStorageImplReservedAliasTaken, error) {
	event := new(ReservedAliasStorageImplReservedAliasTaken)
	if err := _ReservedAliasStorageImpl.contract.UnpackLog(event, "ReservedAliasTaken", log); err != nil {
		return nil, err
	}
	return event, nil
}
