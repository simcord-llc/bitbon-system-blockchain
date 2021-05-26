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

// BitbonStorageImplABI is the input ABI used to generate the binding from.
const BitbonStorageImplABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"assetbox\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"contractStorageAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"assetbox\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"capitalise\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mining\",\"type\":\"uint256\"}],\"name\":\"FeeCharged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"assetbox\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Lock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"assetbox\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"UnLock\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ACCESS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ACCESS_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ADMINABLE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ADMIN_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ASSETBOX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ASSETBOX_INFO\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ASSETBOX_INFO_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ASSETBOX_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_BALANCE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_BITBON\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_BITBON_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_BITBON_SUPPORT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_BONBALANCE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_BONTRANSFER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_DEX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_DISTRIBUTION_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_EXCHANGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_EXCHANGE_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_FEE_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_GENERATOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MINING_AGENT_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MSBON\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MSBON_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_ADD_ADDRESS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_ADD_ADMIN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_ADD_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_ADD_ROLE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_DROP_PERMISSION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_EDIT_ADMIN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_EDIT_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_EDIT_ROLE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REMOVE_ADMIN_BY_ADDRESS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REMOVE_ADMIN_BY_INDEX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REMOVE_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REMOVE_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REMOVE_ROLE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REPLENISH_CAPITALIZATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REPLENISH_CERTIFICATE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REPLENISH_EMISSION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_SET_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_SET_PERMISSION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_OTC\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_OTC_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_REPLENISH_CAPITALIZATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_RESERVED_ALIAS_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ROLE_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_SAFE_TRANSFER_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_TRANSFER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_TRANSFER_FROM_CAPITALIZATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_ACCESS_RESTORATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_ADMIN_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_CONTRACTS_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_CONTRACT_DEPLOY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_EMISSION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_FEES_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_MULTISIG_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_ROLES_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_ACCESS_VERIFIER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_BITBON_ISSUE_VERIFIER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_COMMISSION_VERIFIER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_DEPLOY_ADMIN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_PERMISSION_ADMIN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"contractStorage\",\"outputs\":[{\"internalType\":\"contractContractStorage\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getThisContractIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"assetbox\",\"type\":\"address\"}],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"assetbox\",\"type\":\"address\"}],\"name\":\"getLockBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"assetbox\",\"type\":\"address\"}],\"name\":\"getFullBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTotalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"assetbox\",\"type\":\"address\"}],\"name\":\"setCapitaliseFund\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"assetbox\",\"type\":\"address\"}],\"name\":\"setSimcordAssetbox\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"assetbox\",\"type\":\"address\"}],\"name\":\"setMiningFund\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"assetbox\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"lock\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"assetbox\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"unlock\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"unlockAndTransfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountFee\",\"type\":\"uint256\"}],\"name\":\"chargeFee\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"assetbox\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"setAssetboxBalance\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"setMigrated\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// BitbonStorageImpl is an auto generated Go binding around an Ethereum contract.
type BitbonStorageImpl struct {
	BitbonStorageImplCaller     // Read-only binding to the contract
	BitbonStorageImplTransactor // Write-only binding to the contract
	BitbonStorageImplFilterer   // Log filterer for contract events
}

// BitbonStorageImplCaller is an auto generated read-only Go binding around an Ethereum contract.
type BitbonStorageImplCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BitbonStorageImplTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BitbonStorageImplTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BitbonStorageImplFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BitbonStorageImplFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BitbonStorageImplSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BitbonStorageImplSession struct {
	Contract     *BitbonStorageImpl // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// BitbonStorageImplCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BitbonStorageImplCallerSession struct {
	Contract *BitbonStorageImplCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// BitbonStorageImplTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BitbonStorageImplTransactorSession struct {
	Contract     *BitbonStorageImplTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// BitbonStorageImplRaw is an auto generated low-level Go binding around an Ethereum contract.
type BitbonStorageImplRaw struct {
	Contract *BitbonStorageImpl // Generic contract binding to access the raw methods on
}

// BitbonStorageImplCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BitbonStorageImplCallerRaw struct {
	Contract *BitbonStorageImplCaller // Generic read-only contract binding to access the raw methods on
}

// BitbonStorageImplTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BitbonStorageImplTransactorRaw struct {
	Contract *BitbonStorageImplTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBitbonStorageImpl creates a new instance of BitbonStorageImpl, bound to a specific deployed contract.
func NewBitbonStorageImpl(address common.Address, backend bind.ContractBackend) (*BitbonStorageImpl, error) {
	contract, err := bindBitbonStorageImpl(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BitbonStorageImpl{BitbonStorageImplCaller: BitbonStorageImplCaller{contract: contract}, BitbonStorageImplTransactor: BitbonStorageImplTransactor{contract: contract}, BitbonStorageImplFilterer: BitbonStorageImplFilterer{contract: contract}}, nil
}

// NewBitbonStorageImplCaller creates a new read-only instance of BitbonStorageImpl, bound to a specific deployed contract.
func NewBitbonStorageImplCaller(address common.Address, caller bind.ContractCaller) (*BitbonStorageImplCaller, error) {
	contract, err := bindBitbonStorageImpl(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BitbonStorageImplCaller{contract: contract}, nil
}

// NewBitbonStorageImplTransactor creates a new write-only instance of BitbonStorageImpl, bound to a specific deployed contract.
func NewBitbonStorageImplTransactor(address common.Address, transactor bind.ContractTransactor) (*BitbonStorageImplTransactor, error) {
	contract, err := bindBitbonStorageImpl(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BitbonStorageImplTransactor{contract: contract}, nil
}

// NewBitbonStorageImplFilterer creates a new log filterer instance of BitbonStorageImpl, bound to a specific deployed contract.
func NewBitbonStorageImplFilterer(address common.Address, filterer bind.ContractFilterer) (*BitbonStorageImplFilterer, error) {
	contract, err := bindBitbonStorageImpl(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BitbonStorageImplFilterer{contract: contract}, nil
}

// bindBitbonStorageImpl binds a generic wrapper to an already deployed contract.
func bindBitbonStorageImpl(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BitbonStorageImplABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BitbonStorageImpl *BitbonStorageImplRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BitbonStorageImpl.Contract.BitbonStorageImplCaller.contract.Call(opts, result, method, params...)
}

// Return Abi for BitbonStorageImplCaller
func (_BitbonStorageImpl *BitbonStorageImplCaller) ABI() abi.ABI {
	return _BitbonStorageImpl.contract.Abi
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BitbonStorageImpl *BitbonStorageImplRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BitbonStorageImpl.Contract.BitbonStorageImplTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BitbonStorageImpl *BitbonStorageImplRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BitbonStorageImpl.Contract.BitbonStorageImplTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BitbonStorageImpl *BitbonStorageImplCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BitbonStorageImpl.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BitbonStorageImpl *BitbonStorageImplTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BitbonStorageImpl.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BitbonStorageImpl *BitbonStorageImplTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BitbonStorageImpl.Contract.contract.Transact(opts, method, params...)
}

// CONTRACTACCESS is a free data retrieval call binding the contract method 0x0ac62b0c.
//
// Solidity: function CONTRACT_ACCESS() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCaller) CONTRACTACCESS(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonStorageImpl.contract.Call(opts, out, "CONTRACT_ACCESS")
	return *ret0, err
}

// CONTRACTACCESS is a free data retrieval call binding the contract method 0x0ac62b0c.
//
// Solidity: function CONTRACT_ACCESS() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplSession) CONTRACTACCESS() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.CONTRACTACCESS(&_BitbonStorageImpl.CallOpts)
}

// CONTRACTACCESS is a free data retrieval call binding the contract method 0x0ac62b0c.
//
// Solidity: function CONTRACT_ACCESS() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCallerSession) CONTRACTACCESS() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.CONTRACTACCESS(&_BitbonStorageImpl.CallOpts)
}

// CONTRACTACCESSSTORAGE is a free data retrieval call binding the contract method 0xceac9bd5.
//
// Solidity: function CONTRACT_ACCESS_STORAGE() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCaller) CONTRACTACCESSSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonStorageImpl.contract.Call(opts, out, "CONTRACT_ACCESS_STORAGE")
	return *ret0, err
}

// CONTRACTACCESSSTORAGE is a free data retrieval call binding the contract method 0xceac9bd5.
//
// Solidity: function CONTRACT_ACCESS_STORAGE() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplSession) CONTRACTACCESSSTORAGE() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.CONTRACTACCESSSTORAGE(&_BitbonStorageImpl.CallOpts)
}

// CONTRACTACCESSSTORAGE is a free data retrieval call binding the contract method 0xceac9bd5.
//
// Solidity: function CONTRACT_ACCESS_STORAGE() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCallerSession) CONTRACTACCESSSTORAGE() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.CONTRACTACCESSSTORAGE(&_BitbonStorageImpl.CallOpts)
}

// CONTRACTADMINABLE is a free data retrieval call binding the contract method 0xbbe3a78f.
//
// Solidity: function CONTRACT_ADMINABLE() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCaller) CONTRACTADMINABLE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonStorageImpl.contract.Call(opts, out, "CONTRACT_ADMINABLE")
	return *ret0, err
}

// CONTRACTADMINABLE is a free data retrieval call binding the contract method 0xbbe3a78f.
//
// Solidity: function CONTRACT_ADMINABLE() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplSession) CONTRACTADMINABLE() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.CONTRACTADMINABLE(&_BitbonStorageImpl.CallOpts)
}

// CONTRACTADMINABLE is a free data retrieval call binding the contract method 0xbbe3a78f.
//
// Solidity: function CONTRACT_ADMINABLE() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCallerSession) CONTRACTADMINABLE() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.CONTRACTADMINABLE(&_BitbonStorageImpl.CallOpts)
}

// CONTRACTADMINSTORAGE is a free data retrieval call binding the contract method 0xc9aebae9.
//
// Solidity: function CONTRACT_ADMIN_STORAGE() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCaller) CONTRACTADMINSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonStorageImpl.contract.Call(opts, out, "CONTRACT_ADMIN_STORAGE")
	return *ret0, err
}

// CONTRACTADMINSTORAGE is a free data retrieval call binding the contract method 0xc9aebae9.
//
// Solidity: function CONTRACT_ADMIN_STORAGE() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplSession) CONTRACTADMINSTORAGE() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.CONTRACTADMINSTORAGE(&_BitbonStorageImpl.CallOpts)
}

// CONTRACTADMINSTORAGE is a free data retrieval call binding the contract method 0xc9aebae9.
//
// Solidity: function CONTRACT_ADMIN_STORAGE() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCallerSession) CONTRACTADMINSTORAGE() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.CONTRACTADMINSTORAGE(&_BitbonStorageImpl.CallOpts)
}

// CONTRACTASSETBOX is a free data retrieval call binding the contract method 0x161aa17c.
//
// Solidity: function CONTRACT_ASSETBOX() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCaller) CONTRACTASSETBOX(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonStorageImpl.contract.Call(opts, out, "CONTRACT_ASSETBOX")
	return *ret0, err
}

// CONTRACTASSETBOX is a free data retrieval call binding the contract method 0x161aa17c.
//
// Solidity: function CONTRACT_ASSETBOX() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplSession) CONTRACTASSETBOX() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.CONTRACTASSETBOX(&_BitbonStorageImpl.CallOpts)
}

// CONTRACTASSETBOX is a free data retrieval call binding the contract method 0x161aa17c.
//
// Solidity: function CONTRACT_ASSETBOX() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCallerSession) CONTRACTASSETBOX() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.CONTRACTASSETBOX(&_BitbonStorageImpl.CallOpts)
}

// CONTRACTASSETBOXINFO is a free data retrieval call binding the contract method 0xb572fd3d.
//
// Solidity: function CONTRACT_ASSETBOX_INFO() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCaller) CONTRACTASSETBOXINFO(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonStorageImpl.contract.Call(opts, out, "CONTRACT_ASSETBOX_INFO")
	return *ret0, err
}

// CONTRACTASSETBOXINFO is a free data retrieval call binding the contract method 0xb572fd3d.
//
// Solidity: function CONTRACT_ASSETBOX_INFO() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplSession) CONTRACTASSETBOXINFO() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.CONTRACTASSETBOXINFO(&_BitbonStorageImpl.CallOpts)
}

