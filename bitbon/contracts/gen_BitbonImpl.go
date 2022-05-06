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

// BitbonImplABI is the input ABI used to generate the binding from.
const BitbonImplABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractStorageAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ACCESS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ACCESS_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ADMINABLE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ADMIN_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ASSETBOX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ASSETBOX_INFO\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ASSETBOX_INFO_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ASSETBOX_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_BALANCE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_BITBON\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_BITBON_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_BITBON_SUPPORT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_BONBALANCE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_BONTRANSFER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_DEX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_DISTRIBUTION_GATE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_DISTRIBUTION_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_EXCHANGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_EXCHANGE_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_FEE_GATE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_FEE_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_GENERATOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MINING_AGENT_GATE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MINING_AGENT_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MSBON\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MSBON_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_ADD_ADDRESS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_ADD_ADMIN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_ADD_DISTRIBUTION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_ADD_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_ADD_ROLE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_CLEAR_FEE_EXCEPTIONS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_DROP_PERMISSION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_EDIT_ADMIN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_EDIT_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_EDIT_ROLE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REMOVE_ADMIN_BY_ADDRESS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REMOVE_ADMIN_BY_INDEX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REMOVE_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REMOVE_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REMOVE_ROLE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REPLENISH_CAPITALIZATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REPLENISH_CERTIFICATE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REPLENISH_EMISSION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_SET_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_SET_CONFIRMATION_RATE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_SET_FEE_DISTRIBUTION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_SET_FEE_EXCEPTIONS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_SET_MAX_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_SET_MIN_AMOUNT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_SET_MIN_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_SET_PERMISSION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_OTC\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_OTC_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_REPLENISH_CAPITALIZATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_RESERVED_ALIAS_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ROLE_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_SAFE_TRANSFER_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_TRANSFER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_TRANSFER_FROM_CAPITALIZATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"FEE_OPERATION_CREATE_SAFE_TRANSFER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"FEE_OPERATION_CREATE_SAFE_TRANSFER_ALL\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"FEE_OPERATION_FRAME_TRANSFER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"FEE_OPERATION_QUICK_TRANSFER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"FEE_OPERATION_QUICK_TRANSFER_ALL\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"FEE_OPERATION_WPC_SAFE_TRANSFER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"FEE_OPERATION_WPC_SAFE_TRANSFER_ALL\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_ACCESS_RESTORATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_ADMIN_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_CONTRACTS_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_CONTRACT_DEPLOY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_EMISSION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_FEES_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_MULTISIG_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_ROLES_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_ACCESS_VERIFIER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_BITBON_ISSUE_VERIFIER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_COMMISSION_VERIFIER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_DEPLOY_ADMIN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_PERMISSION_ADMIN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getThisContractIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"scope\",\"type\":\"address[]\"}],\"name\":\"getScopeBalances\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"scope\",\"type\":\"address[]\"}],\"name\":\"getScopeBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"res\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"opType\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"}],\"name\":\"feeTransfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"quickTransfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"quickTransferAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"frameTransfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"vk\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"transferId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"retries\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiresAt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"createSafeTransfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"vk\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"transferId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"retries\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiresAt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"createSafeTransferAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"transferId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"expiresAt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"createWPCSafeTransfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"transferId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"expiresAt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"createWPCSafeTransferAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"transferId\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"protectionCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"approveSafeTransfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"transferId\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"protectionCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"approveSafeTransferAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"transferId\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"approveWPCSafeTransfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"transferId\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"approveWPCSafeTransferAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"transferId\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"cancelSafeTransfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"transferId\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"cancelSafeTransferAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"transferId\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"cancelWPCSafeTransfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"transferId\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"cancelWPCSafeTransferAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"getOperationFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"assetbox\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"assetbox\",\"type\":\"address\"}],\"name\":\"balanceOfLocked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"assetbox\",\"type\":\"address\"}],\"name\":\"balanceOfFull\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"transferId\",\"type\":\"bytes\"}],\"name\":\"transferExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"ids\",\"type\":\"bytes32[]\"}],\"name\":\"expireSafeTransfers\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTotalLockedBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// BitbonImpl is an auto generated Go binding around an Ethereum contract.
type BitbonImpl struct {
	BitbonImplCaller     // Read-only binding to the contract
	BitbonImplTransactor // Write-only binding to the contract
	BitbonImplFilterer   // Log filterer for contract events
}

// BitbonImplCaller is an auto generated read-only Go binding around an Ethereum contract.
type BitbonImplCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BitbonImplTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BitbonImplTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BitbonImplFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BitbonImplFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BitbonImplSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BitbonImplSession struct {
	Contract     *BitbonImpl       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BitbonImplCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BitbonImplCallerSession struct {
	Contract *BitbonImplCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// BitbonImplTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BitbonImplTransactorSession struct {
	Contract     *BitbonImplTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// BitbonImplRaw is an auto generated low-level Go binding around an Ethereum contract.
type BitbonImplRaw struct {
	Contract *BitbonImpl // Generic contract binding to access the raw methods on
}

// BitbonImplCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BitbonImplCallerRaw struct {
	Contract *BitbonImplCaller // Generic read-only contract binding to access the raw methods on
}

// BitbonImplTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BitbonImplTransactorRaw struct {
	Contract *BitbonImplTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBitbonImpl creates a new instance of BitbonImpl, bound to a specific deployed contract.
func NewBitbonImpl(address common.Address, backend bind.ContractBackend) (*BitbonImpl, error) {
	contract, err := bindBitbonImpl(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BitbonImpl{BitbonImplCaller: BitbonImplCaller{contract: contract}, BitbonImplTransactor: BitbonImplTransactor{contract: contract}, BitbonImplFilterer: BitbonImplFilterer{contract: contract}}, nil
}

// NewBitbonImplCaller creates a new read-only instance of BitbonImpl, bound to a specific deployed contract.
func NewBitbonImplCaller(address common.Address, caller bind.ContractCaller) (*BitbonImplCaller, error) {
	contract, err := bindBitbonImpl(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BitbonImplCaller{contract: contract}, nil
}

// NewBitbonImplTransactor creates a new write-only instance of BitbonImpl, bound to a specific deployed contract.
func NewBitbonImplTransactor(address common.Address, transactor bind.ContractTransactor) (*BitbonImplTransactor, error) {
	contract, err := bindBitbonImpl(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BitbonImplTransactor{contract: contract}, nil
}

// NewBitbonImplFilterer creates a new log filterer instance of BitbonImpl, bound to a specific deployed contract.
func NewBitbonImplFilterer(address common.Address, filterer bind.ContractFilterer) (*BitbonImplFilterer, error) {
	contract, err := bindBitbonImpl(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BitbonImplFilterer{contract: contract}, nil
}

// bindBitbonImpl binds a generic wrapper to an already deployed contract.
func bindBitbonImpl(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BitbonImplABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BitbonImpl *BitbonImplRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BitbonImpl.Contract.BitbonImplCaller.contract.Call(opts, result, method, params...)
}

// Return Abi for BitbonImplCaller
func (_BitbonImpl *BitbonImplCaller) ABI() abi.ABI {
	return _BitbonImpl.contract.Abi
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BitbonImpl *BitbonImplRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BitbonImpl.Contract.BitbonImplTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BitbonImpl *BitbonImplRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BitbonImpl.Contract.BitbonImplTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BitbonImpl *BitbonImplCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BitbonImpl.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BitbonImpl *BitbonImplTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BitbonImpl.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BitbonImpl *BitbonImplTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BitbonImpl.Contract.contract.Transact(opts, method, params...)
}

// CONTRACTACCESS is a free data retrieval call binding the contract method 0x0ac62b0c.
//
// Solidity: function CONTRACT_ACCESS() view returns(uint256)
func (_BitbonImpl *BitbonImplCaller) CONTRACTACCESS(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonImpl.contract.Call(opts, out, "CONTRACT_ACCESS")
	return *ret0, err
}

// CONTRACTACCESS is a free data retrieval call binding the contract method 0x0ac62b0c.
//
// Solidity: function CONTRACT_ACCESS() view returns(uint256)
func (_BitbonImpl *BitbonImplSession) CONTRACTACCESS() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTACCESS(&_BitbonImpl.CallOpts)
}

// CONTRACTACCESS is a free data retrieval call binding the contract method 0x0ac62b0c.
//
// Solidity: function CONTRACT_ACCESS() view returns(uint256)
func (_BitbonImpl *BitbonImplCallerSession) CONTRACTACCESS() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTACCESS(&_BitbonImpl.CallOpts)
}

// CONTRACTACCESSSTORAGE is a free data retrieval call binding the contract method 0xceac9bd5.
//
// Solidity: function CONTRACT_ACCESS_STORAGE() view returns(uint256)
func (_BitbonImpl *BitbonImplCaller) CONTRACTACCESSSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonImpl.contract.Call(opts, out, "CONTRACT_ACCESS_STORAGE")
	return *ret0, err
}

// CONTRACTACCESSSTORAGE is a free data retrieval call binding the contract method 0xceac9bd5.
//
// Solidity: function CONTRACT_ACCESS_STORAGE() view returns(uint256)
func (_BitbonImpl *BitbonImplSession) CONTRACTACCESSSTORAGE() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTACCESSSTORAGE(&_BitbonImpl.CallOpts)
}

// CONTRACTACCESSSTORAGE is a free data retrieval call binding the contract method 0xceac9bd5.
//
// Solidity: function CONTRACT_ACCESS_STORAGE() view returns(uint256)
func (_BitbonImpl *BitbonImplCallerSession) CONTRACTACCESSSTORAGE() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTACCESSSTORAGE(&_BitbonImpl.CallOpts)
}

// CONTRACTADMINABLE is a free data retrieval call binding the contract method 0xbbe3a78f.
//
// Solidity: function CONTRACT_ADMINABLE() view returns(uint256)
func (_BitbonImpl *BitbonImplCaller) CONTRACTADMINABLE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonImpl.contract.Call(opts, out, "CONTRACT_ADMINABLE")
	return *ret0, err
}

// CONTRACTADMINABLE is a free data retrieval call binding the contract method 0xbbe3a78f.
//
// Solidity: function CONTRACT_ADMINABLE() view returns(uint256)
func (_BitbonImpl *BitbonImplSession) CONTRACTADMINABLE() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTADMINABLE(&_BitbonImpl.CallOpts)
}

// CONTRACTADMINABLE is a free data retrieval call binding the contract method 0xbbe3a78f.
//
// Solidity: function CONTRACT_ADMINABLE() view returns(uint256)
func (_BitbonImpl *BitbonImplCallerSession) CONTRACTADMINABLE() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTADMINABLE(&_BitbonImpl.CallOpts)
}

// CONTRACTADMINSTORAGE is a free data retrieval call binding the contract method 0xc9aebae9.
//
// Solidity: function CONTRACT_ADMIN_STORAGE() view returns(uint256)
func (_BitbonImpl *BitbonImplCaller) CONTRACTADMINSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonImpl.contract.Call(opts, out, "CONTRACT_ADMIN_STORAGE")
	return *ret0, err
}

// CONTRACTADMINSTORAGE is a free data retrieval call binding the contract method 0xc9aebae9.
//
// Solidity: function CONTRACT_ADMIN_STORAGE() view returns(uint256)
func (_BitbonImpl *BitbonImplSession) CONTRACTADMINSTORAGE() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTADMINSTORAGE(&_BitbonImpl.CallOpts)
}

// CONTRACTADMINSTORAGE is a free data retrieval call binding the contract method 0xc9aebae9.
//
// Solidity: function CONTRACT_ADMIN_STORAGE() view returns(uint256)
func (_BitbonImpl *BitbonImplCallerSession) CONTRACTADMINSTORAGE() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTADMINSTORAGE(&_BitbonImpl.CallOpts)
}

// CONTRACTASSETBOX is a free data retrieval call binding the contract method 0x161aa17c.
//
// Solidity: function CONTRACT_ASSETBOX() view returns(uint256)
func (_BitbonImpl *BitbonImplCaller) CONTRACTASSETBOX(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonImpl.contract.Call(opts, out, "CONTRACT_ASSETBOX")
	return *ret0, err
}

// CONTRACTASSETBOX is a free data retrieval call binding the contract method 0x161aa17c.
//
// Solidity: function CONTRACT_ASSETBOX() view returns(uint256)
func (_BitbonImpl *BitbonImplSession) CONTRACTASSETBOX() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTASSETBOX(&_BitbonImpl.CallOpts)
}

// CONTRACTASSETBOX is a free data retrieval call binding the contract method 0x161aa17c.
//
// Solidity: function CONTRACT_ASSETBOX() view returns(uint256)
func (_BitbonImpl *BitbonImplCallerSession) CONTRACTASSETBOX() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTASSETBOX(&_BitbonImpl.CallOpts)
}

// CONTRACTASSETBOXINFO is a free data retrieval call binding the contract method 0xb572fd3d.
//
// Solidity: function CONTRACT_ASSETBOX_INFO() view returns(uint256)
func (_BitbonImpl *BitbonImplCaller) CONTRACTASSETBOXINFO(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonImpl.contract.Call(opts, out, "CONTRACT_ASSETBOX_INFO")
	return *ret0, err
}

// CONTRACTASSETBOXINFO is a free data retrieval call binding the contract method 0xb572fd3d.
//
// Solidity: function CONTRACT_ASSETBOX_INFO() view returns(uint256)
func (_BitbonImpl *BitbonImplSession) CONTRACTASSETBOXINFO() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTASSETBOXINFO(&_BitbonImpl.CallOpts)
}

// CONTRACTASSETBOXINFO is a free data retrieval call binding the contract method 0xb572fd3d.
//
// Solidity: function CONTRACT_ASSETBOX_INFO() view returns(uint256)
func (_BitbonImpl *BitbonImplCallerSession) CONTRACTASSETBOXINFO() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTASSETBOXINFO(&_BitbonImpl.CallOpts)
}

// CONTRACTASSETBOXINFOSTORAGE is a free data retrieval call binding the contract method 0xf32a115d.
//
// Solidity: function CONTRACT_ASSETBOX_INFO_STORAGE() view returns(uint256)
func (_BitbonImpl *BitbonImplCaller) CONTRACTASSETBOXINFOSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonImpl.contract.Call(opts, out, "CONTRACT_ASSETBOX_INFO_STORAGE")
	return *ret0, err
}

// CONTRACTASSETBOXINFOSTORAGE is a free data retrieval call binding the contract method 0xf32a115d.
//
// Solidity: function CONTRACT_ASSETBOX_INFO_STORAGE() view returns(uint256)
func (_BitbonImpl *BitbonImplSession) CONTRACTASSETBOXINFOSTORAGE() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTASSETBOXINFOSTORAGE(&_BitbonImpl.CallOpts)
}

// CONTRACTASSETBOXINFOSTORAGE is a free data retrieval call binding the contract method 0xf32a115d.
//
// Solidity: function CONTRACT_ASSETBOX_INFO_STORAGE() view returns(uint256)
func (_BitbonImpl *BitbonImplCallerSession) CONTRACTASSETBOXINFOSTORAGE() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTASSETBOXINFOSTORAGE(&_BitbonImpl.CallOpts)
}

// CONTRACTASSETBOXSTORAGE is a free data retrieval call binding the contract method 0xfa5bdafc.
//
// Solidity: function CONTRACT_ASSETBOX_STORAGE() view returns(uint256)
func (_BitbonImpl *BitbonImplCaller) CONTRACTASSETBOXSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonImpl.contract.Call(opts, out, "CONTRACT_ASSETBOX_STORAGE")
	return *ret0, err
}

// CONTRACTASSETBOXSTORAGE is a free data retrieval call binding the contract method 0xfa5bdafc.
//
// Solidity: function CONTRACT_ASSETBOX_STORAGE() view returns(uint256)
func (_BitbonImpl *BitbonImplSession) CONTRACTASSETBOXSTORAGE() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTASSETBOXSTORAGE(&_BitbonImpl.CallOpts)
}

// CONTRACTASSETBOXSTORAGE is a free data retrieval call binding the contract method 0xfa5bdafc.
//
// Solidity: function CONTRACT_ASSETBOX_STORAGE() view returns(uint256)
func (_BitbonImpl *BitbonImplCallerSession) CONTRACTASSETBOXSTORAGE() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTASSETBOXSTORAGE(&_BitbonImpl.CallOpts)
}

// CONTRACTBALANCE is a free data retrieval call binding the contract method 0xd59045f5.
//
// Solidity: function CONTRACT_BALANCE() view returns(uint256)
func (_BitbonImpl *BitbonImplCaller) CONTRACTBALANCE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonImpl.contract.Call(opts, out, "CONTRACT_BALANCE")
	return *ret0, err
}

// CONTRACTBALANCE is a free data retrieval call binding the contract method 0xd59045f5.
//
// Solidity: function CONTRACT_BALANCE() view returns(uint256)
func (_BitbonImpl *BitbonImplSession) CONTRACTBALANCE() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTBALANCE(&_BitbonImpl.CallOpts)
}

// CONTRACTBALANCE is a free data retrieval call binding the contract method 0xd59045f5.
//
// Solidity: function CONTRACT_BALANCE() view returns(uint256)
func (_BitbonImpl *BitbonImplCallerSession) CONTRACTBALANCE() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTBALANCE(&_BitbonImpl.CallOpts)
}

