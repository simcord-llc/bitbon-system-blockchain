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

// AssetboxImplABI is the input ABI used to generate the binding from.
const AssetboxImplABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractStorageAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ACCESS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ACCESS_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ADMINABLE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ADMIN_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ASSETBOX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ASSETBOX_INFO\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ASSETBOX_INFO_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ASSETBOX_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_BALANCE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_BITBON\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_BITBON_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_BITBON_SUPPORT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_BONBALANCE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_BONTRANSFER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_DEX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_DISTRIBUTION_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_EXCHANGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_EXCHANGE_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_FEE_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_GENERATOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MINING_AGENT_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MSBON\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MSBON_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_ADD_ADDRESS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_ADD_ADMIN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_ADD_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_ADD_ROLE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_DROP_PERMISSION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_EDIT_ADMIN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_EDIT_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_EDIT_ROLE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REMOVE_ADMIN_BY_ADDRESS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REMOVE_ADMIN_BY_INDEX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REMOVE_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REMOVE_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REMOVE_ROLE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REPLENISH_CAPITALIZATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REPLENISH_CERTIFICATE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REPLENISH_EMISSION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_SET_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_SET_PERMISSION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_OTC\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_OTC_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_REPLENISH_CAPITALIZATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_RESERVED_ALIAS_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ROLE_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_SAFE_TRANSFER_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_TRANSFER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_TRANSFER_FROM_CAPITALIZATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_ACCESS_RESTORATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_ADMIN_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_CONTRACTS_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_CONTRACT_DEPLOY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_EMISSION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_FEES_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_MULTISIG_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_ROLES_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_ACCESS_VERIFIER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_BITBON_ISSUE_VERIFIER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_COMMISSION_VERIFIER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_DEPLOY_ADMIN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_PERMISSION_ADMIN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getThisContractIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"assetboxAlias\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"serviceId\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"isMining\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"otherInfo\",\"type\":\"bytes\"}],\"name\":\"setAssetboxInfo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"deleteAssetboxInfo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"assetbox\",\"type\":\"address\"}],\"name\":\"deleteAssetboxInfoAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"assetbox\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"assetboxAlias\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"serviceId\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"isMining\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"otherInfo\",\"type\":\"bytes\"}],\"name\":\"setAssetboxInfoAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"assetboxes\",\"type\":\"address[]\"},{\"internalType\":\"string[]\",\"name\":\"assetboxAliases\",\"type\":\"string[]\"},{\"internalType\":\"bytes[]\",\"name\":\"serviceIds\",\"type\":\"bytes[]\"},{\"internalType\":\"bool[]\",\"name\":\"isMiningFlags\",\"type\":\"bool[]\"},{\"internalType\":\"bytes[]\",\"name\":\"otherInfos\",\"type\":\"bytes[]\"}],\"name\":\"setAssetboxInfosAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"assetbox\",\"type\":\"address\"}],\"name\":\"getAssetboxInfoByAssetbox\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"assetboxAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"assetboxAlias\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"serviceId\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"isMining\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"otherInfo\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getMyAssetboxInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"assetboxAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"assetboxAlias\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"serviceId\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"isMining\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"otherInfo\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"assetboxAliasString\",\"type\":\"string\"}],\"name\":\"getAssetboxInfoByAlias\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"assetboxAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"assetboxAlias\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"serviceId\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"isMining\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"otherInfo\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"assetbox\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"serviceId\",\"type\":\"bytes\"}],\"name\":\"isAssetboxBoundToService\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isBound\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"serviceId\",\"type\":\"bytes\"}],\"name\":\"isMyAssetboxBoundToService\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isBound\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"assetboxAlias\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"serviceId\",\"type\":\"bytes\"}],\"name\":\"isAliasBoundToService\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isBound\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"assetboxAlias\",\"type\":\"string\"}],\"name\":\"isAliasTaken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"taken\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"assetbox\",\"type\":\"address\"}],\"name\":\"isAssetboxMining\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"aliasString\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"assetbox\",\"type\":\"address\"}],\"name\":\"approveReservedAlias\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"aliasString\",\"type\":\"string\"}],\"name\":\"reserveAlias\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"oldAlias\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"newAlias\",\"type\":\"string\"}],\"name\":\"editReservedAlias\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"aliasString\",\"type\":\"string\"}],\"name\":\"removeAliasReservation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"removeAliasReservationAt\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// AssetboxImpl is an auto generated Go binding around an Ethereum contract.
type AssetboxImpl struct {
	AssetboxImplCaller     // Read-only binding to the contract
	AssetboxImplTransactor // Write-only binding to the contract
	AssetboxImplFilterer   // Log filterer for contract events
}

// AssetboxImplCaller is an auto generated read-only Go binding around an Ethereum contract.
type AssetboxImplCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssetboxImplTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AssetboxImplTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssetboxImplFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AssetboxImplFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssetboxImplSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AssetboxImplSession struct {
	Contract     *AssetboxImpl     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AssetboxImplCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AssetboxImplCallerSession struct {
	Contract *AssetboxImplCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// AssetboxImplTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AssetboxImplTransactorSession struct {
	Contract     *AssetboxImplTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// AssetboxImplRaw is an auto generated low-level Go binding around an Ethereum contract.
type AssetboxImplRaw struct {
	Contract *AssetboxImpl // Generic contract binding to access the raw methods on
}

// AssetboxImplCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AssetboxImplCallerRaw struct {
	Contract *AssetboxImplCaller // Generic read-only contract binding to access the raw methods on
}

// AssetboxImplTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AssetboxImplTransactorRaw struct {
	Contract *AssetboxImplTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAssetboxImpl creates a new instance of AssetboxImpl, bound to a specific deployed contract.
func NewAssetboxImpl(address common.Address, backend bind.ContractBackend) (*AssetboxImpl, error) {
	contract, err := bindAssetboxImpl(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AssetboxImpl{AssetboxImplCaller: AssetboxImplCaller{contract: contract}, AssetboxImplTransactor: AssetboxImplTransactor{contract: contract}, AssetboxImplFilterer: AssetboxImplFilterer{contract: contract}}, nil
}

// NewAssetboxImplCaller creates a new read-only instance of AssetboxImpl, bound to a specific deployed contract.
func NewAssetboxImplCaller(address common.Address, caller bind.ContractCaller) (*AssetboxImplCaller, error) {
	contract, err := bindAssetboxImpl(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AssetboxImplCaller{contract: contract}, nil
}

// NewAssetboxImplTransactor creates a new write-only instance of AssetboxImpl, bound to a specific deployed contract.
func NewAssetboxImplTransactor(address common.Address, transactor bind.ContractTransactor) (*AssetboxImplTransactor, error) {
	contract, err := bindAssetboxImpl(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AssetboxImplTransactor{contract: contract}, nil
}

// NewAssetboxImplFilterer creates a new log filterer instance of AssetboxImpl, bound to a specific deployed contract.
func NewAssetboxImplFilterer(address common.Address, filterer bind.ContractFilterer) (*AssetboxImplFilterer, error) {
	contract, err := bindAssetboxImpl(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AssetboxImplFilterer{contract: contract}, nil
}

// bindAssetboxImpl binds a generic wrapper to an already deployed contract.
func bindAssetboxImpl(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AssetboxImplABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AssetboxImpl *AssetboxImplRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AssetboxImpl.Contract.AssetboxImplCaller.contract.Call(opts, result, method, params...)
}

// Return Abi for AssetboxImplCaller
func (_AssetboxImpl *AssetboxImplCaller) ABI() abi.ABI {
	return _AssetboxImpl.contract.Abi
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AssetboxImpl *AssetboxImplRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AssetboxImpl.Contract.AssetboxImplTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AssetboxImpl *AssetboxImplRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AssetboxImpl.Contract.AssetboxImplTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AssetboxImpl *AssetboxImplCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AssetboxImpl.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AssetboxImpl *AssetboxImplTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AssetboxImpl.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AssetboxImpl *AssetboxImplTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AssetboxImpl.Contract.contract.Transact(opts, method, params...)
}

// CONTRACTACCESS is a free data retrieval call binding the contract method 0x0ac62b0c.
//
// Solidity: function CONTRACT_ACCESS() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCaller) CONTRACTACCESS(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxImpl.contract.Call(opts, out, "CONTRACT_ACCESS")
	return *ret0, err
}

// CONTRACTACCESS is a free data retrieval call binding the contract method 0x0ac62b0c.
//
// Solidity: function CONTRACT_ACCESS() view returns(uint256)
func (_AssetboxImpl *AssetboxImplSession) CONTRACTACCESS() (*big.Int, error) {
	return _AssetboxImpl.Contract.CONTRACTACCESS(&_AssetboxImpl.CallOpts)
}

// CONTRACTACCESS is a free data retrieval call binding the contract method 0x0ac62b0c.
//
// Solidity: function CONTRACT_ACCESS() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCallerSession) CONTRACTACCESS() (*big.Int, error) {
	return _AssetboxImpl.Contract.CONTRACTACCESS(&_AssetboxImpl.CallOpts)
}

// CONTRACTACCESSSTORAGE is a free data retrieval call binding the contract method 0xceac9bd5.
//
// Solidity: function CONTRACT_ACCESS_STORAGE() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCaller) CONTRACTACCESSSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxImpl.contract.Call(opts, out, "CONTRACT_ACCESS_STORAGE")
	return *ret0, err
}

// CONTRACTACCESSSTORAGE is a free data retrieval call binding the contract method 0xceac9bd5.
//
// Solidity: function CONTRACT_ACCESS_STORAGE() view returns(uint256)
func (_AssetboxImpl *AssetboxImplSession) CONTRACTACCESSSTORAGE() (*big.Int, error) {
	return _AssetboxImpl.Contract.CONTRACTACCESSSTORAGE(&_AssetboxImpl.CallOpts)
}

// CONTRACTACCESSSTORAGE is a free data retrieval call binding the contract method 0xceac9bd5.
//
// Solidity: function CONTRACT_ACCESS_STORAGE() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCallerSession) CONTRACTACCESSSTORAGE() (*big.Int, error) {
	return _AssetboxImpl.Contract.CONTRACTACCESSSTORAGE(&_AssetboxImpl.CallOpts)
}

// CONTRACTADMINABLE is a free data retrieval call binding the contract method 0xbbe3a78f.
//
// Solidity: function CONTRACT_ADMINABLE() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCaller) CONTRACTADMINABLE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxImpl.contract.Call(opts, out, "CONTRACT_ADMINABLE")
	return *ret0, err
}

// CONTRACTADMINABLE is a free data retrieval call binding the contract method 0xbbe3a78f.
//
// Solidity: function CONTRACT_ADMINABLE() view returns(uint256)
func (_AssetboxImpl *AssetboxImplSession) CONTRACTADMINABLE() (*big.Int, error) {
	return _AssetboxImpl.Contract.CONTRACTADMINABLE(&_AssetboxImpl.CallOpts)
}

// CONTRACTADMINABLE is a free data retrieval call binding the contract method 0xbbe3a78f.
//
// Solidity: function CONTRACT_ADMINABLE() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCallerSession) CONTRACTADMINABLE() (*big.Int, error) {
	return _AssetboxImpl.Contract.CONTRACTADMINABLE(&_AssetboxImpl.CallOpts)
}

// CONTRACTADMINSTORAGE is a free data retrieval call binding the contract method 0xc9aebae9.
//
// Solidity: function CONTRACT_ADMIN_STORAGE() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCaller) CONTRACTADMINSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxImpl.contract.Call(opts, out, "CONTRACT_ADMIN_STORAGE")
	return *ret0, err
}

// CONTRACTADMINSTORAGE is a free data retrieval call binding the contract method 0xc9aebae9.
//
// Solidity: function CONTRACT_ADMIN_STORAGE() view returns(uint256)
func (_AssetboxImpl *AssetboxImplSession) CONTRACTADMINSTORAGE() (*big.Int, error) {
	return _AssetboxImpl.Contract.CONTRACTADMINSTORAGE(&_AssetboxImpl.CallOpts)
}

// CONTRACTADMINSTORAGE is a free data retrieval call binding the contract method 0xc9aebae9.
//
// Solidity: function CONTRACT_ADMIN_STORAGE() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCallerSession) CONTRACTADMINSTORAGE() (*big.Int, error) {
	return _AssetboxImpl.Contract.CONTRACTADMINSTORAGE(&_AssetboxImpl.CallOpts)
}

// CONTRACTASSETBOX is a free data retrieval call binding the contract method 0x161aa17c.
//
// Solidity: function CONTRACT_ASSETBOX() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCaller) CONTRACTASSETBOX(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxImpl.contract.Call(opts, out, "CONTRACT_ASSETBOX")
	return *ret0, err
}

// CONTRACTASSETBOX is a free data retrieval call binding the contract method 0x161aa17c.
//
// Solidity: function CONTRACT_ASSETBOX() view returns(uint256)
func (_AssetboxImpl *AssetboxImplSession) CONTRACTASSETBOX() (*big.Int, error) {
	return _AssetboxImpl.Contract.CONTRACTASSETBOX(&_AssetboxImpl.CallOpts)
}

// CONTRACTASSETBOX is a free data retrieval call binding the contract method 0x161aa17c.
//
// Solidity: function CONTRACT_ASSETBOX() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCallerSession) CONTRACTASSETBOX() (*big.Int, error) {
	return _AssetboxImpl.Contract.CONTRACTASSETBOX(&_AssetboxImpl.CallOpts)
}

// CONTRACTASSETBOXINFO is a free data retrieval call binding the contract method 0xb572fd3d.
//
// Solidity: function CONTRACT_ASSETBOX_INFO() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCaller) CONTRACTASSETBOXINFO(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxImpl.contract.Call(opts, out, "CONTRACT_ASSETBOX_INFO")
	return *ret0, err
}

