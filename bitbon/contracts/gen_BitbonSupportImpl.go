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

// BitbonSupportImplABI is the input ABI used to generate the binding from.
const BitbonSupportImplABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"capitalizationAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"simcordAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"contractStorageAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ACCESS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ACCESS_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ADMINABLE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ADMIN_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ASSETBOX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ASSETBOX_INFO\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ASSETBOX_INFO_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ASSETBOX_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_BALANCE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_BITBON\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_BITBON_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_BITBON_SUPPORT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_BONBALANCE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_BONTRANSFER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_DEX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_DISTRIBUTION_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_EXCHANGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_EXCHANGE_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_FEE_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_GENERATOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MINING_AGENT_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MSBON\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MSBON_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_ADD_ADDRESS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_ADD_ADMIN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_ADD_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_ADD_ROLE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_DROP_PERMISSION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_EDIT_ADMIN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_EDIT_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_EDIT_ROLE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REMOVE_ADMIN_BY_ADDRESS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REMOVE_ADMIN_BY_INDEX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REMOVE_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REMOVE_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REMOVE_ROLE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REPLENISH_CAPITALIZATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REPLENISH_CERTIFICATE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REPLENISH_EMISSION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_SET_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_SET_PERMISSION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_OTC\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_OTC_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_REPLENISH_CAPITALIZATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_RESERVED_ALIAS_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ROLE_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_SAFE_TRANSFER_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_TRANSFER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_TRANSFER_FROM_CAPITALIZATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_ACCESS_RESTORATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_ADMIN_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_CONTRACTS_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_CONTRACT_DEPLOY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_EMISSION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_FEES_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_MULTISIG_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_ROLES_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_ACCESS_VERIFIER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_BITBON_ISSUE_VERIFIER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_COMMISSION_VERIFIER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_DEPLOY_ADMIN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_PERMISSION_ADMIN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"capitalizationAccount\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"contractStorage\",\"outputs\":[{\"internalType\":\"contractContractStorage\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"simcordAccount\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getThisContractIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getSimcordAccount\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCapitalizationAccount\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"assetbox\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"setBalanceAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"assetboxes\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"setBalancesAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"setMigrated\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// BitbonSupportImpl is an auto generated Go binding around an Ethereum contract.
type BitbonSupportImpl struct {
	BitbonSupportImplCaller     // Read-only binding to the contract
	BitbonSupportImplTransactor // Write-only binding to the contract
	BitbonSupportImplFilterer   // Log filterer for contract events
}

// BitbonSupportImplCaller is an auto generated read-only Go binding around an Ethereum contract.
type BitbonSupportImplCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BitbonSupportImplTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BitbonSupportImplTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BitbonSupportImplFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BitbonSupportImplFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BitbonSupportImplSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BitbonSupportImplSession struct {
	Contract     *BitbonSupportImpl // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// BitbonSupportImplCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BitbonSupportImplCallerSession struct {
	Contract *BitbonSupportImplCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// BitbonSupportImplTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BitbonSupportImplTransactorSession struct {
	Contract     *BitbonSupportImplTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// BitbonSupportImplRaw is an auto generated low-level Go binding around an Ethereum contract.
type BitbonSupportImplRaw struct {
	Contract *BitbonSupportImpl // Generic contract binding to access the raw methods on
}

// BitbonSupportImplCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BitbonSupportImplCallerRaw struct {
	Contract *BitbonSupportImplCaller // Generic read-only contract binding to access the raw methods on
}

// BitbonSupportImplTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BitbonSupportImplTransactorRaw struct {
	Contract *BitbonSupportImplTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBitbonSupportImpl creates a new instance of BitbonSupportImpl, bound to a specific deployed contract.
func NewBitbonSupportImpl(address common.Address, backend bind.ContractBackend) (*BitbonSupportImpl, error) {
	contract, err := bindBitbonSupportImpl(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BitbonSupportImpl{BitbonSupportImplCaller: BitbonSupportImplCaller{contract: contract}, BitbonSupportImplTransactor: BitbonSupportImplTransactor{contract: contract}, BitbonSupportImplFilterer: BitbonSupportImplFilterer{contract: contract}}, nil
}

// NewBitbonSupportImplCaller creates a new read-only instance of BitbonSupportImpl, bound to a specific deployed contract.
func NewBitbonSupportImplCaller(address common.Address, caller bind.ContractCaller) (*BitbonSupportImplCaller, error) {
	contract, err := bindBitbonSupportImpl(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BitbonSupportImplCaller{contract: contract}, nil
}

// NewBitbonSupportImplTransactor creates a new write-only instance of BitbonSupportImpl, bound to a specific deployed contract.
func NewBitbonSupportImplTransactor(address common.Address, transactor bind.ContractTransactor) (*BitbonSupportImplTransactor, error) {
	contract, err := bindBitbonSupportImpl(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BitbonSupportImplTransactor{contract: contract}, nil
}

// NewBitbonSupportImplFilterer creates a new log filterer instance of BitbonSupportImpl, bound to a specific deployed contract.
func NewBitbonSupportImplFilterer(address common.Address, filterer bind.ContractFilterer) (*BitbonSupportImplFilterer, error) {
	contract, err := bindBitbonSupportImpl(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BitbonSupportImplFilterer{contract: contract}, nil
}

// bindBitbonSupportImpl binds a generic wrapper to an already deployed contract.
func bindBitbonSupportImpl(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BitbonSupportImplABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BitbonSupportImpl *BitbonSupportImplRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BitbonSupportImpl.Contract.BitbonSupportImplCaller.contract.Call(opts, result, method, params...)
}

// Return Abi for BitbonSupportImplCaller
func (_BitbonSupportImpl *BitbonSupportImplCaller) ABI() abi.ABI {
	return _BitbonSupportImpl.contract.Abi
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BitbonSupportImpl *BitbonSupportImplRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BitbonSupportImpl.Contract.BitbonSupportImplTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BitbonSupportImpl *BitbonSupportImplRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BitbonSupportImpl.Contract.BitbonSupportImplTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BitbonSupportImpl *BitbonSupportImplCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BitbonSupportImpl.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BitbonSupportImpl *BitbonSupportImplTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BitbonSupportImpl.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BitbonSupportImpl *BitbonSupportImplTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BitbonSupportImpl.Contract.contract.Transact(opts, method, params...)
}

// CONTRACTACCESS is a free data retrieval call binding the contract method 0x0ac62b0c.
//
// Solidity: function CONTRACT_ACCESS() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCaller) CONTRACTACCESS(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonSupportImpl.contract.Call(opts, out, "CONTRACT_ACCESS")
	return *ret0, err
}

// CONTRACTACCESS is a free data retrieval call binding the contract method 0x0ac62b0c.
//
// Solidity: function CONTRACT_ACCESS() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplSession) CONTRACTACCESS() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.CONTRACTACCESS(&_BitbonSupportImpl.CallOpts)
}

// CONTRACTACCESS is a free data retrieval call binding the contract method 0x0ac62b0c.
//
// Solidity: function CONTRACT_ACCESS() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCallerSession) CONTRACTACCESS() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.CONTRACTACCESS(&_BitbonSupportImpl.CallOpts)
}

// CONTRACTACCESSSTORAGE is a free data retrieval call binding the contract method 0xceac9bd5.
//
// Solidity: function CONTRACT_ACCESS_STORAGE() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCaller) CONTRACTACCESSSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonSupportImpl.contract.Call(opts, out, "CONTRACT_ACCESS_STORAGE")
	return *ret0, err
}

// CONTRACTACCESSSTORAGE is a free data retrieval call binding the contract method 0xceac9bd5.
//
// Solidity: function CONTRACT_ACCESS_STORAGE() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplSession) CONTRACTACCESSSTORAGE() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.CONTRACTACCESSSTORAGE(&_BitbonSupportImpl.CallOpts)
}

// CONTRACTACCESSSTORAGE is a free data retrieval call binding the contract method 0xceac9bd5.
//
// Solidity: function CONTRACT_ACCESS_STORAGE() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCallerSession) CONTRACTACCESSSTORAGE() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.CONTRACTACCESSSTORAGE(&_BitbonSupportImpl.CallOpts)
}

// CONTRACTADMINABLE is a free data retrieval call binding the contract method 0xbbe3a78f.
//
// Solidity: function CONTRACT_ADMINABLE() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCaller) CONTRACTADMINABLE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonSupportImpl.contract.Call(opts, out, "CONTRACT_ADMINABLE")
	return *ret0, err
}

// CONTRACTADMINABLE is a free data retrieval call binding the contract method 0xbbe3a78f.
//
// Solidity: function CONTRACT_ADMINABLE() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplSession) CONTRACTADMINABLE() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.CONTRACTADMINABLE(&_BitbonSupportImpl.CallOpts)
}

// CONTRACTADMINABLE is a free data retrieval call binding the contract method 0xbbe3a78f.
//
// Solidity: function CONTRACT_ADMINABLE() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCallerSession) CONTRACTADMINABLE() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.CONTRACTADMINABLE(&_BitbonSupportImpl.CallOpts)
}

// CONTRACTADMINSTORAGE is a free data retrieval call binding the contract method 0xc9aebae9.
//
// Solidity: function CONTRACT_ADMIN_STORAGE() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCaller) CONTRACTADMINSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonSupportImpl.contract.Call(opts, out, "CONTRACT_ADMIN_STORAGE")
	return *ret0, err
}

// CONTRACTADMINSTORAGE is a free data retrieval call binding the contract method 0xc9aebae9.
//
// Solidity: function CONTRACT_ADMIN_STORAGE() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplSession) CONTRACTADMINSTORAGE() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.CONTRACTADMINSTORAGE(&_BitbonSupportImpl.CallOpts)
}

// CONTRACTADMINSTORAGE is a free data retrieval call binding the contract method 0xc9aebae9.
//
// Solidity: function CONTRACT_ADMIN_STORAGE() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCallerSession) CONTRACTADMINSTORAGE() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.CONTRACTADMINSTORAGE(&_BitbonSupportImpl.CallOpts)
}

