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

// FeeStorageImplABI is the input ABI used to generate the binding from.
const FeeStorageImplABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractStorageAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"confirmationsRequired\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"ExceptionalAccountsChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"opType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"distributionAccounts\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"distributionAmounts\",\"type\":\"uint256[]\"}],\"name\":\"FeeDistributionChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"opType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"FeeValueChanged\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ACCESS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ACCESS_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ADMINABLE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ADMIN_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ASSETBOX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ASSETBOX_INFO\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ASSETBOX_INFO_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ASSETBOX_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_BALANCE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_BITBON\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_BITBON_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_BITBON_SUPPORT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_BONBALANCE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_BONTRANSFER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_DEX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_DISTRIBUTION_GATE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_DISTRIBUTION_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_EXCHANGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_EXCHANGE_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_FEE_GATE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_FEE_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_GENERATOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MINING_AGENT_GATE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MINING_AGENT_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MSBON\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MSBON_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_ADD_ADDRESS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_ADD_ADMIN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_ADD_DISTRIBUTION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_ADD_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_ADD_ROLE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_CLEAR_FEE_EXCEPTIONS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_DROP_PERMISSION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_EDIT_ADMIN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_EDIT_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_EDIT_ROLE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REMOVE_ADMIN_BY_ADDRESS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REMOVE_ADMIN_BY_INDEX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REMOVE_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REMOVE_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REMOVE_ROLE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REPLENISH_CAPITALIZATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REPLENISH_CERTIFICATE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REPLENISH_EMISSION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_SET_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_SET_CONFIRMATION_RATE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_SET_FEE_DISTRIBUTION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_SET_FEE_EXCEPTIONS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_SET_MAX_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_SET_MIN_AMOUNT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_SET_MIN_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_SET_PERMISSION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_OTC\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_OTC_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_REPLENISH_CAPITALIZATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_RESERVED_ALIAS_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ROLE_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_SAFE_TRANSFER_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_TRANSFER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_TRANSFER_FROM_CAPITALIZATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"FEE_OPERATION_CREATE_SAFE_TRANSFER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"FEE_OPERATION_CREATE_SAFE_TRANSFER_ALL\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"FEE_OPERATION_FRAME_TRANSFER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"FEE_OPERATION_QUICK_TRANSFER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"FEE_OPERATION_QUICK_TRANSFER_ALL\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"FEE_OPERATION_WPC_SAFE_TRANSFER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"FEE_OPERATION_WPC_SAFE_TRANSFER_ALL\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_ACCESS_RESTORATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_ADMIN_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_CONTRACTS_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_CONTRACT_DEPLOY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_EMISSION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_FEES_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_MULTISIG_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_ROLES_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_ACCESS_VERIFIER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_BITBON_ISSUE_VERIFIER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_COMMISSION_VERIFIER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_DEPLOY_ADMIN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_PERMISSION_ADMIN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"accountToIsExceptional\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"confirmations\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"contractStorage\",\"outputs\":[{\"internalType\":\"contractContractStorage\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"fees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"opTypeToFeeDistribution\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"opTypeToFeeDistributionAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"opTypeToFeeValueIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"registeredFeeValues\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getThisContractIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"feeType\",\"type\":\"uint256\"}],\"name\":\"get\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getFeeValueSettings\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiresOn\",\"type\":\"uint256\"}],\"name\":\"setMultisig\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"opType\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"feeDistributionAccounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"feeDistributionAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"expiresOn\",\"type\":\"uint256\"}],\"name\":\"setFeeDistributionMultisig\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"opType\",\"type\":\"uint256\"}],\"name\":\"getFeeDistributionAccounts\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"opType\",\"type\":\"uint256\"}],\"name\":\"getFeeDistributionAmounts\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"expiresOn\",\"type\":\"uint256\"}],\"name\":\"setExceptionAccountsMultisig\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getExceptionAccounts\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getIsExceptional\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expiresOn\",\"type\":\"uint256\"}],\"name\":\"clearExceptionAccountsMultisig\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// FeeStorageImpl is an auto generated Go binding around an Ethereum contract.
type FeeStorageImpl struct {
	FeeStorageImplCaller     // Read-only binding to the contract
	FeeStorageImplTransactor // Write-only binding to the contract
	FeeStorageImplFilterer   // Log filterer for contract events
}

// FeeStorageImplCaller is an auto generated read-only Go binding around an Ethereum contract.
type FeeStorageImplCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeeStorageImplTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FeeStorageImplTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeeStorageImplFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FeeStorageImplFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeeStorageImplSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FeeStorageImplSession struct {
	Contract     *FeeStorageImpl   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FeeStorageImplCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FeeStorageImplCallerSession struct {
	Contract *FeeStorageImplCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// FeeStorageImplTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FeeStorageImplTransactorSession struct {
	Contract     *FeeStorageImplTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// FeeStorageImplRaw is an auto generated low-level Go binding around an Ethereum contract.
type FeeStorageImplRaw struct {
	Contract *FeeStorageImpl // Generic contract binding to access the raw methods on
}

// FeeStorageImplCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FeeStorageImplCallerRaw struct {
	Contract *FeeStorageImplCaller // Generic read-only contract binding to access the raw methods on
}

// FeeStorageImplTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FeeStorageImplTransactorRaw struct {
	Contract *FeeStorageImplTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFeeStorageImpl creates a new instance of FeeStorageImpl, bound to a specific deployed contract.
func NewFeeStorageImpl(address common.Address, backend bind.ContractBackend) (*FeeStorageImpl, error) {
	contract, err := bindFeeStorageImpl(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FeeStorageImpl{FeeStorageImplCaller: FeeStorageImplCaller{contract: contract}, FeeStorageImplTransactor: FeeStorageImplTransactor{contract: contract}, FeeStorageImplFilterer: FeeStorageImplFilterer{contract: contract}}, nil
}

// NewFeeStorageImplCaller creates a new read-only instance of FeeStorageImpl, bound to a specific deployed contract.
func NewFeeStorageImplCaller(address common.Address, caller bind.ContractCaller) (*FeeStorageImplCaller, error) {
	contract, err := bindFeeStorageImpl(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FeeStorageImplCaller{contract: contract}, nil
}

// NewFeeStorageImplTransactor creates a new write-only instance of FeeStorageImpl, bound to a specific deployed contract.
func NewFeeStorageImplTransactor(address common.Address, transactor bind.ContractTransactor) (*FeeStorageImplTransactor, error) {
	contract, err := bindFeeStorageImpl(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FeeStorageImplTransactor{contract: contract}, nil
}

// NewFeeStorageImplFilterer creates a new log filterer instance of FeeStorageImpl, bound to a specific deployed contract.
func NewFeeStorageImplFilterer(address common.Address, filterer bind.ContractFilterer) (*FeeStorageImplFilterer, error) {
	contract, err := bindFeeStorageImpl(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FeeStorageImplFilterer{contract: contract}, nil
}

// bindFeeStorageImpl binds a generic wrapper to an already deployed contract.
func bindFeeStorageImpl(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FeeStorageImplABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FeeStorageImpl *FeeStorageImplRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FeeStorageImpl.Contract.FeeStorageImplCaller.contract.Call(opts, result, method, params...)
}

// Return Abi for FeeStorageImplCaller
func (_FeeStorageImpl *FeeStorageImplCaller) ABI() abi.ABI {
	return _FeeStorageImpl.contract.Abi
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FeeStorageImpl *FeeStorageImplRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeStorageImpl.Contract.FeeStorageImplTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FeeStorageImpl *FeeStorageImplRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FeeStorageImpl.Contract.FeeStorageImplTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FeeStorageImpl *FeeStorageImplCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FeeStorageImpl.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FeeStorageImpl *FeeStorageImplTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeStorageImpl.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FeeStorageImpl *FeeStorageImplTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FeeStorageImpl.Contract.contract.Transact(opts, method, params...)
}

// CONTRACTACCESS is a free data retrieval call binding the contract method 0x0ac62b0c.
//
// Solidity: function CONTRACT_ACCESS() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCaller) CONTRACTACCESS(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeStorageImpl.contract.Call(opts, out, "CONTRACT_ACCESS")
	return *ret0, err
}

// CONTRACTACCESS is a free data retrieval call binding the contract method 0x0ac62b0c.
//
// Solidity: function CONTRACT_ACCESS() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplSession) CONTRACTACCESS() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTACCESS(&_FeeStorageImpl.CallOpts)
}

// CONTRACTACCESS is a free data retrieval call binding the contract method 0x0ac62b0c.
//
// Solidity: function CONTRACT_ACCESS() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCallerSession) CONTRACTACCESS() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTACCESS(&_FeeStorageImpl.CallOpts)
}

// CONTRACTACCESSSTORAGE is a free data retrieval call binding the contract method 0xceac9bd5.
//
// Solidity: function CONTRACT_ACCESS_STORAGE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCaller) CONTRACTACCESSSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeStorageImpl.contract.Call(opts, out, "CONTRACT_ACCESS_STORAGE")
	return *ret0, err
}

// CONTRACTACCESSSTORAGE is a free data retrieval call binding the contract method 0xceac9bd5.
//
// Solidity: function CONTRACT_ACCESS_STORAGE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplSession) CONTRACTACCESSSTORAGE() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTACCESSSTORAGE(&_FeeStorageImpl.CallOpts)
}

// CONTRACTACCESSSTORAGE is a free data retrieval call binding the contract method 0xceac9bd5.
//
// Solidity: function CONTRACT_ACCESS_STORAGE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCallerSession) CONTRACTACCESSSTORAGE() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTACCESSSTORAGE(&_FeeStorageImpl.CallOpts)
}

// CONTRACTADMINABLE is a free data retrieval call binding the contract method 0xbbe3a78f.
//
// Solidity: function CONTRACT_ADMINABLE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCaller) CONTRACTADMINABLE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeStorageImpl.contract.Call(opts, out, "CONTRACT_ADMINABLE")
	return *ret0, err
}

// CONTRACTADMINABLE is a free data retrieval call binding the contract method 0xbbe3a78f.
//
// Solidity: function CONTRACT_ADMINABLE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplSession) CONTRACTADMINABLE() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTADMINABLE(&_FeeStorageImpl.CallOpts)
}

// CONTRACTADMINABLE is a free data retrieval call binding the contract method 0xbbe3a78f.
//
// Solidity: function CONTRACT_ADMINABLE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCallerSession) CONTRACTADMINABLE() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTADMINABLE(&_FeeStorageImpl.CallOpts)
}

// CONTRACTADMINSTORAGE is a free data retrieval call binding the contract method 0xc9aebae9.
//
// Solidity: function CONTRACT_ADMIN_STORAGE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCaller) CONTRACTADMINSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeStorageImpl.contract.Call(opts, out, "CONTRACT_ADMIN_STORAGE")
	return *ret0, err
}

// CONTRACTADMINSTORAGE is a free data retrieval call binding the contract method 0xc9aebae9.
//
// Solidity: function CONTRACT_ADMIN_STORAGE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplSession) CONTRACTADMINSTORAGE() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTADMINSTORAGE(&_FeeStorageImpl.CallOpts)
}

// CONTRACTADMINSTORAGE is a free data retrieval call binding the contract method 0xc9aebae9.
//
// Solidity: function CONTRACT_ADMIN_STORAGE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCallerSession) CONTRACTADMINSTORAGE() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTADMINSTORAGE(&_FeeStorageImpl.CallOpts)
}

// CONTRACTASSETBOX is a free data retrieval call binding the contract method 0x161aa17c.
//
// Solidity: function CONTRACT_ASSETBOX() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCaller) CONTRACTASSETBOX(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeStorageImpl.contract.Call(opts, out, "CONTRACT_ASSETBOX")
	return *ret0, err
}

// CONTRACTASSETBOX is a free data retrieval call binding the contract method 0x161aa17c.
//
// Solidity: function CONTRACT_ASSETBOX() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplSession) CONTRACTASSETBOX() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTASSETBOX(&_FeeStorageImpl.CallOpts)
}

// CONTRACTASSETBOX is a free data retrieval call binding the contract method 0x161aa17c.
//
// Solidity: function CONTRACT_ASSETBOX() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCallerSession) CONTRACTASSETBOX() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTASSETBOX(&_FeeStorageImpl.CallOpts)
}

// CONTRACTASSETBOXINFO is a free data retrieval call binding the contract method 0xb572fd3d.
//
// Solidity: function CONTRACT_ASSETBOX_INFO() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCaller) CONTRACTASSETBOXINFO(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeStorageImpl.contract.Call(opts, out, "CONTRACT_ASSETBOX_INFO")
	return *ret0, err
}

// CONTRACTASSETBOXINFO is a free data retrieval call binding the contract method 0xb572fd3d.
//
// Solidity: function CONTRACT_ASSETBOX_INFO() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplSession) CONTRACTASSETBOXINFO() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTASSETBOXINFO(&_FeeStorageImpl.CallOpts)
}

// CONTRACTASSETBOXINFO is a free data retrieval call binding the contract method 0xb572fd3d.
//
// Solidity: function CONTRACT_ASSETBOX_INFO() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCallerSession) CONTRACTASSETBOXINFO() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTASSETBOXINFO(&_FeeStorageImpl.CallOpts)
}

// CONTRACTASSETBOXINFOSTORAGE is a free data retrieval call binding the contract method 0xf32a115d.
//
// Solidity: function CONTRACT_ASSETBOX_INFO_STORAGE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCaller) CONTRACTASSETBOXINFOSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeStorageImpl.contract.Call(opts, out, "CONTRACT_ASSETBOX_INFO_STORAGE")
	return *ret0, err
}

// CONTRACTASSETBOXINFOSTORAGE is a free data retrieval call binding the contract method 0xf32a115d.
//
// Solidity: function CONTRACT_ASSETBOX_INFO_STORAGE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplSession) CONTRACTASSETBOXINFOSTORAGE() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTASSETBOXINFOSTORAGE(&_FeeStorageImpl.CallOpts)
}

// CONTRACTASSETBOXINFOSTORAGE is a free data retrieval call binding the contract method 0xf32a115d.
//
// Solidity: function CONTRACT_ASSETBOX_INFO_STORAGE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCallerSession) CONTRACTASSETBOXINFOSTORAGE() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTASSETBOXINFOSTORAGE(&_FeeStorageImpl.CallOpts)
}

// CONTRACTASSETBOXSTORAGE is a free data retrieval call binding the contract method 0xfa5bdafc.
//
// Solidity: function CONTRACT_ASSETBOX_STORAGE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCaller) CONTRACTASSETBOXSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeStorageImpl.contract.Call(opts, out, "CONTRACT_ASSETBOX_STORAGE")
	return *ret0, err
}

// CONTRACTASSETBOXSTORAGE is a free data retrieval call binding the contract method 0xfa5bdafc.
//
// Solidity: function CONTRACT_ASSETBOX_STORAGE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplSession) CONTRACTASSETBOXSTORAGE() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTASSETBOXSTORAGE(&_FeeStorageImpl.CallOpts)
}

// CONTRACTASSETBOXSTORAGE is a free data retrieval call binding the contract method 0xfa5bdafc.
//
// Solidity: function CONTRACT_ASSETBOX_STORAGE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCallerSession) CONTRACTASSETBOXSTORAGE() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTASSETBOXSTORAGE(&_FeeStorageImpl.CallOpts)
}