// CONTRACTASSETBOXINFO is a free data retrieval call binding the contract method 0xb572fd3d.
//
// Solidity: function CONTRACT_ASSETBOX_INFO() view returns(uint256)
func (_AssetboxImpl *AssetboxImplSession) CONTRACTASSETBOXINFO() (*big.Int, error) {
	return _AssetboxImpl.Contract.CONTRACTASSETBOXINFO(&_AssetboxImpl.CallOpts)
}

// CONTRACTASSETBOXINFO is a free data retrieval call binding the contract method 0xb572fd3d.
//
// Solidity: function CONTRACT_ASSETBOX_INFO() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCallerSession) CONTRACTASSETBOXINFO() (*big.Int, error) {
	return _AssetboxImpl.Contract.CONTRACTASSETBOXINFO(&_AssetboxImpl.CallOpts)
}

// CONTRACTASSETBOXINFOSTORAGE is a free data retrieval call binding the contract method 0xf32a115d.
//
// Solidity: function CONTRACT_ASSETBOX_INFO_STORAGE() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCaller) CONTRACTASSETBOXINFOSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxImpl.contract.Call(opts, out, "CONTRACT_ASSETBOX_INFO_STORAGE")
	return *ret0, err
}

// CONTRACTASSETBOXINFOSTORAGE is a free data retrieval call binding the contract method 0xf32a115d.
//
// Solidity: function CONTRACT_ASSETBOX_INFO_STORAGE() view returns(uint256)
func (_AssetboxImpl *AssetboxImplSession) CONTRACTASSETBOXINFOSTORAGE() (*big.Int, error) {
	return _AssetboxImpl.Contract.CONTRACTASSETBOXINFOSTORAGE(&_AssetboxImpl.CallOpts)
}

// CONTRACTASSETBOXINFOSTORAGE is a free data retrieval call binding the contract method 0xf32a115d.
//
// Solidity: function CONTRACT_ASSETBOX_INFO_STORAGE() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCallerSession) CONTRACTASSETBOXINFOSTORAGE() (*big.Int, error) {
	return _AssetboxImpl.Contract.CONTRACTASSETBOXINFOSTORAGE(&_AssetboxImpl.CallOpts)
}

// CONTRACTASSETBOXSTORAGE is a free data retrieval call binding the contract method 0xfa5bdafc.
//
// Solidity: function CONTRACT_ASSETBOX_STORAGE() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCaller) CONTRACTASSETBOXSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxImpl.contract.Call(opts, out, "CONTRACT_ASSETBOX_STORAGE")
	return *ret0, err
}

// CONTRACTASSETBOXSTORAGE is a free data retrieval call binding the contract method 0xfa5bdafc.
//
// Solidity: function CONTRACT_ASSETBOX_STORAGE() view returns(uint256)
func (_AssetboxImpl *AssetboxImplSession) CONTRACTASSETBOXSTORAGE() (*big.Int, error) {
	return _AssetboxImpl.Contract.CONTRACTASSETBOXSTORAGE(&_AssetboxImpl.CallOpts)
}

// CONTRACTASSETBOXSTORAGE is a free data retrieval call binding the contract method 0xfa5bdafc.
//
// Solidity: function CONTRACT_ASSETBOX_STORAGE() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCallerSession) CONTRACTASSETBOXSTORAGE() (*big.Int, error) {
	return _AssetboxImpl.Contract.CONTRACTASSETBOXSTORAGE(&_AssetboxImpl.CallOpts)
}

// CONTRACTBALANCE is a free data retrieval call binding the contract method 0xd59045f5.
//
// Solidity: function CONTRACT_BALANCE() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCaller) CONTRACTBALANCE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxImpl.contract.Call(opts, out, "CONTRACT_BALANCE")
	return *ret0, err
}

// CONTRACTBALANCE is a free data retrieval call binding the contract method 0xd59045f5.
//
// Solidity: function CONTRACT_BALANCE() view returns(uint256)
func (_AssetboxImpl *AssetboxImplSession) CONTRACTBALANCE() (*big.Int, error) {
	return _AssetboxImpl.Contract.CONTRACTBALANCE(&_AssetboxImpl.CallOpts)
}

// CONTRACTBALANCE is a free data retrieval call binding the contract method 0xd59045f5.
//
// Solidity: function CONTRACT_BALANCE() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCallerSession) CONTRACTBALANCE() (*big.Int, error) {
	return _AssetboxImpl.Contract.CONTRACTBALANCE(&_AssetboxImpl.CallOpts)
}

// CONTRACTBITBON is a free data retrieval call binding the contract method 0xb6b433ab.
//
// Solidity: function CONTRACT_BITBON() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCaller) CONTRACTBITBON(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxImpl.contract.Call(opts, out, "CONTRACT_BITBON")
	return *ret0, err
}

// CONTRACTBITBON is a free data retrieval call binding the contract method 0xb6b433ab.
//
// Solidity: function CONTRACT_BITBON() view returns(uint256)
func (_AssetboxImpl *AssetboxImplSession) CONTRACTBITBON() (*big.Int, error) {
	return _AssetboxImpl.Contract.CONTRACTBITBON(&_AssetboxImpl.CallOpts)
}

// CONTRACTBITBON is a free data retrieval call binding the contract method 0xb6b433ab.
//
// Solidity: function CONTRACT_BITBON() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCallerSession) CONTRACTBITBON() (*big.Int, error) {
	return _AssetboxImpl.Contract.CONTRACTBITBON(&_AssetboxImpl.CallOpts)
}

// CONTRACTBITBONSTORAGE is a free data retrieval call binding the contract method 0xe7f6c30e.
//
// Solidity: function CONTRACT_BITBON_STORAGE() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCaller) CONTRACTBITBONSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxImpl.contract.Call(opts, out, "CONTRACT_BITBON_STORAGE")
	return *ret0, err
}

// CONTRACTBITBONSTORAGE is a free data retrieval call binding the contract method 0xe7f6c30e.
//
// Solidity: function CONTRACT_BITBON_STORAGE() view returns(uint256)
func (_AssetboxImpl *AssetboxImplSession) CONTRACTBITBONSTORAGE() (*big.Int, error) {
	return _AssetboxImpl.Contract.CONTRACTBITBONSTORAGE(&_AssetboxImpl.CallOpts)
}

// CONTRACTBITBONSTORAGE is a free data retrieval call binding the contract method 0xe7f6c30e.
//
// Solidity: function CONTRACT_BITBON_STORAGE() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCallerSession) CONTRACTBITBONSTORAGE() (*big.Int, error) {
	return _AssetboxImpl.Contract.CONTRACTBITBONSTORAGE(&_AssetboxImpl.CallOpts)
}

// CONTRACTBITBONSUPPORT is a free data retrieval call binding the contract method 0x20139c74.
//
// Solidity: function CONTRACT_BITBON_SUPPORT() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCaller) CONTRACTBITBONSUPPORT(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxImpl.contract.Call(opts, out, "CONTRACT_BITBON_SUPPORT")
	return *ret0, err
}

// CONTRACTBITBONSUPPORT is a free data retrieval call binding the contract method 0x20139c74.
//
// Solidity: function CONTRACT_BITBON_SUPPORT() view returns(uint256)
func (_AssetboxImpl *AssetboxImplSession) CONTRACTBITBONSUPPORT() (*big.Int, error) {
	return _AssetboxImpl.Contract.CONTRACTBITBONSUPPORT(&_AssetboxImpl.CallOpts)
}

// CONTRACTBITBONSUPPORT is a free data retrieval call binding the contract method 0x20139c74.
//
// Solidity: function CONTRACT_BITBON_SUPPORT() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCallerSession) CONTRACTBITBONSUPPORT() (*big.Int, error) {
	return _AssetboxImpl.Contract.CONTRACTBITBONSUPPORT(&_AssetboxImpl.CallOpts)
}

// CONTRACTBONBALANCE is a free data retrieval call binding the contract method 0x40e02b7c.
//
// Solidity: function CONTRACT_BONBALANCE() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCaller) CONTRACTBONBALANCE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxImpl.contract.Call(opts, out, "CONTRACT_BONBALANCE")
	return *ret0, err
}

// CONTRACTBONBALANCE is a free data retrieval call binding the contract method 0x40e02b7c.
//
// Solidity: function CONTRACT_BONBALANCE() view returns(uint256)
func (_AssetboxImpl *AssetboxImplSession) CONTRACTBONBALANCE() (*big.Int, error) {
	return _AssetboxImpl.Contract.CONTRACTBONBALANCE(&_AssetboxImpl.CallOpts)
}

// CONTRACTBONBALANCE is a free data retrieval call binding the contract method 0x40e02b7c.
//
// Solidity: function CONTRACT_BONBALANCE() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCallerSession) CONTRACTBONBALANCE() (*big.Int, error) {
	return _AssetboxImpl.Contract.CONTRACTBONBALANCE(&_AssetboxImpl.CallOpts)
}

// CONTRACTBONTRANSFER is a free data retrieval call binding the contract method 0xc18f4a8a.
//
// Solidity: function CONTRACT_BONTRANSFER() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCaller) CONTRACTBONTRANSFER(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxImpl.contract.Call(opts, out, "CONTRACT_BONTRANSFER")
	return *ret0, err
}

// CONTRACTBONTRANSFER is a free data retrieval call binding the contract method 0xc18f4a8a.
//
// Solidity: function CONTRACT_BONTRANSFER() view returns(uint256)
func (_AssetboxImpl *AssetboxImplSession) CONTRACTBONTRANSFER() (*big.Int, error) {
	return _AssetboxImpl.Contract.CONTRACTBONTRANSFER(&_AssetboxImpl.CallOpts)
}

// CONTRACTBONTRANSFER is a free data retrieval call binding the contract method 0xc18f4a8a.
//
// Solidity: function CONTRACT_BONTRANSFER() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCallerSession) CONTRACTBONTRANSFER() (*big.Int, error) {
	return _AssetboxImpl.Contract.CONTRACTBONTRANSFER(&_AssetboxImpl.CallOpts)
}

// CONTRACTDEX is a free data retrieval call binding the contract method 0x52c1f844.
//
// Solidity: function CONTRACT_DEX() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCaller) CONTRACTDEX(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxImpl.contract.Call(opts, out, "CONTRACT_DEX")
	return *ret0, err
}

// CONTRACTDEX is a free data retrieval call binding the contract method 0x52c1f844.
//
// Solidity: function CONTRACT_DEX() view returns(uint256)
func (_AssetboxImpl *AssetboxImplSession) CONTRACTDEX() (*big.Int, error) {
	return _AssetboxImpl.Contract.CONTRACTDEX(&_AssetboxImpl.CallOpts)
}

// CONTRACTDEX is a free data retrieval call binding the contract method 0x52c1f844.
//
// Solidity: function CONTRACT_DEX() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCallerSession) CONTRACTDEX() (*big.Int, error) {
	return _AssetboxImpl.Contract.CONTRACTDEX(&_AssetboxImpl.CallOpts)
}

// CONTRACTDISTRIBUTIONSTORAGE is a free data retrieval call binding the contract method 0x08097327.
//
// Solidity: function CONTRACT_DISTRIBUTION_STORAGE() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCaller) CONTRACTDISTRIBUTIONSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxImpl.contract.Call(opts, out, "CONTRACT_DISTRIBUTION_STORAGE")
	return *ret0, err
}

// CONTRACTDISTRIBUTIONSTORAGE is a free data retrieval call binding the contract method 0x08097327.
//
// Solidity: function CONTRACT_DISTRIBUTION_STORAGE() view returns(uint256)
func (_AssetboxImpl *AssetboxImplSession) CONTRACTDISTRIBUTIONSTORAGE() (*big.Int, error) {
	return _AssetboxImpl.Contract.CONTRACTDISTRIBUTIONSTORAGE(&_AssetboxImpl.CallOpts)
}

// CONTRACTDISTRIBUTIONSTORAGE is a free data retrieval call binding the contract method 0x08097327.
//
// Solidity: function CONTRACT_DISTRIBUTION_STORAGE() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCallerSession) CONTRACTDISTRIBUTIONSTORAGE() (*big.Int, error) {
	return _AssetboxImpl.Contract.CONTRACTDISTRIBUTIONSTORAGE(&_AssetboxImpl.CallOpts)
}

// CONTRACTEXCHANGE is a free data retrieval call binding the contract method 0xa0e3a10e.
//
// Solidity: function CONTRACT_EXCHANGE() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCaller) CONTRACTEXCHANGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxImpl.contract.Call(opts, out, "CONTRACT_EXCHANGE")
	return *ret0, err
}

// CONTRACTEXCHANGE is a free data retrieval call binding the contract method 0xa0e3a10e.
//
// Solidity: function CONTRACT_EXCHANGE() view returns(uint256)
func (_AssetboxImpl *AssetboxImplSession) CONTRACTEXCHANGE() (*big.Int, error) {
	return _AssetboxImpl.Contract.CONTRACTEXCHANGE(&_AssetboxImpl.CallOpts)
}

// CONTRACTEXCHANGE is a free data retrieval call binding the contract method 0xa0e3a10e.
//
// Solidity: function CONTRACT_EXCHANGE() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCallerSession) CONTRACTEXCHANGE() (*big.Int, error) {
	return _AssetboxImpl.Contract.CONTRACTEXCHANGE(&_AssetboxImpl.CallOpts)
}