// CONTRACTASSETBOX is a free data retrieval call binding the contract method 0x161aa17c.
//
// Solidity: function CONTRACT_ASSETBOX() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCaller) CONTRACTASSETBOX(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonSupportImpl.contract.Call(opts, out, "CONTRACT_ASSETBOX")
	return *ret0, err
}

// CONTRACTASSETBOX is a free data retrieval call binding the contract method 0x161aa17c.
//
// Solidity: function CONTRACT_ASSETBOX() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplSession) CONTRACTASSETBOX() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.CONTRACTASSETBOX(&_BitbonSupportImpl.CallOpts)
}

// CONTRACTASSETBOX is a free data retrieval call binding the contract method 0x161aa17c.
//
// Solidity: function CONTRACT_ASSETBOX() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCallerSession) CONTRACTASSETBOX() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.CONTRACTASSETBOX(&_BitbonSupportImpl.CallOpts)
}

// CONTRACTASSETBOXINFO is a free data retrieval call binding the contract method 0xb572fd3d.
//
// Solidity: function CONTRACT_ASSETBOX_INFO() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCaller) CONTRACTASSETBOXINFO(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonSupportImpl.contract.Call(opts, out, "CONTRACT_ASSETBOX_INFO")
	return *ret0, err
}

// CONTRACTASSETBOXINFO is a free data retrieval call binding the contract method 0xb572fd3d.
//
// Solidity: function CONTRACT_ASSETBOX_INFO() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplSession) CONTRACTASSETBOXINFO() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.CONTRACTASSETBOXINFO(&_BitbonSupportImpl.CallOpts)
}

// CONTRACTASSETBOXINFO is a free data retrieval call binding the contract method 0xb572fd3d.
//
// Solidity: function CONTRACT_ASSETBOX_INFO() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCallerSession) CONTRACTASSETBOXINFO() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.CONTRACTASSETBOXINFO(&_BitbonSupportImpl.CallOpts)
}

// CONTRACTASSETBOXINFOSTORAGE is a free data retrieval call binding the contract method 0xf32a115d.
//
// Solidity: function CONTRACT_ASSETBOX_INFO_STORAGE() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCaller) CONTRACTASSETBOXINFOSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonSupportImpl.contract.Call(opts, out, "CONTRACT_ASSETBOX_INFO_STORAGE")
	return *ret0, err
}

// CONTRACTASSETBOXINFOSTORAGE is a free data retrieval call binding the contract method 0xf32a115d.
//
// Solidity: function CONTRACT_ASSETBOX_INFO_STORAGE() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplSession) CONTRACTASSETBOXINFOSTORAGE() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.CONTRACTASSETBOXINFOSTORAGE(&_BitbonSupportImpl.CallOpts)
}

// CONTRACTASSETBOXINFOSTORAGE is a free data retrieval call binding the contract method 0xf32a115d.
//
// Solidity: function CONTRACT_ASSETBOX_INFO_STORAGE() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCallerSession) CONTRACTASSETBOXINFOSTORAGE() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.CONTRACTASSETBOXINFOSTORAGE(&_BitbonSupportImpl.CallOpts)
}

// CONTRACTASSETBOXSTORAGE is a free data retrieval call binding the contract method 0xfa5bdafc.
//
// Solidity: function CONTRACT_ASSETBOX_STORAGE() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCaller) CONTRACTASSETBOXSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonSupportImpl.contract.Call(opts, out, "CONTRACT_ASSETBOX_STORAGE")
	return *ret0, err
}

// CONTRACTASSETBOXSTORAGE is a free data retrieval call binding the contract method 0xfa5bdafc.
//
// Solidity: function CONTRACT_ASSETBOX_STORAGE() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplSession) CONTRACTASSETBOXSTORAGE() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.CONTRACTASSETBOXSTORAGE(&_BitbonSupportImpl.CallOpts)
}

// CONTRACTASSETBOXSTORAGE is a free data retrieval call binding the contract method 0xfa5bdafc.
//
// Solidity: function CONTRACT_ASSETBOX_STORAGE() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCallerSession) CONTRACTASSETBOXSTORAGE() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.CONTRACTASSETBOXSTORAGE(&_BitbonSupportImpl.CallOpts)
}

// CONTRACTBALANCE is a free data retrieval call binding the contract method 0xd59045f5.
//
// Solidity: function CONTRACT_BALANCE() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCaller) CONTRACTBALANCE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonSupportImpl.contract.Call(opts, out, "CONTRACT_BALANCE")
	return *ret0, err
}

// CONTRACTBALANCE is a free data retrieval call binding the contract method 0xd59045f5.
//
// Solidity: function CONTRACT_BALANCE() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplSession) CONTRACTBALANCE() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.CONTRACTBALANCE(&_BitbonSupportImpl.CallOpts)
}

// CONTRACTBALANCE is a free data retrieval call binding the contract method 0xd59045f5.
//
// Solidity: function CONTRACT_BALANCE() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCallerSession) CONTRACTBALANCE() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.CONTRACTBALANCE(&_BitbonSupportImpl.CallOpts)
}

// CONTRACTBITBON is a free data retrieval call binding the contract method 0xb6b433ab.
//
// Solidity: function CONTRACT_BITBON() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCaller) CONTRACTBITBON(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonSupportImpl.contract.Call(opts, out, "CONTRACT_BITBON")
	return *ret0, err
}

// CONTRACTBITBON is a free data retrieval call binding the contract method 0xb6b433ab.
//
// Solidity: function CONTRACT_BITBON() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplSession) CONTRACTBITBON() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.CONTRACTBITBON(&_BitbonSupportImpl.CallOpts)
}

// CONTRACTBITBON is a free data retrieval call binding the contract method 0xb6b433ab.
//
// Solidity: function CONTRACT_BITBON() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCallerSession) CONTRACTBITBON() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.CONTRACTBITBON(&_BitbonSupportImpl.CallOpts)
}

// CONTRACTBITBONSTORAGE is a free data retrieval call binding the contract method 0xe7f6c30e.
//
// Solidity: function CONTRACT_BITBON_STORAGE() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCaller) CONTRACTBITBONSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonSupportImpl.contract.Call(opts, out, "CONTRACT_BITBON_STORAGE")
	return *ret0, err
}

// CONTRACTBITBONSTORAGE is a free data retrieval call binding the contract method 0xe7f6c30e.
//
// Solidity: function CONTRACT_BITBON_STORAGE() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplSession) CONTRACTBITBONSTORAGE() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.CONTRACTBITBONSTORAGE(&_BitbonSupportImpl.CallOpts)
}

// CONTRACTBITBONSTORAGE is a free data retrieval call binding the contract method 0xe7f6c30e.
//
// Solidity: function CONTRACT_BITBON_STORAGE() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCallerSession) CONTRACTBITBONSTORAGE() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.CONTRACTBITBONSTORAGE(&_BitbonSupportImpl.CallOpts)
}

// CONTRACTBITBONSUPPORT is a free data retrieval call binding the contract method 0x20139c74.
//
// Solidity: function CONTRACT_BITBON_SUPPORT() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCaller) CONTRACTBITBONSUPPORT(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonSupportImpl.contract.Call(opts, out, "CONTRACT_BITBON_SUPPORT")
	return *ret0, err
}

// CONTRACTBITBONSUPPORT is a free data retrieval call binding the contract method 0x20139c74.
//
// Solidity: function CONTRACT_BITBON_SUPPORT() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplSession) CONTRACTBITBONSUPPORT() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.CONTRACTBITBONSUPPORT(&_BitbonSupportImpl.CallOpts)
}

// CONTRACTBITBONSUPPORT is a free data retrieval call binding the contract method 0x20139c74.
//
// Solidity: function CONTRACT_BITBON_SUPPORT() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCallerSession) CONTRACTBITBONSUPPORT() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.CONTRACTBITBONSUPPORT(&_BitbonSupportImpl.CallOpts)
}

// CONTRACTBONBALANCE is a free data retrieval call binding the contract method 0x40e02b7c.
//
// Solidity: function CONTRACT_BONBALANCE() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCaller) CONTRACTBONBALANCE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonSupportImpl.contract.Call(opts, out, "CONTRACT_BONBALANCE")
	return *ret0, err
}

// CONTRACTBONBALANCE is a free data retrieval call binding the contract method 0x40e02b7c.
//
// Solidity: function CONTRACT_BONBALANCE() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplSession) CONTRACTBONBALANCE() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.CONTRACTBONBALANCE(&_BitbonSupportImpl.CallOpts)
}

// CONTRACTBONBALANCE is a free data retrieval call binding the contract method 0x40e02b7c.
//
// Solidity: function CONTRACT_BONBALANCE() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCallerSession) CONTRACTBONBALANCE() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.CONTRACTBONBALANCE(&_BitbonSupportImpl.CallOpts)
}

// CONTRACTBONTRANSFER is a free data retrieval call binding the contract method 0xc18f4a8a.
//
// Solidity: function CONTRACT_BONTRANSFER() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCaller) CONTRACTBONTRANSFER(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonSupportImpl.contract.Call(opts, out, "CONTRACT_BONTRANSFER")
	return *ret0, err
}

// CONTRACTBONTRANSFER is a free data retrieval call binding the contract method 0xc18f4a8a.
//
// Solidity: function CONTRACT_BONTRANSFER() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplSession) CONTRACTBONTRANSFER() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.CONTRACTBONTRANSFER(&_BitbonSupportImpl.CallOpts)
}

// CONTRACTBONTRANSFER is a free data retrieval call binding the contract method 0xc18f4a8a.
//
// Solidity: function CONTRACT_BONTRANSFER() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCallerSession) CONTRACTBONTRANSFER() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.CONTRACTBONTRANSFER(&_BitbonSupportImpl.CallOpts)
}

