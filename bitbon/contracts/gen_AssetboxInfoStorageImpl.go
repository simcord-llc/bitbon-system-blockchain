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

// AssetboxInfoStorageImplABI is the input ABI used to generate the binding from.
const AssetboxInfoStorageImplABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"contractStorage\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_DISTRIBUTION_STORAGE\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_ADMIN_STORAGE\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REPLENISH_EMISSION\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ASSETBOX_INFO_GATE\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_ROLES_STORAGE\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_CONTRACTS_STORAGE\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_PERMISSION_ADMIN\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_BITBON_SUPPORT\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_ADD_ROLE\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAdminStorage\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REMOVE_ROLE\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_CONTRACT_DEPLOY\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_TRANSACTION_STORAGE\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_STORAGE\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_DROP_PERMISSION\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"assetboxToOtherInfo\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_EMISSION\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REMOVE_ADMIN_ROLE\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_ACCESS_RESTORATION\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REPLENISH_CAPITALIZATION\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_DEPLOY_ADMIN\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_EDIT_FEE\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_RESERVED_ALIAS_STORAGE\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REMOVE_ADMIN_BY_INDEX\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"contractStorageAddress\",\"type\":\"address\"}],\"name\":\"setContractStorage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"assetboxToServiceId\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_ADD_ADMIN\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MINING_AGENT_STORAGE\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_COMMISSION_VERIFIER\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getContractStorage\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_SET_ADMIN_ROLE\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_EDIT_ADMIN\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_SET_PERMISSION\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_BITBON\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ROLE_STORAGE\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_ADD_FEE\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_ACCESS_VERIFIER\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_EDIT_ROLE\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REMOVE_ADMIN_BY_ADDRESS\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_TRANSFER_FROM_CAPITALIZATION\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ADMIN_STORAGE\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"assetboxToAlias\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_GENERATOR\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_FEE_STORAGE\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"aliasToAssetbox\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REMOVE_FEE\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_SAFE_TRANSFER_STORAGE\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_FEES_STORAGE\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_BITBON_ISSUE_VERIFIER\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_MULTISIG_STORAGE\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ASSETBOX_INFO_STORAGE\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REPLENISH_CERTIFICATE\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_REPLENISH_CAPITALIZATION\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"contractStorageAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"assetbox\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"alias\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"serviceId\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"otherInfo\",\"type\":\"bytes\"}],\"name\":\"AssetboxInfoSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"assetbox\",\"type\":\"address\"}],\"name\":\"AssetboxInfoDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"contractIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"expiresOn\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"MulisigCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"contractIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"expiresOn\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"MulisigSigned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"contractIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"expiresOn\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"MulisigSucceeded\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"assetbox\",\"type\":\"address\"},{\"name\":\"alias\",\"type\":\"string\"},{\"name\":\"serviceId\",\"type\":\"bytes\"},{\"name\":\"otherInfo\",\"type\":\"bytes\"}],\"name\":\"setAssetboxInfo\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"assetbox\",\"type\":\"address\"}],\"name\":\"deleteAssetboxInfo\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"assetbox\",\"type\":\"address\"}],\"name\":\"getAssetboxInfoByAssetbox\",\"outputs\":[{\"name\":\"assetboxAddress\",\"type\":\"address\"},{\"name\":\"assetboxAlias\",\"type\":\"string\"},{\"name\":\"serviceId\",\"type\":\"bytes\"},{\"name\":\"otherInfo\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"alias\",\"type\":\"string\"}],\"name\":\"getAssetboxInfoByAlias\",\"outputs\":[{\"name\":\"assetboxAddress\",\"type\":\"address\"},{\"name\":\"assetboxAlias\",\"type\":\"string\"},{\"name\":\"serviceId\",\"type\":\"bytes\"},{\"name\":\"otherInfo\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"assetbox\",\"type\":\"address\"},{\"name\":\"serviceId\",\"type\":\"bytes\"}],\"name\":\"isAssetboxBoundToService\",\"outputs\":[{\"name\":\"isBound\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"alias\",\"type\":\"string\"},{\"name\":\"serviceId\",\"type\":\"bytes\"}],\"name\":\"isAliasBoundToService\",\"outputs\":[{\"name\":\"isBound\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"alias\",\"type\":\"string\"}],\"name\":\"isAliasTaken\",\"outputs\":[{\"name\":\"taken\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"alias\",\"type\":\"bytes32\"}],\"name\":\"isSanitizedAliasTaken\",\"outputs\":[{\"name\":\"taken\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getThisContractIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// AssetboxInfoStorageImpl is an auto generated Go binding around an Ethereum contract.
type AssetboxInfoStorageImpl struct {
	AssetboxInfoStorageImplCaller     // Read-only binding to the contract
	AssetboxInfoStorageImplTransactor // Write-only binding to the contract
	AssetboxInfoStorageImplFilterer   // Log filterer for contract events
}

// AssetboxInfoStorageImplCaller is an auto generated read-only Go binding around an Ethereum contract.
type AssetboxInfoStorageImplCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssetboxInfoStorageImplTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AssetboxInfoStorageImplTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssetboxInfoStorageImplFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AssetboxInfoStorageImplFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssetboxInfoStorageImplSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AssetboxInfoStorageImplSession struct {
	Contract     *AssetboxInfoStorageImpl // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// AssetboxInfoStorageImplCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AssetboxInfoStorageImplCallerSession struct {
	Contract *AssetboxInfoStorageImplCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// AssetboxInfoStorageImplTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AssetboxInfoStorageImplTransactorSession struct {
	Contract     *AssetboxInfoStorageImplTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// AssetboxInfoStorageImplRaw is an auto generated low-level Go binding around an Ethereum contract.
type AssetboxInfoStorageImplRaw struct {
	Contract *AssetboxInfoStorageImpl // Generic contract binding to access the raw methods on
}

// AssetboxInfoStorageImplCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AssetboxInfoStorageImplCallerRaw struct {
	Contract *AssetboxInfoStorageImplCaller // Generic read-only contract binding to access the raw methods on
}

// AssetboxInfoStorageImplTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AssetboxInfoStorageImplTransactorRaw struct {
	Contract *AssetboxInfoStorageImplTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAssetboxInfoStorageImpl creates a new instance of AssetboxInfoStorageImpl, bound to a specific deployed contract.
func NewAssetboxInfoStorageImpl(address common.Address, backend bind.ContractBackend) (*AssetboxInfoStorageImpl, error) {
	contract, err := bindAssetboxInfoStorageImpl(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AssetboxInfoStorageImpl{AssetboxInfoStorageImplCaller: AssetboxInfoStorageImplCaller{contract: contract}, AssetboxInfoStorageImplTransactor: AssetboxInfoStorageImplTransactor{contract: contract}, AssetboxInfoStorageImplFilterer: AssetboxInfoStorageImplFilterer{contract: contract}}, nil
}

// NewAssetboxInfoStorageImplCaller creates a new read-only instance of AssetboxInfoStorageImpl, bound to a specific deployed contract.
func NewAssetboxInfoStorageImplCaller(address common.Address, caller bind.ContractCaller) (*AssetboxInfoStorageImplCaller, error) {
	contract, err := bindAssetboxInfoStorageImpl(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AssetboxInfoStorageImplCaller{contract: contract}, nil
}

// NewAssetboxInfoStorageImplTransactor creates a new write-only instance of AssetboxInfoStorageImpl, bound to a specific deployed contract.
func NewAssetboxInfoStorageImplTransactor(address common.Address, transactor bind.ContractTransactor) (*AssetboxInfoStorageImplTransactor, error) {
	contract, err := bindAssetboxInfoStorageImpl(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AssetboxInfoStorageImplTransactor{contract: contract}, nil
}

// NewAssetboxInfoStorageImplFilterer creates a new log filterer instance of AssetboxInfoStorageImpl, bound to a specific deployed contract.
func NewAssetboxInfoStorageImplFilterer(address common.Address, filterer bind.ContractFilterer) (*AssetboxInfoStorageImplFilterer, error) {
	contract, err := bindAssetboxInfoStorageImpl(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AssetboxInfoStorageImplFilterer{contract: contract}, nil
}

// bindAssetboxInfoStorageImpl binds a generic wrapper to an already deployed contract.
func bindAssetboxInfoStorageImpl(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AssetboxInfoStorageImplABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AssetboxInfoStorageImpl.Contract.AssetboxInfoStorageImplCaller.contract.Call(opts, result, method, params...)
}

// Return Abi for AssetboxInfoStorageImplCaller
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCaller) ABI() abi.ABI {
	return _AssetboxInfoStorageImpl.contract.Abi
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AssetboxInfoStorageImpl.Contract.AssetboxInfoStorageImplTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AssetboxInfoStorageImpl.Contract.AssetboxInfoStorageImplTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AssetboxInfoStorageImpl.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AssetboxInfoStorageImpl.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AssetboxInfoStorageImpl.Contract.contract.Transact(opts, method, params...)
}

// CONTRACTADMINSTORAGE is a free data retrieval call binding the contract method 0xc9aebae9.
//
// Solidity: function CONTRACT_ADMIN_STORAGE() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCaller) CONTRACTADMINSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxInfoStorageImpl.contract.Call(opts, out, "CONTRACT_ADMIN_STORAGE")
	return *ret0, err
}

// CONTRACTADMINSTORAGE is a free data retrieval call binding the contract method 0xc9aebae9.
//
// Solidity: function CONTRACT_ADMIN_STORAGE() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplSession) CONTRACTADMINSTORAGE() (*big.Int, error) {
	return _AssetboxInfoStorageImpl.Contract.CONTRACTADMINSTORAGE(&_AssetboxInfoStorageImpl.CallOpts)
}

// CONTRACTADMINSTORAGE is a free data retrieval call binding the contract method 0xc9aebae9.
//
// Solidity: function CONTRACT_ADMIN_STORAGE() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCallerSession) CONTRACTADMINSTORAGE() (*big.Int, error) {
	return _AssetboxInfoStorageImpl.Contract.CONTRACTADMINSTORAGE(&_AssetboxInfoStorageImpl.CallOpts)
}

// CONTRACTASSETBOXINFOGATE is a free data retrieval call binding the contract method 0x113a7eb7.
//
// Solidity: function CONTRACT_ASSETBOX_INFO_GATE() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCaller) CONTRACTASSETBOXINFOGATE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxInfoStorageImpl.contract.Call(opts, out, "CONTRACT_ASSETBOX_INFO_GATE")
	return *ret0, err
}

// CONTRACTASSETBOXINFOGATE is a free data retrieval call binding the contract method 0x113a7eb7.
//
// Solidity: function CONTRACT_ASSETBOX_INFO_GATE() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplSession) CONTRACTASSETBOXINFOGATE() (*big.Int, error) {
	return _AssetboxInfoStorageImpl.Contract.CONTRACTASSETBOXINFOGATE(&_AssetboxInfoStorageImpl.CallOpts)
}

// CONTRACTASSETBOXINFOGATE is a free data retrieval call binding the contract method 0x113a7eb7.
//
// Solidity: function CONTRACT_ASSETBOX_INFO_GATE() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCallerSession) CONTRACTASSETBOXINFOGATE() (*big.Int, error) {
	return _AssetboxInfoStorageImpl.Contract.CONTRACTASSETBOXINFOGATE(&_AssetboxInfoStorageImpl.CallOpts)
}

// CONTRACTASSETBOXINFOSTORAGE is a free data retrieval call binding the contract method 0xf32a115d.
//
// Solidity: function CONTRACT_ASSETBOX_INFO_STORAGE() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCaller) CONTRACTASSETBOXINFOSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxInfoStorageImpl.contract.Call(opts, out, "CONTRACT_ASSETBOX_INFO_STORAGE")
	return *ret0, err
}

// CONTRACTASSETBOXINFOSTORAGE is a free data retrieval call binding the contract method 0xf32a115d.
//
// Solidity: function CONTRACT_ASSETBOX_INFO_STORAGE() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplSession) CONTRACTASSETBOXINFOSTORAGE() (*big.Int, error) {
	return _AssetboxInfoStorageImpl.Contract.CONTRACTASSETBOXINFOSTORAGE(&_AssetboxInfoStorageImpl.CallOpts)
}

// CONTRACTASSETBOXINFOSTORAGE is a free data retrieval call binding the contract method 0xf32a115d.
//
// Solidity: function CONTRACT_ASSETBOX_INFO_STORAGE() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCallerSession) CONTRACTASSETBOXINFOSTORAGE() (*big.Int, error) {
	return _AssetboxInfoStorageImpl.Contract.CONTRACTASSETBOXINFOSTORAGE(&_AssetboxInfoStorageImpl.CallOpts)
}

