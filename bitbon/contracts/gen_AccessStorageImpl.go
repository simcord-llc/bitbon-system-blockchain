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

// AccessStorageImplABI is the input ABI used to generate the binding from.
const AccessStorageImplABI = "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"userAddresses\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"userRoles\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"rolesPermission\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"contractStorageAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"AdminAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"AdminRemoved\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ACCESS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ACCESS_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ADMINABLE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ADMIN_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ASSETBOX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ASSETBOX_INFO\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ASSETBOX_INFO_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ASSETBOX_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_BALANCE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_BITBON\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_BITBON_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_BITBON_SUPPORT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_BONBALANCE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_BONTRANSFER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_DEX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_DISTRIBUTION_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_EXCHANGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_EXCHANGE_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_FEE_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_GENERATOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MINING_AGENT_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MSBON\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MSBON_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_ADD_ADDRESS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_ADD_ADMIN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_ADD_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_ADD_ROLE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_DROP_PERMISSION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_EDIT_ADMIN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_EDIT_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_EDIT_ROLE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REMOVE_ADMIN_BY_ADDRESS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REMOVE_ADMIN_BY_INDEX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REMOVE_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REMOVE_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REMOVE_ROLE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REPLENISH_CAPITALIZATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REPLENISH_CERTIFICATE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REPLENISH_EMISSION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_SET_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_SET_PERMISSION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_OTC\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_OTC_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_REPLENISH_CAPITALIZATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_RESERVED_ALIAS_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ROLE_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_SAFE_TRANSFER_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_TRANSFER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_TRANSFER_FROM_CAPITALIZATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_ACCESS_RESTORATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_ADMIN_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_CONTRACTS_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_CONTRACT_DEPLOY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_EMISSION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_FEES_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_MULTISIG_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_ROLES_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_ACCESS_VERIFIER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_BITBON_ISSUE_VERIFIER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_COMMISSION_VERIFIER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_DEPLOY_ADMIN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_PERMISSION_ADMIN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"contractStorage\",\"outputs\":[{\"internalType\":\"contractContractStorage\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"roles\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userRolesIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"users\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getThisContractIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"length\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getUserAt\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAllUsers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getPermission\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"indexOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"getUserRole\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"isUserAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"role\",\"type\":\"uint256\"}],\"name\":\"addUser\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"oldAdmin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"editUser\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"removeUser\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"removeUserAt\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"removeUserRole\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"role\",\"type\":\"uint256\"}],\"name\":\"setUserRole\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"permissions\",\"type\":\"uint256\"}],\"name\":\"addRole\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"permissions\",\"type\":\"uint256\"}],\"name\":\"editRole\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"permissionConstant\",\"type\":\"uint256\"}],\"name\":\"setPermission\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"permissionConstant\",\"type\":\"uint256\"}],\"name\":\"dropPermission\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"removeRole\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"permissionConstant\",\"type\":\"uint256\"}],\"name\":\"isPermitted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// AccessStorageImpl is an auto generated Go binding around an Ethereum contract.
type AccessStorageImpl struct {
	AccessStorageImplCaller     // Read-only binding to the contract
	AccessStorageImplTransactor // Write-only binding to the contract
	AccessStorageImplFilterer   // Log filterer for contract events
}

// AccessStorageImplCaller is an auto generated read-only Go binding around an Ethereum contract.
type AccessStorageImplCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessStorageImplTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AccessStorageImplTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessStorageImplFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AccessStorageImplFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessStorageImplSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AccessStorageImplSession struct {
	Contract     *AccessStorageImpl // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// AccessStorageImplCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AccessStorageImplCallerSession struct {
	Contract *AccessStorageImplCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// AccessStorageImplTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AccessStorageImplTransactorSession struct {
	Contract     *AccessStorageImplTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// AccessStorageImplRaw is an auto generated low-level Go binding around an Ethereum contract.
type AccessStorageImplRaw struct {
	Contract *AccessStorageImpl // Generic contract binding to access the raw methods on
}

// AccessStorageImplCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AccessStorageImplCallerRaw struct {
	Contract *AccessStorageImplCaller // Generic read-only contract binding to access the raw methods on
}

// AccessStorageImplTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AccessStorageImplTransactorRaw struct {
	Contract *AccessStorageImplTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAccessStorageImpl creates a new instance of AccessStorageImpl, bound to a specific deployed contract.
func NewAccessStorageImpl(address common.Address, backend bind.ContractBackend) (*AccessStorageImpl, error) {
	contract, err := bindAccessStorageImpl(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AccessStorageImpl{AccessStorageImplCaller: AccessStorageImplCaller{contract: contract}, AccessStorageImplTransactor: AccessStorageImplTransactor{contract: contract}, AccessStorageImplFilterer: AccessStorageImplFilterer{contract: contract}}, nil
}

// NewAccessStorageImplCaller creates a new read-only instance of AccessStorageImpl, bound to a specific deployed contract.
func NewAccessStorageImplCaller(address common.Address, caller bind.ContractCaller) (*AccessStorageImplCaller, error) {
	contract, err := bindAccessStorageImpl(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AccessStorageImplCaller{contract: contract}, nil
}

// NewAccessStorageImplTransactor creates a new write-only instance of AccessStorageImpl, bound to a specific deployed contract.
func NewAccessStorageImplTransactor(address common.Address, transactor bind.ContractTransactor) (*AccessStorageImplTransactor, error) {
	contract, err := bindAccessStorageImpl(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AccessStorageImplTransactor{contract: contract}, nil
}

// NewAccessStorageImplFilterer creates a new log filterer instance of AccessStorageImpl, bound to a specific deployed contract.
func NewAccessStorageImplFilterer(address common.Address, filterer bind.ContractFilterer) (*AccessStorageImplFilterer, error) {
	contract, err := bindAccessStorageImpl(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AccessStorageImplFilterer{contract: contract}, nil
}

// bindAccessStorageImpl binds a generic wrapper to an already deployed contract.
func bindAccessStorageImpl(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AccessStorageImplABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccessStorageImpl *AccessStorageImplRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AccessStorageImpl.Contract.AccessStorageImplCaller.contract.Call(opts, result, method, params...)
}

// Return Abi for AccessStorageImplCaller
func (_AccessStorageImpl *AccessStorageImplCaller) ABI() abi.ABI {
	return _AccessStorageImpl.contract.Abi
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccessStorageImpl *AccessStorageImplRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccessStorageImpl.Contract.AccessStorageImplTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccessStorageImpl *AccessStorageImplRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccessStorageImpl.Contract.AccessStorageImplTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccessStorageImpl *AccessStorageImplCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AccessStorageImpl.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccessStorageImpl *AccessStorageImplTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccessStorageImpl.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccessStorageImpl *AccessStorageImplTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccessStorageImpl.Contract.contract.Transact(opts, method, params...)
}

// CONTRACTACCESS is a free data retrieval call binding the contract method 0x0ac62b0c.
//
// Solidity: function CONTRACT_ACCESS() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCaller) CONTRACTACCESS(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AccessStorageImpl.contract.Call(opts, out, "CONTRACT_ACCESS")
	return *ret0, err
}

// CONTRACTACCESS is a free data retrieval call binding the contract method 0x0ac62b0c.
//
// Solidity: function CONTRACT_ACCESS() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplSession) CONTRACTACCESS() (*big.Int, error) {
	return _AccessStorageImpl.Contract.CONTRACTACCESS(&_AccessStorageImpl.CallOpts)
}

// CONTRACTACCESS is a free data retrieval call binding the contract method 0x0ac62b0c.
//
// Solidity: function CONTRACT_ACCESS() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCallerSession) CONTRACTACCESS() (*big.Int, error) {
	return _AccessStorageImpl.Contract.CONTRACTACCESS(&_AccessStorageImpl.CallOpts)
}

// CONTRACTACCESSSTORAGE is a free data retrieval call binding the contract method 0xceac9bd5.
//
// Solidity: function CONTRACT_ACCESS_STORAGE() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCaller) CONTRACTACCESSSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AccessStorageImpl.contract.Call(opts, out, "CONTRACT_ACCESS_STORAGE")
	return *ret0, err
}

// CONTRACTACCESSSTORAGE is a free data retrieval call binding the contract method 0xceac9bd5.
//
// Solidity: function CONTRACT_ACCESS_STORAGE() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplSession) CONTRACTACCESSSTORAGE() (*big.Int, error) {
	return _AccessStorageImpl.Contract.CONTRACTACCESSSTORAGE(&_AccessStorageImpl.CallOpts)
}

// CONTRACTACCESSSTORAGE is a free data retrieval call binding the contract method 0xceac9bd5.
//
// Solidity: function CONTRACT_ACCESS_STORAGE() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCallerSession) CONTRACTACCESSSTORAGE() (*big.Int, error) {
	return _AccessStorageImpl.Contract.CONTRACTACCESSSTORAGE(&_AccessStorageImpl.CallOpts)
}

// CONTRACTADMINABLE is a free data retrieval call binding the contract method 0xbbe3a78f.
//
// Solidity: function CONTRACT_ADMINABLE() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCaller) CONTRACTADMINABLE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AccessStorageImpl.contract.Call(opts, out, "CONTRACT_ADMINABLE")
	return *ret0, err
}

// CONTRACTADMINABLE is a free data retrieval call binding the contract method 0xbbe3a78f.
//
// Solidity: function CONTRACT_ADMINABLE() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplSession) CONTRACTADMINABLE() (*big.Int, error) {
	return _AccessStorageImpl.Contract.CONTRACTADMINABLE(&_AccessStorageImpl.CallOpts)
}

// CONTRACTADMINABLE is a free data retrieval call binding the contract method 0xbbe3a78f.
//
// Solidity: function CONTRACT_ADMINABLE() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCallerSession) CONTRACTADMINABLE() (*big.Int, error) {
	return _AccessStorageImpl.Contract.CONTRACTADMINABLE(&_AccessStorageImpl.CallOpts)
}

// CONTRACTADMINSTORAGE is a free data retrieval call binding the contract method 0xc9aebae9.
//
// Solidity: function CONTRACT_ADMIN_STORAGE() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCaller) CONTRACTADMINSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AccessStorageImpl.contract.Call(opts, out, "CONTRACT_ADMIN_STORAGE")
	return *ret0, err
}

// CONTRACTADMINSTORAGE is a free data retrieval call binding the contract method 0xc9aebae9.
//
// Solidity: function CONTRACT_ADMIN_STORAGE() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplSession) CONTRACTADMINSTORAGE() (*big.Int, error) {
	return _AccessStorageImpl.Contract.CONTRACTADMINSTORAGE(&_AccessStorageImpl.CallOpts)
}

// CONTRACTADMINSTORAGE is a free data retrieval call binding the contract method 0xc9aebae9.
//
// Solidity: function CONTRACT_ADMIN_STORAGE() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCallerSession) CONTRACTADMINSTORAGE() (*big.Int, error) {
	return _AccessStorageImpl.Contract.CONTRACTADMINSTORAGE(&_AccessStorageImpl.CallOpts)
}

// CONTRACTASSETBOX is a free data retrieval call binding the contract method 0x161aa17c.
//
// Solidity: function CONTRACT_ASSETBOX() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCaller) CONTRACTASSETBOX(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AccessStorageImpl.contract.Call(opts, out, "CONTRACT_ASSETBOX")
	return *ret0, err
}

// CONTRACTASSETBOX is a free data retrieval call binding the contract method 0x161aa17c.
//
// Solidity: function CONTRACT_ASSETBOX() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplSession) CONTRACTASSETBOX() (*big.Int, error) {
	return _AccessStorageImpl.Contract.CONTRACTASSETBOX(&_AccessStorageImpl.CallOpts)
}

// CONTRACTASSETBOX is a free data retrieval call binding the contract method 0x161aa17c.
//
// Solidity: function CONTRACT_ASSETBOX() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCallerSession) CONTRACTASSETBOX() (*big.Int, error) {
	return _AccessStorageImpl.Contract.CONTRACTASSETBOX(&_AccessStorageImpl.CallOpts)
}

// CONTRACTASSETBOXINFO is a free data retrieval call binding the contract method 0xb572fd3d.
//
// Solidity: function CONTRACT_ASSETBOX_INFO() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCaller) CONTRACTASSETBOXINFO(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AccessStorageImpl.contract.Call(opts, out, "CONTRACT_ASSETBOX_INFO")
	return *ret0, err
}

// CONTRACTASSETBOXINFO is a free data retrieval call binding the contract method 0xb572fd3d.
//
// Solidity: function CONTRACT_ASSETBOX_INFO() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplSession) CONTRACTASSETBOXINFO() (*big.Int, error) {
	return _AccessStorageImpl.Contract.CONTRACTASSETBOXINFO(&_AccessStorageImpl.CallOpts)
}

// CONTRACTASSETBOXINFO is a free data retrieval call binding the contract method 0xb572fd3d.
//
// Solidity: function CONTRACT_ASSETBOX_INFO() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCallerSession) CONTRACTASSETBOXINFO() (*big.Int, error) {
	return _AccessStorageImpl.Contract.CONTRACTASSETBOXINFO(&_AccessStorageImpl.CallOpts)
}

// CONTRACTASSETBOXINFOSTORAGE is a free data retrieval call binding the contract method 0xf32a115d.
//
// Solidity: function CONTRACT_ASSETBOX_INFO_STORAGE() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCaller) CONTRACTASSETBOXINFOSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AccessStorageImpl.contract.Call(opts, out, "CONTRACT_ASSETBOX_INFO_STORAGE")
	return *ret0, err
}