// CONTRACTDEX is a free data retrieval call binding the contract method 0x52c1f844.
//
// Solidity: function CONTRACT_DEX() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCaller) CONTRACTDEX(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonSupportImpl.contract.Call(opts, out, "CONTRACT_DEX")
	return *ret0, err
}

// CONTRACTDEX is a free data retrieval call binding the contract method 0x52c1f844.
//
// Solidity: function CONTRACT_DEX() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplSession) CONTRACTDEX() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.CONTRACTDEX(&_BitbonSupportImpl.CallOpts)
}

// CONTRACTDEX is a free data retrieval call binding the contract method 0x52c1f844.
//
// Solidity: function CONTRACT_DEX() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCallerSession) CONTRACTDEX() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.CONTRACTDEX(&_BitbonSupportImpl.CallOpts)
}

// CONTRACTDISTRIBUTIONSTORAGE is a free data retrieval call binding the contract method 0x08097327.
//
// Solidity: function CONTRACT_DISTRIBUTION_STORAGE() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCaller) CONTRACTDISTRIBUTIONSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonSupportImpl.contract.Call(opts, out, "CONTRACT_DISTRIBUTION_STORAGE")
	return *ret0, err
}

// CONTRACTDISTRIBUTIONSTORAGE is a free data retrieval call binding the contract method 0x08097327.
//
// Solidity: function CONTRACT_DISTRIBUTION_STORAGE() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplSession) CONTRACTDISTRIBUTIONSTORAGE() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.CONTRACTDISTRIBUTIONSTORAGE(&_BitbonSupportImpl.CallOpts)
}

// CONTRACTDISTRIBUTIONSTORAGE is a free data retrieval call binding the contract method 0x08097327.
//
// Solidity: function CONTRACT_DISTRIBUTION_STORAGE() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCallerSession) CONTRACTDISTRIBUTIONSTORAGE() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.CONTRACTDISTRIBUTIONSTORAGE(&_BitbonSupportImpl.CallOpts)
}

// CONTRACTEXCHANGE is a free data retrieval call binding the contract method 0xa0e3a10e.
//
// Solidity: function CONTRACT_EXCHANGE() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCaller) CONTRACTEXCHANGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonSupportImpl.contract.Call(opts, out, "CONTRACT_EXCHANGE")
	return *ret0, err
}

// CONTRACTEXCHANGE is a free data retrieval call binding the contract method 0xa0e3a10e.
//
// Solidity: function CONTRACT_EXCHANGE() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplSession) CONTRACTEXCHANGE() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.CONTRACTEXCHANGE(&_BitbonSupportImpl.CallOpts)
}

// CONTRACTEXCHANGE is a free data retrieval call binding the contract method 0xa0e3a10e.
//
// Solidity: function CONTRACT_EXCHANGE() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCallerSession) CONTRACTEXCHANGE() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.CONTRACTEXCHANGE(&_BitbonSupportImpl.CallOpts)
}

// CONTRACTEXCHANGESTORAGE is a free data retrieval call binding the contract method 0x001570f2.
//
// Solidity: function CONTRACT_EXCHANGE_STORAGE() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCaller) CONTRACTEXCHANGESTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonSupportImpl.contract.Call(opts, out, "CONTRACT_EXCHANGE_STORAGE")
	return *ret0, err
}

// CONTRACTEXCHANGESTORAGE is a free data retrieval call binding the contract method 0x001570f2.
//
// Solidity: function CONTRACT_EXCHANGE_STORAGE() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplSession) CONTRACTEXCHANGESTORAGE() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.CONTRACTEXCHANGESTORAGE(&_BitbonSupportImpl.CallOpts)
}

// CONTRACTEXCHANGESTORAGE is a free data retrieval call binding the contract method 0x001570f2.
//
// Solidity: function CONTRACT_EXCHANGE_STORAGE() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCallerSession) CONTRACTEXCHANGESTORAGE() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.CONTRACTEXCHANGESTORAGE(&_BitbonSupportImpl.CallOpts)
}

// CONTRACTFEESTORAGE is a free data retrieval call binding the contract method 0xcd5e65a4.
//
// Solidity: function CONTRACT_FEE_STORAGE() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCaller) CONTRACTFEESTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonSupportImpl.contract.Call(opts, out, "CONTRACT_FEE_STORAGE")
	return *ret0, err
}

// CONTRACTFEESTORAGE is a free data retrieval call binding the contract method 0xcd5e65a4.
//
// Solidity: function CONTRACT_FEE_STORAGE() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplSession) CONTRACTFEESTORAGE() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.CONTRACTFEESTORAGE(&_BitbonSupportImpl.CallOpts)
}

// CONTRACTFEESTORAGE is a free data retrieval call binding the contract method 0xcd5e65a4.
//
// Solidity: function CONTRACT_FEE_STORAGE() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCallerSession) CONTRACTFEESTORAGE() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.CONTRACTFEESTORAGE(&_BitbonSupportImpl.CallOpts)
}

// CONTRACTGENERATOR is a free data retrieval call binding the contract method 0xccea2b46.
//
// Solidity: function CONTRACT_GENERATOR() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCaller) CONTRACTGENERATOR(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonSupportImpl.contract.Call(opts, out, "CONTRACT_GENERATOR")
	return *ret0, err
}

// CONTRACTGENERATOR is a free data retrieval call binding the contract method 0xccea2b46.
//
// Solidity: function CONTRACT_GENERATOR() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplSession) CONTRACTGENERATOR() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.CONTRACTGENERATOR(&_BitbonSupportImpl.CallOpts)
}

// CONTRACTGENERATOR is a free data retrieval call binding the contract method 0xccea2b46.
//
// Solidity: function CONTRACT_GENERATOR() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCallerSession) CONTRACTGENERATOR() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.CONTRACTGENERATOR(&_BitbonSupportImpl.CallOpts)
}

// CONTRACTMININGAGENTSTORAGE is a free data retrieval call binding the contract method 0x83cc588d.
//
// Solidity: function CONTRACT_MINING_AGENT_STORAGE() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCaller) CONTRACTMININGAGENTSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonSupportImpl.contract.Call(opts, out, "CONTRACT_MINING_AGENT_STORAGE")
	return *ret0, err
}

// CONTRACTMININGAGENTSTORAGE is a free data retrieval call binding the contract method 0x83cc588d.
//
// Solidity: function CONTRACT_MINING_AGENT_STORAGE() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplSession) CONTRACTMININGAGENTSTORAGE() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.CONTRACTMININGAGENTSTORAGE(&_BitbonSupportImpl.CallOpts)
}

// CONTRACTMININGAGENTSTORAGE is a free data retrieval call binding the contract method 0x83cc588d.
//
// Solidity: function CONTRACT_MINING_AGENT_STORAGE() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCallerSession) CONTRACTMININGAGENTSTORAGE() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.CONTRACTMININGAGENTSTORAGE(&_BitbonSupportImpl.CallOpts)
}

// CONTRACTMSBON is a free data retrieval call binding the contract method 0x28b3a976.
//
// Solidity: function CONTRACT_MSBON() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCaller) CONTRACTMSBON(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonSupportImpl.contract.Call(opts, out, "CONTRACT_MSBON")
	return *ret0, err
}

// CONTRACTMSBON is a free data retrieval call binding the contract method 0x28b3a976.
//
// Solidity: function CONTRACT_MSBON() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplSession) CONTRACTMSBON() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.CONTRACTMSBON(&_BitbonSupportImpl.CallOpts)
}

// CONTRACTMSBON is a free data retrieval call binding the contract method 0x28b3a976.
//
// Solidity: function CONTRACT_MSBON() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCallerSession) CONTRACTMSBON() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.CONTRACTMSBON(&_BitbonSupportImpl.CallOpts)
}

// CONTRACTMSBONSTORAGE is a free data retrieval call binding the contract method 0xe04e6061.
//
// Solidity: function CONTRACT_MSBON_STORAGE() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCaller) CONTRACTMSBONSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonSupportImpl.contract.Call(opts, out, "CONTRACT_MSBON_STORAGE")
	return *ret0, err
}

// CONTRACTMSBONSTORAGE is a free data retrieval call binding the contract method 0xe04e6061.
//
// Solidity: function CONTRACT_MSBON_STORAGE() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplSession) CONTRACTMSBONSTORAGE() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.CONTRACTMSBONSTORAGE(&_BitbonSupportImpl.CallOpts)
}

// CONTRACTMSBONSTORAGE is a free data retrieval call binding the contract method 0xe04e6061.
//
// Solidity: function CONTRACT_MSBON_STORAGE() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCallerSession) CONTRACTMSBONSTORAGE() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.CONTRACTMSBONSTORAGE(&_BitbonSupportImpl.CallOpts)
}

// CONTRACTMULTISIGADDADDRESS is a free data retrieval call binding the contract method 0xf4f65a5a.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ADDRESS() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCaller) CONTRACTMULTISIGADDADDRESS(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonSupportImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_ADD_ADDRESS")
	return *ret0, err
}

// CONTRACTMULTISIGADDADDRESS is a free data retrieval call binding the contract method 0xf4f65a5a.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ADDRESS() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplSession) CONTRACTMULTISIGADDADDRESS() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.CONTRACTMULTISIGADDADDRESS(&_BitbonSupportImpl.CallOpts)
}

// CONTRACTMULTISIGADDADDRESS is a free data retrieval call binding the contract method 0xf4f65a5a.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ADDRESS() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCallerSession) CONTRACTMULTISIGADDADDRESS() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.CONTRACTMULTISIGADDADDRESS(&_BitbonSupportImpl.CallOpts)
}

// CONTRACTMULTISIGADDADMIN is a free data retrieval call binding the contract method 0x83888c68.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ADMIN() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCaller) CONTRACTMULTISIGADDADMIN(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonSupportImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_ADD_ADMIN")
	return *ret0, err
}

// CONTRACTMULTISIGADDADMIN is a free data retrieval call binding the contract method 0x83888c68.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ADMIN() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplSession) CONTRACTMULTISIGADDADMIN() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.CONTRACTMULTISIGADDADMIN(&_BitbonSupportImpl.CallOpts)
}

