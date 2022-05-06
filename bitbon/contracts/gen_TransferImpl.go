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

// TransferImplABI is the input ABI used to generate the binding from.
const TransferImplABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractStorageAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"assetbox\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetboxType\",\"type\":\"uint256\"}],\"name\":\"BalanceChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"assetbox\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"balanceAviable\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"balanceLocked\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetboxType\",\"type\":\"uint256\"}],\"name\":\"BalanceLocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"assetbox\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"balanceAviable\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"balanceLocked\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetboxType\",\"type\":\"uint256\"}],\"name\":\"BalanceUnLocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"source\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"dest\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DirectTransferCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"FeeCharged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"source\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"dest\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FrameTransferCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"source\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"dest\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"QuickTransferCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"transferId\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"status\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"currentRetries\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxRetries\",\"type\":\"uint256\"}],\"name\":\"SafeTransferApprovalResult\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"transferId\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"canceller\",\"type\":\"address\"}],\"name\":\"SafeTransferCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"dest\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"transferId\",\"type\":\"bytes\"}],\"name\":\"SafeTransferCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"transferId\",\"type\":\"bytes\"}],\"name\":\"SafeTransferExpired\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"transferId\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"currentRetries\",\"type\":\"uint256\"}],\"name\":\"TransferApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"transferId\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"currentRetries\",\"type\":\"uint256\"}],\"name\":\"TransferWrongProtectionCode\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ACCESS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ACCESS_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ADMINABLE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ADMIN_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ASSETBOX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ASSETBOX_INFO\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ASSETBOX_INFO_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ASSETBOX_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_BALANCE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_BITBON\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_BITBON_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_BITBON_SUPPORT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_BONBALANCE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_BONTRANSFER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_DEX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_DISTRIBUTION_GATE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_DISTRIBUTION_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_EXCHANGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_EXCHANGE_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_FEE_GATE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_FEE_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_GENERATOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MINING_AGENT_GATE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MINING_AGENT_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MSBON\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MSBON_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_ADD_ADDRESS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_ADD_ADMIN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_ADD_DISTRIBUTION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_ADD_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_ADD_ROLE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_CLEAR_FEE_EXCEPTIONS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_DROP_PERMISSION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_EDIT_ADMIN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_EDIT_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_EDIT_ROLE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REMOVE_ADMIN_BY_ADDRESS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REMOVE_ADMIN_BY_INDEX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REMOVE_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REMOVE_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REMOVE_ROLE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REPLENISH_CAPITALIZATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REPLENISH_CERTIFICATE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REPLENISH_EMISSION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_SET_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_SET_CONFIRMATION_RATE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_SET_FEE_DISTRIBUTION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_SET_FEE_EXCEPTIONS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_SET_MAX_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_SET_MIN_AMOUNT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_SET_MIN_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_SET_PERMISSION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_OTC\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_OTC_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_REPLENISH_CAPITALIZATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_RESERVED_ALIAS_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ROLE_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_SAFE_TRANSFER_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_TRANSFER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_TRANSFER_FROM_CAPITALIZATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"FEE_OPERATION_CREATE_SAFE_TRANSFER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"FEE_OPERATION_CREATE_SAFE_TRANSFER_ALL\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"FEE_OPERATION_FRAME_TRANSFER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"FEE_OPERATION_QUICK_TRANSFER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"FEE_OPERATION_QUICK_TRANSFER_ALL\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"FEE_OPERATION_WPC_SAFE_TRANSFER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"FEE_OPERATION_WPC_SAFE_TRANSFER_ALL\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_ACCESS_RESTORATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_ADMIN_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_CONTRACTS_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_CONTRACT_DEPLOY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_EMISSION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_FEES_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_MULTISIG_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_ROLES_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_ACCESS_VERIFIER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_BITBON_ISSUE_VERIFIER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_COMMISSION_VERIFIER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_DEPLOY_ADMIN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_PERMISSION_ADMIN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getThisContractIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"opType\",\"type\":\"uint256\"}],\"name\":\"feeTransfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"quickTransfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"frameTransfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"vk\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"transferId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"retries\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiresAt\",\"type\":\"uint256\"}],\"name\":\"createSafeTransfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"transferId\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"protectionCode\",\"type\":\"bytes\"}],\"name\":\"approveSafeTransfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"transferId\",\"type\":\"bytes\"}],\"name\":\"cancelSafeTransfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"transferIdHash\",\"type\":\"bytes32\"}],\"name\":\"expireSafeTransfersByHash\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"transferId\",\"type\":\"bytes\"}],\"name\":\"transferExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"directTransfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// TransferImpl is an auto generated Go binding around an Ethereum contract.
type TransferImpl struct {
	TransferImplCaller     // Read-only binding to the contract
	TransferImplTransactor // Write-only binding to the contract
	TransferImplFilterer   // Log filterer for contract events
}

// TransferImplCaller is an auto generated read-only Go binding around an Ethereum contract.
type TransferImplCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransferImplTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TransferImplTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransferImplFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TransferImplFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransferImplSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TransferImplSession struct {
	Contract     *TransferImpl     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TransferImplCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TransferImplCallerSession struct {
	Contract *TransferImplCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// TransferImplTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TransferImplTransactorSession struct {
	Contract     *TransferImplTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// TransferImplRaw is an auto generated low-level Go binding around an Ethereum contract.
type TransferImplRaw struct {
	Contract *TransferImpl // Generic contract binding to access the raw methods on
}

// TransferImplCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TransferImplCallerRaw struct {
	Contract *TransferImplCaller // Generic read-only contract binding to access the raw methods on
}

// TransferImplTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TransferImplTransactorRaw struct {
	Contract *TransferImplTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTransferImpl creates a new instance of TransferImpl, bound to a specific deployed contract.
func NewTransferImpl(address common.Address, backend bind.ContractBackend) (*TransferImpl, error) {
	contract, err := bindTransferImpl(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TransferImpl{TransferImplCaller: TransferImplCaller{contract: contract}, TransferImplTransactor: TransferImplTransactor{contract: contract}, TransferImplFilterer: TransferImplFilterer{contract: contract}}, nil
}

// NewTransferImplCaller creates a new read-only instance of TransferImpl, bound to a specific deployed contract.
func NewTransferImplCaller(address common.Address, caller bind.ContractCaller) (*TransferImplCaller, error) {
	contract, err := bindTransferImpl(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TransferImplCaller{contract: contract}, nil
}

// NewTransferImplTransactor creates a new write-only instance of TransferImpl, bound to a specific deployed contract.
func NewTransferImplTransactor(address common.Address, transactor bind.ContractTransactor) (*TransferImplTransactor, error) {
	contract, err := bindTransferImpl(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TransferImplTransactor{contract: contract}, nil
}

// NewTransferImplFilterer creates a new log filterer instance of TransferImpl, bound to a specific deployed contract.
func NewTransferImplFilterer(address common.Address, filterer bind.ContractFilterer) (*TransferImplFilterer, error) {
	contract, err := bindTransferImpl(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TransferImplFilterer{contract: contract}, nil
}

// bindTransferImpl binds a generic wrapper to an already deployed contract.
func bindTransferImpl(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TransferImplABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TransferImpl *TransferImplRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TransferImpl.Contract.TransferImplCaller.contract.Call(opts, result, method, params...)
}

// Return Abi for TransferImplCaller
func (_TransferImpl *TransferImplCaller) ABI() abi.ABI {
	return _TransferImpl.contract.Abi
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TransferImpl *TransferImplRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TransferImpl.Contract.TransferImplTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TransferImpl *TransferImplRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TransferImpl.Contract.TransferImplTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TransferImpl *TransferImplCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TransferImpl.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TransferImpl *TransferImplTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TransferImpl.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TransferImpl *TransferImplTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TransferImpl.Contract.contract.Transact(opts, method, params...)
}

// CONTRACTACCESS is a free data retrieval call binding the contract method 0x0ac62b0c.
//
// Solidity: function CONTRACT_ACCESS() view returns(uint256)
func (_TransferImpl *TransferImplCaller) CONTRACTACCESS(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TransferImpl.contract.Call(opts, out, "CONTRACT_ACCESS")
	return *ret0, err
}

// CONTRACTACCESS is a free data retrieval call binding the contract method 0x0ac62b0c.
//
// Solidity: function CONTRACT_ACCESS() view returns(uint256)
func (_TransferImpl *TransferImplSession) CONTRACTACCESS() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTACCESS(&_TransferImpl.CallOpts)
}

// CONTRACTACCESS is a free data retrieval call binding the contract method 0x0ac62b0c.
//
// Solidity: function CONTRACT_ACCESS() view returns(uint256)
func (_TransferImpl *TransferImplCallerSession) CONTRACTACCESS() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTACCESS(&_TransferImpl.CallOpts)
}

// CONTRACTACCESSSTORAGE is a free data retrieval call binding the contract method 0xceac9bd5.
//
// Solidity: function CONTRACT_ACCESS_STORAGE() view returns(uint256)
func (_TransferImpl *TransferImplCaller) CONTRACTACCESSSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TransferImpl.contract.Call(opts, out, "CONTRACT_ACCESS_STORAGE")
	return *ret0, err
}

// CONTRACTACCESSSTORAGE is a free data retrieval call binding the contract method 0xceac9bd5.
//
// Solidity: function CONTRACT_ACCESS_STORAGE() view returns(uint256)
func (_TransferImpl *TransferImplSession) CONTRACTACCESSSTORAGE() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTACCESSSTORAGE(&_TransferImpl.CallOpts)
}

// CONTRACTACCESSSTORAGE is a free data retrieval call binding the contract method 0xceac9bd5.
//
// Solidity: function CONTRACT_ACCESS_STORAGE() view returns(uint256)
func (_TransferImpl *TransferImplCallerSession) CONTRACTACCESSSTORAGE() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTACCESSSTORAGE(&_TransferImpl.CallOpts)
}

// CONTRACTADMINABLE is a free data retrieval call binding the contract method 0xbbe3a78f.
//
// Solidity: function CONTRACT_ADMINABLE() view returns(uint256)
func (_TransferImpl *TransferImplCaller) CONTRACTADMINABLE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TransferImpl.contract.Call(opts, out, "CONTRACT_ADMINABLE")
	return *ret0, err
}

// CONTRACTADMINABLE is a free data retrieval call binding the contract method 0xbbe3a78f.
//
// Solidity: function CONTRACT_ADMINABLE() view returns(uint256)
func (_TransferImpl *TransferImplSession) CONTRACTADMINABLE() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTADMINABLE(&_TransferImpl.CallOpts)
}

// CONTRACTADMINABLE is a free data retrieval call binding the contract method 0xbbe3a78f.
//
// Solidity: function CONTRACT_ADMINABLE() view returns(uint256)
func (_TransferImpl *TransferImplCallerSession) CONTRACTADMINABLE() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTADMINABLE(&_TransferImpl.CallOpts)
}

// CONTRACTADMINSTORAGE is a free data retrieval call binding the contract method 0xc9aebae9.
//
// Solidity: function CONTRACT_ADMIN_STORAGE() view returns(uint256)
func (_TransferImpl *TransferImplCaller) CONTRACTADMINSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TransferImpl.contract.Call(opts, out, "CONTRACT_ADMIN_STORAGE")
	return *ret0, err
}

// CONTRACTADMINSTORAGE is a free data retrieval call binding the contract method 0xc9aebae9.
//
// Solidity: function CONTRACT_ADMIN_STORAGE() view returns(uint256)
func (_TransferImpl *TransferImplSession) CONTRACTADMINSTORAGE() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTADMINSTORAGE(&_TransferImpl.CallOpts)
}

// CONTRACTADMINSTORAGE is a free data retrieval call binding the contract method 0xc9aebae9.
//
// Solidity: function CONTRACT_ADMIN_STORAGE() view returns(uint256)
func (_TransferImpl *TransferImplCallerSession) CONTRACTADMINSTORAGE() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTADMINSTORAGE(&_TransferImpl.CallOpts)
}

// CONTRACTASSETBOX is a free data retrieval call binding the contract method 0x161aa17c.
//
// Solidity: function CONTRACT_ASSETBOX() view returns(uint256)
func (_TransferImpl *TransferImplCaller) CONTRACTASSETBOX(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TransferImpl.contract.Call(opts, out, "CONTRACT_ASSETBOX")
	return *ret0, err
}

// CONTRACTASSETBOX is a free data retrieval call binding the contract method 0x161aa17c.
//
// Solidity: function CONTRACT_ASSETBOX() view returns(uint256)
func (_TransferImpl *TransferImplSession) CONTRACTASSETBOX() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTASSETBOX(&_TransferImpl.CallOpts)
}

// CONTRACTASSETBOX is a free data retrieval call binding the contract method 0x161aa17c.
//
// Solidity: function CONTRACT_ASSETBOX() view returns(uint256)
func (_TransferImpl *TransferImplCallerSession) CONTRACTASSETBOX() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTASSETBOX(&_TransferImpl.CallOpts)
}

// CONTRACTASSETBOXINFO is a free data retrieval call binding the contract method 0xb572fd3d.
//
// Solidity: function CONTRACT_ASSETBOX_INFO() view returns(uint256)
func (_TransferImpl *TransferImplCaller) CONTRACTASSETBOXINFO(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TransferImpl.contract.Call(opts, out, "CONTRACT_ASSETBOX_INFO")
	return *ret0, err
}

// CONTRACTASSETBOXINFO is a free data retrieval call binding the contract method 0xb572fd3d.
//
// Solidity: function CONTRACT_ASSETBOX_INFO() view returns(uint256)
func (_TransferImpl *TransferImplSession) CONTRACTASSETBOXINFO() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTASSETBOXINFO(&_TransferImpl.CallOpts)
}

// CONTRACTASSETBOXINFO is a free data retrieval call binding the contract method 0xb572fd3d.
//
// Solidity: function CONTRACT_ASSETBOX_INFO() view returns(uint256)
func (_TransferImpl *TransferImplCallerSession) CONTRACTASSETBOXINFO() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTASSETBOXINFO(&_TransferImpl.CallOpts)
}

// CONTRACTASSETBOXINFOSTORAGE is a free data retrieval call binding the contract method 0xf32a115d.
//
// Solidity: function CONTRACT_ASSETBOX_INFO_STORAGE() view returns(uint256)
func (_TransferImpl *TransferImplCaller) CONTRACTASSETBOXINFOSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TransferImpl.contract.Call(opts, out, "CONTRACT_ASSETBOX_INFO_STORAGE")
	return *ret0, err
}

// CONTRACTASSETBOXINFOSTORAGE is a free data retrieval call binding the contract method 0xf32a115d.
//
// Solidity: function CONTRACT_ASSETBOX_INFO_STORAGE() view returns(uint256)
func (_TransferImpl *TransferImplSession) CONTRACTASSETBOXINFOSTORAGE() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTASSETBOXINFOSTORAGE(&_TransferImpl.CallOpts)
}

// CONTRACTASSETBOXINFOSTORAGE is a free data retrieval call binding the contract method 0xf32a115d.
//
// Solidity: function CONTRACT_ASSETBOX_INFO_STORAGE() view returns(uint256)
func (_TransferImpl *TransferImplCallerSession) CONTRACTASSETBOXINFOSTORAGE() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTASSETBOXINFOSTORAGE(&_TransferImpl.CallOpts)
}

// CONTRACTASSETBOXSTORAGE is a free data retrieval call binding the contract method 0xfa5bdafc.
//
// Solidity: function CONTRACT_ASSETBOX_STORAGE() view returns(uint256)
func (_TransferImpl *TransferImplCaller) CONTRACTASSETBOXSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TransferImpl.contract.Call(opts, out, "CONTRACT_ASSETBOX_STORAGE")
	return *ret0, err
}

// CONTRACTASSETBOXSTORAGE is a free data retrieval call binding the contract method 0xfa5bdafc.
//
// Solidity: function CONTRACT_ASSETBOX_STORAGE() view returns(uint256)
func (_TransferImpl *TransferImplSession) CONTRACTASSETBOXSTORAGE() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTASSETBOXSTORAGE(&_TransferImpl.CallOpts)
}

// CONTRACTASSETBOXSTORAGE is a free data retrieval call binding the contract method 0xfa5bdafc.
//
// Solidity: function CONTRACT_ASSETBOX_STORAGE() view returns(uint256)
func (_TransferImpl *TransferImplCallerSession) CONTRACTASSETBOXSTORAGE() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTASSETBOXSTORAGE(&_TransferImpl.CallOpts)
}

// CONTRACTBALANCE is a free data retrieval call binding the contract method 0xd59045f5.
//
// Solidity: function CONTRACT_BALANCE() view returns(uint256)
func (_TransferImpl *TransferImplCaller) CONTRACTBALANCE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TransferImpl.contract.Call(opts, out, "CONTRACT_BALANCE")
	return *ret0, err
}

// CONTRACTBALANCE is a free data retrieval call binding the contract method 0xd59045f5.
//
// Solidity: function CONTRACT_BALANCE() view returns(uint256)
func (_TransferImpl *TransferImplSession) CONTRACTBALANCE() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTBALANCE(&_TransferImpl.CallOpts)
}

// CONTRACTBALANCE is a free data retrieval call binding the contract method 0xd59045f5.
//
// Solidity: function CONTRACT_BALANCE() view returns(uint256)
func (_TransferImpl *TransferImplCallerSession) CONTRACTBALANCE() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTBALANCE(&_TransferImpl.CallOpts)
}

// CONTRACTBITBON is a free data retrieval call binding the contract method 0xb6b433ab.
//
// Solidity: function CONTRACT_BITBON() view returns(uint256)
func (_TransferImpl *TransferImplCaller) CONTRACTBITBON(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TransferImpl.contract.Call(opts, out, "CONTRACT_BITBON")
	return *ret0, err
}

// CONTRACTBITBON is a free data retrieval call binding the contract method 0xb6b433ab.
//
// Solidity: function CONTRACT_BITBON() view returns(uint256)
func (_TransferImpl *TransferImplSession) CONTRACTBITBON() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTBITBON(&_TransferImpl.CallOpts)
}

// CONTRACTBITBON is a free data retrieval call binding the contract method 0xb6b433ab.
//
// Solidity: function CONTRACT_BITBON() view returns(uint256)
func (_TransferImpl *TransferImplCallerSession) CONTRACTBITBON() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTBITBON(&_TransferImpl.CallOpts)
}

// CONTRACTBITBONSTORAGE is a free data retrieval call binding the contract method 0xe7f6c30e.
//
// Solidity: function CONTRACT_BITBON_STORAGE() view returns(uint256)
func (_TransferImpl *TransferImplCaller) CONTRACTBITBONSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TransferImpl.contract.Call(opts, out, "CONTRACT_BITBON_STORAGE")
	return *ret0, err
}