// CONTRACTBALANCE is a free data retrieval call binding the contract method 0xd59045f5.
//
// Solidity: function CONTRACT_BALANCE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCaller) CONTRACTBALANCE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeStorageImpl.contract.Call(opts, out, "CONTRACT_BALANCE")
	return *ret0, err
}

// CONTRACTBALANCE is a free data retrieval call binding the contract method 0xd59045f5.
//
// Solidity: function CONTRACT_BALANCE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplSession) CONTRACTBALANCE() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTBALANCE(&_FeeStorageImpl.CallOpts)
}

// CONTRACTBALANCE is a free data retrieval call binding the contract method 0xd59045f5.
//
// Solidity: function CONTRACT_BALANCE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCallerSession) CONTRACTBALANCE() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTBALANCE(&_FeeStorageImpl.CallOpts)
}

// CONTRACTBITBON is a free data retrieval call binding the contract method 0xb6b433ab.
//
// Solidity: function CONTRACT_BITBON() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCaller) CONTRACTBITBON(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeStorageImpl.contract.Call(opts, out, "CONTRACT_BITBON")
	return *ret0, err
}

// CONTRACTBITBON is a free data retrieval call binding the contract method 0xb6b433ab.
//
// Solidity: function CONTRACT_BITBON() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplSession) CONTRACTBITBON() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTBITBON(&_FeeStorageImpl.CallOpts)
}

// CONTRACTBITBON is a free data retrieval call binding the contract method 0xb6b433ab.
//
// Solidity: function CONTRACT_BITBON() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCallerSession) CONTRACTBITBON() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTBITBON(&_FeeStorageImpl.CallOpts)
}

// CONTRACTBITBONSTORAGE is a free data retrieval call binding the contract method 0xe7f6c30e.
//
// Solidity: function CONTRACT_BITBON_STORAGE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCaller) CONTRACTBITBONSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeStorageImpl.contract.Call(opts, out, "CONTRACT_BITBON_STORAGE")
	return *ret0, err
}

// CONTRACTBITBONSTORAGE is a free data retrieval call binding the contract method 0xe7f6c30e.
//
// Solidity: function CONTRACT_BITBON_STORAGE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplSession) CONTRACTBITBONSTORAGE() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTBITBONSTORAGE(&_FeeStorageImpl.CallOpts)
}

// CONTRACTBITBONSTORAGE is a free data retrieval call binding the contract method 0xe7f6c30e.
//
// Solidity: function CONTRACT_BITBON_STORAGE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCallerSession) CONTRACTBITBONSTORAGE() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTBITBONSTORAGE(&_FeeStorageImpl.CallOpts)
}

// CONTRACTBITBONSUPPORT is a free data retrieval call binding the contract method 0x20139c74.
//
// Solidity: function CONTRACT_BITBON_SUPPORT() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCaller) CONTRACTBITBONSUPPORT(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeStorageImpl.contract.Call(opts, out, "CONTRACT_BITBON_SUPPORT")
	return *ret0, err
}

// CONTRACTBITBONSUPPORT is a free data retrieval call binding the contract method 0x20139c74.
//
// Solidity: function CONTRACT_BITBON_SUPPORT() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplSession) CONTRACTBITBONSUPPORT() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTBITBONSUPPORT(&_FeeStorageImpl.CallOpts)
}

// CONTRACTBITBONSUPPORT is a free data retrieval call binding the contract method 0x20139c74.
//
// Solidity: function CONTRACT_BITBON_SUPPORT() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCallerSession) CONTRACTBITBONSUPPORT() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTBITBONSUPPORT(&_FeeStorageImpl.CallOpts)
}

// CONTRACTBONBALANCE is a free data retrieval call binding the contract method 0x40e02b7c.
//
// Solidity: function CONTRACT_BONBALANCE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCaller) CONTRACTBONBALANCE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeStorageImpl.contract.Call(opts, out, "CONTRACT_BONBALANCE")
	return *ret0, err
}

// CONTRACTBONBALANCE is a free data retrieval call binding the contract method 0x40e02b7c.
//
// Solidity: function CONTRACT_BONBALANCE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplSession) CONTRACTBONBALANCE() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTBONBALANCE(&_FeeStorageImpl.CallOpts)
}

// CONTRACTBONBALANCE is a free data retrieval call binding the contract method 0x40e02b7c.
//
// Solidity: function CONTRACT_BONBALANCE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCallerSession) CONTRACTBONBALANCE() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTBONBALANCE(&_FeeStorageImpl.CallOpts)
}

// CONTRACTBONTRANSFER is a free data retrieval call binding the contract method 0xc18f4a8a.
//
// Solidity: function CONTRACT_BONTRANSFER() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCaller) CONTRACTBONTRANSFER(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeStorageImpl.contract.Call(opts, out, "CONTRACT_BONTRANSFER")
	return *ret0, err
}

// CONTRACTBONTRANSFER is a free data retrieval call binding the contract method 0xc18f4a8a.
//
// Solidity: function CONTRACT_BONTRANSFER() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplSession) CONTRACTBONTRANSFER() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTBONTRANSFER(&_FeeStorageImpl.CallOpts)
}

// CONTRACTBONTRANSFER is a free data retrieval call binding the contract method 0xc18f4a8a.
//
// Solidity: function CONTRACT_BONTRANSFER() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCallerSession) CONTRACTBONTRANSFER() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTBONTRANSFER(&_FeeStorageImpl.CallOpts)
}

// CONTRACTDEX is a free data retrieval call binding the contract method 0x52c1f844.
//
// Solidity: function CONTRACT_DEX() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCaller) CONTRACTDEX(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeStorageImpl.contract.Call(opts, out, "CONTRACT_DEX")
	return *ret0, err
}

// CONTRACTDEX is a free data retrieval call binding the contract method 0x52c1f844.
//
// Solidity: function CONTRACT_DEX() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplSession) CONTRACTDEX() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTDEX(&_FeeStorageImpl.CallOpts)
}

// CONTRACTDEX is a free data retrieval call binding the contract method 0x52c1f844.
//
// Solidity: function CONTRACT_DEX() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCallerSession) CONTRACTDEX() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTDEX(&_FeeStorageImpl.CallOpts)
}

// CONTRACTDISTRIBUTIONGATE is a free data retrieval call binding the contract method 0x6e993588.
//
// Solidity: function CONTRACT_DISTRIBUTION_GATE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCaller) CONTRACTDISTRIBUTIONGATE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeStorageImpl.contract.Call(opts, out, "CONTRACT_DISTRIBUTION_GATE")
	return *ret0, err
}

// CONTRACTDISTRIBUTIONGATE is a free data retrieval call binding the contract method 0x6e993588.
//
// Solidity: function CONTRACT_DISTRIBUTION_GATE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplSession) CONTRACTDISTRIBUTIONGATE() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTDISTRIBUTIONGATE(&_FeeStorageImpl.CallOpts)
}

// CONTRACTDISTRIBUTIONGATE is a free data retrieval call binding the contract method 0x6e993588.
//
// Solidity: function CONTRACT_DISTRIBUTION_GATE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCallerSession) CONTRACTDISTRIBUTIONGATE() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTDISTRIBUTIONGATE(&_FeeStorageImpl.CallOpts)
}

// CONTRACTDISTRIBUTIONSTORAGE is a free data retrieval call binding the contract method 0x08097327.
//
// Solidity: function CONTRACT_DISTRIBUTION_STORAGE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCaller) CONTRACTDISTRIBUTIONSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeStorageImpl.contract.Call(opts, out, "CONTRACT_DISTRIBUTION_STORAGE")
	return *ret0, err
}

// CONTRACTDISTRIBUTIONSTORAGE is a free data retrieval call binding the contract method 0x08097327.
//
// Solidity: function CONTRACT_DISTRIBUTION_STORAGE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplSession) CONTRACTDISTRIBUTIONSTORAGE() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTDISTRIBUTIONSTORAGE(&_FeeStorageImpl.CallOpts)
}

// CONTRACTDISTRIBUTIONSTORAGE is a free data retrieval call binding the contract method 0x08097327.
//
// Solidity: function CONTRACT_DISTRIBUTION_STORAGE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCallerSession) CONTRACTDISTRIBUTIONSTORAGE() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTDISTRIBUTIONSTORAGE(&_FeeStorageImpl.CallOpts)
}

// CONTRACTEXCHANGE is a free data retrieval call binding the contract method 0xa0e3a10e.
//
// Solidity: function CONTRACT_EXCHANGE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCaller) CONTRACTEXCHANGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeStorageImpl.contract.Call(opts, out, "CONTRACT_EXCHANGE")
	return *ret0, err
}

// CONTRACTEXCHANGE is a free data retrieval call binding the contract method 0xa0e3a10e.
//
// Solidity: function CONTRACT_EXCHANGE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplSession) CONTRACTEXCHANGE() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTEXCHANGE(&_FeeStorageImpl.CallOpts)
}

// CONTRACTEXCHANGE is a free data retrieval call binding the contract method 0xa0e3a10e.
//
// Solidity: function CONTRACT_EXCHANGE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCallerSession) CONTRACTEXCHANGE() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTEXCHANGE(&_FeeStorageImpl.CallOpts)
}

// CONTRACTEXCHANGESTORAGE is a free data retrieval call binding the contract method 0x001570f2.
//
// Solidity: function CONTRACT_EXCHANGE_STORAGE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCaller) CONTRACTEXCHANGESTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeStorageImpl.contract.Call(opts, out, "CONTRACT_EXCHANGE_STORAGE")
	return *ret0, err
}

// CONTRACTEXCHANGESTORAGE is a free data retrieval call binding the contract method 0x001570f2.
//
// Solidity: function CONTRACT_EXCHANGE_STORAGE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplSession) CONTRACTEXCHANGESTORAGE() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTEXCHANGESTORAGE(&_FeeStorageImpl.CallOpts)
}

// CONTRACTEXCHANGESTORAGE is a free data retrieval call binding the contract method 0x001570f2.
//
// Solidity: function CONTRACT_EXCHANGE_STORAGE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCallerSession) CONTRACTEXCHANGESTORAGE() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTEXCHANGESTORAGE(&_FeeStorageImpl.CallOpts)
}

// CONTRACTFEEGATE is a free data retrieval call binding the contract method 0x545c17a4.
//
// Solidity: function CONTRACT_FEE_GATE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCaller) CONTRACTFEEGATE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeStorageImpl.contract.Call(opts, out, "CONTRACT_FEE_GATE")
	return *ret0, err
}

// CONTRACTFEEGATE is a free data retrieval call binding the contract method 0x545c17a4.
//
// Solidity: function CONTRACT_FEE_GATE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplSession) CONTRACTFEEGATE() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTFEEGATE(&_FeeStorageImpl.CallOpts)
}

// CONTRACTFEEGATE is a free data retrieval call binding the contract method 0x545c17a4.
//
// Solidity: function CONTRACT_FEE_GATE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCallerSession) CONTRACTFEEGATE() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTFEEGATE(&_FeeStorageImpl.CallOpts)
}

// CONTRACTFEESTORAGE is a free data retrieval call binding the contract method 0xcd5e65a4.
//
// Solidity: function CONTRACT_FEE_STORAGE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCaller) CONTRACTFEESTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeStorageImpl.contract.Call(opts, out, "CONTRACT_FEE_STORAGE")
	return *ret0, err
}

// CONTRACTFEESTORAGE is a free data retrieval call binding the contract method 0xcd5e65a4.
//
// Solidity: function CONTRACT_FEE_STORAGE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplSession) CONTRACTFEESTORAGE() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTFEESTORAGE(&_FeeStorageImpl.CallOpts)
}

// CONTRACTFEESTORAGE is a free data retrieval call binding the contract method 0xcd5e65a4.
//
// Solidity: function CONTRACT_FEE_STORAGE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCallerSession) CONTRACTFEESTORAGE() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTFEESTORAGE(&_FeeStorageImpl.CallOpts)
}

// CONTRACTGENERATOR is a free data retrieval call binding the contract method 0xccea2b46.
//
// Solidity: function CONTRACT_GENERATOR() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCaller) CONTRACTGENERATOR(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeStorageImpl.contract.Call(opts, out, "CONTRACT_GENERATOR")
	return *ret0, err
}

// CONTRACTGENERATOR is a free data retrieval call binding the contract method 0xccea2b46.
//
// Solidity: function CONTRACT_GENERATOR() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplSession) CONTRACTGENERATOR() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTGENERATOR(&_FeeStorageImpl.CallOpts)
}

// CONTRACTGENERATOR is a free data retrieval call binding the contract method 0xccea2b46.
//
// Solidity: function CONTRACT_GENERATOR() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCallerSession) CONTRACTGENERATOR() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTGENERATOR(&_FeeStorageImpl.CallOpts)
}

// CONTRACTMININGAGENTGATE is a free data retrieval call binding the contract method 0xb9b4c0f1.
//
// Solidity: function CONTRACT_MINING_AGENT_GATE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCaller) CONTRACTMININGAGENTGATE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeStorageImpl.contract.Call(opts, out, "CONTRACT_MINING_AGENT_GATE")
	return *ret0, err
}

// CONTRACTMININGAGENTGATE is a free data retrieval call binding the contract method 0xb9b4c0f1.
//
// Solidity: function CONTRACT_MINING_AGENT_GATE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplSession) CONTRACTMININGAGENTGATE() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTMININGAGENTGATE(&_FeeStorageImpl.CallOpts)
}

// CONTRACTMININGAGENTGATE is a free data retrieval call binding the contract method 0xb9b4c0f1.
//
// Solidity: function CONTRACT_MINING_AGENT_GATE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCallerSession) CONTRACTMININGAGENTGATE() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTMININGAGENTGATE(&_FeeStorageImpl.CallOpts)
}

// CONTRACTMININGAGENTSTORAGE is a free data retrieval call binding the contract method 0x83cc588d.
//
// Solidity: function CONTRACT_MINING_AGENT_STORAGE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCaller) CONTRACTMININGAGENTSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeStorageImpl.contract.Call(opts, out, "CONTRACT_MINING_AGENT_STORAGE")
	return *ret0, err
}

// CONTRACTMININGAGENTSTORAGE is a free data retrieval call binding the contract method 0x83cc588d.
//
// Solidity: function CONTRACT_MINING_AGENT_STORAGE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplSession) CONTRACTMININGAGENTSTORAGE() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTMININGAGENTSTORAGE(&_FeeStorageImpl.CallOpts)
}

// CONTRACTMININGAGENTSTORAGE is a free data retrieval call binding the contract method 0x83cc588d.
//
// Solidity: function CONTRACT_MINING_AGENT_STORAGE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCallerSession) CONTRACTMININGAGENTSTORAGE() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTMININGAGENTSTORAGE(&_FeeStorageImpl.CallOpts)
}

// CONTRACTMSBON is a free data retrieval call binding the contract method 0x28b3a976.
//
// Solidity: function CONTRACT_MSBON() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCaller) CONTRACTMSBON(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeStorageImpl.contract.Call(opts, out, "CONTRACT_MSBON")
	return *ret0, err
}

// CONTRACTMSBON is a free data retrieval call binding the contract method 0x28b3a976.
//
// Solidity: function CONTRACT_MSBON() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplSession) CONTRACTMSBON() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTMSBON(&_FeeStorageImpl.CallOpts)
}

