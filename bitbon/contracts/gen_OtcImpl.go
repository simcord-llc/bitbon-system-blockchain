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

// OtcImplABI is the input ABI used to generate the binding from.
const OtcImplABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractStorageAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"assetbox\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"balanceAviable\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"balanceLocked\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetboxType\",\"type\":\"uint256\"}],\"name\":\"BalanceLocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"assetbox\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"balanceAviable\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"balanceLocked\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetboxType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BalanceUnLocked\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ACCESS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ACCESS_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ADMINABLE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ADMIN_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ASSETBOX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ASSETBOX_INFO\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ASSETBOX_INFO_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ASSETBOX_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_BALANCE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_BITBON\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_BITBON_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_BITBON_SUPPORT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_BONBALANCE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_BONTRANSFER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_DEX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_DISTRIBUTION_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_EXCHANGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_EXCHANGE_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_FEE_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_GENERATOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MINING_AGENT_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MSBON\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MSBON_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_ADD_ADDRESS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_ADD_ADMIN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_ADD_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_ADD_ROLE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_DROP_PERMISSION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_EDIT_ADMIN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_EDIT_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_EDIT_ROLE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REMOVE_ADMIN_BY_ADDRESS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REMOVE_ADMIN_BY_INDEX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REMOVE_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REMOVE_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REMOVE_ROLE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REPLENISH_CAPITALIZATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REPLENISH_CERTIFICATE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REPLENISH_EMISSION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_SET_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_SET_PERMISSION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_OTC\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_OTC_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_REPLENISH_CAPITALIZATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_RESERVED_ALIAS_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ROLE_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_SAFE_TRANSFER_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_TRANSFER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_TRANSFER_FROM_CAPITALIZATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_ACCESS_RESTORATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_ADMIN_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_CONTRACTS_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_CONTRACT_DEPLOY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_EMISSION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_FEES_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_MULTISIG_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_ROLES_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_ACCESS_VERIFIER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_BITBON_ISSUE_VERIFIER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_COMMISSION_VERIFIER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_DEPLOY_ADMIN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_PERMISSION_ADMIN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"contractStorage\",\"outputs\":[{\"internalType\":\"contractContractStorage\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"state\",\"outputs\":[{\"internalType\":\"enumDtoDex.ExchangeState\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getThisContractIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"enumDtoDex.ExchangeState\",\"name\":\"newState\",\"type\":\"uint8\"}],\"name\":\"setState\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"offerAssets\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"wantAssets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"offerAmount\",\"type\":\"uint256[]\"},{\"internalType\":\"bool[]\",\"name\":\"tradeType\",\"type\":\"bool[]\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"name\":\"makeBatchOrders\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"offerAssets\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"wantAssets\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"offerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"name\":\"makeBitbonOrder\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"offerAssets\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"wantAssets\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"offerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"name\":\"makeMsbonOrder\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"wantAmount\",\"type\":\"uint256\"}],\"name\":\"setBitbonPrice\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"wantAmount\",\"type\":\"uint256\"}],\"name\":\"setMsbonPrice\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"fillAssets\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"wantAssets\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"offerHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amountToTake\",\"type\":\"uint256\"}],\"name\":\"fillBitbonOrder\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"fillAssets\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"wantAssets\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"offerHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amountToTake\",\"type\":\"uint256\"}],\"name\":\"fillMsbonOrder\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"offerHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"expectedAvailableAmount\",\"type\":\"uint256\"}],\"name\":\"cancelBitbonOrder\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"offerHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"expectedAvailableAmount\",\"type\":\"uint256\"}],\"name\":\"cancelMsbonOrder\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// OtcImpl is an auto generated Go binding around an Ethereum contract.
type OtcImpl struct {
	OtcImplCaller     // Read-only binding to the contract
	OtcImplTransactor // Write-only binding to the contract
	OtcImplFilterer   // Log filterer for contract events
}

// OtcImplCaller is an auto generated read-only Go binding around an Ethereum contract.
type OtcImplCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OtcImplTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OtcImplTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OtcImplFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OtcImplFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OtcImplSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OtcImplSession struct {
	Contract     *OtcImpl          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OtcImplCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OtcImplCallerSession struct {
	Contract *OtcImplCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// OtcImplTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OtcImplTransactorSession struct {
	Contract     *OtcImplTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OtcImplRaw is an auto generated low-level Go binding around an Ethereum contract.
type OtcImplRaw struct {
	Contract *OtcImpl // Generic contract binding to access the raw methods on
}

// OtcImplCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OtcImplCallerRaw struct {
	Contract *OtcImplCaller // Generic read-only contract binding to access the raw methods on
}

// OtcImplTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OtcImplTransactorRaw struct {
	Contract *OtcImplTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOtcImpl creates a new instance of OtcImpl, bound to a specific deployed contract.
func NewOtcImpl(address common.Address, backend bind.ContractBackend) (*OtcImpl, error) {
	contract, err := bindOtcImpl(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OtcImpl{OtcImplCaller: OtcImplCaller{contract: contract}, OtcImplTransactor: OtcImplTransactor{contract: contract}, OtcImplFilterer: OtcImplFilterer{contract: contract}}, nil
}

// NewOtcImplCaller creates a new read-only instance of OtcImpl, bound to a specific deployed contract.
func NewOtcImplCaller(address common.Address, caller bind.ContractCaller) (*OtcImplCaller, error) {
	contract, err := bindOtcImpl(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OtcImplCaller{contract: contract}, nil
}

// NewOtcImplTransactor creates a new write-only instance of OtcImpl, bound to a specific deployed contract.
func NewOtcImplTransactor(address common.Address, transactor bind.ContractTransactor) (*OtcImplTransactor, error) {
	contract, err := bindOtcImpl(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OtcImplTransactor{contract: contract}, nil
}

// NewOtcImplFilterer creates a new log filterer instance of OtcImpl, bound to a specific deployed contract.
func NewOtcImplFilterer(address common.Address, filterer bind.ContractFilterer) (*OtcImplFilterer, error) {
	contract, err := bindOtcImpl(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OtcImplFilterer{contract: contract}, nil
}

// bindOtcImpl binds a generic wrapper to an already deployed contract.
func bindOtcImpl(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OtcImplABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OtcImpl *OtcImplRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _OtcImpl.Contract.OtcImplCaller.contract.Call(opts, result, method, params...)
}

// Return Abi for OtcImplCaller
func (_OtcImpl *OtcImplCaller) ABI() abi.ABI {
	return _OtcImpl.contract.Abi
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OtcImpl *OtcImplRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OtcImpl.Contract.OtcImplTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OtcImpl *OtcImplRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OtcImpl.Contract.OtcImplTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OtcImpl *OtcImplCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _OtcImpl.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OtcImpl *OtcImplTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OtcImpl.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OtcImpl *OtcImplTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OtcImpl.Contract.contract.Transact(opts, method, params...)
}

// CONTRACTACCESS is a free data retrieval call binding the contract method 0x0ac62b0c.
//
// Solidity: function CONTRACT_ACCESS() view returns(uint256)
func (_OtcImpl *OtcImplCaller) CONTRACTACCESS(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OtcImpl.contract.Call(opts, out, "CONTRACT_ACCESS")
	return *ret0, err
}

// CONTRACTACCESS is a free data retrieval call binding the contract method 0x0ac62b0c.
//
// Solidity: function CONTRACT_ACCESS() view returns(uint256)
func (_OtcImpl *OtcImplSession) CONTRACTACCESS() (*big.Int, error) {
	return _OtcImpl.Contract.CONTRACTACCESS(&_OtcImpl.CallOpts)
}

// CONTRACTACCESS is a free data retrieval call binding the contract method 0x0ac62b0c.
//
// Solidity: function CONTRACT_ACCESS() view returns(uint256)
func (_OtcImpl *OtcImplCallerSession) CONTRACTACCESS() (*big.Int, error) {
	return _OtcImpl.Contract.CONTRACTACCESS(&_OtcImpl.CallOpts)
}

// CONTRACTACCESSSTORAGE is a free data retrieval call binding the contract method 0xceac9bd5.
//
// Solidity: function CONTRACT_ACCESS_STORAGE() view returns(uint256)
func (_OtcImpl *OtcImplCaller) CONTRACTACCESSSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OtcImpl.contract.Call(opts, out, "CONTRACT_ACCESS_STORAGE")
	return *ret0, err
}

// CONTRACTACCESSSTORAGE is a free data retrieval call binding the contract method 0xceac9bd5.
//
// Solidity: function CONTRACT_ACCESS_STORAGE() view returns(uint256)
func (_OtcImpl *OtcImplSession) CONTRACTACCESSSTORAGE() (*big.Int, error) {
	return _OtcImpl.Contract.CONTRACTACCESSSTORAGE(&_OtcImpl.CallOpts)
}

// CONTRACTACCESSSTORAGE is a free data retrieval call binding the contract method 0xceac9bd5.
//
// Solidity: function CONTRACT_ACCESS_STORAGE() view returns(uint256)
func (_OtcImpl *OtcImplCallerSession) CONTRACTACCESSSTORAGE() (*big.Int, error) {
	return _OtcImpl.Contract.CONTRACTACCESSSTORAGE(&_OtcImpl.CallOpts)
}

// CONTRACTADMINABLE is a free data retrieval call binding the contract method 0xbbe3a78f.
//
// Solidity: function CONTRACT_ADMINABLE() view returns(uint256)
func (_OtcImpl *OtcImplCaller) CONTRACTADMINABLE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OtcImpl.contract.Call(opts, out, "CONTRACT_ADMINABLE")
	return *ret0, err
}

// CONTRACTADMINABLE is a free data retrieval call binding the contract method 0xbbe3a78f.
//
// Solidity: function CONTRACT_ADMINABLE() view returns(uint256)
func (_OtcImpl *OtcImplSession) CONTRACTADMINABLE() (*big.Int, error) {
	return _OtcImpl.Contract.CONTRACTADMINABLE(&_OtcImpl.CallOpts)
}

// CONTRACTADMINABLE is a free data retrieval call binding the contract method 0xbbe3a78f.
//
// Solidity: function CONTRACT_ADMINABLE() view returns(uint256)
func (_OtcImpl *OtcImplCallerSession) CONTRACTADMINABLE() (*big.Int, error) {
	return _OtcImpl.Contract.CONTRACTADMINABLE(&_OtcImpl.CallOpts)
}

// CONTRACTADMINSTORAGE is a free data retrieval call binding the contract method 0xc9aebae9.
//
// Solidity: function CONTRACT_ADMIN_STORAGE() view returns(uint256)
func (_OtcImpl *OtcImplCaller) CONTRACTADMINSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OtcImpl.contract.Call(opts, out, "CONTRACT_ADMIN_STORAGE")
	return *ret0, err
}

// CONTRACTADMINSTORAGE is a free data retrieval call binding the contract method 0xc9aebae9.
//
// Solidity: function CONTRACT_ADMIN_STORAGE() view returns(uint256)
func (_OtcImpl *OtcImplSession) CONTRACTADMINSTORAGE() (*big.Int, error) {
	return _OtcImpl.Contract.CONTRACTADMINSTORAGE(&_OtcImpl.CallOpts)
}

// CONTRACTADMINSTORAGE is a free data retrieval call binding the contract method 0xc9aebae9.
//
// Solidity: function CONTRACT_ADMIN_STORAGE() view returns(uint256)
func (_OtcImpl *OtcImplCallerSession) CONTRACTADMINSTORAGE() (*big.Int, error) {
	return _OtcImpl.Contract.CONTRACTADMINSTORAGE(&_OtcImpl.CallOpts)
}

// CONTRACTASSETBOX is a free data retrieval call binding the contract method 0x161aa17c.
//
// Solidity: function CONTRACT_ASSETBOX() view returns(uint256)
func (_OtcImpl *OtcImplCaller) CONTRACTASSETBOX(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OtcImpl.contract.Call(opts, out, "CONTRACT_ASSETBOX")
	return *ret0, err
}

// CONTRACTASSETBOX is a free data retrieval call binding the contract method 0x161aa17c.
//
// Solidity: function CONTRACT_ASSETBOX() view returns(uint256)
func (_OtcImpl *OtcImplSession) CONTRACTASSETBOX() (*big.Int, error) {
	return _OtcImpl.Contract.CONTRACTASSETBOX(&_OtcImpl.CallOpts)
}

// CONTRACTASSETBOX is a free data retrieval call binding the contract method 0x161aa17c.
//
// Solidity: function CONTRACT_ASSETBOX() view returns(uint256)
func (_OtcImpl *OtcImplCallerSession) CONTRACTASSETBOX() (*big.Int, error) {
	return _OtcImpl.Contract.CONTRACTASSETBOX(&_OtcImpl.CallOpts)
}

// CONTRACTASSETBOXINFO is a free data retrieval call binding the contract method 0xb572fd3d.
//
// Solidity: function CONTRACT_ASSETBOX_INFO() view returns(uint256)
func (_OtcImpl *OtcImplCaller) CONTRACTASSETBOXINFO(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OtcImpl.contract.Call(opts, out, "CONTRACT_ASSETBOX_INFO")
	return *ret0, err
}

// CONTRACTASSETBOXINFO is a free data retrieval call binding the contract method 0xb572fd3d.
//
// Solidity: function CONTRACT_ASSETBOX_INFO() view returns(uint256)
func (_OtcImpl *OtcImplSession) CONTRACTASSETBOXINFO() (*big.Int, error) {
	return _OtcImpl.Contract.CONTRACTASSETBOXINFO(&_OtcImpl.CallOpts)
}

// CONTRACTASSETBOXINFO is a free data retrieval call binding the contract method 0xb572fd3d.
//
// Solidity: function CONTRACT_ASSETBOX_INFO() view returns(uint256)
func (_OtcImpl *OtcImplCallerSession) CONTRACTASSETBOXINFO() (*big.Int, error) {
	return _OtcImpl.Contract.CONTRACTASSETBOXINFO(&_OtcImpl.CallOpts)
}

// CONTRACTASSETBOXINFOSTORAGE is a free data retrieval call binding the contract method 0xf32a115d.
//
// Solidity: function CONTRACT_ASSETBOX_INFO_STORAGE() view returns(uint256)
func (_OtcImpl *OtcImplCaller) CONTRACTASSETBOXINFOSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OtcImpl.contract.Call(opts, out, "CONTRACT_ASSETBOX_INFO_STORAGE")
	return *ret0, err
}

// CONTRACTASSETBOXINFOSTORAGE is a free data retrieval call binding the contract method 0xf32a115d.
//
// Solidity: function CONTRACT_ASSETBOX_INFO_STORAGE() view returns(uint256)
func (_OtcImpl *OtcImplSession) CONTRACTASSETBOXINFOSTORAGE() (*big.Int, error) {
	return _OtcImpl.Contract.CONTRACTASSETBOXINFOSTORAGE(&_OtcImpl.CallOpts)
}

// CONTRACTASSETBOXINFOSTORAGE is a free data retrieval call binding the contract method 0xf32a115d.
//
// Solidity: function CONTRACT_ASSETBOX_INFO_STORAGE() view returns(uint256)
func (_OtcImpl *OtcImplCallerSession) CONTRACTASSETBOXINFOSTORAGE() (*big.Int, error) {
	return _OtcImpl.Contract.CONTRACTASSETBOXINFOSTORAGE(&_OtcImpl.CallOpts)
}

// CONTRACTASSETBOXSTORAGE is a free data retrieval call binding the contract method 0xfa5bdafc.
//
// Solidity: function CONTRACT_ASSETBOX_STORAGE() view returns(uint256)
func (_OtcImpl *OtcImplCaller) CONTRACTASSETBOXSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OtcImpl.contract.Call(opts, out, "CONTRACT_ASSETBOX_STORAGE")
	return *ret0, err
}