// CONTRACTBITBONSTORAGE is a free data retrieval call binding the contract method 0xe7f6c30e.
//
// Solidity: function CONTRACT_BITBON_STORAGE() view returns(uint256)
func (_TransferImpl *TransferImplSession) CONTRACTBITBONSTORAGE() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTBITBONSTORAGE(&_TransferImpl.CallOpts)
}

// CONTRACTBITBONSTORAGE is a free data retrieval call binding the contract method 0xe7f6c30e.
//
// Solidity: function CONTRACT_BITBON_STORAGE() view returns(uint256)
func (_TransferImpl *TransferImplCallerSession) CONTRACTBITBONSTORAGE() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTBITBONSTORAGE(&_TransferImpl.CallOpts)
}

// CONTRACTBITBONSUPPORT is a free data retrieval call binding the contract method 0x20139c74.
//
// Solidity: function CONTRACT_BITBON_SUPPORT() view returns(uint256)
func (_TransferImpl *TransferImplCaller) CONTRACTBITBONSUPPORT(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TransferImpl.contract.Call(opts, out, "CONTRACT_BITBON_SUPPORT")
	return *ret0, err
}

// CONTRACTBITBONSUPPORT is a free data retrieval call binding the contract method 0x20139c74.
//
// Solidity: function CONTRACT_BITBON_SUPPORT() view returns(uint256)
func (_TransferImpl *TransferImplSession) CONTRACTBITBONSUPPORT() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTBITBONSUPPORT(&_TransferImpl.CallOpts)
}

// CONTRACTBITBONSUPPORT is a free data retrieval call binding the contract method 0x20139c74.
//
// Solidity: function CONTRACT_BITBON_SUPPORT() view returns(uint256)
func (_TransferImpl *TransferImplCallerSession) CONTRACTBITBONSUPPORT() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTBITBONSUPPORT(&_TransferImpl.CallOpts)
}

// CONTRACTBONBALANCE is a free data retrieval call binding the contract method 0x40e02b7c.
//
// Solidity: function CONTRACT_BONBALANCE() view returns(uint256)
func (_TransferImpl *TransferImplCaller) CONTRACTBONBALANCE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TransferImpl.contract.Call(opts, out, "CONTRACT_BONBALANCE")
	return *ret0, err
}

// CONTRACTBONBALANCE is a free data retrieval call binding the contract method 0x40e02b7c.
//
// Solidity: function CONTRACT_BONBALANCE() view returns(uint256)
func (_TransferImpl *TransferImplSession) CONTRACTBONBALANCE() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTBONBALANCE(&_TransferImpl.CallOpts)
}

// CONTRACTBONBALANCE is a free data retrieval call binding the contract method 0x40e02b7c.
//
// Solidity: function CONTRACT_BONBALANCE() view returns(uint256)
func (_TransferImpl *TransferImplCallerSession) CONTRACTBONBALANCE() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTBONBALANCE(&_TransferImpl.CallOpts)
}

// CONTRACTBONTRANSFER is a free data retrieval call binding the contract method 0xc18f4a8a.
//
// Solidity: function CONTRACT_BONTRANSFER() view returns(uint256)
func (_TransferImpl *TransferImplCaller) CONTRACTBONTRANSFER(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TransferImpl.contract.Call(opts, out, "CONTRACT_BONTRANSFER")
	return *ret0, err
}

// CONTRACTBONTRANSFER is a free data retrieval call binding the contract method 0xc18f4a8a.
//
// Solidity: function CONTRACT_BONTRANSFER() view returns(uint256)
func (_TransferImpl *TransferImplSession) CONTRACTBONTRANSFER() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTBONTRANSFER(&_TransferImpl.CallOpts)
}

// CONTRACTBONTRANSFER is a free data retrieval call binding the contract method 0xc18f4a8a.
//
// Solidity: function CONTRACT_BONTRANSFER() view returns(uint256)
func (_TransferImpl *TransferImplCallerSession) CONTRACTBONTRANSFER() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTBONTRANSFER(&_TransferImpl.CallOpts)
}

// CONTRACTDEX is a free data retrieval call binding the contract method 0x52c1f844.
//
// Solidity: function CONTRACT_DEX() view returns(uint256)
func (_TransferImpl *TransferImplCaller) CONTRACTDEX(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TransferImpl.contract.Call(opts, out, "CONTRACT_DEX")
	return *ret0, err
}

// CONTRACTDEX is a free data retrieval call binding the contract method 0x52c1f844.
//
// Solidity: function CONTRACT_DEX() view returns(uint256)
func (_TransferImpl *TransferImplSession) CONTRACTDEX() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTDEX(&_TransferImpl.CallOpts)
}

// CONTRACTDEX is a free data retrieval call binding the contract method 0x52c1f844.
//
// Solidity: function CONTRACT_DEX() view returns(uint256)
func (_TransferImpl *TransferImplCallerSession) CONTRACTDEX() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTDEX(&_TransferImpl.CallOpts)
}

// CONTRACTDISTRIBUTIONGATE is a free data retrieval call binding the contract method 0x6e993588.
//
// Solidity: function CONTRACT_DISTRIBUTION_GATE() view returns(uint256)
func (_TransferImpl *TransferImplCaller) CONTRACTDISTRIBUTIONGATE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TransferImpl.contract.Call(opts, out, "CONTRACT_DISTRIBUTION_GATE")
	return *ret0, err
}

// CONTRACTDISTRIBUTIONGATE is a free data retrieval call binding the contract method 0x6e993588.
//
// Solidity: function CONTRACT_DISTRIBUTION_GATE() view returns(uint256)
func (_TransferImpl *TransferImplSession) CONTRACTDISTRIBUTIONGATE() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTDISTRIBUTIONGATE(&_TransferImpl.CallOpts)
}

// CONTRACTDISTRIBUTIONGATE is a free data retrieval call binding the contract method 0x6e993588.
//
// Solidity: function CONTRACT_DISTRIBUTION_GATE() view returns(uint256)
func (_TransferImpl *TransferImplCallerSession) CONTRACTDISTRIBUTIONGATE() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTDISTRIBUTIONGATE(&_TransferImpl.CallOpts)
}

// CONTRACTDISTRIBUTIONSTORAGE is a free data retrieval call binding the contract method 0x08097327.
//
// Solidity: function CONTRACT_DISTRIBUTION_STORAGE() view returns(uint256)
func (_TransferImpl *TransferImplCaller) CONTRACTDISTRIBUTIONSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TransferImpl.contract.Call(opts, out, "CONTRACT_DISTRIBUTION_STORAGE")
	return *ret0, err
}

// CONTRACTDISTRIBUTIONSTORAGE is a free data retrieval call binding the contract method 0x08097327.
//
// Solidity: function CONTRACT_DISTRIBUTION_STORAGE() view returns(uint256)
func (_TransferImpl *TransferImplSession) CONTRACTDISTRIBUTIONSTORAGE() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTDISTRIBUTIONSTORAGE(&_TransferImpl.CallOpts)
}

// CONTRACTDISTRIBUTIONSTORAGE is a free data retrieval call binding the contract method 0x08097327.
//
// Solidity: function CONTRACT_DISTRIBUTION_STORAGE() view returns(uint256)
func (_TransferImpl *TransferImplCallerSession) CONTRACTDISTRIBUTIONSTORAGE() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTDISTRIBUTIONSTORAGE(&_TransferImpl.CallOpts)
}

// CONTRACTEXCHANGE is a free data retrieval call binding the contract method 0xa0e3a10e.
//
// Solidity: function CONTRACT_EXCHANGE() view returns(uint256)
func (_TransferImpl *TransferImplCaller) CONTRACTEXCHANGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TransferImpl.contract.Call(opts, out, "CONTRACT_EXCHANGE")
	return *ret0, err
}

// CONTRACTEXCHANGE is a free data retrieval call binding the contract method 0xa0e3a10e.
//
// Solidity: function CONTRACT_EXCHANGE() view returns(uint256)
func (_TransferImpl *TransferImplSession) CONTRACTEXCHANGE() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTEXCHANGE(&_TransferImpl.CallOpts)
}

// CONTRACTEXCHANGE is a free data retrieval call binding the contract method 0xa0e3a10e.
//
// Solidity: function CONTRACT_EXCHANGE() view returns(uint256)
func (_TransferImpl *TransferImplCallerSession) CONTRACTEXCHANGE() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTEXCHANGE(&_TransferImpl.CallOpts)
}

// CONTRACTEXCHANGESTORAGE is a free data retrieval call binding the contract method 0x001570f2.
//
// Solidity: function CONTRACT_EXCHANGE_STORAGE() view returns(uint256)
func (_TransferImpl *TransferImplCaller) CONTRACTEXCHANGESTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TransferImpl.contract.Call(opts, out, "CONTRACT_EXCHANGE_STORAGE")
	return *ret0, err
}

// CONTRACTEXCHANGESTORAGE is a free data retrieval call binding the contract method 0x001570f2.
//
// Solidity: function CONTRACT_EXCHANGE_STORAGE() view returns(uint256)
func (_TransferImpl *TransferImplSession) CONTRACTEXCHANGESTORAGE() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTEXCHANGESTORAGE(&_TransferImpl.CallOpts)
}

// CONTRACTEXCHANGESTORAGE is a free data retrieval call binding the contract method 0x001570f2.
//
// Solidity: function CONTRACT_EXCHANGE_STORAGE() view returns(uint256)
func (_TransferImpl *TransferImplCallerSession) CONTRACTEXCHANGESTORAGE() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTEXCHANGESTORAGE(&_TransferImpl.CallOpts)
}

// CONTRACTFEEGATE is a free data retrieval call binding the contract method 0x545c17a4.
//
// Solidity: function CONTRACT_FEE_GATE() view returns(uint256)
func (_TransferImpl *TransferImplCaller) CONTRACTFEEGATE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TransferImpl.contract.Call(opts, out, "CONTRACT_FEE_GATE")
	return *ret0, err
}

// CONTRACTFEEGATE is a free data retrieval call binding the contract method 0x545c17a4.
//
// Solidity: function CONTRACT_FEE_GATE() view returns(uint256)
func (_TransferImpl *TransferImplSession) CONTRACTFEEGATE() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTFEEGATE(&_TransferImpl.CallOpts)
}

// CONTRACTFEEGATE is a free data retrieval call binding the contract method 0x545c17a4.
//
// Solidity: function CONTRACT_FEE_GATE() view returns(uint256)
func (_TransferImpl *TransferImplCallerSession) CONTRACTFEEGATE() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTFEEGATE(&_TransferImpl.CallOpts)
}

// CONTRACTFEESTORAGE is a free data retrieval call binding the contract method 0xcd5e65a4.
//
// Solidity: function CONTRACT_FEE_STORAGE() view returns(uint256)
func (_TransferImpl *TransferImplCaller) CONTRACTFEESTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TransferImpl.contract.Call(opts, out, "CONTRACT_FEE_STORAGE")
	return *ret0, err
}

// CONTRACTFEESTORAGE is a free data retrieval call binding the contract method 0xcd5e65a4.
//
// Solidity: function CONTRACT_FEE_STORAGE() view returns(uint256)
func (_TransferImpl *TransferImplSession) CONTRACTFEESTORAGE() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTFEESTORAGE(&_TransferImpl.CallOpts)
}

// CONTRACTFEESTORAGE is a free data retrieval call binding the contract method 0xcd5e65a4.
//
// Solidity: function CONTRACT_FEE_STORAGE() view returns(uint256)
func (_TransferImpl *TransferImplCallerSession) CONTRACTFEESTORAGE() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTFEESTORAGE(&_TransferImpl.CallOpts)
}

// CONTRACTGENERATOR is a free data retrieval call binding the contract method 0xccea2b46.
//
// Solidity: function CONTRACT_GENERATOR() view returns(uint256)
func (_TransferImpl *TransferImplCaller) CONTRACTGENERATOR(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TransferImpl.contract.Call(opts, out, "CONTRACT_GENERATOR")
	return *ret0, err
}

// CONTRACTGENERATOR is a free data retrieval call binding the contract method 0xccea2b46.
//
// Solidity: function CONTRACT_GENERATOR() view returns(uint256)
func (_TransferImpl *TransferImplSession) CONTRACTGENERATOR() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTGENERATOR(&_TransferImpl.CallOpts)
}

// CONTRACTGENERATOR is a free data retrieval call binding the contract method 0xccea2b46.
//
// Solidity: function CONTRACT_GENERATOR() view returns(uint256)
func (_TransferImpl *TransferImplCallerSession) CONTRACTGENERATOR() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTGENERATOR(&_TransferImpl.CallOpts)
}

// CONTRACTMININGAGENTGATE is a free data retrieval call binding the contract method 0xb9b4c0f1.
//
// Solidity: function CONTRACT_MINING_AGENT_GATE() view returns(uint256)
func (_TransferImpl *TransferImplCaller) CONTRACTMININGAGENTGATE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TransferImpl.contract.Call(opts, out, "CONTRACT_MINING_AGENT_GATE")
	return *ret0, err
}

// CONTRACTMININGAGENTGATE is a free data retrieval call binding the contract method 0xb9b4c0f1.
//
// Solidity: function CONTRACT_MINING_AGENT_GATE() view returns(uint256)
func (_TransferImpl *TransferImplSession) CONTRACTMININGAGENTGATE() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTMININGAGENTGATE(&_TransferImpl.CallOpts)
}

// CONTRACTMININGAGENTGATE is a free data retrieval call binding the contract method 0xb9b4c0f1.
//
// Solidity: function CONTRACT_MINING_AGENT_GATE() view returns(uint256)
func (_TransferImpl *TransferImplCallerSession) CONTRACTMININGAGENTGATE() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTMININGAGENTGATE(&_TransferImpl.CallOpts)
}

// CONTRACTMININGAGENTSTORAGE is a free data retrieval call binding the contract method 0x83cc588d.
//
// Solidity: function CONTRACT_MINING_AGENT_STORAGE() view returns(uint256)
func (_TransferImpl *TransferImplCaller) CONTRACTMININGAGENTSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TransferImpl.contract.Call(opts, out, "CONTRACT_MINING_AGENT_STORAGE")
	return *ret0, err
}

// CONTRACTMININGAGENTSTORAGE is a free data retrieval call binding the contract method 0x83cc588d.
//
// Solidity: function CONTRACT_MINING_AGENT_STORAGE() view returns(uint256)
func (_TransferImpl *TransferImplSession) CONTRACTMININGAGENTSTORAGE() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTMININGAGENTSTORAGE(&_TransferImpl.CallOpts)
}

// CONTRACTMININGAGENTSTORAGE is a free data retrieval call binding the contract method 0x83cc588d.
//
// Solidity: function CONTRACT_MINING_AGENT_STORAGE() view returns(uint256)
func (_TransferImpl *TransferImplCallerSession) CONTRACTMININGAGENTSTORAGE() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTMININGAGENTSTORAGE(&_TransferImpl.CallOpts)
}

// CONTRACTMSBON is a free data retrieval call binding the contract method 0x28b3a976.
//
// Solidity: function CONTRACT_MSBON() view returns(uint256)
func (_TransferImpl *TransferImplCaller) CONTRACTMSBON(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TransferImpl.contract.Call(opts, out, "CONTRACT_MSBON")
	return *ret0, err
}

// CONTRACTMSBON is a free data retrieval call binding the contract method 0x28b3a976.
//
// Solidity: function CONTRACT_MSBON() view returns(uint256)
func (_TransferImpl *TransferImplSession) CONTRACTMSBON() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTMSBON(&_TransferImpl.CallOpts)
}

// CONTRACTMSBON is a free data retrieval call binding the contract method 0x28b3a976.
//
// Solidity: function CONTRACT_MSBON() view returns(uint256)
func (_TransferImpl *TransferImplCallerSession) CONTRACTMSBON() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTMSBON(&_TransferImpl.CallOpts)
}

// CONTRACTMSBONSTORAGE is a free data retrieval call binding the contract method 0xe04e6061.
//
// Solidity: function CONTRACT_MSBON_STORAGE() view returns(uint256)
func (_TransferImpl *TransferImplCaller) CONTRACTMSBONSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TransferImpl.contract.Call(opts, out, "CONTRACT_MSBON_STORAGE")
	return *ret0, err
}

// CONTRACTMSBONSTORAGE is a free data retrieval call binding the contract method 0xe04e6061.
//
// Solidity: function CONTRACT_MSBON_STORAGE() view returns(uint256)
func (_TransferImpl *TransferImplSession) CONTRACTMSBONSTORAGE() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTMSBONSTORAGE(&_TransferImpl.CallOpts)
}

// CONTRACTMSBONSTORAGE is a free data retrieval call binding the contract method 0xe04e6061.
//
// Solidity: function CONTRACT_MSBON_STORAGE() view returns(uint256)
func (_TransferImpl *TransferImplCallerSession) CONTRACTMSBONSTORAGE() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTMSBONSTORAGE(&_TransferImpl.CallOpts)
}

// CONTRACTMULTISIGADDADDRESS is a free data retrieval call binding the contract method 0xf4f65a5a.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ADDRESS() view returns(uint256)
func (_TransferImpl *TransferImplCaller) CONTRACTMULTISIGADDADDRESS(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TransferImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_ADD_ADDRESS")
	return *ret0, err
}

// CONTRACTMULTISIGADDADDRESS is a free data retrieval call binding the contract method 0xf4f65a5a.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ADDRESS() view returns(uint256)
func (_TransferImpl *TransferImplSession) CONTRACTMULTISIGADDADDRESS() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTMULTISIGADDADDRESS(&_TransferImpl.CallOpts)
}

// CONTRACTMULTISIGADDADDRESS is a free data retrieval call binding the contract method 0xf4f65a5a.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ADDRESS() view returns(uint256)
func (_TransferImpl *TransferImplCallerSession) CONTRACTMULTISIGADDADDRESS() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTMULTISIGADDADDRESS(&_TransferImpl.CallOpts)
}

// CONTRACTMULTISIGADDADMIN is a free data retrieval call binding the contract method 0x83888c68.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ADMIN() view returns(uint256)
func (_TransferImpl *TransferImplCaller) CONTRACTMULTISIGADDADMIN(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TransferImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_ADD_ADMIN")
	return *ret0, err
}

// CONTRACTMULTISIGADDADMIN is a free data retrieval call binding the contract method 0x83888c68.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ADMIN() view returns(uint256)
func (_TransferImpl *TransferImplSession) CONTRACTMULTISIGADDADMIN() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTMULTISIGADDADMIN(&_TransferImpl.CallOpts)
}

// CONTRACTMULTISIGADDADMIN is a free data retrieval call binding the contract method 0x83888c68.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ADMIN() view returns(uint256)
func (_TransferImpl *TransferImplCallerSession) CONTRACTMULTISIGADDADMIN() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTMULTISIGADDADMIN(&_TransferImpl.CallOpts)
}

