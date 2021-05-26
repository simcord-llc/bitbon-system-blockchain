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

// AssetboxStorageImplABI is the input ABI used to generate the binding from.
const AssetboxStorageImplABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractStorageAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"assetbox\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"aliasString\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"serviceId\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isMining\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"otherInfo\",\"type\":\"bytes\"}],\"name\":\"AssetboxInfoDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"assetbox\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"aliasString\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"serviceId\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isMining\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"otherInfo\",\"type\":\"bytes\"}],\"name\":\"AssetboxInfoSet\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ACCESS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ACCESS_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ADMINABLE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ADMIN_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ASSETBOX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ASSETBOX_INFO\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ASSETBOX_INFO_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ASSETBOX_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_BALANCE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_BITBON\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_BITBON_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_BITBON_SUPPORT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_BONBALANCE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_BONTRANSFER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_DEX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_DISTRIBUTION_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_EXCHANGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_EXCHANGE_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_FEE_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_GENERATOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MINING_AGENT_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MSBON\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MSBON_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_ADD_ADDRESS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_ADD_ADMIN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_ADD_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_ADD_ROLE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_DROP_PERMISSION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_EDIT_ADMIN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_EDIT_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_EDIT_ROLE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REMOVE_ADMIN_BY_ADDRESS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REMOVE_ADMIN_BY_INDEX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REMOVE_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REMOVE_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REMOVE_ROLE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REPLENISH_CAPITALIZATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REPLENISH_CERTIFICATE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REPLENISH_EMISSION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_SET_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_SET_PERMISSION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_OTC\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_OTC_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_REPLENISH_CAPITALIZATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_RESERVED_ALIAS_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ROLE_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_SAFE_TRANSFER_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_TRANSFER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_TRANSFER_FROM_CAPITALIZATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_ACCESS_RESTORATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_ADMIN_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_CONTRACTS_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_CONTRACT_DEPLOY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_EMISSION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_FEES_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_MULTISIG_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_ROLES_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_ACCESS_VERIFIER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_BITBON_ISSUE_VERIFIER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_COMMISSION_VERIFIER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_DEPLOY_ADMIN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_PERMISSION_ADMIN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"aliasToAssetbox\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"assetboxToAlias\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"assetboxToIsMining\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"assetboxToOtherInfo\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"assetboxToServiceId\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"contractStorage\",\"outputs\":[{\"internalType\":\"contractContractStorage\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getThisContractIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"assetbox\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"aliasString\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"serviceId\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"isMining\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"otherInfo\",\"type\":\"bytes\"}],\"name\":\"setAssetboxInfo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"assetbox\",\"type\":\"address\"}],\"name\":\"deleteAssetboxInfo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"assetbox\",\"type\":\"address\"}],\"name\":\"getAssetboxInfoByAssetbox\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"assetboxAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"assetboxAlias\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"serviceId\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"isMining\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"otherInfo\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"aliases\",\"type\":\"string\"}],\"name\":\"getAssetbox\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"assetbox\",\"type\":\"address\"}],\"name\":\"getServiceId\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"assetbox\",\"type\":\"address\"}],\"name\":\"getIsMining\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"aliasHash\",\"type\":\"bytes32\"}],\"name\":\"isSanitizedAliasTaken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"taken\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// AssetboxStorageImpl is an auto generated Go binding around an Ethereum contract.
type AssetboxStorageImpl struct {
	AssetboxStorageImplCaller     // Read-only binding to the contract
	AssetboxStorageImplTransactor // Write-only binding to the contract
	AssetboxStorageImplFilterer   // Log filterer for contract events
}

// AssetboxStorageImplCaller is an auto generated read-only Go binding around an Ethereum contract.
type AssetboxStorageImplCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssetboxStorageImplTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AssetboxStorageImplTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssetboxStorageImplFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AssetboxStorageImplFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssetboxStorageImplSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AssetboxStorageImplSession struct {
	Contract     *AssetboxStorageImpl // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// AssetboxStorageImplCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AssetboxStorageImplCallerSession struct {
	Contract *AssetboxStorageImplCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// AssetboxStorageImplTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AssetboxStorageImplTransactorSession struct {
	Contract     *AssetboxStorageImplTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// AssetboxStorageImplRaw is an auto generated low-level Go binding around an Ethereum contract.
type AssetboxStorageImplRaw struct {
	Contract *AssetboxStorageImpl // Generic contract binding to access the raw methods on
}

// AssetboxStorageImplCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AssetboxStorageImplCallerRaw struct {
	Contract *AssetboxStorageImplCaller // Generic read-only contract binding to access the raw methods on
}

// AssetboxStorageImplTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AssetboxStorageImplTransactorRaw struct {
	Contract *AssetboxStorageImplTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAssetboxStorageImpl creates a new instance of AssetboxStorageImpl, bound to a specific deployed contract.
func NewAssetboxStorageImpl(address common.Address, backend bind.ContractBackend) (*AssetboxStorageImpl, error) {
	contract, err := bindAssetboxStorageImpl(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AssetboxStorageImpl{AssetboxStorageImplCaller: AssetboxStorageImplCaller{contract: contract}, AssetboxStorageImplTransactor: AssetboxStorageImplTransactor{contract: contract}, AssetboxStorageImplFilterer: AssetboxStorageImplFilterer{contract: contract}}, nil
}

// NewAssetboxStorageImplCaller creates a new read-only instance of AssetboxStorageImpl, bound to a specific deployed contract.
func NewAssetboxStorageImplCaller(address common.Address, caller bind.ContractCaller) (*AssetboxStorageImplCaller, error) {
	contract, err := bindAssetboxStorageImpl(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AssetboxStorageImplCaller{contract: contract}, nil
}

// NewAssetboxStorageImplTransactor creates a new write-only instance of AssetboxStorageImpl, bound to a specific deployed contract.
func NewAssetboxStorageImplTransactor(address common.Address, transactor bind.ContractTransactor) (*AssetboxStorageImplTransactor, error) {
	contract, err := bindAssetboxStorageImpl(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AssetboxStorageImplTransactor{contract: contract}, nil
}

// NewAssetboxStorageImplFilterer creates a new log filterer instance of AssetboxStorageImpl, bound to a specific deployed contract.
func NewAssetboxStorageImplFilterer(address common.Address, filterer bind.ContractFilterer) (*AssetboxStorageImplFilterer, error) {
	contract, err := bindAssetboxStorageImpl(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AssetboxStorageImplFilterer{contract: contract}, nil
}

// bindAssetboxStorageImpl binds a generic wrapper to an already deployed contract.
func bindAssetboxStorageImpl(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AssetboxStorageImplABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AssetboxStorageImpl *AssetboxStorageImplRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AssetboxStorageImpl.Contract.AssetboxStorageImplCaller.contract.Call(opts, result, method, params...)
}

// Return Abi for AssetboxStorageImplCaller
func (_AssetboxStorageImpl *AssetboxStorageImplCaller) ABI() abi.ABI {
	return _AssetboxStorageImpl.contract.Abi
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AssetboxStorageImpl *AssetboxStorageImplRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AssetboxStorageImpl.Contract.AssetboxStorageImplTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AssetboxStorageImpl *AssetboxStorageImplRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AssetboxStorageImpl.Contract.AssetboxStorageImplTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AssetboxStorageImpl *AssetboxStorageImplCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AssetboxStorageImpl.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AssetboxStorageImpl *AssetboxStorageImplTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AssetboxStorageImpl.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AssetboxStorageImpl *AssetboxStorageImplTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AssetboxStorageImpl.Contract.contract.Transact(opts, method, params...)
}

// CONTRACTACCESS is a free data retrieval call binding the contract method 0x0ac62b0c.
//
// Solidity: function CONTRACT_ACCESS() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCaller) CONTRACTACCESS(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxStorageImpl.contract.Call(opts, out, "CONTRACT_ACCESS")
	return *ret0, err
}

// CONTRACTACCESS is a free data retrieval call binding the contract method 0x0ac62b0c.
//
// Solidity: function CONTRACT_ACCESS() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplSession) CONTRACTACCESS() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.CONTRACTACCESS(&_AssetboxStorageImpl.CallOpts)
}

// CONTRACTACCESS is a free data retrieval call binding the contract method 0x0ac62b0c.
//
// Solidity: function CONTRACT_ACCESS() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCallerSession) CONTRACTACCESS() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.CONTRACTACCESS(&_AssetboxStorageImpl.CallOpts)
}

// CONTRACTACCESSSTORAGE is a free data retrieval call binding the contract method 0xceac9bd5.
//
// Solidity: function CONTRACT_ACCESS_STORAGE() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCaller) CONTRACTACCESSSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxStorageImpl.contract.Call(opts, out, "CONTRACT_ACCESS_STORAGE")
	return *ret0, err
}

// CONTRACTACCESSSTORAGE is a free data retrieval call binding the contract method 0xceac9bd5.
//
// Solidity: function CONTRACT_ACCESS_STORAGE() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplSession) CONTRACTACCESSSTORAGE() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.CONTRACTACCESSSTORAGE(&_AssetboxStorageImpl.CallOpts)
}

// CONTRACTACCESSSTORAGE is a free data retrieval call binding the contract method 0xceac9bd5.
//
// Solidity: function CONTRACT_ACCESS_STORAGE() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCallerSession) CONTRACTACCESSSTORAGE() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.CONTRACTACCESSSTORAGE(&_AssetboxStorageImpl.CallOpts)
}

// CONTRACTADMINABLE is a free data retrieval call binding the contract method 0xbbe3a78f.
//
// Solidity: function CONTRACT_ADMINABLE() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCaller) CONTRACTADMINABLE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxStorageImpl.contract.Call(opts, out, "CONTRACT_ADMINABLE")
	return *ret0, err
}

// CONTRACTADMINABLE is a free data retrieval call binding the contract method 0xbbe3a78f.
//
// Solidity: function CONTRACT_ADMINABLE() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplSession) CONTRACTADMINABLE() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.CONTRACTADMINABLE(&_AssetboxStorageImpl.CallOpts)
}

// CONTRACTADMINABLE is a free data retrieval call binding the contract method 0xbbe3a78f.
//
// Solidity: function CONTRACT_ADMINABLE() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCallerSession) CONTRACTADMINABLE() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.CONTRACTADMINABLE(&_AssetboxStorageImpl.CallOpts)
}

// CONTRACTADMINSTORAGE is a free data retrieval call binding the contract method 0xc9aebae9.
//
// Solidity: function CONTRACT_ADMIN_STORAGE() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCaller) CONTRACTADMINSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxStorageImpl.contract.Call(opts, out, "CONTRACT_ADMIN_STORAGE")
	return *ret0, err
}

// CONTRACTADMINSTORAGE is a free data retrieval call binding the contract method 0xc9aebae9.
//
// Solidity: function CONTRACT_ADMIN_STORAGE() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplSession) CONTRACTADMINSTORAGE() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.CONTRACTADMINSTORAGE(&_AssetboxStorageImpl.CallOpts)
}

// CONTRACTADMINSTORAGE is a free data retrieval call binding the contract method 0xc9aebae9.
//
// Solidity: function CONTRACT_ADMIN_STORAGE() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCallerSession) CONTRACTADMINSTORAGE() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.CONTRACTADMINSTORAGE(&_AssetboxStorageImpl.CallOpts)
}

// CONTRACTASSETBOX is a free data retrieval call binding the contract method 0x161aa17c.
//
// Solidity: function CONTRACT_ASSETBOX() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCaller) CONTRACTASSETBOX(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxStorageImpl.contract.Call(opts, out, "CONTRACT_ASSETBOX")
	return *ret0, err
}

// CONTRACTASSETBOX is a free data retrieval call binding the contract method 0x161aa17c.
//
// Solidity: function CONTRACT_ASSETBOX() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplSession) CONTRACTASSETBOX() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.CONTRACTASSETBOX(&_AssetboxStorageImpl.CallOpts)
}