// CONTRACTASSETBOXSTORAGE is a free data retrieval call binding the contract method 0xfa5bdafc.
//
// Solidity: function CONTRACT_ASSETBOX_STORAGE() view returns(uint256)
func (_OtcImpl *OtcImplSession) CONTRACTASSETBOXSTORAGE() (*big.Int, error) {
	return _OtcImpl.Contract.CONTRACTASSETBOXSTORAGE(&_OtcImpl.CallOpts)
}

// CONTRACTASSETBOXSTORAGE is a free data retrieval call binding the contract method 0xfa5bdafc.
//
// Solidity: function CONTRACT_ASSETBOX_STORAGE() view returns(uint256)
func (_OtcImpl *OtcImplCallerSession) CONTRACTASSETBOXSTORAGE() (*big.Int, error) {
	return _OtcImpl.Contract.CONTRACTASSETBOXSTORAGE(&_OtcImpl.CallOpts)
}

// CONTRACTBALANCE is a free data retrieval call binding the contract method 0xd59045f5.
//
// Solidity: function CONTRACT_BALANCE() view returns(uint256)
func (_OtcImpl *OtcImplCaller) CONTRACTBALANCE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OtcImpl.contract.Call(opts, out, "CONTRACT_BALANCE")
	return *ret0, err
}

// CONTRACTBALANCE is a free data retrieval call binding the contract method 0xd59045f5.
//
// Solidity: function CONTRACT_BALANCE() view returns(uint256)
func (_OtcImpl *OtcImplSession) CONTRACTBALANCE() (*big.Int, error) {
	return _OtcImpl.Contract.CONTRACTBALANCE(&_OtcImpl.CallOpts)
}

// CONTRACTBALANCE is a free data retrieval call binding the contract method 0xd59045f5.
//
// Solidity: function CONTRACT_BALANCE() view returns(uint256)
func (_OtcImpl *OtcImplCallerSession) CONTRACTBALANCE() (*big.Int, error) {
	return _OtcImpl.Contract.CONTRACTBALANCE(&_OtcImpl.CallOpts)
}

// CONTRACTBITBON is a free data retrieval call binding the contract method 0xb6b433ab.
//
// Solidity: function CONTRACT_BITBON() view returns(uint256)
func (_OtcImpl *OtcImplCaller) CONTRACTBITBON(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OtcImpl.contract.Call(opts, out, "CONTRACT_BITBON")
	return *ret0, err
}

// CONTRACTBITBON is a free data retrieval call binding the contract method 0xb6b433ab.
//
// Solidity: function CONTRACT_BITBON() view returns(uint256)
func (_OtcImpl *OtcImplSession) CONTRACTBITBON() (*big.Int, error) {
	return _OtcImpl.Contract.CONTRACTBITBON(&_OtcImpl.CallOpts)
}

// CONTRACTBITBON is a free data retrieval call binding the contract method 0xb6b433ab.
//
// Solidity: function CONTRACT_BITBON() view returns(uint256)
func (_OtcImpl *OtcImplCallerSession) CONTRACTBITBON() (*big.Int, error) {
	return _OtcImpl.Contract.CONTRACTBITBON(&_OtcImpl.CallOpts)
}

// CONTRACTBITBONSTORAGE is a free data retrieval call binding the contract method 0xe7f6c30e.
//
// Solidity: function CONTRACT_BITBON_STORAGE() view returns(uint256)
func (_OtcImpl *OtcImplCaller) CONTRACTBITBONSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OtcImpl.contract.Call(opts, out, "CONTRACT_BITBON_STORAGE")
	return *ret0, err
}

// CONTRACTBITBONSTORAGE is a free data retrieval call binding the contract method 0xe7f6c30e.
//
// Solidity: function CONTRACT_BITBON_STORAGE() view returns(uint256)
func (_OtcImpl *OtcImplSession) CONTRACTBITBONSTORAGE() (*big.Int, error) {
	return _OtcImpl.Contract.CONTRACTBITBONSTORAGE(&_OtcImpl.CallOpts)
}

// CONTRACTBITBONSTORAGE is a free data retrieval call binding the contract method 0xe7f6c30e.
//
// Solidity: function CONTRACT_BITBON_STORAGE() view returns(uint256)
func (_OtcImpl *OtcImplCallerSession) CONTRACTBITBONSTORAGE() (*big.Int, error) {
	return _OtcImpl.Contract.CONTRACTBITBONSTORAGE(&_OtcImpl.CallOpts)
}

// CONTRACTBITBONSUPPORT is a free data retrieval call binding the contract method 0x20139c74.
//
// Solidity: function CONTRACT_BITBON_SUPPORT() view returns(uint256)
func (_OtcImpl *OtcImplCaller) CONTRACTBITBONSUPPORT(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OtcImpl.contract.Call(opts, out, "CONTRACT_BITBON_SUPPORT")
	return *ret0, err
}

// CONTRACTBITBONSUPPORT is a free data retrieval call binding the contract method 0x20139c74.
//
// Solidity: function CONTRACT_BITBON_SUPPORT() view returns(uint256)
func (_OtcImpl *OtcImplSession) CONTRACTBITBONSUPPORT() (*big.Int, error) {
	return _OtcImpl.Contract.CONTRACTBITBONSUPPORT(&_OtcImpl.CallOpts)
}

// CONTRACTBITBONSUPPORT is a free data retrieval call binding the contract method 0x20139c74.
//
// Solidity: function CONTRACT_BITBON_SUPPORT() view returns(uint256)
func (_OtcImpl *OtcImplCallerSession) CONTRACTBITBONSUPPORT() (*big.Int, error) {
	return _OtcImpl.Contract.CONTRACTBITBONSUPPORT(&_OtcImpl.CallOpts)
}

// CONTRACTBONBALANCE is a free data retrieval call binding the contract method 0x40e02b7c.
//
// Solidity: function CONTRACT_BONBALANCE() view returns(uint256)
func (_OtcImpl *OtcImplCaller) CONTRACTBONBALANCE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OtcImpl.contract.Call(opts, out, "CONTRACT_BONBALANCE")
	return *ret0, err
}

// CONTRACTBONBALANCE is a free data retrieval call binding the contract method 0x40e02b7c.
//
// Solidity: function CONTRACT_BONBALANCE() view returns(uint256)
func (_OtcImpl *OtcImplSession) CONTRACTBONBALANCE() (*big.Int, error) {
	return _OtcImpl.Contract.CONTRACTBONBALANCE(&_OtcImpl.CallOpts)
}

// CONTRACTBONBALANCE is a free data retrieval call binding the contract method 0x40e02b7c.
//
// Solidity: function CONTRACT_BONBALANCE() view returns(uint256)
func (_OtcImpl *OtcImplCallerSession) CONTRACTBONBALANCE() (*big.Int, error) {
	return _OtcImpl.Contract.CONTRACTBONBALANCE(&_OtcImpl.CallOpts)
}

// CONTRACTBONTRANSFER is a free data retrieval call binding the contract method 0xc18f4a8a.
//
// Solidity: function CONTRACT_BONTRANSFER() view returns(uint256)
func (_OtcImpl *OtcImplCaller) CONTRACTBONTRANSFER(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OtcImpl.contract.Call(opts, out, "CONTRACT_BONTRANSFER")
	return *ret0, err
}

// CONTRACTBONTRANSFER is a free data retrieval call binding the contract method 0xc18f4a8a.
//
// Solidity: function CONTRACT_BONTRANSFER() view returns(uint256)
func (_OtcImpl *OtcImplSession) CONTRACTBONTRANSFER() (*big.Int, error) {
	return _OtcImpl.Contract.CONTRACTBONTRANSFER(&_OtcImpl.CallOpts)
}

// CONTRACTBONTRANSFER is a free data retrieval call binding the contract method 0xc18f4a8a.
//
// Solidity: function CONTRACT_BONTRANSFER() view returns(uint256)
func (_OtcImpl *OtcImplCallerSession) CONTRACTBONTRANSFER() (*big.Int, error) {
	return _OtcImpl.Contract.CONTRACTBONTRANSFER(&_OtcImpl.CallOpts)
}

// CONTRACTDEX is a free data retrieval call binding the contract method 0x52c1f844.
//
// Solidity: function CONTRACT_DEX() view returns(uint256)
func (_OtcImpl *OtcImplCaller) CONTRACTDEX(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OtcImpl.contract.Call(opts, out, "CONTRACT_DEX")
	return *ret0, err
}

// CONTRACTDEX is a free data retrieval call binding the contract method 0x52c1f844.
//
// Solidity: function CONTRACT_DEX() view returns(uint256)
func (_OtcImpl *OtcImplSession) CONTRACTDEX() (*big.Int, error) {
	return _OtcImpl.Contract.CONTRACTDEX(&_OtcImpl.CallOpts)
}

// CONTRACTDEX is a free data retrieval call binding the contract method 0x52c1f844.
//
// Solidity: function CONTRACT_DEX() view returns(uint256)
func (_OtcImpl *OtcImplCallerSession) CONTRACTDEX() (*big.Int, error) {
	return _OtcImpl.Contract.CONTRACTDEX(&_OtcImpl.CallOpts)
}

// CONTRACTDISTRIBUTIONSTORAGE is a free data retrieval call binding the contract method 0x08097327.
//
// Solidity: function CONTRACT_DISTRIBUTION_STORAGE() view returns(uint256)
func (_OtcImpl *OtcImplCaller) CONTRACTDISTRIBUTIONSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OtcImpl.contract.Call(opts, out, "CONTRACT_DISTRIBUTION_STORAGE")
	return *ret0, err
}

// CONTRACTDISTRIBUTIONSTORAGE is a free data retrieval call binding the contract method 0x08097327.
//
// Solidity: function CONTRACT_DISTRIBUTION_STORAGE() view returns(uint256)
func (_OtcImpl *OtcImplSession) CONTRACTDISTRIBUTIONSTORAGE() (*big.Int, error) {
	return _OtcImpl.Contract.CONTRACTDISTRIBUTIONSTORAGE(&_OtcImpl.CallOpts)
}

// CONTRACTDISTRIBUTIONSTORAGE is a free data retrieval call binding the contract method 0x08097327.
//
// Solidity: function CONTRACT_DISTRIBUTION_STORAGE() view returns(uint256)
func (_OtcImpl *OtcImplCallerSession) CONTRACTDISTRIBUTIONSTORAGE() (*big.Int, error) {
	return _OtcImpl.Contract.CONTRACTDISTRIBUTIONSTORAGE(&_OtcImpl.CallOpts)
}

// CONTRACTEXCHANGE is a free data retrieval call binding the contract method 0xa0e3a10e.
//
// Solidity: function CONTRACT_EXCHANGE() view returns(uint256)
func (_OtcImpl *OtcImplCaller) CONTRACTEXCHANGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OtcImpl.contract.Call(opts, out, "CONTRACT_EXCHANGE")
	return *ret0, err
}

// CONTRACTEXCHANGE is a free data retrieval call binding the contract method 0xa0e3a10e.
//
// Solidity: function CONTRACT_EXCHANGE() view returns(uint256)
func (_OtcImpl *OtcImplSession) CONTRACTEXCHANGE() (*big.Int, error) {
	return _OtcImpl.Contract.CONTRACTEXCHANGE(&_OtcImpl.CallOpts)
}