// CONTRACTASSETBOXINFO is a free data retrieval call binding the contract method 0xb572fd3d.
//
// Solidity: function CONTRACT_ASSETBOX_INFO() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCallerSession) CONTRACTASSETBOXINFO() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.CONTRACTASSETBOXINFO(&_BitbonStorageImpl.CallOpts)
}

// CONTRACTASSETBOXINFOSTORAGE is a free data retrieval call binding the contract method 0xf32a115d.
//
// Solidity: function CONTRACT_ASSETBOX_INFO_STORAGE() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCaller) CONTRACTASSETBOXINFOSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonStorageImpl.contract.Call(opts, out, "CONTRACT_ASSETBOX_INFO_STORAGE")
	return *ret0, err
}

// CONTRACTASSETBOXINFOSTORAGE is a free data retrieval call binding the contract method 0xf32a115d.
//
// Solidity: function CONTRACT_ASSETBOX_INFO_STORAGE() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplSession) CONTRACTASSETBOXINFOSTORAGE() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.CONTRACTASSETBOXINFOSTORAGE(&_BitbonStorageImpl.CallOpts)
}

// CONTRACTASSETBOXINFOSTORAGE is a free data retrieval call binding the contract method 0xf32a115d.
//
// Solidity: function CONTRACT_ASSETBOX_INFO_STORAGE() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCallerSession) CONTRACTASSETBOXINFOSTORAGE() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.CONTRACTASSETBOXINFOSTORAGE(&_BitbonStorageImpl.CallOpts)
}

// CONTRACTASSETBOXSTORAGE is a free data retrieval call binding the contract method 0xfa5bdafc.
//
// Solidity: function CONTRACT_ASSETBOX_STORAGE() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCaller) CONTRACTASSETBOXSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonStorageImpl.contract.Call(opts, out, "CONTRACT_ASSETBOX_STORAGE")
	return *ret0, err
}

// CONTRACTASSETBOXSTORAGE is a free data retrieval call binding the contract method 0xfa5bdafc.
//
// Solidity: function CONTRACT_ASSETBOX_STORAGE() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplSession) CONTRACTASSETBOXSTORAGE() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.CONTRACTASSETBOXSTORAGE(&_BitbonStorageImpl.CallOpts)
}

// CONTRACTASSETBOXSTORAGE is a free data retrieval call binding the contract method 0xfa5bdafc.
//
// Solidity: function CONTRACT_ASSETBOX_STORAGE() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCallerSession) CONTRACTASSETBOXSTORAGE() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.CONTRACTASSETBOXSTORAGE(&_BitbonStorageImpl.CallOpts)
}

// CONTRACTBALANCE is a free data retrieval call binding the contract method 0xd59045f5.
//
// Solidity: function CONTRACT_BALANCE() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCaller) CONTRACTBALANCE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonStorageImpl.contract.Call(opts, out, "CONTRACT_BALANCE")
	return *ret0, err
}

// CONTRACTBALANCE is a free data retrieval call binding the contract method 0xd59045f5.
//
// Solidity: function CONTRACT_BALANCE() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplSession) CONTRACTBALANCE() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.CONTRACTBALANCE(&_BitbonStorageImpl.CallOpts)
}

// CONTRACTBALANCE is a free data retrieval call binding the contract method 0xd59045f5.
//
// Solidity: function CONTRACT_BALANCE() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCallerSession) CONTRACTBALANCE() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.CONTRACTBALANCE(&_BitbonStorageImpl.CallOpts)
}

// CONTRACTBITBON is a free data retrieval call binding the contract method 0xb6b433ab.
//
// Solidity: function CONTRACT_BITBON() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCaller) CONTRACTBITBON(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonStorageImpl.contract.Call(opts, out, "CONTRACT_BITBON")
	return *ret0, err
}

// CONTRACTBITBON is a free data retrieval call binding the contract method 0xb6b433ab.
//
// Solidity: function CONTRACT_BITBON() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplSession) CONTRACTBITBON() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.CONTRACTBITBON(&_BitbonStorageImpl.CallOpts)
}

// CONTRACTBITBON is a free data retrieval call binding the contract method 0xb6b433ab.
//
// Solidity: function CONTRACT_BITBON() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCallerSession) CONTRACTBITBON() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.CONTRACTBITBON(&_BitbonStorageImpl.CallOpts)
}

// CONTRACTBITBONSTORAGE is a free data retrieval call binding the contract method 0xe7f6c30e.
//
// Solidity: function CONTRACT_BITBON_STORAGE() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCaller) CONTRACTBITBONSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonStorageImpl.contract.Call(opts, out, "CONTRACT_BITBON_STORAGE")
	return *ret0, err
}

// CONTRACTBITBONSTORAGE is a free data retrieval call binding the contract method 0xe7f6c30e.
//
// Solidity: function CONTRACT_BITBON_STORAGE() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplSession) CONTRACTBITBONSTORAGE() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.CONTRACTBITBONSTORAGE(&_BitbonStorageImpl.CallOpts)
}

// CONTRACTBITBONSTORAGE is a free data retrieval call binding the contract method 0xe7f6c30e.
//
// Solidity: function CONTRACT_BITBON_STORAGE() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCallerSession) CONTRACTBITBONSTORAGE() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.CONTRACTBITBONSTORAGE(&_BitbonStorageImpl.CallOpts)
}

// CONTRACTBITBONSUPPORT is a free data retrieval call binding the contract method 0x20139c74.
//
// Solidity: function CONTRACT_BITBON_SUPPORT() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCaller) CONTRACTBITBONSUPPORT(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonStorageImpl.contract.Call(opts, out, "CONTRACT_BITBON_SUPPORT")
	return *ret0, err
}

// CONTRACTBITBONSUPPORT is a free data retrieval call binding the contract method 0x20139c74.
//
// Solidity: function CONTRACT_BITBON_SUPPORT() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplSession) CONTRACTBITBONSUPPORT() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.CONTRACTBITBONSUPPORT(&_BitbonStorageImpl.CallOpts)
}

// CONTRACTBITBONSUPPORT is a free data retrieval call binding the contract method 0x20139c74.
//
// Solidity: function CONTRACT_BITBON_SUPPORT() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCallerSession) CONTRACTBITBONSUPPORT() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.CONTRACTBITBONSUPPORT(&_BitbonStorageImpl.CallOpts)
}

// CONTRACTBONBALANCE is a free data retrieval call binding the contract method 0x40e02b7c.
//
// Solidity: function CONTRACT_BONBALANCE() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCaller) CONTRACTBONBALANCE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonStorageImpl.contract.Call(opts, out, "CONTRACT_BONBALANCE")
	return *ret0, err
}

// CONTRACTBONBALANCE is a free data retrieval call binding the contract method 0x40e02b7c.
//
// Solidity: function CONTRACT_BONBALANCE() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplSession) CONTRACTBONBALANCE() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.CONTRACTBONBALANCE(&_BitbonStorageImpl.CallOpts)
}

// CONTRACTBONBALANCE is a free data retrieval call binding the contract method 0x40e02b7c.
//
// Solidity: function CONTRACT_BONBALANCE() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCallerSession) CONTRACTBONBALANCE() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.CONTRACTBONBALANCE(&_BitbonStorageImpl.CallOpts)
}

// CONTRACTBONTRANSFER is a free data retrieval call binding the contract method 0xc18f4a8a.
//
// Solidity: function CONTRACT_BONTRANSFER() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCaller) CONTRACTBONTRANSFER(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonStorageImpl.contract.Call(opts, out, "CONTRACT_BONTRANSFER")
	return *ret0, err
}

// CONTRACTBONTRANSFER is a free data retrieval call binding the contract method 0xc18f4a8a.
//
// Solidity: function CONTRACT_BONTRANSFER() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplSession) CONTRACTBONTRANSFER() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.CONTRACTBONTRANSFER(&_BitbonStorageImpl.CallOpts)
}

// CONTRACTBONTRANSFER is a free data retrieval call binding the contract method 0xc18f4a8a.
//
// Solidity: function CONTRACT_BONTRANSFER() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCallerSession) CONTRACTBONTRANSFER() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.CONTRACTBONTRANSFER(&_BitbonStorageImpl.CallOpts)
}

// CONTRACTDEX is a free data retrieval call binding the contract method 0x52c1f844.
//
// Solidity: function CONTRACT_DEX() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCaller) CONTRACTDEX(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonStorageImpl.contract.Call(opts, out, "CONTRACT_DEX")
	return *ret0, err
}

// CONTRACTDEX is a free data retrieval call binding the contract method 0x52c1f844.
//
// Solidity: function CONTRACT_DEX() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplSession) CONTRACTDEX() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.CONTRACTDEX(&_BitbonStorageImpl.CallOpts)
}

// CONTRACTDEX is a free data retrieval call binding the contract method 0x52c1f844.
//
// Solidity: function CONTRACT_DEX() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCallerSession) CONTRACTDEX() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.CONTRACTDEX(&_BitbonStorageImpl.CallOpts)
}

// CONTRACTDISTRIBUTIONSTORAGE is a free data retrieval call binding the contract method 0x08097327.
//
// Solidity: function CONTRACT_DISTRIBUTION_STORAGE() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCaller) CONTRACTDISTRIBUTIONSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonStorageImpl.contract.Call(opts, out, "CONTRACT_DISTRIBUTION_STORAGE")
	return *ret0, err
}

// CONTRACTDISTRIBUTIONSTORAGE is a free data retrieval call binding the contract method 0x08097327.
//
// Solidity: function CONTRACT_DISTRIBUTION_STORAGE() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplSession) CONTRACTDISTRIBUTIONSTORAGE() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.CONTRACTDISTRIBUTIONSTORAGE(&_BitbonStorageImpl.CallOpts)
}

// CONTRACTDISTRIBUTIONSTORAGE is a free data retrieval call binding the contract method 0x08097327.
//
// Solidity: function CONTRACT_DISTRIBUTION_STORAGE() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCallerSession) CONTRACTDISTRIBUTIONSTORAGE() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.CONTRACTDISTRIBUTIONSTORAGE(&_BitbonStorageImpl.CallOpts)
}

// CONTRACTEXCHANGE is a free data retrieval call binding the contract method 0xa0e3a10e.
//
// Solidity: function CONTRACT_EXCHANGE() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCaller) CONTRACTEXCHANGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonStorageImpl.contract.Call(opts, out, "CONTRACT_EXCHANGE")
	return *ret0, err
}

// CONTRACTEXCHANGE is a free data retrieval call binding the contract method 0xa0e3a10e.
//
// Solidity: function CONTRACT_EXCHANGE() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplSession) CONTRACTEXCHANGE() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.CONTRACTEXCHANGE(&_BitbonStorageImpl.CallOpts)
}

// CONTRACTEXCHANGE is a free data retrieval call binding the contract method 0xa0e3a10e.
//
// Solidity: function CONTRACT_EXCHANGE() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCallerSession) CONTRACTEXCHANGE() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.CONTRACTEXCHANGE(&_BitbonStorageImpl.CallOpts)
}

// CONTRACTEXCHANGESTORAGE is a free data retrieval call binding the contract method 0x001570f2.
//
// Solidity: function CONTRACT_EXCHANGE_STORAGE() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCaller) CONTRACTEXCHANGESTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonStorageImpl.contract.Call(opts, out, "CONTRACT_EXCHANGE_STORAGE")
	return *ret0, err
}

// CONTRACTEXCHANGESTORAGE is a free data retrieval call binding the contract method 0x001570f2.
//
// Solidity: function CONTRACT_EXCHANGE_STORAGE() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplSession) CONTRACTEXCHANGESTORAGE() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.CONTRACTEXCHANGESTORAGE(&_BitbonStorageImpl.CallOpts)
}

// CONTRACTEXCHANGESTORAGE is a free data retrieval call binding the contract method 0x001570f2.
//
// Solidity: function CONTRACT_EXCHANGE_STORAGE() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCallerSession) CONTRACTEXCHANGESTORAGE() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.CONTRACTEXCHANGESTORAGE(&_BitbonStorageImpl.CallOpts)
}