// CONTRACTASSETBOX is a free data retrieval call binding the contract method 0x161aa17c.
//
// Solidity: function CONTRACT_ASSETBOX() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCallerSession) CONTRACTASSETBOX() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.CONTRACTASSETBOX(&_AssetboxStorageImpl.CallOpts)
}

// CONTRACTASSETBOXINFO is a free data retrieval call binding the contract method 0xb572fd3d.
//
// Solidity: function CONTRACT_ASSETBOX_INFO() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCaller) CONTRACTASSETBOXINFO(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxStorageImpl.contract.Call(opts, out, "CONTRACT_ASSETBOX_INFO")
	return *ret0, err
}

// CONTRACTASSETBOXINFO is a free data retrieval call binding the contract method 0xb572fd3d.
//
// Solidity: function CONTRACT_ASSETBOX_INFO() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplSession) CONTRACTASSETBOXINFO() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.CONTRACTASSETBOXINFO(&_AssetboxStorageImpl.CallOpts)
}

// CONTRACTASSETBOXINFO is a free data retrieval call binding the contract method 0xb572fd3d.
//
// Solidity: function CONTRACT_ASSETBOX_INFO() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCallerSession) CONTRACTASSETBOXINFO() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.CONTRACTASSETBOXINFO(&_AssetboxStorageImpl.CallOpts)
}

// CONTRACTASSETBOXINFOSTORAGE is a free data retrieval call binding the contract method 0xf32a115d.
//
// Solidity: function CONTRACT_ASSETBOX_INFO_STORAGE() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCaller) CONTRACTASSETBOXINFOSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxStorageImpl.contract.Call(opts, out, "CONTRACT_ASSETBOX_INFO_STORAGE")
	return *ret0, err
}

// CONTRACTASSETBOXINFOSTORAGE is a free data retrieval call binding the contract method 0xf32a115d.
//
// Solidity: function CONTRACT_ASSETBOX_INFO_STORAGE() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplSession) CONTRACTASSETBOXINFOSTORAGE() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.CONTRACTASSETBOXINFOSTORAGE(&_AssetboxStorageImpl.CallOpts)
}

// CONTRACTASSETBOXINFOSTORAGE is a free data retrieval call binding the contract method 0xf32a115d.
//
// Solidity: function CONTRACT_ASSETBOX_INFO_STORAGE() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCallerSession) CONTRACTASSETBOXINFOSTORAGE() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.CONTRACTASSETBOXINFOSTORAGE(&_AssetboxStorageImpl.CallOpts)
}

// CONTRACTASSETBOXSTORAGE is a free data retrieval call binding the contract method 0xfa5bdafc.
//
// Solidity: function CONTRACT_ASSETBOX_STORAGE() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCaller) CONTRACTASSETBOXSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxStorageImpl.contract.Call(opts, out, "CONTRACT_ASSETBOX_STORAGE")
	return *ret0, err
}

// CONTRACTASSETBOXSTORAGE is a free data retrieval call binding the contract method 0xfa5bdafc.
//
// Solidity: function CONTRACT_ASSETBOX_STORAGE() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplSession) CONTRACTASSETBOXSTORAGE() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.CONTRACTASSETBOXSTORAGE(&_AssetboxStorageImpl.CallOpts)
}

// CONTRACTASSETBOXSTORAGE is a free data retrieval call binding the contract method 0xfa5bdafc.
//
// Solidity: function CONTRACT_ASSETBOX_STORAGE() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCallerSession) CONTRACTASSETBOXSTORAGE() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.CONTRACTASSETBOXSTORAGE(&_AssetboxStorageImpl.CallOpts)
}

// CONTRACTBALANCE is a free data retrieval call binding the contract method 0xd59045f5.
//
// Solidity: function CONTRACT_BALANCE() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCaller) CONTRACTBALANCE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxStorageImpl.contract.Call(opts, out, "CONTRACT_BALANCE")
	return *ret0, err
}

// CONTRACTBALANCE is a free data retrieval call binding the contract method 0xd59045f5.
//
// Solidity: function CONTRACT_BALANCE() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplSession) CONTRACTBALANCE() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.CONTRACTBALANCE(&_AssetboxStorageImpl.CallOpts)
}

// CONTRACTBALANCE is a free data retrieval call binding the contract method 0xd59045f5.
//
// Solidity: function CONTRACT_BALANCE() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCallerSession) CONTRACTBALANCE() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.CONTRACTBALANCE(&_AssetboxStorageImpl.CallOpts)
}

// CONTRACTBITBON is a free data retrieval call binding the contract method 0xb6b433ab.
//
// Solidity: function CONTRACT_BITBON() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCaller) CONTRACTBITBON(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxStorageImpl.contract.Call(opts, out, "CONTRACT_BITBON")
	return *ret0, err
}

// CONTRACTBITBON is a free data retrieval call binding the contract method 0xb6b433ab.
//
// Solidity: function CONTRACT_BITBON() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplSession) CONTRACTBITBON() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.CONTRACTBITBON(&_AssetboxStorageImpl.CallOpts)
}

// CONTRACTBITBON is a free data retrieval call binding the contract method 0xb6b433ab.
//
// Solidity: function CONTRACT_BITBON() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCallerSession) CONTRACTBITBON() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.CONTRACTBITBON(&_AssetboxStorageImpl.CallOpts)
}

// CONTRACTBITBONSTORAGE is a free data retrieval call binding the contract method 0xe7f6c30e.
//
// Solidity: function CONTRACT_BITBON_STORAGE() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCaller) CONTRACTBITBONSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxStorageImpl.contract.Call(opts, out, "CONTRACT_BITBON_STORAGE")
	return *ret0, err
}

// CONTRACTBITBONSTORAGE is a free data retrieval call binding the contract method 0xe7f6c30e.
//
// Solidity: function CONTRACT_BITBON_STORAGE() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplSession) CONTRACTBITBONSTORAGE() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.CONTRACTBITBONSTORAGE(&_AssetboxStorageImpl.CallOpts)
}

// CONTRACTBITBONSTORAGE is a free data retrieval call binding the contract method 0xe7f6c30e.
//
// Solidity: function CONTRACT_BITBON_STORAGE() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCallerSession) CONTRACTBITBONSTORAGE() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.CONTRACTBITBONSTORAGE(&_AssetboxStorageImpl.CallOpts)
}

// CONTRACTBITBONSUPPORT is a free data retrieval call binding the contract method 0x20139c74.
//
// Solidity: function CONTRACT_BITBON_SUPPORT() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCaller) CONTRACTBITBONSUPPORT(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxStorageImpl.contract.Call(opts, out, "CONTRACT_BITBON_SUPPORT")
	return *ret0, err
}

// CONTRACTBITBONSUPPORT is a free data retrieval call binding the contract method 0x20139c74.
//
// Solidity: function CONTRACT_BITBON_SUPPORT() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplSession) CONTRACTBITBONSUPPORT() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.CONTRACTBITBONSUPPORT(&_AssetboxStorageImpl.CallOpts)
}

// CONTRACTBITBONSUPPORT is a free data retrieval call binding the contract method 0x20139c74.
//
// Solidity: function CONTRACT_BITBON_SUPPORT() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCallerSession) CONTRACTBITBONSUPPORT() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.CONTRACTBITBONSUPPORT(&_AssetboxStorageImpl.CallOpts)
}

// CONTRACTBONBALANCE is a free data retrieval call binding the contract method 0x40e02b7c.
//
// Solidity: function CONTRACT_BONBALANCE() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCaller) CONTRACTBONBALANCE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxStorageImpl.contract.Call(opts, out, "CONTRACT_BONBALANCE")
	return *ret0, err
}

// CONTRACTBONBALANCE is a free data retrieval call binding the contract method 0x40e02b7c.
//
// Solidity: function CONTRACT_BONBALANCE() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplSession) CONTRACTBONBALANCE() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.CONTRACTBONBALANCE(&_AssetboxStorageImpl.CallOpts)
}

// CONTRACTBONBALANCE is a free data retrieval call binding the contract method 0x40e02b7c.
//
// Solidity: function CONTRACT_BONBALANCE() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCallerSession) CONTRACTBONBALANCE() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.CONTRACTBONBALANCE(&_AssetboxStorageImpl.CallOpts)
}

// CONTRACTBONTRANSFER is a free data retrieval call binding the contract method 0xc18f4a8a.
//
// Solidity: function CONTRACT_BONTRANSFER() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCaller) CONTRACTBONTRANSFER(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxStorageImpl.contract.Call(opts, out, "CONTRACT_BONTRANSFER")
	return *ret0, err
}

// CONTRACTBONTRANSFER is a free data retrieval call binding the contract method 0xc18f4a8a.
//
// Solidity: function CONTRACT_BONTRANSFER() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplSession) CONTRACTBONTRANSFER() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.CONTRACTBONTRANSFER(&_AssetboxStorageImpl.CallOpts)
}

// CONTRACTBONTRANSFER is a free data retrieval call binding the contract method 0xc18f4a8a.
//
// Solidity: function CONTRACT_BONTRANSFER() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCallerSession) CONTRACTBONTRANSFER() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.CONTRACTBONTRANSFER(&_AssetboxStorageImpl.CallOpts)
}

// CONTRACTDEX is a free data retrieval call binding the contract method 0x52c1f844.
//
// Solidity: function CONTRACT_DEX() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCaller) CONTRACTDEX(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxStorageImpl.contract.Call(opts, out, "CONTRACT_DEX")
	return *ret0, err
}

// CONTRACTDEX is a free data retrieval call binding the contract method 0x52c1f844.
//
// Solidity: function CONTRACT_DEX() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplSession) CONTRACTDEX() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.CONTRACTDEX(&_AssetboxStorageImpl.CallOpts)
}

// CONTRACTDEX is a free data retrieval call binding the contract method 0x52c1f844.
//
// Solidity: function CONTRACT_DEX() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCallerSession) CONTRACTDEX() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.CONTRACTDEX(&_AssetboxStorageImpl.CallOpts)
}

// CONTRACTDISTRIBUTIONSTORAGE is a free data retrieval call binding the contract method 0x08097327.
//
// Solidity: function CONTRACT_DISTRIBUTION_STORAGE() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCaller) CONTRACTDISTRIBUTIONSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxStorageImpl.contract.Call(opts, out, "CONTRACT_DISTRIBUTION_STORAGE")
	return *ret0, err
}

// CONTRACTDISTRIBUTIONSTORAGE is a free data retrieval call binding the contract method 0x08097327.
//
// Solidity: function CONTRACT_DISTRIBUTION_STORAGE() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplSession) CONTRACTDISTRIBUTIONSTORAGE() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.CONTRACTDISTRIBUTIONSTORAGE(&_AssetboxStorageImpl.CallOpts)
}

// CONTRACTDISTRIBUTIONSTORAGE is a free data retrieval call binding the contract method 0x08097327.
//
// Solidity: function CONTRACT_DISTRIBUTION_STORAGE() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCallerSession) CONTRACTDISTRIBUTIONSTORAGE() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.CONTRACTDISTRIBUTIONSTORAGE(&_AssetboxStorageImpl.CallOpts)
}

// CONTRACTEXCHANGE is a free data retrieval call binding the contract method 0xa0e3a10e.
//
// Solidity: function CONTRACT_EXCHANGE() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCaller) CONTRACTEXCHANGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxStorageImpl.contract.Call(opts, out, "CONTRACT_EXCHANGE")
	return *ret0, err
}

// CONTRACTEXCHANGE is a free data retrieval call binding the contract method 0xa0e3a10e.
//
// Solidity: function CONTRACT_EXCHANGE() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplSession) CONTRACTEXCHANGE() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.CONTRACTEXCHANGE(&_AssetboxStorageImpl.CallOpts)
}

