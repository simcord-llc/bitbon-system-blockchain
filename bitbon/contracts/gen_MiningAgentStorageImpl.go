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

// MiningAgentStorageImplABI is the input ABI used to generate the binding from.
const MiningAgentStorageImplABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractStorageAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"miningAgentAssetbox\",\"type\":\"address\"}],\"name\":\"MiningAgentAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"miningAgentAssetbox\",\"type\":\"address\"}],\"name\":\"MiningAgentRemoved\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ACCESS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ACCESS_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ADMINABLE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ADMIN_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ASSETBOX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ASSETBOX_INFO\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ASSETBOX_INFO_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ASSETBOX_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_BALANCE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_BITBON\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_BITBON_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_BITBON_SUPPORT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_BONBALANCE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_BONTRANSFER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_DEX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_DISTRIBUTION_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_EXCHANGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_EXCHANGE_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_FEE_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_GENERATOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MINING_AGENT_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MSBON\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MSBON_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_ADD_ADDRESS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_ADD_ADMIN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_ADD_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_ADD_ROLE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_DROP_PERMISSION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_EDIT_ADMIN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_EDIT_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_EDIT_ROLE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REMOVE_ADMIN_BY_ADDRESS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REMOVE_ADMIN_BY_INDEX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REMOVE_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REMOVE_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REMOVE_ROLE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REPLENISH_CAPITALIZATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REPLENISH_CERTIFICATE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REPLENISH_EMISSION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_SET_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_SET_PERMISSION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_OTC\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_OTC_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_REPLENISH_CAPITALIZATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_RESERVED_ALIAS_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ROLE_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_SAFE_TRANSFER_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_TRANSFER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_TRANSFER_FROM_CAPITALIZATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_ACCESS_RESTORATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_ADMIN_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_CONTRACTS_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_CONTRACT_DEPLOY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_EMISSION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_FEES_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_MULTISIG_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_ROLES_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_ACCESS_VERIFIER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_BITBON_ISSUE_VERIFIER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_COMMISSION_VERIFIER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_DEPLOY_ADMIN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_PERMISSION_ADMIN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"contractStorage\",\"outputs\":[{\"internalType\":\"contractContractStorage\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"miningAgentIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"miningAgents\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getThisContractIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"assetbox\",\"type\":\"address\"}],\"name\":\"add\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"length\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getAt\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAll\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"assetbox\",\"type\":\"address\"}],\"name\":\"indexOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"assetbox\",\"type\":\"address\"}],\"name\":\"remove\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"removeAt\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"assetbox\",\"type\":\"address\"}],\"name\":\"isMiningAgent\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// MiningAgentStorageImpl is an auto generated Go binding around an Ethereum contract.
type MiningAgentStorageImpl struct {
	MiningAgentStorageImplCaller     // Read-only binding to the contract
	MiningAgentStorageImplTransactor // Write-only binding to the contract
	MiningAgentStorageImplFilterer   // Log filterer for contract events
}

// MiningAgentStorageImplCaller is an auto generated read-only Go binding around an Ethereum contract.
type MiningAgentStorageImplCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MiningAgentStorageImplTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MiningAgentStorageImplTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MiningAgentStorageImplFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MiningAgentStorageImplFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MiningAgentStorageImplSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MiningAgentStorageImplSession struct {
	Contract     *MiningAgentStorageImpl // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// MiningAgentStorageImplCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MiningAgentStorageImplCallerSession struct {
	Contract *MiningAgentStorageImplCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// MiningAgentStorageImplTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MiningAgentStorageImplTransactorSession struct {
	Contract     *MiningAgentStorageImplTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// MiningAgentStorageImplRaw is an auto generated low-level Go binding around an Ethereum contract.
type MiningAgentStorageImplRaw struct {
	Contract *MiningAgentStorageImpl // Generic contract binding to access the raw methods on
}

// MiningAgentStorageImplCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MiningAgentStorageImplCallerRaw struct {
	Contract *MiningAgentStorageImplCaller // Generic read-only contract binding to access the raw methods on
}

// MiningAgentStorageImplTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MiningAgentStorageImplTransactorRaw struct {
	Contract *MiningAgentStorageImplTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMiningAgentStorageImpl creates a new instance of MiningAgentStorageImpl, bound to a specific deployed contract.
func NewMiningAgentStorageImpl(address common.Address, backend bind.ContractBackend) (*MiningAgentStorageImpl, error) {
	contract, err := bindMiningAgentStorageImpl(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MiningAgentStorageImpl{MiningAgentStorageImplCaller: MiningAgentStorageImplCaller{contract: contract}, MiningAgentStorageImplTransactor: MiningAgentStorageImplTransactor{contract: contract}, MiningAgentStorageImplFilterer: MiningAgentStorageImplFilterer{contract: contract}}, nil
}

// NewMiningAgentStorageImplCaller creates a new read-only instance of MiningAgentStorageImpl, bound to a specific deployed contract.
func NewMiningAgentStorageImplCaller(address common.Address, caller bind.ContractCaller) (*MiningAgentStorageImplCaller, error) {
	contract, err := bindMiningAgentStorageImpl(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MiningAgentStorageImplCaller{contract: contract}, nil
}

// NewMiningAgentStorageImplTransactor creates a new write-only instance of MiningAgentStorageImpl, bound to a specific deployed contract.
func NewMiningAgentStorageImplTransactor(address common.Address, transactor bind.ContractTransactor) (*MiningAgentStorageImplTransactor, error) {
	contract, err := bindMiningAgentStorageImpl(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MiningAgentStorageImplTransactor{contract: contract}, nil
}

// NewMiningAgentStorageImplFilterer creates a new log filterer instance of MiningAgentStorageImpl, bound to a specific deployed contract.
func NewMiningAgentStorageImplFilterer(address common.Address, filterer bind.ContractFilterer) (*MiningAgentStorageImplFilterer, error) {
	contract, err := bindMiningAgentStorageImpl(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MiningAgentStorageImplFilterer{contract: contract}, nil
}

// bindMiningAgentStorageImpl binds a generic wrapper to an already deployed contract.
func bindMiningAgentStorageImpl(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MiningAgentStorageImplABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MiningAgentStorageImpl *MiningAgentStorageImplRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MiningAgentStorageImpl.Contract.MiningAgentStorageImplCaller.contract.Call(opts, result, method, params...)
}

// Return Abi for MiningAgentStorageImplCaller
func (_MiningAgentStorageImpl *MiningAgentStorageImplCaller) ABI() abi.ABI {
	return _MiningAgentStorageImpl.contract.Abi
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MiningAgentStorageImpl *MiningAgentStorageImplRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MiningAgentStorageImpl.Contract.MiningAgentStorageImplTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MiningAgentStorageImpl *MiningAgentStorageImplRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MiningAgentStorageImpl.Contract.MiningAgentStorageImplTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MiningAgentStorageImpl *MiningAgentStorageImplCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MiningAgentStorageImpl.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MiningAgentStorageImpl *MiningAgentStorageImplTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MiningAgentStorageImpl.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MiningAgentStorageImpl *MiningAgentStorageImplTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MiningAgentStorageImpl.Contract.contract.Transact(opts, method, params...)
}

// CONTRACTACCESS is a free data retrieval call binding the contract method 0x0ac62b0c.
//
// Solidity: function CONTRACT_ACCESS() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCaller) CONTRACTACCESS(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MiningAgentStorageImpl.contract.Call(opts, out, "CONTRACT_ACCESS")
	return *ret0, err
}

// CONTRACTACCESS is a free data retrieval call binding the contract method 0x0ac62b0c.
//
// Solidity: function CONTRACT_ACCESS() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplSession) CONTRACTACCESS() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.CONTRACTACCESS(&_MiningAgentStorageImpl.CallOpts)
}

// CONTRACTACCESS is a free data retrieval call binding the contract method 0x0ac62b0c.
//
// Solidity: function CONTRACT_ACCESS() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCallerSession) CONTRACTACCESS() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.CONTRACTACCESS(&_MiningAgentStorageImpl.CallOpts)
}

// CONTRACTACCESSSTORAGE is a free data retrieval call binding the contract method 0xceac9bd5.
//
// Solidity: function CONTRACT_ACCESS_STORAGE() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCaller) CONTRACTACCESSSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MiningAgentStorageImpl.contract.Call(opts, out, "CONTRACT_ACCESS_STORAGE")
	return *ret0, err
}

// CONTRACTACCESSSTORAGE is a free data retrieval call binding the contract method 0xceac9bd5.
//
// Solidity: function CONTRACT_ACCESS_STORAGE() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplSession) CONTRACTACCESSSTORAGE() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.CONTRACTACCESSSTORAGE(&_MiningAgentStorageImpl.CallOpts)
}

// CONTRACTACCESSSTORAGE is a free data retrieval call binding the contract method 0xceac9bd5.
//
// Solidity: function CONTRACT_ACCESS_STORAGE() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCallerSession) CONTRACTACCESSSTORAGE() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.CONTRACTACCESSSTORAGE(&_MiningAgentStorageImpl.CallOpts)
}

// CONTRACTADMINABLE is a free data retrieval call binding the contract method 0xbbe3a78f.
//
// Solidity: function CONTRACT_ADMINABLE() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCaller) CONTRACTADMINABLE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MiningAgentStorageImpl.contract.Call(opts, out, "CONTRACT_ADMINABLE")
	return *ret0, err
}

// CONTRACTADMINABLE is a free data retrieval call binding the contract method 0xbbe3a78f.
//
// Solidity: function CONTRACT_ADMINABLE() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplSession) CONTRACTADMINABLE() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.CONTRACTADMINABLE(&_MiningAgentStorageImpl.CallOpts)
}

// CONTRACTADMINABLE is a free data retrieval call binding the contract method 0xbbe3a78f.
//
// Solidity: function CONTRACT_ADMINABLE() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCallerSession) CONTRACTADMINABLE() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.CONTRACTADMINABLE(&_MiningAgentStorageImpl.CallOpts)
}

// CONTRACTADMINSTORAGE is a free data retrieval call binding the contract method 0xc9aebae9.
//
// Solidity: function CONTRACT_ADMIN_STORAGE() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCaller) CONTRACTADMINSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MiningAgentStorageImpl.contract.Call(opts, out, "CONTRACT_ADMIN_STORAGE")
	return *ret0, err
}

// CONTRACTADMINSTORAGE is a free data retrieval call binding the contract method 0xc9aebae9.
//
// Solidity: function CONTRACT_ADMIN_STORAGE() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplSession) CONTRACTADMINSTORAGE() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.CONTRACTADMINSTORAGE(&_MiningAgentStorageImpl.CallOpts)
}

// CONTRACTADMINSTORAGE is a free data retrieval call binding the contract method 0xc9aebae9.
//
// Solidity: function CONTRACT_ADMIN_STORAGE() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCallerSession) CONTRACTADMINSTORAGE() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.CONTRACTADMINSTORAGE(&_MiningAgentStorageImpl.CallOpts)
}