// CONTRACTEXCHANGESTORAGE is a free data retrieval call binding the contract method 0x001570f2.
//
// Solidity: function CONTRACT_EXCHANGE_STORAGE() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCaller) CONTRACTEXCHANGESTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxImpl.contract.Call(opts, out, "CONTRACT_EXCHANGE_STORAGE")
	return *ret0, err
}

// CONTRACTEXCHANGESTORAGE is a free data retrieval call binding the contract method 0x001570f2.
//
// Solidity: function CONTRACT_EXCHANGE_STORAGE() view returns(uint256)
func (_AssetboxImpl *AssetboxImplSession) CONTRACTEXCHANGESTORAGE() (*big.Int, error) {
	return _AssetboxImpl.Contract.CONTRACTEXCHANGESTORAGE(&_AssetboxImpl.CallOpts)
}

// CONTRACTEXCHANGESTORAGE is a free data retrieval call binding the contract method 0x001570f2.
//
// Solidity: function CONTRACT_EXCHANGE_STORAGE() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCallerSession) CONTRACTEXCHANGESTORAGE() (*big.Int, error) {
	return _AssetboxImpl.Contract.CONTRACTEXCHANGESTORAGE(&_AssetboxImpl.CallOpts)
}

// CONTRACTFEESTORAGE is a free data retrieval call binding the contract method 0xcd5e65a4.
//
// Solidity: function CONTRACT_FEE_STORAGE() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCaller) CONTRACTFEESTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxImpl.contract.Call(opts, out, "CONTRACT_FEE_STORAGE")
	return *ret0, err
}

// CONTRACTFEESTORAGE is a free data retrieval call binding the contract method 0xcd5e65a4.
//
// Solidity: function CONTRACT_FEE_STORAGE() view returns(uint256)
func (_AssetboxImpl *AssetboxImplSession) CONTRACTFEESTORAGE() (*big.Int, error) {
	return _AssetboxImpl.Contract.CONTRACTFEESTORAGE(&_AssetboxImpl.CallOpts)
}

// CONTRACTFEESTORAGE is a free data retrieval call binding the contract method 0xcd5e65a4.
//
// Solidity: function CONTRACT_FEE_STORAGE() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCallerSession) CONTRACTFEESTORAGE() (*big.Int, error) {
	return _AssetboxImpl.Contract.CONTRACTFEESTORAGE(&_AssetboxImpl.CallOpts)
}

// CONTRACTGENERATOR is a free data retrieval call binding the contract method 0xccea2b46.
//
// Solidity: function CONTRACT_GENERATOR() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCaller) CONTRACTGENERATOR(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxImpl.contract.Call(opts, out, "CONTRACT_GENERATOR")
	return *ret0, err
}

// CONTRACTGENERATOR is a free data retrieval call binding the contract method 0xccea2b46.
//
// Solidity: function CONTRACT_GENERATOR() view returns(uint256)
func (_AssetboxImpl *AssetboxImplSession) CONTRACTGENERATOR() (*big.Int, error) {
	return _AssetboxImpl.Contract.CONTRACTGENERATOR(&_AssetboxImpl.CallOpts)
}

// CONTRACTGENERATOR is a free data retrieval call binding the contract method 0xccea2b46.
//
// Solidity: function CONTRACT_GENERATOR() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCallerSession) CONTRACTGENERATOR() (*big.Int, error) {
	return _AssetboxImpl.Contract.CONTRACTGENERATOR(&_AssetboxImpl.CallOpts)
}

// CONTRACTMININGAGENTSTORAGE is a free data retrieval call binding the contract method 0x83cc588d.
//
// Solidity: function CONTRACT_MINING_AGENT_STORAGE() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCaller) CONTRACTMININGAGENTSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxImpl.contract.Call(opts, out, "CONTRACT_MINING_AGENT_STORAGE")
	return *ret0, err
}

// CONTRACTMININGAGENTSTORAGE is a free data retrieval call binding the contract method 0x83cc588d.
//
// Solidity: function CONTRACT_MINING_AGENT_STORAGE() view returns(uint256)
func (_AssetboxImpl *AssetboxImplSession) CONTRACTMININGAGENTSTORAGE() (*big.Int, error) {
	return _AssetboxImpl.Contract.CONTRACTMININGAGENTSTORAGE(&_AssetboxImpl.CallOpts)
}

// CONTRACTMININGAGENTSTORAGE is a free data retrieval call binding the contract method 0x83cc588d.
//
// Solidity: function CONTRACT_MINING_AGENT_STORAGE() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCallerSession) CONTRACTMININGAGENTSTORAGE() (*big.Int, error) {
	return _AssetboxImpl.Contract.CONTRACTMININGAGENTSTORAGE(&_AssetboxImpl.CallOpts)
}

// CONTRACTMSBON is a free data retrieval call binding the contract method 0x28b3a976.
//
// Solidity: function CONTRACT_MSBON() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCaller) CONTRACTMSBON(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxImpl.contract.Call(opts, out, "CONTRACT_MSBON")
	return *ret0, err
}

// CONTRACTMSBON is a free data retrieval call binding the contract method 0x28b3a976.
//
// Solidity: function CONTRACT_MSBON() view returns(uint256)
func (_AssetboxImpl *AssetboxImplSession) CONTRACTMSBON() (*big.Int, error) {
	return _AssetboxImpl.Contract.CONTRACTMSBON(&_AssetboxImpl.CallOpts)
}

// CONTRACTMSBON is a free data retrieval call binding the contract method 0x28b3a976.
//
// Solidity: function CONTRACT_MSBON() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCallerSession) CONTRACTMSBON() (*big.Int, error) {
	return _AssetboxImpl.Contract.CONTRACTMSBON(&_AssetboxImpl.CallOpts)
}

// CONTRACTMSBONSTORAGE is a free data retrieval call binding the contract method 0xe04e6061.
//
// Solidity: function CONTRACT_MSBON_STORAGE() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCaller) CONTRACTMSBONSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxImpl.contract.Call(opts, out, "CONTRACT_MSBON_STORAGE")
	return *ret0, err
}

// CONTRACTMSBONSTORAGE is a free data retrieval call binding the contract method 0xe04e6061.
//
// Solidity: function CONTRACT_MSBON_STORAGE() view returns(uint256)
func (_AssetboxImpl *AssetboxImplSession) CONTRACTMSBONSTORAGE() (*big.Int, error) {
	return _AssetboxImpl.Contract.CONTRACTMSBONSTORAGE(&_AssetboxImpl.CallOpts)
}

// CONTRACTMSBONSTORAGE is a free data retrieval call binding the contract method 0xe04e6061.
//
// Solidity: function CONTRACT_MSBON_STORAGE() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCallerSession) CONTRACTMSBONSTORAGE() (*big.Int, error) {
	return _AssetboxImpl.Contract.CONTRACTMSBONSTORAGE(&_AssetboxImpl.CallOpts)
}

// CONTRACTMULTISIGADDADDRESS is a free data retrieval call binding the contract method 0xf4f65a5a.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ADDRESS() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCaller) CONTRACTMULTISIGADDADDRESS(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_ADD_ADDRESS")
	return *ret0, err
}

// CONTRACTMULTISIGADDADDRESS is a free data retrieval call binding the contract method 0xf4f65a5a.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ADDRESS() view returns(uint256)
func (_AssetboxImpl *AssetboxImplSession) CONTRACTMULTISIGADDADDRESS() (*big.Int, error) {
	return _AssetboxImpl.Contract.CONTRACTMULTISIGADDADDRESS(&_AssetboxImpl.CallOpts)
}

// CONTRACTMULTISIGADDADDRESS is a free data retrieval call binding the contract method 0xf4f65a5a.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ADDRESS() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCallerSession) CONTRACTMULTISIGADDADDRESS() (*big.Int, error) {
	return _AssetboxImpl.Contract.CONTRACTMULTISIGADDADDRESS(&_AssetboxImpl.CallOpts)
}

// CONTRACTMULTISIGADDADMIN is a free data retrieval call binding the contract method 0x83888c68.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ADMIN() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCaller) CONTRACTMULTISIGADDADMIN(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_ADD_ADMIN")
	return *ret0, err
}

// CONTRACTMULTISIGADDADMIN is a free data retrieval call binding the contract method 0x83888c68.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ADMIN() view returns(uint256)
func (_AssetboxImpl *AssetboxImplSession) CONTRACTMULTISIGADDADMIN() (*big.Int, error) {
	return _AssetboxImpl.Contract.CONTRACTMULTISIGADDADMIN(&_AssetboxImpl.CallOpts)
}

// CONTRACTMULTISIGADDADMIN is a free data retrieval call binding the contract method 0x83888c68.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ADMIN() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCallerSession) CONTRACTMULTISIGADDADMIN() (*big.Int, error) {
	return _AssetboxImpl.Contract.CONTRACTMULTISIGADDADMIN(&_AssetboxImpl.CallOpts)
}

// CONTRACTMULTISIGADDFEE is a free data retrieval call binding the contract method 0xc43e8cab.
//
// Solidity: function CONTRACT_MULTISIG_ADD_FEE() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCaller) CONTRACTMULTISIGADDFEE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_ADD_FEE")
	return *ret0, err
}

// CONTRACTMULTISIGADDFEE is a free data retrieval call binding the contract method 0xc43e8cab.
//
// Solidity: function CONTRACT_MULTISIG_ADD_FEE() view returns(uint256)
func (_AssetboxImpl *AssetboxImplSession) CONTRACTMULTISIGADDFEE() (*big.Int, error) {
	return _AssetboxImpl.Contract.CONTRACTMULTISIGADDFEE(&_AssetboxImpl.CallOpts)
}

// CONTRACTMULTISIGADDFEE is a free data retrieval call binding the contract method 0xc43e8cab.
//
// Solidity: function CONTRACT_MULTISIG_ADD_FEE() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCallerSession) CONTRACTMULTISIGADDFEE() (*big.Int, error) {
	return _AssetboxImpl.Contract.CONTRACTMULTISIGADDFEE(&_AssetboxImpl.CallOpts)
}

// CONTRACTMULTISIGADDROLE is a free data retrieval call binding the contract method 0x20a0914b.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ROLE() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCaller) CONTRACTMULTISIGADDROLE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_ADD_ROLE")
	return *ret0, err
}

// CONTRACTMULTISIGADDROLE is a free data retrieval call binding the contract method 0x20a0914b.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ROLE() view returns(uint256)
func (_AssetboxImpl *AssetboxImplSession) CONTRACTMULTISIGADDROLE() (*big.Int, error) {
	return _AssetboxImpl.Contract.CONTRACTMULTISIGADDROLE(&_AssetboxImpl.CallOpts)
}

// CONTRACTMULTISIGADDROLE is a free data retrieval call binding the contract method 0x20a0914b.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ROLE() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCallerSession) CONTRACTMULTISIGADDROLE() (*big.Int, error) {
	return _AssetboxImpl.Contract.CONTRACTMULTISIGADDROLE(&_AssetboxImpl.CallOpts)
}

// CONTRACTMULTISIGDROPPERMISSION is a free data retrieval call binding the contract method 0x3f3fbc39.
//
// Solidity: function CONTRACT_MULTISIG_DROP_PERMISSION() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCaller) CONTRACTMULTISIGDROPPERMISSION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_DROP_PERMISSION")
	return *ret0, err
}

// CONTRACTMULTISIGDROPPERMISSION is a free data retrieval call binding the contract method 0x3f3fbc39.
//
// Solidity: function CONTRACT_MULTISIG_DROP_PERMISSION() view returns(uint256)
func (_AssetboxImpl *AssetboxImplSession) CONTRACTMULTISIGDROPPERMISSION() (*big.Int, error) {
	return _AssetboxImpl.Contract.CONTRACTMULTISIGDROPPERMISSION(&_AssetboxImpl.CallOpts)
}

// CONTRACTMULTISIGDROPPERMISSION is a free data retrieval call binding the contract method 0x3f3fbc39.
//
// Solidity: function CONTRACT_MULTISIG_DROP_PERMISSION() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCallerSession) CONTRACTMULTISIGDROPPERMISSION() (*big.Int, error) {
	return _AssetboxImpl.Contract.CONTRACTMULTISIGDROPPERMISSION(&_AssetboxImpl.CallOpts)
}

// CONTRACTMULTISIGEDITADMIN is a free data retrieval call binding the contract method 0x91b0f051.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_ADMIN() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCaller) CONTRACTMULTISIGEDITADMIN(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_EDIT_ADMIN")
	return *ret0, err
}

// CONTRACTMULTISIGEDITADMIN is a free data retrieval call binding the contract method 0x91b0f051.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_ADMIN() view returns(uint256)
func (_AssetboxImpl *AssetboxImplSession) CONTRACTMULTISIGEDITADMIN() (*big.Int, error) {
	return _AssetboxImpl.Contract.CONTRACTMULTISIGEDITADMIN(&_AssetboxImpl.CallOpts)
}

// CONTRACTMULTISIGEDITADMIN is a free data retrieval call binding the contract method 0x91b0f051.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_ADMIN() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCallerSession) CONTRACTMULTISIGEDITADMIN() (*big.Int, error) {
	return _AssetboxImpl.Contract.CONTRACTMULTISIGEDITADMIN(&_AssetboxImpl.CallOpts)
}

// CONTRACTMULTISIGEDITFEE is a free data retrieval call binding the contract method 0x6b2abe17.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_FEE() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCaller) CONTRACTMULTISIGEDITFEE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_EDIT_FEE")
	return *ret0, err
}

// CONTRACTMULTISIGEDITFEE is a free data retrieval call binding the contract method 0x6b2abe17.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_FEE() view returns(uint256)
func (_AssetboxImpl *AssetboxImplSession) CONTRACTMULTISIGEDITFEE() (*big.Int, error) {
	return _AssetboxImpl.Contract.CONTRACTMULTISIGEDITFEE(&_AssetboxImpl.CallOpts)
}