// CONTRACTBITBON is a free data retrieval call binding the contract method 0xb6b433ab.
//
// Solidity: function CONTRACT_BITBON() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCaller) CONTRACTBITBON(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxInfoStorageImpl.contract.Call(opts, out, "CONTRACT_BITBON")
	return *ret0, err
}

// CONTRACTBITBON is a free data retrieval call binding the contract method 0xb6b433ab.
//
// Solidity: function CONTRACT_BITBON() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplSession) CONTRACTBITBON() (*big.Int, error) {
	return _AssetboxInfoStorageImpl.Contract.CONTRACTBITBON(&_AssetboxInfoStorageImpl.CallOpts)
}

// CONTRACTBITBON is a free data retrieval call binding the contract method 0xb6b433ab.
//
// Solidity: function CONTRACT_BITBON() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCallerSession) CONTRACTBITBON() (*big.Int, error) {
	return _AssetboxInfoStorageImpl.Contract.CONTRACTBITBON(&_AssetboxInfoStorageImpl.CallOpts)
}

// CONTRACTBITBONSUPPORT is a free data retrieval call binding the contract method 0x20139c74.
//
// Solidity: function CONTRACT_BITBON_SUPPORT() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCaller) CONTRACTBITBONSUPPORT(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxInfoStorageImpl.contract.Call(opts, out, "CONTRACT_BITBON_SUPPORT")
	return *ret0, err
}

// CONTRACTBITBONSUPPORT is a free data retrieval call binding the contract method 0x20139c74.
//
// Solidity: function CONTRACT_BITBON_SUPPORT() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplSession) CONTRACTBITBONSUPPORT() (*big.Int, error) {
	return _AssetboxInfoStorageImpl.Contract.CONTRACTBITBONSUPPORT(&_AssetboxInfoStorageImpl.CallOpts)
}

// CONTRACTBITBONSUPPORT is a free data retrieval call binding the contract method 0x20139c74.
//
// Solidity: function CONTRACT_BITBON_SUPPORT() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCallerSession) CONTRACTBITBONSUPPORT() (*big.Int, error) {
	return _AssetboxInfoStorageImpl.Contract.CONTRACTBITBONSUPPORT(&_AssetboxInfoStorageImpl.CallOpts)
}

// CONTRACTDISTRIBUTIONSTORAGE is a free data retrieval call binding the contract method 0x08097327.
//
// Solidity: function CONTRACT_DISTRIBUTION_STORAGE() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCaller) CONTRACTDISTRIBUTIONSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxInfoStorageImpl.contract.Call(opts, out, "CONTRACT_DISTRIBUTION_STORAGE")
	return *ret0, err
}

// CONTRACTDISTRIBUTIONSTORAGE is a free data retrieval call binding the contract method 0x08097327.
//
// Solidity: function CONTRACT_DISTRIBUTION_STORAGE() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplSession) CONTRACTDISTRIBUTIONSTORAGE() (*big.Int, error) {
	return _AssetboxInfoStorageImpl.Contract.CONTRACTDISTRIBUTIONSTORAGE(&_AssetboxInfoStorageImpl.CallOpts)
}

// CONTRACTDISTRIBUTIONSTORAGE is a free data retrieval call binding the contract method 0x08097327.
//
// Solidity: function CONTRACT_DISTRIBUTION_STORAGE() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCallerSession) CONTRACTDISTRIBUTIONSTORAGE() (*big.Int, error) {
	return _AssetboxInfoStorageImpl.Contract.CONTRACTDISTRIBUTIONSTORAGE(&_AssetboxInfoStorageImpl.CallOpts)
}

// CONTRACTFEESTORAGE is a free data retrieval call binding the contract method 0xcd5e65a4.
//
// Solidity: function CONTRACT_FEE_STORAGE() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCaller) CONTRACTFEESTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxInfoStorageImpl.contract.Call(opts, out, "CONTRACT_FEE_STORAGE")
	return *ret0, err
}

// CONTRACTFEESTORAGE is a free data retrieval call binding the contract method 0xcd5e65a4.
//
// Solidity: function CONTRACT_FEE_STORAGE() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplSession) CONTRACTFEESTORAGE() (*big.Int, error) {
	return _AssetboxInfoStorageImpl.Contract.CONTRACTFEESTORAGE(&_AssetboxInfoStorageImpl.CallOpts)
}

// CONTRACTFEESTORAGE is a free data retrieval call binding the contract method 0xcd5e65a4.
//
// Solidity: function CONTRACT_FEE_STORAGE() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCallerSession) CONTRACTFEESTORAGE() (*big.Int, error) {
	return _AssetboxInfoStorageImpl.Contract.CONTRACTFEESTORAGE(&_AssetboxInfoStorageImpl.CallOpts)
}

// CONTRACTGENERATOR is a free data retrieval call binding the contract method 0xccea2b46.
//
// Solidity: function CONTRACT_GENERATOR() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCaller) CONTRACTGENERATOR(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxInfoStorageImpl.contract.Call(opts, out, "CONTRACT_GENERATOR")
	return *ret0, err
}

// CONTRACTGENERATOR is a free data retrieval call binding the contract method 0xccea2b46.
//
// Solidity: function CONTRACT_GENERATOR() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplSession) CONTRACTGENERATOR() (*big.Int, error) {
	return _AssetboxInfoStorageImpl.Contract.CONTRACTGENERATOR(&_AssetboxInfoStorageImpl.CallOpts)
}

// CONTRACTGENERATOR is a free data retrieval call binding the contract method 0xccea2b46.
//
// Solidity: function CONTRACT_GENERATOR() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCallerSession) CONTRACTGENERATOR() (*big.Int, error) {
	return _AssetboxInfoStorageImpl.Contract.CONTRACTGENERATOR(&_AssetboxInfoStorageImpl.CallOpts)
}

// CONTRACTMININGAGENTSTORAGE is a free data retrieval call binding the contract method 0x83cc588d.
//
// Solidity: function CONTRACT_MINING_AGENT_STORAGE() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCaller) CONTRACTMININGAGENTSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxInfoStorageImpl.contract.Call(opts, out, "CONTRACT_MINING_AGENT_STORAGE")
	return *ret0, err
}

// CONTRACTMININGAGENTSTORAGE is a free data retrieval call binding the contract method 0x83cc588d.
//
// Solidity: function CONTRACT_MINING_AGENT_STORAGE() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplSession) CONTRACTMININGAGENTSTORAGE() (*big.Int, error) {
	return _AssetboxInfoStorageImpl.Contract.CONTRACTMININGAGENTSTORAGE(&_AssetboxInfoStorageImpl.CallOpts)
}

// CONTRACTMININGAGENTSTORAGE is a free data retrieval call binding the contract method 0x83cc588d.
//
// Solidity: function CONTRACT_MINING_AGENT_STORAGE() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCallerSession) CONTRACTMININGAGENTSTORAGE() (*big.Int, error) {
	return _AssetboxInfoStorageImpl.Contract.CONTRACTMININGAGENTSTORAGE(&_AssetboxInfoStorageImpl.CallOpts)
}

// CONTRACTMULTISIGADDADMIN is a free data retrieval call binding the contract method 0x83888c68.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ADMIN() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCaller) CONTRACTMULTISIGADDADMIN(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxInfoStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_ADD_ADMIN")
	return *ret0, err
}

// CONTRACTMULTISIGADDADMIN is a free data retrieval call binding the contract method 0x83888c68.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ADMIN() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplSession) CONTRACTMULTISIGADDADMIN() (*big.Int, error) {
	return _AssetboxInfoStorageImpl.Contract.CONTRACTMULTISIGADDADMIN(&_AssetboxInfoStorageImpl.CallOpts)
}

// CONTRACTMULTISIGADDADMIN is a free data retrieval call binding the contract method 0x83888c68.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ADMIN() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCallerSession) CONTRACTMULTISIGADDADMIN() (*big.Int, error) {
	return _AssetboxInfoStorageImpl.Contract.CONTRACTMULTISIGADDADMIN(&_AssetboxInfoStorageImpl.CallOpts)
}

// CONTRACTMULTISIGADDFEE is a free data retrieval call binding the contract method 0xc43e8cab.
//
// Solidity: function CONTRACT_MULTISIG_ADD_FEE() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCaller) CONTRACTMULTISIGADDFEE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxInfoStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_ADD_FEE")
	return *ret0, err
}

// CONTRACTMULTISIGADDFEE is a free data retrieval call binding the contract method 0xc43e8cab.
//
// Solidity: function CONTRACT_MULTISIG_ADD_FEE() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplSession) CONTRACTMULTISIGADDFEE() (*big.Int, error) {
	return _AssetboxInfoStorageImpl.Contract.CONTRACTMULTISIGADDFEE(&_AssetboxInfoStorageImpl.CallOpts)
}

// CONTRACTMULTISIGADDFEE is a free data retrieval call binding the contract method 0xc43e8cab.
//
// Solidity: function CONTRACT_MULTISIG_ADD_FEE() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCallerSession) CONTRACTMULTISIGADDFEE() (*big.Int, error) {
	return _AssetboxInfoStorageImpl.Contract.CONTRACTMULTISIGADDFEE(&_AssetboxInfoStorageImpl.CallOpts)
}

// CONTRACTMULTISIGADDROLE is a free data retrieval call binding the contract method 0x20a0914b.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ROLE() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCaller) CONTRACTMULTISIGADDROLE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxInfoStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_ADD_ROLE")
	return *ret0, err
}

// CONTRACTMULTISIGADDROLE is a free data retrieval call binding the contract method 0x20a0914b.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ROLE() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplSession) CONTRACTMULTISIGADDROLE() (*big.Int, error) {
	return _AssetboxInfoStorageImpl.Contract.CONTRACTMULTISIGADDROLE(&_AssetboxInfoStorageImpl.CallOpts)
}

// CONTRACTMULTISIGADDROLE is a free data retrieval call binding the contract method 0x20a0914b.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ROLE() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCallerSession) CONTRACTMULTISIGADDROLE() (*big.Int, error) {
	return _AssetboxInfoStorageImpl.Contract.CONTRACTMULTISIGADDROLE(&_AssetboxInfoStorageImpl.CallOpts)
}

// CONTRACTMULTISIGDROPPERMISSION is a free data retrieval call binding the contract method 0x3f3fbc39.
//
// Solidity: function CONTRACT_MULTISIG_DROP_PERMISSION() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCaller) CONTRACTMULTISIGDROPPERMISSION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxInfoStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_DROP_PERMISSION")
	return *ret0, err
}

// CONTRACTMULTISIGDROPPERMISSION is a free data retrieval call binding the contract method 0x3f3fbc39.
//
// Solidity: function CONTRACT_MULTISIG_DROP_PERMISSION() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplSession) CONTRACTMULTISIGDROPPERMISSION() (*big.Int, error) {
	return _AssetboxInfoStorageImpl.Contract.CONTRACTMULTISIGDROPPERMISSION(&_AssetboxInfoStorageImpl.CallOpts)
}

// CONTRACTMULTISIGDROPPERMISSION is a free data retrieval call binding the contract method 0x3f3fbc39.
//
// Solidity: function CONTRACT_MULTISIG_DROP_PERMISSION() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCallerSession) CONTRACTMULTISIGDROPPERMISSION() (*big.Int, error) {
	return _AssetboxInfoStorageImpl.Contract.CONTRACTMULTISIGDROPPERMISSION(&_AssetboxInfoStorageImpl.CallOpts)
}

// CONTRACTMULTISIGEDITADMIN is a free data retrieval call binding the contract method 0x91b0f051.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_ADMIN() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCaller) CONTRACTMULTISIGEDITADMIN(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxInfoStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_EDIT_ADMIN")
	return *ret0, err
}

// CONTRACTMULTISIGEDITADMIN is a free data retrieval call binding the contract method 0x91b0f051.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_ADMIN() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplSession) CONTRACTMULTISIGEDITADMIN() (*big.Int, error) {
	return _AssetboxInfoStorageImpl.Contract.CONTRACTMULTISIGEDITADMIN(&_AssetboxInfoStorageImpl.CallOpts)
}

// CONTRACTMULTISIGEDITADMIN is a free data retrieval call binding the contract method 0x91b0f051.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_ADMIN() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCallerSession) CONTRACTMULTISIGEDITADMIN() (*big.Int, error) {
	return _AssetboxInfoStorageImpl.Contract.CONTRACTMULTISIGEDITADMIN(&_AssetboxInfoStorageImpl.CallOpts)
}