// CONTRACTMSBON is a free data retrieval call binding the contract method 0x28b3a976.
//
// Solidity: function CONTRACT_MSBON() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCallerSession) CONTRACTMSBON() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTMSBON(&_FeeStorageImpl.CallOpts)
}

// CONTRACTMSBONSTORAGE is a free data retrieval call binding the contract method 0xe04e6061.
//
// Solidity: function CONTRACT_MSBON_STORAGE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCaller) CONTRACTMSBONSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeStorageImpl.contract.Call(opts, out, "CONTRACT_MSBON_STORAGE")
	return *ret0, err
}

// CONTRACTMSBONSTORAGE is a free data retrieval call binding the contract method 0xe04e6061.
//
// Solidity: function CONTRACT_MSBON_STORAGE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplSession) CONTRACTMSBONSTORAGE() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTMSBONSTORAGE(&_FeeStorageImpl.CallOpts)
}

// CONTRACTMSBONSTORAGE is a free data retrieval call binding the contract method 0xe04e6061.
//
// Solidity: function CONTRACT_MSBON_STORAGE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCallerSession) CONTRACTMSBONSTORAGE() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTMSBONSTORAGE(&_FeeStorageImpl.CallOpts)
}

// CONTRACTMULTISIGADDADDRESS is a free data retrieval call binding the contract method 0xf4f65a5a.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ADDRESS() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCaller) CONTRACTMULTISIGADDADDRESS(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_ADD_ADDRESS")
	return *ret0, err
}

// CONTRACTMULTISIGADDADDRESS is a free data retrieval call binding the contract method 0xf4f65a5a.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ADDRESS() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplSession) CONTRACTMULTISIGADDADDRESS() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTMULTISIGADDADDRESS(&_FeeStorageImpl.CallOpts)
}

// CONTRACTMULTISIGADDADDRESS is a free data retrieval call binding the contract method 0xf4f65a5a.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ADDRESS() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCallerSession) CONTRACTMULTISIGADDADDRESS() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTMULTISIGADDADDRESS(&_FeeStorageImpl.CallOpts)
}

// CONTRACTMULTISIGADDADMIN is a free data retrieval call binding the contract method 0x83888c68.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ADMIN() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCaller) CONTRACTMULTISIGADDADMIN(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_ADD_ADMIN")
	return *ret0, err
}

// CONTRACTMULTISIGADDADMIN is a free data retrieval call binding the contract method 0x83888c68.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ADMIN() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplSession) CONTRACTMULTISIGADDADMIN() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTMULTISIGADDADMIN(&_FeeStorageImpl.CallOpts)
}

// CONTRACTMULTISIGADDADMIN is a free data retrieval call binding the contract method 0x83888c68.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ADMIN() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCallerSession) CONTRACTMULTISIGADDADMIN() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTMULTISIGADDADMIN(&_FeeStorageImpl.CallOpts)
}

// CONTRACTMULTISIGADDDISTRIBUTION is a free data retrieval call binding the contract method 0x8f8b81d1.
//
// Solidity: function CONTRACT_MULTISIG_ADD_DISTRIBUTION() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCaller) CONTRACTMULTISIGADDDISTRIBUTION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_ADD_DISTRIBUTION")
	return *ret0, err
}

// CONTRACTMULTISIGADDDISTRIBUTION is a free data retrieval call binding the contract method 0x8f8b81d1.
//
// Solidity: function CONTRACT_MULTISIG_ADD_DISTRIBUTION() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplSession) CONTRACTMULTISIGADDDISTRIBUTION() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTMULTISIGADDDISTRIBUTION(&_FeeStorageImpl.CallOpts)
}

// CONTRACTMULTISIGADDDISTRIBUTION is a free data retrieval call binding the contract method 0x8f8b81d1.
//
// Solidity: function CONTRACT_MULTISIG_ADD_DISTRIBUTION() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCallerSession) CONTRACTMULTISIGADDDISTRIBUTION() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTMULTISIGADDDISTRIBUTION(&_FeeStorageImpl.CallOpts)
}

// CONTRACTMULTISIGADDFEE is a free data retrieval call binding the contract method 0xc43e8cab.
//
// Solidity: function CONTRACT_MULTISIG_ADD_FEE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCaller) CONTRACTMULTISIGADDFEE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_ADD_FEE")
	return *ret0, err
}

// CONTRACTMULTISIGADDFEE is a free data retrieval call binding the contract method 0xc43e8cab.
//
// Solidity: function CONTRACT_MULTISIG_ADD_FEE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplSession) CONTRACTMULTISIGADDFEE() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTMULTISIGADDFEE(&_FeeStorageImpl.CallOpts)
}

// CONTRACTMULTISIGADDFEE is a free data retrieval call binding the contract method 0xc43e8cab.
//
// Solidity: function CONTRACT_MULTISIG_ADD_FEE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCallerSession) CONTRACTMULTISIGADDFEE() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTMULTISIGADDFEE(&_FeeStorageImpl.CallOpts)
}

// CONTRACTMULTISIGADDROLE is a free data retrieval call binding the contract method 0x20a0914b.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ROLE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCaller) CONTRACTMULTISIGADDROLE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_ADD_ROLE")
	return *ret0, err
}

// CONTRACTMULTISIGADDROLE is a free data retrieval call binding the contract method 0x20a0914b.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ROLE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplSession) CONTRACTMULTISIGADDROLE() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTMULTISIGADDROLE(&_FeeStorageImpl.CallOpts)
}

// CONTRACTMULTISIGADDROLE is a free data retrieval call binding the contract method 0x20a0914b.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ROLE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCallerSession) CONTRACTMULTISIGADDROLE() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTMULTISIGADDROLE(&_FeeStorageImpl.CallOpts)
}

// CONTRACTMULTISIGCLEARFEEEXCEPTIONS is a free data retrieval call binding the contract method 0xf74701f2.
//
// Solidity: function CONTRACT_MULTISIG_CLEAR_FEE_EXCEPTIONS() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCaller) CONTRACTMULTISIGCLEARFEEEXCEPTIONS(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_CLEAR_FEE_EXCEPTIONS")
	return *ret0, err
}

// CONTRACTMULTISIGCLEARFEEEXCEPTIONS is a free data retrieval call binding the contract method 0xf74701f2.
//
// Solidity: function CONTRACT_MULTISIG_CLEAR_FEE_EXCEPTIONS() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplSession) CONTRACTMULTISIGCLEARFEEEXCEPTIONS() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTMULTISIGCLEARFEEEXCEPTIONS(&_FeeStorageImpl.CallOpts)
}

// CONTRACTMULTISIGCLEARFEEEXCEPTIONS is a free data retrieval call binding the contract method 0xf74701f2.
//
// Solidity: function CONTRACT_MULTISIG_CLEAR_FEE_EXCEPTIONS() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCallerSession) CONTRACTMULTISIGCLEARFEEEXCEPTIONS() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTMULTISIGCLEARFEEEXCEPTIONS(&_FeeStorageImpl.CallOpts)
}

// CONTRACTMULTISIGDROPPERMISSION is a free data retrieval call binding the contract method 0x3f3fbc39.
//
// Solidity: function CONTRACT_MULTISIG_DROP_PERMISSION() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCaller) CONTRACTMULTISIGDROPPERMISSION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_DROP_PERMISSION")
	return *ret0, err
}

// CONTRACTMULTISIGDROPPERMISSION is a free data retrieval call binding the contract method 0x3f3fbc39.
//
// Solidity: function CONTRACT_MULTISIG_DROP_PERMISSION() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplSession) CONTRACTMULTISIGDROPPERMISSION() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTMULTISIGDROPPERMISSION(&_FeeStorageImpl.CallOpts)
}

// CONTRACTMULTISIGDROPPERMISSION is a free data retrieval call binding the contract method 0x3f3fbc39.
//
// Solidity: function CONTRACT_MULTISIG_DROP_PERMISSION() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCallerSession) CONTRACTMULTISIGDROPPERMISSION() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTMULTISIGDROPPERMISSION(&_FeeStorageImpl.CallOpts)
}

// CONTRACTMULTISIGEDITADMIN is a free data retrieval call binding the contract method 0x91b0f051.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_ADMIN() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCaller) CONTRACTMULTISIGEDITADMIN(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_EDIT_ADMIN")
	return *ret0, err
}

// CONTRACTMULTISIGEDITADMIN is a free data retrieval call binding the contract method 0x91b0f051.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_ADMIN() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplSession) CONTRACTMULTISIGEDITADMIN() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTMULTISIGEDITADMIN(&_FeeStorageImpl.CallOpts)
}

// CONTRACTMULTISIGEDITADMIN is a free data retrieval call binding the contract method 0x91b0f051.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_ADMIN() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCallerSession) CONTRACTMULTISIGEDITADMIN() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTMULTISIGEDITADMIN(&_FeeStorageImpl.CallOpts)
}

// CONTRACTMULTISIGEDITFEE is a free data retrieval call binding the contract method 0x6b2abe17.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_FEE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCaller) CONTRACTMULTISIGEDITFEE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_EDIT_FEE")
	return *ret0, err
}

// CONTRACTMULTISIGEDITFEE is a free data retrieval call binding the contract method 0x6b2abe17.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_FEE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplSession) CONTRACTMULTISIGEDITFEE() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTMULTISIGEDITFEE(&_FeeStorageImpl.CallOpts)
}

// CONTRACTMULTISIGEDITFEE is a free data retrieval call binding the contract method 0x6b2abe17.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_FEE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCallerSession) CONTRACTMULTISIGEDITFEE() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTMULTISIGEDITFEE(&_FeeStorageImpl.CallOpts)
}

// CONTRACTMULTISIGEDITROLE is a free data retrieval call binding the contract method 0xc5f66412.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_ROLE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCaller) CONTRACTMULTISIGEDITROLE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_EDIT_ROLE")
	return *ret0, err
}

// CONTRACTMULTISIGEDITROLE is a free data retrieval call binding the contract method 0xc5f66412.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_ROLE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplSession) CONTRACTMULTISIGEDITROLE() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTMULTISIGEDITROLE(&_FeeStorageImpl.CallOpts)
}

// CONTRACTMULTISIGEDITROLE is a free data retrieval call binding the contract method 0xc5f66412.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_ROLE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCallerSession) CONTRACTMULTISIGEDITROLE() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTMULTISIGEDITROLE(&_FeeStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEADMINBYADDRESS is a free data retrieval call binding the contract method 0xc816b091.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_BY_ADDRESS() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCaller) CONTRACTMULTISIGREMOVEADMINBYADDRESS(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REMOVE_ADMIN_BY_ADDRESS")
	return *ret0, err
}

// CONTRACTMULTISIGREMOVEADMINBYADDRESS is a free data retrieval call binding the contract method 0xc816b091.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_BY_ADDRESS() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplSession) CONTRACTMULTISIGREMOVEADMINBYADDRESS() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTMULTISIGREMOVEADMINBYADDRESS(&_FeeStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEADMINBYADDRESS is a free data retrieval call binding the contract method 0xc816b091.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_BY_ADDRESS() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCallerSession) CONTRACTMULTISIGREMOVEADMINBYADDRESS() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTMULTISIGREMOVEADMINBYADDRESS(&_FeeStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEADMINBYINDEX is a free data retrieval call binding the contract method 0x751fbea2.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_BY_INDEX() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCaller) CONTRACTMULTISIGREMOVEADMINBYINDEX(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REMOVE_ADMIN_BY_INDEX")
	return *ret0, err
}

// CONTRACTMULTISIGREMOVEADMINBYINDEX is a free data retrieval call binding the contract method 0x751fbea2.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_BY_INDEX() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplSession) CONTRACTMULTISIGREMOVEADMINBYINDEX() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTMULTISIGREMOVEADMINBYINDEX(&_FeeStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEADMINBYINDEX is a free data retrieval call binding the contract method 0x751fbea2.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_BY_INDEX() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCallerSession) CONTRACTMULTISIGREMOVEADMINBYINDEX() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTMULTISIGREMOVEADMINBYINDEX(&_FeeStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEADMINROLE is a free data retrieval call binding the contract method 0x53c00200.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_ROLE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCaller) CONTRACTMULTISIGREMOVEADMINROLE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REMOVE_ADMIN_ROLE")
	return *ret0, err
}

// CONTRACTMULTISIGREMOVEADMINROLE is a free data retrieval call binding the contract method 0x53c00200.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_ROLE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplSession) CONTRACTMULTISIGREMOVEADMINROLE() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTMULTISIGREMOVEADMINROLE(&_FeeStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEADMINROLE is a free data retrieval call binding the contract method 0x53c00200.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_ROLE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCallerSession) CONTRACTMULTISIGREMOVEADMINROLE() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTMULTISIGREMOVEADMINROLE(&_FeeStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEFEE is a free data retrieval call binding the contract method 0xdb614d25.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_FEE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCaller) CONTRACTMULTISIGREMOVEFEE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REMOVE_FEE")
	return *ret0, err
}

// CONTRACTMULTISIGREMOVEFEE is a free data retrieval call binding the contract method 0xdb614d25.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_FEE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplSession) CONTRACTMULTISIGREMOVEFEE() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTMULTISIGREMOVEFEE(&_FeeStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEFEE is a free data retrieval call binding the contract method 0xdb614d25.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_FEE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCallerSession) CONTRACTMULTISIGREMOVEFEE() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTMULTISIGREMOVEFEE(&_FeeStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEROLE is a free data retrieval call binding the contract method 0x23648496.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ROLE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCaller) CONTRACTMULTISIGREMOVEROLE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REMOVE_ROLE")
	return *ret0, err
}

// CONTRACTMULTISIGREMOVEROLE is a free data retrieval call binding the contract method 0x23648496.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ROLE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplSession) CONTRACTMULTISIGREMOVEROLE() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTMULTISIGREMOVEROLE(&_FeeStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEROLE is a free data retrieval call binding the contract method 0x23648496.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ROLE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCallerSession) CONTRACTMULTISIGREMOVEROLE() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTMULTISIGREMOVEROLE(&_FeeStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREPLENISHCAPITALIZATION is a free data retrieval call binding the contract method 0x62d5d523.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_CAPITALIZATION() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCaller) CONTRACTMULTISIGREPLENISHCAPITALIZATION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REPLENISH_CAPITALIZATION")
	return *ret0, err
}

// CONTRACTMULTISIGREPLENISHCAPITALIZATION is a free data retrieval call binding the contract method 0x62d5d523.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_CAPITALIZATION() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplSession) CONTRACTMULTISIGREPLENISHCAPITALIZATION() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTMULTISIGREPLENISHCAPITALIZATION(&_FeeStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREPLENISHCAPITALIZATION is a free data retrieval call binding the contract method 0x62d5d523.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_CAPITALIZATION() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCallerSession) CONTRACTMULTISIGREPLENISHCAPITALIZATION() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTMULTISIGREPLENISHCAPITALIZATION(&_FeeStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREPLENISHCERTIFICATE is a free data retrieval call binding the contract method 0xfd19b7e1.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_CERTIFICATE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCaller) CONTRACTMULTISIGREPLENISHCERTIFICATE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REPLENISH_CERTIFICATE")
	return *ret0, err
}