// CONTRACTMULTISIGADDDISTRIBUTION is a free data retrieval call binding the contract method 0x8f8b81d1.
//
// Solidity: function CONTRACT_MULTISIG_ADD_DISTRIBUTION() view returns(uint256)
func (_TransferImpl *TransferImplCaller) CONTRACTMULTISIGADDDISTRIBUTION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TransferImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_ADD_DISTRIBUTION")
	return *ret0, err
}

// CONTRACTMULTISIGADDDISTRIBUTION is a free data retrieval call binding the contract method 0x8f8b81d1.
//
// Solidity: function CONTRACT_MULTISIG_ADD_DISTRIBUTION() view returns(uint256)
func (_TransferImpl *TransferImplSession) CONTRACTMULTISIGADDDISTRIBUTION() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTMULTISIGADDDISTRIBUTION(&_TransferImpl.CallOpts)
}

// CONTRACTMULTISIGADDDISTRIBUTION is a free data retrieval call binding the contract method 0x8f8b81d1.
//
// Solidity: function CONTRACT_MULTISIG_ADD_DISTRIBUTION() view returns(uint256)
func (_TransferImpl *TransferImplCallerSession) CONTRACTMULTISIGADDDISTRIBUTION() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTMULTISIGADDDISTRIBUTION(&_TransferImpl.CallOpts)
}

// CONTRACTMULTISIGADDFEE is a free data retrieval call binding the contract method 0xc43e8cab.
//
// Solidity: function CONTRACT_MULTISIG_ADD_FEE() view returns(uint256)
func (_TransferImpl *TransferImplCaller) CONTRACTMULTISIGADDFEE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TransferImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_ADD_FEE")
	return *ret0, err
}

// CONTRACTMULTISIGADDFEE is a free data retrieval call binding the contract method 0xc43e8cab.
//
// Solidity: function CONTRACT_MULTISIG_ADD_FEE() view returns(uint256)
func (_TransferImpl *TransferImplSession) CONTRACTMULTISIGADDFEE() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTMULTISIGADDFEE(&_TransferImpl.CallOpts)
}

// CONTRACTMULTISIGADDFEE is a free data retrieval call binding the contract method 0xc43e8cab.
//
// Solidity: function CONTRACT_MULTISIG_ADD_FEE() view returns(uint256)
func (_TransferImpl *TransferImplCallerSession) CONTRACTMULTISIGADDFEE() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTMULTISIGADDFEE(&_TransferImpl.CallOpts)
}

// CONTRACTMULTISIGADDROLE is a free data retrieval call binding the contract method 0x20a0914b.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ROLE() view returns(uint256)
func (_TransferImpl *TransferImplCaller) CONTRACTMULTISIGADDROLE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TransferImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_ADD_ROLE")
	return *ret0, err
}

// CONTRACTMULTISIGADDROLE is a free data retrieval call binding the contract method 0x20a0914b.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ROLE() view returns(uint256)
func (_TransferImpl *TransferImplSession) CONTRACTMULTISIGADDROLE() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTMULTISIGADDROLE(&_TransferImpl.CallOpts)
}

// CONTRACTMULTISIGADDROLE is a free data retrieval call binding the contract method 0x20a0914b.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ROLE() view returns(uint256)
func (_TransferImpl *TransferImplCallerSession) CONTRACTMULTISIGADDROLE() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTMULTISIGADDROLE(&_TransferImpl.CallOpts)
}

// CONTRACTMULTISIGCLEARFEEEXCEPTIONS is a free data retrieval call binding the contract method 0xf74701f2.
//
// Solidity: function CONTRACT_MULTISIG_CLEAR_FEE_EXCEPTIONS() view returns(uint256)
func (_TransferImpl *TransferImplCaller) CONTRACTMULTISIGCLEARFEEEXCEPTIONS(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TransferImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_CLEAR_FEE_EXCEPTIONS")
	return *ret0, err
}

// CONTRACTMULTISIGCLEARFEEEXCEPTIONS is a free data retrieval call binding the contract method 0xf74701f2.
//
// Solidity: function CONTRACT_MULTISIG_CLEAR_FEE_EXCEPTIONS() view returns(uint256)
func (_TransferImpl *TransferImplSession) CONTRACTMULTISIGCLEARFEEEXCEPTIONS() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTMULTISIGCLEARFEEEXCEPTIONS(&_TransferImpl.CallOpts)
}

// CONTRACTMULTISIGCLEARFEEEXCEPTIONS is a free data retrieval call binding the contract method 0xf74701f2.
//
// Solidity: function CONTRACT_MULTISIG_CLEAR_FEE_EXCEPTIONS() view returns(uint256)
func (_TransferImpl *TransferImplCallerSession) CONTRACTMULTISIGCLEARFEEEXCEPTIONS() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTMULTISIGCLEARFEEEXCEPTIONS(&_TransferImpl.CallOpts)
}

// CONTRACTMULTISIGDROPPERMISSION is a free data retrieval call binding the contract method 0x3f3fbc39.
//
// Solidity: function CONTRACT_MULTISIG_DROP_PERMISSION() view returns(uint256)
func (_TransferImpl *TransferImplCaller) CONTRACTMULTISIGDROPPERMISSION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TransferImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_DROP_PERMISSION")
	return *ret0, err
}

// CONTRACTMULTISIGDROPPERMISSION is a free data retrieval call binding the contract method 0x3f3fbc39.
//
// Solidity: function CONTRACT_MULTISIG_DROP_PERMISSION() view returns(uint256)
func (_TransferImpl *TransferImplSession) CONTRACTMULTISIGDROPPERMISSION() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTMULTISIGDROPPERMISSION(&_TransferImpl.CallOpts)
}

// CONTRACTMULTISIGDROPPERMISSION is a free data retrieval call binding the contract method 0x3f3fbc39.
//
// Solidity: function CONTRACT_MULTISIG_DROP_PERMISSION() view returns(uint256)
func (_TransferImpl *TransferImplCallerSession) CONTRACTMULTISIGDROPPERMISSION() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTMULTISIGDROPPERMISSION(&_TransferImpl.CallOpts)
}

// CONTRACTMULTISIGEDITADMIN is a free data retrieval call binding the contract method 0x91b0f051.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_ADMIN() view returns(uint256)
func (_TransferImpl *TransferImplCaller) CONTRACTMULTISIGEDITADMIN(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TransferImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_EDIT_ADMIN")
	return *ret0, err
}

// CONTRACTMULTISIGEDITADMIN is a free data retrieval call binding the contract method 0x91b0f051.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_ADMIN() view returns(uint256)
func (_TransferImpl *TransferImplSession) CONTRACTMULTISIGEDITADMIN() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTMULTISIGEDITADMIN(&_TransferImpl.CallOpts)
}

// CONTRACTMULTISIGEDITADMIN is a free data retrieval call binding the contract method 0x91b0f051.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_ADMIN() view returns(uint256)
func (_TransferImpl *TransferImplCallerSession) CONTRACTMULTISIGEDITADMIN() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTMULTISIGEDITADMIN(&_TransferImpl.CallOpts)
}

// CONTRACTMULTISIGEDITFEE is a free data retrieval call binding the contract method 0x6b2abe17.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_FEE() view returns(uint256)
func (_TransferImpl *TransferImplCaller) CONTRACTMULTISIGEDITFEE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TransferImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_EDIT_FEE")
	return *ret0, err
}

// CONTRACTMULTISIGEDITFEE is a free data retrieval call binding the contract method 0x6b2abe17.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_FEE() view returns(uint256)
func (_TransferImpl *TransferImplSession) CONTRACTMULTISIGEDITFEE() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTMULTISIGEDITFEE(&_TransferImpl.CallOpts)
}

// CONTRACTMULTISIGEDITFEE is a free data retrieval call binding the contract method 0x6b2abe17.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_FEE() view returns(uint256)
func (_TransferImpl *TransferImplCallerSession) CONTRACTMULTISIGEDITFEE() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTMULTISIGEDITFEE(&_TransferImpl.CallOpts)
}

// CONTRACTMULTISIGEDITROLE is a free data retrieval call binding the contract method 0xc5f66412.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_ROLE() view returns(uint256)
func (_TransferImpl *TransferImplCaller) CONTRACTMULTISIGEDITROLE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TransferImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_EDIT_ROLE")
	return *ret0, err
}

// CONTRACTMULTISIGEDITROLE is a free data retrieval call binding the contract method 0xc5f66412.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_ROLE() view returns(uint256)
func (_TransferImpl *TransferImplSession) CONTRACTMULTISIGEDITROLE() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTMULTISIGEDITROLE(&_TransferImpl.CallOpts)
}

// CONTRACTMULTISIGEDITROLE is a free data retrieval call binding the contract method 0xc5f66412.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_ROLE() view returns(uint256)
func (_TransferImpl *TransferImplCallerSession) CONTRACTMULTISIGEDITROLE() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTMULTISIGEDITROLE(&_TransferImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEADMINBYADDRESS is a free data retrieval call binding the contract method 0xc816b091.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_BY_ADDRESS() view returns(uint256)
func (_TransferImpl *TransferImplCaller) CONTRACTMULTISIGREMOVEADMINBYADDRESS(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TransferImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REMOVE_ADMIN_BY_ADDRESS")
	return *ret0, err
}

// CONTRACTMULTISIGREMOVEADMINBYADDRESS is a free data retrieval call binding the contract method 0xc816b091.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_BY_ADDRESS() view returns(uint256)
func (_TransferImpl *TransferImplSession) CONTRACTMULTISIGREMOVEADMINBYADDRESS() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTMULTISIGREMOVEADMINBYADDRESS(&_TransferImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEADMINBYADDRESS is a free data retrieval call binding the contract method 0xc816b091.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_BY_ADDRESS() view returns(uint256)
func (_TransferImpl *TransferImplCallerSession) CONTRACTMULTISIGREMOVEADMINBYADDRESS() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTMULTISIGREMOVEADMINBYADDRESS(&_TransferImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEADMINBYINDEX is a free data retrieval call binding the contract method 0x751fbea2.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_BY_INDEX() view returns(uint256)
func (_TransferImpl *TransferImplCaller) CONTRACTMULTISIGREMOVEADMINBYINDEX(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TransferImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REMOVE_ADMIN_BY_INDEX")
	return *ret0, err
}

// CONTRACTMULTISIGREMOVEADMINBYINDEX is a free data retrieval call binding the contract method 0x751fbea2.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_BY_INDEX() view returns(uint256)
func (_TransferImpl *TransferImplSession) CONTRACTMULTISIGREMOVEADMINBYINDEX() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTMULTISIGREMOVEADMINBYINDEX(&_TransferImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEADMINBYINDEX is a free data retrieval call binding the contract method 0x751fbea2.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_BY_INDEX() view returns(uint256)
func (_TransferImpl *TransferImplCallerSession) CONTRACTMULTISIGREMOVEADMINBYINDEX() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTMULTISIGREMOVEADMINBYINDEX(&_TransferImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEADMINROLE is a free data retrieval call binding the contract method 0x53c00200.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_ROLE() view returns(uint256)
func (_TransferImpl *TransferImplCaller) CONTRACTMULTISIGREMOVEADMINROLE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TransferImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REMOVE_ADMIN_ROLE")
	return *ret0, err
}

// CONTRACTMULTISIGREMOVEADMINROLE is a free data retrieval call binding the contract method 0x53c00200.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_ROLE() view returns(uint256)
func (_TransferImpl *TransferImplSession) CONTRACTMULTISIGREMOVEADMINROLE() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTMULTISIGREMOVEADMINROLE(&_TransferImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEADMINROLE is a free data retrieval call binding the contract method 0x53c00200.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_ROLE() view returns(uint256)
func (_TransferImpl *TransferImplCallerSession) CONTRACTMULTISIGREMOVEADMINROLE() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTMULTISIGREMOVEADMINROLE(&_TransferImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEFEE is a free data retrieval call binding the contract method 0xdb614d25.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_FEE() view returns(uint256)
func (_TransferImpl *TransferImplCaller) CONTRACTMULTISIGREMOVEFEE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TransferImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REMOVE_FEE")
	return *ret0, err
}

// CONTRACTMULTISIGREMOVEFEE is a free data retrieval call binding the contract method 0xdb614d25.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_FEE() view returns(uint256)
func (_TransferImpl *TransferImplSession) CONTRACTMULTISIGREMOVEFEE() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTMULTISIGREMOVEFEE(&_TransferImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEFEE is a free data retrieval call binding the contract method 0xdb614d25.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_FEE() view returns(uint256)
func (_TransferImpl *TransferImplCallerSession) CONTRACTMULTISIGREMOVEFEE() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTMULTISIGREMOVEFEE(&_TransferImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEROLE is a free data retrieval call binding the contract method 0x23648496.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ROLE() view returns(uint256)
func (_TransferImpl *TransferImplCaller) CONTRACTMULTISIGREMOVEROLE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TransferImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REMOVE_ROLE")
	return *ret0, err
}

// CONTRACTMULTISIGREMOVEROLE is a free data retrieval call binding the contract method 0x23648496.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ROLE() view returns(uint256)
func (_TransferImpl *TransferImplSession) CONTRACTMULTISIGREMOVEROLE() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTMULTISIGREMOVEROLE(&_TransferImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEROLE is a free data retrieval call binding the contract method 0x23648496.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ROLE() view returns(uint256)
func (_TransferImpl *TransferImplCallerSession) CONTRACTMULTISIGREMOVEROLE() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTMULTISIGREMOVEROLE(&_TransferImpl.CallOpts)
}

// CONTRACTMULTISIGREPLENISHCAPITALIZATION is a free data retrieval call binding the contract method 0x62d5d523.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_CAPITALIZATION() view returns(uint256)
func (_TransferImpl *TransferImplCaller) CONTRACTMULTISIGREPLENISHCAPITALIZATION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TransferImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REPLENISH_CAPITALIZATION")
	return *ret0, err
}

// CONTRACTMULTISIGREPLENISHCAPITALIZATION is a free data retrieval call binding the contract method 0x62d5d523.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_CAPITALIZATION() view returns(uint256)
func (_TransferImpl *TransferImplSession) CONTRACTMULTISIGREPLENISHCAPITALIZATION() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTMULTISIGREPLENISHCAPITALIZATION(&_TransferImpl.CallOpts)
}

// CONTRACTMULTISIGREPLENISHCAPITALIZATION is a free data retrieval call binding the contract method 0x62d5d523.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_CAPITALIZATION() view returns(uint256)
func (_TransferImpl *TransferImplCallerSession) CONTRACTMULTISIGREPLENISHCAPITALIZATION() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTMULTISIGREPLENISHCAPITALIZATION(&_TransferImpl.CallOpts)
}

// CONTRACTMULTISIGREPLENISHCERTIFICATE is a free data retrieval call binding the contract method 0xfd19b7e1.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_CERTIFICATE() view returns(uint256)
func (_TransferImpl *TransferImplCaller) CONTRACTMULTISIGREPLENISHCERTIFICATE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TransferImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REPLENISH_CERTIFICATE")
	return *ret0, err
}

// CONTRACTMULTISIGREPLENISHCERTIFICATE is a free data retrieval call binding the contract method 0xfd19b7e1.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_CERTIFICATE() view returns(uint256)
func (_TransferImpl *TransferImplSession) CONTRACTMULTISIGREPLENISHCERTIFICATE() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTMULTISIGREPLENISHCERTIFICATE(&_TransferImpl.CallOpts)
}

// CONTRACTMULTISIGREPLENISHCERTIFICATE is a free data retrieval call binding the contract method 0xfd19b7e1.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_CERTIFICATE() view returns(uint256)
func (_TransferImpl *TransferImplCallerSession) CONTRACTMULTISIGREPLENISHCERTIFICATE() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTMULTISIGREPLENISHCERTIFICATE(&_TransferImpl.CallOpts)
}

// CONTRACTMULTISIGREPLENISHEMISSION is a free data retrieval call binding the contract method 0x0cb59fde.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_EMISSION() view returns(uint256)
func (_TransferImpl *TransferImplCaller) CONTRACTMULTISIGREPLENISHEMISSION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TransferImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REPLENISH_EMISSION")
	return *ret0, err
}

// CONTRACTMULTISIGREPLENISHEMISSION is a free data retrieval call binding the contract method 0x0cb59fde.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_EMISSION() view returns(uint256)
func (_TransferImpl *TransferImplSession) CONTRACTMULTISIGREPLENISHEMISSION() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTMULTISIGREPLENISHEMISSION(&_TransferImpl.CallOpts)
}

// CONTRACTMULTISIGREPLENISHEMISSION is a free data retrieval call binding the contract method 0x0cb59fde.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_EMISSION() view returns(uint256)
func (_TransferImpl *TransferImplCallerSession) CONTRACTMULTISIGREPLENISHEMISSION() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTMULTISIGREPLENISHEMISSION(&_TransferImpl.CallOpts)
}

// CONTRACTMULTISIGSETADMINROLE is a free data retrieval call binding the contract method 0x8f209ab7.
//
// Solidity: function CONTRACT_MULTISIG_SET_ADMIN_ROLE() view returns(uint256)
func (_TransferImpl *TransferImplCaller) CONTRACTMULTISIGSETADMINROLE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TransferImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_SET_ADMIN_ROLE")
	return *ret0, err
}

// CONTRACTMULTISIGSETADMINROLE is a free data retrieval call binding the contract method 0x8f209ab7.
//
// Solidity: function CONTRACT_MULTISIG_SET_ADMIN_ROLE() view returns(uint256)
func (_TransferImpl *TransferImplSession) CONTRACTMULTISIGSETADMINROLE() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTMULTISIGSETADMINROLE(&_TransferImpl.CallOpts)
}

// CONTRACTMULTISIGSETADMINROLE is a free data retrieval call binding the contract method 0x8f209ab7.
//
// Solidity: function CONTRACT_MULTISIG_SET_ADMIN_ROLE() view returns(uint256)
func (_TransferImpl *TransferImplCallerSession) CONTRACTMULTISIGSETADMINROLE() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTMULTISIGSETADMINROLE(&_TransferImpl.CallOpts)
}

// CONTRACTMULTISIGSETCONFIRMATIONRATE is a free data retrieval call binding the contract method 0xa997279b.
//
// Solidity: function CONTRACT_MULTISIG_SET_CONFIRMATION_RATE() view returns(uint256)
func (_TransferImpl *TransferImplCaller) CONTRACTMULTISIGSETCONFIRMATIONRATE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TransferImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_SET_CONFIRMATION_RATE")
	return *ret0, err
}

// CONTRACTMULTISIGSETCONFIRMATIONRATE is a free data retrieval call binding the contract method 0xa997279b.
//
// Solidity: function CONTRACT_MULTISIG_SET_CONFIRMATION_RATE() view returns(uint256)
func (_TransferImpl *TransferImplSession) CONTRACTMULTISIGSETCONFIRMATIONRATE() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTMULTISIGSETCONFIRMATIONRATE(&_TransferImpl.CallOpts)
}