// CONTRACTMULTISIGEDITFEE is a free data retrieval call binding the contract method 0x6b2abe17.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_FEE() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCaller) CONTRACTMULTISIGEDITFEE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxInfoStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_EDIT_FEE")
	return *ret0, err
}

// CONTRACTMULTISIGEDITFEE is a free data retrieval call binding the contract method 0x6b2abe17.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_FEE() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplSession) CONTRACTMULTISIGEDITFEE() (*big.Int, error) {
	return _AssetboxInfoStorageImpl.Contract.CONTRACTMULTISIGEDITFEE(&_AssetboxInfoStorageImpl.CallOpts)
}

// CONTRACTMULTISIGEDITFEE is a free data retrieval call binding the contract method 0x6b2abe17.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_FEE() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCallerSession) CONTRACTMULTISIGEDITFEE() (*big.Int, error) {
	return _AssetboxInfoStorageImpl.Contract.CONTRACTMULTISIGEDITFEE(&_AssetboxInfoStorageImpl.CallOpts)
}

// CONTRACTMULTISIGEDITROLE is a free data retrieval call binding the contract method 0xc5f66412.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_ROLE() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCaller) CONTRACTMULTISIGEDITROLE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxInfoStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_EDIT_ROLE")
	return *ret0, err
}

// CONTRACTMULTISIGEDITROLE is a free data retrieval call binding the contract method 0xc5f66412.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_ROLE() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplSession) CONTRACTMULTISIGEDITROLE() (*big.Int, error) {
	return _AssetboxInfoStorageImpl.Contract.CONTRACTMULTISIGEDITROLE(&_AssetboxInfoStorageImpl.CallOpts)
}

// CONTRACTMULTISIGEDITROLE is a free data retrieval call binding the contract method 0xc5f66412.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_ROLE() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCallerSession) CONTRACTMULTISIGEDITROLE() (*big.Int, error) {
	return _AssetboxInfoStorageImpl.Contract.CONTRACTMULTISIGEDITROLE(&_AssetboxInfoStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEADMINBYADDRESS is a free data retrieval call binding the contract method 0xc816b091.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_BY_ADDRESS() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCaller) CONTRACTMULTISIGREMOVEADMINBYADDRESS(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxInfoStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REMOVE_ADMIN_BY_ADDRESS")
	return *ret0, err
}

// CONTRACTMULTISIGREMOVEADMINBYADDRESS is a free data retrieval call binding the contract method 0xc816b091.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_BY_ADDRESS() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplSession) CONTRACTMULTISIGREMOVEADMINBYADDRESS() (*big.Int, error) {
	return _AssetboxInfoStorageImpl.Contract.CONTRACTMULTISIGREMOVEADMINBYADDRESS(&_AssetboxInfoStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEADMINBYADDRESS is a free data retrieval call binding the contract method 0xc816b091.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_BY_ADDRESS() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCallerSession) CONTRACTMULTISIGREMOVEADMINBYADDRESS() (*big.Int, error) {
	return _AssetboxInfoStorageImpl.Contract.CONTRACTMULTISIGREMOVEADMINBYADDRESS(&_AssetboxInfoStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEADMINBYINDEX is a free data retrieval call binding the contract method 0x751fbea2.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_BY_INDEX() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCaller) CONTRACTMULTISIGREMOVEADMINBYINDEX(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxInfoStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REMOVE_ADMIN_BY_INDEX")
	return *ret0, err
}

// CONTRACTMULTISIGREMOVEADMINBYINDEX is a free data retrieval call binding the contract method 0x751fbea2.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_BY_INDEX() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplSession) CONTRACTMULTISIGREMOVEADMINBYINDEX() (*big.Int, error) {
	return _AssetboxInfoStorageImpl.Contract.CONTRACTMULTISIGREMOVEADMINBYINDEX(&_AssetboxInfoStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEADMINBYINDEX is a free data retrieval call binding the contract method 0x751fbea2.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_BY_INDEX() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCallerSession) CONTRACTMULTISIGREMOVEADMINBYINDEX() (*big.Int, error) {
	return _AssetboxInfoStorageImpl.Contract.CONTRACTMULTISIGREMOVEADMINBYINDEX(&_AssetboxInfoStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEADMINROLE is a free data retrieval call binding the contract method 0x53c00200.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_ROLE() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCaller) CONTRACTMULTISIGREMOVEADMINROLE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxInfoStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REMOVE_ADMIN_ROLE")
	return *ret0, err
}

// CONTRACTMULTISIGREMOVEADMINROLE is a free data retrieval call binding the contract method 0x53c00200.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_ROLE() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplSession) CONTRACTMULTISIGREMOVEADMINROLE() (*big.Int, error) {
	return _AssetboxInfoStorageImpl.Contract.CONTRACTMULTISIGREMOVEADMINROLE(&_AssetboxInfoStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEADMINROLE is a free data retrieval call binding the contract method 0x53c00200.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_ROLE() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCallerSession) CONTRACTMULTISIGREMOVEADMINROLE() (*big.Int, error) {
	return _AssetboxInfoStorageImpl.Contract.CONTRACTMULTISIGREMOVEADMINROLE(&_AssetboxInfoStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEFEE is a free data retrieval call binding the contract method 0xdb614d25.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_FEE() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCaller) CONTRACTMULTISIGREMOVEFEE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxInfoStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REMOVE_FEE")
	return *ret0, err
}

// CONTRACTMULTISIGREMOVEFEE is a free data retrieval call binding the contract method 0xdb614d25.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_FEE() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplSession) CONTRACTMULTISIGREMOVEFEE() (*big.Int, error) {
	return _AssetboxInfoStorageImpl.Contract.CONTRACTMULTISIGREMOVEFEE(&_AssetboxInfoStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEFEE is a free data retrieval call binding the contract method 0xdb614d25.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_FEE() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCallerSession) CONTRACTMULTISIGREMOVEFEE() (*big.Int, error) {
	return _AssetboxInfoStorageImpl.Contract.CONTRACTMULTISIGREMOVEFEE(&_AssetboxInfoStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEROLE is a free data retrieval call binding the contract method 0x23648496.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ROLE() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCaller) CONTRACTMULTISIGREMOVEROLE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxInfoStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REMOVE_ROLE")
	return *ret0, err
}

// CONTRACTMULTISIGREMOVEROLE is a free data retrieval call binding the contract method 0x23648496.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ROLE() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplSession) CONTRACTMULTISIGREMOVEROLE() (*big.Int, error) {
	return _AssetboxInfoStorageImpl.Contract.CONTRACTMULTISIGREMOVEROLE(&_AssetboxInfoStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEROLE is a free data retrieval call binding the contract method 0x23648496.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ROLE() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCallerSession) CONTRACTMULTISIGREMOVEROLE() (*big.Int, error) {
	return _AssetboxInfoStorageImpl.Contract.CONTRACTMULTISIGREMOVEROLE(&_AssetboxInfoStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREPLENISHCAPITALIZATION is a free data retrieval call binding the contract method 0x62d5d523.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_CAPITALIZATION() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCaller) CONTRACTMULTISIGREPLENISHCAPITALIZATION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxInfoStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REPLENISH_CAPITALIZATION")
	return *ret0, err
}

// CONTRACTMULTISIGREPLENISHCAPITALIZATION is a free data retrieval call binding the contract method 0x62d5d523.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_CAPITALIZATION() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplSession) CONTRACTMULTISIGREPLENISHCAPITALIZATION() (*big.Int, error) {
	return _AssetboxInfoStorageImpl.Contract.CONTRACTMULTISIGREPLENISHCAPITALIZATION(&_AssetboxInfoStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREPLENISHCAPITALIZATION is a free data retrieval call binding the contract method 0x62d5d523.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_CAPITALIZATION() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCallerSession) CONTRACTMULTISIGREPLENISHCAPITALIZATION() (*big.Int, error) {
	return _AssetboxInfoStorageImpl.Contract.CONTRACTMULTISIGREPLENISHCAPITALIZATION(&_AssetboxInfoStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREPLENISHCERTIFICATE is a free data retrieval call binding the contract method 0xfd19b7e1.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_CERTIFICATE() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCaller) CONTRACTMULTISIGREPLENISHCERTIFICATE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxInfoStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REPLENISH_CERTIFICATE")
	return *ret0, err
}

// CONTRACTMULTISIGREPLENISHCERTIFICATE is a free data retrieval call binding the contract method 0xfd19b7e1.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_CERTIFICATE() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplSession) CONTRACTMULTISIGREPLENISHCERTIFICATE() (*big.Int, error) {
	return _AssetboxInfoStorageImpl.Contract.CONTRACTMULTISIGREPLENISHCERTIFICATE(&_AssetboxInfoStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREPLENISHCERTIFICATE is a free data retrieval call binding the contract method 0xfd19b7e1.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_CERTIFICATE() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCallerSession) CONTRACTMULTISIGREPLENISHCERTIFICATE() (*big.Int, error) {
	return _AssetboxInfoStorageImpl.Contract.CONTRACTMULTISIGREPLENISHCERTIFICATE(&_AssetboxInfoStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREPLENISHEMISSION is a free data retrieval call binding the contract method 0x0cb59fde.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_EMISSION() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCaller) CONTRACTMULTISIGREPLENISHEMISSION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxInfoStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REPLENISH_EMISSION")
	return *ret0, err
}

// CONTRACTMULTISIGREPLENISHEMISSION is a free data retrieval call binding the contract method 0x0cb59fde.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_EMISSION() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplSession) CONTRACTMULTISIGREPLENISHEMISSION() (*big.Int, error) {
	return _AssetboxInfoStorageImpl.Contract.CONTRACTMULTISIGREPLENISHEMISSION(&_AssetboxInfoStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREPLENISHEMISSION is a free data retrieval call binding the contract method 0x0cb59fde.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_EMISSION() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCallerSession) CONTRACTMULTISIGREPLENISHEMISSION() (*big.Int, error) {
	return _AssetboxInfoStorageImpl.Contract.CONTRACTMULTISIGREPLENISHEMISSION(&_AssetboxInfoStorageImpl.CallOpts)
}

// CONTRACTMULTISIGSETADMINROLE is a free data retrieval call binding the contract method 0x8f209ab7.
//
// Solidity: function CONTRACT_MULTISIG_SET_ADMIN_ROLE() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCaller) CONTRACTMULTISIGSETADMINROLE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxInfoStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_SET_ADMIN_ROLE")
	return *ret0, err
}

// CONTRACTMULTISIGSETADMINROLE is a free data retrieval call binding the contract method 0x8f209ab7.
//
// Solidity: function CONTRACT_MULTISIG_SET_ADMIN_ROLE() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplSession) CONTRACTMULTISIGSETADMINROLE() (*big.Int, error) {
	return _AssetboxInfoStorageImpl.Contract.CONTRACTMULTISIGSETADMINROLE(&_AssetboxInfoStorageImpl.CallOpts)
}

// CONTRACTMULTISIGSETADMINROLE is a free data retrieval call binding the contract method 0x8f209ab7.
//
// Solidity: function CONTRACT_MULTISIG_SET_ADMIN_ROLE() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCallerSession) CONTRACTMULTISIGSETADMINROLE() (*big.Int, error) {
	return _AssetboxInfoStorageImpl.Contract.CONTRACTMULTISIGSETADMINROLE(&_AssetboxInfoStorageImpl.CallOpts)
}

// CONTRACTMULTISIGSETPERMISSION is a free data retrieval call binding the contract method 0xb0e91914.
//
// Solidity: function CONTRACT_MULTISIG_SET_PERMISSION() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCaller) CONTRACTMULTISIGSETPERMISSION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxInfoStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_SET_PERMISSION")
	return *ret0, err
}

// CONTRACTMULTISIGSETPERMISSION is a free data retrieval call binding the contract method 0xb0e91914.
//
// Solidity: function CONTRACT_MULTISIG_SET_PERMISSION() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplSession) CONTRACTMULTISIGSETPERMISSION() (*big.Int, error) {
	return _AssetboxInfoStorageImpl.Contract.CONTRACTMULTISIGSETPERMISSION(&_AssetboxInfoStorageImpl.CallOpts)
}

// CONTRACTMULTISIGSETPERMISSION is a free data retrieval call binding the contract method 0xb0e91914.
//
// Solidity: function CONTRACT_MULTISIG_SET_PERMISSION() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCallerSession) CONTRACTMULTISIGSETPERMISSION() (*big.Int, error) {
	return _AssetboxInfoStorageImpl.Contract.CONTRACTMULTISIGSETPERMISSION(&_AssetboxInfoStorageImpl.CallOpts)
}