// CONTRACTMULTISIGEDITFEE is a free data retrieval call binding the contract method 0x6b2abe17.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_FEE() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCallerSession) CONTRACTMULTISIGEDITFEE() (*big.Int, error) {
	return _AssetboxImpl.Contract.CONTRACTMULTISIGEDITFEE(&_AssetboxImpl.CallOpts)
}

// CONTRACTMULTISIGEDITROLE is a free data retrieval call binding the contract method 0xc5f66412.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_ROLE() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCaller) CONTRACTMULTISIGEDITROLE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_EDIT_ROLE")
	return *ret0, err
}

// CONTRACTMULTISIGEDITROLE is a free data retrieval call binding the contract method 0xc5f66412.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_ROLE() view returns(uint256)
func (_AssetboxImpl *AssetboxImplSession) CONTRACTMULTISIGEDITROLE() (*big.Int, error) {
	return _AssetboxImpl.Contract.CONTRACTMULTISIGEDITROLE(&_AssetboxImpl.CallOpts)
}

// CONTRACTMULTISIGEDITROLE is a free data retrieval call binding the contract method 0xc5f66412.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_ROLE() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCallerSession) CONTRACTMULTISIGEDITROLE() (*big.Int, error) {
	return _AssetboxImpl.Contract.CONTRACTMULTISIGEDITROLE(&_AssetboxImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEADMINBYADDRESS is a free data retrieval call binding the contract method 0xc816b091.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_BY_ADDRESS() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCaller) CONTRACTMULTISIGREMOVEADMINBYADDRESS(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REMOVE_ADMIN_BY_ADDRESS")
	return *ret0, err
}

// CONTRACTMULTISIGREMOVEADMINBYADDRESS is a free data retrieval call binding the contract method 0xc816b091.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_BY_ADDRESS() view returns(uint256)
func (_AssetboxImpl *AssetboxImplSession) CONTRACTMULTISIGREMOVEADMINBYADDRESS() (*big.Int, error) {
	return _AssetboxImpl.Contract.CONTRACTMULTISIGREMOVEADMINBYADDRESS(&_AssetboxImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEADMINBYADDRESS is a free data retrieval call binding the contract method 0xc816b091.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_BY_ADDRESS() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCallerSession) CONTRACTMULTISIGREMOVEADMINBYADDRESS() (*big.Int, error) {
	return _AssetboxImpl.Contract.CONTRACTMULTISIGREMOVEADMINBYADDRESS(&_AssetboxImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEADMINBYINDEX is a free data retrieval call binding the contract method 0x751fbea2.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_BY_INDEX() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCaller) CONTRACTMULTISIGREMOVEADMINBYINDEX(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REMOVE_ADMIN_BY_INDEX")
	return *ret0, err
}

// CONTRACTMULTISIGREMOVEADMINBYINDEX is a free data retrieval call binding the contract method 0x751fbea2.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_BY_INDEX() view returns(uint256)
func (_AssetboxImpl *AssetboxImplSession) CONTRACTMULTISIGREMOVEADMINBYINDEX() (*big.Int, error) {
	return _AssetboxImpl.Contract.CONTRACTMULTISIGREMOVEADMINBYINDEX(&_AssetboxImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEADMINBYINDEX is a free data retrieval call binding the contract method 0x751fbea2.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_BY_INDEX() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCallerSession) CONTRACTMULTISIGREMOVEADMINBYINDEX() (*big.Int, error) {
	return _AssetboxImpl.Contract.CONTRACTMULTISIGREMOVEADMINBYINDEX(&_AssetboxImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEADMINROLE is a free data retrieval call binding the contract method 0x53c00200.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_ROLE() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCaller) CONTRACTMULTISIGREMOVEADMINROLE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REMOVE_ADMIN_ROLE")
	return *ret0, err
}

// CONTRACTMULTISIGREMOVEADMINROLE is a free data retrieval call binding the contract method 0x53c00200.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_ROLE() view returns(uint256)
func (_AssetboxImpl *AssetboxImplSession) CONTRACTMULTISIGREMOVEADMINROLE() (*big.Int, error) {
	return _AssetboxImpl.Contract.CONTRACTMULTISIGREMOVEADMINROLE(&_AssetboxImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEADMINROLE is a free data retrieval call binding the contract method 0x53c00200.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_ROLE() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCallerSession) CONTRACTMULTISIGREMOVEADMINROLE() (*big.Int, error) {
	return _AssetboxImpl.Contract.CONTRACTMULTISIGREMOVEADMINROLE(&_AssetboxImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEFEE is a free data retrieval call binding the contract method 0xdb614d25.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_FEE() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCaller) CONTRACTMULTISIGREMOVEFEE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REMOVE_FEE")
	return *ret0, err
}

// CONTRACTMULTISIGREMOVEFEE is a free data retrieval call binding the contract method 0xdb614d25.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_FEE() view returns(uint256)
func (_AssetboxImpl *AssetboxImplSession) CONTRACTMULTISIGREMOVEFEE() (*big.Int, error) {
	return _AssetboxImpl.Contract.CONTRACTMULTISIGREMOVEFEE(&_AssetboxImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEFEE is a free data retrieval call binding the contract method 0xdb614d25.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_FEE() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCallerSession) CONTRACTMULTISIGREMOVEFEE() (*big.Int, error) {
	return _AssetboxImpl.Contract.CONTRACTMULTISIGREMOVEFEE(&_AssetboxImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEROLE is a free data retrieval call binding the contract method 0x23648496.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ROLE() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCaller) CONTRACTMULTISIGREMOVEROLE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REMOVE_ROLE")
	return *ret0, err
}

// CONTRACTMULTISIGREMOVEROLE is a free data retrieval call binding the contract method 0x23648496.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ROLE() view returns(uint256)
func (_AssetboxImpl *AssetboxImplSession) CONTRACTMULTISIGREMOVEROLE() (*big.Int, error) {
	return _AssetboxImpl.Contract.CONTRACTMULTISIGREMOVEROLE(&_AssetboxImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEROLE is a free data retrieval call binding the contract method 0x23648496.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ROLE() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCallerSession) CONTRACTMULTISIGREMOVEROLE() (*big.Int, error) {
	return _AssetboxImpl.Contract.CONTRACTMULTISIGREMOVEROLE(&_AssetboxImpl.CallOpts)
}

// CONTRACTMULTISIGREPLENISHCAPITALIZATION is a free data retrieval call binding the contract method 0x62d5d523.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_CAPITALIZATION() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCaller) CONTRACTMULTISIGREPLENISHCAPITALIZATION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REPLENISH_CAPITALIZATION")
	return *ret0, err
}

// CONTRACTMULTISIGREPLENISHCAPITALIZATION is a free data retrieval call binding the contract method 0x62d5d523.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_CAPITALIZATION() view returns(uint256)
func (_AssetboxImpl *AssetboxImplSession) CONTRACTMULTISIGREPLENISHCAPITALIZATION() (*big.Int, error) {
	return _AssetboxImpl.Contract.CONTRACTMULTISIGREPLENISHCAPITALIZATION(&_AssetboxImpl.CallOpts)
}

// CONTRACTMULTISIGREPLENISHCAPITALIZATION is a free data retrieval call binding the contract method 0x62d5d523.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_CAPITALIZATION() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCallerSession) CONTRACTMULTISIGREPLENISHCAPITALIZATION() (*big.Int, error) {
	return _AssetboxImpl.Contract.CONTRACTMULTISIGREPLENISHCAPITALIZATION(&_AssetboxImpl.CallOpts)
}

// CONTRACTMULTISIGREPLENISHCERTIFICATE is a free data retrieval call binding the contract method 0xfd19b7e1.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_CERTIFICATE() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCaller) CONTRACTMULTISIGREPLENISHCERTIFICATE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REPLENISH_CERTIFICATE")
	return *ret0, err
}

// CONTRACTMULTISIGREPLENISHCERTIFICATE is a free data retrieval call binding the contract method 0xfd19b7e1.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_CERTIFICATE() view returns(uint256)
func (_AssetboxImpl *AssetboxImplSession) CONTRACTMULTISIGREPLENISHCERTIFICATE() (*big.Int, error) {
	return _AssetboxImpl.Contract.CONTRACTMULTISIGREPLENISHCERTIFICATE(&_AssetboxImpl.CallOpts)
}

// CONTRACTMULTISIGREPLENISHCERTIFICATE is a free data retrieval call binding the contract method 0xfd19b7e1.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_CERTIFICATE() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCallerSession) CONTRACTMULTISIGREPLENISHCERTIFICATE() (*big.Int, error) {
	return _AssetboxImpl.Contract.CONTRACTMULTISIGREPLENISHCERTIFICATE(&_AssetboxImpl.CallOpts)
}

// CONTRACTMULTISIGREPLENISHEMISSION is a free data retrieval call binding the contract method 0x0cb59fde.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_EMISSION() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCaller) CONTRACTMULTISIGREPLENISHEMISSION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REPLENISH_EMISSION")
	return *ret0, err
}

// CONTRACTMULTISIGREPLENISHEMISSION is a free data retrieval call binding the contract method 0x0cb59fde.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_EMISSION() view returns(uint256)
func (_AssetboxImpl *AssetboxImplSession) CONTRACTMULTISIGREPLENISHEMISSION() (*big.Int, error) {
	return _AssetboxImpl.Contract.CONTRACTMULTISIGREPLENISHEMISSION(&_AssetboxImpl.CallOpts)
}

// CONTRACTMULTISIGREPLENISHEMISSION is a free data retrieval call binding the contract method 0x0cb59fde.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_EMISSION() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCallerSession) CONTRACTMULTISIGREPLENISHEMISSION() (*big.Int, error) {
	return _AssetboxImpl.Contract.CONTRACTMULTISIGREPLENISHEMISSION(&_AssetboxImpl.CallOpts)
}

// CONTRACTMULTISIGSETADMINROLE is a free data retrieval call binding the contract method 0x8f209ab7.
//
// Solidity: function CONTRACT_MULTISIG_SET_ADMIN_ROLE() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCaller) CONTRACTMULTISIGSETADMINROLE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_SET_ADMIN_ROLE")
	return *ret0, err
}

// CONTRACTMULTISIGSETADMINROLE is a free data retrieval call binding the contract method 0x8f209ab7.
//
// Solidity: function CONTRACT_MULTISIG_SET_ADMIN_ROLE() view returns(uint256)
func (_AssetboxImpl *AssetboxImplSession) CONTRACTMULTISIGSETADMINROLE() (*big.Int, error) {
	return _AssetboxImpl.Contract.CONTRACTMULTISIGSETADMINROLE(&_AssetboxImpl.CallOpts)
}

// CONTRACTMULTISIGSETADMINROLE is a free data retrieval call binding the contract method 0x8f209ab7.
//
// Solidity: function CONTRACT_MULTISIG_SET_ADMIN_ROLE() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCallerSession) CONTRACTMULTISIGSETADMINROLE() (*big.Int, error) {
	return _AssetboxImpl.Contract.CONTRACTMULTISIGSETADMINROLE(&_AssetboxImpl.CallOpts)
}

// CONTRACTMULTISIGSETPERMISSION is a free data retrieval call binding the contract method 0xb0e91914.
//
// Solidity: function CONTRACT_MULTISIG_SET_PERMISSION() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCaller) CONTRACTMULTISIGSETPERMISSION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_SET_PERMISSION")
	return *ret0, err
}

// CONTRACTMULTISIGSETPERMISSION is a free data retrieval call binding the contract method 0xb0e91914.
//
// Solidity: function CONTRACT_MULTISIG_SET_PERMISSION() view returns(uint256)
func (_AssetboxImpl *AssetboxImplSession) CONTRACTMULTISIGSETPERMISSION() (*big.Int, error) {
	return _AssetboxImpl.Contract.CONTRACTMULTISIGSETPERMISSION(&_AssetboxImpl.CallOpts)
}

// CONTRACTMULTISIGSETPERMISSION is a free data retrieval call binding the contract method 0xb0e91914.
//
// Solidity: function CONTRACT_MULTISIG_SET_PERMISSION() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCallerSession) CONTRACTMULTISIGSETPERMISSION() (*big.Int, error) {
	return _AssetboxImpl.Contract.CONTRACTMULTISIGSETPERMISSION(&_AssetboxImpl.CallOpts)
}

// CONTRACTMULTISIGSTORAGE is a free data retrieval call binding the contract method 0xe7e1a303.
//
// Solidity: function CONTRACT_MULTISIG_STORAGE() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCaller) CONTRACTMULTISIGSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_STORAGE")
	return *ret0, err
}

// CONTRACTMULTISIGSTORAGE is a free data retrieval call binding the contract method 0xe7e1a303.
//
// Solidity: function CONTRACT_MULTISIG_STORAGE() view returns(uint256)
func (_AssetboxImpl *AssetboxImplSession) CONTRACTMULTISIGSTORAGE() (*big.Int, error) {
	return _AssetboxImpl.Contract.CONTRACTMULTISIGSTORAGE(&_AssetboxImpl.CallOpts)
}

// CONTRACTMULTISIGSTORAGE is a free data retrieval call binding the contract method 0xe7e1a303.
//
// Solidity: function CONTRACT_MULTISIG_STORAGE() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCallerSession) CONTRACTMULTISIGSTORAGE() (*big.Int, error) {
	return _AssetboxImpl.Contract.CONTRACTMULTISIGSTORAGE(&_AssetboxImpl.CallOpts)
}