// CONTRACTASSETBOX is a free data retrieval call binding the contract method 0x161aa17c.
//
// Solidity: function CONTRACT_ASSETBOX() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCaller) CONTRACTASSETBOX(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MiningAgentStorageImpl.contract.Call(opts, out, "CONTRACT_ASSETBOX")
	return *ret0, err
}

// CONTRACTASSETBOX is a free data retrieval call binding the contract method 0x161aa17c.
//
// Solidity: function CONTRACT_ASSETBOX() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplSession) CONTRACTASSETBOX() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.CONTRACTASSETBOX(&_MiningAgentStorageImpl.CallOpts)
}

// CONTRACTASSETBOX is a free data retrieval call binding the contract method 0x161aa17c.
//
// Solidity: function CONTRACT_ASSETBOX() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCallerSession) CONTRACTASSETBOX() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.CONTRACTASSETBOX(&_MiningAgentStorageImpl.CallOpts)
}

// CONTRACTASSETBOXINFO is a free data retrieval call binding the contract method 0xb572fd3d.
//
// Solidity: function CONTRACT_ASSETBOX_INFO() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCaller) CONTRACTASSETBOXINFO(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MiningAgentStorageImpl.contract.Call(opts, out, "CONTRACT_ASSETBOX_INFO")
	return *ret0, err
}

// CONTRACTASSETBOXINFO is a free data retrieval call binding the contract method 0xb572fd3d.
//
// Solidity: function CONTRACT_ASSETBOX_INFO() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplSession) CONTRACTASSETBOXINFO() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.CONTRACTASSETBOXINFO(&_MiningAgentStorageImpl.CallOpts)
}

// CONTRACTASSETBOXINFO is a free data retrieval call binding the contract method 0xb572fd3d.
//
// Solidity: function CONTRACT_ASSETBOX_INFO() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCallerSession) CONTRACTASSETBOXINFO() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.CONTRACTASSETBOXINFO(&_MiningAgentStorageImpl.CallOpts)
}

// CONTRACTASSETBOXINFOSTORAGE is a free data retrieval call binding the contract method 0xf32a115d.
//
// Solidity: function CONTRACT_ASSETBOX_INFO_STORAGE() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCaller) CONTRACTASSETBOXINFOSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MiningAgentStorageImpl.contract.Call(opts, out, "CONTRACT_ASSETBOX_INFO_STORAGE")
	return *ret0, err
}

// CONTRACTASSETBOXINFOSTORAGE is a free data retrieval call binding the contract method 0xf32a115d.
//
// Solidity: function CONTRACT_ASSETBOX_INFO_STORAGE() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplSession) CONTRACTASSETBOXINFOSTORAGE() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.CONTRACTASSETBOXINFOSTORAGE(&_MiningAgentStorageImpl.CallOpts)
}

// CONTRACTASSETBOXINFOSTORAGE is a free data retrieval call binding the contract method 0xf32a115d.
//
// Solidity: function CONTRACT_ASSETBOX_INFO_STORAGE() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCallerSession) CONTRACTASSETBOXINFOSTORAGE() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.CONTRACTASSETBOXINFOSTORAGE(&_MiningAgentStorageImpl.CallOpts)
}

// CONTRACTASSETBOXSTORAGE is a free data retrieval call binding the contract method 0xfa5bdafc.
//
// Solidity: function CONTRACT_ASSETBOX_STORAGE() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCaller) CONTRACTASSETBOXSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MiningAgentStorageImpl.contract.Call(opts, out, "CONTRACT_ASSETBOX_STORAGE")
	return *ret0, err
}

// CONTRACTASSETBOXSTORAGE is a free data retrieval call binding the contract method 0xfa5bdafc.
//
// Solidity: function CONTRACT_ASSETBOX_STORAGE() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplSession) CONTRACTASSETBOXSTORAGE() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.CONTRACTASSETBOXSTORAGE(&_MiningAgentStorageImpl.CallOpts)
}

// CONTRACTASSETBOXSTORAGE is a free data retrieval call binding the contract method 0xfa5bdafc.
//
// Solidity: function CONTRACT_ASSETBOX_STORAGE() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCallerSession) CONTRACTASSETBOXSTORAGE() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.CONTRACTASSETBOXSTORAGE(&_MiningAgentStorageImpl.CallOpts)
}

// CONTRACTBALANCE is a free data retrieval call binding the contract method 0xd59045f5.
//
// Solidity: function CONTRACT_BALANCE() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCaller) CONTRACTBALANCE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MiningAgentStorageImpl.contract.Call(opts, out, "CONTRACT_BALANCE")
	return *ret0, err
}

// CONTRACTBALANCE is a free data retrieval call binding the contract method 0xd59045f5.
//
// Solidity: function CONTRACT_BALANCE() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplSession) CONTRACTBALANCE() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.CONTRACTBALANCE(&_MiningAgentStorageImpl.CallOpts)
}

// CONTRACTBALANCE is a free data retrieval call binding the contract method 0xd59045f5.
//
// Solidity: function CONTRACT_BALANCE() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCallerSession) CONTRACTBALANCE() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.CONTRACTBALANCE(&_MiningAgentStorageImpl.CallOpts)
}

// CONTRACTBITBON is a free data retrieval call binding the contract method 0xb6b433ab.
//
// Solidity: function CONTRACT_BITBON() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCaller) CONTRACTBITBON(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MiningAgentStorageImpl.contract.Call(opts, out, "CONTRACT_BITBON")
	return *ret0, err
}

// CONTRACTBITBON is a free data retrieval call binding the contract method 0xb6b433ab.
//
// Solidity: function CONTRACT_BITBON() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplSession) CONTRACTBITBON() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.CONTRACTBITBON(&_MiningAgentStorageImpl.CallOpts)
}

// CONTRACTBITBON is a free data retrieval call binding the contract method 0xb6b433ab.
//
// Solidity: function CONTRACT_BITBON() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCallerSession) CONTRACTBITBON() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.CONTRACTBITBON(&_MiningAgentStorageImpl.CallOpts)
}

// CONTRACTBITBONSTORAGE is a free data retrieval call binding the contract method 0xe7f6c30e.
//
// Solidity: function CONTRACT_BITBON_STORAGE() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCaller) CONTRACTBITBONSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MiningAgentStorageImpl.contract.Call(opts, out, "CONTRACT_BITBON_STORAGE")
	return *ret0, err
}

// CONTRACTBITBONSTORAGE is a free data retrieval call binding the contract method 0xe7f6c30e.
//
// Solidity: function CONTRACT_BITBON_STORAGE() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplSession) CONTRACTBITBONSTORAGE() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.CONTRACTBITBONSTORAGE(&_MiningAgentStorageImpl.CallOpts)
}

// CONTRACTBITBONSTORAGE is a free data retrieval call binding the contract method 0xe7f6c30e.
//
// Solidity: function CONTRACT_BITBON_STORAGE() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCallerSession) CONTRACTBITBONSTORAGE() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.CONTRACTBITBONSTORAGE(&_MiningAgentStorageImpl.CallOpts)
}

// CONTRACTBITBONSUPPORT is a free data retrieval call binding the contract method 0x20139c74.
//
// Solidity: function CONTRACT_BITBON_SUPPORT() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCaller) CONTRACTBITBONSUPPORT(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MiningAgentStorageImpl.contract.Call(opts, out, "CONTRACT_BITBON_SUPPORT")
	return *ret0, err
}

// CONTRACTBITBONSUPPORT is a free data retrieval call binding the contract method 0x20139c74.
//
// Solidity: function CONTRACT_BITBON_SUPPORT() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplSession) CONTRACTBITBONSUPPORT() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.CONTRACTBITBONSUPPORT(&_MiningAgentStorageImpl.CallOpts)
}

// CONTRACTBITBONSUPPORT is a free data retrieval call binding the contract method 0x20139c74.
//
// Solidity: function CONTRACT_BITBON_SUPPORT() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCallerSession) CONTRACTBITBONSUPPORT() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.CONTRACTBITBONSUPPORT(&_MiningAgentStorageImpl.CallOpts)
}

// CONTRACTBONBALANCE is a free data retrieval call binding the contract method 0x40e02b7c.
//
// Solidity: function CONTRACT_BONBALANCE() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCaller) CONTRACTBONBALANCE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MiningAgentStorageImpl.contract.Call(opts, out, "CONTRACT_BONBALANCE")
	return *ret0, err
}

// CONTRACTBONBALANCE is a free data retrieval call binding the contract method 0x40e02b7c.
//
// Solidity: function CONTRACT_BONBALANCE() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplSession) CONTRACTBONBALANCE() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.CONTRACTBONBALANCE(&_MiningAgentStorageImpl.CallOpts)
}

// CONTRACTBONBALANCE is a free data retrieval call binding the contract method 0x40e02b7c.
//
// Solidity: function CONTRACT_BONBALANCE() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCallerSession) CONTRACTBONBALANCE() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.CONTRACTBONBALANCE(&_MiningAgentStorageImpl.CallOpts)
}

// CONTRACTBONTRANSFER is a free data retrieval call binding the contract method 0xc18f4a8a.
//
// Solidity: function CONTRACT_BONTRANSFER() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCaller) CONTRACTBONTRANSFER(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MiningAgentStorageImpl.contract.Call(opts, out, "CONTRACT_BONTRANSFER")
	return *ret0, err
}

// CONTRACTBONTRANSFER is a free data retrieval call binding the contract method 0xc18f4a8a.
//
// Solidity: function CONTRACT_BONTRANSFER() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplSession) CONTRACTBONTRANSFER() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.CONTRACTBONTRANSFER(&_MiningAgentStorageImpl.CallOpts)
}

// CONTRACTBONTRANSFER is a free data retrieval call binding the contract method 0xc18f4a8a.
//
// Solidity: function CONTRACT_BONTRANSFER() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCallerSession) CONTRACTBONTRANSFER() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.CONTRACTBONTRANSFER(&_MiningAgentStorageImpl.CallOpts)
}

// CONTRACTDEX is a free data retrieval call binding the contract method 0x52c1f844.
//
// Solidity: function CONTRACT_DEX() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCaller) CONTRACTDEX(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MiningAgentStorageImpl.contract.Call(opts, out, "CONTRACT_DEX")
	return *ret0, err
}

// CONTRACTDEX is a free data retrieval call binding the contract method 0x52c1f844.
//
// Solidity: function CONTRACT_DEX() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplSession) CONTRACTDEX() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.CONTRACTDEX(&_MiningAgentStorageImpl.CallOpts)
}

// CONTRACTDEX is a free data retrieval call binding the contract method 0x52c1f844.
//
// Solidity: function CONTRACT_DEX() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCallerSession) CONTRACTDEX() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.CONTRACTDEX(&_MiningAgentStorageImpl.CallOpts)
}

// CONTRACTDISTRIBUTIONSTORAGE is a free data retrieval call binding the contract method 0x08097327.
//
// Solidity: function CONTRACT_DISTRIBUTION_STORAGE() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCaller) CONTRACTDISTRIBUTIONSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MiningAgentStorageImpl.contract.Call(opts, out, "CONTRACT_DISTRIBUTION_STORAGE")
	return *ret0, err
}

