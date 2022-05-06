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

// ExchangeImplABI is the input ABI used to generate the binding from.
const ExchangeImplABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractStorageAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"assetbox\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"balanceAviable\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"balanceLocked\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetboxType\",\"type\":\"uint256\"}],\"name\":\"BalanceLocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"assetbox\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"balanceAviable\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"balanceLocked\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetboxType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BalanceUnLocked\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ACCESS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ACCESS_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ADMINABLE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ADMIN_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ASSETBOX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ASSETBOX_INFO\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ASSETBOX_INFO_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ASSETBOX_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_BALANCE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_BITBON\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_BITBON_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_BITBON_SUPPORT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_BONBALANCE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_BONTRANSFER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_DEX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_DISTRIBUTION_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_EXCHANGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_EXCHANGE_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_FEE_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_GENERATOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MINING_AGENT_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MSBON\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MSBON_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_ADD_ADDRESS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_ADD_ADMIN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_ADD_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_ADD_ROLE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_DROP_PERMISSION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_EDIT_ADMIN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_EDIT_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_EDIT_ROLE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REMOVE_ADMIN_BY_ADDRESS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REMOVE_ADMIN_BY_INDEX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REMOVE_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REMOVE_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REMOVE_ROLE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REPLENISH_CAPITALIZATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REPLENISH_CERTIFICATE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REPLENISH_EMISSION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_SET_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_SET_PERMISSION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_OTC\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_OTC_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_REPLENISH_CAPITALIZATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_RESERVED_ALIAS_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ROLE_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_SAFE_TRANSFER_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_TRANSFER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_TRANSFER_FROM_CAPITALIZATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_ACCESS_RESTORATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_ADMIN_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_CONTRACTS_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_CONTRACT_DEPLOY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_EMISSION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_FEES_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_MULTISIG_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_ROLES_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_ACCESS_VERIFIER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_BITBON_ISSUE_VERIFIER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_COMMISSION_VERIFIER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_DEPLOY_ADMIN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_PERMISSION_ADMIN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"contractStorage\",\"outputs\":[{\"internalType\":\"contractContractStorage\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"state\",\"outputs\":[{\"internalType\":\"enumDtoDex.ExchangeState\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getThisContractIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"enumDtoDex.ExchangeState\",\"name\":\"newState\",\"type\":\"uint8\"}],\"name\":\"setState\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"offerAssets\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"wantAssets\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"offerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"wantAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"name\":\"makeBitbonOrder\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"offerAssets\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"wantAssets\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"offerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"wantAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"name\":\"makeMsbonOrder\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"fillAssets\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"wantAssets\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"offerHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amountToTake\",\"type\":\"uint256\"}],\"name\":\"fillBitbonOrder\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"fillAssets\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"wantAssets\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"offerHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amountToTake\",\"type\":\"uint256\"}],\"name\":\"fillMsbonOrder\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"offerHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"expectedAvailableAmount\",\"type\":\"uint256\"}],\"name\":\"cancelBitbonOrder\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"offerHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"expectedAvailableAmount\",\"type\":\"uint256\"}],\"name\":\"cancelMsbonOrder\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ExchangeImpl is an auto generated Go binding around an Ethereum contract.
type ExchangeImpl struct {
	ExchangeImplCaller     // Read-only binding to the contract
	ExchangeImplTransactor // Write-only binding to the contract
	ExchangeImplFilterer   // Log filterer for contract events
}

// ExchangeImplCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExchangeImplCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangeImplTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExchangeImplTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangeImplFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExchangeImplFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangeImplSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExchangeImplSession struct {
	Contract     *ExchangeImpl     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ExchangeImplCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExchangeImplCallerSession struct {
	Contract *ExchangeImplCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ExchangeImplTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExchangeImplTransactorSession struct {
	Contract     *ExchangeImplTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ExchangeImplRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExchangeImplRaw struct {
	Contract *ExchangeImpl // Generic contract binding to access the raw methods on
}

// ExchangeImplCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExchangeImplCallerRaw struct {
	Contract *ExchangeImplCaller // Generic read-only contract binding to access the raw methods on
}

// ExchangeImplTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExchangeImplTransactorRaw struct {
	Contract *ExchangeImplTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExchangeImpl creates a new instance of ExchangeImpl, bound to a specific deployed contract.
func NewExchangeImpl(address common.Address, backend bind.ContractBackend) (*ExchangeImpl, error) {
	contract, err := bindExchangeImpl(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ExchangeImpl{ExchangeImplCaller: ExchangeImplCaller{contract: contract}, ExchangeImplTransactor: ExchangeImplTransactor{contract: contract}, ExchangeImplFilterer: ExchangeImplFilterer{contract: contract}}, nil
}

// NewExchangeImplCaller creates a new read-only instance of ExchangeImpl, bound to a specific deployed contract.
func NewExchangeImplCaller(address common.Address, caller bind.ContractCaller) (*ExchangeImplCaller, error) {
	contract, err := bindExchangeImpl(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExchangeImplCaller{contract: contract}, nil
}

// NewExchangeImplTransactor creates a new write-only instance of ExchangeImpl, bound to a specific deployed contract.
func NewExchangeImplTransactor(address common.Address, transactor bind.ContractTransactor) (*ExchangeImplTransactor, error) {
	contract, err := bindExchangeImpl(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExchangeImplTransactor{contract: contract}, nil
}

// NewExchangeImplFilterer creates a new log filterer instance of ExchangeImpl, bound to a specific deployed contract.
func NewExchangeImplFilterer(address common.Address, filterer bind.ContractFilterer) (*ExchangeImplFilterer, error) {
	contract, err := bindExchangeImpl(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExchangeImplFilterer{contract: contract}, nil
}

// bindExchangeImpl binds a generic wrapper to an already deployed contract.
func bindExchangeImpl(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ExchangeImplABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExchangeImpl *ExchangeImplRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ExchangeImpl.Contract.ExchangeImplCaller.contract.Call(opts, result, method, params...)
}

// Return Abi for ExchangeImplCaller
func (_ExchangeImpl *ExchangeImplCaller) ABI() abi.ABI {
	return _ExchangeImpl.contract.Abi
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExchangeImpl *ExchangeImplRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExchangeImpl.Contract.ExchangeImplTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExchangeImpl *ExchangeImplRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExchangeImpl.Contract.ExchangeImplTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExchangeImpl *ExchangeImplCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ExchangeImpl.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExchangeImpl *ExchangeImplTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExchangeImpl.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExchangeImpl *ExchangeImplTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExchangeImpl.Contract.contract.Transact(opts, method, params...)
}

// CONTRACTACCESS is a free data retrieval call binding the contract method 0x0ac62b0c.
//
// Solidity: function CONTRACT_ACCESS() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCaller) CONTRACTACCESS(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ExchangeImpl.contract.Call(opts, out, "CONTRACT_ACCESS")
	return *ret0, err
}

// CONTRACTACCESS is a free data retrieval call binding the contract method 0x0ac62b0c.
//
// Solidity: function CONTRACT_ACCESS() view returns(uint256)
func (_ExchangeImpl *ExchangeImplSession) CONTRACTACCESS() (*big.Int, error) {
	return _ExchangeImpl.Contract.CONTRACTACCESS(&_ExchangeImpl.CallOpts)
}

// CONTRACTACCESS is a free data retrieval call binding the contract method 0x0ac62b0c.
//
// Solidity: function CONTRACT_ACCESS() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCallerSession) CONTRACTACCESS() (*big.Int, error) {
	return _ExchangeImpl.Contract.CONTRACTACCESS(&_ExchangeImpl.CallOpts)
}

// CONTRACTACCESSSTORAGE is a free data retrieval call binding the contract method 0xceac9bd5.
//
// Solidity: function CONTRACT_ACCESS_STORAGE() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCaller) CONTRACTACCESSSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ExchangeImpl.contract.Call(opts, out, "CONTRACT_ACCESS_STORAGE")
	return *ret0, err
}

// CONTRACTACCESSSTORAGE is a free data retrieval call binding the contract method 0xceac9bd5.
//
// Solidity: function CONTRACT_ACCESS_STORAGE() view returns(uint256)
func (_ExchangeImpl *ExchangeImplSession) CONTRACTACCESSSTORAGE() (*big.Int, error) {
	return _ExchangeImpl.Contract.CONTRACTACCESSSTORAGE(&_ExchangeImpl.CallOpts)
}

// CONTRACTACCESSSTORAGE is a free data retrieval call binding the contract method 0xceac9bd5.
//
// Solidity: function CONTRACT_ACCESS_STORAGE() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCallerSession) CONTRACTACCESSSTORAGE() (*big.Int, error) {
	return _ExchangeImpl.Contract.CONTRACTACCESSSTORAGE(&_ExchangeImpl.CallOpts)
}

// CONTRACTADMINABLE is a free data retrieval call binding the contract method 0xbbe3a78f.
//
// Solidity: function CONTRACT_ADMINABLE() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCaller) CONTRACTADMINABLE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ExchangeImpl.contract.Call(opts, out, "CONTRACT_ADMINABLE")
	return *ret0, err
}

// CONTRACTADMINABLE is a free data retrieval call binding the contract method 0xbbe3a78f.
//
// Solidity: function CONTRACT_ADMINABLE() view returns(uint256)
func (_ExchangeImpl *ExchangeImplSession) CONTRACTADMINABLE() (*big.Int, error) {
	return _ExchangeImpl.Contract.CONTRACTADMINABLE(&_ExchangeImpl.CallOpts)
}

// CONTRACTADMINABLE is a free data retrieval call binding the contract method 0xbbe3a78f.
//
// Solidity: function CONTRACT_ADMINABLE() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCallerSession) CONTRACTADMINABLE() (*big.Int, error) {
	return _ExchangeImpl.Contract.CONTRACTADMINABLE(&_ExchangeImpl.CallOpts)
}

// CONTRACTADMINSTORAGE is a free data retrieval call binding the contract method 0xc9aebae9.
//
// Solidity: function CONTRACT_ADMIN_STORAGE() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCaller) CONTRACTADMINSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ExchangeImpl.contract.Call(opts, out, "CONTRACT_ADMIN_STORAGE")
	return *ret0, err
}

// CONTRACTADMINSTORAGE is a free data retrieval call binding the contract method 0xc9aebae9.
//
// Solidity: function CONTRACT_ADMIN_STORAGE() view returns(uint256)
func (_ExchangeImpl *ExchangeImplSession) CONTRACTADMINSTORAGE() (*big.Int, error) {
	return _ExchangeImpl.Contract.CONTRACTADMINSTORAGE(&_ExchangeImpl.CallOpts)
}

// CONTRACTADMINSTORAGE is a free data retrieval call binding the contract method 0xc9aebae9.
//
// Solidity: function CONTRACT_ADMIN_STORAGE() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCallerSession) CONTRACTADMINSTORAGE() (*big.Int, error) {
	return _ExchangeImpl.Contract.CONTRACTADMINSTORAGE(&_ExchangeImpl.CallOpts)
}

// CONTRACTASSETBOX is a free data retrieval call binding the contract method 0x161aa17c.
//
// Solidity: function CONTRACT_ASSETBOX() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCaller) CONTRACTASSETBOX(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ExchangeImpl.contract.Call(opts, out, "CONTRACT_ASSETBOX")
	return *ret0, err
}

// CONTRACTASSETBOX is a free data retrieval call binding the contract method 0x161aa17c.
//
// Solidity: function CONTRACT_ASSETBOX() view returns(uint256)
func (_ExchangeImpl *ExchangeImplSession) CONTRACTASSETBOX() (*big.Int, error) {
	return _ExchangeImpl.Contract.CONTRACTASSETBOX(&_ExchangeImpl.CallOpts)
}

// CONTRACTASSETBOX is a free data retrieval call binding the contract method 0x161aa17c.
//
// Solidity: function CONTRACT_ASSETBOX() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCallerSession) CONTRACTASSETBOX() (*big.Int, error) {
	return _ExchangeImpl.Contract.CONTRACTASSETBOX(&_ExchangeImpl.CallOpts)
}

// CONTRACTASSETBOXINFO is a free data retrieval call binding the contract method 0xb572fd3d.
//
// Solidity: function CONTRACT_ASSETBOX_INFO() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCaller) CONTRACTASSETBOXINFO(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ExchangeImpl.contract.Call(opts, out, "CONTRACT_ASSETBOX_INFO")
	return *ret0, err
}

