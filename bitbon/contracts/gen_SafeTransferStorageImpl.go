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

// SafeTransferStorageImplABI is the input ABI used to generate the binding from.
const SafeTransferStorageImplABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractStorageAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"transferId\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"source\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"dest\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expiresAt\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxRetries\",\"type\":\"uint256\"}],\"name\":\"TransferAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"transferId\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"canceller\",\"type\":\"address\"}],\"name\":\"TransferCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"transferId\",\"type\":\"bytes\"}],\"name\":\"TransferExpired\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"transferIdHash\",\"type\":\"bytes32\"}],\"name\":\"TransferExpiredByHash\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ACCESS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ACCESS_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ADMINABLE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ADMIN_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ASSETBOX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ASSETBOX_INFO\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ASSETBOX_INFO_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ASSETBOX_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_BALANCE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_BITBON\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_BITBON_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_BITBON_SUPPORT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_BONBALANCE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_BONTRANSFER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_DEX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_DISTRIBUTION_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_EXCHANGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_EXCHANGE_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_FEE_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_GENERATOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MINING_AGENT_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MSBON\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MSBON_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_ADD_ADDRESS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_ADD_ADMIN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_ADD_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_ADD_ROLE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_DROP_PERMISSION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_EDIT_ADMIN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_EDIT_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_EDIT_ROLE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REMOVE_ADMIN_BY_ADDRESS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REMOVE_ADMIN_BY_INDEX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REMOVE_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REMOVE_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REMOVE_ROLE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REPLENISH_CAPITALIZATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REPLENISH_CERTIFICATE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_REPLENISH_EMISSION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_SET_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_SET_PERMISSION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_MULTISIG_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_OTC\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_OTC_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_REPLENISH_CAPITALIZATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_RESERVED_ALIAS_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_ROLE_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_SAFE_TRANSFER_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_TRANSFER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_TRANSFER_FROM_CAPITALIZATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_ACCESS_RESTORATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_ADMIN_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_CONTRACTS_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_CONTRACT_DEPLOY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_EMISSION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_FEES_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_MULTISIG_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMISSION_ROLES_STORAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_ACCESS_VERIFIER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_BITBON_ISSUE_VERIFIER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_COMMISSION_VERIFIER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_DEPLOY_ADMIN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_PERMISSION_ADMIN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"contractStorage\",\"outputs\":[{\"internalType\":\"contractContractStorage\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"oldestPending\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"transferIdOrig\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"transferIds\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"transfers\",\"outputs\":[{\"internalType\":\"enumDtoAware.BitbonTransferState\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"transferIdHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"source\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dest\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"vk\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiresAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentRetries\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxRetries\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"canceller\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"transfersFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getThisContractIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"transferIdHash\",\"type\":\"bytes32\"}],\"name\":\"getTransferIdByHash\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"transferId\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"source\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dest\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"vk\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"expiresAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxRetries\",\"type\":\"uint256\"}],\"name\":\"addSafeTransfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"transferId\",\"type\":\"bytes\"}],\"name\":\"cancelTransfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"transferId\",\"type\":\"bytes\"}],\"name\":\"expireTransfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"transferIdHash\",\"type\":\"bytes32\"}],\"name\":\"expireTransferByHash\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"transferId\",\"type\":\"bytes\"}],\"name\":\"isPending\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"transferId\",\"type\":\"bytes\"}],\"name\":\"isSuccessful\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"transferId\",\"type\":\"bytes\"}],\"name\":\"isCancelled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"transferId\",\"type\":\"bytes\"}],\"name\":\"isExpired\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"transferId\",\"type\":\"bytes\"}],\"name\":\"getTransfer\",\"outputs\":[{\"internalType\":\"enumDtoAware.BitbonTransferState\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"transferIdHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"source\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dest\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"vk\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiresAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentRetries\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxRetries\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"canceller\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"transferIdHashString\",\"type\":\"bytes32\"}],\"name\":\"getTransferByHash\",\"outputs\":[{\"internalType\":\"enumDtoAware.BitbonTransferState\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"transferIdHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"source\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dest\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"vk\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiresAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentRetries\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxRetries\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"canceller\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"transferId\",\"type\":\"bytes\"}],\"name\":\"setFailedState\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"transferId\",\"type\":\"bytes\"}],\"name\":\"setExecutedState\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"transferId\",\"type\":\"bytes\"}],\"name\":\"getTransferSource\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"transferIdHash\",\"type\":\"bytes32\"}],\"name\":\"getTransferSourceByHash\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"transferId\",\"type\":\"bytes\"}],\"name\":\"getTransferDest\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"transferId\",\"type\":\"bytes\"}],\"name\":\"getTransferAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"transferIdHash\",\"type\":\"bytes32\"}],\"name\":\"getTransferAmountByHash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"transferId\",\"type\":\"bytes\"}],\"name\":\"transferExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"transferId\",\"type\":\"bytes\"}],\"name\":\"getState\",\"outputs\":[{\"internalType\":\"enumDtoAware.BitbonTransferState\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"transferId\",\"type\":\"bytes\"}],\"name\":\"getProtectionCode\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"transferId\",\"type\":\"bytes\"}],\"name\":\"checkRetries\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"transferId\",\"type\":\"bytes\"}],\"name\":\"getRetries\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"transferId\",\"type\":\"bytes\"}],\"name\":\"getFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTransfers\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTransferLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"transferIdHash\",\"type\":\"bytes32\"}],\"name\":\"isPendingByHash\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"transferIdHash\",\"type\":\"bytes32\"}],\"name\":\"isExpiredByHash\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"first\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"last\",\"type\":\"uint256\"}],\"name\":\"searchOldestPending\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isPresent\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"first\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"last\",\"type\":\"uint256\"}],\"name\":\"getExpiredTransfers\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"oldestPendingIndex\",\"type\":\"uint256\"}],\"name\":\"setOldestPending\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOldestPending\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// SafeTransferStorageImpl is an auto generated Go binding around an Ethereum contract.
type SafeTransferStorageImpl struct {
	SafeTransferStorageImplCaller     // Read-only binding to the contract
	SafeTransferStorageImplTransactor // Write-only binding to the contract
	SafeTransferStorageImplFilterer   // Log filterer for contract events
}

// SafeTransferStorageImplCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeTransferStorageImplCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeTransferStorageImplTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeTransferStorageImplTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeTransferStorageImplFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeTransferStorageImplFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeTransferStorageImplSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeTransferStorageImplSession struct {
	Contract     *SafeTransferStorageImpl // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// SafeTransferStorageImplCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeTransferStorageImplCallerSession struct {
	Contract *SafeTransferStorageImplCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// SafeTransferStorageImplTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeTransferStorageImplTransactorSession struct {
	Contract     *SafeTransferStorageImplTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// SafeTransferStorageImplRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeTransferStorageImplRaw struct {
	Contract *SafeTransferStorageImpl // Generic contract binding to access the raw methods on
}

// SafeTransferStorageImplCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeTransferStorageImplCallerRaw struct {
	Contract *SafeTransferStorageImplCaller // Generic read-only contract binding to access the raw methods on
}

// SafeTransferStorageImplTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeTransferStorageImplTransactorRaw struct {
	Contract *SafeTransferStorageImplTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeTransferStorageImpl creates a new instance of SafeTransferStorageImpl, bound to a specific deployed contract.
func NewSafeTransferStorageImpl(address common.Address, backend bind.ContractBackend) (*SafeTransferStorageImpl, error) {
	contract, err := bindSafeTransferStorageImpl(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeTransferStorageImpl{SafeTransferStorageImplCaller: SafeTransferStorageImplCaller{contract: contract}, SafeTransferStorageImplTransactor: SafeTransferStorageImplTransactor{contract: contract}, SafeTransferStorageImplFilterer: SafeTransferStorageImplFilterer{contract: contract}}, nil
}

// NewSafeTransferStorageImplCaller creates a new read-only instance of SafeTransferStorageImpl, bound to a specific deployed contract.
func NewSafeTransferStorageImplCaller(address common.Address, caller bind.ContractCaller) (*SafeTransferStorageImplCaller, error) {
	contract, err := bindSafeTransferStorageImpl(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeTransferStorageImplCaller{contract: contract}, nil
}

// NewSafeTransferStorageImplTransactor creates a new write-only instance of SafeTransferStorageImpl, bound to a specific deployed contract.
func NewSafeTransferStorageImplTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeTransferStorageImplTransactor, error) {
	contract, err := bindSafeTransferStorageImpl(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeTransferStorageImplTransactor{contract: contract}, nil
}

// NewSafeTransferStorageImplFilterer creates a new log filterer instance of SafeTransferStorageImpl, bound to a specific deployed contract.
func NewSafeTransferStorageImplFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeTransferStorageImplFilterer, error) {
	contract, err := bindSafeTransferStorageImpl(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeTransferStorageImplFilterer{contract: contract}, nil
}

// bindSafeTransferStorageImpl binds a generic wrapper to an already deployed contract.
func bindSafeTransferStorageImpl(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeTransferStorageImplABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeTransferStorageImpl *SafeTransferStorageImplRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeTransferStorageImpl.Contract.SafeTransferStorageImplCaller.contract.Call(opts, result, method, params...)
}

// Return Abi for SafeTransferStorageImplCaller
func (_SafeTransferStorageImpl *SafeTransferStorageImplCaller) ABI() abi.ABI {
	return _SafeTransferStorageImpl.contract.Abi
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeTransferStorageImpl *SafeTransferStorageImplRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeTransferStorageImpl.Contract.SafeTransferStorageImplTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeTransferStorageImpl *SafeTransferStorageImplRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeTransferStorageImpl.Contract.SafeTransferStorageImplTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeTransferStorageImpl *SafeTransferStorageImplCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeTransferStorageImpl.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeTransferStorageImpl *SafeTransferStorageImplTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeTransferStorageImpl.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeTransferStorageImpl *SafeTransferStorageImplTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeTransferStorageImpl.Contract.contract.Transact(opts, method, params...)
}

// CONTRACTACCESS is a free data retrieval call binding the contract method 0x0ac62b0c.
//
// Solidity: function CONTRACT_ACCESS() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCaller) CONTRACTACCESS(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SafeTransferStorageImpl.contract.Call(opts, out, "CONTRACT_ACCESS")
	return *ret0, err
}

// CONTRACTACCESS is a free data retrieval call binding the contract method 0x0ac62b0c.
//
// Solidity: function CONTRACT_ACCESS() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplSession) CONTRACTACCESS() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.CONTRACTACCESS(&_SafeTransferStorageImpl.CallOpts)
}

// CONTRACTACCESS is a free data retrieval call binding the contract method 0x0ac62b0c.
//
// Solidity: function CONTRACT_ACCESS() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCallerSession) CONTRACTACCESS() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.CONTRACTACCESS(&_SafeTransferStorageImpl.CallOpts)
}

// CONTRACTACCESSSTORAGE is a free data retrieval call binding the contract method 0xceac9bd5.
//
// Solidity: function CONTRACT_ACCESS_STORAGE() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCaller) CONTRACTACCESSSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SafeTransferStorageImpl.contract.Call(opts, out, "CONTRACT_ACCESS_STORAGE")
	return *ret0, err
}

// CONTRACTACCESSSTORAGE is a free data retrieval call binding the contract method 0xceac9bd5.
//
// Solidity: function CONTRACT_ACCESS_STORAGE() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplSession) CONTRACTACCESSSTORAGE() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.CONTRACTACCESSSTORAGE(&_SafeTransferStorageImpl.CallOpts)
}

// CONTRACTACCESSSTORAGE is a free data retrieval call binding the contract method 0xceac9bd5.
//
// Solidity: function CONTRACT_ACCESS_STORAGE() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCallerSession) CONTRACTACCESSSTORAGE() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.CONTRACTACCESSSTORAGE(&_SafeTransferStorageImpl.CallOpts)
}

// CONTRACTADMINABLE is a free data retrieval call binding the contract method 0xbbe3a78f.
//
// Solidity: function CONTRACT_ADMINABLE() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCaller) CONTRACTADMINABLE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SafeTransferStorageImpl.contract.Call(opts, out, "CONTRACT_ADMINABLE")
	return *ret0, err
}

// CONTRACTADMINABLE is a free data retrieval call binding the contract method 0xbbe3a78f.
//
// Solidity: function CONTRACT_ADMINABLE() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplSession) CONTRACTADMINABLE() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.CONTRACTADMINABLE(&_SafeTransferStorageImpl.CallOpts)
}

// CONTRACTADMINABLE is a free data retrieval call binding the contract method 0xbbe3a78f.
//
// Solidity: function CONTRACT_ADMINABLE() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCallerSession) CONTRACTADMINABLE() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.CONTRACTADMINABLE(&_SafeTransferStorageImpl.CallOpts)
}

// CONTRACTADMINSTORAGE is a free data retrieval call binding the contract method 0xc9aebae9.
//
// Solidity: function CONTRACT_ADMIN_STORAGE() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCaller) CONTRACTADMINSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SafeTransferStorageImpl.contract.Call(opts, out, "CONTRACT_ADMIN_STORAGE")
	return *ret0, err
}

// CONTRACTADMINSTORAGE is a free data retrieval call binding the contract method 0xc9aebae9.
//
// Solidity: function CONTRACT_ADMIN_STORAGE() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplSession) CONTRACTADMINSTORAGE() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.CONTRACTADMINSTORAGE(&_SafeTransferStorageImpl.CallOpts)
}

// CONTRACTADMINSTORAGE is a free data retrieval call binding the contract method 0xc9aebae9.
//
// Solidity: function CONTRACT_ADMIN_STORAGE() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCallerSession) CONTRACTADMINSTORAGE() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.CONTRACTADMINSTORAGE(&_SafeTransferStorageImpl.CallOpts)
}

// CONTRACTASSETBOX is a free data retrieval call binding the contract method 0x161aa17c.
//
// Solidity: function CONTRACT_ASSETBOX() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCaller) CONTRACTASSETBOX(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SafeTransferStorageImpl.contract.Call(opts, out, "CONTRACT_ASSETBOX")
	return *ret0, err
}

// CONTRACTASSETBOX is a free data retrieval call binding the contract method 0x161aa17c.
//
// Solidity: function CONTRACT_ASSETBOX() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplSession) CONTRACTASSETBOX() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.CONTRACTASSETBOX(&_SafeTransferStorageImpl.CallOpts)
}

// CONTRACTASSETBOX is a free data retrieval call binding the contract method 0x161aa17c.
//
// Solidity: function CONTRACT_ASSETBOX() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCallerSession) CONTRACTASSETBOX() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.CONTRACTASSETBOX(&_SafeTransferStorageImpl.CallOpts)
}

// CONTRACTASSETBOXINFO is a free data retrieval call binding the contract method 0xb572fd3d.
//
// Solidity: function CONTRACT_ASSETBOX_INFO() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCaller) CONTRACTASSETBOXINFO(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SafeTransferStorageImpl.contract.Call(opts, out, "CONTRACT_ASSETBOX_INFO")
	return *ret0, err
}

// CONTRACTASSETBOXINFO is a free data retrieval call binding the contract method 0xb572fd3d.
//
// Solidity: function CONTRACT_ASSETBOX_INFO() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplSession) CONTRACTASSETBOXINFO() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.CONTRACTASSETBOXINFO(&_SafeTransferStorageImpl.CallOpts)
}

// CONTRACTASSETBOXINFO is a free data retrieval call binding the contract method 0xb572fd3d.
//
// Solidity: function CONTRACT_ASSETBOX_INFO() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCallerSession) CONTRACTASSETBOXINFO() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.CONTRACTASSETBOXINFO(&_SafeTransferStorageImpl.CallOpts)
}

// CONTRACTASSETBOXINFOSTORAGE is a free data retrieval call binding the contract method 0xf32a115d.
//
// Solidity: function CONTRACT_ASSETBOX_INFO_STORAGE() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCaller) CONTRACTASSETBOXINFOSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SafeTransferStorageImpl.contract.Call(opts, out, "CONTRACT_ASSETBOX_INFO_STORAGE")
	return *ret0, err
}

// CONTRACTASSETBOXINFOSTORAGE is a free data retrieval call binding the contract method 0xf32a115d.
//
// Solidity: function CONTRACT_ASSETBOX_INFO_STORAGE() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplSession) CONTRACTASSETBOXINFOSTORAGE() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.CONTRACTASSETBOXINFOSTORAGE(&_SafeTransferStorageImpl.CallOpts)
}

// CONTRACTASSETBOXINFOSTORAGE is a free data retrieval call binding the contract method 0xf32a115d.
//
// Solidity: function CONTRACT_ASSETBOX_INFO_STORAGE() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCallerSession) CONTRACTASSETBOXINFOSTORAGE() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.CONTRACTASSETBOXINFOSTORAGE(&_SafeTransferStorageImpl.CallOpts)
}

// CONTRACTASSETBOXSTORAGE is a free data retrieval call binding the contract method 0xfa5bdafc.
//
// Solidity: function CONTRACT_ASSETBOX_STORAGE() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCaller) CONTRACTASSETBOXSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SafeTransferStorageImpl.contract.Call(opts, out, "CONTRACT_ASSETBOX_STORAGE")
	return *ret0, err
}