// CONTRACTFEESTORAGE is a free data retrieval call binding the contract method 0xcd5e65a4.
//
// Solidity: function CONTRACT_FEE_STORAGE() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCaller) CONTRACTFEESTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonStorageImpl.contract.Call(opts, out, "CONTRACT_FEE_STORAGE")
	return *ret0, err
}

// CONTRACTFEESTORAGE is a free data retrieval call binding the contract method 0xcd5e65a4.
//
// Solidity: function CONTRACT_FEE_STORAGE() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplSession) CONTRACTFEESTORAGE() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.CONTRACTFEESTORAGE(&_BitbonStorageImpl.CallOpts)
}

// CONTRACTFEESTORAGE is a free data retrieval call binding the contract method 0xcd5e65a4.
//
// Solidity: function CONTRACT_FEE_STORAGE() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCallerSession) CONTRACTFEESTORAGE() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.CONTRACTFEESTORAGE(&_BitbonStorageImpl.CallOpts)
}

// CONTRACTGENERATOR is a free data retrieval call binding the contract method 0xccea2b46.
//
// Solidity: function CONTRACT_GENERATOR() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCaller) CONTRACTGENERATOR(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonStorageImpl.contract.Call(opts, out, "CONTRACT_GENERATOR")
	return *ret0, err
}

// CONTRACTGENERATOR is a free data retrieval call binding the contract method 0xccea2b46.
//
// Solidity: function CONTRACT_GENERATOR() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplSession) CONTRACTGENERATOR() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.CONTRACTGENERATOR(&_BitbonStorageImpl.CallOpts)
}

// CONTRACTGENERATOR is a free data retrieval call binding the contract method 0xccea2b46.
//
// Solidity: function CONTRACT_GENERATOR() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCallerSession) CONTRACTGENERATOR() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.CONTRACTGENERATOR(&_BitbonStorageImpl.CallOpts)
}

// CONTRACTMININGAGENTSTORAGE is a free data retrieval call binding the contract method 0x83cc588d.
//
// Solidity: function CONTRACT_MINING_AGENT_STORAGE() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCaller) CONTRACTMININGAGENTSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonStorageImpl.contract.Call(opts, out, "CONTRACT_MINING_AGENT_STORAGE")
	return *ret0, err
}

// CONTRACTMININGAGENTSTORAGE is a free data retrieval call binding the contract method 0x83cc588d.
//
// Solidity: function CONTRACT_MINING_AGENT_STORAGE() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplSession) CONTRACTMININGAGENTSTORAGE() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.CONTRACTMININGAGENTSTORAGE(&_BitbonStorageImpl.CallOpts)
}

// CONTRACTMININGAGENTSTORAGE is a free data retrieval call binding the contract method 0x83cc588d.
//
// Solidity: function CONTRACT_MINING_AGENT_STORAGE() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCallerSession) CONTRACTMININGAGENTSTORAGE() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.CONTRACTMININGAGENTSTORAGE(&_BitbonStorageImpl.CallOpts)
}

// CONTRACTMSBON is a free data retrieval call binding the contract method 0x28b3a976.
//
// Solidity: function CONTRACT_MSBON() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCaller) CONTRACTMSBON(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonStorageImpl.contract.Call(opts, out, "CONTRACT_MSBON")
	return *ret0, err
}

// CONTRACTMSBON is a free data retrieval call binding the contract method 0x28b3a976.
//
// Solidity: function CONTRACT_MSBON() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplSession) CONTRACTMSBON() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.CONTRACTMSBON(&_BitbonStorageImpl.CallOpts)
}

// CONTRACTMSBON is a free data retrieval call binding the contract method 0x28b3a976.
//
// Solidity: function CONTRACT_MSBON() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCallerSession) CONTRACTMSBON() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.CONTRACTMSBON(&_BitbonStorageImpl.CallOpts)
}

// CONTRACTMSBONSTORAGE is a free data retrieval call binding the contract method 0xe04e6061.
//
// Solidity: function CONTRACT_MSBON_STORAGE() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCaller) CONTRACTMSBONSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonStorageImpl.contract.Call(opts, out, "CONTRACT_MSBON_STORAGE")
	return *ret0, err
}

// CONTRACTMSBONSTORAGE is a free data retrieval call binding the contract method 0xe04e6061.
//
// Solidity: function CONTRACT_MSBON_STORAGE() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplSession) CONTRACTMSBONSTORAGE() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.CONTRACTMSBONSTORAGE(&_BitbonStorageImpl.CallOpts)
}

// CONTRACTMSBONSTORAGE is a free data retrieval call binding the contract method 0xe04e6061.
//
// Solidity: function CONTRACT_MSBON_STORAGE() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCallerSession) CONTRACTMSBONSTORAGE() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.CONTRACTMSBONSTORAGE(&_BitbonStorageImpl.CallOpts)
}

// CONTRACTMULTISIGADDADDRESS is a free data retrieval call binding the contract method 0xf4f65a5a.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ADDRESS() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCaller) CONTRACTMULTISIGADDADDRESS(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_ADD_ADDRESS")
	return *ret0, err
}

// CONTRACTMULTISIGADDADDRESS is a free data retrieval call binding the contract method 0xf4f65a5a.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ADDRESS() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplSession) CONTRACTMULTISIGADDADDRESS() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.CONTRACTMULTISIGADDADDRESS(&_BitbonStorageImpl.CallOpts)
}

// CONTRACTMULTISIGADDADDRESS is a free data retrieval call binding the contract method 0xf4f65a5a.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ADDRESS() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCallerSession) CONTRACTMULTISIGADDADDRESS() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.CONTRACTMULTISIGADDADDRESS(&_BitbonStorageImpl.CallOpts)
}

// CONTRACTMULTISIGADDADMIN is a free data retrieval call binding the contract method 0x83888c68.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ADMIN() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCaller) CONTRACTMULTISIGADDADMIN(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_ADD_ADMIN")
	return *ret0, err
}

// CONTRACTMULTISIGADDADMIN is a free data retrieval call binding the contract method 0x83888c68.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ADMIN() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplSession) CONTRACTMULTISIGADDADMIN() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.CONTRACTMULTISIGADDADMIN(&_BitbonStorageImpl.CallOpts)
}

// CONTRACTMULTISIGADDADMIN is a free data retrieval call binding the contract method 0x83888c68.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ADMIN() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCallerSession) CONTRACTMULTISIGADDADMIN() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.CONTRACTMULTISIGADDADMIN(&_BitbonStorageImpl.CallOpts)
}

// CONTRACTMULTISIGADDFEE is a free data retrieval call binding the contract method 0xc43e8cab.
//
// Solidity: function CONTRACT_MULTISIG_ADD_FEE() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCaller) CONTRACTMULTISIGADDFEE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_ADD_FEE")
	return *ret0, err
}

// CONTRACTMULTISIGADDFEE is a free data retrieval call binding the contract method 0xc43e8cab.
//
// Solidity: function CONTRACT_MULTISIG_ADD_FEE() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplSession) CONTRACTMULTISIGADDFEE() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.CONTRACTMULTISIGADDFEE(&_BitbonStorageImpl.CallOpts)
}

// CONTRACTMULTISIGADDFEE is a free data retrieval call binding the contract method 0xc43e8cab.
//
// Solidity: function CONTRACT_MULTISIG_ADD_FEE() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCallerSession) CONTRACTMULTISIGADDFEE() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.CONTRACTMULTISIGADDFEE(&_BitbonStorageImpl.CallOpts)
}

// CONTRACTMULTISIGADDROLE is a free data retrieval call binding the contract method 0x20a0914b.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ROLE() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCaller) CONTRACTMULTISIGADDROLE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_ADD_ROLE")
	return *ret0, err
}

// CONTRACTMULTISIGADDROLE is a free data retrieval call binding the contract method 0x20a0914b.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ROLE() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplSession) CONTRACTMULTISIGADDROLE() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.CONTRACTMULTISIGADDROLE(&_BitbonStorageImpl.CallOpts)
}

// CONTRACTMULTISIGADDROLE is a free data retrieval call binding the contract method 0x20a0914b.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ROLE() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCallerSession) CONTRACTMULTISIGADDROLE() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.CONTRACTMULTISIGADDROLE(&_BitbonStorageImpl.CallOpts)
}

// CONTRACTMULTISIGDROPPERMISSION is a free data retrieval call binding the contract method 0x3f3fbc39.
//
// Solidity: function CONTRACT_MULTISIG_DROP_PERMISSION() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCaller) CONTRACTMULTISIGDROPPERMISSION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_DROP_PERMISSION")
	return *ret0, err
}

// CONTRACTMULTISIGDROPPERMISSION is a free data retrieval call binding the contract method 0x3f3fbc39.
//
// Solidity: function CONTRACT_MULTISIG_DROP_PERMISSION() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplSession) CONTRACTMULTISIGDROPPERMISSION() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.CONTRACTMULTISIGDROPPERMISSION(&_BitbonStorageImpl.CallOpts)
}

// CONTRACTMULTISIGDROPPERMISSION is a free data retrieval call binding the contract method 0x3f3fbc39.
//
// Solidity: function CONTRACT_MULTISIG_DROP_PERMISSION() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCallerSession) CONTRACTMULTISIGDROPPERMISSION() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.CONTRACTMULTISIGDROPPERMISSION(&_BitbonStorageImpl.CallOpts)
}

// CONTRACTMULTISIGEDITADMIN is a free data retrieval call binding the contract method 0x91b0f051.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_ADMIN() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCaller) CONTRACTMULTISIGEDITADMIN(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_EDIT_ADMIN")
	return *ret0, err
}

// CONTRACTMULTISIGEDITADMIN is a free data retrieval call binding the contract method 0x91b0f051.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_ADMIN() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplSession) CONTRACTMULTISIGEDITADMIN() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.CONTRACTMULTISIGEDITADMIN(&_BitbonStorageImpl.CallOpts)
}

// CONTRACTMULTISIGEDITADMIN is a free data retrieval call binding the contract method 0x91b0f051.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_ADMIN() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCallerSession) CONTRACTMULTISIGEDITADMIN() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.CONTRACTMULTISIGEDITADMIN(&_BitbonStorageImpl.CallOpts)
}

// CONTRACTMULTISIGEDITFEE is a free data retrieval call binding the contract method 0x6b2abe17.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_FEE() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCaller) CONTRACTMULTISIGEDITFEE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_EDIT_FEE")
	return *ret0, err
}

// CONTRACTMULTISIGEDITFEE is a free data retrieval call binding the contract method 0x6b2abe17.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_FEE() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplSession) CONTRACTMULTISIGEDITFEE() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.CONTRACTMULTISIGEDITFEE(&_BitbonStorageImpl.CallOpts)
}

// CONTRACTMULTISIGEDITFEE is a free data retrieval call binding the contract method 0x6b2abe17.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_FEE() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCallerSession) CONTRACTMULTISIGEDITFEE() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.CONTRACTMULTISIGEDITFEE(&_BitbonStorageImpl.CallOpts)
}

// CONTRACTMULTISIGEDITROLE is a free data retrieval call binding the contract method 0xc5f66412.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_ROLE() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCaller) CONTRACTMULTISIGEDITROLE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_EDIT_ROLE")
	return *ret0, err
}

// CONTRACTMULTISIGEDITROLE is a free data retrieval call binding the contract method 0xc5f66412.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_ROLE() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplSession) CONTRACTMULTISIGEDITROLE() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.CONTRACTMULTISIGEDITROLE(&_BitbonStorageImpl.CallOpts)
}

// CONTRACTMULTISIGEDITROLE is a free data retrieval call binding the contract method 0xc5f66412.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_ROLE() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCallerSession) CONTRACTMULTISIGEDITROLE() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.CONTRACTMULTISIGEDITROLE(&_BitbonStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEADMINBYADDRESS is a free data retrieval call binding the contract method 0xc816b091.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_BY_ADDRESS() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCaller) CONTRACTMULTISIGREMOVEADMINBYADDRESS(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REMOVE_ADMIN_BY_ADDRESS")
	return *ret0, err
}