// CONTRACTMULTISIGREPLENISHCERTIFICATE is a free data retrieval call binding the contract method 0xfd19b7e1.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_CERTIFICATE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplSession) CONTRACTMULTISIGREPLENISHCERTIFICATE() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTMULTISIGREPLENISHCERTIFICATE(&_FeeStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREPLENISHCERTIFICATE is a free data retrieval call binding the contract method 0xfd19b7e1.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_CERTIFICATE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCallerSession) CONTRACTMULTISIGREPLENISHCERTIFICATE() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTMULTISIGREPLENISHCERTIFICATE(&_FeeStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREPLENISHEMISSION is a free data retrieval call binding the contract method 0x0cb59fde.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_EMISSION() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCaller) CONTRACTMULTISIGREPLENISHEMISSION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REPLENISH_EMISSION")
	return *ret0, err
}

// CONTRACTMULTISIGREPLENISHEMISSION is a free data retrieval call binding the contract method 0x0cb59fde.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_EMISSION() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplSession) CONTRACTMULTISIGREPLENISHEMISSION() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTMULTISIGREPLENISHEMISSION(&_FeeStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREPLENISHEMISSION is a free data retrieval call binding the contract method 0x0cb59fde.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_EMISSION() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCallerSession) CONTRACTMULTISIGREPLENISHEMISSION() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTMULTISIGREPLENISHEMISSION(&_FeeStorageImpl.CallOpts)
}

// CONTRACTMULTISIGSETADMINROLE is a free data retrieval call binding the contract method 0x8f209ab7.
//
// Solidity: function CONTRACT_MULTISIG_SET_ADMIN_ROLE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCaller) CONTRACTMULTISIGSETADMINROLE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_SET_ADMIN_ROLE")
	return *ret0, err
}

// CONTRACTMULTISIGSETADMINROLE is a free data retrieval call binding the contract method 0x8f209ab7.
//
// Solidity: function CONTRACT_MULTISIG_SET_ADMIN_ROLE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplSession) CONTRACTMULTISIGSETADMINROLE() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTMULTISIGSETADMINROLE(&_FeeStorageImpl.CallOpts)
}

// CONTRACTMULTISIGSETADMINROLE is a free data retrieval call binding the contract method 0x8f209ab7.
//
// Solidity: function CONTRACT_MULTISIG_SET_ADMIN_ROLE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCallerSession) CONTRACTMULTISIGSETADMINROLE() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTMULTISIGSETADMINROLE(&_FeeStorageImpl.CallOpts)
}

// CONTRACTMULTISIGSETCONFIRMATIONRATE is a free data retrieval call binding the contract method 0xa997279b.
//
// Solidity: function CONTRACT_MULTISIG_SET_CONFIRMATION_RATE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCaller) CONTRACTMULTISIGSETCONFIRMATIONRATE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_SET_CONFIRMATION_RATE")
	return *ret0, err
}

// CONTRACTMULTISIGSETCONFIRMATIONRATE is a free data retrieval call binding the contract method 0xa997279b.
//
// Solidity: function CONTRACT_MULTISIG_SET_CONFIRMATION_RATE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplSession) CONTRACTMULTISIGSETCONFIRMATIONRATE() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTMULTISIGSETCONFIRMATIONRATE(&_FeeStorageImpl.CallOpts)
}

// CONTRACTMULTISIGSETCONFIRMATIONRATE is a free data retrieval call binding the contract method 0xa997279b.
//
// Solidity: function CONTRACT_MULTISIG_SET_CONFIRMATION_RATE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCallerSession) CONTRACTMULTISIGSETCONFIRMATIONRATE() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTMULTISIGSETCONFIRMATIONRATE(&_FeeStorageImpl.CallOpts)
}

// CONTRACTMULTISIGSETFEEDISTRIBUTION is a free data retrieval call binding the contract method 0x0b76e64d.
//
// Solidity: function CONTRACT_MULTISIG_SET_FEE_DISTRIBUTION() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCaller) CONTRACTMULTISIGSETFEEDISTRIBUTION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_SET_FEE_DISTRIBUTION")
	return *ret0, err
}

// CONTRACTMULTISIGSETFEEDISTRIBUTION is a free data retrieval call binding the contract method 0x0b76e64d.
//
// Solidity: function CONTRACT_MULTISIG_SET_FEE_DISTRIBUTION() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplSession) CONTRACTMULTISIGSETFEEDISTRIBUTION() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTMULTISIGSETFEEDISTRIBUTION(&_FeeStorageImpl.CallOpts)
}

// CONTRACTMULTISIGSETFEEDISTRIBUTION is a free data retrieval call binding the contract method 0x0b76e64d.
//
// Solidity: function CONTRACT_MULTISIG_SET_FEE_DISTRIBUTION() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCallerSession) CONTRACTMULTISIGSETFEEDISTRIBUTION() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTMULTISIGSETFEEDISTRIBUTION(&_FeeStorageImpl.CallOpts)
}

// CONTRACTMULTISIGSETFEEEXCEPTIONS is a free data retrieval call binding the contract method 0x5e6b28a0.
//
// Solidity: function CONTRACT_MULTISIG_SET_FEE_EXCEPTIONS() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCaller) CONTRACTMULTISIGSETFEEEXCEPTIONS(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_SET_FEE_EXCEPTIONS")
	return *ret0, err
}

// CONTRACTMULTISIGSETFEEEXCEPTIONS is a free data retrieval call binding the contract method 0x5e6b28a0.
//
// Solidity: function CONTRACT_MULTISIG_SET_FEE_EXCEPTIONS() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplSession) CONTRACTMULTISIGSETFEEEXCEPTIONS() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTMULTISIGSETFEEEXCEPTIONS(&_FeeStorageImpl.CallOpts)
}

// CONTRACTMULTISIGSETFEEEXCEPTIONS is a free data retrieval call binding the contract method 0x5e6b28a0.
//
// Solidity: function CONTRACT_MULTISIG_SET_FEE_EXCEPTIONS() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCallerSession) CONTRACTMULTISIGSETFEEEXCEPTIONS() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTMULTISIGSETFEEEXCEPTIONS(&_FeeStorageImpl.CallOpts)
}

// CONTRACTMULTISIGSETMAXFEE is a free data retrieval call binding the contract method 0xa5f59013.
//
// Solidity: function CONTRACT_MULTISIG_SET_MAX_FEE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCaller) CONTRACTMULTISIGSETMAXFEE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_SET_MAX_FEE")
	return *ret0, err
}

// CONTRACTMULTISIGSETMAXFEE is a free data retrieval call binding the contract method 0xa5f59013.
//
// Solidity: function CONTRACT_MULTISIG_SET_MAX_FEE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplSession) CONTRACTMULTISIGSETMAXFEE() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTMULTISIGSETMAXFEE(&_FeeStorageImpl.CallOpts)
}

// CONTRACTMULTISIGSETMAXFEE is a free data retrieval call binding the contract method 0xa5f59013.
//
// Solidity: function CONTRACT_MULTISIG_SET_MAX_FEE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCallerSession) CONTRACTMULTISIGSETMAXFEE() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTMULTISIGSETMAXFEE(&_FeeStorageImpl.CallOpts)
}

// CONTRACTMULTISIGSETMINAMOUNT is a free data retrieval call binding the contract method 0xaba147b6.
//
// Solidity: function CONTRACT_MULTISIG_SET_MIN_AMOUNT() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCaller) CONTRACTMULTISIGSETMINAMOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_SET_MIN_AMOUNT")
	return *ret0, err
}

// CONTRACTMULTISIGSETMINAMOUNT is a free data retrieval call binding the contract method 0xaba147b6.
//
// Solidity: function CONTRACT_MULTISIG_SET_MIN_AMOUNT() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplSession) CONTRACTMULTISIGSETMINAMOUNT() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTMULTISIGSETMINAMOUNT(&_FeeStorageImpl.CallOpts)
}

// CONTRACTMULTISIGSETMINAMOUNT is a free data retrieval call binding the contract method 0xaba147b6.
//
// Solidity: function CONTRACT_MULTISIG_SET_MIN_AMOUNT() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCallerSession) CONTRACTMULTISIGSETMINAMOUNT() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTMULTISIGSETMINAMOUNT(&_FeeStorageImpl.CallOpts)
}

// CONTRACTMULTISIGSETMINFEE is a free data retrieval call binding the contract method 0xf2d1904f.
//
// Solidity: function CONTRACT_MULTISIG_SET_MIN_FEE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCaller) CONTRACTMULTISIGSETMINFEE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_SET_MIN_FEE")
	return *ret0, err
}

// CONTRACTMULTISIGSETMINFEE is a free data retrieval call binding the contract method 0xf2d1904f.
//
// Solidity: function CONTRACT_MULTISIG_SET_MIN_FEE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplSession) CONTRACTMULTISIGSETMINFEE() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTMULTISIGSETMINFEE(&_FeeStorageImpl.CallOpts)
}

// CONTRACTMULTISIGSETMINFEE is a free data retrieval call binding the contract method 0xf2d1904f.
//
// Solidity: function CONTRACT_MULTISIG_SET_MIN_FEE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCallerSession) CONTRACTMULTISIGSETMINFEE() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTMULTISIGSETMINFEE(&_FeeStorageImpl.CallOpts)
}

// CONTRACTMULTISIGSETPERMISSION is a free data retrieval call binding the contract method 0xb0e91914.
//
// Solidity: function CONTRACT_MULTISIG_SET_PERMISSION() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCaller) CONTRACTMULTISIGSETPERMISSION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_SET_PERMISSION")
	return *ret0, err
}

// CONTRACTMULTISIGSETPERMISSION is a free data retrieval call binding the contract method 0xb0e91914.
//
// Solidity: function CONTRACT_MULTISIG_SET_PERMISSION() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplSession) CONTRACTMULTISIGSETPERMISSION() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTMULTISIGSETPERMISSION(&_FeeStorageImpl.CallOpts)
}

// CONTRACTMULTISIGSETPERMISSION is a free data retrieval call binding the contract method 0xb0e91914.
//
// Solidity: function CONTRACT_MULTISIG_SET_PERMISSION() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCallerSession) CONTRACTMULTISIGSETPERMISSION() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTMULTISIGSETPERMISSION(&_FeeStorageImpl.CallOpts)
}

// CONTRACTMULTISIGSTORAGE is a free data retrieval call binding the contract method 0xe7e1a303.
//
// Solidity: function CONTRACT_MULTISIG_STORAGE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCaller) CONTRACTMULTISIGSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_STORAGE")
	return *ret0, err
}

// CONTRACTMULTISIGSTORAGE is a free data retrieval call binding the contract method 0xe7e1a303.
//
// Solidity: function CONTRACT_MULTISIG_STORAGE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplSession) CONTRACTMULTISIGSTORAGE() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTMULTISIGSTORAGE(&_FeeStorageImpl.CallOpts)
}

// CONTRACTMULTISIGSTORAGE is a free data retrieval call binding the contract method 0xe7e1a303.
//
// Solidity: function CONTRACT_MULTISIG_STORAGE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCallerSession) CONTRACTMULTISIGSTORAGE() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTMULTISIGSTORAGE(&_FeeStorageImpl.CallOpts)
}

// CONTRACTOTC is a free data retrieval call binding the contract method 0x5eceac70.
//
// Solidity: function CONTRACT_OTC() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCaller) CONTRACTOTC(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeStorageImpl.contract.Call(opts, out, "CONTRACT_OTC")
	return *ret0, err
}

// CONTRACTOTC is a free data retrieval call binding the contract method 0x5eceac70.
//
// Solidity: function CONTRACT_OTC() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplSession) CONTRACTOTC() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTOTC(&_FeeStorageImpl.CallOpts)
}

// CONTRACTOTC is a free data retrieval call binding the contract method 0x5eceac70.
//
// Solidity: function CONTRACT_OTC() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCallerSession) CONTRACTOTC() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTOTC(&_FeeStorageImpl.CallOpts)
}

// CONTRACTOTCSTORAGE is a free data retrieval call binding the contract method 0x825b476f.
//
// Solidity: function CONTRACT_OTC_STORAGE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCaller) CONTRACTOTCSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeStorageImpl.contract.Call(opts, out, "CONTRACT_OTC_STORAGE")
	return *ret0, err
}

// CONTRACTOTCSTORAGE is a free data retrieval call binding the contract method 0x825b476f.
//
// Solidity: function CONTRACT_OTC_STORAGE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplSession) CONTRACTOTCSTORAGE() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTOTCSTORAGE(&_FeeStorageImpl.CallOpts)
}

// CONTRACTOTCSTORAGE is a free data retrieval call binding the contract method 0x825b476f.
//
// Solidity: function CONTRACT_OTC_STORAGE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCallerSession) CONTRACTOTCSTORAGE() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTOTCSTORAGE(&_FeeStorageImpl.CallOpts)
}

// CONTRACTREPLENISHCAPITALIZATION is a free data retrieval call binding the contract method 0xff5e09d8.
//
// Solidity: function CONTRACT_REPLENISH_CAPITALIZATION() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCaller) CONTRACTREPLENISHCAPITALIZATION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeStorageImpl.contract.Call(opts, out, "CONTRACT_REPLENISH_CAPITALIZATION")
	return *ret0, err
}

// CONTRACTREPLENISHCAPITALIZATION is a free data retrieval call binding the contract method 0xff5e09d8.
//
// Solidity: function CONTRACT_REPLENISH_CAPITALIZATION() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplSession) CONTRACTREPLENISHCAPITALIZATION() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTREPLENISHCAPITALIZATION(&_FeeStorageImpl.CallOpts)
}

// CONTRACTREPLENISHCAPITALIZATION is a free data retrieval call binding the contract method 0xff5e09d8.
//
// Solidity: function CONTRACT_REPLENISH_CAPITALIZATION() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCallerSession) CONTRACTREPLENISHCAPITALIZATION() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTREPLENISHCAPITALIZATION(&_FeeStorageImpl.CallOpts)
}

// CONTRACTRESERVEDALIASSTORAGE is a free data retrieval call binding the contract method 0x70a94aa7.
//
// Solidity: function CONTRACT_RESERVED_ALIAS_STORAGE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCaller) CONTRACTRESERVEDALIASSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeStorageImpl.contract.Call(opts, out, "CONTRACT_RESERVED_ALIAS_STORAGE")
	return *ret0, err
}

// CONTRACTRESERVEDALIASSTORAGE is a free data retrieval call binding the contract method 0x70a94aa7.
//
// Solidity: function CONTRACT_RESERVED_ALIAS_STORAGE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplSession) CONTRACTRESERVEDALIASSTORAGE() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTRESERVEDALIASSTORAGE(&_FeeStorageImpl.CallOpts)
}

// CONTRACTRESERVEDALIASSTORAGE is a free data retrieval call binding the contract method 0x70a94aa7.
//
// Solidity: function CONTRACT_RESERVED_ALIAS_STORAGE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCallerSession) CONTRACTRESERVEDALIASSTORAGE() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTRESERVEDALIASSTORAGE(&_FeeStorageImpl.CallOpts)
}

// CONTRACTROLESTORAGE is a free data retrieval call binding the contract method 0xb938235b.
//
// Solidity: function CONTRACT_ROLE_STORAGE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCaller) CONTRACTROLESTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeStorageImpl.contract.Call(opts, out, "CONTRACT_ROLE_STORAGE")
	return *ret0, err
}