// CONTRACTBITBON is a free data retrieval call binding the contract method 0xb6b433ab.
//
// Solidity: function CONTRACT_BITBON() view returns(uint256)
func (_BitbonImpl *BitbonImplCaller) CONTRACTBITBON(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonImpl.contract.Call(opts, out, "CONTRACT_BITBON")
	return *ret0, err
}

// CONTRACTBITBON is a free data retrieval call binding the contract method 0xb6b433ab.
//
// Solidity: function CONTRACT_BITBON() view returns(uint256)
func (_BitbonImpl *BitbonImplSession) CONTRACTBITBON() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTBITBON(&_BitbonImpl.CallOpts)
}

// CONTRACTBITBON is a free data retrieval call binding the contract method 0xb6b433ab.
//
// Solidity: function CONTRACT_BITBON() view returns(uint256)
func (_BitbonImpl *BitbonImplCallerSession) CONTRACTBITBON() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTBITBON(&_BitbonImpl.CallOpts)
}

// CONTRACTBITBONSTORAGE is a free data retrieval call binding the contract method 0xe7f6c30e.
//
// Solidity: function CONTRACT_BITBON_STORAGE() view returns(uint256)
func (_BitbonImpl *BitbonImplCaller) CONTRACTBITBONSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonImpl.contract.Call(opts, out, "CONTRACT_BITBON_STORAGE")
	return *ret0, err
}

// CONTRACTBITBONSTORAGE is a free data retrieval call binding the contract method 0xe7f6c30e.
//
// Solidity: function CONTRACT_BITBON_STORAGE() view returns(uint256)
func (_BitbonImpl *BitbonImplSession) CONTRACTBITBONSTORAGE() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTBITBONSTORAGE(&_BitbonImpl.CallOpts)
}

// CONTRACTBITBONSTORAGE is a free data retrieval call binding the contract method 0xe7f6c30e.
//
// Solidity: function CONTRACT_BITBON_STORAGE() view returns(uint256)
func (_BitbonImpl *BitbonImplCallerSession) CONTRACTBITBONSTORAGE() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTBITBONSTORAGE(&_BitbonImpl.CallOpts)
}

// CONTRACTBITBONSUPPORT is a free data retrieval call binding the contract method 0x20139c74.
//
// Solidity: function CONTRACT_BITBON_SUPPORT() view returns(uint256)
func (_BitbonImpl *BitbonImplCaller) CONTRACTBITBONSUPPORT(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonImpl.contract.Call(opts, out, "CONTRACT_BITBON_SUPPORT")
	return *ret0, err
}

// CONTRACTBITBONSUPPORT is a free data retrieval call binding the contract method 0x20139c74.
//
// Solidity: function CONTRACT_BITBON_SUPPORT() view returns(uint256)
func (_BitbonImpl *BitbonImplSession) CONTRACTBITBONSUPPORT() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTBITBONSUPPORT(&_BitbonImpl.CallOpts)
}

// CONTRACTBITBONSUPPORT is a free data retrieval call binding the contract method 0x20139c74.
//
// Solidity: function CONTRACT_BITBON_SUPPORT() view returns(uint256)
func (_BitbonImpl *BitbonImplCallerSession) CONTRACTBITBONSUPPORT() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTBITBONSUPPORT(&_BitbonImpl.CallOpts)
}

// CONTRACTBONBALANCE is a free data retrieval call binding the contract method 0x40e02b7c.
//
// Solidity: function CONTRACT_BONBALANCE() view returns(uint256)
func (_BitbonImpl *BitbonImplCaller) CONTRACTBONBALANCE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonImpl.contract.Call(opts, out, "CONTRACT_BONBALANCE")
	return *ret0, err
}

// CONTRACTBONBALANCE is a free data retrieval call binding the contract method 0x40e02b7c.
//
// Solidity: function CONTRACT_BONBALANCE() view returns(uint256)
func (_BitbonImpl *BitbonImplSession) CONTRACTBONBALANCE() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTBONBALANCE(&_BitbonImpl.CallOpts)
}

// CONTRACTBONBALANCE is a free data retrieval call binding the contract method 0x40e02b7c.
//
// Solidity: function CONTRACT_BONBALANCE() view returns(uint256)
func (_BitbonImpl *BitbonImplCallerSession) CONTRACTBONBALANCE() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTBONBALANCE(&_BitbonImpl.CallOpts)
}

// CONTRACTBONTRANSFER is a free data retrieval call binding the contract method 0xc18f4a8a.
//
// Solidity: function CONTRACT_BONTRANSFER() view returns(uint256)
func (_BitbonImpl *BitbonImplCaller) CONTRACTBONTRANSFER(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonImpl.contract.Call(opts, out, "CONTRACT_BONTRANSFER")
	return *ret0, err
}

// CONTRACTBONTRANSFER is a free data retrieval call binding the contract method 0xc18f4a8a.
//
// Solidity: function CONTRACT_BONTRANSFER() view returns(uint256)
func (_BitbonImpl *BitbonImplSession) CONTRACTBONTRANSFER() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTBONTRANSFER(&_BitbonImpl.CallOpts)
}

// CONTRACTBONTRANSFER is a free data retrieval call binding the contract method 0xc18f4a8a.
//
// Solidity: function CONTRACT_BONTRANSFER() view returns(uint256)
func (_BitbonImpl *BitbonImplCallerSession) CONTRACTBONTRANSFER() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTBONTRANSFER(&_BitbonImpl.CallOpts)
}

// CONTRACTDEX is a free data retrieval call binding the contract method 0x52c1f844.
//
// Solidity: function CONTRACT_DEX() view returns(uint256)
func (_BitbonImpl *BitbonImplCaller) CONTRACTDEX(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonImpl.contract.Call(opts, out, "CONTRACT_DEX")
	return *ret0, err
}

// CONTRACTDEX is a free data retrieval call binding the contract method 0x52c1f844.
//
// Solidity: function CONTRACT_DEX() view returns(uint256)
func (_BitbonImpl *BitbonImplSession) CONTRACTDEX() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTDEX(&_BitbonImpl.CallOpts)
}

// CONTRACTDEX is a free data retrieval call binding the contract method 0x52c1f844.
//
// Solidity: function CONTRACT_DEX() view returns(uint256)
func (_BitbonImpl *BitbonImplCallerSession) CONTRACTDEX() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTDEX(&_BitbonImpl.CallOpts)
}

// CONTRACTDISTRIBUTIONGATE is a free data retrieval call binding the contract method 0x6e993588.
//
// Solidity: function CONTRACT_DISTRIBUTION_GATE() view returns(uint256)
func (_BitbonImpl *BitbonImplCaller) CONTRACTDISTRIBUTIONGATE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonImpl.contract.Call(opts, out, "CONTRACT_DISTRIBUTION_GATE")
	return *ret0, err
}

// CONTRACTDISTRIBUTIONGATE is a free data retrieval call binding the contract method 0x6e993588.
//
// Solidity: function CONTRACT_DISTRIBUTION_GATE() view returns(uint256)
func (_BitbonImpl *BitbonImplSession) CONTRACTDISTRIBUTIONGATE() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTDISTRIBUTIONGATE(&_BitbonImpl.CallOpts)
}

// CONTRACTDISTRIBUTIONGATE is a free data retrieval call binding the contract method 0x6e993588.
//
// Solidity: function CONTRACT_DISTRIBUTION_GATE() view returns(uint256)
func (_BitbonImpl *BitbonImplCallerSession) CONTRACTDISTRIBUTIONGATE() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTDISTRIBUTIONGATE(&_BitbonImpl.CallOpts)
}

// CONTRACTDISTRIBUTIONSTORAGE is a free data retrieval call binding the contract method 0x08097327.
//
// Solidity: function CONTRACT_DISTRIBUTION_STORAGE() view returns(uint256)
func (_BitbonImpl *BitbonImplCaller) CONTRACTDISTRIBUTIONSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonImpl.contract.Call(opts, out, "CONTRACT_DISTRIBUTION_STORAGE")
	return *ret0, err
}

// CONTRACTDISTRIBUTIONSTORAGE is a free data retrieval call binding the contract method 0x08097327.
//
// Solidity: function CONTRACT_DISTRIBUTION_STORAGE() view returns(uint256)
func (_BitbonImpl *BitbonImplSession) CONTRACTDISTRIBUTIONSTORAGE() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTDISTRIBUTIONSTORAGE(&_BitbonImpl.CallOpts)
}

// CONTRACTDISTRIBUTIONSTORAGE is a free data retrieval call binding the contract method 0x08097327.
//
// Solidity: function CONTRACT_DISTRIBUTION_STORAGE() view returns(uint256)
func (_BitbonImpl *BitbonImplCallerSession) CONTRACTDISTRIBUTIONSTORAGE() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTDISTRIBUTIONSTORAGE(&_BitbonImpl.CallOpts)
}

// CONTRACTEXCHANGE is a free data retrieval call binding the contract method 0xa0e3a10e.
//
// Solidity: function CONTRACT_EXCHANGE() view returns(uint256)
func (_BitbonImpl *BitbonImplCaller) CONTRACTEXCHANGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonImpl.contract.Call(opts, out, "CONTRACT_EXCHANGE")
	return *ret0, err
}

// CONTRACTEXCHANGE is a free data retrieval call binding the contract method 0xa0e3a10e.
//
// Solidity: function CONTRACT_EXCHANGE() view returns(uint256)
func (_BitbonImpl *BitbonImplSession) CONTRACTEXCHANGE() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTEXCHANGE(&_BitbonImpl.CallOpts)
}

// CONTRACTEXCHANGE is a free data retrieval call binding the contract method 0xa0e3a10e.
//
// Solidity: function CONTRACT_EXCHANGE() view returns(uint256)
func (_BitbonImpl *BitbonImplCallerSession) CONTRACTEXCHANGE() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTEXCHANGE(&_BitbonImpl.CallOpts)
}

// CONTRACTEXCHANGESTORAGE is a free data retrieval call binding the contract method 0x001570f2.
//
// Solidity: function CONTRACT_EXCHANGE_STORAGE() view returns(uint256)
func (_BitbonImpl *BitbonImplCaller) CONTRACTEXCHANGESTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonImpl.contract.Call(opts, out, "CONTRACT_EXCHANGE_STORAGE")
	return *ret0, err
}

// CONTRACTEXCHANGESTORAGE is a free data retrieval call binding the contract method 0x001570f2.
//
// Solidity: function CONTRACT_EXCHANGE_STORAGE() view returns(uint256)
func (_BitbonImpl *BitbonImplSession) CONTRACTEXCHANGESTORAGE() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTEXCHANGESTORAGE(&_BitbonImpl.CallOpts)
}

// CONTRACTEXCHANGESTORAGE is a free data retrieval call binding the contract method 0x001570f2.
//
// Solidity: function CONTRACT_EXCHANGE_STORAGE() view returns(uint256)
func (_BitbonImpl *BitbonImplCallerSession) CONTRACTEXCHANGESTORAGE() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTEXCHANGESTORAGE(&_BitbonImpl.CallOpts)
}

// CONTRACTFEEGATE is a free data retrieval call binding the contract method 0x545c17a4.
//
// Solidity: function CONTRACT_FEE_GATE() view returns(uint256)
func (_BitbonImpl *BitbonImplCaller) CONTRACTFEEGATE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonImpl.contract.Call(opts, out, "CONTRACT_FEE_GATE")
	return *ret0, err
}

// CONTRACTFEEGATE is a free data retrieval call binding the contract method 0x545c17a4.
//
// Solidity: function CONTRACT_FEE_GATE() view returns(uint256)
func (_BitbonImpl *BitbonImplSession) CONTRACTFEEGATE() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTFEEGATE(&_BitbonImpl.CallOpts)
}

// CONTRACTFEEGATE is a free data retrieval call binding the contract method 0x545c17a4.
//
// Solidity: function CONTRACT_FEE_GATE() view returns(uint256)
func (_BitbonImpl *BitbonImplCallerSession) CONTRACTFEEGATE() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTFEEGATE(&_BitbonImpl.CallOpts)
}

// CONTRACTFEESTORAGE is a free data retrieval call binding the contract method 0xcd5e65a4.
//
// Solidity: function CONTRACT_FEE_STORAGE() view returns(uint256)
func (_BitbonImpl *BitbonImplCaller) CONTRACTFEESTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonImpl.contract.Call(opts, out, "CONTRACT_FEE_STORAGE")
	return *ret0, err
}

// CONTRACTFEESTORAGE is a free data retrieval call binding the contract method 0xcd5e65a4.
//
// Solidity: function CONTRACT_FEE_STORAGE() view returns(uint256)
func (_BitbonImpl *BitbonImplSession) CONTRACTFEESTORAGE() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTFEESTORAGE(&_BitbonImpl.CallOpts)
}

// CONTRACTFEESTORAGE is a free data retrieval call binding the contract method 0xcd5e65a4.
//
// Solidity: function CONTRACT_FEE_STORAGE() view returns(uint256)
func (_BitbonImpl *BitbonImplCallerSession) CONTRACTFEESTORAGE() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTFEESTORAGE(&_BitbonImpl.CallOpts)
}

// CONTRACTGENERATOR is a free data retrieval call binding the contract method 0xccea2b46.
//
// Solidity: function CONTRACT_GENERATOR() view returns(uint256)
func (_BitbonImpl *BitbonImplCaller) CONTRACTGENERATOR(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonImpl.contract.Call(opts, out, "CONTRACT_GENERATOR")
	return *ret0, err
}

// CONTRACTGENERATOR is a free data retrieval call binding the contract method 0xccea2b46.
//
// Solidity: function CONTRACT_GENERATOR() view returns(uint256)
func (_BitbonImpl *BitbonImplSession) CONTRACTGENERATOR() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTGENERATOR(&_BitbonImpl.CallOpts)
}

// CONTRACTGENERATOR is a free data retrieval call binding the contract method 0xccea2b46.
//
// Solidity: function CONTRACT_GENERATOR() view returns(uint256)
func (_BitbonImpl *BitbonImplCallerSession) CONTRACTGENERATOR() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTGENERATOR(&_BitbonImpl.CallOpts)
}

// CONTRACTMININGAGENTGATE is a free data retrieval call binding the contract method 0xb9b4c0f1.
//
// Solidity: function CONTRACT_MINING_AGENT_GATE() view returns(uint256)
func (_BitbonImpl *BitbonImplCaller) CONTRACTMININGAGENTGATE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonImpl.contract.Call(opts, out, "CONTRACT_MINING_AGENT_GATE")
	return *ret0, err
}

// CONTRACTMININGAGENTGATE is a free data retrieval call binding the contract method 0xb9b4c0f1.
//
// Solidity: function CONTRACT_MINING_AGENT_GATE() view returns(uint256)
func (_BitbonImpl *BitbonImplSession) CONTRACTMININGAGENTGATE() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTMININGAGENTGATE(&_BitbonImpl.CallOpts)
}

// CONTRACTMININGAGENTGATE is a free data retrieval call binding the contract method 0xb9b4c0f1.
//
// Solidity: function CONTRACT_MINING_AGENT_GATE() view returns(uint256)
func (_BitbonImpl *BitbonImplCallerSession) CONTRACTMININGAGENTGATE() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTMININGAGENTGATE(&_BitbonImpl.CallOpts)
}

// CONTRACTMININGAGENTSTORAGE is a free data retrieval call binding the contract method 0x83cc588d.
//
// Solidity: function CONTRACT_MINING_AGENT_STORAGE() view returns(uint256)
func (_BitbonImpl *BitbonImplCaller) CONTRACTMININGAGENTSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonImpl.contract.Call(opts, out, "CONTRACT_MINING_AGENT_STORAGE")
	return *ret0, err
}

// CONTRACTMININGAGENTSTORAGE is a free data retrieval call binding the contract method 0x83cc588d.
//
// Solidity: function CONTRACT_MINING_AGENT_STORAGE() view returns(uint256)
func (_BitbonImpl *BitbonImplSession) CONTRACTMININGAGENTSTORAGE() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTMININGAGENTSTORAGE(&_BitbonImpl.CallOpts)
}

// CONTRACTMININGAGENTSTORAGE is a free data retrieval call binding the contract method 0x83cc588d.
//
// Solidity: function CONTRACT_MINING_AGENT_STORAGE() view returns(uint256)
func (_BitbonImpl *BitbonImplCallerSession) CONTRACTMININGAGENTSTORAGE() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTMININGAGENTSTORAGE(&_BitbonImpl.CallOpts)
}

// CONTRACTMSBON is a free data retrieval call binding the contract method 0x28b3a976.
//
// Solidity: function CONTRACT_MSBON() view returns(uint256)
func (_BitbonImpl *BitbonImplCaller) CONTRACTMSBON(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonImpl.contract.Call(opts, out, "CONTRACT_MSBON")
	return *ret0, err
}

// CONTRACTMSBON is a free data retrieval call binding the contract method 0x28b3a976.
//
// Solidity: function CONTRACT_MSBON() view returns(uint256)
func (_BitbonImpl *BitbonImplSession) CONTRACTMSBON() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTMSBON(&_BitbonImpl.CallOpts)
}

// CONTRACTMSBON is a free data retrieval call binding the contract method 0x28b3a976.
//
// Solidity: function CONTRACT_MSBON() view returns(uint256)
func (_BitbonImpl *BitbonImplCallerSession) CONTRACTMSBON() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTMSBON(&_BitbonImpl.CallOpts)
}