// CONTRACTMULTISIGREMOVEADMINBYADDRESS is a free data retrieval call binding the contract method 0xc816b091.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_BY_ADDRESS() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplSession) CONTRACTMULTISIGREMOVEADMINBYADDRESS() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.CONTRACTMULTISIGREMOVEADMINBYADDRESS(&_BitbonStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEADMINBYADDRESS is a free data retrieval call binding the contract method 0xc816b091.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_BY_ADDRESS() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCallerSession) CONTRACTMULTISIGREMOVEADMINBYADDRESS() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.CONTRACTMULTISIGREMOVEADMINBYADDRESS(&_BitbonStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEADMINBYINDEX is a free data retrieval call binding the contract method 0x751fbea2.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_BY_INDEX() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCaller) CONTRACTMULTISIGREMOVEADMINBYINDEX(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REMOVE_ADMIN_BY_INDEX")
	return *ret0, err
}

// CONTRACTMULTISIGREMOVEADMINBYINDEX is a free data retrieval call binding the contract method 0x751fbea2.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_BY_INDEX() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplSession) CONTRACTMULTISIGREMOVEADMINBYINDEX() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.CONTRACTMULTISIGREMOVEADMINBYINDEX(&_BitbonStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEADMINBYINDEX is a free data retrieval call binding the contract method 0x751fbea2.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_BY_INDEX() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCallerSession) CONTRACTMULTISIGREMOVEADMINBYINDEX() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.CONTRACTMULTISIGREMOVEADMINBYINDEX(&_BitbonStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEADMINROLE is a free data retrieval call binding the contract method 0x53c00200.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_ROLE() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCaller) CONTRACTMULTISIGREMOVEADMINROLE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REMOVE_ADMIN_ROLE")
	return *ret0, err
}

// CONTRACTMULTISIGREMOVEADMINROLE is a free data retrieval call binding the contract method 0x53c00200.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_ROLE() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplSession) CONTRACTMULTISIGREMOVEADMINROLE() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.CONTRACTMULTISIGREMOVEADMINROLE(&_BitbonStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEADMINROLE is a free data retrieval call binding the contract method 0x53c00200.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_ROLE() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCallerSession) CONTRACTMULTISIGREMOVEADMINROLE() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.CONTRACTMULTISIGREMOVEADMINROLE(&_BitbonStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEFEE is a free data retrieval call binding the contract method 0xdb614d25.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_FEE() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCaller) CONTRACTMULTISIGREMOVEFEE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REMOVE_FEE")
	return *ret0, err
}

// CONTRACTMULTISIGREMOVEFEE is a free data retrieval call binding the contract method 0xdb614d25.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_FEE() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplSession) CONTRACTMULTISIGREMOVEFEE() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.CONTRACTMULTISIGREMOVEFEE(&_BitbonStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEFEE is a free data retrieval call binding the contract method 0xdb614d25.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_FEE() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCallerSession) CONTRACTMULTISIGREMOVEFEE() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.CONTRACTMULTISIGREMOVEFEE(&_BitbonStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEROLE is a free data retrieval call binding the contract method 0x23648496.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ROLE() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCaller) CONTRACTMULTISIGREMOVEROLE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REMOVE_ROLE")
	return *ret0, err
}

// CONTRACTMULTISIGREMOVEROLE is a free data retrieval call binding the contract method 0x23648496.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ROLE() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplSession) CONTRACTMULTISIGREMOVEROLE() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.CONTRACTMULTISIGREMOVEROLE(&_BitbonStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEROLE is a free data retrieval call binding the contract method 0x23648496.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ROLE() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCallerSession) CONTRACTMULTISIGREMOVEROLE() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.CONTRACTMULTISIGREMOVEROLE(&_BitbonStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREPLENISHCAPITALIZATION is a free data retrieval call binding the contract method 0x62d5d523.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_CAPITALIZATION() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCaller) CONTRACTMULTISIGREPLENISHCAPITALIZATION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REPLENISH_CAPITALIZATION")
	return *ret0, err
}

// CONTRACTMULTISIGREPLENISHCAPITALIZATION is a free data retrieval call binding the contract method 0x62d5d523.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_CAPITALIZATION() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplSession) CONTRACTMULTISIGREPLENISHCAPITALIZATION() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.CONTRACTMULTISIGREPLENISHCAPITALIZATION(&_BitbonStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREPLENISHCAPITALIZATION is a free data retrieval call binding the contract method 0x62d5d523.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_CAPITALIZATION() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCallerSession) CONTRACTMULTISIGREPLENISHCAPITALIZATION() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.CONTRACTMULTISIGREPLENISHCAPITALIZATION(&_BitbonStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREPLENISHCERTIFICATE is a free data retrieval call binding the contract method 0xfd19b7e1.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_CERTIFICATE() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCaller) CONTRACTMULTISIGREPLENISHCERTIFICATE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REPLENISH_CERTIFICATE")
	return *ret0, err
}

// CONTRACTMULTISIGREPLENISHCERTIFICATE is a free data retrieval call binding the contract method 0xfd19b7e1.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_CERTIFICATE() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplSession) CONTRACTMULTISIGREPLENISHCERTIFICATE() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.CONTRACTMULTISIGREPLENISHCERTIFICATE(&_BitbonStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREPLENISHCERTIFICATE is a free data retrieval call binding the contract method 0xfd19b7e1.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_CERTIFICATE() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCallerSession) CONTRACTMULTISIGREPLENISHCERTIFICATE() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.CONTRACTMULTISIGREPLENISHCERTIFICATE(&_BitbonStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREPLENISHEMISSION is a free data retrieval call binding the contract method 0x0cb59fde.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_EMISSION() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCaller) CONTRACTMULTISIGREPLENISHEMISSION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REPLENISH_EMISSION")
	return *ret0, err
}

// CONTRACTMULTISIGREPLENISHEMISSION is a free data retrieval call binding the contract method 0x0cb59fde.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_EMISSION() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplSession) CONTRACTMULTISIGREPLENISHEMISSION() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.CONTRACTMULTISIGREPLENISHEMISSION(&_BitbonStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREPLENISHEMISSION is a free data retrieval call binding the contract method 0x0cb59fde.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_EMISSION() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCallerSession) CONTRACTMULTISIGREPLENISHEMISSION() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.CONTRACTMULTISIGREPLENISHEMISSION(&_BitbonStorageImpl.CallOpts)
}

// CONTRACTMULTISIGSETADMINROLE is a free data retrieval call binding the contract method 0x8f209ab7.
//
// Solidity: function CONTRACT_MULTISIG_SET_ADMIN_ROLE() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCaller) CONTRACTMULTISIGSETADMINROLE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_SET_ADMIN_ROLE")
	return *ret0, err
}

// CONTRACTMULTISIGSETADMINROLE is a free data retrieval call binding the contract method 0x8f209ab7.
//
// Solidity: function CONTRACT_MULTISIG_SET_ADMIN_ROLE() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplSession) CONTRACTMULTISIGSETADMINROLE() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.CONTRACTMULTISIGSETADMINROLE(&_BitbonStorageImpl.CallOpts)
}

// CONTRACTMULTISIGSETADMINROLE is a free data retrieval call binding the contract method 0x8f209ab7.
//
// Solidity: function CONTRACT_MULTISIG_SET_ADMIN_ROLE() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCallerSession) CONTRACTMULTISIGSETADMINROLE() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.CONTRACTMULTISIGSETADMINROLE(&_BitbonStorageImpl.CallOpts)
}

// CONTRACTMULTISIGSETPERMISSION is a free data retrieval call binding the contract method 0xb0e91914.
//
// Solidity: function CONTRACT_MULTISIG_SET_PERMISSION() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCaller) CONTRACTMULTISIGSETPERMISSION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_SET_PERMISSION")
	return *ret0, err
}

// CONTRACTMULTISIGSETPERMISSION is a free data retrieval call binding the contract method 0xb0e91914.
//
// Solidity: function CONTRACT_MULTISIG_SET_PERMISSION() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplSession) CONTRACTMULTISIGSETPERMISSION() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.CONTRACTMULTISIGSETPERMISSION(&_BitbonStorageImpl.CallOpts)
}

// CONTRACTMULTISIGSETPERMISSION is a free data retrieval call binding the contract method 0xb0e91914.
//
// Solidity: function CONTRACT_MULTISIG_SET_PERMISSION() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCallerSession) CONTRACTMULTISIGSETPERMISSION() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.CONTRACTMULTISIGSETPERMISSION(&_BitbonStorageImpl.CallOpts)
}

// CONTRACTMULTISIGSTORAGE is a free data retrieval call binding the contract method 0xe7e1a303.
//
// Solidity: function CONTRACT_MULTISIG_STORAGE() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCaller) CONTRACTMULTISIGSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_STORAGE")
	return *ret0, err
}

// CONTRACTMULTISIGSTORAGE is a free data retrieval call binding the contract method 0xe7e1a303.
//
// Solidity: function CONTRACT_MULTISIG_STORAGE() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplSession) CONTRACTMULTISIGSTORAGE() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.CONTRACTMULTISIGSTORAGE(&_BitbonStorageImpl.CallOpts)
}

// CONTRACTMULTISIGSTORAGE is a free data retrieval call binding the contract method 0xe7e1a303.
//
// Solidity: function CONTRACT_MULTISIG_STORAGE() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCallerSession) CONTRACTMULTISIGSTORAGE() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.CONTRACTMULTISIGSTORAGE(&_BitbonStorageImpl.CallOpts)
}

// CONTRACTOTC is a free data retrieval call binding the contract method 0x5eceac70.
//
// Solidity: function CONTRACT_OTC() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCaller) CONTRACTOTC(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonStorageImpl.contract.Call(opts, out, "CONTRACT_OTC")
	return *ret0, err
}

// CONTRACTOTC is a free data retrieval call binding the contract method 0x5eceac70.
//
// Solidity: function CONTRACT_OTC() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplSession) CONTRACTOTC() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.CONTRACTOTC(&_BitbonStorageImpl.CallOpts)
}

// CONTRACTOTC is a free data retrieval call binding the contract method 0x5eceac70.
//
// Solidity: function CONTRACT_OTC() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCallerSession) CONTRACTOTC() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.CONTRACTOTC(&_BitbonStorageImpl.CallOpts)
}

// CONTRACTOTCSTORAGE is a free data retrieval call binding the contract method 0x825b476f.
//
// Solidity: function CONTRACT_OTC_STORAGE() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCaller) CONTRACTOTCSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonStorageImpl.contract.Call(opts, out, "CONTRACT_OTC_STORAGE")
	return *ret0, err
}

// CONTRACTOTCSTORAGE is a free data retrieval call binding the contract method 0x825b476f.
//
// Solidity: function CONTRACT_OTC_STORAGE() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplSession) CONTRACTOTCSTORAGE() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.CONTRACTOTCSTORAGE(&_BitbonStorageImpl.CallOpts)
}

// CONTRACTOTCSTORAGE is a free data retrieval call binding the contract method 0x825b476f.
//
// Solidity: function CONTRACT_OTC_STORAGE() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCallerSession) CONTRACTOTCSTORAGE() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.CONTRACTOTCSTORAGE(&_BitbonStorageImpl.CallOpts)
}

// CONTRACTREPLENISHCAPITALIZATION is a free data retrieval call binding the contract method 0xff5e09d8.
//
// Solidity: function CONTRACT_REPLENISH_CAPITALIZATION() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCaller) CONTRACTREPLENISHCAPITALIZATION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonStorageImpl.contract.Call(opts, out, "CONTRACT_REPLENISH_CAPITALIZATION")
	return *ret0, err
}

// CONTRACTREPLENISHCAPITALIZATION is a free data retrieval call binding the contract method 0xff5e09d8.
//
// Solidity: function CONTRACT_REPLENISH_CAPITALIZATION() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplSession) CONTRACTREPLENISHCAPITALIZATION() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.CONTRACTREPLENISHCAPITALIZATION(&_BitbonStorageImpl.CallOpts)
}

// CONTRACTREPLENISHCAPITALIZATION is a free data retrieval call binding the contract method 0xff5e09d8.
//
// Solidity: function CONTRACT_REPLENISH_CAPITALIZATION() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCallerSession) CONTRACTREPLENISHCAPITALIZATION() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.CONTRACTREPLENISHCAPITALIZATION(&_BitbonStorageImpl.CallOpts)
}