// CONTRACTASSETBOXINFO is a free data retrieval call binding the contract method 0xb572fd3d.
//
// Solidity: function CONTRACT_ASSETBOX_INFO() view returns(uint256)
func (_ExchangeImpl *ExchangeImplSession) CONTRACTASSETBOXINFO() (*big.Int, error) {
	return _ExchangeImpl.Contract.CONTRACTASSETBOXINFO(&_ExchangeImpl.CallOpts)
}

// CONTRACTASSETBOXINFO is a free data retrieval call binding the contract method 0xb572fd3d.
//
// Solidity: function CONTRACT_ASSETBOX_INFO() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCallerSession) CONTRACTASSETBOXINFO() (*big.Int, error) {
	return _ExchangeImpl.Contract.CONTRACTASSETBOXINFO(&_ExchangeImpl.CallOpts)
}

// CONTRACTASSETBOXINFOSTORAGE is a free data retrieval call binding the contract method 0xf32a115d.
//
// Solidity: function CONTRACT_ASSETBOX_INFO_STORAGE() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCaller) CONTRACTASSETBOXINFOSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ExchangeImpl.contract.Call(opts, out, "CONTRACT_ASSETBOX_INFO_STORAGE")
	return *ret0, err
}

// CONTRACTASSETBOXINFOSTORAGE is a free data retrieval call binding the contract method 0xf32a115d.
//
// Solidity: function CONTRACT_ASSETBOX_INFO_STORAGE() view returns(uint256)
func (_ExchangeImpl *ExchangeImplSession) CONTRACTASSETBOXINFOSTORAGE() (*big.Int, error) {
	return _ExchangeImpl.Contract.CONTRACTASSETBOXINFOSTORAGE(&_ExchangeImpl.CallOpts)
}

// CONTRACTASSETBOXINFOSTORAGE is a free data retrieval call binding the contract method 0xf32a115d.
//
// Solidity: function CONTRACT_ASSETBOX_INFO_STORAGE() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCallerSession) CONTRACTASSETBOXINFOSTORAGE() (*big.Int, error) {
	return _ExchangeImpl.Contract.CONTRACTASSETBOXINFOSTORAGE(&_ExchangeImpl.CallOpts)
}

// CONTRACTASSETBOXSTORAGE is a free data retrieval call binding the contract method 0xfa5bdafc.
//
// Solidity: function CONTRACT_ASSETBOX_STORAGE() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCaller) CONTRACTASSETBOXSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ExchangeImpl.contract.Call(opts, out, "CONTRACT_ASSETBOX_STORAGE")
	return *ret0, err
}

// CONTRACTASSETBOXSTORAGE is a free data retrieval call binding the contract method 0xfa5bdafc.
//
// Solidity: function CONTRACT_ASSETBOX_STORAGE() view returns(uint256)
func (_ExchangeImpl *ExchangeImplSession) CONTRACTASSETBOXSTORAGE() (*big.Int, error) {
	return _ExchangeImpl.Contract.CONTRACTASSETBOXSTORAGE(&_ExchangeImpl.CallOpts)
}

// CONTRACTASSETBOXSTORAGE is a free data retrieval call binding the contract method 0xfa5bdafc.
//
// Solidity: function CONTRACT_ASSETBOX_STORAGE() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCallerSession) CONTRACTASSETBOXSTORAGE() (*big.Int, error) {
	return _ExchangeImpl.Contract.CONTRACTASSETBOXSTORAGE(&_ExchangeImpl.CallOpts)
}

// CONTRACTBALANCE is a free data retrieval call binding the contract method 0xd59045f5.
//
// Solidity: function CONTRACT_BALANCE() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCaller) CONTRACTBALANCE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ExchangeImpl.contract.Call(opts, out, "CONTRACT_BALANCE")
	return *ret0, err
}

// CONTRACTBALANCE is a free data retrieval call binding the contract method 0xd59045f5.
//
// Solidity: function CONTRACT_BALANCE() view returns(uint256)
func (_ExchangeImpl *ExchangeImplSession) CONTRACTBALANCE() (*big.Int, error) {
	return _ExchangeImpl.Contract.CONTRACTBALANCE(&_ExchangeImpl.CallOpts)
}

// CONTRACTBALANCE is a free data retrieval call binding the contract method 0xd59045f5.
//
// Solidity: function CONTRACT_BALANCE() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCallerSession) CONTRACTBALANCE() (*big.Int, error) {
	return _ExchangeImpl.Contract.CONTRACTBALANCE(&_ExchangeImpl.CallOpts)
}

// CONTRACTBITBON is a free data retrieval call binding the contract method 0xb6b433ab.
//
// Solidity: function CONTRACT_BITBON() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCaller) CONTRACTBITBON(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ExchangeImpl.contract.Call(opts, out, "CONTRACT_BITBON")
	return *ret0, err
}

// CONTRACTBITBON is a free data retrieval call binding the contract method 0xb6b433ab.
//
// Solidity: function CONTRACT_BITBON() view returns(uint256)
func (_ExchangeImpl *ExchangeImplSession) CONTRACTBITBON() (*big.Int, error) {
	return _ExchangeImpl.Contract.CONTRACTBITBON(&_ExchangeImpl.CallOpts)
}

// CONTRACTBITBON is a free data retrieval call binding the contract method 0xb6b433ab.
//
// Solidity: function CONTRACT_BITBON() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCallerSession) CONTRACTBITBON() (*big.Int, error) {
	return _ExchangeImpl.Contract.CONTRACTBITBON(&_ExchangeImpl.CallOpts)
}

// CONTRACTBITBONSTORAGE is a free data retrieval call binding the contract method 0xe7f6c30e.
//
// Solidity: function CONTRACT_BITBON_STORAGE() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCaller) CONTRACTBITBONSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ExchangeImpl.contract.Call(opts, out, "CONTRACT_BITBON_STORAGE")
	return *ret0, err
}

// CONTRACTBITBONSTORAGE is a free data retrieval call binding the contract method 0xe7f6c30e.
//
// Solidity: function CONTRACT_BITBON_STORAGE() view returns(uint256)
func (_ExchangeImpl *ExchangeImplSession) CONTRACTBITBONSTORAGE() (*big.Int, error) {
	return _ExchangeImpl.Contract.CONTRACTBITBONSTORAGE(&_ExchangeImpl.CallOpts)
}

// CONTRACTBITBONSTORAGE is a free data retrieval call binding the contract method 0xe7f6c30e.
//
// Solidity: function CONTRACT_BITBON_STORAGE() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCallerSession) CONTRACTBITBONSTORAGE() (*big.Int, error) {
	return _ExchangeImpl.Contract.CONTRACTBITBONSTORAGE(&_ExchangeImpl.CallOpts)
}

// CONTRACTBITBONSUPPORT is a free data retrieval call binding the contract method 0x20139c74.
//
// Solidity: function CONTRACT_BITBON_SUPPORT() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCaller) CONTRACTBITBONSUPPORT(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ExchangeImpl.contract.Call(opts, out, "CONTRACT_BITBON_SUPPORT")
	return *ret0, err
}

// CONTRACTBITBONSUPPORT is a free data retrieval call binding the contract method 0x20139c74.
//
// Solidity: function CONTRACT_BITBON_SUPPORT() view returns(uint256)
func (_ExchangeImpl *ExchangeImplSession) CONTRACTBITBONSUPPORT() (*big.Int, error) {
	return _ExchangeImpl.Contract.CONTRACTBITBONSUPPORT(&_ExchangeImpl.CallOpts)
}

// CONTRACTBITBONSUPPORT is a free data retrieval call binding the contract method 0x20139c74.
//
// Solidity: function CONTRACT_BITBON_SUPPORT() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCallerSession) CONTRACTBITBONSUPPORT() (*big.Int, error) {
	return _ExchangeImpl.Contract.CONTRACTBITBONSUPPORT(&_ExchangeImpl.CallOpts)
}

// CONTRACTBONBALANCE is a free data retrieval call binding the contract method 0x40e02b7c.
//
// Solidity: function CONTRACT_BONBALANCE() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCaller) CONTRACTBONBALANCE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ExchangeImpl.contract.Call(opts, out, "CONTRACT_BONBALANCE")
	return *ret0, err
}

// CONTRACTBONBALANCE is a free data retrieval call binding the contract method 0x40e02b7c.
//
// Solidity: function CONTRACT_BONBALANCE() view returns(uint256)
func (_ExchangeImpl *ExchangeImplSession) CONTRACTBONBALANCE() (*big.Int, error) {
	return _ExchangeImpl.Contract.CONTRACTBONBALANCE(&_ExchangeImpl.CallOpts)
}

// CONTRACTBONBALANCE is a free data retrieval call binding the contract method 0x40e02b7c.
//
// Solidity: function CONTRACT_BONBALANCE() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCallerSession) CONTRACTBONBALANCE() (*big.Int, error) {
	return _ExchangeImpl.Contract.CONTRACTBONBALANCE(&_ExchangeImpl.CallOpts)
}

// CONTRACTBONTRANSFER is a free data retrieval call binding the contract method 0xc18f4a8a.
//
// Solidity: function CONTRACT_BONTRANSFER() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCaller) CONTRACTBONTRANSFER(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ExchangeImpl.contract.Call(opts, out, "CONTRACT_BONTRANSFER")
	return *ret0, err
}

// CONTRACTBONTRANSFER is a free data retrieval call binding the contract method 0xc18f4a8a.
//
// Solidity: function CONTRACT_BONTRANSFER() view returns(uint256)
func (_ExchangeImpl *ExchangeImplSession) CONTRACTBONTRANSFER() (*big.Int, error) {
	return _ExchangeImpl.Contract.CONTRACTBONTRANSFER(&_ExchangeImpl.CallOpts)
}

// CONTRACTBONTRANSFER is a free data retrieval call binding the contract method 0xc18f4a8a.
//
// Solidity: function CONTRACT_BONTRANSFER() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCallerSession) CONTRACTBONTRANSFER() (*big.Int, error) {
	return _ExchangeImpl.Contract.CONTRACTBONTRANSFER(&_ExchangeImpl.CallOpts)
}

// CONTRACTDEX is a free data retrieval call binding the contract method 0x52c1f844.
//
// Solidity: function CONTRACT_DEX() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCaller) CONTRACTDEX(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ExchangeImpl.contract.Call(opts, out, "CONTRACT_DEX")
	return *ret0, err
}

// CONTRACTDEX is a free data retrieval call binding the contract method 0x52c1f844.
//
// Solidity: function CONTRACT_DEX() view returns(uint256)
func (_ExchangeImpl *ExchangeImplSession) CONTRACTDEX() (*big.Int, error) {
	return _ExchangeImpl.Contract.CONTRACTDEX(&_ExchangeImpl.CallOpts)
}

// CONTRACTDEX is a free data retrieval call binding the contract method 0x52c1f844.
//
// Solidity: function CONTRACT_DEX() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCallerSession) CONTRACTDEX() (*big.Int, error) {
	return _ExchangeImpl.Contract.CONTRACTDEX(&_ExchangeImpl.CallOpts)
}

// CONTRACTDISTRIBUTIONSTORAGE is a free data retrieval call binding the contract method 0x08097327.
//
// Solidity: function CONTRACT_DISTRIBUTION_STORAGE() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCaller) CONTRACTDISTRIBUTIONSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ExchangeImpl.contract.Call(opts, out, "CONTRACT_DISTRIBUTION_STORAGE")
	return *ret0, err
}

// CONTRACTDISTRIBUTIONSTORAGE is a free data retrieval call binding the contract method 0x08097327.
//
// Solidity: function CONTRACT_DISTRIBUTION_STORAGE() view returns(uint256)
func (_ExchangeImpl *ExchangeImplSession) CONTRACTDISTRIBUTIONSTORAGE() (*big.Int, error) {
	return _ExchangeImpl.Contract.CONTRACTDISTRIBUTIONSTORAGE(&_ExchangeImpl.CallOpts)
}

// CONTRACTDISTRIBUTIONSTORAGE is a free data retrieval call binding the contract method 0x08097327.
//
// Solidity: function CONTRACT_DISTRIBUTION_STORAGE() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCallerSession) CONTRACTDISTRIBUTIONSTORAGE() (*big.Int, error) {
	return _ExchangeImpl.Contract.CONTRACTDISTRIBUTIONSTORAGE(&_ExchangeImpl.CallOpts)
}