// CONTRACTMSBONSTORAGE is a free data retrieval call binding the contract method 0xe04e6061.
//
// Solidity: function CONTRACT_MSBON_STORAGE() view returns(uint256)
func (_BitbonImpl *BitbonImplCaller) CONTRACTMSBONSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonImpl.contract.Call(opts, out, "CONTRACT_MSBON_STORAGE")
	return *ret0, err
}

// CONTRACTMSBONSTORAGE is a free data retrieval call binding the contract method 0xe04e6061.
//
// Solidity: function CONTRACT_MSBON_STORAGE() view returns(uint256)
func (_BitbonImpl *BitbonImplSession) CONTRACTMSBONSTORAGE() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTMSBONSTORAGE(&_BitbonImpl.CallOpts)
}

// CONTRACTMSBONSTORAGE is a free data retrieval call binding the contract method 0xe04e6061.
//
// Solidity: function CONTRACT_MSBON_STORAGE() view returns(uint256)
func (_BitbonImpl *BitbonImplCallerSession) CONTRACTMSBONSTORAGE() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTMSBONSTORAGE(&_BitbonImpl.CallOpts)
}

// CONTRACTMULTISIGADDADDRESS is a free data retrieval call binding the contract method 0xf4f65a5a.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ADDRESS() view returns(uint256)
func (_BitbonImpl *BitbonImplCaller) CONTRACTMULTISIGADDADDRESS(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_ADD_ADDRESS")
	return *ret0, err
}

// CONTRACTMULTISIGADDADDRESS is a free data retrieval call binding the contract method 0xf4f65a5a.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ADDRESS() view returns(uint256)
func (_BitbonImpl *BitbonImplSession) CONTRACTMULTISIGADDADDRESS() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTMULTISIGADDADDRESS(&_BitbonImpl.CallOpts)
}

// CONTRACTMULTISIGADDADDRESS is a free data retrieval call binding the contract method 0xf4f65a5a.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ADDRESS() view returns(uint256)
func (_BitbonImpl *BitbonImplCallerSession) CONTRACTMULTISIGADDADDRESS() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTMULTISIGADDADDRESS(&_BitbonImpl.CallOpts)
}

// CONTRACTMULTISIGADDADMIN is a free data retrieval call binding the contract method 0x83888c68.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ADMIN() view returns(uint256)
func (_BitbonImpl *BitbonImplCaller) CONTRACTMULTISIGADDADMIN(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_ADD_ADMIN")
	return *ret0, err
}

// CONTRACTMULTISIGADDADMIN is a free data retrieval call binding the contract method 0x83888c68.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ADMIN() view returns(uint256)
func (_BitbonImpl *BitbonImplSession) CONTRACTMULTISIGADDADMIN() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTMULTISIGADDADMIN(&_BitbonImpl.CallOpts)
}

// CONTRACTMULTISIGADDADMIN is a free data retrieval call binding the contract method 0x83888c68.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ADMIN() view returns(uint256)
func (_BitbonImpl *BitbonImplCallerSession) CONTRACTMULTISIGADDADMIN() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTMULTISIGADDADMIN(&_BitbonImpl.CallOpts)
}

// CONTRACTMULTISIGADDDISTRIBUTION is a free data retrieval call binding the contract method 0x8f8b81d1.
//
// Solidity: function CONTRACT_MULTISIG_ADD_DISTRIBUTION() view returns(uint256)
func (_BitbonImpl *BitbonImplCaller) CONTRACTMULTISIGADDDISTRIBUTION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_ADD_DISTRIBUTION")
	return *ret0, err
}

// CONTRACTMULTISIGADDDISTRIBUTION is a free data retrieval call binding the contract method 0x8f8b81d1.
//
// Solidity: function CONTRACT_MULTISIG_ADD_DISTRIBUTION() view returns(uint256)
func (_BitbonImpl *BitbonImplSession) CONTRACTMULTISIGADDDISTRIBUTION() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTMULTISIGADDDISTRIBUTION(&_BitbonImpl.CallOpts)
}

// CONTRACTMULTISIGADDDISTRIBUTION is a free data retrieval call binding the contract method 0x8f8b81d1.
//
// Solidity: function CONTRACT_MULTISIG_ADD_DISTRIBUTION() view returns(uint256)
func (_BitbonImpl *BitbonImplCallerSession) CONTRACTMULTISIGADDDISTRIBUTION() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTMULTISIGADDDISTRIBUTION(&_BitbonImpl.CallOpts)
}

// CONTRACTMULTISIGADDFEE is a free data retrieval call binding the contract method 0xc43e8cab.
//
// Solidity: function CONTRACT_MULTISIG_ADD_FEE() view returns(uint256)
func (_BitbonImpl *BitbonImplCaller) CONTRACTMULTISIGADDFEE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_ADD_FEE")
	return *ret0, err
}

// CONTRACTMULTISIGADDFEE is a free data retrieval call binding the contract method 0xc43e8cab.
//
// Solidity: function CONTRACT_MULTISIG_ADD_FEE() view returns(uint256)
func (_BitbonImpl *BitbonImplSession) CONTRACTMULTISIGADDFEE() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTMULTISIGADDFEE(&_BitbonImpl.CallOpts)
}

// CONTRACTMULTISIGADDFEE is a free data retrieval call binding the contract method 0xc43e8cab.
//
// Solidity: function CONTRACT_MULTISIG_ADD_FEE() view returns(uint256)
func (_BitbonImpl *BitbonImplCallerSession) CONTRACTMULTISIGADDFEE() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTMULTISIGADDFEE(&_BitbonImpl.CallOpts)
}

// CONTRACTMULTISIGADDROLE is a free data retrieval call binding the contract method 0x20a0914b.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ROLE() view returns(uint256)
func (_BitbonImpl *BitbonImplCaller) CONTRACTMULTISIGADDROLE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_ADD_ROLE")
	return *ret0, err
}

// CONTRACTMULTISIGADDROLE is a free data retrieval call binding the contract method 0x20a0914b.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ROLE() view returns(uint256)
func (_BitbonImpl *BitbonImplSession) CONTRACTMULTISIGADDROLE() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTMULTISIGADDROLE(&_BitbonImpl.CallOpts)
}

// CONTRACTMULTISIGADDROLE is a free data retrieval call binding the contract method 0x20a0914b.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ROLE() view returns(uint256)
func (_BitbonImpl *BitbonImplCallerSession) CONTRACTMULTISIGADDROLE() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTMULTISIGADDROLE(&_BitbonImpl.CallOpts)
}

// CONTRACTMULTISIGCLEARFEEEXCEPTIONS is a free data retrieval call binding the contract method 0xf74701f2.
//
// Solidity: function CONTRACT_MULTISIG_CLEAR_FEE_EXCEPTIONS() view returns(uint256)
func (_BitbonImpl *BitbonImplCaller) CONTRACTMULTISIGCLEARFEEEXCEPTIONS(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_CLEAR_FEE_EXCEPTIONS")
	return *ret0, err
}

// CONTRACTMULTISIGCLEARFEEEXCEPTIONS is a free data retrieval call binding the contract method 0xf74701f2.
//
// Solidity: function CONTRACT_MULTISIG_CLEAR_FEE_EXCEPTIONS() view returns(uint256)
func (_BitbonImpl *BitbonImplSession) CONTRACTMULTISIGCLEARFEEEXCEPTIONS() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTMULTISIGCLEARFEEEXCEPTIONS(&_BitbonImpl.CallOpts)
}

// CONTRACTMULTISIGCLEARFEEEXCEPTIONS is a free data retrieval call binding the contract method 0xf74701f2.
//
// Solidity: function CONTRACT_MULTISIG_CLEAR_FEE_EXCEPTIONS() view returns(uint256)
func (_BitbonImpl *BitbonImplCallerSession) CONTRACTMULTISIGCLEARFEEEXCEPTIONS() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTMULTISIGCLEARFEEEXCEPTIONS(&_BitbonImpl.CallOpts)
}

// CONTRACTMULTISIGDROPPERMISSION is a free data retrieval call binding the contract method 0x3f3fbc39.
//
// Solidity: function CONTRACT_MULTISIG_DROP_PERMISSION() view returns(uint256)
func (_BitbonImpl *BitbonImplCaller) CONTRACTMULTISIGDROPPERMISSION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_DROP_PERMISSION")
	return *ret0, err
}

// CONTRACTMULTISIGDROPPERMISSION is a free data retrieval call binding the contract method 0x3f3fbc39.
//
// Solidity: function CONTRACT_MULTISIG_DROP_PERMISSION() view returns(uint256)
func (_BitbonImpl *BitbonImplSession) CONTRACTMULTISIGDROPPERMISSION() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTMULTISIGDROPPERMISSION(&_BitbonImpl.CallOpts)
}

// CONTRACTMULTISIGDROPPERMISSION is a free data retrieval call binding the contract method 0x3f3fbc39.
//
// Solidity: function CONTRACT_MULTISIG_DROP_PERMISSION() view returns(uint256)
func (_BitbonImpl *BitbonImplCallerSession) CONTRACTMULTISIGDROPPERMISSION() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTMULTISIGDROPPERMISSION(&_BitbonImpl.CallOpts)
}

// CONTRACTMULTISIGEDITADMIN is a free data retrieval call binding the contract method 0x91b0f051.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_ADMIN() view returns(uint256)
func (_BitbonImpl *BitbonImplCaller) CONTRACTMULTISIGEDITADMIN(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_EDIT_ADMIN")
	return *ret0, err
}

// CONTRACTMULTISIGEDITADMIN is a free data retrieval call binding the contract method 0x91b0f051.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_ADMIN() view returns(uint256)
func (_BitbonImpl *BitbonImplSession) CONTRACTMULTISIGEDITADMIN() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTMULTISIGEDITADMIN(&_BitbonImpl.CallOpts)
}

// CONTRACTMULTISIGEDITADMIN is a free data retrieval call binding the contract method 0x91b0f051.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_ADMIN() view returns(uint256)
func (_BitbonImpl *BitbonImplCallerSession) CONTRACTMULTISIGEDITADMIN() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTMULTISIGEDITADMIN(&_BitbonImpl.CallOpts)
}

// CONTRACTMULTISIGEDITFEE is a free data retrieval call binding the contract method 0x6b2abe17.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_FEE() view returns(uint256)
func (_BitbonImpl *BitbonImplCaller) CONTRACTMULTISIGEDITFEE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_EDIT_FEE")
	return *ret0, err
}

// CONTRACTMULTISIGEDITFEE is a free data retrieval call binding the contract method 0x6b2abe17.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_FEE() view returns(uint256)
func (_BitbonImpl *BitbonImplSession) CONTRACTMULTISIGEDITFEE() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTMULTISIGEDITFEE(&_BitbonImpl.CallOpts)
}

// CONTRACTMULTISIGEDITFEE is a free data retrieval call binding the contract method 0x6b2abe17.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_FEE() view returns(uint256)
func (_BitbonImpl *BitbonImplCallerSession) CONTRACTMULTISIGEDITFEE() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTMULTISIGEDITFEE(&_BitbonImpl.CallOpts)
}

// CONTRACTMULTISIGEDITROLE is a free data retrieval call binding the contract method 0xc5f66412.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_ROLE() view returns(uint256)
func (_BitbonImpl *BitbonImplCaller) CONTRACTMULTISIGEDITROLE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_EDIT_ROLE")
	return *ret0, err
}

// CONTRACTMULTISIGEDITROLE is a free data retrieval call binding the contract method 0xc5f66412.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_ROLE() view returns(uint256)
func (_BitbonImpl *BitbonImplSession) CONTRACTMULTISIGEDITROLE() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTMULTISIGEDITROLE(&_BitbonImpl.CallOpts)
}

// CONTRACTMULTISIGEDITROLE is a free data retrieval call binding the contract method 0xc5f66412.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_ROLE() view returns(uint256)
func (_BitbonImpl *BitbonImplCallerSession) CONTRACTMULTISIGEDITROLE() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTMULTISIGEDITROLE(&_BitbonImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEADMINBYADDRESS is a free data retrieval call binding the contract method 0xc816b091.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_BY_ADDRESS() view returns(uint256)
func (_BitbonImpl *BitbonImplCaller) CONTRACTMULTISIGREMOVEADMINBYADDRESS(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REMOVE_ADMIN_BY_ADDRESS")
	return *ret0, err
}

// CONTRACTMULTISIGREMOVEADMINBYADDRESS is a free data retrieval call binding the contract method 0xc816b091.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_BY_ADDRESS() view returns(uint256)
func (_BitbonImpl *BitbonImplSession) CONTRACTMULTISIGREMOVEADMINBYADDRESS() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTMULTISIGREMOVEADMINBYADDRESS(&_BitbonImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEADMINBYADDRESS is a free data retrieval call binding the contract method 0xc816b091.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_BY_ADDRESS() view returns(uint256)
func (_BitbonImpl *BitbonImplCallerSession) CONTRACTMULTISIGREMOVEADMINBYADDRESS() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTMULTISIGREMOVEADMINBYADDRESS(&_BitbonImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEADMINBYINDEX is a free data retrieval call binding the contract method 0x751fbea2.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_BY_INDEX() view returns(uint256)
func (_BitbonImpl *BitbonImplCaller) CONTRACTMULTISIGREMOVEADMINBYINDEX(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REMOVE_ADMIN_BY_INDEX")
	return *ret0, err
}

// CONTRACTMULTISIGREMOVEADMINBYINDEX is a free data retrieval call binding the contract method 0x751fbea2.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_BY_INDEX() view returns(uint256)
func (_BitbonImpl *BitbonImplSession) CONTRACTMULTISIGREMOVEADMINBYINDEX() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTMULTISIGREMOVEADMINBYINDEX(&_BitbonImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEADMINBYINDEX is a free data retrieval call binding the contract method 0x751fbea2.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_BY_INDEX() view returns(uint256)
func (_BitbonImpl *BitbonImplCallerSession) CONTRACTMULTISIGREMOVEADMINBYINDEX() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTMULTISIGREMOVEADMINBYINDEX(&_BitbonImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEADMINROLE is a free data retrieval call binding the contract method 0x53c00200.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_ROLE() view returns(uint256)
func (_BitbonImpl *BitbonImplCaller) CONTRACTMULTISIGREMOVEADMINROLE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REMOVE_ADMIN_ROLE")
	return *ret0, err
}

// CONTRACTMULTISIGREMOVEADMINROLE is a free data retrieval call binding the contract method 0x53c00200.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_ROLE() view returns(uint256)
func (_BitbonImpl *BitbonImplSession) CONTRACTMULTISIGREMOVEADMINROLE() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTMULTISIGREMOVEADMINROLE(&_BitbonImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEADMINROLE is a free data retrieval call binding the contract method 0x53c00200.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_ROLE() view returns(uint256)
func (_BitbonImpl *BitbonImplCallerSession) CONTRACTMULTISIGREMOVEADMINROLE() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTMULTISIGREMOVEADMINROLE(&_BitbonImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEFEE is a free data retrieval call binding the contract method 0xdb614d25.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_FEE() view returns(uint256)
func (_BitbonImpl *BitbonImplCaller) CONTRACTMULTISIGREMOVEFEE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REMOVE_FEE")
	return *ret0, err
}

// CONTRACTMULTISIGREMOVEFEE is a free data retrieval call binding the contract method 0xdb614d25.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_FEE() view returns(uint256)
func (_BitbonImpl *BitbonImplSession) CONTRACTMULTISIGREMOVEFEE() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTMULTISIGREMOVEFEE(&_BitbonImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEFEE is a free data retrieval call binding the contract method 0xdb614d25.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_FEE() view returns(uint256)
func (_BitbonImpl *BitbonImplCallerSession) CONTRACTMULTISIGREMOVEFEE() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTMULTISIGREMOVEFEE(&_BitbonImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEROLE is a free data retrieval call binding the contract method 0x23648496.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ROLE() view returns(uint256)
func (_BitbonImpl *BitbonImplCaller) CONTRACTMULTISIGREMOVEROLE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REMOVE_ROLE")
	return *ret0, err
}

// CONTRACTMULTISIGREMOVEROLE is a free data retrieval call binding the contract method 0x23648496.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ROLE() view returns(uint256)
func (_BitbonImpl *BitbonImplSession) CONTRACTMULTISIGREMOVEROLE() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTMULTISIGREMOVEROLE(&_BitbonImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEROLE is a free data retrieval call binding the contract method 0x23648496.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ROLE() view returns(uint256)
func (_BitbonImpl *BitbonImplCallerSession) CONTRACTMULTISIGREMOVEROLE() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTMULTISIGREMOVEROLE(&_BitbonImpl.CallOpts)
}

// CONTRACTMULTISIGREPLENISHCAPITALIZATION is a free data retrieval call binding the contract method 0x62d5d523.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_CAPITALIZATION() view returns(uint256)
func (_BitbonImpl *BitbonImplCaller) CONTRACTMULTISIGREPLENISHCAPITALIZATION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REPLENISH_CAPITALIZATION")
	return *ret0, err
}

// CONTRACTMULTISIGREPLENISHCAPITALIZATION is a free data retrieval call binding the contract method 0x62d5d523.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_CAPITALIZATION() view returns(uint256)
func (_BitbonImpl *BitbonImplSession) CONTRACTMULTISIGREPLENISHCAPITALIZATION() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTMULTISIGREPLENISHCAPITALIZATION(&_BitbonImpl.CallOpts)
}

// CONTRACTMULTISIGREPLENISHCAPITALIZATION is a free data retrieval call binding the contract method 0x62d5d523.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_CAPITALIZATION() view returns(uint256)
func (_BitbonImpl *BitbonImplCallerSession) CONTRACTMULTISIGREPLENISHCAPITALIZATION() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTMULTISIGREPLENISHCAPITALIZATION(&_BitbonImpl.CallOpts)
}

// CONTRACTMULTISIGREPLENISHCERTIFICATE is a free data retrieval call binding the contract method 0xfd19b7e1.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_CERTIFICATE() view returns(uint256)
func (_BitbonImpl *BitbonImplCaller) CONTRACTMULTISIGREPLENISHCERTIFICATE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REPLENISH_CERTIFICATE")
	return *ret0, err
}

// CONTRACTMULTISIGREPLENISHCERTIFICATE is a free data retrieval call binding the contract method 0xfd19b7e1.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_CERTIFICATE() view returns(uint256)
func (_BitbonImpl *BitbonImplSession) CONTRACTMULTISIGREPLENISHCERTIFICATE() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTMULTISIGREPLENISHCERTIFICATE(&_BitbonImpl.CallOpts)
}

// CONTRACTMULTISIGREPLENISHCERTIFICATE is a free data retrieval call binding the contract method 0xfd19b7e1.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_CERTIFICATE() view returns(uint256)
func (_BitbonImpl *BitbonImplCallerSession) CONTRACTMULTISIGREPLENISHCERTIFICATE() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTMULTISIGREPLENISHCERTIFICATE(&_BitbonImpl.CallOpts)
}