// CONTRACTREPLENISHCAPITALIZATION is a free data retrieval call binding the contract method 0xff5e09d8.
//
// Solidity: function CONTRACT_REPLENISH_CAPITALIZATION() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCaller) CONTRACTREPLENISHCAPITALIZATION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxInfoStorageImpl.contract.Call(opts, out, "CONTRACT_REPLENISH_CAPITALIZATION")
	return *ret0, err
}

// CONTRACTREPLENISHCAPITALIZATION is a free data retrieval call binding the contract method 0xff5e09d8.
//
// Solidity: function CONTRACT_REPLENISH_CAPITALIZATION() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplSession) CONTRACTREPLENISHCAPITALIZATION() (*big.Int, error) {
	return _AssetboxInfoStorageImpl.Contract.CONTRACTREPLENISHCAPITALIZATION(&_AssetboxInfoStorageImpl.CallOpts)
}

// CONTRACTREPLENISHCAPITALIZATION is a free data retrieval call binding the contract method 0xff5e09d8.
//
// Solidity: function CONTRACT_REPLENISH_CAPITALIZATION() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCallerSession) CONTRACTREPLENISHCAPITALIZATION() (*big.Int, error) {
	return _AssetboxInfoStorageImpl.Contract.CONTRACTREPLENISHCAPITALIZATION(&_AssetboxInfoStorageImpl.CallOpts)
}

// CONTRACTRESERVEDALIASSTORAGE is a free data retrieval call binding the contract method 0x70a94aa7.
//
// Solidity: function CONTRACT_RESERVED_ALIAS_STORAGE() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCaller) CONTRACTRESERVEDALIASSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxInfoStorageImpl.contract.Call(opts, out, "CONTRACT_RESERVED_ALIAS_STORAGE")
	return *ret0, err
}

// CONTRACTRESERVEDALIASSTORAGE is a free data retrieval call binding the contract method 0x70a94aa7.
//
// Solidity: function CONTRACT_RESERVED_ALIAS_STORAGE() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplSession) CONTRACTRESERVEDALIASSTORAGE() (*big.Int, error) {
	return _AssetboxInfoStorageImpl.Contract.CONTRACTRESERVEDALIASSTORAGE(&_AssetboxInfoStorageImpl.CallOpts)
}

// CONTRACTRESERVEDALIASSTORAGE is a free data retrieval call binding the contract method 0x70a94aa7.
//
// Solidity: function CONTRACT_RESERVED_ALIAS_STORAGE() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCallerSession) CONTRACTRESERVEDALIASSTORAGE() (*big.Int, error) {
	return _AssetboxInfoStorageImpl.Contract.CONTRACTRESERVEDALIASSTORAGE(&_AssetboxInfoStorageImpl.CallOpts)
}

// CONTRACTROLESTORAGE is a free data retrieval call binding the contract method 0xb938235b.
//
// Solidity: function CONTRACT_ROLE_STORAGE() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCaller) CONTRACTROLESTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxInfoStorageImpl.contract.Call(opts, out, "CONTRACT_ROLE_STORAGE")
	return *ret0, err
}

// CONTRACTROLESTORAGE is a free data retrieval call binding the contract method 0xb938235b.
//
// Solidity: function CONTRACT_ROLE_STORAGE() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplSession) CONTRACTROLESTORAGE() (*big.Int, error) {
	return _AssetboxInfoStorageImpl.Contract.CONTRACTROLESTORAGE(&_AssetboxInfoStorageImpl.CallOpts)
}

// CONTRACTROLESTORAGE is a free data retrieval call binding the contract method 0xb938235b.
//
// Solidity: function CONTRACT_ROLE_STORAGE() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCallerSession) CONTRACTROLESTORAGE() (*big.Int, error) {
	return _AssetboxInfoStorageImpl.Contract.CONTRACTROLESTORAGE(&_AssetboxInfoStorageImpl.CallOpts)
}

// CONTRACTSAFETRANSFERSTORAGE is a free data retrieval call binding the contract method 0xe7e45392.
//
// Solidity: function CONTRACT_SAFE_TRANSFER_STORAGE() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCaller) CONTRACTSAFETRANSFERSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxInfoStorageImpl.contract.Call(opts, out, "CONTRACT_SAFE_TRANSFER_STORAGE")
	return *ret0, err
}

// CONTRACTSAFETRANSFERSTORAGE is a free data retrieval call binding the contract method 0xe7e45392.
//
// Solidity: function CONTRACT_SAFE_TRANSFER_STORAGE() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplSession) CONTRACTSAFETRANSFERSTORAGE() (*big.Int, error) {
	return _AssetboxInfoStorageImpl.Contract.CONTRACTSAFETRANSFERSTORAGE(&_AssetboxInfoStorageImpl.CallOpts)
}

// CONTRACTSAFETRANSFERSTORAGE is a free data retrieval call binding the contract method 0xe7e45392.
//
// Solidity: function CONTRACT_SAFE_TRANSFER_STORAGE() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCallerSession) CONTRACTSAFETRANSFERSTORAGE() (*big.Int, error) {
	return _AssetboxInfoStorageImpl.Contract.CONTRACTSAFETRANSFERSTORAGE(&_AssetboxInfoStorageImpl.CallOpts)
}

// CONTRACTSTORAGE is a free data retrieval call binding the contract method 0x3b2f595f.
//
// Solidity: function CONTRACT_STORAGE() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCaller) CONTRACTSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxInfoStorageImpl.contract.Call(opts, out, "CONTRACT_STORAGE")
	return *ret0, err
}

// CONTRACTSTORAGE is a free data retrieval call binding the contract method 0x3b2f595f.
//
// Solidity: function CONTRACT_STORAGE() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplSession) CONTRACTSTORAGE() (*big.Int, error) {
	return _AssetboxInfoStorageImpl.Contract.CONTRACTSTORAGE(&_AssetboxInfoStorageImpl.CallOpts)
}

// CONTRACTSTORAGE is a free data retrieval call binding the contract method 0x3b2f595f.
//
// Solidity: function CONTRACT_STORAGE() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCallerSession) CONTRACTSTORAGE() (*big.Int, error) {
	return _AssetboxInfoStorageImpl.Contract.CONTRACTSTORAGE(&_AssetboxInfoStorageImpl.CallOpts)
}

// CONTRACTTRANSACTIONSTORAGE is a free data retrieval call binding the contract method 0x327279dd.
//
// Solidity: function CONTRACT_TRANSACTION_STORAGE() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCaller) CONTRACTTRANSACTIONSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxInfoStorageImpl.contract.Call(opts, out, "CONTRACT_TRANSACTION_STORAGE")
	return *ret0, err
}

// CONTRACTTRANSACTIONSTORAGE is a free data retrieval call binding the contract method 0x327279dd.
//
// Solidity: function CONTRACT_TRANSACTION_STORAGE() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplSession) CONTRACTTRANSACTIONSTORAGE() (*big.Int, error) {
	return _AssetboxInfoStorageImpl.Contract.CONTRACTTRANSACTIONSTORAGE(&_AssetboxInfoStorageImpl.CallOpts)
}

// CONTRACTTRANSACTIONSTORAGE is a free data retrieval call binding the contract method 0x327279dd.
//
// Solidity: function CONTRACT_TRANSACTION_STORAGE() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCallerSession) CONTRACTTRANSACTIONSTORAGE() (*big.Int, error) {
	return _AssetboxInfoStorageImpl.Contract.CONTRACTTRANSACTIONSTORAGE(&_AssetboxInfoStorageImpl.CallOpts)
}

// CONTRACTTRANSFERFROMCAPITALIZATION is a free data retrieval call binding the contract method 0xc8f8decb.
//
// Solidity: function CONTRACT_TRANSFER_FROM_CAPITALIZATION() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCaller) CONTRACTTRANSFERFROMCAPITALIZATION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxInfoStorageImpl.contract.Call(opts, out, "CONTRACT_TRANSFER_FROM_CAPITALIZATION")
	return *ret0, err
}

// CONTRACTTRANSFERFROMCAPITALIZATION is a free data retrieval call binding the contract method 0xc8f8decb.
//
// Solidity: function CONTRACT_TRANSFER_FROM_CAPITALIZATION() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplSession) CONTRACTTRANSFERFROMCAPITALIZATION() (*big.Int, error) {
	return _AssetboxInfoStorageImpl.Contract.CONTRACTTRANSFERFROMCAPITALIZATION(&_AssetboxInfoStorageImpl.CallOpts)
}

// CONTRACTTRANSFERFROMCAPITALIZATION is a free data retrieval call binding the contract method 0xc8f8decb.
//
// Solidity: function CONTRACT_TRANSFER_FROM_CAPITALIZATION() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCallerSession) CONTRACTTRANSFERFROMCAPITALIZATION() (*big.Int, error) {
	return _AssetboxInfoStorageImpl.Contract.CONTRACTTRANSFERFROMCAPITALIZATION(&_AssetboxInfoStorageImpl.CallOpts)
}

// PERMISSIONACCESSRESTORATION is a free data retrieval call binding the contract method 0x5c3053bf.
//
// Solidity: function PERMISSION_ACCESS_RESTORATION() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCaller) PERMISSIONACCESSRESTORATION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxInfoStorageImpl.contract.Call(opts, out, "PERMISSION_ACCESS_RESTORATION")
	return *ret0, err
}

// PERMISSIONACCESSRESTORATION is a free data retrieval call binding the contract method 0x5c3053bf.
//
// Solidity: function PERMISSION_ACCESS_RESTORATION() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplSession) PERMISSIONACCESSRESTORATION() (*big.Int, error) {
	return _AssetboxInfoStorageImpl.Contract.PERMISSIONACCESSRESTORATION(&_AssetboxInfoStorageImpl.CallOpts)
}

// PERMISSIONACCESSRESTORATION is a free data retrieval call binding the contract method 0x5c3053bf.
//
// Solidity: function PERMISSION_ACCESS_RESTORATION() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCallerSession) PERMISSIONACCESSRESTORATION() (*big.Int, error) {
	return _AssetboxInfoStorageImpl.Contract.PERMISSIONACCESSRESTORATION(&_AssetboxInfoStorageImpl.CallOpts)
}

// PERMISSIONADMINSTORAGE is a free data retrieval call binding the contract method 0x0c9fc0db.
//
// Solidity: function PERMISSION_ADMIN_STORAGE() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCaller) PERMISSIONADMINSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxInfoStorageImpl.contract.Call(opts, out, "PERMISSION_ADMIN_STORAGE")
	return *ret0, err
}

// PERMISSIONADMINSTORAGE is a free data retrieval call binding the contract method 0x0c9fc0db.
//
// Solidity: function PERMISSION_ADMIN_STORAGE() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplSession) PERMISSIONADMINSTORAGE() (*big.Int, error) {
	return _AssetboxInfoStorageImpl.Contract.PERMISSIONADMINSTORAGE(&_AssetboxInfoStorageImpl.CallOpts)
}

// PERMISSIONADMINSTORAGE is a free data retrieval call binding the contract method 0x0c9fc0db.
//
// Solidity: function PERMISSION_ADMIN_STORAGE() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCallerSession) PERMISSIONADMINSTORAGE() (*big.Int, error) {
	return _AssetboxInfoStorageImpl.Contract.PERMISSIONADMINSTORAGE(&_AssetboxInfoStorageImpl.CallOpts)
}

// PERMISSIONCONTRACTSSTORAGE is a free data retrieval call binding the contract method 0x1516e83b.
//
// Solidity: function PERMISSION_CONTRACTS_STORAGE() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCaller) PERMISSIONCONTRACTSSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxInfoStorageImpl.contract.Call(opts, out, "PERMISSION_CONTRACTS_STORAGE")
	return *ret0, err
}

// PERMISSIONCONTRACTSSTORAGE is a free data retrieval call binding the contract method 0x1516e83b.
//
// Solidity: function PERMISSION_CONTRACTS_STORAGE() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplSession) PERMISSIONCONTRACTSSTORAGE() (*big.Int, error) {
	return _AssetboxInfoStorageImpl.Contract.PERMISSIONCONTRACTSSTORAGE(&_AssetboxInfoStorageImpl.CallOpts)
}

// PERMISSIONCONTRACTSSTORAGE is a free data retrieval call binding the contract method 0x1516e83b.
//
// Solidity: function PERMISSION_CONTRACTS_STORAGE() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCallerSession) PERMISSIONCONTRACTSSTORAGE() (*big.Int, error) {
	return _AssetboxInfoStorageImpl.Contract.PERMISSIONCONTRACTSSTORAGE(&_AssetboxInfoStorageImpl.CallOpts)
}