// CONTRACTEXCHANGE is a free data retrieval call binding the contract method 0xa0e3a10e.
//
// Solidity: function CONTRACT_EXCHANGE() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCallerSession) CONTRACTEXCHANGE() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.CONTRACTEXCHANGE(&_AssetboxStorageImpl.CallOpts)
}

// CONTRACTEXCHANGESTORAGE is a free data retrieval call binding the contract method 0x001570f2.
//
// Solidity: function CONTRACT_EXCHANGE_STORAGE() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCaller) CONTRACTEXCHANGESTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxStorageImpl.contract.Call(opts, out, "CONTRACT_EXCHANGE_STORAGE")
	return *ret0, err
}

// CONTRACTEXCHANGESTORAGE is a free data retrieval call binding the contract method 0x001570f2.
//
// Solidity: function CONTRACT_EXCHANGE_STORAGE() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplSession) CONTRACTEXCHANGESTORAGE() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.CONTRACTEXCHANGESTORAGE(&_AssetboxStorageImpl.CallOpts)
}

// CONTRACTEXCHANGESTORAGE is a free data retrieval call binding the contract method 0x001570f2.
//
// Solidity: function CONTRACT_EXCHANGE_STORAGE() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCallerSession) CONTRACTEXCHANGESTORAGE() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.CONTRACTEXCHANGESTORAGE(&_AssetboxStorageImpl.CallOpts)
}

// CONTRACTFEESTORAGE is a free data retrieval call binding the contract method 0xcd5e65a4.
//
// Solidity: function CONTRACT_FEE_STORAGE() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCaller) CONTRACTFEESTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxStorageImpl.contract.Call(opts, out, "CONTRACT_FEE_STORAGE")
	return *ret0, err
}

// CONTRACTFEESTORAGE is a free data retrieval call binding the contract method 0xcd5e65a4.
//
// Solidity: function CONTRACT_FEE_STORAGE() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplSession) CONTRACTFEESTORAGE() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.CONTRACTFEESTORAGE(&_AssetboxStorageImpl.CallOpts)
}

// CONTRACTFEESTORAGE is a free data retrieval call binding the contract method 0xcd5e65a4.
//
// Solidity: function CONTRACT_FEE_STORAGE() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCallerSession) CONTRACTFEESTORAGE() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.CONTRACTFEESTORAGE(&_AssetboxStorageImpl.CallOpts)
}

// CONTRACTGENERATOR is a free data retrieval call binding the contract method 0xccea2b46.
//
// Solidity: function CONTRACT_GENERATOR() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCaller) CONTRACTGENERATOR(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxStorageImpl.contract.Call(opts, out, "CONTRACT_GENERATOR")
	return *ret0, err
}

// CONTRACTGENERATOR is a free data retrieval call binding the contract method 0xccea2b46.
//
// Solidity: function CONTRACT_GENERATOR() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplSession) CONTRACTGENERATOR() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.CONTRACTGENERATOR(&_AssetboxStorageImpl.CallOpts)
}

// CONTRACTGENERATOR is a free data retrieval call binding the contract method 0xccea2b46.
//
// Solidity: function CONTRACT_GENERATOR() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCallerSession) CONTRACTGENERATOR() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.CONTRACTGENERATOR(&_AssetboxStorageImpl.CallOpts)
}

// CONTRACTMININGAGENTSTORAGE is a free data retrieval call binding the contract method 0x83cc588d.
//
// Solidity: function CONTRACT_MINING_AGENT_STORAGE() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCaller) CONTRACTMININGAGENTSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxStorageImpl.contract.Call(opts, out, "CONTRACT_MINING_AGENT_STORAGE")
	return *ret0, err
}

// CONTRACTMININGAGENTSTORAGE is a free data retrieval call binding the contract method 0x83cc588d.
//
// Solidity: function CONTRACT_MINING_AGENT_STORAGE() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplSession) CONTRACTMININGAGENTSTORAGE() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.CONTRACTMININGAGENTSTORAGE(&_AssetboxStorageImpl.CallOpts)
}

// CONTRACTMININGAGENTSTORAGE is a free data retrieval call binding the contract method 0x83cc588d.
//
// Solidity: function CONTRACT_MINING_AGENT_STORAGE() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCallerSession) CONTRACTMININGAGENTSTORAGE() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.CONTRACTMININGAGENTSTORAGE(&_AssetboxStorageImpl.CallOpts)
}

// CONTRACTMSBON is a free data retrieval call binding the contract method 0x28b3a976.
//
// Solidity: function CONTRACT_MSBON() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCaller) CONTRACTMSBON(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxStorageImpl.contract.Call(opts, out, "CONTRACT_MSBON")
	return *ret0, err
}

// CONTRACTMSBON is a free data retrieval call binding the contract method 0x28b3a976.
//
// Solidity: function CONTRACT_MSBON() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplSession) CONTRACTMSBON() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.CONTRACTMSBON(&_AssetboxStorageImpl.CallOpts)
}

// CONTRACTMSBON is a free data retrieval call binding the contract method 0x28b3a976.
//
// Solidity: function CONTRACT_MSBON() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCallerSession) CONTRACTMSBON() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.CONTRACTMSBON(&_AssetboxStorageImpl.CallOpts)
}

// CONTRACTMSBONSTORAGE is a free data retrieval call binding the contract method 0xe04e6061.
//
// Solidity: function CONTRACT_MSBON_STORAGE() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCaller) CONTRACTMSBONSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxStorageImpl.contract.Call(opts, out, "CONTRACT_MSBON_STORAGE")
	return *ret0, err
}

// CONTRACTMSBONSTORAGE is a free data retrieval call binding the contract method 0xe04e6061.
//
// Solidity: function CONTRACT_MSBON_STORAGE() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplSession) CONTRACTMSBONSTORAGE() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.CONTRACTMSBONSTORAGE(&_AssetboxStorageImpl.CallOpts)
}

// CONTRACTMSBONSTORAGE is a free data retrieval call binding the contract method 0xe04e6061.
//
// Solidity: function CONTRACT_MSBON_STORAGE() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCallerSession) CONTRACTMSBONSTORAGE() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.CONTRACTMSBONSTORAGE(&_AssetboxStorageImpl.CallOpts)
}

// CONTRACTMULTISIGADDADDRESS is a free data retrieval call binding the contract method 0xf4f65a5a.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ADDRESS() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCaller) CONTRACTMULTISIGADDADDRESS(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_ADD_ADDRESS")
	return *ret0, err
}

// CONTRACTMULTISIGADDADDRESS is a free data retrieval call binding the contract method 0xf4f65a5a.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ADDRESS() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplSession) CONTRACTMULTISIGADDADDRESS() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.CONTRACTMULTISIGADDADDRESS(&_AssetboxStorageImpl.CallOpts)
}

// CONTRACTMULTISIGADDADDRESS is a free data retrieval call binding the contract method 0xf4f65a5a.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ADDRESS() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCallerSession) CONTRACTMULTISIGADDADDRESS() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.CONTRACTMULTISIGADDADDRESS(&_AssetboxStorageImpl.CallOpts)
}

// CONTRACTMULTISIGADDADMIN is a free data retrieval call binding the contract method 0x83888c68.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ADMIN() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCaller) CONTRACTMULTISIGADDADMIN(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_ADD_ADMIN")
	return *ret0, err
}

// CONTRACTMULTISIGADDADMIN is a free data retrieval call binding the contract method 0x83888c68.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ADMIN() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplSession) CONTRACTMULTISIGADDADMIN() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.CONTRACTMULTISIGADDADMIN(&_AssetboxStorageImpl.CallOpts)
}

// CONTRACTMULTISIGADDADMIN is a free data retrieval call binding the contract method 0x83888c68.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ADMIN() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCallerSession) CONTRACTMULTISIGADDADMIN() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.CONTRACTMULTISIGADDADMIN(&_AssetboxStorageImpl.CallOpts)
}

// CONTRACTMULTISIGADDFEE is a free data retrieval call binding the contract method 0xc43e8cab.
//
// Solidity: function CONTRACT_MULTISIG_ADD_FEE() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCaller) CONTRACTMULTISIGADDFEE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_ADD_FEE")
	return *ret0, err
}

// CONTRACTMULTISIGADDFEE is a free data retrieval call binding the contract method 0xc43e8cab.
//
// Solidity: function CONTRACT_MULTISIG_ADD_FEE() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplSession) CONTRACTMULTISIGADDFEE() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.CONTRACTMULTISIGADDFEE(&_AssetboxStorageImpl.CallOpts)
}

// CONTRACTMULTISIGADDFEE is a free data retrieval call binding the contract method 0xc43e8cab.
//
// Solidity: function CONTRACT_MULTISIG_ADD_FEE() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCallerSession) CONTRACTMULTISIGADDFEE() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.CONTRACTMULTISIGADDFEE(&_AssetboxStorageImpl.CallOpts)
}

// CONTRACTMULTISIGADDROLE is a free data retrieval call binding the contract method 0x20a0914b.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ROLE() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCaller) CONTRACTMULTISIGADDROLE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_ADD_ROLE")
	return *ret0, err
}

// CONTRACTMULTISIGADDROLE is a free data retrieval call binding the contract method 0x20a0914b.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ROLE() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplSession) CONTRACTMULTISIGADDROLE() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.CONTRACTMULTISIGADDROLE(&_AssetboxStorageImpl.CallOpts)
}

// CONTRACTMULTISIGADDROLE is a free data retrieval call binding the contract method 0x20a0914b.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ROLE() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCallerSession) CONTRACTMULTISIGADDROLE() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.CONTRACTMULTISIGADDROLE(&_AssetboxStorageImpl.CallOpts)
}

// CONTRACTMULTISIGDROPPERMISSION is a free data retrieval call binding the contract method 0x3f3fbc39.
//
// Solidity: function CONTRACT_MULTISIG_DROP_PERMISSION() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCaller) CONTRACTMULTISIGDROPPERMISSION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_DROP_PERMISSION")
	return *ret0, err
}

// CONTRACTMULTISIGDROPPERMISSION is a free data retrieval call binding the contract method 0x3f3fbc39.
//
// Solidity: function CONTRACT_MULTISIG_DROP_PERMISSION() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplSession) CONTRACTMULTISIGDROPPERMISSION() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.CONTRACTMULTISIGDROPPERMISSION(&_AssetboxStorageImpl.CallOpts)
}

// CONTRACTMULTISIGDROPPERMISSION is a free data retrieval call binding the contract method 0x3f3fbc39.
//
// Solidity: function CONTRACT_MULTISIG_DROP_PERMISSION() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCallerSession) CONTRACTMULTISIGDROPPERMISSION() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.CONTRACTMULTISIGDROPPERMISSION(&_AssetboxStorageImpl.CallOpts)
}

// CONTRACTMULTISIGEDITADMIN is a free data retrieval call binding the contract method 0x91b0f051.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_ADMIN() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCaller) CONTRACTMULTISIGEDITADMIN(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_EDIT_ADMIN")
	return *ret0, err
}

// CONTRACTMULTISIGEDITADMIN is a free data retrieval call binding the contract method 0x91b0f051.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_ADMIN() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplSession) CONTRACTMULTISIGEDITADMIN() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.CONTRACTMULTISIGEDITADMIN(&_AssetboxStorageImpl.CallOpts)
}

// CONTRACTMULTISIGEDITADMIN is a free data retrieval call binding the contract method 0x91b0f051.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_ADMIN() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCallerSession) CONTRACTMULTISIGEDITADMIN() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.CONTRACTMULTISIGEDITADMIN(&_AssetboxStorageImpl.CallOpts)
}

// CONTRACTMULTISIGEDITFEE is a free data retrieval call binding the contract method 0x6b2abe17.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_FEE() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCaller) CONTRACTMULTISIGEDITFEE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_EDIT_FEE")
	return *ret0, err
}

// CONTRACTMULTISIGEDITFEE is a free data retrieval call binding the contract method 0x6b2abe17.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_FEE() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplSession) CONTRACTMULTISIGEDITFEE() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.CONTRACTMULTISIGEDITFEE(&_AssetboxStorageImpl.CallOpts)
}