// CONTRACTMULTISIGSETCONFIRMATIONRATE is a free data retrieval call binding the contract method 0xa997279b.
//
// Solidity: function CONTRACT_MULTISIG_SET_CONFIRMATION_RATE() view returns(uint256)
func (_TransferImpl *TransferImplCallerSession) CONTRACTMULTISIGSETCONFIRMATIONRATE() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTMULTISIGSETCONFIRMATIONRATE(&_TransferImpl.CallOpts)
}

// CONTRACTMULTISIGSETFEEDISTRIBUTION is a free data retrieval call binding the contract method 0x0b76e64d.
//
// Solidity: function CONTRACT_MULTISIG_SET_FEE_DISTRIBUTION() view returns(uint256)
func (_TransferImpl *TransferImplCaller) CONTRACTMULTISIGSETFEEDISTRIBUTION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TransferImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_SET_FEE_DISTRIBUTION")
	return *ret0, err
}

// CONTRACTMULTISIGSETFEEDISTRIBUTION is a free data retrieval call binding the contract method 0x0b76e64d.
//
// Solidity: function CONTRACT_MULTISIG_SET_FEE_DISTRIBUTION() view returns(uint256)
func (_TransferImpl *TransferImplSession) CONTRACTMULTISIGSETFEEDISTRIBUTION() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTMULTISIGSETFEEDISTRIBUTION(&_TransferImpl.CallOpts)
}

// CONTRACTMULTISIGSETFEEDISTRIBUTION is a free data retrieval call binding the contract method 0x0b76e64d.
//
// Solidity: function CONTRACT_MULTISIG_SET_FEE_DISTRIBUTION() view returns(uint256)
func (_TransferImpl *TransferImplCallerSession) CONTRACTMULTISIGSETFEEDISTRIBUTION() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTMULTISIGSETFEEDISTRIBUTION(&_TransferImpl.CallOpts)
}

// CONTRACTMULTISIGSETFEEEXCEPTIONS is a free data retrieval call binding the contract method 0x5e6b28a0.
//
// Solidity: function CONTRACT_MULTISIG_SET_FEE_EXCEPTIONS() view returns(uint256)
func (_TransferImpl *TransferImplCaller) CONTRACTMULTISIGSETFEEEXCEPTIONS(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TransferImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_SET_FEE_EXCEPTIONS")
	return *ret0, err
}

// CONTRACTMULTISIGSETFEEEXCEPTIONS is a free data retrieval call binding the contract method 0x5e6b28a0.
//
// Solidity: function CONTRACT_MULTISIG_SET_FEE_EXCEPTIONS() view returns(uint256)
func (_TransferImpl *TransferImplSession) CONTRACTMULTISIGSETFEEEXCEPTIONS() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTMULTISIGSETFEEEXCEPTIONS(&_TransferImpl.CallOpts)
}

// CONTRACTMULTISIGSETFEEEXCEPTIONS is a free data retrieval call binding the contract method 0x5e6b28a0.
//
// Solidity: function CONTRACT_MULTISIG_SET_FEE_EXCEPTIONS() view returns(uint256)
func (_TransferImpl *TransferImplCallerSession) CONTRACTMULTISIGSETFEEEXCEPTIONS() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTMULTISIGSETFEEEXCEPTIONS(&_TransferImpl.CallOpts)
}

// CONTRACTMULTISIGSETMAXFEE is a free data retrieval call binding the contract method 0xa5f59013.
//
// Solidity: function CONTRACT_MULTISIG_SET_MAX_FEE() view returns(uint256)
func (_TransferImpl *TransferImplCaller) CONTRACTMULTISIGSETMAXFEE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TransferImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_SET_MAX_FEE")
	return *ret0, err
}

// CONTRACTMULTISIGSETMAXFEE is a free data retrieval call binding the contract method 0xa5f59013.
//
// Solidity: function CONTRACT_MULTISIG_SET_MAX_FEE() view returns(uint256)
func (_TransferImpl *TransferImplSession) CONTRACTMULTISIGSETMAXFEE() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTMULTISIGSETMAXFEE(&_TransferImpl.CallOpts)
}

// CONTRACTMULTISIGSETMAXFEE is a free data retrieval call binding the contract method 0xa5f59013.
//
// Solidity: function CONTRACT_MULTISIG_SET_MAX_FEE() view returns(uint256)
func (_TransferImpl *TransferImplCallerSession) CONTRACTMULTISIGSETMAXFEE() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTMULTISIGSETMAXFEE(&_TransferImpl.CallOpts)
}

// CONTRACTMULTISIGSETMINAMOUNT is a free data retrieval call binding the contract method 0xaba147b6.
//
// Solidity: function CONTRACT_MULTISIG_SET_MIN_AMOUNT() view returns(uint256)
func (_TransferImpl *TransferImplCaller) CONTRACTMULTISIGSETMINAMOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TransferImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_SET_MIN_AMOUNT")
	return *ret0, err
}

// CONTRACTMULTISIGSETMINAMOUNT is a free data retrieval call binding the contract method 0xaba147b6.
//
// Solidity: function CONTRACT_MULTISIG_SET_MIN_AMOUNT() view returns(uint256)
func (_TransferImpl *TransferImplSession) CONTRACTMULTISIGSETMINAMOUNT() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTMULTISIGSETMINAMOUNT(&_TransferImpl.CallOpts)
}

// CONTRACTMULTISIGSETMINAMOUNT is a free data retrieval call binding the contract method 0xaba147b6.
//
// Solidity: function CONTRACT_MULTISIG_SET_MIN_AMOUNT() view returns(uint256)
func (_TransferImpl *TransferImplCallerSession) CONTRACTMULTISIGSETMINAMOUNT() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTMULTISIGSETMINAMOUNT(&_TransferImpl.CallOpts)
}

// CONTRACTMULTISIGSETMINFEE is a free data retrieval call binding the contract method 0xf2d1904f.
//
// Solidity: function CONTRACT_MULTISIG_SET_MIN_FEE() view returns(uint256)
func (_TransferImpl *TransferImplCaller) CONTRACTMULTISIGSETMINFEE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TransferImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_SET_MIN_FEE")
	return *ret0, err
}

// CONTRACTMULTISIGSETMINFEE is a free data retrieval call binding the contract method 0xf2d1904f.
//
// Solidity: function CONTRACT_MULTISIG_SET_MIN_FEE() view returns(uint256)
func (_TransferImpl *TransferImplSession) CONTRACTMULTISIGSETMINFEE() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTMULTISIGSETMINFEE(&_TransferImpl.CallOpts)
}

// CONTRACTMULTISIGSETMINFEE is a free data retrieval call binding the contract method 0xf2d1904f.
//
// Solidity: function CONTRACT_MULTISIG_SET_MIN_FEE() view returns(uint256)
func (_TransferImpl *TransferImplCallerSession) CONTRACTMULTISIGSETMINFEE() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTMULTISIGSETMINFEE(&_TransferImpl.CallOpts)
}

// CONTRACTMULTISIGSETPERMISSION is a free data retrieval call binding the contract method 0xb0e91914.
//
// Solidity: function CONTRACT_MULTISIG_SET_PERMISSION() view returns(uint256)
func (_TransferImpl *TransferImplCaller) CONTRACTMULTISIGSETPERMISSION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TransferImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_SET_PERMISSION")
	return *ret0, err
}

// CONTRACTMULTISIGSETPERMISSION is a free data retrieval call binding the contract method 0xb0e91914.
//
// Solidity: function CONTRACT_MULTISIG_SET_PERMISSION() view returns(uint256)
func (_TransferImpl *TransferImplSession) CONTRACTMULTISIGSETPERMISSION() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTMULTISIGSETPERMISSION(&_TransferImpl.CallOpts)
}

// CONTRACTMULTISIGSETPERMISSION is a free data retrieval call binding the contract method 0xb0e91914.
//
// Solidity: function CONTRACT_MULTISIG_SET_PERMISSION() view returns(uint256)
func (_TransferImpl *TransferImplCallerSession) CONTRACTMULTISIGSETPERMISSION() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTMULTISIGSETPERMISSION(&_TransferImpl.CallOpts)
}

// CONTRACTMULTISIGSTORAGE is a free data retrieval call binding the contract method 0xe7e1a303.
//
// Solidity: function CONTRACT_MULTISIG_STORAGE() view returns(uint256)
func (_TransferImpl *TransferImplCaller) CONTRACTMULTISIGSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TransferImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_STORAGE")
	return *ret0, err
}

// CONTRACTMULTISIGSTORAGE is a free data retrieval call binding the contract method 0xe7e1a303.
//
// Solidity: function CONTRACT_MULTISIG_STORAGE() view returns(uint256)
func (_TransferImpl *TransferImplSession) CONTRACTMULTISIGSTORAGE() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTMULTISIGSTORAGE(&_TransferImpl.CallOpts)
}

// CONTRACTMULTISIGSTORAGE is a free data retrieval call binding the contract method 0xe7e1a303.
//
// Solidity: function CONTRACT_MULTISIG_STORAGE() view returns(uint256)
func (_TransferImpl *TransferImplCallerSession) CONTRACTMULTISIGSTORAGE() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTMULTISIGSTORAGE(&_TransferImpl.CallOpts)
}

// CONTRACTOTC is a free data retrieval call binding the contract method 0x5eceac70.
//
// Solidity: function CONTRACT_OTC() view returns(uint256)
func (_TransferImpl *TransferImplCaller) CONTRACTOTC(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TransferImpl.contract.Call(opts, out, "CONTRACT_OTC")
	return *ret0, err
}

// CONTRACTOTC is a free data retrieval call binding the contract method 0x5eceac70.
//
// Solidity: function CONTRACT_OTC() view returns(uint256)
func (_TransferImpl *TransferImplSession) CONTRACTOTC() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTOTC(&_TransferImpl.CallOpts)
}

// CONTRACTOTC is a free data retrieval call binding the contract method 0x5eceac70.
//
// Solidity: function CONTRACT_OTC() view returns(uint256)
func (_TransferImpl *TransferImplCallerSession) CONTRACTOTC() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTOTC(&_TransferImpl.CallOpts)
}

// CONTRACTOTCSTORAGE is a free data retrieval call binding the contract method 0x825b476f.
//
// Solidity: function CONTRACT_OTC_STORAGE() view returns(uint256)
func (_TransferImpl *TransferImplCaller) CONTRACTOTCSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TransferImpl.contract.Call(opts, out, "CONTRACT_OTC_STORAGE")
	return *ret0, err
}

// CONTRACTOTCSTORAGE is a free data retrieval call binding the contract method 0x825b476f.
//
// Solidity: function CONTRACT_OTC_STORAGE() view returns(uint256)
func (_TransferImpl *TransferImplSession) CONTRACTOTCSTORAGE() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTOTCSTORAGE(&_TransferImpl.CallOpts)
}

// CONTRACTOTCSTORAGE is a free data retrieval call binding the contract method 0x825b476f.
//
// Solidity: function CONTRACT_OTC_STORAGE() view returns(uint256)
func (_TransferImpl *TransferImplCallerSession) CONTRACTOTCSTORAGE() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTOTCSTORAGE(&_TransferImpl.CallOpts)
}

// CONTRACTREPLENISHCAPITALIZATION is a free data retrieval call binding the contract method 0xff5e09d8.
//
// Solidity: function CONTRACT_REPLENISH_CAPITALIZATION() view returns(uint256)
func (_TransferImpl *TransferImplCaller) CONTRACTREPLENISHCAPITALIZATION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TransferImpl.contract.Call(opts, out, "CONTRACT_REPLENISH_CAPITALIZATION")
	return *ret0, err
}

// CONTRACTREPLENISHCAPITALIZATION is a free data retrieval call binding the contract method 0xff5e09d8.
//
// Solidity: function CONTRACT_REPLENISH_CAPITALIZATION() view returns(uint256)
func (_TransferImpl *TransferImplSession) CONTRACTREPLENISHCAPITALIZATION() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTREPLENISHCAPITALIZATION(&_TransferImpl.CallOpts)
}

// CONTRACTREPLENISHCAPITALIZATION is a free data retrieval call binding the contract method 0xff5e09d8.
//
// Solidity: function CONTRACT_REPLENISH_CAPITALIZATION() view returns(uint256)
func (_TransferImpl *TransferImplCallerSession) CONTRACTREPLENISHCAPITALIZATION() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTREPLENISHCAPITALIZATION(&_TransferImpl.CallOpts)
}

// CONTRACTRESERVEDALIASSTORAGE is a free data retrieval call binding the contract method 0x70a94aa7.
//
// Solidity: function CONTRACT_RESERVED_ALIAS_STORAGE() view returns(uint256)
func (_TransferImpl *TransferImplCaller) CONTRACTRESERVEDALIASSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TransferImpl.contract.Call(opts, out, "CONTRACT_RESERVED_ALIAS_STORAGE")
	return *ret0, err
}

// CONTRACTRESERVEDALIASSTORAGE is a free data retrieval call binding the contract method 0x70a94aa7.
//
// Solidity: function CONTRACT_RESERVED_ALIAS_STORAGE() view returns(uint256)
func (_TransferImpl *TransferImplSession) CONTRACTRESERVEDALIASSTORAGE() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTRESERVEDALIASSTORAGE(&_TransferImpl.CallOpts)
}

// CONTRACTRESERVEDALIASSTORAGE is a free data retrieval call binding the contract method 0x70a94aa7.
//
// Solidity: function CONTRACT_RESERVED_ALIAS_STORAGE() view returns(uint256)
func (_TransferImpl *TransferImplCallerSession) CONTRACTRESERVEDALIASSTORAGE() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTRESERVEDALIASSTORAGE(&_TransferImpl.CallOpts)
}

// CONTRACTROLESTORAGE is a free data retrieval call binding the contract method 0xb938235b.
//
// Solidity: function CONTRACT_ROLE_STORAGE() view returns(uint256)
func (_TransferImpl *TransferImplCaller) CONTRACTROLESTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TransferImpl.contract.Call(opts, out, "CONTRACT_ROLE_STORAGE")
	return *ret0, err
}

// CONTRACTROLESTORAGE is a free data retrieval call binding the contract method 0xb938235b.
//
// Solidity: function CONTRACT_ROLE_STORAGE() view returns(uint256)
func (_TransferImpl *TransferImplSession) CONTRACTROLESTORAGE() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTROLESTORAGE(&_TransferImpl.CallOpts)
}

// CONTRACTROLESTORAGE is a free data retrieval call binding the contract method 0xb938235b.
//
// Solidity: function CONTRACT_ROLE_STORAGE() view returns(uint256)
func (_TransferImpl *TransferImplCallerSession) CONTRACTROLESTORAGE() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTROLESTORAGE(&_TransferImpl.CallOpts)
}

// CONTRACTSAFETRANSFERSTORAGE is a free data retrieval call binding the contract method 0xe7e45392.
//
// Solidity: function CONTRACT_SAFE_TRANSFER_STORAGE() view returns(uint256)
func (_TransferImpl *TransferImplCaller) CONTRACTSAFETRANSFERSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TransferImpl.contract.Call(opts, out, "CONTRACT_SAFE_TRANSFER_STORAGE")
	return *ret0, err
}

// CONTRACTSAFETRANSFERSTORAGE is a free data retrieval call binding the contract method 0xe7e45392.
//
// Solidity: function CONTRACT_SAFE_TRANSFER_STORAGE() view returns(uint256)
func (_TransferImpl *TransferImplSession) CONTRACTSAFETRANSFERSTORAGE() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTSAFETRANSFERSTORAGE(&_TransferImpl.CallOpts)
}

// CONTRACTSAFETRANSFERSTORAGE is a free data retrieval call binding the contract method 0xe7e45392.
//
// Solidity: function CONTRACT_SAFE_TRANSFER_STORAGE() view returns(uint256)
func (_TransferImpl *TransferImplCallerSession) CONTRACTSAFETRANSFERSTORAGE() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTSAFETRANSFERSTORAGE(&_TransferImpl.CallOpts)
}

// CONTRACTTRANSFER is a free data retrieval call binding the contract method 0xef94fe38.
//
// Solidity: function CONTRACT_TRANSFER() view returns(uint256)
func (_TransferImpl *TransferImplCaller) CONTRACTTRANSFER(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TransferImpl.contract.Call(opts, out, "CONTRACT_TRANSFER")
	return *ret0, err
}

// CONTRACTTRANSFER is a free data retrieval call binding the contract method 0xef94fe38.
//
// Solidity: function CONTRACT_TRANSFER() view returns(uint256)
func (_TransferImpl *TransferImplSession) CONTRACTTRANSFER() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTTRANSFER(&_TransferImpl.CallOpts)
}

// CONTRACTTRANSFER is a free data retrieval call binding the contract method 0xef94fe38.
//
// Solidity: function CONTRACT_TRANSFER() view returns(uint256)
func (_TransferImpl *TransferImplCallerSession) CONTRACTTRANSFER() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTTRANSFER(&_TransferImpl.CallOpts)
}

// CONTRACTTRANSFERFROMCAPITALIZATION is a free data retrieval call binding the contract method 0xc8f8decb.
//
// Solidity: function CONTRACT_TRANSFER_FROM_CAPITALIZATION() view returns(uint256)
func (_TransferImpl *TransferImplCaller) CONTRACTTRANSFERFROMCAPITALIZATION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TransferImpl.contract.Call(opts, out, "CONTRACT_TRANSFER_FROM_CAPITALIZATION")
	return *ret0, err
}

// CONTRACTTRANSFERFROMCAPITALIZATION is a free data retrieval call binding the contract method 0xc8f8decb.
//
// Solidity: function CONTRACT_TRANSFER_FROM_CAPITALIZATION() view returns(uint256)
func (_TransferImpl *TransferImplSession) CONTRACTTRANSFERFROMCAPITALIZATION() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTTRANSFERFROMCAPITALIZATION(&_TransferImpl.CallOpts)
}

// CONTRACTTRANSFERFROMCAPITALIZATION is a free data retrieval call binding the contract method 0xc8f8decb.
//
// Solidity: function CONTRACT_TRANSFER_FROM_CAPITALIZATION() view returns(uint256)
func (_TransferImpl *TransferImplCallerSession) CONTRACTTRANSFERFROMCAPITALIZATION() (*big.Int, error) {
	return _TransferImpl.Contract.CONTRACTTRANSFERFROMCAPITALIZATION(&_TransferImpl.CallOpts)
}

// FEEOPERATIONCREATESAFETRANSFER is a free data retrieval call binding the contract method 0xf46d727d.
//
// Solidity: function FEE_OPERATION_CREATE_SAFE_TRANSFER() view returns(uint256)
func (_TransferImpl *TransferImplCaller) FEEOPERATIONCREATESAFETRANSFER(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TransferImpl.contract.Call(opts, out, "FEE_OPERATION_CREATE_SAFE_TRANSFER")
	return *ret0, err
}