// CONTRACTDISTRIBUTIONSTORAGE is a free data retrieval call binding the contract method 0x08097327.
//
// Solidity: function CONTRACT_DISTRIBUTION_STORAGE() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplSession) CONTRACTDISTRIBUTIONSTORAGE() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.CONTRACTDISTRIBUTIONSTORAGE(&_MiningAgentStorageImpl.CallOpts)
}

// CONTRACTDISTRIBUTIONSTORAGE is a free data retrieval call binding the contract method 0x08097327.
//
// Solidity: function CONTRACT_DISTRIBUTION_STORAGE() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCallerSession) CONTRACTDISTRIBUTIONSTORAGE() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.CONTRACTDISTRIBUTIONSTORAGE(&_MiningAgentStorageImpl.CallOpts)
}

// CONTRACTEXCHANGE is a free data retrieval call binding the contract method 0xa0e3a10e.
//
// Solidity: function CONTRACT_EXCHANGE() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCaller) CONTRACTEXCHANGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MiningAgentStorageImpl.contract.Call(opts, out, "CONTRACT_EXCHANGE")
	return *ret0, err
}

// CONTRACTEXCHANGE is a free data retrieval call binding the contract method 0xa0e3a10e.
//
// Solidity: function CONTRACT_EXCHANGE() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplSession) CONTRACTEXCHANGE() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.CONTRACTEXCHANGE(&_MiningAgentStorageImpl.CallOpts)
}

// CONTRACTEXCHANGE is a free data retrieval call binding the contract method 0xa0e3a10e.
//
// Solidity: function CONTRACT_EXCHANGE() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCallerSession) CONTRACTEXCHANGE() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.CONTRACTEXCHANGE(&_MiningAgentStorageImpl.CallOpts)
}

// CONTRACTEXCHANGESTORAGE is a free data retrieval call binding the contract method 0x001570f2.
//
// Solidity: function CONTRACT_EXCHANGE_STORAGE() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCaller) CONTRACTEXCHANGESTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MiningAgentStorageImpl.contract.Call(opts, out, "CONTRACT_EXCHANGE_STORAGE")
	return *ret0, err
}

// CONTRACTEXCHANGESTORAGE is a free data retrieval call binding the contract method 0x001570f2.
//
// Solidity: function CONTRACT_EXCHANGE_STORAGE() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplSession) CONTRACTEXCHANGESTORAGE() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.CONTRACTEXCHANGESTORAGE(&_MiningAgentStorageImpl.CallOpts)
}

// CONTRACTEXCHANGESTORAGE is a free data retrieval call binding the contract method 0x001570f2.
//
// Solidity: function CONTRACT_EXCHANGE_STORAGE() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCallerSession) CONTRACTEXCHANGESTORAGE() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.CONTRACTEXCHANGESTORAGE(&_MiningAgentStorageImpl.CallOpts)
}

// CONTRACTFEESTORAGE is a free data retrieval call binding the contract method 0xcd5e65a4.
//
// Solidity: function CONTRACT_FEE_STORAGE() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCaller) CONTRACTFEESTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MiningAgentStorageImpl.contract.Call(opts, out, "CONTRACT_FEE_STORAGE")
	return *ret0, err
}

// CONTRACTFEESTORAGE is a free data retrieval call binding the contract method 0xcd5e65a4.
//
// Solidity: function CONTRACT_FEE_STORAGE() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplSession) CONTRACTFEESTORAGE() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.CONTRACTFEESTORAGE(&_MiningAgentStorageImpl.CallOpts)
}

// CONTRACTFEESTORAGE is a free data retrieval call binding the contract method 0xcd5e65a4.
//
// Solidity: function CONTRACT_FEE_STORAGE() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCallerSession) CONTRACTFEESTORAGE() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.CONTRACTFEESTORAGE(&_MiningAgentStorageImpl.CallOpts)
}

// CONTRACTGENERATOR is a free data retrieval call binding the contract method 0xccea2b46.
//
// Solidity: function CONTRACT_GENERATOR() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCaller) CONTRACTGENERATOR(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MiningAgentStorageImpl.contract.Call(opts, out, "CONTRACT_GENERATOR")
	return *ret0, err
}

// CONTRACTGENERATOR is a free data retrieval call binding the contract method 0xccea2b46.
//
// Solidity: function CONTRACT_GENERATOR() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplSession) CONTRACTGENERATOR() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.CONTRACTGENERATOR(&_MiningAgentStorageImpl.CallOpts)
}

// CONTRACTGENERATOR is a free data retrieval call binding the contract method 0xccea2b46.
//
// Solidity: function CONTRACT_GENERATOR() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCallerSession) CONTRACTGENERATOR() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.CONTRACTGENERATOR(&_MiningAgentStorageImpl.CallOpts)
}

// CONTRACTMININGAGENTSTORAGE is a free data retrieval call binding the contract method 0x83cc588d.
//
// Solidity: function CONTRACT_MINING_AGENT_STORAGE() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCaller) CONTRACTMININGAGENTSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MiningAgentStorageImpl.contract.Call(opts, out, "CONTRACT_MINING_AGENT_STORAGE")
	return *ret0, err
}

// CONTRACTMININGAGENTSTORAGE is a free data retrieval call binding the contract method 0x83cc588d.
//
// Solidity: function CONTRACT_MINING_AGENT_STORAGE() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplSession) CONTRACTMININGAGENTSTORAGE() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.CONTRACTMININGAGENTSTORAGE(&_MiningAgentStorageImpl.CallOpts)
}

// CONTRACTMININGAGENTSTORAGE is a free data retrieval call binding the contract method 0x83cc588d.
//
// Solidity: function CONTRACT_MINING_AGENT_STORAGE() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCallerSession) CONTRACTMININGAGENTSTORAGE() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.CONTRACTMININGAGENTSTORAGE(&_MiningAgentStorageImpl.CallOpts)
}

// CONTRACTMSBON is a free data retrieval call binding the contract method 0x28b3a976.
//
// Solidity: function CONTRACT_MSBON() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCaller) CONTRACTMSBON(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MiningAgentStorageImpl.contract.Call(opts, out, "CONTRACT_MSBON")
	return *ret0, err
}

// CONTRACTMSBON is a free data retrieval call binding the contract method 0x28b3a976.
//
// Solidity: function CONTRACT_MSBON() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplSession) CONTRACTMSBON() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.CONTRACTMSBON(&_MiningAgentStorageImpl.CallOpts)
}

// CONTRACTMSBON is a free data retrieval call binding the contract method 0x28b3a976.
//
// Solidity: function CONTRACT_MSBON() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCallerSession) CONTRACTMSBON() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.CONTRACTMSBON(&_MiningAgentStorageImpl.CallOpts)
}

// CONTRACTMSBONSTORAGE is a free data retrieval call binding the contract method 0xe04e6061.
//
// Solidity: function CONTRACT_MSBON_STORAGE() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCaller) CONTRACTMSBONSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MiningAgentStorageImpl.contract.Call(opts, out, "CONTRACT_MSBON_STORAGE")
	return *ret0, err
}

// CONTRACTMSBONSTORAGE is a free data retrieval call binding the contract method 0xe04e6061.
//
// Solidity: function CONTRACT_MSBON_STORAGE() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplSession) CONTRACTMSBONSTORAGE() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.CONTRACTMSBONSTORAGE(&_MiningAgentStorageImpl.CallOpts)
}

// CONTRACTMSBONSTORAGE is a free data retrieval call binding the contract method 0xe04e6061.
//
// Solidity: function CONTRACT_MSBON_STORAGE() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCallerSession) CONTRACTMSBONSTORAGE() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.CONTRACTMSBONSTORAGE(&_MiningAgentStorageImpl.CallOpts)
}

// CONTRACTMULTISIGADDADDRESS is a free data retrieval call binding the contract method 0xf4f65a5a.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ADDRESS() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCaller) CONTRACTMULTISIGADDADDRESS(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MiningAgentStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_ADD_ADDRESS")
	return *ret0, err
}

// CONTRACTMULTISIGADDADDRESS is a free data retrieval call binding the contract method 0xf4f65a5a.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ADDRESS() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplSession) CONTRACTMULTISIGADDADDRESS() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.CONTRACTMULTISIGADDADDRESS(&_MiningAgentStorageImpl.CallOpts)
}

// CONTRACTMULTISIGADDADDRESS is a free data retrieval call binding the contract method 0xf4f65a5a.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ADDRESS() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCallerSession) CONTRACTMULTISIGADDADDRESS() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.CONTRACTMULTISIGADDADDRESS(&_MiningAgentStorageImpl.CallOpts)
}

// CONTRACTMULTISIGADDADMIN is a free data retrieval call binding the contract method 0x83888c68.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ADMIN() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCaller) CONTRACTMULTISIGADDADMIN(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MiningAgentStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_ADD_ADMIN")
	return *ret0, err
}

// CONTRACTMULTISIGADDADMIN is a free data retrieval call binding the contract method 0x83888c68.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ADMIN() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplSession) CONTRACTMULTISIGADDADMIN() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.CONTRACTMULTISIGADDADMIN(&_MiningAgentStorageImpl.CallOpts)
}

// CONTRACTMULTISIGADDADMIN is a free data retrieval call binding the contract method 0x83888c68.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ADMIN() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCallerSession) CONTRACTMULTISIGADDADMIN() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.CONTRACTMULTISIGADDADMIN(&_MiningAgentStorageImpl.CallOpts)
}

// CONTRACTMULTISIGADDFEE is a free data retrieval call binding the contract method 0xc43e8cab.
//
// Solidity: function CONTRACT_MULTISIG_ADD_FEE() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCaller) CONTRACTMULTISIGADDFEE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MiningAgentStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_ADD_FEE")
	return *ret0, err
}

// CONTRACTMULTISIGADDFEE is a free data retrieval call binding the contract method 0xc43e8cab.
//
// Solidity: function CONTRACT_MULTISIG_ADD_FEE() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplSession) CONTRACTMULTISIGADDFEE() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.CONTRACTMULTISIGADDFEE(&_MiningAgentStorageImpl.CallOpts)
}

// CONTRACTMULTISIGADDFEE is a free data retrieval call binding the contract method 0xc43e8cab.
//
// Solidity: function CONTRACT_MULTISIG_ADD_FEE() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCallerSession) CONTRACTMULTISIGADDFEE() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.CONTRACTMULTISIGADDFEE(&_MiningAgentStorageImpl.CallOpts)
}

// CONTRACTMULTISIGADDROLE is a free data retrieval call binding the contract method 0x20a0914b.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ROLE() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCaller) CONTRACTMULTISIGADDROLE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MiningAgentStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_ADD_ROLE")
	return *ret0, err
}

// CONTRACTMULTISIGADDROLE is a free data retrieval call binding the contract method 0x20a0914b.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ROLE() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplSession) CONTRACTMULTISIGADDROLE() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.CONTRACTMULTISIGADDROLE(&_MiningAgentStorageImpl.CallOpts)
}