// CONTRACTMULTISIGEDITFEE is a free data retrieval call binding the contract method 0x6b2abe17.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_FEE() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCallerSession) CONTRACTMULTISIGEDITFEE() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.CONTRACTMULTISIGEDITFEE(&_AssetboxStorageImpl.CallOpts)
}

// CONTRACTMULTISIGEDITROLE is a free data retrieval call binding the contract method 0xc5f66412.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_ROLE() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCaller) CONTRACTMULTISIGEDITROLE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_EDIT_ROLE")
	return *ret0, err
}

// CONTRACTMULTISIGEDITROLE is a free data retrieval call binding the contract method 0xc5f66412.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_ROLE() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplSession) CONTRACTMULTISIGEDITROLE() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.CONTRACTMULTISIGEDITROLE(&_AssetboxStorageImpl.CallOpts)
}

// CONTRACTMULTISIGEDITROLE is a free data retrieval call binding the contract method 0xc5f66412.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_ROLE() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCallerSession) CONTRACTMULTISIGEDITROLE() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.CONTRACTMULTISIGEDITROLE(&_AssetboxStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEADMINBYADDRESS is a free data retrieval call binding the contract method 0xc816b091.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_BY_ADDRESS() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCaller) CONTRACTMULTISIGREMOVEADMINBYADDRESS(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REMOVE_ADMIN_BY_ADDRESS")
	return *ret0, err
}

// CONTRACTMULTISIGREMOVEADMINBYADDRESS is a free data retrieval call binding the contract method 0xc816b091.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_BY_ADDRESS() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplSession) CONTRACTMULTISIGREMOVEADMINBYADDRESS() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.CONTRACTMULTISIGREMOVEADMINBYADDRESS(&_AssetboxStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEADMINBYADDRESS is a free data retrieval call binding the contract method 0xc816b091.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_BY_ADDRESS() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCallerSession) CONTRACTMULTISIGREMOVEADMINBYADDRESS() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.CONTRACTMULTISIGREMOVEADMINBYADDRESS(&_AssetboxStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEADMINBYINDEX is a free data retrieval call binding the contract method 0x751fbea2.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_BY_INDEX() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCaller) CONTRACTMULTISIGREMOVEADMINBYINDEX(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REMOVE_ADMIN_BY_INDEX")
	return *ret0, err
}

// CONTRACTMULTISIGREMOVEADMINBYINDEX is a free data retrieval call binding the contract method 0x751fbea2.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_BY_INDEX() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplSession) CONTRACTMULTISIGREMOVEADMINBYINDEX() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.CONTRACTMULTISIGREMOVEADMINBYINDEX(&_AssetboxStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEADMINBYINDEX is a free data retrieval call binding the contract method 0x751fbea2.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_BY_INDEX() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCallerSession) CONTRACTMULTISIGREMOVEADMINBYINDEX() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.CONTRACTMULTISIGREMOVEADMINBYINDEX(&_AssetboxStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEADMINROLE is a free data retrieval call binding the contract method 0x53c00200.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_ROLE() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCaller) CONTRACTMULTISIGREMOVEADMINROLE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REMOVE_ADMIN_ROLE")
	return *ret0, err
}

// CONTRACTMULTISIGREMOVEADMINROLE is a free data retrieval call binding the contract method 0x53c00200.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_ROLE() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplSession) CONTRACTMULTISIGREMOVEADMINROLE() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.CONTRACTMULTISIGREMOVEADMINROLE(&_AssetboxStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEADMINROLE is a free data retrieval call binding the contract method 0x53c00200.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_ROLE() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCallerSession) CONTRACTMULTISIGREMOVEADMINROLE() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.CONTRACTMULTISIGREMOVEADMINROLE(&_AssetboxStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEFEE is a free data retrieval call binding the contract method 0xdb614d25.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_FEE() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCaller) CONTRACTMULTISIGREMOVEFEE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REMOVE_FEE")
	return *ret0, err
}

// CONTRACTMULTISIGREMOVEFEE is a free data retrieval call binding the contract method 0xdb614d25.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_FEE() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplSession) CONTRACTMULTISIGREMOVEFEE() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.CONTRACTMULTISIGREMOVEFEE(&_AssetboxStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEFEE is a free data retrieval call binding the contract method 0xdb614d25.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_FEE() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCallerSession) CONTRACTMULTISIGREMOVEFEE() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.CONTRACTMULTISIGREMOVEFEE(&_AssetboxStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEROLE is a free data retrieval call binding the contract method 0x23648496.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ROLE() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCaller) CONTRACTMULTISIGREMOVEROLE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REMOVE_ROLE")
	return *ret0, err
}

// CONTRACTMULTISIGREMOVEROLE is a free data retrieval call binding the contract method 0x23648496.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ROLE() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplSession) CONTRACTMULTISIGREMOVEROLE() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.CONTRACTMULTISIGREMOVEROLE(&_AssetboxStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEROLE is a free data retrieval call binding the contract method 0x23648496.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ROLE() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCallerSession) CONTRACTMULTISIGREMOVEROLE() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.CONTRACTMULTISIGREMOVEROLE(&_AssetboxStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREPLENISHCAPITALIZATION is a free data retrieval call binding the contract method 0x62d5d523.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_CAPITALIZATION() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCaller) CONTRACTMULTISIGREPLENISHCAPITALIZATION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REPLENISH_CAPITALIZATION")
	return *ret0, err
}

// CONTRACTMULTISIGREPLENISHCAPITALIZATION is a free data retrieval call binding the contract method 0x62d5d523.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_CAPITALIZATION() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplSession) CONTRACTMULTISIGREPLENISHCAPITALIZATION() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.CONTRACTMULTISIGREPLENISHCAPITALIZATION(&_AssetboxStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREPLENISHCAPITALIZATION is a free data retrieval call binding the contract method 0x62d5d523.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_CAPITALIZATION() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCallerSession) CONTRACTMULTISIGREPLENISHCAPITALIZATION() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.CONTRACTMULTISIGREPLENISHCAPITALIZATION(&_AssetboxStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREPLENISHCERTIFICATE is a free data retrieval call binding the contract method 0xfd19b7e1.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_CERTIFICATE() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCaller) CONTRACTMULTISIGREPLENISHCERTIFICATE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REPLENISH_CERTIFICATE")
	return *ret0, err
}

// CONTRACTMULTISIGREPLENISHCERTIFICATE is a free data retrieval call binding the contract method 0xfd19b7e1.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_CERTIFICATE() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplSession) CONTRACTMULTISIGREPLENISHCERTIFICATE() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.CONTRACTMULTISIGREPLENISHCERTIFICATE(&_AssetboxStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREPLENISHCERTIFICATE is a free data retrieval call binding the contract method 0xfd19b7e1.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_CERTIFICATE() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCallerSession) CONTRACTMULTISIGREPLENISHCERTIFICATE() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.CONTRACTMULTISIGREPLENISHCERTIFICATE(&_AssetboxStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREPLENISHEMISSION is a free data retrieval call binding the contract method 0x0cb59fde.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_EMISSION() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCaller) CONTRACTMULTISIGREPLENISHEMISSION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REPLENISH_EMISSION")
	return *ret0, err
}

// CONTRACTMULTISIGREPLENISHEMISSION is a free data retrieval call binding the contract method 0x0cb59fde.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_EMISSION() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplSession) CONTRACTMULTISIGREPLENISHEMISSION() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.CONTRACTMULTISIGREPLENISHEMISSION(&_AssetboxStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREPLENISHEMISSION is a free data retrieval call binding the contract method 0x0cb59fde.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_EMISSION() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCallerSession) CONTRACTMULTISIGREPLENISHEMISSION() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.CONTRACTMULTISIGREPLENISHEMISSION(&_AssetboxStorageImpl.CallOpts)
}

// CONTRACTMULTISIGSETADMINROLE is a free data retrieval call binding the contract method 0x8f209ab7.
//
// Solidity: function CONTRACT_MULTISIG_SET_ADMIN_ROLE() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCaller) CONTRACTMULTISIGSETADMINROLE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_SET_ADMIN_ROLE")
	return *ret0, err
}

// CONTRACTMULTISIGSETADMINROLE is a free data retrieval call binding the contract method 0x8f209ab7.
//
// Solidity: function CONTRACT_MULTISIG_SET_ADMIN_ROLE() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplSession) CONTRACTMULTISIGSETADMINROLE() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.CONTRACTMULTISIGSETADMINROLE(&_AssetboxStorageImpl.CallOpts)
}

// CONTRACTMULTISIGSETADMINROLE is a free data retrieval call binding the contract method 0x8f209ab7.
//
// Solidity: function CONTRACT_MULTISIG_SET_ADMIN_ROLE() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCallerSession) CONTRACTMULTISIGSETADMINROLE() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.CONTRACTMULTISIGSETADMINROLE(&_AssetboxStorageImpl.CallOpts)
}

// CONTRACTMULTISIGSETPERMISSION is a free data retrieval call binding the contract method 0xb0e91914.
//
// Solidity: function CONTRACT_MULTISIG_SET_PERMISSION() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCaller) CONTRACTMULTISIGSETPERMISSION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_SET_PERMISSION")
	return *ret0, err
}

// CONTRACTMULTISIGSETPERMISSION is a free data retrieval call binding the contract method 0xb0e91914.
//
// Solidity: function CONTRACT_MULTISIG_SET_PERMISSION() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplSession) CONTRACTMULTISIGSETPERMISSION() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.CONTRACTMULTISIGSETPERMISSION(&_AssetboxStorageImpl.CallOpts)
}

// CONTRACTMULTISIGSETPERMISSION is a free data retrieval call binding the contract method 0xb0e91914.
//
// Solidity: function CONTRACT_MULTISIG_SET_PERMISSION() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCallerSession) CONTRACTMULTISIGSETPERMISSION() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.CONTRACTMULTISIGSETPERMISSION(&_AssetboxStorageImpl.CallOpts)
}

// CONTRACTMULTISIGSTORAGE is a free data retrieval call binding the contract method 0xe7e1a303.
//
// Solidity: function CONTRACT_MULTISIG_STORAGE() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCaller) CONTRACTMULTISIGSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_STORAGE")
	return *ret0, err
}

// CONTRACTMULTISIGSTORAGE is a free data retrieval call binding the contract method 0xe7e1a303.
//
// Solidity: function CONTRACT_MULTISIG_STORAGE() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplSession) CONTRACTMULTISIGSTORAGE() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.CONTRACTMULTISIGSTORAGE(&_AssetboxStorageImpl.CallOpts)
}

// CONTRACTMULTISIGSTORAGE is a free data retrieval call binding the contract method 0xe7e1a303.
//
// Solidity: function CONTRACT_MULTISIG_STORAGE() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCallerSession) CONTRACTMULTISIGSTORAGE() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.CONTRACTMULTISIGSTORAGE(&_AssetboxStorageImpl.CallOpts)
}

// CONTRACTOTC is a free data retrieval call binding the contract method 0x5eceac70.
//
// Solidity: function CONTRACT_OTC() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCaller) CONTRACTOTC(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxStorageImpl.contract.Call(opts, out, "CONTRACT_OTC")
	return *ret0, err
}

// CONTRACTOTC is a free data retrieval call binding the contract method 0x5eceac70.
//
// Solidity: function CONTRACT_OTC() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplSession) CONTRACTOTC() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.CONTRACTOTC(&_AssetboxStorageImpl.CallOpts)
}

// CONTRACTOTC is a free data retrieval call binding the contract method 0x5eceac70.
//
// Solidity: function CONTRACT_OTC() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCallerSession) CONTRACTOTC() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.CONTRACTOTC(&_AssetboxStorageImpl.CallOpts)
}