// PERMISSIONCONTRACTDEPLOY is a free data retrieval call binding the contract method 0x257a4a04.
//
// Solidity: function PERMISSION_CONTRACT_DEPLOY() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCaller) PERMISSIONCONTRACTDEPLOY(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxInfoStorageImpl.contract.Call(opts, out, "PERMISSION_CONTRACT_DEPLOY")
	return *ret0, err
}

// PERMISSIONCONTRACTDEPLOY is a free data retrieval call binding the contract method 0x257a4a04.
//
// Solidity: function PERMISSION_CONTRACT_DEPLOY() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplSession) PERMISSIONCONTRACTDEPLOY() (*big.Int, error) {
	return _AssetboxInfoStorageImpl.Contract.PERMISSIONCONTRACTDEPLOY(&_AssetboxInfoStorageImpl.CallOpts)
}

// PERMISSIONCONTRACTDEPLOY is a free data retrieval call binding the contract method 0x257a4a04.
//
// Solidity: function PERMISSION_CONTRACT_DEPLOY() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCallerSession) PERMISSIONCONTRACTDEPLOY() (*big.Int, error) {
	return _AssetboxInfoStorageImpl.Contract.PERMISSIONCONTRACTDEPLOY(&_AssetboxInfoStorageImpl.CallOpts)
}

// PERMISSIONEMISSION is a free data retrieval call binding the contract method 0x4ccd92d1.
//
// Solidity: function PERMISSION_EMISSION() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCaller) PERMISSIONEMISSION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxInfoStorageImpl.contract.Call(opts, out, "PERMISSION_EMISSION")
	return *ret0, err
}

// PERMISSIONEMISSION is a free data retrieval call binding the contract method 0x4ccd92d1.
//
// Solidity: function PERMISSION_EMISSION() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplSession) PERMISSIONEMISSION() (*big.Int, error) {
	return _AssetboxInfoStorageImpl.Contract.PERMISSIONEMISSION(&_AssetboxInfoStorageImpl.CallOpts)
}

// PERMISSIONEMISSION is a free data retrieval call binding the contract method 0x4ccd92d1.
//
// Solidity: function PERMISSION_EMISSION() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCallerSession) PERMISSIONEMISSION() (*big.Int, error) {
	return _AssetboxInfoStorageImpl.Contract.PERMISSIONEMISSION(&_AssetboxInfoStorageImpl.CallOpts)
}

// PERMISSIONFEESSTORAGE is a free data retrieval call binding the contract method 0xe837eb87.
//
// Solidity: function PERMISSION_FEES_STORAGE() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCaller) PERMISSIONFEESSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxInfoStorageImpl.contract.Call(opts, out, "PERMISSION_FEES_STORAGE")
	return *ret0, err
}

// PERMISSIONFEESSTORAGE is a free data retrieval call binding the contract method 0xe837eb87.
//
// Solidity: function PERMISSION_FEES_STORAGE() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplSession) PERMISSIONFEESSTORAGE() (*big.Int, error) {
	return _AssetboxInfoStorageImpl.Contract.PERMISSIONFEESSTORAGE(&_AssetboxInfoStorageImpl.CallOpts)
}

// PERMISSIONFEESSTORAGE is a free data retrieval call binding the contract method 0xe837eb87.
//
// Solidity: function PERMISSION_FEES_STORAGE() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCallerSession) PERMISSIONFEESSTORAGE() (*big.Int, error) {
	return _AssetboxInfoStorageImpl.Contract.PERMISSIONFEESSTORAGE(&_AssetboxInfoStorageImpl.CallOpts)
}

// PERMISSIONMULTISIGSTORAGE is a free data retrieval call binding the contract method 0xf1bb7aad.
//
// Solidity: function PERMISSION_MULTISIG_STORAGE() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCaller) PERMISSIONMULTISIGSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxInfoStorageImpl.contract.Call(opts, out, "PERMISSION_MULTISIG_STORAGE")
	return *ret0, err
}

// PERMISSIONMULTISIGSTORAGE is a free data retrieval call binding the contract method 0xf1bb7aad.
//
// Solidity: function PERMISSION_MULTISIG_STORAGE() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplSession) PERMISSIONMULTISIGSTORAGE() (*big.Int, error) {
	return _AssetboxInfoStorageImpl.Contract.PERMISSIONMULTISIGSTORAGE(&_AssetboxInfoStorageImpl.CallOpts)
}

// PERMISSIONMULTISIGSTORAGE is a free data retrieval call binding the contract method 0xf1bb7aad.
//
// Solidity: function PERMISSION_MULTISIG_STORAGE() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCallerSession) PERMISSIONMULTISIGSTORAGE() (*big.Int, error) {
	return _AssetboxInfoStorageImpl.Contract.PERMISSIONMULTISIGSTORAGE(&_AssetboxInfoStorageImpl.CallOpts)
}

// PERMISSIONROLESSTORAGE is a free data retrieval call binding the contract method 0x12a16032.
//
// Solidity: function PERMISSION_ROLES_STORAGE() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCaller) PERMISSIONROLESSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxInfoStorageImpl.contract.Call(opts, out, "PERMISSION_ROLES_STORAGE")
	return *ret0, err
}

// PERMISSIONROLESSTORAGE is a free data retrieval call binding the contract method 0x12a16032.
//
// Solidity: function PERMISSION_ROLES_STORAGE() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplSession) PERMISSIONROLESSTORAGE() (*big.Int, error) {
	return _AssetboxInfoStorageImpl.Contract.PERMISSIONROLESSTORAGE(&_AssetboxInfoStorageImpl.CallOpts)
}

// PERMISSIONROLESSTORAGE is a free data retrieval call binding the contract method 0x12a16032.
//
// Solidity: function PERMISSION_ROLES_STORAGE() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCallerSession) PERMISSIONROLESSTORAGE() (*big.Int, error) {
	return _AssetboxInfoStorageImpl.Contract.PERMISSIONROLESSTORAGE(&_AssetboxInfoStorageImpl.CallOpts)
}

// ROLEACCESSVERIFIER is a free data retrieval call binding the contract method 0xc5405f92.
//
// Solidity: function ROLE_ACCESS_VERIFIER() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCaller) ROLEACCESSVERIFIER(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxInfoStorageImpl.contract.Call(opts, out, "ROLE_ACCESS_VERIFIER")
	return *ret0, err
}

// ROLEACCESSVERIFIER is a free data retrieval call binding the contract method 0xc5405f92.
//
// Solidity: function ROLE_ACCESS_VERIFIER() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplSession) ROLEACCESSVERIFIER() (*big.Int, error) {
	return _AssetboxInfoStorageImpl.Contract.ROLEACCESSVERIFIER(&_AssetboxInfoStorageImpl.CallOpts)
}

// ROLEACCESSVERIFIER is a free data retrieval call binding the contract method 0xc5405f92.
//
// Solidity: function ROLE_ACCESS_VERIFIER() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCallerSession) ROLEACCESSVERIFIER() (*big.Int, error) {
	return _AssetboxInfoStorageImpl.Contract.ROLEACCESSVERIFIER(&_AssetboxInfoStorageImpl.CallOpts)
}

// ROLEBITBONISSUEVERIFIER is a free data retrieval call binding the contract method 0xf0a012a6.
//
// Solidity: function ROLE_BITBON_ISSUE_VERIFIER() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCaller) ROLEBITBONISSUEVERIFIER(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxInfoStorageImpl.contract.Call(opts, out, "ROLE_BITBON_ISSUE_VERIFIER")
	return *ret0, err
}

// ROLEBITBONISSUEVERIFIER is a free data retrieval call binding the contract method 0xf0a012a6.
//
// Solidity: function ROLE_BITBON_ISSUE_VERIFIER() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplSession) ROLEBITBONISSUEVERIFIER() (*big.Int, error) {
	return _AssetboxInfoStorageImpl.Contract.ROLEBITBONISSUEVERIFIER(&_AssetboxInfoStorageImpl.CallOpts)
}

// ROLEBITBONISSUEVERIFIER is a free data retrieval call binding the contract method 0xf0a012a6.
//
// Solidity: function ROLE_BITBON_ISSUE_VERIFIER() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCallerSession) ROLEBITBONISSUEVERIFIER() (*big.Int, error) {
	return _AssetboxInfoStorageImpl.Contract.ROLEBITBONISSUEVERIFIER(&_AssetboxInfoStorageImpl.CallOpts)
}

// ROLECOMMISSIONVERIFIER is a free data retrieval call binding the contract method 0x8ab3de8c.
//
// Solidity: function ROLE_COMMISSION_VERIFIER() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCaller) ROLECOMMISSIONVERIFIER(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxInfoStorageImpl.contract.Call(opts, out, "ROLE_COMMISSION_VERIFIER")
	return *ret0, err
}

// ROLECOMMISSIONVERIFIER is a free data retrieval call binding the contract method 0x8ab3de8c.
//
// Solidity: function ROLE_COMMISSION_VERIFIER() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplSession) ROLECOMMISSIONVERIFIER() (*big.Int, error) {
	return _AssetboxInfoStorageImpl.Contract.ROLECOMMISSIONVERIFIER(&_AssetboxInfoStorageImpl.CallOpts)
}

// ROLECOMMISSIONVERIFIER is a free data retrieval call binding the contract method 0x8ab3de8c.
//
// Solidity: function ROLE_COMMISSION_VERIFIER() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCallerSession) ROLECOMMISSIONVERIFIER() (*big.Int, error) {
	return _AssetboxInfoStorageImpl.Contract.ROLECOMMISSIONVERIFIER(&_AssetboxInfoStorageImpl.CallOpts)
}

// ROLEDEPLOYADMIN is a free data retrieval call binding the contract method 0x667bf6e5.
//
// Solidity: function ROLE_DEPLOY_ADMIN() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCaller) ROLEDEPLOYADMIN(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxInfoStorageImpl.contract.Call(opts, out, "ROLE_DEPLOY_ADMIN")
	return *ret0, err
}

// ROLEDEPLOYADMIN is a free data retrieval call binding the contract method 0x667bf6e5.
//
// Solidity: function ROLE_DEPLOY_ADMIN() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplSession) ROLEDEPLOYADMIN() (*big.Int, error) {
	return _AssetboxInfoStorageImpl.Contract.ROLEDEPLOYADMIN(&_AssetboxInfoStorageImpl.CallOpts)
}

// ROLEDEPLOYADMIN is a free data retrieval call binding the contract method 0x667bf6e5.
//
// Solidity: function ROLE_DEPLOY_ADMIN() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCallerSession) ROLEDEPLOYADMIN() (*big.Int, error) {
	return _AssetboxInfoStorageImpl.Contract.ROLEDEPLOYADMIN(&_AssetboxInfoStorageImpl.CallOpts)
}

// ROLEPERMISSIONADMIN is a free data retrieval call binding the contract method 0x1b239996.
//
// Solidity: function ROLE_PERMISSION_ADMIN() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCaller) ROLEPERMISSIONADMIN(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxInfoStorageImpl.contract.Call(opts, out, "ROLE_PERMISSION_ADMIN")
	return *ret0, err
}

// ROLEPERMISSIONADMIN is a free data retrieval call binding the contract method 0x1b239996.
//
// Solidity: function ROLE_PERMISSION_ADMIN() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplSession) ROLEPERMISSIONADMIN() (*big.Int, error) {
	return _AssetboxInfoStorageImpl.Contract.ROLEPERMISSIONADMIN(&_AssetboxInfoStorageImpl.CallOpts)
}

// ROLEPERMISSIONADMIN is a free data retrieval call binding the contract method 0x1b239996.
//
// Solidity: function ROLE_PERMISSION_ADMIN() view returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCallerSession) ROLEPERMISSIONADMIN() (*big.Int, error) {
	return _AssetboxInfoStorageImpl.Contract.ROLEPERMISSIONADMIN(&_AssetboxInfoStorageImpl.CallOpts)
}

// AliasToAssetbox is a free data retrieval call binding the contract method 0xd08800dc.
//
// Solidity: function aliasToAssetbox(bytes32 ) view returns(address)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCaller) AliasToAssetbox(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AssetboxInfoStorageImpl.contract.Call(opts, out, "aliasToAssetbox", arg0)
	return *ret0, err
}

// AliasToAssetbox is a free data retrieval call binding the contract method 0xd08800dc.
//
// Solidity: function aliasToAssetbox(bytes32 ) view returns(address)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplSession) AliasToAssetbox(arg0 [32]byte) (common.Address, error) {
	return _AssetboxInfoStorageImpl.Contract.AliasToAssetbox(&_AssetboxInfoStorageImpl.CallOpts, arg0)
}