// CONTRACTRESERVEDALIASSTORAGE is a free data retrieval call binding the contract method 0x70a94aa7.
//
// Solidity: function CONTRACT_RESERVED_ALIAS_STORAGE() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCaller) CONTRACTRESERVEDALIASSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonStorageImpl.contract.Call(opts, out, "CONTRACT_RESERVED_ALIAS_STORAGE")
	return *ret0, err
}

// CONTRACTRESERVEDALIASSTORAGE is a free data retrieval call binding the contract method 0x70a94aa7.
//
// Solidity: function CONTRACT_RESERVED_ALIAS_STORAGE() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplSession) CONTRACTRESERVEDALIASSTORAGE() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.CONTRACTRESERVEDALIASSTORAGE(&_BitbonStorageImpl.CallOpts)
}

// CONTRACTRESERVEDALIASSTORAGE is a free data retrieval call binding the contract method 0x70a94aa7.
//
// Solidity: function CONTRACT_RESERVED_ALIAS_STORAGE() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCallerSession) CONTRACTRESERVEDALIASSTORAGE() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.CONTRACTRESERVEDALIASSTORAGE(&_BitbonStorageImpl.CallOpts)
}

// CONTRACTROLESTORAGE is a free data retrieval call binding the contract method 0xb938235b.
//
// Solidity: function CONTRACT_ROLE_STORAGE() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCaller) CONTRACTROLESTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonStorageImpl.contract.Call(opts, out, "CONTRACT_ROLE_STORAGE")
	return *ret0, err
}

// CONTRACTROLESTORAGE is a free data retrieval call binding the contract method 0xb938235b.
//
// Solidity: function CONTRACT_ROLE_STORAGE() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplSession) CONTRACTROLESTORAGE() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.CONTRACTROLESTORAGE(&_BitbonStorageImpl.CallOpts)
}

// CONTRACTROLESTORAGE is a free data retrieval call binding the contract method 0xb938235b.
//
// Solidity: function CONTRACT_ROLE_STORAGE() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCallerSession) CONTRACTROLESTORAGE() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.CONTRACTROLESTORAGE(&_BitbonStorageImpl.CallOpts)
}

// CONTRACTSAFETRANSFERSTORAGE is a free data retrieval call binding the contract method 0xe7e45392.
//
// Solidity: function CONTRACT_SAFE_TRANSFER_STORAGE() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCaller) CONTRACTSAFETRANSFERSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonStorageImpl.contract.Call(opts, out, "CONTRACT_SAFE_TRANSFER_STORAGE")
	return *ret0, err
}

// CONTRACTSAFETRANSFERSTORAGE is a free data retrieval call binding the contract method 0xe7e45392.
//
// Solidity: function CONTRACT_SAFE_TRANSFER_STORAGE() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplSession) CONTRACTSAFETRANSFERSTORAGE() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.CONTRACTSAFETRANSFERSTORAGE(&_BitbonStorageImpl.CallOpts)
}

// CONTRACTSAFETRANSFERSTORAGE is a free data retrieval call binding the contract method 0xe7e45392.
//
// Solidity: function CONTRACT_SAFE_TRANSFER_STORAGE() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCallerSession) CONTRACTSAFETRANSFERSTORAGE() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.CONTRACTSAFETRANSFERSTORAGE(&_BitbonStorageImpl.CallOpts)
}

// CONTRACTTRANSFER is a free data retrieval call binding the contract method 0xef94fe38.
//
// Solidity: function CONTRACT_TRANSFER() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCaller) CONTRACTTRANSFER(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonStorageImpl.contract.Call(opts, out, "CONTRACT_TRANSFER")
	return *ret0, err
}

// CONTRACTTRANSFER is a free data retrieval call binding the contract method 0xef94fe38.
//
// Solidity: function CONTRACT_TRANSFER() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplSession) CONTRACTTRANSFER() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.CONTRACTTRANSFER(&_BitbonStorageImpl.CallOpts)
}

// CONTRACTTRANSFER is a free data retrieval call binding the contract method 0xef94fe38.
//
// Solidity: function CONTRACT_TRANSFER() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCallerSession) CONTRACTTRANSFER() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.CONTRACTTRANSFER(&_BitbonStorageImpl.CallOpts)
}

// CONTRACTTRANSFERFROMCAPITALIZATION is a free data retrieval call binding the contract method 0xc8f8decb.
//
// Solidity: function CONTRACT_TRANSFER_FROM_CAPITALIZATION() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCaller) CONTRACTTRANSFERFROMCAPITALIZATION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonStorageImpl.contract.Call(opts, out, "CONTRACT_TRANSFER_FROM_CAPITALIZATION")
	return *ret0, err
}

// CONTRACTTRANSFERFROMCAPITALIZATION is a free data retrieval call binding the contract method 0xc8f8decb.
//
// Solidity: function CONTRACT_TRANSFER_FROM_CAPITALIZATION() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplSession) CONTRACTTRANSFERFROMCAPITALIZATION() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.CONTRACTTRANSFERFROMCAPITALIZATION(&_BitbonStorageImpl.CallOpts)
}

// CONTRACTTRANSFERFROMCAPITALIZATION is a free data retrieval call binding the contract method 0xc8f8decb.
//
// Solidity: function CONTRACT_TRANSFER_FROM_CAPITALIZATION() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCallerSession) CONTRACTTRANSFERFROMCAPITALIZATION() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.CONTRACTTRANSFERFROMCAPITALIZATION(&_BitbonStorageImpl.CallOpts)
}

// PERMISSIONACCESSRESTORATION is a free data retrieval call binding the contract method 0x5c3053bf.
//
// Solidity: function PERMISSION_ACCESS_RESTORATION() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCaller) PERMISSIONACCESSRESTORATION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonStorageImpl.contract.Call(opts, out, "PERMISSION_ACCESS_RESTORATION")
	return *ret0, err
}

// PERMISSIONACCESSRESTORATION is a free data retrieval call binding the contract method 0x5c3053bf.
//
// Solidity: function PERMISSION_ACCESS_RESTORATION() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplSession) PERMISSIONACCESSRESTORATION() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.PERMISSIONACCESSRESTORATION(&_BitbonStorageImpl.CallOpts)
}

// PERMISSIONACCESSRESTORATION is a free data retrieval call binding the contract method 0x5c3053bf.
//
// Solidity: function PERMISSION_ACCESS_RESTORATION() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCallerSession) PERMISSIONACCESSRESTORATION() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.PERMISSIONACCESSRESTORATION(&_BitbonStorageImpl.CallOpts)
}

// PERMISSIONADMINSTORAGE is a free data retrieval call binding the contract method 0x0c9fc0db.
//
// Solidity: function PERMISSION_ADMIN_STORAGE() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCaller) PERMISSIONADMINSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonStorageImpl.contract.Call(opts, out, "PERMISSION_ADMIN_STORAGE")
	return *ret0, err
}

// PERMISSIONADMINSTORAGE is a free data retrieval call binding the contract method 0x0c9fc0db.
//
// Solidity: function PERMISSION_ADMIN_STORAGE() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplSession) PERMISSIONADMINSTORAGE() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.PERMISSIONADMINSTORAGE(&_BitbonStorageImpl.CallOpts)
}

// PERMISSIONADMINSTORAGE is a free data retrieval call binding the contract method 0x0c9fc0db.
//
// Solidity: function PERMISSION_ADMIN_STORAGE() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCallerSession) PERMISSIONADMINSTORAGE() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.PERMISSIONADMINSTORAGE(&_BitbonStorageImpl.CallOpts)
}

// PERMISSIONCONTRACTSSTORAGE is a free data retrieval call binding the contract method 0x1516e83b.
//
// Solidity: function PERMISSION_CONTRACTS_STORAGE() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCaller) PERMISSIONCONTRACTSSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonStorageImpl.contract.Call(opts, out, "PERMISSION_CONTRACTS_STORAGE")
	return *ret0, err
}

// PERMISSIONCONTRACTSSTORAGE is a free data retrieval call binding the contract method 0x1516e83b.
//
// Solidity: function PERMISSION_CONTRACTS_STORAGE() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplSession) PERMISSIONCONTRACTSSTORAGE() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.PERMISSIONCONTRACTSSTORAGE(&_BitbonStorageImpl.CallOpts)
}

// PERMISSIONCONTRACTSSTORAGE is a free data retrieval call binding the contract method 0x1516e83b.
//
// Solidity: function PERMISSION_CONTRACTS_STORAGE() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCallerSession) PERMISSIONCONTRACTSSTORAGE() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.PERMISSIONCONTRACTSSTORAGE(&_BitbonStorageImpl.CallOpts)
}

// PERMISSIONCONTRACTDEPLOY is a free data retrieval call binding the contract method 0x257a4a04.
//
// Solidity: function PERMISSION_CONTRACT_DEPLOY() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCaller) PERMISSIONCONTRACTDEPLOY(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonStorageImpl.contract.Call(opts, out, "PERMISSION_CONTRACT_DEPLOY")
	return *ret0, err
}

// PERMISSIONCONTRACTDEPLOY is a free data retrieval call binding the contract method 0x257a4a04.
//
// Solidity: function PERMISSION_CONTRACT_DEPLOY() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplSession) PERMISSIONCONTRACTDEPLOY() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.PERMISSIONCONTRACTDEPLOY(&_BitbonStorageImpl.CallOpts)
}

// PERMISSIONCONTRACTDEPLOY is a free data retrieval call binding the contract method 0x257a4a04.
//
// Solidity: function PERMISSION_CONTRACT_DEPLOY() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCallerSession) PERMISSIONCONTRACTDEPLOY() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.PERMISSIONCONTRACTDEPLOY(&_BitbonStorageImpl.CallOpts)
}

// PERMISSIONEMISSION is a free data retrieval call binding the contract method 0x4ccd92d1.
//
// Solidity: function PERMISSION_EMISSION() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCaller) PERMISSIONEMISSION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonStorageImpl.contract.Call(opts, out, "PERMISSION_EMISSION")
	return *ret0, err
}

// PERMISSIONEMISSION is a free data retrieval call binding the contract method 0x4ccd92d1.
//
// Solidity: function PERMISSION_EMISSION() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplSession) PERMISSIONEMISSION() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.PERMISSIONEMISSION(&_BitbonStorageImpl.CallOpts)
}

// PERMISSIONEMISSION is a free data retrieval call binding the contract method 0x4ccd92d1.
//
// Solidity: function PERMISSION_EMISSION() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCallerSession) PERMISSIONEMISSION() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.PERMISSIONEMISSION(&_BitbonStorageImpl.CallOpts)
}

// PERMISSIONFEESSTORAGE is a free data retrieval call binding the contract method 0xe837eb87.
//
// Solidity: function PERMISSION_FEES_STORAGE() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCaller) PERMISSIONFEESSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonStorageImpl.contract.Call(opts, out, "PERMISSION_FEES_STORAGE")
	return *ret0, err
}

// PERMISSIONFEESSTORAGE is a free data retrieval call binding the contract method 0xe837eb87.
//
// Solidity: function PERMISSION_FEES_STORAGE() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplSession) PERMISSIONFEESSTORAGE() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.PERMISSIONFEESSTORAGE(&_BitbonStorageImpl.CallOpts)
}

// PERMISSIONFEESSTORAGE is a free data retrieval call binding the contract method 0xe837eb87.
//
// Solidity: function PERMISSION_FEES_STORAGE() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCallerSession) PERMISSIONFEESSTORAGE() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.PERMISSIONFEESSTORAGE(&_BitbonStorageImpl.CallOpts)
}

// PERMISSIONMULTISIGSTORAGE is a free data retrieval call binding the contract method 0xf1bb7aad.
//
// Solidity: function PERMISSION_MULTISIG_STORAGE() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCaller) PERMISSIONMULTISIGSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonStorageImpl.contract.Call(opts, out, "PERMISSION_MULTISIG_STORAGE")
	return *ret0, err
}