// CONTRACTOTCSTORAGE is a free data retrieval call binding the contract method 0x825b476f.
//
// Solidity: function CONTRACT_OTC_STORAGE() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCaller) CONTRACTOTCSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxStorageImpl.contract.Call(opts, out, "CONTRACT_OTC_STORAGE")
	return *ret0, err
}

// CONTRACTOTCSTORAGE is a free data retrieval call binding the contract method 0x825b476f.
//
// Solidity: function CONTRACT_OTC_STORAGE() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplSession) CONTRACTOTCSTORAGE() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.CONTRACTOTCSTORAGE(&_AssetboxStorageImpl.CallOpts)
}

// CONTRACTOTCSTORAGE is a free data retrieval call binding the contract method 0x825b476f.
//
// Solidity: function CONTRACT_OTC_STORAGE() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCallerSession) CONTRACTOTCSTORAGE() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.CONTRACTOTCSTORAGE(&_AssetboxStorageImpl.CallOpts)
}

// CONTRACTREPLENISHCAPITALIZATION is a free data retrieval call binding the contract method 0xff5e09d8.
//
// Solidity: function CONTRACT_REPLENISH_CAPITALIZATION() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCaller) CONTRACTREPLENISHCAPITALIZATION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxStorageImpl.contract.Call(opts, out, "CONTRACT_REPLENISH_CAPITALIZATION")
	return *ret0, err
}

// CONTRACTREPLENISHCAPITALIZATION is a free data retrieval call binding the contract method 0xff5e09d8.
//
// Solidity: function CONTRACT_REPLENISH_CAPITALIZATION() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplSession) CONTRACTREPLENISHCAPITALIZATION() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.CONTRACTREPLENISHCAPITALIZATION(&_AssetboxStorageImpl.CallOpts)
}

// CONTRACTREPLENISHCAPITALIZATION is a free data retrieval call binding the contract method 0xff5e09d8.
//
// Solidity: function CONTRACT_REPLENISH_CAPITALIZATION() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCallerSession) CONTRACTREPLENISHCAPITALIZATION() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.CONTRACTREPLENISHCAPITALIZATION(&_AssetboxStorageImpl.CallOpts)
}

// CONTRACTRESERVEDALIASSTORAGE is a free data retrieval call binding the contract method 0x70a94aa7.
//
// Solidity: function CONTRACT_RESERVED_ALIAS_STORAGE() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCaller) CONTRACTRESERVEDALIASSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxStorageImpl.contract.Call(opts, out, "CONTRACT_RESERVED_ALIAS_STORAGE")
	return *ret0, err
}

// CONTRACTRESERVEDALIASSTORAGE is a free data retrieval call binding the contract method 0x70a94aa7.
//
// Solidity: function CONTRACT_RESERVED_ALIAS_STORAGE() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplSession) CONTRACTRESERVEDALIASSTORAGE() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.CONTRACTRESERVEDALIASSTORAGE(&_AssetboxStorageImpl.CallOpts)
}

// CONTRACTRESERVEDALIASSTORAGE is a free data retrieval call binding the contract method 0x70a94aa7.
//
// Solidity: function CONTRACT_RESERVED_ALIAS_STORAGE() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCallerSession) CONTRACTRESERVEDALIASSTORAGE() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.CONTRACTRESERVEDALIASSTORAGE(&_AssetboxStorageImpl.CallOpts)
}

// CONTRACTROLESTORAGE is a free data retrieval call binding the contract method 0xb938235b.
//
// Solidity: function CONTRACT_ROLE_STORAGE() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCaller) CONTRACTROLESTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxStorageImpl.contract.Call(opts, out, "CONTRACT_ROLE_STORAGE")
	return *ret0, err
}

// CONTRACTROLESTORAGE is a free data retrieval call binding the contract method 0xb938235b.
//
// Solidity: function CONTRACT_ROLE_STORAGE() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplSession) CONTRACTROLESTORAGE() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.CONTRACTROLESTORAGE(&_AssetboxStorageImpl.CallOpts)
}

// CONTRACTROLESTORAGE is a free data retrieval call binding the contract method 0xb938235b.
//
// Solidity: function CONTRACT_ROLE_STORAGE() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCallerSession) CONTRACTROLESTORAGE() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.CONTRACTROLESTORAGE(&_AssetboxStorageImpl.CallOpts)
}

// CONTRACTSAFETRANSFERSTORAGE is a free data retrieval call binding the contract method 0xe7e45392.
//
// Solidity: function CONTRACT_SAFE_TRANSFER_STORAGE() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCaller) CONTRACTSAFETRANSFERSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxStorageImpl.contract.Call(opts, out, "CONTRACT_SAFE_TRANSFER_STORAGE")
	return *ret0, err
}

// CONTRACTSAFETRANSFERSTORAGE is a free data retrieval call binding the contract method 0xe7e45392.
//
// Solidity: function CONTRACT_SAFE_TRANSFER_STORAGE() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplSession) CONTRACTSAFETRANSFERSTORAGE() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.CONTRACTSAFETRANSFERSTORAGE(&_AssetboxStorageImpl.CallOpts)
}

// CONTRACTSAFETRANSFERSTORAGE is a free data retrieval call binding the contract method 0xe7e45392.
//
// Solidity: function CONTRACT_SAFE_TRANSFER_STORAGE() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCallerSession) CONTRACTSAFETRANSFERSTORAGE() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.CONTRACTSAFETRANSFERSTORAGE(&_AssetboxStorageImpl.CallOpts)
}

// CONTRACTTRANSFER is a free data retrieval call binding the contract method 0xef94fe38.
//
// Solidity: function CONTRACT_TRANSFER() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCaller) CONTRACTTRANSFER(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxStorageImpl.contract.Call(opts, out, "CONTRACT_TRANSFER")
	return *ret0, err
}

// CONTRACTTRANSFER is a free data retrieval call binding the contract method 0xef94fe38.
//
// Solidity: function CONTRACT_TRANSFER() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplSession) CONTRACTTRANSFER() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.CONTRACTTRANSFER(&_AssetboxStorageImpl.CallOpts)
}

// CONTRACTTRANSFER is a free data retrieval call binding the contract method 0xef94fe38.
//
// Solidity: function CONTRACT_TRANSFER() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCallerSession) CONTRACTTRANSFER() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.CONTRACTTRANSFER(&_AssetboxStorageImpl.CallOpts)
}

// CONTRACTTRANSFERFROMCAPITALIZATION is a free data retrieval call binding the contract method 0xc8f8decb.
//
// Solidity: function CONTRACT_TRANSFER_FROM_CAPITALIZATION() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCaller) CONTRACTTRANSFERFROMCAPITALIZATION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxStorageImpl.contract.Call(opts, out, "CONTRACT_TRANSFER_FROM_CAPITALIZATION")
	return *ret0, err
}

// CONTRACTTRANSFERFROMCAPITALIZATION is a free data retrieval call binding the contract method 0xc8f8decb.
//
// Solidity: function CONTRACT_TRANSFER_FROM_CAPITALIZATION() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplSession) CONTRACTTRANSFERFROMCAPITALIZATION() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.CONTRACTTRANSFERFROMCAPITALIZATION(&_AssetboxStorageImpl.CallOpts)
}

// CONTRACTTRANSFERFROMCAPITALIZATION is a free data retrieval call binding the contract method 0xc8f8decb.
//
// Solidity: function CONTRACT_TRANSFER_FROM_CAPITALIZATION() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCallerSession) CONTRACTTRANSFERFROMCAPITALIZATION() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.CONTRACTTRANSFERFROMCAPITALIZATION(&_AssetboxStorageImpl.CallOpts)
}

// PERMISSIONACCESSRESTORATION is a free data retrieval call binding the contract method 0x5c3053bf.
//
// Solidity: function PERMISSION_ACCESS_RESTORATION() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCaller) PERMISSIONACCESSRESTORATION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxStorageImpl.contract.Call(opts, out, "PERMISSION_ACCESS_RESTORATION")
	return *ret0, err
}

// PERMISSIONACCESSRESTORATION is a free data retrieval call binding the contract method 0x5c3053bf.
//
// Solidity: function PERMISSION_ACCESS_RESTORATION() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplSession) PERMISSIONACCESSRESTORATION() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.PERMISSIONACCESSRESTORATION(&_AssetboxStorageImpl.CallOpts)
}

// PERMISSIONACCESSRESTORATION is a free data retrieval call binding the contract method 0x5c3053bf.
//
// Solidity: function PERMISSION_ACCESS_RESTORATION() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCallerSession) PERMISSIONACCESSRESTORATION() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.PERMISSIONACCESSRESTORATION(&_AssetboxStorageImpl.CallOpts)
}

// PERMISSIONADMINSTORAGE is a free data retrieval call binding the contract method 0x0c9fc0db.
//
// Solidity: function PERMISSION_ADMIN_STORAGE() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCaller) PERMISSIONADMINSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxStorageImpl.contract.Call(opts, out, "PERMISSION_ADMIN_STORAGE")
	return *ret0, err
}

// PERMISSIONADMINSTORAGE is a free data retrieval call binding the contract method 0x0c9fc0db.
//
// Solidity: function PERMISSION_ADMIN_STORAGE() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplSession) PERMISSIONADMINSTORAGE() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.PERMISSIONADMINSTORAGE(&_AssetboxStorageImpl.CallOpts)
}

// PERMISSIONADMINSTORAGE is a free data retrieval call binding the contract method 0x0c9fc0db.
//
// Solidity: function PERMISSION_ADMIN_STORAGE() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCallerSession) PERMISSIONADMINSTORAGE() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.PERMISSIONADMINSTORAGE(&_AssetboxStorageImpl.CallOpts)
}

// PERMISSIONCONTRACTSSTORAGE is a free data retrieval call binding the contract method 0x1516e83b.
//
// Solidity: function PERMISSION_CONTRACTS_STORAGE() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCaller) PERMISSIONCONTRACTSSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxStorageImpl.contract.Call(opts, out, "PERMISSION_CONTRACTS_STORAGE")
	return *ret0, err
}

// PERMISSIONCONTRACTSSTORAGE is a free data retrieval call binding the contract method 0x1516e83b.
//
// Solidity: function PERMISSION_CONTRACTS_STORAGE() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplSession) PERMISSIONCONTRACTSSTORAGE() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.PERMISSIONCONTRACTSSTORAGE(&_AssetboxStorageImpl.CallOpts)
}

// PERMISSIONCONTRACTSSTORAGE is a free data retrieval call binding the contract method 0x1516e83b.
//
// Solidity: function PERMISSION_CONTRACTS_STORAGE() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCallerSession) PERMISSIONCONTRACTSSTORAGE() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.PERMISSIONCONTRACTSSTORAGE(&_AssetboxStorageImpl.CallOpts)
}

// PERMISSIONCONTRACTDEPLOY is a free data retrieval call binding the contract method 0x257a4a04.
//
// Solidity: function PERMISSION_CONTRACT_DEPLOY() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCaller) PERMISSIONCONTRACTDEPLOY(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxStorageImpl.contract.Call(opts, out, "PERMISSION_CONTRACT_DEPLOY")
	return *ret0, err
}

// PERMISSIONCONTRACTDEPLOY is a free data retrieval call binding the contract method 0x257a4a04.
//
// Solidity: function PERMISSION_CONTRACT_DEPLOY() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplSession) PERMISSIONCONTRACTDEPLOY() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.PERMISSIONCONTRACTDEPLOY(&_AssetboxStorageImpl.CallOpts)
}

// PERMISSIONCONTRACTDEPLOY is a free data retrieval call binding the contract method 0x257a4a04.
//
// Solidity: function PERMISSION_CONTRACT_DEPLOY() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCallerSession) PERMISSIONCONTRACTDEPLOY() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.PERMISSIONCONTRACTDEPLOY(&_AssetboxStorageImpl.CallOpts)
}

// PERMISSIONEMISSION is a free data retrieval call binding the contract method 0x4ccd92d1.
//
// Solidity: function PERMISSION_EMISSION() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCaller) PERMISSIONEMISSION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxStorageImpl.contract.Call(opts, out, "PERMISSION_EMISSION")
	return *ret0, err
}