// AliasToAssetbox is a free data retrieval call binding the contract method 0xd08800dc.
//
// Solidity: function aliasToAssetbox(bytes32 ) view returns(address)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCallerSession) AliasToAssetbox(arg0 [32]byte) (common.Address, error) {
	return _AssetboxInfoStorageImpl.Contract.AliasToAssetbox(&_AssetboxInfoStorageImpl.CallOpts, arg0)
}

// AssetboxToAlias is a free data retrieval call binding the contract method 0xcb7807f8.
//
// Solidity: function assetboxToAlias(address ) view returns(string)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCaller) AssetboxToAlias(opts *bind.CallOpts, arg0 common.Address) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _AssetboxInfoStorageImpl.contract.Call(opts, out, "assetboxToAlias", arg0)
	return *ret0, err
}

// AssetboxToAlias is a free data retrieval call binding the contract method 0xcb7807f8.
//
// Solidity: function assetboxToAlias(address ) view returns(string)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplSession) AssetboxToAlias(arg0 common.Address) (string, error) {
	return _AssetboxInfoStorageImpl.Contract.AssetboxToAlias(&_AssetboxInfoStorageImpl.CallOpts, arg0)
}

// AssetboxToAlias is a free data retrieval call binding the contract method 0xcb7807f8.
//
// Solidity: function assetboxToAlias(address ) view returns(string)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCallerSession) AssetboxToAlias(arg0 common.Address) (string, error) {
	return _AssetboxInfoStorageImpl.Contract.AssetboxToAlias(&_AssetboxInfoStorageImpl.CallOpts, arg0)
}

// AssetboxToOtherInfo is a free data retrieval call binding the contract method 0x4a0a23b0.
//
// Solidity: function assetboxToOtherInfo(address ) view returns(bytes)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCaller) AssetboxToOtherInfo(opts *bind.CallOpts, arg0 common.Address) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _AssetboxInfoStorageImpl.contract.Call(opts, out, "assetboxToOtherInfo", arg0)
	return *ret0, err
}

// AssetboxToOtherInfo is a free data retrieval call binding the contract method 0x4a0a23b0.
//
// Solidity: function assetboxToOtherInfo(address ) view returns(bytes)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplSession) AssetboxToOtherInfo(arg0 common.Address) ([]byte, error) {
	return _AssetboxInfoStorageImpl.Contract.AssetboxToOtherInfo(&_AssetboxInfoStorageImpl.CallOpts, arg0)
}

// AssetboxToOtherInfo is a free data retrieval call binding the contract method 0x4a0a23b0.
//
// Solidity: function assetboxToOtherInfo(address ) view returns(bytes)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCallerSession) AssetboxToOtherInfo(arg0 common.Address) ([]byte, error) {
	return _AssetboxInfoStorageImpl.Contract.AssetboxToOtherInfo(&_AssetboxInfoStorageImpl.CallOpts, arg0)
}

// AssetboxToServiceId is a free data retrieval call binding the contract method 0x828f2e27.
//
// Solidity: function assetboxToServiceId(address ) view returns(bytes)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCaller) AssetboxToServiceId(opts *bind.CallOpts, arg0 common.Address) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _AssetboxInfoStorageImpl.contract.Call(opts, out, "assetboxToServiceId", arg0)
	return *ret0, err
}

// AssetboxToServiceId is a free data retrieval call binding the contract method 0x828f2e27.
//
// Solidity: function assetboxToServiceId(address ) view returns(bytes)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplSession) AssetboxToServiceId(arg0 common.Address) ([]byte, error) {
	return _AssetboxInfoStorageImpl.Contract.AssetboxToServiceId(&_AssetboxInfoStorageImpl.CallOpts, arg0)
}

// AssetboxToServiceId is a free data retrieval call binding the contract method 0x828f2e27.
//
// Solidity: function assetboxToServiceId(address ) view returns(bytes)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCallerSession) AssetboxToServiceId(arg0 common.Address) ([]byte, error) {
	return _AssetboxInfoStorageImpl.Contract.AssetboxToServiceId(&_AssetboxInfoStorageImpl.CallOpts, arg0)
}

// ContractStorage is a free data retrieval call binding the contract method 0x0549b27e.
//
// Solidity: function contractStorage() view returns(address)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCaller) ContractStorage(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AssetboxInfoStorageImpl.contract.Call(opts, out, "contractStorage")
	return *ret0, err
}

// ContractStorage is a free data retrieval call binding the contract method 0x0549b27e.
//
// Solidity: function contractStorage() view returns(address)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplSession) ContractStorage() (common.Address, error) {
	return _AssetboxInfoStorageImpl.Contract.ContractStorage(&_AssetboxInfoStorageImpl.CallOpts)
}

// ContractStorage is a free data retrieval call binding the contract method 0x0549b27e.
//
// Solidity: function contractStorage() view returns(address)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCallerSession) ContractStorage() (common.Address, error) {
	return _AssetboxInfoStorageImpl.Contract.ContractStorage(&_AssetboxInfoStorageImpl.CallOpts)
}

// GetAdminStorage is a free data retrieval call binding the contract method 0x2182f050.
//
// Solidity: function getAdminStorage() view returns(address)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCaller) GetAdminStorage(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AssetboxInfoStorageImpl.contract.Call(opts, out, "getAdminStorage")
	return *ret0, err
}

// GetAdminStorage is a free data retrieval call binding the contract method 0x2182f050.
//
// Solidity: function getAdminStorage() view returns(address)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplSession) GetAdminStorage() (common.Address, error) {
	return _AssetboxInfoStorageImpl.Contract.GetAdminStorage(&_AssetboxInfoStorageImpl.CallOpts)
}

// GetAdminStorage is a free data retrieval call binding the contract method 0x2182f050.
//
// Solidity: function getAdminStorage() view returns(address)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCallerSession) GetAdminStorage() (common.Address, error) {
	return _AssetboxInfoStorageImpl.Contract.GetAdminStorage(&_AssetboxInfoStorageImpl.CallOpts)
}

// GetAssetboxInfoByAlias is a free data retrieval call binding the contract method 0xdd97a8fb.
//
// Solidity: function getAssetboxInfoByAlias(string alias) view returns(address assetboxAddress, string assetboxAlias, bytes serviceId, bytes otherInfo)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCaller) GetAssetboxInfoByAlias(opts *bind.CallOpts, alias string) (struct {
	AssetboxAddress common.Address
	AssetboxAlias   string
	ServiceId       []byte
	OtherInfo       []byte
}, error) {
	ret := new(struct {
		AssetboxAddress common.Address
		AssetboxAlias   string
		ServiceId       []byte
		OtherInfo       []byte
	})
	out := ret
	err := _AssetboxInfoStorageImpl.contract.Call(opts, out, "getAssetboxInfoByAlias", alias)
	return *ret, err
}

// GetAssetboxInfoByAlias is a free data retrieval call binding the contract method 0xdd97a8fb.
//
// Solidity: function getAssetboxInfoByAlias(string alias) view returns(address assetboxAddress, string assetboxAlias, bytes serviceId, bytes otherInfo)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplSession) GetAssetboxInfoByAlias(alias string) (struct {
	AssetboxAddress common.Address
	AssetboxAlias   string
	ServiceId       []byte
	OtherInfo       []byte
}, error) {
	return _AssetboxInfoStorageImpl.Contract.GetAssetboxInfoByAlias(&_AssetboxInfoStorageImpl.CallOpts, alias)
}

// GetAssetboxInfoByAlias is a free data retrieval call binding the contract method 0xdd97a8fb.
//
// Solidity: function getAssetboxInfoByAlias(string alias) view returns(address assetboxAddress, string assetboxAlias, bytes serviceId, bytes otherInfo)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCallerSession) GetAssetboxInfoByAlias(alias string) (struct {
	AssetboxAddress common.Address
	AssetboxAlias   string
	ServiceId       []byte
	OtherInfo       []byte
}, error) {
	return _AssetboxInfoStorageImpl.Contract.GetAssetboxInfoByAlias(&_AssetboxInfoStorageImpl.CallOpts, alias)
}

// GetAssetboxInfoByAssetbox is a free data retrieval call binding the contract method 0x839e0fe1.
//
// Solidity: function getAssetboxInfoByAssetbox(address assetbox) view returns(address assetboxAddress, string assetboxAlias, bytes serviceId, bytes otherInfo)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCaller) GetAssetboxInfoByAssetbox(opts *bind.CallOpts, assetbox common.Address) (struct {
	AssetboxAddress common.Address
	AssetboxAlias   string
	ServiceId       []byte
	OtherInfo       []byte
}, error) {
	ret := new(struct {
		AssetboxAddress common.Address
		AssetboxAlias   string
		ServiceId       []byte
		OtherInfo       []byte
	})
	out := ret
	err := _AssetboxInfoStorageImpl.contract.Call(opts, out, "getAssetboxInfoByAssetbox", assetbox)
	return *ret, err
}

// GetAssetboxInfoByAssetbox is a free data retrieval call binding the contract method 0x839e0fe1.
//
// Solidity: function getAssetboxInfoByAssetbox(address assetbox) view returns(address assetboxAddress, string assetboxAlias, bytes serviceId, bytes otherInfo)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplSession) GetAssetboxInfoByAssetbox(assetbox common.Address) (struct {
	AssetboxAddress common.Address
	AssetboxAlias   string
	ServiceId       []byte
	OtherInfo       []byte
}, error) {
	return _AssetboxInfoStorageImpl.Contract.GetAssetboxInfoByAssetbox(&_AssetboxInfoStorageImpl.CallOpts, assetbox)
}

// GetAssetboxInfoByAssetbox is a free data retrieval call binding the contract method 0x839e0fe1.
//
// Solidity: function getAssetboxInfoByAssetbox(address assetbox) view returns(address assetboxAddress, string assetboxAlias, bytes serviceId, bytes otherInfo)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCallerSession) GetAssetboxInfoByAssetbox(assetbox common.Address) (struct {
	AssetboxAddress common.Address
	AssetboxAlias   string
	ServiceId       []byte
	OtherInfo       []byte
}, error) {
	return _AssetboxInfoStorageImpl.Contract.GetAssetboxInfoByAssetbox(&_AssetboxInfoStorageImpl.CallOpts, assetbox)
}

// GetContractStorage is a free data retrieval call binding the contract method 0x8c5eeeca.
//
// Solidity: function getContractStorage() view returns(address)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCaller) GetContractStorage(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AssetboxInfoStorageImpl.contract.Call(opts, out, "getContractStorage")
	return *ret0, err
}

// GetContractStorage is a free data retrieval call binding the contract method 0x8c5eeeca.
//
// Solidity: function getContractStorage() view returns(address)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplSession) GetContractStorage() (common.Address, error) {
	return _AssetboxInfoStorageImpl.Contract.GetContractStorage(&_AssetboxInfoStorageImpl.CallOpts)
}

// GetContractStorage is a free data retrieval call binding the contract method 0x8c5eeeca.
//
// Solidity: function getContractStorage() view returns(address)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCallerSession) GetContractStorage() (common.Address, error) {
	return _AssetboxInfoStorageImpl.Contract.GetContractStorage(&_AssetboxInfoStorageImpl.CallOpts)
}

// GetThisContractIndex is a free data retrieval call binding the contract method 0x70c6d367.
//
// Solidity: function getThisContractIndex() pure returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCaller) GetThisContractIndex(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxInfoStorageImpl.contract.Call(opts, out, "getThisContractIndex")
	return *ret0, err
}

// GetThisContractIndex is a free data retrieval call binding the contract method 0x70c6d367.
//
// Solidity: function getThisContractIndex() pure returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplSession) GetThisContractIndex() (*big.Int, error) {
	return _AssetboxInfoStorageImpl.Contract.GetThisContractIndex(&_AssetboxInfoStorageImpl.CallOpts)
}

// GetThisContractIndex is a free data retrieval call binding the contract method 0x70c6d367.
//
// Solidity: function getThisContractIndex() pure returns(uint256)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCallerSession) GetThisContractIndex() (*big.Int, error) {
	return _AssetboxInfoStorageImpl.Contract.GetThisContractIndex(&_AssetboxInfoStorageImpl.CallOpts)
}