// CONTRACTMULTISIGADDROLE is a free data retrieval call binding the contract method 0x20a0914b.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ROLE() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCallerSession) CONTRACTMULTISIGADDROLE() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.CONTRACTMULTISIGADDROLE(&_MiningAgentStorageImpl.CallOpts)
}

// CONTRACTMULTISIGDROPPERMISSION is a free data retrieval call binding the contract method 0x3f3fbc39.
//
// Solidity: function CONTRACT_MULTISIG_DROP_PERMISSION() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCaller) CONTRACTMULTISIGDROPPERMISSION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MiningAgentStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_DROP_PERMISSION")
	return *ret0, err
}

// CONTRACTMULTISIGDROPPERMISSION is a free data retrieval call binding the contract method 0x3f3fbc39.
//
// Solidity: function CONTRACT_MULTISIG_DROP_PERMISSION() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplSession) CONTRACTMULTISIGDROPPERMISSION() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.CONTRACTMULTISIGDROPPERMISSION(&_MiningAgentStorageImpl.CallOpts)
}

// CONTRACTMULTISIGDROPPERMISSION is a free data retrieval call binding the contract method 0x3f3fbc39.
//
// Solidity: function CONTRACT_MULTISIG_DROP_PERMISSION() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCallerSession) CONTRACTMULTISIGDROPPERMISSION() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.CONTRACTMULTISIGDROPPERMISSION(&_MiningAgentStorageImpl.CallOpts)
}

// CONTRACTMULTISIGEDITADMIN is a free data retrieval call binding the contract method 0x91b0f051.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_ADMIN() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCaller) CONTRACTMULTISIGEDITADMIN(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MiningAgentStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_EDIT_ADMIN")
	return *ret0, err
}

// CONTRACTMULTISIGEDITADMIN is a free data retrieval call binding the contract method 0x91b0f051.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_ADMIN() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplSession) CONTRACTMULTISIGEDITADMIN() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.CONTRACTMULTISIGEDITADMIN(&_MiningAgentStorageImpl.CallOpts)
}

// CONTRACTMULTISIGEDITADMIN is a free data retrieval call binding the contract method 0x91b0f051.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_ADMIN() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCallerSession) CONTRACTMULTISIGEDITADMIN() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.CONTRACTMULTISIGEDITADMIN(&_MiningAgentStorageImpl.CallOpts)
}

// CONTRACTMULTISIGEDITFEE is a free data retrieval call binding the contract method 0x6b2abe17.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_FEE() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCaller) CONTRACTMULTISIGEDITFEE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MiningAgentStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_EDIT_FEE")
	return *ret0, err
}

// CONTRACTMULTISIGEDITFEE is a free data retrieval call binding the contract method 0x6b2abe17.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_FEE() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplSession) CONTRACTMULTISIGEDITFEE() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.CONTRACTMULTISIGEDITFEE(&_MiningAgentStorageImpl.CallOpts)
}

// CONTRACTMULTISIGEDITFEE is a free data retrieval call binding the contract method 0x6b2abe17.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_FEE() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCallerSession) CONTRACTMULTISIGEDITFEE() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.CONTRACTMULTISIGEDITFEE(&_MiningAgentStorageImpl.CallOpts)
}

// CONTRACTMULTISIGEDITROLE is a free data retrieval call binding the contract method 0xc5f66412.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_ROLE() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCaller) CONTRACTMULTISIGEDITROLE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MiningAgentStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_EDIT_ROLE")
	return *ret0, err
}

// CONTRACTMULTISIGEDITROLE is a free data retrieval call binding the contract method 0xc5f66412.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_ROLE() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplSession) CONTRACTMULTISIGEDITROLE() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.CONTRACTMULTISIGEDITROLE(&_MiningAgentStorageImpl.CallOpts)
}

// CONTRACTMULTISIGEDITROLE is a free data retrieval call binding the contract method 0xc5f66412.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_ROLE() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCallerSession) CONTRACTMULTISIGEDITROLE() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.CONTRACTMULTISIGEDITROLE(&_MiningAgentStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEADMINBYADDRESS is a free data retrieval call binding the contract method 0xc816b091.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_BY_ADDRESS() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCaller) CONTRACTMULTISIGREMOVEADMINBYADDRESS(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MiningAgentStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REMOVE_ADMIN_BY_ADDRESS")
	return *ret0, err
}

// CONTRACTMULTISIGREMOVEADMINBYADDRESS is a free data retrieval call binding the contract method 0xc816b091.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_BY_ADDRESS() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplSession) CONTRACTMULTISIGREMOVEADMINBYADDRESS() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.CONTRACTMULTISIGREMOVEADMINBYADDRESS(&_MiningAgentStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEADMINBYADDRESS is a free data retrieval call binding the contract method 0xc816b091.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_BY_ADDRESS() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCallerSession) CONTRACTMULTISIGREMOVEADMINBYADDRESS() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.CONTRACTMULTISIGREMOVEADMINBYADDRESS(&_MiningAgentStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEADMINBYINDEX is a free data retrieval call binding the contract method 0x751fbea2.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_BY_INDEX() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCaller) CONTRACTMULTISIGREMOVEADMINBYINDEX(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MiningAgentStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REMOVE_ADMIN_BY_INDEX")
	return *ret0, err
}

// CONTRACTMULTISIGREMOVEADMINBYINDEX is a free data retrieval call binding the contract method 0x751fbea2.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_BY_INDEX() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplSession) CONTRACTMULTISIGREMOVEADMINBYINDEX() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.CONTRACTMULTISIGREMOVEADMINBYINDEX(&_MiningAgentStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEADMINBYINDEX is a free data retrieval call binding the contract method 0x751fbea2.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_BY_INDEX() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCallerSession) CONTRACTMULTISIGREMOVEADMINBYINDEX() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.CONTRACTMULTISIGREMOVEADMINBYINDEX(&_MiningAgentStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEADMINROLE is a free data retrieval call binding the contract method 0x53c00200.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_ROLE() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCaller) CONTRACTMULTISIGREMOVEADMINROLE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MiningAgentStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REMOVE_ADMIN_ROLE")
	return *ret0, err
}

// CONTRACTMULTISIGREMOVEADMINROLE is a free data retrieval call binding the contract method 0x53c00200.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_ROLE() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplSession) CONTRACTMULTISIGREMOVEADMINROLE() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.CONTRACTMULTISIGREMOVEADMINROLE(&_MiningAgentStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEADMINROLE is a free data retrieval call binding the contract method 0x53c00200.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_ROLE() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCallerSession) CONTRACTMULTISIGREMOVEADMINROLE() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.CONTRACTMULTISIGREMOVEADMINROLE(&_MiningAgentStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEFEE is a free data retrieval call binding the contract method 0xdb614d25.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_FEE() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCaller) CONTRACTMULTISIGREMOVEFEE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MiningAgentStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REMOVE_FEE")
	return *ret0, err
}

// CONTRACTMULTISIGREMOVEFEE is a free data retrieval call binding the contract method 0xdb614d25.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_FEE() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplSession) CONTRACTMULTISIGREMOVEFEE() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.CONTRACTMULTISIGREMOVEFEE(&_MiningAgentStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEFEE is a free data retrieval call binding the contract method 0xdb614d25.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_FEE() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCallerSession) CONTRACTMULTISIGREMOVEFEE() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.CONTRACTMULTISIGREMOVEFEE(&_MiningAgentStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEROLE is a free data retrieval call binding the contract method 0x23648496.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ROLE() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCaller) CONTRACTMULTISIGREMOVEROLE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MiningAgentStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REMOVE_ROLE")
	return *ret0, err
}

// CONTRACTMULTISIGREMOVEROLE is a free data retrieval call binding the contract method 0x23648496.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ROLE() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplSession) CONTRACTMULTISIGREMOVEROLE() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.CONTRACTMULTISIGREMOVEROLE(&_MiningAgentStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEROLE is a free data retrieval call binding the contract method 0x23648496.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ROLE() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCallerSession) CONTRACTMULTISIGREMOVEROLE() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.CONTRACTMULTISIGREMOVEROLE(&_MiningAgentStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREPLENISHCAPITALIZATION is a free data retrieval call binding the contract method 0x62d5d523.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_CAPITALIZATION() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCaller) CONTRACTMULTISIGREPLENISHCAPITALIZATION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MiningAgentStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REPLENISH_CAPITALIZATION")
	return *ret0, err
}

// CONTRACTMULTISIGREPLENISHCAPITALIZATION is a free data retrieval call binding the contract method 0x62d5d523.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_CAPITALIZATION() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplSession) CONTRACTMULTISIGREPLENISHCAPITALIZATION() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.CONTRACTMULTISIGREPLENISHCAPITALIZATION(&_MiningAgentStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREPLENISHCAPITALIZATION is a free data retrieval call binding the contract method 0x62d5d523.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_CAPITALIZATION() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCallerSession) CONTRACTMULTISIGREPLENISHCAPITALIZATION() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.CONTRACTMULTISIGREPLENISHCAPITALIZATION(&_MiningAgentStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREPLENISHCERTIFICATE is a free data retrieval call binding the contract method 0xfd19b7e1.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_CERTIFICATE() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCaller) CONTRACTMULTISIGREPLENISHCERTIFICATE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MiningAgentStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REPLENISH_CERTIFICATE")
	return *ret0, err
}

// CONTRACTMULTISIGREPLENISHCERTIFICATE is a free data retrieval call binding the contract method 0xfd19b7e1.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_CERTIFICATE() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplSession) CONTRACTMULTISIGREPLENISHCERTIFICATE() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.CONTRACTMULTISIGREPLENISHCERTIFICATE(&_MiningAgentStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREPLENISHCERTIFICATE is a free data retrieval call binding the contract method 0xfd19b7e1.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_CERTIFICATE() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCallerSession) CONTRACTMULTISIGREPLENISHCERTIFICATE() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.CONTRACTMULTISIGREPLENISHCERTIFICATE(&_MiningAgentStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREPLENISHEMISSION is a free data retrieval call binding the contract method 0x0cb59fde.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_EMISSION() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCaller) CONTRACTMULTISIGREPLENISHEMISSION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MiningAgentStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REPLENISH_EMISSION")
	return *ret0, err
}

// CONTRACTMULTISIGREPLENISHEMISSION is a free data retrieval call binding the contract method 0x0cb59fde.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_EMISSION() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplSession) CONTRACTMULTISIGREPLENISHEMISSION() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.CONTRACTMULTISIGREPLENISHEMISSION(&_MiningAgentStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREPLENISHEMISSION is a free data retrieval call binding the contract method 0x0cb59fde.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_EMISSION() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCallerSession) CONTRACTMULTISIGREPLENISHEMISSION() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.CONTRACTMULTISIGREPLENISHEMISSION(&_MiningAgentStorageImpl.CallOpts)
}