// PERMISSIONEMISSION is a free data retrieval call binding the contract method 0x4ccd92d1.
//
// Solidity: function PERMISSION_EMISSION() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplSession) PERMISSIONEMISSION() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.PERMISSIONEMISSION(&_AssetboxStorageImpl.CallOpts)
}

// PERMISSIONEMISSION is a free data retrieval call binding the contract method 0x4ccd92d1.
//
// Solidity: function PERMISSION_EMISSION() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCallerSession) PERMISSIONEMISSION() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.PERMISSIONEMISSION(&_AssetboxStorageImpl.CallOpts)
}

// PERMISSIONFEESSTORAGE is a free data retrieval call binding the contract method 0xe837eb87.
//
// Solidity: function PERMISSION_FEES_STORAGE() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCaller) PERMISSIONFEESSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxStorageImpl.contract.Call(opts, out, "PERMISSION_FEES_STORAGE")
	return *ret0, err
}

// PERMISSIONFEESSTORAGE is a free data retrieval call binding the contract method 0xe837eb87.
//
// Solidity: function PERMISSION_FEES_STORAGE() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplSession) PERMISSIONFEESSTORAGE() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.PERMISSIONFEESSTORAGE(&_AssetboxStorageImpl.CallOpts)
}

// PERMISSIONFEESSTORAGE is a free data retrieval call binding the contract method 0xe837eb87.
//
// Solidity: function PERMISSION_FEES_STORAGE() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCallerSession) PERMISSIONFEESSTORAGE() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.PERMISSIONFEESSTORAGE(&_AssetboxStorageImpl.CallOpts)
}

// PERMISSIONMULTISIGSTORAGE is a free data retrieval call binding the contract method 0xf1bb7aad.
//
// Solidity: function PERMISSION_MULTISIG_STORAGE() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCaller) PERMISSIONMULTISIGSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxStorageImpl.contract.Call(opts, out, "PERMISSION_MULTISIG_STORAGE")
	return *ret0, err
}

// PERMISSIONMULTISIGSTORAGE is a free data retrieval call binding the contract method 0xf1bb7aad.
//
// Solidity: function PERMISSION_MULTISIG_STORAGE() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplSession) PERMISSIONMULTISIGSTORAGE() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.PERMISSIONMULTISIGSTORAGE(&_AssetboxStorageImpl.CallOpts)
}

// PERMISSIONMULTISIGSTORAGE is a free data retrieval call binding the contract method 0xf1bb7aad.
//
// Solidity: function PERMISSION_MULTISIG_STORAGE() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCallerSession) PERMISSIONMULTISIGSTORAGE() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.PERMISSIONMULTISIGSTORAGE(&_AssetboxStorageImpl.CallOpts)
}

// PERMISSIONROLESSTORAGE is a free data retrieval call binding the contract method 0x12a16032.
//
// Solidity: function PERMISSION_ROLES_STORAGE() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCaller) PERMISSIONROLESSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxStorageImpl.contract.Call(opts, out, "PERMISSION_ROLES_STORAGE")
	return *ret0, err
}

// PERMISSIONROLESSTORAGE is a free data retrieval call binding the contract method 0x12a16032.
//
// Solidity: function PERMISSION_ROLES_STORAGE() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplSession) PERMISSIONROLESSTORAGE() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.PERMISSIONROLESSTORAGE(&_AssetboxStorageImpl.CallOpts)
}

// PERMISSIONROLESSTORAGE is a free data retrieval call binding the contract method 0x12a16032.
//
// Solidity: function PERMISSION_ROLES_STORAGE() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCallerSession) PERMISSIONROLESSTORAGE() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.PERMISSIONROLESSTORAGE(&_AssetboxStorageImpl.CallOpts)
}

// ROLEACCESSVERIFIER is a free data retrieval call binding the contract method 0xc5405f92.
//
// Solidity: function ROLE_ACCESS_VERIFIER() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCaller) ROLEACCESSVERIFIER(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxStorageImpl.contract.Call(opts, out, "ROLE_ACCESS_VERIFIER")
	return *ret0, err
}

// ROLEACCESSVERIFIER is a free data retrieval call binding the contract method 0xc5405f92.
//
// Solidity: function ROLE_ACCESS_VERIFIER() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplSession) ROLEACCESSVERIFIER() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.ROLEACCESSVERIFIER(&_AssetboxStorageImpl.CallOpts)
}

// ROLEACCESSVERIFIER is a free data retrieval call binding the contract method 0xc5405f92.
//
// Solidity: function ROLE_ACCESS_VERIFIER() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCallerSession) ROLEACCESSVERIFIER() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.ROLEACCESSVERIFIER(&_AssetboxStorageImpl.CallOpts)
}

// ROLEBITBONISSUEVERIFIER is a free data retrieval call binding the contract method 0xf0a012a6.
//
// Solidity: function ROLE_BITBON_ISSUE_VERIFIER() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCaller) ROLEBITBONISSUEVERIFIER(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxStorageImpl.contract.Call(opts, out, "ROLE_BITBON_ISSUE_VERIFIER")
	return *ret0, err
}

// ROLEBITBONISSUEVERIFIER is a free data retrieval call binding the contract method 0xf0a012a6.
//
// Solidity: function ROLE_BITBON_ISSUE_VERIFIER() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplSession) ROLEBITBONISSUEVERIFIER() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.ROLEBITBONISSUEVERIFIER(&_AssetboxStorageImpl.CallOpts)
}

// ROLEBITBONISSUEVERIFIER is a free data retrieval call binding the contract method 0xf0a012a6.
//
// Solidity: function ROLE_BITBON_ISSUE_VERIFIER() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCallerSession) ROLEBITBONISSUEVERIFIER() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.ROLEBITBONISSUEVERIFIER(&_AssetboxStorageImpl.CallOpts)
}

// ROLECOMMISSIONVERIFIER is a free data retrieval call binding the contract method 0x8ab3de8c.
//
// Solidity: function ROLE_COMMISSION_VERIFIER() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCaller) ROLECOMMISSIONVERIFIER(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxStorageImpl.contract.Call(opts, out, "ROLE_COMMISSION_VERIFIER")
	return *ret0, err
}

// ROLECOMMISSIONVERIFIER is a free data retrieval call binding the contract method 0x8ab3de8c.
//
// Solidity: function ROLE_COMMISSION_VERIFIER() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplSession) ROLECOMMISSIONVERIFIER() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.ROLECOMMISSIONVERIFIER(&_AssetboxStorageImpl.CallOpts)
}

// ROLECOMMISSIONVERIFIER is a free data retrieval call binding the contract method 0x8ab3de8c.
//
// Solidity: function ROLE_COMMISSION_VERIFIER() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCallerSession) ROLECOMMISSIONVERIFIER() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.ROLECOMMISSIONVERIFIER(&_AssetboxStorageImpl.CallOpts)
}

// ROLEDEPLOYADMIN is a free data retrieval call binding the contract method 0x667bf6e5.
//
// Solidity: function ROLE_DEPLOY_ADMIN() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCaller) ROLEDEPLOYADMIN(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxStorageImpl.contract.Call(opts, out, "ROLE_DEPLOY_ADMIN")
	return *ret0, err
}

// ROLEDEPLOYADMIN is a free data retrieval call binding the contract method 0x667bf6e5.
//
// Solidity: function ROLE_DEPLOY_ADMIN() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplSession) ROLEDEPLOYADMIN() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.ROLEDEPLOYADMIN(&_AssetboxStorageImpl.CallOpts)
}

// ROLEDEPLOYADMIN is a free data retrieval call binding the contract method 0x667bf6e5.
//
// Solidity: function ROLE_DEPLOY_ADMIN() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCallerSession) ROLEDEPLOYADMIN() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.ROLEDEPLOYADMIN(&_AssetboxStorageImpl.CallOpts)
}

// ROLEPERMISSIONADMIN is a free data retrieval call binding the contract method 0x1b239996.
//
// Solidity: function ROLE_PERMISSION_ADMIN() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCaller) ROLEPERMISSIONADMIN(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxStorageImpl.contract.Call(opts, out, "ROLE_PERMISSION_ADMIN")
	return *ret0, err
}

// ROLEPERMISSIONADMIN is a free data retrieval call binding the contract method 0x1b239996.
//
// Solidity: function ROLE_PERMISSION_ADMIN() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplSession) ROLEPERMISSIONADMIN() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.ROLEPERMISSIONADMIN(&_AssetboxStorageImpl.CallOpts)
}

// ROLEPERMISSIONADMIN is a free data retrieval call binding the contract method 0x1b239996.
//
// Solidity: function ROLE_PERMISSION_ADMIN() view returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCallerSession) ROLEPERMISSIONADMIN() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.ROLEPERMISSIONADMIN(&_AssetboxStorageImpl.CallOpts)
}

// AliasToAssetbox is a free data retrieval call binding the contract method 0xd08800dc.
//
// Solidity: function aliasToAssetbox(bytes32 ) view returns(address)
func (_AssetboxStorageImpl *AssetboxStorageImplCaller) AliasToAssetbox(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AssetboxStorageImpl.contract.Call(opts, out, "aliasToAssetbox", arg0)
	return *ret0, err
}

// AliasToAssetbox is a free data retrieval call binding the contract method 0xd08800dc.
//
// Solidity: function aliasToAssetbox(bytes32 ) view returns(address)
func (_AssetboxStorageImpl *AssetboxStorageImplSession) AliasToAssetbox(arg0 [32]byte) (common.Address, error) {
	return _AssetboxStorageImpl.Contract.AliasToAssetbox(&_AssetboxStorageImpl.CallOpts, arg0)
}

// AliasToAssetbox is a free data retrieval call binding the contract method 0xd08800dc.
//
// Solidity: function aliasToAssetbox(bytes32 ) view returns(address)
func (_AssetboxStorageImpl *AssetboxStorageImplCallerSession) AliasToAssetbox(arg0 [32]byte) (common.Address, error) {
	return _AssetboxStorageImpl.Contract.AliasToAssetbox(&_AssetboxStorageImpl.CallOpts, arg0)
}

// AssetboxToAlias is a free data retrieval call binding the contract method 0xcb7807f8.
//
// Solidity: function assetboxToAlias(address ) view returns(string)
func (_AssetboxStorageImpl *AssetboxStorageImplCaller) AssetboxToAlias(opts *bind.CallOpts, arg0 common.Address) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _AssetboxStorageImpl.contract.Call(opts, out, "assetboxToAlias", arg0)
	return *ret0, err
}

// AssetboxToAlias is a free data retrieval call binding the contract method 0xcb7807f8.
//
// Solidity: function assetboxToAlias(address ) view returns(string)
func (_AssetboxStorageImpl *AssetboxStorageImplSession) AssetboxToAlias(arg0 common.Address) (string, error) {
	return _AssetboxStorageImpl.Contract.AssetboxToAlias(&_AssetboxStorageImpl.CallOpts, arg0)
}

// AssetboxToAlias is a free data retrieval call binding the contract method 0xcb7807f8.
//
// Solidity: function assetboxToAlias(address ) view returns(string)
func (_AssetboxStorageImpl *AssetboxStorageImplCallerSession) AssetboxToAlias(arg0 common.Address) (string, error) {
	return _AssetboxStorageImpl.Contract.AssetboxToAlias(&_AssetboxStorageImpl.CallOpts, arg0)
}

// AssetboxToIsMining is a free data retrieval call binding the contract method 0x856ac331.
//
// Solidity: function assetboxToIsMining(address ) view returns(bool)
func (_AssetboxStorageImpl *AssetboxStorageImplCaller) AssetboxToIsMining(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _AssetboxStorageImpl.contract.Call(opts, out, "assetboxToIsMining", arg0)
	return *ret0, err
}