// FEEOPERATIONCREATESAFETRANSFER is a free data retrieval call binding the contract method 0xf46d727d.
//
// Solidity: function FEE_OPERATION_CREATE_SAFE_TRANSFER() view returns(uint256)
func (_TransferImpl *TransferImplSession) FEEOPERATIONCREATESAFETRANSFER() (*big.Int, error) {
	return _TransferImpl.Contract.FEEOPERATIONCREATESAFETRANSFER(&_TransferImpl.CallOpts)
}

// FEEOPERATIONCREATESAFETRANSFER is a free data retrieval call binding the contract method 0xf46d727d.
//
// Solidity: function FEE_OPERATION_CREATE_SAFE_TRANSFER() view returns(uint256)
func (_TransferImpl *TransferImplCallerSession) FEEOPERATIONCREATESAFETRANSFER() (*big.Int, error) {
	return _TransferImpl.Contract.FEEOPERATIONCREATESAFETRANSFER(&_TransferImpl.CallOpts)
}

// FEEOPERATIONCREATESAFETRANSFERALL is a free data retrieval call binding the contract method 0x9d558bef.
//
// Solidity: function FEE_OPERATION_CREATE_SAFE_TRANSFER_ALL() view returns(uint256)
func (_TransferImpl *TransferImplCaller) FEEOPERATIONCREATESAFETRANSFERALL(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TransferImpl.contract.Call(opts, out, "FEE_OPERATION_CREATE_SAFE_TRANSFER_ALL")
	return *ret0, err
}

// FEEOPERATIONCREATESAFETRANSFERALL is a free data retrieval call binding the contract method 0x9d558bef.
//
// Solidity: function FEE_OPERATION_CREATE_SAFE_TRANSFER_ALL() view returns(uint256)
func (_TransferImpl *TransferImplSession) FEEOPERATIONCREATESAFETRANSFERALL() (*big.Int, error) {
	return _TransferImpl.Contract.FEEOPERATIONCREATESAFETRANSFERALL(&_TransferImpl.CallOpts)
}

// FEEOPERATIONCREATESAFETRANSFERALL is a free data retrieval call binding the contract method 0x9d558bef.
//
// Solidity: function FEE_OPERATION_CREATE_SAFE_TRANSFER_ALL() view returns(uint256)
func (_TransferImpl *TransferImplCallerSession) FEEOPERATIONCREATESAFETRANSFERALL() (*big.Int, error) {
	return _TransferImpl.Contract.FEEOPERATIONCREATESAFETRANSFERALL(&_TransferImpl.CallOpts)
}

// FEEOPERATIONFRAMETRANSFER is a free data retrieval call binding the contract method 0xc0c92120.
//
// Solidity: function FEE_OPERATION_FRAME_TRANSFER() view returns(uint256)
func (_TransferImpl *TransferImplCaller) FEEOPERATIONFRAMETRANSFER(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TransferImpl.contract.Call(opts, out, "FEE_OPERATION_FRAME_TRANSFER")
	return *ret0, err
}

// FEEOPERATIONFRAMETRANSFER is a free data retrieval call binding the contract method 0xc0c92120.
//
// Solidity: function FEE_OPERATION_FRAME_TRANSFER() view returns(uint256)
func (_TransferImpl *TransferImplSession) FEEOPERATIONFRAMETRANSFER() (*big.Int, error) {
	return _TransferImpl.Contract.FEEOPERATIONFRAMETRANSFER(&_TransferImpl.CallOpts)
}

// FEEOPERATIONFRAMETRANSFER is a free data retrieval call binding the contract method 0xc0c92120.
//
// Solidity: function FEE_OPERATION_FRAME_TRANSFER() view returns(uint256)
func (_TransferImpl *TransferImplCallerSession) FEEOPERATIONFRAMETRANSFER() (*big.Int, error) {
	return _TransferImpl.Contract.FEEOPERATIONFRAMETRANSFER(&_TransferImpl.CallOpts)
}

// FEEOPERATIONQUICKTRANSFER is a free data retrieval call binding the contract method 0xbd64b596.
//
// Solidity: function FEE_OPERATION_QUICK_TRANSFER() view returns(uint256)
func (_TransferImpl *TransferImplCaller) FEEOPERATIONQUICKTRANSFER(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TransferImpl.contract.Call(opts, out, "FEE_OPERATION_QUICK_TRANSFER")
	return *ret0, err
}

// FEEOPERATIONQUICKTRANSFER is a free data retrieval call binding the contract method 0xbd64b596.
//
// Solidity: function FEE_OPERATION_QUICK_TRANSFER() view returns(uint256)
func (_TransferImpl *TransferImplSession) FEEOPERATIONQUICKTRANSFER() (*big.Int, error) {
	return _TransferImpl.Contract.FEEOPERATIONQUICKTRANSFER(&_TransferImpl.CallOpts)
}

// FEEOPERATIONQUICKTRANSFER is a free data retrieval call binding the contract method 0xbd64b596.
//
// Solidity: function FEE_OPERATION_QUICK_TRANSFER() view returns(uint256)
func (_TransferImpl *TransferImplCallerSession) FEEOPERATIONQUICKTRANSFER() (*big.Int, error) {
	return _TransferImpl.Contract.FEEOPERATIONQUICKTRANSFER(&_TransferImpl.CallOpts)
}

// FEEOPERATIONQUICKTRANSFERALL is a free data retrieval call binding the contract method 0x69cc9631.
//
// Solidity: function FEE_OPERATION_QUICK_TRANSFER_ALL() view returns(uint256)
func (_TransferImpl *TransferImplCaller) FEEOPERATIONQUICKTRANSFERALL(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TransferImpl.contract.Call(opts, out, "FEE_OPERATION_QUICK_TRANSFER_ALL")
	return *ret0, err
}

// FEEOPERATIONQUICKTRANSFERALL is a free data retrieval call binding the contract method 0x69cc9631.
//
// Solidity: function FEE_OPERATION_QUICK_TRANSFER_ALL() view returns(uint256)
func (_TransferImpl *TransferImplSession) FEEOPERATIONQUICKTRANSFERALL() (*big.Int, error) {
	return _TransferImpl.Contract.FEEOPERATIONQUICKTRANSFERALL(&_TransferImpl.CallOpts)
}

// FEEOPERATIONQUICKTRANSFERALL is a free data retrieval call binding the contract method 0x69cc9631.
//
// Solidity: function FEE_OPERATION_QUICK_TRANSFER_ALL() view returns(uint256)
func (_TransferImpl *TransferImplCallerSession) FEEOPERATIONQUICKTRANSFERALL() (*big.Int, error) {
	return _TransferImpl.Contract.FEEOPERATIONQUICKTRANSFERALL(&_TransferImpl.CallOpts)
}

// FEEOPERATIONWPCSAFETRANSFER is a free data retrieval call binding the contract method 0x33594855.
//
// Solidity: function FEE_OPERATION_WPC_SAFE_TRANSFER() view returns(uint256)
func (_TransferImpl *TransferImplCaller) FEEOPERATIONWPCSAFETRANSFER(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TransferImpl.contract.Call(opts, out, "FEE_OPERATION_WPC_SAFE_TRANSFER")
	return *ret0, err
}

// FEEOPERATIONWPCSAFETRANSFER is a free data retrieval call binding the contract method 0x33594855.
//
// Solidity: function FEE_OPERATION_WPC_SAFE_TRANSFER() view returns(uint256)
func (_TransferImpl *TransferImplSession) FEEOPERATIONWPCSAFETRANSFER() (*big.Int, error) {
	return _TransferImpl.Contract.FEEOPERATIONWPCSAFETRANSFER(&_TransferImpl.CallOpts)
}

// FEEOPERATIONWPCSAFETRANSFER is a free data retrieval call binding the contract method 0x33594855.
//
// Solidity: function FEE_OPERATION_WPC_SAFE_TRANSFER() view returns(uint256)
func (_TransferImpl *TransferImplCallerSession) FEEOPERATIONWPCSAFETRANSFER() (*big.Int, error) {
	return _TransferImpl.Contract.FEEOPERATIONWPCSAFETRANSFER(&_TransferImpl.CallOpts)
}

// FEEOPERATIONWPCSAFETRANSFERALL is a free data retrieval call binding the contract method 0xf0ef7f20.
//
// Solidity: function FEE_OPERATION_WPC_SAFE_TRANSFER_ALL() view returns(uint256)
func (_TransferImpl *TransferImplCaller) FEEOPERATIONWPCSAFETRANSFERALL(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TransferImpl.contract.Call(opts, out, "FEE_OPERATION_WPC_SAFE_TRANSFER_ALL")
	return *ret0, err
}

// FEEOPERATIONWPCSAFETRANSFERALL is a free data retrieval call binding the contract method 0xf0ef7f20.
//
// Solidity: function FEE_OPERATION_WPC_SAFE_TRANSFER_ALL() view returns(uint256)
func (_TransferImpl *TransferImplSession) FEEOPERATIONWPCSAFETRANSFERALL() (*big.Int, error) {
	return _TransferImpl.Contract.FEEOPERATIONWPCSAFETRANSFERALL(&_TransferImpl.CallOpts)
}

// FEEOPERATIONWPCSAFETRANSFERALL is a free data retrieval call binding the contract method 0xf0ef7f20.
//
// Solidity: function FEE_OPERATION_WPC_SAFE_TRANSFER_ALL() view returns(uint256)
func (_TransferImpl *TransferImplCallerSession) FEEOPERATIONWPCSAFETRANSFERALL() (*big.Int, error) {
	return _TransferImpl.Contract.FEEOPERATIONWPCSAFETRANSFERALL(&_TransferImpl.CallOpts)
}

// PERMISSIONACCESSRESTORATION is a free data retrieval call binding the contract method 0x5c3053bf.
//
// Solidity: function PERMISSION_ACCESS_RESTORATION() view returns(uint256)
func (_TransferImpl *TransferImplCaller) PERMISSIONACCESSRESTORATION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TransferImpl.contract.Call(opts, out, "PERMISSION_ACCESS_RESTORATION")
	return *ret0, err
}

// PERMISSIONACCESSRESTORATION is a free data retrieval call binding the contract method 0x5c3053bf.
//
// Solidity: function PERMISSION_ACCESS_RESTORATION() view returns(uint256)
func (_TransferImpl *TransferImplSession) PERMISSIONACCESSRESTORATION() (*big.Int, error) {
	return _TransferImpl.Contract.PERMISSIONACCESSRESTORATION(&_TransferImpl.CallOpts)
}

// PERMISSIONACCESSRESTORATION is a free data retrieval call binding the contract method 0x5c3053bf.
//
// Solidity: function PERMISSION_ACCESS_RESTORATION() view returns(uint256)
func (_TransferImpl *TransferImplCallerSession) PERMISSIONACCESSRESTORATION() (*big.Int, error) {
	return _TransferImpl.Contract.PERMISSIONACCESSRESTORATION(&_TransferImpl.CallOpts)
}

// PERMISSIONADMINSTORAGE is a free data retrieval call binding the contract method 0x0c9fc0db.
//
// Solidity: function PERMISSION_ADMIN_STORAGE() view returns(uint256)
func (_TransferImpl *TransferImplCaller) PERMISSIONADMINSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TransferImpl.contract.Call(opts, out, "PERMISSION_ADMIN_STORAGE")
	return *ret0, err
}

// PERMISSIONADMINSTORAGE is a free data retrieval call binding the contract method 0x0c9fc0db.
//
// Solidity: function PERMISSION_ADMIN_STORAGE() view returns(uint256)
func (_TransferImpl *TransferImplSession) PERMISSIONADMINSTORAGE() (*big.Int, error) {
	return _TransferImpl.Contract.PERMISSIONADMINSTORAGE(&_TransferImpl.CallOpts)
}

// PERMISSIONADMINSTORAGE is a free data retrieval call binding the contract method 0x0c9fc0db.
//
// Solidity: function PERMISSION_ADMIN_STORAGE() view returns(uint256)
func (_TransferImpl *TransferImplCallerSession) PERMISSIONADMINSTORAGE() (*big.Int, error) {
	return _TransferImpl.Contract.PERMISSIONADMINSTORAGE(&_TransferImpl.CallOpts)
}

// PERMISSIONCONTRACTSSTORAGE is a free data retrieval call binding the contract method 0x1516e83b.
//
// Solidity: function PERMISSION_CONTRACTS_STORAGE() view returns(uint256)
func (_TransferImpl *TransferImplCaller) PERMISSIONCONTRACTSSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TransferImpl.contract.Call(opts, out, "PERMISSION_CONTRACTS_STORAGE")
	return *ret0, err
}

// PERMISSIONCONTRACTSSTORAGE is a free data retrieval call binding the contract method 0x1516e83b.
//
// Solidity: function PERMISSION_CONTRACTS_STORAGE() view returns(uint256)
func (_TransferImpl *TransferImplSession) PERMISSIONCONTRACTSSTORAGE() (*big.Int, error) {
	return _TransferImpl.Contract.PERMISSIONCONTRACTSSTORAGE(&_TransferImpl.CallOpts)
}

// PERMISSIONCONTRACTSSTORAGE is a free data retrieval call binding the contract method 0x1516e83b.
//
// Solidity: function PERMISSION_CONTRACTS_STORAGE() view returns(uint256)
func (_TransferImpl *TransferImplCallerSession) PERMISSIONCONTRACTSSTORAGE() (*big.Int, error) {
	return _TransferImpl.Contract.PERMISSIONCONTRACTSSTORAGE(&_TransferImpl.CallOpts)
}

// PERMISSIONCONTRACTDEPLOY is a free data retrieval call binding the contract method 0x257a4a04.
//
// Solidity: function PERMISSION_CONTRACT_DEPLOY() view returns(uint256)
func (_TransferImpl *TransferImplCaller) PERMISSIONCONTRACTDEPLOY(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TransferImpl.contract.Call(opts, out, "PERMISSION_CONTRACT_DEPLOY")
	return *ret0, err
}

// PERMISSIONCONTRACTDEPLOY is a free data retrieval call binding the contract method 0x257a4a04.
//
// Solidity: function PERMISSION_CONTRACT_DEPLOY() view returns(uint256)
func (_TransferImpl *TransferImplSession) PERMISSIONCONTRACTDEPLOY() (*big.Int, error) {
	return _TransferImpl.Contract.PERMISSIONCONTRACTDEPLOY(&_TransferImpl.CallOpts)
}

// PERMISSIONCONTRACTDEPLOY is a free data retrieval call binding the contract method 0x257a4a04.
//
// Solidity: function PERMISSION_CONTRACT_DEPLOY() view returns(uint256)
func (_TransferImpl *TransferImplCallerSession) PERMISSIONCONTRACTDEPLOY() (*big.Int, error) {
	return _TransferImpl.Contract.PERMISSIONCONTRACTDEPLOY(&_TransferImpl.CallOpts)
}

// PERMISSIONEMISSION is a free data retrieval call binding the contract method 0x4ccd92d1.
//
// Solidity: function PERMISSION_EMISSION() view returns(uint256)
func (_TransferImpl *TransferImplCaller) PERMISSIONEMISSION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TransferImpl.contract.Call(opts, out, "PERMISSION_EMISSION")
	return *ret0, err
}

// PERMISSIONEMISSION is a free data retrieval call binding the contract method 0x4ccd92d1.
//
// Solidity: function PERMISSION_EMISSION() view returns(uint256)
func (_TransferImpl *TransferImplSession) PERMISSIONEMISSION() (*big.Int, error) {
	return _TransferImpl.Contract.PERMISSIONEMISSION(&_TransferImpl.CallOpts)
}

// PERMISSIONEMISSION is a free data retrieval call binding the contract method 0x4ccd92d1.
//
// Solidity: function PERMISSION_EMISSION() view returns(uint256)
func (_TransferImpl *TransferImplCallerSession) PERMISSIONEMISSION() (*big.Int, error) {
	return _TransferImpl.Contract.PERMISSIONEMISSION(&_TransferImpl.CallOpts)
}

// PERMISSIONFEESSTORAGE is a free data retrieval call binding the contract method 0xe837eb87.
//
// Solidity: function PERMISSION_FEES_STORAGE() view returns(uint256)
func (_TransferImpl *TransferImplCaller) PERMISSIONFEESSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TransferImpl.contract.Call(opts, out, "PERMISSION_FEES_STORAGE")
	return *ret0, err
}

// PERMISSIONFEESSTORAGE is a free data retrieval call binding the contract method 0xe837eb87.
//
// Solidity: function PERMISSION_FEES_STORAGE() view returns(uint256)
func (_TransferImpl *TransferImplSession) PERMISSIONFEESSTORAGE() (*big.Int, error) {
	return _TransferImpl.Contract.PERMISSIONFEESSTORAGE(&_TransferImpl.CallOpts)
}

// PERMISSIONFEESSTORAGE is a free data retrieval call binding the contract method 0xe837eb87.
//
// Solidity: function PERMISSION_FEES_STORAGE() view returns(uint256)
func (_TransferImpl *TransferImplCallerSession) PERMISSIONFEESSTORAGE() (*big.Int, error) {
	return _TransferImpl.Contract.PERMISSIONFEESSTORAGE(&_TransferImpl.CallOpts)
}

// PERMISSIONMULTISIGSTORAGE is a free data retrieval call binding the contract method 0xf1bb7aad.
//
// Solidity: function PERMISSION_MULTISIG_STORAGE() view returns(uint256)
func (_TransferImpl *TransferImplCaller) PERMISSIONMULTISIGSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TransferImpl.contract.Call(opts, out, "PERMISSION_MULTISIG_STORAGE")
	return *ret0, err
}

// PERMISSIONMULTISIGSTORAGE is a free data retrieval call binding the contract method 0xf1bb7aad.
//
// Solidity: function PERMISSION_MULTISIG_STORAGE() view returns(uint256)
func (_TransferImpl *TransferImplSession) PERMISSIONMULTISIGSTORAGE() (*big.Int, error) {
	return _TransferImpl.Contract.PERMISSIONMULTISIGSTORAGE(&_TransferImpl.CallOpts)
}

// PERMISSIONMULTISIGSTORAGE is a free data retrieval call binding the contract method 0xf1bb7aad.
//
// Solidity: function PERMISSION_MULTISIG_STORAGE() view returns(uint256)
func (_TransferImpl *TransferImplCallerSession) PERMISSIONMULTISIGSTORAGE() (*big.Int, error) {
	return _TransferImpl.Contract.PERMISSIONMULTISIGSTORAGE(&_TransferImpl.CallOpts)
}

// PERMISSIONROLESSTORAGE is a free data retrieval call binding the contract method 0x12a16032.
//
// Solidity: function PERMISSION_ROLES_STORAGE() view returns(uint256)
func (_TransferImpl *TransferImplCaller) PERMISSIONROLESSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TransferImpl.contract.Call(opts, out, "PERMISSION_ROLES_STORAGE")
	return *ret0, err
}