// CONTRACTMULTISIGREPLENISHEMISSION is a free data retrieval call binding the contract method 0x0cb59fde.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_EMISSION() view returns(uint256)
func (_BitbonImpl *BitbonImplCaller) CONTRACTMULTISIGREPLENISHEMISSION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REPLENISH_EMISSION")
	return *ret0, err
}

// CONTRACTMULTISIGREPLENISHEMISSION is a free data retrieval call binding the contract method 0x0cb59fde.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_EMISSION() view returns(uint256)
func (_BitbonImpl *BitbonImplSession) CONTRACTMULTISIGREPLENISHEMISSION() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTMULTISIGREPLENISHEMISSION(&_BitbonImpl.CallOpts)
}

// CONTRACTMULTISIGREPLENISHEMISSION is a free data retrieval call binding the contract method 0x0cb59fde.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_EMISSION() view returns(uint256)
func (_BitbonImpl *BitbonImplCallerSession) CONTRACTMULTISIGREPLENISHEMISSION() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTMULTISIGREPLENISHEMISSION(&_BitbonImpl.CallOpts)
}

// CONTRACTMULTISIGSETADMINROLE is a free data retrieval call binding the contract method 0x8f209ab7.
//
// Solidity: function CONTRACT_MULTISIG_SET_ADMIN_ROLE() view returns(uint256)
func (_BitbonImpl *BitbonImplCaller) CONTRACTMULTISIGSETADMINROLE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_SET_ADMIN_ROLE")
	return *ret0, err
}

// CONTRACTMULTISIGSETADMINROLE is a free data retrieval call binding the contract method 0x8f209ab7.
//
// Solidity: function CONTRACT_MULTISIG_SET_ADMIN_ROLE() view returns(uint256)
func (_BitbonImpl *BitbonImplSession) CONTRACTMULTISIGSETADMINROLE() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTMULTISIGSETADMINROLE(&_BitbonImpl.CallOpts)
}

// CONTRACTMULTISIGSETADMINROLE is a free data retrieval call binding the contract method 0x8f209ab7.
//
// Solidity: function CONTRACT_MULTISIG_SET_ADMIN_ROLE() view returns(uint256)
func (_BitbonImpl *BitbonImplCallerSession) CONTRACTMULTISIGSETADMINROLE() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTMULTISIGSETADMINROLE(&_BitbonImpl.CallOpts)
}

// CONTRACTMULTISIGSETCONFIRMATIONRATE is a free data retrieval call binding the contract method 0xa997279b.
//
// Solidity: function CONTRACT_MULTISIG_SET_CONFIRMATION_RATE() view returns(uint256)
func (_BitbonImpl *BitbonImplCaller) CONTRACTMULTISIGSETCONFIRMATIONRATE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_SET_CONFIRMATION_RATE")
	return *ret0, err
}

// CONTRACTMULTISIGSETCONFIRMATIONRATE is a free data retrieval call binding the contract method 0xa997279b.
//
// Solidity: function CONTRACT_MULTISIG_SET_CONFIRMATION_RATE() view returns(uint256)
func (_BitbonImpl *BitbonImplSession) CONTRACTMULTISIGSETCONFIRMATIONRATE() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTMULTISIGSETCONFIRMATIONRATE(&_BitbonImpl.CallOpts)
}

// CONTRACTMULTISIGSETCONFIRMATIONRATE is a free data retrieval call binding the contract method 0xa997279b.
//
// Solidity: function CONTRACT_MULTISIG_SET_CONFIRMATION_RATE() view returns(uint256)
func (_BitbonImpl *BitbonImplCallerSession) CONTRACTMULTISIGSETCONFIRMATIONRATE() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTMULTISIGSETCONFIRMATIONRATE(&_BitbonImpl.CallOpts)
}

// CONTRACTMULTISIGSETFEEDISTRIBUTION is a free data retrieval call binding the contract method 0x0b76e64d.
//
// Solidity: function CONTRACT_MULTISIG_SET_FEE_DISTRIBUTION() view returns(uint256)
func (_BitbonImpl *BitbonImplCaller) CONTRACTMULTISIGSETFEEDISTRIBUTION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_SET_FEE_DISTRIBUTION")
	return *ret0, err
}

// CONTRACTMULTISIGSETFEEDISTRIBUTION is a free data retrieval call binding the contract method 0x0b76e64d.
//
// Solidity: function CONTRACT_MULTISIG_SET_FEE_DISTRIBUTION() view returns(uint256)
func (_BitbonImpl *BitbonImplSession) CONTRACTMULTISIGSETFEEDISTRIBUTION() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTMULTISIGSETFEEDISTRIBUTION(&_BitbonImpl.CallOpts)
}

// CONTRACTMULTISIGSETFEEDISTRIBUTION is a free data retrieval call binding the contract method 0x0b76e64d.
//
// Solidity: function CONTRACT_MULTISIG_SET_FEE_DISTRIBUTION() view returns(uint256)
func (_BitbonImpl *BitbonImplCallerSession) CONTRACTMULTISIGSETFEEDISTRIBUTION() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTMULTISIGSETFEEDISTRIBUTION(&_BitbonImpl.CallOpts)
}

// CONTRACTMULTISIGSETFEEEXCEPTIONS is a free data retrieval call binding the contract method 0x5e6b28a0.
//
// Solidity: function CONTRACT_MULTISIG_SET_FEE_EXCEPTIONS() view returns(uint256)
func (_BitbonImpl *BitbonImplCaller) CONTRACTMULTISIGSETFEEEXCEPTIONS(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_SET_FEE_EXCEPTIONS")
	return *ret0, err
}

// CONTRACTMULTISIGSETFEEEXCEPTIONS is a free data retrieval call binding the contract method 0x5e6b28a0.
//
// Solidity: function CONTRACT_MULTISIG_SET_FEE_EXCEPTIONS() view returns(uint256)
func (_BitbonImpl *BitbonImplSession) CONTRACTMULTISIGSETFEEEXCEPTIONS() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTMULTISIGSETFEEEXCEPTIONS(&_BitbonImpl.CallOpts)
}

// CONTRACTMULTISIGSETFEEEXCEPTIONS is a free data retrieval call binding the contract method 0x5e6b28a0.
//
// Solidity: function CONTRACT_MULTISIG_SET_FEE_EXCEPTIONS() view returns(uint256)
func (_BitbonImpl *BitbonImplCallerSession) CONTRACTMULTISIGSETFEEEXCEPTIONS() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTMULTISIGSETFEEEXCEPTIONS(&_BitbonImpl.CallOpts)
}

// CONTRACTMULTISIGSETMAXFEE is a free data retrieval call binding the contract method 0xa5f59013.
//
// Solidity: function CONTRACT_MULTISIG_SET_MAX_FEE() view returns(uint256)
func (_BitbonImpl *BitbonImplCaller) CONTRACTMULTISIGSETMAXFEE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_SET_MAX_FEE")
	return *ret0, err
}

// CONTRACTMULTISIGSETMAXFEE is a free data retrieval call binding the contract method 0xa5f59013.
//
// Solidity: function CONTRACT_MULTISIG_SET_MAX_FEE() view returns(uint256)
func (_BitbonImpl *BitbonImplSession) CONTRACTMULTISIGSETMAXFEE() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTMULTISIGSETMAXFEE(&_BitbonImpl.CallOpts)
}

// CONTRACTMULTISIGSETMAXFEE is a free data retrieval call binding the contract method 0xa5f59013.
//
// Solidity: function CONTRACT_MULTISIG_SET_MAX_FEE() view returns(uint256)
func (_BitbonImpl *BitbonImplCallerSession) CONTRACTMULTISIGSETMAXFEE() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTMULTISIGSETMAXFEE(&_BitbonImpl.CallOpts)
}

// CONTRACTMULTISIGSETMINAMOUNT is a free data retrieval call binding the contract method 0xaba147b6.
//
// Solidity: function CONTRACT_MULTISIG_SET_MIN_AMOUNT() view returns(uint256)
func (_BitbonImpl *BitbonImplCaller) CONTRACTMULTISIGSETMINAMOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_SET_MIN_AMOUNT")
	return *ret0, err
}

// CONTRACTMULTISIGSETMINAMOUNT is a free data retrieval call binding the contract method 0xaba147b6.
//
// Solidity: function CONTRACT_MULTISIG_SET_MIN_AMOUNT() view returns(uint256)
func (_BitbonImpl *BitbonImplSession) CONTRACTMULTISIGSETMINAMOUNT() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTMULTISIGSETMINAMOUNT(&_BitbonImpl.CallOpts)
}

// CONTRACTMULTISIGSETMINAMOUNT is a free data retrieval call binding the contract method 0xaba147b6.
//
// Solidity: function CONTRACT_MULTISIG_SET_MIN_AMOUNT() view returns(uint256)
func (_BitbonImpl *BitbonImplCallerSession) CONTRACTMULTISIGSETMINAMOUNT() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTMULTISIGSETMINAMOUNT(&_BitbonImpl.CallOpts)
}

// CONTRACTMULTISIGSETMINFEE is a free data retrieval call binding the contract method 0xf2d1904f.
//
// Solidity: function CONTRACT_MULTISIG_SET_MIN_FEE() view returns(uint256)
func (_BitbonImpl *BitbonImplCaller) CONTRACTMULTISIGSETMINFEE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_SET_MIN_FEE")
	return *ret0, err
}

// CONTRACTMULTISIGSETMINFEE is a free data retrieval call binding the contract method 0xf2d1904f.
//
// Solidity: function CONTRACT_MULTISIG_SET_MIN_FEE() view returns(uint256)
func (_BitbonImpl *BitbonImplSession) CONTRACTMULTISIGSETMINFEE() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTMULTISIGSETMINFEE(&_BitbonImpl.CallOpts)
}

// CONTRACTMULTISIGSETMINFEE is a free data retrieval call binding the contract method 0xf2d1904f.
//
// Solidity: function CONTRACT_MULTISIG_SET_MIN_FEE() view returns(uint256)
func (_BitbonImpl *BitbonImplCallerSession) CONTRACTMULTISIGSETMINFEE() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTMULTISIGSETMINFEE(&_BitbonImpl.CallOpts)
}

// CONTRACTMULTISIGSETPERMISSION is a free data retrieval call binding the contract method 0xb0e91914.
//
// Solidity: function CONTRACT_MULTISIG_SET_PERMISSION() view returns(uint256)
func (_BitbonImpl *BitbonImplCaller) CONTRACTMULTISIGSETPERMISSION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_SET_PERMISSION")
	return *ret0, err
}

// CONTRACTMULTISIGSETPERMISSION is a free data retrieval call binding the contract method 0xb0e91914.
//
// Solidity: function CONTRACT_MULTISIG_SET_PERMISSION() view returns(uint256)
func (_BitbonImpl *BitbonImplSession) CONTRACTMULTISIGSETPERMISSION() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTMULTISIGSETPERMISSION(&_BitbonImpl.CallOpts)
}

// CONTRACTMULTISIGSETPERMISSION is a free data retrieval call binding the contract method 0xb0e91914.
//
// Solidity: function CONTRACT_MULTISIG_SET_PERMISSION() view returns(uint256)
func (_BitbonImpl *BitbonImplCallerSession) CONTRACTMULTISIGSETPERMISSION() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTMULTISIGSETPERMISSION(&_BitbonImpl.CallOpts)
}

// CONTRACTMULTISIGSTORAGE is a free data retrieval call binding the contract method 0xe7e1a303.
//
// Solidity: function CONTRACT_MULTISIG_STORAGE() view returns(uint256)
func (_BitbonImpl *BitbonImplCaller) CONTRACTMULTISIGSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_STORAGE")
	return *ret0, err
}

// CONTRACTMULTISIGSTORAGE is a free data retrieval call binding the contract method 0xe7e1a303.
//
// Solidity: function CONTRACT_MULTISIG_STORAGE() view returns(uint256)
func (_BitbonImpl *BitbonImplSession) CONTRACTMULTISIGSTORAGE() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTMULTISIGSTORAGE(&_BitbonImpl.CallOpts)
}

// CONTRACTMULTISIGSTORAGE is a free data retrieval call binding the contract method 0xe7e1a303.
//
// Solidity: function CONTRACT_MULTISIG_STORAGE() view returns(uint256)
func (_BitbonImpl *BitbonImplCallerSession) CONTRACTMULTISIGSTORAGE() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTMULTISIGSTORAGE(&_BitbonImpl.CallOpts)
}

// CONTRACTOTC is a free data retrieval call binding the contract method 0x5eceac70.
//
// Solidity: function CONTRACT_OTC() view returns(uint256)
func (_BitbonImpl *BitbonImplCaller) CONTRACTOTC(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonImpl.contract.Call(opts, out, "CONTRACT_OTC")
	return *ret0, err
}

// CONTRACTOTC is a free data retrieval call binding the contract method 0x5eceac70.
//
// Solidity: function CONTRACT_OTC() view returns(uint256)
func (_BitbonImpl *BitbonImplSession) CONTRACTOTC() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTOTC(&_BitbonImpl.CallOpts)
}

// CONTRACTOTC is a free data retrieval call binding the contract method 0x5eceac70.
//
// Solidity: function CONTRACT_OTC() view returns(uint256)
func (_BitbonImpl *BitbonImplCallerSession) CONTRACTOTC() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTOTC(&_BitbonImpl.CallOpts)
}

// CONTRACTOTCSTORAGE is a free data retrieval call binding the contract method 0x825b476f.
//
// Solidity: function CONTRACT_OTC_STORAGE() view returns(uint256)
func (_BitbonImpl *BitbonImplCaller) CONTRACTOTCSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonImpl.contract.Call(opts, out, "CONTRACT_OTC_STORAGE")
	return *ret0, err
}

// CONTRACTOTCSTORAGE is a free data retrieval call binding the contract method 0x825b476f.
//
// Solidity: function CONTRACT_OTC_STORAGE() view returns(uint256)
func (_BitbonImpl *BitbonImplSession) CONTRACTOTCSTORAGE() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTOTCSTORAGE(&_BitbonImpl.CallOpts)
}

// CONTRACTOTCSTORAGE is a free data retrieval call binding the contract method 0x825b476f.
//
// Solidity: function CONTRACT_OTC_STORAGE() view returns(uint256)
func (_BitbonImpl *BitbonImplCallerSession) CONTRACTOTCSTORAGE() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTOTCSTORAGE(&_BitbonImpl.CallOpts)
}

// CONTRACTREPLENISHCAPITALIZATION is a free data retrieval call binding the contract method 0xff5e09d8.
//
// Solidity: function CONTRACT_REPLENISH_CAPITALIZATION() view returns(uint256)
func (_BitbonImpl *BitbonImplCaller) CONTRACTREPLENISHCAPITALIZATION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonImpl.contract.Call(opts, out, "CONTRACT_REPLENISH_CAPITALIZATION")
	return *ret0, err
}

// CONTRACTREPLENISHCAPITALIZATION is a free data retrieval call binding the contract method 0xff5e09d8.
//
// Solidity: function CONTRACT_REPLENISH_CAPITALIZATION() view returns(uint256)
func (_BitbonImpl *BitbonImplSession) CONTRACTREPLENISHCAPITALIZATION() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTREPLENISHCAPITALIZATION(&_BitbonImpl.CallOpts)
}

// CONTRACTREPLENISHCAPITALIZATION is a free data retrieval call binding the contract method 0xff5e09d8.
//
// Solidity: function CONTRACT_REPLENISH_CAPITALIZATION() view returns(uint256)
func (_BitbonImpl *BitbonImplCallerSession) CONTRACTREPLENISHCAPITALIZATION() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTREPLENISHCAPITALIZATION(&_BitbonImpl.CallOpts)
}

// CONTRACTRESERVEDALIASSTORAGE is a free data retrieval call binding the contract method 0x70a94aa7.
//
// Solidity: function CONTRACT_RESERVED_ALIAS_STORAGE() view returns(uint256)
func (_BitbonImpl *BitbonImplCaller) CONTRACTRESERVEDALIASSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonImpl.contract.Call(opts, out, "CONTRACT_RESERVED_ALIAS_STORAGE")
	return *ret0, err
}