// AssetboxToIsMining is a free data retrieval call binding the contract method 0x856ac331.
//
// Solidity: function assetboxToIsMining(address ) view returns(bool)
func (_AssetboxStorageImpl *AssetboxStorageImplSession) AssetboxToIsMining(arg0 common.Address) (bool, error) {
	return _AssetboxStorageImpl.Contract.AssetboxToIsMining(&_AssetboxStorageImpl.CallOpts, arg0)
}

// AssetboxToIsMining is a free data retrieval call binding the contract method 0x856ac331.
//
// Solidity: function assetboxToIsMining(address ) view returns(bool)
func (_AssetboxStorageImpl *AssetboxStorageImplCallerSession) AssetboxToIsMining(arg0 common.Address) (bool, error) {
	return _AssetboxStorageImpl.Contract.AssetboxToIsMining(&_AssetboxStorageImpl.CallOpts, arg0)
}

// AssetboxToOtherInfo is a free data retrieval call binding the contract method 0x4a0a23b0.
//
// Solidity: function assetboxToOtherInfo(address ) view returns(bytes)
func (_AssetboxStorageImpl *AssetboxStorageImplCaller) AssetboxToOtherInfo(opts *bind.CallOpts, arg0 common.Address) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _AssetboxStorageImpl.contract.Call(opts, out, "assetboxToOtherInfo", arg0)
	return *ret0, err
}

// AssetboxToOtherInfo is a free data retrieval call binding the contract method 0x4a0a23b0.
//
// Solidity: function assetboxToOtherInfo(address ) view returns(bytes)
func (_AssetboxStorageImpl *AssetboxStorageImplSession) AssetboxToOtherInfo(arg0 common.Address) ([]byte, error) {
	return _AssetboxStorageImpl.Contract.AssetboxToOtherInfo(&_AssetboxStorageImpl.CallOpts, arg0)
}

// AssetboxToOtherInfo is a free data retrieval call binding the contract method 0x4a0a23b0.
//
// Solidity: function assetboxToOtherInfo(address ) view returns(bytes)
func (_AssetboxStorageImpl *AssetboxStorageImplCallerSession) AssetboxToOtherInfo(arg0 common.Address) ([]byte, error) {
	return _AssetboxStorageImpl.Contract.AssetboxToOtherInfo(&_AssetboxStorageImpl.CallOpts, arg0)
}

// AssetboxToServiceId is a free data retrieval call binding the contract method 0x828f2e27.
//
// Solidity: function assetboxToServiceId(address ) view returns(bytes)
func (_AssetboxStorageImpl *AssetboxStorageImplCaller) AssetboxToServiceId(opts *bind.CallOpts, arg0 common.Address) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _AssetboxStorageImpl.contract.Call(opts, out, "assetboxToServiceId", arg0)
	return *ret0, err
}

// AssetboxToServiceId is a free data retrieval call binding the contract method 0x828f2e27.
//
// Solidity: function assetboxToServiceId(address ) view returns(bytes)
func (_AssetboxStorageImpl *AssetboxStorageImplSession) AssetboxToServiceId(arg0 common.Address) ([]byte, error) {
	return _AssetboxStorageImpl.Contract.AssetboxToServiceId(&_AssetboxStorageImpl.CallOpts, arg0)
}

// AssetboxToServiceId is a free data retrieval call binding the contract method 0x828f2e27.
//
// Solidity: function assetboxToServiceId(address ) view returns(bytes)
func (_AssetboxStorageImpl *AssetboxStorageImplCallerSession) AssetboxToServiceId(arg0 common.Address) ([]byte, error) {
	return _AssetboxStorageImpl.Contract.AssetboxToServiceId(&_AssetboxStorageImpl.CallOpts, arg0)
}

// ContractStorage is a free data retrieval call binding the contract method 0x0549b27e.
//
// Solidity: function contractStorage() view returns(address)
func (_AssetboxStorageImpl *AssetboxStorageImplCaller) ContractStorage(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AssetboxStorageImpl.contract.Call(opts, out, "contractStorage")
	return *ret0, err
}

// ContractStorage is a free data retrieval call binding the contract method 0x0549b27e.
//
// Solidity: function contractStorage() view returns(address)
func (_AssetboxStorageImpl *AssetboxStorageImplSession) ContractStorage() (common.Address, error) {
	return _AssetboxStorageImpl.Contract.ContractStorage(&_AssetboxStorageImpl.CallOpts)
}

// ContractStorage is a free data retrieval call binding the contract method 0x0549b27e.
//
// Solidity: function contractStorage() view returns(address)
func (_AssetboxStorageImpl *AssetboxStorageImplCallerSession) ContractStorage() (common.Address, error) {
	return _AssetboxStorageImpl.Contract.ContractStorage(&_AssetboxStorageImpl.CallOpts)
}

// GetAssetbox is a free data retrieval call binding the contract method 0xe27bfa37.
//
// Solidity: function getAssetbox(string aliases) view returns(address)
func (_AssetboxStorageImpl *AssetboxStorageImplCaller) GetAssetbox(opts *bind.CallOpts, aliases string) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AssetboxStorageImpl.contract.Call(opts, out, "getAssetbox", aliases)
	return *ret0, err
}

// GetAssetbox is a free data retrieval call binding the contract method 0xe27bfa37.
//
// Solidity: function getAssetbox(string aliases) view returns(address)
func (_AssetboxStorageImpl *AssetboxStorageImplSession) GetAssetbox(aliases string) (common.Address, error) {
	return _AssetboxStorageImpl.Contract.GetAssetbox(&_AssetboxStorageImpl.CallOpts, aliases)
}

// GetAssetbox is a free data retrieval call binding the contract method 0xe27bfa37.
//
// Solidity: function getAssetbox(string aliases) view returns(address)
func (_AssetboxStorageImpl *AssetboxStorageImplCallerSession) GetAssetbox(aliases string) (common.Address, error) {
	return _AssetboxStorageImpl.Contract.GetAssetbox(&_AssetboxStorageImpl.CallOpts, aliases)
}

// GetAssetboxInfoByAssetbox is a free data retrieval call binding the contract method 0x839e0fe1.
//
// Solidity: function getAssetboxInfoByAssetbox(address assetbox) view returns(address assetboxAddress, string assetboxAlias, bytes serviceId, bool isMining, bytes otherInfo)
func (_AssetboxStorageImpl *AssetboxStorageImplCaller) GetAssetboxInfoByAssetbox(opts *bind.CallOpts, assetbox common.Address) (struct {
	AssetboxAddress common.Address
	AssetboxAlias   string
	ServiceId       []byte
	IsMining        bool
	OtherInfo       []byte
}, error) {
	ret := new(struct {
		AssetboxAddress common.Address
		AssetboxAlias   string
		ServiceId       []byte
		IsMining        bool
		OtherInfo       []byte
	})
	out := ret
	err := _AssetboxStorageImpl.contract.Call(opts, out, "getAssetboxInfoByAssetbox", assetbox)
	return *ret, err
}

// GetAssetboxInfoByAssetbox is a free data retrieval call binding the contract method 0x839e0fe1.
//
// Solidity: function getAssetboxInfoByAssetbox(address assetbox) view returns(address assetboxAddress, string assetboxAlias, bytes serviceId, bool isMining, bytes otherInfo)
func (_AssetboxStorageImpl *AssetboxStorageImplSession) GetAssetboxInfoByAssetbox(assetbox common.Address) (struct {
	AssetboxAddress common.Address
	AssetboxAlias   string
	ServiceId       []byte
	IsMining        bool
	OtherInfo       []byte
}, error) {
	return _AssetboxStorageImpl.Contract.GetAssetboxInfoByAssetbox(&_AssetboxStorageImpl.CallOpts, assetbox)
}

// GetAssetboxInfoByAssetbox is a free data retrieval call binding the contract method 0x839e0fe1.
//
// Solidity: function getAssetboxInfoByAssetbox(address assetbox) view returns(address assetboxAddress, string assetboxAlias, bytes serviceId, bool isMining, bytes otherInfo)
func (_AssetboxStorageImpl *AssetboxStorageImplCallerSession) GetAssetboxInfoByAssetbox(assetbox common.Address) (struct {
	AssetboxAddress common.Address
	AssetboxAlias   string
	ServiceId       []byte
	IsMining        bool
	OtherInfo       []byte
}, error) {
	return _AssetboxStorageImpl.Contract.GetAssetboxInfoByAssetbox(&_AssetboxStorageImpl.CallOpts, assetbox)
}

// GetIsMining is a free data retrieval call binding the contract method 0x5d04614a.
//
// Solidity: function getIsMining(address assetbox) view returns(bool)
func (_AssetboxStorageImpl *AssetboxStorageImplCaller) GetIsMining(opts *bind.CallOpts, assetbox common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _AssetboxStorageImpl.contract.Call(opts, out, "getIsMining", assetbox)
	return *ret0, err
}

// GetIsMining is a free data retrieval call binding the contract method 0x5d04614a.
//
// Solidity: function getIsMining(address assetbox) view returns(bool)
func (_AssetboxStorageImpl *AssetboxStorageImplSession) GetIsMining(assetbox common.Address) (bool, error) {
	return _AssetboxStorageImpl.Contract.GetIsMining(&_AssetboxStorageImpl.CallOpts, assetbox)
}

// GetIsMining is a free data retrieval call binding the contract method 0x5d04614a.
//
// Solidity: function getIsMining(address assetbox) view returns(bool)
func (_AssetboxStorageImpl *AssetboxStorageImplCallerSession) GetIsMining(assetbox common.Address) (bool, error) {
	return _AssetboxStorageImpl.Contract.GetIsMining(&_AssetboxStorageImpl.CallOpts, assetbox)
}

// GetServiceId is a free data retrieval call binding the contract method 0x1f797f27.
//
// Solidity: function getServiceId(address assetbox) view returns(bytes)
func (_AssetboxStorageImpl *AssetboxStorageImplCaller) GetServiceId(opts *bind.CallOpts, assetbox common.Address) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _AssetboxStorageImpl.contract.Call(opts, out, "getServiceId", assetbox)
	return *ret0, err
}

// GetServiceId is a free data retrieval call binding the contract method 0x1f797f27.
//
// Solidity: function getServiceId(address assetbox) view returns(bytes)
func (_AssetboxStorageImpl *AssetboxStorageImplSession) GetServiceId(assetbox common.Address) ([]byte, error) {
	return _AssetboxStorageImpl.Contract.GetServiceId(&_AssetboxStorageImpl.CallOpts, assetbox)
}

// GetServiceId is a free data retrieval call binding the contract method 0x1f797f27.
//
// Solidity: function getServiceId(address assetbox) view returns(bytes)
func (_AssetboxStorageImpl *AssetboxStorageImplCallerSession) GetServiceId(assetbox common.Address) ([]byte, error) {
	return _AssetboxStorageImpl.Contract.GetServiceId(&_AssetboxStorageImpl.CallOpts, assetbox)
}

// GetThisContractIndex is a free data retrieval call binding the contract method 0x70c6d367.
//
// Solidity: function getThisContractIndex() pure returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCaller) GetThisContractIndex(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxStorageImpl.contract.Call(opts, out, "getThisContractIndex")
	return *ret0, err
}

// GetThisContractIndex is a free data retrieval call binding the contract method 0x70c6d367.
//
// Solidity: function getThisContractIndex() pure returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplSession) GetThisContractIndex() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.GetThisContractIndex(&_AssetboxStorageImpl.CallOpts)
}

// GetThisContractIndex is a free data retrieval call binding the contract method 0x70c6d367.
//
// Solidity: function getThisContractIndex() pure returns(uint256)
func (_AssetboxStorageImpl *AssetboxStorageImplCallerSession) GetThisContractIndex() (*big.Int, error) {
	return _AssetboxStorageImpl.Contract.GetThisContractIndex(&_AssetboxStorageImpl.CallOpts)
}