// PERMISSIONROLESSTORAGE is a free data retrieval call binding the contract method 0x12a16032.
//
// Solidity: function PERMISSION_ROLES_STORAGE() view returns(uint256)
func (_TransferImpl *TransferImplSession) PERMISSIONROLESSTORAGE() (*big.Int, error) {
	return _TransferImpl.Contract.PERMISSIONROLESSTORAGE(&_TransferImpl.CallOpts)
}

// PERMISSIONROLESSTORAGE is a free data retrieval call binding the contract method 0x12a16032.
//
// Solidity: function PERMISSION_ROLES_STORAGE() view returns(uint256)
func (_TransferImpl *TransferImplCallerSession) PERMISSIONROLESSTORAGE() (*big.Int, error) {
	return _TransferImpl.Contract.PERMISSIONROLESSTORAGE(&_TransferImpl.CallOpts)
}

// ROLEACCESSVERIFIER is a free data retrieval call binding the contract method 0xc5405f92.
//
// Solidity: function ROLE_ACCESS_VERIFIER() view returns(uint256)
func (_TransferImpl *TransferImplCaller) ROLEACCESSVERIFIER(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TransferImpl.contract.Call(opts, out, "ROLE_ACCESS_VERIFIER")
	return *ret0, err
}

// ROLEACCESSVERIFIER is a free data retrieval call binding the contract method 0xc5405f92.
//
// Solidity: function ROLE_ACCESS_VERIFIER() view returns(uint256)
func (_TransferImpl *TransferImplSession) ROLEACCESSVERIFIER() (*big.Int, error) {
	return _TransferImpl.Contract.ROLEACCESSVERIFIER(&_TransferImpl.CallOpts)
}

// ROLEACCESSVERIFIER is a free data retrieval call binding the contract method 0xc5405f92.
//
// Solidity: function ROLE_ACCESS_VERIFIER() view returns(uint256)
func (_TransferImpl *TransferImplCallerSession) ROLEACCESSVERIFIER() (*big.Int, error) {
	return _TransferImpl.Contract.ROLEACCESSVERIFIER(&_TransferImpl.CallOpts)
}

// ROLEBITBONISSUEVERIFIER is a free data retrieval call binding the contract method 0xf0a012a6.
//
// Solidity: function ROLE_BITBON_ISSUE_VERIFIER() view returns(uint256)
func (_TransferImpl *TransferImplCaller) ROLEBITBONISSUEVERIFIER(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TransferImpl.contract.Call(opts, out, "ROLE_BITBON_ISSUE_VERIFIER")
	return *ret0, err
}

// ROLEBITBONISSUEVERIFIER is a free data retrieval call binding the contract method 0xf0a012a6.
//
// Solidity: function ROLE_BITBON_ISSUE_VERIFIER() view returns(uint256)
func (_TransferImpl *TransferImplSession) ROLEBITBONISSUEVERIFIER() (*big.Int, error) {
	return _TransferImpl.Contract.ROLEBITBONISSUEVERIFIER(&_TransferImpl.CallOpts)
}

// ROLEBITBONISSUEVERIFIER is a free data retrieval call binding the contract method 0xf0a012a6.
//
// Solidity: function ROLE_BITBON_ISSUE_VERIFIER() view returns(uint256)
func (_TransferImpl *TransferImplCallerSession) ROLEBITBONISSUEVERIFIER() (*big.Int, error) {
	return _TransferImpl.Contract.ROLEBITBONISSUEVERIFIER(&_TransferImpl.CallOpts)
}

// ROLECOMMISSIONVERIFIER is a free data retrieval call binding the contract method 0x8ab3de8c.
//
// Solidity: function ROLE_COMMISSION_VERIFIER() view returns(uint256)
func (_TransferImpl *TransferImplCaller) ROLECOMMISSIONVERIFIER(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TransferImpl.contract.Call(opts, out, "ROLE_COMMISSION_VERIFIER")
	return *ret0, err
}

// ROLECOMMISSIONVERIFIER is a free data retrieval call binding the contract method 0x8ab3de8c.
//
// Solidity: function ROLE_COMMISSION_VERIFIER() view returns(uint256)
func (_TransferImpl *TransferImplSession) ROLECOMMISSIONVERIFIER() (*big.Int, error) {
	return _TransferImpl.Contract.ROLECOMMISSIONVERIFIER(&_TransferImpl.CallOpts)
}

// ROLECOMMISSIONVERIFIER is a free data retrieval call binding the contract method 0x8ab3de8c.
//
// Solidity: function ROLE_COMMISSION_VERIFIER() view returns(uint256)
func (_TransferImpl *TransferImplCallerSession) ROLECOMMISSIONVERIFIER() (*big.Int, error) {
	return _TransferImpl.Contract.ROLECOMMISSIONVERIFIER(&_TransferImpl.CallOpts)
}

// ROLEDEPLOYADMIN is a free data retrieval call binding the contract method 0x667bf6e5.
//
// Solidity: function ROLE_DEPLOY_ADMIN() view returns(uint256)
func (_TransferImpl *TransferImplCaller) ROLEDEPLOYADMIN(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TransferImpl.contract.Call(opts, out, "ROLE_DEPLOY_ADMIN")
	return *ret0, err
}

// ROLEDEPLOYADMIN is a free data retrieval call binding the contract method 0x667bf6e5.
//
// Solidity: function ROLE_DEPLOY_ADMIN() view returns(uint256)
func (_TransferImpl *TransferImplSession) ROLEDEPLOYADMIN() (*big.Int, error) {
	return _TransferImpl.Contract.ROLEDEPLOYADMIN(&_TransferImpl.CallOpts)
}

// ROLEDEPLOYADMIN is a free data retrieval call binding the contract method 0x667bf6e5.
//
// Solidity: function ROLE_DEPLOY_ADMIN() view returns(uint256)
func (_TransferImpl *TransferImplCallerSession) ROLEDEPLOYADMIN() (*big.Int, error) {
	return _TransferImpl.Contract.ROLEDEPLOYADMIN(&_TransferImpl.CallOpts)
}

// ROLEPERMISSIONADMIN is a free data retrieval call binding the contract method 0x1b239996.
//
// Solidity: function ROLE_PERMISSION_ADMIN() view returns(uint256)
func (_TransferImpl *TransferImplCaller) ROLEPERMISSIONADMIN(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TransferImpl.contract.Call(opts, out, "ROLE_PERMISSION_ADMIN")
	return *ret0, err
}

// ROLEPERMISSIONADMIN is a free data retrieval call binding the contract method 0x1b239996.
//
// Solidity: function ROLE_PERMISSION_ADMIN() view returns(uint256)
func (_TransferImpl *TransferImplSession) ROLEPERMISSIONADMIN() (*big.Int, error) {
	return _TransferImpl.Contract.ROLEPERMISSIONADMIN(&_TransferImpl.CallOpts)
}

// ROLEPERMISSIONADMIN is a free data retrieval call binding the contract method 0x1b239996.
//
// Solidity: function ROLE_PERMISSION_ADMIN() view returns(uint256)
func (_TransferImpl *TransferImplCallerSession) ROLEPERMISSIONADMIN() (*big.Int, error) {
	return _TransferImpl.Contract.ROLEPERMISSIONADMIN(&_TransferImpl.CallOpts)
}

// GetThisContractIndex is a free data retrieval call binding the contract method 0x70c6d367.
//
// Solidity: function getThisContractIndex() pure returns(uint256)
func (_TransferImpl *TransferImplCaller) GetThisContractIndex(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TransferImpl.contract.Call(opts, out, "getThisContractIndex")
	return *ret0, err
}

// GetThisContractIndex is a free data retrieval call binding the contract method 0x70c6d367.
//
// Solidity: function getThisContractIndex() pure returns(uint256)
func (_TransferImpl *TransferImplSession) GetThisContractIndex() (*big.Int, error) {
	return _TransferImpl.Contract.GetThisContractIndex(&_TransferImpl.CallOpts)
}

// GetThisContractIndex is a free data retrieval call binding the contract method 0x70c6d367.
//
// Solidity: function getThisContractIndex() pure returns(uint256)
func (_TransferImpl *TransferImplCallerSession) GetThisContractIndex() (*big.Int, error) {
	return _TransferImpl.Contract.GetThisContractIndex(&_TransferImpl.CallOpts)
}

// TransferExists is a free data retrieval call binding the contract method 0x168f3e26.
//
// Solidity: function transferExists(bytes transferId) view returns(bool)
func (_TransferImpl *TransferImplCaller) TransferExists(opts *bind.CallOpts, transferId []byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TransferImpl.contract.Call(opts, out, "transferExists", transferId)
	return *ret0, err
}

// TransferExists is a free data retrieval call binding the contract method 0x168f3e26.
//
// Solidity: function transferExists(bytes transferId) view returns(bool)
func (_TransferImpl *TransferImplSession) TransferExists(transferId []byte) (bool, error) {
	return _TransferImpl.Contract.TransferExists(&_TransferImpl.CallOpts, transferId)
}

// TransferExists is a free data retrieval call binding the contract method 0x168f3e26.
//
// Solidity: function transferExists(bytes transferId) view returns(bool)
func (_TransferImpl *TransferImplCallerSession) TransferExists(transferId []byte) (bool, error) {
	return _TransferImpl.Contract.TransferExists(&_TransferImpl.CallOpts, transferId)
}

// ApproveSafeTransfer is a paid mutator transaction binding the contract method 0xd1408c8d.
//
// Solidity: function approveSafeTransfer(bytes transferId, bytes protectionCode) returns(bool)
func (_TransferImpl *TransferImplTransactor) ApproveSafeTransfer(opts *bind.TransactOpts, transferId []byte, protectionCode []byte) (*types.Transaction, error) {
	return _TransferImpl.contract.Transact(opts, "approveSafeTransfer", transferId, protectionCode)
}

// ApproveSafeTransfer is a paid mutator transaction binding the contract method 0xd1408c8d.
//
// Solidity: function approveSafeTransfer(bytes transferId, bytes protectionCode) returns(bool)
func (_TransferImpl *TransferImplSession) ApproveSafeTransfer(transferId []byte, protectionCode []byte) (*types.Transaction, error) {
	return _TransferImpl.Contract.ApproveSafeTransfer(&_TransferImpl.TransactOpts, transferId, protectionCode)
}

// ApproveSafeTransfer is a paid mutator transaction binding the contract method 0xd1408c8d.
//
// Solidity: function approveSafeTransfer(bytes transferId, bytes protectionCode) returns(bool)
func (_TransferImpl *TransferImplTransactorSession) ApproveSafeTransfer(transferId []byte, protectionCode []byte) (*types.Transaction, error) {
	return _TransferImpl.Contract.ApproveSafeTransfer(&_TransferImpl.TransactOpts, transferId, protectionCode)
}

// CancelSafeTransfer is a paid mutator transaction binding the contract method 0x8cac0630.
//
// Solidity: function cancelSafeTransfer(bytes transferId) returns(bool)
func (_TransferImpl *TransferImplTransactor) CancelSafeTransfer(opts *bind.TransactOpts, transferId []byte) (*types.Transaction, error) {
	return _TransferImpl.contract.Transact(opts, "cancelSafeTransfer", transferId)
}

// CancelSafeTransfer is a paid mutator transaction binding the contract method 0x8cac0630.
//
// Solidity: function cancelSafeTransfer(bytes transferId) returns(bool)
func (_TransferImpl *TransferImplSession) CancelSafeTransfer(transferId []byte) (*types.Transaction, error) {
	return _TransferImpl.Contract.CancelSafeTransfer(&_TransferImpl.TransactOpts, transferId)
}

// CancelSafeTransfer is a paid mutator transaction binding the contract method 0x8cac0630.
//
// Solidity: function cancelSafeTransfer(bytes transferId) returns(bool)
func (_TransferImpl *TransferImplTransactorSession) CancelSafeTransfer(transferId []byte) (*types.Transaction, error) {
	return _TransferImpl.Contract.CancelSafeTransfer(&_TransferImpl.TransactOpts, transferId)
}

// CreateSafeTransfer is a paid mutator transaction binding the contract method 0x61c779f6.
//
// Solidity: function createSafeTransfer(address from, address to, uint256 value, uint256 fee, bytes proof, bytes vk, bytes transferId, uint256 retries, uint256 expiresAt) returns(bool)
func (_TransferImpl *TransferImplTransactor) CreateSafeTransfer(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int, fee *big.Int, proof []byte, vk []byte, transferId []byte, retries *big.Int, expiresAt *big.Int) (*types.Transaction, error) {
	return _TransferImpl.contract.Transact(opts, "createSafeTransfer", from, to, value, fee, proof, vk, transferId, retries, expiresAt)
}

// CreateSafeTransfer is a paid mutator transaction binding the contract method 0x61c779f6.
//
// Solidity: function createSafeTransfer(address from, address to, uint256 value, uint256 fee, bytes proof, bytes vk, bytes transferId, uint256 retries, uint256 expiresAt) returns(bool)
func (_TransferImpl *TransferImplSession) CreateSafeTransfer(from common.Address, to common.Address, value *big.Int, fee *big.Int, proof []byte, vk []byte, transferId []byte, retries *big.Int, expiresAt *big.Int) (*types.Transaction, error) {
	return _TransferImpl.Contract.CreateSafeTransfer(&_TransferImpl.TransactOpts, from, to, value, fee, proof, vk, transferId, retries, expiresAt)
}

// CreateSafeTransfer is a paid mutator transaction binding the contract method 0x61c779f6.
//
// Solidity: function createSafeTransfer(address from, address to, uint256 value, uint256 fee, bytes proof, bytes vk, bytes transferId, uint256 retries, uint256 expiresAt) returns(bool)
func (_TransferImpl *TransferImplTransactorSession) CreateSafeTransfer(from common.Address, to common.Address, value *big.Int, fee *big.Int, proof []byte, vk []byte, transferId []byte, retries *big.Int, expiresAt *big.Int) (*types.Transaction, error) {
	return _TransferImpl.Contract.CreateSafeTransfer(&_TransferImpl.TransactOpts, from, to, value, fee, proof, vk, transferId, retries, expiresAt)
}

// DirectTransfer is a paid mutator transaction binding the contract method 0xf83d1791.
//
// Solidity: function directTransfer(address from, address to, uint256 value) returns(bool)
func (_TransferImpl *TransferImplTransactor) DirectTransfer(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _TransferImpl.contract.Transact(opts, "directTransfer", from, to, value)
}

// DirectTransfer is a paid mutator transaction binding the contract method 0xf83d1791.
//
// Solidity: function directTransfer(address from, address to, uint256 value) returns(bool)
func (_TransferImpl *TransferImplSession) DirectTransfer(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _TransferImpl.Contract.DirectTransfer(&_TransferImpl.TransactOpts, from, to, value)
}

// DirectTransfer is a paid mutator transaction binding the contract method 0xf83d1791.
//
// Solidity: function directTransfer(address from, address to, uint256 value) returns(bool)
func (_TransferImpl *TransferImplTransactorSession) DirectTransfer(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _TransferImpl.Contract.DirectTransfer(&_TransferImpl.TransactOpts, from, to, value)
}

// ExpireSafeTransfersByHash is a paid mutator transaction binding the contract method 0x8fbc97d4.
//
// Solidity: function expireSafeTransfersByHash(bytes32 transferIdHash) returns(bool)
func (_TransferImpl *TransferImplTransactor) ExpireSafeTransfersByHash(opts *bind.TransactOpts, transferIdHash [32]byte) (*types.Transaction, error) {
	return _TransferImpl.contract.Transact(opts, "expireSafeTransfersByHash", transferIdHash)
}

// ExpireSafeTransfersByHash is a paid mutator transaction binding the contract method 0x8fbc97d4.
//
// Solidity: function expireSafeTransfersByHash(bytes32 transferIdHash) returns(bool)
func (_TransferImpl *TransferImplSession) ExpireSafeTransfersByHash(transferIdHash [32]byte) (*types.Transaction, error) {
	return _TransferImpl.Contract.ExpireSafeTransfersByHash(&_TransferImpl.TransactOpts, transferIdHash)
}

// ExpireSafeTransfersByHash is a paid mutator transaction binding the contract method 0x8fbc97d4.
//
// Solidity: function expireSafeTransfersByHash(bytes32 transferIdHash) returns(bool)
func (_TransferImpl *TransferImplTransactorSession) ExpireSafeTransfersByHash(transferIdHash [32]byte) (*types.Transaction, error) {
	return _TransferImpl.Contract.ExpireSafeTransfersByHash(&_TransferImpl.TransactOpts, transferIdHash)
}

// FeeTransfer is a paid mutator transaction binding the contract method 0x2403e3ee.
//
// Solidity: function feeTransfer(address from, uint256 opType) returns(bool)
func (_TransferImpl *TransferImplTransactor) FeeTransfer(opts *bind.TransactOpts, from common.Address, opType *big.Int) (*types.Transaction, error) {
	return _TransferImpl.contract.Transact(opts, "feeTransfer", from, opType)
}

// FeeTransfer is a paid mutator transaction binding the contract method 0x2403e3ee.
//
// Solidity: function feeTransfer(address from, uint256 opType) returns(bool)
func (_TransferImpl *TransferImplSession) FeeTransfer(from common.Address, opType *big.Int) (*types.Transaction, error) {
	return _TransferImpl.Contract.FeeTransfer(&_TransferImpl.TransactOpts, from, opType)
}

// FeeTransfer is a paid mutator transaction binding the contract method 0x2403e3ee.
//
// Solidity: function feeTransfer(address from, uint256 opType) returns(bool)
func (_TransferImpl *TransferImplTransactorSession) FeeTransfer(from common.Address, opType *big.Int) (*types.Transaction, error) {
	return _TransferImpl.Contract.FeeTransfer(&_TransferImpl.TransactOpts, from, opType)
}

// FrameTransfer is a paid mutator transaction binding the contract method 0xb0092c89.
//
// Solidity: function frameTransfer(address from, address to, uint256 value, uint256 fee) returns(bool)
func (_TransferImpl *TransferImplTransactor) FrameTransfer(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int, fee *big.Int) (*types.Transaction, error) {
	return _TransferImpl.contract.Transact(opts, "frameTransfer", from, to, value, fee)
}

// FrameTransfer is a paid mutator transaction binding the contract method 0xb0092c89.
//
// Solidity: function frameTransfer(address from, address to, uint256 value, uint256 fee) returns(bool)
func (_TransferImpl *TransferImplSession) FrameTransfer(from common.Address, to common.Address, value *big.Int, fee *big.Int) (*types.Transaction, error) {
	return _TransferImpl.Contract.FrameTransfer(&_TransferImpl.TransactOpts, from, to, value, fee)
}

// FrameTransfer is a paid mutator transaction binding the contract method 0xb0092c89.
//
// Solidity: function frameTransfer(address from, address to, uint256 value, uint256 fee) returns(bool)
func (_TransferImpl *TransferImplTransactorSession) FrameTransfer(from common.Address, to common.Address, value *big.Int, fee *big.Int) (*types.Transaction, error) {
	return _TransferImpl.Contract.FrameTransfer(&_TransferImpl.TransactOpts, from, to, value, fee)
}