// CONTRACTASSETBOXSTORAGE is a free data retrieval call binding the contract method 0xfa5bdafc.
//
// Solidity: function CONTRACT_ASSETBOX_STORAGE() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplSession) CONTRACTASSETBOXSTORAGE() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.CONTRACTASSETBOXSTORAGE(&_SafeTransferStorageImpl.CallOpts)
}

// CONTRACTASSETBOXSTORAGE is a free data retrieval call binding the contract method 0xfa5bdafc.
//
// Solidity: function CONTRACT_ASSETBOX_STORAGE() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCallerSession) CONTRACTASSETBOXSTORAGE() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.CONTRACTASSETBOXSTORAGE(&_SafeTransferStorageImpl.CallOpts)
}

// CONTRACTBALANCE is a free data retrieval call binding the contract method 0xd59045f5.
//
// Solidity: function CONTRACT_BALANCE() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCaller) CONTRACTBALANCE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SafeTransferStorageImpl.contract.Call(opts, out, "CONTRACT_BALANCE")
	return *ret0, err
}

// CONTRACTBALANCE is a free data retrieval call binding the contract method 0xd59045f5.
//
// Solidity: function CONTRACT_BALANCE() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplSession) CONTRACTBALANCE() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.CONTRACTBALANCE(&_SafeTransferStorageImpl.CallOpts)
}

// CONTRACTBALANCE is a free data retrieval call binding the contract method 0xd59045f5.
//
// Solidity: function CONTRACT_BALANCE() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCallerSession) CONTRACTBALANCE() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.CONTRACTBALANCE(&_SafeTransferStorageImpl.CallOpts)
}

// CONTRACTBITBON is a free data retrieval call binding the contract method 0xb6b433ab.
//
// Solidity: function CONTRACT_BITBON() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCaller) CONTRACTBITBON(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SafeTransferStorageImpl.contract.Call(opts, out, "CONTRACT_BITBON")
	return *ret0, err
}

// CONTRACTBITBON is a free data retrieval call binding the contract method 0xb6b433ab.
//
// Solidity: function CONTRACT_BITBON() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplSession) CONTRACTBITBON() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.CONTRACTBITBON(&_SafeTransferStorageImpl.CallOpts)
}

// CONTRACTBITBON is a free data retrieval call binding the contract method 0xb6b433ab.
//
// Solidity: function CONTRACT_BITBON() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCallerSession) CONTRACTBITBON() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.CONTRACTBITBON(&_SafeTransferStorageImpl.CallOpts)
}

// CONTRACTBITBONSTORAGE is a free data retrieval call binding the contract method 0xe7f6c30e.
//
// Solidity: function CONTRACT_BITBON_STORAGE() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCaller) CONTRACTBITBONSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SafeTransferStorageImpl.contract.Call(opts, out, "CONTRACT_BITBON_STORAGE")
	return *ret0, err
}

// CONTRACTBITBONSTORAGE is a free data retrieval call binding the contract method 0xe7f6c30e.
//
// Solidity: function CONTRACT_BITBON_STORAGE() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplSession) CONTRACTBITBONSTORAGE() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.CONTRACTBITBONSTORAGE(&_SafeTransferStorageImpl.CallOpts)
}

// CONTRACTBITBONSTORAGE is a free data retrieval call binding the contract method 0xe7f6c30e.
//
// Solidity: function CONTRACT_BITBON_STORAGE() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCallerSession) CONTRACTBITBONSTORAGE() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.CONTRACTBITBONSTORAGE(&_SafeTransferStorageImpl.CallOpts)
}

// CONTRACTBITBONSUPPORT is a free data retrieval call binding the contract method 0x20139c74.
//
// Solidity: function CONTRACT_BITBON_SUPPORT() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCaller) CONTRACTBITBONSUPPORT(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SafeTransferStorageImpl.contract.Call(opts, out, "CONTRACT_BITBON_SUPPORT")
	return *ret0, err
}

// CONTRACTBITBONSUPPORT is a free data retrieval call binding the contract method 0x20139c74.
//
// Solidity: function CONTRACT_BITBON_SUPPORT() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplSession) CONTRACTBITBONSUPPORT() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.CONTRACTBITBONSUPPORT(&_SafeTransferStorageImpl.CallOpts)
}

// CONTRACTBITBONSUPPORT is a free data retrieval call binding the contract method 0x20139c74.
//
// Solidity: function CONTRACT_BITBON_SUPPORT() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCallerSession) CONTRACTBITBONSUPPORT() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.CONTRACTBITBONSUPPORT(&_SafeTransferStorageImpl.CallOpts)
}

// CONTRACTBONBALANCE is a free data retrieval call binding the contract method 0x40e02b7c.
//
// Solidity: function CONTRACT_BONBALANCE() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCaller) CONTRACTBONBALANCE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SafeTransferStorageImpl.contract.Call(opts, out, "CONTRACT_BONBALANCE")
	return *ret0, err
}

// CONTRACTBONBALANCE is a free data retrieval call binding the contract method 0x40e02b7c.
//
// Solidity: function CONTRACT_BONBALANCE() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplSession) CONTRACTBONBALANCE() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.CONTRACTBONBALANCE(&_SafeTransferStorageImpl.CallOpts)
}

// CONTRACTBONBALANCE is a free data retrieval call binding the contract method 0x40e02b7c.
//
// Solidity: function CONTRACT_BONBALANCE() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCallerSession) CONTRACTBONBALANCE() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.CONTRACTBONBALANCE(&_SafeTransferStorageImpl.CallOpts)
}

// CONTRACTBONTRANSFER is a free data retrieval call binding the contract method 0xc18f4a8a.
//
// Solidity: function CONTRACT_BONTRANSFER() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCaller) CONTRACTBONTRANSFER(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SafeTransferStorageImpl.contract.Call(opts, out, "CONTRACT_BONTRANSFER")
	return *ret0, err
}

// CONTRACTBONTRANSFER is a free data retrieval call binding the contract method 0xc18f4a8a.
//
// Solidity: function CONTRACT_BONTRANSFER() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplSession) CONTRACTBONTRANSFER() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.CONTRACTBONTRANSFER(&_SafeTransferStorageImpl.CallOpts)
}

// CONTRACTBONTRANSFER is a free data retrieval call binding the contract method 0xc18f4a8a.
//
// Solidity: function CONTRACT_BONTRANSFER() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCallerSession) CONTRACTBONTRANSFER() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.CONTRACTBONTRANSFER(&_SafeTransferStorageImpl.CallOpts)
}

// CONTRACTDEX is a free data retrieval call binding the contract method 0x52c1f844.
//
// Solidity: function CONTRACT_DEX() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCaller) CONTRACTDEX(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SafeTransferStorageImpl.contract.Call(opts, out, "CONTRACT_DEX")
	return *ret0, err
}

// CONTRACTDEX is a free data retrieval call binding the contract method 0x52c1f844.
//
// Solidity: function CONTRACT_DEX() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplSession) CONTRACTDEX() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.CONTRACTDEX(&_SafeTransferStorageImpl.CallOpts)
}

// CONTRACTDEX is a free data retrieval call binding the contract method 0x52c1f844.
//
// Solidity: function CONTRACT_DEX() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCallerSession) CONTRACTDEX() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.CONTRACTDEX(&_SafeTransferStorageImpl.CallOpts)
}

// CONTRACTDISTRIBUTIONSTORAGE is a free data retrieval call binding the contract method 0x08097327.
//
// Solidity: function CONTRACT_DISTRIBUTION_STORAGE() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCaller) CONTRACTDISTRIBUTIONSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SafeTransferStorageImpl.contract.Call(opts, out, "CONTRACT_DISTRIBUTION_STORAGE")
	return *ret0, err
}

// CONTRACTDISTRIBUTIONSTORAGE is a free data retrieval call binding the contract method 0x08097327.
//
// Solidity: function CONTRACT_DISTRIBUTION_STORAGE() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplSession) CONTRACTDISTRIBUTIONSTORAGE() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.CONTRACTDISTRIBUTIONSTORAGE(&_SafeTransferStorageImpl.CallOpts)
}

// CONTRACTDISTRIBUTIONSTORAGE is a free data retrieval call binding the contract method 0x08097327.
//
// Solidity: function CONTRACT_DISTRIBUTION_STORAGE() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCallerSession) CONTRACTDISTRIBUTIONSTORAGE() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.CONTRACTDISTRIBUTIONSTORAGE(&_SafeTransferStorageImpl.CallOpts)
}

// CONTRACTEXCHANGE is a free data retrieval call binding the contract method 0xa0e3a10e.
//
// Solidity: function CONTRACT_EXCHANGE() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCaller) CONTRACTEXCHANGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SafeTransferStorageImpl.contract.Call(opts, out, "CONTRACT_EXCHANGE")
	return *ret0, err
}

// CONTRACTEXCHANGE is a free data retrieval call binding the contract method 0xa0e3a10e.
//
// Solidity: function CONTRACT_EXCHANGE() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplSession) CONTRACTEXCHANGE() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.CONTRACTEXCHANGE(&_SafeTransferStorageImpl.CallOpts)
}

// CONTRACTEXCHANGE is a free data retrieval call binding the contract method 0xa0e3a10e.
//
// Solidity: function CONTRACT_EXCHANGE() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCallerSession) CONTRACTEXCHANGE() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.CONTRACTEXCHANGE(&_SafeTransferStorageImpl.CallOpts)
}

// CONTRACTEXCHANGESTORAGE is a free data retrieval call binding the contract method 0x001570f2.
//
// Solidity: function CONTRACT_EXCHANGE_STORAGE() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCaller) CONTRACTEXCHANGESTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SafeTransferStorageImpl.contract.Call(opts, out, "CONTRACT_EXCHANGE_STORAGE")
	return *ret0, err
}

// CONTRACTEXCHANGESTORAGE is a free data retrieval call binding the contract method 0x001570f2.
//
// Solidity: function CONTRACT_EXCHANGE_STORAGE() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplSession) CONTRACTEXCHANGESTORAGE() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.CONTRACTEXCHANGESTORAGE(&_SafeTransferStorageImpl.CallOpts)
}

// CONTRACTEXCHANGESTORAGE is a free data retrieval call binding the contract method 0x001570f2.
//
// Solidity: function CONTRACT_EXCHANGE_STORAGE() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCallerSession) CONTRACTEXCHANGESTORAGE() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.CONTRACTEXCHANGESTORAGE(&_SafeTransferStorageImpl.CallOpts)
}

// CONTRACTFEESTORAGE is a free data retrieval call binding the contract method 0xcd5e65a4.
//
// Solidity: function CONTRACT_FEE_STORAGE() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCaller) CONTRACTFEESTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SafeTransferStorageImpl.contract.Call(opts, out, "CONTRACT_FEE_STORAGE")
	return *ret0, err
}

// CONTRACTFEESTORAGE is a free data retrieval call binding the contract method 0xcd5e65a4.
//
// Solidity: function CONTRACT_FEE_STORAGE() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplSession) CONTRACTFEESTORAGE() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.CONTRACTFEESTORAGE(&_SafeTransferStorageImpl.CallOpts)
}

// CONTRACTFEESTORAGE is a free data retrieval call binding the contract method 0xcd5e65a4.
//
// Solidity: function CONTRACT_FEE_STORAGE() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCallerSession) CONTRACTFEESTORAGE() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.CONTRACTFEESTORAGE(&_SafeTransferStorageImpl.CallOpts)
}

// CONTRACTGENERATOR is a free data retrieval call binding the contract method 0xccea2b46.
//
// Solidity: function CONTRACT_GENERATOR() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCaller) CONTRACTGENERATOR(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SafeTransferStorageImpl.contract.Call(opts, out, "CONTRACT_GENERATOR")
	return *ret0, err
}

// CONTRACTGENERATOR is a free data retrieval call binding the contract method 0xccea2b46.
//
// Solidity: function CONTRACT_GENERATOR() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplSession) CONTRACTGENERATOR() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.CONTRACTGENERATOR(&_SafeTransferStorageImpl.CallOpts)
}

// CONTRACTGENERATOR is a free data retrieval call binding the contract method 0xccea2b46.
//
// Solidity: function CONTRACT_GENERATOR() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCallerSession) CONTRACTGENERATOR() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.CONTRACTGENERATOR(&_SafeTransferStorageImpl.CallOpts)
}

// CONTRACTMININGAGENTSTORAGE is a free data retrieval call binding the contract method 0x83cc588d.
//
// Solidity: function CONTRACT_MINING_AGENT_STORAGE() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCaller) CONTRACTMININGAGENTSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SafeTransferStorageImpl.contract.Call(opts, out, "CONTRACT_MINING_AGENT_STORAGE")
	return *ret0, err
}

// CONTRACTMININGAGENTSTORAGE is a free data retrieval call binding the contract method 0x83cc588d.
//
// Solidity: function CONTRACT_MINING_AGENT_STORAGE() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplSession) CONTRACTMININGAGENTSTORAGE() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.CONTRACTMININGAGENTSTORAGE(&_SafeTransferStorageImpl.CallOpts)
}

// CONTRACTMININGAGENTSTORAGE is a free data retrieval call binding the contract method 0x83cc588d.
//
// Solidity: function CONTRACT_MINING_AGENT_STORAGE() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCallerSession) CONTRACTMININGAGENTSTORAGE() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.CONTRACTMININGAGENTSTORAGE(&_SafeTransferStorageImpl.CallOpts)
}

// CONTRACTMSBON is a free data retrieval call binding the contract method 0x28b3a976.
//
// Solidity: function CONTRACT_MSBON() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCaller) CONTRACTMSBON(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SafeTransferStorageImpl.contract.Call(opts, out, "CONTRACT_MSBON")
	return *ret0, err
}

// CONTRACTMSBON is a free data retrieval call binding the contract method 0x28b3a976.
//
// Solidity: function CONTRACT_MSBON() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplSession) CONTRACTMSBON() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.CONTRACTMSBON(&_SafeTransferStorageImpl.CallOpts)
}

// CONTRACTMSBON is a free data retrieval call binding the contract method 0x28b3a976.
//
// Solidity: function CONTRACT_MSBON() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCallerSession) CONTRACTMSBON() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.CONTRACTMSBON(&_SafeTransferStorageImpl.CallOpts)
}

// CONTRACTMSBONSTORAGE is a free data retrieval call binding the contract method 0xe04e6061.
//
// Solidity: function CONTRACT_MSBON_STORAGE() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCaller) CONTRACTMSBONSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SafeTransferStorageImpl.contract.Call(opts, out, "CONTRACT_MSBON_STORAGE")
	return *ret0, err
}

// CONTRACTMSBONSTORAGE is a free data retrieval call binding the contract method 0xe04e6061.
//
// Solidity: function CONTRACT_MSBON_STORAGE() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplSession) CONTRACTMSBONSTORAGE() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.CONTRACTMSBONSTORAGE(&_SafeTransferStorageImpl.CallOpts)
}

// CONTRACTMSBONSTORAGE is a free data retrieval call binding the contract method 0xe04e6061.
//
// Solidity: function CONTRACT_MSBON_STORAGE() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCallerSession) CONTRACTMSBONSTORAGE() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.CONTRACTMSBONSTORAGE(&_SafeTransferStorageImpl.CallOpts)
}

// CONTRACTMULTISIGADDADDRESS is a free data retrieval call binding the contract method 0xf4f65a5a.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ADDRESS() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCaller) CONTRACTMULTISIGADDADDRESS(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SafeTransferStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_ADD_ADDRESS")
	return *ret0, err
}

// CONTRACTMULTISIGADDADDRESS is a free data retrieval call binding the contract method 0xf4f65a5a.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ADDRESS() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplSession) CONTRACTMULTISIGADDADDRESS() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.CONTRACTMULTISIGADDADDRESS(&_SafeTransferStorageImpl.CallOpts)
}

// CONTRACTMULTISIGADDADDRESS is a free data retrieval call binding the contract method 0xf4f65a5a.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ADDRESS() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCallerSession) CONTRACTMULTISIGADDADDRESS() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.CONTRACTMULTISIGADDADDRESS(&_SafeTransferStorageImpl.CallOpts)
}

// CONTRACTMULTISIGADDADMIN is a free data retrieval call binding the contract method 0x83888c68.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ADMIN() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCaller) CONTRACTMULTISIGADDADMIN(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SafeTransferStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_ADD_ADMIN")
	return *ret0, err
}

// CONTRACTMULTISIGADDADMIN is a free data retrieval call binding the contract method 0x83888c68.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ADMIN() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplSession) CONTRACTMULTISIGADDADMIN() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.CONTRACTMULTISIGADDADMIN(&_SafeTransferStorageImpl.CallOpts)
}

// CONTRACTMULTISIGADDADMIN is a free data retrieval call binding the contract method 0x83888c68.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ADMIN() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCallerSession) CONTRACTMULTISIGADDADMIN() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.CONTRACTMULTISIGADDADMIN(&_SafeTransferStorageImpl.CallOpts)
}