// CONTRACTASSETBOXINFOSTORAGE is a free data retrieval call binding the contract method 0xf32a115d.
//
// Solidity: function CONTRACT_ASSETBOX_INFO_STORAGE() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplSession) CONTRACTASSETBOXINFOSTORAGE() (*big.Int, error) {
	return _AccessStorageImpl.Contract.CONTRACTASSETBOXINFOSTORAGE(&_AccessStorageImpl.CallOpts)
}

// CONTRACTASSETBOXINFOSTORAGE is a free data retrieval call binding the contract method 0xf32a115d.
//
// Solidity: function CONTRACT_ASSETBOX_INFO_STORAGE() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCallerSession) CONTRACTASSETBOXINFOSTORAGE() (*big.Int, error) {
	return _AccessStorageImpl.Contract.CONTRACTASSETBOXINFOSTORAGE(&_AccessStorageImpl.CallOpts)
}

// CONTRACTASSETBOXSTORAGE is a free data retrieval call binding the contract method 0xfa5bdafc.
//
// Solidity: function CONTRACT_ASSETBOX_STORAGE() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCaller) CONTRACTASSETBOXSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AccessStorageImpl.contract.Call(opts, out, "CONTRACT_ASSETBOX_STORAGE")
	return *ret0, err
}

// CONTRACTASSETBOXSTORAGE is a free data retrieval call binding the contract method 0xfa5bdafc.
//
// Solidity: function CONTRACT_ASSETBOX_STORAGE() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplSession) CONTRACTASSETBOXSTORAGE() (*big.Int, error) {
	return _AccessStorageImpl.Contract.CONTRACTASSETBOXSTORAGE(&_AccessStorageImpl.CallOpts)
}

// CONTRACTASSETBOXSTORAGE is a free data retrieval call binding the contract method 0xfa5bdafc.
//
// Solidity: function CONTRACT_ASSETBOX_STORAGE() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCallerSession) CONTRACTASSETBOXSTORAGE() (*big.Int, error) {
	return _AccessStorageImpl.Contract.CONTRACTASSETBOXSTORAGE(&_AccessStorageImpl.CallOpts)
}

// CONTRACTBALANCE is a free data retrieval call binding the contract method 0xd59045f5.
//
// Solidity: function CONTRACT_BALANCE() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCaller) CONTRACTBALANCE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AccessStorageImpl.contract.Call(opts, out, "CONTRACT_BALANCE")
	return *ret0, err
}

// CONTRACTBALANCE is a free data retrieval call binding the contract method 0xd59045f5.
//
// Solidity: function CONTRACT_BALANCE() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplSession) CONTRACTBALANCE() (*big.Int, error) {
	return _AccessStorageImpl.Contract.CONTRACTBALANCE(&_AccessStorageImpl.CallOpts)
}

// CONTRACTBALANCE is a free data retrieval call binding the contract method 0xd59045f5.
//
// Solidity: function CONTRACT_BALANCE() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCallerSession) CONTRACTBALANCE() (*big.Int, error) {
	return _AccessStorageImpl.Contract.CONTRACTBALANCE(&_AccessStorageImpl.CallOpts)
}

// CONTRACTBITBON is a free data retrieval call binding the contract method 0xb6b433ab.
//
// Solidity: function CONTRACT_BITBON() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCaller) CONTRACTBITBON(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AccessStorageImpl.contract.Call(opts, out, "CONTRACT_BITBON")
	return *ret0, err
}

// CONTRACTBITBON is a free data retrieval call binding the contract method 0xb6b433ab.
//
// Solidity: function CONTRACT_BITBON() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplSession) CONTRACTBITBON() (*big.Int, error) {
	return _AccessStorageImpl.Contract.CONTRACTBITBON(&_AccessStorageImpl.CallOpts)
}

// CONTRACTBITBON is a free data retrieval call binding the contract method 0xb6b433ab.
//
// Solidity: function CONTRACT_BITBON() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCallerSession) CONTRACTBITBON() (*big.Int, error) {
	return _AccessStorageImpl.Contract.CONTRACTBITBON(&_AccessStorageImpl.CallOpts)
}

// CONTRACTBITBONSTORAGE is a free data retrieval call binding the contract method 0xe7f6c30e.
//
// Solidity: function CONTRACT_BITBON_STORAGE() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCaller) CONTRACTBITBONSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AccessStorageImpl.contract.Call(opts, out, "CONTRACT_BITBON_STORAGE")
	return *ret0, err
}

// CONTRACTBITBONSTORAGE is a free data retrieval call binding the contract method 0xe7f6c30e.
//
// Solidity: function CONTRACT_BITBON_STORAGE() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplSession) CONTRACTBITBONSTORAGE() (*big.Int, error) {
	return _AccessStorageImpl.Contract.CONTRACTBITBONSTORAGE(&_AccessStorageImpl.CallOpts)
}

// CONTRACTBITBONSTORAGE is a free data retrieval call binding the contract method 0xe7f6c30e.
//
// Solidity: function CONTRACT_BITBON_STORAGE() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCallerSession) CONTRACTBITBONSTORAGE() (*big.Int, error) {
	return _AccessStorageImpl.Contract.CONTRACTBITBONSTORAGE(&_AccessStorageImpl.CallOpts)
}

// CONTRACTBITBONSUPPORT is a free data retrieval call binding the contract method 0x20139c74.
//
// Solidity: function CONTRACT_BITBON_SUPPORT() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCaller) CONTRACTBITBONSUPPORT(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AccessStorageImpl.contract.Call(opts, out, "CONTRACT_BITBON_SUPPORT")
	return *ret0, err
}

// CONTRACTBITBONSUPPORT is a free data retrieval call binding the contract method 0x20139c74.
//
// Solidity: function CONTRACT_BITBON_SUPPORT() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplSession) CONTRACTBITBONSUPPORT() (*big.Int, error) {
	return _AccessStorageImpl.Contract.CONTRACTBITBONSUPPORT(&_AccessStorageImpl.CallOpts)
}

// CONTRACTBITBONSUPPORT is a free data retrieval call binding the contract method 0x20139c74.
//
// Solidity: function CONTRACT_BITBON_SUPPORT() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCallerSession) CONTRACTBITBONSUPPORT() (*big.Int, error) {
	return _AccessStorageImpl.Contract.CONTRACTBITBONSUPPORT(&_AccessStorageImpl.CallOpts)
}

// CONTRACTBONBALANCE is a free data retrieval call binding the contract method 0x40e02b7c.
//
// Solidity: function CONTRACT_BONBALANCE() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCaller) CONTRACTBONBALANCE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AccessStorageImpl.contract.Call(opts, out, "CONTRACT_BONBALANCE")
	return *ret0, err
}

// CONTRACTBONBALANCE is a free data retrieval call binding the contract method 0x40e02b7c.
//
// Solidity: function CONTRACT_BONBALANCE() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplSession) CONTRACTBONBALANCE() (*big.Int, error) {
	return _AccessStorageImpl.Contract.CONTRACTBONBALANCE(&_AccessStorageImpl.CallOpts)
}

// CONTRACTBONBALANCE is a free data retrieval call binding the contract method 0x40e02b7c.
//
// Solidity: function CONTRACT_BONBALANCE() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCallerSession) CONTRACTBONBALANCE() (*big.Int, error) {
	return _AccessStorageImpl.Contract.CONTRACTBONBALANCE(&_AccessStorageImpl.CallOpts)
}

// CONTRACTBONTRANSFER is a free data retrieval call binding the contract method 0xc18f4a8a.
//
// Solidity: function CONTRACT_BONTRANSFER() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCaller) CONTRACTBONTRANSFER(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AccessStorageImpl.contract.Call(opts, out, "CONTRACT_BONTRANSFER")
	return *ret0, err
}

// CONTRACTBONTRANSFER is a free data retrieval call binding the contract method 0xc18f4a8a.
//
// Solidity: function CONTRACT_BONTRANSFER() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplSession) CONTRACTBONTRANSFER() (*big.Int, error) {
	return _AccessStorageImpl.Contract.CONTRACTBONTRANSFER(&_AccessStorageImpl.CallOpts)
}

// CONTRACTBONTRANSFER is a free data retrieval call binding the contract method 0xc18f4a8a.
//
// Solidity: function CONTRACT_BONTRANSFER() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCallerSession) CONTRACTBONTRANSFER() (*big.Int, error) {
	return _AccessStorageImpl.Contract.CONTRACTBONTRANSFER(&_AccessStorageImpl.CallOpts)
}

// CONTRACTDEX is a free data retrieval call binding the contract method 0x52c1f844.
//
// Solidity: function CONTRACT_DEX() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCaller) CONTRACTDEX(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AccessStorageImpl.contract.Call(opts, out, "CONTRACT_DEX")
	return *ret0, err
}

// CONTRACTDEX is a free data retrieval call binding the contract method 0x52c1f844.
//
// Solidity: function CONTRACT_DEX() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplSession) CONTRACTDEX() (*big.Int, error) {
	return _AccessStorageImpl.Contract.CONTRACTDEX(&_AccessStorageImpl.CallOpts)
}

// CONTRACTDEX is a free data retrieval call binding the contract method 0x52c1f844.
//
// Solidity: function CONTRACT_DEX() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCallerSession) CONTRACTDEX() (*big.Int, error) {
	return _AccessStorageImpl.Contract.CONTRACTDEX(&_AccessStorageImpl.CallOpts)
}

// CONTRACTDISTRIBUTIONSTORAGE is a free data retrieval call binding the contract method 0x08097327.
//
// Solidity: function CONTRACT_DISTRIBUTION_STORAGE() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCaller) CONTRACTDISTRIBUTIONSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AccessStorageImpl.contract.Call(opts, out, "CONTRACT_DISTRIBUTION_STORAGE")
	return *ret0, err
}

// CONTRACTDISTRIBUTIONSTORAGE is a free data retrieval call binding the contract method 0x08097327.
//
// Solidity: function CONTRACT_DISTRIBUTION_STORAGE() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplSession) CONTRACTDISTRIBUTIONSTORAGE() (*big.Int, error) {
	return _AccessStorageImpl.Contract.CONTRACTDISTRIBUTIONSTORAGE(&_AccessStorageImpl.CallOpts)
}

// CONTRACTDISTRIBUTIONSTORAGE is a free data retrieval call binding the contract method 0x08097327.
//
// Solidity: function CONTRACT_DISTRIBUTION_STORAGE() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCallerSession) CONTRACTDISTRIBUTIONSTORAGE() (*big.Int, error) {
	return _AccessStorageImpl.Contract.CONTRACTDISTRIBUTIONSTORAGE(&_AccessStorageImpl.CallOpts)
}

// CONTRACTEXCHANGE is a free data retrieval call binding the contract method 0xa0e3a10e.
//
// Solidity: function CONTRACT_EXCHANGE() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCaller) CONTRACTEXCHANGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AccessStorageImpl.contract.Call(opts, out, "CONTRACT_EXCHANGE")
	return *ret0, err
}

// CONTRACTEXCHANGE is a free data retrieval call binding the contract method 0xa0e3a10e.
//
// Solidity: function CONTRACT_EXCHANGE() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplSession) CONTRACTEXCHANGE() (*big.Int, error) {
	return _AccessStorageImpl.Contract.CONTRACTEXCHANGE(&_AccessStorageImpl.CallOpts)
}

// CONTRACTEXCHANGE is a free data retrieval call binding the contract method 0xa0e3a10e.
//
// Solidity: function CONTRACT_EXCHANGE() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCallerSession) CONTRACTEXCHANGE() (*big.Int, error) {
	return _AccessStorageImpl.Contract.CONTRACTEXCHANGE(&_AccessStorageImpl.CallOpts)
}

// CONTRACTEXCHANGESTORAGE is a free data retrieval call binding the contract method 0x001570f2.
//
// Solidity: function CONTRACT_EXCHANGE_STORAGE() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCaller) CONTRACTEXCHANGESTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AccessStorageImpl.contract.Call(opts, out, "CONTRACT_EXCHANGE_STORAGE")
	return *ret0, err
}

// CONTRACTEXCHANGESTORAGE is a free data retrieval call binding the contract method 0x001570f2.
//
// Solidity: function CONTRACT_EXCHANGE_STORAGE() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplSession) CONTRACTEXCHANGESTORAGE() (*big.Int, error) {
	return _AccessStorageImpl.Contract.CONTRACTEXCHANGESTORAGE(&_AccessStorageImpl.CallOpts)
}

// CONTRACTEXCHANGESTORAGE is a free data retrieval call binding the contract method 0x001570f2.
//
// Solidity: function CONTRACT_EXCHANGE_STORAGE() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCallerSession) CONTRACTEXCHANGESTORAGE() (*big.Int, error) {
	return _AccessStorageImpl.Contract.CONTRACTEXCHANGESTORAGE(&_AccessStorageImpl.CallOpts)
}

// CONTRACTFEESTORAGE is a free data retrieval call binding the contract method 0xcd5e65a4.
//
// Solidity: function CONTRACT_FEE_STORAGE() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCaller) CONTRACTFEESTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AccessStorageImpl.contract.Call(opts, out, "CONTRACT_FEE_STORAGE")
	return *ret0, err
}

// CONTRACTFEESTORAGE is a free data retrieval call binding the contract method 0xcd5e65a4.
//
// Solidity: function CONTRACT_FEE_STORAGE() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplSession) CONTRACTFEESTORAGE() (*big.Int, error) {
	return _AccessStorageImpl.Contract.CONTRACTFEESTORAGE(&_AccessStorageImpl.CallOpts)
}

// CONTRACTFEESTORAGE is a free data retrieval call binding the contract method 0xcd5e65a4.
//
// Solidity: function CONTRACT_FEE_STORAGE() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCallerSession) CONTRACTFEESTORAGE() (*big.Int, error) {
	return _AccessStorageImpl.Contract.CONTRACTFEESTORAGE(&_AccessStorageImpl.CallOpts)
}

