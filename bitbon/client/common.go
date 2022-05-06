// Copyright Simcord LLC
// This file is part of the Bitbon System Blockchain Node library.
//
// The Bitbon System Blockchain Node library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The Bitbon System Blockchain Node library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the Bitbon System Blockchain Node library. If not, see <http://www.gnu.org/licenses/>.

package client

import (
	"runtime"
	"time"

	"github.com/simcord-llc/bitbon-system-blockchain/log"
)

type BitbonClientMode int

const (
	BitbonClientUndefinedMode BitbonClientMode = iota
	BitbonClientAllMode
	BitbonClientMutableMode
	BitbonClientReadableMode
)

// Bitbon constants
const (
	ProtocolVersionStr = "1.0.0"        // The same, as a string
	ProtocolName       = "bitbonclient" // Nickname of the protocol in geth

	DefaultJournalName    = ".bitbon.client.journal.LDB"
	DefaultJournalTimeout = 10 * time.Minute
	DefaultConfirmTimeout = 5 * time.Second
	DefaultPublishTimeout = 3 * time.Second

	DefaultAssetboxExchange    = "e.assetbox.forward"
	DefaultTransferExchange    = "e.transfer.forward"
	DefaultBlockExchange       = "e.block.forward"
	DefaultTransactionExchange = "e.transaction.forward"
	DefaultMiningExchange      = "e.mining.forward"
	DefaultFeeExchange         = "e.fee.forward"

	LastErrorCleanInterval = 20 * time.Second

	// Rabbit routing keys
	prepareAssetboxRequestRK  = "r.assetbox.#.PrepareAssetboxesRequest"
	prepareAssetboxResponseRK = "r.assetbox.bitbon-node.PrepareAssetboxesResponse"

	deleteAssetboxRequestRK  = "r.assetbox.#.DeleteAssetboxRequest"
	deleteAssetboxResponseRK = "r.assetbox.bitbon-node.DeleteAssetboxResponse"

	setPublicAssetboxInfoRequestRK  = "r.assetbox.#.SetPublicAssetboxInfoRequest"
	setPublicAssetboxInfoResponseRK = "r.assetbox.bitbon-node.SetPublicAssetboxInfoResponse"

	assetboxBalancesRequestRK  = "r.assetbox.#.AssetboxBalancesRequest"
	assetboxBalancesResponseRK = "r.assetbox.bitbon-node.AssetboxBalancesResponse"

	quickTransferRequestRK  = "r.transfer.#.QuickTransferRequest"
	quickTransferResponseRK = "r.transfer.bitbon-node.QuickTransferResponse"

	frameTransferRequestRK  = "r.transfer.#.FrameTransferRequest"
	frameTransferResponseRK = "r.transfer.bitbon-node.FrameTransferResponse"

	serviceFeeTransferRequestRK  = "r.transfer.#.ServiceFeeTransferRequest"
	serviceFeeTransferResponseRK = "r.transfer.bitbon-node.ServiceFeeTransferResponse"

	createWPCSafeTransferRequestRK  = "r.transfer.#.CreateWPCSafeTransferRequest"
	createWPCSafeTransferResponseRK = "r.transfer.bitbon-node.CreateWPCSafeTransferResponse"

	approveWPCSafeTransferRequestRK  = "r.transfer.#.ApproveWPCSafeTransferRequest"
	approveWPCSafeTransferResponseRK = "r.transfer.bitbon-node.ApproveWPCSafeTransferResponse"

	cancelWPCSafeTransferRequestRK  = "r.transfer.#.CancelWPCSafeTransferRequest"
	cancelWPCSafeTransferResponseRK = "r.transfer.bitbon-node.CancelWPCSafeTransferResponse"

	createSafeTransferRequestRK  = "r.transfer.#.CreateSafeTransferRequest"
	createSafeTransferResponseRK = "r.transfer.bitbon-node.CreateSafeTransferResponse"

	approveSafeTransferRequestRK  = "r.transfer.#.ApproveSafeTransferRequest"
	approveSafeTransferResponseRK = "r.transfer.bitbon-node.ApproveSafeTransferResponse"

	cancelSafeTransferRequestRK  = "r.transfer.#.CancelSafeTransferRequest"
	cancelSafeTransferResponseRK = "r.transfer.bitbon-node.CancelSafeTransferResponse"

	fullBalanceQuickTransferRequestRK  = "r.transfer.#.FullBalanceQuickTransferRequest"
	fullBalanceQuickTransferResponseRK = "r.transfer.bitbon-node.FullBalanceQuickTransferResponse"

	createFullBalanceWPCSafeTransferRequestRK  = "r.transfer.#.CreateFullBalanceWPCSafeTransferRequest"
	createFullBalanceWPCSafeTransferResponseRK = "r.transfer.bitbon-node.CreateFullBalanceWPCSafeTransferResponse"

	approveFullBalanceWPCSafeTransferRequestRK  = "r.transfer.#.ApproveFullBalanceWPCSafeTransferRequest"
	approveFullBalanceWPCSafeTransferResponseRK = "r.transfer.bitbon-node.ApproveFullBalanceWPCSafeTransferResponse"

	cancelFullBalanceWPCSafeTransferRequestRK  = "r.transfer.#.CancelFullBalanceWPCSafeTransferRequest"
	cancelFullBalanceWPCSafeTransferResponseRK = "r.transfer.bitbon-node.CancelFullBalanceWPCSafeTransferResponse"

	createFullBalanceSafeTransferRequestRK  = "r.transfer.#.CreateFullBalanceSafeTransferRequest"
	createFullBalanceSafeTransferResponseRK = "r.transfer.bitbon-node.CreateFullBalanceSafeTransferResponse"

	approveFullBalanceSafeTransferRequestRK  = "r.transfer.#.ApproveFullBalanceSafeTransferRequest"
	approveFullBalanceSafeTransferResponseRK = "r.transfer.bitbon-node.ApproveFullBalanceSafeTransferResponse"

	cancelFullBalanceSafeTransferRequestRK  = "r.transfer.#.CancelFullBalanceSafeTransferRequest"
	cancelFullBalanceSafeTransferResponseRK = "r.transfer.bitbon-node.CancelFullBalanceSafeTransferResponse"

	directTransferRequestRK  = "r.transfer.#.DirectTransferRequest"
	directTransferResponseRK = "r.transfer.bitbon-node.DirectTransferResponse"

	quickTransferRequestAsyncRK  = "r.transfer.#.QuickTransferAsyncRequest"
	quickTransferResponseAsyncRK = "r.transfer.bitbon-node.QuickTransferAsyncResponse"

	frameTransferRequestAsyncRK  = "r.transfer.#.FrameTransferAsyncRequest"
	frameTransferResponseAsyncRK = "r.transfer.bitbon-node.FrameTransferAsyncResponse"

	serviceFeeTransferRequestAsyncRK  = "r.transfer.#.ServiceFeeTransferAsyncRequest"
	serviceFeeTransferResponseAsyncRK = "r.transfer.bitbon-node.ServiceFeeTransferAsyncResponse"

	createWPCSafeTransferRequestAsyncRK  = "r.transfer.#.CreateWPCSafeTransferAsyncRequest"
	createWPCSafeTransferResponseAsyncRK = "r.transfer.bitbon-node.CreateWPCSafeTransferAsyncResponse"

	approveWPCSafeTransferRequestAsyncRK  = "r.transfer.#.ApproveWPCSafeTransferAsyncRequest"
	approveWPCSafeTransferResponseAsyncRK = "r.transfer.bitbon-node.ApproveWPCSafeTransferAsyncResponse"

	cancelWPCSafeTransferRequestAsyncRK  = "r.transfer.#.CancelWPCSafeTransferAsyncRequest"
	cancelWPCSafeTransferResponseAsyncRK = "r.transfer.bitbon-node.CancelWPCSafeTransferAsyncResponse"

	createSafeTransferRequestAsyncRK  = "r.transfer.#.CreateSafeTransferAsyncRequest"
	createSafeTransferResponseAsyncRK = "r.transfer.bitbon-node.CreateSafeTransferAsyncResponse"

	approveSafeTransferRequestAsyncRK  = "r.transfer.#.ApproveSafeTransferAsyncRequest"
	approveSafeTransferResponseAsyncRK = "r.transfer.bitbon-node.ApproveSafeTransferAsyncResponse"

	cancelSafeTransferRequestAsyncRK  = "r.transfer.#.CancelSafeTransferAsyncRequest"
	cancelSafeTransferResponseAsyncRK = "r.transfer.bitbon-node.CancelSafeTransferAsyncResponse"

	fullBalanceQuickTransferRequestAsyncRK  = "r.transfer.#.FullBalanceQuickTransferAsyncRequest"
	fullBalanceQuickTransferResponseAsyncRK = "r.transfer.bitbon-node.FullBalanceQuickTransferAsyncResponse"

	createFullBalanceWPCSafeTransferRequestAsyncRK  = "r.transfer.#.CreateFullBalanceWPCSafeTransferAsyncRequest"
	createFullBalanceWPCSafeTransferResponseAsyncRK = "r.transfer.bitbon-node.CreateFullBalanceWPCSafeTransferAsyncResponse"

	approveFullBalanceWPCSafeTransferRequestAsyncRK  = "r.transfer.#.ApproveFullBalanceWPCSafeTransferAsyncRequest"
	approveFullBalanceWPCSafeTransferResponseAsyncRK = "r.transfer.bitbon-node.ApproveFullBalanceWPCSafeTransferAsyncResponse"

	cancelFullBalanceWPCSafeTransferRequestAsyncRK  = "r.transfer.#.CancelFullBalanceWPCSafeTransferAsyncRequest"
	cancelFullBalanceWPCSafeTransferResponseAsyncRK = "r.transfer.bitbon-node.CancelFullBalanceWPCSafeTransferAsyncResponse"

	createFullBalanceSafeTransferRequestAsyncRK  = "r.transfer.#.CreateFullBalanceSafeTransferAsyncRequest"
	createFullBalanceSafeTransferResponseAsyncRK = "r.transfer.bitbon-node.CreateFullBalanceSafeTransferAsyncResponse"

	approveFullBalanceSafeTransferRequestAsyncRK  = "r.transfer.#.ApproveFullBalanceSafeTransferAsyncRequest"
	approveFullBalanceSafeTransferResponseAsyncRK = "r.transfer.bitbon-node.ApproveFullBalanceSafeTransferAsyncResponse"

	cancelFullBalanceSafeTransferRequestAsyncRK  = "r.transfer.#.CancelFullBalanceSafeTransferAsyncRequest"
	cancelFullBalanceSafeTransferResponseAsyncRK = "r.transfer.bitbon-node.CancelFullBalanceSafeTransferAsyncResponse"

	directTransferRequestAsyncRK  = "r.transfer.#.DirectTransferAsyncRequest"
	directTransferResponseAsyncRK = "r.transfer.bitbon-node.DirectTransferAsyncResponse"

	expireTransfersRequestRK  = "r.transfer.#.ExpireTransfersRequest"
	expireTransfersResponseRK = "r.transfer.bitbon-node.ExpireTransfersResponse"

	// Blocks routing config
	blockByHashRequestRK  = "r.block.#.BlockByHashRequest"
	blockByHashResponseRK = "r.block.bitbon-node.BlockByHashResponse"

	blockByNumberRequestRK  = "r.block.#.BlockByNumberRequest"
	blockByNumberResponseRK = "r.block.bitbon-node.BlockByNumberResponse"

	rangeBlocksByNumberRequestRK  = "r.block.#.RangeBlocksByNumberRequest"
	rangeBlocksByNumberResponseRK = "r.block.bitbon-node.RangeBlocksByNumberResponse"

	blocksByTimePeriodRequestRK  = "r.block.#.BlocksByTimePeriodRequest"
	blocksByTimePeriodResponseRK = "r.block.bitbon-node.BlocksByTimePeriodResponse"

	// Transaction routing config
	blockTransactionCountByHashRequestRK  = "r.transaction.#.BlockTransactionCountByHashRequest"
	blockTransactionCountByHashResponseRK = "r.transaction.bitbon-node.BlockTransactionCountByHashResponse"

	blockTransactionCountByNumberRequestRK  = "r.transaction.#.BlockTransactionCountByNumberRequest"
	blockTransactionCountByNumberResponseRK = "r.transaction.bitbon-node.BlockTransactionCountByNumberResponse"

	transactionByBlockHashAndIndexRequestRK  = "r.transaction.#.TransactionByBlockHashAndIndexRequest"
	transactionByBlockHashAndIndexResponseRK = "r.transaction.bitbon-node.TransactionByBlockHashAndIndexResponse"

	transactionByBlockNumberAndIndexRequestRK  = "r.transaction.#.TransactionByBlockNumberAndIndexRequest"
	transactionByBlockNumberAndIndexResponseRK = "r.transaction.bitbon-node.TransactionByBlockNumberAndIndexResponse"

	transactionRequestRK  = "r.transaction.#.TransactionRequest"
	transactionResponseRK = "r.transaction.bitbon-node.TransactionResponse"

	transactionsByTimePeriodRequestRK  = "r.transaction.#.TransactionsByTimePeriodRequest"
	transactionsByTimePeriodResponseRK = "r.transaction.bitbon-node.TransactionsByTimePeriodResponse"

	checkTransactionsByTimePeriodRequestRK  = "r.transaction.#.CheckTransactionsByTimePeriodRequest"
	checkTransactionsByTimePeriodResponseRK = "r.transaction.bitbon-node.CheckTransactionsByTimePeriodResponse"

	// Mining agent routing config
	proposeDistributionRequestRK  = "r.mining-agent.#.ProposeDistributionRequest"
	proposeDistributionResponseRK = "r.mining-agent.bitbon-node.ProposeDistributionResponse"

	getCurrentDistributionRequestRK  = "r.mining-agent.#.GetCurrentDistributionRequest"
	getCurrentDistributionResponseRK = "r.mining-agent.bitbon-node.GetCurrentDistributionResponse"

	getMinerNodesRequestRK  = "r.mining-agent.#.GetMinerNodesRequest"
	getMinerNodesResponseRK = "r.mining-agent.bitbon-node.GetMinerNodesResponse"

	// Fee routing config
	feeSettingsRequestRK  = "r.fee.#.FeeSettingsRequest"
	feeSettingsResponseRK = "r.fee.bitbon-node.FeeSettingsResponse"

	// keys for configs
	prepareAssetboxesKey                 = "prepareAssetboxes"
	deleteAssetboxKey                    = "deleteAssetbox"
	setPublicAssetboxInfoKey             = "setPublicAssetboxInfo"
	assetboxBalancesKey                  = "assetboxBalances"
	quickTransferKey                     = "quickTransfer"
	frameTransferKey                     = "frameTransfer"
	serviceFeeTransferKey                = "serviceFeeTransfer"
	createWPCSafeTransferKey             = "createWPCSafeTransfer"
	approveWPCSafeTransferKey            = "approveWPCSafeTransfer"
	cancelWPCSafeTransferKey             = "cancelWPCSafeTransfer"
	createSafeTransferKey                = "createSafeTransfer"
	approveSafeTransferKey               = "approveSafeTransfer"
	cancelSafeTransferKey                = "cancelSafeTransfer"
	directTransferKey                    = "directTransfer"
	monetizeEmissionKey                  = "monetizeEmission"
	monetizeCertificateKey               = "monetizeCertificate"
	fullBalanceQuickTransferKey          = "fullBalanceQuickTransfer"
	createFullBalanceWPCSafeTransferKey  = "createFullBalanceWPCSafeTransfer"
	approveFullBalanceWPCSafeTransferKey = "approveFullBalanceWPCSafeTransfer"
	cancelFullBalanceWPCSafeTransferKey  = "cancelFullBalanceWPCSafeTransfer"
	createFullBalanceSafeTransferKey     = "createFullBalanceSafeTransfer"
	approveFullBalanceSafeTransferKey    = "approveFullBalanceSafeTransfer"
	cancelFullBalanceSafeTransferKey     = "cancelFullBalanceSafeTransfer"
	// Async transfers
	quickTransferAsyncKey                     = "quickTransferAsync"
	frameTransferAsyncKey                     = "frameTransferAsync"
	serviceFeeTransferAsyncKey                = "serviceFeeTransferAsync"
	createWPCTransferAsyncKey                 = "createWPCTransferAsync"
	approveWPCSafeTransferAsyncKey            = "approveWPCSafeTransferAsync"
	cancelWPCSafeTransferAsyncKey             = "cancelWPCSafeTransferAsync"
	createSafeTransferAsyncKey                = "createSafeTransferAsync"
	approveSafeTransferAsyncKey               = "approveSafeTransferAsync"
	cancelSafeTransferAsyncKey                = "cancelSafeTransferAsync"
	directTransferAsyncKey                    = "directTransferAsync"
	fullBalanceQuickAllTransferAsyncKey       = "fullBalanceQuickTransferAsync"
	createFullBalanceWPCSafeTransferAsyncKey  = "createFullBalanceWPCTransferAsync"
	approveFullBalanceWPCSafeTransferAsyncKey = "approveFullBalanceWPCSafeTransferAsync"
	cancelFullBalanceWPCSafeTransferAsyncKey  = "cancelFullBalanceWPCSafeTransferAsync"
	createFullBalanceSafeTransferAsyncKey     = "createFullBalanceSafeTransferAsync"
	approveFullBalanceSafeTransferAsyncKey    = "approveFullBalanceSafeTransferAsync"
	cancelFullBalanceSafeTransferAsyncKey     = "cancelFullBalanceSafeTransferAsync"

	// Block key config
	blockByHashKey         = "blockByHash"
	blockByNumberKey       = "blockByNumber"
	blocksByTimePeriodKey  = "blocksByTimePeriod"
	rangeBlocksByNumberKey = "rangeBlocksByNumber"

	// Transaction key config
	transactionKey                      = "transaction"
	transactionsByTimePeriodKey         = "transactionsByTimePeriod"
	blockTransactionCountByHashKey      = "blockTransactionCountByHash"
	blockTransactionCountByNumberKey    = "blockTransactionCountByNumber"
	transactionByBlockHashAndIndexKey   = "transactionByBlockHashAndIndex"
	transactionByBlockNumberAndIndexKey = "transactionByBlockNumberAndIndex"
	checkTransactionsByTimePeriodKey    = "checkTransactionsByTimePeriod"
	expireTransfersKey                  = "expireTransfers"

	// Mining Agent key config
	proposeDistributionKey    = "proposeDistribution"
	getCurrentDistributionKey = "getCurrentDistribution"
	getMinerNodesKey          = "getMinerNodes"

	// Fee key config
	feeValueSettingsKey = "feeValueSettings"
)