// PERMISSIONMULTISIGSTORAGE is a free data retrieval call binding the contract method 0xf1bb7aad.
//
// Solidity: function PERMISSION_MULTISIG_STORAGE() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplSession) PERMISSIONMULTISIGSTORAGE() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.PERMISSIONMULTISIGSTORAGE(&_BitbonStorageImpl.CallOpts)
}

// PERMISSIONMULTISIGSTORAGE is a free data retrieval call binding the contract method 0xf1bb7aad.
//
// Solidity: function PERMISSION_MULTISIG_STORAGE() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCallerSession) PERMISSIONMULTISIGSTORAGE() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.PERMISSIONMULTISIGSTORAGE(&_BitbonStorageImpl.CallOpts)
}

// PERMISSIONROLESSTORAGE is a free data retrieval call binding the contract method 0x12a16032.
//
// Solidity: function PERMISSION_ROLES_STORAGE() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCaller) PERMISSIONROLESSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonStorageImpl.contract.Call(opts, out, "PERMISSION_ROLES_STORAGE")
	return *ret0, err
}

// PERMISSIONROLESSTORAGE is a free data retrieval call binding the contract method 0x12a16032.
//
// Solidity: function PERMISSION_ROLES_STORAGE() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplSession) PERMISSIONROLESSTORAGE() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.PERMISSIONROLESSTORAGE(&_BitbonStorageImpl.CallOpts)
}

// PERMISSIONROLESSTORAGE is a free data retrieval call binding the contract method 0x12a16032.
//
// Solidity: function PERMISSION_ROLES_STORAGE() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCallerSession) PERMISSIONROLESSTORAGE() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.PERMISSIONROLESSTORAGE(&_BitbonStorageImpl.CallOpts)
}

// ROLEACCESSVERIFIER is a free data retrieval call binding the contract method 0xc5405f92.
//
// Solidity: function ROLE_ACCESS_VERIFIER() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCaller) ROLEACCESSVERIFIER(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonStorageImpl.contract.Call(opts, out, "ROLE_ACCESS_VERIFIER")
	return *ret0, err
}

// ROLEACCESSVERIFIER is a free data retrieval call binding the contract method 0xc5405f92.
//
// Solidity: function ROLE_ACCESS_VERIFIER() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplSession) ROLEACCESSVERIFIER() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.ROLEACCESSVERIFIER(&_BitbonStorageImpl.CallOpts)
}

// ROLEACCESSVERIFIER is a free data retrieval call binding the contract method 0xc5405f92.
//
// Solidity: function ROLE_ACCESS_VERIFIER() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCallerSession) ROLEACCESSVERIFIER() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.ROLEACCESSVERIFIER(&_BitbonStorageImpl.CallOpts)
}

// ROLEBITBONISSUEVERIFIER is a free data retrieval call binding the contract method 0xf0a012a6.
//
// Solidity: function ROLE_BITBON_ISSUE_VERIFIER() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCaller) ROLEBITBONISSUEVERIFIER(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonStorageImpl.contract.Call(opts, out, "ROLE_BITBON_ISSUE_VERIFIER")
	return *ret0, err
}

// ROLEBITBONISSUEVERIFIER is a free data retrieval call binding the contract method 0xf0a012a6.
//
// Solidity: function ROLE_BITBON_ISSUE_VERIFIER() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplSession) ROLEBITBONISSUEVERIFIER() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.ROLEBITBONISSUEVERIFIER(&_BitbonStorageImpl.CallOpts)
}

// ROLEBITBONISSUEVERIFIER is a free data retrieval call binding the contract method 0xf0a012a6.
//
// Solidity: function ROLE_BITBON_ISSUE_VERIFIER() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCallerSession) ROLEBITBONISSUEVERIFIER() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.ROLEBITBONISSUEVERIFIER(&_BitbonStorageImpl.CallOpts)
}

// ROLECOMMISSIONVERIFIER is a free data retrieval call binding the contract method 0x8ab3de8c.
//
// Solidity: function ROLE_COMMISSION_VERIFIER() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCaller) ROLECOMMISSIONVERIFIER(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonStorageImpl.contract.Call(opts, out, "ROLE_COMMISSION_VERIFIER")
	return *ret0, err
}

// ROLECOMMISSIONVERIFIER is a free data retrieval call binding the contract method 0x8ab3de8c.
//
// Solidity: function ROLE_COMMISSION_VERIFIER() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplSession) ROLECOMMISSIONVERIFIER() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.ROLECOMMISSIONVERIFIER(&_BitbonStorageImpl.CallOpts)
}

// ROLECOMMISSIONVERIFIER is a free data retrieval call binding the contract method 0x8ab3de8c.
//
// Solidity: function ROLE_COMMISSION_VERIFIER() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCallerSession) ROLECOMMISSIONVERIFIER() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.ROLECOMMISSIONVERIFIER(&_BitbonStorageImpl.CallOpts)
}

// ROLEDEPLOYADMIN is a free data retrieval call binding the contract method 0x667bf6e5.
//
// Solidity: function ROLE_DEPLOY_ADMIN() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCaller) ROLEDEPLOYADMIN(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonStorageImpl.contract.Call(opts, out, "ROLE_DEPLOY_ADMIN")
	return *ret0, err
}

// ROLEDEPLOYADMIN is a free data retrieval call binding the contract method 0x667bf6e5.
//
// Solidity: function ROLE_DEPLOY_ADMIN() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplSession) ROLEDEPLOYADMIN() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.ROLEDEPLOYADMIN(&_BitbonStorageImpl.CallOpts)
}

// ROLEDEPLOYADMIN is a free data retrieval call binding the contract method 0x667bf6e5.
//
// Solidity: function ROLE_DEPLOY_ADMIN() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCallerSession) ROLEDEPLOYADMIN() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.ROLEDEPLOYADMIN(&_BitbonStorageImpl.CallOpts)
}

// ROLEPERMISSIONADMIN is a free data retrieval call binding the contract method 0x1b239996.
//
// Solidity: function ROLE_PERMISSION_ADMIN() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCaller) ROLEPERMISSIONADMIN(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonStorageImpl.contract.Call(opts, out, "ROLE_PERMISSION_ADMIN")
	return *ret0, err
}

// ROLEPERMISSIONADMIN is a free data retrieval call binding the contract method 0x1b239996.
//
// Solidity: function ROLE_PERMISSION_ADMIN() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplSession) ROLEPERMISSIONADMIN() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.ROLEPERMISSIONADMIN(&_BitbonStorageImpl.CallOpts)
}

// ROLEPERMISSIONADMIN is a free data retrieval call binding the contract method 0x1b239996.
//
// Solidity: function ROLE_PERMISSION_ADMIN() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCallerSession) ROLEPERMISSIONADMIN() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.ROLEPERMISSIONADMIN(&_BitbonStorageImpl.CallOpts)
}

// ContractStorage is a free data retrieval call binding the contract method 0x0549b27e.
//
// Solidity: function contractStorage() view returns(address)
func (_BitbonStorageImpl *BitbonStorageImplCaller) ContractStorage(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BitbonStorageImpl.contract.Call(opts, out, "contractStorage")
	return *ret0, err
}

// ContractStorage is a free data retrieval call binding the contract method 0x0549b27e.
//
// Solidity: function contractStorage() view returns(address)
func (_BitbonStorageImpl *BitbonStorageImplSession) ContractStorage() (common.Address, error) {
	return _BitbonStorageImpl.Contract.ContractStorage(&_BitbonStorageImpl.CallOpts)
}

// ContractStorage is a free data retrieval call binding the contract method 0x0549b27e.
//
// Solidity: function contractStorage() view returns(address)
func (_BitbonStorageImpl *BitbonStorageImplCallerSession) ContractStorage() (common.Address, error) {
	return _BitbonStorageImpl.Contract.ContractStorage(&_BitbonStorageImpl.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address assetbox) view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCaller) GetBalance(opts *bind.CallOpts, assetbox common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonStorageImpl.contract.Call(opts, out, "getBalance", assetbox)
	return *ret0, err
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address assetbox) view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplSession) GetBalance(assetbox common.Address) (*big.Int, error) {
	return _BitbonStorageImpl.Contract.GetBalance(&_BitbonStorageImpl.CallOpts, assetbox)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address assetbox) view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCallerSession) GetBalance(assetbox common.Address) (*big.Int, error) {
	return _BitbonStorageImpl.Contract.GetBalance(&_BitbonStorageImpl.CallOpts, assetbox)
}

// GetFullBalance is a free data retrieval call binding the contract method 0xd1eac5cc.
//
// Solidity: function getFullBalance(address assetbox) view returns(uint256, uint256)
func (_BitbonStorageImpl *BitbonStorageImplCaller) GetFullBalance(opts *bind.CallOpts, assetbox common.Address) (*big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _BitbonStorageImpl.contract.Call(opts, out, "getFullBalance", assetbox)
	return *ret0, *ret1, err
}

// GetFullBalance is a free data retrieval call binding the contract method 0xd1eac5cc.
//
// Solidity: function getFullBalance(address assetbox) view returns(uint256, uint256)
func (_BitbonStorageImpl *BitbonStorageImplSession) GetFullBalance(assetbox common.Address) (*big.Int, *big.Int, error) {
	return _BitbonStorageImpl.Contract.GetFullBalance(&_BitbonStorageImpl.CallOpts, assetbox)
}

// GetFullBalance is a free data retrieval call binding the contract method 0xd1eac5cc.
//
// Solidity: function getFullBalance(address assetbox) view returns(uint256, uint256)
func (_BitbonStorageImpl *BitbonStorageImplCallerSession) GetFullBalance(assetbox common.Address) (*big.Int, *big.Int, error) {
	return _BitbonStorageImpl.Contract.GetFullBalance(&_BitbonStorageImpl.CallOpts, assetbox)
}

// GetLockBalance is a free data retrieval call binding the contract method 0x55f78af8.
//
// Solidity: function getLockBalance(address assetbox) view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCaller) GetLockBalance(opts *bind.CallOpts, assetbox common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonStorageImpl.contract.Call(opts, out, "getLockBalance", assetbox)
	return *ret0, err
}

// GetLockBalance is a free data retrieval call binding the contract method 0x55f78af8.
//
// Solidity: function getLockBalance(address assetbox) view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplSession) GetLockBalance(assetbox common.Address) (*big.Int, error) {
	return _BitbonStorageImpl.Contract.GetLockBalance(&_BitbonStorageImpl.CallOpts, assetbox)
}

// GetLockBalance is a free data retrieval call binding the contract method 0x55f78af8.
//
// Solidity: function getLockBalance(address assetbox) view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCallerSession) GetLockBalance(assetbox common.Address) (*big.Int, error) {
	return _BitbonStorageImpl.Contract.GetLockBalance(&_BitbonStorageImpl.CallOpts, assetbox)
}

// GetThisContractIndex is a free data retrieval call binding the contract method 0x70c6d367.
//
// Solidity: function getThisContractIndex() pure returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCaller) GetThisContractIndex(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonStorageImpl.contract.Call(opts, out, "getThisContractIndex")
	return *ret0, err
}

// GetThisContractIndex is a free data retrieval call binding the contract method 0x70c6d367.
//
// Solidity: function getThisContractIndex() pure returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplSession) GetThisContractIndex() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.GetThisContractIndex(&_BitbonStorageImpl.CallOpts)
}

// GetThisContractIndex is a free data retrieval call binding the contract method 0x70c6d367.
//
// Solidity: function getThisContractIndex() pure returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCallerSession) GetThisContractIndex() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.GetThisContractIndex(&_BitbonStorageImpl.CallOpts)
}

// GetTotalSupply is a free data retrieval call binding the contract method 0xc4e41b22.
//
// Solidity: function getTotalSupply() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCaller) GetTotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonStorageImpl.contract.Call(opts, out, "getTotalSupply")
	return *ret0, err
}

// GetTotalSupply is a free data retrieval call binding the contract method 0xc4e41b22.
//
// Solidity: function getTotalSupply() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplSession) GetTotalSupply() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.GetTotalSupply(&_BitbonStorageImpl.CallOpts)
}