// CONTRACTGENERATOR is a free data retrieval call binding the contract method 0xccea2b46.
//
// Solidity: function CONTRACT_GENERATOR() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCaller) CONTRACTGENERATOR(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AccessStorageImpl.contract.Call(opts, out, "CONTRACT_GENERATOR")
	return *ret0, err
}

// CONTRACTGENERATOR is a free data retrieval call binding the contract method 0xccea2b46.
//
// Solidity: function CONTRACT_GENERATOR() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplSession) CONTRACTGENERATOR() (*big.Int, error) {
	return _AccessStorageImpl.Contract.CONTRACTGENERATOR(&_AccessStorageImpl.CallOpts)
}

// CONTRACTGENERATOR is a free data retrieval call binding the contract method 0xccea2b46.
//
// Solidity: function CONTRACT_GENERATOR() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCallerSession) CONTRACTGENERATOR() (*big.Int, error) {
	return _AccessStorageImpl.Contract.CONTRACTGENERATOR(&_AccessStorageImpl.CallOpts)
}

// CONTRACTMININGAGENTSTORAGE is a free data retrieval call binding the contract method 0x83cc588d.
//
// Solidity: function CONTRACT_MINING_AGENT_STORAGE() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCaller) CONTRACTMININGAGENTSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AccessStorageImpl.contract.Call(opts, out, "CONTRACT_MINING_AGENT_STORAGE")
	return *ret0, err
}

// CONTRACTMININGAGENTSTORAGE is a free data retrieval call binding the contract method 0x83cc588d.
//
// Solidity: function CONTRACT_MINING_AGENT_STORAGE() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplSession) CONTRACTMININGAGENTSTORAGE() (*big.Int, error) {
	return _AccessStorageImpl.Contract.CONTRACTMININGAGENTSTORAGE(&_AccessStorageImpl.CallOpts)
}

// CONTRACTMININGAGENTSTORAGE is a free data retrieval call binding the contract method 0x83cc588d.
//
// Solidity: function CONTRACT_MINING_AGENT_STORAGE() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCallerSession) CONTRACTMININGAGENTSTORAGE() (*big.Int, error) {
	return _AccessStorageImpl.Contract.CONTRACTMININGAGENTSTORAGE(&_AccessStorageImpl.CallOpts)
}

// CONTRACTMSBON is a free data retrieval call binding the contract method 0x28b3a976.
//
// Solidity: function CONTRACT_MSBON() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCaller) CONTRACTMSBON(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AccessStorageImpl.contract.Call(opts, out, "CONTRACT_MSBON")
	return *ret0, err
}

// CONTRACTMSBON is a free data retrieval call binding the contract method 0x28b3a976.
//
// Solidity: function CONTRACT_MSBON() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplSession) CONTRACTMSBON() (*big.Int, error) {
	return _AccessStorageImpl.Contract.CONTRACTMSBON(&_AccessStorageImpl.CallOpts)
}

// CONTRACTMSBON is a free data retrieval call binding the contract method 0x28b3a976.
//
// Solidity: function CONTRACT_MSBON() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCallerSession) CONTRACTMSBON() (*big.Int, error) {
	return _AccessStorageImpl.Contract.CONTRACTMSBON(&_AccessStorageImpl.CallOpts)
}

// CONTRACTMSBONSTORAGE is a free data retrieval call binding the contract method 0xe04e6061.
//
// Solidity: function CONTRACT_MSBON_STORAGE() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCaller) CONTRACTMSBONSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AccessStorageImpl.contract.Call(opts, out, "CONTRACT_MSBON_STORAGE")
	return *ret0, err
}

// CONTRACTMSBONSTORAGE is a free data retrieval call binding the contract method 0xe04e6061.
//
// Solidity: function CONTRACT_MSBON_STORAGE() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplSession) CONTRACTMSBONSTORAGE() (*big.Int, error) {
	return _AccessStorageImpl.Contract.CONTRACTMSBONSTORAGE(&_AccessStorageImpl.CallOpts)
}

// CONTRACTMSBONSTORAGE is a free data retrieval call binding the contract method 0xe04e6061.
//
// Solidity: function CONTRACT_MSBON_STORAGE() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCallerSession) CONTRACTMSBONSTORAGE() (*big.Int, error) {
	return _AccessStorageImpl.Contract.CONTRACTMSBONSTORAGE(&_AccessStorageImpl.CallOpts)
}

// CONTRACTMULTISIGADDADDRESS is a free data retrieval call binding the contract method 0xf4f65a5a.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ADDRESS() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCaller) CONTRACTMULTISIGADDADDRESS(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AccessStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_ADD_ADDRESS")
	return *ret0, err
}

// CONTRACTMULTISIGADDADDRESS is a free data retrieval call binding the contract method 0xf4f65a5a.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ADDRESS() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplSession) CONTRACTMULTISIGADDADDRESS() (*big.Int, error) {
	return _AccessStorageImpl.Contract.CONTRACTMULTISIGADDADDRESS(&_AccessStorageImpl.CallOpts)
}

// CONTRACTMULTISIGADDADDRESS is a free data retrieval call binding the contract method 0xf4f65a5a.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ADDRESS() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCallerSession) CONTRACTMULTISIGADDADDRESS() (*big.Int, error) {
	return _AccessStorageImpl.Contract.CONTRACTMULTISIGADDADDRESS(&_AccessStorageImpl.CallOpts)
}

// CONTRACTMULTISIGADDADMIN is a free data retrieval call binding the contract method 0x83888c68.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ADMIN() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCaller) CONTRACTMULTISIGADDADMIN(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AccessStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_ADD_ADMIN")
	return *ret0, err
}

// CONTRACTMULTISIGADDADMIN is a free data retrieval call binding the contract method 0x83888c68.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ADMIN() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplSession) CONTRACTMULTISIGADDADMIN() (*big.Int, error) {
	return _AccessStorageImpl.Contract.CONTRACTMULTISIGADDADMIN(&_AccessStorageImpl.CallOpts)
}

// CONTRACTMULTISIGADDADMIN is a free data retrieval call binding the contract method 0x83888c68.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ADMIN() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCallerSession) CONTRACTMULTISIGADDADMIN() (*big.Int, error) {
	return _AccessStorageImpl.Contract.CONTRACTMULTISIGADDADMIN(&_AccessStorageImpl.CallOpts)
}

// CONTRACTMULTISIGADDFEE is a free data retrieval call binding the contract method 0xc43e8cab.
//
// Solidity: function CONTRACT_MULTISIG_ADD_FEE() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCaller) CONTRACTMULTISIGADDFEE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AccessStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_ADD_FEE")
	return *ret0, err
}

// CONTRACTMULTISIGADDFEE is a free data retrieval call binding the contract method 0xc43e8cab.
//
// Solidity: function CONTRACT_MULTISIG_ADD_FEE() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplSession) CONTRACTMULTISIGADDFEE() (*big.Int, error) {
	return _AccessStorageImpl.Contract.CONTRACTMULTISIGADDFEE(&_AccessStorageImpl.CallOpts)
}

// CONTRACTMULTISIGADDFEE is a free data retrieval call binding the contract method 0xc43e8cab.
//
// Solidity: function CONTRACT_MULTISIG_ADD_FEE() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCallerSession) CONTRACTMULTISIGADDFEE() (*big.Int, error) {
	return _AccessStorageImpl.Contract.CONTRACTMULTISIGADDFEE(&_AccessStorageImpl.CallOpts)
}

// CONTRACTMULTISIGADDROLE is a free data retrieval call binding the contract method 0x20a0914b.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ROLE() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCaller) CONTRACTMULTISIGADDROLE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AccessStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_ADD_ROLE")
	return *ret0, err
}

// CONTRACTMULTISIGADDROLE is a free data retrieval call binding the contract method 0x20a0914b.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ROLE() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplSession) CONTRACTMULTISIGADDROLE() (*big.Int, error) {
	return _AccessStorageImpl.Contract.CONTRACTMULTISIGADDROLE(&_AccessStorageImpl.CallOpts)
}

// CONTRACTMULTISIGADDROLE is a free data retrieval call binding the contract method 0x20a0914b.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ROLE() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCallerSession) CONTRACTMULTISIGADDROLE() (*big.Int, error) {
	return _AccessStorageImpl.Contract.CONTRACTMULTISIGADDROLE(&_AccessStorageImpl.CallOpts)
}

// CONTRACTMULTISIGDROPPERMISSION is a free data retrieval call binding the contract method 0x3f3fbc39.
//
// Solidity: function CONTRACT_MULTISIG_DROP_PERMISSION() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCaller) CONTRACTMULTISIGDROPPERMISSION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AccessStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_DROP_PERMISSION")
	return *ret0, err
}

// CONTRACTMULTISIGDROPPERMISSION is a free data retrieval call binding the contract method 0x3f3fbc39.
//
// Solidity: function CONTRACT_MULTISIG_DROP_PERMISSION() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplSession) CONTRACTMULTISIGDROPPERMISSION() (*big.Int, error) {
	return _AccessStorageImpl.Contract.CONTRACTMULTISIGDROPPERMISSION(&_AccessStorageImpl.CallOpts)
}

// CONTRACTMULTISIGDROPPERMISSION is a free data retrieval call binding the contract method 0x3f3fbc39.
//
// Solidity: function CONTRACT_MULTISIG_DROP_PERMISSION() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCallerSession) CONTRACTMULTISIGDROPPERMISSION() (*big.Int, error) {
	return _AccessStorageImpl.Contract.CONTRACTMULTISIGDROPPERMISSION(&_AccessStorageImpl.CallOpts)
}

// CONTRACTMULTISIGEDITADMIN is a free data retrieval call binding the contract method 0x91b0f051.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_ADMIN() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCaller) CONTRACTMULTISIGEDITADMIN(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AccessStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_EDIT_ADMIN")
	return *ret0, err
}

// CONTRACTMULTISIGEDITADMIN is a free data retrieval call binding the contract method 0x91b0f051.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_ADMIN() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplSession) CONTRACTMULTISIGEDITADMIN() (*big.Int, error) {
	return _AccessStorageImpl.Contract.CONTRACTMULTISIGEDITADMIN(&_AccessStorageImpl.CallOpts)
}

// CONTRACTMULTISIGEDITADMIN is a free data retrieval call binding the contract method 0x91b0f051.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_ADMIN() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCallerSession) CONTRACTMULTISIGEDITADMIN() (*big.Int, error) {
	return _AccessStorageImpl.Contract.CONTRACTMULTISIGEDITADMIN(&_AccessStorageImpl.CallOpts)
}

// CONTRACTMULTISIGEDITFEE is a free data retrieval call binding the contract method 0x6b2abe17.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_FEE() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCaller) CONTRACTMULTISIGEDITFEE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AccessStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_EDIT_FEE")
	return *ret0, err
}

// CONTRACTMULTISIGEDITFEE is a free data retrieval call binding the contract method 0x6b2abe17.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_FEE() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplSession) CONTRACTMULTISIGEDITFEE() (*big.Int, error) {
	return _AccessStorageImpl.Contract.CONTRACTMULTISIGEDITFEE(&_AccessStorageImpl.CallOpts)
}

// CONTRACTMULTISIGEDITFEE is a free data retrieval call binding the contract method 0x6b2abe17.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_FEE() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCallerSession) CONTRACTMULTISIGEDITFEE() (*big.Int, error) {
	return _AccessStorageImpl.Contract.CONTRACTMULTISIGEDITFEE(&_AccessStorageImpl.CallOpts)
}

// CONTRACTMULTISIGEDITROLE is a free data retrieval call binding the contract method 0xc5f66412.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_ROLE() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCaller) CONTRACTMULTISIGEDITROLE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AccessStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_EDIT_ROLE")
	return *ret0, err
}

// CONTRACTMULTISIGEDITROLE is a free data retrieval call binding the contract method 0xc5f66412.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_ROLE() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplSession) CONTRACTMULTISIGEDITROLE() (*big.Int, error) {
	return _AccessStorageImpl.Contract.CONTRACTMULTISIGEDITROLE(&_AccessStorageImpl.CallOpts)
}

// CONTRACTMULTISIGEDITROLE is a free data retrieval call binding the contract method 0xc5f66412.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_ROLE() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCallerSession) CONTRACTMULTISIGEDITROLE() (*big.Int, error) {
	return _AccessStorageImpl.Contract.CONTRACTMULTISIGEDITROLE(&_AccessStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEADMINBYADDRESS is a free data retrieval call binding the contract method 0xc816b091.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_BY_ADDRESS() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCaller) CONTRACTMULTISIGREMOVEADMINBYADDRESS(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AccessStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REMOVE_ADMIN_BY_ADDRESS")
	return *ret0, err
}