// CONTRACTEXCHANGE is a free data retrieval call binding the contract method 0xa0e3a10e.
//
// Solidity: function CONTRACT_EXCHANGE() view returns(uint256)
func (_OtcImpl *OtcImplCallerSession) CONTRACTEXCHANGE() (*big.Int, error) {
	return _OtcImpl.Contract.CONTRACTEXCHANGE(&_OtcImpl.CallOpts)
}

// CONTRACTEXCHANGESTORAGE is a free data retrieval call binding the contract method 0x001570f2.
//
// Solidity: function CONTRACT_EXCHANGE_STORAGE() view returns(uint256)
func (_OtcImpl *OtcImplCaller) CONTRACTEXCHANGESTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OtcImpl.contract.Call(opts, out, "CONTRACT_EXCHANGE_STORAGE")
	return *ret0, err
}

// CONTRACTEXCHANGESTORAGE is a free data retrieval call binding the contract method 0x001570f2.
//
// Solidity: function CONTRACT_EXCHANGE_STORAGE() view returns(uint256)
func (_OtcImpl *OtcImplSession) CONTRACTEXCHANGESTORAGE() (*big.Int, error) {
	return _OtcImpl.Contract.CONTRACTEXCHANGESTORAGE(&_OtcImpl.CallOpts)
}

// CONTRACTEXCHANGESTORAGE is a free data retrieval call binding the contract method 0x001570f2.
//
// Solidity: function CONTRACT_EXCHANGE_STORAGE() view returns(uint256)
func (_OtcImpl *OtcImplCallerSession) CONTRACTEXCHANGESTORAGE() (*big.Int, error) {
	return _OtcImpl.Contract.CONTRACTEXCHANGESTORAGE(&_OtcImpl.CallOpts)
}

// CONTRACTFEESTORAGE is a free data retrieval call binding the contract method 0xcd5e65a4.
//
// Solidity: function CONTRACT_FEE_STORAGE() view returns(uint256)
func (_OtcImpl *OtcImplCaller) CONTRACTFEESTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OtcImpl.contract.Call(opts, out, "CONTRACT_FEE_STORAGE")
	return *ret0, err
}

// CONTRACTFEESTORAGE is a free data retrieval call binding the contract method 0xcd5e65a4.
//
// Solidity: function CONTRACT_FEE_STORAGE() view returns(uint256)
func (_OtcImpl *OtcImplSession) CONTRACTFEESTORAGE() (*big.Int, error) {
	return _OtcImpl.Contract.CONTRACTFEESTORAGE(&_OtcImpl.CallOpts)
}

// CONTRACTFEESTORAGE is a free data retrieval call binding the contract method 0xcd5e65a4.
//
// Solidity: function CONTRACT_FEE_STORAGE() view returns(uint256)
func (_OtcImpl *OtcImplCallerSession) CONTRACTFEESTORAGE() (*big.Int, error) {
	return _OtcImpl.Contract.CONTRACTFEESTORAGE(&_OtcImpl.CallOpts)
}

// CONTRACTGENERATOR is a free data retrieval call binding the contract method 0xccea2b46.
//
// Solidity: function CONTRACT_GENERATOR() view returns(uint256)
func (_OtcImpl *OtcImplCaller) CONTRACTGENERATOR(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OtcImpl.contract.Call(opts, out, "CONTRACT_GENERATOR")
	return *ret0, err
}

// CONTRACTGENERATOR is a free data retrieval call binding the contract method 0xccea2b46.
//
// Solidity: function CONTRACT_GENERATOR() view returns(uint256)
func (_OtcImpl *OtcImplSession) CONTRACTGENERATOR() (*big.Int, error) {
	return _OtcImpl.Contract.CONTRACTGENERATOR(&_OtcImpl.CallOpts)
}

// CONTRACTGENERATOR is a free data retrieval call binding the contract method 0xccea2b46.
//
// Solidity: function CONTRACT_GENERATOR() view returns(uint256)
func (_OtcImpl *OtcImplCallerSession) CONTRACTGENERATOR() (*big.Int, error) {
	return _OtcImpl.Contract.CONTRACTGENERATOR(&_OtcImpl.CallOpts)
}

// CONTRACTMININGAGENTSTORAGE is a free data retrieval call binding the contract method 0x83cc588d.
//
// Solidity: function CONTRACT_MINING_AGENT_STORAGE() view returns(uint256)
func (_OtcImpl *OtcImplCaller) CONTRACTMININGAGENTSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OtcImpl.contract.Call(opts, out, "CONTRACT_MINING_AGENT_STORAGE")
	return *ret0, err
}

// CONTRACTMININGAGENTSTORAGE is a free data retrieval call binding the contract method 0x83cc588d.
//
// Solidity: function CONTRACT_MINING_AGENT_STORAGE() view returns(uint256)
func (_OtcImpl *OtcImplSession) CONTRACTMININGAGENTSTORAGE() (*big.Int, error) {
	return _OtcImpl.Contract.CONTRACTMININGAGENTSTORAGE(&_OtcImpl.CallOpts)
}

// CONTRACTMININGAGENTSTORAGE is a free data retrieval call binding the contract method 0x83cc588d.
//
// Solidity: function CONTRACT_MINING_AGENT_STORAGE() view returns(uint256)
func (_OtcImpl *OtcImplCallerSession) CONTRACTMININGAGENTSTORAGE() (*big.Int, error) {
	return _OtcImpl.Contract.CONTRACTMININGAGENTSTORAGE(&_OtcImpl.CallOpts)
}

// CONTRACTMSBON is a free data retrieval call binding the contract method 0x28b3a976.
//
// Solidity: function CONTRACT_MSBON() view returns(uint256)
func (_OtcImpl *OtcImplCaller) CONTRACTMSBON(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OtcImpl.contract.Call(opts, out, "CONTRACT_MSBON")
	return *ret0, err
}

// CONTRACTMSBON is a free data retrieval call binding the contract method 0x28b3a976.
//
// Solidity: function CONTRACT_MSBON() view returns(uint256)
func (_OtcImpl *OtcImplSession) CONTRACTMSBON() (*big.Int, error) {
	return _OtcImpl.Contract.CONTRACTMSBON(&_OtcImpl.CallOpts)
}

// CONTRACTMSBON is a free data retrieval call binding the contract method 0x28b3a976.
//
// Solidity: function CONTRACT_MSBON() view returns(uint256)
func (_OtcImpl *OtcImplCallerSession) CONTRACTMSBON() (*big.Int, error) {
	return _OtcImpl.Contract.CONTRACTMSBON(&_OtcImpl.CallOpts)
}

// CONTRACTMSBONSTORAGE is a free data retrieval call binding the contract method 0xe04e6061.
//
// Solidity: function CONTRACT_MSBON_STORAGE() view returns(uint256)
func (_OtcImpl *OtcImplCaller) CONTRACTMSBONSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OtcImpl.contract.Call(opts, out, "CONTRACT_MSBON_STORAGE")
	return *ret0, err
}

// CONTRACTMSBONSTORAGE is a free data retrieval call binding the contract method 0xe04e6061.
//
// Solidity: function CONTRACT_MSBON_STORAGE() view returns(uint256)
func (_OtcImpl *OtcImplSession) CONTRACTMSBONSTORAGE() (*big.Int, error) {
	return _OtcImpl.Contract.CONTRACTMSBONSTORAGE(&_OtcImpl.CallOpts)
}

// CONTRACTMSBONSTORAGE is a free data retrieval call binding the contract method 0xe04e6061.
//
// Solidity: function CONTRACT_MSBON_STORAGE() view returns(uint256)
func (_OtcImpl *OtcImplCallerSession) CONTRACTMSBONSTORAGE() (*big.Int, error) {
	return _OtcImpl.Contract.CONTRACTMSBONSTORAGE(&_OtcImpl.CallOpts)
}

// CONTRACTMULTISIGADDADDRESS is a free data retrieval call binding the contract method 0xf4f65a5a.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ADDRESS() view returns(uint256)
func (_OtcImpl *OtcImplCaller) CONTRACTMULTISIGADDADDRESS(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OtcImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_ADD_ADDRESS")
	return *ret0, err
}

// CONTRACTMULTISIGADDADDRESS is a free data retrieval call binding the contract method 0xf4f65a5a.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ADDRESS() view returns(uint256)
func (_OtcImpl *OtcImplSession) CONTRACTMULTISIGADDADDRESS() (*big.Int, error) {
	return _OtcImpl.Contract.CONTRACTMULTISIGADDADDRESS(&_OtcImpl.CallOpts)
}

// CONTRACTMULTISIGADDADDRESS is a free data retrieval call binding the contract method 0xf4f65a5a.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ADDRESS() view returns(uint256)
func (_OtcImpl *OtcImplCallerSession) CONTRACTMULTISIGADDADDRESS() (*big.Int, error) {
	return _OtcImpl.Contract.CONTRACTMULTISIGADDADDRESS(&_OtcImpl.CallOpts)
}

// CONTRACTMULTISIGADDADMIN is a free data retrieval call binding the contract method 0x83888c68.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ADMIN() view returns(uint256)
func (_OtcImpl *OtcImplCaller) CONTRACTMULTISIGADDADMIN(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OtcImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_ADD_ADMIN")
	return *ret0, err
}

// CONTRACTMULTISIGADDADMIN is a free data retrieval call binding the contract method 0x83888c68.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ADMIN() view returns(uint256)
func (_OtcImpl *OtcImplSession) CONTRACTMULTISIGADDADMIN() (*big.Int, error) {
	return _OtcImpl.Contract.CONTRACTMULTISIGADDADMIN(&_OtcImpl.CallOpts)
}

// CONTRACTMULTISIGADDADMIN is a free data retrieval call binding the contract method 0x83888c68.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ADMIN() view returns(uint256)
func (_OtcImpl *OtcImplCallerSession) CONTRACTMULTISIGADDADMIN() (*big.Int, error) {
	return _OtcImpl.Contract.CONTRACTMULTISIGADDADMIN(&_OtcImpl.CallOpts)
}

// CONTRACTMULTISIGADDFEE is a free data retrieval call binding the contract method 0xc43e8cab.
//
// Solidity: function CONTRACT_MULTISIG_ADD_FEE() view returns(uint256)
func (_OtcImpl *OtcImplCaller) CONTRACTMULTISIGADDFEE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OtcImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_ADD_FEE")
	return *ret0, err
}

// CONTRACTMULTISIGADDFEE is a free data retrieval call binding the contract method 0xc43e8cab.
//
// Solidity: function CONTRACT_MULTISIG_ADD_FEE() view returns(uint256)
func (_OtcImpl *OtcImplSession) CONTRACTMULTISIGADDFEE() (*big.Int, error) {
	return _OtcImpl.Contract.CONTRACTMULTISIGADDFEE(&_OtcImpl.CallOpts)
}

// CONTRACTMULTISIGADDFEE is a free data retrieval call binding the contract method 0xc43e8cab.
//
// Solidity: function CONTRACT_MULTISIG_ADD_FEE() view returns(uint256)
func (_OtcImpl *OtcImplCallerSession) CONTRACTMULTISIGADDFEE() (*big.Int, error) {
	return _OtcImpl.Contract.CONTRACTMULTISIGADDFEE(&_OtcImpl.CallOpts)
}

// CONTRACTMULTISIGADDROLE is a free data retrieval call binding the contract method 0x20a0914b.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ROLE() view returns(uint256)
func (_OtcImpl *OtcImplCaller) CONTRACTMULTISIGADDROLE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OtcImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_ADD_ROLE")
	return *ret0, err
}

// CONTRACTMULTISIGADDROLE is a free data retrieval call binding the contract method 0x20a0914b.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ROLE() view returns(uint256)
func (_OtcImpl *OtcImplSession) CONTRACTMULTISIGADDROLE() (*big.Int, error) {
	return _OtcImpl.Contract.CONTRACTMULTISIGADDROLE(&_OtcImpl.CallOpts)
}

// CONTRACTMULTISIGADDROLE is a free data retrieval call binding the contract method 0x20a0914b.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ROLE() view returns(uint256)
func (_OtcImpl *OtcImplCallerSession) CONTRACTMULTISIGADDROLE() (*big.Int, error) {
	return _OtcImpl.Contract.CONTRACTMULTISIGADDROLE(&_OtcImpl.CallOpts)
}

// CONTRACTMULTISIGDROPPERMISSION is a free data retrieval call binding the contract method 0x3f3fbc39.
//
// Solidity: function CONTRACT_MULTISIG_DROP_PERMISSION() view returns(uint256)
func (_OtcImpl *OtcImplCaller) CONTRACTMULTISIGDROPPERMISSION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OtcImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_DROP_PERMISSION")
	return *ret0, err
}

// CONTRACTMULTISIGDROPPERMISSION is a free data retrieval call binding the contract method 0x3f3fbc39.
//
// Solidity: function CONTRACT_MULTISIG_DROP_PERMISSION() view returns(uint256)
func (_OtcImpl *OtcImplSession) CONTRACTMULTISIGDROPPERMISSION() (*big.Int, error) {
	return _OtcImpl.Contract.CONTRACTMULTISIGDROPPERMISSION(&_OtcImpl.CallOpts)
}

// CONTRACTMULTISIGDROPPERMISSION is a free data retrieval call binding the contract method 0x3f3fbc39.
//
// Solidity: function CONTRACT_MULTISIG_DROP_PERMISSION() view returns(uint256)
func (_OtcImpl *OtcImplCallerSession) CONTRACTMULTISIGDROPPERMISSION() (*big.Int, error) {
	return _OtcImpl.Contract.CONTRACTMULTISIGDROPPERMISSION(&_OtcImpl.CallOpts)
}