func (c *Client) getAllModes() map[string]BitbonClientMode {
	return map[string]BitbonClientMode{
		"Undefined": BitbonClientUndefinedMode,
		"All":       BitbonClientAllMode,
		"Mutable":   BitbonClientMutableMode,
		"Readable":  BitbonClientReadableMode,
	}
}

func IsAllowed() bool {
	return true
}

// nolint:funlen
func (c *Client) getMutableProcessConfig() map[string]workerConfig {
	return map[string]workerConfig{
		prepareAssetboxesKey: {
			exchange:       c.cfg.AmqpAssetboxExchange,
			queueName:      "q.bitbon-node.handlePrepareAssetbox",
			reqRoutingKey:  prepareAssetboxRequestRK,
			respRoutingKey: prepareAssetboxResponseRK,
			handle:         c.handlePrepareAssetbox,
			threads:        1,
		},
		deleteAssetboxKey: {
			exchange:       c.cfg.AmqpAssetboxExchange,
			queueName:      "q.bitbon-node.handleDeleteAssetbox",
			reqRoutingKey:  deleteAssetboxRequestRK,
			respRoutingKey: deleteAssetboxResponseRK,
			handle:         c.handleDeleteAssetbox,
			threads:        1,
		},
		setPublicAssetboxInfoKey: {
			exchange:       c.cfg.AmqpAssetboxExchange,
			queueName:      "q.bitbon-node.handleSetPublicAssetboxInfo",
			reqRoutingKey:  setPublicAssetboxInfoRequestRK,
			respRoutingKey: setPublicAssetboxInfoResponseRK,
			handle:         c.handleSetPublicAssetboxInfo,
			threads:        1,
		},
		quickTransferKey: {
			exchange:       c.cfg.AmqpTransferExchange,
			queueName:      "q.bitbon-node.handleQuickTransfer",
			reqRoutingKey:  quickTransferRequestRK,
			respRoutingKey: quickTransferResponseRK,
			handle:         c.handleQuickTransfer,
			threads:        runtime.NumCPU(),
		},
		frameTransferKey: {
			exchange:       c.cfg.AmqpTransferExchange,
			queueName:      "q.bitbon-node.handleFrameTransfer",
			reqRoutingKey:  frameTransferRequestRK,
			respRoutingKey: frameTransferResponseRK,
			handle:         c.handleFrameTransfer,
			threads:        runtime.NumCPU(),
		},
		serviceFeeTransferKey: {
			exchange:       c.cfg.AmqpTransferExchange,
			queueName:      "q.bitbon-node.handleServiceFeeTransfer",
			reqRoutingKey:  serviceFeeTransferRequestRK,
			respRoutingKey: serviceFeeTransferResponseRK,
			handle:         c.handleServiceFeeTransfer,
			threads:        runtime.NumCPU(),
		},
		createWPCSafeTransferKey: {
			exchange:       c.cfg.AmqpTransferExchange,
			queueName:      "q.bitbon-node.handleCreateWPCSafeTransfer",
			reqRoutingKey:  createWPCSafeTransferRequestRK,
			respRoutingKey: createWPCSafeTransferResponseRK,
			handle:         c.handleCreateWPCSafeTransfer,
			threads:        runtime.NumCPU(),
		},
		approveWPCSafeTransferKey: {
			exchange:       c.cfg.AmqpTransferExchange,
			queueName:      "q.bitbon-node.handleApproveWPCSafeTransfer",
			reqRoutingKey:  approveWPCSafeTransferRequestRK,
			respRoutingKey: approveWPCSafeTransferResponseRK,
			handle:         c.handleApproveWPCSafeTransfer,
			threads:        runtime.NumCPU(),
		},
		cancelWPCSafeTransferKey: {
			exchange:       c.cfg.AmqpTransferExchange,
			queueName:      "q.bitbon-node.handleCancelWPCSafeTransfer",
			reqRoutingKey:  cancelWPCSafeTransferRequestRK,
			respRoutingKey: cancelWPCSafeTransferResponseRK,
			handle:         c.handleCancelWPCSafeTransfer,
			threads:        runtime.NumCPU(),
		},
		createSafeTransferKey: {
			exchange:       c.cfg.AmqpTransferExchange,
			queueName:      "q.bitbon-node.handleCreateSafeTransfer",
			reqRoutingKey:  createSafeTransferRequestRK,
			respRoutingKey: createSafeTransferResponseRK,
			handle:         c.handleCreateSafeTransfer,
			threads:        runtime.NumCPU(),
		},
		approveSafeTransferKey: {
			exchange:       c.cfg.AmqpTransferExchange,
			queueName:      "q.bitbon-node.handleApproveSafeTransfer",
			reqRoutingKey:  approveSafeTransferRequestRK,
			respRoutingKey: approveSafeTransferResponseRK,
			handle:         c.handleApproveSafeTransfer,
			threads:        runtime.NumCPU(),
		},
		cancelSafeTransferKey: {
			exchange:       c.cfg.AmqpTransferExchange,
			queueName:      "q.bitbon-node.handleCancelSafeTransfer",
			reqRoutingKey:  cancelSafeTransferRequestRK,
			respRoutingKey: cancelSafeTransferResponseRK,
			handle:         c.handleCancelSafeTransfer,
			threads:        runtime.NumCPU(),
		},
		directTransferKey: {
			exchange:       c.cfg.AmqpTransferExchange,
			queueName:      "q.bitbon-node.handleDirectTransfer",
			reqRoutingKey:  directTransferRequestRK,
			respRoutingKey: directTransferResponseRK,
			handle:         c.handleDirectTransfer,
			threads:        1,
		},
		fullBalanceQuickTransferKey: {
			exchange:       c.cfg.AmqpTransferExchange,
			queueName:      "q.bitbon-node.handleFullBalanceQuickTransfer",
			reqRoutingKey:  fullBalanceQuickTransferRequestRK,
			respRoutingKey: fullBalanceQuickTransferResponseRK,
			handle:         c.handleFullBalanceQuickTransfer,
			threads:        runtime.NumCPU(),
		},
		createFullBalanceWPCSafeTransferKey: {
			exchange:       c.cfg.AmqpTransferExchange,
			queueName:      "q.bitbon-node.handleCreateFullBalanceWPCSafeTransfer",
			reqRoutingKey:  createFullBalanceWPCSafeTransferRequestRK,
			respRoutingKey: createFullBalanceWPCSafeTransferResponseRK,
			handle:         c.handleCreateFullBalanceWPCSafeTransfer,
			threads:        runtime.NumCPU(),
		},
		approveFullBalanceWPCSafeTransferKey: {
			exchange:       c.cfg.AmqpTransferExchange,
			queueName:      "q.bitbon-node.handleApproveFullBalanceWPCSafeTransfer",
			reqRoutingKey:  approveFullBalanceWPCSafeTransferRequestRK,
			respRoutingKey: approveFullBalanceWPCSafeTransferResponseRK,
			handle:         c.handleApproveFullBalanceWPCSafeTransfer,
			threads:        runtime.NumCPU(),
		},
		cancelFullBalanceWPCSafeTransferKey: {
			exchange:       c.cfg.AmqpTransferExchange,
			queueName:      "q.bitbon-node.handleCancelFullBalanceWPCSafeTransfer",
			reqRoutingKey:  cancelFullBalanceWPCSafeTransferRequestRK,
			respRoutingKey: cancelFullBalanceWPCSafeTransferResponseRK,
			handle:         c.handleCancelFullBalanceWPCSafeTransfer,
			threads:        runtime.NumCPU(),
		},
		createFullBalanceSafeTransferKey: {
			exchange:       c.cfg.AmqpTransferExchange,
			queueName:      "q.bitbon-node.handleCreateFullBalanceSafeTransfer",
			reqRoutingKey:  createFullBalanceSafeTransferRequestRK,
			respRoutingKey: createFullBalanceSafeTransferResponseRK,
			handle:         c.handleCreateFullBalanceSafeTransfer,
			threads:        runtime.NumCPU(),
		},
		approveFullBalanceSafeTransferKey: {
			exchange:       c.cfg.AmqpTransferExchange,
			queueName:      "q.bitbon-node.handleApproveFullBalanceSafeTransfer",
			reqRoutingKey:  approveFullBalanceSafeTransferRequestRK,
			respRoutingKey: approveFullBalanceSafeTransferResponseRK,
			handle:         c.handleApproveFullBalanceSafeTransfer,
			threads:        runtime.NumCPU(),
		},
		cancelFullBalanceSafeTransferKey: {
			exchange:       c.cfg.AmqpTransferExchange,
			queueName:      "q.bitbon-node.handleCancelFullBalanceSafeTransfer",
			reqRoutingKey:  cancelFullBalanceSafeTransferRequestRK,
			respRoutingKey: cancelFullBalanceSafeTransferResponseRK,
			handle:         c.handleCancelFullBalanceSafeTransfer,
			threads:        runtime.NumCPU(),
		},
		// Async transfers
		quickTransferAsyncKey: {
			exchange:       c.cfg.AmqpTransferExchange,
			queueName:      "q.bitbon-node.handleQuickTransferAsync",
			reqRoutingKey:  quickTransferRequestAsyncRK,
			respRoutingKey: quickTransferResponseAsyncRK,
			handle:         c.handleQuickTransferAsync,
			threads:        runtime.NumCPU(),
		},
		frameTransferAsyncKey: {
			exchange:       c.cfg.AmqpTransferExchange,
			queueName:      "q.bitbon-node.handleFrameTransferAsync",
			reqRoutingKey:  frameTransferRequestAsyncRK,
			respRoutingKey: frameTransferResponseAsyncRK,
			handle:         c.handleFrameTransferAsync,
			threads:        runtime.NumCPU(),
		},
		serviceFeeTransferAsyncKey: {
			exchange:       c.cfg.AmqpTransferExchange,
			queueName:      "q.bitbon-node.handleServiceFeeTransferAsync",
			reqRoutingKey:  serviceFeeTransferRequestAsyncRK,
			respRoutingKey: serviceFeeTransferResponseAsyncRK,
			handle:         c.handleServiceFeeTransferAsync,
			threads:        runtime.NumCPU(),
		},
		createWPCTransferAsyncKey: {
			exchange:       c.cfg.AmqpTransferExchange,
			queueName:      "q.bitbon-node.handleCreateQuickTransferAsync",
			reqRoutingKey:  createWPCSafeTransferRequestAsyncRK,
			respRoutingKey: createWPCSafeTransferResponseAsyncRK,
			handle:         c.handleCreateWPCSafeTransferAsync,
			threads:        runtime.NumCPU(),
		},
		approveWPCSafeTransferAsyncKey: {
			exchange:       c.cfg.AmqpTransferExchange,
			queueName:      "q.bitbon-node.handleApproveWPCSafeTransferAsync",
			reqRoutingKey:  approveWPCSafeTransferRequestAsyncRK,
			respRoutingKey: approveWPCSafeTransferResponseAsyncRK,
			handle:         c.handleApproveWPCSafeTransferAsync,
			threads:        runtime.NumCPU(),
		},
		cancelWPCSafeTransferAsyncKey: {
			exchange:       c.cfg.AmqpTransferExchange,
			queueName:      "q.bitbon-node.handleCancelWPCSafeTransferAsync",
			reqRoutingKey:  cancelWPCSafeTransferRequestAsyncRK,
			respRoutingKey: cancelWPCSafeTransferResponseAsyncRK,
			handle:         c.handleCancelWPCSafeTransferAsync,
			threads:        runtime.NumCPU(),
		},
		createSafeTransferAsyncKey: {
			exchange:       c.cfg.AmqpTransferExchange,
			queueName:      "q.bitbon-node.handleCreateSafeTransferAsync",
			reqRoutingKey:  createSafeTransferRequestAsyncRK,
			respRoutingKey: createSafeTransferResponseAsyncRK,
			handle:         c.handleCreateSafeTransferAsync,
			threads:        runtime.NumCPU(),
		},
		approveSafeTransferAsyncKey: {
			exchange:       c.cfg.AmqpTransferExchange,
			queueName:      "q.bitbon-node.handleApproveSafeTransferAsync",
			reqRoutingKey:  approveSafeTransferRequestAsyncRK,
			respRoutingKey: approveSafeTransferResponseAsyncRK,
			handle:         c.handleApproveSafeTransferAsync,
			threads:        runtime.NumCPU(),
		},
		cancelSafeTransferAsyncKey: {
			exchange:       c.cfg.AmqpTransferExchange,
			queueName:      "q.bitbon-node.handleCancelSafeTransferAsync",
			reqRoutingKey:  cancelSafeTransferRequestAsyncRK,
			respRoutingKey: cancelSafeTransferResponseAsyncRK,
			handle:         c.handleCancelSafeTransferAsync,
			threads:        runtime.NumCPU(),
		},
		directTransferAsyncKey: {
			exchange:       c.cfg.AmqpTransferExchange,
			queueName:      "q.bitbon-node.handleDirectTransferAsync",
			reqRoutingKey:  directTransferRequestAsyncRK,
			respRoutingKey: directTransferResponseAsyncRK,
			handle:         c.handleDirectTransferAsync,
			threads:        1,
		},
		fullBalanceQuickAllTransferAsyncKey: {
			exchange:       c.cfg.AmqpTransferExchange,
			queueName:      "q.bitbon-node.handleFullBalanceQuickTransferAsync",
			reqRoutingKey:  fullBalanceQuickTransferRequestAsyncRK,
			respRoutingKey: fullBalanceQuickTransferResponseAsyncRK,
			handle:         c.handleFullBalanceQuickTransferAsync,
			threads:        runtime.NumCPU(),
		},
		createFullBalanceWPCSafeTransferAsyncKey: {
			exchange:       c.cfg.AmqpTransferExchange,
			queueName:      "q.bitbon-node.handleCreateFullBalanceWPCSafeTransferAsync",
			reqRoutingKey:  createFullBalanceWPCSafeTransferRequestAsyncRK,
			respRoutingKey: createFullBalanceWPCSafeTransferResponseAsyncRK,
			handle:         c.handleCreateFullBalanceWPCSafeTransferAsync,
			threads:        runtime.NumCPU(),
		},
		approveFullBalanceWPCSafeTransferAsyncKey: {
			exchange:       c.cfg.AmqpTransferExchange,
			queueName:      "q.bitbon-node.handleApproveFullBalanceWPCSafeTransferAsync",
			reqRoutingKey:  approveFullBalanceWPCSafeTransferRequestAsyncRK,
			respRoutingKey: approveFullBalanceWPCSafeTransferResponseAsyncRK,
			handle:         c.handleApproveFullBalanceWPCSafeTransferAsync,
			threads:        runtime.NumCPU(),
		},
		cancelFullBalanceWPCSafeTransferAsyncKey: {
			exchange:       c.cfg.AmqpTransferExchange,
			queueName:      "q.bitbon-node.handleCancelFullBalanceWPCSafeTransferAsync",
			reqRoutingKey:  cancelFullBalanceWPCSafeTransferRequestAsyncRK,
			respRoutingKey: cancelFullBalanceWPCSafeTransferResponseAsyncRK,
			handle:         c.handleCancelFullBalanceWPCSafeTransferAsync,
			threads:        runtime.NumCPU(),
		},

		createFullBalanceSafeTransferAsyncKey: {
			exchange:       c.cfg.AmqpTransferExchange,
			queueName:      "q.bitbon-node.handleCreateFullBalanceSafeTransferAsync",
			reqRoutingKey:  createFullBalanceSafeTransferRequestAsyncRK,
			respRoutingKey: createFullBalanceSafeTransferResponseAsyncRK,
			handle:         c.handleCreateFullBalanceSafeTransferAsync,
			threads:        runtime.NumCPU(),
		},
		approveFullBalanceSafeTransferAsyncKey: {
			exchange:       c.cfg.AmqpTransferExchange,
			queueName:      "q.bitbon-node.handleApproveFullBalanceSafeTransferAsync",
			reqRoutingKey:  approveFullBalanceSafeTransferRequestAsyncRK,
			respRoutingKey: approveFullBalanceSafeTransferResponseAsyncRK,
			handle:         c.handleApproveFullBalanceSafeTransferAsync,
			threads:        runtime.NumCPU(),
		},
		cancelFullBalanceSafeTransferAsyncKey: {
			exchange:       c.cfg.AmqpTransferExchange,
			queueName:      "q.bitbon-node.handleCancelFullBalanceSafeTransferAsync",
			reqRoutingKey:  cancelFullBalanceSafeTransferRequestAsyncRK,
			respRoutingKey: cancelFullBalanceSafeTransferResponseAsyncRK,
			handle:         c.handleCancelFullBalanceSafeTransferAsync,
			threads:        runtime.NumCPU(),
		},
		expireTransfersKey: {
			exchange:       c.cfg.AmqpTransferExchange,
			queueName:      "q.bitbon-node.handleExpireTransfers",
			reqRoutingKey:  expireTransfersRequestRK,
			respRoutingKey: expireTransfersResponseRK,
			handle:         c.handleExpireTransfersAsync,
			threads:        1,
		},
		proposeDistributionKey: {
			exchange:       c.cfg.AmqpMiningExchange,
			queueName:      "q.bitbon-node.handleProposeDistribution",
			reqRoutingKey:  proposeDistributionRequestRK,
			respRoutingKey: proposeDistributionResponseRK,
			handle:         c.handleProposeDistribution,
			threads:        1,
		},
	}
}