// CONTRACTOTC is a free data retrieval call binding the contract method 0x5eceac70.
//
// Solidity: function CONTRACT_OTC() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCaller) CONTRACTOTC(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxImpl.contract.Call(opts, out, "CONTRACT_OTC")
	return *ret0, err
}

// CONTRACTOTC is a free data retrieval call binding the contract method 0x5eceac70.
//
// Solidity: function CONTRACT_OTC() view returns(uint256)
func (_AssetboxImpl *AssetboxImplSession) CONTRACTOTC() (*big.Int, error) {
	return _AssetboxImpl.Contract.CONTRACTOTC(&_AssetboxImpl.CallOpts)
}

// CONTRACTOTC is a free data retrieval call binding the contract method 0x5eceac70.
//
// Solidity: function CONTRACT_OTC() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCallerSession) CONTRACTOTC() (*big.Int, error) {
	return _AssetboxImpl.Contract.CONTRACTOTC(&_AssetboxImpl.CallOpts)
}

// CONTRACTOTCSTORAGE is a free data retrieval call binding the contract method 0x825b476f.
//
// Solidity: function CONTRACT_OTC_STORAGE() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCaller) CONTRACTOTCSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxImpl.contract.Call(opts, out, "CONTRACT_OTC_STORAGE")
	return *ret0, err
}

// CONTRACTOTCSTORAGE is a free data retrieval call binding the contract method 0x825b476f.
//
// Solidity: function CONTRACT_OTC_STORAGE() view returns(uint256)
func (_AssetboxImpl *AssetboxImplSession) CONTRACTOTCSTORAGE() (*big.Int, error) {
	return _AssetboxImpl.Contract.CONTRACTOTCSTORAGE(&_AssetboxImpl.CallOpts)
}

// CONTRACTOTCSTORAGE is a free data retrieval call binding the contract method 0x825b476f.
//
// Solidity: function CONTRACT_OTC_STORAGE() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCallerSession) CONTRACTOTCSTORAGE() (*big.Int, error) {
	return _AssetboxImpl.Contract.CONTRACTOTCSTORAGE(&_AssetboxImpl.CallOpts)
}

// CONTRACTREPLENISHCAPITALIZATION is a free data retrieval call binding the contract method 0xff5e09d8.
//
// Solidity: function CONTRACT_REPLENISH_CAPITALIZATION() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCaller) CONTRACTREPLENISHCAPITALIZATION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxImpl.contract.Call(opts, out, "CONTRACT_REPLENISH_CAPITALIZATION")
	return *ret0, err
}

// CONTRACTREPLENISHCAPITALIZATION is a free data retrieval call binding the contract method 0xff5e09d8.
//
// Solidity: function CONTRACT_REPLENISH_CAPITALIZATION() view returns(uint256)
func (_AssetboxImpl *AssetboxImplSession) CONTRACTREPLENISHCAPITALIZATION() (*big.Int, error) {
	return _AssetboxImpl.Contract.CONTRACTREPLENISHCAPITALIZATION(&_AssetboxImpl.CallOpts)
}

// CONTRACTREPLENISHCAPITALIZATION is a free data retrieval call binding the contract method 0xff5e09d8.
//
// Solidity: function CONTRACT_REPLENISH_CAPITALIZATION() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCallerSession) CONTRACTREPLENISHCAPITALIZATION() (*big.Int, error) {
	return _AssetboxImpl.Contract.CONTRACTREPLENISHCAPITALIZATION(&_AssetboxImpl.CallOpts)
}

// CONTRACTRESERVEDALIASSTORAGE is a free data retrieval call binding the contract method 0x70a94aa7.
//
// Solidity: function CONTRACT_RESERVED_ALIAS_STORAGE() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCaller) CONTRACTRESERVEDALIASSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxImpl.contract.Call(opts, out, "CONTRACT_RESERVED_ALIAS_STORAGE")
	return *ret0, err
}

// CONTRACTRESERVEDALIASSTORAGE is a free data retrieval call binding the contract method 0x70a94aa7.
//
// Solidity: function CONTRACT_RESERVED_ALIAS_STORAGE() view returns(uint256)
func (_AssetboxImpl *AssetboxImplSession) CONTRACTRESERVEDALIASSTORAGE() (*big.Int, error) {
	return _AssetboxImpl.Contract.CONTRACTRESERVEDALIASSTORAGE(&_AssetboxImpl.CallOpts)
}

// CONTRACTRESERVEDALIASSTORAGE is a free data retrieval call binding the contract method 0x70a94aa7.
//
// Solidity: function CONTRACT_RESERVED_ALIAS_STORAGE() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCallerSession) CONTRACTRESERVEDALIASSTORAGE() (*big.Int, error) {
	return _AssetboxImpl.Contract.CONTRACTRESERVEDALIASSTORAGE(&_AssetboxImpl.CallOpts)
}

// CONTRACTROLESTORAGE is a free data retrieval call binding the contract method 0xb938235b.
//
// Solidity: function CONTRACT_ROLE_STORAGE() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCaller) CONTRACTROLESTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxImpl.contract.Call(opts, out, "CONTRACT_ROLE_STORAGE")
	return *ret0, err
}

// CONTRACTROLESTORAGE is a free data retrieval call binding the contract method 0xb938235b.
//
// Solidity: function CONTRACT_ROLE_STORAGE() view returns(uint256)
func (_AssetboxImpl *AssetboxImplSession) CONTRACTROLESTORAGE() (*big.Int, error) {
	return _AssetboxImpl.Contract.CONTRACTROLESTORAGE(&_AssetboxImpl.CallOpts)
}

// CONTRACTROLESTORAGE is a free data retrieval call binding the contract method 0xb938235b.
//
// Solidity: function CONTRACT_ROLE_STORAGE() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCallerSession) CONTRACTROLESTORAGE() (*big.Int, error) {
	return _AssetboxImpl.Contract.CONTRACTROLESTORAGE(&_AssetboxImpl.CallOpts)
}

// CONTRACTSAFETRANSFERSTORAGE is a free data retrieval call binding the contract method 0xe7e45392.
//
// Solidity: function CONTRACT_SAFE_TRANSFER_STORAGE() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCaller) CONTRACTSAFETRANSFERSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxImpl.contract.Call(opts, out, "CONTRACT_SAFE_TRANSFER_STORAGE")
	return *ret0, err
}

// CONTRACTSAFETRANSFERSTORAGE is a free data retrieval call binding the contract method 0xe7e45392.
//
// Solidity: function CONTRACT_SAFE_TRANSFER_STORAGE() view returns(uint256)
func (_AssetboxImpl *AssetboxImplSession) CONTRACTSAFETRANSFERSTORAGE() (*big.Int, error) {
	return _AssetboxImpl.Contract.CONTRACTSAFETRANSFERSTORAGE(&_AssetboxImpl.CallOpts)
}

// CONTRACTSAFETRANSFERSTORAGE is a free data retrieval call binding the contract method 0xe7e45392.
//
// Solidity: function CONTRACT_SAFE_TRANSFER_STORAGE() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCallerSession) CONTRACTSAFETRANSFERSTORAGE() (*big.Int, error) {
	return _AssetboxImpl.Contract.CONTRACTSAFETRANSFERSTORAGE(&_AssetboxImpl.CallOpts)
}

// CONTRACTTRANSFER is a free data retrieval call binding the contract method 0xef94fe38.
//
// Solidity: function CONTRACT_TRANSFER() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCaller) CONTRACTTRANSFER(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxImpl.contract.Call(opts, out, "CONTRACT_TRANSFER")
	return *ret0, err
}

// CONTRACTTRANSFER is a free data retrieval call binding the contract method 0xef94fe38.
//
// Solidity: function CONTRACT_TRANSFER() view returns(uint256)
func (_AssetboxImpl *AssetboxImplSession) CONTRACTTRANSFER() (*big.Int, error) {
	return _AssetboxImpl.Contract.CONTRACTTRANSFER(&_AssetboxImpl.CallOpts)
}

// CONTRACTTRANSFER is a free data retrieval call binding the contract method 0xef94fe38.
//
// Solidity: function CONTRACT_TRANSFER() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCallerSession) CONTRACTTRANSFER() (*big.Int, error) {
	return _AssetboxImpl.Contract.CONTRACTTRANSFER(&_AssetboxImpl.CallOpts)
}

// CONTRACTTRANSFERFROMCAPITALIZATION is a free data retrieval call binding the contract method 0xc8f8decb.
//
// Solidity: function CONTRACT_TRANSFER_FROM_CAPITALIZATION() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCaller) CONTRACTTRANSFERFROMCAPITALIZATION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxImpl.contract.Call(opts, out, "CONTRACT_TRANSFER_FROM_CAPITALIZATION")
	return *ret0, err
}

// CONTRACTTRANSFERFROMCAPITALIZATION is a free data retrieval call binding the contract method 0xc8f8decb.
//
// Solidity: function CONTRACT_TRANSFER_FROM_CAPITALIZATION() view returns(uint256)
func (_AssetboxImpl *AssetboxImplSession) CONTRACTTRANSFERFROMCAPITALIZATION() (*big.Int, error) {
	return _AssetboxImpl.Contract.CONTRACTTRANSFERFROMCAPITALIZATION(&_AssetboxImpl.CallOpts)
}

// CONTRACTTRANSFERFROMCAPITALIZATION is a free data retrieval call binding the contract method 0xc8f8decb.
//
// Solidity: function CONTRACT_TRANSFER_FROM_CAPITALIZATION() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCallerSession) CONTRACTTRANSFERFROMCAPITALIZATION() (*big.Int, error) {
	return _AssetboxImpl.Contract.CONTRACTTRANSFERFROMCAPITALIZATION(&_AssetboxImpl.CallOpts)
}

// PERMISSIONACCESSRESTORATION is a free data retrieval call binding the contract method 0x5c3053bf.
//
// Solidity: function PERMISSION_ACCESS_RESTORATION() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCaller) PERMISSIONACCESSRESTORATION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxImpl.contract.Call(opts, out, "PERMISSION_ACCESS_RESTORATION")
	return *ret0, err
}

// PERMISSIONACCESSRESTORATION is a free data retrieval call binding the contract method 0x5c3053bf.
//
// Solidity: function PERMISSION_ACCESS_RESTORATION() view returns(uint256)
func (_AssetboxImpl *AssetboxImplSession) PERMISSIONACCESSRESTORATION() (*big.Int, error) {
	return _AssetboxImpl.Contract.PERMISSIONACCESSRESTORATION(&_AssetboxImpl.CallOpts)
}

// PERMISSIONACCESSRESTORATION is a free data retrieval call binding the contract method 0x5c3053bf.
//
// Solidity: function PERMISSION_ACCESS_RESTORATION() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCallerSession) PERMISSIONACCESSRESTORATION() (*big.Int, error) {
	return _AssetboxImpl.Contract.PERMISSIONACCESSRESTORATION(&_AssetboxImpl.CallOpts)
}

// PERMISSIONADMINSTORAGE is a free data retrieval call binding the contract method 0x0c9fc0db.
//
// Solidity: function PERMISSION_ADMIN_STORAGE() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCaller) PERMISSIONADMINSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxImpl.contract.Call(opts, out, "PERMISSION_ADMIN_STORAGE")
	return *ret0, err
}

// PERMISSIONADMINSTORAGE is a free data retrieval call binding the contract method 0x0c9fc0db.
//
// Solidity: function PERMISSION_ADMIN_STORAGE() view returns(uint256)
func (_AssetboxImpl *AssetboxImplSession) PERMISSIONADMINSTORAGE() (*big.Int, error) {
	return _AssetboxImpl.Contract.PERMISSIONADMINSTORAGE(&_AssetboxImpl.CallOpts)
}

// PERMISSIONADMINSTORAGE is a free data retrieval call binding the contract method 0x0c9fc0db.
//
// Solidity: function PERMISSION_ADMIN_STORAGE() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCallerSession) PERMISSIONADMINSTORAGE() (*big.Int, error) {
	return _AssetboxImpl.Contract.PERMISSIONADMINSTORAGE(&_AssetboxImpl.CallOpts)
}

// PERMISSIONCONTRACTSSTORAGE is a free data retrieval call binding the contract method 0x1516e83b.
//
// Solidity: function PERMISSION_CONTRACTS_STORAGE() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCaller) PERMISSIONCONTRACTSSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxImpl.contract.Call(opts, out, "PERMISSION_CONTRACTS_STORAGE")
	return *ret0, err
}

// PERMISSIONCONTRACTSSTORAGE is a free data retrieval call binding the contract method 0x1516e83b.
//
// Solidity: function PERMISSION_CONTRACTS_STORAGE() view returns(uint256)
func (_AssetboxImpl *AssetboxImplSession) PERMISSIONCONTRACTSSTORAGE() (*big.Int, error) {
	return _AssetboxImpl.Contract.PERMISSIONCONTRACTSSTORAGE(&_AssetboxImpl.CallOpts)
}

// PERMISSIONCONTRACTSSTORAGE is a free data retrieval call binding the contract method 0x1516e83b.
//
// Solidity: function PERMISSION_CONTRACTS_STORAGE() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCallerSession) PERMISSIONCONTRACTSSTORAGE() (*big.Int, error) {
	return _AssetboxImpl.Contract.PERMISSIONCONTRACTSSTORAGE(&_AssetboxImpl.CallOpts)
}

// PERMISSIONCONTRACTDEPLOY is a free data retrieval call binding the contract method 0x257a4a04.
//
// Solidity: function PERMISSION_CONTRACT_DEPLOY() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCaller) PERMISSIONCONTRACTDEPLOY(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxImpl.contract.Call(opts, out, "PERMISSION_CONTRACT_DEPLOY")
	return *ret0, err
}