// CONTRACTMULTISIGADDADMIN is a free data retrieval call binding the contract method 0x83888c68.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ADMIN() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCallerSession) CONTRACTMULTISIGADDADMIN() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.CONTRACTMULTISIGADDADMIN(&_BitbonSupportImpl.CallOpts)
}

// CONTRACTMULTISIGADDFEE is a free data retrieval call binding the contract method 0xc43e8cab.
//
// Solidity: function CONTRACT_MULTISIG_ADD_FEE() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCaller) CONTRACTMULTISIGADDFEE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonSupportImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_ADD_FEE")
	return *ret0, err
}

// CONTRACTMULTISIGADDFEE is a free data retrieval call binding the contract method 0xc43e8cab.
//
// Solidity: function CONTRACT_MULTISIG_ADD_FEE() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplSession) CONTRACTMULTISIGADDFEE() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.CONTRACTMULTISIGADDFEE(&_BitbonSupportImpl.CallOpts)
}

// CONTRACTMULTISIGADDFEE is a free data retrieval call binding the contract method 0xc43e8cab.
//
// Solidity: function CONTRACT_MULTISIG_ADD_FEE() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCallerSession) CONTRACTMULTISIGADDFEE() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.CONTRACTMULTISIGADDFEE(&_BitbonSupportImpl.CallOpts)
}

// CONTRACTMULTISIGADDROLE is a free data retrieval call binding the contract method 0x20a0914b.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ROLE() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCaller) CONTRACTMULTISIGADDROLE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonSupportImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_ADD_ROLE")
	return *ret0, err
}

// CONTRACTMULTISIGADDROLE is a free data retrieval call binding the contract method 0x20a0914b.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ROLE() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplSession) CONTRACTMULTISIGADDROLE() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.CONTRACTMULTISIGADDROLE(&_BitbonSupportImpl.CallOpts)
}

// CONTRACTMULTISIGADDROLE is a free data retrieval call binding the contract method 0x20a0914b.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ROLE() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCallerSession) CONTRACTMULTISIGADDROLE() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.CONTRACTMULTISIGADDROLE(&_BitbonSupportImpl.CallOpts)
}

// CONTRACTMULTISIGDROPPERMISSION is a free data retrieval call binding the contract method 0x3f3fbc39.
//
// Solidity: function CONTRACT_MULTISIG_DROP_PERMISSION() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCaller) CONTRACTMULTISIGDROPPERMISSION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonSupportImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_DROP_PERMISSION")
	return *ret0, err
}

// CONTRACTMULTISIGDROPPERMISSION is a free data retrieval call binding the contract method 0x3f3fbc39.
//
// Solidity: function CONTRACT_MULTISIG_DROP_PERMISSION() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplSession) CONTRACTMULTISIGDROPPERMISSION() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.CONTRACTMULTISIGDROPPERMISSION(&_BitbonSupportImpl.CallOpts)
}

// CONTRACTMULTISIGDROPPERMISSION is a free data retrieval call binding the contract method 0x3f3fbc39.
//
// Solidity: function CONTRACT_MULTISIG_DROP_PERMISSION() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCallerSession) CONTRACTMULTISIGDROPPERMISSION() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.CONTRACTMULTISIGDROPPERMISSION(&_BitbonSupportImpl.CallOpts)
}

// CONTRACTMULTISIGEDITADMIN is a free data retrieval call binding the contract method 0x91b0f051.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_ADMIN() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCaller) CONTRACTMULTISIGEDITADMIN(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonSupportImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_EDIT_ADMIN")
	return *ret0, err
}

// CONTRACTMULTISIGEDITADMIN is a free data retrieval call binding the contract method 0x91b0f051.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_ADMIN() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplSession) CONTRACTMULTISIGEDITADMIN() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.CONTRACTMULTISIGEDITADMIN(&_BitbonSupportImpl.CallOpts)
}

// CONTRACTMULTISIGEDITADMIN is a free data retrieval call binding the contract method 0x91b0f051.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_ADMIN() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCallerSession) CONTRACTMULTISIGEDITADMIN() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.CONTRACTMULTISIGEDITADMIN(&_BitbonSupportImpl.CallOpts)
}

// CONTRACTMULTISIGEDITFEE is a free data retrieval call binding the contract method 0x6b2abe17.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_FEE() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCaller) CONTRACTMULTISIGEDITFEE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonSupportImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_EDIT_FEE")
	return *ret0, err
}

// CONTRACTMULTISIGEDITFEE is a free data retrieval call binding the contract method 0x6b2abe17.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_FEE() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplSession) CONTRACTMULTISIGEDITFEE() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.CONTRACTMULTISIGEDITFEE(&_BitbonSupportImpl.CallOpts)
}

// CONTRACTMULTISIGEDITFEE is a free data retrieval call binding the contract method 0x6b2abe17.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_FEE() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCallerSession) CONTRACTMULTISIGEDITFEE() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.CONTRACTMULTISIGEDITFEE(&_BitbonSupportImpl.CallOpts)
}

// CONTRACTMULTISIGEDITROLE is a free data retrieval call binding the contract method 0xc5f66412.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_ROLE() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCaller) CONTRACTMULTISIGEDITROLE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonSupportImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_EDIT_ROLE")
	return *ret0, err
}

// CONTRACTMULTISIGEDITROLE is a free data retrieval call binding the contract method 0xc5f66412.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_ROLE() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplSession) CONTRACTMULTISIGEDITROLE() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.CONTRACTMULTISIGEDITROLE(&_BitbonSupportImpl.CallOpts)
}

// CONTRACTMULTISIGEDITROLE is a free data retrieval call binding the contract method 0xc5f66412.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_ROLE() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCallerSession) CONTRACTMULTISIGEDITROLE() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.CONTRACTMULTISIGEDITROLE(&_BitbonSupportImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEADMINBYADDRESS is a free data retrieval call binding the contract method 0xc816b091.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_BY_ADDRESS() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCaller) CONTRACTMULTISIGREMOVEADMINBYADDRESS(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonSupportImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REMOVE_ADMIN_BY_ADDRESS")
	return *ret0, err
}

// CONTRACTMULTISIGREMOVEADMINBYADDRESS is a free data retrieval call binding the contract method 0xc816b091.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_BY_ADDRESS() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplSession) CONTRACTMULTISIGREMOVEADMINBYADDRESS() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.CONTRACTMULTISIGREMOVEADMINBYADDRESS(&_BitbonSupportImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEADMINBYADDRESS is a free data retrieval call binding the contract method 0xc816b091.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_BY_ADDRESS() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCallerSession) CONTRACTMULTISIGREMOVEADMINBYADDRESS() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.CONTRACTMULTISIGREMOVEADMINBYADDRESS(&_BitbonSupportImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEADMINBYINDEX is a free data retrieval call binding the contract method 0x751fbea2.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_BY_INDEX() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCaller) CONTRACTMULTISIGREMOVEADMINBYINDEX(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonSupportImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REMOVE_ADMIN_BY_INDEX")
	return *ret0, err
}

// CONTRACTMULTISIGREMOVEADMINBYINDEX is a free data retrieval call binding the contract method 0x751fbea2.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_BY_INDEX() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplSession) CONTRACTMULTISIGREMOVEADMINBYINDEX() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.CONTRACTMULTISIGREMOVEADMINBYINDEX(&_BitbonSupportImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEADMINBYINDEX is a free data retrieval call binding the contract method 0x751fbea2.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_BY_INDEX() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCallerSession) CONTRACTMULTISIGREMOVEADMINBYINDEX() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.CONTRACTMULTISIGREMOVEADMINBYINDEX(&_BitbonSupportImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEADMINROLE is a free data retrieval call binding the contract method 0x53c00200.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_ROLE() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCaller) CONTRACTMULTISIGREMOVEADMINROLE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonSupportImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REMOVE_ADMIN_ROLE")
	return *ret0, err
}

// CONTRACTMULTISIGREMOVEADMINROLE is a free data retrieval call binding the contract method 0x53c00200.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_ROLE() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplSession) CONTRACTMULTISIGREMOVEADMINROLE() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.CONTRACTMULTISIGREMOVEADMINROLE(&_BitbonSupportImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEADMINROLE is a free data retrieval call binding the contract method 0x53c00200.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_ROLE() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCallerSession) CONTRACTMULTISIGREMOVEADMINROLE() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.CONTRACTMULTISIGREMOVEADMINROLE(&_BitbonSupportImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEFEE is a free data retrieval call binding the contract method 0xdb614d25.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_FEE() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCaller) CONTRACTMULTISIGREMOVEFEE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonSupportImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REMOVE_FEE")
	return *ret0, err
}

// CONTRACTMULTISIGREMOVEFEE is a free data retrieval call binding the contract method 0xdb614d25.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_FEE() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplSession) CONTRACTMULTISIGREMOVEFEE() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.CONTRACTMULTISIGREMOVEFEE(&_BitbonSupportImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEFEE is a free data retrieval call binding the contract method 0xdb614d25.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_FEE() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCallerSession) CONTRACTMULTISIGREMOVEFEE() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.CONTRACTMULTISIGREMOVEFEE(&_BitbonSupportImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEROLE is a free data retrieval call binding the contract method 0x23648496.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ROLE() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCaller) CONTRACTMULTISIGREMOVEROLE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonSupportImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REMOVE_ROLE")
	return *ret0, err
}

// CONTRACTMULTISIGREMOVEROLE is a free data retrieval call binding the contract method 0x23648496.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ROLE() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplSession) CONTRACTMULTISIGREMOVEROLE() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.CONTRACTMULTISIGREMOVEROLE(&_BitbonSupportImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEROLE is a free data retrieval call binding the contract method 0x23648496.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ROLE() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCallerSession) CONTRACTMULTISIGREMOVEROLE() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.CONTRACTMULTISIGREMOVEROLE(&_BitbonSupportImpl.CallOpts)
}