// QuickTransfer is a paid mutator transaction binding the contract method 0x671d0455.
//
// Solidity: function quickTransfer(address from, address to, uint256 value, uint256 fee) returns(bool)
func (_TransferImpl *TransferImplTransactor) QuickTransfer(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int, fee *big.Int) (*types.Transaction, error) {
	return _TransferImpl.contract.Transact(opts, "quickTransfer", from, to, value, fee)
}

// QuickTransfer is a paid mutator transaction binding the contract method 0x671d0455.
//
// Solidity: function quickTransfer(address from, address to, uint256 value, uint256 fee) returns(bool)
func (_TransferImpl *TransferImplSession) QuickTransfer(from common.Address, to common.Address, value *big.Int, fee *big.Int) (*types.Transaction, error) {
	return _TransferImpl.Contract.QuickTransfer(&_TransferImpl.TransactOpts, from, to, value, fee)
}

// QuickTransfer is a paid mutator transaction binding the contract method 0x671d0455.
//
// Solidity: function quickTransfer(address from, address to, uint256 value, uint256 fee) returns(bool)
func (_TransferImpl *TransferImplTransactorSession) QuickTransfer(from common.Address, to common.Address, value *big.Int, fee *big.Int) (*types.Transaction, error) {
	return _TransferImpl.Contract.QuickTransfer(&_TransferImpl.TransactOpts, from, to, value, fee)
}