// CONTRACTRESERVEDALIASSTORAGE is a free data retrieval call binding the contract method 0x70a94aa7.
//
// Solidity: function CONTRACT_RESERVED_ALIAS_STORAGE() view returns(uint256)
func (_BitbonImpl *BitbonImplSession) CONTRACTRESERVEDALIASSTORAGE() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTRESERVEDALIASSTORAGE(&_BitbonImpl.CallOpts)
}

// CONTRACTRESERVEDALIASSTORAGE is a free data retrieval call binding the contract method 0x70a94aa7.
//
// Solidity: function CONTRACT_RESERVED_ALIAS_STORAGE() view returns(uint256)
func (_BitbonImpl *BitbonImplCallerSession) CONTRACTRESERVEDALIASSTORAGE() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTRESERVEDALIASSTORAGE(&_BitbonImpl.CallOpts)
}

// CONTRACTROLESTORAGE is a free data retrieval call binding the contract method 0xb938235b.
//
// Solidity: function CONTRACT_ROLE_STORAGE() view returns(uint256)
func (_BitbonImpl *BitbonImplCaller) CONTRACTROLESTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonImpl.contract.Call(opts, out, "CONTRACT_ROLE_STORAGE")
	return *ret0, err
}

// CONTRACTROLESTORAGE is a free data retrieval call binding the contract method 0xb938235b.
//
// Solidity: function CONTRACT_ROLE_STORAGE() view returns(uint256)
func (_BitbonImpl *BitbonImplSession) CONTRACTROLESTORAGE() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTROLESTORAGE(&_BitbonImpl.CallOpts)
}

// CONTRACTROLESTORAGE is a free data retrieval call binding the contract method 0xb938235b.
//
// Solidity: function CONTRACT_ROLE_STORAGE() view returns(uint256)
func (_BitbonImpl *BitbonImplCallerSession) CONTRACTROLESTORAGE() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTROLESTORAGE(&_BitbonImpl.CallOpts)
}

// CONTRACTSAFETRANSFERSTORAGE is a free data retrieval call binding the contract method 0xe7e45392.
//
// Solidity: function CONTRACT_SAFE_TRANSFER_STORAGE() view returns(uint256)
func (_BitbonImpl *BitbonImplCaller) CONTRACTSAFETRANSFERSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonImpl.contract.Call(opts, out, "CONTRACT_SAFE_TRANSFER_STORAGE")
	return *ret0, err
}

// CONTRACTSAFETRANSFERSTORAGE is a free data retrieval call binding the contract method 0xe7e45392.
//
// Solidity: function CONTRACT_SAFE_TRANSFER_STORAGE() view returns(uint256)
func (_BitbonImpl *BitbonImplSession) CONTRACTSAFETRANSFERSTORAGE() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTSAFETRANSFERSTORAGE(&_BitbonImpl.CallOpts)
}

// CONTRACTSAFETRANSFERSTORAGE is a free data retrieval call binding the contract method 0xe7e45392.
//
// Solidity: function CONTRACT_SAFE_TRANSFER_STORAGE() view returns(uint256)
func (_BitbonImpl *BitbonImplCallerSession) CONTRACTSAFETRANSFERSTORAGE() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTSAFETRANSFERSTORAGE(&_BitbonImpl.CallOpts)
}

// CONTRACTTRANSFER is a free data retrieval call binding the contract method 0xef94fe38.
//
// Solidity: function CONTRACT_TRANSFER() view returns(uint256)
func (_BitbonImpl *BitbonImplCaller) CONTRACTTRANSFER(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonImpl.contract.Call(opts, out, "CONTRACT_TRANSFER")
	return *ret0, err
}

// CONTRACTTRANSFER is a free data retrieval call binding the contract method 0xef94fe38.
//
// Solidity: function CONTRACT_TRANSFER() view returns(uint256)
func (_BitbonImpl *BitbonImplSession) CONTRACTTRANSFER() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTTRANSFER(&_BitbonImpl.CallOpts)
}

// CONTRACTTRANSFER is a free data retrieval call binding the contract method 0xef94fe38.
//
// Solidity: function CONTRACT_TRANSFER() view returns(uint256)
func (_BitbonImpl *BitbonImplCallerSession) CONTRACTTRANSFER() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTTRANSFER(&_BitbonImpl.CallOpts)
}

// CONTRACTTRANSFERFROMCAPITALIZATION is a free data retrieval call binding the contract method 0xc8f8decb.
//
// Solidity: function CONTRACT_TRANSFER_FROM_CAPITALIZATION() view returns(uint256)
func (_BitbonImpl *BitbonImplCaller) CONTRACTTRANSFERFROMCAPITALIZATION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonImpl.contract.Call(opts, out, "CONTRACT_TRANSFER_FROM_CAPITALIZATION")
	return *ret0, err
}

// CONTRACTTRANSFERFROMCAPITALIZATION is a free data retrieval call binding the contract method 0xc8f8decb.
//
// Solidity: function CONTRACT_TRANSFER_FROM_CAPITALIZATION() view returns(uint256)
func (_BitbonImpl *BitbonImplSession) CONTRACTTRANSFERFROMCAPITALIZATION() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTTRANSFERFROMCAPITALIZATION(&_BitbonImpl.CallOpts)
}

// CONTRACTTRANSFERFROMCAPITALIZATION is a free data retrieval call binding the contract method 0xc8f8decb.
//
// Solidity: function CONTRACT_TRANSFER_FROM_CAPITALIZATION() view returns(uint256)
func (_BitbonImpl *BitbonImplCallerSession) CONTRACTTRANSFERFROMCAPITALIZATION() (*big.Int, error) {
	return _BitbonImpl.Contract.CONTRACTTRANSFERFROMCAPITALIZATION(&_BitbonImpl.CallOpts)
}

// FEEOPERATIONCREATESAFETRANSFER is a free data retrieval call binding the contract method 0xf46d727d.
//
// Solidity: function FEE_OPERATION_CREATE_SAFE_TRANSFER() view returns(uint256)
func (_BitbonImpl *BitbonImplCaller) FEEOPERATIONCREATESAFETRANSFER(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonImpl.contract.Call(opts, out, "FEE_OPERATION_CREATE_SAFE_TRANSFER")
	return *ret0, err
}

// FEEOPERATIONCREATESAFETRANSFER is a free data retrieval call binding the contract method 0xf46d727d.
//
// Solidity: function FEE_OPERATION_CREATE_SAFE_TRANSFER() view returns(uint256)
func (_BitbonImpl *BitbonImplSession) FEEOPERATIONCREATESAFETRANSFER() (*big.Int, error) {
	return _BitbonImpl.Contract.FEEOPERATIONCREATESAFETRANSFER(&_BitbonImpl.CallOpts)
}

// FEEOPERATIONCREATESAFETRANSFER is a free data retrieval call binding the contract method 0xf46d727d.
//
// Solidity: function FEE_OPERATION_CREATE_SAFE_TRANSFER() view returns(uint256)
func (_BitbonImpl *BitbonImplCallerSession) FEEOPERATIONCREATESAFETRANSFER() (*big.Int, error) {
	return _BitbonImpl.Contract.FEEOPERATIONCREATESAFETRANSFER(&_BitbonImpl.CallOpts)
}

// FEEOPERATIONCREATESAFETRANSFERALL is a free data retrieval call binding the contract method 0x9d558bef.
//
// Solidity: function FEE_OPERATION_CREATE_SAFE_TRANSFER_ALL() view returns(uint256)
func (_BitbonImpl *BitbonImplCaller) FEEOPERATIONCREATESAFETRANSFERALL(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonImpl.contract.Call(opts, out, "FEE_OPERATION_CREATE_SAFE_TRANSFER_ALL")
	return *ret0, err
}

// FEEOPERATIONCREATESAFETRANSFERALL is a free data retrieval call binding the contract method 0x9d558bef.
//
// Solidity: function FEE_OPERATION_CREATE_SAFE_TRANSFER_ALL() view returns(uint256)
func (_BitbonImpl *BitbonImplSession) FEEOPERATIONCREATESAFETRANSFERALL() (*big.Int, error) {
	return _BitbonImpl.Contract.FEEOPERATIONCREATESAFETRANSFERALL(&_BitbonImpl.CallOpts)
}

// FEEOPERATIONCREATESAFETRANSFERALL is a free data retrieval call binding the contract method 0x9d558bef.
//
// Solidity: function FEE_OPERATION_CREATE_SAFE_TRANSFER_ALL() view returns(uint256)
func (_BitbonImpl *BitbonImplCallerSession) FEEOPERATIONCREATESAFETRANSFERALL() (*big.Int, error) {
	return _BitbonImpl.Contract.FEEOPERATIONCREATESAFETRANSFERALL(&_BitbonImpl.CallOpts)
}

// FEEOPERATIONFRAMETRANSFER is a free data retrieval call binding the contract method 0xc0c92120.
//
// Solidity: function FEE_OPERATION_FRAME_TRANSFER() view returns(uint256)
func (_BitbonImpl *BitbonImplCaller) FEEOPERATIONFRAMETRANSFER(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonImpl.contract.Call(opts, out, "FEE_OPERATION_FRAME_TRANSFER")
	return *ret0, err
}

// FEEOPERATIONFRAMETRANSFER is a free data retrieval call binding the contract method 0xc0c92120.
//
// Solidity: function FEE_OPERATION_FRAME_TRANSFER() view returns(uint256)
func (_BitbonImpl *BitbonImplSession) FEEOPERATIONFRAMETRANSFER() (*big.Int, error) {
	return _BitbonImpl.Contract.FEEOPERATIONFRAMETRANSFER(&_BitbonImpl.CallOpts)
}

// FEEOPERATIONFRAMETRANSFER is a free data retrieval call binding the contract method 0xc0c92120.
//
// Solidity: function FEE_OPERATION_FRAME_TRANSFER() view returns(uint256)
func (_BitbonImpl *BitbonImplCallerSession) FEEOPERATIONFRAMETRANSFER() (*big.Int, error) {
	return _BitbonImpl.Contract.FEEOPERATIONFRAMETRANSFER(&_BitbonImpl.CallOpts)
}

// FEEOPERATIONQUICKTRANSFER is a free data retrieval call binding the contract method 0xbd64b596.
//
// Solidity: function FEE_OPERATION_QUICK_TRANSFER() view returns(uint256)
func (_BitbonImpl *BitbonImplCaller) FEEOPERATIONQUICKTRANSFER(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonImpl.contract.Call(opts, out, "FEE_OPERATION_QUICK_TRANSFER")
	return *ret0, err
}

// FEEOPERATIONQUICKTRANSFER is a free data retrieval call binding the contract method 0xbd64b596.
//
// Solidity: function FEE_OPERATION_QUICK_TRANSFER() view returns(uint256)
func (_BitbonImpl *BitbonImplSession) FEEOPERATIONQUICKTRANSFER() (*big.Int, error) {
	return _BitbonImpl.Contract.FEEOPERATIONQUICKTRANSFER(&_BitbonImpl.CallOpts)
}

// FEEOPERATIONQUICKTRANSFER is a free data retrieval call binding the contract method 0xbd64b596.
//
// Solidity: function FEE_OPERATION_QUICK_TRANSFER() view returns(uint256)
func (_BitbonImpl *BitbonImplCallerSession) FEEOPERATIONQUICKTRANSFER() (*big.Int, error) {
	return _BitbonImpl.Contract.FEEOPERATIONQUICKTRANSFER(&_BitbonImpl.CallOpts)
}

// FEEOPERATIONQUICKTRANSFERALL is a free data retrieval call binding the contract method 0x69cc9631.
//
// Solidity: function FEE_OPERATION_QUICK_TRANSFER_ALL() view returns(uint256)
func (_BitbonImpl *BitbonImplCaller) FEEOPERATIONQUICKTRANSFERALL(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonImpl.contract.Call(opts, out, "FEE_OPERATION_QUICK_TRANSFER_ALL")
	return *ret0, err
}

// FEEOPERATIONQUICKTRANSFERALL is a free data retrieval call binding the contract method 0x69cc9631.
//
// Solidity: function FEE_OPERATION_QUICK_TRANSFER_ALL() view returns(uint256)
func (_BitbonImpl *BitbonImplSession) FEEOPERATIONQUICKTRANSFERALL() (*big.Int, error) {
	return _BitbonImpl.Contract.FEEOPERATIONQUICKTRANSFERALL(&_BitbonImpl.CallOpts)
}

// FEEOPERATIONQUICKTRANSFERALL is a free data retrieval call binding the contract method 0x69cc9631.
//
// Solidity: function FEE_OPERATION_QUICK_TRANSFER_ALL() view returns(uint256)
func (_BitbonImpl *BitbonImplCallerSession) FEEOPERATIONQUICKTRANSFERALL() (*big.Int, error) {
	return _BitbonImpl.Contract.FEEOPERATIONQUICKTRANSFERALL(&_BitbonImpl.CallOpts)
}

// FEEOPERATIONWPCSAFETRANSFER is a free data retrieval call binding the contract method 0x33594855.
//
// Solidity: function FEE_OPERATION_WPC_SAFE_TRANSFER() view returns(uint256)
func (_BitbonImpl *BitbonImplCaller) FEEOPERATIONWPCSAFETRANSFER(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonImpl.contract.Call(opts, out, "FEE_OPERATION_WPC_SAFE_TRANSFER")
	return *ret0, err
}

// FEEOPERATIONWPCSAFETRANSFER is a free data retrieval call binding the contract method 0x33594855.
//
// Solidity: function FEE_OPERATION_WPC_SAFE_TRANSFER() view returns(uint256)
func (_BitbonImpl *BitbonImplSession) FEEOPERATIONWPCSAFETRANSFER() (*big.Int, error) {
	return _BitbonImpl.Contract.FEEOPERATIONWPCSAFETRANSFER(&_BitbonImpl.CallOpts)
}

// FEEOPERATIONWPCSAFETRANSFER is a free data retrieval call binding the contract method 0x33594855.
//
// Solidity: function FEE_OPERATION_WPC_SAFE_TRANSFER() view returns(uint256)
func (_BitbonImpl *BitbonImplCallerSession) FEEOPERATIONWPCSAFETRANSFER() (*big.Int, error) {
	return _BitbonImpl.Contract.FEEOPERATIONWPCSAFETRANSFER(&_BitbonImpl.CallOpts)
}

// FEEOPERATIONWPCSAFETRANSFERALL is a free data retrieval call binding the contract method 0xf0ef7f20.
//
// Solidity: function FEE_OPERATION_WPC_SAFE_TRANSFER_ALL() view returns(uint256)
func (_BitbonImpl *BitbonImplCaller) FEEOPERATIONWPCSAFETRANSFERALL(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonImpl.contract.Call(opts, out, "FEE_OPERATION_WPC_SAFE_TRANSFER_ALL")
	return *ret0, err
}

// FEEOPERATIONWPCSAFETRANSFERALL is a free data retrieval call binding the contract method 0xf0ef7f20.
//
// Solidity: function FEE_OPERATION_WPC_SAFE_TRANSFER_ALL() view returns(uint256)
func (_BitbonImpl *BitbonImplSession) FEEOPERATIONWPCSAFETRANSFERALL() (*big.Int, error) {
	return _BitbonImpl.Contract.FEEOPERATIONWPCSAFETRANSFERALL(&_BitbonImpl.CallOpts)
}

// FEEOPERATIONWPCSAFETRANSFERALL is a free data retrieval call binding the contract method 0xf0ef7f20.
//
// Solidity: function FEE_OPERATION_WPC_SAFE_TRANSFER_ALL() view returns(uint256)
func (_BitbonImpl *BitbonImplCallerSession) FEEOPERATIONWPCSAFETRANSFERALL() (*big.Int, error) {
	return _BitbonImpl.Contract.FEEOPERATIONWPCSAFETRANSFERALL(&_BitbonImpl.CallOpts)
}

// PERMISSIONACCESSRESTORATION is a free data retrieval call binding the contract method 0x5c3053bf.
//
// Solidity: function PERMISSION_ACCESS_RESTORATION() view returns(uint256)
func (_BitbonImpl *BitbonImplCaller) PERMISSIONACCESSRESTORATION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonImpl.contract.Call(opts, out, "PERMISSION_ACCESS_RESTORATION")
	return *ret0, err
}

// PERMISSIONACCESSRESTORATION is a free data retrieval call binding the contract method 0x5c3053bf.
//
// Solidity: function PERMISSION_ACCESS_RESTORATION() view returns(uint256)
func (_BitbonImpl *BitbonImplSession) PERMISSIONACCESSRESTORATION() (*big.Int, error) {
	return _BitbonImpl.Contract.PERMISSIONACCESSRESTORATION(&_BitbonImpl.CallOpts)
}

// PERMISSIONACCESSRESTORATION is a free data retrieval call binding the contract method 0x5c3053bf.
//
// Solidity: function PERMISSION_ACCESS_RESTORATION() view returns(uint256)
func (_BitbonImpl *BitbonImplCallerSession) PERMISSIONACCESSRESTORATION() (*big.Int, error) {
	return _BitbonImpl.Contract.PERMISSIONACCESSRESTORATION(&_BitbonImpl.CallOpts)
}

// PERMISSIONADMINSTORAGE is a free data retrieval call binding the contract method 0x0c9fc0db.
//
// Solidity: function PERMISSION_ADMIN_STORAGE() view returns(uint256)
func (_BitbonImpl *BitbonImplCaller) PERMISSIONADMINSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonImpl.contract.Call(opts, out, "PERMISSION_ADMIN_STORAGE")
	return *ret0, err
}