// CONTRACTMULTISIGREMOVEADMINBYADDRESS is a free data retrieval call binding the contract method 0xc816b091.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_BY_ADDRESS() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplSession) CONTRACTMULTISIGREMOVEADMINBYADDRESS() (*big.Int, error) {
	return _AccessStorageImpl.Contract.CONTRACTMULTISIGREMOVEADMINBYADDRESS(&_AccessStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEADMINBYADDRESS is a free data retrieval call binding the contract method 0xc816b091.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_BY_ADDRESS() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCallerSession) CONTRACTMULTISIGREMOVEADMINBYADDRESS() (*big.Int, error) {
	return _AccessStorageImpl.Contract.CONTRACTMULTISIGREMOVEADMINBYADDRESS(&_AccessStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEADMINBYINDEX is a free data retrieval call binding the contract method 0x751fbea2.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_BY_INDEX() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCaller) CONTRACTMULTISIGREMOVEADMINBYINDEX(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AccessStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REMOVE_ADMIN_BY_INDEX")
	return *ret0, err
}

// CONTRACTMULTISIGREMOVEADMINBYINDEX is a free data retrieval call binding the contract method 0x751fbea2.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_BY_INDEX() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplSession) CONTRACTMULTISIGREMOVEADMINBYINDEX() (*big.Int, error) {
	return _AccessStorageImpl.Contract.CONTRACTMULTISIGREMOVEADMINBYINDEX(&_AccessStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEADMINBYINDEX is a free data retrieval call binding the contract method 0x751fbea2.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_BY_INDEX() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCallerSession) CONTRACTMULTISIGREMOVEADMINBYINDEX() (*big.Int, error) {
	return _AccessStorageImpl.Contract.CONTRACTMULTISIGREMOVEADMINBYINDEX(&_AccessStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEADMINROLE is a free data retrieval call binding the contract method 0x53c00200.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_ROLE() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCaller) CONTRACTMULTISIGREMOVEADMINROLE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AccessStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REMOVE_ADMIN_ROLE")
	return *ret0, err
}

// CONTRACTMULTISIGREMOVEADMINROLE is a free data retrieval call binding the contract method 0x53c00200.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_ROLE() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplSession) CONTRACTMULTISIGREMOVEADMINROLE() (*big.Int, error) {
	return _AccessStorageImpl.Contract.CONTRACTMULTISIGREMOVEADMINROLE(&_AccessStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEADMINROLE is a free data retrieval call binding the contract method 0x53c00200.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_ROLE() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCallerSession) CONTRACTMULTISIGREMOVEADMINROLE() (*big.Int, error) {
	return _AccessStorageImpl.Contract.CONTRACTMULTISIGREMOVEADMINROLE(&_AccessStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEFEE is a free data retrieval call binding the contract method 0xdb614d25.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_FEE() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCaller) CONTRACTMULTISIGREMOVEFEE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AccessStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REMOVE_FEE")
	return *ret0, err
}

// CONTRACTMULTISIGREMOVEFEE is a free data retrieval call binding the contract method 0xdb614d25.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_FEE() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplSession) CONTRACTMULTISIGREMOVEFEE() (*big.Int, error) {
	return _AccessStorageImpl.Contract.CONTRACTMULTISIGREMOVEFEE(&_AccessStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEFEE is a free data retrieval call binding the contract method 0xdb614d25.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_FEE() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCallerSession) CONTRACTMULTISIGREMOVEFEE() (*big.Int, error) {
	return _AccessStorageImpl.Contract.CONTRACTMULTISIGREMOVEFEE(&_AccessStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEROLE is a free data retrieval call binding the contract method 0x23648496.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ROLE() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCaller) CONTRACTMULTISIGREMOVEROLE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AccessStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REMOVE_ROLE")
	return *ret0, err
}

// CONTRACTMULTISIGREMOVEROLE is a free data retrieval call binding the contract method 0x23648496.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ROLE() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplSession) CONTRACTMULTISIGREMOVEROLE() (*big.Int, error) {
	return _AccessStorageImpl.Contract.CONTRACTMULTISIGREMOVEROLE(&_AccessStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEROLE is a free data retrieval call binding the contract method 0x23648496.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ROLE() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCallerSession) CONTRACTMULTISIGREMOVEROLE() (*big.Int, error) {
	return _AccessStorageImpl.Contract.CONTRACTMULTISIGREMOVEROLE(&_AccessStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREPLENISHCAPITALIZATION is a free data retrieval call binding the contract method 0x62d5d523.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_CAPITALIZATION() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCaller) CONTRACTMULTISIGREPLENISHCAPITALIZATION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AccessStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REPLENISH_CAPITALIZATION")
	return *ret0, err
}

// CONTRACTMULTISIGREPLENISHCAPITALIZATION is a free data retrieval call binding the contract method 0x62d5d523.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_CAPITALIZATION() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplSession) CONTRACTMULTISIGREPLENISHCAPITALIZATION() (*big.Int, error) {
	return _AccessStorageImpl.Contract.CONTRACTMULTISIGREPLENISHCAPITALIZATION(&_AccessStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREPLENISHCAPITALIZATION is a free data retrieval call binding the contract method 0x62d5d523.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_CAPITALIZATION() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCallerSession) CONTRACTMULTISIGREPLENISHCAPITALIZATION() (*big.Int, error) {
	return _AccessStorageImpl.Contract.CONTRACTMULTISIGREPLENISHCAPITALIZATION(&_AccessStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREPLENISHCERTIFICATE is a free data retrieval call binding the contract method 0xfd19b7e1.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_CERTIFICATE() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCaller) CONTRACTMULTISIGREPLENISHCERTIFICATE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AccessStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REPLENISH_CERTIFICATE")
	return *ret0, err
}

// CONTRACTMULTISIGREPLENISHCERTIFICATE is a free data retrieval call binding the contract method 0xfd19b7e1.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_CERTIFICATE() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplSession) CONTRACTMULTISIGREPLENISHCERTIFICATE() (*big.Int, error) {
	return _AccessStorageImpl.Contract.CONTRACTMULTISIGREPLENISHCERTIFICATE(&_AccessStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREPLENISHCERTIFICATE is a free data retrieval call binding the contract method 0xfd19b7e1.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_CERTIFICATE() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCallerSession) CONTRACTMULTISIGREPLENISHCERTIFICATE() (*big.Int, error) {
	return _AccessStorageImpl.Contract.CONTRACTMULTISIGREPLENISHCERTIFICATE(&_AccessStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREPLENISHEMISSION is a free data retrieval call binding the contract method 0x0cb59fde.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_EMISSION() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCaller) CONTRACTMULTISIGREPLENISHEMISSION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AccessStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REPLENISH_EMISSION")
	return *ret0, err
}

// CONTRACTMULTISIGREPLENISHEMISSION is a free data retrieval call binding the contract method 0x0cb59fde.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_EMISSION() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplSession) CONTRACTMULTISIGREPLENISHEMISSION() (*big.Int, error) {
	return _AccessStorageImpl.Contract.CONTRACTMULTISIGREPLENISHEMISSION(&_AccessStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREPLENISHEMISSION is a free data retrieval call binding the contract method 0x0cb59fde.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_EMISSION() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCallerSession) CONTRACTMULTISIGREPLENISHEMISSION() (*big.Int, error) {
	return _AccessStorageImpl.Contract.CONTRACTMULTISIGREPLENISHEMISSION(&_AccessStorageImpl.CallOpts)
}

// CONTRACTMULTISIGSETADMINROLE is a free data retrieval call binding the contract method 0x8f209ab7.
//
// Solidity: function CONTRACT_MULTISIG_SET_ADMIN_ROLE() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCaller) CONTRACTMULTISIGSETADMINROLE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AccessStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_SET_ADMIN_ROLE")
	return *ret0, err
}

// CONTRACTMULTISIGSETADMINROLE is a free data retrieval call binding the contract method 0x8f209ab7.
//
// Solidity: function CONTRACT_MULTISIG_SET_ADMIN_ROLE() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplSession) CONTRACTMULTISIGSETADMINROLE() (*big.Int, error) {
	return _AccessStorageImpl.Contract.CONTRACTMULTISIGSETADMINROLE(&_AccessStorageImpl.CallOpts)
}

// CONTRACTMULTISIGSETADMINROLE is a free data retrieval call binding the contract method 0x8f209ab7.
//
// Solidity: function CONTRACT_MULTISIG_SET_ADMIN_ROLE() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCallerSession) CONTRACTMULTISIGSETADMINROLE() (*big.Int, error) {
	return _AccessStorageImpl.Contract.CONTRACTMULTISIGSETADMINROLE(&_AccessStorageImpl.CallOpts)
}

// CONTRACTMULTISIGSETPERMISSION is a free data retrieval call binding the contract method 0xb0e91914.
//
// Solidity: function CONTRACT_MULTISIG_SET_PERMISSION() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCaller) CONTRACTMULTISIGSETPERMISSION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AccessStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_SET_PERMISSION")
	return *ret0, err
}

// CONTRACTMULTISIGSETPERMISSION is a free data retrieval call binding the contract method 0xb0e91914.
//
// Solidity: function CONTRACT_MULTISIG_SET_PERMISSION() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplSession) CONTRACTMULTISIGSETPERMISSION() (*big.Int, error) {
	return _AccessStorageImpl.Contract.CONTRACTMULTISIGSETPERMISSION(&_AccessStorageImpl.CallOpts)
}

// CONTRACTMULTISIGSETPERMISSION is a free data retrieval call binding the contract method 0xb0e91914.
//
// Solidity: function CONTRACT_MULTISIG_SET_PERMISSION() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCallerSession) CONTRACTMULTISIGSETPERMISSION() (*big.Int, error) {
	return _AccessStorageImpl.Contract.CONTRACTMULTISIGSETPERMISSION(&_AccessStorageImpl.CallOpts)
}

// CONTRACTMULTISIGSTORAGE is a free data retrieval call binding the contract method 0xe7e1a303.
//
// Solidity: function CONTRACT_MULTISIG_STORAGE() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCaller) CONTRACTMULTISIGSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AccessStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_STORAGE")
	return *ret0, err
}

// CONTRACTMULTISIGSTORAGE is a free data retrieval call binding the contract method 0xe7e1a303.
//
// Solidity: function CONTRACT_MULTISIG_STORAGE() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplSession) CONTRACTMULTISIGSTORAGE() (*big.Int, error) {
	return _AccessStorageImpl.Contract.CONTRACTMULTISIGSTORAGE(&_AccessStorageImpl.CallOpts)
}

// CONTRACTMULTISIGSTORAGE is a free data retrieval call binding the contract method 0xe7e1a303.
//
// Solidity: function CONTRACT_MULTISIG_STORAGE() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCallerSession) CONTRACTMULTISIGSTORAGE() (*big.Int, error) {
	return _AccessStorageImpl.Contract.CONTRACTMULTISIGSTORAGE(&_AccessStorageImpl.CallOpts)
}

// CONTRACTOTC is a free data retrieval call binding the contract method 0x5eceac70.
//
// Solidity: function CONTRACT_OTC() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCaller) CONTRACTOTC(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AccessStorageImpl.contract.Call(opts, out, "CONTRACT_OTC")
	return *ret0, err
}

// CONTRACTOTC is a free data retrieval call binding the contract method 0x5eceac70.
//
// Solidity: function CONTRACT_OTC() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplSession) CONTRACTOTC() (*big.Int, error) {
	return _AccessStorageImpl.Contract.CONTRACTOTC(&_AccessStorageImpl.CallOpts)
}

// CONTRACTOTC is a free data retrieval call binding the contract method 0x5eceac70.
//
// Solidity: function CONTRACT_OTC() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCallerSession) CONTRACTOTC() (*big.Int, error) {
	return _AccessStorageImpl.Contract.CONTRACTOTC(&_AccessStorageImpl.CallOpts)
}

// CONTRACTOTCSTORAGE is a free data retrieval call binding the contract method 0x825b476f.
//
// Solidity: function CONTRACT_OTC_STORAGE() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCaller) CONTRACTOTCSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AccessStorageImpl.contract.Call(opts, out, "CONTRACT_OTC_STORAGE")
	return *ret0, err
}

// CONTRACTOTCSTORAGE is a free data retrieval call binding the contract method 0x825b476f.
//
// Solidity: function CONTRACT_OTC_STORAGE() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplSession) CONTRACTOTCSTORAGE() (*big.Int, error) {
	return _AccessStorageImpl.Contract.CONTRACTOTCSTORAGE(&_AccessStorageImpl.CallOpts)
}

// CONTRACTOTCSTORAGE is a free data retrieval call binding the contract method 0x825b476f.
//
// Solidity: function CONTRACT_OTC_STORAGE() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCallerSession) CONTRACTOTCSTORAGE() (*big.Int, error) {
	return _AccessStorageImpl.Contract.CONTRACTOTCSTORAGE(&_AccessStorageImpl.CallOpts)
}

// CONTRACTREPLENISHCAPITALIZATION is a free data retrieval call binding the contract method 0xff5e09d8.
//
// Solidity: function CONTRACT_REPLENISH_CAPITALIZATION() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCaller) CONTRACTREPLENISHCAPITALIZATION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AccessStorageImpl.contract.Call(opts, out, "CONTRACT_REPLENISH_CAPITALIZATION")
	return *ret0, err
}

// CONTRACTREPLENISHCAPITALIZATION is a free data retrieval call binding the contract method 0xff5e09d8.
//
// Solidity: function CONTRACT_REPLENISH_CAPITALIZATION() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplSession) CONTRACTREPLENISHCAPITALIZATION() (*big.Int, error) {
	return _AccessStorageImpl.Contract.CONTRACTREPLENISHCAPITALIZATION(&_AccessStorageImpl.CallOpts)
}

// CONTRACTREPLENISHCAPITALIZATION is a free data retrieval call binding the contract method 0xff5e09d8.
//
// Solidity: function CONTRACT_REPLENISH_CAPITALIZATION() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCallerSession) CONTRACTREPLENISHCAPITALIZATION() (*big.Int, error) {
	return _AccessStorageImpl.Contract.CONTRACTREPLENISHCAPITALIZATION(&_AccessStorageImpl.CallOpts)
}

// CONTRACTRESERVEDALIASSTORAGE is a free data retrieval call binding the contract method 0x70a94aa7.
//
// Solidity: function CONTRACT_RESERVED_ALIAS_STORAGE() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCaller) CONTRACTRESERVEDALIASSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AccessStorageImpl.contract.Call(opts, out, "CONTRACT_RESERVED_ALIAS_STORAGE")
	return *ret0, err
}