// PERMISSIONCONTRACTDEPLOY is a free data retrieval call binding the contract method 0x257a4a04.
//
// Solidity: function PERMISSION_CONTRACT_DEPLOY() view returns(uint256)
func (_AssetboxImpl *AssetboxImplSession) PERMISSIONCONTRACTDEPLOY() (*big.Int, error) {
	return _AssetboxImpl.Contract.PERMISSIONCONTRACTDEPLOY(&_AssetboxImpl.CallOpts)
}

// PERMISSIONCONTRACTDEPLOY is a free data retrieval call binding the contract method 0x257a4a04.
//
// Solidity: function PERMISSION_CONTRACT_DEPLOY() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCallerSession) PERMISSIONCONTRACTDEPLOY() (*big.Int, error) {
	return _AssetboxImpl.Contract.PERMISSIONCONTRACTDEPLOY(&_AssetboxImpl.CallOpts)
}

// PERMISSIONEMISSION is a free data retrieval call binding the contract method 0x4ccd92d1.
//
// Solidity: function PERMISSION_EMISSION() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCaller) PERMISSIONEMISSION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxImpl.contract.Call(opts, out, "PERMISSION_EMISSION")
	return *ret0, err
}

// PERMISSIONEMISSION is a free data retrieval call binding the contract method 0x4ccd92d1.
//
// Solidity: function PERMISSION_EMISSION() view returns(uint256)
func (_AssetboxImpl *AssetboxImplSession) PERMISSIONEMISSION() (*big.Int, error) {
	return _AssetboxImpl.Contract.PERMISSIONEMISSION(&_AssetboxImpl.CallOpts)
}

// PERMISSIONEMISSION is a free data retrieval call binding the contract method 0x4ccd92d1.
//
// Solidity: function PERMISSION_EMISSION() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCallerSession) PERMISSIONEMISSION() (*big.Int, error) {
	return _AssetboxImpl.Contract.PERMISSIONEMISSION(&_AssetboxImpl.CallOpts)
}

// PERMISSIONFEESSTORAGE is a free data retrieval call binding the contract method 0xe837eb87.
//
// Solidity: function PERMISSION_FEES_STORAGE() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCaller) PERMISSIONFEESSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxImpl.contract.Call(opts, out, "PERMISSION_FEES_STORAGE")
	return *ret0, err
}

// PERMISSIONFEESSTORAGE is a free data retrieval call binding the contract method 0xe837eb87.
//
// Solidity: function PERMISSION_FEES_STORAGE() view returns(uint256)
func (_AssetboxImpl *AssetboxImplSession) PERMISSIONFEESSTORAGE() (*big.Int, error) {
	return _AssetboxImpl.Contract.PERMISSIONFEESSTORAGE(&_AssetboxImpl.CallOpts)
}

// PERMISSIONFEESSTORAGE is a free data retrieval call binding the contract method 0xe837eb87.
//
// Solidity: function PERMISSION_FEES_STORAGE() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCallerSession) PERMISSIONFEESSTORAGE() (*big.Int, error) {
	return _AssetboxImpl.Contract.PERMISSIONFEESSTORAGE(&_AssetboxImpl.CallOpts)
}

// PERMISSIONMULTISIGSTORAGE is a free data retrieval call binding the contract method 0xf1bb7aad.
//
// Solidity: function PERMISSION_MULTISIG_STORAGE() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCaller) PERMISSIONMULTISIGSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxImpl.contract.Call(opts, out, "PERMISSION_MULTISIG_STORAGE")
	return *ret0, err
}

// PERMISSIONMULTISIGSTORAGE is a free data retrieval call binding the contract method 0xf1bb7aad.
//
// Solidity: function PERMISSION_MULTISIG_STORAGE() view returns(uint256)
func (_AssetboxImpl *AssetboxImplSession) PERMISSIONMULTISIGSTORAGE() (*big.Int, error) {
	return _AssetboxImpl.Contract.PERMISSIONMULTISIGSTORAGE(&_AssetboxImpl.CallOpts)
}

// PERMISSIONMULTISIGSTORAGE is a free data retrieval call binding the contract method 0xf1bb7aad.
//
// Solidity: function PERMISSION_MULTISIG_STORAGE() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCallerSession) PERMISSIONMULTISIGSTORAGE() (*big.Int, error) {
	return _AssetboxImpl.Contract.PERMISSIONMULTISIGSTORAGE(&_AssetboxImpl.CallOpts)
}

// PERMISSIONROLESSTORAGE is a free data retrieval call binding the contract method 0x12a16032.
//
// Solidity: function PERMISSION_ROLES_STORAGE() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCaller) PERMISSIONROLESSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxImpl.contract.Call(opts, out, "PERMISSION_ROLES_STORAGE")
	return *ret0, err
}

// PERMISSIONROLESSTORAGE is a free data retrieval call binding the contract method 0x12a16032.
//
// Solidity: function PERMISSION_ROLES_STORAGE() view returns(uint256)
func (_AssetboxImpl *AssetboxImplSession) PERMISSIONROLESSTORAGE() (*big.Int, error) {
	return _AssetboxImpl.Contract.PERMISSIONROLESSTORAGE(&_AssetboxImpl.CallOpts)
}

// PERMISSIONROLESSTORAGE is a free data retrieval call binding the contract method 0x12a16032.
//
// Solidity: function PERMISSION_ROLES_STORAGE() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCallerSession) PERMISSIONROLESSTORAGE() (*big.Int, error) {
	return _AssetboxImpl.Contract.PERMISSIONROLESSTORAGE(&_AssetboxImpl.CallOpts)
}

// ROLEACCESSVERIFIER is a free data retrieval call binding the contract method 0xc5405f92.
//
// Solidity: function ROLE_ACCESS_VERIFIER() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCaller) ROLEACCESSVERIFIER(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxImpl.contract.Call(opts, out, "ROLE_ACCESS_VERIFIER")
	return *ret0, err
}

// ROLEACCESSVERIFIER is a free data retrieval call binding the contract method 0xc5405f92.
//
// Solidity: function ROLE_ACCESS_VERIFIER() view returns(uint256)
func (_AssetboxImpl *AssetboxImplSession) ROLEACCESSVERIFIER() (*big.Int, error) {
	return _AssetboxImpl.Contract.ROLEACCESSVERIFIER(&_AssetboxImpl.CallOpts)
}

// ROLEACCESSVERIFIER is a free data retrieval call binding the contract method 0xc5405f92.
//
// Solidity: function ROLE_ACCESS_VERIFIER() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCallerSession) ROLEACCESSVERIFIER() (*big.Int, error) {
	return _AssetboxImpl.Contract.ROLEACCESSVERIFIER(&_AssetboxImpl.CallOpts)
}

// ROLEBITBONISSUEVERIFIER is a free data retrieval call binding the contract method 0xf0a012a6.
//
// Solidity: function ROLE_BITBON_ISSUE_VERIFIER() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCaller) ROLEBITBONISSUEVERIFIER(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxImpl.contract.Call(opts, out, "ROLE_BITBON_ISSUE_VERIFIER")
	return *ret0, err
}

// ROLEBITBONISSUEVERIFIER is a free data retrieval call binding the contract method 0xf0a012a6.
//
// Solidity: function ROLE_BITBON_ISSUE_VERIFIER() view returns(uint256)
func (_AssetboxImpl *AssetboxImplSession) ROLEBITBONISSUEVERIFIER() (*big.Int, error) {
	return _AssetboxImpl.Contract.ROLEBITBONISSUEVERIFIER(&_AssetboxImpl.CallOpts)
}

// ROLEBITBONISSUEVERIFIER is a free data retrieval call binding the contract method 0xf0a012a6.
//
// Solidity: function ROLE_BITBON_ISSUE_VERIFIER() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCallerSession) ROLEBITBONISSUEVERIFIER() (*big.Int, error) {
	return _AssetboxImpl.Contract.ROLEBITBONISSUEVERIFIER(&_AssetboxImpl.CallOpts)
}

// ROLECOMMISSIONVERIFIER is a free data retrieval call binding the contract method 0x8ab3de8c.
//
// Solidity: function ROLE_COMMISSION_VERIFIER() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCaller) ROLECOMMISSIONVERIFIER(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxImpl.contract.Call(opts, out, "ROLE_COMMISSION_VERIFIER")
	return *ret0, err
}

// ROLECOMMISSIONVERIFIER is a free data retrieval call binding the contract method 0x8ab3de8c.
//
// Solidity: function ROLE_COMMISSION_VERIFIER() view returns(uint256)
func (_AssetboxImpl *AssetboxImplSession) ROLECOMMISSIONVERIFIER() (*big.Int, error) {
	return _AssetboxImpl.Contract.ROLECOMMISSIONVERIFIER(&_AssetboxImpl.CallOpts)
}

// ROLECOMMISSIONVERIFIER is a free data retrieval call binding the contract method 0x8ab3de8c.
//
// Solidity: function ROLE_COMMISSION_VERIFIER() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCallerSession) ROLECOMMISSIONVERIFIER() (*big.Int, error) {
	return _AssetboxImpl.Contract.ROLECOMMISSIONVERIFIER(&_AssetboxImpl.CallOpts)
}

// ROLEDEPLOYADMIN is a free data retrieval call binding the contract method 0x667bf6e5.
//
// Solidity: function ROLE_DEPLOY_ADMIN() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCaller) ROLEDEPLOYADMIN(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxImpl.contract.Call(opts, out, "ROLE_DEPLOY_ADMIN")
	return *ret0, err
}

// ROLEDEPLOYADMIN is a free data retrieval call binding the contract method 0x667bf6e5.
//
// Solidity: function ROLE_DEPLOY_ADMIN() view returns(uint256)
func (_AssetboxImpl *AssetboxImplSession) ROLEDEPLOYADMIN() (*big.Int, error) {
	return _AssetboxImpl.Contract.ROLEDEPLOYADMIN(&_AssetboxImpl.CallOpts)
}

// ROLEDEPLOYADMIN is a free data retrieval call binding the contract method 0x667bf6e5.
//
// Solidity: function ROLE_DEPLOY_ADMIN() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCallerSession) ROLEDEPLOYADMIN() (*big.Int, error) {
	return _AssetboxImpl.Contract.ROLEDEPLOYADMIN(&_AssetboxImpl.CallOpts)
}

// ROLEPERMISSIONADMIN is a free data retrieval call binding the contract method 0x1b239996.
//
// Solidity: function ROLE_PERMISSION_ADMIN() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCaller) ROLEPERMISSIONADMIN(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxImpl.contract.Call(opts, out, "ROLE_PERMISSION_ADMIN")
	return *ret0, err
}

// ROLEPERMISSIONADMIN is a free data retrieval call binding the contract method 0x1b239996.
//
// Solidity: function ROLE_PERMISSION_ADMIN() view returns(uint256)
func (_AssetboxImpl *AssetboxImplSession) ROLEPERMISSIONADMIN() (*big.Int, error) {
	return _AssetboxImpl.Contract.ROLEPERMISSIONADMIN(&_AssetboxImpl.CallOpts)
}

// ROLEPERMISSIONADMIN is a free data retrieval call binding the contract method 0x1b239996.
//
// Solidity: function ROLE_PERMISSION_ADMIN() view returns(uint256)
func (_AssetboxImpl *AssetboxImplCallerSession) ROLEPERMISSIONADMIN() (*big.Int, error) {
	return _AssetboxImpl.Contract.ROLEPERMISSIONADMIN(&_AssetboxImpl.CallOpts)
}