// CONTRACTMULTISIGEDITADMIN is a free data retrieval call binding the contract method 0x91b0f051.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_ADMIN() view returns(uint256)
func (_OtcImpl *OtcImplCaller) CONTRACTMULTISIGEDITADMIN(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OtcImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_EDIT_ADMIN")
	return *ret0, err
}

// CONTRACTMULTISIGEDITADMIN is a free data retrieval call binding the contract method 0x91b0f051.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_ADMIN() view returns(uint256)
func (_OtcImpl *OtcImplSession) CONTRACTMULTISIGEDITADMIN() (*big.Int, error) {
	return _OtcImpl.Contract.CONTRACTMULTISIGEDITADMIN(&_OtcImpl.CallOpts)
}

// CONTRACTMULTISIGEDITADMIN is a free data retrieval call binding the contract method 0x91b0f051.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_ADMIN() view returns(uint256)
func (_OtcImpl *OtcImplCallerSession) CONTRACTMULTISIGEDITADMIN() (*big.Int, error) {
	return _OtcImpl.Contract.CONTRACTMULTISIGEDITADMIN(&_OtcImpl.CallOpts)
}

// CONTRACTMULTISIGEDITFEE is a free data retrieval call binding the contract method 0x6b2abe17.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_FEE() view returns(uint256)
func (_OtcImpl *OtcImplCaller) CONTRACTMULTISIGEDITFEE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OtcImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_EDIT_FEE")
	return *ret0, err
}

// CONTRACTMULTISIGEDITFEE is a free data retrieval call binding the contract method 0x6b2abe17.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_FEE() view returns(uint256)
func (_OtcImpl *OtcImplSession) CONTRACTMULTISIGEDITFEE() (*big.Int, error) {
	return _OtcImpl.Contract.CONTRACTMULTISIGEDITFEE(&_OtcImpl.CallOpts)
}

// CONTRACTMULTISIGEDITFEE is a free data retrieval call binding the contract method 0x6b2abe17.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_FEE() view returns(uint256)
func (_OtcImpl *OtcImplCallerSession) CONTRACTMULTISIGEDITFEE() (*big.Int, error) {
	return _OtcImpl.Contract.CONTRACTMULTISIGEDITFEE(&_OtcImpl.CallOpts)
}

// CONTRACTMULTISIGEDITROLE is a free data retrieval call binding the contract method 0xc5f66412.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_ROLE() view returns(uint256)
func (_OtcImpl *OtcImplCaller) CONTRACTMULTISIGEDITROLE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OtcImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_EDIT_ROLE")
	return *ret0, err
}

// CONTRACTMULTISIGEDITROLE is a free data retrieval call binding the contract method 0xc5f66412.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_ROLE() view returns(uint256)
func (_OtcImpl *OtcImplSession) CONTRACTMULTISIGEDITROLE() (*big.Int, error) {
	return _OtcImpl.Contract.CONTRACTMULTISIGEDITROLE(&_OtcImpl.CallOpts)
}