// CONTRACTROLESTORAGE is a free data retrieval call binding the contract method 0xb938235b.
//
// Solidity: function CONTRACT_ROLE_STORAGE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplSession) CONTRACTROLESTORAGE() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTROLESTORAGE(&_FeeStorageImpl.CallOpts)
}

// CONTRACTROLESTORAGE is a free data retrieval call binding the contract method 0xb938235b.
//
// Solidity: function CONTRACT_ROLE_STORAGE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCallerSession) CONTRACTROLESTORAGE() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTROLESTORAGE(&_FeeStorageImpl.CallOpts)
}

// CONTRACTSAFETRANSFERSTORAGE is a free data retrieval call binding the contract method 0xe7e45392.
//
// Solidity: function CONTRACT_SAFE_TRANSFER_STORAGE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCaller) CONTRACTSAFETRANSFERSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeStorageImpl.contract.Call(opts, out, "CONTRACT_SAFE_TRANSFER_STORAGE")
	return *ret0, err
}

// CONTRACTSAFETRANSFERSTORAGE is a free data retrieval call binding the contract method 0xe7e45392.
//
// Solidity: function CONTRACT_SAFE_TRANSFER_STORAGE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplSession) CONTRACTSAFETRANSFERSTORAGE() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTSAFETRANSFERSTORAGE(&_FeeStorageImpl.CallOpts)
}

// CONTRACTSAFETRANSFERSTORAGE is a free data retrieval call binding the contract method 0xe7e45392.
//
// Solidity: function CONTRACT_SAFE_TRANSFER_STORAGE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCallerSession) CONTRACTSAFETRANSFERSTORAGE() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTSAFETRANSFERSTORAGE(&_FeeStorageImpl.CallOpts)
}

// CONTRACTTRANSFER is a free data retrieval call binding the contract method 0xef94fe38.
//
// Solidity: function CONTRACT_TRANSFER() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCaller) CONTRACTTRANSFER(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeStorageImpl.contract.Call(opts, out, "CONTRACT_TRANSFER")
	return *ret0, err
}

// CONTRACTTRANSFER is a free data retrieval call binding the contract method 0xef94fe38.
//
// Solidity: function CONTRACT_TRANSFER() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplSession) CONTRACTTRANSFER() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTTRANSFER(&_FeeStorageImpl.CallOpts)
}

// CONTRACTTRANSFER is a free data retrieval call binding the contract method 0xef94fe38.
//
// Solidity: function CONTRACT_TRANSFER() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCallerSession) CONTRACTTRANSFER() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTTRANSFER(&_FeeStorageImpl.CallOpts)
}

// CONTRACTTRANSFERFROMCAPITALIZATION is a free data retrieval call binding the contract method 0xc8f8decb.
//
// Solidity: function CONTRACT_TRANSFER_FROM_CAPITALIZATION() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCaller) CONTRACTTRANSFERFROMCAPITALIZATION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeStorageImpl.contract.Call(opts, out, "CONTRACT_TRANSFER_FROM_CAPITALIZATION")
	return *ret0, err
}

// CONTRACTTRANSFERFROMCAPITALIZATION is a free data retrieval call binding the contract method 0xc8f8decb.
//
// Solidity: function CONTRACT_TRANSFER_FROM_CAPITALIZATION() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplSession) CONTRACTTRANSFERFROMCAPITALIZATION() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTTRANSFERFROMCAPITALIZATION(&_FeeStorageImpl.CallOpts)
}

// CONTRACTTRANSFERFROMCAPITALIZATION is a free data retrieval call binding the contract method 0xc8f8decb.
//
// Solidity: function CONTRACT_TRANSFER_FROM_CAPITALIZATION() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCallerSession) CONTRACTTRANSFERFROMCAPITALIZATION() (*big.Int, error) {
	return _FeeStorageImpl.Contract.CONTRACTTRANSFERFROMCAPITALIZATION(&_FeeStorageImpl.CallOpts)
}

// FEEOPERATIONCREATESAFETRANSFER is a free data retrieval call binding the contract method 0xf46d727d.
//
// Solidity: function FEE_OPERATION_CREATE_SAFE_TRANSFER() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCaller) FEEOPERATIONCREATESAFETRANSFER(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeStorageImpl.contract.Call(opts, out, "FEE_OPERATION_CREATE_SAFE_TRANSFER")
	return *ret0, err
}

// FEEOPERATIONCREATESAFETRANSFER is a free data retrieval call binding the contract method 0xf46d727d.
//
// Solidity: function FEE_OPERATION_CREATE_SAFE_TRANSFER() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplSession) FEEOPERATIONCREATESAFETRANSFER() (*big.Int, error) {
	return _FeeStorageImpl.Contract.FEEOPERATIONCREATESAFETRANSFER(&_FeeStorageImpl.CallOpts)
}

// FEEOPERATIONCREATESAFETRANSFER is a free data retrieval call binding the contract method 0xf46d727d.
//
// Solidity: function FEE_OPERATION_CREATE_SAFE_TRANSFER() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCallerSession) FEEOPERATIONCREATESAFETRANSFER() (*big.Int, error) {
	return _FeeStorageImpl.Contract.FEEOPERATIONCREATESAFETRANSFER(&_FeeStorageImpl.CallOpts)
}

// FEEOPERATIONCREATESAFETRANSFERALL is a free data retrieval call binding the contract method 0x9d558bef.
//
// Solidity: function FEE_OPERATION_CREATE_SAFE_TRANSFER_ALL() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCaller) FEEOPERATIONCREATESAFETRANSFERALL(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeStorageImpl.contract.Call(opts, out, "FEE_OPERATION_CREATE_SAFE_TRANSFER_ALL")
	return *ret0, err
}

// FEEOPERATIONCREATESAFETRANSFERALL is a free data retrieval call binding the contract method 0x9d558bef.
//
// Solidity: function FEE_OPERATION_CREATE_SAFE_TRANSFER_ALL() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplSession) FEEOPERATIONCREATESAFETRANSFERALL() (*big.Int, error) {
	return _FeeStorageImpl.Contract.FEEOPERATIONCREATESAFETRANSFERALL(&_FeeStorageImpl.CallOpts)
}

// FEEOPERATIONCREATESAFETRANSFERALL is a free data retrieval call binding the contract method 0x9d558bef.
//
// Solidity: function FEE_OPERATION_CREATE_SAFE_TRANSFER_ALL() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCallerSession) FEEOPERATIONCREATESAFETRANSFERALL() (*big.Int, error) {
	return _FeeStorageImpl.Contract.FEEOPERATIONCREATESAFETRANSFERALL(&_FeeStorageImpl.CallOpts)
}

// FEEOPERATIONFRAMETRANSFER is a free data retrieval call binding the contract method 0xc0c92120.
//
// Solidity: function FEE_OPERATION_FRAME_TRANSFER() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCaller) FEEOPERATIONFRAMETRANSFER(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeStorageImpl.contract.Call(opts, out, "FEE_OPERATION_FRAME_TRANSFER")
	return *ret0, err
}

// FEEOPERATIONFRAMETRANSFER is a free data retrieval call binding the contract method 0xc0c92120.
//
// Solidity: function FEE_OPERATION_FRAME_TRANSFER() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplSession) FEEOPERATIONFRAMETRANSFER() (*big.Int, error) {
	return _FeeStorageImpl.Contract.FEEOPERATIONFRAMETRANSFER(&_FeeStorageImpl.CallOpts)
}

// FEEOPERATIONFRAMETRANSFER is a free data retrieval call binding the contract method 0xc0c92120.
//
// Solidity: function FEE_OPERATION_FRAME_TRANSFER() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCallerSession) FEEOPERATIONFRAMETRANSFER() (*big.Int, error) {
	return _FeeStorageImpl.Contract.FEEOPERATIONFRAMETRANSFER(&_FeeStorageImpl.CallOpts)
}

// FEEOPERATIONQUICKTRANSFER is a free data retrieval call binding the contract method 0xbd64b596.
//
// Solidity: function FEE_OPERATION_QUICK_TRANSFER() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCaller) FEEOPERATIONQUICKTRANSFER(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeStorageImpl.contract.Call(opts, out, "FEE_OPERATION_QUICK_TRANSFER")
	return *ret0, err
}

// FEEOPERATIONQUICKTRANSFER is a free data retrieval call binding the contract method 0xbd64b596.
//
// Solidity: function FEE_OPERATION_QUICK_TRANSFER() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplSession) FEEOPERATIONQUICKTRANSFER() (*big.Int, error) {
	return _FeeStorageImpl.Contract.FEEOPERATIONQUICKTRANSFER(&_FeeStorageImpl.CallOpts)
}

// FEEOPERATIONQUICKTRANSFER is a free data retrieval call binding the contract method 0xbd64b596.
//
// Solidity: function FEE_OPERATION_QUICK_TRANSFER() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCallerSession) FEEOPERATIONQUICKTRANSFER() (*big.Int, error) {
	return _FeeStorageImpl.Contract.FEEOPERATIONQUICKTRANSFER(&_FeeStorageImpl.CallOpts)
}

// FEEOPERATIONQUICKTRANSFERALL is a free data retrieval call binding the contract method 0x69cc9631.
//
// Solidity: function FEE_OPERATION_QUICK_TRANSFER_ALL() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCaller) FEEOPERATIONQUICKTRANSFERALL(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeStorageImpl.contract.Call(opts, out, "FEE_OPERATION_QUICK_TRANSFER_ALL")
	return *ret0, err
}

// FEEOPERATIONQUICKTRANSFERALL is a free data retrieval call binding the contract method 0x69cc9631.
//
// Solidity: function FEE_OPERATION_QUICK_TRANSFER_ALL() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplSession) FEEOPERATIONQUICKTRANSFERALL() (*big.Int, error) {
	return _FeeStorageImpl.Contract.FEEOPERATIONQUICKTRANSFERALL(&_FeeStorageImpl.CallOpts)
}

// FEEOPERATIONQUICKTRANSFERALL is a free data retrieval call binding the contract method 0x69cc9631.
//
// Solidity: function FEE_OPERATION_QUICK_TRANSFER_ALL() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCallerSession) FEEOPERATIONQUICKTRANSFERALL() (*big.Int, error) {
	return _FeeStorageImpl.Contract.FEEOPERATIONQUICKTRANSFERALL(&_FeeStorageImpl.CallOpts)
}

// FEEOPERATIONWPCSAFETRANSFER is a free data retrieval call binding the contract method 0x33594855.
//
// Solidity: function FEE_OPERATION_WPC_SAFE_TRANSFER() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCaller) FEEOPERATIONWPCSAFETRANSFER(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeStorageImpl.contract.Call(opts, out, "FEE_OPERATION_WPC_SAFE_TRANSFER")
	return *ret0, err
}

// FEEOPERATIONWPCSAFETRANSFER is a free data retrieval call binding the contract method 0x33594855.
//
// Solidity: function FEE_OPERATION_WPC_SAFE_TRANSFER() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplSession) FEEOPERATIONWPCSAFETRANSFER() (*big.Int, error) {
	return _FeeStorageImpl.Contract.FEEOPERATIONWPCSAFETRANSFER(&_FeeStorageImpl.CallOpts)
}

// FEEOPERATIONWPCSAFETRANSFER is a free data retrieval call binding the contract method 0x33594855.
//
// Solidity: function FEE_OPERATION_WPC_SAFE_TRANSFER() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCallerSession) FEEOPERATIONWPCSAFETRANSFER() (*big.Int, error) {
	return _FeeStorageImpl.Contract.FEEOPERATIONWPCSAFETRANSFER(&_FeeStorageImpl.CallOpts)
}

// FEEOPERATIONWPCSAFETRANSFERALL is a free data retrieval call binding the contract method 0xf0ef7f20.
//
// Solidity: function FEE_OPERATION_WPC_SAFE_TRANSFER_ALL() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCaller) FEEOPERATIONWPCSAFETRANSFERALL(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeStorageImpl.contract.Call(opts, out, "FEE_OPERATION_WPC_SAFE_TRANSFER_ALL")
	return *ret0, err
}

// FEEOPERATIONWPCSAFETRANSFERALL is a free data retrieval call binding the contract method 0xf0ef7f20.
//
// Solidity: function FEE_OPERATION_WPC_SAFE_TRANSFER_ALL() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplSession) FEEOPERATIONWPCSAFETRANSFERALL() (*big.Int, error) {
	return _FeeStorageImpl.Contract.FEEOPERATIONWPCSAFETRANSFERALL(&_FeeStorageImpl.CallOpts)
}

// FEEOPERATIONWPCSAFETRANSFERALL is a free data retrieval call binding the contract method 0xf0ef7f20.
//
// Solidity: function FEE_OPERATION_WPC_SAFE_TRANSFER_ALL() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCallerSession) FEEOPERATIONWPCSAFETRANSFERALL() (*big.Int, error) {
	return _FeeStorageImpl.Contract.FEEOPERATIONWPCSAFETRANSFERALL(&_FeeStorageImpl.CallOpts)
}

// PERMISSIONACCESSRESTORATION is a free data retrieval call binding the contract method 0x5c3053bf.
//
// Solidity: function PERMISSION_ACCESS_RESTORATION() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCaller) PERMISSIONACCESSRESTORATION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeStorageImpl.contract.Call(opts, out, "PERMISSION_ACCESS_RESTORATION")
	return *ret0, err
}

// PERMISSIONACCESSRESTORATION is a free data retrieval call binding the contract method 0x5c3053bf.
//
// Solidity: function PERMISSION_ACCESS_RESTORATION() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplSession) PERMISSIONACCESSRESTORATION() (*big.Int, error) {
	return _FeeStorageImpl.Contract.PERMISSIONACCESSRESTORATION(&_FeeStorageImpl.CallOpts)
}

// PERMISSIONACCESSRESTORATION is a free data retrieval call binding the contract method 0x5c3053bf.
//
// Solidity: function PERMISSION_ACCESS_RESTORATION() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCallerSession) PERMISSIONACCESSRESTORATION() (*big.Int, error) {
	return _FeeStorageImpl.Contract.PERMISSIONACCESSRESTORATION(&_FeeStorageImpl.CallOpts)
}

// PERMISSIONADMINSTORAGE is a free data retrieval call binding the contract method 0x0c9fc0db.
//
// Solidity: function PERMISSION_ADMIN_STORAGE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCaller) PERMISSIONADMINSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeStorageImpl.contract.Call(opts, out, "PERMISSION_ADMIN_STORAGE")
	return *ret0, err
}

// PERMISSIONADMINSTORAGE is a free data retrieval call binding the contract method 0x0c9fc0db.
//
// Solidity: function PERMISSION_ADMIN_STORAGE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplSession) PERMISSIONADMINSTORAGE() (*big.Int, error) {
	return _FeeStorageImpl.Contract.PERMISSIONADMINSTORAGE(&_FeeStorageImpl.CallOpts)
}

// PERMISSIONADMINSTORAGE is a free data retrieval call binding the contract method 0x0c9fc0db.
//
// Solidity: function PERMISSION_ADMIN_STORAGE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCallerSession) PERMISSIONADMINSTORAGE() (*big.Int, error) {
	return _FeeStorageImpl.Contract.PERMISSIONADMINSTORAGE(&_FeeStorageImpl.CallOpts)
}