// CONTRACTMULTISIGREPLENISHCAPITALIZATION is a free data retrieval call binding the contract method 0x62d5d523.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_CAPITALIZATION() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCaller) CONTRACTMULTISIGREPLENISHCAPITALIZATION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonSupportImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REPLENISH_CAPITALIZATION")
	return *ret0, err
}

// CONTRACTMULTISIGREPLENISHCAPITALIZATION is a free data retrieval call binding the contract method 0x62d5d523.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_CAPITALIZATION() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplSession) CONTRACTMULTISIGREPLENISHCAPITALIZATION() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.CONTRACTMULTISIGREPLENISHCAPITALIZATION(&_BitbonSupportImpl.CallOpts)
}

// CONTRACTMULTISIGREPLENISHCAPITALIZATION is a free data retrieval call binding the contract method 0x62d5d523.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_CAPITALIZATION() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCallerSession) CONTRACTMULTISIGREPLENISHCAPITALIZATION() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.CONTRACTMULTISIGREPLENISHCAPITALIZATION(&_BitbonSupportImpl.CallOpts)
}

// CONTRACTMULTISIGREPLENISHCERTIFICATE is a free data retrieval call binding the contract method 0xfd19b7e1.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_CERTIFICATE() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCaller) CONTRACTMULTISIGREPLENISHCERTIFICATE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonSupportImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REPLENISH_CERTIFICATE")
	return *ret0, err
}

// CONTRACTMULTISIGREPLENISHCERTIFICATE is a free data retrieval call binding the contract method 0xfd19b7e1.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_CERTIFICATE() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplSession) CONTRACTMULTISIGREPLENISHCERTIFICATE() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.CONTRACTMULTISIGREPLENISHCERTIFICATE(&_BitbonSupportImpl.CallOpts)
}

// CONTRACTMULTISIGREPLENISHCERTIFICATE is a free data retrieval call binding the contract method 0xfd19b7e1.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_CERTIFICATE() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCallerSession) CONTRACTMULTISIGREPLENISHCERTIFICATE() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.CONTRACTMULTISIGREPLENISHCERTIFICATE(&_BitbonSupportImpl.CallOpts)
}

// CONTRACTMULTISIGREPLENISHEMISSION is a free data retrieval call binding the contract method 0x0cb59fde.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_EMISSION() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCaller) CONTRACTMULTISIGREPLENISHEMISSION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonSupportImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REPLENISH_EMISSION")
	return *ret0, err
}

// CONTRACTMULTISIGREPLENISHEMISSION is a free data retrieval call binding the contract method 0x0cb59fde.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_EMISSION() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplSession) CONTRACTMULTISIGREPLENISHEMISSION() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.CONTRACTMULTISIGREPLENISHEMISSION(&_BitbonSupportImpl.CallOpts)
}

// CONTRACTMULTISIGREPLENISHEMISSION is a free data retrieval call binding the contract method 0x0cb59fde.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_EMISSION() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCallerSession) CONTRACTMULTISIGREPLENISHEMISSION() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.CONTRACTMULTISIGREPLENISHEMISSION(&_BitbonSupportImpl.CallOpts)
}

// CONTRACTMULTISIGSETADMINROLE is a free data retrieval call binding the contract method 0x8f209ab7.
//
// Solidity: function CONTRACT_MULTISIG_SET_ADMIN_ROLE() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCaller) CONTRACTMULTISIGSETADMINROLE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonSupportImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_SET_ADMIN_ROLE")
	return *ret0, err
}

// CONTRACTMULTISIGSETADMINROLE is a free data retrieval call binding the contract method 0x8f209ab7.
//
// Solidity: function CONTRACT_MULTISIG_SET_ADMIN_ROLE() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplSession) CONTRACTMULTISIGSETADMINROLE() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.CONTRACTMULTISIGSETADMINROLE(&_BitbonSupportImpl.CallOpts)
}

// CONTRACTMULTISIGSETADMINROLE is a free data retrieval call binding the contract method 0x8f209ab7.
//
// Solidity: function CONTRACT_MULTISIG_SET_ADMIN_ROLE() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCallerSession) CONTRACTMULTISIGSETADMINROLE() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.CONTRACTMULTISIGSETADMINROLE(&_BitbonSupportImpl.CallOpts)
}

// CONTRACTMULTISIGSETPERMISSION is a free data retrieval call binding the contract method 0xb0e91914.
//
// Solidity: function CONTRACT_MULTISIG_SET_PERMISSION() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCaller) CONTRACTMULTISIGSETPERMISSION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonSupportImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_SET_PERMISSION")
	return *ret0, err
}

// CONTRACTMULTISIGSETPERMISSION is a free data retrieval call binding the contract method 0xb0e91914.
//
// Solidity: function CONTRACT_MULTISIG_SET_PERMISSION() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplSession) CONTRACTMULTISIGSETPERMISSION() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.CONTRACTMULTISIGSETPERMISSION(&_BitbonSupportImpl.CallOpts)
}

// CONTRACTMULTISIGSETPERMISSION is a free data retrieval call binding the contract method 0xb0e91914.
//
// Solidity: function CONTRACT_MULTISIG_SET_PERMISSION() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCallerSession) CONTRACTMULTISIGSETPERMISSION() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.CONTRACTMULTISIGSETPERMISSION(&_BitbonSupportImpl.CallOpts)
}

// CONTRACTMULTISIGSTORAGE is a free data retrieval call binding the contract method 0xe7e1a303.
//
// Solidity: function CONTRACT_MULTISIG_STORAGE() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCaller) CONTRACTMULTISIGSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonSupportImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_STORAGE")
	return *ret0, err
}

// CONTRACTMULTISIGSTORAGE is a free data retrieval call binding the contract method 0xe7e1a303.
//
// Solidity: function CONTRACT_MULTISIG_STORAGE() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplSession) CONTRACTMULTISIGSTORAGE() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.CONTRACTMULTISIGSTORAGE(&_BitbonSupportImpl.CallOpts)
}

// CONTRACTMULTISIGSTORAGE is a free data retrieval call binding the contract method 0xe7e1a303.
//
// Solidity: function CONTRACT_MULTISIG_STORAGE() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCallerSession) CONTRACTMULTISIGSTORAGE() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.CONTRACTMULTISIGSTORAGE(&_BitbonSupportImpl.CallOpts)
}

// CONTRACTOTC is a free data retrieval call binding the contract method 0x5eceac70.
//
// Solidity: function CONTRACT_OTC() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCaller) CONTRACTOTC(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonSupportImpl.contract.Call(opts, out, "CONTRACT_OTC")
	return *ret0, err
}

// CONTRACTOTC is a free data retrieval call binding the contract method 0x5eceac70.
//
// Solidity: function CONTRACT_OTC() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplSession) CONTRACTOTC() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.CONTRACTOTC(&_BitbonSupportImpl.CallOpts)
}

// CONTRACTOTC is a free data retrieval call binding the contract method 0x5eceac70.
//
// Solidity: function CONTRACT_OTC() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCallerSession) CONTRACTOTC() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.CONTRACTOTC(&_BitbonSupportImpl.CallOpts)
}

// CONTRACTOTCSTORAGE is a free data retrieval call binding the contract method 0x825b476f.
//
// Solidity: function CONTRACT_OTC_STORAGE() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCaller) CONTRACTOTCSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonSupportImpl.contract.Call(opts, out, "CONTRACT_OTC_STORAGE")
	return *ret0, err
}

// CONTRACTOTCSTORAGE is a free data retrieval call binding the contract method 0x825b476f.
//
// Solidity: function CONTRACT_OTC_STORAGE() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplSession) CONTRACTOTCSTORAGE() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.CONTRACTOTCSTORAGE(&_BitbonSupportImpl.CallOpts)
}

// CONTRACTOTCSTORAGE is a free data retrieval call binding the contract method 0x825b476f.
//
// Solidity: function CONTRACT_OTC_STORAGE() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCallerSession) CONTRACTOTCSTORAGE() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.CONTRACTOTCSTORAGE(&_BitbonSupportImpl.CallOpts)
}

// CONTRACTREPLENISHCAPITALIZATION is a free data retrieval call binding the contract method 0xff5e09d8.
//
// Solidity: function CONTRACT_REPLENISH_CAPITALIZATION() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCaller) CONTRACTREPLENISHCAPITALIZATION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonSupportImpl.contract.Call(opts, out, "CONTRACT_REPLENISH_CAPITALIZATION")
	return *ret0, err
}

// CONTRACTREPLENISHCAPITALIZATION is a free data retrieval call binding the contract method 0xff5e09d8.
//
// Solidity: function CONTRACT_REPLENISH_CAPITALIZATION() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplSession) CONTRACTREPLENISHCAPITALIZATION() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.CONTRACTREPLENISHCAPITALIZATION(&_BitbonSupportImpl.CallOpts)
}

// CONTRACTREPLENISHCAPITALIZATION is a free data retrieval call binding the contract method 0xff5e09d8.
//
// Solidity: function CONTRACT_REPLENISH_CAPITALIZATION() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCallerSession) CONTRACTREPLENISHCAPITALIZATION() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.CONTRACTREPLENISHCAPITALIZATION(&_BitbonSupportImpl.CallOpts)
}

// CONTRACTRESERVEDALIASSTORAGE is a free data retrieval call binding the contract method 0x70a94aa7.
//
// Solidity: function CONTRACT_RESERVED_ALIAS_STORAGE() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCaller) CONTRACTRESERVEDALIASSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonSupportImpl.contract.Call(opts, out, "CONTRACT_RESERVED_ALIAS_STORAGE")
	return *ret0, err
}

// CONTRACTRESERVEDALIASSTORAGE is a free data retrieval call binding the contract method 0x70a94aa7.
//
// Solidity: function CONTRACT_RESERVED_ALIAS_STORAGE() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplSession) CONTRACTRESERVEDALIASSTORAGE() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.CONTRACTRESERVEDALIASSTORAGE(&_BitbonSupportImpl.CallOpts)
}