// CONTRACTRESERVEDALIASSTORAGE is a free data retrieval call binding the contract method 0x70a94aa7.
//
// Solidity: function CONTRACT_RESERVED_ALIAS_STORAGE() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplSession) CONTRACTRESERVEDALIASSTORAGE() (*big.Int, error) {
	return _AccessStorageImpl.Contract.CONTRACTRESERVEDALIASSTORAGE(&_AccessStorageImpl.CallOpts)
}

// CONTRACTRESERVEDALIASSTORAGE is a free data retrieval call binding the contract method 0x70a94aa7.
//
// Solidity: function CONTRACT_RESERVED_ALIAS_STORAGE() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCallerSession) CONTRACTRESERVEDALIASSTORAGE() (*big.Int, error) {
	return _AccessStorageImpl.Contract.CONTRACTRESERVEDALIASSTORAGE(&_AccessStorageImpl.CallOpts)
}

// CONTRACTROLESTORAGE is a free data retrieval call binding the contract method 0xb938235b.
//
// Solidity: function CONTRACT_ROLE_STORAGE() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCaller) CONTRACTROLESTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AccessStorageImpl.contract.Call(opts, out, "CONTRACT_ROLE_STORAGE")
	return *ret0, err
}

// CONTRACTROLESTORAGE is a free data retrieval call binding the contract method 0xb938235b.
//
// Solidity: function CONTRACT_ROLE_STORAGE() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplSession) CONTRACTROLESTORAGE() (*big.Int, error) {
	return _AccessStorageImpl.Contract.CONTRACTROLESTORAGE(&_AccessStorageImpl.CallOpts)
}

// CONTRACTROLESTORAGE is a free data retrieval call binding the contract method 0xb938235b.
//
// Solidity: function CONTRACT_ROLE_STORAGE() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCallerSession) CONTRACTROLESTORAGE() (*big.Int, error) {
	return _AccessStorageImpl.Contract.CONTRACTROLESTORAGE(&_AccessStorageImpl.CallOpts)
}

// CONTRACTSAFETRANSFERSTORAGE is a free data retrieval call binding the contract method 0xe7e45392.
//
// Solidity: function CONTRACT_SAFE_TRANSFER_STORAGE() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCaller) CONTRACTSAFETRANSFERSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AccessStorageImpl.contract.Call(opts, out, "CONTRACT_SAFE_TRANSFER_STORAGE")
	return *ret0, err
}

// CONTRACTSAFETRANSFERSTORAGE is a free data retrieval call binding the contract method 0xe7e45392.
//
// Solidity: function CONTRACT_SAFE_TRANSFER_STORAGE() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplSession) CONTRACTSAFETRANSFERSTORAGE() (*big.Int, error) {
	return _AccessStorageImpl.Contract.CONTRACTSAFETRANSFERSTORAGE(&_AccessStorageImpl.CallOpts)
}

// CONTRACTSAFETRANSFERSTORAGE is a free data retrieval call binding the contract method 0xe7e45392.
//
// Solidity: function CONTRACT_SAFE_TRANSFER_STORAGE() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCallerSession) CONTRACTSAFETRANSFERSTORAGE() (*big.Int, error) {
	return _AccessStorageImpl.Contract.CONTRACTSAFETRANSFERSTORAGE(&_AccessStorageImpl.CallOpts)
}

// CONTRACTTRANSFER is a free data retrieval call binding the contract method 0xef94fe38.
//
// Solidity: function CONTRACT_TRANSFER() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCaller) CONTRACTTRANSFER(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AccessStorageImpl.contract.Call(opts, out, "CONTRACT_TRANSFER")
	return *ret0, err
}

// CONTRACTTRANSFER is a free data retrieval call binding the contract method 0xef94fe38.
//
// Solidity: function CONTRACT_TRANSFER() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplSession) CONTRACTTRANSFER() (*big.Int, error) {
	return _AccessStorageImpl.Contract.CONTRACTTRANSFER(&_AccessStorageImpl.CallOpts)
}

// CONTRACTTRANSFER is a free data retrieval call binding the contract method 0xef94fe38.
//
// Solidity: function CONTRACT_TRANSFER() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCallerSession) CONTRACTTRANSFER() (*big.Int, error) {
	return _AccessStorageImpl.Contract.CONTRACTTRANSFER(&_AccessStorageImpl.CallOpts)
}

// CONTRACTTRANSFERFROMCAPITALIZATION is a free data retrieval call binding the contract method 0xc8f8decb.
//
// Solidity: function CONTRACT_TRANSFER_FROM_CAPITALIZATION() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCaller) CONTRACTTRANSFERFROMCAPITALIZATION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AccessStorageImpl.contract.Call(opts, out, "CONTRACT_TRANSFER_FROM_CAPITALIZATION")
	return *ret0, err
}

// CONTRACTTRANSFERFROMCAPITALIZATION is a free data retrieval call binding the contract method 0xc8f8decb.
//
// Solidity: function CONTRACT_TRANSFER_FROM_CAPITALIZATION() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplSession) CONTRACTTRANSFERFROMCAPITALIZATION() (*big.Int, error) {
	return _AccessStorageImpl.Contract.CONTRACTTRANSFERFROMCAPITALIZATION(&_AccessStorageImpl.CallOpts)
}

// CONTRACTTRANSFERFROMCAPITALIZATION is a free data retrieval call binding the contract method 0xc8f8decb.
//
// Solidity: function CONTRACT_TRANSFER_FROM_CAPITALIZATION() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCallerSession) CONTRACTTRANSFERFROMCAPITALIZATION() (*big.Int, error) {
	return _AccessStorageImpl.Contract.CONTRACTTRANSFERFROMCAPITALIZATION(&_AccessStorageImpl.CallOpts)
}

// PERMISSIONACCESSRESTORATION is a free data retrieval call binding the contract method 0x5c3053bf.
//
// Solidity: function PERMISSION_ACCESS_RESTORATION() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCaller) PERMISSIONACCESSRESTORATION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AccessStorageImpl.contract.Call(opts, out, "PERMISSION_ACCESS_RESTORATION")
	return *ret0, err
}

// PERMISSIONACCESSRESTORATION is a free data retrieval call binding the contract method 0x5c3053bf.
//
// Solidity: function PERMISSION_ACCESS_RESTORATION() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplSession) PERMISSIONACCESSRESTORATION() (*big.Int, error) {
	return _AccessStorageImpl.Contract.PERMISSIONACCESSRESTORATION(&_AccessStorageImpl.CallOpts)
}

// PERMISSIONACCESSRESTORATION is a free data retrieval call binding the contract method 0x5c3053bf.
//
// Solidity: function PERMISSION_ACCESS_RESTORATION() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCallerSession) PERMISSIONACCESSRESTORATION() (*big.Int, error) {
	return _AccessStorageImpl.Contract.PERMISSIONACCESSRESTORATION(&_AccessStorageImpl.CallOpts)
}

// PERMISSIONADMINSTORAGE is a free data retrieval call binding the contract method 0x0c9fc0db.
//
// Solidity: function PERMISSION_ADMIN_STORAGE() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCaller) PERMISSIONADMINSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AccessStorageImpl.contract.Call(opts, out, "PERMISSION_ADMIN_STORAGE")
	return *ret0, err
}

// PERMISSIONADMINSTORAGE is a free data retrieval call binding the contract method 0x0c9fc0db.
//
// Solidity: function PERMISSION_ADMIN_STORAGE() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplSession) PERMISSIONADMINSTORAGE() (*big.Int, error) {
	return _AccessStorageImpl.Contract.PERMISSIONADMINSTORAGE(&_AccessStorageImpl.CallOpts)
}

// PERMISSIONADMINSTORAGE is a free data retrieval call binding the contract method 0x0c9fc0db.
//
// Solidity: function PERMISSION_ADMIN_STORAGE() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCallerSession) PERMISSIONADMINSTORAGE() (*big.Int, error) {
	return _AccessStorageImpl.Contract.PERMISSIONADMINSTORAGE(&_AccessStorageImpl.CallOpts)
}

// PERMISSIONCONTRACTSSTORAGE is a free data retrieval call binding the contract method 0x1516e83b.
//
// Solidity: function PERMISSION_CONTRACTS_STORAGE() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCaller) PERMISSIONCONTRACTSSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AccessStorageImpl.contract.Call(opts, out, "PERMISSION_CONTRACTS_STORAGE")
	return *ret0, err
}

// PERMISSIONCONTRACTSSTORAGE is a free data retrieval call binding the contract method 0x1516e83b.
//
// Solidity: function PERMISSION_CONTRACTS_STORAGE() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplSession) PERMISSIONCONTRACTSSTORAGE() (*big.Int, error) {
	return _AccessStorageImpl.Contract.PERMISSIONCONTRACTSSTORAGE(&_AccessStorageImpl.CallOpts)
}

// PERMISSIONCONTRACTSSTORAGE is a free data retrieval call binding the contract method 0x1516e83b.
//
// Solidity: function PERMISSION_CONTRACTS_STORAGE() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCallerSession) PERMISSIONCONTRACTSSTORAGE() (*big.Int, error) {
	return _AccessStorageImpl.Contract.PERMISSIONCONTRACTSSTORAGE(&_AccessStorageImpl.CallOpts)
}

// PERMISSIONCONTRACTDEPLOY is a free data retrieval call binding the contract method 0x257a4a04.
//
// Solidity: function PERMISSION_CONTRACT_DEPLOY() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCaller) PERMISSIONCONTRACTDEPLOY(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AccessStorageImpl.contract.Call(opts, out, "PERMISSION_CONTRACT_DEPLOY")
	return *ret0, err
}

// PERMISSIONCONTRACTDEPLOY is a free data retrieval call binding the contract method 0x257a4a04.
//
// Solidity: function PERMISSION_CONTRACT_DEPLOY() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplSession) PERMISSIONCONTRACTDEPLOY() (*big.Int, error) {
	return _AccessStorageImpl.Contract.PERMISSIONCONTRACTDEPLOY(&_AccessStorageImpl.CallOpts)
}

// PERMISSIONCONTRACTDEPLOY is a free data retrieval call binding the contract method 0x257a4a04.
//
// Solidity: function PERMISSION_CONTRACT_DEPLOY() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCallerSession) PERMISSIONCONTRACTDEPLOY() (*big.Int, error) {
	return _AccessStorageImpl.Contract.PERMISSIONCONTRACTDEPLOY(&_AccessStorageImpl.CallOpts)
}

// PERMISSIONEMISSION is a free data retrieval call binding the contract method 0x4ccd92d1.
//
// Solidity: function PERMISSION_EMISSION() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCaller) PERMISSIONEMISSION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AccessStorageImpl.contract.Call(opts, out, "PERMISSION_EMISSION")
	return *ret0, err
}

// PERMISSIONEMISSION is a free data retrieval call binding the contract method 0x4ccd92d1.
//
// Solidity: function PERMISSION_EMISSION() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplSession) PERMISSIONEMISSION() (*big.Int, error) {
	return _AccessStorageImpl.Contract.PERMISSIONEMISSION(&_AccessStorageImpl.CallOpts)
}

// PERMISSIONEMISSION is a free data retrieval call binding the contract method 0x4ccd92d1.
//
// Solidity: function PERMISSION_EMISSION() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCallerSession) PERMISSIONEMISSION() (*big.Int, error) {
	return _AccessStorageImpl.Contract.PERMISSIONEMISSION(&_AccessStorageImpl.CallOpts)
}

// PERMISSIONFEESSTORAGE is a free data retrieval call binding the contract method 0xe837eb87.
//
// Solidity: function PERMISSION_FEES_STORAGE() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCaller) PERMISSIONFEESSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AccessStorageImpl.contract.Call(opts, out, "PERMISSION_FEES_STORAGE")
	return *ret0, err
}

// PERMISSIONFEESSTORAGE is a free data retrieval call binding the contract method 0xe837eb87.
//
// Solidity: function PERMISSION_FEES_STORAGE() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplSession) PERMISSIONFEESSTORAGE() (*big.Int, error) {
	return _AccessStorageImpl.Contract.PERMISSIONFEESSTORAGE(&_AccessStorageImpl.CallOpts)
}

// PERMISSIONFEESSTORAGE is a free data retrieval call binding the contract method 0xe837eb87.
//
// Solidity: function PERMISSION_FEES_STORAGE() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCallerSession) PERMISSIONFEESSTORAGE() (*big.Int, error) {
	return _AccessStorageImpl.Contract.PERMISSIONFEESSTORAGE(&_AccessStorageImpl.CallOpts)
}

// PERMISSIONMULTISIGSTORAGE is a free data retrieval call binding the contract method 0xf1bb7aad.
//
// Solidity: function PERMISSION_MULTISIG_STORAGE() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCaller) PERMISSIONMULTISIGSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AccessStorageImpl.contract.Call(opts, out, "PERMISSION_MULTISIG_STORAGE")
	return *ret0, err
}

// PERMISSIONMULTISIGSTORAGE is a free data retrieval call binding the contract method 0xf1bb7aad.
//
// Solidity: function PERMISSION_MULTISIG_STORAGE() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplSession) PERMISSIONMULTISIGSTORAGE() (*big.Int, error) {
	return _AccessStorageImpl.Contract.PERMISSIONMULTISIGSTORAGE(&_AccessStorageImpl.CallOpts)
}

// PERMISSIONMULTISIGSTORAGE is a free data retrieval call binding the contract method 0xf1bb7aad.
//
// Solidity: function PERMISSION_MULTISIG_STORAGE() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCallerSession) PERMISSIONMULTISIGSTORAGE() (*big.Int, error) {
	return _AccessStorageImpl.Contract.PERMISSIONMULTISIGSTORAGE(&_AccessStorageImpl.CallOpts)
}

// PERMISSIONROLESSTORAGE is a free data retrieval call binding the contract method 0x12a16032.
//
// Solidity: function PERMISSION_ROLES_STORAGE() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCaller) PERMISSIONROLESSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AccessStorageImpl.contract.Call(opts, out, "PERMISSION_ROLES_STORAGE")
	return *ret0, err
}