// CONTRACTEXCHANGE is a free data retrieval call binding the contract method 0xa0e3a10e.
//
// Solidity: function CONTRACT_EXCHANGE() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCaller) CONTRACTEXCHANGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ExchangeImpl.contract.Call(opts, out, "CONTRACT_EXCHANGE")
	return *ret0, err
}

// CONTRACTEXCHANGE is a free data retrieval call binding the contract method 0xa0e3a10e.
//
// Solidity: function CONTRACT_EXCHANGE() view returns(uint256)
func (_ExchangeImpl *ExchangeImplSession) CONTRACTEXCHANGE() (*big.Int, error) {
	return _ExchangeImpl.Contract.CONTRACTEXCHANGE(&_ExchangeImpl.CallOpts)
}

// CONTRACTEXCHANGE is a free data retrieval call binding the contract method 0xa0e3a10e.
//
// Solidity: function CONTRACT_EXCHANGE() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCallerSession) CONTRACTEXCHANGE() (*big.Int, error) {
	return _ExchangeImpl.Contract.CONTRACTEXCHANGE(&_ExchangeImpl.CallOpts)
}

// CONTRACTEXCHANGESTORAGE is a free data retrieval call binding the contract method 0x001570f2.
//
// Solidity: function CONTRACT_EXCHANGE_STORAGE() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCaller) CONTRACTEXCHANGESTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ExchangeImpl.contract.Call(opts, out, "CONTRACT_EXCHANGE_STORAGE")
	return *ret0, err
}

// CONTRACTEXCHANGESTORAGE is a free data retrieval call binding the contract method 0x001570f2.
//
// Solidity: function CONTRACT_EXCHANGE_STORAGE() view returns(uint256)
func (_ExchangeImpl *ExchangeImplSession) CONTRACTEXCHANGESTORAGE() (*big.Int, error) {
	return _ExchangeImpl.Contract.CONTRACTEXCHANGESTORAGE(&_ExchangeImpl.CallOpts)
}

// CONTRACTEXCHANGESTORAGE is a free data retrieval call binding the contract method 0x001570f2.
//
// Solidity: function CONTRACT_EXCHANGE_STORAGE() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCallerSession) CONTRACTEXCHANGESTORAGE() (*big.Int, error) {
	return _ExchangeImpl.Contract.CONTRACTEXCHANGESTORAGE(&_ExchangeImpl.CallOpts)
}

// CONTRACTFEESTORAGE is a free data retrieval call binding the contract method 0xcd5e65a4.
//
// Solidity: function CONTRACT_FEE_STORAGE() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCaller) CONTRACTFEESTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ExchangeImpl.contract.Call(opts, out, "CONTRACT_FEE_STORAGE")
	return *ret0, err
}

// CONTRACTFEESTORAGE is a free data retrieval call binding the contract method 0xcd5e65a4.
//
// Solidity: function CONTRACT_FEE_STORAGE() view returns(uint256)
func (_ExchangeImpl *ExchangeImplSession) CONTRACTFEESTORAGE() (*big.Int, error) {
	return _ExchangeImpl.Contract.CONTRACTFEESTORAGE(&_ExchangeImpl.CallOpts)
}

// CONTRACTFEESTORAGE is a free data retrieval call binding the contract method 0xcd5e65a4.
//
// Solidity: function CONTRACT_FEE_STORAGE() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCallerSession) CONTRACTFEESTORAGE() (*big.Int, error) {
	return _ExchangeImpl.Contract.CONTRACTFEESTORAGE(&_ExchangeImpl.CallOpts)
}

// CONTRACTGENERATOR is a free data retrieval call binding the contract method 0xccea2b46.
//
// Solidity: function CONTRACT_GENERATOR() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCaller) CONTRACTGENERATOR(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ExchangeImpl.contract.Call(opts, out, "CONTRACT_GENERATOR")
	return *ret0, err
}

// CONTRACTGENERATOR is a free data retrieval call binding the contract method 0xccea2b46.
//
// Solidity: function CONTRACT_GENERATOR() view returns(uint256)
func (_ExchangeImpl *ExchangeImplSession) CONTRACTGENERATOR() (*big.Int, error) {
	return _ExchangeImpl.Contract.CONTRACTGENERATOR(&_ExchangeImpl.CallOpts)
}

// CONTRACTGENERATOR is a free data retrieval call binding the contract method 0xccea2b46.
//
// Solidity: function CONTRACT_GENERATOR() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCallerSession) CONTRACTGENERATOR() (*big.Int, error) {
	return _ExchangeImpl.Contract.CONTRACTGENERATOR(&_ExchangeImpl.CallOpts)
}

// CONTRACTMININGAGENTSTORAGE is a free data retrieval call binding the contract method 0x83cc588d.
//
// Solidity: function CONTRACT_MINING_AGENT_STORAGE() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCaller) CONTRACTMININGAGENTSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ExchangeImpl.contract.Call(opts, out, "CONTRACT_MINING_AGENT_STORAGE")
	return *ret0, err
}

// CONTRACTMININGAGENTSTORAGE is a free data retrieval call binding the contract method 0x83cc588d.
//
// Solidity: function CONTRACT_MINING_AGENT_STORAGE() view returns(uint256)
func (_ExchangeImpl *ExchangeImplSession) CONTRACTMININGAGENTSTORAGE() (*big.Int, error) {
	return _ExchangeImpl.Contract.CONTRACTMININGAGENTSTORAGE(&_ExchangeImpl.CallOpts)
}

// CONTRACTMININGAGENTSTORAGE is a free data retrieval call binding the contract method 0x83cc588d.
//
// Solidity: function CONTRACT_MINING_AGENT_STORAGE() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCallerSession) CONTRACTMININGAGENTSTORAGE() (*big.Int, error) {
	return _ExchangeImpl.Contract.CONTRACTMININGAGENTSTORAGE(&_ExchangeImpl.CallOpts)
}

// CONTRACTMSBON is a free data retrieval call binding the contract method 0x28b3a976.
//
// Solidity: function CONTRACT_MSBON() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCaller) CONTRACTMSBON(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ExchangeImpl.contract.Call(opts, out, "CONTRACT_MSBON")
	return *ret0, err
}

// CONTRACTMSBON is a free data retrieval call binding the contract method 0x28b3a976.
//
// Solidity: function CONTRACT_MSBON() view returns(uint256)
func (_ExchangeImpl *ExchangeImplSession) CONTRACTMSBON() (*big.Int, error) {
	return _ExchangeImpl.Contract.CONTRACTMSBON(&_ExchangeImpl.CallOpts)
}

// CONTRACTMSBON is a free data retrieval call binding the contract method 0x28b3a976.
//
// Solidity: function CONTRACT_MSBON() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCallerSession) CONTRACTMSBON() (*big.Int, error) {
	return _ExchangeImpl.Contract.CONTRACTMSBON(&_ExchangeImpl.CallOpts)
}

// CONTRACTMSBONSTORAGE is a free data retrieval call binding the contract method 0xe04e6061.
//
// Solidity: function CONTRACT_MSBON_STORAGE() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCaller) CONTRACTMSBONSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ExchangeImpl.contract.Call(opts, out, "CONTRACT_MSBON_STORAGE")
	return *ret0, err
}

// CONTRACTMSBONSTORAGE is a free data retrieval call binding the contract method 0xe04e6061.
//
// Solidity: function CONTRACT_MSBON_STORAGE() view returns(uint256)
func (_ExchangeImpl *ExchangeImplSession) CONTRACTMSBONSTORAGE() (*big.Int, error) {
	return _ExchangeImpl.Contract.CONTRACTMSBONSTORAGE(&_ExchangeImpl.CallOpts)
}

// CONTRACTMSBONSTORAGE is a free data retrieval call binding the contract method 0xe04e6061.
//
// Solidity: function CONTRACT_MSBON_STORAGE() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCallerSession) CONTRACTMSBONSTORAGE() (*big.Int, error) {
	return _ExchangeImpl.Contract.CONTRACTMSBONSTORAGE(&_ExchangeImpl.CallOpts)
}

// CONTRACTMULTISIGADDADDRESS is a free data retrieval call binding the contract method 0xf4f65a5a.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ADDRESS() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCaller) CONTRACTMULTISIGADDADDRESS(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ExchangeImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_ADD_ADDRESS")
	return *ret0, err
}

// CONTRACTMULTISIGADDADDRESS is a free data retrieval call binding the contract method 0xf4f65a5a.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ADDRESS() view returns(uint256)
func (_ExchangeImpl *ExchangeImplSession) CONTRACTMULTISIGADDADDRESS() (*big.Int, error) {
	return _ExchangeImpl.Contract.CONTRACTMULTISIGADDADDRESS(&_ExchangeImpl.CallOpts)
}

// CONTRACTMULTISIGADDADDRESS is a free data retrieval call binding the contract method 0xf4f65a5a.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ADDRESS() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCallerSession) CONTRACTMULTISIGADDADDRESS() (*big.Int, error) {
	return _ExchangeImpl.Contract.CONTRACTMULTISIGADDADDRESS(&_ExchangeImpl.CallOpts)
}

// CONTRACTMULTISIGADDADMIN is a free data retrieval call binding the contract method 0x83888c68.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ADMIN() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCaller) CONTRACTMULTISIGADDADMIN(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ExchangeImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_ADD_ADMIN")
	return *ret0, err
}

// CONTRACTMULTISIGADDADMIN is a free data retrieval call binding the contract method 0x83888c68.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ADMIN() view returns(uint256)
func (_ExchangeImpl *ExchangeImplSession) CONTRACTMULTISIGADDADMIN() (*big.Int, error) {
	return _ExchangeImpl.Contract.CONTRACTMULTISIGADDADMIN(&_ExchangeImpl.CallOpts)
}

// CONTRACTMULTISIGADDADMIN is a free data retrieval call binding the contract method 0x83888c68.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ADMIN() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCallerSession) CONTRACTMULTISIGADDADMIN() (*big.Int, error) {
	return _ExchangeImpl.Contract.CONTRACTMULTISIGADDADMIN(&_ExchangeImpl.CallOpts)
}

// CONTRACTMULTISIGADDFEE is a free data retrieval call binding the contract method 0xc43e8cab.
//
// Solidity: function CONTRACT_MULTISIG_ADD_FEE() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCaller) CONTRACTMULTISIGADDFEE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ExchangeImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_ADD_FEE")
	return *ret0, err
}

// CONTRACTMULTISIGADDFEE is a free data retrieval call binding the contract method 0xc43e8cab.
//
// Solidity: function CONTRACT_MULTISIG_ADD_FEE() view returns(uint256)
func (_ExchangeImpl *ExchangeImplSession) CONTRACTMULTISIGADDFEE() (*big.Int, error) {
	return _ExchangeImpl.Contract.CONTRACTMULTISIGADDFEE(&_ExchangeImpl.CallOpts)
}

// CONTRACTMULTISIGADDFEE is a free data retrieval call binding the contract method 0xc43e8cab.
//
// Solidity: function CONTRACT_MULTISIG_ADD_FEE() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCallerSession) CONTRACTMULTISIGADDFEE() (*big.Int, error) {
	return _ExchangeImpl.Contract.CONTRACTMULTISIGADDFEE(&_ExchangeImpl.CallOpts)
}

// CONTRACTMULTISIGADDROLE is a free data retrieval call binding the contract method 0x20a0914b.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ROLE() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCaller) CONTRACTMULTISIGADDROLE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ExchangeImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_ADD_ROLE")
	return *ret0, err
}

// CONTRACTMULTISIGADDROLE is a free data retrieval call binding the contract method 0x20a0914b.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ROLE() view returns(uint256)
func (_ExchangeImpl *ExchangeImplSession) CONTRACTMULTISIGADDROLE() (*big.Int, error) {
	return _ExchangeImpl.Contract.CONTRACTMULTISIGADDROLE(&_ExchangeImpl.CallOpts)
}

// CONTRACTMULTISIGADDROLE is a free data retrieval call binding the contract method 0x20a0914b.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ROLE() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCallerSession) CONTRACTMULTISIGADDROLE() (*big.Int, error) {
	return _ExchangeImpl.Contract.CONTRACTMULTISIGADDROLE(&_ExchangeImpl.CallOpts)
}