// CONTRACTMULTISIGEDITROLE is a free data retrieval call binding the contract method 0xc5f66412.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_ROLE() view returns(uint256)
func (_OtcImpl *OtcImplCallerSession) CONTRACTMULTISIGEDITROLE() (*big.Int, error) {
	return _OtcImpl.Contract.CONTRACTMULTISIGEDITROLE(&_OtcImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEADMINBYADDRESS is a free data retrieval call binding the contract method 0xc816b091.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_BY_ADDRESS() view returns(uint256)
func (_OtcImpl *OtcImplCaller) CONTRACTMULTISIGREMOVEADMINBYADDRESS(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OtcImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REMOVE_ADMIN_BY_ADDRESS")
	return *ret0, err
}

// CONTRACTMULTISIGREMOVEADMINBYADDRESS is a free data retrieval call binding the contract method 0xc816b091.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_BY_ADDRESS() view returns(uint256)
func (_OtcImpl *OtcImplSession) CONTRACTMULTISIGREMOVEADMINBYADDRESS() (*big.Int, error) {
	return _OtcImpl.Contract.CONTRACTMULTISIGREMOVEADMINBYADDRESS(&_OtcImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEADMINBYADDRESS is a free data retrieval call binding the contract method 0xc816b091.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_BY_ADDRESS() view returns(uint256)
func (_OtcImpl *OtcImplCallerSession) CONTRACTMULTISIGREMOVEADMINBYADDRESS() (*big.Int, error) {
	return _OtcImpl.Contract.CONTRACTMULTISIGREMOVEADMINBYADDRESS(&_OtcImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEADMINBYINDEX is a free data retrieval call binding the contract method 0x751fbea2.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_BY_INDEX() view returns(uint256)
func (_OtcImpl *OtcImplCaller) CONTRACTMULTISIGREMOVEADMINBYINDEX(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OtcImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REMOVE_ADMIN_BY_INDEX")
	return *ret0, err
}

// CONTRACTMULTISIGREMOVEADMINBYINDEX is a free data retrieval call binding the contract method 0x751fbea2.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_BY_INDEX() view returns(uint256)
func (_OtcImpl *OtcImplSession) CONTRACTMULTISIGREMOVEADMINBYINDEX() (*big.Int, error) {
	return _OtcImpl.Contract.CONTRACTMULTISIGREMOVEADMINBYINDEX(&_OtcImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEADMINBYINDEX is a free data retrieval call binding the contract method 0x751fbea2.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_BY_INDEX() view returns(uint256)
func (_OtcImpl *OtcImplCallerSession) CONTRACTMULTISIGREMOVEADMINBYINDEX() (*big.Int, error) {
	return _OtcImpl.Contract.CONTRACTMULTISIGREMOVEADMINBYINDEX(&_OtcImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEADMINROLE is a free data retrieval call binding the contract method 0x53c00200.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_ROLE() view returns(uint256)
func (_OtcImpl *OtcImplCaller) CONTRACTMULTISIGREMOVEADMINROLE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OtcImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REMOVE_ADMIN_ROLE")
	return *ret0, err
}

// CONTRACTMULTISIGREMOVEADMINROLE is a free data retrieval call binding the contract method 0x53c00200.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_ROLE() view returns(uint256)
func (_OtcImpl *OtcImplSession) CONTRACTMULTISIGREMOVEADMINROLE() (*big.Int, error) {
	return _OtcImpl.Contract.CONTRACTMULTISIGREMOVEADMINROLE(&_OtcImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEADMINROLE is a free data retrieval call binding the contract method 0x53c00200.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_ROLE() view returns(uint256)
func (_OtcImpl *OtcImplCallerSession) CONTRACTMULTISIGREMOVEADMINROLE() (*big.Int, error) {
	return _OtcImpl.Contract.CONTRACTMULTISIGREMOVEADMINROLE(&_OtcImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEFEE is a free data retrieval call binding the contract method 0xdb614d25.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_FEE() view returns(uint256)
func (_OtcImpl *OtcImplCaller) CONTRACTMULTISIGREMOVEFEE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OtcImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REMOVE_FEE")
	return *ret0, err
}

// CONTRACTMULTISIGREMOVEFEE is a free data retrieval call binding the contract method 0xdb614d25.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_FEE() view returns(uint256)
func (_OtcImpl *OtcImplSession) CONTRACTMULTISIGREMOVEFEE() (*big.Int, error) {
	return _OtcImpl.Contract.CONTRACTMULTISIGREMOVEFEE(&_OtcImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEFEE is a free data retrieval call binding the contract method 0xdb614d25.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_FEE() view returns(uint256)
func (_OtcImpl *OtcImplCallerSession) CONTRACTMULTISIGREMOVEFEE() (*big.Int, error) {
	return _OtcImpl.Contract.CONTRACTMULTISIGREMOVEFEE(&_OtcImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEROLE is a free data retrieval call binding the contract method 0x23648496.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ROLE() view returns(uint256)
func (_OtcImpl *OtcImplCaller) CONTRACTMULTISIGREMOVEROLE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OtcImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REMOVE_ROLE")
	return *ret0, err
}

// CONTRACTMULTISIGREMOVEROLE is a free data retrieval call binding the contract method 0x23648496.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ROLE() view returns(uint256)
func (_OtcImpl *OtcImplSession) CONTRACTMULTISIGREMOVEROLE() (*big.Int, error) {
	return _OtcImpl.Contract.CONTRACTMULTISIGREMOVEROLE(&_OtcImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEROLE is a free data retrieval call binding the contract method 0x23648496.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ROLE() view returns(uint256)
func (_OtcImpl *OtcImplCallerSession) CONTRACTMULTISIGREMOVEROLE() (*big.Int, error) {
	return _OtcImpl.Contract.CONTRACTMULTISIGREMOVEROLE(&_OtcImpl.CallOpts)
}

// CONTRACTMULTISIGREPLENISHCAPITALIZATION is a free data retrieval call binding the contract method 0x62d5d523.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_CAPITALIZATION() view returns(uint256)
func (_OtcImpl *OtcImplCaller) CONTRACTMULTISIGREPLENISHCAPITALIZATION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OtcImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REPLENISH_CAPITALIZATION")
	return *ret0, err
}

// CONTRACTMULTISIGREPLENISHCAPITALIZATION is a free data retrieval call binding the contract method 0x62d5d523.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_CAPITALIZATION() view returns(uint256)
func (_OtcImpl *OtcImplSession) CONTRACTMULTISIGREPLENISHCAPITALIZATION() (*big.Int, error) {
	return _OtcImpl.Contract.CONTRACTMULTISIGREPLENISHCAPITALIZATION(&_OtcImpl.CallOpts)
}

// CONTRACTMULTISIGREPLENISHCAPITALIZATION is a free data retrieval call binding the contract method 0x62d5d523.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_CAPITALIZATION() view returns(uint256)
func (_OtcImpl *OtcImplCallerSession) CONTRACTMULTISIGREPLENISHCAPITALIZATION() (*big.Int, error) {
	return _OtcImpl.Contract.CONTRACTMULTISIGREPLENISHCAPITALIZATION(&_OtcImpl.CallOpts)
}

// CONTRACTMULTISIGREPLENISHCERTIFICATE is a free data retrieval call binding the contract method 0xfd19b7e1.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_CERTIFICATE() view returns(uint256)
func (_OtcImpl *OtcImplCaller) CONTRACTMULTISIGREPLENISHCERTIFICATE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OtcImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REPLENISH_CERTIFICATE")
	return *ret0, err
}

// CONTRACTMULTISIGREPLENISHCERTIFICATE is a free data retrieval call binding the contract method 0xfd19b7e1.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_CERTIFICATE() view returns(uint256)
func (_OtcImpl *OtcImplSession) CONTRACTMULTISIGREPLENISHCERTIFICATE() (*big.Int, error) {
	return _OtcImpl.Contract.CONTRACTMULTISIGREPLENISHCERTIFICATE(&_OtcImpl.CallOpts)
}

// CONTRACTMULTISIGREPLENISHCERTIFICATE is a free data retrieval call binding the contract method 0xfd19b7e1.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_CERTIFICATE() view returns(uint256)
func (_OtcImpl *OtcImplCallerSession) CONTRACTMULTISIGREPLENISHCERTIFICATE() (*big.Int, error) {
	return _OtcImpl.Contract.CONTRACTMULTISIGREPLENISHCERTIFICATE(&_OtcImpl.CallOpts)
}

// CONTRACTMULTISIGREPLENISHEMISSION is a free data retrieval call binding the contract method 0x0cb59fde.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_EMISSION() view returns(uint256)
func (_OtcImpl *OtcImplCaller) CONTRACTMULTISIGREPLENISHEMISSION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OtcImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REPLENISH_EMISSION")
	return *ret0, err
}

// CONTRACTMULTISIGREPLENISHEMISSION is a free data retrieval call binding the contract method 0x0cb59fde.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_EMISSION() view returns(uint256)
func (_OtcImpl *OtcImplSession) CONTRACTMULTISIGREPLENISHEMISSION() (*big.Int, error) {
	return _OtcImpl.Contract.CONTRACTMULTISIGREPLENISHEMISSION(&_OtcImpl.CallOpts)
}

// CONTRACTMULTISIGREPLENISHEMISSION is a free data retrieval call binding the contract method 0x0cb59fde.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_EMISSION() view returns(uint256)
func (_OtcImpl *OtcImplCallerSession) CONTRACTMULTISIGREPLENISHEMISSION() (*big.Int, error) {
	return _OtcImpl.Contract.CONTRACTMULTISIGREPLENISHEMISSION(&_OtcImpl.CallOpts)
}

// CONTRACTMULTISIGSETADMINROLE is a free data retrieval call binding the contract method 0x8f209ab7.
//
// Solidity: function CONTRACT_MULTISIG_SET_ADMIN_ROLE() view returns(uint256)
func (_OtcImpl *OtcImplCaller) CONTRACTMULTISIGSETADMINROLE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OtcImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_SET_ADMIN_ROLE")
	return *ret0, err
}

// CONTRACTMULTISIGSETADMINROLE is a free data retrieval call binding the contract method 0x8f209ab7.
//
// Solidity: function CONTRACT_MULTISIG_SET_ADMIN_ROLE() view returns(uint256)
func (_OtcImpl *OtcImplSession) CONTRACTMULTISIGSETADMINROLE() (*big.Int, error) {
	return _OtcImpl.Contract.CONTRACTMULTISIGSETADMINROLE(&_OtcImpl.CallOpts)
}

// CONTRACTMULTISIGSETADMINROLE is a free data retrieval call binding the contract method 0x8f209ab7.
//
// Solidity: function CONTRACT_MULTISIG_SET_ADMIN_ROLE() view returns(uint256)
func (_OtcImpl *OtcImplCallerSession) CONTRACTMULTISIGSETADMINROLE() (*big.Int, error) {
	return _OtcImpl.Contract.CONTRACTMULTISIGSETADMINROLE(&_OtcImpl.CallOpts)
}

// CONTRACTMULTISIGSETPERMISSION is a free data retrieval call binding the contract method 0xb0e91914.
//
// Solidity: function CONTRACT_MULTISIG_SET_PERMISSION() view returns(uint256)
func (_OtcImpl *OtcImplCaller) CONTRACTMULTISIGSETPERMISSION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OtcImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_SET_PERMISSION")
	return *ret0, err
}

// CONTRACTMULTISIGSETPERMISSION is a free data retrieval call binding the contract method 0xb0e91914.
//
// Solidity: function CONTRACT_MULTISIG_SET_PERMISSION() view returns(uint256)
func (_OtcImpl *OtcImplSession) CONTRACTMULTISIGSETPERMISSION() (*big.Int, error) {
	return _OtcImpl.Contract.CONTRACTMULTISIGSETPERMISSION(&_OtcImpl.CallOpts)
}

// CONTRACTMULTISIGSETPERMISSION is a free data retrieval call binding the contract method 0xb0e91914.
//
// Solidity: function CONTRACT_MULTISIG_SET_PERMISSION() view returns(uint256)
func (_OtcImpl *OtcImplCallerSession) CONTRACTMULTISIGSETPERMISSION() (*big.Int, error) {
	return _OtcImpl.Contract.CONTRACTMULTISIGSETPERMISSION(&_OtcImpl.CallOpts)
}

// CONTRACTMULTISIGSTORAGE is a free data retrieval call binding the contract method 0xe7e1a303.
//
// Solidity: function CONTRACT_MULTISIG_STORAGE() view returns(uint256)
func (_OtcImpl *OtcImplCaller) CONTRACTMULTISIGSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OtcImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_STORAGE")
	return *ret0, err
}

// CONTRACTMULTISIGSTORAGE is a free data retrieval call binding the contract method 0xe7e1a303.
//
// Solidity: function CONTRACT_MULTISIG_STORAGE() view returns(uint256)
func (_OtcImpl *OtcImplSession) CONTRACTMULTISIGSTORAGE() (*big.Int, error) {
	return _OtcImpl.Contract.CONTRACTMULTISIGSTORAGE(&_OtcImpl.CallOpts)
}

// CONTRACTMULTISIGSTORAGE is a free data retrieval call binding the contract method 0xe7e1a303.
//
// Solidity: function CONTRACT_MULTISIG_STORAGE() view returns(uint256)
func (_OtcImpl *OtcImplCallerSession) CONTRACTMULTISIGSTORAGE() (*big.Int, error) {
	return _OtcImpl.Contract.CONTRACTMULTISIGSTORAGE(&_OtcImpl.CallOpts)
}

// CONTRACTOTC is a free data retrieval call binding the contract method 0x5eceac70.
//
// Solidity: function CONTRACT_OTC() view returns(uint256)
func (_OtcImpl *OtcImplCaller) CONTRACTOTC(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OtcImpl.contract.Call(opts, out, "CONTRACT_OTC")
	return *ret0, err
}

// CONTRACTOTC is a free data retrieval call binding the contract method 0x5eceac70.
//
// Solidity: function CONTRACT_OTC() view returns(uint256)
func (_OtcImpl *OtcImplSession) CONTRACTOTC() (*big.Int, error) {
	return _OtcImpl.Contract.CONTRACTOTC(&_OtcImpl.CallOpts)
}

// CONTRACTOTC is a free data retrieval call binding the contract method 0x5eceac70.
//
// Solidity: function CONTRACT_OTC() view returns(uint256)
func (_OtcImpl *OtcImplCallerSession) CONTRACTOTC() (*big.Int, error) {
	return _OtcImpl.Contract.CONTRACTOTC(&_OtcImpl.CallOpts)
}

// CONTRACTOTCSTORAGE is a free data retrieval call binding the contract method 0x825b476f.
//
// Solidity: function CONTRACT_OTC_STORAGE() view returns(uint256)
func (_OtcImpl *OtcImplCaller) CONTRACTOTCSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OtcImpl.contract.Call(opts, out, "CONTRACT_OTC_STORAGE")
	return *ret0, err
}

// CONTRACTOTCSTORAGE is a free data retrieval call binding the contract method 0x825b476f.
//
// Solidity: function CONTRACT_OTC_STORAGE() view returns(uint256)
func (_OtcImpl *OtcImplSession) CONTRACTOTCSTORAGE() (*big.Int, error) {
	return _OtcImpl.Contract.CONTRACTOTCSTORAGE(&_OtcImpl.CallOpts)
}

// CONTRACTOTCSTORAGE is a free data retrieval call binding the contract method 0x825b476f.
//
// Solidity: function CONTRACT_OTC_STORAGE() view returns(uint256)
func (_OtcImpl *OtcImplCallerSession) CONTRACTOTCSTORAGE() (*big.Int, error) {
	return _OtcImpl.Contract.CONTRACTOTCSTORAGE(&_OtcImpl.CallOpts)
}

// CONTRACTREPLENISHCAPITALIZATION is a free data retrieval call binding the contract method 0xff5e09d8.
//
// Solidity: function CONTRACT_REPLENISH_CAPITALIZATION() view returns(uint256)
func (_OtcImpl *OtcImplCaller) CONTRACTREPLENISHCAPITALIZATION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OtcImpl.contract.Call(opts, out, "CONTRACT_REPLENISH_CAPITALIZATION")
	return *ret0, err
}

// CONTRACTREPLENISHCAPITALIZATION is a free data retrieval call binding the contract method 0xff5e09d8.
//
// Solidity: function CONTRACT_REPLENISH_CAPITALIZATION() view returns(uint256)
func (_OtcImpl *OtcImplSession) CONTRACTREPLENISHCAPITALIZATION() (*big.Int, error) {
	return _OtcImpl.Contract.CONTRACTREPLENISHCAPITALIZATION(&_OtcImpl.CallOpts)
}

// CONTRACTREPLENISHCAPITALIZATION is a free data retrieval call binding the contract method 0xff5e09d8.
//
// Solidity: function CONTRACT_REPLENISH_CAPITALIZATION() view returns(uint256)
func (_OtcImpl *OtcImplCallerSession) CONTRACTREPLENISHCAPITALIZATION() (*big.Int, error) {
	return _OtcImpl.Contract.CONTRACTREPLENISHCAPITALIZATION(&_OtcImpl.CallOpts)
}

// CONTRACTRESERVEDALIASSTORAGE is a free data retrieval call binding the contract method 0x70a94aa7.
//
// Solidity: function CONTRACT_RESERVED_ALIAS_STORAGE() view returns(uint256)
func (_OtcImpl *OtcImplCaller) CONTRACTRESERVEDALIASSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OtcImpl.contract.Call(opts, out, "CONTRACT_RESERVED_ALIAS_STORAGE")
	return *ret0, err
}

// CONTRACTRESERVEDALIASSTORAGE is a free data retrieval call binding the contract method 0x70a94aa7.
//
// Solidity: function CONTRACT_RESERVED_ALIAS_STORAGE() view returns(uint256)
func (_OtcImpl *OtcImplSession) CONTRACTRESERVEDALIASSTORAGE() (*big.Int, error) {
	return _OtcImpl.Contract.CONTRACTRESERVEDALIASSTORAGE(&_OtcImpl.CallOpts)
}

// CONTRACTRESERVEDALIASSTORAGE is a free data retrieval call binding the contract method 0x70a94aa7.
//
// Solidity: function CONTRACT_RESERVED_ALIAS_STORAGE() view returns(uint256)
func (_OtcImpl *OtcImplCallerSession) CONTRACTRESERVEDALIASSTORAGE() (*big.Int, error) {
	return _OtcImpl.Contract.CONTRACTRESERVEDALIASSTORAGE(&_OtcImpl.CallOpts)
}

// CONTRACTROLESTORAGE is a free data retrieval call binding the contract method 0xb938235b.
//
// Solidity: function CONTRACT_ROLE_STORAGE() view returns(uint256)
func (_OtcImpl *OtcImplCaller) CONTRACTROLESTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OtcImpl.contract.Call(opts, out, "CONTRACT_ROLE_STORAGE")
	return *ret0, err
}

// CONTRACTROLESTORAGE is a free data retrieval call binding the contract method 0xb938235b.
//
// Solidity: function CONTRACT_ROLE_STORAGE() view returns(uint256)
func (_OtcImpl *OtcImplSession) CONTRACTROLESTORAGE() (*big.Int, error) {
	return _OtcImpl.Contract.CONTRACTROLESTORAGE(&_OtcImpl.CallOpts)
}

// CONTRACTROLESTORAGE is a free data retrieval call binding the contract method 0xb938235b.
//
// Solidity: function CONTRACT_ROLE_STORAGE() view returns(uint256)
func (_OtcImpl *OtcImplCallerSession) CONTRACTROLESTORAGE() (*big.Int, error) {
	return _OtcImpl.Contract.CONTRACTROLESTORAGE(&_OtcImpl.CallOpts)
}

// CONTRACTSAFETRANSFERSTORAGE is a free data retrieval call binding the contract method 0xe7e45392.
//
// Solidity: function CONTRACT_SAFE_TRANSFER_STORAGE() view returns(uint256)
func (_OtcImpl *OtcImplCaller) CONTRACTSAFETRANSFERSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OtcImpl.contract.Call(opts, out, "CONTRACT_SAFE_TRANSFER_STORAGE")
	return *ret0, err
}

// CONTRACTSAFETRANSFERSTORAGE is a free data retrieval call binding the contract method 0xe7e45392.
//
// Solidity: function CONTRACT_SAFE_TRANSFER_STORAGE() view returns(uint256)
func (_OtcImpl *OtcImplSession) CONTRACTSAFETRANSFERSTORAGE() (*big.Int, error) {
	return _OtcImpl.Contract.CONTRACTSAFETRANSFERSTORAGE(&_OtcImpl.CallOpts)
}

// CONTRACTSAFETRANSFERSTORAGE is a free data retrieval call binding the contract method 0xe7e45392.
//
// Solidity: function CONTRACT_SAFE_TRANSFER_STORAGE() view returns(uint256)
func (_OtcImpl *OtcImplCallerSession) CONTRACTSAFETRANSFERSTORAGE() (*big.Int, error) {
	return _OtcImpl.Contract.CONTRACTSAFETRANSFERSTORAGE(&_OtcImpl.CallOpts)
}

// CONTRACTTRANSFER is a free data retrieval call binding the contract method 0xef94fe38.
//
// Solidity: function CONTRACT_TRANSFER() view returns(uint256)
func (_OtcImpl *OtcImplCaller) CONTRACTTRANSFER(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OtcImpl.contract.Call(opts, out, "CONTRACT_TRANSFER")
	return *ret0, err
}

// CONTRACTTRANSFER is a free data retrieval call binding the contract method 0xef94fe38.
//
// Solidity: function CONTRACT_TRANSFER() view returns(uint256)
func (_OtcImpl *OtcImplSession) CONTRACTTRANSFER() (*big.Int, error) {
	return _OtcImpl.Contract.CONTRACTTRANSFER(&_OtcImpl.CallOpts)
}

// CONTRACTTRANSFER is a free data retrieval call binding the contract method 0xef94fe38.
//
// Solidity: function CONTRACT_TRANSFER() view returns(uint256)
func (_OtcImpl *OtcImplCallerSession) CONTRACTTRANSFER() (*big.Int, error) {
	return _OtcImpl.Contract.CONTRACTTRANSFER(&_OtcImpl.CallOpts)
}

// CONTRACTTRANSFERFROMCAPITALIZATION is a free data retrieval call binding the contract method 0xc8f8decb.
//
// Solidity: function CONTRACT_TRANSFER_FROM_CAPITALIZATION() view returns(uint256)
func (_OtcImpl *OtcImplCaller) CONTRACTTRANSFERFROMCAPITALIZATION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OtcImpl.contract.Call(opts, out, "CONTRACT_TRANSFER_FROM_CAPITALIZATION")
	return *ret0, err
}

// CONTRACTTRANSFERFROMCAPITALIZATION is a free data retrieval call binding the contract method 0xc8f8decb.
//
// Solidity: function CONTRACT_TRANSFER_FROM_CAPITALIZATION() view returns(uint256)
func (_OtcImpl *OtcImplSession) CONTRACTTRANSFERFROMCAPITALIZATION() (*big.Int, error) {
	return _OtcImpl.Contract.CONTRACTTRANSFERFROMCAPITALIZATION(&_OtcImpl.CallOpts)
}

// CONTRACTTRANSFERFROMCAPITALIZATION is a free data retrieval call binding the contract method 0xc8f8decb.
//
// Solidity: function CONTRACT_TRANSFER_FROM_CAPITALIZATION() view returns(uint256)
func (_OtcImpl *OtcImplCallerSession) CONTRACTTRANSFERFROMCAPITALIZATION() (*big.Int, error) {
	return _OtcImpl.Contract.CONTRACTTRANSFERFROMCAPITALIZATION(&_OtcImpl.CallOpts)
}

// PERMISSIONACCESSRESTORATION is a free data retrieval call binding the contract method 0x5c3053bf.
//
// Solidity: function PERMISSION_ACCESS_RESTORATION() view returns(uint256)
func (_OtcImpl *OtcImplCaller) PERMISSIONACCESSRESTORATION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OtcImpl.contract.Call(opts, out, "PERMISSION_ACCESS_RESTORATION")
	return *ret0, err
}

// PERMISSIONACCESSRESTORATION is a free data retrieval call binding the contract method 0x5c3053bf.
//
// Solidity: function PERMISSION_ACCESS_RESTORATION() view returns(uint256)
func (_OtcImpl *OtcImplSession) PERMISSIONACCESSRESTORATION() (*big.Int, error) {
	return _OtcImpl.Contract.PERMISSIONACCESSRESTORATION(&_OtcImpl.CallOpts)
}

// PERMISSIONACCESSRESTORATION is a free data retrieval call binding the contract method 0x5c3053bf.
//
// Solidity: function PERMISSION_ACCESS_RESTORATION() view returns(uint256)
func (_OtcImpl *OtcImplCallerSession) PERMISSIONACCESSRESTORATION() (*big.Int, error) {
	return _OtcImpl.Contract.PERMISSIONACCESSRESTORATION(&_OtcImpl.CallOpts)
}

// PERMISSIONADMINSTORAGE is a free data retrieval call binding the contract method 0x0c9fc0db.
//
// Solidity: function PERMISSION_ADMIN_STORAGE() view returns(uint256)
func (_OtcImpl *OtcImplCaller) PERMISSIONADMINSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OtcImpl.contract.Call(opts, out, "PERMISSION_ADMIN_STORAGE")
	return *ret0, err
}

// PERMISSIONADMINSTORAGE is a free data retrieval call binding the contract method 0x0c9fc0db.
//
// Solidity: function PERMISSION_ADMIN_STORAGE() view returns(uint256)
func (_OtcImpl *OtcImplSession) PERMISSIONADMINSTORAGE() (*big.Int, error) {
	return _OtcImpl.Contract.PERMISSIONADMINSTORAGE(&_OtcImpl.CallOpts)
}

// PERMISSIONADMINSTORAGE is a free data retrieval call binding the contract method 0x0c9fc0db.
//
// Solidity: function PERMISSION_ADMIN_STORAGE() view returns(uint256)
func (_OtcImpl *OtcImplCallerSession) PERMISSIONADMINSTORAGE() (*big.Int, error) {
	return _OtcImpl.Contract.PERMISSIONADMINSTORAGE(&_OtcImpl.CallOpts)
}

// PERMISSIONCONTRACTSSTORAGE is a free data retrieval call binding the contract method 0x1516e83b.
//
// Solidity: function PERMISSION_CONTRACTS_STORAGE() view returns(uint256)
func (_OtcImpl *OtcImplCaller) PERMISSIONCONTRACTSSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OtcImpl.contract.Call(opts, out, "PERMISSION_CONTRACTS_STORAGE")
	return *ret0, err
}

// PERMISSIONCONTRACTSSTORAGE is a free data retrieval call binding the contract method 0x1516e83b.
//
// Solidity: function PERMISSION_CONTRACTS_STORAGE() view returns(uint256)
func (_OtcImpl *OtcImplSession) PERMISSIONCONTRACTSSTORAGE() (*big.Int, error) {
	return _OtcImpl.Contract.PERMISSIONCONTRACTSSTORAGE(&_OtcImpl.CallOpts)
}

// PERMISSIONCONTRACTSSTORAGE is a free data retrieval call binding the contract method 0x1516e83b.
//
// Solidity: function PERMISSION_CONTRACTS_STORAGE() view returns(uint256)
func (_OtcImpl *OtcImplCallerSession) PERMISSIONCONTRACTSSTORAGE() (*big.Int, error) {
	return _OtcImpl.Contract.PERMISSIONCONTRACTSSTORAGE(&_OtcImpl.CallOpts)
}

// PERMISSIONCONTRACTDEPLOY is a free data retrieval call binding the contract method 0x257a4a04.
//
// Solidity: function PERMISSION_CONTRACT_DEPLOY() view returns(uint256)
func (_OtcImpl *OtcImplCaller) PERMISSIONCONTRACTDEPLOY(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OtcImpl.contract.Call(opts, out, "PERMISSION_CONTRACT_DEPLOY")
	return *ret0, err
}

// PERMISSIONCONTRACTDEPLOY is a free data retrieval call binding the contract method 0x257a4a04.
//
// Solidity: function PERMISSION_CONTRACT_DEPLOY() view returns(uint256)
func (_OtcImpl *OtcImplSession) PERMISSIONCONTRACTDEPLOY() (*big.Int, error) {
	return _OtcImpl.Contract.PERMISSIONCONTRACTDEPLOY(&_OtcImpl.CallOpts)
}

// PERMISSIONCONTRACTDEPLOY is a free data retrieval call binding the contract method 0x257a4a04.
//
// Solidity: function PERMISSION_CONTRACT_DEPLOY() view returns(uint256)
func (_OtcImpl *OtcImplCallerSession) PERMISSIONCONTRACTDEPLOY() (*big.Int, error) {
	return _OtcImpl.Contract.PERMISSIONCONTRACTDEPLOY(&_OtcImpl.CallOpts)
}

// PERMISSIONEMISSION is a free data retrieval call binding the contract method 0x4ccd92d1.
//
// Solidity: function PERMISSION_EMISSION() view returns(uint256)
func (_OtcImpl *OtcImplCaller) PERMISSIONEMISSION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OtcImpl.contract.Call(opts, out, "PERMISSION_EMISSION")
	return *ret0, err
}

// PERMISSIONEMISSION is a free data retrieval call binding the contract method 0x4ccd92d1.
//
// Solidity: function PERMISSION_EMISSION() view returns(uint256)
func (_OtcImpl *OtcImplSession) PERMISSIONEMISSION() (*big.Int, error) {
	return _OtcImpl.Contract.PERMISSIONEMISSION(&_OtcImpl.CallOpts)
}

// PERMISSIONEMISSION is a free data retrieval call binding the contract method 0x4ccd92d1.
//
// Solidity: function PERMISSION_EMISSION() view returns(uint256)
func (_OtcImpl *OtcImplCallerSession) PERMISSIONEMISSION() (*big.Int, error) {
	return _OtcImpl.Contract.PERMISSIONEMISSION(&_OtcImpl.CallOpts)
}

// PERMISSIONFEESSTORAGE is a free data retrieval call binding the contract method 0xe837eb87.
//
// Solidity: function PERMISSION_FEES_STORAGE() view returns(uint256)
func (_OtcImpl *OtcImplCaller) PERMISSIONFEESSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OtcImpl.contract.Call(opts, out, "PERMISSION_FEES_STORAGE")
	return *ret0, err
}