// CONTRACTMULTISIGSETADMINROLE is a free data retrieval call binding the contract method 0x8f209ab7.
//
// Solidity: function CONTRACT_MULTISIG_SET_ADMIN_ROLE() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCaller) CONTRACTMULTISIGSETADMINROLE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MiningAgentStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_SET_ADMIN_ROLE")
	return *ret0, err
}

// CONTRACTMULTISIGSETADMINROLE is a free data retrieval call binding the contract method 0x8f209ab7.
//
// Solidity: function CONTRACT_MULTISIG_SET_ADMIN_ROLE() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplSession) CONTRACTMULTISIGSETADMINROLE() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.CONTRACTMULTISIGSETADMINROLE(&_MiningAgentStorageImpl.CallOpts)
}

// CONTRACTMULTISIGSETADMINROLE is a free data retrieval call binding the contract method 0x8f209ab7.
//
// Solidity: function CONTRACT_MULTISIG_SET_ADMIN_ROLE() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCallerSession) CONTRACTMULTISIGSETADMINROLE() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.CONTRACTMULTISIGSETADMINROLE(&_MiningAgentStorageImpl.CallOpts)
}

// CONTRACTMULTISIGSETPERMISSION is a free data retrieval call binding the contract method 0xb0e91914.
//
// Solidity: function CONTRACT_MULTISIG_SET_PERMISSION() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCaller) CONTRACTMULTISIGSETPERMISSION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MiningAgentStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_SET_PERMISSION")
	return *ret0, err
}

// CONTRACTMULTISIGSETPERMISSION is a free data retrieval call binding the contract method 0xb0e91914.
//
// Solidity: function CONTRACT_MULTISIG_SET_PERMISSION() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplSession) CONTRACTMULTISIGSETPERMISSION() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.CONTRACTMULTISIGSETPERMISSION(&_MiningAgentStorageImpl.CallOpts)
}

// CONTRACTMULTISIGSETPERMISSION is a free data retrieval call binding the contract method 0xb0e91914.
//
// Solidity: function CONTRACT_MULTISIG_SET_PERMISSION() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCallerSession) CONTRACTMULTISIGSETPERMISSION() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.CONTRACTMULTISIGSETPERMISSION(&_MiningAgentStorageImpl.CallOpts)
}

// CONTRACTMULTISIGSTORAGE is a free data retrieval call binding the contract method 0xe7e1a303.
//
// Solidity: function CONTRACT_MULTISIG_STORAGE() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCaller) CONTRACTMULTISIGSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MiningAgentStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_STORAGE")
	return *ret0, err
}

// CONTRACTMULTISIGSTORAGE is a free data retrieval call binding the contract method 0xe7e1a303.
//
// Solidity: function CONTRACT_MULTISIG_STORAGE() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplSession) CONTRACTMULTISIGSTORAGE() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.CONTRACTMULTISIGSTORAGE(&_MiningAgentStorageImpl.CallOpts)
}

// CONTRACTMULTISIGSTORAGE is a free data retrieval call binding the contract method 0xe7e1a303.
//
// Solidity: function CONTRACT_MULTISIG_STORAGE() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCallerSession) CONTRACTMULTISIGSTORAGE() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.CONTRACTMULTISIGSTORAGE(&_MiningAgentStorageImpl.CallOpts)
}

// CONTRACTOTC is a free data retrieval call binding the contract method 0x5eceac70.
//
// Solidity: function CONTRACT_OTC() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCaller) CONTRACTOTC(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MiningAgentStorageImpl.contract.Call(opts, out, "CONTRACT_OTC")
	return *ret0, err
}

// CONTRACTOTC is a free data retrieval call binding the contract method 0x5eceac70.
//
// Solidity: function CONTRACT_OTC() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplSession) CONTRACTOTC() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.CONTRACTOTC(&_MiningAgentStorageImpl.CallOpts)
}

// CONTRACTOTC is a free data retrieval call binding the contract method 0x5eceac70.
//
// Solidity: function CONTRACT_OTC() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCallerSession) CONTRACTOTC() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.CONTRACTOTC(&_MiningAgentStorageImpl.CallOpts)
}

// CONTRACTOTCSTORAGE is a free data retrieval call binding the contract method 0x825b476f.
//
// Solidity: function CONTRACT_OTC_STORAGE() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCaller) CONTRACTOTCSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MiningAgentStorageImpl.contract.Call(opts, out, "CONTRACT_OTC_STORAGE")
	return *ret0, err
}

// CONTRACTOTCSTORAGE is a free data retrieval call binding the contract method 0x825b476f.
//
// Solidity: function CONTRACT_OTC_STORAGE() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplSession) CONTRACTOTCSTORAGE() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.CONTRACTOTCSTORAGE(&_MiningAgentStorageImpl.CallOpts)
}

// CONTRACTOTCSTORAGE is a free data retrieval call binding the contract method 0x825b476f.
//
// Solidity: function CONTRACT_OTC_STORAGE() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCallerSession) CONTRACTOTCSTORAGE() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.CONTRACTOTCSTORAGE(&_MiningAgentStorageImpl.CallOpts)
}

// CONTRACTREPLENISHCAPITALIZATION is a free data retrieval call binding the contract method 0xff5e09d8.
//
// Solidity: function CONTRACT_REPLENISH_CAPITALIZATION() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCaller) CONTRACTREPLENISHCAPITALIZATION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MiningAgentStorageImpl.contract.Call(opts, out, "CONTRACT_REPLENISH_CAPITALIZATION")
	return *ret0, err
}

// CONTRACTREPLENISHCAPITALIZATION is a free data retrieval call binding the contract method 0xff5e09d8.
//
// Solidity: function CONTRACT_REPLENISH_CAPITALIZATION() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplSession) CONTRACTREPLENISHCAPITALIZATION() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.CONTRACTREPLENISHCAPITALIZATION(&_MiningAgentStorageImpl.CallOpts)
}

// CONTRACTREPLENISHCAPITALIZATION is a free data retrieval call binding the contract method 0xff5e09d8.
//
// Solidity: function CONTRACT_REPLENISH_CAPITALIZATION() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCallerSession) CONTRACTREPLENISHCAPITALIZATION() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.CONTRACTREPLENISHCAPITALIZATION(&_MiningAgentStorageImpl.CallOpts)
}

// CONTRACTRESERVEDALIASSTORAGE is a free data retrieval call binding the contract method 0x70a94aa7.
//
// Solidity: function CONTRACT_RESERVED_ALIAS_STORAGE() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCaller) CONTRACTRESERVEDALIASSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MiningAgentStorageImpl.contract.Call(opts, out, "CONTRACT_RESERVED_ALIAS_STORAGE")
	return *ret0, err
}

// CONTRACTRESERVEDALIASSTORAGE is a free data retrieval call binding the contract method 0x70a94aa7.
//
// Solidity: function CONTRACT_RESERVED_ALIAS_STORAGE() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplSession) CONTRACTRESERVEDALIASSTORAGE() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.CONTRACTRESERVEDALIASSTORAGE(&_MiningAgentStorageImpl.CallOpts)
}

// CONTRACTRESERVEDALIASSTORAGE is a free data retrieval call binding the contract method 0x70a94aa7.
//
// Solidity: function CONTRACT_RESERVED_ALIAS_STORAGE() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCallerSession) CONTRACTRESERVEDALIASSTORAGE() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.CONTRACTRESERVEDALIASSTORAGE(&_MiningAgentStorageImpl.CallOpts)
}

// CONTRACTROLESTORAGE is a free data retrieval call binding the contract method 0xb938235b.
//
// Solidity: function CONTRACT_ROLE_STORAGE() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCaller) CONTRACTROLESTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MiningAgentStorageImpl.contract.Call(opts, out, "CONTRACT_ROLE_STORAGE")
	return *ret0, err
}

// CONTRACTROLESTORAGE is a free data retrieval call binding the contract method 0xb938235b.
//
// Solidity: function CONTRACT_ROLE_STORAGE() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplSession) CONTRACTROLESTORAGE() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.CONTRACTROLESTORAGE(&_MiningAgentStorageImpl.CallOpts)
}

// CONTRACTROLESTORAGE is a free data retrieval call binding the contract method 0xb938235b.
//
// Solidity: function CONTRACT_ROLE_STORAGE() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCallerSession) CONTRACTROLESTORAGE() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.CONTRACTROLESTORAGE(&_MiningAgentStorageImpl.CallOpts)
}

// CONTRACTSAFETRANSFERSTORAGE is a free data retrieval call binding the contract method 0xe7e45392.
//
// Solidity: function CONTRACT_SAFE_TRANSFER_STORAGE() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCaller) CONTRACTSAFETRANSFERSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MiningAgentStorageImpl.contract.Call(opts, out, "CONTRACT_SAFE_TRANSFER_STORAGE")
	return *ret0, err
}

// CONTRACTSAFETRANSFERSTORAGE is a free data retrieval call binding the contract method 0xe7e45392.
//
// Solidity: function CONTRACT_SAFE_TRANSFER_STORAGE() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplSession) CONTRACTSAFETRANSFERSTORAGE() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.CONTRACTSAFETRANSFERSTORAGE(&_MiningAgentStorageImpl.CallOpts)
}

// CONTRACTSAFETRANSFERSTORAGE is a free data retrieval call binding the contract method 0xe7e45392.
//
// Solidity: function CONTRACT_SAFE_TRANSFER_STORAGE() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCallerSession) CONTRACTSAFETRANSFERSTORAGE() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.CONTRACTSAFETRANSFERSTORAGE(&_MiningAgentStorageImpl.CallOpts)
}

// CONTRACTTRANSFER is a free data retrieval call binding the contract method 0xef94fe38.
//
// Solidity: function CONTRACT_TRANSFER() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCaller) CONTRACTTRANSFER(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MiningAgentStorageImpl.contract.Call(opts, out, "CONTRACT_TRANSFER")
	return *ret0, err
}

// CONTRACTTRANSFER is a free data retrieval call binding the contract method 0xef94fe38.
//
// Solidity: function CONTRACT_TRANSFER() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplSession) CONTRACTTRANSFER() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.CONTRACTTRANSFER(&_MiningAgentStorageImpl.CallOpts)
}

// CONTRACTTRANSFER is a free data retrieval call binding the contract method 0xef94fe38.
//
// Solidity: function CONTRACT_TRANSFER() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCallerSession) CONTRACTTRANSFER() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.CONTRACTTRANSFER(&_MiningAgentStorageImpl.CallOpts)
}

// CONTRACTTRANSFERFROMCAPITALIZATION is a free data retrieval call binding the contract method 0xc8f8decb.
//
// Solidity: function CONTRACT_TRANSFER_FROM_CAPITALIZATION() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCaller) CONTRACTTRANSFERFROMCAPITALIZATION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MiningAgentStorageImpl.contract.Call(opts, out, "CONTRACT_TRANSFER_FROM_CAPITALIZATION")
	return *ret0, err
}