// PERMISSIONCONTRACTSSTORAGE is a free data retrieval call binding the contract method 0x1516e83b.
//
// Solidity: function PERMISSION_CONTRACTS_STORAGE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCaller) PERMISSIONCONTRACTSSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeStorageImpl.contract.Call(opts, out, "PERMISSION_CONTRACTS_STORAGE")
	return *ret0, err
}

// PERMISSIONCONTRACTSSTORAGE is a free data retrieval call binding the contract method 0x1516e83b.
//
// Solidity: function PERMISSION_CONTRACTS_STORAGE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplSession) PERMISSIONCONTRACTSSTORAGE() (*big.Int, error) {
	return _FeeStorageImpl.Contract.PERMISSIONCONTRACTSSTORAGE(&_FeeStorageImpl.CallOpts)
}

// PERMISSIONCONTRACTSSTORAGE is a free data retrieval call binding the contract method 0x1516e83b.
//
// Solidity: function PERMISSION_CONTRACTS_STORAGE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCallerSession) PERMISSIONCONTRACTSSTORAGE() (*big.Int, error) {
	return _FeeStorageImpl.Contract.PERMISSIONCONTRACTSSTORAGE(&_FeeStorageImpl.CallOpts)
}

// PERMISSIONCONTRACTDEPLOY is a free data retrieval call binding the contract method 0x257a4a04.
//
// Solidity: function PERMISSION_CONTRACT_DEPLOY() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCaller) PERMISSIONCONTRACTDEPLOY(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeStorageImpl.contract.Call(opts, out, "PERMISSION_CONTRACT_DEPLOY")
	return *ret0, err
}

// PERMISSIONCONTRACTDEPLOY is a free data retrieval call binding the contract method 0x257a4a04.
//
// Solidity: function PERMISSION_CONTRACT_DEPLOY() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplSession) PERMISSIONCONTRACTDEPLOY() (*big.Int, error) {
	return _FeeStorageImpl.Contract.PERMISSIONCONTRACTDEPLOY(&_FeeStorageImpl.CallOpts)
}

// PERMISSIONCONTRACTDEPLOY is a free data retrieval call binding the contract method 0x257a4a04.
//
// Solidity: function PERMISSION_CONTRACT_DEPLOY() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCallerSession) PERMISSIONCONTRACTDEPLOY() (*big.Int, error) {
	return _FeeStorageImpl.Contract.PERMISSIONCONTRACTDEPLOY(&_FeeStorageImpl.CallOpts)
}

// PERMISSIONEMISSION is a free data retrieval call binding the contract method 0x4ccd92d1.
//
// Solidity: function PERMISSION_EMISSION() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCaller) PERMISSIONEMISSION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeStorageImpl.contract.Call(opts, out, "PERMISSION_EMISSION")
	return *ret0, err
}

// PERMISSIONEMISSION is a free data retrieval call binding the contract method 0x4ccd92d1.
//
// Solidity: function PERMISSION_EMISSION() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplSession) PERMISSIONEMISSION() (*big.Int, error) {
	return _FeeStorageImpl.Contract.PERMISSIONEMISSION(&_FeeStorageImpl.CallOpts)
}

// PERMISSIONEMISSION is a free data retrieval call binding the contract method 0x4ccd92d1.
//
// Solidity: function PERMISSION_EMISSION() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCallerSession) PERMISSIONEMISSION() (*big.Int, error) {
	return _FeeStorageImpl.Contract.PERMISSIONEMISSION(&_FeeStorageImpl.CallOpts)
}

// PERMISSIONFEESSTORAGE is a free data retrieval call binding the contract method 0xe837eb87.
//
// Solidity: function PERMISSION_FEES_STORAGE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCaller) PERMISSIONFEESSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeStorageImpl.contract.Call(opts, out, "PERMISSION_FEES_STORAGE")
	return *ret0, err
}

// PERMISSIONFEESSTORAGE is a free data retrieval call binding the contract method 0xe837eb87.
//
// Solidity: function PERMISSION_FEES_STORAGE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplSession) PERMISSIONFEESSTORAGE() (*big.Int, error) {
	return _FeeStorageImpl.Contract.PERMISSIONFEESSTORAGE(&_FeeStorageImpl.CallOpts)
}

// PERMISSIONFEESSTORAGE is a free data retrieval call binding the contract method 0xe837eb87.
//
// Solidity: function PERMISSION_FEES_STORAGE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCallerSession) PERMISSIONFEESSTORAGE() (*big.Int, error) {
	return _FeeStorageImpl.Contract.PERMISSIONFEESSTORAGE(&_FeeStorageImpl.CallOpts)
}

// PERMISSIONMULTISIGSTORAGE is a free data retrieval call binding the contract method 0xf1bb7aad.
//
// Solidity: function PERMISSION_MULTISIG_STORAGE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCaller) PERMISSIONMULTISIGSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeStorageImpl.contract.Call(opts, out, "PERMISSION_MULTISIG_STORAGE")
	return *ret0, err
}

// PERMISSIONMULTISIGSTORAGE is a free data retrieval call binding the contract method 0xf1bb7aad.
//
// Solidity: function PERMISSION_MULTISIG_STORAGE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplSession) PERMISSIONMULTISIGSTORAGE() (*big.Int, error) {
	return _FeeStorageImpl.Contract.PERMISSIONMULTISIGSTORAGE(&_FeeStorageImpl.CallOpts)
}

// PERMISSIONMULTISIGSTORAGE is a free data retrieval call binding the contract method 0xf1bb7aad.
//
// Solidity: function PERMISSION_MULTISIG_STORAGE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCallerSession) PERMISSIONMULTISIGSTORAGE() (*big.Int, error) {
	return _FeeStorageImpl.Contract.PERMISSIONMULTISIGSTORAGE(&_FeeStorageImpl.CallOpts)
}

// PERMISSIONROLESSTORAGE is a free data retrieval call binding the contract method 0x12a16032.
//
// Solidity: function PERMISSION_ROLES_STORAGE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCaller) PERMISSIONROLESSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeStorageImpl.contract.Call(opts, out, "PERMISSION_ROLES_STORAGE")
	return *ret0, err
}

// PERMISSIONROLESSTORAGE is a free data retrieval call binding the contract method 0x12a16032.
//
// Solidity: function PERMISSION_ROLES_STORAGE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplSession) PERMISSIONROLESSTORAGE() (*big.Int, error) {
	return _FeeStorageImpl.Contract.PERMISSIONROLESSTORAGE(&_FeeStorageImpl.CallOpts)
}

// PERMISSIONROLESSTORAGE is a free data retrieval call binding the contract method 0x12a16032.
//
// Solidity: function PERMISSION_ROLES_STORAGE() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCallerSession) PERMISSIONROLESSTORAGE() (*big.Int, error) {
	return _FeeStorageImpl.Contract.PERMISSIONROLESSTORAGE(&_FeeStorageImpl.CallOpts)
}

// ROLEACCESSVERIFIER is a free data retrieval call binding the contract method 0xc5405f92.
//
// Solidity: function ROLE_ACCESS_VERIFIER() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCaller) ROLEACCESSVERIFIER(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeStorageImpl.contract.Call(opts, out, "ROLE_ACCESS_VERIFIER")
	return *ret0, err
}

// ROLEACCESSVERIFIER is a free data retrieval call binding the contract method 0xc5405f92.
//
// Solidity: function ROLE_ACCESS_VERIFIER() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplSession) ROLEACCESSVERIFIER() (*big.Int, error) {
	return _FeeStorageImpl.Contract.ROLEACCESSVERIFIER(&_FeeStorageImpl.CallOpts)
}

// ROLEACCESSVERIFIER is a free data retrieval call binding the contract method 0xc5405f92.
//
// Solidity: function ROLE_ACCESS_VERIFIER() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCallerSession) ROLEACCESSVERIFIER() (*big.Int, error) {
	return _FeeStorageImpl.Contract.ROLEACCESSVERIFIER(&_FeeStorageImpl.CallOpts)
}

// ROLEBITBONISSUEVERIFIER is a free data retrieval call binding the contract method 0xf0a012a6.
//
// Solidity: function ROLE_BITBON_ISSUE_VERIFIER() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCaller) ROLEBITBONISSUEVERIFIER(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeStorageImpl.contract.Call(opts, out, "ROLE_BITBON_ISSUE_VERIFIER")
	return *ret0, err
}

// ROLEBITBONISSUEVERIFIER is a free data retrieval call binding the contract method 0xf0a012a6.
//
// Solidity: function ROLE_BITBON_ISSUE_VERIFIER() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplSession) ROLEBITBONISSUEVERIFIER() (*big.Int, error) {
	return _FeeStorageImpl.Contract.ROLEBITBONISSUEVERIFIER(&_FeeStorageImpl.CallOpts)
}

// ROLEBITBONISSUEVERIFIER is a free data retrieval call binding the contract method 0xf0a012a6.
//
// Solidity: function ROLE_BITBON_ISSUE_VERIFIER() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCallerSession) ROLEBITBONISSUEVERIFIER() (*big.Int, error) {
	return _FeeStorageImpl.Contract.ROLEBITBONISSUEVERIFIER(&_FeeStorageImpl.CallOpts)
}

// ROLECOMMISSIONVERIFIER is a free data retrieval call binding the contract method 0x8ab3de8c.
//
// Solidity: function ROLE_COMMISSION_VERIFIER() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCaller) ROLECOMMISSIONVERIFIER(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeStorageImpl.contract.Call(opts, out, "ROLE_COMMISSION_VERIFIER")
	return *ret0, err
}

// ROLECOMMISSIONVERIFIER is a free data retrieval call binding the contract method 0x8ab3de8c.
//
// Solidity: function ROLE_COMMISSION_VERIFIER() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplSession) ROLECOMMISSIONVERIFIER() (*big.Int, error) {
	return _FeeStorageImpl.Contract.ROLECOMMISSIONVERIFIER(&_FeeStorageImpl.CallOpts)
}

// ROLECOMMISSIONVERIFIER is a free data retrieval call binding the contract method 0x8ab3de8c.
//
// Solidity: function ROLE_COMMISSION_VERIFIER() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCallerSession) ROLECOMMISSIONVERIFIER() (*big.Int, error) {
	return _FeeStorageImpl.Contract.ROLECOMMISSIONVERIFIER(&_FeeStorageImpl.CallOpts)
}

// ROLEDEPLOYADMIN is a free data retrieval call binding the contract method 0x667bf6e5.
//
// Solidity: function ROLE_DEPLOY_ADMIN() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCaller) ROLEDEPLOYADMIN(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeStorageImpl.contract.Call(opts, out, "ROLE_DEPLOY_ADMIN")
	return *ret0, err
}

// ROLEDEPLOYADMIN is a free data retrieval call binding the contract method 0x667bf6e5.
//
// Solidity: function ROLE_DEPLOY_ADMIN() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplSession) ROLEDEPLOYADMIN() (*big.Int, error) {
	return _FeeStorageImpl.Contract.ROLEDEPLOYADMIN(&_FeeStorageImpl.CallOpts)
}

// ROLEDEPLOYADMIN is a free data retrieval call binding the contract method 0x667bf6e5.
//
// Solidity: function ROLE_DEPLOY_ADMIN() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCallerSession) ROLEDEPLOYADMIN() (*big.Int, error) {
	return _FeeStorageImpl.Contract.ROLEDEPLOYADMIN(&_FeeStorageImpl.CallOpts)
}

// ROLEPERMISSIONADMIN is a free data retrieval call binding the contract method 0x1b239996.
//
// Solidity: function ROLE_PERMISSION_ADMIN() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCaller) ROLEPERMISSIONADMIN(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeStorageImpl.contract.Call(opts, out, "ROLE_PERMISSION_ADMIN")
	return *ret0, err
}

// ROLEPERMISSIONADMIN is a free data retrieval call binding the contract method 0x1b239996.
//
// Solidity: function ROLE_PERMISSION_ADMIN() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplSession) ROLEPERMISSIONADMIN() (*big.Int, error) {
	return _FeeStorageImpl.Contract.ROLEPERMISSIONADMIN(&_FeeStorageImpl.CallOpts)
}

// ROLEPERMISSIONADMIN is a free data retrieval call binding the contract method 0x1b239996.
//
// Solidity: function ROLE_PERMISSION_ADMIN() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCallerSession) ROLEPERMISSIONADMIN() (*big.Int, error) {
	return _FeeStorageImpl.Contract.ROLEPERMISSIONADMIN(&_FeeStorageImpl.CallOpts)
}

// AccountToIsExceptional is a free data retrieval call binding the contract method 0xb0f3371c.
//
// Solidity: function accountToIsExceptional(address ) view returns(bool)
func (_FeeStorageImpl *FeeStorageImplCaller) AccountToIsExceptional(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _FeeStorageImpl.contract.Call(opts, out, "accountToIsExceptional", arg0)
	return *ret0, err
}

// AccountToIsExceptional is a free data retrieval call binding the contract method 0xb0f3371c.
//
// Solidity: function accountToIsExceptional(address ) view returns(bool)
func (_FeeStorageImpl *FeeStorageImplSession) AccountToIsExceptional(arg0 common.Address) (bool, error) {
	return _FeeStorageImpl.Contract.AccountToIsExceptional(&_FeeStorageImpl.CallOpts, arg0)
}

// AccountToIsExceptional is a free data retrieval call binding the contract method 0xb0f3371c.
//
// Solidity: function accountToIsExceptional(address ) view returns(bool)
func (_FeeStorageImpl *FeeStorageImplCallerSession) AccountToIsExceptional(arg0 common.Address) (bool, error) {
	return _FeeStorageImpl.Contract.AccountToIsExceptional(&_FeeStorageImpl.CallOpts, arg0)
}

// Confirmations is a free data retrieval call binding the contract method 0x9cf5d607.
//
// Solidity: function confirmations() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCaller) Confirmations(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeStorageImpl.contract.Call(opts, out, "confirmations")
	return *ret0, err
}

// Confirmations is a free data retrieval call binding the contract method 0x9cf5d607.
//
// Solidity: function confirmations() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplSession) Confirmations() (*big.Int, error) {
	return _FeeStorageImpl.Contract.Confirmations(&_FeeStorageImpl.CallOpts)
}

// Confirmations is a free data retrieval call binding the contract method 0x9cf5d607.
//
// Solidity: function confirmations() view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCallerSession) Confirmations() (*big.Int, error) {
	return _FeeStorageImpl.Contract.Confirmations(&_FeeStorageImpl.CallOpts)
}

// ContractStorage is a free data retrieval call binding the contract method 0x0549b27e.
//
// Solidity: function contractStorage() view returns(address)
func (_FeeStorageImpl *FeeStorageImplCaller) ContractStorage(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _FeeStorageImpl.contract.Call(opts, out, "contractStorage")
	return *ret0, err
}

// ContractStorage is a free data retrieval call binding the contract method 0x0549b27e.
//
// Solidity: function contractStorage() view returns(address)
func (_FeeStorageImpl *FeeStorageImplSession) ContractStorage() (common.Address, error) {
	return _FeeStorageImpl.Contract.ContractStorage(&_FeeStorageImpl.CallOpts)
}