// CONTRACTMULTISIGADDFEE is a free data retrieval call binding the contract method 0xc43e8cab.
//
// Solidity: function CONTRACT_MULTISIG_ADD_FEE() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCaller) CONTRACTMULTISIGADDFEE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SafeTransferStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_ADD_FEE")
	return *ret0, err
}

// CONTRACTMULTISIGADDFEE is a free data retrieval call binding the contract method 0xc43e8cab.
//
// Solidity: function CONTRACT_MULTISIG_ADD_FEE() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplSession) CONTRACTMULTISIGADDFEE() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.CONTRACTMULTISIGADDFEE(&_SafeTransferStorageImpl.CallOpts)
}

// CONTRACTMULTISIGADDFEE is a free data retrieval call binding the contract method 0xc43e8cab.
//
// Solidity: function CONTRACT_MULTISIG_ADD_FEE() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCallerSession) CONTRACTMULTISIGADDFEE() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.CONTRACTMULTISIGADDFEE(&_SafeTransferStorageImpl.CallOpts)
}

// CONTRACTMULTISIGADDROLE is a free data retrieval call binding the contract method 0x20a0914b.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ROLE() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCaller) CONTRACTMULTISIGADDROLE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SafeTransferStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_ADD_ROLE")
	return *ret0, err
}

// CONTRACTMULTISIGADDROLE is a free data retrieval call binding the contract method 0x20a0914b.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ROLE() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplSession) CONTRACTMULTISIGADDROLE() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.CONTRACTMULTISIGADDROLE(&_SafeTransferStorageImpl.CallOpts)
}

// CONTRACTMULTISIGADDROLE is a free data retrieval call binding the contract method 0x20a0914b.
//
// Solidity: function CONTRACT_MULTISIG_ADD_ROLE() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCallerSession) CONTRACTMULTISIGADDROLE() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.CONTRACTMULTISIGADDROLE(&_SafeTransferStorageImpl.CallOpts)
}

// CONTRACTMULTISIGDROPPERMISSION is a free data retrieval call binding the contract method 0x3f3fbc39.
//
// Solidity: function CONTRACT_MULTISIG_DROP_PERMISSION() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCaller) CONTRACTMULTISIGDROPPERMISSION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SafeTransferStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_DROP_PERMISSION")
	return *ret0, err
}

// CONTRACTMULTISIGDROPPERMISSION is a free data retrieval call binding the contract method 0x3f3fbc39.
//
// Solidity: function CONTRACT_MULTISIG_DROP_PERMISSION() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplSession) CONTRACTMULTISIGDROPPERMISSION() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.CONTRACTMULTISIGDROPPERMISSION(&_SafeTransferStorageImpl.CallOpts)
}

// CONTRACTMULTISIGDROPPERMISSION is a free data retrieval call binding the contract method 0x3f3fbc39.
//
// Solidity: function CONTRACT_MULTISIG_DROP_PERMISSION() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCallerSession) CONTRACTMULTISIGDROPPERMISSION() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.CONTRACTMULTISIGDROPPERMISSION(&_SafeTransferStorageImpl.CallOpts)
}

// CONTRACTMULTISIGEDITADMIN is a free data retrieval call binding the contract method 0x91b0f051.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_ADMIN() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCaller) CONTRACTMULTISIGEDITADMIN(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SafeTransferStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_EDIT_ADMIN")
	return *ret0, err
}

// CONTRACTMULTISIGEDITADMIN is a free data retrieval call binding the contract method 0x91b0f051.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_ADMIN() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplSession) CONTRACTMULTISIGEDITADMIN() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.CONTRACTMULTISIGEDITADMIN(&_SafeTransferStorageImpl.CallOpts)
}

// CONTRACTMULTISIGEDITADMIN is a free data retrieval call binding the contract method 0x91b0f051.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_ADMIN() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCallerSession) CONTRACTMULTISIGEDITADMIN() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.CONTRACTMULTISIGEDITADMIN(&_SafeTransferStorageImpl.CallOpts)
}

// CONTRACTMULTISIGEDITFEE is a free data retrieval call binding the contract method 0x6b2abe17.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_FEE() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCaller) CONTRACTMULTISIGEDITFEE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SafeTransferStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_EDIT_FEE")
	return *ret0, err
}

// CONTRACTMULTISIGEDITFEE is a free data retrieval call binding the contract method 0x6b2abe17.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_FEE() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplSession) CONTRACTMULTISIGEDITFEE() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.CONTRACTMULTISIGEDITFEE(&_SafeTransferStorageImpl.CallOpts)
}

// CONTRACTMULTISIGEDITFEE is a free data retrieval call binding the contract method 0x6b2abe17.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_FEE() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCallerSession) CONTRACTMULTISIGEDITFEE() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.CONTRACTMULTISIGEDITFEE(&_SafeTransferStorageImpl.CallOpts)
}

// CONTRACTMULTISIGEDITROLE is a free data retrieval call binding the contract method 0xc5f66412.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_ROLE() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCaller) CONTRACTMULTISIGEDITROLE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SafeTransferStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_EDIT_ROLE")
	return *ret0, err
}

// CONTRACTMULTISIGEDITROLE is a free data retrieval call binding the contract method 0xc5f66412.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_ROLE() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplSession) CONTRACTMULTISIGEDITROLE() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.CONTRACTMULTISIGEDITROLE(&_SafeTransferStorageImpl.CallOpts)
}

// CONTRACTMULTISIGEDITROLE is a free data retrieval call binding the contract method 0xc5f66412.
//
// Solidity: function CONTRACT_MULTISIG_EDIT_ROLE() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCallerSession) CONTRACTMULTISIGEDITROLE() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.CONTRACTMULTISIGEDITROLE(&_SafeTransferStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEADMINBYADDRESS is a free data retrieval call binding the contract method 0xc816b091.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_BY_ADDRESS() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCaller) CONTRACTMULTISIGREMOVEADMINBYADDRESS(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SafeTransferStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REMOVE_ADMIN_BY_ADDRESS")
	return *ret0, err
}

// CONTRACTMULTISIGREMOVEADMINBYADDRESS is a free data retrieval call binding the contract method 0xc816b091.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_BY_ADDRESS() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplSession) CONTRACTMULTISIGREMOVEADMINBYADDRESS() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.CONTRACTMULTISIGREMOVEADMINBYADDRESS(&_SafeTransferStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEADMINBYADDRESS is a free data retrieval call binding the contract method 0xc816b091.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_BY_ADDRESS() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCallerSession) CONTRACTMULTISIGREMOVEADMINBYADDRESS() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.CONTRACTMULTISIGREMOVEADMINBYADDRESS(&_SafeTransferStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEADMINBYINDEX is a free data retrieval call binding the contract method 0x751fbea2.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_BY_INDEX() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCaller) CONTRACTMULTISIGREMOVEADMINBYINDEX(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SafeTransferStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REMOVE_ADMIN_BY_INDEX")
	return *ret0, err
}

// CONTRACTMULTISIGREMOVEADMINBYINDEX is a free data retrieval call binding the contract method 0x751fbea2.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_BY_INDEX() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplSession) CONTRACTMULTISIGREMOVEADMINBYINDEX() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.CONTRACTMULTISIGREMOVEADMINBYINDEX(&_SafeTransferStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEADMINBYINDEX is a free data retrieval call binding the contract method 0x751fbea2.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_BY_INDEX() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCallerSession) CONTRACTMULTISIGREMOVEADMINBYINDEX() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.CONTRACTMULTISIGREMOVEADMINBYINDEX(&_SafeTransferStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEADMINROLE is a free data retrieval call binding the contract method 0x53c00200.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_ROLE() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCaller) CONTRACTMULTISIGREMOVEADMINROLE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SafeTransferStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REMOVE_ADMIN_ROLE")
	return *ret0, err
}

// CONTRACTMULTISIGREMOVEADMINROLE is a free data retrieval call binding the contract method 0x53c00200.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_ROLE() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplSession) CONTRACTMULTISIGREMOVEADMINROLE() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.CONTRACTMULTISIGREMOVEADMINROLE(&_SafeTransferStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEADMINROLE is a free data retrieval call binding the contract method 0x53c00200.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ADMIN_ROLE() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCallerSession) CONTRACTMULTISIGREMOVEADMINROLE() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.CONTRACTMULTISIGREMOVEADMINROLE(&_SafeTransferStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEFEE is a free data retrieval call binding the contract method 0xdb614d25.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_FEE() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCaller) CONTRACTMULTISIGREMOVEFEE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SafeTransferStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REMOVE_FEE")
	return *ret0, err
}

// CONTRACTMULTISIGREMOVEFEE is a free data retrieval call binding the contract method 0xdb614d25.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_FEE() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplSession) CONTRACTMULTISIGREMOVEFEE() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.CONTRACTMULTISIGREMOVEFEE(&_SafeTransferStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEFEE is a free data retrieval call binding the contract method 0xdb614d25.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_FEE() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCallerSession) CONTRACTMULTISIGREMOVEFEE() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.CONTRACTMULTISIGREMOVEFEE(&_SafeTransferStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEROLE is a free data retrieval call binding the contract method 0x23648496.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ROLE() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCaller) CONTRACTMULTISIGREMOVEROLE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SafeTransferStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REMOVE_ROLE")
	return *ret0, err
}

// CONTRACTMULTISIGREMOVEROLE is a free data retrieval call binding the contract method 0x23648496.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ROLE() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplSession) CONTRACTMULTISIGREMOVEROLE() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.CONTRACTMULTISIGREMOVEROLE(&_SafeTransferStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREMOVEROLE is a free data retrieval call binding the contract method 0x23648496.
//
// Solidity: function CONTRACT_MULTISIG_REMOVE_ROLE() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCallerSession) CONTRACTMULTISIGREMOVEROLE() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.CONTRACTMULTISIGREMOVEROLE(&_SafeTransferStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREPLENISHCAPITALIZATION is a free data retrieval call binding the contract method 0x62d5d523.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_CAPITALIZATION() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCaller) CONTRACTMULTISIGREPLENISHCAPITALIZATION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SafeTransferStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REPLENISH_CAPITALIZATION")
	return *ret0, err
}

// CONTRACTMULTISIGREPLENISHCAPITALIZATION is a free data retrieval call binding the contract method 0x62d5d523.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_CAPITALIZATION() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplSession) CONTRACTMULTISIGREPLENISHCAPITALIZATION() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.CONTRACTMULTISIGREPLENISHCAPITALIZATION(&_SafeTransferStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREPLENISHCAPITALIZATION is a free data retrieval call binding the contract method 0x62d5d523.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_CAPITALIZATION() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCallerSession) CONTRACTMULTISIGREPLENISHCAPITALIZATION() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.CONTRACTMULTISIGREPLENISHCAPITALIZATION(&_SafeTransferStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREPLENISHCERTIFICATE is a free data retrieval call binding the contract method 0xfd19b7e1.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_CERTIFICATE() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCaller) CONTRACTMULTISIGREPLENISHCERTIFICATE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SafeTransferStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REPLENISH_CERTIFICATE")
	return *ret0, err
}

// CONTRACTMULTISIGREPLENISHCERTIFICATE is a free data retrieval call binding the contract method 0xfd19b7e1.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_CERTIFICATE() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplSession) CONTRACTMULTISIGREPLENISHCERTIFICATE() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.CONTRACTMULTISIGREPLENISHCERTIFICATE(&_SafeTransferStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREPLENISHCERTIFICATE is a free data retrieval call binding the contract method 0xfd19b7e1.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_CERTIFICATE() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCallerSession) CONTRACTMULTISIGREPLENISHCERTIFICATE() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.CONTRACTMULTISIGREPLENISHCERTIFICATE(&_SafeTransferStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREPLENISHEMISSION is a free data retrieval call binding the contract method 0x0cb59fde.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_EMISSION() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCaller) CONTRACTMULTISIGREPLENISHEMISSION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SafeTransferStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_REPLENISH_EMISSION")
	return *ret0, err
}

// CONTRACTMULTISIGREPLENISHEMISSION is a free data retrieval call binding the contract method 0x0cb59fde.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_EMISSION() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplSession) CONTRACTMULTISIGREPLENISHEMISSION() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.CONTRACTMULTISIGREPLENISHEMISSION(&_SafeTransferStorageImpl.CallOpts)
}

// CONTRACTMULTISIGREPLENISHEMISSION is a free data retrieval call binding the contract method 0x0cb59fde.
//
// Solidity: function CONTRACT_MULTISIG_REPLENISH_EMISSION() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCallerSession) CONTRACTMULTISIGREPLENISHEMISSION() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.CONTRACTMULTISIGREPLENISHEMISSION(&_SafeTransferStorageImpl.CallOpts)
}

// CONTRACTMULTISIGSETADMINROLE is a free data retrieval call binding the contract method 0x8f209ab7.
//
// Solidity: function CONTRACT_MULTISIG_SET_ADMIN_ROLE() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCaller) CONTRACTMULTISIGSETADMINROLE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SafeTransferStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_SET_ADMIN_ROLE")
	return *ret0, err
}

// CONTRACTMULTISIGSETADMINROLE is a free data retrieval call binding the contract method 0x8f209ab7.
//
// Solidity: function CONTRACT_MULTISIG_SET_ADMIN_ROLE() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplSession) CONTRACTMULTISIGSETADMINROLE() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.CONTRACTMULTISIGSETADMINROLE(&_SafeTransferStorageImpl.CallOpts)
}

// CONTRACTMULTISIGSETADMINROLE is a free data retrieval call binding the contract method 0x8f209ab7.
//
// Solidity: function CONTRACT_MULTISIG_SET_ADMIN_ROLE() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCallerSession) CONTRACTMULTISIGSETADMINROLE() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.CONTRACTMULTISIGSETADMINROLE(&_SafeTransferStorageImpl.CallOpts)
}

// CONTRACTMULTISIGSETPERMISSION is a free data retrieval call binding the contract method 0xb0e91914.
//
// Solidity: function CONTRACT_MULTISIG_SET_PERMISSION() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCaller) CONTRACTMULTISIGSETPERMISSION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SafeTransferStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_SET_PERMISSION")
	return *ret0, err
}

// CONTRACTMULTISIGSETPERMISSION is a free data retrieval call binding the contract method 0xb0e91914.
//
// Solidity: function CONTRACT_MULTISIG_SET_PERMISSION() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplSession) CONTRACTMULTISIGSETPERMISSION() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.CONTRACTMULTISIGSETPERMISSION(&_SafeTransferStorageImpl.CallOpts)
}

// CONTRACTMULTISIGSETPERMISSION is a free data retrieval call binding the contract method 0xb0e91914.
//
// Solidity: function CONTRACT_MULTISIG_SET_PERMISSION() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCallerSession) CONTRACTMULTISIGSETPERMISSION() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.CONTRACTMULTISIGSETPERMISSION(&_SafeTransferStorageImpl.CallOpts)
}

// CONTRACTMULTISIGSTORAGE is a free data retrieval call binding the contract method 0xe7e1a303.
//
// Solidity: function CONTRACT_MULTISIG_STORAGE() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCaller) CONTRACTMULTISIGSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SafeTransferStorageImpl.contract.Call(opts, out, "CONTRACT_MULTISIG_STORAGE")
	return *ret0, err
}

// CONTRACTMULTISIGSTORAGE is a free data retrieval call binding the contract method 0xe7e1a303.
//
// Solidity: function CONTRACT_MULTISIG_STORAGE() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplSession) CONTRACTMULTISIGSTORAGE() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.CONTRACTMULTISIGSTORAGE(&_SafeTransferStorageImpl.CallOpts)
}

// CONTRACTMULTISIGSTORAGE is a free data retrieval call binding the contract method 0xe7e1a303.
//
// Solidity: function CONTRACT_MULTISIG_STORAGE() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCallerSession) CONTRACTMULTISIGSTORAGE() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.CONTRACTMULTISIGSTORAGE(&_SafeTransferStorageImpl.CallOpts)
}

// CONTRACTOTC is a free data retrieval call binding the contract method 0x5eceac70.
//
// Solidity: function CONTRACT_OTC() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCaller) CONTRACTOTC(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SafeTransferStorageImpl.contract.Call(opts, out, "CONTRACT_OTC")
	return *ret0, err
}

// CONTRACTOTC is a free data retrieval call binding the contract method 0x5eceac70.
//
// Solidity: function CONTRACT_OTC() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplSession) CONTRACTOTC() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.CONTRACTOTC(&_SafeTransferStorageImpl.CallOpts)
}

// CONTRACTOTC is a free data retrieval call binding the contract method 0x5eceac70.
//
// Solidity: function CONTRACT_OTC() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCallerSession) CONTRACTOTC() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.CONTRACTOTC(&_SafeTransferStorageImpl.CallOpts)
}

// CONTRACTOTCSTORAGE is a free data retrieval call binding the contract method 0x825b476f.
//
// Solidity: function CONTRACT_OTC_STORAGE() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCaller) CONTRACTOTCSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SafeTransferStorageImpl.contract.Call(opts, out, "CONTRACT_OTC_STORAGE")
	return *ret0, err
}