// PERMISSIONFEESSTORAGE is a free data retrieval call binding the contract method 0xe837eb87.
//
// Solidity: function PERMISSION_FEES_STORAGE() view returns(uint256)
func (_OtcImpl *OtcImplSession) PERMISSIONFEESSTORAGE() (*big.Int, error) {
	return _OtcImpl.Contract.PERMISSIONFEESSTORAGE(&_OtcImpl.CallOpts)
}

// PERMISSIONFEESSTORAGE is a free data retrieval call binding the contract method 0xe837eb87.
//
// Solidity: function PERMISSION_FEES_STORAGE() view returns(uint256)
func (_OtcImpl *OtcImplCallerSession) PERMISSIONFEESSTORAGE() (*big.Int, error) {
	return _OtcImpl.Contract.PERMISSIONFEESSTORAGE(&_OtcImpl.CallOpts)
}

// PERMISSIONMULTISIGSTORAGE is a free data retrieval call binding the contract method 0xf1bb7aad.
//
// Solidity: function PERMISSION_MULTISIG_STORAGE() view returns(uint256)
func (_OtcImpl *OtcImplCaller) PERMISSIONMULTISIGSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OtcImpl.contract.Call(opts, out, "PERMISSION_MULTISIG_STORAGE")
	return *ret0, err
}

// PERMISSIONMULTISIGSTORAGE is a free data retrieval call binding the contract method 0xf1bb7aad.
//
// Solidity: function PERMISSION_MULTISIG_STORAGE() view returns(uint256)
func (_OtcImpl *OtcImplSession) PERMISSIONMULTISIGSTORAGE() (*big.Int, error) {
	return _OtcImpl.Contract.PERMISSIONMULTISIGSTORAGE(&_OtcImpl.CallOpts)
}

// PERMISSIONMULTISIGSTORAGE is a free data retrieval call binding the contract method 0xf1bb7aad.
//
// Solidity: function PERMISSION_MULTISIG_STORAGE() view returns(uint256)
func (_OtcImpl *OtcImplCallerSession) PERMISSIONMULTISIGSTORAGE() (*big.Int, error) {
	return _OtcImpl.Contract.PERMISSIONMULTISIGSTORAGE(&_OtcImpl.CallOpts)
}

// PERMISSIONROLESSTORAGE is a free data retrieval call binding the contract method 0x12a16032.
//
// Solidity: function PERMISSION_ROLES_STORAGE() view returns(uint256)
func (_OtcImpl *OtcImplCaller) PERMISSIONROLESSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OtcImpl.contract.Call(opts, out, "PERMISSION_ROLES_STORAGE")
	return *ret0, err
}

// PERMISSIONROLESSTORAGE is a free data retrieval call binding the contract method 0x12a16032.
//
// Solidity: function PERMISSION_ROLES_STORAGE() view returns(uint256)
func (_OtcImpl *OtcImplSession) PERMISSIONROLESSTORAGE() (*big.Int, error) {
	return _OtcImpl.Contract.PERMISSIONROLESSTORAGE(&_OtcImpl.CallOpts)
}

// PERMISSIONROLESSTORAGE is a free data retrieval call binding the contract method 0x12a16032.
//
// Solidity: function PERMISSION_ROLES_STORAGE() view returns(uint256)
func (_OtcImpl *OtcImplCallerSession) PERMISSIONROLESSTORAGE() (*big.Int, error) {
	return _OtcImpl.Contract.PERMISSIONROLESSTORAGE(&_OtcImpl.CallOpts)
}

// ROLEACCESSVERIFIER is a free data retrieval call binding the contract method 0xc5405f92.
//
// Solidity: function ROLE_ACCESS_VERIFIER() view returns(uint256)
func (_OtcImpl *OtcImplCaller) ROLEACCESSVERIFIER(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OtcImpl.contract.Call(opts, out, "ROLE_ACCESS_VERIFIER")
	return *ret0, err
}

// ROLEACCESSVERIFIER is a free data retrieval call binding the contract method 0xc5405f92.
//
// Solidity: function ROLE_ACCESS_VERIFIER() view returns(uint256)
func (_OtcImpl *OtcImplSession) ROLEACCESSVERIFIER() (*big.Int, error) {
	return _OtcImpl.Contract.ROLEACCESSVERIFIER(&_OtcImpl.CallOpts)
}

// ROLEACCESSVERIFIER is a free data retrieval call binding the contract method 0xc5405f92.
//
// Solidity: function ROLE_ACCESS_VERIFIER() view returns(uint256)
func (_OtcImpl *OtcImplCallerSession) ROLEACCESSVERIFIER() (*big.Int, error) {
	return _OtcImpl.Contract.ROLEACCESSVERIFIER(&_OtcImpl.CallOpts)
}

// ROLEBITBONISSUEVERIFIER is a free data retrieval call binding the contract method 0xf0a012a6.
//
// Solidity: function ROLE_BITBON_ISSUE_VERIFIER() view returns(uint256)
func (_OtcImpl *OtcImplCaller) ROLEBITBONISSUEVERIFIER(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OtcImpl.contract.Call(opts, out, "ROLE_BITBON_ISSUE_VERIFIER")
	return *ret0, err
}

// ROLEBITBONISSUEVERIFIER is a free data retrieval call binding the contract method 0xf0a012a6.
//
// Solidity: function ROLE_BITBON_ISSUE_VERIFIER() view returns(uint256)
func (_OtcImpl *OtcImplSession) ROLEBITBONISSUEVERIFIER() (*big.Int, error) {
	return _OtcImpl.Contract.ROLEBITBONISSUEVERIFIER(&_OtcImpl.CallOpts)
}

// ROLEBITBONISSUEVERIFIER is a free data retrieval call binding the contract method 0xf0a012a6.
//
// Solidity: function ROLE_BITBON_ISSUE_VERIFIER() view returns(uint256)
func (_OtcImpl *OtcImplCallerSession) ROLEBITBONISSUEVERIFIER() (*big.Int, error) {
	return _OtcImpl.Contract.ROLEBITBONISSUEVERIFIER(&_OtcImpl.CallOpts)
}

// ROLECOMMISSIONVERIFIER is a free data retrieval call binding the contract method 0x8ab3de8c.
//
// Solidity: function ROLE_COMMISSION_VERIFIER() view returns(uint256)
func (_OtcImpl *OtcImplCaller) ROLECOMMISSIONVERIFIER(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OtcImpl.contract.Call(opts, out, "ROLE_COMMISSION_VERIFIER")
	return *ret0, err
}

// ROLECOMMISSIONVERIFIER is a free data retrieval call binding the contract method 0x8ab3de8c.
//
// Solidity: function ROLE_COMMISSION_VERIFIER() view returns(uint256)
func (_OtcImpl *OtcImplSession) ROLECOMMISSIONVERIFIER() (*big.Int, error) {
	return _OtcImpl.Contract.ROLECOMMISSIONVERIFIER(&_OtcImpl.CallOpts)
}

// ROLECOMMISSIONVERIFIER is a free data retrieval call binding the contract method 0x8ab3de8c.
//
// Solidity: function ROLE_COMMISSION_VERIFIER() view returns(uint256)
func (_OtcImpl *OtcImplCallerSession) ROLECOMMISSIONVERIFIER() (*big.Int, error) {
	return _OtcImpl.Contract.ROLECOMMISSIONVERIFIER(&_OtcImpl.CallOpts)
}