// ContractStorage is a free data retrieval call binding the contract method 0x0549b27e.
//
// Solidity: function contractStorage() view returns(address)
func (_FeeStorageImpl *FeeStorageImplCallerSession) ContractStorage() (common.Address, error) {
	return _FeeStorageImpl.Contract.ContractStorage(&_FeeStorageImpl.CallOpts)
}

// Fees is a free data retrieval call binding the contract method 0x4acc79ed.
//
// Solidity: function fees(uint256 ) view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCaller) Fees(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeStorageImpl.contract.Call(opts, out, "fees", arg0)
	return *ret0, err
}

// Fees is a free data retrieval call binding the contract method 0x4acc79ed.
//
// Solidity: function fees(uint256 ) view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplSession) Fees(arg0 *big.Int) (*big.Int, error) {
	return _FeeStorageImpl.Contract.Fees(&_FeeStorageImpl.CallOpts, arg0)
}

// Fees is a free data retrieval call binding the contract method 0x4acc79ed.
//
// Solidity: function fees(uint256 ) view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCallerSession) Fees(arg0 *big.Int) (*big.Int, error) {
	return _FeeStorageImpl.Contract.Fees(&_FeeStorageImpl.CallOpts, arg0)
}

// Get is a free data retrieval call binding the contract method 0x9507d39a.
//
// Solidity: function get(uint256 feeType) view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCaller) Get(opts *bind.CallOpts, feeType *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeStorageImpl.contract.Call(opts, out, "get", feeType)
	return *ret0, err
}

// Get is a free data retrieval call binding the contract method 0x9507d39a.
//
// Solidity: function get(uint256 feeType) view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplSession) Get(feeType *big.Int) (*big.Int, error) {
	return _FeeStorageImpl.Contract.Get(&_FeeStorageImpl.CallOpts, feeType)
}

// Get is a free data retrieval call binding the contract method 0x9507d39a.
//
// Solidity: function get(uint256 feeType) view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCallerSession) Get(feeType *big.Int) (*big.Int, error) {
	return _FeeStorageImpl.Contract.Get(&_FeeStorageImpl.CallOpts, feeType)
}

// GetExceptionAccounts is a free data retrieval call binding the contract method 0xd9dbfa78.
//
// Solidity: function getExceptionAccounts() view returns(address[])
func (_FeeStorageImpl *FeeStorageImplCaller) GetExceptionAccounts(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _FeeStorageImpl.contract.Call(opts, out, "getExceptionAccounts")
	return *ret0, err
}

// GetExceptionAccounts is a free data retrieval call binding the contract method 0xd9dbfa78.
//
// Solidity: function getExceptionAccounts() view returns(address[])
func (_FeeStorageImpl *FeeStorageImplSession) GetExceptionAccounts() ([]common.Address, error) {
	return _FeeStorageImpl.Contract.GetExceptionAccounts(&_FeeStorageImpl.CallOpts)
}

// GetExceptionAccounts is a free data retrieval call binding the contract method 0xd9dbfa78.
//
// Solidity: function getExceptionAccounts() view returns(address[])
func (_FeeStorageImpl *FeeStorageImplCallerSession) GetExceptionAccounts() ([]common.Address, error) {
	return _FeeStorageImpl.Contract.GetExceptionAccounts(&_FeeStorageImpl.CallOpts)
}

// GetFeeDistributionAccounts is a free data retrieval call binding the contract method 0xb3284c48.
//
// Solidity: function getFeeDistributionAccounts(uint256 opType) view returns(address[])
func (_FeeStorageImpl *FeeStorageImplCaller) GetFeeDistributionAccounts(opts *bind.CallOpts, opType *big.Int) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _FeeStorageImpl.contract.Call(opts, out, "getFeeDistributionAccounts", opType)
	return *ret0, err
}

// GetFeeDistributionAccounts is a free data retrieval call binding the contract method 0xb3284c48.
//
// Solidity: function getFeeDistributionAccounts(uint256 opType) view returns(address[])
func (_FeeStorageImpl *FeeStorageImplSession) GetFeeDistributionAccounts(opType *big.Int) ([]common.Address, error) {
	return _FeeStorageImpl.Contract.GetFeeDistributionAccounts(&_FeeStorageImpl.CallOpts, opType)
}

// GetFeeDistributionAccounts is a free data retrieval call binding the contract method 0xb3284c48.
//
// Solidity: function getFeeDistributionAccounts(uint256 opType) view returns(address[])
func (_FeeStorageImpl *FeeStorageImplCallerSession) GetFeeDistributionAccounts(opType *big.Int) ([]common.Address, error) {
	return _FeeStorageImpl.Contract.GetFeeDistributionAccounts(&_FeeStorageImpl.CallOpts, opType)
}

// GetFeeDistributionAmounts is a free data retrieval call binding the contract method 0x757d77a4.
//
// Solidity: function getFeeDistributionAmounts(uint256 opType) view returns(uint256[])
func (_FeeStorageImpl *FeeStorageImplCaller) GetFeeDistributionAmounts(opts *bind.CallOpts, opType *big.Int) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _FeeStorageImpl.contract.Call(opts, out, "getFeeDistributionAmounts", opType)
	return *ret0, err
}

// GetFeeDistributionAmounts is a free data retrieval call binding the contract method 0x757d77a4.
//
// Solidity: function getFeeDistributionAmounts(uint256 opType) view returns(uint256[])
func (_FeeStorageImpl *FeeStorageImplSession) GetFeeDistributionAmounts(opType *big.Int) ([]*big.Int, error) {
	return _FeeStorageImpl.Contract.GetFeeDistributionAmounts(&_FeeStorageImpl.CallOpts, opType)
}

// GetFeeDistributionAmounts is a free data retrieval call binding the contract method 0x757d77a4.
//
// Solidity: function getFeeDistributionAmounts(uint256 opType) view returns(uint256[])
func (_FeeStorageImpl *FeeStorageImplCallerSession) GetFeeDistributionAmounts(opType *big.Int) ([]*big.Int, error) {
	return _FeeStorageImpl.Contract.GetFeeDistributionAmounts(&_FeeStorageImpl.CallOpts, opType)
}

// GetFeeValueSettings is a free data retrieval call binding the contract method 0x2a674186.
//
// Solidity: function getFeeValueSettings() view returns(uint256[], uint256[])
func (_FeeStorageImpl *FeeStorageImplCaller) GetFeeValueSettings(opts *bind.CallOpts) ([]*big.Int, []*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
		ret1 = new([]*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _FeeStorageImpl.contract.Call(opts, out, "getFeeValueSettings")
	return *ret0, *ret1, err
}

// GetFeeValueSettings is a free data retrieval call binding the contract method 0x2a674186.
//
// Solidity: function getFeeValueSettings() view returns(uint256[], uint256[])
func (_FeeStorageImpl *FeeStorageImplSession) GetFeeValueSettings() ([]*big.Int, []*big.Int, error) {
	return _FeeStorageImpl.Contract.GetFeeValueSettings(&_FeeStorageImpl.CallOpts)
}

// GetFeeValueSettings is a free data retrieval call binding the contract method 0x2a674186.
//
// Solidity: function getFeeValueSettings() view returns(uint256[], uint256[])
func (_FeeStorageImpl *FeeStorageImplCallerSession) GetFeeValueSettings() ([]*big.Int, []*big.Int, error) {
	return _FeeStorageImpl.Contract.GetFeeValueSettings(&_FeeStorageImpl.CallOpts)
}

// GetIsExceptional is a free data retrieval call binding the contract method 0x01a9255b.
//
// Solidity: function getIsExceptional(address account) view returns(bool)
func (_FeeStorageImpl *FeeStorageImplCaller) GetIsExceptional(opts *bind.CallOpts, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _FeeStorageImpl.contract.Call(opts, out, "getIsExceptional", account)
	return *ret0, err
}

// GetIsExceptional is a free data retrieval call binding the contract method 0x01a9255b.
//
// Solidity: function getIsExceptional(address account) view returns(bool)
func (_FeeStorageImpl *FeeStorageImplSession) GetIsExceptional(account common.Address) (bool, error) {
	return _FeeStorageImpl.Contract.GetIsExceptional(&_FeeStorageImpl.CallOpts, account)
}

// GetIsExceptional is a free data retrieval call binding the contract method 0x01a9255b.
//
// Solidity: function getIsExceptional(address account) view returns(bool)
func (_FeeStorageImpl *FeeStorageImplCallerSession) GetIsExceptional(account common.Address) (bool, error) {
	return _FeeStorageImpl.Contract.GetIsExceptional(&_FeeStorageImpl.CallOpts, account)
}

// GetThisContractIndex is a free data retrieval call binding the contract method 0x70c6d367.
//
// Solidity: function getThisContractIndex() pure returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCaller) GetThisContractIndex(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeStorageImpl.contract.Call(opts, out, "getThisContractIndex")
	return *ret0, err
}

// GetThisContractIndex is a free data retrieval call binding the contract method 0x70c6d367.
//
// Solidity: function getThisContractIndex() pure returns(uint256)
func (_FeeStorageImpl *FeeStorageImplSession) GetThisContractIndex() (*big.Int, error) {
	return _FeeStorageImpl.Contract.GetThisContractIndex(&_FeeStorageImpl.CallOpts)
}

// GetThisContractIndex is a free data retrieval call binding the contract method 0x70c6d367.
//
// Solidity: function getThisContractIndex() pure returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCallerSession) GetThisContractIndex() (*big.Int, error) {
	return _FeeStorageImpl.Contract.GetThisContractIndex(&_FeeStorageImpl.CallOpts)
}

// OpTypeToFeeDistribution is a free data retrieval call binding the contract method 0xec4e3fbd.
//
// Solidity: function opTypeToFeeDistribution(uint256 , uint256 ) view returns(address)
func (_FeeStorageImpl *FeeStorageImplCaller) OpTypeToFeeDistribution(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _FeeStorageImpl.contract.Call(opts, out, "opTypeToFeeDistribution", arg0, arg1)
	return *ret0, err
}

// OpTypeToFeeDistribution is a free data retrieval call binding the contract method 0xec4e3fbd.
//
// Solidity: function opTypeToFeeDistribution(uint256 , uint256 ) view returns(address)
func (_FeeStorageImpl *FeeStorageImplSession) OpTypeToFeeDistribution(arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	return _FeeStorageImpl.Contract.OpTypeToFeeDistribution(&_FeeStorageImpl.CallOpts, arg0, arg1)
}

// OpTypeToFeeDistribution is a free data retrieval call binding the contract method 0xec4e3fbd.
//
// Solidity: function opTypeToFeeDistribution(uint256 , uint256 ) view returns(address)
func (_FeeStorageImpl *FeeStorageImplCallerSession) OpTypeToFeeDistribution(arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	return _FeeStorageImpl.Contract.OpTypeToFeeDistribution(&_FeeStorageImpl.CallOpts, arg0, arg1)
}

// OpTypeToFeeDistributionAmounts is a free data retrieval call binding the contract method 0x87e39654.
//
// Solidity: function opTypeToFeeDistributionAmounts(uint256 , uint256 ) view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCaller) OpTypeToFeeDistributionAmounts(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeStorageImpl.contract.Call(opts, out, "opTypeToFeeDistributionAmounts", arg0, arg1)
	return *ret0, err
}

// OpTypeToFeeDistributionAmounts is a free data retrieval call binding the contract method 0x87e39654.
//
// Solidity: function opTypeToFeeDistributionAmounts(uint256 , uint256 ) view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplSession) OpTypeToFeeDistributionAmounts(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _FeeStorageImpl.Contract.OpTypeToFeeDistributionAmounts(&_FeeStorageImpl.CallOpts, arg0, arg1)
}

// OpTypeToFeeDistributionAmounts is a free data retrieval call binding the contract method 0x87e39654.
//
// Solidity: function opTypeToFeeDistributionAmounts(uint256 , uint256 ) view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCallerSession) OpTypeToFeeDistributionAmounts(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _FeeStorageImpl.Contract.OpTypeToFeeDistributionAmounts(&_FeeStorageImpl.CallOpts, arg0, arg1)
}

// OpTypeToFeeValueIndex is a free data retrieval call binding the contract method 0x8669a05c.
//
// Solidity: function opTypeToFeeValueIndex(uint256 ) view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCaller) OpTypeToFeeValueIndex(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeStorageImpl.contract.Call(opts, out, "opTypeToFeeValueIndex", arg0)
	return *ret0, err
}

// OpTypeToFeeValueIndex is a free data retrieval call binding the contract method 0x8669a05c.
//
// Solidity: function opTypeToFeeValueIndex(uint256 ) view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplSession) OpTypeToFeeValueIndex(arg0 *big.Int) (*big.Int, error) {
	return _FeeStorageImpl.Contract.OpTypeToFeeValueIndex(&_FeeStorageImpl.CallOpts, arg0)
}

// OpTypeToFeeValueIndex is a free data retrieval call binding the contract method 0x8669a05c.
//
// Solidity: function opTypeToFeeValueIndex(uint256 ) view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCallerSession) OpTypeToFeeValueIndex(arg0 *big.Int) (*big.Int, error) {
	return _FeeStorageImpl.Contract.OpTypeToFeeValueIndex(&_FeeStorageImpl.CallOpts, arg0)
}

// RegisteredFeeValues is a free data retrieval call binding the contract method 0xf6fb088d.
//
// Solidity: function registeredFeeValues(uint256 ) view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCaller) RegisteredFeeValues(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeStorageImpl.contract.Call(opts, out, "registeredFeeValues", arg0)
	return *ret0, err
}

// RegisteredFeeValues is a free data retrieval call binding the contract method 0xf6fb088d.
//
// Solidity: function registeredFeeValues(uint256 ) view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplSession) RegisteredFeeValues(arg0 *big.Int) (*big.Int, error) {
	return _FeeStorageImpl.Contract.RegisteredFeeValues(&_FeeStorageImpl.CallOpts, arg0)
}

// RegisteredFeeValues is a free data retrieval call binding the contract method 0xf6fb088d.
//
// Solidity: function registeredFeeValues(uint256 ) view returns(uint256)
func (_FeeStorageImpl *FeeStorageImplCallerSession) RegisteredFeeValues(arg0 *big.Int) (*big.Int, error) {
	return _FeeStorageImpl.Contract.RegisteredFeeValues(&_FeeStorageImpl.CallOpts, arg0)
}

// ClearExceptionAccountsMultisig is a paid mutator transaction binding the contract method 0xffc3c15b.
//
// Solidity: function clearExceptionAccountsMultisig(uint256 expiresOn) returns(bool)
func (_FeeStorageImpl *FeeStorageImplTransactor) ClearExceptionAccountsMultisig(opts *bind.TransactOpts, expiresOn *big.Int) (*types.Transaction, error) {
	return _FeeStorageImpl.contract.Transact(opts, "clearExceptionAccountsMultisig", expiresOn)
}

// ClearExceptionAccountsMultisig is a paid mutator transaction binding the contract method 0xffc3c15b.
//
// Solidity: function clearExceptionAccountsMultisig(uint256 expiresOn) returns(bool)
func (_FeeStorageImpl *FeeStorageImplSession) ClearExceptionAccountsMultisig(expiresOn *big.Int) (*types.Transaction, error) {
	return _FeeStorageImpl.Contract.ClearExceptionAccountsMultisig(&_FeeStorageImpl.TransactOpts, expiresOn)
}