// IsSanitizedAliasTaken is a free data retrieval call binding the contract method 0xc128d587.
//
// Solidity: function isSanitizedAliasTaken(bytes32 aliasHash) view returns(bool taken)
func (_AssetboxStorageImpl *AssetboxStorageImplCaller) IsSanitizedAliasTaken(opts *bind.CallOpts, aliasHash [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _AssetboxStorageImpl.contract.Call(opts, out, "isSanitizedAliasTaken", aliasHash)
	return *ret0, err
}

// IsSanitizedAliasTaken is a free data retrieval call binding the contract method 0xc128d587.
//
// Solidity: function isSanitizedAliasTaken(bytes32 aliasHash) view returns(bool taken)
func (_AssetboxStorageImpl *AssetboxStorageImplSession) IsSanitizedAliasTaken(aliasHash [32]byte) (bool, error) {
	return _AssetboxStorageImpl.Contract.IsSanitizedAliasTaken(&_AssetboxStorageImpl.CallOpts, aliasHash)
}

// IsSanitizedAliasTaken is a free data retrieval call binding the contract method 0xc128d587.
//
// Solidity: function isSanitizedAliasTaken(bytes32 aliasHash) view returns(bool taken)
func (_AssetboxStorageImpl *AssetboxStorageImplCallerSession) IsSanitizedAliasTaken(aliasHash [32]byte) (bool, error) {
	return _AssetboxStorageImpl.Contract.IsSanitizedAliasTaken(&_AssetboxStorageImpl.CallOpts, aliasHash)
}

// DeleteAssetboxInfo is a paid mutator transaction binding the contract method 0xf843b90d.
//
// Solidity: function deleteAssetboxInfo(address assetbox) returns(bool)
func (_AssetboxStorageImpl *AssetboxStorageImplTransactor) DeleteAssetboxInfo(opts *bind.TransactOpts, assetbox common.Address) (*types.Transaction, error) {
	return _AssetboxStorageImpl.contract.Transact(opts, "deleteAssetboxInfo", assetbox)
}

// DeleteAssetboxInfo is a paid mutator transaction binding the contract method 0xf843b90d.
//
// Solidity: function deleteAssetboxInfo(address assetbox) returns(bool)
func (_AssetboxStorageImpl *AssetboxStorageImplSession) DeleteAssetboxInfo(assetbox common.Address) (*types.Transaction, error) {
	return _AssetboxStorageImpl.Contract.DeleteAssetboxInfo(&_AssetboxStorageImpl.TransactOpts, assetbox)
}

// DeleteAssetboxInfo is a paid mutator transaction binding the contract method 0xf843b90d.
//
// Solidity: function deleteAssetboxInfo(address assetbox) returns(bool)
func (_AssetboxStorageImpl *AssetboxStorageImplTransactorSession) DeleteAssetboxInfo(assetbox common.Address) (*types.Transaction, error) {
	return _AssetboxStorageImpl.Contract.DeleteAssetboxInfo(&_AssetboxStorageImpl.TransactOpts, assetbox)
}

// SetAssetboxInfo is a paid mutator transaction binding the contract method 0x6efc7a66.
//
// Solidity: function setAssetboxInfo(address assetbox, string aliasString, bytes serviceId, bool isMining, bytes otherInfo) returns(bool)
func (_AssetboxStorageImpl *AssetboxStorageImplTransactor) SetAssetboxInfo(opts *bind.TransactOpts, assetbox common.Address, aliasString string, serviceId []byte, isMining bool, otherInfo []byte) (*types.Transaction, error) {
	return _AssetboxStorageImpl.contract.Transact(opts, "setAssetboxInfo", assetbox, aliasString, serviceId, isMining, otherInfo)
}

// SetAssetboxInfo is a paid mutator transaction binding the contract method 0x6efc7a66.
//
// Solidity: function setAssetboxInfo(address assetbox, string aliasString, bytes serviceId, bool isMining, bytes otherInfo) returns(bool)
func (_AssetboxStorageImpl *AssetboxStorageImplSession) SetAssetboxInfo(assetbox common.Address, aliasString string, serviceId []byte, isMining bool, otherInfo []byte) (*types.Transaction, error) {
	return _AssetboxStorageImpl.Contract.SetAssetboxInfo(&_AssetboxStorageImpl.TransactOpts, assetbox, aliasString, serviceId, isMining, otherInfo)
}

// SetAssetboxInfo is a paid mutator transaction binding the contract method 0x6efc7a66.
//
// Solidity: function setAssetboxInfo(address assetbox, string aliasString, bytes serviceId, bool isMining, bytes otherInfo) returns(bool)
func (_AssetboxStorageImpl *AssetboxStorageImplTransactorSession) SetAssetboxInfo(assetbox common.Address, aliasString string, serviceId []byte, isMining bool, otherInfo []byte) (*types.Transaction, error) {
	return _AssetboxStorageImpl.Contract.SetAssetboxInfo(&_AssetboxStorageImpl.TransactOpts, assetbox, aliasString, serviceId, isMining, otherInfo)
}

// AssetboxStorageImplAssetboxInfoDeletedIterator is returned from FilterAssetboxInfoDeleted and is used to iterate over the raw logs and unpacked data for AssetboxInfoDeleted events raised by the AssetboxStorageImpl contract.
type AssetboxStorageImplAssetboxInfoDeletedIterator struct {
	Event *AssetboxStorageImplAssetboxInfoDeleted // Event containing the contract specifics and raw log

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
func (it *AssetboxStorageImplAssetboxInfoDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetboxStorageImplAssetboxInfoDeleted)
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
		it.Event = new(AssetboxStorageImplAssetboxInfoDeleted)
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
func (it *AssetboxStorageImplAssetboxInfoDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetboxStorageImplAssetboxInfoDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetboxStorageImplAssetboxInfoDeleted represents a AssetboxInfoDeleted event raised by the AssetboxStorageImpl contract.
type AssetboxStorageImplAssetboxInfoDeleted struct {
	Assetbox    common.Address
	AliasString string
	ServiceId   []byte
	IsMining    bool
	OtherInfo   []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAssetboxInfoDeleted is a free log retrieval operation binding the contract event 0xcb4d96ae79873aa2cf1499974681961922466fdc2868f9af67e2a7f99e57f2fa.
//
// Solidity: event AssetboxInfoDeleted(address assetbox, string aliasString, bytes serviceId, bool isMining, bytes otherInfo)
func (_AssetboxStorageImpl *AssetboxStorageImplFilterer) FilterAssetboxInfoDeleted(opts *bind.FilterOpts) (*AssetboxStorageImplAssetboxInfoDeletedIterator, error) {

	logs, sub, err := _AssetboxStorageImpl.contract.FilterLogs(opts, "AssetboxInfoDeleted")
	if err != nil {
		return nil, err
	}
	return &AssetboxStorageImplAssetboxInfoDeletedIterator{contract: _AssetboxStorageImpl.contract, event: "AssetboxInfoDeleted", logs: logs, sub: sub}, nil
}

// WatchAssetboxInfoDeleted is a free log subscription operation binding the contract event 0xcb4d96ae79873aa2cf1499974681961922466fdc2868f9af67e2a7f99e57f2fa.
//
// Solidity: event AssetboxInfoDeleted(address assetbox, string aliasString, bytes serviceId, bool isMining, bytes otherInfo)
func (_AssetboxStorageImpl *AssetboxStorageImplFilterer) WatchAssetboxInfoDeleted(opts *bind.WatchOpts, sink chan<- *AssetboxStorageImplAssetboxInfoDeleted) (event.Subscription, error) {

	logs, sub, err := _AssetboxStorageImpl.contract.WatchLogs(opts, "AssetboxInfoDeleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetboxStorageImplAssetboxInfoDeleted)
				if err := _AssetboxStorageImpl.contract.UnpackLog(event, "AssetboxInfoDeleted", log); err != nil {
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

// ParseAssetboxInfoDeleted is a log parse operation binding the contract event 0xcb4d96ae79873aa2cf1499974681961922466fdc2868f9af67e2a7f99e57f2fa.
//
// Solidity: event AssetboxInfoDeleted(address assetbox, string aliasString, bytes serviceId, bool isMining, bytes otherInfo)
func (_AssetboxStorageImpl *AssetboxStorageImplFilterer) ParseAssetboxInfoDeleted(log types.Log) (*AssetboxStorageImplAssetboxInfoDeleted, error) {
	event := new(AssetboxStorageImplAssetboxInfoDeleted)
	if err := _AssetboxStorageImpl.contract.UnpackLog(event, "AssetboxInfoDeleted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AssetboxStorageImplAssetboxInfoSetIterator is returned from FilterAssetboxInfoSet and is used to iterate over the raw logs and unpacked data for AssetboxInfoSet events raised by the AssetboxStorageImpl contract.
type AssetboxStorageImplAssetboxInfoSetIterator struct {
	Event *AssetboxStorageImplAssetboxInfoSet // Event containing the contract specifics and raw log

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
func (it *AssetboxStorageImplAssetboxInfoSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetboxStorageImplAssetboxInfoSet)
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
		it.Event = new(AssetboxStorageImplAssetboxInfoSet)
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
func (it *AssetboxStorageImplAssetboxInfoSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetboxStorageImplAssetboxInfoSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetboxStorageImplAssetboxInfoSet represents a AssetboxInfoSet event raised by the AssetboxStorageImpl contract.
type AssetboxStorageImplAssetboxInfoSet struct {
	Assetbox    common.Address
	AliasString string
	ServiceId   []byte
	IsMining    bool
	OtherInfo   []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAssetboxInfoSet is a free log retrieval operation binding the contract event 0x7965991c18e7e693baca70bc19efa1cc9c4565284f80a953d9cc540ff08d65d8.
//
// Solidity: event AssetboxInfoSet(address assetbox, string aliasString, bytes serviceId, bool isMining, bytes otherInfo)
func (_AssetboxStorageImpl *AssetboxStorageImplFilterer) FilterAssetboxInfoSet(opts *bind.FilterOpts) (*AssetboxStorageImplAssetboxInfoSetIterator, error) {

	logs, sub, err := _AssetboxStorageImpl.contract.FilterLogs(opts, "AssetboxInfoSet")
	if err != nil {
		return nil, err
	}
	return &AssetboxStorageImplAssetboxInfoSetIterator{contract: _AssetboxStorageImpl.contract, event: "AssetboxInfoSet", logs: logs, sub: sub}, nil
}

// WatchAssetboxInfoSet is a free log subscription operation binding the contract event 0x7965991c18e7e693baca70bc19efa1cc9c4565284f80a953d9cc540ff08d65d8.
//
// Solidity: event AssetboxInfoSet(address assetbox, string aliasString, bytes serviceId, bool isMining, bytes otherInfo)
func (_AssetboxStorageImpl *AssetboxStorageImplFilterer) WatchAssetboxInfoSet(opts *bind.WatchOpts, sink chan<- *AssetboxStorageImplAssetboxInfoSet) (event.Subscription, error) {

	logs, sub, err := _AssetboxStorageImpl.contract.WatchLogs(opts, "AssetboxInfoSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetboxStorageImplAssetboxInfoSet)
				if err := _AssetboxStorageImpl.contract.UnpackLog(event, "AssetboxInfoSet", log); err != nil {
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

// ParseAssetboxInfoSet is a log parse operation binding the contract event 0x7965991c18e7e693baca70bc19efa1cc9c4565284f80a953d9cc540ff08d65d8.
//
// Solidity: event AssetboxInfoSet(address assetbox, string aliasString, bytes serviceId, bool isMining, bytes otherInfo)
func (_AssetboxStorageImpl *AssetboxStorageImplFilterer) ParseAssetboxInfoSet(log types.Log) (*AssetboxStorageImplAssetboxInfoSet, error) {
	event := new(AssetboxStorageImplAssetboxInfoSet)
	if err := _AssetboxStorageImpl.contract.UnpackLog(event, "AssetboxInfoSet", log); err != nil {
		return nil, err
	}
	return event, nil
}