// CONTRACTOTCSTORAGE is a free data retrieval call binding the contract method 0x825b476f.
//
// Solidity: function CONTRACT_OTC_STORAGE() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplSession) CONTRACTOTCSTORAGE() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.CONTRACTOTCSTORAGE(&_SafeTransferStorageImpl.CallOpts)
}

// CONTRACTOTCSTORAGE is a free data retrieval call binding the contract method 0x825b476f.
//
// Solidity: function CONTRACT_OTC_STORAGE() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCallerSession) CONTRACTOTCSTORAGE() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.CONTRACTOTCSTORAGE(&_SafeTransferStorageImpl.CallOpts)
}

// CONTRACTREPLENISHCAPITALIZATION is a free data retrieval call binding the contract method 0xff5e09d8.
//
// Solidity: function CONTRACT_REPLENISH_CAPITALIZATION() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCaller) CONTRACTREPLENISHCAPITALIZATION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SafeTransferStorageImpl.contract.Call(opts, out, "CONTRACT_REPLENISH_CAPITALIZATION")
	return *ret0, err
}

// CONTRACTREPLENISHCAPITALIZATION is a free data retrieval call binding the contract method 0xff5e09d8.
//
// Solidity: function CONTRACT_REPLENISH_CAPITALIZATION() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplSession) CONTRACTREPLENISHCAPITALIZATION() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.CONTRACTREPLENISHCAPITALIZATION(&_SafeTransferStorageImpl.CallOpts)
}

// CONTRACTREPLENISHCAPITALIZATION is a free data retrieval call binding the contract method 0xff5e09d8.
//
// Solidity: function CONTRACT_REPLENISH_CAPITALIZATION() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCallerSession) CONTRACTREPLENISHCAPITALIZATION() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.CONTRACTREPLENISHCAPITALIZATION(&_SafeTransferStorageImpl.CallOpts)
}

// CONTRACTRESERVEDALIASSTORAGE is a free data retrieval call binding the contract method 0x70a94aa7.
//
// Solidity: function CONTRACT_RESERVED_ALIAS_STORAGE() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCaller) CONTRACTRESERVEDALIASSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SafeTransferStorageImpl.contract.Call(opts, out, "CONTRACT_RESERVED_ALIAS_STORAGE")
	return *ret0, err
}

// CONTRACTRESERVEDALIASSTORAGE is a free data retrieval call binding the contract method 0x70a94aa7.
//
// Solidity: function CONTRACT_RESERVED_ALIAS_STORAGE() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplSession) CONTRACTRESERVEDALIASSTORAGE() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.CONTRACTRESERVEDALIASSTORAGE(&_SafeTransferStorageImpl.CallOpts)
}

// CONTRACTRESERVEDALIASSTORAGE is a free data retrieval call binding the contract method 0x70a94aa7.
//
// Solidity: function CONTRACT_RESERVED_ALIAS_STORAGE() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCallerSession) CONTRACTRESERVEDALIASSTORAGE() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.CONTRACTRESERVEDALIASSTORAGE(&_SafeTransferStorageImpl.CallOpts)
}

// CONTRACTROLESTORAGE is a free data retrieval call binding the contract method 0xb938235b.
//
// Solidity: function CONTRACT_ROLE_STORAGE() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCaller) CONTRACTROLESTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SafeTransferStorageImpl.contract.Call(opts, out, "CONTRACT_ROLE_STORAGE")
	return *ret0, err
}

// CONTRACTROLESTORAGE is a free data retrieval call binding the contract method 0xb938235b.
//
// Solidity: function CONTRACT_ROLE_STORAGE() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplSession) CONTRACTROLESTORAGE() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.CONTRACTROLESTORAGE(&_SafeTransferStorageImpl.CallOpts)
}

// CONTRACTROLESTORAGE is a free data retrieval call binding the contract method 0xb938235b.
//
// Solidity: function CONTRACT_ROLE_STORAGE() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCallerSession) CONTRACTROLESTORAGE() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.CONTRACTROLESTORAGE(&_SafeTransferStorageImpl.CallOpts)
}

// CONTRACTSAFETRANSFERSTORAGE is a free data retrieval call binding the contract method 0xe7e45392.
//
// Solidity: function CONTRACT_SAFE_TRANSFER_STORAGE() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCaller) CONTRACTSAFETRANSFERSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SafeTransferStorageImpl.contract.Call(opts, out, "CONTRACT_SAFE_TRANSFER_STORAGE")
	return *ret0, err
}

// CONTRACTSAFETRANSFERSTORAGE is a free data retrieval call binding the contract method 0xe7e45392.
//
// Solidity: function CONTRACT_SAFE_TRANSFER_STORAGE() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplSession) CONTRACTSAFETRANSFERSTORAGE() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.CONTRACTSAFETRANSFERSTORAGE(&_SafeTransferStorageImpl.CallOpts)
}

// CONTRACTSAFETRANSFERSTORAGE is a free data retrieval call binding the contract method 0xe7e45392.
//
// Solidity: function CONTRACT_SAFE_TRANSFER_STORAGE() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCallerSession) CONTRACTSAFETRANSFERSTORAGE() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.CONTRACTSAFETRANSFERSTORAGE(&_SafeTransferStorageImpl.CallOpts)
}

// CONTRACTTRANSFER is a free data retrieval call binding the contract method 0xef94fe38.
//
// Solidity: function CONTRACT_TRANSFER() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCaller) CONTRACTTRANSFER(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SafeTransferStorageImpl.contract.Call(opts, out, "CONTRACT_TRANSFER")
	return *ret0, err
}

// CONTRACTTRANSFER is a free data retrieval call binding the contract method 0xef94fe38.
//
// Solidity: function CONTRACT_TRANSFER() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplSession) CONTRACTTRANSFER() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.CONTRACTTRANSFER(&_SafeTransferStorageImpl.CallOpts)
}

// CONTRACTTRANSFER is a free data retrieval call binding the contract method 0xef94fe38.
//
// Solidity: function CONTRACT_TRANSFER() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCallerSession) CONTRACTTRANSFER() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.CONTRACTTRANSFER(&_SafeTransferStorageImpl.CallOpts)
}

// CONTRACTTRANSFERFROMCAPITALIZATION is a free data retrieval call binding the contract method 0xc8f8decb.
//
// Solidity: function CONTRACT_TRANSFER_FROM_CAPITALIZATION() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCaller) CONTRACTTRANSFERFROMCAPITALIZATION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SafeTransferStorageImpl.contract.Call(opts, out, "CONTRACT_TRANSFER_FROM_CAPITALIZATION")
	return *ret0, err
}

// CONTRACTTRANSFERFROMCAPITALIZATION is a free data retrieval call binding the contract method 0xc8f8decb.
//
// Solidity: function CONTRACT_TRANSFER_FROM_CAPITALIZATION() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplSession) CONTRACTTRANSFERFROMCAPITALIZATION() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.CONTRACTTRANSFERFROMCAPITALIZATION(&_SafeTransferStorageImpl.CallOpts)
}

// CONTRACTTRANSFERFROMCAPITALIZATION is a free data retrieval call binding the contract method 0xc8f8decb.
//
// Solidity: function CONTRACT_TRANSFER_FROM_CAPITALIZATION() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCallerSession) CONTRACTTRANSFERFROMCAPITALIZATION() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.CONTRACTTRANSFERFROMCAPITALIZATION(&_SafeTransferStorageImpl.CallOpts)
}

// PERMISSIONACCESSRESTORATION is a free data retrieval call binding the contract method 0x5c3053bf.
//
// Solidity: function PERMISSION_ACCESS_RESTORATION() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCaller) PERMISSIONACCESSRESTORATION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SafeTransferStorageImpl.contract.Call(opts, out, "PERMISSION_ACCESS_RESTORATION")
	return *ret0, err
}

// PERMISSIONACCESSRESTORATION is a free data retrieval call binding the contract method 0x5c3053bf.
//
// Solidity: function PERMISSION_ACCESS_RESTORATION() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplSession) PERMISSIONACCESSRESTORATION() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.PERMISSIONACCESSRESTORATION(&_SafeTransferStorageImpl.CallOpts)
}

// PERMISSIONACCESSRESTORATION is a free data retrieval call binding the contract method 0x5c3053bf.
//
// Solidity: function PERMISSION_ACCESS_RESTORATION() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCallerSession) PERMISSIONACCESSRESTORATION() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.PERMISSIONACCESSRESTORATION(&_SafeTransferStorageImpl.CallOpts)
}

// PERMISSIONADMINSTORAGE is a free data retrieval call binding the contract method 0x0c9fc0db.
//
// Solidity: function PERMISSION_ADMIN_STORAGE() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCaller) PERMISSIONADMINSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SafeTransferStorageImpl.contract.Call(opts, out, "PERMISSION_ADMIN_STORAGE")
	return *ret0, err
}

// PERMISSIONADMINSTORAGE is a free data retrieval call binding the contract method 0x0c9fc0db.
//
// Solidity: function PERMISSION_ADMIN_STORAGE() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplSession) PERMISSIONADMINSTORAGE() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.PERMISSIONADMINSTORAGE(&_SafeTransferStorageImpl.CallOpts)
}

// PERMISSIONADMINSTORAGE is a free data retrieval call binding the contract method 0x0c9fc0db.
//
// Solidity: function PERMISSION_ADMIN_STORAGE() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCallerSession) PERMISSIONADMINSTORAGE() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.PERMISSIONADMINSTORAGE(&_SafeTransferStorageImpl.CallOpts)
}

// PERMISSIONCONTRACTSSTORAGE is a free data retrieval call binding the contract method 0x1516e83b.
//
// Solidity: function PERMISSION_CONTRACTS_STORAGE() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCaller) PERMISSIONCONTRACTSSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SafeTransferStorageImpl.contract.Call(opts, out, "PERMISSION_CONTRACTS_STORAGE")
	return *ret0, err
}

// PERMISSIONCONTRACTSSTORAGE is a free data retrieval call binding the contract method 0x1516e83b.
//
// Solidity: function PERMISSION_CONTRACTS_STORAGE() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplSession) PERMISSIONCONTRACTSSTORAGE() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.PERMISSIONCONTRACTSSTORAGE(&_SafeTransferStorageImpl.CallOpts)
}

// PERMISSIONCONTRACTSSTORAGE is a free data retrieval call binding the contract method 0x1516e83b.
//
// Solidity: function PERMISSION_CONTRACTS_STORAGE() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCallerSession) PERMISSIONCONTRACTSSTORAGE() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.PERMISSIONCONTRACTSSTORAGE(&_SafeTransferStorageImpl.CallOpts)
}

// PERMISSIONCONTRACTDEPLOY is a free data retrieval call binding the contract method 0x257a4a04.
//
// Solidity: function PERMISSION_CONTRACT_DEPLOY() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCaller) PERMISSIONCONTRACTDEPLOY(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SafeTransferStorageImpl.contract.Call(opts, out, "PERMISSION_CONTRACT_DEPLOY")
	return *ret0, err
}

// PERMISSIONCONTRACTDEPLOY is a free data retrieval call binding the contract method 0x257a4a04.
//
// Solidity: function PERMISSION_CONTRACT_DEPLOY() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplSession) PERMISSIONCONTRACTDEPLOY() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.PERMISSIONCONTRACTDEPLOY(&_SafeTransferStorageImpl.CallOpts)
}

// PERMISSIONCONTRACTDEPLOY is a free data retrieval call binding the contract method 0x257a4a04.
//
// Solidity: function PERMISSION_CONTRACT_DEPLOY() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCallerSession) PERMISSIONCONTRACTDEPLOY() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.PERMISSIONCONTRACTDEPLOY(&_SafeTransferStorageImpl.CallOpts)
}

// PERMISSIONEMISSION is a free data retrieval call binding the contract method 0x4ccd92d1.
//
// Solidity: function PERMISSION_EMISSION() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCaller) PERMISSIONEMISSION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SafeTransferStorageImpl.contract.Call(opts, out, "PERMISSION_EMISSION")
	return *ret0, err
}

// PERMISSIONEMISSION is a free data retrieval call binding the contract method 0x4ccd92d1.
//
// Solidity: function PERMISSION_EMISSION() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplSession) PERMISSIONEMISSION() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.PERMISSIONEMISSION(&_SafeTransferStorageImpl.CallOpts)
}

// PERMISSIONEMISSION is a free data retrieval call binding the contract method 0x4ccd92d1.
//
// Solidity: function PERMISSION_EMISSION() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCallerSession) PERMISSIONEMISSION() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.PERMISSIONEMISSION(&_SafeTransferStorageImpl.CallOpts)
}

// PERMISSIONFEESSTORAGE is a free data retrieval call binding the contract method 0xe837eb87.
//
// Solidity: function PERMISSION_FEES_STORAGE() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCaller) PERMISSIONFEESSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SafeTransferStorageImpl.contract.Call(opts, out, "PERMISSION_FEES_STORAGE")
	return *ret0, err
}

// PERMISSIONFEESSTORAGE is a free data retrieval call binding the contract method 0xe837eb87.
//
// Solidity: function PERMISSION_FEES_STORAGE() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplSession) PERMISSIONFEESSTORAGE() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.PERMISSIONFEESSTORAGE(&_SafeTransferStorageImpl.CallOpts)
}

// PERMISSIONFEESSTORAGE is a free data retrieval call binding the contract method 0xe837eb87.
//
// Solidity: function PERMISSION_FEES_STORAGE() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCallerSession) PERMISSIONFEESSTORAGE() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.PERMISSIONFEESSTORAGE(&_SafeTransferStorageImpl.CallOpts)
}

// PERMISSIONMULTISIGSTORAGE is a free data retrieval call binding the contract method 0xf1bb7aad.
//
// Solidity: function PERMISSION_MULTISIG_STORAGE() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCaller) PERMISSIONMULTISIGSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SafeTransferStorageImpl.contract.Call(opts, out, "PERMISSION_MULTISIG_STORAGE")
	return *ret0, err
}

// PERMISSIONMULTISIGSTORAGE is a free data retrieval call binding the contract method 0xf1bb7aad.
//
// Solidity: function PERMISSION_MULTISIG_STORAGE() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplSession) PERMISSIONMULTISIGSTORAGE() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.PERMISSIONMULTISIGSTORAGE(&_SafeTransferStorageImpl.CallOpts)
}

// PERMISSIONMULTISIGSTORAGE is a free data retrieval call binding the contract method 0xf1bb7aad.
//
// Solidity: function PERMISSION_MULTISIG_STORAGE() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCallerSession) PERMISSIONMULTISIGSTORAGE() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.PERMISSIONMULTISIGSTORAGE(&_SafeTransferStorageImpl.CallOpts)
}

// PERMISSIONROLESSTORAGE is a free data retrieval call binding the contract method 0x12a16032.
//
// Solidity: function PERMISSION_ROLES_STORAGE() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCaller) PERMISSIONROLESSTORAGE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SafeTransferStorageImpl.contract.Call(opts, out, "PERMISSION_ROLES_STORAGE")
	return *ret0, err
}

// PERMISSIONROLESSTORAGE is a free data retrieval call binding the contract method 0x12a16032.
//
// Solidity: function PERMISSION_ROLES_STORAGE() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplSession) PERMISSIONROLESSTORAGE() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.PERMISSIONROLESSTORAGE(&_SafeTransferStorageImpl.CallOpts)
}

// PERMISSIONROLESSTORAGE is a free data retrieval call binding the contract method 0x12a16032.
//
// Solidity: function PERMISSION_ROLES_STORAGE() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCallerSession) PERMISSIONROLESSTORAGE() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.PERMISSIONROLESSTORAGE(&_SafeTransferStorageImpl.CallOpts)
}

// ROLEACCESSVERIFIER is a free data retrieval call binding the contract method 0xc5405f92.
//
// Solidity: function ROLE_ACCESS_VERIFIER() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCaller) ROLEACCESSVERIFIER(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SafeTransferStorageImpl.contract.Call(opts, out, "ROLE_ACCESS_VERIFIER")
	return *ret0, err
}

// ROLEACCESSVERIFIER is a free data retrieval call binding the contract method 0xc5405f92.
//
// Solidity: function ROLE_ACCESS_VERIFIER() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplSession) ROLEACCESSVERIFIER() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.ROLEACCESSVERIFIER(&_SafeTransferStorageImpl.CallOpts)
}

// ROLEACCESSVERIFIER is a free data retrieval call binding the contract method 0xc5405f92.
//
// Solidity: function ROLE_ACCESS_VERIFIER() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCallerSession) ROLEACCESSVERIFIER() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.ROLEACCESSVERIFIER(&_SafeTransferStorageImpl.CallOpts)
}

// ROLEBITBONISSUEVERIFIER is a free data retrieval call binding the contract method 0xf0a012a6.
//
// Solidity: function ROLE_BITBON_ISSUE_VERIFIER() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCaller) ROLEBITBONISSUEVERIFIER(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SafeTransferStorageImpl.contract.Call(opts, out, "ROLE_BITBON_ISSUE_VERIFIER")
	return *ret0, err
}