// PERMISSIONADMINSTORAGE is a free data retrieval call binding the contract method 0x0c9fc0db.
//
// Solidity: function PERMISSION_ADMIN_STORAGE() view returns(uint256)
func (_BitbonImpl *BitbonImplSession) PERMISSIONADMINSTORAGE() (*big.Int, error) {
	return _BitbonImpl.Contract.PERMISSIONADMINSTORAGE(&_BitbonImpl.CallOpts)
}

// PERMISSIONADMINSTORAGE is a free data retrieval call binding the contract method 0x0c9fc0db.
//
// Solidity: function PERMISSION_ADMIN_STORAGE() view returns(uint256)
func (_BitbonImpl *BitbonImplCallerSession) PERMISSIONADMINSTORAGE() (*big.Int, error) {
	return _BitbonImpl.Contract.PERMISSIONADMINSTORAGE(&_BitbonImpl.CallOpts)
}

// PERMISSIONCONTRACTSSTORAGE is a free data retrieval call binding the contract method 0x1516e83b.
//
// Solidity: function PERMISSION_CONTRACTS_STORAGE() view returns(uint256)
func (_BitbonImpl *BitbonImplCaller) PERMISSIONCONTRACTSSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonImpl.contract.Call(opts, out, "PERMISSION_CONTRACTS_STORAGE")
	return *ret0, err
}

// PERMISSIONCONTRACTSSTORAGE is a free data retrieval call binding the contract method 0x1516e83b.
//
// Solidity: function PERMISSION_CONTRACTS_STORAGE() view returns(uint256)
func (_BitbonImpl *BitbonImplSession) PERMISSIONCONTRACTSSTORAGE() (*big.Int, error) {
	return _BitbonImpl.Contract.PERMISSIONCONTRACTSSTORAGE(&_BitbonImpl.CallOpts)
}

// PERMISSIONCONTRACTSSTORAGE is a free data retrieval call binding the contract method 0x1516e83b.
//
// Solidity: function PERMISSION_CONTRACTS_STORAGE() view returns(uint256)
func (_BitbonImpl *BitbonImplCallerSession) PERMISSIONCONTRACTSSTORAGE() (*big.Int, error) {
	return _BitbonImpl.Contract.PERMISSIONCONTRACTSSTORAGE(&_BitbonImpl.CallOpts)
}

// PERMISSIONCONTRACTDEPLOY is a free data retrieval call binding the contract method 0x257a4a04.
//
// Solidity: function PERMISSION_CONTRACT_DEPLOY() view returns(uint256)
func (_BitbonImpl *BitbonImplCaller) PERMISSIONCONTRACTDEPLOY(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonImpl.contract.Call(opts, out, "PERMISSION_CONTRACT_DEPLOY")
	return *ret0, err
}

// PERMISSIONCONTRACTDEPLOY is a free data retrieval call binding the contract method 0x257a4a04.
//
// Solidity: function PERMISSION_CONTRACT_DEPLOY() view returns(uint256)
func (_BitbonImpl *BitbonImplSession) PERMISSIONCONTRACTDEPLOY() (*big.Int, error) {
	return _BitbonImpl.Contract.PERMISSIONCONTRACTDEPLOY(&_BitbonImpl.CallOpts)
}

// PERMISSIONCONTRACTDEPLOY is a free data retrieval call binding the contract method 0x257a4a04.
//
// Solidity: function PERMISSION_CONTRACT_DEPLOY() view returns(uint256)
func (_BitbonImpl *BitbonImplCallerSession) PERMISSIONCONTRACTDEPLOY() (*big.Int, error) {
	return _BitbonImpl.Contract.PERMISSIONCONTRACTDEPLOY(&_BitbonImpl.CallOpts)
}

// PERMISSIONEMISSION is a free data retrieval call binding the contract method 0x4ccd92d1.
//
// Solidity: function PERMISSION_EMISSION() view returns(uint256)
func (_BitbonImpl *BitbonImplCaller) PERMISSIONEMISSION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonImpl.contract.Call(opts, out, "PERMISSION_EMISSION")
	return *ret0, err
}

// PERMISSIONEMISSION is a free data retrieval call binding the contract method 0x4ccd92d1.
//
// Solidity: function PERMISSION_EMISSION() view returns(uint256)
func (_BitbonImpl *BitbonImplSession) PERMISSIONEMISSION() (*big.Int, error) {
	return _BitbonImpl.Contract.PERMISSIONEMISSION(&_BitbonImpl.CallOpts)
}

// PERMISSIONEMISSION is a free data retrieval call binding the contract method 0x4ccd92d1.
//
// Solidity: function PERMISSION_EMISSION() view returns(uint256)
func (_BitbonImpl *BitbonImplCallerSession) PERMISSIONEMISSION() (*big.Int, error) {
	return _BitbonImpl.Contract.PERMISSIONEMISSION(&_BitbonImpl.CallOpts)
}

// PERMISSIONFEESSTORAGE is a free data retrieval call binding the contract method 0xe837eb87.
//
// Solidity: function PERMISSION_FEES_STORAGE() view returns(uint256)
func (_BitbonImpl *BitbonImplCaller) PERMISSIONFEESSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonImpl.contract.Call(opts, out, "PERMISSION_FEES_STORAGE")
	return *ret0, err
}

// PERMISSIONFEESSTORAGE is a free data retrieval call binding the contract method 0xe837eb87.
//
// Solidity: function PERMISSION_FEES_STORAGE() view returns(uint256)
func (_BitbonImpl *BitbonImplSession) PERMISSIONFEESSTORAGE() (*big.Int, error) {
	return _BitbonImpl.Contract.PERMISSIONFEESSTORAGE(&_BitbonImpl.CallOpts)
}

// PERMISSIONFEESSTORAGE is a free data retrieval call binding the contract method 0xe837eb87.
//
// Solidity: function PERMISSION_FEES_STORAGE() view returns(uint256)
func (_BitbonImpl *BitbonImplCallerSession) PERMISSIONFEESSTORAGE() (*big.Int, error) {
	return _BitbonImpl.Contract.PERMISSIONFEESSTORAGE(&_BitbonImpl.CallOpts)
}

// PERMISSIONMULTISIGSTORAGE is a free data retrieval call binding the contract method 0xf1bb7aad.
//
// Solidity: function PERMISSION_MULTISIG_STORAGE() view returns(uint256)
func (_BitbonImpl *BitbonImplCaller) PERMISSIONMULTISIGSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonImpl.contract.Call(opts, out, "PERMISSION_MULTISIG_STORAGE")
	return *ret0, err
}

// PERMISSIONMULTISIGSTORAGE is a free data retrieval call binding the contract method 0xf1bb7aad.
//
// Solidity: function PERMISSION_MULTISIG_STORAGE() view returns(uint256)
func (_BitbonImpl *BitbonImplSession) PERMISSIONMULTISIGSTORAGE() (*big.Int, error) {
	return _BitbonImpl.Contract.PERMISSIONMULTISIGSTORAGE(&_BitbonImpl.CallOpts)
}

// PERMISSIONMULTISIGSTORAGE is a free data retrieval call binding the contract method 0xf1bb7aad.
//
// Solidity: function PERMISSION_MULTISIG_STORAGE() view returns(uint256)
func (_BitbonImpl *BitbonImplCallerSession) PERMISSIONMULTISIGSTORAGE() (*big.Int, error) {
	return _BitbonImpl.Contract.PERMISSIONMULTISIGSTORAGE(&_BitbonImpl.CallOpts)
}

// PERMISSIONROLESSTORAGE is a free data retrieval call binding the contract method 0x12a16032.
//
// Solidity: function PERMISSION_ROLES_STORAGE() view returns(uint256)
func (_BitbonImpl *BitbonImplCaller) PERMISSIONROLESSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonImpl.contract.Call(opts, out, "PERMISSION_ROLES_STORAGE")
	return *ret0, err
}

// PERMISSIONROLESSTORAGE is a free data retrieval call binding the contract method 0x12a16032.
//
// Solidity: function PERMISSION_ROLES_STORAGE() view returns(uint256)
func (_BitbonImpl *BitbonImplSession) PERMISSIONROLESSTORAGE() (*big.Int, error) {
	return _BitbonImpl.Contract.PERMISSIONROLESSTORAGE(&_BitbonImpl.CallOpts)
}

// PERMISSIONROLESSTORAGE is a free data retrieval call binding the contract method 0x12a16032.
//
// Solidity: function PERMISSION_ROLES_STORAGE() view returns(uint256)
func (_BitbonImpl *BitbonImplCallerSession) PERMISSIONROLESSTORAGE() (*big.Int, error) {
	return _BitbonImpl.Contract.PERMISSIONROLESSTORAGE(&_BitbonImpl.CallOpts)
}

// ROLEACCESSVERIFIER is a free data retrieval call binding the contract method 0xc5405f92.
//
// Solidity: function ROLE_ACCESS_VERIFIER() view returns(uint256)
func (_BitbonImpl *BitbonImplCaller) ROLEACCESSVERIFIER(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonImpl.contract.Call(opts, out, "ROLE_ACCESS_VERIFIER")
	return *ret0, err
}

// ROLEACCESSVERIFIER is a free data retrieval call binding the contract method 0xc5405f92.
//
// Solidity: function ROLE_ACCESS_VERIFIER() view returns(uint256)
func (_BitbonImpl *BitbonImplSession) ROLEACCESSVERIFIER() (*big.Int, error) {
	return _BitbonImpl.Contract.ROLEACCESSVERIFIER(&_BitbonImpl.CallOpts)
}

// ROLEACCESSVERIFIER is a free data retrieval call binding the contract method 0xc5405f92.
//
// Solidity: function ROLE_ACCESS_VERIFIER() view returns(uint256)
func (_BitbonImpl *BitbonImplCallerSession) ROLEACCESSVERIFIER() (*big.Int, error) {
	return _BitbonImpl.Contract.ROLEACCESSVERIFIER(&_BitbonImpl.CallOpts)
}

// ROLEBITBONISSUEVERIFIER is a free data retrieval call binding the contract method 0xf0a012a6.
//
// Solidity: function ROLE_BITBON_ISSUE_VERIFIER() view returns(uint256)
func (_BitbonImpl *BitbonImplCaller) ROLEBITBONISSUEVERIFIER(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonImpl.contract.Call(opts, out, "ROLE_BITBON_ISSUE_VERIFIER")
	return *ret0, err
}

// ROLEBITBONISSUEVERIFIER is a free data retrieval call binding the contract method 0xf0a012a6.
//
// Solidity: function ROLE_BITBON_ISSUE_VERIFIER() view returns(uint256)
func (_BitbonImpl *BitbonImplSession) ROLEBITBONISSUEVERIFIER() (*big.Int, error) {
	return _BitbonImpl.Contract.ROLEBITBONISSUEVERIFIER(&_BitbonImpl.CallOpts)
}

// ROLEBITBONISSUEVERIFIER is a free data retrieval call binding the contract method 0xf0a012a6.
//
// Solidity: function ROLE_BITBON_ISSUE_VERIFIER() view returns(uint256)
func (_BitbonImpl *BitbonImplCallerSession) ROLEBITBONISSUEVERIFIER() (*big.Int, error) {
	return _BitbonImpl.Contract.ROLEBITBONISSUEVERIFIER(&_BitbonImpl.CallOpts)
}

// ROLECOMMISSIONVERIFIER is a free data retrieval call binding the contract method 0x8ab3de8c.
//
// Solidity: function ROLE_COMMISSION_VERIFIER() view returns(uint256)
func (_BitbonImpl *BitbonImplCaller) ROLECOMMISSIONVERIFIER(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonImpl.contract.Call(opts, out, "ROLE_COMMISSION_VERIFIER")
	return *ret0, err
}

// ROLECOMMISSIONVERIFIER is a free data retrieval call binding the contract method 0x8ab3de8c.
//
// Solidity: function ROLE_COMMISSION_VERIFIER() view returns(uint256)
func (_BitbonImpl *BitbonImplSession) ROLECOMMISSIONVERIFIER() (*big.Int, error) {
	return _BitbonImpl.Contract.ROLECOMMISSIONVERIFIER(&_BitbonImpl.CallOpts)
}

// ROLECOMMISSIONVERIFIER is a free data retrieval call binding the contract method 0x8ab3de8c.
//
// Solidity: function ROLE_COMMISSION_VERIFIER() view returns(uint256)
func (_BitbonImpl *BitbonImplCallerSession) ROLECOMMISSIONVERIFIER() (*big.Int, error) {
	return _BitbonImpl.Contract.ROLECOMMISSIONVERIFIER(&_BitbonImpl.CallOpts)
}

// ROLEDEPLOYADMIN is a free data retrieval call binding the contract method 0x667bf6e5.
//
// Solidity: function ROLE_DEPLOY_ADMIN() view returns(uint256)
func (_BitbonImpl *BitbonImplCaller) ROLEDEPLOYADMIN(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonImpl.contract.Call(opts, out, "ROLE_DEPLOY_ADMIN")
	return *ret0, err
}

// ROLEDEPLOYADMIN is a free data retrieval call binding the contract method 0x667bf6e5.
//
// Solidity: function ROLE_DEPLOY_ADMIN() view returns(uint256)
func (_BitbonImpl *BitbonImplSession) ROLEDEPLOYADMIN() (*big.Int, error) {
	return _BitbonImpl.Contract.ROLEDEPLOYADMIN(&_BitbonImpl.CallOpts)
}

// ROLEDEPLOYADMIN is a free data retrieval call binding the contract method 0x667bf6e5.
//
// Solidity: function ROLE_DEPLOY_ADMIN() view returns(uint256)
func (_BitbonImpl *BitbonImplCallerSession) ROLEDEPLOYADMIN() (*big.Int, error) {
	return _BitbonImpl.Contract.ROLEDEPLOYADMIN(&_BitbonImpl.CallOpts)
}

// ROLEPERMISSIONADMIN is a free data retrieval call binding the contract method 0x1b239996.
//
// Solidity: function ROLE_PERMISSION_ADMIN() view returns(uint256)
func (_BitbonImpl *BitbonImplCaller) ROLEPERMISSIONADMIN(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonImpl.contract.Call(opts, out, "ROLE_PERMISSION_ADMIN")
	return *ret0, err
}

// ROLEPERMISSIONADMIN is a free data retrieval call binding the contract method 0x1b239996.
//
// Solidity: function ROLE_PERMISSION_ADMIN() view returns(uint256)
func (_BitbonImpl *BitbonImplSession) ROLEPERMISSIONADMIN() (*big.Int, error) {
	return _BitbonImpl.Contract.ROLEPERMISSIONADMIN(&_BitbonImpl.CallOpts)
}