// PERMISSIONROLESSTORAGE is a free data retrieval call binding the contract method 0x12a16032.
//
// Solidity: function PERMISSION_ROLES_STORAGE() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplSession) PERMISSIONROLESSTORAGE() (*big.Int, error) {
	return _AccessStorageImpl.Contract.PERMISSIONROLESSTORAGE(&_AccessStorageImpl.CallOpts)
}

// PERMISSIONROLESSTORAGE is a free data retrieval call binding the contract method 0x12a16032.
//
// Solidity: function PERMISSION_ROLES_STORAGE() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCallerSession) PERMISSIONROLESSTORAGE() (*big.Int, error) {
	return _AccessStorageImpl.Contract.PERMISSIONROLESSTORAGE(&_AccessStorageImpl.CallOpts)
}

// ROLEACCESSVERIFIER is a free data retrieval call binding the contract method 0xc5405f92.
//
// Solidity: function ROLE_ACCESS_VERIFIER() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCaller) ROLEACCESSVERIFIER(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AccessStorageImpl.contract.Call(opts, out, "ROLE_ACCESS_VERIFIER")
	return *ret0, err
}

// ROLEACCESSVERIFIER is a free data retrieval call binding the contract method 0xc5405f92.
//
// Solidity: function ROLE_ACCESS_VERIFIER() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplSession) ROLEACCESSVERIFIER() (*big.Int, error) {
	return _AccessStorageImpl.Contract.ROLEACCESSVERIFIER(&_AccessStorageImpl.CallOpts)
}

// ROLEACCESSVERIFIER is a free data retrieval call binding the contract method 0xc5405f92.
//
// Solidity: function ROLE_ACCESS_VERIFIER() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCallerSession) ROLEACCESSVERIFIER() (*big.Int, error) {
	return _AccessStorageImpl.Contract.ROLEACCESSVERIFIER(&_AccessStorageImpl.CallOpts)
}

// ROLEBITBONISSUEVERIFIER is a free data retrieval call binding the contract method 0xf0a012a6.
//
// Solidity: function ROLE_BITBON_ISSUE_VERIFIER() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCaller) ROLEBITBONISSUEVERIFIER(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AccessStorageImpl.contract.Call(opts, out, "ROLE_BITBON_ISSUE_VERIFIER")
	return *ret0, err
}

// ROLEBITBONISSUEVERIFIER is a free data retrieval call binding the contract method 0xf0a012a6.
//
// Solidity: function ROLE_BITBON_ISSUE_VERIFIER() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplSession) ROLEBITBONISSUEVERIFIER() (*big.Int, error) {
	return _AccessStorageImpl.Contract.ROLEBITBONISSUEVERIFIER(&_AccessStorageImpl.CallOpts)
}

// ROLEBITBONISSUEVERIFIER is a free data retrieval call binding the contract method 0xf0a012a6.
//
// Solidity: function ROLE_BITBON_ISSUE_VERIFIER() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCallerSession) ROLEBITBONISSUEVERIFIER() (*big.Int, error) {
	return _AccessStorageImpl.Contract.ROLEBITBONISSUEVERIFIER(&_AccessStorageImpl.CallOpts)
}

// ROLECOMMISSIONVERIFIER is a free data retrieval call binding the contract method 0x8ab3de8c.
//
// Solidity: function ROLE_COMMISSION_VERIFIER() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCaller) ROLECOMMISSIONVERIFIER(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AccessStorageImpl.contract.Call(opts, out, "ROLE_COMMISSION_VERIFIER")
	return *ret0, err
}

// ROLECOMMISSIONVERIFIER is a free data retrieval call binding the contract method 0x8ab3de8c.
//
// Solidity: function ROLE_COMMISSION_VERIFIER() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplSession) ROLECOMMISSIONVERIFIER() (*big.Int, error) {
	return _AccessStorageImpl.Contract.ROLECOMMISSIONVERIFIER(&_AccessStorageImpl.CallOpts)
}

// ROLECOMMISSIONVERIFIER is a free data retrieval call binding the contract method 0x8ab3de8c.
//
// Solidity: function ROLE_COMMISSION_VERIFIER() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCallerSession) ROLECOMMISSIONVERIFIER() (*big.Int, error) {
	return _AccessStorageImpl.Contract.ROLECOMMISSIONVERIFIER(&_AccessStorageImpl.CallOpts)
}

// ROLEDEPLOYADMIN is a free data retrieval call binding the contract method 0x667bf6e5.
//
// Solidity: function ROLE_DEPLOY_ADMIN() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCaller) ROLEDEPLOYADMIN(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AccessStorageImpl.contract.Call(opts, out, "ROLE_DEPLOY_ADMIN")
	return *ret0, err
}

// ROLEDEPLOYADMIN is a free data retrieval call binding the contract method 0x667bf6e5.
//
// Solidity: function ROLE_DEPLOY_ADMIN() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplSession) ROLEDEPLOYADMIN() (*big.Int, error) {
	return _AccessStorageImpl.Contract.ROLEDEPLOYADMIN(&_AccessStorageImpl.CallOpts)
}

// ROLEDEPLOYADMIN is a free data retrieval call binding the contract method 0x667bf6e5.
//
// Solidity: function ROLE_DEPLOY_ADMIN() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCallerSession) ROLEDEPLOYADMIN() (*big.Int, error) {
	return _AccessStorageImpl.Contract.ROLEDEPLOYADMIN(&_AccessStorageImpl.CallOpts)
}

// ROLEPERMISSIONADMIN is a free data retrieval call binding the contract method 0x1b239996.
//
// Solidity: function ROLE_PERMISSION_ADMIN() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCaller) ROLEPERMISSIONADMIN(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AccessStorageImpl.contract.Call(opts, out, "ROLE_PERMISSION_ADMIN")
	return *ret0, err
}

// ROLEPERMISSIONADMIN is a free data retrieval call binding the contract method 0x1b239996.
//
// Solidity: function ROLE_PERMISSION_ADMIN() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplSession) ROLEPERMISSIONADMIN() (*big.Int, error) {
	return _AccessStorageImpl.Contract.ROLEPERMISSIONADMIN(&_AccessStorageImpl.CallOpts)
}

// ROLEPERMISSIONADMIN is a free data retrieval call binding the contract method 0x1b239996.
//
// Solidity: function ROLE_PERMISSION_ADMIN() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCallerSession) ROLEPERMISSIONADMIN() (*big.Int, error) {
	return _AccessStorageImpl.Contract.ROLEPERMISSIONADMIN(&_AccessStorageImpl.CallOpts)
}

// ContractStorage is a free data retrieval call binding the contract method 0x0549b27e.
//
// Solidity: function contractStorage() view returns(address)
func (_AccessStorageImpl *AccessStorageImplCaller) ContractStorage(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AccessStorageImpl.contract.Call(opts, out, "contractStorage")
	return *ret0, err
}

// ContractStorage is a free data retrieval call binding the contract method 0x0549b27e.
//
// Solidity: function contractStorage() view returns(address)
func (_AccessStorageImpl *AccessStorageImplSession) ContractStorage() (common.Address, error) {
	return _AccessStorageImpl.Contract.ContractStorage(&_AccessStorageImpl.CallOpts)
}

// ContractStorage is a free data retrieval call binding the contract method 0x0549b27e.
//
// Solidity: function contractStorage() view returns(address)
func (_AccessStorageImpl *AccessStorageImplCallerSession) ContractStorage() (common.Address, error) {
	return _AccessStorageImpl.Contract.ContractStorage(&_AccessStorageImpl.CallOpts)
}

// GetAllUsers is a free data retrieval call binding the contract method 0xe2842d79.
//
// Solidity: function getAllUsers() view returns(address[])
func (_AccessStorageImpl *AccessStorageImplCaller) GetAllUsers(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _AccessStorageImpl.contract.Call(opts, out, "getAllUsers")
	return *ret0, err
}

// GetAllUsers is a free data retrieval call binding the contract method 0xe2842d79.
//
// Solidity: function getAllUsers() view returns(address[])
func (_AccessStorageImpl *AccessStorageImplSession) GetAllUsers() ([]common.Address, error) {
	return _AccessStorageImpl.Contract.GetAllUsers(&_AccessStorageImpl.CallOpts)
}

// GetAllUsers is a free data retrieval call binding the contract method 0xe2842d79.
//
// Solidity: function getAllUsers() view returns(address[])
func (_AccessStorageImpl *AccessStorageImplCallerSession) GetAllUsers() ([]common.Address, error) {
	return _AccessStorageImpl.Contract.GetAllUsers(&_AccessStorageImpl.CallOpts)
}

// GetPermission is a free data retrieval call binding the contract method 0xffb5e276.
//
// Solidity: function getPermission(uint256 id) view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCaller) GetPermission(opts *bind.CallOpts, id *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AccessStorageImpl.contract.Call(opts, out, "getPermission", id)
	return *ret0, err
}

// GetPermission is a free data retrieval call binding the contract method 0xffb5e276.
//
// Solidity: function getPermission(uint256 id) view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplSession) GetPermission(id *big.Int) (*big.Int, error) {
	return _AccessStorageImpl.Contract.GetPermission(&_AccessStorageImpl.CallOpts, id)
}

// GetPermission is a free data retrieval call binding the contract method 0xffb5e276.
//
// Solidity: function getPermission(uint256 id) view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCallerSession) GetPermission(id *big.Int) (*big.Int, error) {
	return _AccessStorageImpl.Contract.GetPermission(&_AccessStorageImpl.CallOpts, id)
}

// GetThisContractIndex is a free data retrieval call binding the contract method 0x70c6d367.
//
// Solidity: function getThisContractIndex() pure returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCaller) GetThisContractIndex(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AccessStorageImpl.contract.Call(opts, out, "getThisContractIndex")
	return *ret0, err
}

// GetThisContractIndex is a free data retrieval call binding the contract method 0x70c6d367.
//
// Solidity: function getThisContractIndex() pure returns(uint256)
func (_AccessStorageImpl *AccessStorageImplSession) GetThisContractIndex() (*big.Int, error) {
	return _AccessStorageImpl.Contract.GetThisContractIndex(&_AccessStorageImpl.CallOpts)
}

// GetThisContractIndex is a free data retrieval call binding the contract method 0x70c6d367.
//
// Solidity: function getThisContractIndex() pure returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCallerSession) GetThisContractIndex() (*big.Int, error) {
	return _AccessStorageImpl.Contract.GetThisContractIndex(&_AccessStorageImpl.CallOpts)
}

// GetUserAt is a free data retrieval call binding the contract method 0xc6e064ad.
//
// Solidity: function getUserAt(uint256 index) view returns(address)
func (_AccessStorageImpl *AccessStorageImplCaller) GetUserAt(opts *bind.CallOpts, index *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AccessStorageImpl.contract.Call(opts, out, "getUserAt", index)
	return *ret0, err
}

// GetUserAt is a free data retrieval call binding the contract method 0xc6e064ad.
//
// Solidity: function getUserAt(uint256 index) view returns(address)
func (_AccessStorageImpl *AccessStorageImplSession) GetUserAt(index *big.Int) (common.Address, error) {
	return _AccessStorageImpl.Contract.GetUserAt(&_AccessStorageImpl.CallOpts, index)
}

// GetUserAt is a free data retrieval call binding the contract method 0xc6e064ad.
//
// Solidity: function getUserAt(uint256 index) view returns(address)
func (_AccessStorageImpl *AccessStorageImplCallerSession) GetUserAt(index *big.Int) (common.Address, error) {
	return _AccessStorageImpl.Contract.GetUserAt(&_AccessStorageImpl.CallOpts, index)
}

// GetUserRole is a free data retrieval call binding the contract method 0x27820851.
//
// Solidity: function getUserRole(address admin) view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCaller) GetUserRole(opts *bind.CallOpts, admin common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AccessStorageImpl.contract.Call(opts, out, "getUserRole", admin)
	return *ret0, err
}

// GetUserRole is a free data retrieval call binding the contract method 0x27820851.
//
// Solidity: function getUserRole(address admin) view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplSession) GetUserRole(admin common.Address) (*big.Int, error) {
	return _AccessStorageImpl.Contract.GetUserRole(&_AccessStorageImpl.CallOpts, admin)
}

// GetUserRole is a free data retrieval call binding the contract method 0x27820851.
//
// Solidity: function getUserRole(address admin) view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCallerSession) GetUserRole(admin common.Address) (*big.Int, error) {
	return _AccessStorageImpl.Contract.GetUserRole(&_AccessStorageImpl.CallOpts, admin)
}

// IndexOf is a free data retrieval call binding the contract method 0xfd6aad25.
//
// Solidity: function indexOf(address user) view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCaller) IndexOf(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AccessStorageImpl.contract.Call(opts, out, "indexOf", user)
	return *ret0, err
}

// IndexOf is a free data retrieval call binding the contract method 0xfd6aad25.
//
// Solidity: function indexOf(address user) view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplSession) IndexOf(user common.Address) (*big.Int, error) {
	return _AccessStorageImpl.Contract.IndexOf(&_AccessStorageImpl.CallOpts, user)
}

// IndexOf is a free data retrieval call binding the contract method 0xfd6aad25.
//
// Solidity: function indexOf(address user) view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCallerSession) IndexOf(user common.Address) (*big.Int, error) {
	return _AccessStorageImpl.Contract.IndexOf(&_AccessStorageImpl.CallOpts, user)
}