// ROLEBITBONISSUEVERIFIER is a free data retrieval call binding the contract method 0xf0a012a6.
//
// Solidity: function ROLE_BITBON_ISSUE_VERIFIER() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplSession) ROLEBITBONISSUEVERIFIER() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.ROLEBITBONISSUEVERIFIER(&_SafeTransferStorageImpl.CallOpts)
}

// ROLEBITBONISSUEVERIFIER is a free data retrieval call binding the contract method 0xf0a012a6.
//
// Solidity: function ROLE_BITBON_ISSUE_VERIFIER() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCallerSession) ROLEBITBONISSUEVERIFIER() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.ROLEBITBONISSUEVERIFIER(&_SafeTransferStorageImpl.CallOpts)
}

// ROLECOMMISSIONVERIFIER is a free data retrieval call binding the contract method 0x8ab3de8c.
//
// Solidity: function ROLE_COMMISSION_VERIFIER() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCaller) ROLECOMMISSIONVERIFIER(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SafeTransferStorageImpl.contract.Call(opts, out, "ROLE_COMMISSION_VERIFIER")
	return *ret0, err
}

// ROLECOMMISSIONVERIFIER is a free data retrieval call binding the contract method 0x8ab3de8c.
//
// Solidity: function ROLE_COMMISSION_VERIFIER() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplSession) ROLECOMMISSIONVERIFIER() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.ROLECOMMISSIONVERIFIER(&_SafeTransferStorageImpl.CallOpts)
}

// ROLECOMMISSIONVERIFIER is a free data retrieval call binding the contract method 0x8ab3de8c.
//
// Solidity: function ROLE_COMMISSION_VERIFIER() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCallerSession) ROLECOMMISSIONVERIFIER() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.ROLECOMMISSIONVERIFIER(&_SafeTransferStorageImpl.CallOpts)
}

// ROLEDEPLOYADMIN is a free data retrieval call binding the contract method 0x667bf6e5.
//
// Solidity: function ROLE_DEPLOY_ADMIN() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCaller) ROLEDEPLOYADMIN(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SafeTransferStorageImpl.contract.Call(opts, out, "ROLE_DEPLOY_ADMIN")
	return *ret0, err
}

// ROLEDEPLOYADMIN is a free data retrieval call binding the contract method 0x667bf6e5.
//
// Solidity: function ROLE_DEPLOY_ADMIN() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplSession) ROLEDEPLOYADMIN() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.ROLEDEPLOYADMIN(&_SafeTransferStorageImpl.CallOpts)
}

// ROLEDEPLOYADMIN is a free data retrieval call binding the contract method 0x667bf6e5.
//
// Solidity: function ROLE_DEPLOY_ADMIN() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCallerSession) ROLEDEPLOYADMIN() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.ROLEDEPLOYADMIN(&_SafeTransferStorageImpl.CallOpts)
}

// ROLEPERMISSIONADMIN is a free data retrieval call binding the contract method 0x1b239996.
//
// Solidity: function ROLE_PERMISSION_ADMIN() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCaller) ROLEPERMISSIONADMIN(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SafeTransferStorageImpl.contract.Call(opts, out, "ROLE_PERMISSION_ADMIN")
	return *ret0, err
}

// ROLEPERMISSIONADMIN is a free data retrieval call binding the contract method 0x1b239996.
//
// Solidity: function ROLE_PERMISSION_ADMIN() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplSession) ROLEPERMISSIONADMIN() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.ROLEPERMISSIONADMIN(&_SafeTransferStorageImpl.CallOpts)
}

// ROLEPERMISSIONADMIN is a free data retrieval call binding the contract method 0x1b239996.
//
// Solidity: function ROLE_PERMISSION_ADMIN() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCallerSession) ROLEPERMISSIONADMIN() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.ROLEPERMISSIONADMIN(&_SafeTransferStorageImpl.CallOpts)
}

// CheckRetries is a free data retrieval call binding the contract method 0x48e16acc.
//
// Solidity: function checkRetries(bytes transferId) view returns(bool)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCaller) CheckRetries(opts *bind.CallOpts, transferId []byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SafeTransferStorageImpl.contract.Call(opts, out, "checkRetries", transferId)
	return *ret0, err
}

// CheckRetries is a free data retrieval call binding the contract method 0x48e16acc.
//
// Solidity: function checkRetries(bytes transferId) view returns(bool)
func (_SafeTransferStorageImpl *SafeTransferStorageImplSession) CheckRetries(transferId []byte) (bool, error) {
	return _SafeTransferStorageImpl.Contract.CheckRetries(&_SafeTransferStorageImpl.CallOpts, transferId)
}

// CheckRetries is a free data retrieval call binding the contract method 0x48e16acc.
//
// Solidity: function checkRetries(bytes transferId) view returns(bool)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCallerSession) CheckRetries(transferId []byte) (bool, error) {
	return _SafeTransferStorageImpl.Contract.CheckRetries(&_SafeTransferStorageImpl.CallOpts, transferId)
}

// ContractStorage is a free data retrieval call binding the contract method 0x0549b27e.
//
// Solidity: function contractStorage() view returns(address)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCaller) ContractStorage(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SafeTransferStorageImpl.contract.Call(opts, out, "contractStorage")
	return *ret0, err
}

// ContractStorage is a free data retrieval call binding the contract method 0x0549b27e.
//
// Solidity: function contractStorage() view returns(address)
func (_SafeTransferStorageImpl *SafeTransferStorageImplSession) ContractStorage() (common.Address, error) {
	return _SafeTransferStorageImpl.Contract.ContractStorage(&_SafeTransferStorageImpl.CallOpts)
}

// ContractStorage is a free data retrieval call binding the contract method 0x0549b27e.
//
// Solidity: function contractStorage() view returns(address)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCallerSession) ContractStorage() (common.Address, error) {
	return _SafeTransferStorageImpl.Contract.ContractStorage(&_SafeTransferStorageImpl.CallOpts)
}

// GetExpiredTransfers is a free data retrieval call binding the contract method 0x538a46ab.
//
// Solidity: function getExpiredTransfers(uint256 first, uint256 last) view returns(bytes32[])
func (_SafeTransferStorageImpl *SafeTransferStorageImplCaller) GetExpiredTransfers(opts *bind.CallOpts, first *big.Int, last *big.Int) ([][32]byte, error) {
	var (
		ret0 = new([][32]byte)
	)
	out := ret0
	err := _SafeTransferStorageImpl.contract.Call(opts, out, "getExpiredTransfers", first, last)
	return *ret0, err
}

// GetExpiredTransfers is a free data retrieval call binding the contract method 0x538a46ab.
//
// Solidity: function getExpiredTransfers(uint256 first, uint256 last) view returns(bytes32[])
func (_SafeTransferStorageImpl *SafeTransferStorageImplSession) GetExpiredTransfers(first *big.Int, last *big.Int) ([][32]byte, error) {
	return _SafeTransferStorageImpl.Contract.GetExpiredTransfers(&_SafeTransferStorageImpl.CallOpts, first, last)
}

// GetExpiredTransfers is a free data retrieval call binding the contract method 0x538a46ab.
//
// Solidity: function getExpiredTransfers(uint256 first, uint256 last) view returns(bytes32[])
func (_SafeTransferStorageImpl *SafeTransferStorageImplCallerSession) GetExpiredTransfers(first *big.Int, last *big.Int) ([][32]byte, error) {
	return _SafeTransferStorageImpl.Contract.GetExpiredTransfers(&_SafeTransferStorageImpl.CallOpts, first, last)
}

// GetFee is a free data retrieval call binding the contract method 0x8204ecdd.
//
// Solidity: function getFee(bytes transferId) view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCaller) GetFee(opts *bind.CallOpts, transferId []byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SafeTransferStorageImpl.contract.Call(opts, out, "getFee", transferId)
	return *ret0, err
}

// GetFee is a free data retrieval call binding the contract method 0x8204ecdd.
//
// Solidity: function getFee(bytes transferId) view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplSession) GetFee(transferId []byte) (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.GetFee(&_SafeTransferStorageImpl.CallOpts, transferId)
}

// GetFee is a free data retrieval call binding the contract method 0x8204ecdd.
//
// Solidity: function getFee(bytes transferId) view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCallerSession) GetFee(transferId []byte) (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.GetFee(&_SafeTransferStorageImpl.CallOpts, transferId)
}

// GetOldestPending is a free data retrieval call binding the contract method 0x02f8fade.
//
// Solidity: function getOldestPending() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCaller) GetOldestPending(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SafeTransferStorageImpl.contract.Call(opts, out, "getOldestPending")
	return *ret0, err
}

// GetOldestPending is a free data retrieval call binding the contract method 0x02f8fade.
//
// Solidity: function getOldestPending() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplSession) GetOldestPending() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.GetOldestPending(&_SafeTransferStorageImpl.CallOpts)
}

// GetOldestPending is a free data retrieval call binding the contract method 0x02f8fade.
//
// Solidity: function getOldestPending() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCallerSession) GetOldestPending() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.GetOldestPending(&_SafeTransferStorageImpl.CallOpts)
}

// GetProtectionCode is a free data retrieval call binding the contract method 0x2eeb79f9.
//
// Solidity: function getProtectionCode(bytes transferId) view returns(bytes, bytes)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCaller) GetProtectionCode(opts *bind.CallOpts, transferId []byte) ([]byte, []byte, error) {
	var (
		ret0 = new([]byte)
		ret1 = new([]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _SafeTransferStorageImpl.contract.Call(opts, out, "getProtectionCode", transferId)
	return *ret0, *ret1, err
}

// GetProtectionCode is a free data retrieval call binding the contract method 0x2eeb79f9.
//
// Solidity: function getProtectionCode(bytes transferId) view returns(bytes, bytes)
func (_SafeTransferStorageImpl *SafeTransferStorageImplSession) GetProtectionCode(transferId []byte) ([]byte, []byte, error) {
	return _SafeTransferStorageImpl.Contract.GetProtectionCode(&_SafeTransferStorageImpl.CallOpts, transferId)
}

// GetProtectionCode is a free data retrieval call binding the contract method 0x2eeb79f9.
//
// Solidity: function getProtectionCode(bytes transferId) view returns(bytes, bytes)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCallerSession) GetProtectionCode(transferId []byte) ([]byte, []byte, error) {
	return _SafeTransferStorageImpl.Contract.GetProtectionCode(&_SafeTransferStorageImpl.CallOpts, transferId)
}

// GetRetries is a free data retrieval call binding the contract method 0x553f41e2.
//
// Solidity: function getRetries(bytes transferId) view returns(uint256, uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCaller) GetRetries(opts *bind.CallOpts, transferId []byte) (*big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _SafeTransferStorageImpl.contract.Call(opts, out, "getRetries", transferId)
	return *ret0, *ret1, err
}

// GetRetries is a free data retrieval call binding the contract method 0x553f41e2.
//
// Solidity: function getRetries(bytes transferId) view returns(uint256, uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplSession) GetRetries(transferId []byte) (*big.Int, *big.Int, error) {
	return _SafeTransferStorageImpl.Contract.GetRetries(&_SafeTransferStorageImpl.CallOpts, transferId)
}

// GetRetries is a free data retrieval call binding the contract method 0x553f41e2.
//
// Solidity: function getRetries(bytes transferId) view returns(uint256, uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCallerSession) GetRetries(transferId []byte) (*big.Int, *big.Int, error) {
	return _SafeTransferStorageImpl.Contract.GetRetries(&_SafeTransferStorageImpl.CallOpts, transferId)
}

// GetState is a free data retrieval call binding the contract method 0xe35f157c.
//
// Solidity: function getState(bytes transferId) view returns(uint8)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCaller) GetState(opts *bind.CallOpts, transferId []byte) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _SafeTransferStorageImpl.contract.Call(opts, out, "getState", transferId)
	return *ret0, err
}

// GetState is a free data retrieval call binding the contract method 0xe35f157c.
//
// Solidity: function getState(bytes transferId) view returns(uint8)
func (_SafeTransferStorageImpl *SafeTransferStorageImplSession) GetState(transferId []byte) (uint8, error) {
	return _SafeTransferStorageImpl.Contract.GetState(&_SafeTransferStorageImpl.CallOpts, transferId)
}

// GetState is a free data retrieval call binding the contract method 0xe35f157c.
//
// Solidity: function getState(bytes transferId) view returns(uint8)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCallerSession) GetState(transferId []byte) (uint8, error) {
	return _SafeTransferStorageImpl.Contract.GetState(&_SafeTransferStorageImpl.CallOpts, transferId)
}

// GetThisContractIndex is a free data retrieval call binding the contract method 0x70c6d367.
//
// Solidity: function getThisContractIndex() pure returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCaller) GetThisContractIndex(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SafeTransferStorageImpl.contract.Call(opts, out, "getThisContractIndex")
	return *ret0, err
}

// GetThisContractIndex is a free data retrieval call binding the contract method 0x70c6d367.
//
// Solidity: function getThisContractIndex() pure returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplSession) GetThisContractIndex() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.GetThisContractIndex(&_SafeTransferStorageImpl.CallOpts)
}

// GetThisContractIndex is a free data retrieval call binding the contract method 0x70c6d367.
//
// Solidity: function getThisContractIndex() pure returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCallerSession) GetThisContractIndex() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.GetThisContractIndex(&_SafeTransferStorageImpl.CallOpts)
}

// GetTransfer is a free data retrieval call binding the contract method 0xf7786d31.
//
// Solidity: function getTransfer(bytes transferId) view returns(uint8 state, bytes32 transferIdHash, address source, address dest, uint256 amount, bytes proof, bytes vk, uint256 createdAt, uint256 expiresAt, uint256 currentRetries, uint256 maxRetries, address canceller)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCaller) GetTransfer(opts *bind.CallOpts, transferId []byte) (struct {
	State          uint8
	TransferIdHash [32]byte
	Source         common.Address
	Dest           common.Address
	Amount         *big.Int
	Proof          []byte
	Vk             []byte
	CreatedAt      *big.Int
	ExpiresAt      *big.Int
	CurrentRetries *big.Int
	MaxRetries     *big.Int
	Canceller      common.Address
}, error) {
	ret := new(struct {
		State          uint8
		TransferIdHash [32]byte
		Source         common.Address
		Dest           common.Address
		Amount         *big.Int
		Proof          []byte
		Vk             []byte
		CreatedAt      *big.Int
		ExpiresAt      *big.Int
		CurrentRetries *big.Int
		MaxRetries     *big.Int
		Canceller      common.Address
	})
	out := ret
	err := _SafeTransferStorageImpl.contract.Call(opts, out, "getTransfer", transferId)
	return *ret, err
}

// GetTransfer is a free data retrieval call binding the contract method 0xf7786d31.
//
// Solidity: function getTransfer(bytes transferId) view returns(uint8 state, bytes32 transferIdHash, address source, address dest, uint256 amount, bytes proof, bytes vk, uint256 createdAt, uint256 expiresAt, uint256 currentRetries, uint256 maxRetries, address canceller)
func (_SafeTransferStorageImpl *SafeTransferStorageImplSession) GetTransfer(transferId []byte) (struct {
	State          uint8
	TransferIdHash [32]byte
	Source         common.Address
	Dest           common.Address
	Amount         *big.Int
	Proof          []byte
	Vk             []byte
	CreatedAt      *big.Int
	ExpiresAt      *big.Int
	CurrentRetries *big.Int
	MaxRetries     *big.Int
	Canceller      common.Address
}, error) {
	return _SafeTransferStorageImpl.Contract.GetTransfer(&_SafeTransferStorageImpl.CallOpts, transferId)
}

// GetTransfer is a free data retrieval call binding the contract method 0xf7786d31.
//
// Solidity: function getTransfer(bytes transferId) view returns(uint8 state, bytes32 transferIdHash, address source, address dest, uint256 amount, bytes proof, bytes vk, uint256 createdAt, uint256 expiresAt, uint256 currentRetries, uint256 maxRetries, address canceller)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCallerSession) GetTransfer(transferId []byte) (struct {
	State          uint8
	TransferIdHash [32]byte
	Source         common.Address
	Dest           common.Address
	Amount         *big.Int
	Proof          []byte
	Vk             []byte
	CreatedAt      *big.Int
	ExpiresAt      *big.Int
	CurrentRetries *big.Int
	MaxRetries     *big.Int
	Canceller      common.Address
}, error) {
	return _SafeTransferStorageImpl.Contract.GetTransfer(&_SafeTransferStorageImpl.CallOpts, transferId)
}

// GetTransferAmount is a free data retrieval call binding the contract method 0xeda33719.
//
// Solidity: function getTransferAmount(bytes transferId) view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCaller) GetTransferAmount(opts *bind.CallOpts, transferId []byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SafeTransferStorageImpl.contract.Call(opts, out, "getTransferAmount", transferId)
	return *ret0, err
}

// GetTransferAmount is a free data retrieval call binding the contract method 0xeda33719.
//
// Solidity: function getTransferAmount(bytes transferId) view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplSession) GetTransferAmount(transferId []byte) (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.GetTransferAmount(&_SafeTransferStorageImpl.CallOpts, transferId)
}