// CONTRACTMULTISIGDROPPERMISSION is a free data retrieval call binding the contract method 0x3f3fbc39.
//
// Solidity: function CONTRACT_MULTISIG_DROP_PERMISSION() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCaller) CONTRACTMULTISIGDROPPERMISSION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ExchangeImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_DROP_PERMISSION")
	return *ret0, err
}

// CONTRACTMULTISIGDROPPERMISSION is a free data retrieval call binding the contract method 0x3f3fbc39.
//
// Solidity: function CONTRACT_MULTISIG_DROP_PERMISSION() view returns(uint256)
func (_ExchangeImpl *ExchangeImplSession) CONTRACTMULTISIGDROPPERMISSION() (*big.Int, error) {
	return _ExchangeImpl.Contract.CONTRACTMULTISIGDROPPERMISSION(&_ExchangeImpl.CallOpts)
}

// CONTRACTMULTISIGDROPPERMISSION is a free data retrieval call binding the contract method 0x3f3fbc39.
//
// Solidity: function CONTRACT_MULTISIG_DROP_PERMISSION() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCallerSession) CONTRACTMULTISIGDROPPERMISSION() (*big.Int, error) {
	return _ExchangeImpl.Contract.CONTRACTMULTISIGDROPPERMISSION(&_ExchangeImpl.CallOpts)
}

// CONTRACTMULTISIGEDITADMIN is a free data retrieval call binding the contract method 0x91b0f051.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_ADMIN() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCaller) CONTRACTMULTISIGEDITADMIN(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ExchangeImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_EDIT_ADMIN")
	return *ret0, err
}

// CONTRACTMULTISIGEDITADMIN is a free data retrieval call binding the contract method 0x91b0f051.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_ADMIN() view returns(uint256)
func (_ExchangeImpl *ExchangeImplSession) CONTRACTMULTISIGEDITADMIN() (*big.Int, error) {
	return _ExchangeImpl.Contract.CONTRACTMULTISIGEDITADMIN(&_ExchangeImpl.CallOpts)
}

// CONTRACTMULTISIGEDITADMIN is a free data retrieval call binding the contract method 0x91b0f051.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_ADMIN() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCallerSession) CONTRACTMULTISIGEDITADMIN() (*big.Int, error) {
	return _ExchangeImpl.Contract.CONTRACTMULTISIGEDITADMIN(&_ExchangeImpl.CallOpts)
}

// CONTRACTMULTISIGEDITFEE is a free data retrieval call binding the contract method 0x6b2abe17.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_FEE() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCaller) CONTRACTMULTISIGEDITFEE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ExchangeImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_EDIT_FEE")
	return *ret0, err
}

// CONTRACTMULTISIGEDITFEE is a free data retrieval call binding the contract method 0x6b2abe17.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_FEE() view returns(uint256)
func (_ExchangeImpl *ExchangeImplSession) CONTRACTMULTISIGEDITFEE() (*big.Int, error) {
	return _ExchangeImpl.Contract.CONTRACTMULTISIGEDITFEE(&_ExchangeImpl.CallOpts)
}

// CONTRACTMULTISIGEDITFEE is a free data retrieval call binding the contract method 0x6b2abe17.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_FEE() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCallerSession) CONTRACTMULTISIGEDITFEE() (*big.Int, error) {
	return _ExchangeImpl.Contract.CONTRACTMULTISIGEDITFEE(&_ExchangeImpl.CallOpts)
}

// CONTRACTMULTISIGEDITROLE is a free data retrieval call binding the contract method 0xc5f66412.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_ROLE() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCaller) CONTRACTMULTISIGEDITROLE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ExchangeImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_EDIT_ROLE")
	return *ret0, err
}

// CONTRACTMULTISIGEDITROLE is a free data retrieval call binding the contract method 0xc5f66412.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_ROLE() view returns(uint256)
func (_ExchangeImpl *ExchangeImplSession) CONTRACTMULTISIGEDITROLE() (*big.Int, error) {
	return _ExchangeImpl.Contract.CONTRACTMULTISIGEDITROLE(&_ExchangeImpl.CallOpts)
}

// CONTRACTMULTISIGEDITROLE is a free data retrieval call binding the contract method 0xc5f66412.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_ROLE() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCallerSession) CONTRACTMULTISIGEDITROLE() (*big.Int, error) {
	return _ExchangeImpl.Contract.CONTRACTMULTISIGEDITROLE(&_ExchangeImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEADMINBYADDRESS is a free data retrieval call binding the contract method 0xc816b091.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_BY_ADDRESS() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCaller) CONTRACTMULTISIGREMOVEADMINBYADDRESS(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ExchangeImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REMOVE_ADMIN_BY_ADDRESS")
	return *ret0, err
}

// CONTRACTMULTISIGREMOVEADMINBYADDRESS is a free data retrieval call binding the contract method 0xc816b091.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_BY_ADDRESS() view returns(uint256)
func (_ExchangeImpl *ExchangeImplSession) CONTRACTMULTISIGREMOVEADMINBYADDRESS() (*big.Int, error) {
	return _ExchangeImpl.Contract.CONTRACTMULTISIGREMOVEADMINBYADDRESS(&_ExchangeImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEADMINBYADDRESS is a free data retrieval call binding the contract method 0xc816b091.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_BY_ADDRESS() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCallerSession) CONTRACTMULTISIGREMOVEADMINBYADDRESS() (*big.Int, error) {
	return _ExchangeImpl.Contract.CONTRACTMULTISIGREMOVEADMINBYADDRESS(&_ExchangeImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEADMINBYINDEX is a free data retrieval call binding the contract method 0x751fbea2.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_BY_INDEX() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCaller) CONTRACTMULTISIGREMOVEADMINBYINDEX(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ExchangeImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REMOVE_ADMIN_BY_INDEX")
	return *ret0, err
}

// CONTRACTMULTISIGREMOVEADMINBYINDEX is a free data retrieval call binding the contract method 0x751fbea2.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_BY_INDEX() view returns(uint256)
func (_ExchangeImpl *ExchangeImplSession) CONTRACTMULTISIGREMOVEADMINBYINDEX() (*big.Int, error) {
	return _ExchangeImpl.Contract.CONTRACTMULTISIGREMOVEADMINBYINDEX(&_ExchangeImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEADMINBYINDEX is a free data retrieval call binding the contract method 0x751fbea2.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_BY_INDEX() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCallerSession) CONTRACTMULTISIGREMOVEADMINBYINDEX() (*big.Int, error) {
	return _ExchangeImpl.Contract.CONTRACTMULTISIGREMOVEADMINBYINDEX(&_ExchangeImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEADMINROLE is a free data retrieval call binding the contract method 0x53c00200.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_ROLE() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCaller) CONTRACTMULTISIGREMOVEADMINROLE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ExchangeImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REMOVE_ADMIN_ROLE")
	return *ret0, err
}

// CONTRACTMULTISIGREMOVEADMINROLE is a free data retrieval call binding the contract method 0x53c00200.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_ROLE() view returns(uint256)
func (_ExchangeImpl *ExchangeImplSession) CONTRACTMULTISIGREMOVEADMINROLE() (*big.Int, error) {
	return _ExchangeImpl.Contract.CONTRACTMULTISIGREMOVEADMINROLE(&_ExchangeImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEADMINROLE is a free data retrieval call binding the contract method 0x53c00200.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_ROLE() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCallerSession) CONTRACTMULTISIGREMOVEADMINROLE() (*big.Int, error) {
	return _ExchangeImpl.Contract.CONTRACTMULTISIGREMOVEADMINROLE(&_ExchangeImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEFEE is a free data retrieval call binding the contract method 0xdb614d25.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_FEE() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCaller) CONTRACTMULTISIGREMOVEFEE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ExchangeImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REMOVE_FEE")
	return *ret0, err
}

// CONTRACTMULTISIGREMOVEFEE is a free data retrieval call binding the contract method 0xdb614d25.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_FEE() view returns(uint256)
func (_ExchangeImpl *ExchangeImplSession) CONTRACTMULTISIGREMOVEFEE() (*big.Int, error) {
	return _ExchangeImpl.Contract.CONTRACTMULTISIGREMOVEFEE(&_ExchangeImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEFEE is a free data retrieval call binding the contract method 0xdb614d25.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_FEE() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCallerSession) CONTRACTMULTISIGREMOVEFEE() (*big.Int, error) {
	return _ExchangeImpl.Contract.CONTRACTMULTISIGREMOVEFEE(&_ExchangeImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEROLE is a free data retrieval call binding the contract method 0x23648496.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ROLE() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCaller) CONTRACTMULTISIGREMOVEROLE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ExchangeImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REMOVE_ROLE")
	return *ret0, err
}

// CONTRACTMULTISIGREMOVEROLE is a free data retrieval call binding the contract method 0x23648496.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ROLE() view returns(uint256)
func (_ExchangeImpl *ExchangeImplSession) CONTRACTMULTISIGREMOVEROLE() (*big.Int, error) {
	return _ExchangeImpl.Contract.CONTRACTMULTISIGREMOVEROLE(&_ExchangeImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEROLE is a free data retrieval call binding the contract method 0x23648496.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ROLE() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCallerSession) CONTRACTMULTISIGREMOVEROLE() (*big.Int, error) {
	return _ExchangeImpl.Contract.CONTRACTMULTISIGREMOVEROLE(&_ExchangeImpl.CallOpts)
}

// CONTRACTMULTISIGREPLENISHCAPITALIZATION is a free data retrieval call binding the contract method 0x62d5d523.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_CAPITALIZATION() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCaller) CONTRACTMULTISIGREPLENISHCAPITALIZATION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ExchangeImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REPLENISH_CAPITALIZATION")
	return *ret0, err
}

// CONTRACTMULTISIGREPLENISHCAPITALIZATION is a free data retrieval call binding the contract method 0x62d5d523.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_CAPITALIZATION() view returns(uint256)
func (_ExchangeImpl *ExchangeImplSession) CONTRACTMULTISIGREPLENISHCAPITALIZATION() (*big.Int, error) {
	return _ExchangeImpl.Contract.CONTRACTMULTISIGREPLENISHCAPITALIZATION(&_ExchangeImpl.CallOpts)
}

// CONTRACTMULTISIGREPLENISHCAPITALIZATION is a free data retrieval call binding the contract method 0x62d5d523.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_CAPITALIZATION() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCallerSession) CONTRACTMULTISIGREPLENISHCAPITALIZATION() (*big.Int, error) {
	return _ExchangeImpl.Contract.CONTRACTMULTISIGREPLENISHCAPITALIZATION(&_ExchangeImpl.CallOpts)
}

// CONTRACTMULTISIGREPLENISHCERTIFICATE is a free data retrieval call binding the contract method 0xfd19b7e1.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_CERTIFICATE() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCaller) CONTRACTMULTISIGREPLENISHCERTIFICATE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ExchangeImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REPLENISH_CERTIFICATE")
	return *ret0, err
}

// CONTRACTMULTISIGREPLENISHCERTIFICATE is a free data retrieval call binding the contract method 0xfd19b7e1.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_CERTIFICATE() view returns(uint256)
func (_ExchangeImpl *ExchangeImplSession) CONTRACTMULTISIGREPLENISHCERTIFICATE() (*big.Int, error) {
	return _ExchangeImpl.Contract.CONTRACTMULTISIGREPLENISHCERTIFICATE(&_ExchangeImpl.CallOpts)
}

// CONTRACTMULTISIGREPLENISHCERTIFICATE is a free data retrieval call binding the contract method 0xfd19b7e1.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_CERTIFICATE() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCallerSession) CONTRACTMULTISIGREPLENISHCERTIFICATE() (*big.Int, error) {
	return _ExchangeImpl.Contract.CONTRACTMULTISIGREPLENISHCERTIFICATE(&_ExchangeImpl.CallOpts)
}

// CONTRACTMULTISIGREPLENISHEMISSION is a free data retrieval call binding the contract method 0x0cb59fde.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_EMISSION() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCaller) CONTRACTMULTISIGREPLENISHEMISSION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ExchangeImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REPLENISH_EMISSION")
	return *ret0, err
}

// CONTRACTMULTISIGREPLENISHEMISSION is a free data retrieval call binding the contract method 0x0cb59fde.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_EMISSION() view returns(uint256)
func (_ExchangeImpl *ExchangeImplSession) CONTRACTMULTISIGREPLENISHEMISSION() (*big.Int, error) {
	return _ExchangeImpl.Contract.CONTRACTMULTISIGREPLENISHEMISSION(&_ExchangeImpl.CallOpts)
}

// CONTRACTMULTISIGREPLENISHEMISSION is a free data retrieval call binding the contract method 0x0cb59fde.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_EMISSION() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCallerSession) CONTRACTMULTISIGREPLENISHEMISSION() (*big.Int, error) {
	return _ExchangeImpl.Contract.CONTRACTMULTISIGREPLENISHEMISSION(&_ExchangeImpl.CallOpts)
}