// ROLEDEPLOYADMIN is a free data retrieval call binding the contract method 0x667bf6e5.
//
// Solidity: function ROLE_DEPLOY_ADMIN() view returns(uint256)
func (_OtcImpl *OtcImplCaller) ROLEDEPLOYADMIN(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OtcImpl.contract.Call(opts, out, "ROLE_DEPLOY_ADMIN")
	return *ret0, err
}

// ROLEDEPLOYADMIN is a free data retrieval call binding the contract method 0x667bf6e5.
//
// Solidity: function ROLE_DEPLOY_ADMIN() view returns(uint256)
func (_OtcImpl *OtcImplSession) ROLEDEPLOYADMIN() (*big.Int, error) {
	return _OtcImpl.Contract.ROLEDEPLOYADMIN(&_OtcImpl.CallOpts)
}

// ROLEDEPLOYADMIN is a free data retrieval call binding the contract method 0x667bf6e5.
//
// Solidity: function ROLE_DEPLOY_ADMIN() view returns(uint256)
func (_OtcImpl *OtcImplCallerSession) ROLEDEPLOYADMIN() (*big.Int, error) {
	return _OtcImpl.Contract.ROLEDEPLOYADMIN(&_OtcImpl.CallOpts)
}

// ROLEPERMISSIONADMIN is a free data retrieval call binding the contract method 0x1b239996.
//
// Solidity: function ROLE_PERMISSION_ADMIN() view returns(uint256)
func (_OtcImpl *OtcImplCaller) ROLEPERMISSIONADMIN(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OtcImpl.contract.Call(opts, out, "ROLE_PERMISSION_ADMIN")
	return *ret0, err
}

// ROLEPERMISSIONADMIN is a free data retrieval call binding the contract method 0x1b239996.
//
// Solidity: function ROLE_PERMISSION_ADMIN() view returns(uint256)
func (_OtcImpl *OtcImplSession) ROLEPERMISSIONADMIN() (*big.Int, error) {
	return _OtcImpl.Contract.ROLEPERMISSIONADMIN(&_OtcImpl.CallOpts)
}

// ROLEPERMISSIONADMIN is a free data retrieval call binding the contract method 0x1b239996.
//
// Solidity: function ROLE_PERMISSION_ADMIN() view returns(uint256)
func (_OtcImpl *OtcImplCallerSession) ROLEPERMISSIONADMIN() (*big.Int, error) {
	return _OtcImpl.Contract.ROLEPERMISSIONADMIN(&_OtcImpl.CallOpts)
}

// ContractStorage is a free data retrieval call binding the contract method 0x0549b27e.
//
// Solidity: function contractStorage() view returns(address)
func (_OtcImpl *OtcImplCaller) ContractStorage(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _OtcImpl.contract.Call(opts, out, "contractStorage")
	return *ret0, err
}

// ContractStorage is a free data retrieval call binding the contract method 0x0549b27e.
//
// Solidity: function contractStorage() view returns(address)
func (_OtcImpl *OtcImplSession) ContractStorage() (common.Address, error) {
	return _OtcImpl.Contract.ContractStorage(&_OtcImpl.CallOpts)
}

// ContractStorage is a free data retrieval call binding the contract method 0x0549b27e.
//
// Solidity: function contractStorage() view returns(address)
func (_OtcImpl *OtcImplCallerSession) ContractStorage() (common.Address, error) {
	return _OtcImpl.Contract.ContractStorage(&_OtcImpl.CallOpts)
}

// GetThisContractIndex is a free data retrieval call binding the contract method 0x70c6d367.
//
// Solidity: function getThisContractIndex() pure returns(uint256)
func (_OtcImpl *OtcImplCaller) GetThisContractIndex(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OtcImpl.contract.Call(opts, out, "getThisContractIndex")
	return *ret0, err
}

// GetThisContractIndex is a free data retrieval call binding the contract method 0x70c6d367.
//
// Solidity: function getThisContractIndex() pure returns(uint256)
func (_OtcImpl *OtcImplSession) GetThisContractIndex() (*big.Int, error) {
	return _OtcImpl.Contract.GetThisContractIndex(&_OtcImpl.CallOpts)
}

// GetThisContractIndex is a free data retrieval call binding the contract method 0x70c6d367.
//
// Solidity: function getThisContractIndex() pure returns(uint256)
func (_OtcImpl *OtcImplCallerSession) GetThisContractIndex() (*big.Int, error) {
	return _OtcImpl.Contract.GetThisContractIndex(&_OtcImpl.CallOpts)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint8)
func (_OtcImpl *OtcImplCaller) State(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _OtcImpl.contract.Call(opts, out, "state")
	return *ret0, err
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint8)
func (_OtcImpl *OtcImplSession) State() (uint8, error) {
	return _OtcImpl.Contract.State(&_OtcImpl.CallOpts)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint8)
func (_OtcImpl *OtcImplCallerSession) State() (uint8, error) {
	return _OtcImpl.Contract.State(&_OtcImpl.CallOpts)
}

// CancelBitbonOrder is a paid mutator transaction binding the contract method 0x3ac822c0.
//
// Solidity: function cancelBitbonOrder(bytes32 offerHash, uint256 expectedAvailableAmount) returns()
func (_OtcImpl *OtcImplTransactor) CancelBitbonOrder(opts *bind.TransactOpts, offerHash [32]byte, expectedAvailableAmount *big.Int) (*types.Transaction, error) {
	return _OtcImpl.contract.Transact(opts, "cancelBitbonOrder", offerHash, expectedAvailableAmount)
}

// CancelBitbonOrder is a paid mutator transaction binding the contract method 0x3ac822c0.
//
// Solidity: function cancelBitbonOrder(bytes32 offerHash, uint256 expectedAvailableAmount) returns()
func (_OtcImpl *OtcImplSession) CancelBitbonOrder(offerHash [32]byte, expectedAvailableAmount *big.Int) (*types.Transaction, error) {
	return _OtcImpl.Contract.CancelBitbonOrder(&_OtcImpl.TransactOpts, offerHash, expectedAvailableAmount)
}

// CancelBitbonOrder is a paid mutator transaction binding the contract method 0x3ac822c0.
//
// Solidity: function cancelBitbonOrder(bytes32 offerHash, uint256 expectedAvailableAmount) returns()
func (_OtcImpl *OtcImplTransactorSession) CancelBitbonOrder(offerHash [32]byte, expectedAvailableAmount *big.Int) (*types.Transaction, error) {
	return _OtcImpl.Contract.CancelBitbonOrder(&_OtcImpl.TransactOpts, offerHash, expectedAvailableAmount)
}

// CancelMsbonOrder is a paid mutator transaction binding the contract method 0x78b80a87.
//
// Solidity: function cancelMsbonOrder(bytes32 offerHash, uint256 expectedAvailableAmount) returns()
func (_OtcImpl *OtcImplTransactor) CancelMsbonOrder(opts *bind.TransactOpts, offerHash [32]byte, expectedAvailableAmount *big.Int) (*types.Transaction, error) {
	return _OtcImpl.contract.Transact(opts, "cancelMsbonOrder", offerHash, expectedAvailableAmount)
}

// CancelMsbonOrder is a paid mutator transaction binding the contract method 0x78b80a87.
//
// Solidity: function cancelMsbonOrder(bytes32 offerHash, uint256 expectedAvailableAmount) returns()
func (_OtcImpl *OtcImplSession) CancelMsbonOrder(offerHash [32]byte, expectedAvailableAmount *big.Int) (*types.Transaction, error) {
	return _OtcImpl.Contract.CancelMsbonOrder(&_OtcImpl.TransactOpts, offerHash, expectedAvailableAmount)
}

// CancelMsbonOrder is a paid mutator transaction binding the contract method 0x78b80a87.
//
// Solidity: function cancelMsbonOrder(bytes32 offerHash, uint256 expectedAvailableAmount) returns()
func (_OtcImpl *OtcImplTransactorSession) CancelMsbonOrder(offerHash [32]byte, expectedAvailableAmount *big.Int) (*types.Transaction, error) {
	return _OtcImpl.Contract.CancelMsbonOrder(&_OtcImpl.TransactOpts, offerHash, expectedAvailableAmount)
}

// FillBitbonOrder is a paid mutator transaction binding the contract method 0x79b1ae4b.
//
// Solidity: function fillBitbonOrder(address fillAssets, address wantAssets, bytes32 offerHash, uint256 amountToTake) returns()
func (_OtcImpl *OtcImplTransactor) FillBitbonOrder(opts *bind.TransactOpts, fillAssets common.Address, wantAssets common.Address, offerHash [32]byte, amountToTake *big.Int) (*types.Transaction, error) {
	return _OtcImpl.contract.Transact(opts, "fillBitbonOrder", fillAssets, wantAssets, offerHash, amountToTake)
}

// FillBitbonOrder is a paid mutator transaction binding the contract method 0x79b1ae4b.
//
// Solidity: function fillBitbonOrder(address fillAssets, address wantAssets, bytes32 offerHash, uint256 amountToTake) returns()
func (_OtcImpl *OtcImplSession) FillBitbonOrder(fillAssets common.Address, wantAssets common.Address, offerHash [32]byte, amountToTake *big.Int) (*types.Transaction, error) {
	return _OtcImpl.Contract.FillBitbonOrder(&_OtcImpl.TransactOpts, fillAssets, wantAssets, offerHash, amountToTake)
}

// FillBitbonOrder is a paid mutator transaction binding the contract method 0x79b1ae4b.
//
// Solidity: function fillBitbonOrder(address fillAssets, address wantAssets, bytes32 offerHash, uint256 amountToTake) returns()
func (_OtcImpl *OtcImplTransactorSession) FillBitbonOrder(fillAssets common.Address, wantAssets common.Address, offerHash [32]byte, amountToTake *big.Int) (*types.Transaction, error) {
	return _OtcImpl.Contract.FillBitbonOrder(&_OtcImpl.TransactOpts, fillAssets, wantAssets, offerHash, amountToTake)
}

// FillMsbonOrder is a paid mutator transaction binding the contract method 0x41b8d73e.
//
// Solidity: function fillMsbonOrder(address fillAssets, address wantAssets, bytes32 offerHash, uint256 amountToTake) returns()
func (_OtcImpl *OtcImplTransactor) FillMsbonOrder(opts *bind.TransactOpts, fillAssets common.Address, wantAssets common.Address, offerHash [32]byte, amountToTake *big.Int) (*types.Transaction, error) {
	return _OtcImpl.contract.Transact(opts, "fillMsbonOrder", fillAssets, wantAssets, offerHash, amountToTake)
}

// FillMsbonOrder is a paid mutator transaction binding the contract method 0x41b8d73e.
//
// Solidity: function fillMsbonOrder(address fillAssets, address wantAssets, bytes32 offerHash, uint256 amountToTake) returns()
func (_OtcImpl *OtcImplSession) FillMsbonOrder(fillAssets common.Address, wantAssets common.Address, offerHash [32]byte, amountToTake *big.Int) (*types.Transaction, error) {
	return _OtcImpl.Contract.FillMsbonOrder(&_OtcImpl.TransactOpts, fillAssets, wantAssets, offerHash, amountToTake)
}

// FillMsbonOrder is a paid mutator transaction binding the contract method 0x41b8d73e.
//
// Solidity: function fillMsbonOrder(address fillAssets, address wantAssets, bytes32 offerHash, uint256 amountToTake) returns()
func (_OtcImpl *OtcImplTransactorSession) FillMsbonOrder(fillAssets common.Address, wantAssets common.Address, offerHash [32]byte, amountToTake *big.Int) (*types.Transaction, error) {
	return _OtcImpl.Contract.FillMsbonOrder(&_OtcImpl.TransactOpts, fillAssets, wantAssets, offerHash, amountToTake)
}

// MakeBatchOrders is a paid mutator transaction binding the contract method 0x3a06df97.
//
// Solidity: function makeBatchOrders(address[] offerAssets, address[] wantAssets, uint256[] offerAmount, bool[] tradeType, uint64 nonce) returns()
func (_OtcImpl *OtcImplTransactor) MakeBatchOrders(opts *bind.TransactOpts, offerAssets []common.Address, wantAssets []common.Address, offerAmount []*big.Int, tradeType []bool, nonce uint64) (*types.Transaction, error) {
	return _OtcImpl.contract.Transact(opts, "makeBatchOrders", offerAssets, wantAssets, offerAmount, tradeType, nonce)
}

// MakeBatchOrders is a paid mutator transaction binding the contract method 0x3a06df97.
//
// Solidity: function makeBatchOrders(address[] offerAssets, address[] wantAssets, uint256[] offerAmount, bool[] tradeType, uint64 nonce) returns()
func (_OtcImpl *OtcImplSession) MakeBatchOrders(offerAssets []common.Address, wantAssets []common.Address, offerAmount []*big.Int, tradeType []bool, nonce uint64) (*types.Transaction, error) {
	return _OtcImpl.Contract.MakeBatchOrders(&_OtcImpl.TransactOpts, offerAssets, wantAssets, offerAmount, tradeType, nonce)
}

// MakeBatchOrders is a paid mutator transaction binding the contract method 0x3a06df97.
//
// Solidity: function makeBatchOrders(address[] offerAssets, address[] wantAssets, uint256[] offerAmount, bool[] tradeType, uint64 nonce) returns()
func (_OtcImpl *OtcImplTransactorSession) MakeBatchOrders(offerAssets []common.Address, wantAssets []common.Address, offerAmount []*big.Int, tradeType []bool, nonce uint64) (*types.Transaction, error) {
	return _OtcImpl.Contract.MakeBatchOrders(&_OtcImpl.TransactOpts, offerAssets, wantAssets, offerAmount, tradeType, nonce)
}