// TransferImplBalanceChangedIterator is returned from FilterBalanceChanged and is used to iterate over the raw logs and unpacked data for BalanceChanged events raised by the TransferImpl contract.
type TransferImplBalanceChangedIterator struct {
	Event *TransferImplBalanceChanged // Event containing the contract specifics and raw log

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
func (it *TransferImplBalanceChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TransferImplBalanceChanged)
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
		it.Event = new(TransferImplBalanceChanged)
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
func (it *TransferImplBalanceChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TransferImplBalanceChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TransferImplBalanceChanged represents a BalanceChanged event raised by the TransferImpl contract.
type TransferImplBalanceChanged struct {
	Assetbox     common.Address
	Balance      *big.Int
	AssetboxType *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterBalanceChanged is a free log retrieval operation binding the contract event 0xaeeb0cb16f299136e7e5467ea84217150fe83008833064528f360cde7b7b54c3.
//
// Solidity: event BalanceChanged(address assetbox, uint256 balance, uint256 assetboxType)
func (_TransferImpl *TransferImplFilterer) FilterBalanceChanged(opts *bind.FilterOpts) (*TransferImplBalanceChangedIterator, error) {

	logs, sub, err := _TransferImpl.contract.FilterLogs(opts, "BalanceChanged")
	if err != nil {
		return nil, err
	}
	return &TransferImplBalanceChangedIterator{contract: _TransferImpl.contract, event: "BalanceChanged", logs: logs, sub: sub}, nil
}

// WatchBalanceChanged is a free log subscription operation binding the contract event 0xaeeb0cb16f299136e7e5467ea84217150fe83008833064528f360cde7b7b54c3.
//
// Solidity: event BalanceChanged(address assetbox, uint256 balance, uint256 assetboxType)
func (_TransferImpl *TransferImplFilterer) WatchBalanceChanged(opts *bind.WatchOpts, sink chan<- *TransferImplBalanceChanged) (event.Subscription, error) {

	logs, sub, err := _TransferImpl.contract.WatchLogs(opts, "BalanceChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TransferImplBalanceChanged)
				if err := _TransferImpl.contract.UnpackLog(event, "BalanceChanged", log); err != nil {
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

// ParseBalanceChanged is a log parse operation binding the contract event 0xaeeb0cb16f299136e7e5467ea84217150fe83008833064528f360cde7b7b54c3.
//
// Solidity: event BalanceChanged(address assetbox, uint256 balance, uint256 assetboxType)
func (_TransferImpl *TransferImplFilterer) ParseBalanceChanged(log types.Log) (*TransferImplBalanceChanged, error) {
	event := new(TransferImplBalanceChanged)
	if err := _TransferImpl.contract.UnpackLog(event, "BalanceChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TransferImplBalanceLockedIterator is returned from FilterBalanceLocked and is used to iterate over the raw logs and unpacked data for BalanceLocked events raised by the TransferImpl contract.
type TransferImplBalanceLockedIterator struct {
	Event *TransferImplBalanceLocked // Event containing the contract specifics and raw log

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
func (it *TransferImplBalanceLockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TransferImplBalanceLocked)
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
		it.Event = new(TransferImplBalanceLocked)
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
func (it *TransferImplBalanceLockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TransferImplBalanceLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TransferImplBalanceLocked represents a BalanceLocked event raised by the TransferImpl contract.
type TransferImplBalanceLocked struct {
	Assetbox       common.Address
	BalanceAviable *big.Int
	BalanceLocked  *big.Int
	AssetboxType   *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBalanceLocked is a free log retrieval operation binding the contract event 0x89f85a4bd38f70943757e43dedd843409e565220cb52ba80fc297d1246b3b9bb.
//
// Solidity: event BalanceLocked(address assetbox, uint256 balanceAviable, uint256 balanceLocked, uint256 assetboxType)
func (_TransferImpl *TransferImplFilterer) FilterBalanceLocked(opts *bind.FilterOpts) (*TransferImplBalanceLockedIterator, error) {

	logs, sub, err := _TransferImpl.contract.FilterLogs(opts, "BalanceLocked")
	if err != nil {
		return nil, err
	}
	return &TransferImplBalanceLockedIterator{contract: _TransferImpl.contract, event: "BalanceLocked", logs: logs, sub: sub}, nil
}

// WatchBalanceLocked is a free log subscription operation binding the contract event 0x89f85a4bd38f70943757e43dedd843409e565220cb52ba80fc297d1246b3b9bb.
//
// Solidity: event BalanceLocked(address assetbox, uint256 balanceAviable, uint256 balanceLocked, uint256 assetboxType)
func (_TransferImpl *TransferImplFilterer) WatchBalanceLocked(opts *bind.WatchOpts, sink chan<- *TransferImplBalanceLocked) (event.Subscription, error) {

	logs, sub, err := _TransferImpl.contract.WatchLogs(opts, "BalanceLocked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TransferImplBalanceLocked)
				if err := _TransferImpl.contract.UnpackLog(event, "BalanceLocked", log); err != nil {
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
func (_TransferImpl *TransferImplFilterer) ParseBalanceLocked(log types.Log) (*TransferImplBalanceLocked, error) {
	event := new(TransferImplBalanceLocked)
	if err := _TransferImpl.contract.UnpackLog(event, "BalanceLocked", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TransferImplBalanceUnLockedIterator is returned from FilterBalanceUnLocked and is used to iterate over the raw logs and unpacked data for BalanceUnLocked events raised by the TransferImpl contract.
type TransferImplBalanceUnLockedIterator struct {
	Event *TransferImplBalanceUnLocked // Event containing the contract specifics and raw log

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
func (it *TransferImplBalanceUnLockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TransferImplBalanceUnLocked)
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
		it.Event = new(TransferImplBalanceUnLocked)
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
func (it *TransferImplBalanceUnLockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TransferImplBalanceUnLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TransferImplBalanceUnLocked represents a BalanceUnLocked event raised by the TransferImpl contract.
type TransferImplBalanceUnLocked struct {
	Assetbox       common.Address
	BalanceAviable *big.Int
	BalanceLocked  *big.Int
	AssetboxType   *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBalanceUnLocked is a free log retrieval operation binding the contract event 0x49489f24cf65bd74f17eeccd070da8bb29c8a9eff5b3106025bae2a90ab67ffd.
//
// Solidity: event BalanceUnLocked(address assetbox, uint256 balanceAviable, uint256 balanceLocked, uint256 assetboxType)
func (_TransferImpl *TransferImplFilterer) FilterBalanceUnLocked(opts *bind.FilterOpts) (*TransferImplBalanceUnLockedIterator, error) {

	logs, sub, err := _TransferImpl.contract.FilterLogs(opts, "BalanceUnLocked")
	if err != nil {
		return nil, err
	}
	return &TransferImplBalanceUnLockedIterator{contract: _TransferImpl.contract, event: "BalanceUnLocked", logs: logs, sub: sub}, nil
}

// WatchBalanceUnLocked is a free log subscription operation binding the contract event 0x49489f24cf65bd74f17eeccd070da8bb29c8a9eff5b3106025bae2a90ab67ffd.
//
// Solidity: event BalanceUnLocked(address assetbox, uint256 balanceAviable, uint256 balanceLocked, uint256 assetboxType)
func (_TransferImpl *TransferImplFilterer) WatchBalanceUnLocked(opts *bind.WatchOpts, sink chan<- *TransferImplBalanceUnLocked) (event.Subscription, error) {

	logs, sub, err := _TransferImpl.contract.WatchLogs(opts, "BalanceUnLocked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TransferImplBalanceUnLocked)
				if err := _TransferImpl.contract.UnpackLog(event, "BalanceUnLocked", log); err != nil {
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

// ParseBalanceUnLocked is a log parse operation binding the contract event 0x49489f24cf65bd74f17eeccd070da8bb29c8a9eff5b3106025bae2a90ab67ffd.
//
// Solidity: event BalanceUnLocked(address assetbox, uint256 balanceAviable, uint256 balanceLocked, uint256 assetboxType)
func (_TransferImpl *TransferImplFilterer) ParseBalanceUnLocked(log types.Log) (*TransferImplBalanceUnLocked, error) {
	event := new(TransferImplBalanceUnLocked)
	if err := _TransferImpl.contract.UnpackLog(event, "BalanceUnLocked", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TransferImplDirectTransferCompletedIterator is returned from FilterDirectTransferCompleted and is used to iterate over the raw logs and unpacked data for DirectTransferCompleted events raised by the TransferImpl contract.
type TransferImplDirectTransferCompletedIterator struct {
	Event *TransferImplDirectTransferCompleted // Event containing the contract specifics and raw log

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
func (it *TransferImplDirectTransferCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TransferImplDirectTransferCompleted)
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
		it.Event = new(TransferImplDirectTransferCompleted)
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
func (it *TransferImplDirectTransferCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TransferImplDirectTransferCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TransferImplDirectTransferCompleted represents a DirectTransferCompleted event raised by the TransferImpl contract.
type TransferImplDirectTransferCompleted struct {
	Source common.Address
	Dest   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDirectTransferCompleted is a free log retrieval operation binding the contract event 0x50c303ddb75f64db23b6fa6377525d8dc6110226bf12a063750206f1f2537f7f.
//
// Solidity: event DirectTransferCompleted(address source, address dest, uint256 amount)
func (_TransferImpl *TransferImplFilterer) FilterDirectTransferCompleted(opts *bind.FilterOpts) (*TransferImplDirectTransferCompletedIterator, error) {

	logs, sub, err := _TransferImpl.contract.FilterLogs(opts, "DirectTransferCompleted")
	if err != nil {
		return nil, err
	}
	return &TransferImplDirectTransferCompletedIterator{contract: _TransferImpl.contract, event: "DirectTransferCompleted", logs: logs, sub: sub}, nil
}

// WatchDirectTransferCompleted is a free log subscription operation binding the contract event 0x50c303ddb75f64db23b6fa6377525d8dc6110226bf12a063750206f1f2537f7f.
//
// Solidity: event DirectTransferCompleted(address source, address dest, uint256 amount)
func (_TransferImpl *TransferImplFilterer) WatchDirectTransferCompleted(opts *bind.WatchOpts, sink chan<- *TransferImplDirectTransferCompleted) (event.Subscription, error) {

	logs, sub, err := _TransferImpl.contract.WatchLogs(opts, "DirectTransferCompleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TransferImplDirectTransferCompleted)
				if err := _TransferImpl.contract.UnpackLog(event, "DirectTransferCompleted", log); err != nil {
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

// ParseDirectTransferCompleted is a log parse operation binding the contract event 0x50c303ddb75f64db23b6fa6377525d8dc6110226bf12a063750206f1f2537f7f.
//
// Solidity: event DirectTransferCompleted(address source, address dest, uint256 amount)
func (_TransferImpl *TransferImplFilterer) ParseDirectTransferCompleted(log types.Log) (*TransferImplDirectTransferCompleted, error) {
	event := new(TransferImplDirectTransferCompleted)
	if err := _TransferImpl.contract.UnpackLog(event, "DirectTransferCompleted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TransferImplFeeChargedIterator is returned from FilterFeeCharged and is used to iterate over the raw logs and unpacked data for FeeCharged events raised by the TransferImpl contract.
type TransferImplFeeChargedIterator struct {
	Event *TransferImplFeeCharged // Event containing the contract specifics and raw log

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
func (it *TransferImplFeeChargedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TransferImplFeeCharged)
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
		it.Event = new(TransferImplFeeCharged)
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
func (it *TransferImplFeeChargedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TransferImplFeeChargedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TransferImplFeeCharged represents a FeeCharged event raised by the TransferImpl contract.
type TransferImplFeeCharged struct {
	From common.Address
	To   common.Address
	Fee  *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterFeeCharged is a free log retrieval operation binding the contract event 0x945458c62aa39df7a4d87d6c4dbaaab7de5d870c9a1fe40e2b7571d84f158a8d.
//
// Solidity: event FeeCharged(address from, address to, uint256 fee)
func (_TransferImpl *TransferImplFilterer) FilterFeeCharged(opts *bind.FilterOpts) (*TransferImplFeeChargedIterator, error) {

	logs, sub, err := _TransferImpl.contract.FilterLogs(opts, "FeeCharged")
	if err != nil {
		return nil, err
	}
	return &TransferImplFeeChargedIterator{contract: _TransferImpl.contract, event: "FeeCharged", logs: logs, sub: sub}, nil
}

// WatchFeeCharged is a free log subscription operation binding the contract event 0x945458c62aa39df7a4d87d6c4dbaaab7de5d870c9a1fe40e2b7571d84f158a8d.
//
// Solidity: event FeeCharged(address from, address to, uint256 fee)
func (_TransferImpl *TransferImplFilterer) WatchFeeCharged(opts *bind.WatchOpts, sink chan<- *TransferImplFeeCharged) (event.Subscription, error) {

	logs, sub, err := _TransferImpl.contract.WatchLogs(opts, "FeeCharged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TransferImplFeeCharged)
				if err := _TransferImpl.contract.UnpackLog(event, "FeeCharged", log); err != nil {
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

// ParseFeeCharged is a log parse operation binding the contract event 0x945458c62aa39df7a4d87d6c4dbaaab7de5d870c9a1fe40e2b7571d84f158a8d.
//
// Solidity: event FeeCharged(address from, address to, uint256 fee)
func (_TransferImpl *TransferImplFilterer) ParseFeeCharged(log types.Log) (*TransferImplFeeCharged, error) {
	event := new(TransferImplFeeCharged)
	if err := _TransferImpl.contract.UnpackLog(event, "FeeCharged", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TransferImplFrameTransferCompletedIterator is returned from FilterFrameTransferCompleted and is used to iterate over the raw logs and unpacked data for FrameTransferCompleted events raised by the TransferImpl contract.
type TransferImplFrameTransferCompletedIterator struct {
	Event *TransferImplFrameTransferCompleted // Event containing the contract specifics and raw log

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
func (it *TransferImplFrameTransferCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TransferImplFrameTransferCompleted)
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
		it.Event = new(TransferImplFrameTransferCompleted)
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
func (it *TransferImplFrameTransferCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TransferImplFrameTransferCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TransferImplFrameTransferCompleted represents a FrameTransferCompleted event raised by the TransferImpl contract.
type TransferImplFrameTransferCompleted struct {
	Source common.Address
	Dest   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFrameTransferCompleted is a free log retrieval operation binding the contract event 0x43389500c78a5c9c4b70b031455339c185181c630410e2e72baf05171517cb2d.
//
// Solidity: event FrameTransferCompleted(address source, address dest, uint256 amount)
func (_TransferImpl *TransferImplFilterer) FilterFrameTransferCompleted(opts *bind.FilterOpts) (*TransferImplFrameTransferCompletedIterator, error) {

	logs, sub, err := _TransferImpl.contract.FilterLogs(opts, "FrameTransferCompleted")
	if err != nil {
		return nil, err
	}
	return &TransferImplFrameTransferCompletedIterator{contract: _TransferImpl.contract, event: "FrameTransferCompleted", logs: logs, sub: sub}, nil
}

// WatchFrameTransferCompleted is a free log subscription operation binding the contract event 0x43389500c78a5c9c4b70b031455339c185181c630410e2e72baf05171517cb2d.
//
// Solidity: event FrameTransferCompleted(address source, address dest, uint256 amount)
func (_TransferImpl *TransferImplFilterer) WatchFrameTransferCompleted(opts *bind.WatchOpts, sink chan<- *TransferImplFrameTransferCompleted) (event.Subscription, error) {

	logs, sub, err := _TransferImpl.contract.WatchLogs(opts, "FrameTransferCompleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TransferImplFrameTransferCompleted)
				if err := _TransferImpl.contract.UnpackLog(event, "FrameTransferCompleted", log); err != nil {
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

// ParseFrameTransferCompleted is a log parse operation binding the contract event 0x43389500c78a5c9c4b70b031455339c185181c630410e2e72baf05171517cb2d.
//
// Solidity: event FrameTransferCompleted(address source, address dest, uint256 amount)
func (_TransferImpl *TransferImplFilterer) ParseFrameTransferCompleted(log types.Log) (*TransferImplFrameTransferCompleted, error) {
	event := new(TransferImplFrameTransferCompleted)
	if err := _TransferImpl.contract.UnpackLog(event, "FrameTransferCompleted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TransferImplQuickTransferCompletedIterator is returned from FilterQuickTransferCompleted and is used to iterate over the raw logs and unpacked data for QuickTransferCompleted events raised by the TransferImpl contract.
type TransferImplQuickTransferCompletedIterator struct {
	Event *TransferImplQuickTransferCompleted // Event containing the contract specifics and raw log

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
func (it *TransferImplQuickTransferCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TransferImplQuickTransferCompleted)
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
		it.Event = new(TransferImplQuickTransferCompleted)
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
func (it *TransferImplQuickTransferCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TransferImplQuickTransferCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TransferImplQuickTransferCompleted represents a QuickTransferCompleted event raised by the TransferImpl contract.
type TransferImplQuickTransferCompleted struct {
	Source common.Address
	Dest   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterQuickTransferCompleted is a free log retrieval operation binding the contract event 0xecd4ea7d3fcc63a6f799d99d2ac97b86203a229bfe4ae29446867f44dad8fa50.
//
// Solidity: event QuickTransferCompleted(address source, address dest, uint256 amount)
func (_TransferImpl *TransferImplFilterer) FilterQuickTransferCompleted(opts *bind.FilterOpts) (*TransferImplQuickTransferCompletedIterator, error) {

	logs, sub, err := _TransferImpl.contract.FilterLogs(opts, "QuickTransferCompleted")
	if err != nil {
		return nil, err
	}
	return &TransferImplQuickTransferCompletedIterator{contract: _TransferImpl.contract, event: "QuickTransferCompleted", logs: logs, sub: sub}, nil
}

// WatchQuickTransferCompleted is a free log subscription operation binding the contract event 0xecd4ea7d3fcc63a6f799d99d2ac97b86203a229bfe4ae29446867f44dad8fa50.
//
// Solidity: event QuickTransferCompleted(address source, address dest, uint256 amount)
func (_TransferImpl *TransferImplFilterer) WatchQuickTransferCompleted(opts *bind.WatchOpts, sink chan<- *TransferImplQuickTransferCompleted) (event.Subscription, error) {

	logs, sub, err := _TransferImpl.contract.WatchLogs(opts, "QuickTransferCompleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TransferImplQuickTransferCompleted)
				if err := _TransferImpl.contract.UnpackLog(event, "QuickTransferCompleted", log); err != nil {
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

// ParseQuickTransferCompleted is a log parse operation binding the contract event 0xecd4ea7d3fcc63a6f799d99d2ac97b86203a229bfe4ae29446867f44dad8fa50.
//
// Solidity: event QuickTransferCompleted(address source, address dest, uint256 amount)
func (_TransferImpl *TransferImplFilterer) ParseQuickTransferCompleted(log types.Log) (*TransferImplQuickTransferCompleted, error) {
	event := new(TransferImplQuickTransferCompleted)
	if err := _TransferImpl.contract.UnpackLog(event, "QuickTransferCompleted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TransferImplSafeTransferApprovalResultIterator is returned from FilterSafeTransferApprovalResult and is used to iterate over the raw logs and unpacked data for SafeTransferApprovalResult events raised by the TransferImpl contract.
type TransferImplSafeTransferApprovalResultIterator struct {
	Event *TransferImplSafeTransferApprovalResult // Event containing the contract specifics and raw log

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
func (it *TransferImplSafeTransferApprovalResultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TransferImplSafeTransferApprovalResult)
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
		it.Event = new(TransferImplSafeTransferApprovalResult)
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
func (it *TransferImplSafeTransferApprovalResultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TransferImplSafeTransferApprovalResultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TransferImplSafeTransferApprovalResult represents a SafeTransferApprovalResult event raised by the TransferImpl contract.
type TransferImplSafeTransferApprovalResult struct {
	TransferId     []byte
	Status         uint8
	CurrentRetries *big.Int
	MaxRetries     *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSafeTransferApprovalResult is a free log retrieval operation binding the contract event 0x59aa4628288e8c8d57c40755ed76826015a7adc2b1770c8e86d230c58b43fd16.
//
// Solidity: event SafeTransferApprovalResult(bytes transferId, uint8 status, uint256 currentRetries, uint256 maxRetries)
func (_TransferImpl *TransferImplFilterer) FilterSafeTransferApprovalResult(opts *bind.FilterOpts) (*TransferImplSafeTransferApprovalResultIterator, error) {

	logs, sub, err := _TransferImpl.contract.FilterLogs(opts, "SafeTransferApprovalResult")
	if err != nil {
		return nil, err
	}
	return &TransferImplSafeTransferApprovalResultIterator{contract: _TransferImpl.contract, event: "SafeTransferApprovalResult", logs: logs, sub: sub}, nil
}

// WatchSafeTransferApprovalResult is a free log subscription operation binding the contract event 0x59aa4628288e8c8d57c40755ed76826015a7adc2b1770c8e86d230c58b43fd16.
//
// Solidity: event SafeTransferApprovalResult(bytes transferId, uint8 status, uint256 currentRetries, uint256 maxRetries)
func (_TransferImpl *TransferImplFilterer) WatchSafeTransferApprovalResult(opts *bind.WatchOpts, sink chan<- *TransferImplSafeTransferApprovalResult) (event.Subscription, error) {

	logs, sub, err := _TransferImpl.contract.WatchLogs(opts, "SafeTransferApprovalResult")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TransferImplSafeTransferApprovalResult)
				if err := _TransferImpl.contract.UnpackLog(event, "SafeTransferApprovalResult", log); err != nil {
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

// ParseSafeTransferApprovalResult is a log parse operation binding the contract event 0x59aa4628288e8c8d57c40755ed76826015a7adc2b1770c8e86d230c58b43fd16.
//
// Solidity: event SafeTransferApprovalResult(bytes transferId, uint8 status, uint256 currentRetries, uint256 maxRetries)
func (_TransferImpl *TransferImplFilterer) ParseSafeTransferApprovalResult(log types.Log) (*TransferImplSafeTransferApprovalResult, error) {
	event := new(TransferImplSafeTransferApprovalResult)
	if err := _TransferImpl.contract.UnpackLog(event, "SafeTransferApprovalResult", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TransferImplSafeTransferCancelledIterator is returned from FilterSafeTransferCancelled and is used to iterate over the raw logs and unpacked data for SafeTransferCancelled events raised by the TransferImpl contract.
type TransferImplSafeTransferCancelledIterator struct {
	Event *TransferImplSafeTransferCancelled // Event containing the contract specifics and raw log

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
func (it *TransferImplSafeTransferCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TransferImplSafeTransferCancelled)
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
		it.Event = new(TransferImplSafeTransferCancelled)
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
func (it *TransferImplSafeTransferCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TransferImplSafeTransferCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TransferImplSafeTransferCancelled represents a SafeTransferCancelled event raised by the TransferImpl contract.
type TransferImplSafeTransferCancelled struct {
	TransferId []byte
	Canceller  common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSafeTransferCancelled is a free log retrieval operation binding the contract event 0x8f0a9256e163085bafc9d35c91a118a17505595656a8fddfd0a39e40450669b4.
//
// Solidity: event SafeTransferCancelled(bytes transferId, address canceller)
func (_TransferImpl *TransferImplFilterer) FilterSafeTransferCancelled(opts *bind.FilterOpts) (*TransferImplSafeTransferCancelledIterator, error) {

	logs, sub, err := _TransferImpl.contract.FilterLogs(opts, "SafeTransferCancelled")
	if err != nil {
		return nil, err
	}
	return &TransferImplSafeTransferCancelledIterator{contract: _TransferImpl.contract, event: "SafeTransferCancelled", logs: logs, sub: sub}, nil
}

// WatchSafeTransferCancelled is a free log subscription operation binding the contract event 0x8f0a9256e163085bafc9d35c91a118a17505595656a8fddfd0a39e40450669b4.
//
// Solidity: event SafeTransferCancelled(bytes transferId, address canceller)
func (_TransferImpl *TransferImplFilterer) WatchSafeTransferCancelled(opts *bind.WatchOpts, sink chan<- *TransferImplSafeTransferCancelled) (event.Subscription, error) {

	logs, sub, err := _TransferImpl.contract.WatchLogs(opts, "SafeTransferCancelled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TransferImplSafeTransferCancelled)
				if err := _TransferImpl.contract.UnpackLog(event, "SafeTransferCancelled", log); err != nil {
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

// ParseSafeTransferCancelled is a log parse operation binding the contract event 0x8f0a9256e163085bafc9d35c91a118a17505595656a8fddfd0a39e40450669b4.
//
// Solidity: event SafeTransferCancelled(bytes transferId, address canceller)
func (_TransferImpl *TransferImplFilterer) ParseSafeTransferCancelled(log types.Log) (*TransferImplSafeTransferCancelled, error) {
	event := new(TransferImplSafeTransferCancelled)
	if err := _TransferImpl.contract.UnpackLog(event, "SafeTransferCancelled", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TransferImplSafeTransferCreatedIterator is returned from FilterSafeTransferCreated and is used to iterate over the raw logs and unpacked data for SafeTransferCreated events raised by the TransferImpl contract.
type TransferImplSafeTransferCreatedIterator struct {
	Event *TransferImplSafeTransferCreated // Event containing the contract specifics and raw log

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
func (it *TransferImplSafeTransferCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TransferImplSafeTransferCreated)
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
		it.Event = new(TransferImplSafeTransferCreated)
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
func (it *TransferImplSafeTransferCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TransferImplSafeTransferCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TransferImplSafeTransferCreated represents a SafeTransferCreated event raised by the TransferImpl contract.
type TransferImplSafeTransferCreated struct {
	Src        common.Address
	Dest       common.Address
	Value      *big.Int
	TransferId []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSafeTransferCreated is a free log retrieval operation binding the contract event 0x08b767302d6dc832fb9914380dd151cb6d341d3c2483f110a0b44127c0259df8.
//
// Solidity: event SafeTransferCreated(address src, address dest, uint256 value, bytes transferId)
func (_TransferImpl *TransferImplFilterer) FilterSafeTransferCreated(opts *bind.FilterOpts) (*TransferImplSafeTransferCreatedIterator, error) {

	logs, sub, err := _TransferImpl.contract.FilterLogs(opts, "SafeTransferCreated")
	if err != nil {
		return nil, err
	}
	return &TransferImplSafeTransferCreatedIterator{contract: _TransferImpl.contract, event: "SafeTransferCreated", logs: logs, sub: sub}, nil
}

// WatchSafeTransferCreated is a free log subscription operation binding the contract event 0x08b767302d6dc832fb9914380dd151cb6d341d3c2483f110a0b44127c0259df8.
//
// Solidity: event SafeTransferCreated(address src, address dest, uint256 value, bytes transferId)
func (_TransferImpl *TransferImplFilterer) WatchSafeTransferCreated(opts *bind.WatchOpts, sink chan<- *TransferImplSafeTransferCreated) (event.Subscription, error) {

	logs, sub, err := _TransferImpl.contract.WatchLogs(opts, "SafeTransferCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TransferImplSafeTransferCreated)
				if err := _TransferImpl.contract.UnpackLog(event, "SafeTransferCreated", log); err != nil {
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

// ParseSafeTransferCreated is a log parse operation binding the contract event 0x08b767302d6dc832fb9914380dd151cb6d341d3c2483f110a0b44127c0259df8.
//
// Solidity: event SafeTransferCreated(address src, address dest, uint256 value, bytes transferId)
func (_TransferImpl *TransferImplFilterer) ParseSafeTransferCreated(log types.Log) (*TransferImplSafeTransferCreated, error) {
	event := new(TransferImplSafeTransferCreated)
	if err := _TransferImpl.contract.UnpackLog(event, "SafeTransferCreated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TransferImplSafeTransferExpiredIterator is returned from FilterSafeTransferExpired and is used to iterate over the raw logs and unpacked data for SafeTransferExpired events raised by the TransferImpl contract.
type TransferImplSafeTransferExpiredIterator struct {
	Event *TransferImplSafeTransferExpired // Event containing the contract specifics and raw log

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
func (it *TransferImplSafeTransferExpiredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TransferImplSafeTransferExpired)
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
		it.Event = new(TransferImplSafeTransferExpired)
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
func (it *TransferImplSafeTransferExpiredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TransferImplSafeTransferExpiredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TransferImplSafeTransferExpired represents a SafeTransferExpired event raised by the TransferImpl contract.
type TransferImplSafeTransferExpired struct {
	TransferId []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSafeTransferExpired is a free log retrieval operation binding the contract event 0x502dc83449ece50b468bec4310462dbe9e08f7c869fb644e469f3195224aa880.
//
// Solidity: event SafeTransferExpired(bytes transferId)
func (_TransferImpl *TransferImplFilterer) FilterSafeTransferExpired(opts *bind.FilterOpts) (*TransferImplSafeTransferExpiredIterator, error) {

	logs, sub, err := _TransferImpl.contract.FilterLogs(opts, "SafeTransferExpired")
	if err != nil {
		return nil, err
	}
	return &TransferImplSafeTransferExpiredIterator{contract: _TransferImpl.contract, event: "SafeTransferExpired", logs: logs, sub: sub}, nil
}

// WatchSafeTransferExpired is a free log subscription operation binding the contract event 0x502dc83449ece50b468bec4310462dbe9e08f7c869fb644e469f3195224aa880.
//
// Solidity: event SafeTransferExpired(bytes transferId)
func (_TransferImpl *TransferImplFilterer) WatchSafeTransferExpired(opts *bind.WatchOpts, sink chan<- *TransferImplSafeTransferExpired) (event.Subscription, error) {

	logs, sub, err := _TransferImpl.contract.WatchLogs(opts, "SafeTransferExpired")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TransferImplSafeTransferExpired)
				if err := _TransferImpl.contract.UnpackLog(event, "SafeTransferExpired", log); err != nil {
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

// ParseSafeTransferExpired is a log parse operation binding the contract event 0x502dc83449ece50b468bec4310462dbe9e08f7c869fb644e469f3195224aa880.
//
// Solidity: event SafeTransferExpired(bytes transferId)
func (_TransferImpl *TransferImplFilterer) ParseSafeTransferExpired(log types.Log) (*TransferImplSafeTransferExpired, error) {
	event := new(TransferImplSafeTransferExpired)
	if err := _TransferImpl.contract.UnpackLog(event, "SafeTransferExpired", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TransferImplTransferApprovedIterator is returned from FilterTransferApproved and is used to iterate over the raw logs and unpacked data for TransferApproved events raised by the TransferImpl contract.
type TransferImplTransferApprovedIterator struct {
	Event *TransferImplTransferApproved // Event containing the contract specifics and raw log

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
func (it *TransferImplTransferApprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TransferImplTransferApproved)
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
		it.Event = new(TransferImplTransferApproved)
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
func (it *TransferImplTransferApprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TransferImplTransferApprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TransferImplTransferApproved represents a TransferApproved event raised by the TransferImpl contract.
type TransferImplTransferApproved struct {
	TransferId     []byte
	Caller         common.Address
	CurrentRetries *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterTransferApproved is a free log retrieval operation binding the contract event 0x88bdca62e2a463deb25af4f77505c5c6832d6e2ac9fd4e567be3ce2b35cba980.
//
// Solidity: event TransferApproved(bytes transferId, address caller, uint256 currentRetries)
func (_TransferImpl *TransferImplFilterer) FilterTransferApproved(opts *bind.FilterOpts) (*TransferImplTransferApprovedIterator, error) {

	logs, sub, err := _TransferImpl.contract.FilterLogs(opts, "TransferApproved")
	if err != nil {
		return nil, err
	}
	return &TransferImplTransferApprovedIterator{contract: _TransferImpl.contract, event: "TransferApproved", logs: logs, sub: sub}, nil
}

// WatchTransferApproved is a free log subscription operation binding the contract event 0x88bdca62e2a463deb25af4f77505c5c6832d6e2ac9fd4e567be3ce2b35cba980.
//
// Solidity: event TransferApproved(bytes transferId, address caller, uint256 currentRetries)
func (_TransferImpl *TransferImplFilterer) WatchTransferApproved(opts *bind.WatchOpts, sink chan<- *TransferImplTransferApproved) (event.Subscription, error) {

	logs, sub, err := _TransferImpl.contract.WatchLogs(opts, "TransferApproved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TransferImplTransferApproved)
				if err := _TransferImpl.contract.UnpackLog(event, "TransferApproved", log); err != nil {
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

// ParseTransferApproved is a log parse operation binding the contract event 0x88bdca62e2a463deb25af4f77505c5c6832d6e2ac9fd4e567be3ce2b35cba980.
//
// Solidity: event TransferApproved(bytes transferId, address caller, uint256 currentRetries)
func (_TransferImpl *TransferImplFilterer) ParseTransferApproved(log types.Log) (*TransferImplTransferApproved, error) {
	event := new(TransferImplTransferApproved)
	if err := _TransferImpl.contract.UnpackLog(event, "TransferApproved", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TransferImplTransferWrongProtectionCodeIterator is returned from FilterTransferWrongProtectionCode and is used to iterate over the raw logs and unpacked data for TransferWrongProtectionCode events raised by the TransferImpl contract.
type TransferImplTransferWrongProtectionCodeIterator struct {
	Event *TransferImplTransferWrongProtectionCode // Event containing the contract specifics and raw log

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
func (it *TransferImplTransferWrongProtectionCodeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TransferImplTransferWrongProtectionCode)
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
		it.Event = new(TransferImplTransferWrongProtectionCode)
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
func (it *TransferImplTransferWrongProtectionCodeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TransferImplTransferWrongProtectionCodeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TransferImplTransferWrongProtectionCode represents a TransferWrongProtectionCode event raised by the TransferImpl contract.
type TransferImplTransferWrongProtectionCode struct {
	TransferId     []byte
	Caller         common.Address
	CurrentRetries *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterTransferWrongProtectionCode is a free log retrieval operation binding the contract event 0x487d724419de60e113166369bc4672eaf3fd1363d0f4eb6c6c76256a0e80b867.
//
// Solidity: event TransferWrongProtectionCode(bytes transferId, address caller, uint256 currentRetries)
func (_TransferImpl *TransferImplFilterer) FilterTransferWrongProtectionCode(opts *bind.FilterOpts) (*TransferImplTransferWrongProtectionCodeIterator, error) {

	logs, sub, err := _TransferImpl.contract.FilterLogs(opts, "TransferWrongProtectionCode")
	if err != nil {
		return nil, err
	}
	return &TransferImplTransferWrongProtectionCodeIterator{contract: _TransferImpl.contract, event: "TransferWrongProtectionCode", logs: logs, sub: sub}, nil
}

// WatchTransferWrongProtectionCode is a free log subscription operation binding the contract event 0x487d724419de60e113166369bc4672eaf3fd1363d0f4eb6c6c76256a0e80b867.
//
// Solidity: event TransferWrongProtectionCode(bytes transferId, address caller, uint256 currentRetries)
func (_TransferImpl *TransferImplFilterer) WatchTransferWrongProtectionCode(opts *bind.WatchOpts, sink chan<- *TransferImplTransferWrongProtectionCode) (event.Subscription, error) {

	logs, sub, err := _TransferImpl.contract.WatchLogs(opts, "TransferWrongProtectionCode")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TransferImplTransferWrongProtectionCode)
				if err := _TransferImpl.contract.UnpackLog(event, "TransferWrongProtectionCode", log); err != nil {
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

// ParseTransferWrongProtectionCode is a log parse operation binding the contract event 0x487d724419de60e113166369bc4672eaf3fd1363d0f4eb6c6c76256a0e80b867.
//
// Solidity: event TransferWrongProtectionCode(bytes transferId, address caller, uint256 currentRetries)
func (_TransferImpl *TransferImplFilterer) ParseTransferWrongProtectionCode(log types.Log) (*TransferImplTransferWrongProtectionCode, error) {
	event := new(TransferImplTransferWrongProtectionCode)
	if err := _TransferImpl.contract.UnpackLog(event, "TransferWrongProtectionCode", log); err != nil {
		return nil, err
	}
	return event, nil
}