// CONTRACTMULTISIGSETADMINROLE is a free data retrieval call binding the contract method 0x8f209ab7.
//
// Solidity: function CONTRACT_MULTISIG_SET_ADMIN_ROLE() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCaller) CONTRACTMULTISIGSETADMINROLE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ExchangeImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_SET_ADMIN_ROLE")
	return *ret0, err
}

// CONTRACTMULTISIGSETADMINROLE is a free data retrieval call binding the contract method 0x8f209ab7.
//
// Solidity: function CONTRACT_MULTISIG_SET_ADMIN_ROLE() view returns(uint256)
func (_ExchangeImpl *ExchangeImplSession) CONTRACTMULTISIGSETADMINROLE() (*big.Int, error) {
	return _ExchangeImpl.Contract.CONTRACTMULTISIGSETADMINROLE(&_ExchangeImpl.CallOpts)
}

// CONTRACTMULTISIGSETADMINROLE is a free data retrieval call binding the contract method 0x8f209ab7.
//
// Solidity: function CONTRACT_MULTISIG_SET_ADMIN_ROLE() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCallerSession) CONTRACTMULTISIGSETADMINROLE() (*big.Int, error) {
	return _ExchangeImpl.Contract.CONTRACTMULTISIGSETADMINROLE(&_ExchangeImpl.CallOpts)
}

// CONTRACTMULTISIGSETPERMISSION is a free data retrieval call binding the contract method 0xb0e91914.
//
// Solidity: function CONTRACT_MULTISIG_SET_PERMISSION() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCaller) CONTRACTMULTISIGSETPERMISSION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ExchangeImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_SET_PERMISSION")
	return *ret0, err
}

// CONTRACTMULTISIGSETPERMISSION is a free data retrieval call binding the contract method 0xb0e91914.
//
// Solidity: function CONTRACT_MULTISIG_SET_PERMISSION() view returns(uint256)
func (_ExchangeImpl *ExchangeImplSession) CONTRACTMULTISIGSETPERMISSION() (*big.Int, error) {
	return _ExchangeImpl.Contract.CONTRACTMULTISIGSETPERMISSION(&_ExchangeImpl.CallOpts)
}

// CONTRACTMULTISIGSETPERMISSION is a free data retrieval call binding the contract method 0xb0e91914.
//
// Solidity: function CONTRACT_MULTISIG_SET_PERMISSION() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCallerSession) CONTRACTMULTISIGSETPERMISSION() (*big.Int, error) {
	return _ExchangeImpl.Contract.CONTRACTMULTISIGSETPERMISSION(&_ExchangeImpl.CallOpts)
}

// CONTRACTMULTISIGSTORAGE is a free data retrieval call binding the contract method 0xe7e1a303.
//
// Solidity: function CONTRACT_MULTISIG_STORAGE() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCaller) CONTRACTMULTISIGSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ExchangeImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_STORAGE")
	return *ret0, err
}

// CONTRACTMULTISIGSTORAGE is a free data retrieval call binding the contract method 0xe7e1a303.
//
// Solidity: function CONTRACT_MULTISIG_STORAGE() view returns(uint256)
func (_ExchangeImpl *ExchangeImplSession) CONTRACTMULTISIGSTORAGE() (*big.Int, error) {
	return _ExchangeImpl.Contract.CONTRACTMULTISIGSTORAGE(&_ExchangeImpl.CallOpts)
}

// CONTRACTMULTISIGSTORAGE is a free data retrieval call binding the contract method 0xe7e1a303.
//
// Solidity: function CONTRACT_MULTISIG_STORAGE() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCallerSession) CONTRACTMULTISIGSTORAGE() (*big.Int, error) {
	return _ExchangeImpl.Contract.CONTRACTMULTISIGSTORAGE(&_ExchangeImpl.CallOpts)
}

// CONTRACTOTC is a free data retrieval call binding the contract method 0x5eceac70.
//
// Solidity: function CONTRACT_OTC() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCaller) CONTRACTOTC(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ExchangeImpl.contract.Call(opts, out, "CONTRACT_OTC")
	return *ret0, err
}

// CONTRACTOTC is a free data retrieval call binding the contract method 0x5eceac70.
//
// Solidity: function CONTRACT_OTC() view returns(uint256)
func (_ExchangeImpl *ExchangeImplSession) CONTRACTOTC() (*big.Int, error) {
	return _ExchangeImpl.Contract.CONTRACTOTC(&_ExchangeImpl.CallOpts)
}

// CONTRACTOTC is a free data retrieval call binding the contract method 0x5eceac70.
//
// Solidity: function CONTRACT_OTC() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCallerSession) CONTRACTOTC() (*big.Int, error) {
	return _ExchangeImpl.Contract.CONTRACTOTC(&_ExchangeImpl.CallOpts)
}

// CONTRACTOTCSTORAGE is a free data retrieval call binding the contract method 0x825b476f.
//
// Solidity: function CONTRACT_OTC_STORAGE() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCaller) CONTRACTOTCSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ExchangeImpl.contract.Call(opts, out, "CONTRACT_OTC_STORAGE")
	return *ret0, err
}

// CONTRACTOTCSTORAGE is a free data retrieval call binding the contract method 0x825b476f.
//
// Solidity: function CONTRACT_OTC_STORAGE() view returns(uint256)
func (_ExchangeImpl *ExchangeImplSession) CONTRACTOTCSTORAGE() (*big.Int, error) {
	return _ExchangeImpl.Contract.CONTRACTOTCSTORAGE(&_ExchangeImpl.CallOpts)
}

// CONTRACTOTCSTORAGE is a free data retrieval call binding the contract method 0x825b476f.
//
// Solidity: function CONTRACT_OTC_STORAGE() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCallerSession) CONTRACTOTCSTORAGE() (*big.Int, error) {
	return _ExchangeImpl.Contract.CONTRACTOTCSTORAGE(&_ExchangeImpl.CallOpts)
}

// CONTRACTREPLENISHCAPITALIZATION is a free data retrieval call binding the contract method 0xff5e09d8.
//
// Solidity: function CONTRACT_REPLENISH_CAPITALIZATION() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCaller) CONTRACTREPLENISHCAPITALIZATION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ExchangeImpl.contract.Call(opts, out, "CONTRACT_REPLENISH_CAPITALIZATION")
	return *ret0, err
}

// CONTRACTREPLENISHCAPITALIZATION is a free data retrieval call binding the contract method 0xff5e09d8.
//
// Solidity: function CONTRACT_REPLENISH_CAPITALIZATION() view returns(uint256)
func (_ExchangeImpl *ExchangeImplSession) CONTRACTREPLENISHCAPITALIZATION() (*big.Int, error) {
	return _ExchangeImpl.Contract.CONTRACTREPLENISHCAPITALIZATION(&_ExchangeImpl.CallOpts)
}

// CONTRACTREPLENISHCAPITALIZATION is a free data retrieval call binding the contract method 0xff5e09d8.
//
// Solidity: function CONTRACT_REPLENISH_CAPITALIZATION() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCallerSession) CONTRACTREPLENISHCAPITALIZATION() (*big.Int, error) {
	return _ExchangeImpl.Contract.CONTRACTREPLENISHCAPITALIZATION(&_ExchangeImpl.CallOpts)
}

// CONTRACTRESERVEDALIASSTORAGE is a free data retrieval call binding the contract method 0x70a94aa7.
//
// Solidity: function CONTRACT_RESERVED_ALIAS_STORAGE() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCaller) CONTRACTRESERVEDALIASSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ExchangeImpl.contract.Call(opts, out, "CONTRACT_RESERVED_ALIAS_STORAGE")
	return *ret0, err
}

// CONTRACTRESERVEDALIASSTORAGE is a free data retrieval call binding the contract method 0x70a94aa7.
//
// Solidity: function CONTRACT_RESERVED_ALIAS_STORAGE() view returns(uint256)
func (_ExchangeImpl *ExchangeImplSession) CONTRACTRESERVEDALIASSTORAGE() (*big.Int, error) {
	return _ExchangeImpl.Contract.CONTRACTRESERVEDALIASSTORAGE(&_ExchangeImpl.CallOpts)
}

// CONTRACTRESERVEDALIASSTORAGE is a free data retrieval call binding the contract method 0x70a94aa7.
//
// Solidity: function CONTRACT_RESERVED_ALIAS_STORAGE() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCallerSession) CONTRACTRESERVEDALIASSTORAGE() (*big.Int, error) {
	return _ExchangeImpl.Contract.CONTRACTRESERVEDALIASSTORAGE(&_ExchangeImpl.CallOpts)
}

// CONTRACTROLESTORAGE is a free data retrieval call binding the contract method 0xb938235b.
//
// Solidity: function CONTRACT_ROLE_STORAGE() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCaller) CONTRACTROLESTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ExchangeImpl.contract.Call(opts, out, "CONTRACT_ROLE_STORAGE")
	return *ret0, err
}

// CONTRACTROLESTORAGE is a free data retrieval call binding the contract method 0xb938235b.
//
// Solidity: function CONTRACT_ROLE_STORAGE() view returns(uint256)
func (_ExchangeImpl *ExchangeImplSession) CONTRACTROLESTORAGE() (*big.Int, error) {
	return _ExchangeImpl.Contract.CONTRACTROLESTORAGE(&_ExchangeImpl.CallOpts)
}

// CONTRACTROLESTORAGE is a free data retrieval call binding the contract method 0xb938235b.
//
// Solidity: function CONTRACT_ROLE_STORAGE() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCallerSession) CONTRACTROLESTORAGE() (*big.Int, error) {
	return _ExchangeImpl.Contract.CONTRACTROLESTORAGE(&_ExchangeImpl.CallOpts)
}

// CONTRACTSAFETRANSFERSTORAGE is a free data retrieval call binding the contract method 0xe7e45392.
//
// Solidity: function CONTRACT_SAFE_TRANSFER_STORAGE() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCaller) CONTRACTSAFETRANSFERSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ExchangeImpl.contract.Call(opts, out, "CONTRACT_SAFE_TRANSFER_STORAGE")
	return *ret0, err
}

// CONTRACTSAFETRANSFERSTORAGE is a free data retrieval call binding the contract method 0xe7e45392.
//
// Solidity: function CONTRACT_SAFE_TRANSFER_STORAGE() view returns(uint256)
func (_ExchangeImpl *ExchangeImplSession) CONTRACTSAFETRANSFERSTORAGE() (*big.Int, error) {
	return _ExchangeImpl.Contract.CONTRACTSAFETRANSFERSTORAGE(&_ExchangeImpl.CallOpts)
}

// CONTRACTSAFETRANSFERSTORAGE is a free data retrieval call binding the contract method 0xe7e45392.
//
// Solidity: function CONTRACT_SAFE_TRANSFER_STORAGE() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCallerSession) CONTRACTSAFETRANSFERSTORAGE() (*big.Int, error) {
	return _ExchangeImpl.Contract.CONTRACTSAFETRANSFERSTORAGE(&_ExchangeImpl.CallOpts)
}

// CONTRACTTRANSFER is a free data retrieval call binding the contract method 0xef94fe38.
//
// Solidity: function CONTRACT_TRANSFER() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCaller) CONTRACTTRANSFER(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ExchangeImpl.contract.Call(opts, out, "CONTRACT_TRANSFER")
	return *ret0, err
}

// CONTRACTTRANSFER is a free data retrieval call binding the contract method 0xef94fe38.
//
// Solidity: function CONTRACT_TRANSFER() view returns(uint256)
func (_ExchangeImpl *ExchangeImplSession) CONTRACTTRANSFER() (*big.Int, error) {
	return _ExchangeImpl.Contract.CONTRACTTRANSFER(&_ExchangeImpl.CallOpts)
}

// CONTRACTTRANSFER is a free data retrieval call binding the contract method 0xef94fe38.
//
// Solidity: function CONTRACT_TRANSFER() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCallerSession) CONTRACTTRANSFER() (*big.Int, error) {
	return _ExchangeImpl.Contract.CONTRACTTRANSFER(&_ExchangeImpl.CallOpts)
}