// IsPermitted is a free data retrieval call binding the contract method 0x41ec69d3.
//
// Solidity: function isPermitted(uint256 id, uint256 permissionConstant) view returns(bool)
func (_AccessStorageImpl *AccessStorageImplCaller) IsPermitted(opts *bind.CallOpts, id *big.Int, permissionConstant *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _AccessStorageImpl.contract.Call(opts, out, "isPermitted", id, permissionConstant)
	return *ret0, err
}

// IsPermitted is a free data retrieval call binding the contract method 0x41ec69d3.
//
// Solidity: function isPermitted(uint256 id, uint256 permissionConstant) view returns(bool)
func (_AccessStorageImpl *AccessStorageImplSession) IsPermitted(id *big.Int, permissionConstant *big.Int) (bool, error) {
	return _AccessStorageImpl.Contract.IsPermitted(&_AccessStorageImpl.CallOpts, id, permissionConstant)
}

// IsPermitted is a free data retrieval call binding the contract method 0x41ec69d3.
//
// Solidity: function isPermitted(uint256 id, uint256 permissionConstant) view returns(bool)
func (_AccessStorageImpl *AccessStorageImplCallerSession) IsPermitted(id *big.Int, permissionConstant *big.Int) (bool, error) {
	return _AccessStorageImpl.Contract.IsPermitted(&_AccessStorageImpl.CallOpts, id, permissionConstant)
}

// IsUserAdmin is a free data retrieval call binding the contract method 0x74d523a8.
//
// Solidity: function isUserAdmin(address user) view returns(bool)
func (_AccessStorageImpl *AccessStorageImplCaller) IsUserAdmin(opts *bind.CallOpts, user common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _AccessStorageImpl.contract.Call(opts, out, "isUserAdmin", user)
	return *ret0, err
}

// IsUserAdmin is a free data retrieval call binding the contract method 0x74d523a8.
//
// Solidity: function isUserAdmin(address user) view returns(bool)
func (_AccessStorageImpl *AccessStorageImplSession) IsUserAdmin(user common.Address) (bool, error) {
	return _AccessStorageImpl.Contract.IsUserAdmin(&_AccessStorageImpl.CallOpts, user)
}

// IsUserAdmin is a free data retrieval call binding the contract method 0x74d523a8.
//
// Solidity: function isUserAdmin(address user) view returns(bool)
func (_AccessStorageImpl *AccessStorageImplCallerSession) IsUserAdmin(user common.Address) (bool, error) {
	return _AccessStorageImpl.Contract.IsUserAdmin(&_AccessStorageImpl.CallOpts, user)
}

// Length is a free data retrieval call binding the contract method 0x1f7b6d32.
//
// Solidity: function length() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCaller) Length(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AccessStorageImpl.contract.Call(opts, out, "length")
	return *ret0, err
}

// Length is a free data retrieval call binding the contract method 0x1f7b6d32.
//
// Solidity: function length() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplSession) Length() (*big.Int, error) {
	return _AccessStorageImpl.Contract.Length(&_AccessStorageImpl.CallOpts)
}

// Length is a free data retrieval call binding the contract method 0x1f7b6d32.
//
// Solidity: function length() view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCallerSession) Length() (*big.Int, error) {
	return _AccessStorageImpl.Contract.Length(&_AccessStorageImpl.CallOpts)
}

// Roles is a free data retrieval call binding the contract method 0xbfda4a49.
//
// Solidity: function roles(uint256 ) view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCaller) Roles(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AccessStorageImpl.contract.Call(opts, out, "roles", arg0)
	return *ret0, err
}

// Roles is a free data retrieval call binding the contract method 0xbfda4a49.
//
// Solidity: function roles(uint256 ) view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplSession) Roles(arg0 *big.Int) (*big.Int, error) {
	return _AccessStorageImpl.Contract.Roles(&_AccessStorageImpl.CallOpts, arg0)
}

// Roles is a free data retrieval call binding the contract method 0xbfda4a49.
//
// Solidity: function roles(uint256 ) view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCallerSession) Roles(arg0 *big.Int) (*big.Int, error) {
	return _AccessStorageImpl.Contract.Roles(&_AccessStorageImpl.CallOpts, arg0)
}

// UserIndex is a free data retrieval call binding the contract method 0xc96679fe.
//
// Solidity: function userIndex(address ) view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCaller) UserIndex(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AccessStorageImpl.contract.Call(opts, out, "userIndex", arg0)
	return *ret0, err
}

// UserIndex is a free data retrieval call binding the contract method 0xc96679fe.
//
// Solidity: function userIndex(address ) view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplSession) UserIndex(arg0 common.Address) (*big.Int, error) {
	return _AccessStorageImpl.Contract.UserIndex(&_AccessStorageImpl.CallOpts, arg0)
}

// UserIndex is a free data retrieval call binding the contract method 0xc96679fe.
//
// Solidity: function userIndex(address ) view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCallerSession) UserIndex(arg0 common.Address) (*big.Int, error) {
	return _AccessStorageImpl.Contract.UserIndex(&_AccessStorageImpl.CallOpts, arg0)
}

// UserRolesIndex is a free data retrieval call binding the contract method 0xfffbf340.
//
// Solidity: function userRolesIndex(address ) view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCaller) UserRolesIndex(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AccessStorageImpl.contract.Call(opts, out, "userRolesIndex", arg0)
	return *ret0, err
}

// UserRolesIndex is a free data retrieval call binding the contract method 0xfffbf340.
//
// Solidity: function userRolesIndex(address ) view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplSession) UserRolesIndex(arg0 common.Address) (*big.Int, error) {
	return _AccessStorageImpl.Contract.UserRolesIndex(&_AccessStorageImpl.CallOpts, arg0)
}

// UserRolesIndex is a free data retrieval call binding the contract method 0xfffbf340.
//
// Solidity: function userRolesIndex(address ) view returns(uint256)
func (_AccessStorageImpl *AccessStorageImplCallerSession) UserRolesIndex(arg0 common.Address) (*big.Int, error) {
	return _AccessStorageImpl.Contract.UserRolesIndex(&_AccessStorageImpl.CallOpts, arg0)
}

// Users is a free data retrieval call binding the contract method 0x365b98b2.
//
// Solidity: function users(uint256 ) view returns(address)
func (_AccessStorageImpl *AccessStorageImplCaller) Users(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AccessStorageImpl.contract.Call(opts, out, "users", arg0)
	return *ret0, err
}

// Users is a free data retrieval call binding the contract method 0x365b98b2.
//
// Solidity: function users(uint256 ) view returns(address)
func (_AccessStorageImpl *AccessStorageImplSession) Users(arg0 *big.Int) (common.Address, error) {
	return _AccessStorageImpl.Contract.Users(&_AccessStorageImpl.CallOpts, arg0)
}

// Users is a free data retrieval call binding the contract method 0x365b98b2.
//
// Solidity: function users(uint256 ) view returns(address)
func (_AccessStorageImpl *AccessStorageImplCallerSession) Users(arg0 *big.Int) (common.Address, error) {
	return _AccessStorageImpl.Contract.Users(&_AccessStorageImpl.CallOpts, arg0)
}

// AddRole is a paid mutator transaction binding the contract method 0x0446d6fa.
//
// Solidity: function addRole(uint256 id, uint256 permissions) returns()
func (_AccessStorageImpl *AccessStorageImplTransactor) AddRole(opts *bind.TransactOpts, id *big.Int, permissions *big.Int) (*types.Transaction, error) {
	return _AccessStorageImpl.contract.Transact(opts, "addRole", id, permissions)
}

// AddRole is a paid mutator transaction binding the contract method 0x0446d6fa.
//
// Solidity: function addRole(uint256 id, uint256 permissions) returns()
func (_AccessStorageImpl *AccessStorageImplSession) AddRole(id *big.Int, permissions *big.Int) (*types.Transaction, error) {
	return _AccessStorageImpl.Contract.AddRole(&_AccessStorageImpl.TransactOpts, id, permissions)
}

// AddRole is a paid mutator transaction binding the contract method 0x0446d6fa.
//
// Solidity: function addRole(uint256 id, uint256 permissions) returns()
func (_AccessStorageImpl *AccessStorageImplTransactorSession) AddRole(id *big.Int, permissions *big.Int) (*types.Transaction, error) {
	return _AccessStorageImpl.Contract.AddRole(&_AccessStorageImpl.TransactOpts, id, permissions)
}

// AddUser is a paid mutator transaction binding the contract method 0xd3017193.
//
// Solidity: function addUser(address admin, uint256 role) returns(bool)
func (_AccessStorageImpl *AccessStorageImplTransactor) AddUser(opts *bind.TransactOpts, admin common.Address, role *big.Int) (*types.Transaction, error) {
	return _AccessStorageImpl.contract.Transact(opts, "addUser", admin, role)
}

// AddUser is a paid mutator transaction binding the contract method 0xd3017193.
//
// Solidity: function addUser(address admin, uint256 role) returns(bool)
func (_AccessStorageImpl *AccessStorageImplSession) AddUser(admin common.Address, role *big.Int) (*types.Transaction, error) {
	return _AccessStorageImpl.Contract.AddUser(&_AccessStorageImpl.TransactOpts, admin, role)
}

// AddUser is a paid mutator transaction binding the contract method 0xd3017193.
//
// Solidity: function addUser(address admin, uint256 role) returns(bool)
func (_AccessStorageImpl *AccessStorageImplTransactorSession) AddUser(admin common.Address, role *big.Int) (*types.Transaction, error) {
	return _AccessStorageImpl.Contract.AddUser(&_AccessStorageImpl.TransactOpts, admin, role)
}

// DropPermission is a paid mutator transaction binding the contract method 0xf37254ee.
//
// Solidity: function dropPermission(uint256 id, uint256 permissionConstant) returns()
func (_AccessStorageImpl *AccessStorageImplTransactor) DropPermission(opts *bind.TransactOpts, id *big.Int, permissionConstant *big.Int) (*types.Transaction, error) {
	return _AccessStorageImpl.contract.Transact(opts, "dropPermission", id, permissionConstant)
}

// DropPermission is a paid mutator transaction binding the contract method 0xf37254ee.
//
// Solidity: function dropPermission(uint256 id, uint256 permissionConstant) returns()
func (_AccessStorageImpl *AccessStorageImplSession) DropPermission(id *big.Int, permissionConstant *big.Int) (*types.Transaction, error) {
	return _AccessStorageImpl.Contract.DropPermission(&_AccessStorageImpl.TransactOpts, id, permissionConstant)
}

// DropPermission is a paid mutator transaction binding the contract method 0xf37254ee.
//
// Solidity: function dropPermission(uint256 id, uint256 permissionConstant) returns()
func (_AccessStorageImpl *AccessStorageImplTransactorSession) DropPermission(id *big.Int, permissionConstant *big.Int) (*types.Transaction, error) {
	return _AccessStorageImpl.Contract.DropPermission(&_AccessStorageImpl.TransactOpts, id, permissionConstant)
}

// EditRole is a paid mutator transaction binding the contract method 0xd12f769a.
//
// Solidity: function editRole(uint256 id, uint256 permissions) returns()
func (_AccessStorageImpl *AccessStorageImplTransactor) EditRole(opts *bind.TransactOpts, id *big.Int, permissions *big.Int) (*types.Transaction, error) {
	return _AccessStorageImpl.contract.Transact(opts, "editRole", id, permissions)
}

// EditRole is a paid mutator transaction binding the contract method 0xd12f769a.
//
// Solidity: function editRole(uint256 id, uint256 permissions) returns()
func (_AccessStorageImpl *AccessStorageImplSession) EditRole(id *big.Int, permissions *big.Int) (*types.Transaction, error) {
	return _AccessStorageImpl.Contract.EditRole(&_AccessStorageImpl.TransactOpts, id, permissions)
}

// EditRole is a paid mutator transaction binding the contract method 0xd12f769a.
//
// Solidity: function editRole(uint256 id, uint256 permissions) returns()
func (_AccessStorageImpl *AccessStorageImplTransactorSession) EditRole(id *big.Int, permissions *big.Int) (*types.Transaction, error) {
	return _AccessStorageImpl.Contract.EditRole(&_AccessStorageImpl.TransactOpts, id, permissions)
}

// EditUser is a paid mutator transaction binding the contract method 0x908df1f6.
//
// Solidity: function editUser(address oldAdmin, address newAdmin) returns(bool)
func (_AccessStorageImpl *AccessStorageImplTransactor) EditUser(opts *bind.TransactOpts, oldAdmin common.Address, newAdmin common.Address) (*types.Transaction, error) {
	return _AccessStorageImpl.contract.Transact(opts, "editUser", oldAdmin, newAdmin)
}

// EditUser is a paid mutator transaction binding the contract method 0x908df1f6.
//
// Solidity: function editUser(address oldAdmin, address newAdmin) returns(bool)
func (_AccessStorageImpl *AccessStorageImplSession) EditUser(oldAdmin common.Address, newAdmin common.Address) (*types.Transaction, error) {
	return _AccessStorageImpl.Contract.EditUser(&_AccessStorageImpl.TransactOpts, oldAdmin, newAdmin)
}

// EditUser is a paid mutator transaction binding the contract method 0x908df1f6.
//
// Solidity: function editUser(address oldAdmin, address newAdmin) returns(bool)
func (_AccessStorageImpl *AccessStorageImplTransactorSession) EditUser(oldAdmin common.Address, newAdmin common.Address) (*types.Transaction, error) {
	return _AccessStorageImpl.Contract.EditUser(&_AccessStorageImpl.TransactOpts, oldAdmin, newAdmin)
}

// RemoveRole is a paid mutator transaction binding the contract method 0x92691821.
//
// Solidity: function removeRole(uint256 id) returns()
func (_AccessStorageImpl *AccessStorageImplTransactor) RemoveRole(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _AccessStorageImpl.contract.Transact(opts, "removeRole", id)
}