// CONTRACTTRANSFERFROMCAPITALIZATION is a free data retrieval call binding the contract method 0xc8f8decb.
//
// Solidity: function CONTRACT_TRANSFER_FROM_CAPITALIZATION() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplSession) CONTRACTTRANSFERFROMCAPITALIZATION() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.CONTRACTTRANSFERFROMCAPITALIZATION(&_MiningAgentStorageImpl.CallOpts)
}

// CONTRACTTRANSFERFROMCAPITALIZATION is a free data retrieval call binding the contract method 0xc8f8decb.
//
// Solidity: function CONTRACT_TRANSFER_FROM_CAPITALIZATION() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCallerSession) CONTRACTTRANSFERFROMCAPITALIZATION() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.CONTRACTTRANSFERFROMCAPITALIZATION(&_MiningAgentStorageImpl.CallOpts)
}

// PERMISSIONACCESSRESTORATION is a free data retrieval call binding the contract method 0x5c3053bf.
//
// Solidity: function PERMISSION_ACCESS_RESTORATION() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCaller) PERMISSIONACCESSRESTORATION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MiningAgentStorageImpl.contract.Call(opts, out, "PERMISSION_ACCESS_RESTORATION")
	return *ret0, err
}

// PERMISSIONACCESSRESTORATION is a free data retrieval call binding the contract method 0x5c3053bf.
//
// Solidity: function PERMISSION_ACCESS_RESTORATION() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplSession) PERMISSIONACCESSRESTORATION() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.PERMISSIONACCESSRESTORATION(&_MiningAgentStorageImpl.CallOpts)
}

// PERMISSIONACCESSRESTORATION is a free data retrieval call binding the contract method 0x5c3053bf.
//
// Solidity: function PERMISSION_ACCESS_RESTORATION() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCallerSession) PERMISSIONACCESSRESTORATION() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.PERMISSIONACCESSRESTORATION(&_MiningAgentStorageImpl.CallOpts)
}

// PERMISSIONADMINSTORAGE is a free data retrieval call binding the contract method 0x0c9fc0db.
//
// Solidity: function PERMISSION_ADMIN_STORAGE() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCaller) PERMISSIONADMINSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MiningAgentStorageImpl.contract.Call(opts, out, "PERMISSION_ADMIN_STORAGE")
	return *ret0, err
}

// PERMISSIONADMINSTORAGE is a free data retrieval call binding the contract method 0x0c9fc0db.
//
// Solidity: function PERMISSION_ADMIN_STORAGE() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplSession) PERMISSIONADMINSTORAGE() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.PERMISSIONADMINSTORAGE(&_MiningAgentStorageImpl.CallOpts)
}

// PERMISSIONADMINSTORAGE is a free data retrieval call binding the contract method 0x0c9fc0db.
//
// Solidity: function PERMISSION_ADMIN_STORAGE() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCallerSession) PERMISSIONADMINSTORAGE() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.PERMISSIONADMINSTORAGE(&_MiningAgentStorageImpl.CallOpts)
}

// PERMISSIONCONTRACTSSTORAGE is a free data retrieval call binding the contract method 0x1516e83b.
//
// Solidity: function PERMISSION_CONTRACTS_STORAGE() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCaller) PERMISSIONCONTRACTSSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MiningAgentStorageImpl.contract.Call(opts, out, "PERMISSION_CONTRACTS_STORAGE")
	return *ret0, err
}

// PERMISSIONCONTRACTSSTORAGE is a free data retrieval call binding the contract method 0x1516e83b.
//
// Solidity: function PERMISSION_CONTRACTS_STORAGE() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplSession) PERMISSIONCONTRACTSSTORAGE() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.PERMISSIONCONTRACTSSTORAGE(&_MiningAgentStorageImpl.CallOpts)
}

// PERMISSIONCONTRACTSSTORAGE is a free data retrieval call binding the contract method 0x1516e83b.
//
// Solidity: function PERMISSION_CONTRACTS_STORAGE() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCallerSession) PERMISSIONCONTRACTSSTORAGE() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.PERMISSIONCONTRACTSSTORAGE(&_MiningAgentStorageImpl.CallOpts)
}

// PERMISSIONCONTRACTDEPLOY is a free data retrieval call binding the contract method 0x257a4a04.
//
// Solidity: function PERMISSION_CONTRACT_DEPLOY() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCaller) PERMISSIONCONTRACTDEPLOY(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MiningAgentStorageImpl.contract.Call(opts, out, "PERMISSION_CONTRACT_DEPLOY")
	return *ret0, err
}

// PERMISSIONCONTRACTDEPLOY is a free data retrieval call binding the contract method 0x257a4a04.
//
// Solidity: function PERMISSION_CONTRACT_DEPLOY() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplSession) PERMISSIONCONTRACTDEPLOY() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.PERMISSIONCONTRACTDEPLOY(&_MiningAgentStorageImpl.CallOpts)
}

// PERMISSIONCONTRACTDEPLOY is a free data retrieval call binding the contract method 0x257a4a04.
//
// Solidity: function PERMISSION_CONTRACT_DEPLOY() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCallerSession) PERMISSIONCONTRACTDEPLOY() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.PERMISSIONCONTRACTDEPLOY(&_MiningAgentStorageImpl.CallOpts)
}

// PERMISSIONEMISSION is a free data retrieval call binding the contract method 0x4ccd92d1.
//
// Solidity: function PERMISSION_EMISSION() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCaller) PERMISSIONEMISSION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MiningAgentStorageImpl.contract.Call(opts, out, "PERMISSION_EMISSION")
	return *ret0, err
}

// PERMISSIONEMISSION is a free data retrieval call binding the contract method 0x4ccd92d1.
//
// Solidity: function PERMISSION_EMISSION() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplSession) PERMISSIONEMISSION() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.PERMISSIONEMISSION(&_MiningAgentStorageImpl.CallOpts)
}

// PERMISSIONEMISSION is a free data retrieval call binding the contract method 0x4ccd92d1.
//
// Solidity: function PERMISSION_EMISSION() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCallerSession) PERMISSIONEMISSION() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.PERMISSIONEMISSION(&_MiningAgentStorageImpl.CallOpts)
}

// PERMISSIONFEESSTORAGE is a free data retrieval call binding the contract method 0xe837eb87.
//
// Solidity: function PERMISSION_FEES_STORAGE() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCaller) PERMISSIONFEESSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MiningAgentStorageImpl.contract.Call(opts, out, "PERMISSION_FEES_STORAGE")
	return *ret0, err
}

// PERMISSIONFEESSTORAGE is a free data retrieval call binding the contract method 0xe837eb87.
//
// Solidity: function PERMISSION_FEES_STORAGE() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplSession) PERMISSIONFEESSTORAGE() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.PERMISSIONFEESSTORAGE(&_MiningAgentStorageImpl.CallOpts)
}

// PERMISSIONFEESSTORAGE is a free data retrieval call binding the contract method 0xe837eb87.
//
// Solidity: function PERMISSION_FEES_STORAGE() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCallerSession) PERMISSIONFEESSTORAGE() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.PERMISSIONFEESSTORAGE(&_MiningAgentStorageImpl.CallOpts)
}

// PERMISSIONMULTISIGSTORAGE is a free data retrieval call binding the contract method 0xf1bb7aad.
//
// Solidity: function PERMISSION_MULTISIG_STORAGE() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCaller) PERMISSIONMULTISIGSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MiningAgentStorageImpl.contract.Call(opts, out, "PERMISSION_MULTISIG_STORAGE")
	return *ret0, err
}

// PERMISSIONMULTISIGSTORAGE is a free data retrieval call binding the contract method 0xf1bb7aad.
//
// Solidity: function PERMISSION_MULTISIG_STORAGE() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplSession) PERMISSIONMULTISIGSTORAGE() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.PERMISSIONMULTISIGSTORAGE(&_MiningAgentStorageImpl.CallOpts)
}

// PERMISSIONMULTISIGSTORAGE is a free data retrieval call binding the contract method 0xf1bb7aad.
//
// Solidity: function PERMISSION_MULTISIG_STORAGE() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCallerSession) PERMISSIONMULTISIGSTORAGE() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.PERMISSIONMULTISIGSTORAGE(&_MiningAgentStorageImpl.CallOpts)
}

// PERMISSIONROLESSTORAGE is a free data retrieval call binding the contract method 0x12a16032.
//
// Solidity: function PERMISSION_ROLES_STORAGE() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCaller) PERMISSIONROLESSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MiningAgentStorageImpl.contract.Call(opts, out, "PERMISSION_ROLES_STORAGE")
	return *ret0, err
}

// PERMISSIONROLESSTORAGE is a free data retrieval call binding the contract method 0x12a16032.
//
// Solidity: function PERMISSION_ROLES_STORAGE() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplSession) PERMISSIONROLESSTORAGE() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.PERMISSIONROLESSTORAGE(&_MiningAgentStorageImpl.CallOpts)
}

// PERMISSIONROLESSTORAGE is a free data retrieval call binding the contract method 0x12a16032.
//
// Solidity: function PERMISSION_ROLES_STORAGE() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCallerSession) PERMISSIONROLESSTORAGE() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.PERMISSIONROLESSTORAGE(&_MiningAgentStorageImpl.CallOpts)
}

// ROLEACCESSVERIFIER is a free data retrieval call binding the contract method 0xc5405f92.
//
// Solidity: function ROLE_ACCESS_VERIFIER() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCaller) ROLEACCESSVERIFIER(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MiningAgentStorageImpl.contract.Call(opts, out, "ROLE_ACCESS_VERIFIER")
	return *ret0, err
}

// ROLEACCESSVERIFIER is a free data retrieval call binding the contract method 0xc5405f92.
//
// Solidity: function ROLE_ACCESS_VERIFIER() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplSession) ROLEACCESSVERIFIER() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.ROLEACCESSVERIFIER(&_MiningAgentStorageImpl.CallOpts)
}

// ROLEACCESSVERIFIER is a free data retrieval call binding the contract method 0xc5405f92.
//
// Solidity: function ROLE_ACCESS_VERIFIER() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCallerSession) ROLEACCESSVERIFIER() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.ROLEACCESSVERIFIER(&_MiningAgentStorageImpl.CallOpts)
}

// ROLEBITBONISSUEVERIFIER is a free data retrieval call binding the contract method 0xf0a012a6.
//
// Solidity: function ROLE_BITBON_ISSUE_VERIFIER() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCaller) ROLEBITBONISSUEVERIFIER(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MiningAgentStorageImpl.contract.Call(opts, out, "ROLE_BITBON_ISSUE_VERIFIER")
	return *ret0, err
}

// ROLEBITBONISSUEVERIFIER is a free data retrieval call binding the contract method 0xf0a012a6.
//
// Solidity: function ROLE_BITBON_ISSUE_VERIFIER() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplSession) ROLEBITBONISSUEVERIFIER() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.ROLEBITBONISSUEVERIFIER(&_MiningAgentStorageImpl.CallOpts)
}

// ROLEBITBONISSUEVERIFIER is a free data retrieval call binding the contract method 0xf0a012a6.
//
// Solidity: function ROLE_BITBON_ISSUE_VERIFIER() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCallerSession) ROLEBITBONISSUEVERIFIER() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.ROLEBITBONISSUEVERIFIER(&_MiningAgentStorageImpl.CallOpts)
}