// CONTRACTTRANSFERFROMCAPITALIZATION is a free data retrieval call binding the contract method 0xc8f8decb.
//
// Solidity: function CONTRACT_TRANSFER_FROM_CAPITALIZATION() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCaller) CONTRACTTRANSFERFROMCAPITALIZATION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ExchangeImpl.contract.Call(opts, out, "CONTRACT_TRANSFER_FROM_CAPITALIZATION")
	return *ret0, err
}

// CONTRACTTRANSFERFROMCAPITALIZATION is a free data retrieval call binding the contract method 0xc8f8decb.
//
// Solidity: function CONTRACT_TRANSFER_FROM_CAPITALIZATION() view returns(uint256)
func (_ExchangeImpl *ExchangeImplSession) CONTRACTTRANSFERFROMCAPITALIZATION() (*big.Int, error) {
	return _ExchangeImpl.Contract.CONTRACTTRANSFERFROMCAPITALIZATION(&_ExchangeImpl.CallOpts)
}

// CONTRACTTRANSFERFROMCAPITALIZATION is a free data retrieval call binding the contract method 0xc8f8decb.
//
// Solidity: function CONTRACT_TRANSFER_FROM_CAPITALIZATION() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCallerSession) CONTRACTTRANSFERFROMCAPITALIZATION() (*big.Int, error) {
	return _ExchangeImpl.Contract.CONTRACTTRANSFERFROMCAPITALIZATION(&_ExchangeImpl.CallOpts)
}

// PERMISSIONACCESSRESTORATION is a free data retrieval call binding the contract method 0x5c3053bf.
//
// Solidity: function PERMISSION_ACCESS_RESTORATION() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCaller) PERMISSIONACCESSRESTORATION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ExchangeImpl.contract.Call(opts, out, "PERMISSION_ACCESS_RESTORATION")
	return *ret0, err
}

// PERMISSIONACCESSRESTORATION is a free data retrieval call binding the contract method 0x5c3053bf.
//
// Solidity: function PERMISSION_ACCESS_RESTORATION() view returns(uint256)
func (_ExchangeImpl *ExchangeImplSession) PERMISSIONACCESSRESTORATION() (*big.Int, error) {
	return _ExchangeImpl.Contract.PERMISSIONACCESSRESTORATION(&_ExchangeImpl.CallOpts)
}

// PERMISSIONACCESSRESTORATION is a free data retrieval call binding the contract method 0x5c3053bf.
//
// Solidity: function PERMISSION_ACCESS_RESTORATION() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCallerSession) PERMISSIONACCESSRESTORATION() (*big.Int, error) {
	return _ExchangeImpl.Contract.PERMISSIONACCESSRESTORATION(&_ExchangeImpl.CallOpts)
}

// PERMISSIONADMINSTORAGE is a free data retrieval call binding the contract method 0x0c9fc0db.
//
// Solidity: function PERMISSION_ADMIN_STORAGE() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCaller) PERMISSIONADMINSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ExchangeImpl.contract.Call(opts, out, "PERMISSION_ADMIN_STORAGE")
	return *ret0, err
}

// PERMISSIONADMINSTORAGE is a free data retrieval call binding the contract method 0x0c9fc0db.
//
// Solidity: function PERMISSION_ADMIN_STORAGE() view returns(uint256)
func (_ExchangeImpl *ExchangeImplSession) PERMISSIONADMINSTORAGE() (*big.Int, error) {
	return _ExchangeImpl.Contract.PERMISSIONADMINSTORAGE(&_ExchangeImpl.CallOpts)
}

// PERMISSIONADMINSTORAGE is a free data retrieval call binding the contract method 0x0c9fc0db.
//
// Solidity: function PERMISSION_ADMIN_STORAGE() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCallerSession) PERMISSIONADMINSTORAGE() (*big.Int, error) {
	return _ExchangeImpl.Contract.PERMISSIONADMINSTORAGE(&_ExchangeImpl.CallOpts)
}

// PERMISSIONCONTRACTSSTORAGE is a free data retrieval call binding the contract method 0x1516e83b.
//
// Solidity: function PERMISSION_CONTRACTS_STORAGE() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCaller) PERMISSIONCONTRACTSSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ExchangeImpl.contract.Call(opts, out, "PERMISSION_CONTRACTS_STORAGE")
	return *ret0, err
}

// PERMISSIONCONTRACTSSTORAGE is a free data retrieval call binding the contract method 0x1516e83b.
//
// Solidity: function PERMISSION_CONTRACTS_STORAGE() view returns(uint256)
func (_ExchangeImpl *ExchangeImplSession) PERMISSIONCONTRACTSSTORAGE() (*big.Int, error) {
	return _ExchangeImpl.Contract.PERMISSIONCONTRACTSSTORAGE(&_ExchangeImpl.CallOpts)
}

// PERMISSIONCONTRACTSSTORAGE is a free data retrieval call binding the contract method 0x1516e83b.
//
// Solidity: function PERMISSION_CONTRACTS_STORAGE() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCallerSession) PERMISSIONCONTRACTSSTORAGE() (*big.Int, error) {
	return _ExchangeImpl.Contract.PERMISSIONCONTRACTSSTORAGE(&_ExchangeImpl.CallOpts)
}

// PERMISSIONCONTRACTDEPLOY is a free data retrieval call binding the contract method 0x257a4a04.
//
// Solidity: function PERMISSION_CONTRACT_DEPLOY() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCaller) PERMISSIONCONTRACTDEPLOY(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ExchangeImpl.contract.Call(opts, out, "PERMISSION_CONTRACT_DEPLOY")
	return *ret0, err
}

// PERMISSIONCONTRACTDEPLOY is a free data retrieval call binding the contract method 0x257a4a04.
//
// Solidity: function PERMISSION_CONTRACT_DEPLOY() view returns(uint256)
func (_ExchangeImpl *ExchangeImplSession) PERMISSIONCONTRACTDEPLOY() (*big.Int, error) {
	return _ExchangeImpl.Contract.PERMISSIONCONTRACTDEPLOY(&_ExchangeImpl.CallOpts)
}

// PERMISSIONCONTRACTDEPLOY is a free data retrieval call binding the contract method 0x257a4a04.
//
// Solidity: function PERMISSION_CONTRACT_DEPLOY() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCallerSession) PERMISSIONCONTRACTDEPLOY() (*big.Int, error) {
	return _ExchangeImpl.Contract.PERMISSIONCONTRACTDEPLOY(&_ExchangeImpl.CallOpts)
}

// PERMISSIONEMISSION is a free data retrieval call binding the contract method 0x4ccd92d1.
//
// Solidity: function PERMISSION_EMISSION() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCaller) PERMISSIONEMISSION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ExchangeImpl.contract.Call(opts, out, "PERMISSION_EMISSION")
	return *ret0, err
}

// PERMISSIONEMISSION is a free data retrieval call binding the contract method 0x4ccd92d1.
//
// Solidity: function PERMISSION_EMISSION() view returns(uint256)
func (_ExchangeImpl *ExchangeImplSession) PERMISSIONEMISSION() (*big.Int, error) {
	return _ExchangeImpl.Contract.PERMISSIONEMISSION(&_ExchangeImpl.CallOpts)
}

// PERMISSIONEMISSION is a free data retrieval call binding the contract method 0x4ccd92d1.
//
// Solidity: function PERMISSION_EMISSION() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCallerSession) PERMISSIONEMISSION() (*big.Int, error) {
	return _ExchangeImpl.Contract.PERMISSIONEMISSION(&_ExchangeImpl.CallOpts)
}

// PERMISSIONFEESSTORAGE is a free data retrieval call binding the contract method 0xe837eb87.
//
// Solidity: function PERMISSION_FEES_STORAGE() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCaller) PERMISSIONFEESSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ExchangeImpl.contract.Call(opts, out, "PERMISSION_FEES_STORAGE")
	return *ret0, err
}

// PERMISSIONFEESSTORAGE is a free data retrieval call binding the contract method 0xe837eb87.
//
// Solidity: function PERMISSION_FEES_STORAGE() view returns(uint256)
func (_ExchangeImpl *ExchangeImplSession) PERMISSIONFEESSTORAGE() (*big.Int, error) {
	return _ExchangeImpl.Contract.PERMISSIONFEESSTORAGE(&_ExchangeImpl.CallOpts)
}

// PERMISSIONFEESSTORAGE is a free data retrieval call binding the contract method 0xe837eb87.
//
// Solidity: function PERMISSION_FEES_STORAGE() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCallerSession) PERMISSIONFEESSTORAGE() (*big.Int, error) {
	return _ExchangeImpl.Contract.PERMISSIONFEESSTORAGE(&_ExchangeImpl.CallOpts)
}

// PERMISSIONMULTISIGSTORAGE is a free data retrieval call binding the contract method 0xf1bb7aad.
//
// Solidity: function PERMISSION_MULTISIG_STORAGE() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCaller) PERMISSIONMULTISIGSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ExchangeImpl.contract.Call(opts, out, "PERMISSION_MULTISIG_STORAGE")
	return *ret0, err
}

// PERMISSIONMULTISIGSTORAGE is a free data retrieval call binding the contract method 0xf1bb7aad.
//
// Solidity: function PERMISSION_MULTISIG_STORAGE() view returns(uint256)
func (_ExchangeImpl *ExchangeImplSession) PERMISSIONMULTISIGSTORAGE() (*big.Int, error) {
	return _ExchangeImpl.Contract.PERMISSIONMULTISIGSTORAGE(&_ExchangeImpl.CallOpts)
}

// PERMISSIONMULTISIGSTORAGE is a free data retrieval call binding the contract method 0xf1bb7aad.
//
// Solidity: function PERMISSION_MULTISIG_STORAGE() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCallerSession) PERMISSIONMULTISIGSTORAGE() (*big.Int, error) {
	return _ExchangeImpl.Contract.PERMISSIONMULTISIGSTORAGE(&_ExchangeImpl.CallOpts)
}

// PERMISSIONROLESSTORAGE is a free data retrieval call binding the contract method 0x12a16032.
//
// Solidity: function PERMISSION_ROLES_STORAGE() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCaller) PERMISSIONROLESSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ExchangeImpl.contract.Call(opts, out, "PERMISSION_ROLES_STORAGE")
	return *ret0, err
}

// PERMISSIONROLESSTORAGE is a free data retrieval call binding the contract method 0x12a16032.
//
// Solidity: function PERMISSION_ROLES_STORAGE() view returns(uint256)
func (_ExchangeImpl *ExchangeImplSession) PERMISSIONROLESSTORAGE() (*big.Int, error) {
	return _ExchangeImpl.Contract.PERMISSIONROLESSTORAGE(&_ExchangeImpl.CallOpts)
}

// PERMISSIONROLESSTORAGE is a free data retrieval call binding the contract method 0x12a16032.
//
// Solidity: function PERMISSION_ROLES_STORAGE() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCallerSession) PERMISSIONROLESSTORAGE() (*big.Int, error) {
	return _ExchangeImpl.Contract.PERMISSIONROLESSTORAGE(&_ExchangeImpl.CallOpts)
}

// ROLEACCESSVERIFIER is a free data retrieval call binding the contract method 0xc5405f92.
//
// Solidity: function ROLE_ACCESS_VERIFIER() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCaller) ROLEACCESSVERIFIER(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ExchangeImpl.contract.Call(opts, out, "ROLE_ACCESS_VERIFIER")
	return *ret0, err
}

// ROLEACCESSVERIFIER is a free data retrieval call binding the contract method 0xc5405f92.
//
// Solidity: function ROLE_ACCESS_VERIFIER() view returns(uint256)
func (_ExchangeImpl *ExchangeImplSession) ROLEACCESSVERIFIER() (*big.Int, error) {
	return _ExchangeImpl.Contract.ROLEACCESSVERIFIER(&_ExchangeImpl.CallOpts)
}

// ROLEACCESSVERIFIER is a free data retrieval call binding the contract method 0xc5405f92.
//
// Solidity: function ROLE_ACCESS_VERIFIER() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCallerSession) ROLEACCESSVERIFIER() (*big.Int, error) {
	return _ExchangeImpl.Contract.ROLEACCESSVERIFIER(&_ExchangeImpl.CallOpts)
}