// GetTotalSupply is a free data retrieval call binding the contract method 0xc4e41b22.
//
// Solidity: function getTotalSupply() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCallerSession) GetTotalSupply() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.GetTotalSupply(&_BitbonStorageImpl.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonStorageImpl.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplSession) TotalSupply() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.TotalSupply(&_BitbonStorageImpl.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BitbonStorageImpl *BitbonStorageImplCallerSession) TotalSupply() (*big.Int, error) {
	return _BitbonStorageImpl.Contract.TotalSupply(&_BitbonStorageImpl.CallOpts)
}

// ChargeFee is a paid mutator transaction binding the contract method 0xd3beb298.
//
// Solidity: function chargeFee(address from, uint256 amountFee) returns(bool)
func (_BitbonStorageImpl *BitbonStorageImplTransactor) ChargeFee(opts *bind.TransactOpts, from common.Address, amountFee *big.Int) (*types.Transaction, error) {
	return _BitbonStorageImpl.contract.Transact(opts, "chargeFee", from, amountFee)
}

// ChargeFee is a paid mutator transaction binding the contract method 0xd3beb298.
//
// Solidity: function chargeFee(address from, uint256 amountFee) returns(bool)
func (_BitbonStorageImpl *BitbonStorageImplSession) ChargeFee(from common.Address, amountFee *big.Int) (*types.Transaction, error) {
	return _BitbonStorageImpl.Contract.ChargeFee(&_BitbonStorageImpl.TransactOpts, from, amountFee)
}

// ChargeFee is a paid mutator transaction binding the contract method 0xd3beb298.
//
// Solidity: function chargeFee(address from, uint256 amountFee) returns(bool)
func (_BitbonStorageImpl *BitbonStorageImplTransactorSession) ChargeFee(from common.Address, amountFee *big.Int) (*types.Transaction, error) {
	return _BitbonStorageImpl.Contract.ChargeFee(&_BitbonStorageImpl.TransactOpts, from, amountFee)
}

// Lock is a paid mutator transaction binding the contract method 0x282d3fdf.
//
// Solidity: function lock(address assetbox, uint256 amount) returns(bool)
func (_BitbonStorageImpl *BitbonStorageImplTransactor) Lock(opts *bind.TransactOpts, assetbox common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BitbonStorageImpl.contract.Transact(opts, "lock", assetbox, amount)
}

// Lock is a paid mutator transaction binding the contract method 0x282d3fdf.
//
// Solidity: function lock(address assetbox, uint256 amount) returns(bool)
func (_BitbonStorageImpl *BitbonStorageImplSession) Lock(assetbox common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BitbonStorageImpl.Contract.Lock(&_BitbonStorageImpl.TransactOpts, assetbox, amount)
}

// Lock is a paid mutator transaction binding the contract method 0x282d3fdf.
//
// Solidity: function lock(address assetbox, uint256 amount) returns(bool)
func (_BitbonStorageImpl *BitbonStorageImplTransactorSession) Lock(assetbox common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BitbonStorageImpl.Contract.Lock(&_BitbonStorageImpl.TransactOpts, assetbox, amount)
}

// SetAssetboxBalance is a paid mutator transaction binding the contract method 0xdc7254d3.
//
// Solidity: function setAssetboxBalance(address assetbox, uint256 amount) returns()
func (_BitbonStorageImpl *BitbonStorageImplTransactor) SetAssetboxBalance(opts *bind.TransactOpts, assetbox common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BitbonStorageImpl.contract.Transact(opts, "setAssetboxBalance", assetbox, amount)
}

// SetAssetboxBalance is a paid mutator transaction binding the contract method 0xdc7254d3.
//
// Solidity: function setAssetboxBalance(address assetbox, uint256 amount) returns()
func (_BitbonStorageImpl *BitbonStorageImplSession) SetAssetboxBalance(assetbox common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BitbonStorageImpl.Contract.SetAssetboxBalance(&_BitbonStorageImpl.TransactOpts, assetbox, amount)
}

// SetAssetboxBalance is a paid mutator transaction binding the contract method 0xdc7254d3.
//
// Solidity: function setAssetboxBalance(address assetbox, uint256 amount) returns()
func (_BitbonStorageImpl *BitbonStorageImplTransactorSession) SetAssetboxBalance(assetbox common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BitbonStorageImpl.Contract.SetAssetboxBalance(&_BitbonStorageImpl.TransactOpts, assetbox, amount)
}

// SetCapitaliseFund is a paid mutator transaction binding the contract method 0xb8355814.
//
// Solidity: function setCapitaliseFund(address assetbox) returns()
func (_BitbonStorageImpl *BitbonStorageImplTransactor) SetCapitaliseFund(opts *bind.TransactOpts, assetbox common.Address) (*types.Transaction, error) {
	return _BitbonStorageImpl.contract.Transact(opts, "setCapitaliseFund", assetbox)
}

// SetCapitaliseFund is a paid mutator transaction binding the contract method 0xb8355814.
//
// Solidity: function setCapitaliseFund(address assetbox) returns()
func (_BitbonStorageImpl *BitbonStorageImplSession) SetCapitaliseFund(assetbox common.Address) (*types.Transaction, error) {
	return _BitbonStorageImpl.Contract.SetCapitaliseFund(&_BitbonStorageImpl.TransactOpts, assetbox)
}

// SetCapitaliseFund is a paid mutator transaction binding the contract method 0xb8355814.
//
// Solidity: function setCapitaliseFund(address assetbox) returns()
func (_BitbonStorageImpl *BitbonStorageImplTransactorSession) SetCapitaliseFund(assetbox common.Address) (*types.Transaction, error) {
	return _BitbonStorageImpl.Contract.SetCapitaliseFund(&_BitbonStorageImpl.TransactOpts, assetbox)
}

// SetMigrated is a paid mutator transaction binding the contract method 0x925382c0.
//
// Solidity: function setMigrated() returns()
func (_BitbonStorageImpl *BitbonStorageImplTransactor) SetMigrated(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BitbonStorageImpl.contract.Transact(opts, "setMigrated")
}

// SetMigrated is a paid mutator transaction binding the contract method 0x925382c0.
//
// Solidity: function setMigrated() returns()
func (_BitbonStorageImpl *BitbonStorageImplSession) SetMigrated() (*types.Transaction, error) {
	return _BitbonStorageImpl.Contract.SetMigrated(&_BitbonStorageImpl.TransactOpts)
}

// SetMigrated is a paid mutator transaction binding the contract method 0x925382c0.
//
// Solidity: function setMigrated() returns()
func (_BitbonStorageImpl *BitbonStorageImplTransactorSession) SetMigrated() (*types.Transaction, error) {
	return _BitbonStorageImpl.Contract.SetMigrated(&_BitbonStorageImpl.TransactOpts)
}

// SetMiningFund is a paid mutator transaction binding the contract method 0x6d089128.
//
// Solidity: function setMiningFund(address assetbox) returns()
func (_BitbonStorageImpl *BitbonStorageImplTransactor) SetMiningFund(opts *bind.TransactOpts, assetbox common.Address) (*types.Transaction, error) {
	return _BitbonStorageImpl.contract.Transact(opts, "setMiningFund", assetbox)
}

// SetMiningFund is a paid mutator transaction binding the contract method 0x6d089128.
//
// Solidity: function setMiningFund(address assetbox) returns()
func (_BitbonStorageImpl *BitbonStorageImplSession) SetMiningFund(assetbox common.Address) (*types.Transaction, error) {
	return _BitbonStorageImpl.Contract.SetMiningFund(&_BitbonStorageImpl.TransactOpts, assetbox)
}

// SetMiningFund is a paid mutator transaction binding the contract method 0x6d089128.
//
// Solidity: function setMiningFund(address assetbox) returns()
func (_BitbonStorageImpl *BitbonStorageImplTransactorSession) SetMiningFund(assetbox common.Address) (*types.Transaction, error) {
	return _BitbonStorageImpl.Contract.SetMiningFund(&_BitbonStorageImpl.TransactOpts, assetbox)
}

// SetSimcordAssetbox is a paid mutator transaction binding the contract method 0x81648555.
//
// Solidity: function setSimcordAssetbox(address assetbox) returns()
func (_BitbonStorageImpl *BitbonStorageImplTransactor) SetSimcordAssetbox(opts *bind.TransactOpts, assetbox common.Address) (*types.Transaction, error) {
	return _BitbonStorageImpl.contract.Transact(opts, "setSimcordAssetbox", assetbox)
}

// SetSimcordAssetbox is a paid mutator transaction binding the contract method 0x81648555.
//
// Solidity: function setSimcordAssetbox(address assetbox) returns()
func (_BitbonStorageImpl *BitbonStorageImplSession) SetSimcordAssetbox(assetbox common.Address) (*types.Transaction, error) {
	return _BitbonStorageImpl.Contract.SetSimcordAssetbox(&_BitbonStorageImpl.TransactOpts, assetbox)
}

// SetSimcordAssetbox is a paid mutator transaction binding the contract method 0x81648555.
//
// Solidity: function setSimcordAssetbox(address assetbox) returns()
func (_BitbonStorageImpl *BitbonStorageImplTransactorSession) SetSimcordAssetbox(assetbox common.Address) (*types.Transaction, error) {
	return _BitbonStorageImpl.Contract.SetSimcordAssetbox(&_BitbonStorageImpl.TransactOpts, assetbox)
}