// IsAliasBoundToService is a free data retrieval call binding the contract method 0x29b08235.
//
// Solidity: function isAliasBoundToService(string alias, bytes serviceId) view returns(bool isBound)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCaller) IsAliasBoundToService(opts *bind.CallOpts, alias string, serviceId []byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _AssetboxInfoStorageImpl.contract.Call(opts, out, "isAliasBoundToService", alias, serviceId)
	return *ret0, err
}

// IsAliasBoundToService is a free data retrieval call binding the contract method 0x29b08235.
//
// Solidity: function isAliasBoundToService(string alias, bytes serviceId) view returns(bool isBound)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplSession) IsAliasBoundToService(alias string, serviceId []byte) (bool, error) {
	return _AssetboxInfoStorageImpl.Contract.IsAliasBoundToService(&_AssetboxInfoStorageImpl.CallOpts, alias, serviceId)
}

// IsAliasBoundToService is a free data retrieval call binding the contract method 0x29b08235.
//
// Solidity: function isAliasBoundToService(string alias, bytes serviceId) view returns(bool isBound)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCallerSession) IsAliasBoundToService(alias string, serviceId []byte) (bool, error) {
	return _AssetboxInfoStorageImpl.Contract.IsAliasBoundToService(&_AssetboxInfoStorageImpl.CallOpts, alias, serviceId)
}

// IsAliasTaken is a free data retrieval call binding the contract method 0x7cf7252f.
//
// Solidity: function isAliasTaken(string alias) view returns(bool taken)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCaller) IsAliasTaken(opts *bind.CallOpts, alias string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _AssetboxInfoStorageImpl.contract.Call(opts, out, "isAliasTaken", alias)
	return *ret0, err
}

// IsAliasTaken is a free data retrieval call binding the contract method 0x7cf7252f.
//
// Solidity: function isAliasTaken(string alias) view returns(bool taken)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplSession) IsAliasTaken(alias string) (bool, error) {
	return _AssetboxInfoStorageImpl.Contract.IsAliasTaken(&_AssetboxInfoStorageImpl.CallOpts, alias)
}

// IsAliasTaken is a free data retrieval call binding the contract method 0x7cf7252f.
//
// Solidity: function isAliasTaken(string alias) view returns(bool taken)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCallerSession) IsAliasTaken(alias string) (bool, error) {
	return _AssetboxInfoStorageImpl.Contract.IsAliasTaken(&_AssetboxInfoStorageImpl.CallOpts, alias)
}

// IsAssetboxBoundToService is a free data retrieval call binding the contract method 0x9ae049d6.
//
// Solidity: function isAssetboxBoundToService(address assetbox, bytes serviceId) view returns(bool isBound)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCaller) IsAssetboxBoundToService(opts *bind.CallOpts, assetbox common.Address, serviceId []byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _AssetboxInfoStorageImpl.contract.Call(opts, out, "isAssetboxBoundToService", assetbox, serviceId)
	return *ret0, err
}

// IsAssetboxBoundToService is a free data retrieval call binding the contract method 0x9ae049d6.
//
// Solidity: function isAssetboxBoundToService(address assetbox, bytes serviceId) view returns(bool isBound)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplSession) IsAssetboxBoundToService(assetbox common.Address, serviceId []byte) (bool, error) {
	return _AssetboxInfoStorageImpl.Contract.IsAssetboxBoundToService(&_AssetboxInfoStorageImpl.CallOpts, assetbox, serviceId)
}

// IsAssetboxBoundToService is a free data retrieval call binding the contract method 0x9ae049d6.
//
// Solidity: function isAssetboxBoundToService(address assetbox, bytes serviceId) view returns(bool isBound)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCallerSession) IsAssetboxBoundToService(assetbox common.Address, serviceId []byte) (bool, error) {
	return _AssetboxInfoStorageImpl.Contract.IsAssetboxBoundToService(&_AssetboxInfoStorageImpl.CallOpts, assetbox, serviceId)
}

// IsSanitizedAliasTaken is a free data retrieval call binding the contract method 0xc128d587.
//
// Solidity: function isSanitizedAliasTaken(bytes32 alias) view returns(bool taken)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCaller) IsSanitizedAliasTaken(opts *bind.CallOpts, alias [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _AssetboxInfoStorageImpl.contract.Call(opts, out, "isSanitizedAliasTaken", alias)
	return *ret0, err
}

// IsSanitizedAliasTaken is a free data retrieval call binding the contract method 0xc128d587.
//
// Solidity: function isSanitizedAliasTaken(bytes32 alias) view returns(bool taken)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplSession) IsSanitizedAliasTaken(alias [32]byte) (bool, error) {
	return _AssetboxInfoStorageImpl.Contract.IsSanitizedAliasTaken(&_AssetboxInfoStorageImpl.CallOpts, alias)
}

// IsSanitizedAliasTaken is a free data retrieval call binding the contract method 0xc128d587.
//
// Solidity: function isSanitizedAliasTaken(bytes32 alias) view returns(bool taken)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplCallerSession) IsSanitizedAliasTaken(alias [32]byte) (bool, error) {
	return _AssetboxInfoStorageImpl.Contract.IsSanitizedAliasTaken(&_AssetboxInfoStorageImpl.CallOpts, alias)
}

// DeleteAssetboxInfo is a paid mutator transaction binding the contract method 0xf843b90d.
//
// Solidity: function deleteAssetboxInfo(address assetbox) returns(bool)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplTransactor) DeleteAssetboxInfo(opts *bind.TransactOpts, assetbox common.Address) (*types.Transaction, error) {
	return _AssetboxInfoStorageImpl.contract.Transact(opts, "deleteAssetboxInfo", assetbox)
}

// DeleteAssetboxInfo is a paid mutator transaction binding the contract method 0xf843b90d.
//
// Solidity: function deleteAssetboxInfo(address assetbox) returns(bool)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplSession) DeleteAssetboxInfo(assetbox common.Address) (*types.Transaction, error) {
	return _AssetboxInfoStorageImpl.Contract.DeleteAssetboxInfo(&_AssetboxInfoStorageImpl.TransactOpts, assetbox)
}

// DeleteAssetboxInfo is a paid mutator transaction binding the contract method 0xf843b90d.
//
// Solidity: function deleteAssetboxInfo(address assetbox) returns(bool)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplTransactorSession) DeleteAssetboxInfo(assetbox common.Address) (*types.Transaction, error) {
	return _AssetboxInfoStorageImpl.Contract.DeleteAssetboxInfo(&_AssetboxInfoStorageImpl.TransactOpts, assetbox)
}

// SetAssetboxInfo is a paid mutator transaction binding the contract method 0x2d9b8f21.
//
// Solidity: function setAssetboxInfo(address assetbox, string alias, bytes serviceId, bytes otherInfo) returns(bool)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplTransactor) SetAssetboxInfo(opts *bind.TransactOpts, assetbox common.Address, alias string, serviceId []byte, otherInfo []byte) (*types.Transaction, error) {
	return _AssetboxInfoStorageImpl.contract.Transact(opts, "setAssetboxInfo", assetbox, alias, serviceId, otherInfo)
}

// SetAssetboxInfo is a paid mutator transaction binding the contract method 0x2d9b8f21.
//
// Solidity: function setAssetboxInfo(address assetbox, string alias, bytes serviceId, bytes otherInfo) returns(bool)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplSession) SetAssetboxInfo(assetbox common.Address, alias string, serviceId []byte, otherInfo []byte) (*types.Transaction, error) {
	return _AssetboxInfoStorageImpl.Contract.SetAssetboxInfo(&_AssetboxInfoStorageImpl.TransactOpts, assetbox, alias, serviceId, otherInfo)
}

// SetAssetboxInfo is a paid mutator transaction binding the contract method 0x2d9b8f21.
//
// Solidity: function setAssetboxInfo(address assetbox, string alias, bytes serviceId, bytes otherInfo) returns(bool)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplTransactorSession) SetAssetboxInfo(assetbox common.Address, alias string, serviceId []byte, otherInfo []byte) (*types.Transaction, error) {
	return _AssetboxInfoStorageImpl.Contract.SetAssetboxInfo(&_AssetboxInfoStorageImpl.TransactOpts, assetbox, alias, serviceId, otherInfo)
}

// SetContractStorage is a paid mutator transaction binding the contract method 0x824de02f.
//
// Solidity: function setContractStorage(address contractStorageAddress) returns()
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplTransactor) SetContractStorage(opts *bind.TransactOpts, contractStorageAddress common.Address) (*types.Transaction, error) {
	return _AssetboxInfoStorageImpl.contract.Transact(opts, "setContractStorage", contractStorageAddress)
}

// SetContractStorage is a paid mutator transaction binding the contract method 0x824de02f.
//
// Solidity: function setContractStorage(address contractStorageAddress) returns()
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplSession) SetContractStorage(contractStorageAddress common.Address) (*types.Transaction, error) {
	return _AssetboxInfoStorageImpl.Contract.SetContractStorage(&_AssetboxInfoStorageImpl.TransactOpts, contractStorageAddress)
}

// SetContractStorage is a paid mutator transaction binding the contract method 0x824de02f.
//
// Solidity: function setContractStorage(address contractStorageAddress) returns()
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplTransactorSession) SetContractStorage(contractStorageAddress common.Address) (*types.Transaction, error) {
	return _AssetboxInfoStorageImpl.Contract.SetContractStorage(&_AssetboxInfoStorageImpl.TransactOpts, contractStorageAddress)
}