// ClearExceptionAccountsMultisig is a paid mutator transaction binding the contract method 0xffc3c15b.
//
// Solidity: function clearExceptionAccountsMultisig(uint256 expiresOn) returns(bool)
func (_FeeStorageImpl *FeeStorageImplTransactorSession) ClearExceptionAccountsMultisig(expiresOn *big.Int) (*types.Transaction, error) {
	return _FeeStorageImpl.Contract.ClearExceptionAccountsMultisig(&_FeeStorageImpl.TransactOpts, expiresOn)
}

// SetExceptionAccountsMultisig is a paid mutator transaction binding the contract method 0x39cfa441.
//
// Solidity: function setExceptionAccountsMultisig(address[] accounts, uint256 expiresOn) returns(bool)
func (_FeeStorageImpl *FeeStorageImplTransactor) SetExceptionAccountsMultisig(opts *bind.TransactOpts, accounts []common.Address, expiresOn *big.Int) (*types.Transaction, error) {
	return _FeeStorageImpl.contract.Transact(opts, "setExceptionAccountsMultisig", accounts, expiresOn)
}

// SetExceptionAccountsMultisig is a paid mutator transaction binding the contract method 0x39cfa441.
//
// Solidity: function setExceptionAccountsMultisig(address[] accounts, uint256 expiresOn) returns(bool)
func (_FeeStorageImpl *FeeStorageImplSession) SetExceptionAccountsMultisig(accounts []common.Address, expiresOn *big.Int) (*types.Transaction, error) {
	return _FeeStorageImpl.Contract.SetExceptionAccountsMultisig(&_FeeStorageImpl.TransactOpts, accounts, expiresOn)
}

// SetExceptionAccountsMultisig is a paid mutator transaction binding the contract method 0x39cfa441.
//
// Solidity: function setExceptionAccountsMultisig(address[] accounts, uint256 expiresOn) returns(bool)
func (_FeeStorageImpl *FeeStorageImplTransactorSession) SetExceptionAccountsMultisig(accounts []common.Address, expiresOn *big.Int) (*types.Transaction, error) {
	return _FeeStorageImpl.Contract.SetExceptionAccountsMultisig(&_FeeStorageImpl.TransactOpts, accounts, expiresOn)
}

// SetFeeDistributionMultisig is a paid mutator transaction binding the contract method 0xc39bdc64.
//
// Solidity: function setFeeDistributionMultisig(uint256 opType, address[] feeDistributionAccounts, uint256[] feeDistributionAmounts, uint256 expiresOn) returns(bool)
func (_FeeStorageImpl *FeeStorageImplTransactor) SetFeeDistributionMultisig(opts *bind.TransactOpts, opType *big.Int, feeDistributionAccounts []common.Address, feeDistributionAmounts []*big.Int, expiresOn *big.Int) (*types.Transaction, error) {
	return _FeeStorageImpl.contract.Transact(opts, "setFeeDistributionMultisig", opType, feeDistributionAccounts, feeDistributionAmounts, expiresOn)
}

// SetFeeDistributionMultisig is a paid mutator transaction binding the contract method 0xc39bdc64.
//
// Solidity: function setFeeDistributionMultisig(uint256 opType, address[] feeDistributionAccounts, uint256[] feeDistributionAmounts, uint256 expiresOn) returns(bool)
func (_FeeStorageImpl *FeeStorageImplSession) SetFeeDistributionMultisig(opType *big.Int, feeDistributionAccounts []common.Address, feeDistributionAmounts []*big.Int, expiresOn *big.Int) (*types.Transaction, error) {
	return _FeeStorageImpl.Contract.SetFeeDistributionMultisig(&_FeeStorageImpl.TransactOpts, opType, feeDistributionAccounts, feeDistributionAmounts, expiresOn)
}

// SetFeeDistributionMultisig is a paid mutator transaction binding the contract method 0xc39bdc64.
//
// Solidity: function setFeeDistributionMultisig(uint256 opType, address[] feeDistributionAccounts, uint256[] feeDistributionAmounts, uint256 expiresOn) returns(bool)
func (_FeeStorageImpl *FeeStorageImplTransactorSession) SetFeeDistributionMultisig(opType *big.Int, feeDistributionAccounts []common.Address, feeDistributionAmounts []*big.Int, expiresOn *big.Int) (*types.Transaction, error) {
	return _FeeStorageImpl.Contract.SetFeeDistributionMultisig(&_FeeStorageImpl.TransactOpts, opType, feeDistributionAccounts, feeDistributionAmounts, expiresOn)
}

// SetMultisig is a paid mutator transaction binding the contract method 0xccd809e8.
//
// Solidity: function setMultisig(uint256 amount, uint256 feeType, uint256 expiresOn) returns(bool)
func (_FeeStorageImpl *FeeStorageImplTransactor) SetMultisig(opts *bind.TransactOpts, amount *big.Int, feeType *big.Int, expiresOn *big.Int) (*types.Transaction, error) {
	return _FeeStorageImpl.contract.Transact(opts, "setMultisig", amount, feeType, expiresOn)
}

// SetMultisig is a paid mutator transaction binding the contract method 0xccd809e8.
//
// Solidity: function setMultisig(uint256 amount, uint256 feeType, uint256 expiresOn) returns(bool)
func (_FeeStorageImpl *FeeStorageImplSession) SetMultisig(amount *big.Int, feeType *big.Int, expiresOn *big.Int) (*types.Transaction, error) {
	return _FeeStorageImpl.Contract.SetMultisig(&_FeeStorageImpl.TransactOpts, amount, feeType, expiresOn)
}

// SetMultisig is a paid mutator transaction binding the contract method 0xccd809e8.
//
// Solidity: function setMultisig(uint256 amount, uint256 feeType, uint256 expiresOn) returns(bool)
func (_FeeStorageImpl *FeeStorageImplTransactorSession) SetMultisig(amount *big.Int, feeType *big.Int, expiresOn *big.Int) (*types.Transaction, error) {
	return _FeeStorageImpl.Contract.SetMultisig(&_FeeStorageImpl.TransactOpts, amount, feeType, expiresOn)
}

// FeeStorageImplExceptionalAccountsChangedIterator is returned from FilterExceptionalAccountsChanged and is used to iterate over the raw logs and unpacked data for ExceptionalAccountsChanged events raised by the FeeStorageImpl contract.
type FeeStorageImplExceptionalAccountsChangedIterator struct {
	Event *FeeStorageImplExceptionalAccountsChanged // Event containing the contract specifics and raw log

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
func (it *FeeStorageImplExceptionalAccountsChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeStorageImplExceptionalAccountsChanged)
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
		it.Event = new(FeeStorageImplExceptionalAccountsChanged)
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
func (it *FeeStorageImplExceptionalAccountsChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeStorageImplExceptionalAccountsChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeStorageImplExceptionalAccountsChanged represents a ExceptionalAccountsChanged event raised by the FeeStorageImpl contract.
type FeeStorageImplExceptionalAccountsChanged struct {
	Accounts []common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterExceptionalAccountsChanged is a free log retrieval operation binding the contract event 0x611f5aa5c7f095ba927124ccc58a8b565c42d284c5d2d89b3ee1905ff8cd698c.
//
// Solidity: event ExceptionalAccountsChanged(address[] accounts)
func (_FeeStorageImpl *FeeStorageImplFilterer) FilterExceptionalAccountsChanged(opts *bind.FilterOpts) (*FeeStorageImplExceptionalAccountsChangedIterator, error) {

	logs, sub, err := _FeeStorageImpl.contract.FilterLogs(opts, "ExceptionalAccountsChanged")
	if err != nil {
		return nil, err
	}
	return &FeeStorageImplExceptionalAccountsChangedIterator{contract: _FeeStorageImpl.contract, event: "ExceptionalAccountsChanged", logs: logs, sub: sub}, nil
}

// WatchExceptionalAccountsChanged is a free log subscription operation binding the contract event 0x611f5aa5c7f095ba927124ccc58a8b565c42d284c5d2d89b3ee1905ff8cd698c.
//
// Solidity: event ExceptionalAccountsChanged(address[] accounts)
func (_FeeStorageImpl *FeeStorageImplFilterer) WatchExceptionalAccountsChanged(opts *bind.WatchOpts, sink chan<- *FeeStorageImplExceptionalAccountsChanged) (event.Subscription, error) {

	logs, sub, err := _FeeStorageImpl.contract.WatchLogs(opts, "ExceptionalAccountsChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeStorageImplExceptionalAccountsChanged)
				if err := _FeeStorageImpl.contract.UnpackLog(event, "ExceptionalAccountsChanged", log); err != nil {
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

// ParseExceptionalAccountsChanged is a log parse operation binding the contract event 0x611f5aa5c7f095ba927124ccc58a8b565c42d284c5d2d89b3ee1905ff8cd698c.
//
// Solidity: event ExceptionalAccountsChanged(address[] accounts)
func (_FeeStorageImpl *FeeStorageImplFilterer) ParseExceptionalAccountsChanged(log types.Log) (*FeeStorageImplExceptionalAccountsChanged, error) {
	event := new(FeeStorageImplExceptionalAccountsChanged)
	if err := _FeeStorageImpl.contract.UnpackLog(event, "ExceptionalAccountsChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}

// FeeStorageImplFeeDistributionChangedIterator is returned from FilterFeeDistributionChanged and is used to iterate over the raw logs and unpacked data for FeeDistributionChanged events raised by the FeeStorageImpl contract.
type FeeStorageImplFeeDistributionChangedIterator struct {
	Event *FeeStorageImplFeeDistributionChanged // Event containing the contract specifics and raw log

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
func (it *FeeStorageImplFeeDistributionChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeStorageImplFeeDistributionChanged)
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
		it.Event = new(FeeStorageImplFeeDistributionChanged)
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
func (it *FeeStorageImplFeeDistributionChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeStorageImplFeeDistributionChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeStorageImplFeeDistributionChanged represents a FeeDistributionChanged event raised by the FeeStorageImpl contract.
type FeeStorageImplFeeDistributionChanged struct {
	OpType               *big.Int
	DistributionAccounts []common.Address
	DistributionAmounts  []*big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterFeeDistributionChanged is a free log retrieval operation binding the contract event 0xe6181c29f8c547e7f73c9d71a0bee0d9e5b60598064776fa7b7c858bdb5c41cf.
//
// Solidity: event FeeDistributionChanged(uint256 opType, address[] distributionAccounts, uint256[] distributionAmounts)
func (_FeeStorageImpl *FeeStorageImplFilterer) FilterFeeDistributionChanged(opts *bind.FilterOpts) (*FeeStorageImplFeeDistributionChangedIterator, error) {

	logs, sub, err := _FeeStorageImpl.contract.FilterLogs(opts, "FeeDistributionChanged")
	if err != nil {
		return nil, err
	}
	return &FeeStorageImplFeeDistributionChangedIterator{contract: _FeeStorageImpl.contract, event: "FeeDistributionChanged", logs: logs, sub: sub}, nil
}

// WatchFeeDistributionChanged is a free log subscription operation binding the contract event 0xe6181c29f8c547e7f73c9d71a0bee0d9e5b60598064776fa7b7c858bdb5c41cf.
//
// Solidity: event FeeDistributionChanged(uint256 opType, address[] distributionAccounts, uint256[] distributionAmounts)
func (_FeeStorageImpl *FeeStorageImplFilterer) WatchFeeDistributionChanged(opts *bind.WatchOpts, sink chan<- *FeeStorageImplFeeDistributionChanged) (event.Subscription, error) {

	logs, sub, err := _FeeStorageImpl.contract.WatchLogs(opts, "FeeDistributionChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeStorageImplFeeDistributionChanged)
				if err := _FeeStorageImpl.contract.UnpackLog(event, "FeeDistributionChanged", log); err != nil {
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

// ParseFeeDistributionChanged is a log parse operation binding the contract event 0xe6181c29f8c547e7f73c9d71a0bee0d9e5b60598064776fa7b7c858bdb5c41cf.
//
// Solidity: event FeeDistributionChanged(uint256 opType, address[] distributionAccounts, uint256[] distributionAmounts)
func (_FeeStorageImpl *FeeStorageImplFilterer) ParseFeeDistributionChanged(log types.Log) (*FeeStorageImplFeeDistributionChanged, error) {
	event := new(FeeStorageImplFeeDistributionChanged)
	if err := _FeeStorageImpl.contract.UnpackLog(event, "FeeDistributionChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}

// FeeStorageImplFeeValueChangedIterator is returned from FilterFeeValueChanged and is used to iterate over the raw logs and unpacked data for FeeValueChanged events raised by the FeeStorageImpl contract.
type FeeStorageImplFeeValueChangedIterator struct {
	Event *FeeStorageImplFeeValueChanged // Event containing the contract specifics and raw log

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
func (it *FeeStorageImplFeeValueChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeStorageImplFeeValueChanged)
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
		it.Event = new(FeeStorageImplFeeValueChanged)
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
func (it *FeeStorageImplFeeValueChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeStorageImplFeeValueChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeStorageImplFeeValueChanged represents a FeeValueChanged event raised by the FeeStorageImpl contract.
type FeeStorageImplFeeValueChanged struct {
	OpType *big.Int
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFeeValueChanged is a free log retrieval operation binding the contract event 0xa2e24a12c2ff86d8a4117735073797adfa6d262f904f32179993d2834f423132.
//
// Solidity: event FeeValueChanged(uint256 opType, uint256 value)
func (_FeeStorageImpl *FeeStorageImplFilterer) FilterFeeValueChanged(opts *bind.FilterOpts) (*FeeStorageImplFeeValueChangedIterator, error) {

	logs, sub, err := _FeeStorageImpl.contract.FilterLogs(opts, "FeeValueChanged")
	if err != nil {
		return nil, err
	}
	return &FeeStorageImplFeeValueChangedIterator{contract: _FeeStorageImpl.contract, event: "FeeValueChanged", logs: logs, sub: sub}, nil
}

// WatchFeeValueChanged is a free log subscription operation binding the contract event 0xa2e24a12c2ff86d8a4117735073797adfa6d262f904f32179993d2834f423132.
//
// Solidity: event FeeValueChanged(uint256 opType, uint256 value)
func (_FeeStorageImpl *FeeStorageImplFilterer) WatchFeeValueChanged(opts *bind.WatchOpts, sink chan<- *FeeStorageImplFeeValueChanged) (event.Subscription, error) {

	logs, sub, err := _FeeStorageImpl.contract.WatchLogs(opts, "FeeValueChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeStorageImplFeeValueChanged)
				if err := _FeeStorageImpl.contract.UnpackLog(event, "FeeValueChanged", log); err != nil {
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

// ParseFeeValueChanged is a log parse operation binding the contract event 0xa2e24a12c2ff86d8a4117735073797adfa6d262f904f32179993d2834f423132.
//
// Solidity: event FeeValueChanged(uint256 opType, uint256 value)
func (_FeeStorageImpl *FeeStorageImplFilterer) ParseFeeValueChanged(log types.Log) (*FeeStorageImplFeeValueChanged, error) {
	event := new(FeeStorageImplFeeValueChanged)
	if err := _FeeStorageImpl.contract.UnpackLog(event, "FeeValueChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}