// ROLEBITBONISSUEVERIFIER is a free data retrieval call binding the contract method 0xf0a012a6.
//
// Solidity: function ROLE_BITBON_ISSUE_VERIFIER() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCaller) ROLEBITBONISSUEVERIFIER(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ExchangeImpl.contract.Call(opts, out, "ROLE_BITBON_ISSUE_VERIFIER")
	return *ret0, err
}

// ROLEBITBONISSUEVERIFIER is a free data retrieval call binding the contract method 0xf0a012a6.
//
// Solidity: function ROLE_BITBON_ISSUE_VERIFIER() view returns(uint256)
func (_ExchangeImpl *ExchangeImplSession) ROLEBITBONISSUEVERIFIER() (*big.Int, error) {
	return _ExchangeImpl.Contract.ROLEBITBONISSUEVERIFIER(&_ExchangeImpl.CallOpts)
}

// ROLEBITBONISSUEVERIFIER is a free data retrieval call binding the contract method 0xf0a012a6.
//
// Solidity: function ROLE_BITBON_ISSUE_VERIFIER() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCallerSession) ROLEBITBONISSUEVERIFIER() (*big.Int, error) {
	return _ExchangeImpl.Contract.ROLEBITBONISSUEVERIFIER(&_ExchangeImpl.CallOpts)
}

// ROLECOMMISSIONVERIFIER is a free data retrieval call binding the contract method 0x8ab3de8c.
//
// Solidity: function ROLE_COMMISSION_VERIFIER() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCaller) ROLECOMMISSIONVERIFIER(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ExchangeImpl.contract.Call(opts, out, "ROLE_COMMISSION_VERIFIER")
	return *ret0, err
}

// ROLECOMMISSIONVERIFIER is a free data retrieval call binding the contract method 0x8ab3de8c.
//
// Solidity: function ROLE_COMMISSION_VERIFIER() view returns(uint256)
func (_ExchangeImpl *ExchangeImplSession) ROLECOMMISSIONVERIFIER() (*big.Int, error) {
	return _ExchangeImpl.Contract.ROLECOMMISSIONVERIFIER(&_ExchangeImpl.CallOpts)
}

// ROLECOMMISSIONVERIFIER is a free data retrieval call binding the contract method 0x8ab3de8c.
//
// Solidity: function ROLE_COMMISSION_VERIFIER() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCallerSession) ROLECOMMISSIONVERIFIER() (*big.Int, error) {
	return _ExchangeImpl.Contract.ROLECOMMISSIONVERIFIER(&_ExchangeImpl.CallOpts)
}

// ROLEDEPLOYADMIN is a free data retrieval call binding the contract method 0x667bf6e5.
//
// Solidity: function ROLE_DEPLOY_ADMIN() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCaller) ROLEDEPLOYADMIN(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ExchangeImpl.contract.Call(opts, out, "ROLE_DEPLOY_ADMIN")
	return *ret0, err
}

// ROLEDEPLOYADMIN is a free data retrieval call binding the contract method 0x667bf6e5.
//
// Solidity: function ROLE_DEPLOY_ADMIN() view returns(uint256)
func (_ExchangeImpl *ExchangeImplSession) ROLEDEPLOYADMIN() (*big.Int, error) {
	return _ExchangeImpl.Contract.ROLEDEPLOYADMIN(&_ExchangeImpl.CallOpts)
}

// ROLEDEPLOYADMIN is a free data retrieval call binding the contract method 0x667bf6e5.
//
// Solidity: function ROLE_DEPLOY_ADMIN() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCallerSession) ROLEDEPLOYADMIN() (*big.Int, error) {
	return _ExchangeImpl.Contract.ROLEDEPLOYADMIN(&_ExchangeImpl.CallOpts)
}

// ROLEPERMISSIONADMIN is a free data retrieval call binding the contract method 0x1b239996.
//
// Solidity: function ROLE_PERMISSION_ADMIN() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCaller) ROLEPERMISSIONADMIN(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ExchangeImpl.contract.Call(opts, out, "ROLE_PERMISSION_ADMIN")
	return *ret0, err
}

// ROLEPERMISSIONADMIN is a free data retrieval call binding the contract method 0x1b239996.
//
// Solidity: function ROLE_PERMISSION_ADMIN() view returns(uint256)
func (_ExchangeImpl *ExchangeImplSession) ROLEPERMISSIONADMIN() (*big.Int, error) {
	return _ExchangeImpl.Contract.ROLEPERMISSIONADMIN(&_ExchangeImpl.CallOpts)
}

// ROLEPERMISSIONADMIN is a free data retrieval call binding the contract method 0x1b239996.
//
// Solidity: function ROLE_PERMISSION_ADMIN() view returns(uint256)
func (_ExchangeImpl *ExchangeImplCallerSession) ROLEPERMISSIONADMIN() (*big.Int, error) {
	return _ExchangeImpl.Contract.ROLEPERMISSIONADMIN(&_ExchangeImpl.CallOpts)
}

// ContractStorage is a free data retrieval call binding the contract method 0x0549b27e.
//
// Solidity: function contractStorage() view returns(address)
func (_ExchangeImpl *ExchangeImplCaller) ContractStorage(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ExchangeImpl.contract.Call(opts, out, "contractStorage")
	return *ret0, err
}

// ContractStorage is a free data retrieval call binding the contract method 0x0549b27e.
//
// Solidity: function contractStorage() view returns(address)
func (_ExchangeImpl *ExchangeImplSession) ContractStorage() (common.Address, error) {
	return _ExchangeImpl.Contract.ContractStorage(&_ExchangeImpl.CallOpts)
}

// ContractStorage is a free data retrieval call binding the contract method 0x0549b27e.
//
// Solidity: function contractStorage() view returns(address)
func (_ExchangeImpl *ExchangeImplCallerSession) ContractStorage() (common.Address, error) {
	return _ExchangeImpl.Contract.ContractStorage(&_ExchangeImpl.CallOpts)
}

// GetThisContractIndex is a free data retrieval call binding the contract method 0x70c6d367.
//
// Solidity: function getThisContractIndex() pure returns(uint256)
func (_ExchangeImpl *ExchangeImplCaller) GetThisContractIndex(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ExchangeImpl.contract.Call(opts, out, "getThisContractIndex")
	return *ret0, err
}

// GetThisContractIndex is a free data retrieval call binding the contract method 0x70c6d367.
//
// Solidity: function getThisContractIndex() pure returns(uint256)
func (_ExchangeImpl *ExchangeImplSession) GetThisContractIndex() (*big.Int, error) {
	return _ExchangeImpl.Contract.GetThisContractIndex(&_ExchangeImpl.CallOpts)
}

// GetThisContractIndex is a free data retrieval call binding the contract method 0x70c6d367.
//
// Solidity: function getThisContractIndex() pure returns(uint256)
func (_ExchangeImpl *ExchangeImplCallerSession) GetThisContractIndex() (*big.Int, error) {
	return _ExchangeImpl.Contract.GetThisContractIndex(&_ExchangeImpl.CallOpts)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint8)
func (_ExchangeImpl *ExchangeImplCaller) State(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _ExchangeImpl.contract.Call(opts, out, "state")
	return *ret0, err
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint8)
func (_ExchangeImpl *ExchangeImplSession) State() (uint8, error) {
	return _ExchangeImpl.Contract.State(&_ExchangeImpl.CallOpts)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint8)
func (_ExchangeImpl *ExchangeImplCallerSession) State() (uint8, error) {
	return _ExchangeImpl.Contract.State(&_ExchangeImpl.CallOpts)
}

// CancelBitbonOrder is a paid mutator transaction binding the contract method 0x3ac822c0.
//
// Solidity: function cancelBitbonOrder(bytes32 offerHash, uint256 expectedAvailableAmount) returns()
func (_ExchangeImpl *ExchangeImplTransactor) CancelBitbonOrder(opts *bind.TransactOpts, offerHash [32]byte, expectedAvailableAmount *big.Int) (*types.Transaction, error) {
	return _ExchangeImpl.contract.Transact(opts, "cancelBitbonOrder", offerHash, expectedAvailableAmount)
}

// CancelBitbonOrder is a paid mutator transaction binding the contract method 0x3ac822c0.
//
// Solidity: function cancelBitbonOrder(bytes32 offerHash, uint256 expectedAvailableAmount) returns()
func (_ExchangeImpl *ExchangeImplSession) CancelBitbonOrder(offerHash [32]byte, expectedAvailableAmount *big.Int) (*types.Transaction, error) {
	return _ExchangeImpl.Contract.CancelBitbonOrder(&_ExchangeImpl.TransactOpts, offerHash, expectedAvailableAmount)
}

// CancelBitbonOrder is a paid mutator transaction binding the contract method 0x3ac822c0.
//
// Solidity: function cancelBitbonOrder(bytes32 offerHash, uint256 expectedAvailableAmount) returns()
func (_ExchangeImpl *ExchangeImplTransactorSession) CancelBitbonOrder(offerHash [32]byte, expectedAvailableAmount *big.Int) (*types.Transaction, error) {
	return _ExchangeImpl.Contract.CancelBitbonOrder(&_ExchangeImpl.TransactOpts, offerHash, expectedAvailableAmount)
}

// CancelMsbonOrder is a paid mutator transaction binding the contract method 0x78b80a87.
//
// Solidity: function cancelMsbonOrder(bytes32 offerHash, uint256 expectedAvailableAmount) returns()
func (_ExchangeImpl *ExchangeImplTransactor) CancelMsbonOrder(opts *bind.TransactOpts, offerHash [32]byte, expectedAvailableAmount *big.Int) (*types.Transaction, error) {
	return _ExchangeImpl.contract.Transact(opts, "cancelMsbonOrder", offerHash, expectedAvailableAmount)
}

// CancelMsbonOrder is a paid mutator transaction binding the contract method 0x78b80a87.
//
// Solidity: function cancelMsbonOrder(bytes32 offerHash, uint256 expectedAvailableAmount) returns()
func (_ExchangeImpl *ExchangeImplSession) CancelMsbonOrder(offerHash [32]byte, expectedAvailableAmount *big.Int) (*types.Transaction, error) {
	return _ExchangeImpl.Contract.CancelMsbonOrder(&_ExchangeImpl.TransactOpts, offerHash, expectedAvailableAmount)
}

// CancelMsbonOrder is a paid mutator transaction binding the contract method 0x78b80a87.
//
// Solidity: function cancelMsbonOrder(bytes32 offerHash, uint256 expectedAvailableAmount) returns()
func (_ExchangeImpl *ExchangeImplTransactorSession) CancelMsbonOrder(offerHash [32]byte, expectedAvailableAmount *big.Int) (*types.Transaction, error) {
	return _ExchangeImpl.Contract.CancelMsbonOrder(&_ExchangeImpl.TransactOpts, offerHash, expectedAvailableAmount)
}

// FillBitbonOrder is a paid mutator transaction binding the contract method 0x79b1ae4b.
//
// Solidity: function fillBitbonOrder(address fillAssets, address wantAssets, bytes32 offerHash, uint256 amountToTake) returns()
func (_ExchangeImpl *ExchangeImplTransactor) FillBitbonOrder(opts *bind.TransactOpts, fillAssets common.Address, wantAssets common.Address, offerHash [32]byte, amountToTake *big.Int) (*types.Transaction, error) {
	return _ExchangeImpl.contract.Transact(opts, "fillBitbonOrder", fillAssets, wantAssets, offerHash, amountToTake)
}

// FillBitbonOrder is a paid mutator transaction binding the contract method 0x79b1ae4b.
//
// Solidity: function fillBitbonOrder(address fillAssets, address wantAssets, bytes32 offerHash, uint256 amountToTake) returns()
func (_ExchangeImpl *ExchangeImplSession) FillBitbonOrder(fillAssets common.Address, wantAssets common.Address, offerHash [32]byte, amountToTake *big.Int) (*types.Transaction, error) {
	return _ExchangeImpl.Contract.FillBitbonOrder(&_ExchangeImpl.TransactOpts, fillAssets, wantAssets, offerHash, amountToTake)
}

// FillBitbonOrder is a paid mutator transaction binding the contract method 0x79b1ae4b.
//
// Solidity: function fillBitbonOrder(address fillAssets, address wantAssets, bytes32 offerHash, uint256 amountToTake) returns()
func (_ExchangeImpl *ExchangeImplTransactorSession) FillBitbonOrder(fillAssets common.Address, wantAssets common.Address, offerHash [32]byte, amountToTake *big.Int) (*types.Transaction, error) {
	return _ExchangeImpl.Contract.FillBitbonOrder(&_ExchangeImpl.TransactOpts, fillAssets, wantAssets, offerHash, amountToTake)
}