// AssetboxInfoStorageImplAssetboxInfoDeletedIterator is returned from FilterAssetboxInfoDeleted and is used to iterate over the raw logs and unpacked data for AssetboxInfoDeleted events raised by the AssetboxInfoStorageImpl contract.
type AssetboxInfoStorageImplAssetboxInfoDeletedIterator struct {
	Event *AssetboxInfoStorageImplAssetboxInfoDeleted // Event containing the contract specifics and raw log

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
func (it *AssetboxInfoStorageImplAssetboxInfoDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetboxInfoStorageImplAssetboxInfoDeleted)
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
		it.Event = new(AssetboxInfoStorageImplAssetboxInfoDeleted)
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
func (it *AssetboxInfoStorageImplAssetboxInfoDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetboxInfoStorageImplAssetboxInfoDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetboxInfoStorageImplAssetboxInfoDeleted represents a AssetboxInfoDeleted event raised by the AssetboxInfoStorageImpl contract.
type AssetboxInfoStorageImplAssetboxInfoDeleted struct {
	Assetbox common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAssetboxInfoDeleted is a free log retrieval operation binding the contract event 0xec7164478831d34ae4812bf43e7b99973007063374e7f6a44ddc1b97f155ceb3.
//
// Solidity: event AssetboxInfoDeleted(address assetbox)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplFilterer) FilterAssetboxInfoDeleted(opts *bind.FilterOpts) (*AssetboxInfoStorageImplAssetboxInfoDeletedIterator, error) {

	logs, sub, err := _AssetboxInfoStorageImpl.contract.FilterLogs(opts, "AssetboxInfoDeleted")
	if err != nil {
		return nil, err
	}
	return &AssetboxInfoStorageImplAssetboxInfoDeletedIterator{contract: _AssetboxInfoStorageImpl.contract, event: "AssetboxInfoDeleted", logs: logs, sub: sub}, nil
}

// WatchAssetboxInfoDeleted is a free log subscription operation binding the contract event 0xec7164478831d34ae4812bf43e7b99973007063374e7f6a44ddc1b97f155ceb3.
//
// Solidity: event AssetboxInfoDeleted(address assetbox)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplFilterer) WatchAssetboxInfoDeleted(opts *bind.WatchOpts, sink chan<- *AssetboxInfoStorageImplAssetboxInfoDeleted) (event.Subscription, error) {

	logs, sub, err := _AssetboxInfoStorageImpl.contract.WatchLogs(opts, "AssetboxInfoDeleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetboxInfoStorageImplAssetboxInfoDeleted)
				if err := _AssetboxInfoStorageImpl.contract.UnpackLog(event, "AssetboxInfoDeleted", log); err != nil {
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

// ParseAssetboxInfoDeleted is a log parse operation binding the contract event 0xec7164478831d34ae4812bf43e7b99973007063374e7f6a44ddc1b97f155ceb3.
//
// Solidity: event AssetboxInfoDeleted(address assetbox)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplFilterer) ParseAssetboxInfoDeleted(log types.Log) (*AssetboxInfoStorageImplAssetboxInfoDeleted, error) {
	event := new(AssetboxInfoStorageImplAssetboxInfoDeleted)
	if err := _AssetboxInfoStorageImpl.contract.UnpackLog(event, "AssetboxInfoDeleted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AssetboxInfoStorageImplAssetboxInfoSetIterator is returned from FilterAssetboxInfoSet and is used to iterate over the raw logs and unpacked data for AssetboxInfoSet events raised by the AssetboxInfoStorageImpl contract.
type AssetboxInfoStorageImplAssetboxInfoSetIterator struct {
	Event *AssetboxInfoStorageImplAssetboxInfoSet // Event containing the contract specifics and raw log

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
func (it *AssetboxInfoStorageImplAssetboxInfoSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetboxInfoStorageImplAssetboxInfoSet)
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
		it.Event = new(AssetboxInfoStorageImplAssetboxInfoSet)
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
func (it *AssetboxInfoStorageImplAssetboxInfoSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetboxInfoStorageImplAssetboxInfoSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetboxInfoStorageImplAssetboxInfoSet represents a AssetboxInfoSet event raised by the AssetboxInfoStorageImpl contract.
type AssetboxInfoStorageImplAssetboxInfoSet struct {
	Assetbox  common.Address
	Alias     string
	ServiceId []byte
	OtherInfo []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAssetboxInfoSet is a free log retrieval operation binding the contract event 0xa102a154ca061e91c5f0db4883539d37f7db6b8c9ae76809bdccb017d1006a74.
//
// Solidity: event AssetboxInfoSet(address assetbox, string alias, bytes serviceId, bytes otherInfo)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplFilterer) FilterAssetboxInfoSet(opts *bind.FilterOpts) (*AssetboxInfoStorageImplAssetboxInfoSetIterator, error) {

	logs, sub, err := _AssetboxInfoStorageImpl.contract.FilterLogs(opts, "AssetboxInfoSet")
	if err != nil {
		return nil, err
	}
	return &AssetboxInfoStorageImplAssetboxInfoSetIterator{contract: _AssetboxInfoStorageImpl.contract, event: "AssetboxInfoSet", logs: logs, sub: sub}, nil
}

// WatchAssetboxInfoSet is a free log subscription operation binding the contract event 0xa102a154ca061e91c5f0db4883539d37f7db6b8c9ae76809bdccb017d1006a74.
//
// Solidity: event AssetboxInfoSet(address assetbox, string alias, bytes serviceId, bytes otherInfo)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplFilterer) WatchAssetboxInfoSet(opts *bind.WatchOpts, sink chan<- *AssetboxInfoStorageImplAssetboxInfoSet) (event.Subscription, error) {

	logs, sub, err := _AssetboxInfoStorageImpl.contract.WatchLogs(opts, "AssetboxInfoSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetboxInfoStorageImplAssetboxInfoSet)
				if err := _AssetboxInfoStorageImpl.contract.UnpackLog(event, "AssetboxInfoSet", log); err != nil {
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

// ParseAssetboxInfoSet is a log parse operation binding the contract event 0xa102a154ca061e91c5f0db4883539d37f7db6b8c9ae76809bdccb017d1006a74.
//
// Solidity: event AssetboxInfoSet(address assetbox, string alias, bytes serviceId, bytes otherInfo)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplFilterer) ParseAssetboxInfoSet(log types.Log) (*AssetboxInfoStorageImplAssetboxInfoSet, error) {
	event := new(AssetboxInfoStorageImplAssetboxInfoSet)
	if err := _AssetboxInfoStorageImpl.contract.UnpackLog(event, "AssetboxInfoSet", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AssetboxInfoStorageImplMulisigCreatedIterator is returned from FilterMulisigCreated and is used to iterate over the raw logs and unpacked data for MulisigCreated events raised by the AssetboxInfoStorageImpl contract.
type AssetboxInfoStorageImplMulisigCreatedIterator struct {
	Event *AssetboxInfoStorageImplMulisigCreated // Event containing the contract specifics and raw log

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
func (it *AssetboxInfoStorageImplMulisigCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetboxInfoStorageImplMulisigCreated)
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
		it.Event = new(AssetboxInfoStorageImplMulisigCreated)
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
func (it *AssetboxInfoStorageImplMulisigCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetboxInfoStorageImplMulisigCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetboxInfoStorageImplMulisigCreated represents a MulisigCreated event raised by the AssetboxInfoStorageImpl contract.
type AssetboxInfoStorageImplMulisigCreated struct {
	ContractIndex *big.Int
	ExpiresOn     *big.Int
	Signer        common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterMulisigCreated is a free log retrieval operation binding the contract event 0xc9aa4fdac693d13cda7102406a3807cd7d382eac9dcb5132318f7bab3ba241d1.
//
// Solidity: event MulisigCreated(uint256 contractIndex, uint256 expiresOn, address signer)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplFilterer) FilterMulisigCreated(opts *bind.FilterOpts) (*AssetboxInfoStorageImplMulisigCreatedIterator, error) {

	logs, sub, err := _AssetboxInfoStorageImpl.contract.FilterLogs(opts, "MulisigCreated")
	if err != nil {
		return nil, err
	}
	return &AssetboxInfoStorageImplMulisigCreatedIterator{contract: _AssetboxInfoStorageImpl.contract, event: "MulisigCreated", logs: logs, sub: sub}, nil
}

// WatchMulisigCreated is a free log subscription operation binding the contract event 0xc9aa4fdac693d13cda7102406a3807cd7d382eac9dcb5132318f7bab3ba241d1.
//
// Solidity: event MulisigCreated(uint256 contractIndex, uint256 expiresOn, address signer)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplFilterer) WatchMulisigCreated(opts *bind.WatchOpts, sink chan<- *AssetboxInfoStorageImplMulisigCreated) (event.Subscription, error) {

	logs, sub, err := _AssetboxInfoStorageImpl.contract.WatchLogs(opts, "MulisigCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetboxInfoStorageImplMulisigCreated)
				if err := _AssetboxInfoStorageImpl.contract.UnpackLog(event, "MulisigCreated", log); err != nil {
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

// ParseMulisigCreated is a log parse operation binding the contract event 0xc9aa4fdac693d13cda7102406a3807cd7d382eac9dcb5132318f7bab3ba241d1.
//
// Solidity: event MulisigCreated(uint256 contractIndex, uint256 expiresOn, address signer)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplFilterer) ParseMulisigCreated(log types.Log) (*AssetboxInfoStorageImplMulisigCreated, error) {
	event := new(AssetboxInfoStorageImplMulisigCreated)
	if err := _AssetboxInfoStorageImpl.contract.UnpackLog(event, "MulisigCreated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AssetboxInfoStorageImplMulisigSignedIterator is returned from FilterMulisigSigned and is used to iterate over the raw logs and unpacked data for MulisigSigned events raised by the AssetboxInfoStorageImpl contract.
type AssetboxInfoStorageImplMulisigSignedIterator struct {
	Event *AssetboxInfoStorageImplMulisigSigned // Event containing the contract specifics and raw log

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
func (it *AssetboxInfoStorageImplMulisigSignedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetboxInfoStorageImplMulisigSigned)
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
		it.Event = new(AssetboxInfoStorageImplMulisigSigned)
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
func (it *AssetboxInfoStorageImplMulisigSignedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetboxInfoStorageImplMulisigSignedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetboxInfoStorageImplMulisigSigned represents a MulisigSigned event raised by the AssetboxInfoStorageImpl contract.
type AssetboxInfoStorageImplMulisigSigned struct {
	ContractIndex *big.Int
	ExpiresOn     *big.Int
	Signer        common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterMulisigSigned is a free log retrieval operation binding the contract event 0x18f36c61b768a61554ed308cc5de15b5730f714f92327240eb1471d8a9498c88.
//
// Solidity: event MulisigSigned(uint256 contractIndex, uint256 expiresOn, address signer)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplFilterer) FilterMulisigSigned(opts *bind.FilterOpts) (*AssetboxInfoStorageImplMulisigSignedIterator, error) {

	logs, sub, err := _AssetboxInfoStorageImpl.contract.FilterLogs(opts, "MulisigSigned")
	if err != nil {
		return nil, err
	}
	return &AssetboxInfoStorageImplMulisigSignedIterator{contract: _AssetboxInfoStorageImpl.contract, event: "MulisigSigned", logs: logs, sub: sub}, nil
}

// WatchMulisigSigned is a free log subscription operation binding the contract event 0x18f36c61b768a61554ed308cc5de15b5730f714f92327240eb1471d8a9498c88.
//
// Solidity: event MulisigSigned(uint256 contractIndex, uint256 expiresOn, address signer)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplFilterer) WatchMulisigSigned(opts *bind.WatchOpts, sink chan<- *AssetboxInfoStorageImplMulisigSigned) (event.Subscription, error) {

	logs, sub, err := _AssetboxInfoStorageImpl.contract.WatchLogs(opts, "MulisigSigned")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetboxInfoStorageImplMulisigSigned)
				if err := _AssetboxInfoStorageImpl.contract.UnpackLog(event, "MulisigSigned", log); err != nil {
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

// ParseMulisigSigned is a log parse operation binding the contract event 0x18f36c61b768a61554ed308cc5de15b5730f714f92327240eb1471d8a9498c88.
//
// Solidity: event MulisigSigned(uint256 contractIndex, uint256 expiresOn, address signer)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplFilterer) ParseMulisigSigned(log types.Log) (*AssetboxInfoStorageImplMulisigSigned, error) {
	event := new(AssetboxInfoStorageImplMulisigSigned)
	if err := _AssetboxInfoStorageImpl.contract.UnpackLog(event, "MulisigSigned", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AssetboxInfoStorageImplMulisigSucceededIterator is returned from FilterMulisigSucceeded and is used to iterate over the raw logs and unpacked data for MulisigSucceeded events raised by the AssetboxInfoStorageImpl contract.
type AssetboxInfoStorageImplMulisigSucceededIterator struct {
	Event *AssetboxInfoStorageImplMulisigSucceeded // Event containing the contract specifics and raw log

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
func (it *AssetboxInfoStorageImplMulisigSucceededIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetboxInfoStorageImplMulisigSucceeded)
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
		it.Event = new(AssetboxInfoStorageImplMulisigSucceeded)
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
func (it *AssetboxInfoStorageImplMulisigSucceededIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetboxInfoStorageImplMulisigSucceededIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetboxInfoStorageImplMulisigSucceeded represents a MulisigSucceeded event raised by the AssetboxInfoStorageImpl contract.
type AssetboxInfoStorageImplMulisigSucceeded struct {
	ContractIndex *big.Int
	ExpiresOn     *big.Int
	Signer        common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterMulisigSucceeded is a free log retrieval operation binding the contract event 0xe505452fbd2b40208939598519b0abcce3dccaf156ddff770c8879423894178a.
//
// Solidity: event MulisigSucceeded(uint256 contractIndex, uint256 expiresOn, address signer)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplFilterer) FilterMulisigSucceeded(opts *bind.FilterOpts) (*AssetboxInfoStorageImplMulisigSucceededIterator, error) {

	logs, sub, err := _AssetboxInfoStorageImpl.contract.FilterLogs(opts, "MulisigSucceeded")
	if err != nil {
		return nil, err
	}
	return &AssetboxInfoStorageImplMulisigSucceededIterator{contract: _AssetboxInfoStorageImpl.contract, event: "MulisigSucceeded", logs: logs, sub: sub}, nil
}

// WatchMulisigSucceeded is a free log subscription operation binding the contract event 0xe505452fbd2b40208939598519b0abcce3dccaf156ddff770c8879423894178a.
//
// Solidity: event MulisigSucceeded(uint256 contractIndex, uint256 expiresOn, address signer)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplFilterer) WatchMulisigSucceeded(opts *bind.WatchOpts, sink chan<- *AssetboxInfoStorageImplMulisigSucceeded) (event.Subscription, error) {

	logs, sub, err := _AssetboxInfoStorageImpl.contract.WatchLogs(opts, "MulisigSucceeded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetboxInfoStorageImplMulisigSucceeded)
				if err := _AssetboxInfoStorageImpl.contract.UnpackLog(event, "MulisigSucceeded", log); err != nil {
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

// ParseMulisigSucceeded is a log parse operation binding the contract event 0xe505452fbd2b40208939598519b0abcce3dccaf156ddff770c8879423894178a.
//
// Solidity: event MulisigSucceeded(uint256 contractIndex, uint256 expiresOn, address signer)
func (_AssetboxInfoStorageImpl *AssetboxInfoStorageImplFilterer) ParseMulisigSucceeded(log types.Log) (*AssetboxInfoStorageImplMulisigSucceeded, error) {
	event := new(AssetboxInfoStorageImplMulisigSucceeded)
	if err := _AssetboxInfoStorageImpl.contract.UnpackLog(event, "MulisigSucceeded", log); err != nil {
		return nil, err
	}
	return event, nil
}