// CONTRACTRESERVEDALIASSTORAGE is a free data retrieval call binding the contract method 0x70a94aa7.
//
// Solidity: function CONTRACT_RESERVED_ALIAS_STORAGE() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCallerSession) CONTRACTRESERVEDALIASSTORAGE() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.CONTRACTRESERVEDALIASSTORAGE(&_BitbonSupportImpl.CallOpts)
}

// CONTRACTROLESTORAGE is a free data retrieval call binding the contract method 0xb938235b.
//
// Solidity: function CONTRACT_ROLE_STORAGE() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCaller) CONTRACTROLESTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonSupportImpl.contract.Call(opts, out, "CONTRACT_ROLE_STORAGE")
	return *ret0, err
}

// CONTRACTROLESTORAGE is a free data retrieval call binding the contract method 0xb938235b.
//
// Solidity: function CONTRACT_ROLE_STORAGE() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplSession) CONTRACTROLESTORAGE() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.CONTRACTROLESTORAGE(&_BitbonSupportImpl.CallOpts)
}

// CONTRACTROLESTORAGE is a free data retrieval call binding the contract method 0xb938235b.
//
// Solidity: function CONTRACT_ROLE_STORAGE() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCallerSession) CONTRACTROLESTORAGE() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.CONTRACTROLESTORAGE(&_BitbonSupportImpl.CallOpts)
}

// CONTRACTSAFETRANSFERSTORAGE is a free data retrieval call binding the contract method 0xe7e45392.
//
// Solidity: function CONTRACT_SAFE_TRANSFER_STORAGE() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCaller) CONTRACTSAFETRANSFERSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonSupportImpl.contract.Call(opts, out, "CONTRACT_SAFE_TRANSFER_STORAGE")
	return *ret0, err
}

// CONTRACTSAFETRANSFERSTORAGE is a free data retrieval call binding the contract method 0xe7e45392.
//
// Solidity: function CONTRACT_SAFE_TRANSFER_STORAGE() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplSession) CONTRACTSAFETRANSFERSTORAGE() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.CONTRACTSAFETRANSFERSTORAGE(&_BitbonSupportImpl.CallOpts)
}

// CONTRACTSAFETRANSFERSTORAGE is a free data retrieval call binding the contract method 0xe7e45392.
//
// Solidity: function CONTRACT_SAFE_TRANSFER_STORAGE() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCallerSession) CONTRACTSAFETRANSFERSTORAGE() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.CONTRACTSAFETRANSFERSTORAGE(&_BitbonSupportImpl.CallOpts)
}

// CONTRACTTRANSFER is a free data retrieval call binding the contract method 0xef94fe38.
//
// Solidity: function CONTRACT_TRANSFER() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCaller) CONTRACTTRANSFER(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonSupportImpl.contract.Call(opts, out, "CONTRACT_TRANSFER")
	return *ret0, err
}

// CONTRACTTRANSFER is a free data retrieval call binding the contract method 0xef94fe38.
//
// Solidity: function CONTRACT_TRANSFER() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplSession) CONTRACTTRANSFER() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.CONTRACTTRANSFER(&_BitbonSupportImpl.CallOpts)
}

// CONTRACTTRANSFER is a free data retrieval call binding the contract method 0xef94fe38.
//
// Solidity: function CONTRACT_TRANSFER() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCallerSession) CONTRACTTRANSFER() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.CONTRACTTRANSFER(&_BitbonSupportImpl.CallOpts)
}

// CONTRACTTRANSFERFROMCAPITALIZATION is a free data retrieval call binding the contract method 0xc8f8decb.
//
// Solidity: function CONTRACT_TRANSFER_FROM_CAPITALIZATION() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCaller) CONTRACTTRANSFERFROMCAPITALIZATION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonSupportImpl.contract.Call(opts, out, "CONTRACT_TRANSFER_FROM_CAPITALIZATION")
	return *ret0, err
}

// CONTRACTTRANSFERFROMCAPITALIZATION is a free data retrieval call binding the contract method 0xc8f8decb.
//
// Solidity: function CONTRACT_TRANSFER_FROM_CAPITALIZATION() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplSession) CONTRACTTRANSFERFROMCAPITALIZATION() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.CONTRACTTRANSFERFROMCAPITALIZATION(&_BitbonSupportImpl.CallOpts)
}

// CONTRACTTRANSFERFROMCAPITALIZATION is a free data retrieval call binding the contract method 0xc8f8decb.
//
// Solidity: function CONTRACT_TRANSFER_FROM_CAPITALIZATION() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCallerSession) CONTRACTTRANSFERFROMCAPITALIZATION() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.CONTRACTTRANSFERFROMCAPITALIZATION(&_BitbonSupportImpl.CallOpts)
}

// PERMISSIONACCESSRESTORATION is a free data retrieval call binding the contract method 0x5c3053bf.
//
// Solidity: function PERMISSION_ACCESS_RESTORATION() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCaller) PERMISSIONACCESSRESTORATION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonSupportImpl.contract.Call(opts, out, "PERMISSION_ACCESS_RESTORATION")
	return *ret0, err
}

// PERMISSIONACCESSRESTORATION is a free data retrieval call binding the contract method 0x5c3053bf.
//
// Solidity: function PERMISSION_ACCESS_RESTORATION() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplSession) PERMISSIONACCESSRESTORATION() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.PERMISSIONACCESSRESTORATION(&_BitbonSupportImpl.CallOpts)
}

// PERMISSIONACCESSRESTORATION is a free data retrieval call binding the contract method 0x5c3053bf.
//
// Solidity: function PERMISSION_ACCESS_RESTORATION() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCallerSession) PERMISSIONACCESSRESTORATION() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.PERMISSIONACCESSRESTORATION(&_BitbonSupportImpl.CallOpts)
}

// PERMISSIONADMINSTORAGE is a free data retrieval call binding the contract method 0x0c9fc0db.
//
// Solidity: function PERMISSION_ADMIN_STORAGE() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCaller) PERMISSIONADMINSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonSupportImpl.contract.Call(opts, out, "PERMISSION_ADMIN_STORAGE")
	return *ret0, err
}

// PERMISSIONADMINSTORAGE is a free data retrieval call binding the contract method 0x0c9fc0db.
//
// Solidity: function PERMISSION_ADMIN_STORAGE() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplSession) PERMISSIONADMINSTORAGE() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.PERMISSIONADMINSTORAGE(&_BitbonSupportImpl.CallOpts)
}

// PERMISSIONADMINSTORAGE is a free data retrieval call binding the contract method 0x0c9fc0db.
//
// Solidity: function PERMISSION_ADMIN_STORAGE() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCallerSession) PERMISSIONADMINSTORAGE() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.PERMISSIONADMINSTORAGE(&_BitbonSupportImpl.CallOpts)
}

// PERMISSIONCONTRACTSSTORAGE is a free data retrieval call binding the contract method 0x1516e83b.
//
// Solidity: function PERMISSION_CONTRACTS_STORAGE() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCaller) PERMISSIONCONTRACTSSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonSupportImpl.contract.Call(opts, out, "PERMISSION_CONTRACTS_STORAGE")
	return *ret0, err
}

// PERMISSIONCONTRACTSSTORAGE is a free data retrieval call binding the contract method 0x1516e83b.
//
// Solidity: function PERMISSION_CONTRACTS_STORAGE() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplSession) PERMISSIONCONTRACTSSTORAGE() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.PERMISSIONCONTRACTSSTORAGE(&_BitbonSupportImpl.CallOpts)
}

// PERMISSIONCONTRACTSSTORAGE is a free data retrieval call binding the contract method 0x1516e83b.
//
// Solidity: function PERMISSION_CONTRACTS_STORAGE() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCallerSession) PERMISSIONCONTRACTSSTORAGE() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.PERMISSIONCONTRACTSSTORAGE(&_BitbonSupportImpl.CallOpts)
}

// PERMISSIONCONTRACTDEPLOY is a free data retrieval call binding the contract method 0x257a4a04.
//
// Solidity: function PERMISSION_CONTRACT_DEPLOY() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCaller) PERMISSIONCONTRACTDEPLOY(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonSupportImpl.contract.Call(opts, out, "PERMISSION_CONTRACT_DEPLOY")
	return *ret0, err
}

// PERMISSIONCONTRACTDEPLOY is a free data retrieval call binding the contract method 0x257a4a04.
//
// Solidity: function PERMISSION_CONTRACT_DEPLOY() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplSession) PERMISSIONCONTRACTDEPLOY() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.PERMISSIONCONTRACTDEPLOY(&_BitbonSupportImpl.CallOpts)
}

// PERMISSIONCONTRACTDEPLOY is a free data retrieval call binding the contract method 0x257a4a04.
//
// Solidity: function PERMISSION_CONTRACT_DEPLOY() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCallerSession) PERMISSIONCONTRACTDEPLOY() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.PERMISSIONCONTRACTDEPLOY(&_BitbonSupportImpl.CallOpts)
}

// PERMISSIONEMISSION is a free data retrieval call binding the contract method 0x4ccd92d1.
//
// Solidity: function PERMISSION_EMISSION() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCaller) PERMISSIONEMISSION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonSupportImpl.contract.Call(opts, out, "PERMISSION_EMISSION")
	return *ret0, err
}

// PERMISSIONEMISSION is a free data retrieval call binding the contract method 0x4ccd92d1.
//
// Solidity: function PERMISSION_EMISSION() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplSession) PERMISSIONEMISSION() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.PERMISSIONEMISSION(&_BitbonSupportImpl.CallOpts)
}

// PERMISSIONEMISSION is a free data retrieval call binding the contract method 0x4ccd92d1.
//
// Solidity: function PERMISSION_EMISSION() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCallerSession) PERMISSIONEMISSION() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.PERMISSIONEMISSION(&_BitbonSupportImpl.CallOpts)
}