// GetTransferAmount is a free data retrieval call binding the contract method 0xeda33719.
//
// Solidity: function getTransferAmount(bytes transferId) view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCallerSession) GetTransferAmount(transferId []byte) (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.GetTransferAmount(&_SafeTransferStorageImpl.CallOpts, transferId)
}

// GetTransferAmountByHash is a free data retrieval call binding the contract method 0x27796898.
//
// Solidity: function getTransferAmountByHash(bytes32 transferIdHash) view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCaller) GetTransferAmountByHash(opts *bind.CallOpts, transferIdHash [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SafeTransferStorageImpl.contract.Call(opts, out, "getTransferAmountByHash", transferIdHash)
	return *ret0, err
}

// GetTransferAmountByHash is a free data retrieval call binding the contract method 0x27796898.
//
// Solidity: function getTransferAmountByHash(bytes32 transferIdHash) view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplSession) GetTransferAmountByHash(transferIdHash [32]byte) (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.GetTransferAmountByHash(&_SafeTransferStorageImpl.CallOpts, transferIdHash)
}

// GetTransferAmountByHash is a free data retrieval call binding the contract method 0x27796898.
//
// Solidity: function getTransferAmountByHash(bytes32 transferIdHash) view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCallerSession) GetTransferAmountByHash(transferIdHash [32]byte) (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.GetTransferAmountByHash(&_SafeTransferStorageImpl.CallOpts, transferIdHash)
}

// GetTransferByHash is a free data retrieval call binding the contract method 0xee290d9e.
//
// Solidity: function getTransferByHash(bytes32 transferIdHashString) view returns(uint8 state, bytes32 transferIdHash, address source, address dest, uint256 amount, bytes proof, bytes vk, uint256 createdAt, uint256 expiresAt, uint256 currentRetries, uint256 maxRetries, address canceller)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCaller) GetTransferByHash(opts *bind.CallOpts, transferIdHashString [32]byte) (struct {
	State          uint8
	TransferIdHash [32]byte
	Source         common.Address
	Dest           common.Address
	Amount         *big.Int
	Proof          []byte
	Vk             []byte
	CreatedAt      *big.Int
	ExpiresAt      *big.Int
	CurrentRetries *big.Int
	MaxRetries     *big.Int
	Canceller      common.Address
}, error) {
	ret := new(struct {
		State          uint8
		TransferIdHash [32]byte
		Source         common.Address
		Dest           common.Address
		Amount         *big.Int
		Proof          []byte
		Vk             []byte
		CreatedAt      *big.Int
		ExpiresAt      *big.Int
		CurrentRetries *big.Int
		MaxRetries     *big.Int
		Canceller      common.Address
	})
	out := ret
	err := _SafeTransferStorageImpl.contract.Call(opts, out, "getTransferByHash", transferIdHashString)
	return *ret, err
}

// GetTransferByHash is a free data retrieval call binding the contract method 0xee290d9e.
//
// Solidity: function getTransferByHash(bytes32 transferIdHashString) view returns(uint8 state, bytes32 transferIdHash, address source, address dest, uint256 amount, bytes proof, bytes vk, uint256 createdAt, uint256 expiresAt, uint256 currentRetries, uint256 maxRetries, address canceller)
func (_SafeTransferStorageImpl *SafeTransferStorageImplSession) GetTransferByHash(transferIdHashString [32]byte) (struct {
	State          uint8
	TransferIdHash [32]byte
	Source         common.Address
	Dest           common.Address
	Amount         *big.Int
	Proof          []byte
	Vk             []byte
	CreatedAt      *big.Int
	ExpiresAt      *big.Int
	CurrentRetries *big.Int
	MaxRetries     *big.Int
	Canceller      common.Address
}, error) {
	return _SafeTransferStorageImpl.Contract.GetTransferByHash(&_SafeTransferStorageImpl.CallOpts, transferIdHashString)
}

// GetTransferByHash is a free data retrieval call binding the contract method 0xee290d9e.
//
// Solidity: function getTransferByHash(bytes32 transferIdHashString) view returns(uint8 state, bytes32 transferIdHash, address source, address dest, uint256 amount, bytes proof, bytes vk, uint256 createdAt, uint256 expiresAt, uint256 currentRetries, uint256 maxRetries, address canceller)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCallerSession) GetTransferByHash(transferIdHashString [32]byte) (struct {
	State          uint8
	TransferIdHash [32]byte
	Source         common.Address
	Dest           common.Address
	Amount         *big.Int
	Proof          []byte
	Vk             []byte
	CreatedAt      *big.Int
	ExpiresAt      *big.Int
	CurrentRetries *big.Int
	MaxRetries     *big.Int
	Canceller      common.Address
}, error) {
	return _SafeTransferStorageImpl.Contract.GetTransferByHash(&_SafeTransferStorageImpl.CallOpts, transferIdHashString)
}

// GetTransferDest is a free data retrieval call binding the contract method 0x8d0c743e.
//
// Solidity: function getTransferDest(bytes transferId) view returns(address)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCaller) GetTransferDest(opts *bind.CallOpts, transferId []byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SafeTransferStorageImpl.contract.Call(opts, out, "getTransferDest", transferId)
	return *ret0, err
}

// GetTransferDest is a free data retrieval call binding the contract method 0x8d0c743e.
//
// Solidity: function getTransferDest(bytes transferId) view returns(address)
func (_SafeTransferStorageImpl *SafeTransferStorageImplSession) GetTransferDest(transferId []byte) (common.Address, error) {
	return _SafeTransferStorageImpl.Contract.GetTransferDest(&_SafeTransferStorageImpl.CallOpts, transferId)
}

// GetTransferDest is a free data retrieval call binding the contract method 0x8d0c743e.
//
// Solidity: function getTransferDest(bytes transferId) view returns(address)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCallerSession) GetTransferDest(transferId []byte) (common.Address, error) {
	return _SafeTransferStorageImpl.Contract.GetTransferDest(&_SafeTransferStorageImpl.CallOpts, transferId)
}

// GetTransferIdByHash is a free data retrieval call binding the contract method 0x903b8a72.
//
// Solidity: function getTransferIdByHash(bytes32 transferIdHash) view returns(bytes)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCaller) GetTransferIdByHash(opts *bind.CallOpts, transferIdHash [32]byte) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _SafeTransferStorageImpl.contract.Call(opts, out, "getTransferIdByHash", transferIdHash)
	return *ret0, err
}

// GetTransferIdByHash is a free data retrieval call binding the contract method 0x903b8a72.
//
// Solidity: function getTransferIdByHash(bytes32 transferIdHash) view returns(bytes)
func (_SafeTransferStorageImpl *SafeTransferStorageImplSession) GetTransferIdByHash(transferIdHash [32]byte) ([]byte, error) {
	return _SafeTransferStorageImpl.Contract.GetTransferIdByHash(&_SafeTransferStorageImpl.CallOpts, transferIdHash)
}

// GetTransferIdByHash is a free data retrieval call binding the contract method 0x903b8a72.
//
// Solidity: function getTransferIdByHash(bytes32 transferIdHash) view returns(bytes)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCallerSession) GetTransferIdByHash(transferIdHash [32]byte) ([]byte, error) {
	return _SafeTransferStorageImpl.Contract.GetTransferIdByHash(&_SafeTransferStorageImpl.CallOpts, transferIdHash)
}

// GetTransferLength is a free data retrieval call binding the contract method 0x36f3f96f.
//
// Solidity: function getTransferLength() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCaller) GetTransferLength(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SafeTransferStorageImpl.contract.Call(opts, out, "getTransferLength")
	return *ret0, err
}

// GetTransferLength is a free data retrieval call binding the contract method 0x36f3f96f.
//
// Solidity: function getTransferLength() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplSession) GetTransferLength() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.GetTransferLength(&_SafeTransferStorageImpl.CallOpts)
}

// GetTransferLength is a free data retrieval call binding the contract method 0x36f3f96f.
//
// Solidity: function getTransferLength() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCallerSession) GetTransferLength() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.GetTransferLength(&_SafeTransferStorageImpl.CallOpts)
}

// GetTransferSource is a free data retrieval call binding the contract method 0x1827be86.
//
// Solidity: function getTransferSource(bytes transferId) view returns(address)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCaller) GetTransferSource(opts *bind.CallOpts, transferId []byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SafeTransferStorageImpl.contract.Call(opts, out, "getTransferSource", transferId)
	return *ret0, err
}

// GetTransferSource is a free data retrieval call binding the contract method 0x1827be86.
//
// Solidity: function getTransferSource(bytes transferId) view returns(address)
func (_SafeTransferStorageImpl *SafeTransferStorageImplSession) GetTransferSource(transferId []byte) (common.Address, error) {
	return _SafeTransferStorageImpl.Contract.GetTransferSource(&_SafeTransferStorageImpl.CallOpts, transferId)
}

// GetTransferSource is a free data retrieval call binding the contract method 0x1827be86.
//
// Solidity: function getTransferSource(bytes transferId) view returns(address)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCallerSession) GetTransferSource(transferId []byte) (common.Address, error) {
	return _SafeTransferStorageImpl.Contract.GetTransferSource(&_SafeTransferStorageImpl.CallOpts, transferId)
}

// GetTransferSourceByHash is a free data retrieval call binding the contract method 0xa2f753cc.
//
// Solidity: function getTransferSourceByHash(bytes32 transferIdHash) view returns(address)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCaller) GetTransferSourceByHash(opts *bind.CallOpts, transferIdHash [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SafeTransferStorageImpl.contract.Call(opts, out, "getTransferSourceByHash", transferIdHash)
	return *ret0, err
}

// GetTransferSourceByHash is a free data retrieval call binding the contract method 0xa2f753cc.
//
// Solidity: function getTransferSourceByHash(bytes32 transferIdHash) view returns(address)
func (_SafeTransferStorageImpl *SafeTransferStorageImplSession) GetTransferSourceByHash(transferIdHash [32]byte) (common.Address, error) {
	return _SafeTransferStorageImpl.Contract.GetTransferSourceByHash(&_SafeTransferStorageImpl.CallOpts, transferIdHash)
}

// GetTransferSourceByHash is a free data retrieval call binding the contract method 0xa2f753cc.
//
// Solidity: function getTransferSourceByHash(bytes32 transferIdHash) view returns(address)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCallerSession) GetTransferSourceByHash(transferIdHash [32]byte) (common.Address, error) {
	return _SafeTransferStorageImpl.Contract.GetTransferSourceByHash(&_SafeTransferStorageImpl.CallOpts, transferIdHash)
}

// GetTransfers is a free data retrieval call binding the contract method 0x2f65142c.
//
// Solidity: function getTransfers() view returns(bytes32[])
func (_SafeTransferStorageImpl *SafeTransferStorageImplCaller) GetTransfers(opts *bind.CallOpts) ([][32]byte, error) {
	var (
		ret0 = new([][32]byte)
	)
	out := ret0
	err := _SafeTransferStorageImpl.contract.Call(opts, out, "getTransfers")
	return *ret0, err
}

// GetTransfers is a free data retrieval call binding the contract method 0x2f65142c.
//
// Solidity: function getTransfers() view returns(bytes32[])
func (_SafeTransferStorageImpl *SafeTransferStorageImplSession) GetTransfers() ([][32]byte, error) {
	return _SafeTransferStorageImpl.Contract.GetTransfers(&_SafeTransferStorageImpl.CallOpts)
}

// GetTransfers is a free data retrieval call binding the contract method 0x2f65142c.
//
// Solidity: function getTransfers() view returns(bytes32[])
func (_SafeTransferStorageImpl *SafeTransferStorageImplCallerSession) GetTransfers() ([][32]byte, error) {
	return _SafeTransferStorageImpl.Contract.GetTransfers(&_SafeTransferStorageImpl.CallOpts)
}

// IsCancelled is a free data retrieval call binding the contract method 0x703fcdd4.
//
// Solidity: function isCancelled(bytes transferId) view returns(bool)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCaller) IsCancelled(opts *bind.CallOpts, transferId []byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SafeTransferStorageImpl.contract.Call(opts, out, "isCancelled", transferId)
	return *ret0, err
}

// IsCancelled is a free data retrieval call binding the contract method 0x703fcdd4.
//
// Solidity: function isCancelled(bytes transferId) view returns(bool)
func (_SafeTransferStorageImpl *SafeTransferStorageImplSession) IsCancelled(transferId []byte) (bool, error) {
	return _SafeTransferStorageImpl.Contract.IsCancelled(&_SafeTransferStorageImpl.CallOpts, transferId)
}

// IsCancelled is a free data retrieval call binding the contract method 0x703fcdd4.
//
// Solidity: function isCancelled(bytes transferId) view returns(bool)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCallerSession) IsCancelled(transferId []byte) (bool, error) {
	return _SafeTransferStorageImpl.Contract.IsCancelled(&_SafeTransferStorageImpl.CallOpts, transferId)
}

// IsExpired is a free data retrieval call binding the contract method 0xc1f18e38.
//
// Solidity: function isExpired(bytes transferId) view returns(bool)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCaller) IsExpired(opts *bind.CallOpts, transferId []byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SafeTransferStorageImpl.contract.Call(opts, out, "isExpired", transferId)
	return *ret0, err
}

// IsExpired is a free data retrieval call binding the contract method 0xc1f18e38.
//
// Solidity: function isExpired(bytes transferId) view returns(bool)
func (_SafeTransferStorageImpl *SafeTransferStorageImplSession) IsExpired(transferId []byte) (bool, error) {
	return _SafeTransferStorageImpl.Contract.IsExpired(&_SafeTransferStorageImpl.CallOpts, transferId)
}

// IsExpired is a free data retrieval call binding the contract method 0xc1f18e38.
//
// Solidity: function isExpired(bytes transferId) view returns(bool)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCallerSession) IsExpired(transferId []byte) (bool, error) {
	return _SafeTransferStorageImpl.Contract.IsExpired(&_SafeTransferStorageImpl.CallOpts, transferId)
}

// IsExpiredByHash is a free data retrieval call binding the contract method 0x56657500.
//
// Solidity: function isExpiredByHash(bytes32 transferIdHash) view returns(bool)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCaller) IsExpiredByHash(opts *bind.CallOpts, transferIdHash [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SafeTransferStorageImpl.contract.Call(opts, out, "isExpiredByHash", transferIdHash)
	return *ret0, err
}

// IsExpiredByHash is a free data retrieval call binding the contract method 0x56657500.
//
// Solidity: function isExpiredByHash(bytes32 transferIdHash) view returns(bool)
func (_SafeTransferStorageImpl *SafeTransferStorageImplSession) IsExpiredByHash(transferIdHash [32]byte) (bool, error) {
	return _SafeTransferStorageImpl.Contract.IsExpiredByHash(&_SafeTransferStorageImpl.CallOpts, transferIdHash)
}

// IsExpiredByHash is a free data retrieval call binding the contract method 0x56657500.
//
// Solidity: function isExpiredByHash(bytes32 transferIdHash) view returns(bool)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCallerSession) IsExpiredByHash(transferIdHash [32]byte) (bool, error) {
	return _SafeTransferStorageImpl.Contract.IsExpiredByHash(&_SafeTransferStorageImpl.CallOpts, transferIdHash)
}

// IsPending is a free data retrieval call binding the contract method 0xb3adb4fa.
//
// Solidity: function isPending(bytes transferId) view returns(bool)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCaller) IsPending(opts *bind.CallOpts, transferId []byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SafeTransferStorageImpl.contract.Call(opts, out, "isPending", transferId)
	return *ret0, err
}

// IsPending is a free data retrieval call binding the contract method 0xb3adb4fa.
//
// Solidity: function isPending(bytes transferId) view returns(bool)
func (_SafeTransferStorageImpl *SafeTransferStorageImplSession) IsPending(transferId []byte) (bool, error) {
	return _SafeTransferStorageImpl.Contract.IsPending(&_SafeTransferStorageImpl.CallOpts, transferId)
}

// IsPending is a free data retrieval call binding the contract method 0xb3adb4fa.
//
// Solidity: function isPending(bytes transferId) view returns(bool)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCallerSession) IsPending(transferId []byte) (bool, error) {
	return _SafeTransferStorageImpl.Contract.IsPending(&_SafeTransferStorageImpl.CallOpts, transferId)
}

// IsPendingByHash is a free data retrieval call binding the contract method 0x59234057.
//
// Solidity: function isPendingByHash(bytes32 transferIdHash) view returns(bool)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCaller) IsPendingByHash(opts *bind.CallOpts, transferIdHash [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SafeTransferStorageImpl.contract.Call(opts, out, "isPendingByHash", transferIdHash)
	return *ret0, err
}

// IsPendingByHash is a free data retrieval call binding the contract method 0x59234057.
//
// Solidity: function isPendingByHash(bytes32 transferIdHash) view returns(bool)
func (_SafeTransferStorageImpl *SafeTransferStorageImplSession) IsPendingByHash(transferIdHash [32]byte) (bool, error) {
	return _SafeTransferStorageImpl.Contract.IsPendingByHash(&_SafeTransferStorageImpl.CallOpts, transferIdHash)
}