// RemoveRole is a paid mutator transaction binding the contract method 0x92691821.
//
// Solidity: function removeRole(uint256 id) returns()
func (_AccessStorageImpl *AccessStorageImplSession) RemoveRole(id *big.Int) (*types.Transaction, error) {
	return _AccessStorageImpl.Contract.RemoveRole(&_AccessStorageImpl.TransactOpts, id)
}

// RemoveRole is a paid mutator transaction binding the contract method 0x92691821.
//
// Solidity: function removeRole(uint256 id) returns()
func (_AccessStorageImpl *AccessStorageImplTransactorSession) RemoveRole(id *big.Int) (*types.Transaction, error) {
	return _AccessStorageImpl.Contract.RemoveRole(&_AccessStorageImpl.TransactOpts, id)
}

// RemoveUser is a paid mutator transaction binding the contract method 0x98575188.
//
// Solidity: function removeUser(address admin) returns(bool)
func (_AccessStorageImpl *AccessStorageImplTransactor) RemoveUser(opts *bind.TransactOpts, admin common.Address) (*types.Transaction, error) {
	return _AccessStorageImpl.contract.Transact(opts, "removeUser", admin)
}

// RemoveUser is a paid mutator transaction binding the contract method 0x98575188.
//
// Solidity: function removeUser(address admin) returns(bool)
func (_AccessStorageImpl *AccessStorageImplSession) RemoveUser(admin common.Address) (*types.Transaction, error) {
	return _AccessStorageImpl.Contract.RemoveUser(&_AccessStorageImpl.TransactOpts, admin)
}

// RemoveUser is a paid mutator transaction binding the contract method 0x98575188.
//
// Solidity: function removeUser(address admin) returns(bool)
func (_AccessStorageImpl *AccessStorageImplTransactorSession) RemoveUser(admin common.Address) (*types.Transaction, error) {
	return _AccessStorageImpl.Contract.RemoveUser(&_AccessStorageImpl.TransactOpts, admin)
}

// RemoveUserAt is a paid mutator transaction binding the contract method 0x71ec1c7c.
//
// Solidity: function removeUserAt(uint256 index) returns(bool)
func (_AccessStorageImpl *AccessStorageImplTransactor) RemoveUserAt(opts *bind.TransactOpts, index *big.Int) (*types.Transaction, error) {
	return _AccessStorageImpl.contract.Transact(opts, "removeUserAt", index)
}

// RemoveUserAt is a paid mutator transaction binding the contract method 0x71ec1c7c.
//
// Solidity: function removeUserAt(uint256 index) returns(bool)
func (_AccessStorageImpl *AccessStorageImplSession) RemoveUserAt(index *big.Int) (*types.Transaction, error) {
	return _AccessStorageImpl.Contract.RemoveUserAt(&_AccessStorageImpl.TransactOpts, index)
}

// RemoveUserAt is a paid mutator transaction binding the contract method 0x71ec1c7c.
//
// Solidity: function removeUserAt(uint256 index) returns(bool)
func (_AccessStorageImpl *AccessStorageImplTransactorSession) RemoveUserAt(index *big.Int) (*types.Transaction, error) {
	return _AccessStorageImpl.Contract.RemoveUserAt(&_AccessStorageImpl.TransactOpts, index)
}

// RemoveUserRole is a paid mutator transaction binding the contract method 0x11cb1991.
//
// Solidity: function removeUserRole(address admin) returns()
func (_AccessStorageImpl *AccessStorageImplTransactor) RemoveUserRole(opts *bind.TransactOpts, admin common.Address) (*types.Transaction, error) {
	return _AccessStorageImpl.contract.Transact(opts, "removeUserRole", admin)
}

// RemoveUserRole is a paid mutator transaction binding the contract method 0x11cb1991.
//
// Solidity: function removeUserRole(address admin) returns()
func (_AccessStorageImpl *AccessStorageImplSession) RemoveUserRole(admin common.Address) (*types.Transaction, error) {
	return _AccessStorageImpl.Contract.RemoveUserRole(&_AccessStorageImpl.TransactOpts, admin)
}

// RemoveUserRole is a paid mutator transaction binding the contract method 0x11cb1991.
//
// Solidity: function removeUserRole(address admin) returns()
func (_AccessStorageImpl *AccessStorageImplTransactorSession) RemoveUserRole(admin common.Address) (*types.Transaction, error) {
	return _AccessStorageImpl.Contract.RemoveUserRole(&_AccessStorageImpl.TransactOpts, admin)
}

// SetPermission is a paid mutator transaction binding the contract method 0x2837c405.
//
// Solidity: function setPermission(uint256 id, uint256 permissionConstant) returns()
func (_AccessStorageImpl *AccessStorageImplTransactor) SetPermission(opts *bind.TransactOpts, id *big.Int, permissionConstant *big.Int) (*types.Transaction, error) {
	return _AccessStorageImpl.contract.Transact(opts, "setPermission", id, permissionConstant)
}

// SetPermission is a paid mutator transaction binding the contract method 0x2837c405.
//
// Solidity: function setPermission(uint256 id, uint256 permissionConstant) returns()
func (_AccessStorageImpl *AccessStorageImplSession) SetPermission(id *big.Int, permissionConstant *big.Int) (*types.Transaction, error) {
	return _AccessStorageImpl.Contract.SetPermission(&_AccessStorageImpl.TransactOpts, id, permissionConstant)
}

// SetPermission is a paid mutator transaction binding the contract method 0x2837c405.
//
// Solidity: function setPermission(uint256 id, uint256 permissionConstant) returns()
func (_AccessStorageImpl *AccessStorageImplTransactorSession) SetPermission(id *big.Int, permissionConstant *big.Int) (*types.Transaction, error) {
	return _AccessStorageImpl.Contract.SetPermission(&_AccessStorageImpl.TransactOpts, id, permissionConstant)
}

// SetUserRole is a paid mutator transaction binding the contract method 0xc8444031.
//
// Solidity: function setUserRole(address user, uint256 role) returns()
func (_AccessStorageImpl *AccessStorageImplTransactor) SetUserRole(opts *bind.TransactOpts, user common.Address, role *big.Int) (*types.Transaction, error) {
	return _AccessStorageImpl.contract.Transact(opts, "setUserRole", user, role)
}

// SetUserRole is a paid mutator transaction binding the contract method 0xc8444031.
//
// Solidity: function setUserRole(address user, uint256 role) returns()
func (_AccessStorageImpl *AccessStorageImplSession) SetUserRole(user common.Address, role *big.Int) (*types.Transaction, error) {
	return _AccessStorageImpl.Contract.SetUserRole(&_AccessStorageImpl.TransactOpts, user, role)
}

// SetUserRole is a paid mutator transaction binding the contract method 0xc8444031.
//
// Solidity: function setUserRole(address user, uint256 role) returns()
func (_AccessStorageImpl *AccessStorageImplTransactorSession) SetUserRole(user common.Address, role *big.Int) (*types.Transaction, error) {
	return _AccessStorageImpl.Contract.SetUserRole(&_AccessStorageImpl.TransactOpts, user, role)
}

// AccessStorageImplAdminAddedIterator is returned from FilterAdminAdded and is used to iterate over the raw logs and unpacked data for AdminAdded events raised by the AccessStorageImpl contract.
type AccessStorageImplAdminAddedIterator struct {
	Event *AccessStorageImplAdminAdded // Event containing the contract specifics and raw log

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
func (it *AccessStorageImplAdminAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessStorageImplAdminAdded)
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
		it.Event = new(AccessStorageImplAdminAdded)
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
func (it *AccessStorageImplAdminAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccessStorageImplAdminAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccessStorageImplAdminAdded represents a AdminAdded event raised by the AccessStorageImpl contract.
type AccessStorageImplAdminAdded struct {
	Admin common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAdminAdded is a free log retrieval operation binding the contract event 0x44d6d25963f097ad14f29f06854a01f575648a1ef82f30e562ccd3889717e339.
//
// Solidity: event AdminAdded(address admin)
func (_AccessStorageImpl *AccessStorageImplFilterer) FilterAdminAdded(opts *bind.FilterOpts) (*AccessStorageImplAdminAddedIterator, error) {

	logs, sub, err := _AccessStorageImpl.contract.FilterLogs(opts, "AdminAdded")
	if err != nil {
		return nil, err
	}
	return &AccessStorageImplAdminAddedIterator{contract: _AccessStorageImpl.contract, event: "AdminAdded", logs: logs, sub: sub}, nil
}

// WatchAdminAdded is a free log subscription operation binding the contract event 0x44d6d25963f097ad14f29f06854a01f575648a1ef82f30e562ccd3889717e339.
//
// Solidity: event AdminAdded(address admin)
func (_AccessStorageImpl *AccessStorageImplFilterer) WatchAdminAdded(opts *bind.WatchOpts, sink chan<- *AccessStorageImplAdminAdded) (event.Subscription, error) {

	logs, sub, err := _AccessStorageImpl.contract.WatchLogs(opts, "AdminAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccessStorageImplAdminAdded)
				if err := _AccessStorageImpl.contract.UnpackLog(event, "AdminAdded", log); err != nil {
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

// ParseAdminAdded is a log parse operation binding the contract event 0x44d6d25963f097ad14f29f06854a01f575648a1ef82f30e562ccd3889717e339.
//
// Solidity: event AdminAdded(address admin)
func (_AccessStorageImpl *AccessStorageImplFilterer) ParseAdminAdded(log types.Log) (*AccessStorageImplAdminAdded, error) {
	event := new(AccessStorageImplAdminAdded)
	if err := _AccessStorageImpl.contract.UnpackLog(event, "AdminAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AccessStorageImplAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the AccessStorageImpl contract.
type AccessStorageImplAdminChangedIterator struct {
	Event *AccessStorageImplAdminChanged // Event containing the contract specifics and raw log

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
func (it *AccessStorageImplAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessStorageImplAdminChanged)
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
		it.Event = new(AccessStorageImplAdminChanged)
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
func (it *AccessStorageImplAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccessStorageImplAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccessStorageImplAdminChanged represents a AdminChanged event raised by the AccessStorageImpl contract.
type AccessStorageImplAdminChanged struct {
	OldAdmin common.Address
	NewAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address oldAdmin, address newAdmin)
func (_AccessStorageImpl *AccessStorageImplFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*AccessStorageImplAdminChangedIterator, error) {

	logs, sub, err := _AccessStorageImpl.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &AccessStorageImplAdminChangedIterator{contract: _AccessStorageImpl.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address oldAdmin, address newAdmin)
func (_AccessStorageImpl *AccessStorageImplFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *AccessStorageImplAdminChanged) (event.Subscription, error) {

	logs, sub, err := _AccessStorageImpl.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccessStorageImplAdminChanged)
				if err := _AccessStorageImpl.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address oldAdmin, address newAdmin)
func (_AccessStorageImpl *AccessStorageImplFilterer) ParseAdminChanged(log types.Log) (*AccessStorageImplAdminChanged, error) {
	event := new(AccessStorageImplAdminChanged)
	if err := _AccessStorageImpl.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AccessStorageImplAdminRemovedIterator is returned from FilterAdminRemoved and is used to iterate over the raw logs and unpacked data for AdminRemoved events raised by the AccessStorageImpl contract.
type AccessStorageImplAdminRemovedIterator struct {
	Event *AccessStorageImplAdminRemoved // Event containing the contract specifics and raw log

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
func (it *AccessStorageImplAdminRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessStorageImplAdminRemoved)
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
		it.Event = new(AccessStorageImplAdminRemoved)
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
func (it *AccessStorageImplAdminRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccessStorageImplAdminRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccessStorageImplAdminRemoved represents a AdminRemoved event raised by the AccessStorageImpl contract.
type AccessStorageImplAdminRemoved struct {
	Admin common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAdminRemoved is a free log retrieval operation binding the contract event 0xa3b62bc36326052d97ea62d63c3d60308ed4c3ea8ac079dd8499f1e9c4f80c0f.
//
// Solidity: event AdminRemoved(address admin)
func (_AccessStorageImpl *AccessStorageImplFilterer) FilterAdminRemoved(opts *bind.FilterOpts) (*AccessStorageImplAdminRemovedIterator, error) {

	logs, sub, err := _AccessStorageImpl.contract.FilterLogs(opts, "AdminRemoved")
	if err != nil {
		return nil, err
	}
	return &AccessStorageImplAdminRemovedIterator{contract: _AccessStorageImpl.contract, event: "AdminRemoved", logs: logs, sub: sub}, nil
}

// WatchAdminRemoved is a free log subscription operation binding the contract event 0xa3b62bc36326052d97ea62d63c3d60308ed4c3ea8ac079dd8499f1e9c4f80c0f.
//
// Solidity: event AdminRemoved(address admin)
func (_AccessStorageImpl *AccessStorageImplFilterer) WatchAdminRemoved(opts *bind.WatchOpts, sink chan<- *AccessStorageImplAdminRemoved) (event.Subscription, error) {

	logs, sub, err := _AccessStorageImpl.contract.WatchLogs(opts, "AdminRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccessStorageImplAdminRemoved)
				if err := _AccessStorageImpl.contract.UnpackLog(event, "AdminRemoved", log); err != nil {
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

// ParseAdminRemoved is a log parse operation binding the contract event 0xa3b62bc36326052d97ea62d63c3d60308ed4c3ea8ac079dd8499f1e9c4f80c0f.
//
// Solidity: event AdminRemoved(address admin)
func (_AccessStorageImpl *AccessStorageImplFilterer) ParseAdminRemoved(log types.Log) (*AccessStorageImplAdminRemoved, error) {
	event := new(AccessStorageImplAdminRemoved)
	if err := _AccessStorageImpl.contract.UnpackLog(event, "AdminRemoved", log); err != nil {
		return nil, err
	}
	return event, nil
}