// PERMISSIONFEESSTORAGE is a free data retrieval call binding the contract method 0xe837eb87.
//
// Solidity: function PERMISSION_FEES_STORAGE() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCaller) PERMISSIONFEESSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonSupportImpl.contract.Call(opts, out, "PERMISSION_FEES_STORAGE")
	return *ret0, err
}

// PERMISSIONFEESSTORAGE is a free data retrieval call binding the contract method 0xe837eb87.
//
// Solidity: function PERMISSION_FEES_STORAGE() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplSession) PERMISSIONFEESSTORAGE() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.PERMISSIONFEESSTORAGE(&_BitbonSupportImpl.CallOpts)
}

// PERMISSIONFEESSTORAGE is a free data retrieval call binding the contract method 0xe837eb87.
//
// Solidity: function PERMISSION_FEES_STORAGE() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCallerSession) PERMISSIONFEESSTORAGE() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.PERMISSIONFEESSTORAGE(&_BitbonSupportImpl.CallOpts)
}

// PERMISSIONMULTISIGSTORAGE is a free data retrieval call binding the contract method 0xf1bb7aad.
//
// Solidity: function PERMISSION_MULTISIG_STORAGE() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCaller) PERMISSIONMULTISIGSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonSupportImpl.contract.Call(opts, out, "PERMISSION_MULTISIG_STORAGE")
	return *ret0, err
}

// PERMISSIONMULTISIGSTORAGE is a free data retrieval call binding the contract method 0xf1bb7aad.
//
// Solidity: function PERMISSION_MULTISIG_STORAGE() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplSession) PERMISSIONMULTISIGSTORAGE() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.PERMISSIONMULTISIGSTORAGE(&_BitbonSupportImpl.CallOpts)
}

// PERMISSIONMULTISIGSTORAGE is a free data retrieval call binding the contract method 0xf1bb7aad.
//
// Solidity: function PERMISSION_MULTISIG_STORAGE() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCallerSession) PERMISSIONMULTISIGSTORAGE() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.PERMISSIONMULTISIGSTORAGE(&_BitbonSupportImpl.CallOpts)
}

// PERMISSIONROLESSTORAGE is a free data retrieval call binding the contract method 0x12a16032.
//
// Solidity: function PERMISSION_ROLES_STORAGE() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCaller) PERMISSIONROLESSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonSupportImpl.contract.Call(opts, out, "PERMISSION_ROLES_STORAGE")
	return *ret0, err
}

// PERMISSIONROLESSTORAGE is a free data retrieval call binding the contract method 0x12a16032.
//
// Solidity: function PERMISSION_ROLES_STORAGE() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplSession) PERMISSIONROLESSTORAGE() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.PERMISSIONROLESSTORAGE(&_BitbonSupportImpl.CallOpts)
}

// PERMISSIONROLESSTORAGE is a free data retrieval call binding the contract method 0x12a16032.
//
// Solidity: function PERMISSION_ROLES_STORAGE() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCallerSession) PERMISSIONROLESSTORAGE() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.PERMISSIONROLESSTORAGE(&_BitbonSupportImpl.CallOpts)
}

// ROLEACCESSVERIFIER is a free data retrieval call binding the contract method 0xc5405f92.
//
// Solidity: function ROLE_ACCESS_VERIFIER() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCaller) ROLEACCESSVERIFIER(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonSupportImpl.contract.Call(opts, out, "ROLE_ACCESS_VERIFIER")
	return *ret0, err
}

// ROLEACCESSVERIFIER is a free data retrieval call binding the contract method 0xc5405f92.
//
// Solidity: function ROLE_ACCESS_VERIFIER() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplSession) ROLEACCESSVERIFIER() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.ROLEACCESSVERIFIER(&_BitbonSupportImpl.CallOpts)
}

// ROLEACCESSVERIFIER is a free data retrieval call binding the contract method 0xc5405f92.
//
// Solidity: function ROLE_ACCESS_VERIFIER() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCallerSession) ROLEACCESSVERIFIER() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.ROLEACCESSVERIFIER(&_BitbonSupportImpl.CallOpts)
}

// ROLEBITBONISSUEVERIFIER is a free data retrieval call binding the contract method 0xf0a012a6.
//
// Solidity: function ROLE_BITBON_ISSUE_VERIFIER() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCaller) ROLEBITBONISSUEVERIFIER(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonSupportImpl.contract.Call(opts, out, "ROLE_BITBON_ISSUE_VERIFIER")
	return *ret0, err
}

// ROLEBITBONISSUEVERIFIER is a free data retrieval call binding the contract method 0xf0a012a6.
//
// Solidity: function ROLE_BITBON_ISSUE_VERIFIER() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplSession) ROLEBITBONISSUEVERIFIER() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.ROLEBITBONISSUEVERIFIER(&_BitbonSupportImpl.CallOpts)
}

// ROLEBITBONISSUEVERIFIER is a free data retrieval call binding the contract method 0xf0a012a6.
//
// Solidity: function ROLE_BITBON_ISSUE_VERIFIER() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCallerSession) ROLEBITBONISSUEVERIFIER() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.ROLEBITBONISSUEVERIFIER(&_BitbonSupportImpl.CallOpts)
}

// ROLECOMMISSIONVERIFIER is a free data retrieval call binding the contract method 0x8ab3de8c.
//
// Solidity: function ROLE_COMMISSION_VERIFIER() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCaller) ROLECOMMISSIONVERIFIER(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonSupportImpl.contract.Call(opts, out, "ROLE_COMMISSION_VERIFIER")
	return *ret0, err
}

// ROLECOMMISSIONVERIFIER is a free data retrieval call binding the contract method 0x8ab3de8c.
//
// Solidity: function ROLE_COMMISSION_VERIFIER() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplSession) ROLECOMMISSIONVERIFIER() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.ROLECOMMISSIONVERIFIER(&_BitbonSupportImpl.CallOpts)
}

// ROLECOMMISSIONVERIFIER is a free data retrieval call binding the contract method 0x8ab3de8c.
//
// Solidity: function ROLE_COMMISSION_VERIFIER() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCallerSession) ROLECOMMISSIONVERIFIER() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.ROLECOMMISSIONVERIFIER(&_BitbonSupportImpl.CallOpts)
}

// ROLEDEPLOYADMIN is a free data retrieval call binding the contract method 0x667bf6e5.
//
// Solidity: function ROLE_DEPLOY_ADMIN() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCaller) ROLEDEPLOYADMIN(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonSupportImpl.contract.Call(opts, out, "ROLE_DEPLOY_ADMIN")
	return *ret0, err
}

// ROLEDEPLOYADMIN is a free data retrieval call binding the contract method 0x667bf6e5.
//
// Solidity: function ROLE_DEPLOY_ADMIN() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplSession) ROLEDEPLOYADMIN() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.ROLEDEPLOYADMIN(&_BitbonSupportImpl.CallOpts)
}

// ROLEDEPLOYADMIN is a free data retrieval call binding the contract method 0x667bf6e5.
//
// Solidity: function ROLE_DEPLOY_ADMIN() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCallerSession) ROLEDEPLOYADMIN() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.ROLEDEPLOYADMIN(&_BitbonSupportImpl.CallOpts)
}

// ROLEPERMISSIONADMIN is a free data retrieval call binding the contract method 0x1b239996.
//
// Solidity: function ROLE_PERMISSION_ADMIN() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCaller) ROLEPERMISSIONADMIN(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonSupportImpl.contract.Call(opts, out, "ROLE_PERMISSION_ADMIN")
	return *ret0, err
}

// ROLEPERMISSIONADMIN is a free data retrieval call binding the contract method 0x1b239996.
//
// Solidity: function ROLE_PERMISSION_ADMIN() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplSession) ROLEPERMISSIONADMIN() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.ROLEPERMISSIONADMIN(&_BitbonSupportImpl.CallOpts)
}

// ROLEPERMISSIONADMIN is a free data retrieval call binding the contract method 0x1b239996.
//
// Solidity: function ROLE_PERMISSION_ADMIN() view returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCallerSession) ROLEPERMISSIONADMIN() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.ROLEPERMISSIONADMIN(&_BitbonSupportImpl.CallOpts)
}

// CapitalizationAccount is a free data retrieval call binding the contract method 0x3da8ca46.
//
// Solidity: function capitalizationAccount() view returns(address)
func (_BitbonSupportImpl *BitbonSupportImplCaller) CapitalizationAccount(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BitbonSupportImpl.contract.Call(opts, out, "capitalizationAccount")
	return *ret0, err
}

// CapitalizationAccount is a free data retrieval call binding the contract method 0x3da8ca46.
//
// Solidity: function capitalizationAccount() view returns(address)
func (_BitbonSupportImpl *BitbonSupportImplSession) CapitalizationAccount() (common.Address, error) {
	return _BitbonSupportImpl.Contract.CapitalizationAccount(&_BitbonSupportImpl.CallOpts)
}

// CapitalizationAccount is a free data retrieval call binding the contract method 0x3da8ca46.
//
// Solidity: function capitalizationAccount() view returns(address)
func (_BitbonSupportImpl *BitbonSupportImplCallerSession) CapitalizationAccount() (common.Address, error) {
	return _BitbonSupportImpl.Contract.CapitalizationAccount(&_BitbonSupportImpl.CallOpts)
}

// ContractStorage is a free data retrieval call binding the contract method 0x0549b27e.
//
// Solidity: function contractStorage() view returns(address)
func (_BitbonSupportImpl *BitbonSupportImplCaller) ContractStorage(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BitbonSupportImpl.contract.Call(opts, out, "contractStorage")
	return *ret0, err
}

// ContractStorage is a free data retrieval call binding the contract method 0x0549b27e.
//
// Solidity: function contractStorage() view returns(address)
func (_BitbonSupportImpl *BitbonSupportImplSession) ContractStorage() (common.Address, error) {
	return _BitbonSupportImpl.Contract.ContractStorage(&_BitbonSupportImpl.CallOpts)
}