// IsPendingByHash is a free data retrieval call binding the contract method 0x59234057.
//
// Solidity: function isPendingByHash(bytes32 transferIdHash) view returns(bool)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCallerSession) IsPendingByHash(transferIdHash [32]byte) (bool, error) {
	return _SafeTransferStorageImpl.Contract.IsPendingByHash(&_SafeTransferStorageImpl.CallOpts, transferIdHash)
}

// IsSuccessful is a free data retrieval call binding the contract method 0x2cd03ea5.
//
// Solidity: function isSuccessful(bytes transferId) view returns(bool)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCaller) IsSuccessful(opts *bind.CallOpts, transferId []byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SafeTransferStorageImpl.contract.Call(opts, out, "isSuccessful", transferId)
	return *ret0, err
}

// IsSuccessful is a free data retrieval call binding the contract method 0x2cd03ea5.
//
// Solidity: function isSuccessful(bytes transferId) view returns(bool)
func (_SafeTransferStorageImpl *SafeTransferStorageImplSession) IsSuccessful(transferId []byte) (bool, error) {
	return _SafeTransferStorageImpl.Contract.IsSuccessful(&_SafeTransferStorageImpl.CallOpts, transferId)
}

// IsSuccessful is a free data retrieval call binding the contract method 0x2cd03ea5.
//
// Solidity: function isSuccessful(bytes transferId) view returns(bool)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCallerSession) IsSuccessful(transferId []byte) (bool, error) {
	return _SafeTransferStorageImpl.Contract.IsSuccessful(&_SafeTransferStorageImpl.CallOpts, transferId)
}

// OldestPending is a free data retrieval call binding the contract method 0x10a5dd5e.
//
// Solidity: function oldestPending() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCaller) OldestPending(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SafeTransferStorageImpl.contract.Call(opts, out, "oldestPending")
	return *ret0, err
}

// OldestPending is a free data retrieval call binding the contract method 0x10a5dd5e.
//
// Solidity: function oldestPending() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplSession) OldestPending() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.OldestPending(&_SafeTransferStorageImpl.CallOpts)
}

// OldestPending is a free data retrieval call binding the contract method 0x10a5dd5e.
//
// Solidity: function oldestPending() view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCallerSession) OldestPending() (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.OldestPending(&_SafeTransferStorageImpl.CallOpts)
}

// SearchOldestPending is a free data retrieval call binding the contract method 0xaaa27718.
//
// Solidity: function searchOldestPending(uint256 first, uint256 last) view returns(uint256 index, bool isPresent)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCaller) SearchOldestPending(opts *bind.CallOpts, first *big.Int, last *big.Int) (struct {
	Index     *big.Int
	IsPresent bool
}, error) {
	ret := new(struct {
		Index     *big.Int
		IsPresent bool
	})
	out := ret
	err := _SafeTransferStorageImpl.contract.Call(opts, out, "searchOldestPending", first, last)
	return *ret, err
}

// SearchOldestPending is a free data retrieval call binding the contract method 0xaaa27718.
//
// Solidity: function searchOldestPending(uint256 first, uint256 last) view returns(uint256 index, bool isPresent)
func (_SafeTransferStorageImpl *SafeTransferStorageImplSession) SearchOldestPending(first *big.Int, last *big.Int) (struct {
	Index     *big.Int
	IsPresent bool
}, error) {
	return _SafeTransferStorageImpl.Contract.SearchOldestPending(&_SafeTransferStorageImpl.CallOpts, first, last)
}

// SearchOldestPending is a free data retrieval call binding the contract method 0xaaa27718.
//
// Solidity: function searchOldestPending(uint256 first, uint256 last) view returns(uint256 index, bool isPresent)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCallerSession) SearchOldestPending(first *big.Int, last *big.Int) (struct {
	Index     *big.Int
	IsPresent bool
}, error) {
	return _SafeTransferStorageImpl.Contract.SearchOldestPending(&_SafeTransferStorageImpl.CallOpts, first, last)
}

// TransferExists is a free data retrieval call binding the contract method 0x168f3e26.
//
// Solidity: function transferExists(bytes transferId) view returns(bool)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCaller) TransferExists(opts *bind.CallOpts, transferId []byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SafeTransferStorageImpl.contract.Call(opts, out, "transferExists", transferId)
	return *ret0, err
}

// TransferExists is a free data retrieval call binding the contract method 0x168f3e26.
//
// Solidity: function transferExists(bytes transferId) view returns(bool)
func (_SafeTransferStorageImpl *SafeTransferStorageImplSession) TransferExists(transferId []byte) (bool, error) {
	return _SafeTransferStorageImpl.Contract.TransferExists(&_SafeTransferStorageImpl.CallOpts, transferId)
}

// TransferExists is a free data retrieval call binding the contract method 0x168f3e26.
//
// Solidity: function transferExists(bytes transferId) view returns(bool)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCallerSession) TransferExists(transferId []byte) (bool, error) {
	return _SafeTransferStorageImpl.Contract.TransferExists(&_SafeTransferStorageImpl.CallOpts, transferId)
}

// TransferIdOrig is a free data retrieval call binding the contract method 0xdb775dca.
//
// Solidity: function transferIdOrig(bytes32 ) view returns(bytes)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCaller) TransferIdOrig(opts *bind.CallOpts, arg0 [32]byte) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _SafeTransferStorageImpl.contract.Call(opts, out, "transferIdOrig", arg0)
	return *ret0, err
}

// TransferIdOrig is a free data retrieval call binding the contract method 0xdb775dca.
//
// Solidity: function transferIdOrig(bytes32 ) view returns(bytes)
func (_SafeTransferStorageImpl *SafeTransferStorageImplSession) TransferIdOrig(arg0 [32]byte) ([]byte, error) {
	return _SafeTransferStorageImpl.Contract.TransferIdOrig(&_SafeTransferStorageImpl.CallOpts, arg0)
}

// TransferIdOrig is a free data retrieval call binding the contract method 0xdb775dca.
//
// Solidity: function transferIdOrig(bytes32 ) view returns(bytes)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCallerSession) TransferIdOrig(arg0 [32]byte) ([]byte, error) {
	return _SafeTransferStorageImpl.Contract.TransferIdOrig(&_SafeTransferStorageImpl.CallOpts, arg0)
}

// TransferIds is a free data retrieval call binding the contract method 0x512de5d9.
//
// Solidity: function transferIds(uint256 ) view returns(bytes32)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCaller) TransferIds(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _SafeTransferStorageImpl.contract.Call(opts, out, "transferIds", arg0)
	return *ret0, err
}

// TransferIds is a free data retrieval call binding the contract method 0x512de5d9.
//
// Solidity: function transferIds(uint256 ) view returns(bytes32)
func (_SafeTransferStorageImpl *SafeTransferStorageImplSession) TransferIds(arg0 *big.Int) ([32]byte, error) {
	return _SafeTransferStorageImpl.Contract.TransferIds(&_SafeTransferStorageImpl.CallOpts, arg0)
}

// TransferIds is a free data retrieval call binding the contract method 0x512de5d9.
//
// Solidity: function transferIds(uint256 ) view returns(bytes32)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCallerSession) TransferIds(arg0 *big.Int) ([32]byte, error) {
	return _SafeTransferStorageImpl.Contract.TransferIds(&_SafeTransferStorageImpl.CallOpts, arg0)
}

// Transfers is a free data retrieval call binding the contract method 0x3c64f04b.
//
// Solidity: function transfers(bytes32 ) view returns(uint8 state, bytes32 transferIdHash, address source, address dest, uint256 amount, bytes proof, bytes vk, uint256 createdAt, uint256 expiresAt, uint256 currentRetries, uint256 maxRetries, address canceller)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCaller) Transfers(opts *bind.CallOpts, arg0 [32]byte) (struct {
	State          uint8
	TransferIdHash [32]byte
	Source         common.Address
	Dest           common.Address
	Amount         *big.Int
	Proof          []byte
	Vk             []byte
	CreatedAt      *big.Int
	ExpiresAt      *big.Int
	CurrentRetries *big.Int
	MaxRetries     *big.Int
	Canceller      common.Address
}, error) {
	ret := new(struct {
		State          uint8
		TransferIdHash [32]byte
		Source         common.Address
		Dest           common.Address
		Amount         *big.Int
		Proof          []byte
		Vk             []byte
		CreatedAt      *big.Int
		ExpiresAt      *big.Int
		CurrentRetries *big.Int
		MaxRetries     *big.Int
		Canceller      common.Address
	})
	out := ret
	err := _SafeTransferStorageImpl.contract.Call(opts, out, "transfers", arg0)
	return *ret, err
}

// Transfers is a free data retrieval call binding the contract method 0x3c64f04b.
//
// Solidity: function transfers(bytes32 ) view returns(uint8 state, bytes32 transferIdHash, address source, address dest, uint256 amount, bytes proof, bytes vk, uint256 createdAt, uint256 expiresAt, uint256 currentRetries, uint256 maxRetries, address canceller)
func (_SafeTransferStorageImpl *SafeTransferStorageImplSession) Transfers(arg0 [32]byte) (struct {
	State          uint8
	TransferIdHash [32]byte
	Source         common.Address
	Dest           common.Address
	Amount         *big.Int
	Proof          []byte
	Vk             []byte
	CreatedAt      *big.Int
	ExpiresAt      *big.Int
	CurrentRetries *big.Int
	MaxRetries     *big.Int
	Canceller      common.Address
}, error) {
	return _SafeTransferStorageImpl.Contract.Transfers(&_SafeTransferStorageImpl.CallOpts, arg0)
}

// Transfers is a free data retrieval call binding the contract method 0x3c64f04b.
//
// Solidity: function transfers(bytes32 ) view returns(uint8 state, bytes32 transferIdHash, address source, address dest, uint256 amount, bytes proof, bytes vk, uint256 createdAt, uint256 expiresAt, uint256 currentRetries, uint256 maxRetries, address canceller)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCallerSession) Transfers(arg0 [32]byte) (struct {
	State          uint8
	TransferIdHash [32]byte
	Source         common.Address
	Dest           common.Address
	Amount         *big.Int
	Proof          []byte
	Vk             []byte
	CreatedAt      *big.Int
	ExpiresAt      *big.Int
	CurrentRetries *big.Int
	MaxRetries     *big.Int
	Canceller      common.Address
}, error) {
	return _SafeTransferStorageImpl.Contract.Transfers(&_SafeTransferStorageImpl.CallOpts, arg0)
}

// TransfersFee is a free data retrieval call binding the contract method 0x2b935c35.
//
// Solidity: function transfersFee(bytes32 ) view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCaller) TransfersFee(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SafeTransferStorageImpl.contract.Call(opts, out, "transfersFee", arg0)
	return *ret0, err
}

// TransfersFee is a free data retrieval call binding the contract method 0x2b935c35.
//
// Solidity: function transfersFee(bytes32 ) view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplSession) TransfersFee(arg0 [32]byte) (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.TransfersFee(&_SafeTransferStorageImpl.CallOpts, arg0)
}

// TransfersFee is a free data retrieval call binding the contract method 0x2b935c35.
//
// Solidity: function transfersFee(bytes32 ) view returns(uint256)
func (_SafeTransferStorageImpl *SafeTransferStorageImplCallerSession) TransfersFee(arg0 [32]byte) (*big.Int, error) {
	return _SafeTransferStorageImpl.Contract.TransfersFee(&_SafeTransferStorageImpl.CallOpts, arg0)
}

// AddSafeTransfer is a paid mutator transaction binding the contract method 0x38af6a91.
//
// Solidity: function addSafeTransfer(bytes transferId, address source, address dest, uint256 amount, uint256 feeAmount, bytes proof, bytes vk, uint256 expiresAt, uint256 maxRetries) returns(bool)
func (_SafeTransferStorageImpl *SafeTransferStorageImplTransactor) AddSafeTransfer(opts *bind.TransactOpts, transferId []byte, source common.Address, dest common.Address, amount *big.Int, feeAmount *big.Int, proof []byte, vk []byte, expiresAt *big.Int, maxRetries *big.Int) (*types.Transaction, error) {
	return _SafeTransferStorageImpl.contract.Transact(opts, "addSafeTransfer", transferId, source, dest, amount, feeAmount, proof, vk, expiresAt, maxRetries)
}

// AddSafeTransfer is a paid mutator transaction binding the contract method 0x38af6a91.
//
// Solidity: function addSafeTransfer(bytes transferId, address source, address dest, uint256 amount, uint256 feeAmount, bytes proof, bytes vk, uint256 expiresAt, uint256 maxRetries) returns(bool)
func (_SafeTransferStorageImpl *SafeTransferStorageImplSession) AddSafeTransfer(transferId []byte, source common.Address, dest common.Address, amount *big.Int, feeAmount *big.Int, proof []byte, vk []byte, expiresAt *big.Int, maxRetries *big.Int) (*types.Transaction, error) {
	return _SafeTransferStorageImpl.Contract.AddSafeTransfer(&_SafeTransferStorageImpl.TransactOpts, transferId, source, dest, amount, feeAmount, proof, vk, expiresAt, maxRetries)
}

// AddSafeTransfer is a paid mutator transaction binding the contract method 0x38af6a91.
//
// Solidity: function addSafeTransfer(bytes transferId, address source, address dest, uint256 amount, uint256 feeAmount, bytes proof, bytes vk, uint256 expiresAt, uint256 maxRetries) returns(bool)
func (_SafeTransferStorageImpl *SafeTransferStorageImplTransactorSession) AddSafeTransfer(transferId []byte, source common.Address, dest common.Address, amount *big.Int, feeAmount *big.Int, proof []byte, vk []byte, expiresAt *big.Int, maxRetries *big.Int) (*types.Transaction, error) {
	return _SafeTransferStorageImpl.Contract.AddSafeTransfer(&_SafeTransferStorageImpl.TransactOpts, transferId, source, dest, amount, feeAmount, proof, vk, expiresAt, maxRetries)
}

// CancelTransfer is a paid mutator transaction binding the contract method 0xc98acd03.
//
// Solidity: function cancelTransfer(bytes transferId) returns(bool)
func (_SafeTransferStorageImpl *SafeTransferStorageImplTransactor) CancelTransfer(opts *bind.TransactOpts, transferId []byte) (*types.Transaction, error) {
	return _SafeTransferStorageImpl.contract.Transact(opts, "cancelTransfer", transferId)
}

// CancelTransfer is a paid mutator transaction binding the contract method 0xc98acd03.
//
// Solidity: function cancelTransfer(bytes transferId) returns(bool)
func (_SafeTransferStorageImpl *SafeTransferStorageImplSession) CancelTransfer(transferId []byte) (*types.Transaction, error) {
	return _SafeTransferStorageImpl.Contract.CancelTransfer(&_SafeTransferStorageImpl.TransactOpts, transferId)
}

// CancelTransfer is a paid mutator transaction binding the contract method 0xc98acd03.
//
// Solidity: function cancelTransfer(bytes transferId) returns(bool)
func (_SafeTransferStorageImpl *SafeTransferStorageImplTransactorSession) CancelTransfer(transferId []byte) (*types.Transaction, error) {
	return _SafeTransferStorageImpl.Contract.CancelTransfer(&_SafeTransferStorageImpl.TransactOpts, transferId)
}

// ExpireTransfer is a paid mutator transaction binding the contract method 0x6cf522b5.
//
// Solidity: function expireTransfer(bytes transferId) returns(bool)
func (_SafeTransferStorageImpl *SafeTransferStorageImplTransactor) ExpireTransfer(opts *bind.TransactOpts, transferId []byte) (*types.Transaction, error) {
	return _SafeTransferStorageImpl.contract.Transact(opts, "expireTransfer", transferId)
}

// ExpireTransfer is a paid mutator transaction binding the contract method 0x6cf522b5.
//
// Solidity: function expireTransfer(bytes transferId) returns(bool)
func (_SafeTransferStorageImpl *SafeTransferStorageImplSession) ExpireTransfer(transferId []byte) (*types.Transaction, error) {
	return _SafeTransferStorageImpl.Contract.ExpireTransfer(&_SafeTransferStorageImpl.TransactOpts, transferId)
}

// ExpireTransfer is a paid mutator transaction binding the contract method 0x6cf522b5.
//
// Solidity: function expireTransfer(bytes transferId) returns(bool)
func (_SafeTransferStorageImpl *SafeTransferStorageImplTransactorSession) ExpireTransfer(transferId []byte) (*types.Transaction, error) {
	return _SafeTransferStorageImpl.Contract.ExpireTransfer(&_SafeTransferStorageImpl.TransactOpts, transferId)
}

// ExpireTransferByHash is a paid mutator transaction binding the contract method 0x32af9f6c.
//
// Solidity: function expireTransferByHash(bytes32 transferIdHash) returns(bool)
func (_SafeTransferStorageImpl *SafeTransferStorageImplTransactor) ExpireTransferByHash(opts *bind.TransactOpts, transferIdHash [32]byte) (*types.Transaction, error) {
	return _SafeTransferStorageImpl.contract.Transact(opts, "expireTransferByHash", transferIdHash)
}