// MakeBitbonOrder is a paid mutator transaction binding the contract method 0xac882251.
//
// Solidity: function makeBitbonOrder(address offerAssets, address wantAssets, uint256 offerAmount, uint64 nonce) returns()
func (_OtcImpl *OtcImplTransactor) MakeBitbonOrder(opts *bind.TransactOpts, offerAssets common.Address, wantAssets common.Address, offerAmount *big.Int, nonce uint64) (*types.Transaction, error) {
	return _OtcImpl.contract.Transact(opts, "makeBitbonOrder", offerAssets, wantAssets, offerAmount, nonce)
}

// MakeBitbonOrder is a paid mutator transaction binding the contract method 0xac882251.
//
// Solidity: function makeBitbonOrder(address offerAssets, address wantAssets, uint256 offerAmount, uint64 nonce) returns()
func (_OtcImpl *OtcImplSession) MakeBitbonOrder(offerAssets common.Address, wantAssets common.Address, offerAmount *big.Int, nonce uint64) (*types.Transaction, error) {
	return _OtcImpl.Contract.MakeBitbonOrder(&_OtcImpl.TransactOpts, offerAssets, wantAssets, offerAmount, nonce)
}

// MakeBitbonOrder is a paid mutator transaction binding the contract method 0xac882251.
//
// Solidity: function makeBitbonOrder(address offerAssets, address wantAssets, uint256 offerAmount, uint64 nonce) returns()
func (_OtcImpl *OtcImplTransactorSession) MakeBitbonOrder(offerAssets common.Address, wantAssets common.Address, offerAmount *big.Int, nonce uint64) (*types.Transaction, error) {
	return _OtcImpl.Contract.MakeBitbonOrder(&_OtcImpl.TransactOpts, offerAssets, wantAssets, offerAmount, nonce)
}

// MakeMsbonOrder is a paid mutator transaction binding the contract method 0xc4b1709a.
//
// Solidity: function makeMsbonOrder(address offerAssets, address wantAssets, uint256 offerAmount, uint64 nonce) returns()
func (_OtcImpl *OtcImplTransactor) MakeMsbonOrder(opts *bind.TransactOpts, offerAssets common.Address, wantAssets common.Address, offerAmount *big.Int, nonce uint64) (*types.Transaction, error) {
	return _OtcImpl.contract.Transact(opts, "makeMsbonOrder", offerAssets, wantAssets, offerAmount, nonce)
}

// MakeMsbonOrder is a paid mutator transaction binding the contract method 0xc4b1709a.
//
// Solidity: function makeMsbonOrder(address offerAssets, address wantAssets, uint256 offerAmount, uint64 nonce) returns()
func (_OtcImpl *OtcImplSession) MakeMsbonOrder(offerAssets common.Address, wantAssets common.Address, offerAmount *big.Int, nonce uint64) (*types.Transaction, error) {
	return _OtcImpl.Contract.MakeMsbonOrder(&_OtcImpl.TransactOpts, offerAssets, wantAssets, offerAmount, nonce)
}

// MakeMsbonOrder is a paid mutator transaction binding the contract method 0xc4b1709a.
//
// Solidity: function makeMsbonOrder(address offerAssets, address wantAssets, uint256 offerAmount, uint64 nonce) returns()
func (_OtcImpl *OtcImplTransactorSession) MakeMsbonOrder(offerAssets common.Address, wantAssets common.Address, offerAmount *big.Int, nonce uint64) (*types.Transaction, error) {
	return _OtcImpl.Contract.MakeMsbonOrder(&_OtcImpl.TransactOpts, offerAssets, wantAssets, offerAmount, nonce)
}

// SetBitbonPrice is a paid mutator transaction binding the contract method 0x6940d394.
//
// Solidity: function setBitbonPrice(uint256 offerAmount, uint256 wantAmount) returns()
func (_OtcImpl *OtcImplTransactor) SetBitbonPrice(opts *bind.TransactOpts, offerAmount *big.Int, wantAmount *big.Int) (*types.Transaction, error) {
	return _OtcImpl.contract.Transact(opts, "setBitbonPrice", offerAmount, wantAmount)
}

// SetBitbonPrice is a paid mutator transaction binding the contract method 0x6940d394.
//
// Solidity: function setBitbonPrice(uint256 offerAmount, uint256 wantAmount) returns()
func (_OtcImpl *OtcImplSession) SetBitbonPrice(offerAmount *big.Int, wantAmount *big.Int) (*types.Transaction, error) {
	return _OtcImpl.Contract.SetBitbonPrice(&_OtcImpl.TransactOpts, offerAmount, wantAmount)
}

// SetBitbonPrice is a paid mutator transaction binding the contract method 0x6940d394.
//
// Solidity: function setBitbonPrice(uint256 offerAmount, uint256 wantAmount) returns()
func (_OtcImpl *OtcImplTransactorSession) SetBitbonPrice(offerAmount *big.Int, wantAmount *big.Int) (*types.Transaction, error) {
	return _OtcImpl.Contract.SetBitbonPrice(&_OtcImpl.TransactOpts, offerAmount, wantAmount)
}

// SetMsbonPrice is a paid mutator transaction binding the contract method 0x0b8b3201.
//
// Solidity: function setMsbonPrice(uint256 offerAmount, uint256 wantAmount) returns()
func (_OtcImpl *OtcImplTransactor) SetMsbonPrice(opts *bind.TransactOpts, offerAmount *big.Int, wantAmount *big.Int) (*types.Transaction, error) {
	return _OtcImpl.contract.Transact(opts, "setMsbonPrice", offerAmount, wantAmount)
}

// SetMsbonPrice is a paid mutator transaction binding the contract method 0x0b8b3201.
//
// Solidity: function setMsbonPrice(uint256 offerAmount, uint256 wantAmount) returns()
func (_OtcImpl *OtcImplSession) SetMsbonPrice(offerAmount *big.Int, wantAmount *big.Int) (*types.Transaction, error) {
	return _OtcImpl.Contract.SetMsbonPrice(&_OtcImpl.TransactOpts, offerAmount, wantAmount)
}

// SetMsbonPrice is a paid mutator transaction binding the contract method 0x0b8b3201.
//
// Solidity: function setMsbonPrice(uint256 offerAmount, uint256 wantAmount) returns()
func (_OtcImpl *OtcImplTransactorSession) SetMsbonPrice(offerAmount *big.Int, wantAmount *big.Int) (*types.Transaction, error) {
	return _OtcImpl.Contract.SetMsbonPrice(&_OtcImpl.TransactOpts, offerAmount, wantAmount)
}

// SetState is a paid mutator transaction binding the contract method 0x56de96db.
//
// Solidity: function setState(uint8 newState) returns()
func (_OtcImpl *OtcImplTransactor) SetState(opts *bind.TransactOpts, newState uint8) (*types.Transaction, error) {
	return _OtcImpl.contract.Transact(opts, "setState", newState)
}

// SetState is a paid mutator transaction binding the contract method 0x56de96db.
//
// Solidity: function setState(uint8 newState) returns()
func (_OtcImpl *OtcImplSession) SetState(newState uint8) (*types.Transaction, error) {
	return _OtcImpl.Contract.SetState(&_OtcImpl.TransactOpts, newState)
}

// SetState is a paid mutator transaction binding the contract method 0x56de96db.
//
// Solidity: function setState(uint8 newState) returns()
func (_OtcImpl *OtcImplTransactorSession) SetState(newState uint8) (*types.Transaction, error) {
	return _OtcImpl.Contract.SetState(&_OtcImpl.TransactOpts, newState)
}

// OtcImplBalanceLockedIterator is returned from FilterBalanceLocked and is used to iterate over the raw logs and unpacked data for BalanceLocked events raised by the OtcImpl contract.
type OtcImplBalanceLockedIterator struct {
	Event *OtcImplBalanceLocked // Event containing the contract specifics and raw log

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
func (it *OtcImplBalanceLockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OtcImplBalanceLocked)
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
		it.Event = new(OtcImplBalanceLocked)
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
func (it *OtcImplBalanceLockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OtcImplBalanceLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OtcImplBalanceLocked represents a BalanceLocked event raised by the OtcImpl contract.
type OtcImplBalanceLocked struct {
	Assetbox       common.Address
	BalanceAviable *big.Int
	BalanceLocked  *big.Int
	AssetboxType   *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBalanceLocked is a free log retrieval operation binding the contract event 0x89f85a4bd38f70943757e43dedd843409e565220cb52ba80fc297d1246b3b9bb.
//
// Solidity: event BalanceLocked(address assetbox, uint256 balanceAviable, uint256 balanceLocked, uint256 assetboxType)
func (_OtcImpl *OtcImplFilterer) FilterBalanceLocked(opts *bind.FilterOpts) (*OtcImplBalanceLockedIterator, error) {

	logs, sub, err := _OtcImpl.contract.FilterLogs(opts, "BalanceLocked")
	if err != nil {
		return nil, err
	}
	return &OtcImplBalanceLockedIterator{contract: _OtcImpl.contract, event: "BalanceLocked", logs: logs, sub: sub}, nil
}

// WatchBalanceLocked is a free log subscription operation binding the contract event 0x89f85a4bd38f70943757e43dedd843409e565220cb52ba80fc297d1246b3b9bb.
//
// Solidity: event BalanceLocked(address assetbox, uint256 balanceAviable, uint256 balanceLocked, uint256 assetboxType)
func (_OtcImpl *OtcImplFilterer) WatchBalanceLocked(opts *bind.WatchOpts, sink chan<- *OtcImplBalanceLocked) (event.Subscription, error) {

	logs, sub, err := _OtcImpl.contract.WatchLogs(opts, "BalanceLocked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OtcImplBalanceLocked)
				if err := _OtcImpl.contract.UnpackLog(event, "BalanceLocked", log); err != nil {
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

// ParseBalanceLocked is a log parse operation binding the contract event 0x89f85a4bd38f70943757e43dedd843409e565220cb52ba80fc297d1246b3b9bb.
//
// Solidity: event BalanceLocked(address assetbox, uint256 balanceAviable, uint256 balanceLocked, uint256 assetboxType)
func (_OtcImpl *OtcImplFilterer) ParseBalanceLocked(log types.Log) (*OtcImplBalanceLocked, error) {
	event := new(OtcImplBalanceLocked)
	if err := _OtcImpl.contract.UnpackLog(event, "BalanceLocked", log); err != nil {
		return nil, err
	}
	return event, nil
}

// OtcImplBalanceUnLockedIterator is returned from FilterBalanceUnLocked and is used to iterate over the raw logs and unpacked data for BalanceUnLocked events raised by the OtcImpl contract.
type OtcImplBalanceUnLockedIterator struct {
	Event *OtcImplBalanceUnLocked // Event containing the contract specifics and raw log

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
func (it *OtcImplBalanceUnLockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OtcImplBalanceUnLocked)
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
		it.Event = new(OtcImplBalanceUnLocked)
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
func (it *OtcImplBalanceUnLockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OtcImplBalanceUnLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OtcImplBalanceUnLocked represents a BalanceUnLocked event raised by the OtcImpl contract.
type OtcImplBalanceUnLocked struct {
	Assetbox       common.Address
	BalanceAviable *big.Int
	BalanceLocked  *big.Int
	AssetboxType   *big.Int
	Amount         *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBalanceUnLocked is a free log retrieval operation binding the contract event 0x486bd0991e37e48fc69ea3be6a0b11058d29ef8a89bbc475b139a00a681b4c9b.
//
// Solidity: event BalanceUnLocked(address assetbox, uint256 balanceAviable, uint256 balanceLocked, uint256 assetboxType, uint256 amount)
func (_OtcImpl *OtcImplFilterer) FilterBalanceUnLocked(opts *bind.FilterOpts) (*OtcImplBalanceUnLockedIterator, error) {

	logs, sub, err := _OtcImpl.contract.FilterLogs(opts, "BalanceUnLocked")
	if err != nil {
		return nil, err
	}
	return &OtcImplBalanceUnLockedIterator{contract: _OtcImpl.contract, event: "BalanceUnLocked", logs: logs, sub: sub}, nil
}

// WatchBalanceUnLocked is a free log subscription operation binding the contract event 0x486bd0991e37e48fc69ea3be6a0b11058d29ef8a89bbc475b139a00a681b4c9b.
//
// Solidity: event BalanceUnLocked(address assetbox, uint256 balanceAviable, uint256 balanceLocked, uint256 assetboxType, uint256 amount)
func (_OtcImpl *OtcImplFilterer) WatchBalanceUnLocked(opts *bind.WatchOpts, sink chan<- *OtcImplBalanceUnLocked) (event.Subscription, error) {

	logs, sub, err := _OtcImpl.contract.WatchLogs(opts, "BalanceUnLocked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OtcImplBalanceUnLocked)
				if err := _OtcImpl.contract.UnpackLog(event, "BalanceUnLocked", log); err != nil {
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

// ParseBalanceUnLocked is a log parse operation binding the contract event 0x486bd0991e37e48fc69ea3be6a0b11058d29ef8a89bbc475b139a00a681b4c9b.
//
// Solidity: event BalanceUnLocked(address assetbox, uint256 balanceAviable, uint256 balanceLocked, uint256 assetboxType, uint256 amount)
func (_OtcImpl *OtcImplFilterer) ParseBalanceUnLocked(log types.Log) (*OtcImplBalanceUnLocked, error) {
	event := new(OtcImplBalanceUnLocked)
	if err := _OtcImpl.contract.UnpackLog(event, "BalanceUnLocked", log); err != nil {
		return nil, err
	}
	return event, nil
}