// ROLECOMMISSIONVERIFIER is a free data retrieval call binding the contract method 0x8ab3de8c.
//
// Solidity: function ROLE_COMMISSION_VERIFIER() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCaller) ROLECOMMISSIONVERIFIER(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MiningAgentStorageImpl.contract.Call(opts, out, "ROLE_COMMISSION_VERIFIER")
	return *ret0, err
}

// ROLECOMMISSIONVERIFIER is a free data retrieval call binding the contract method 0x8ab3de8c.
//
// Solidity: function ROLE_COMMISSION_VERIFIER() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplSession) ROLECOMMISSIONVERIFIER() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.ROLECOMMISSIONVERIFIER(&_MiningAgentStorageImpl.CallOpts)
}

// ROLECOMMISSIONVERIFIER is a free data retrieval call binding the contract method 0x8ab3de8c.
//
// Solidity: function ROLE_COMMISSION_VERIFIER() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCallerSession) ROLECOMMISSIONVERIFIER() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.ROLECOMMISSIONVERIFIER(&_MiningAgentStorageImpl.CallOpts)
}

// ROLEDEPLOYADMIN is a free data retrieval call binding the contract method 0x667bf6e5.
//
// Solidity: function ROLE_DEPLOY_ADMIN() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCaller) ROLEDEPLOYADMIN(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MiningAgentStorageImpl.contract.Call(opts, out, "ROLE_DEPLOY_ADMIN")
	return *ret0, err
}

// ROLEDEPLOYADMIN is a free data retrieval call binding the contract method 0x667bf6e5.
//
// Solidity: function ROLE_DEPLOY_ADMIN() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplSession) ROLEDEPLOYADMIN() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.ROLEDEPLOYADMIN(&_MiningAgentStorageImpl.CallOpts)
}

// ROLEDEPLOYADMIN is a free data retrieval call binding the contract method 0x667bf6e5.
//
// Solidity: function ROLE_DEPLOY_ADMIN() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCallerSession) ROLEDEPLOYADMIN() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.ROLEDEPLOYADMIN(&_MiningAgentStorageImpl.CallOpts)
}

// ROLEPERMISSIONADMIN is a free data retrieval call binding the contract method 0x1b239996.
//
// Solidity: function ROLE_PERMISSION_ADMIN() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCaller) ROLEPERMISSIONADMIN(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MiningAgentStorageImpl.contract.Call(opts, out, "ROLE_PERMISSION_ADMIN")
	return *ret0, err
}

// ROLEPERMISSIONADMIN is a free data retrieval call binding the contract method 0x1b239996.
//
// Solidity: function ROLE_PERMISSION_ADMIN() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplSession) ROLEPERMISSIONADMIN() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.ROLEPERMISSIONADMIN(&_MiningAgentStorageImpl.CallOpts)
}

// ROLEPERMISSIONADMIN is a free data retrieval call binding the contract method 0x1b239996.
//
// Solidity: function ROLE_PERMISSION_ADMIN() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCallerSession) ROLEPERMISSIONADMIN() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.ROLEPERMISSIONADMIN(&_MiningAgentStorageImpl.CallOpts)
}

// ContractStorage is a free data retrieval call binding the contract method 0x0549b27e.
//
// Solidity: function contractStorage() view returns(address)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCaller) ContractStorage(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MiningAgentStorageImpl.contract.Call(opts, out, "contractStorage")
	return *ret0, err
}

// ContractStorage is a free data retrieval call binding the contract method 0x0549b27e.
//
// Solidity: function contractStorage() view returns(address)
func (_MiningAgentStorageImpl *MiningAgentStorageImplSession) ContractStorage() (common.Address, error) {
	return _MiningAgentStorageImpl.Contract.ContractStorage(&_MiningAgentStorageImpl.CallOpts)
}

// ContractStorage is a free data retrieval call binding the contract method 0x0549b27e.
//
// Solidity: function contractStorage() view returns(address)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCallerSession) ContractStorage() (common.Address, error) {
	return _MiningAgentStorageImpl.Contract.ContractStorage(&_MiningAgentStorageImpl.CallOpts)
}

// GetAll is a free data retrieval call binding the contract method 0x53ed5143.
//
// Solidity: function getAll() view returns(address[])
func (_MiningAgentStorageImpl *MiningAgentStorageImplCaller) GetAll(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _MiningAgentStorageImpl.contract.Call(opts, out, "getAll")
	return *ret0, err
}

// GetAll is a free data retrieval call binding the contract method 0x53ed5143.
//
// Solidity: function getAll() view returns(address[])
func (_MiningAgentStorageImpl *MiningAgentStorageImplSession) GetAll() ([]common.Address, error) {
	return _MiningAgentStorageImpl.Contract.GetAll(&_MiningAgentStorageImpl.CallOpts)
}

// GetAll is a free data retrieval call binding the contract method 0x53ed5143.
//
// Solidity: function getAll() view returns(address[])
func (_MiningAgentStorageImpl *MiningAgentStorageImplCallerSession) GetAll() ([]common.Address, error) {
	return _MiningAgentStorageImpl.Contract.GetAll(&_MiningAgentStorageImpl.CallOpts)
}

// GetAt is a free data retrieval call binding the contract method 0x303eaeed.
//
// Solidity: function getAt(uint256 index) view returns(address)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCaller) GetAt(opts *bind.CallOpts, index *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MiningAgentStorageImpl.contract.Call(opts, out, "getAt", index)
	return *ret0, err
}

// GetAt is a free data retrieval call binding the contract method 0x303eaeed.
//
// Solidity: function getAt(uint256 index) view returns(address)
func (_MiningAgentStorageImpl *MiningAgentStorageImplSession) GetAt(index *big.Int) (common.Address, error) {
	return _MiningAgentStorageImpl.Contract.GetAt(&_MiningAgentStorageImpl.CallOpts, index)
}

// GetAt is a free data retrieval call binding the contract method 0x303eaeed.
//
// Solidity: function getAt(uint256 index) view returns(address)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCallerSession) GetAt(index *big.Int) (common.Address, error) {
	return _MiningAgentStorageImpl.Contract.GetAt(&_MiningAgentStorageImpl.CallOpts, index)
}

// GetThisContractIndex is a free data retrieval call binding the contract method 0x70c6d367.
//
// Solidity: function getThisContractIndex() pure returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCaller) GetThisContractIndex(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MiningAgentStorageImpl.contract.Call(opts, out, "getThisContractIndex")
	return *ret0, err
}

// GetThisContractIndex is a free data retrieval call binding the contract method 0x70c6d367.
//
// Solidity: function getThisContractIndex() pure returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplSession) GetThisContractIndex() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.GetThisContractIndex(&_MiningAgentStorageImpl.CallOpts)
}

// GetThisContractIndex is a free data retrieval call binding the contract method 0x70c6d367.
//
// Solidity: function getThisContractIndex() pure returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCallerSession) GetThisContractIndex() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.GetThisContractIndex(&_MiningAgentStorageImpl.CallOpts)
}

// IndexOf is a free data retrieval call binding the contract method 0xfd6aad25.
//
// Solidity: function indexOf(address assetbox) view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCaller) IndexOf(opts *bind.CallOpts, assetbox common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MiningAgentStorageImpl.contract.Call(opts, out, "indexOf", assetbox)
	return *ret0, err
}

// IndexOf is a free data retrieval call binding the contract method 0xfd6aad25.
//
// Solidity: function indexOf(address assetbox) view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplSession) IndexOf(assetbox common.Address) (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.IndexOf(&_MiningAgentStorageImpl.CallOpts, assetbox)
}

// IndexOf is a free data retrieval call binding the contract method 0xfd6aad25.
//
// Solidity: function indexOf(address assetbox) view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCallerSession) IndexOf(assetbox common.Address) (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.IndexOf(&_MiningAgentStorageImpl.CallOpts, assetbox)
}

// IsMiningAgent is a free data retrieval call binding the contract method 0x9d2043d0.
//
// Solidity: function isMiningAgent(address assetbox) view returns(bool)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCaller) IsMiningAgent(opts *bind.CallOpts, assetbox common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _MiningAgentStorageImpl.contract.Call(opts, out, "isMiningAgent", assetbox)
	return *ret0, err
}

// IsMiningAgent is a free data retrieval call binding the contract method 0x9d2043d0.
//
// Solidity: function isMiningAgent(address assetbox) view returns(bool)
func (_MiningAgentStorageImpl *MiningAgentStorageImplSession) IsMiningAgent(assetbox common.Address) (bool, error) {
	return _MiningAgentStorageImpl.Contract.IsMiningAgent(&_MiningAgentStorageImpl.CallOpts, assetbox)
}

// IsMiningAgent is a free data retrieval call binding the contract method 0x9d2043d0.
//
// Solidity: function isMiningAgent(address assetbox) view returns(bool)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCallerSession) IsMiningAgent(assetbox common.Address) (bool, error) {
	return _MiningAgentStorageImpl.Contract.IsMiningAgent(&_MiningAgentStorageImpl.CallOpts, assetbox)
}

// Length is a free data retrieval call binding the contract method 0x1f7b6d32.
//
// Solidity: function length() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCaller) Length(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MiningAgentStorageImpl.contract.Call(opts, out, "length")
	return *ret0, err
}

// Length is a free data retrieval call binding the contract method 0x1f7b6d32.
//
// Solidity: function length() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplSession) Length() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.Length(&_MiningAgentStorageImpl.CallOpts)
}

// Length is a free data retrieval call binding the contract method 0x1f7b6d32.
//
// Solidity: function length() view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCallerSession) Length() (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.Length(&_MiningAgentStorageImpl.CallOpts)
}

// MiningAgentIndex is a free data retrieval call binding the contract method 0x9f70a846.
//
// Solidity: function miningAgentIndex(address ) view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCaller) MiningAgentIndex(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MiningAgentStorageImpl.contract.Call(opts, out, "miningAgentIndex", arg0)
	return *ret0, err
}

// MiningAgentIndex is a free data retrieval call binding the contract method 0x9f70a846.
//
// Solidity: function miningAgentIndex(address ) view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplSession) MiningAgentIndex(arg0 common.Address) (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.MiningAgentIndex(&_MiningAgentStorageImpl.CallOpts, arg0)
}

// MiningAgentIndex is a free data retrieval call binding the contract method 0x9f70a846.
//
// Solidity: function miningAgentIndex(address ) view returns(uint256)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCallerSession) MiningAgentIndex(arg0 common.Address) (*big.Int, error) {
	return _MiningAgentStorageImpl.Contract.MiningAgentIndex(&_MiningAgentStorageImpl.CallOpts, arg0)
}

// MiningAgents is a free data retrieval call binding the contract method 0x636a5c2e.
//
// Solidity: function miningAgents(uint256 ) view returns(address)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCaller) MiningAgents(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MiningAgentStorageImpl.contract.Call(opts, out, "miningAgents", arg0)
	return *ret0, err
}