// ContractStorage is a free data retrieval call binding the contract method 0x0549b27e.
//
// Solidity: function contractStorage() view returns(address)
func (_BitbonSupportImpl *BitbonSupportImplCallerSession) ContractStorage() (common.Address, error) {
	return _BitbonSupportImpl.Contract.ContractStorage(&_BitbonSupportImpl.CallOpts)
}

// GetCapitalizationAccount is a free data retrieval call binding the contract method 0x2c41f29e.
//
// Solidity: function getCapitalizationAccount() view returns(address)
func (_BitbonSupportImpl *BitbonSupportImplCaller) GetCapitalizationAccount(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BitbonSupportImpl.contract.Call(opts, out, "getCapitalizationAccount")
	return *ret0, err
}

// GetCapitalizationAccount is a free data retrieval call binding the contract method 0x2c41f29e.
//
// Solidity: function getCapitalizationAccount() view returns(address)
func (_BitbonSupportImpl *BitbonSupportImplSession) GetCapitalizationAccount() (common.Address, error) {
	return _BitbonSupportImpl.Contract.GetCapitalizationAccount(&_BitbonSupportImpl.CallOpts)
}

// GetCapitalizationAccount is a free data retrieval call binding the contract method 0x2c41f29e.
//
// Solidity: function getCapitalizationAccount() view returns(address)
func (_BitbonSupportImpl *BitbonSupportImplCallerSession) GetCapitalizationAccount() (common.Address, error) {
	return _BitbonSupportImpl.Contract.GetCapitalizationAccount(&_BitbonSupportImpl.CallOpts)
}

// GetSimcordAccount is a free data retrieval call binding the contract method 0x2927a396.
//
// Solidity: function getSimcordAccount() view returns(address)
func (_BitbonSupportImpl *BitbonSupportImplCaller) GetSimcordAccount(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BitbonSupportImpl.contract.Call(opts, out, "getSimcordAccount")
	return *ret0, err
}

// GetSimcordAccount is a free data retrieval call binding the contract method 0x2927a396.
//
// Solidity: function getSimcordAccount() view returns(address)
func (_BitbonSupportImpl *BitbonSupportImplSession) GetSimcordAccount() (common.Address, error) {
	return _BitbonSupportImpl.Contract.GetSimcordAccount(&_BitbonSupportImpl.CallOpts)
}

// GetSimcordAccount is a free data retrieval call binding the contract method 0x2927a396.
//
// Solidity: function getSimcordAccount() view returns(address)
func (_BitbonSupportImpl *BitbonSupportImplCallerSession) GetSimcordAccount() (common.Address, error) {
	return _BitbonSupportImpl.Contract.GetSimcordAccount(&_BitbonSupportImpl.CallOpts)
}

// GetThisContractIndex is a free data retrieval call binding the contract method 0x70c6d367.
//
// Solidity: function getThisContractIndex() pure returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCaller) GetThisContractIndex(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonSupportImpl.contract.Call(opts, out, "getThisContractIndex")
	return *ret0, err
}

// GetThisContractIndex is a free data retrieval call binding the contract method 0x70c6d367.
//
// Solidity: function getThisContractIndex() pure returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplSession) GetThisContractIndex() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.GetThisContractIndex(&_BitbonSupportImpl.CallOpts)
}

// GetThisContractIndex is a free data retrieval call binding the contract method 0x70c6d367.
//
// Solidity: function getThisContractIndex() pure returns(uint256)
func (_BitbonSupportImpl *BitbonSupportImplCallerSession) GetThisContractIndex() (*big.Int, error) {
	return _BitbonSupportImpl.Contract.GetThisContractIndex(&_BitbonSupportImpl.CallOpts)
}

// SimcordAccount is a free data retrieval call binding the contract method 0x95f1125b.
//
// Solidity: function simcordAccount() view returns(address)
func (_BitbonSupportImpl *BitbonSupportImplCaller) SimcordAccount(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BitbonSupportImpl.contract.Call(opts, out, "simcordAccount")
	return *ret0, err
}

// SimcordAccount is a free data retrieval call binding the contract method 0x95f1125b.
//
// Solidity: function simcordAccount() view returns(address)
func (_BitbonSupportImpl *BitbonSupportImplSession) SimcordAccount() (common.Address, error) {
	return _BitbonSupportImpl.Contract.SimcordAccount(&_BitbonSupportImpl.CallOpts)
}

// SimcordAccount is a free data retrieval call binding the contract method 0x95f1125b.
//
// Solidity: function simcordAccount() view returns(address)
func (_BitbonSupportImpl *BitbonSupportImplCallerSession) SimcordAccount() (common.Address, error) {
	return _BitbonSupportImpl.Contract.SimcordAccount(&_BitbonSupportImpl.CallOpts)
}

// SetBalanceAdmin is a paid mutator transaction binding the contract method 0x9580811c.
//
// Solidity: function setBalanceAdmin(address assetbox, uint256 value) returns()
func (_BitbonSupportImpl *BitbonSupportImplTransactor) SetBalanceAdmin(opts *bind.TransactOpts, assetbox common.Address, value *big.Int) (*types.Transaction, error) {
	return _BitbonSupportImpl.contract.Transact(opts, "setBalanceAdmin", assetbox, value)
}

// SetBalanceAdmin is a paid mutator transaction binding the contract method 0x9580811c.
//
// Solidity: function setBalanceAdmin(address assetbox, uint256 value) returns()
func (_BitbonSupportImpl *BitbonSupportImplSession) SetBalanceAdmin(assetbox common.Address, value *big.Int) (*types.Transaction, error) {
	return _BitbonSupportImpl.Contract.SetBalanceAdmin(&_BitbonSupportImpl.TransactOpts, assetbox, value)
}

// SetBalanceAdmin is a paid mutator transaction binding the contract method 0x9580811c.
//
// Solidity: function setBalanceAdmin(address assetbox, uint256 value) returns()
func (_BitbonSupportImpl *BitbonSupportImplTransactorSession) SetBalanceAdmin(assetbox common.Address, value *big.Int) (*types.Transaction, error) {
	return _BitbonSupportImpl.Contract.SetBalanceAdmin(&_BitbonSupportImpl.TransactOpts, assetbox, value)
}

// SetBalancesAdmin is a paid mutator transaction binding the contract method 0x05741fed.
//
// Solidity: function setBalancesAdmin(address[] assetboxes, uint256[] values) returns()
func (_BitbonSupportImpl *BitbonSupportImplTransactor) SetBalancesAdmin(opts *bind.TransactOpts, assetboxes []common.Address, values []*big.Int) (*types.Transaction, error) {
	return _BitbonSupportImpl.contract.Transact(opts, "setBalancesAdmin", assetboxes, values)
}

// SetBalancesAdmin is a paid mutator transaction binding the contract method 0x05741fed.
//
// Solidity: function setBalancesAdmin(address[] assetboxes, uint256[] values) returns()
func (_BitbonSupportImpl *BitbonSupportImplSession) SetBalancesAdmin(assetboxes []common.Address, values []*big.Int) (*types.Transaction, error) {
	return _BitbonSupportImpl.Contract.SetBalancesAdmin(&_BitbonSupportImpl.TransactOpts, assetboxes, values)
}

// SetBalancesAdmin is a paid mutator transaction binding the contract method 0x05741fed.
//
// Solidity: function setBalancesAdmin(address[] assetboxes, uint256[] values) returns()
func (_BitbonSupportImpl *BitbonSupportImplTransactorSession) SetBalancesAdmin(assetboxes []common.Address, values []*big.Int) (*types.Transaction, error) {
	return _BitbonSupportImpl.Contract.SetBalancesAdmin(&_BitbonSupportImpl.TransactOpts, assetboxes, values)
}

// SetMigrated is a paid mutator transaction binding the contract method 0x925382c0.
//
// Solidity: function setMigrated() returns()
func (_BitbonSupportImpl *BitbonSupportImplTransactor) SetMigrated(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BitbonSupportImpl.contract.Transact(opts, "setMigrated")
}

// SetMigrated is a paid mutator transaction binding the contract method 0x925382c0.
//
// Solidity: function setMigrated() returns()
func (_BitbonSupportImpl *BitbonSupportImplSession) SetMigrated() (*types.Transaction, error) {
	return _BitbonSupportImpl.Contract.SetMigrated(&_BitbonSupportImpl.TransactOpts)
}

// SetMigrated is a paid mutator transaction binding the contract method 0x925382c0.
//
// Solidity: function setMigrated() returns()
func (_BitbonSupportImpl *BitbonSupportImplTransactorSession) SetMigrated() (*types.Transaction, error) {
	return _BitbonSupportImpl.Contract.SetMigrated(&_BitbonSupportImpl.TransactOpts)
}

// TransferFrom is a paid mutator transaction binding the contract method 0xab67aa58.
//
// Solidity: function transferFrom(address from, address to, uint256 value, bytes extraData) returns(bool)
func (_BitbonSupportImpl *BitbonSupportImplTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int, extraData []byte) (*types.Transaction, error) {
	return _BitbonSupportImpl.contract.Transact(opts, "transferFrom", from, to, value, extraData)
}

// TransferFrom is a paid mutator transaction binding the contract method 0xab67aa58.
//
// Solidity: function transferFrom(address from, address to, uint256 value, bytes extraData) returns(bool)
func (_BitbonSupportImpl *BitbonSupportImplSession) TransferFrom(from common.Address, to common.Address, value *big.Int, extraData []byte) (*types.Transaction, error) {
	return _BitbonSupportImpl.Contract.TransferFrom(&_BitbonSupportImpl.TransactOpts, from, to, value, extraData)
}

// TransferFrom is a paid mutator transaction binding the contract method 0xab67aa58.
//
// Solidity: function transferFrom(address from, address to, uint256 value, bytes extraData) returns(bool)
func (_BitbonSupportImpl *BitbonSupportImplTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int, extraData []byte) (*types.Transaction, error) {
	return _BitbonSupportImpl.Contract.TransferFrom(&_BitbonSupportImpl.TransactOpts, from, to, value, extraData)
}