// GetAssetboxInfoByAlias is a free data retrieval call binding the contract method 0xdd97a8fb.
//
// Solidity: function getAssetboxInfoByAlias(string assetboxAliasString) view returns(address assetboxAddress, string assetboxAlias, bytes serviceId, bool isMining, bytes otherInfo)
func (_AssetboxImpl *AssetboxImplCaller) GetAssetboxInfoByAlias(opts *bind.CallOpts, assetboxAliasString string) (struct {
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
	err := _AssetboxImpl.contract.Call(opts, out, "getAssetboxInfoByAlias", assetboxAliasString)
	return *ret, err
}

// GetAssetboxInfoByAlias is a free data retrieval call binding the contract method 0xdd97a8fb.
//
// Solidity: function getAssetboxInfoByAlias(string assetboxAliasString) view returns(address assetboxAddress, string assetboxAlias, bytes serviceId, bool isMining, bytes otherInfo)
func (_AssetboxImpl *AssetboxImplSession) GetAssetboxInfoByAlias(assetboxAliasString string) (struct {
	AssetboxAddress common.Address
	AssetboxAlias   string
	ServiceId       []byte
	IsMining        bool
	OtherInfo       []byte
}, error) {
	return _AssetboxImpl.Contract.GetAssetboxInfoByAlias(&_AssetboxImpl.CallOpts, assetboxAliasString)
}

// GetAssetboxInfoByAlias is a free data retrieval call binding the contract method 0xdd97a8fb.
//
// Solidity: function getAssetboxInfoByAlias(string assetboxAliasString) view returns(address assetboxAddress, string assetboxAlias, bytes serviceId, bool isMining, bytes otherInfo)
func (_AssetboxImpl *AssetboxImplCallerSession) GetAssetboxInfoByAlias(assetboxAliasString string) (struct {
	AssetboxAddress common.Address
	AssetboxAlias   string
	ServiceId       []byte
	IsMining        bool
	OtherInfo       []byte
}, error) {
	return _AssetboxImpl.Contract.GetAssetboxInfoByAlias(&_AssetboxImpl.CallOpts, assetboxAliasString)
}

// GetAssetboxInfoByAssetbox is a free data retrieval call binding the contract method 0x839e0fe1.
//
// Solidity: function getAssetboxInfoByAssetbox(address assetbox) view returns(address assetboxAddress, string assetboxAlias, bytes serviceId, bool isMining, bytes otherInfo)
func (_AssetboxImpl *AssetboxImplCaller) GetAssetboxInfoByAssetbox(opts *bind.CallOpts, assetbox common.Address) (struct {
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
	err := _AssetboxImpl.contract.Call(opts, out, "getAssetboxInfoByAssetbox", assetbox)
	return *ret, err
}

// GetAssetboxInfoByAssetbox is a free data retrieval call binding the contract method 0x839e0fe1.
//
// Solidity: function getAssetboxInfoByAssetbox(address assetbox) view returns(address assetboxAddress, string assetboxAlias, bytes serviceId, bool isMining, bytes otherInfo)
func (_AssetboxImpl *AssetboxImplSession) GetAssetboxInfoByAssetbox(assetbox common.Address) (struct {
	AssetboxAddress common.Address
	AssetboxAlias   string
	ServiceId       []byte
	IsMining        bool
	OtherInfo       []byte
}, error) {
	return _AssetboxImpl.Contract.GetAssetboxInfoByAssetbox(&_AssetboxImpl.CallOpts, assetbox)
}

// GetAssetboxInfoByAssetbox is a free data retrieval call binding the contract method 0x839e0fe1.
//
// Solidity: function getAssetboxInfoByAssetbox(address assetbox) view returns(address assetboxAddress, string assetboxAlias, bytes serviceId, bool isMining, bytes otherInfo)
func (_AssetboxImpl *AssetboxImplCallerSession) GetAssetboxInfoByAssetbox(assetbox common.Address) (struct {
	AssetboxAddress common.Address
	AssetboxAlias   string
	ServiceId       []byte
	IsMining        bool
	OtherInfo       []byte
}, error) {
	return _AssetboxImpl.Contract.GetAssetboxInfoByAssetbox(&_AssetboxImpl.CallOpts, assetbox)
}

// GetMyAssetboxInfo is a free data retrieval call binding the contract method 0xf9ee94d8.
//
// Solidity: function getMyAssetboxInfo() view returns(address assetboxAddress, string assetboxAlias, bytes serviceId, bool isMining, bytes otherInfo)
func (_AssetboxImpl *AssetboxImplCaller) GetMyAssetboxInfo(opts *bind.CallOpts) (struct {
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
	err := _AssetboxImpl.contract.Call(opts, out, "getMyAssetboxInfo")
	return *ret, err
}

// GetMyAssetboxInfo is a free data retrieval call binding the contract method 0xf9ee94d8.
//
// Solidity: function getMyAssetboxInfo() view returns(address assetboxAddress, string assetboxAlias, bytes serviceId, bool isMining, bytes otherInfo)
func (_AssetboxImpl *AssetboxImplSession) GetMyAssetboxInfo() (struct {
	AssetboxAddress common.Address
	AssetboxAlias   string
	ServiceId       []byte
	IsMining        bool
	OtherInfo       []byte
}, error) {
	return _AssetboxImpl.Contract.GetMyAssetboxInfo(&_AssetboxImpl.CallOpts)
}

// GetMyAssetboxInfo is a free data retrieval call binding the contract method 0xf9ee94d8.
//
// Solidity: function getMyAssetboxInfo() view returns(address assetboxAddress, string assetboxAlias, bytes serviceId, bool isMining, bytes otherInfo)
func (_AssetboxImpl *AssetboxImplCallerSession) GetMyAssetboxInfo() (struct {
	AssetboxAddress common.Address
	AssetboxAlias   string
	ServiceId       []byte
	IsMining        bool
	OtherInfo       []byte
}, error) {
	return _AssetboxImpl.Contract.GetMyAssetboxInfo(&_AssetboxImpl.CallOpts)
}

// GetThisContractIndex is a free data retrieval call binding the contract method 0x70c6d367.
//
// Solidity: function getThisContractIndex() pure returns(uint256)
func (_AssetboxImpl *AssetboxImplCaller) GetThisContractIndex(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetboxImpl.contract.Call(opts, out, "getThisContractIndex")
	return *ret0, err
}

// GetThisContractIndex is a free data retrieval call binding the contract method 0x70c6d367.
//
// Solidity: function getThisContractIndex() pure returns(uint256)
func (_AssetboxImpl *AssetboxImplSession) GetThisContractIndex() (*big.Int, error) {
	return _AssetboxImpl.Contract.GetThisContractIndex(&_AssetboxImpl.CallOpts)
}

// GetThisContractIndex is a free data retrieval call binding the contract method 0x70c6d367.
//
// Solidity: function getThisContractIndex() pure returns(uint256)
func (_AssetboxImpl *AssetboxImplCallerSession) GetThisContractIndex() (*big.Int, error) {
	return _AssetboxImpl.Contract.GetThisContractIndex(&_AssetboxImpl.CallOpts)
}

// IsAliasBoundToService is a free data retrieval call binding the contract method 0x29b08235.
//
// Solidity: function isAliasBoundToService(string assetboxAlias, bytes serviceId) view returns(bool isBound)
func (_AssetboxImpl *AssetboxImplCaller) IsAliasBoundToService(opts *bind.CallOpts, assetboxAlias string, serviceId []byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _AssetboxImpl.contract.Call(opts, out, "isAliasBoundToService", assetboxAlias, serviceId)
	return *ret0, err
}

// IsAliasBoundToService is a free data retrieval call binding the contract method 0x29b08235.
//
// Solidity: function isAliasBoundToService(string assetboxAlias, bytes serviceId) view returns(bool isBound)
func (_AssetboxImpl *AssetboxImplSession) IsAliasBoundToService(assetboxAlias string, serviceId []byte) (bool, error) {
	return _AssetboxImpl.Contract.IsAliasBoundToService(&_AssetboxImpl.CallOpts, assetboxAlias, serviceId)
}

// IsAliasBoundToService is a free data retrieval call binding the contract method 0x29b08235.
//
// Solidity: function isAliasBoundToService(string assetboxAlias, bytes serviceId) view returns(bool isBound)
func (_AssetboxImpl *AssetboxImplCallerSession) IsAliasBoundToService(assetboxAlias string, serviceId []byte) (bool, error) {
	return _AssetboxImpl.Contract.IsAliasBoundToService(&_AssetboxImpl.CallOpts, assetboxAlias, serviceId)
}

// IsAliasTaken is a free data retrieval call binding the contract method 0x7cf7252f.
//
// Solidity: function isAliasTaken(string assetboxAlias) view returns(bool taken)
func (_AssetboxImpl *AssetboxImplCaller) IsAliasTaken(opts *bind.CallOpts, assetboxAlias string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _AssetboxImpl.contract.Call(opts, out, "isAliasTaken", assetboxAlias)
	return *ret0, err
}

// IsAliasTaken is a free data retrieval call binding the contract method 0x7cf7252f.
//
// Solidity: function isAliasTaken(string assetboxAlias) view returns(bool taken)
func (_AssetboxImpl *AssetboxImplSession) IsAliasTaken(assetboxAlias string) (bool, error) {
	return _AssetboxImpl.Contract.IsAliasTaken(&_AssetboxImpl.CallOpts, assetboxAlias)
}

// IsAliasTaken is a free data retrieval call binding the contract method 0x7cf7252f.
//
// Solidity: function isAliasTaken(string assetboxAlias) view returns(bool taken)
func (_AssetboxImpl *AssetboxImplCallerSession) IsAliasTaken(assetboxAlias string) (bool, error) {
	return _AssetboxImpl.Contract.IsAliasTaken(&_AssetboxImpl.CallOpts, assetboxAlias)
}

// IsAssetboxBoundToService is a free data retrieval call binding the contract method 0x9ae049d6.
//
// Solidity: function isAssetboxBoundToService(address assetbox, bytes serviceId) view returns(bool isBound)
func (_AssetboxImpl *AssetboxImplCaller) IsAssetboxBoundToService(opts *bind.CallOpts, assetbox common.Address, serviceId []byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _AssetboxImpl.contract.Call(opts, out, "isAssetboxBoundToService", assetbox, serviceId)
	return *ret0, err
}

// IsAssetboxBoundToService is a free data retrieval call binding the contract method 0x9ae049d6.
//
// Solidity: function isAssetboxBoundToService(address assetbox, bytes serviceId) view returns(bool isBound)
func (_AssetboxImpl *AssetboxImplSession) IsAssetboxBoundToService(assetbox common.Address, serviceId []byte) (bool, error) {
	return _AssetboxImpl.Contract.IsAssetboxBoundToService(&_AssetboxImpl.CallOpts, assetbox, serviceId)
}

// IsAssetboxBoundToService is a free data retrieval call binding the contract method 0x9ae049d6.
//
// Solidity: function isAssetboxBoundToService(address assetbox, bytes serviceId) view returns(bool isBound)
func (_AssetboxImpl *AssetboxImplCallerSession) IsAssetboxBoundToService(assetbox common.Address, serviceId []byte) (bool, error) {
	return _AssetboxImpl.Contract.IsAssetboxBoundToService(&_AssetboxImpl.CallOpts, assetbox, serviceId)
}

// IsAssetboxMining is a free data retrieval call binding the contract method 0xc8d8a000.
//
// Solidity: function isAssetboxMining(address assetbox) view returns(bool)
func (_AssetboxImpl *AssetboxImplCaller) IsAssetboxMining(opts *bind.CallOpts, assetbox common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _AssetboxImpl.contract.Call(opts, out, "isAssetboxMining", assetbox)
	return *ret0, err
}

// IsAssetboxMining is a free data retrieval call binding the contract method 0xc8d8a000.
//
// Solidity: function isAssetboxMining(address assetbox) view returns(bool)
func (_AssetboxImpl *AssetboxImplSession) IsAssetboxMining(assetbox common.Address) (bool, error) {
	return _AssetboxImpl.Contract.IsAssetboxMining(&_AssetboxImpl.CallOpts, assetbox)
}

// IsAssetboxMining is a free data retrieval call binding the contract method 0xc8d8a000.
//
// Solidity: function isAssetboxMining(address assetbox) view returns(bool)
func (_AssetboxImpl *AssetboxImplCallerSession) IsAssetboxMining(assetbox common.Address) (bool, error) {
	return _AssetboxImpl.Contract.IsAssetboxMining(&_AssetboxImpl.CallOpts, assetbox)
}

// IsMyAssetboxBoundToService is a free data retrieval call binding the contract method 0x7d879734.
//
// Solidity: function isMyAssetboxBoundToService(bytes serviceId) view returns(bool isBound)
func (_AssetboxImpl *AssetboxImplCaller) IsMyAssetboxBoundToService(opts *bind.CallOpts, serviceId []byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _AssetboxImpl.contract.Call(opts, out, "isMyAssetboxBoundToService", serviceId)
	return *ret0, err
}

// IsMyAssetboxBoundToService is a free data retrieval call binding the contract method 0x7d879734.
//
// Solidity: function isMyAssetboxBoundToService(bytes serviceId) view returns(bool isBound)
func (_AssetboxImpl *AssetboxImplSession) IsMyAssetboxBoundToService(serviceId []byte) (bool, error) {
	return _AssetboxImpl.Contract.IsMyAssetboxBoundToService(&_AssetboxImpl.CallOpts, serviceId)
}

// IsMyAssetboxBoundToService is a free data retrieval call binding the contract method 0x7d879734.
//
// Solidity: function isMyAssetboxBoundToService(bytes serviceId) view returns(bool isBound)
func (_AssetboxImpl *AssetboxImplCallerSession) IsMyAssetboxBoundToService(serviceId []byte) (bool, error) {
	return _AssetboxImpl.Contract.IsMyAssetboxBoundToService(&_AssetboxImpl.CallOpts, serviceId)
}

// ApproveReservedAlias is a paid mutator transaction binding the contract method 0x3cdac2a7.
//
// Solidity: function approveReservedAlias(string aliasString, address assetbox) returns(bool)
func (_AssetboxImpl *AssetboxImplTransactor) ApproveReservedAlias(opts *bind.TransactOpts, aliasString string, assetbox common.Address) (*types.Transaction, error) {
	return _AssetboxImpl.contract.Transact(opts, "approveReservedAlias", aliasString, assetbox)
}

// ApproveReservedAlias is a paid mutator transaction binding the contract method 0x3cdac2a7.
//
// Solidity: function approveReservedAlias(string aliasString, address assetbox) returns(bool)
func (_AssetboxImpl *AssetboxImplSession) ApproveReservedAlias(aliasString string, assetbox common.Address) (*types.Transaction, error) {
	return _AssetboxImpl.Contract.ApproveReservedAlias(&_AssetboxImpl.TransactOpts, aliasString, assetbox)
}

// ApproveReservedAlias is a paid mutator transaction binding the contract method 0x3cdac2a7.
//
// Solidity: function approveReservedAlias(string aliasString, address assetbox) returns(bool)
func (_AssetboxImpl *AssetboxImplTransactorSession) ApproveReservedAlias(aliasString string, assetbox common.Address) (*types.Transaction, error) {
	return _AssetboxImpl.Contract.ApproveReservedAlias(&_AssetboxImpl.TransactOpts, aliasString, assetbox)
}

// DeleteAssetboxInfo is a paid mutator transaction binding the contract method 0x3624843b.
//
// Solidity: function deleteAssetboxInfo() returns()
func (_AssetboxImpl *AssetboxImplTransactor) DeleteAssetboxInfo(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AssetboxImpl.contract.Transact(opts, "deleteAssetboxInfo")
}

// DeleteAssetboxInfo is a paid mutator transaction binding the contract method 0x3624843b.
//
// Solidity: function deleteAssetboxInfo() returns()
func (_AssetboxImpl *AssetboxImplSession) DeleteAssetboxInfo() (*types.Transaction, error) {
	return _AssetboxImpl.Contract.DeleteAssetboxInfo(&_AssetboxImpl.TransactOpts)
}

// DeleteAssetboxInfo is a paid mutator transaction binding the contract method 0x3624843b.
//
// Solidity: function deleteAssetboxInfo() returns()
func (_AssetboxImpl *AssetboxImplTransactorSession) DeleteAssetboxInfo() (*types.Transaction, error) {
	return _AssetboxImpl.Contract.DeleteAssetboxInfo(&_AssetboxImpl.TransactOpts)
}

// DeleteAssetboxInfoAdmin is a paid mutator transaction binding the contract method 0x191bf945.
//
// Solidity: function deleteAssetboxInfoAdmin(address assetbox) returns()
func (_AssetboxImpl *AssetboxImplTransactor) DeleteAssetboxInfoAdmin(opts *bind.TransactOpts, assetbox common.Address) (*types.Transaction, error) {
	return _AssetboxImpl.contract.Transact(opts, "deleteAssetboxInfoAdmin", assetbox)
}

// DeleteAssetboxInfoAdmin is a paid mutator transaction binding the contract method 0x191bf945.
//
// Solidity: function deleteAssetboxInfoAdmin(address assetbox) returns()
func (_AssetboxImpl *AssetboxImplSession) DeleteAssetboxInfoAdmin(assetbox common.Address) (*types.Transaction, error) {
	return _AssetboxImpl.Contract.DeleteAssetboxInfoAdmin(&_AssetboxImpl.TransactOpts, assetbox)
}

// DeleteAssetboxInfoAdmin is a paid mutator transaction binding the contract method 0x191bf945.
//
// Solidity: function deleteAssetboxInfoAdmin(address assetbox) returns()
func (_AssetboxImpl *AssetboxImplTransactorSession) DeleteAssetboxInfoAdmin(assetbox common.Address) (*types.Transaction, error) {
	return _AssetboxImpl.Contract.DeleteAssetboxInfoAdmin(&_AssetboxImpl.TransactOpts, assetbox)
}

// EditReservedAlias is a paid mutator transaction binding the contract method 0x845d48a4.
//
// Solidity: function editReservedAlias(string oldAlias, string newAlias) returns(bool)
func (_AssetboxImpl *AssetboxImplTransactor) EditReservedAlias(opts *bind.TransactOpts, oldAlias string, newAlias string) (*types.Transaction, error) {
	return _AssetboxImpl.contract.Transact(opts, "editReservedAlias", oldAlias, newAlias)
}

// EditReservedAlias is a paid mutator transaction binding the contract method 0x845d48a4.
//
// Solidity: function editReservedAlias(string oldAlias, string newAlias) returns(bool)
func (_AssetboxImpl *AssetboxImplSession) EditReservedAlias(oldAlias string, newAlias string) (*types.Transaction, error) {
	return _AssetboxImpl.Contract.EditReservedAlias(&_AssetboxImpl.TransactOpts, oldAlias, newAlias)
}

// EditReservedAlias is a paid mutator transaction binding the contract method 0x845d48a4.
//
// Solidity: function editReservedAlias(string oldAlias, string newAlias) returns(bool)
func (_AssetboxImpl *AssetboxImplTransactorSession) EditReservedAlias(oldAlias string, newAlias string) (*types.Transaction, error) {
	return _AssetboxImpl.Contract.EditReservedAlias(&_AssetboxImpl.TransactOpts, oldAlias, newAlias)
}

// RemoveAliasReservation is a paid mutator transaction binding the contract method 0xef234326.
//
// Solidity: function removeAliasReservation(string aliasString) returns(bool)
func (_AssetboxImpl *AssetboxImplTransactor) RemoveAliasReservation(opts *bind.TransactOpts, aliasString string) (*types.Transaction, error) {
	return _AssetboxImpl.contract.Transact(opts, "removeAliasReservation", aliasString)
}

// RemoveAliasReservation is a paid mutator transaction binding the contract method 0xef234326.
//
// Solidity: function removeAliasReservation(string aliasString) returns(bool)
func (_AssetboxImpl *AssetboxImplSession) RemoveAliasReservation(aliasString string) (*types.Transaction, error) {
	return _AssetboxImpl.Contract.RemoveAliasReservation(&_AssetboxImpl.TransactOpts, aliasString)
}

// RemoveAliasReservation is a paid mutator transaction binding the contract method 0xef234326.
//
// Solidity: function removeAliasReservation(string aliasString) returns(bool)
func (_AssetboxImpl *AssetboxImplTransactorSession) RemoveAliasReservation(aliasString string) (*types.Transaction, error) {
	return _AssetboxImpl.Contract.RemoveAliasReservation(&_AssetboxImpl.TransactOpts, aliasString)
}

// RemoveAliasReservationAt is a paid mutator transaction binding the contract method 0x01e7b297.
//
// Solidity: function removeAliasReservationAt(uint256 index) returns(bool)
func (_AssetboxImpl *AssetboxImplTransactor) RemoveAliasReservationAt(opts *bind.TransactOpts, index *big.Int) (*types.Transaction, error) {
	return _AssetboxImpl.contract.Transact(opts, "removeAliasReservationAt", index)
}

// RemoveAliasReservationAt is a paid mutator transaction binding the contract method 0x01e7b297.
//
// Solidity: function removeAliasReservationAt(uint256 index) returns(bool)
func (_AssetboxImpl *AssetboxImplSession) RemoveAliasReservationAt(index *big.Int) (*types.Transaction, error) {
	return _AssetboxImpl.Contract.RemoveAliasReservationAt(&_AssetboxImpl.TransactOpts, index)
}

// RemoveAliasReservationAt is a paid mutator transaction binding the contract method 0x01e7b297.
//
// Solidity: function removeAliasReservationAt(uint256 index) returns(bool)
func (_AssetboxImpl *AssetboxImplTransactorSession) RemoveAliasReservationAt(index *big.Int) (*types.Transaction, error) {
	return _AssetboxImpl.Contract.RemoveAliasReservationAt(&_AssetboxImpl.TransactOpts, index)
}

// ReserveAlias is a paid mutator transaction binding the contract method 0x67f7eaba.
//
// Solidity: function reserveAlias(string aliasString) returns(bool)
func (_AssetboxImpl *AssetboxImplTransactor) ReserveAlias(opts *bind.TransactOpts, aliasString string) (*types.Transaction, error) {
	return _AssetboxImpl.contract.Transact(opts, "reserveAlias", aliasString)
}

// ReserveAlias is a paid mutator transaction binding the contract method 0x67f7eaba.
//
// Solidity: function reserveAlias(string aliasString) returns(bool)
func (_AssetboxImpl *AssetboxImplSession) ReserveAlias(aliasString string) (*types.Transaction, error) {
	return _AssetboxImpl.Contract.ReserveAlias(&_AssetboxImpl.TransactOpts, aliasString)
}

// ReserveAlias is a paid mutator transaction binding the contract method 0x67f7eaba.
//
// Solidity: function reserveAlias(string aliasString) returns(bool)
func (_AssetboxImpl *AssetboxImplTransactorSession) ReserveAlias(aliasString string) (*types.Transaction, error) {
	return _AssetboxImpl.Contract.ReserveAlias(&_AssetboxImpl.TransactOpts, aliasString)
}

// SetAssetboxInfo is a paid mutator transaction binding the contract method 0x8192fc54.
//
// Solidity: function setAssetboxInfo(string assetboxAlias, bytes serviceId, bool isMining, bytes otherInfo) returns()
func (_AssetboxImpl *AssetboxImplTransactor) SetAssetboxInfo(opts *bind.TransactOpts, assetboxAlias string, serviceId []byte, isMining bool, otherInfo []byte) (*types.Transaction, error) {
	return _AssetboxImpl.contract.Transact(opts, "setAssetboxInfo", assetboxAlias, serviceId, isMining, otherInfo)
}

// SetAssetboxInfo is a paid mutator transaction binding the contract method 0x8192fc54.
//
// Solidity: function setAssetboxInfo(string assetboxAlias, bytes serviceId, bool isMining, bytes otherInfo) returns()
func (_AssetboxImpl *AssetboxImplSession) SetAssetboxInfo(assetboxAlias string, serviceId []byte, isMining bool, otherInfo []byte) (*types.Transaction, error) {
	return _AssetboxImpl.Contract.SetAssetboxInfo(&_AssetboxImpl.TransactOpts, assetboxAlias, serviceId, isMining, otherInfo)
}

// SetAssetboxInfo is a paid mutator transaction binding the contract method 0x8192fc54.
//
// Solidity: function setAssetboxInfo(string assetboxAlias, bytes serviceId, bool isMining, bytes otherInfo) returns()
func (_AssetboxImpl *AssetboxImplTransactorSession) SetAssetboxInfo(assetboxAlias string, serviceId []byte, isMining bool, otherInfo []byte) (*types.Transaction, error) {
	return _AssetboxImpl.Contract.SetAssetboxInfo(&_AssetboxImpl.TransactOpts, assetboxAlias, serviceId, isMining, otherInfo)
}

// SetAssetboxInfoAdmin is a paid mutator transaction binding the contract method 0xbbff78a9.
//
// Solidity: function setAssetboxInfoAdmin(address assetbox, string assetboxAlias, bytes serviceId, bool isMining, bytes otherInfo) returns()
func (_AssetboxImpl *AssetboxImplTransactor) SetAssetboxInfoAdmin(opts *bind.TransactOpts, assetbox common.Address, assetboxAlias string, serviceId []byte, isMining bool, otherInfo []byte) (*types.Transaction, error) {
	return _AssetboxImpl.contract.Transact(opts, "setAssetboxInfoAdmin", assetbox, assetboxAlias, serviceId, isMining, otherInfo)
}

// SetAssetboxInfoAdmin is a paid mutator transaction binding the contract method 0xbbff78a9.
//
// Solidity: function setAssetboxInfoAdmin(address assetbox, string assetboxAlias, bytes serviceId, bool isMining, bytes otherInfo) returns()
func (_AssetboxImpl *AssetboxImplSession) SetAssetboxInfoAdmin(assetbox common.Address, assetboxAlias string, serviceId []byte, isMining bool, otherInfo []byte) (*types.Transaction, error) {
	return _AssetboxImpl.Contract.SetAssetboxInfoAdmin(&_AssetboxImpl.TransactOpts, assetbox, assetboxAlias, serviceId, isMining, otherInfo)
}

// SetAssetboxInfoAdmin is a paid mutator transaction binding the contract method 0xbbff78a9.
//
// Solidity: function setAssetboxInfoAdmin(address assetbox, string assetboxAlias, bytes serviceId, bool isMining, bytes otherInfo) returns()
func (_AssetboxImpl *AssetboxImplTransactorSession) SetAssetboxInfoAdmin(assetbox common.Address, assetboxAlias string, serviceId []byte, isMining bool, otherInfo []byte) (*types.Transaction, error) {
	return _AssetboxImpl.Contract.SetAssetboxInfoAdmin(&_AssetboxImpl.TransactOpts, assetbox, assetboxAlias, serviceId, isMining, otherInfo)
}

// SetAssetboxInfosAdmin is a paid mutator transaction binding the contract method 0xecc59e1e.
//
// Solidity: function setAssetboxInfosAdmin(address[] assetboxes, string[] assetboxAliases, bytes[] serviceIds, bool[] isMiningFlags, bytes[] otherInfos) returns()
func (_AssetboxImpl *AssetboxImplTransactor) SetAssetboxInfosAdmin(opts *bind.TransactOpts, assetboxes []common.Address, assetboxAliases []string, serviceIds [][]byte, isMiningFlags []bool, otherInfos [][]byte) (*types.Transaction, error) {
	return _AssetboxImpl.contract.Transact(opts, "setAssetboxInfosAdmin", assetboxes, assetboxAliases, serviceIds, isMiningFlags, otherInfos)
}

// SetAssetboxInfosAdmin is a paid mutator transaction binding the contract method 0xecc59e1e.
//
// Solidity: function setAssetboxInfosAdmin(address[] assetboxes, string[] assetboxAliases, bytes[] serviceIds, bool[] isMiningFlags, bytes[] otherInfos) returns()
func (_AssetboxImpl *AssetboxImplSession) SetAssetboxInfosAdmin(assetboxes []common.Address, assetboxAliases []string, serviceIds [][]byte, isMiningFlags []bool, otherInfos [][]byte) (*types.Transaction, error) {
	return _AssetboxImpl.Contract.SetAssetboxInfosAdmin(&_AssetboxImpl.TransactOpts, assetboxes, assetboxAliases, serviceIds, isMiningFlags, otherInfos)
}

// SetAssetboxInfosAdmin is a paid mutator transaction binding the contract method 0xecc59e1e.
//
// Solidity: function setAssetboxInfosAdmin(address[] assetboxes, string[] assetboxAliases, bytes[] serviceIds, bool[] isMiningFlags, bytes[] otherInfos) returns()
func (_AssetboxImpl *AssetboxImplTransactorSession) SetAssetboxInfosAdmin(assetboxes []common.Address, assetboxAliases []string, serviceIds [][]byte, isMiningFlags []bool, otherInfos [][]byte) (*types.Transaction, error) {
	return _AssetboxImpl.Contract.SetAssetboxInfosAdmin(&_AssetboxImpl.TransactOpts, assetboxes, assetboxAliases, serviceIds, isMiningFlags, otherInfos)
}