// FillMsbonOrder is a paid mutator transaction binding the contract method 0x41b8d73e.
//
// Solidity: function fillMsbonOrder(address fillAssets, address wantAssets, bytes32 offerHash, uint256 amountToTake) returns()
func (_ExchangeImpl *ExchangeImplTransactor) FillMsbonOrder(opts *bind.TransactOpts, fillAssets common.Address, wantAssets common.Address, offerHash [32]byte, amountToTake *big.Int) (*types.Transaction, error) {
	return _ExchangeImpl.contract.Transact(opts, "fillMsbonOrder", fillAssets, wantAssets, offerHash, amountToTake)
}

// FillMsbonOrder is a paid mutator transaction binding the contract method 0x41b8d73e.
//
// Solidity: function fillMsbonOrder(address fillAssets, address wantAssets, bytes32 offerHash, uint256 amountToTake) returns()
func (_ExchangeImpl *ExchangeImplSession) FillMsbonOrder(fillAssets common.Address, wantAssets common.Address, offerHash [32]byte, amountToTake *big.Int) (*types.Transaction, error) {
	return _ExchangeImpl.Contract.FillMsbonOrder(&_ExchangeImpl.TransactOpts, fillAssets, wantAssets, offerHash, amountToTake)
}

// FillMsbonOrder is a paid mutator transaction binding the contract method 0x41b8d73e.
//
// Solidity: function fillMsbonOrder(address fillAssets, address wantAssets, bytes32 offerHash, uint256 amountToTake) returns()
func (_ExchangeImpl *ExchangeImplTransactorSession) FillMsbonOrder(fillAssets common.Address, wantAssets common.Address, offerHash [32]byte, amountToTake *big.Int) (*types.Transaction, error) {
	return _ExchangeImpl.Contract.FillMsbonOrder(&_ExchangeImpl.TransactOpts, fillAssets, wantAssets, offerHash, amountToTake)
}

// MakeBitbonOrder is a paid mutator transaction binding the contract method 0x4b8b0e16.
//
// Solidity: function makeBitbonOrder(address offerAssets, address wantAssets, uint256 offerAmount, uint256 wantAmount, uint64 nonce) returns()
func (_ExchangeImpl *ExchangeImplTransactor) MakeBitbonOrder(opts *bind.TransactOpts, offerAssets common.Address, wantAssets common.Address, offerAmount *big.Int, wantAmount *big.Int, nonce uint64) (*types.Transaction, error) {
	return _ExchangeImpl.contract.Transact(opts, "makeBitbonOrder", offerAssets, wantAssets, offerAmount, wantAmount, nonce)
}

// MakeBitbonOrder is a paid mutator transaction binding the contract method 0x4b8b0e16.
//
// Solidity: function makeBitbonOrder(address offerAssets, address wantAssets, uint256 offerAmount, uint256 wantAmount, uint64 nonce) returns()
func (_ExchangeImpl *ExchangeImplSession) MakeBitbonOrder(offerAssets common.Address, wantAssets common.Address, offerAmount *big.Int, wantAmount *big.Int, nonce uint64) (*types.Transaction, error) {
	return _ExchangeImpl.Contract.MakeBitbonOrder(&_ExchangeImpl.TransactOpts, offerAssets, wantAssets, offerAmount, wantAmount, nonce)
}

// MakeBitbonOrder is a paid mutator transaction binding the contract method 0x4b8b0e16.
//
// Solidity: function makeBitbonOrder(address offerAssets, address wantAssets, uint256 offerAmount, uint256 wantAmount, uint64 nonce) returns()
func (_ExchangeImpl *ExchangeImplTransactorSession) MakeBitbonOrder(offerAssets common.Address, wantAssets common.Address, offerAmount *big.Int, wantAmount *big.Int, nonce uint64) (*types.Transaction, error) {
	return _ExchangeImpl.Contract.MakeBitbonOrder(&_ExchangeImpl.TransactOpts, offerAssets, wantAssets, offerAmount, wantAmount, nonce)
}

// MakeMsbonOrder is a paid mutator transaction binding the contract method 0xb80ad34c.
//
// Solidity: function makeMsbonOrder(address offerAssets, address wantAssets, uint256 offerAmount, uint256 wantAmount, uint64 nonce) returns()
func (_ExchangeImpl *ExchangeImplTransactor) MakeMsbonOrder(opts *bind.TransactOpts, offerAssets common.Address, wantAssets common.Address, offerAmount *big.Int, wantAmount *big.Int, nonce uint64) (*types.Transaction, error) {
	return _ExchangeImpl.contract.Transact(opts, "makeMsbonOrder", offerAssets, wantAssets, offerAmount, wantAmount, nonce)
}

// MakeMsbonOrder is a paid mutator transaction binding the contract method 0xb80ad34c.
//
// Solidity: function makeMsbonOrder(address offerAssets, address wantAssets, uint256 offerAmount, uint256 wantAmount, uint64 nonce) returns()
func (_ExchangeImpl *ExchangeImplSession) MakeMsbonOrder(offerAssets common.Address, wantAssets common.Address, offerAmount *big.Int, wantAmount *big.Int, nonce uint64) (*types.Transaction, error) {
	return _ExchangeImpl.Contract.MakeMsbonOrder(&_ExchangeImpl.TransactOpts, offerAssets, wantAssets, offerAmount, wantAmount, nonce)
}

// MakeMsbonOrder is a paid mutator transaction binding the contract method 0xb80ad34c.
//
// Solidity: function makeMsbonOrder(address offerAssets, address wantAssets, uint256 offerAmount, uint256 wantAmount, uint64 nonce) returns()
func (_ExchangeImpl *ExchangeImplTransactorSession) MakeMsbonOrder(offerAssets common.Address, wantAssets common.Address, offerAmount *big.Int, wantAmount *big.Int, nonce uint64) (*types.Transaction, error) {
	return _ExchangeImpl.Contract.MakeMsbonOrder(&_ExchangeImpl.TransactOpts, offerAssets, wantAssets, offerAmount, wantAmount, nonce)
}

// SetState is a paid mutator transaction binding the contract method 0x56de96db.
//
// Solidity: function setState(uint8 newState) returns()
func (_ExchangeImpl *ExchangeImplTransactor) SetState(opts *bind.TransactOpts, newState uint8) (*types.Transaction, error) {
	return _ExchangeImpl.contract.Transact(opts, "setState", newState)
}

// SetState is a paid mutator transaction binding the contract method 0x56de96db.
//
// Solidity: function setState(uint8 newState) returns()
func (_ExchangeImpl *ExchangeImplSession) SetState(newState uint8) (*types.Transaction, error) {
	return _ExchangeImpl.Contract.SetState(&_ExchangeImpl.TransactOpts, newState)
}

// SetState is a paid mutator transaction binding the contract method 0x56de96db.
//
// Solidity: function setState(uint8 newState) returns()
func (_ExchangeImpl *ExchangeImplTransactorSession) SetState(newState uint8) (*types.Transaction, error) {
	return _ExchangeImpl.Contract.SetState(&_ExchangeImpl.TransactOpts, newState)
}

// ExchangeImplBalanceLockedIterator is returned from FilterBalanceLocked and is used to iterate over the raw logs and unpacked data for BalanceLocked events raised by the ExchangeImpl contract.
type ExchangeImplBalanceLockedIterator struct {
	Event *ExchangeImplBalanceLocked // Event containing the contract specifics and raw log

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
func (it *ExchangeImplBalanceLockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeImplBalanceLocked)
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
		it.Event = new(ExchangeImplBalanceLocked)
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
func (it *ExchangeImplBalanceLockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeImplBalanceLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeImplBalanceLocked represents a BalanceLocked event raised by the ExchangeImpl contract.
type ExchangeImplBalanceLocked struct {
	Assetbox       common.Address
	BalanceAviable *big.Int
	BalanceLocked  *big.Int
	AssetboxType   *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBalanceLocked is a free log retrieval operation binding the contract event 0x89f85a4bd38f70943757e43dedd843409e565220cb52ba80fc297d1246b3b9bb.
//
// Solidity: event BalanceLocked(address assetbox, uint256 balanceAviable, uint256 balanceLocked, uint256 assetboxType)
func (_ExchangeImpl *ExchangeImplFilterer) FilterBalanceLocked(opts *bind.FilterOpts) (*ExchangeImplBalanceLockedIterator, error) {

	logs, sub, err := _ExchangeImpl.contract.FilterLogs(opts, "BalanceLocked")
	if err != nil {
		return nil, err
	}
	return &ExchangeImplBalanceLockedIterator{contract: _ExchangeImpl.contract, event: "BalanceLocked", logs: logs, sub: sub}, nil
}

// WatchBalanceLocked is a free log subscription operation binding the contract event 0x89f85a4bd38f70943757e43dedd843409e565220cb52ba80fc297d1246b3b9bb.
//
// Solidity: event BalanceLocked(address assetbox, uint256 balanceAviable, uint256 balanceLocked, uint256 assetboxType)
func (_ExchangeImpl *ExchangeImplFilterer) WatchBalanceLocked(opts *bind.WatchOpts, sink chan<- *ExchangeImplBalanceLocked) (event.Subscription, error) {

	logs, sub, err := _ExchangeImpl.contract.WatchLogs(opts, "BalanceLocked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeImplBalanceLocked)
				if err := _ExchangeImpl.contract.UnpackLog(event, "BalanceLocked", log); err != nil {
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
func (_ExchangeImpl *ExchangeImplFilterer) ParseBalanceLocked(log types.Log) (*ExchangeImplBalanceLocked, error) {
	event := new(ExchangeImplBalanceLocked)
	if err := _ExchangeImpl.contract.UnpackLog(event, "BalanceLocked", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ExchangeImplBalanceUnLockedIterator is returned from FilterBalanceUnLocked and is used to iterate over the raw logs and unpacked data for BalanceUnLocked events raised by the ExchangeImpl contract.
type ExchangeImplBalanceUnLockedIterator struct {
	Event *ExchangeImplBalanceUnLocked // Event containing the contract specifics and raw log

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
func (it *ExchangeImplBalanceUnLockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeImplBalanceUnLocked)
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
		it.Event = new(ExchangeImplBalanceUnLocked)
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
func (it *ExchangeImplBalanceUnLockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeImplBalanceUnLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeImplBalanceUnLocked represents a BalanceUnLocked event raised by the ExchangeImpl contract.
type ExchangeImplBalanceUnLocked struct {
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
func (_ExchangeImpl *ExchangeImplFilterer) FilterBalanceUnLocked(opts *bind.FilterOpts) (*ExchangeImplBalanceUnLockedIterator, error) {

	logs, sub, err := _ExchangeImpl.contract.FilterLogs(opts, "BalanceUnLocked")
	if err != nil {
		return nil, err
	}
	return &ExchangeImplBalanceUnLockedIterator{contract: _ExchangeImpl.contract, event: "BalanceUnLocked", logs: logs, sub: sub}, nil
}

// WatchBalanceUnLocked is a free log subscription operation binding the contract event 0x486bd0991e37e48fc69ea3be6a0b11058d29ef8a89bbc475b139a00a681b4c9b.
//
// Solidity: event BalanceUnLocked(address assetbox, uint256 balanceAviable, uint256 balanceLocked, uint256 assetboxType, uint256 amount)
func (_ExchangeImpl *ExchangeImplFilterer) WatchBalanceUnLocked(opts *bind.WatchOpts, sink chan<- *ExchangeImplBalanceUnLocked) (event.Subscription, error) {

	logs, sub, err := _ExchangeImpl.contract.WatchLogs(opts, "BalanceUnLocked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeImplBalanceUnLocked)
				if err := _ExchangeImpl.contract.UnpackLog(event, "BalanceUnLocked", log); err != nil {
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
func (_ExchangeImpl *ExchangeImplFilterer) ParseBalanceUnLocked(log types.Log) (*ExchangeImplBalanceUnLocked, error) {
	event := new(ExchangeImplBalanceUnLocked)
	if err := _ExchangeImpl.contract.UnpackLog(event, "BalanceUnLocked", log); err != nil {
		return nil, err
	}
	return event, nil
}