// Transfer is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(address from, address to, uint256 amount) returns(bool)
func (_BitbonStorageImpl *BitbonStorageImplTransactor) Transfer(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BitbonStorageImpl.contract.Transact(opts, "transfer", from, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(address from, address to, uint256 amount) returns(bool)
func (_BitbonStorageImpl *BitbonStorageImplSession) Transfer(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BitbonStorageImpl.Contract.Transfer(&_BitbonStorageImpl.TransactOpts, from, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(address from, address to, uint256 amount) returns(bool)
func (_BitbonStorageImpl *BitbonStorageImplTransactorSession) Transfer(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BitbonStorageImpl.Contract.Transfer(&_BitbonStorageImpl.TransactOpts, from, to, amount)
}

// Unlock is a paid mutator transaction binding the contract method 0x7eee288d.
//
// Solidity: function unlock(address assetbox, uint256 amount) returns(bool)
func (_BitbonStorageImpl *BitbonStorageImplTransactor) Unlock(opts *bind.TransactOpts, assetbox common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BitbonStorageImpl.contract.Transact(opts, "unlock", assetbox, amount)
}

// Unlock is a paid mutator transaction binding the contract method 0x7eee288d.
//
// Solidity: function unlock(address assetbox, uint256 amount) returns(bool)
func (_BitbonStorageImpl *BitbonStorageImplSession) Unlock(assetbox common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BitbonStorageImpl.Contract.Unlock(&_BitbonStorageImpl.TransactOpts, assetbox, amount)
}

// Unlock is a paid mutator transaction binding the contract method 0x7eee288d.
//
// Solidity: function unlock(address assetbox, uint256 amount) returns(bool)
func (_BitbonStorageImpl *BitbonStorageImplTransactorSession) Unlock(assetbox common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BitbonStorageImpl.Contract.Unlock(&_BitbonStorageImpl.TransactOpts, assetbox, amount)
}

// UnlockAndTransfer is a paid mutator transaction binding the contract method 0xf074b3cc.
//
// Solidity: function unlockAndTransfer(address from, address to, uint256 amount) returns(bool)
func (_BitbonStorageImpl *BitbonStorageImplTransactor) UnlockAndTransfer(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BitbonStorageImpl.contract.Transact(opts, "unlockAndTransfer", from, to, amount)
}

// UnlockAndTransfer is a paid mutator transaction binding the contract method 0xf074b3cc.
//
// Solidity: function unlockAndTransfer(address from, address to, uint256 amount) returns(bool)
func (_BitbonStorageImpl *BitbonStorageImplSession) UnlockAndTransfer(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BitbonStorageImpl.Contract.UnlockAndTransfer(&_BitbonStorageImpl.TransactOpts, from, to, amount)
}

// UnlockAndTransfer is a paid mutator transaction binding the contract method 0xf074b3cc.
//
// Solidity: function unlockAndTransfer(address from, address to, uint256 amount) returns(bool)
func (_BitbonStorageImpl *BitbonStorageImplTransactorSession) UnlockAndTransfer(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BitbonStorageImpl.Contract.UnlockAndTransfer(&_BitbonStorageImpl.TransactOpts, from, to, amount)
}

// BitbonStorageImplFeeChargedIterator is returned from FilterFeeCharged and is used to iterate over the raw logs and unpacked data for FeeCharged events raised by the BitbonStorageImpl contract.
type BitbonStorageImplFeeChargedIterator struct {
	Event *BitbonStorageImplFeeCharged // Event containing the contract specifics and raw log

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
func (it *BitbonStorageImplFeeChargedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BitbonStorageImplFeeCharged)
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
		it.Event = new(BitbonStorageImplFeeCharged)
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
func (it *BitbonStorageImplFeeChargedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BitbonStorageImplFeeChargedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BitbonStorageImplFeeCharged represents a FeeCharged event raised by the BitbonStorageImpl contract.
type BitbonStorageImplFeeCharged struct {
	Assetbox   common.Address
	Capitalise *big.Int
	Value      *big.Int
	Mining     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterFeeCharged is a free log retrieval operation binding the contract event 0xf95c12eeb9ffa470bcf7f32b50e2998d150717d9661a85a7b44818e9abae4972.
//
// Solidity: event FeeCharged(address indexed assetbox, uint256 capitalise, uint256 value, uint256 mining)
func (_BitbonStorageImpl *BitbonStorageImplFilterer) FilterFeeCharged(opts *bind.FilterOpts, assetbox []common.Address) (*BitbonStorageImplFeeChargedIterator, error) {

	var assetboxRule []interface{}
	for _, assetboxItem := range assetbox {
		assetboxRule = append(assetboxRule, assetboxItem)
	}

	logs, sub, err := _BitbonStorageImpl.contract.FilterLogs(opts, "FeeCharged", assetboxRule)
	if err != nil {
		return nil, err
	}
	return &BitbonStorageImplFeeChargedIterator{contract: _BitbonStorageImpl.contract, event: "FeeCharged", logs: logs, sub: sub}, nil
}

// WatchFeeCharged is a free log subscription operation binding the contract event 0xf95c12eeb9ffa470bcf7f32b50e2998d150717d9661a85a7b44818e9abae4972.
//
// Solidity: event FeeCharged(address indexed assetbox, uint256 capitalise, uint256 value, uint256 mining)
func (_BitbonStorageImpl *BitbonStorageImplFilterer) WatchFeeCharged(opts *bind.WatchOpts, sink chan<- *BitbonStorageImplFeeCharged, assetbox []common.Address) (event.Subscription, error) {

	var assetboxRule []interface{}
	for _, assetboxItem := range assetbox {
		assetboxRule = append(assetboxRule, assetboxItem)
	}

	logs, sub, err := _BitbonStorageImpl.contract.WatchLogs(opts, "FeeCharged", assetboxRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BitbonStorageImplFeeCharged)
				if err := _BitbonStorageImpl.contract.UnpackLog(event, "FeeCharged", log); err != nil {
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

// ParseFeeCharged is a log parse operation binding the contract event 0xf95c12eeb9ffa470bcf7f32b50e2998d150717d9661a85a7b44818e9abae4972.
//
// Solidity: event FeeCharged(address indexed assetbox, uint256 capitalise, uint256 value, uint256 mining)
func (_BitbonStorageImpl *BitbonStorageImplFilterer) ParseFeeCharged(log types.Log) (*BitbonStorageImplFeeCharged, error) {
	event := new(BitbonStorageImplFeeCharged)
	if err := _BitbonStorageImpl.contract.UnpackLog(event, "FeeCharged", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BitbonStorageImplLockIterator is returned from FilterLock and is used to iterate over the raw logs and unpacked data for Lock events raised by the BitbonStorageImpl contract.
type BitbonStorageImplLockIterator struct {
	Event *BitbonStorageImplLock // Event containing the contract specifics and raw log

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
func (it *BitbonStorageImplLockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BitbonStorageImplLock)
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
		it.Event = new(BitbonStorageImplLock)
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
func (it *BitbonStorageImplLockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BitbonStorageImplLockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BitbonStorageImplLock represents a Lock event raised by the BitbonStorageImpl contract.
type BitbonStorageImplLock struct {
	Assetbox common.Address
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLock is a free log retrieval operation binding the contract event 0x625fed9875dada8643f2418b838ae0bc78d9a148a18eee4ee1979ff0f3f5d427.
//
// Solidity: event Lock(address indexed assetbox, uint256 value)
func (_BitbonStorageImpl *BitbonStorageImplFilterer) FilterLock(opts *bind.FilterOpts, assetbox []common.Address) (*BitbonStorageImplLockIterator, error) {

	var assetboxRule []interface{}
	for _, assetboxItem := range assetbox {
		assetboxRule = append(assetboxRule, assetboxItem)
	}

	logs, sub, err := _BitbonStorageImpl.contract.FilterLogs(opts, "Lock", assetboxRule)
	if err != nil {
		return nil, err
	}
	return &BitbonStorageImplLockIterator{contract: _BitbonStorageImpl.contract, event: "Lock", logs: logs, sub: sub}, nil
}

// WatchLock is a free log subscription operation binding the contract event 0x625fed9875dada8643f2418b838ae0bc78d9a148a18eee4ee1979ff0f3f5d427.
//
// Solidity: event Lock(address indexed assetbox, uint256 value)
func (_BitbonStorageImpl *BitbonStorageImplFilterer) WatchLock(opts *bind.WatchOpts, sink chan<- *BitbonStorageImplLock, assetbox []common.Address) (event.Subscription, error) {

	var assetboxRule []interface{}
	for _, assetboxItem := range assetbox {
		assetboxRule = append(assetboxRule, assetboxItem)
	}

	logs, sub, err := _BitbonStorageImpl.contract.WatchLogs(opts, "Lock", assetboxRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BitbonStorageImplLock)
				if err := _BitbonStorageImpl.contract.UnpackLog(event, "Lock", log); err != nil {
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

// ParseLock is a log parse operation binding the contract event 0x625fed9875dada8643f2418b838ae0bc78d9a148a18eee4ee1979ff0f3f5d427.
//
// Solidity: event Lock(address indexed assetbox, uint256 value)
func (_BitbonStorageImpl *BitbonStorageImplFilterer) ParseLock(log types.Log) (*BitbonStorageImplLock, error) {
	event := new(BitbonStorageImplLock)
	if err := _BitbonStorageImpl.contract.UnpackLog(event, "Lock", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BitbonStorageImplMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the BitbonStorageImpl contract.
type BitbonStorageImplMintIterator struct {
	Event *BitbonStorageImplMint // Event containing the contract specifics and raw log

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
func (it *BitbonStorageImplMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BitbonStorageImplMint)
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
		it.Event = new(BitbonStorageImplMint)
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
func (it *BitbonStorageImplMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BitbonStorageImplMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BitbonStorageImplMint represents a Mint event raised by the BitbonStorageImpl contract.
type BitbonStorageImplMint struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed to, uint256 amount)
func (_BitbonStorageImpl *BitbonStorageImplFilterer) FilterMint(opts *bind.FilterOpts, to []common.Address) (*BitbonStorageImplMintIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BitbonStorageImpl.contract.FilterLogs(opts, "Mint", toRule)
	if err != nil {
		return nil, err
	}
	return &BitbonStorageImplMintIterator{contract: _BitbonStorageImpl.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed to, uint256 amount)
func (_BitbonStorageImpl *BitbonStorageImplFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *BitbonStorageImplMint, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BitbonStorageImpl.contract.WatchLogs(opts, "Mint", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BitbonStorageImplMint)
				if err := _BitbonStorageImpl.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed to, uint256 amount)
func (_BitbonStorageImpl *BitbonStorageImplFilterer) ParseMint(log types.Log) (*BitbonStorageImplMint, error) {
	event := new(BitbonStorageImplMint)
	if err := _BitbonStorageImpl.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BitbonStorageImplTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the BitbonStorageImpl contract.
type BitbonStorageImplTransferIterator struct {
	Event *BitbonStorageImplTransfer // Event containing the contract specifics and raw log

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
func (it *BitbonStorageImplTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BitbonStorageImplTransfer)
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
		it.Event = new(BitbonStorageImplTransfer)
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
func (it *BitbonStorageImplTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BitbonStorageImplTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BitbonStorageImplTransfer represents a Transfer event raised by the BitbonStorageImpl contract.
type BitbonStorageImplTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_BitbonStorageImpl *BitbonStorageImplFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BitbonStorageImplTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BitbonStorageImpl.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BitbonStorageImplTransferIterator{contract: _BitbonStorageImpl.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_BitbonStorageImpl *BitbonStorageImplFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BitbonStorageImplTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BitbonStorageImpl.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BitbonStorageImplTransfer)
				if err := _BitbonStorageImpl.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_BitbonStorageImpl *BitbonStorageImplFilterer) ParseTransfer(log types.Log) (*BitbonStorageImplTransfer, error) {
	event := new(BitbonStorageImplTransfer)
	if err := _BitbonStorageImpl.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BitbonStorageImplUnLockIterator is returned from FilterUnLock and is used to iterate over the raw logs and unpacked data for UnLock events raised by the BitbonStorageImpl contract.
type BitbonStorageImplUnLockIterator struct {
	Event *BitbonStorageImplUnLock // Event containing the contract specifics and raw log

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
func (it *BitbonStorageImplUnLockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BitbonStorageImplUnLock)
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
		it.Event = new(BitbonStorageImplUnLock)
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
func (it *BitbonStorageImplUnLockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BitbonStorageImplUnLockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BitbonStorageImplUnLock represents a UnLock event raised by the BitbonStorageImpl contract.
type BitbonStorageImplUnLock struct {
	Assetbox common.Address
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUnLock is a free log retrieval operation binding the contract event 0xb371d42b3715509a27f3109f6ac1ef6b7d7e7f8e9232b738ed17338be6cf9580.
//
// Solidity: event UnLock(address indexed assetbox, uint256 value)
func (_BitbonStorageImpl *BitbonStorageImplFilterer) FilterUnLock(opts *bind.FilterOpts, assetbox []common.Address) (*BitbonStorageImplUnLockIterator, error) {

	var assetboxRule []interface{}
	for _, assetboxItem := range assetbox {
		assetboxRule = append(assetboxRule, assetboxItem)
	}

	logs, sub, err := _BitbonStorageImpl.contract.FilterLogs(opts, "UnLock", assetboxRule)
	if err != nil {
		return nil, err
	}
	return &BitbonStorageImplUnLockIterator{contract: _BitbonStorageImpl.contract, event: "UnLock", logs: logs, sub: sub}, nil
}

// WatchUnLock is a free log subscription operation binding the contract event 0xb371d42b3715509a27f3109f6ac1ef6b7d7e7f8e9232b738ed17338be6cf9580.
//
// Solidity: event UnLock(address indexed assetbox, uint256 value)
func (_BitbonStorageImpl *BitbonStorageImplFilterer) WatchUnLock(opts *bind.WatchOpts, sink chan<- *BitbonStorageImplUnLock, assetbox []common.Address) (event.Subscription, error) {

	var assetboxRule []interface{}
	for _, assetboxItem := range assetbox {
		assetboxRule = append(assetboxRule, assetboxItem)
	}

	logs, sub, err := _BitbonStorageImpl.contract.WatchLogs(opts, "UnLock", assetboxRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BitbonStorageImplUnLock)
				if err := _BitbonStorageImpl.contract.UnpackLog(event, "UnLock", log); err != nil {
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

// ParseUnLock is a log parse operation binding the contract event 0xb371d42b3715509a27f3109f6ac1ef6b7d7e7f8e9232b738ed17338be6cf9580.
//
// Solidity: event UnLock(address indexed assetbox, uint256 value)
func (_BitbonStorageImpl *BitbonStorageImplFilterer) ParseUnLock(log types.Log) (*BitbonStorageImplUnLock, error) {
	event := new(BitbonStorageImplUnLock)
	if err := _BitbonStorageImpl.contract.UnpackLog(event, "UnLock", log); err != nil {
		return nil, err
	}
	return event, nil
}