// nolint:funlen
func (c *Client) getReadableProcessConfig() map[string]workerConfig {
	return map[string]workerConfig{
		assetboxBalancesKey: {
			exchange:       c.cfg.AmqpAssetboxExchange,
			queueName:      "q.bitbon-node.handleAssetboxBalances",
			reqRoutingKey:  assetboxBalancesRequestRK,
			respRoutingKey: assetboxBalancesResponseRK,
			handle:         c.handleAssetboxBalances,
			threads:        1,
		},

		// Blocks
		blockByHashKey: {
			exchange:       c.cfg.AmqpBlockExchange,
			queueName:      "q.bitbon-node.handleBlockByHash",
			reqRoutingKey:  blockByHashRequestRK,
			respRoutingKey: blockByHashResponseRK,
			handle:         c.handleBlockByHash,
			threads:        1,
		},
		blockByNumberKey: {
			exchange:       c.cfg.AmqpBlockExchange,
			queueName:      "q.bitbon-node.handleBlockByNumber",
			reqRoutingKey:  blockByNumberRequestRK,
			respRoutingKey: blockByNumberResponseRK,
			handle:         c.handleBlockByNumber,
			threads:        1,
		},
		rangeBlocksByNumberKey: {
			exchange:       c.cfg.AmqpBlockExchange,
			queueName:      "q.bitbon-node.handleRangeBlocksByNumber",
			reqRoutingKey:  rangeBlocksByNumberRequestRK,
			respRoutingKey: rangeBlocksByNumberResponseRK,
			handle:         c.handleRangeBlocksByNumber,
			threads:        1,
		},
		blocksByTimePeriodKey: {
			exchange:       c.cfg.AmqpBlockExchange,
			queueName:      "q.bitbon-node.handleBlocksByTimePeriod",
			reqRoutingKey:  blocksByTimePeriodRequestRK,
			respRoutingKey: blocksByTimePeriodResponseRK,
			handle:         c.handleBlocksByTimePeriod,
			threads:        1,
		},

		// Transactions
		blockTransactionCountByHashKey: {
			exchange:       c.cfg.AmqpTransactionExchange,
			queueName:      "q.bitbon-node.handleBlockTransactionCountByBlockHash",
			reqRoutingKey:  blockTransactionCountByHashRequestRK,
			respRoutingKey: blockTransactionCountByHashResponseRK,
			handle:         c.handleBlockTransactionCountByBlockHash,
			threads:        1,
		},
		blockTransactionCountByNumberKey: {
			exchange:       c.cfg.AmqpTransactionExchange,
			queueName:      "q.bitbon-node.handleBlockTransactionCountByNumber",
			reqRoutingKey:  blockTransactionCountByNumberRequestRK,
			respRoutingKey: blockTransactionCountByNumberResponseRK,
			handle:         c.handleBlockTransactionCountByNumber,
			threads:        1,
		},
		transactionByBlockHashAndIndexKey: {
			exchange:       c.cfg.AmqpTransactionExchange,
			queueName:      "q.bitbon-node.handleTransactionByBlockHashAndIndex",
			reqRoutingKey:  transactionByBlockHashAndIndexRequestRK,
			respRoutingKey: transactionByBlockHashAndIndexResponseRK,
			handle:         c.handleTransactionByBlockHashAndIndex,
			threads:        1,
		},
		transactionByBlockNumberAndIndexKey: {
			exchange:       c.cfg.AmqpTransactionExchange,
			queueName:      "q.bitbon-node.handleTransactionByBlockNumberAndIndex",
			reqRoutingKey:  transactionByBlockNumberAndIndexRequestRK,
			respRoutingKey: transactionByBlockNumberAndIndexResponseRK,
			handle:         c.handleTransactionByBlockNumberAndIndex,
			threads:        1,
		},
		transactionKey: {
			exchange:       c.cfg.AmqpTransactionExchange,
			queueName:      "q.bitbon-node.handleTransaction",
			reqRoutingKey:  transactionRequestRK,
			respRoutingKey: transactionResponseRK,
			handle:         c.handleTransaction,
			threads:        1,
		},
		transactionsByTimePeriodKey: {
			exchange:       c.cfg.AmqpTransactionExchange,
			queueName:      "q.bitbon-node.handleTransactionsByTimePeriod",
			reqRoutingKey:  transactionsByTimePeriodRequestRK,
			respRoutingKey: transactionsByTimePeriodResponseRK,
			handle:         c.handleTransactionsByPeriod,
			threads:        1,
		},
		checkTransactionsByTimePeriodKey: {
			exchange:       c.cfg.AmqpTransactionExchange,
			queueName:      "q.bitbon-node.handleCheckTransactionsByTimePeriod",
			reqRoutingKey:  checkTransactionsByTimePeriodRequestRK,
			respRoutingKey: checkTransactionsByTimePeriodResponseRK,
			handle:         c.handleCheckTransactionsByPeriod,
			threads:        1,
		},

		getCurrentDistributionKey: {
			exchange:       c.cfg.AmqpMiningExchange,
			queueName:      "q.bitbon-node.handleGetCurrentDistribution",
			reqRoutingKey:  getCurrentDistributionRequestRK,
			respRoutingKey: getCurrentDistributionResponseRK,
			handle:         c.handleGetCurrentDistribution,
			threads:        1,
		},
		getMinerNodesKey: {
			exchange:       c.cfg.AmqpMiningExchange,
			queueName:      "q.bitbon-node.handleGetMinerNodes",
			reqRoutingKey:  getMinerNodesRequestRK,
			respRoutingKey: getMinerNodesResponseRK,
			handle:         c.handleGetMinerNodes,
			threads:        1,
		},

		// Fee's
		feeValueSettingsKey: {
			exchange:       c.cfg.AmqpFeeExchange,
			queueName:      "q.bitbon-node.handleFeeValueSettings",
			reqRoutingKey:  feeSettingsRequestRK,
			respRoutingKey: feeSettingsResponseRK,
			handle:         c.handleFeeSettingsRequest,
			threads:        1,
		},
	}
}

// nolint:funlen
func (c *Client) processConfigs(mode BitbonClientMode) map[string]workerConfig {
	switch mode {
	case BitbonClientReadableMode:
		return c.getReadableProcessConfig()
	case BitbonClientMutableMode:
		return c.getMutableProcessConfig()
	case BitbonClientAllMode:
		mutable := c.getMutableProcessConfig()
		readable := c.getReadableProcessConfig()
		for key, val := range readable {
			mutable[key] = val
		}
		return mutable
	}
	log.Error("invalid bitbon client mode", "val", mode)
	return nil
}