// MiningAgents is a free data retrieval call binding the contract method 0x636a5c2e.
//
// Solidity: function miningAgents(uint256 ) view returns(address)
func (_MiningAgentStorageImpl *MiningAgentStorageImplSession) MiningAgents(arg0 *big.Int) (common.Address, error) {
	return _MiningAgentStorageImpl.Contract.MiningAgents(&_MiningAgentStorageImpl.CallOpts, arg0)
}

// MiningAgents is a free data retrieval call binding the contract method 0x636a5c2e.
//
// Solidity: function miningAgents(uint256 ) view returns(address)
func (_MiningAgentStorageImpl *MiningAgentStorageImplCallerSession) MiningAgents(arg0 *big.Int) (common.Address, error) {
	return _MiningAgentStorageImpl.Contract.MiningAgents(&_MiningAgentStorageImpl.CallOpts, arg0)
}

// Add is a paid mutator transaction binding the contract method 0x0a3b0a4f.
//
// Solidity: function add(address assetbox) returns(bool)
func (_MiningAgentStorageImpl *MiningAgentStorageImplTransactor) Add(opts *bind.TransactOpts, assetbox common.Address) (*types.Transaction, error) {
	return _MiningAgentStorageImpl.contract.Transact(opts, "add", assetbox)
}

// Add is a paid mutator transaction binding the contract method 0x0a3b0a4f.
//
// Solidity: function add(address assetbox) returns(bool)
func (_MiningAgentStorageImpl *MiningAgentStorageImplSession) Add(assetbox common.Address) (*types.Transaction, error) {
	return _MiningAgentStorageImpl.Contract.Add(&_MiningAgentStorageImpl.TransactOpts, assetbox)
}

// Add is a paid mutator transaction binding the contract method 0x0a3b0a4f.
//
// Solidity: function add(address assetbox) returns(bool)
func (_MiningAgentStorageImpl *MiningAgentStorageImplTransactorSession) Add(assetbox common.Address) (*types.Transaction, error) {
	return _MiningAgentStorageImpl.Contract.Add(&_MiningAgentStorageImpl.TransactOpts, assetbox)
}

// Remove is a paid mutator transaction binding the contract method 0x29092d0e.
//
// Solidity: function remove(address assetbox) returns(bool)
func (_MiningAgentStorageImpl *MiningAgentStorageImplTransactor) Remove(opts *bind.TransactOpts, assetbox common.Address) (*types.Transaction, error) {
	return _MiningAgentStorageImpl.contract.Transact(opts, "remove", assetbox)
}

// Remove is a paid mutator transaction binding the contract method 0x29092d0e.
//
// Solidity: function remove(address assetbox) returns(bool)
func (_MiningAgentStorageImpl *MiningAgentStorageImplSession) Remove(assetbox common.Address) (*types.Transaction, error) {
	return _MiningAgentStorageImpl.Contract.Remove(&_MiningAgentStorageImpl.TransactOpts, assetbox)
}

// Remove is a paid mutator transaction binding the contract method 0x29092d0e.
//
// Solidity: function remove(address assetbox) returns(bool)
func (_MiningAgentStorageImpl *MiningAgentStorageImplTransactorSession) Remove(assetbox common.Address) (*types.Transaction, error) {
	return _MiningAgentStorageImpl.Contract.Remove(&_MiningAgentStorageImpl.TransactOpts, assetbox)
}

// RemoveAt is a paid mutator transaction binding the contract method 0xc624c6f2.
//
// Solidity: function removeAt(uint256 index) returns(bool)
func (_MiningAgentStorageImpl *MiningAgentStorageImplTransactor) RemoveAt(opts *bind.TransactOpts, index *big.Int) (*types.Transaction, error) {
	return _MiningAgentStorageImpl.contract.Transact(opts, "removeAt", index)
}

// RemoveAt is a paid mutator transaction binding the contract method 0xc624c6f2.
//
// Solidity: function removeAt(uint256 index) returns(bool)
func (_MiningAgentStorageImpl *MiningAgentStorageImplSession) RemoveAt(index *big.Int) (*types.Transaction, error) {
	return _MiningAgentStorageImpl.Contract.RemoveAt(&_MiningAgentStorageImpl.TransactOpts, index)
}

// RemoveAt is a paid mutator transaction binding the contract method 0xc624c6f2.
//
// Solidity: function removeAt(uint256 index) returns(bool)
func (_MiningAgentStorageImpl *MiningAgentStorageImplTransactorSession) RemoveAt(index *big.Int) (*types.Transaction, error) {
	return _MiningAgentStorageImpl.Contract.RemoveAt(&_MiningAgentStorageImpl.TransactOpts, index)
}

// MiningAgentStorageImplMiningAgentAddedIterator is returned from FilterMiningAgentAdded and is used to iterate over the raw logs and unpacked data for MiningAgentAdded events raised by the MiningAgentStorageImpl contract.
type MiningAgentStorageImplMiningAgentAddedIterator struct {
	Event *MiningAgentStorageImplMiningAgentAdded // Event containing the contract specifics and raw log

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
func (it *MiningAgentStorageImplMiningAgentAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MiningAgentStorageImplMiningAgentAdded)
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
		it.Event = new(MiningAgentStorageImplMiningAgentAdded)
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
func (it *MiningAgentStorageImplMiningAgentAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MiningAgentStorageImplMiningAgentAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MiningAgentStorageImplMiningAgentAdded represents a MiningAgentAdded event raised by the MiningAgentStorageImpl contract.
type MiningAgentStorageImplMiningAgentAdded struct {
	MiningAgentAssetbox common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterMiningAgentAdded is a free log retrieval operation binding the contract event 0x6c5e9b8242ac8385a0529f451a7de9c384369b8df284e64aa30f78c7fa3324c0.
//
// Solidity: event MiningAgentAdded(address miningAgentAssetbox)
func (_MiningAgentStorageImpl *MiningAgentStorageImplFilterer) FilterMiningAgentAdded(opts *bind.FilterOpts) (*MiningAgentStorageImplMiningAgentAddedIterator, error) {

	logs, sub, err := _MiningAgentStorageImpl.contract.FilterLogs(opts, "MiningAgentAdded")
	if err != nil {
		return nil, err
	}
	return &MiningAgentStorageImplMiningAgentAddedIterator{contract: _MiningAgentStorageImpl.contract, event: "MiningAgentAdded", logs: logs, sub: sub}, nil
}

// WatchMiningAgentAdded is a free log subscription operation binding the contract event 0x6c5e9b8242ac8385a0529f451a7de9c384369b8df284e64aa30f78c7fa3324c0.
//
// Solidity: event MiningAgentAdded(address miningAgentAssetbox)
func (_MiningAgentStorageImpl *MiningAgentStorageImplFilterer) WatchMiningAgentAdded(opts *bind.WatchOpts, sink chan<- *MiningAgentStorageImplMiningAgentAdded) (event.Subscription, error) {

	logs, sub, err := _MiningAgentStorageImpl.contract.WatchLogs(opts, "MiningAgentAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MiningAgentStorageImplMiningAgentAdded)
				if err := _MiningAgentStorageImpl.contract.UnpackLog(event, "MiningAgentAdded", log); err != nil {
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

// ParseMiningAgentAdded is a log parse operation binding the contract event 0x6c5e9b8242ac8385a0529f451a7de9c384369b8df284e64aa30f78c7fa3324c0.
//
// Solidity: event MiningAgentAdded(address miningAgentAssetbox)
func (_MiningAgentStorageImpl *MiningAgentStorageImplFilterer) ParseMiningAgentAdded(log types.Log) (*MiningAgentStorageImplMiningAgentAdded, error) {
	event := new(MiningAgentStorageImplMiningAgentAdded)
	if err := _MiningAgentStorageImpl.contract.UnpackLog(event, "MiningAgentAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// MiningAgentStorageImplMiningAgentRemovedIterator is returned from FilterMiningAgentRemoved and is used to iterate over the raw logs and unpacked data for MiningAgentRemoved events raised by the MiningAgentStorageImpl contract.
type MiningAgentStorageImplMiningAgentRemovedIterator struct {
	Event *MiningAgentStorageImplMiningAgentRemoved // Event containing the contract specifics and raw log

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
func (it *MiningAgentStorageImplMiningAgentRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MiningAgentStorageImplMiningAgentRemoved)
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
		it.Event = new(MiningAgentStorageImplMiningAgentRemoved)
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
func (it *MiningAgentStorageImplMiningAgentRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MiningAgentStorageImplMiningAgentRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MiningAgentStorageImplMiningAgentRemoved represents a MiningAgentRemoved event raised by the MiningAgentStorageImpl contract.
type MiningAgentStorageImplMiningAgentRemoved struct {
	MiningAgentAssetbox common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterMiningAgentRemoved is a free log retrieval operation binding the contract event 0x7b993c0010ede270c5461171300777e772478b38aeb4ff8802781da5315388bd.
//
// Solidity: event MiningAgentRemoved(address miningAgentAssetbox)
func (_MiningAgentStorageImpl *MiningAgentStorageImplFilterer) FilterMiningAgentRemoved(opts *bind.FilterOpts) (*MiningAgentStorageImplMiningAgentRemovedIterator, error) {

	logs, sub, err := _MiningAgentStorageImpl.contract.FilterLogs(opts, "MiningAgentRemoved")
	if err != nil {
		return nil, err
	}
	return &MiningAgentStorageImplMiningAgentRemovedIterator{contract: _MiningAgentStorageImpl.contract, event: "MiningAgentRemoved", logs: logs, sub: sub}, nil
}

// WatchMiningAgentRemoved is a free log subscription operation binding the contract event 0x7b993c0010ede270c5461171300777e772478b38aeb4ff8802781da5315388bd.
//
// Solidity: event MiningAgentRemoved(address miningAgentAssetbox)
func (_MiningAgentStorageImpl *MiningAgentStorageImplFilterer) WatchMiningAgentRemoved(opts *bind.WatchOpts, sink chan<- *MiningAgentStorageImplMiningAgentRemoved) (event.Subscription, error) {

	logs, sub, err := _MiningAgentStorageImpl.contract.WatchLogs(opts, "MiningAgentRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MiningAgentStorageImplMiningAgentRemoved)
				if err := _MiningAgentStorageImpl.contract.UnpackLog(event, "MiningAgentRemoved", log); err != nil {
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

// ParseMiningAgentRemoved is a log parse operation binding the contract event 0x7b993c0010ede270c5461171300777e772478b38aeb4ff8802781da5315388bd.
//
// Solidity: event MiningAgentRemoved(address miningAgentAssetbox)
func (_MiningAgentStorageImpl *MiningAgentStorageImplFilterer) ParseMiningAgentRemoved(log types.Log) (*MiningAgentStorageImplMiningAgentRemoved, error) {
	event := new(MiningAgentStorageImplMiningAgentRemoved)
	if err := _MiningAgentStorageImpl.contract.UnpackLog(event, "MiningAgentRemoved", log); err != nil {
		return nil, err
	}
	return event, nil
}