// ExpireTransferByHash is a paid mutator transaction binding the contract method 0x32af9f6c.
//
// Solidity: function expireTransferByHash(bytes32 transferIdHash) returns(bool)
func (_SafeTransferStorageImpl *SafeTransferStorageImplSession) ExpireTransferByHash(transferIdHash [32]byte) (*types.Transaction, error) {
	return _SafeTransferStorageImpl.Contract.ExpireTransferByHash(&_SafeTransferStorageImpl.TransactOpts, transferIdHash)
}

// ExpireTransferByHash is a paid mutator transaction binding the contract method 0x32af9f6c.
//
// Solidity: function expireTransferByHash(bytes32 transferIdHash) returns(bool)
func (_SafeTransferStorageImpl *SafeTransferStorageImplTransactorSession) ExpireTransferByHash(transferIdHash [32]byte) (*types.Transaction, error) {
	return _SafeTransferStorageImpl.Contract.ExpireTransferByHash(&_SafeTransferStorageImpl.TransactOpts, transferIdHash)
}

// SetExecutedState is a paid mutator transaction binding the contract method 0x1468e1c4.
//
// Solidity: function setExecutedState(bytes transferId) returns()
func (_SafeTransferStorageImpl *SafeTransferStorageImplTransactor) SetExecutedState(opts *bind.TransactOpts, transferId []byte) (*types.Transaction, error) {
	return _SafeTransferStorageImpl.contract.Transact(opts, "setExecutedState", transferId)
}

// SetExecutedState is a paid mutator transaction binding the contract method 0x1468e1c4.
//
// Solidity: function setExecutedState(bytes transferId) returns()
func (_SafeTransferStorageImpl *SafeTransferStorageImplSession) SetExecutedState(transferId []byte) (*types.Transaction, error) {
	return _SafeTransferStorageImpl.Contract.SetExecutedState(&_SafeTransferStorageImpl.TransactOpts, transferId)
}

// SetExecutedState is a paid mutator transaction binding the contract method 0x1468e1c4.
//
// Solidity: function setExecutedState(bytes transferId) returns()
func (_SafeTransferStorageImpl *SafeTransferStorageImplTransactorSession) SetExecutedState(transferId []byte) (*types.Transaction, error) {
	return _SafeTransferStorageImpl.Contract.SetExecutedState(&_SafeTransferStorageImpl.TransactOpts, transferId)
}

// SetFailedState is a paid mutator transaction binding the contract method 0x006ec50b.
//
// Solidity: function setFailedState(bytes transferId) returns()
func (_SafeTransferStorageImpl *SafeTransferStorageImplTransactor) SetFailedState(opts *bind.TransactOpts, transferId []byte) (*types.Transaction, error) {
	return _SafeTransferStorageImpl.contract.Transact(opts, "setFailedState", transferId)
}

// SetFailedState is a paid mutator transaction binding the contract method 0x006ec50b.
//
// Solidity: function setFailedState(bytes transferId) returns()
func (_SafeTransferStorageImpl *SafeTransferStorageImplSession) SetFailedState(transferId []byte) (*types.Transaction, error) {
	return _SafeTransferStorageImpl.Contract.SetFailedState(&_SafeTransferStorageImpl.TransactOpts, transferId)
}

// SetFailedState is a paid mutator transaction binding the contract method 0x006ec50b.
//
// Solidity: function setFailedState(bytes transferId) returns()
func (_SafeTransferStorageImpl *SafeTransferStorageImplTransactorSession) SetFailedState(transferId []byte) (*types.Transaction, error) {
	return _SafeTransferStorageImpl.Contract.SetFailedState(&_SafeTransferStorageImpl.TransactOpts, transferId)
}

// SetOldestPending is a paid mutator transaction binding the contract method 0x0b340e31.
//
// Solidity: function setOldestPending(uint256 oldestPendingIndex) returns()
func (_SafeTransferStorageImpl *SafeTransferStorageImplTransactor) SetOldestPending(opts *bind.TransactOpts, oldestPendingIndex *big.Int) (*types.Transaction, error) {
	return _SafeTransferStorageImpl.contract.Transact(opts, "setOldestPending", oldestPendingIndex)
}

// SetOldestPending is a paid mutator transaction binding the contract method 0x0b340e31.
//
// Solidity: function setOldestPending(uint256 oldestPendingIndex) returns()
func (_SafeTransferStorageImpl *SafeTransferStorageImplSession) SetOldestPending(oldestPendingIndex *big.Int) (*types.Transaction, error) {
	return _SafeTransferStorageImpl.Contract.SetOldestPending(&_SafeTransferStorageImpl.TransactOpts, oldestPendingIndex)
}

// SetOldestPending is a paid mutator transaction binding the contract method 0x0b340e31.
//
// Solidity: function setOldestPending(uint256 oldestPendingIndex) returns()
func (_SafeTransferStorageImpl *SafeTransferStorageImplTransactorSession) SetOldestPending(oldestPendingIndex *big.Int) (*types.Transaction, error) {
	return _SafeTransferStorageImpl.Contract.SetOldestPending(&_SafeTransferStorageImpl.TransactOpts, oldestPendingIndex)
}

// SafeTransferStorageImplTransferAddedIterator is returned from FilterTransferAdded and is used to iterate over the raw logs and unpacked data for TransferAdded events raised by the SafeTransferStorageImpl contract.
type SafeTransferStorageImplTransferAddedIterator struct {
	Event *SafeTransferStorageImplTransferAdded // Event containing the contract specifics and raw log

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
func (it *SafeTransferStorageImplTransferAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SafeTransferStorageImplTransferAdded)
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
		it.Event = new(SafeTransferStorageImplTransferAdded)
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
func (it *SafeTransferStorageImplTransferAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SafeTransferStorageImplTransferAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SafeTransferStorageImplTransferAdded represents a TransferAdded event raised by the SafeTransferStorageImpl contract.
type SafeTransferStorageImplTransferAdded struct {
	TransferId []byte
	Source     common.Address
	Dest       common.Address
	Amount     *big.Int
	ExpiresAt  *big.Int
	MaxRetries *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTransferAdded is a free log retrieval operation binding the contract event 0x61889633ca0f7246f5c4acc1891cd92065ddf1894e05e9920015dc5318b18a27.
//
// Solidity: event TransferAdded(bytes transferId, address source, address dest, uint256 amount, uint256 expiresAt, uint256 maxRetries)
func (_SafeTransferStorageImpl *SafeTransferStorageImplFilterer) FilterTransferAdded(opts *bind.FilterOpts) (*SafeTransferStorageImplTransferAddedIterator, error) {

	logs, sub, err := _SafeTransferStorageImpl.contract.FilterLogs(opts, "TransferAdded")
	if err != nil {
		return nil, err
	}
	return &SafeTransferStorageImplTransferAddedIterator{contract: _SafeTransferStorageImpl.contract, event: "TransferAdded", logs: logs, sub: sub}, nil
}

// WatchTransferAdded is a free log subscription operation binding the contract event 0x61889633ca0f7246f5c4acc1891cd92065ddf1894e05e9920015dc5318b18a27.
//
// Solidity: event TransferAdded(bytes transferId, address source, address dest, uint256 amount, uint256 expiresAt, uint256 maxRetries)
func (_SafeTransferStorageImpl *SafeTransferStorageImplFilterer) WatchTransferAdded(opts *bind.WatchOpts, sink chan<- *SafeTransferStorageImplTransferAdded) (event.Subscription, error) {

	logs, sub, err := _SafeTransferStorageImpl.contract.WatchLogs(opts, "TransferAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SafeTransferStorageImplTransferAdded)
				if err := _SafeTransferStorageImpl.contract.UnpackLog(event, "TransferAdded", log); err != nil {
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

// ParseTransferAdded is a log parse operation binding the contract event 0x61889633ca0f7246f5c4acc1891cd92065ddf1894e05e9920015dc5318b18a27.
//
// Solidity: event TransferAdded(bytes transferId, address source, address dest, uint256 amount, uint256 expiresAt, uint256 maxRetries)
func (_SafeTransferStorageImpl *SafeTransferStorageImplFilterer) ParseTransferAdded(log types.Log) (*SafeTransferStorageImplTransferAdded, error) {
	event := new(SafeTransferStorageImplTransferAdded)
	if err := _SafeTransferStorageImpl.contract.UnpackLog(event, "TransferAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SafeTransferStorageImplTransferCancelledIterator is returned from FilterTransferCancelled and is used to iterate over the raw logs and unpacked data for TransferCancelled events raised by the SafeTransferStorageImpl contract.
type SafeTransferStorageImplTransferCancelledIterator struct {
	Event *SafeTransferStorageImplTransferCancelled // Event containing the contract specifics and raw log

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
func (it *SafeTransferStorageImplTransferCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SafeTransferStorageImplTransferCancelled)
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
		it.Event = new(SafeTransferStorageImplTransferCancelled)
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
func (it *SafeTransferStorageImplTransferCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SafeTransferStorageImplTransferCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SafeTransferStorageImplTransferCancelled represents a TransferCancelled event raised by the SafeTransferStorageImpl contract.
type SafeTransferStorageImplTransferCancelled struct {
	TransferId []byte
	Canceller  common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTransferCancelled is a free log retrieval operation binding the contract event 0xb218c8a05e9932efae26d2f2c58dbc38c991451cb170427260817b5454134b21.
//
// Solidity: event TransferCancelled(bytes transferId, address canceller)
func (_SafeTransferStorageImpl *SafeTransferStorageImplFilterer) FilterTransferCancelled(opts *bind.FilterOpts) (*SafeTransferStorageImplTransferCancelledIterator, error) {

	logs, sub, err := _SafeTransferStorageImpl.contract.FilterLogs(opts, "TransferCancelled")
	if err != nil {
		return nil, err
	}
	return &SafeTransferStorageImplTransferCancelledIterator{contract: _SafeTransferStorageImpl.contract, event: "TransferCancelled", logs: logs, sub: sub}, nil
}

// WatchTransferCancelled is a free log subscription operation binding the contract event 0xb218c8a05e9932efae26d2f2c58dbc38c991451cb170427260817b5454134b21.
//
// Solidity: event TransferCancelled(bytes transferId, address canceller)
func (_SafeTransferStorageImpl *SafeTransferStorageImplFilterer) WatchTransferCancelled(opts *bind.WatchOpts, sink chan<- *SafeTransferStorageImplTransferCancelled) (event.Subscription, error) {

	logs, sub, err := _SafeTransferStorageImpl.contract.WatchLogs(opts, "TransferCancelled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SafeTransferStorageImplTransferCancelled)
				if err := _SafeTransferStorageImpl.contract.UnpackLog(event, "TransferCancelled", log); err != nil {
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

// ParseTransferCancelled is a log parse operation binding the contract event 0xb218c8a05e9932efae26d2f2c58dbc38c991451cb170427260817b5454134b21.
//
// Solidity: event TransferCancelled(bytes transferId, address canceller)
func (_SafeTransferStorageImpl *SafeTransferStorageImplFilterer) ParseTransferCancelled(log types.Log) (*SafeTransferStorageImplTransferCancelled, error) {
	event := new(SafeTransferStorageImplTransferCancelled)
	if err := _SafeTransferStorageImpl.contract.UnpackLog(event, "TransferCancelled", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SafeTransferStorageImplTransferExpiredIterator is returned from FilterTransferExpired and is used to iterate over the raw logs and unpacked data for TransferExpired events raised by the SafeTransferStorageImpl contract.
type SafeTransferStorageImplTransferExpiredIterator struct {
	Event *SafeTransferStorageImplTransferExpired // Event containing the contract specifics and raw log

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
func (it *SafeTransferStorageImplTransferExpiredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SafeTransferStorageImplTransferExpired)
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
		it.Event = new(SafeTransferStorageImplTransferExpired)
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
func (it *SafeTransferStorageImplTransferExpiredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SafeTransferStorageImplTransferExpiredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SafeTransferStorageImplTransferExpired represents a TransferExpired event raised by the SafeTransferStorageImpl contract.
type SafeTransferStorageImplTransferExpired struct {
	TransferId []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTransferExpired is a free log retrieval operation binding the contract event 0xa7217951342b9ae9225d6cf462243813f3d24ecea136332e91f42637e678be2d.
//
// Solidity: event TransferExpired(bytes transferId)
func (_SafeTransferStorageImpl *SafeTransferStorageImplFilterer) FilterTransferExpired(opts *bind.FilterOpts) (*SafeTransferStorageImplTransferExpiredIterator, error) {

	logs, sub, err := _SafeTransferStorageImpl.contract.FilterLogs(opts, "TransferExpired")
	if err != nil {
		return nil, err
	}
	return &SafeTransferStorageImplTransferExpiredIterator{contract: _SafeTransferStorageImpl.contract, event: "TransferExpired", logs: logs, sub: sub}, nil
}

// WatchTransferExpired is a free log subscription operation binding the contract event 0xa7217951342b9ae9225d6cf462243813f3d24ecea136332e91f42637e678be2d.
//
// Solidity: event TransferExpired(bytes transferId)
func (_SafeTransferStorageImpl *SafeTransferStorageImplFilterer) WatchTransferExpired(opts *bind.WatchOpts, sink chan<- *SafeTransferStorageImplTransferExpired) (event.Subscription, error) {

	logs, sub, err := _SafeTransferStorageImpl.contract.WatchLogs(opts, "TransferExpired")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SafeTransferStorageImplTransferExpired)
				if err := _SafeTransferStorageImpl.contract.UnpackLog(event, "TransferExpired", log); err != nil {
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

// ParseTransferExpired is a log parse operation binding the contract event 0xa7217951342b9ae9225d6cf462243813f3d24ecea136332e91f42637e678be2d.
//
// Solidity: event TransferExpired(bytes transferId)
func (_SafeTransferStorageImpl *SafeTransferStorageImplFilterer) ParseTransferExpired(log types.Log) (*SafeTransferStorageImplTransferExpired, error) {
	event := new(SafeTransferStorageImplTransferExpired)
	if err := _SafeTransferStorageImpl.contract.UnpackLog(event, "TransferExpired", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SafeTransferStorageImplTransferExpiredByHashIterator is returned from FilterTransferExpiredByHash and is used to iterate over the raw logs and unpacked data for TransferExpiredByHash events raised by the SafeTransferStorageImpl contract.
type SafeTransferStorageImplTransferExpiredByHashIterator struct {
	Event *SafeTransferStorageImplTransferExpiredByHash // Event containing the contract specifics and raw log

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
func (it *SafeTransferStorageImplTransferExpiredByHashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SafeTransferStorageImplTransferExpiredByHash)
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
		it.Event = new(SafeTransferStorageImplTransferExpiredByHash)
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
func (it *SafeTransferStorageImplTransferExpiredByHashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SafeTransferStorageImplTransferExpiredByHashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SafeTransferStorageImplTransferExpiredByHash represents a TransferExpiredByHash event raised by the SafeTransferStorageImpl contract.
type SafeTransferStorageImplTransferExpiredByHash struct {
	TransferIdHash [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterTransferExpiredByHash is a free log retrieval operation binding the contract event 0xb713225e6a8374d53acfc600983e9e1cf6810636c4c173595107383e549e2af1.
//
// Solidity: event TransferExpiredByHash(bytes32 transferIdHash)
func (_SafeTransferStorageImpl *SafeTransferStorageImplFilterer) FilterTransferExpiredByHash(opts *bind.FilterOpts) (*SafeTransferStorageImplTransferExpiredByHashIterator, error) {

	logs, sub, err := _SafeTransferStorageImpl.contract.FilterLogs(opts, "TransferExpiredByHash")
	if err != nil {
		return nil, err
	}
	return &SafeTransferStorageImplTransferExpiredByHashIterator{contract: _SafeTransferStorageImpl.contract, event: "TransferExpiredByHash", logs: logs, sub: sub}, nil
}

// WatchTransferExpiredByHash is a free log subscription operation binding the contract event 0xb713225e6a8374d53acfc600983e9e1cf6810636c4c173595107383e549e2af1.
//
// Solidity: event TransferExpiredByHash(bytes32 transferIdHash)
func (_SafeTransferStorageImpl *SafeTransferStorageImplFilterer) WatchTransferExpiredByHash(opts *bind.WatchOpts, sink chan<- *SafeTransferStorageImplTransferExpiredByHash) (event.Subscription, error) {

	logs, sub, err := _SafeTransferStorageImpl.contract.WatchLogs(opts, "TransferExpiredByHash")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SafeTransferStorageImplTransferExpiredByHash)
				if err := _SafeTransferStorageImpl.contract.UnpackLog(event, "TransferExpiredByHash", log); err != nil {
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

// ParseTransferExpiredByHash is a log parse operation binding the contract event 0xb713225e6a8374d53acfc600983e9e1cf6810636c4c173595107383e549e2af1.
//
// Solidity: event TransferExpiredByHash(bytes32 transferIdHash)
func (_SafeTransferStorageImpl *SafeTransferStorageImplFilterer) ParseTransferExpiredByHash(log types.Log) (*SafeTransferStorageImplTransferExpiredByHash, error) {
	event := new(SafeTransferStorageImplTransferExpiredByHash)
	if err := _SafeTransferStorageImpl.contract.UnpackLog(event, "TransferExpiredByHash", log); err != nil {
		return nil, err
	}
	return event, nil
}