// ROLEPERMISSIONADMIN is a free data retrieval call binding the contract method 0x1b239996.
//
// Solidity: function ROLE_PERMISSION_ADMIN() view returns(uint256)
func (_BitbonImpl *BitbonImplCallerSession) ROLEPERMISSIONADMIN() (*big.Int, error) {
	return _BitbonImpl.Contract.ROLEPERMISSIONADMIN(&_BitbonImpl.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address assetbox) view returns(uint256)
func (_BitbonImpl *BitbonImplCaller) BalanceOf(opts *bind.CallOpts, assetbox common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonImpl.contract.Call(opts, out, "balanceOf", assetbox)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address assetbox) view returns(uint256)
func (_BitbonImpl *BitbonImplSession) BalanceOf(assetbox common.Address) (*big.Int, error) {
	return _BitbonImpl.Contract.BalanceOf(&_BitbonImpl.CallOpts, assetbox)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address assetbox) view returns(uint256)
func (_BitbonImpl *BitbonImplCallerSession) BalanceOf(assetbox common.Address) (*big.Int, error) {
	return _BitbonImpl.Contract.BalanceOf(&_BitbonImpl.CallOpts, assetbox)
}

// BalanceOfFull is a free data retrieval call binding the contract method 0xfe025902.
//
// Solidity: function balanceOfFull(address assetbox) view returns(uint256, uint256)
func (_BitbonImpl *BitbonImplCaller) BalanceOfFull(opts *bind.CallOpts, assetbox common.Address) (*big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _BitbonImpl.contract.Call(opts, out, "balanceOfFull", assetbox)
	return *ret0, *ret1, err
}

// BalanceOfFull is a free data retrieval call binding the contract method 0xfe025902.
//
// Solidity: function balanceOfFull(address assetbox) view returns(uint256, uint256)
func (_BitbonImpl *BitbonImplSession) BalanceOfFull(assetbox common.Address) (*big.Int, *big.Int, error) {
	return _BitbonImpl.Contract.BalanceOfFull(&_BitbonImpl.CallOpts, assetbox)
}

// BalanceOfFull is a free data retrieval call binding the contract method 0xfe025902.
//
// Solidity: function balanceOfFull(address assetbox) view returns(uint256, uint256)
func (_BitbonImpl *BitbonImplCallerSession) BalanceOfFull(assetbox common.Address) (*big.Int, *big.Int, error) {
	return _BitbonImpl.Contract.BalanceOfFull(&_BitbonImpl.CallOpts, assetbox)
}

// BalanceOfLocked is a free data retrieval call binding the contract method 0xe960bb48.
//
// Solidity: function balanceOfLocked(address assetbox) view returns(uint256)
func (_BitbonImpl *BitbonImplCaller) BalanceOfLocked(opts *bind.CallOpts, assetbox common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonImpl.contract.Call(opts, out, "balanceOfLocked", assetbox)
	return *ret0, err
}

// BalanceOfLocked is a free data retrieval call binding the contract method 0xe960bb48.
//
// Solidity: function balanceOfLocked(address assetbox) view returns(uint256)
func (_BitbonImpl *BitbonImplSession) BalanceOfLocked(assetbox common.Address) (*big.Int, error) {
	return _BitbonImpl.Contract.BalanceOfLocked(&_BitbonImpl.CallOpts, assetbox)
}

// BalanceOfLocked is a free data retrieval call binding the contract method 0xe960bb48.
//
// Solidity: function balanceOfLocked(address assetbox) view returns(uint256)
func (_BitbonImpl *BitbonImplCallerSession) BalanceOfLocked(assetbox common.Address) (*big.Int, error) {
	return _BitbonImpl.Contract.BalanceOfLocked(&_BitbonImpl.CallOpts, assetbox)
}

// GetOperationFee is a free data retrieval call binding the contract method 0xb6e9e4bc.
//
// Solidity: function getOperationFee(uint256 value) view returns(uint256)
func (_BitbonImpl *BitbonImplCaller) GetOperationFee(opts *bind.CallOpts, value *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonImpl.contract.Call(opts, out, "getOperationFee", value)
	return *ret0, err
}

// GetOperationFee is a free data retrieval call binding the contract method 0xb6e9e4bc.
//
// Solidity: function getOperationFee(uint256 value) view returns(uint256)
func (_BitbonImpl *BitbonImplSession) GetOperationFee(value *big.Int) (*big.Int, error) {
	return _BitbonImpl.Contract.GetOperationFee(&_BitbonImpl.CallOpts, value)
}

// GetOperationFee is a free data retrieval call binding the contract method 0xb6e9e4bc.
//
// Solidity: function getOperationFee(uint256 value) view returns(uint256)
func (_BitbonImpl *BitbonImplCallerSession) GetOperationFee(value *big.Int) (*big.Int, error) {
	return _BitbonImpl.Contract.GetOperationFee(&_BitbonImpl.CallOpts, value)
}

// GetScopeBalance is a free data retrieval call binding the contract method 0xee94b911.
//
// Solidity: function getScopeBalance(address[] scope) view returns(uint256 res)
func (_BitbonImpl *BitbonImplCaller) GetScopeBalance(opts *bind.CallOpts, scope []common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonImpl.contract.Call(opts, out, "getScopeBalance", scope)
	return *ret0, err
}

// GetScopeBalance is a free data retrieval call binding the contract method 0xee94b911.
//
// Solidity: function getScopeBalance(address[] scope) view returns(uint256 res)
func (_BitbonImpl *BitbonImplSession) GetScopeBalance(scope []common.Address) (*big.Int, error) {
	return _BitbonImpl.Contract.GetScopeBalance(&_BitbonImpl.CallOpts, scope)
}

// GetScopeBalance is a free data retrieval call binding the contract method 0xee94b911.
//
// Solidity: function getScopeBalance(address[] scope) view returns(uint256 res)
func (_BitbonImpl *BitbonImplCallerSession) GetScopeBalance(scope []common.Address) (*big.Int, error) {
	return _BitbonImpl.Contract.GetScopeBalance(&_BitbonImpl.CallOpts, scope)
}

// GetScopeBalances is a free data retrieval call binding the contract method 0x25ffcef1.
//
// Solidity: function getScopeBalances(address[] scope) view returns(uint256[])
func (_BitbonImpl *BitbonImplCaller) GetScopeBalances(opts *bind.CallOpts, scope []common.Address) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _BitbonImpl.contract.Call(opts, out, "getScopeBalances", scope)
	return *ret0, err
}

// GetScopeBalances is a free data retrieval call binding the contract method 0x25ffcef1.
//
// Solidity: function getScopeBalances(address[] scope) view returns(uint256[])
func (_BitbonImpl *BitbonImplSession) GetScopeBalances(scope []common.Address) ([]*big.Int, error) {
	return _BitbonImpl.Contract.GetScopeBalances(&_BitbonImpl.CallOpts, scope)
}

// GetScopeBalances is a free data retrieval call binding the contract method 0x25ffcef1.
//
// Solidity: function getScopeBalances(address[] scope) view returns(uint256[])
func (_BitbonImpl *BitbonImplCallerSession) GetScopeBalances(scope []common.Address) ([]*big.Int, error) {
	return _BitbonImpl.Contract.GetScopeBalances(&_BitbonImpl.CallOpts, scope)
}

// GetThisContractIndex is a free data retrieval call binding the contract method 0x70c6d367.
//
// Solidity: function getThisContractIndex() pure returns(uint256)
func (_BitbonImpl *BitbonImplCaller) GetThisContractIndex(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonImpl.contract.Call(opts, out, "getThisContractIndex")
	return *ret0, err
}

// GetThisContractIndex is a free data retrieval call binding the contract method 0x70c6d367.
//
// Solidity: function getThisContractIndex() pure returns(uint256)
func (_BitbonImpl *BitbonImplSession) GetThisContractIndex() (*big.Int, error) {
	return _BitbonImpl.Contract.GetThisContractIndex(&_BitbonImpl.CallOpts)
}

// GetThisContractIndex is a free data retrieval call binding the contract method 0x70c6d367.
//
// Solidity: function getThisContractIndex() pure returns(uint256)
func (_BitbonImpl *BitbonImplCallerSession) GetThisContractIndex() (*big.Int, error) {
	return _BitbonImpl.Contract.GetThisContractIndex(&_BitbonImpl.CallOpts)
}

// GetTotalLockedBalance is a free data retrieval call binding the contract method 0xd5c25890.
//
// Solidity: function getTotalLockedBalance() view returns(uint256)
func (_BitbonImpl *BitbonImplCaller) GetTotalLockedBalance(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonImpl.contract.Call(opts, out, "getTotalLockedBalance")
	return *ret0, err
}

// GetTotalLockedBalance is a free data retrieval call binding the contract method 0xd5c25890.
//
// Solidity: function getTotalLockedBalance() view returns(uint256)
func (_BitbonImpl *BitbonImplSession) GetTotalLockedBalance() (*big.Int, error) {
	return _BitbonImpl.Contract.GetTotalLockedBalance(&_BitbonImpl.CallOpts)
}

// GetTotalLockedBalance is a free data retrieval call binding the contract method 0xd5c25890.
//
// Solidity: function getTotalLockedBalance() view returns(uint256)
func (_BitbonImpl *BitbonImplCallerSession) GetTotalLockedBalance() (*big.Int, error) {
	return _BitbonImpl.Contract.GetTotalLockedBalance(&_BitbonImpl.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BitbonImpl *BitbonImplCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BitbonImpl.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BitbonImpl *BitbonImplSession) TotalSupply() (*big.Int, error) {
	return _BitbonImpl.Contract.TotalSupply(&_BitbonImpl.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BitbonImpl *BitbonImplCallerSession) TotalSupply() (*big.Int, error) {
	return _BitbonImpl.Contract.TotalSupply(&_BitbonImpl.CallOpts)
}

// TransferExists is a free data retrieval call binding the contract method 0x168f3e26.
//
// Solidity: function transferExists(bytes transferId) view returns(bool)
func (_BitbonImpl *BitbonImplCaller) TransferExists(opts *bind.CallOpts, transferId []byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BitbonImpl.contract.Call(opts, out, "transferExists", transferId)
	return *ret0, err
}

// TransferExists is a free data retrieval call binding the contract method 0x168f3e26.
//
// Solidity: function transferExists(bytes transferId) view returns(bool)
func (_BitbonImpl *BitbonImplSession) TransferExists(transferId []byte) (bool, error) {
	return _BitbonImpl.Contract.TransferExists(&_BitbonImpl.CallOpts, transferId)
}

// TransferExists is a free data retrieval call binding the contract method 0x168f3e26.
//
// Solidity: function transferExists(bytes transferId) view returns(bool)
func (_BitbonImpl *BitbonImplCallerSession) TransferExists(transferId []byte) (bool, error) {
	return _BitbonImpl.Contract.TransferExists(&_BitbonImpl.CallOpts, transferId)
}

// ApproveSafeTransfer is a paid mutator transaction binding the contract method 0xfd821f3f.
//
// Solidity: function approveSafeTransfer(bytes transferId, bytes protectionCode, bytes extraData) returns(bool)
func (_BitbonImpl *BitbonImplTransactor) ApproveSafeTransfer(opts *bind.TransactOpts, transferId []byte, protectionCode []byte, extraData []byte) (*types.Transaction, error) {
	return _BitbonImpl.contract.Transact(opts, "approveSafeTransfer", transferId, protectionCode, extraData)
}

// ApproveSafeTransfer is a paid mutator transaction binding the contract method 0xfd821f3f.
//
// Solidity: function approveSafeTransfer(bytes transferId, bytes protectionCode, bytes extraData) returns(bool)
func (_BitbonImpl *BitbonImplSession) ApproveSafeTransfer(transferId []byte, protectionCode []byte, extraData []byte) (*types.Transaction, error) {
	return _BitbonImpl.Contract.ApproveSafeTransfer(&_BitbonImpl.TransactOpts, transferId, protectionCode, extraData)
}

// ApproveSafeTransfer is a paid mutator transaction binding the contract method 0xfd821f3f.
//
// Solidity: function approveSafeTransfer(bytes transferId, bytes protectionCode, bytes extraData) returns(bool)
func (_BitbonImpl *BitbonImplTransactorSession) ApproveSafeTransfer(transferId []byte, protectionCode []byte, extraData []byte) (*types.Transaction, error) {
	return _BitbonImpl.Contract.ApproveSafeTransfer(&_BitbonImpl.TransactOpts, transferId, protectionCode, extraData)
}

// ApproveSafeTransferAll is a paid mutator transaction binding the contract method 0x47a17ecb.
//
// Solidity: function approveSafeTransferAll(bytes transferId, bytes protectionCode, bytes extraData) returns(bool)
func (_BitbonImpl *BitbonImplTransactor) ApproveSafeTransferAll(opts *bind.TransactOpts, transferId []byte, protectionCode []byte, extraData []byte) (*types.Transaction, error) {
	return _BitbonImpl.contract.Transact(opts, "approveSafeTransferAll", transferId, protectionCode, extraData)
}

// ApproveSafeTransferAll is a paid mutator transaction binding the contract method 0x47a17ecb.
//
// Solidity: function approveSafeTransferAll(bytes transferId, bytes protectionCode, bytes extraData) returns(bool)
func (_BitbonImpl *BitbonImplSession) ApproveSafeTransferAll(transferId []byte, protectionCode []byte, extraData []byte) (*types.Transaction, error) {
	return _BitbonImpl.Contract.ApproveSafeTransferAll(&_BitbonImpl.TransactOpts, transferId, protectionCode, extraData)
}

// ApproveSafeTransferAll is a paid mutator transaction binding the contract method 0x47a17ecb.
//
// Solidity: function approveSafeTransferAll(bytes transferId, bytes protectionCode, bytes extraData) returns(bool)
func (_BitbonImpl *BitbonImplTransactorSession) ApproveSafeTransferAll(transferId []byte, protectionCode []byte, extraData []byte) (*types.Transaction, error) {
	return _BitbonImpl.Contract.ApproveSafeTransferAll(&_BitbonImpl.TransactOpts, transferId, protectionCode, extraData)
}

// ApproveWPCSafeTransfer is a paid mutator transaction binding the contract method 0xc48bb729.
//
// Solidity: function approveWPCSafeTransfer(bytes transferId, bytes extraData) returns(bool)
func (_BitbonImpl *BitbonImplTransactor) ApproveWPCSafeTransfer(opts *bind.TransactOpts, transferId []byte, extraData []byte) (*types.Transaction, error) {
	return _BitbonImpl.contract.Transact(opts, "approveWPCSafeTransfer", transferId, extraData)
}

// ApproveWPCSafeTransfer is a paid mutator transaction binding the contract method 0xc48bb729.
//
// Solidity: function approveWPCSafeTransfer(bytes transferId, bytes extraData) returns(bool)
func (_BitbonImpl *BitbonImplSession) ApproveWPCSafeTransfer(transferId []byte, extraData []byte) (*types.Transaction, error) {
	return _BitbonImpl.Contract.ApproveWPCSafeTransfer(&_BitbonImpl.TransactOpts, transferId, extraData)
}

// ApproveWPCSafeTransfer is a paid mutator transaction binding the contract method 0xc48bb729.
//
// Solidity: function approveWPCSafeTransfer(bytes transferId, bytes extraData) returns(bool)
func (_BitbonImpl *BitbonImplTransactorSession) ApproveWPCSafeTransfer(transferId []byte, extraData []byte) (*types.Transaction, error) {
	return _BitbonImpl.Contract.ApproveWPCSafeTransfer(&_BitbonImpl.TransactOpts, transferId, extraData)
}

// ApproveWPCSafeTransferAll is a paid mutator transaction binding the contract method 0x8313ea89.
//
// Solidity: function approveWPCSafeTransferAll(bytes transferId, bytes extraData) returns(bool)
func (_BitbonImpl *BitbonImplTransactor) ApproveWPCSafeTransferAll(opts *bind.TransactOpts, transferId []byte, extraData []byte) (*types.Transaction, error) {
	return _BitbonImpl.contract.Transact(opts, "approveWPCSafeTransferAll", transferId, extraData)
}

// ApproveWPCSafeTransferAll is a paid mutator transaction binding the contract method 0x8313ea89.
//
// Solidity: function approveWPCSafeTransferAll(bytes transferId, bytes extraData) returns(bool)
func (_BitbonImpl *BitbonImplSession) ApproveWPCSafeTransferAll(transferId []byte, extraData []byte) (*types.Transaction, error) {
	return _BitbonImpl.Contract.ApproveWPCSafeTransferAll(&_BitbonImpl.TransactOpts, transferId, extraData)
}

// ApproveWPCSafeTransferAll is a paid mutator transaction binding the contract method 0x8313ea89.
//
// Solidity: function approveWPCSafeTransferAll(bytes transferId, bytes extraData) returns(bool)
func (_BitbonImpl *BitbonImplTransactorSession) ApproveWPCSafeTransferAll(transferId []byte, extraData []byte) (*types.Transaction, error) {
	return _BitbonImpl.Contract.ApproveWPCSafeTransferAll(&_BitbonImpl.TransactOpts, transferId, extraData)
}

// CancelSafeTransfer is a paid mutator transaction binding the contract method 0xa8389ace.
//
// Solidity: function cancelSafeTransfer(bytes transferId, bytes extraData) returns(bool)
func (_BitbonImpl *BitbonImplTransactor) CancelSafeTransfer(opts *bind.TransactOpts, transferId []byte, extraData []byte) (*types.Transaction, error) {
	return _BitbonImpl.contract.Transact(opts, "cancelSafeTransfer", transferId, extraData)
}

// CancelSafeTransfer is a paid mutator transaction binding the contract method 0xa8389ace.
//
// Solidity: function cancelSafeTransfer(bytes transferId, bytes extraData) returns(bool)
func (_BitbonImpl *BitbonImplSession) CancelSafeTransfer(transferId []byte, extraData []byte) (*types.Transaction, error) {
	return _BitbonImpl.Contract.CancelSafeTransfer(&_BitbonImpl.TransactOpts, transferId, extraData)
}

// CancelSafeTransfer is a paid mutator transaction binding the contract method 0xa8389ace.
//
// Solidity: function cancelSafeTransfer(bytes transferId, bytes extraData) returns(bool)
func (_BitbonImpl *BitbonImplTransactorSession) CancelSafeTransfer(transferId []byte, extraData []byte) (*types.Transaction, error) {
	return _BitbonImpl.Contract.CancelSafeTransfer(&_BitbonImpl.TransactOpts, transferId, extraData)
}

// CancelSafeTransferAll is a paid mutator transaction binding the contract method 0xe48a66f7.
//
// Solidity: function cancelSafeTransferAll(bytes transferId, bytes extraData) returns(bool)
func (_BitbonImpl *BitbonImplTransactor) CancelSafeTransferAll(opts *bind.TransactOpts, transferId []byte, extraData []byte) (*types.Transaction, error) {
	return _BitbonImpl.contract.Transact(opts, "cancelSafeTransferAll", transferId, extraData)
}

// CancelSafeTransferAll is a paid mutator transaction binding the contract method 0xe48a66f7.
//
// Solidity: function cancelSafeTransferAll(bytes transferId, bytes extraData) returns(bool)
func (_BitbonImpl *BitbonImplSession) CancelSafeTransferAll(transferId []byte, extraData []byte) (*types.Transaction, error) {
	return _BitbonImpl.Contract.CancelSafeTransferAll(&_BitbonImpl.TransactOpts, transferId, extraData)
}

// CancelSafeTransferAll is a paid mutator transaction binding the contract method 0xe48a66f7.
//
// Solidity: function cancelSafeTransferAll(bytes transferId, bytes extraData) returns(bool)
func (_BitbonImpl *BitbonImplTransactorSession) CancelSafeTransferAll(transferId []byte, extraData []byte) (*types.Transaction, error) {
	return _BitbonImpl.Contract.CancelSafeTransferAll(&_BitbonImpl.TransactOpts, transferId, extraData)
}

// CancelWPCSafeTransfer is a paid mutator transaction binding the contract method 0x28b5f788.
//
// Solidity: function cancelWPCSafeTransfer(bytes transferId, bytes extraData) returns(bool)
func (_BitbonImpl *BitbonImplTransactor) CancelWPCSafeTransfer(opts *bind.TransactOpts, transferId []byte, extraData []byte) (*types.Transaction, error) {
	return _BitbonImpl.contract.Transact(opts, "cancelWPCSafeTransfer", transferId, extraData)
}

// CancelWPCSafeTransfer is a paid mutator transaction binding the contract method 0x28b5f788.
//
// Solidity: function cancelWPCSafeTransfer(bytes transferId, bytes extraData) returns(bool)
func (_BitbonImpl *BitbonImplSession) CancelWPCSafeTransfer(transferId []byte, extraData []byte) (*types.Transaction, error) {
	return _BitbonImpl.Contract.CancelWPCSafeTransfer(&_BitbonImpl.TransactOpts, transferId, extraData)
}

// CancelWPCSafeTransfer is a paid mutator transaction binding the contract method 0x28b5f788.
//
// Solidity: function cancelWPCSafeTransfer(bytes transferId, bytes extraData) returns(bool)
func (_BitbonImpl *BitbonImplTransactorSession) CancelWPCSafeTransfer(transferId []byte, extraData []byte) (*types.Transaction, error) {
	return _BitbonImpl.Contract.CancelWPCSafeTransfer(&_BitbonImpl.TransactOpts, transferId, extraData)
}

// CancelWPCSafeTransferAll is a paid mutator transaction binding the contract method 0xc8c2e567.
//
// Solidity: function cancelWPCSafeTransferAll(bytes transferId, bytes extraData) returns(bool)
func (_BitbonImpl *BitbonImplTransactor) CancelWPCSafeTransferAll(opts *bind.TransactOpts, transferId []byte, extraData []byte) (*types.Transaction, error) {
	return _BitbonImpl.contract.Transact(opts, "cancelWPCSafeTransferAll", transferId, extraData)
}

// CancelWPCSafeTransferAll is a paid mutator transaction binding the contract method 0xc8c2e567.
//
// Solidity: function cancelWPCSafeTransferAll(bytes transferId, bytes extraData) returns(bool)
func (_BitbonImpl *BitbonImplSession) CancelWPCSafeTransferAll(transferId []byte, extraData []byte) (*types.Transaction, error) {
	return _BitbonImpl.Contract.CancelWPCSafeTransferAll(&_BitbonImpl.TransactOpts, transferId, extraData)
}

// CancelWPCSafeTransferAll is a paid mutator transaction binding the contract method 0xc8c2e567.
//
// Solidity: function cancelWPCSafeTransferAll(bytes transferId, bytes extraData) returns(bool)
func (_BitbonImpl *BitbonImplTransactorSession) CancelWPCSafeTransferAll(transferId []byte, extraData []byte) (*types.Transaction, error) {
	return _BitbonImpl.Contract.CancelWPCSafeTransferAll(&_BitbonImpl.TransactOpts, transferId, extraData)
}

// CreateSafeTransfer is a paid mutator transaction binding the contract method 0x6c9bcf6f.
//
// Solidity: function createSafeTransfer(address to, uint256 value, bytes proof, bytes vk, bytes transferId, uint256 retries, uint256 expiresAt, bytes extraData) returns(bool)
func (_BitbonImpl *BitbonImplTransactor) CreateSafeTransfer(opts *bind.TransactOpts, to common.Address, value *big.Int, proof []byte, vk []byte, transferId []byte, retries *big.Int, expiresAt *big.Int, extraData []byte) (*types.Transaction, error) {
	return _BitbonImpl.contract.Transact(opts, "createSafeTransfer", to, value, proof, vk, transferId, retries, expiresAt, extraData)
}

// CreateSafeTransfer is a paid mutator transaction binding the contract method 0x6c9bcf6f.
//
// Solidity: function createSafeTransfer(address to, uint256 value, bytes proof, bytes vk, bytes transferId, uint256 retries, uint256 expiresAt, bytes extraData) returns(bool)
func (_BitbonImpl *BitbonImplSession) CreateSafeTransfer(to common.Address, value *big.Int, proof []byte, vk []byte, transferId []byte, retries *big.Int, expiresAt *big.Int, extraData []byte) (*types.Transaction, error) {
	return _BitbonImpl.Contract.CreateSafeTransfer(&_BitbonImpl.TransactOpts, to, value, proof, vk, transferId, retries, expiresAt, extraData)
}

// CreateSafeTransfer is a paid mutator transaction binding the contract method 0x6c9bcf6f.
//
// Solidity: function createSafeTransfer(address to, uint256 value, bytes proof, bytes vk, bytes transferId, uint256 retries, uint256 expiresAt, bytes extraData) returns(bool)
func (_BitbonImpl *BitbonImplTransactorSession) CreateSafeTransfer(to common.Address, value *big.Int, proof []byte, vk []byte, transferId []byte, retries *big.Int, expiresAt *big.Int, extraData []byte) (*types.Transaction, error) {
	return _BitbonImpl.Contract.CreateSafeTransfer(&_BitbonImpl.TransactOpts, to, value, proof, vk, transferId, retries, expiresAt, extraData)
}

// CreateSafeTransferAll is a paid mutator transaction binding the contract method 0xe33fe72c.
//
// Solidity: function createSafeTransferAll(address to, bytes proof, bytes vk, bytes transferId, uint256 retries, uint256 expiresAt, bytes extraData) returns(bool)
func (_BitbonImpl *BitbonImplTransactor) CreateSafeTransferAll(opts *bind.TransactOpts, to common.Address, proof []byte, vk []byte, transferId []byte, retries *big.Int, expiresAt *big.Int, extraData []byte) (*types.Transaction, error) {
	return _BitbonImpl.contract.Transact(opts, "createSafeTransferAll", to, proof, vk, transferId, retries, expiresAt, extraData)
}

// CreateSafeTransferAll is a paid mutator transaction binding the contract method 0xe33fe72c.
//
// Solidity: function createSafeTransferAll(address to, bytes proof, bytes vk, bytes transferId, uint256 retries, uint256 expiresAt, bytes extraData) returns(bool)
func (_BitbonImpl *BitbonImplSession) CreateSafeTransferAll(to common.Address, proof []byte, vk []byte, transferId []byte, retries *big.Int, expiresAt *big.Int, extraData []byte) (*types.Transaction, error) {
	return _BitbonImpl.Contract.CreateSafeTransferAll(&_BitbonImpl.TransactOpts, to, proof, vk, transferId, retries, expiresAt, extraData)
}

// CreateSafeTransferAll is a paid mutator transaction binding the contract method 0xe33fe72c.
//
// Solidity: function createSafeTransferAll(address to, bytes proof, bytes vk, bytes transferId, uint256 retries, uint256 expiresAt, bytes extraData) returns(bool)
func (_BitbonImpl *BitbonImplTransactorSession) CreateSafeTransferAll(to common.Address, proof []byte, vk []byte, transferId []byte, retries *big.Int, expiresAt *big.Int, extraData []byte) (*types.Transaction, error) {
	return _BitbonImpl.Contract.CreateSafeTransferAll(&_BitbonImpl.TransactOpts, to, proof, vk, transferId, retries, expiresAt, extraData)
}

// CreateWPCSafeTransfer is a paid mutator transaction binding the contract method 0x9ca9dfa7.
//
// Solidity: function createWPCSafeTransfer(address to, uint256 value, bytes transferId, uint256 expiresAt, bytes extraData) returns(bool)
func (_BitbonImpl *BitbonImplTransactor) CreateWPCSafeTransfer(opts *bind.TransactOpts, to common.Address, value *big.Int, transferId []byte, expiresAt *big.Int, extraData []byte) (*types.Transaction, error) {
	return _BitbonImpl.contract.Transact(opts, "createWPCSafeTransfer", to, value, transferId, expiresAt, extraData)
}

// CreateWPCSafeTransfer is a paid mutator transaction binding the contract method 0x9ca9dfa7.
//
// Solidity: function createWPCSafeTransfer(address to, uint256 value, bytes transferId, uint256 expiresAt, bytes extraData) returns(bool)
func (_BitbonImpl *BitbonImplSession) CreateWPCSafeTransfer(to common.Address, value *big.Int, transferId []byte, expiresAt *big.Int, extraData []byte) (*types.Transaction, error) {
	return _BitbonImpl.Contract.CreateWPCSafeTransfer(&_BitbonImpl.TransactOpts, to, value, transferId, expiresAt, extraData)
}

// CreateWPCSafeTransfer is a paid mutator transaction binding the contract method 0x9ca9dfa7.
//
// Solidity: function createWPCSafeTransfer(address to, uint256 value, bytes transferId, uint256 expiresAt, bytes extraData) returns(bool)
func (_BitbonImpl *BitbonImplTransactorSession) CreateWPCSafeTransfer(to common.Address, value *big.Int, transferId []byte, expiresAt *big.Int, extraData []byte) (*types.Transaction, error) {
	return _BitbonImpl.Contract.CreateWPCSafeTransfer(&_BitbonImpl.TransactOpts, to, value, transferId, expiresAt, extraData)
}

// CreateWPCSafeTransferAll is a paid mutator transaction binding the contract method 0x1a9b6d93.
//
// Solidity: function createWPCSafeTransferAll(address to, bytes transferId, uint256 expiresAt, bytes extraData) returns(bool)
func (_BitbonImpl *BitbonImplTransactor) CreateWPCSafeTransferAll(opts *bind.TransactOpts, to common.Address, transferId []byte, expiresAt *big.Int, extraData []byte) (*types.Transaction, error) {
	return _BitbonImpl.contract.Transact(opts, "createWPCSafeTransferAll", to, transferId, expiresAt, extraData)
}

// CreateWPCSafeTransferAll is a paid mutator transaction binding the contract method 0x1a9b6d93.
//
// Solidity: function createWPCSafeTransferAll(address to, bytes transferId, uint256 expiresAt, bytes extraData) returns(bool)
func (_BitbonImpl *BitbonImplSession) CreateWPCSafeTransferAll(to common.Address, transferId []byte, expiresAt *big.Int, extraData []byte) (*types.Transaction, error) {
	return _BitbonImpl.Contract.CreateWPCSafeTransferAll(&_BitbonImpl.TransactOpts, to, transferId, expiresAt, extraData)
}

// CreateWPCSafeTransferAll is a paid mutator transaction binding the contract method 0x1a9b6d93.
//
// Solidity: function createWPCSafeTransferAll(address to, bytes transferId, uint256 expiresAt, bytes extraData) returns(bool)
func (_BitbonImpl *BitbonImplTransactorSession) CreateWPCSafeTransferAll(to common.Address, transferId []byte, expiresAt *big.Int, extraData []byte) (*types.Transaction, error) {
	return _BitbonImpl.Contract.CreateWPCSafeTransferAll(&_BitbonImpl.TransactOpts, to, transferId, expiresAt, extraData)
}

// ExpireSafeTransfers is a paid mutator transaction binding the contract method 0xa5471f24.
//
// Solidity: function expireSafeTransfers(bytes32[] ids) returns()
func (_BitbonImpl *BitbonImplTransactor) ExpireSafeTransfers(opts *bind.TransactOpts, ids [][32]byte) (*types.Transaction, error) {
	return _BitbonImpl.contract.Transact(opts, "expireSafeTransfers", ids)
}

// ExpireSafeTransfers is a paid mutator transaction binding the contract method 0xa5471f24.
//
// Solidity: function expireSafeTransfers(bytes32[] ids) returns()
func (_BitbonImpl *BitbonImplSession) ExpireSafeTransfers(ids [][32]byte) (*types.Transaction, error) {
	return _BitbonImpl.Contract.ExpireSafeTransfers(&_BitbonImpl.TransactOpts, ids)
}

// ExpireSafeTransfers is a paid mutator transaction binding the contract method 0xa5471f24.
//
// Solidity: function expireSafeTransfers(bytes32[] ids) returns()
func (_BitbonImpl *BitbonImplTransactorSession) ExpireSafeTransfers(ids [][32]byte) (*types.Transaction, error) {
	return _BitbonImpl.Contract.ExpireSafeTransfers(&_BitbonImpl.TransactOpts, ids)
}

// FeeTransfer is a paid mutator transaction binding the contract method 0x4fe06442.
//
// Solidity: function feeTransfer(uint256 opType, address from) returns(bool)
func (_BitbonImpl *BitbonImplTransactor) FeeTransfer(opts *bind.TransactOpts, opType *big.Int, from common.Address) (*types.Transaction, error) {
	return _BitbonImpl.contract.Transact(opts, "feeTransfer", opType, from)
}

// FeeTransfer is a paid mutator transaction binding the contract method 0x4fe06442.
//
// Solidity: function feeTransfer(uint256 opType, address from) returns(bool)
func (_BitbonImpl *BitbonImplSession) FeeTransfer(opType *big.Int, from common.Address) (*types.Transaction, error) {
	return _BitbonImpl.Contract.FeeTransfer(&_BitbonImpl.TransactOpts, opType, from)
}

// FeeTransfer is a paid mutator transaction binding the contract method 0x4fe06442.
//
// Solidity: function feeTransfer(uint256 opType, address from) returns(bool)
func (_BitbonImpl *BitbonImplTransactorSession) FeeTransfer(opType *big.Int, from common.Address) (*types.Transaction, error) {
	return _BitbonImpl.Contract.FeeTransfer(&_BitbonImpl.TransactOpts, opType, from)
}

// FrameTransfer is a paid mutator transaction binding the contract method 0x2e7cfa0c.
//
// Solidity: function frameTransfer(address to, uint256 value, bytes extraData) returns(bool)
func (_BitbonImpl *BitbonImplTransactor) FrameTransfer(opts *bind.TransactOpts, to common.Address, value *big.Int, extraData []byte) (*types.Transaction, error) {
	return _BitbonImpl.contract.Transact(opts, "frameTransfer", to, value, extraData)
}

// FrameTransfer is a paid mutator transaction binding the contract method 0x2e7cfa0c.
//
// Solidity: function frameTransfer(address to, uint256 value, bytes extraData) returns(bool)
func (_BitbonImpl *BitbonImplSession) FrameTransfer(to common.Address, value *big.Int, extraData []byte) (*types.Transaction, error) {
	return _BitbonImpl.Contract.FrameTransfer(&_BitbonImpl.TransactOpts, to, value, extraData)
}

// FrameTransfer is a paid mutator transaction binding the contract method 0x2e7cfa0c.
//
// Solidity: function frameTransfer(address to, uint256 value, bytes extraData) returns(bool)
func (_BitbonImpl *BitbonImplTransactorSession) FrameTransfer(to common.Address, value *big.Int, extraData []byte) (*types.Transaction, error) {
	return _BitbonImpl.Contract.FrameTransfer(&_BitbonImpl.TransactOpts, to, value, extraData)
}

// QuickTransfer is a paid mutator transaction binding the contract method 0x536684e7.
//
// Solidity: function quickTransfer(address to, uint256 value, bytes extraData) returns(bool)
func (_BitbonImpl *BitbonImplTransactor) QuickTransfer(opts *bind.TransactOpts, to common.Address, value *big.Int, extraData []byte) (*types.Transaction, error) {
	return _BitbonImpl.contract.Transact(opts, "quickTransfer", to, value, extraData)
}

// QuickTransfer is a paid mutator transaction binding the contract method 0x536684e7.
//
// Solidity: function quickTransfer(address to, uint256 value, bytes extraData) returns(bool)
func (_BitbonImpl *BitbonImplSession) QuickTransfer(to common.Address, value *big.Int, extraData []byte) (*types.Transaction, error) {
	return _BitbonImpl.Contract.QuickTransfer(&_BitbonImpl.TransactOpts, to, value, extraData)
}

// QuickTransfer is a paid mutator transaction binding the contract method 0x536684e7.
//
// Solidity: function quickTransfer(address to, uint256 value, bytes extraData) returns(bool)
func (_BitbonImpl *BitbonImplTransactorSession) QuickTransfer(to common.Address, value *big.Int, extraData []byte) (*types.Transaction, error) {
	return _BitbonImpl.Contract.QuickTransfer(&_BitbonImpl.TransactOpts, to, value, extraData)
}

// QuickTransferAll is a paid mutator transaction binding the contract method 0xff0833f9.
//
// Solidity: function quickTransferAll(address to, bytes extraData) returns(bool)
func (_BitbonImpl *BitbonImplTransactor) QuickTransferAll(opts *bind.TransactOpts, to common.Address, extraData []byte) (*types.Transaction, error) {
	return _BitbonImpl.contract.Transact(opts, "quickTransferAll", to, extraData)
}

// QuickTransferAll is a paid mutator transaction binding the contract method 0xff0833f9.
//
// Solidity: function quickTransferAll(address to, bytes extraData) returns(bool)
func (_BitbonImpl *BitbonImplSession) QuickTransferAll(to common.Address, extraData []byte) (*types.Transaction, error) {
	return _BitbonImpl.Contract.QuickTransferAll(&_BitbonImpl.TransactOpts, to, extraData)
}

// QuickTransferAll is a paid mutator transaction binding the contract method 0xff0833f9.
//
// Solidity: function quickTransferAll(address to, bytes extraData) returns(bool)
func (_BitbonImpl *BitbonImplTransactorSession) QuickTransferAll(to common.Address, extraData []byte) (*types.Transaction, error) {
	return _BitbonImpl.Contract.QuickTransferAll(&_BitbonImpl.TransactOpts, to, extraData)
}
