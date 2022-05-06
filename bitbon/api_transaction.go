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

package bitbon

import (
	"context"

	"github.com/pkg/errors"
	ethereum "github.com/simcord-llc/bitbon-system-blockchain"
	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/dto"
	bitbonErrors "github.com/simcord-llc/bitbon-system-blockchain/bitbon/errors"
	"github.com/simcord-llc/bitbon-system-blockchain/core/types"

	"github.com/simcord-llc/bitbon-system-blockchain/common"
)

func (api *API) BlockTransactionCountByBlockHash(ctx context.Context, hash common.Hash) (uint64, error) {
	if api.bitbon.APIStopped() {
		return 0, ErrAPIStopped
	}
	api.bitbon.apiWG.Add(1)
	defer api.bitbon.apiWG.Done()

	if hash == (common.Hash{}) {
		return 0, bitbonErrors.NewInvalidRequestError(errEmptyRequest)
	}

	block := api.bitbon.GetEth().BlockChain().GetBlockByHash(hash)
	if block == nil {
		return 0, bitbonErrors.NewNotFoundError(ethereum.NotFound)
	}

	return uint64(block.Transactions().Len()), nil
}

func (api *API) BlockTransactionCountByBlockNumber(ctx context.Context, number uint64) (uint64, error) {
	if api.bitbon.APIStopped() {
		return 0, ErrAPIStopped
	}
	api.bitbon.apiWG.Add(1)
	defer api.bitbon.apiWG.Done()

	if number == 0 {
		return 0, nil
	}

	block := api.bitbon.GetEth().BlockChain().GetBlockByNumber(number)
	if block == nil {
		return 0, bitbonErrors.NewNotFoundError(ethereum.NotFound)
	}

	return uint64(block.Transactions().Len()), nil
}

func (api *API) TransactionByBlockHashAndIndex(ctx context.Context,
	request *dto.TransactionByBlockHashAndIndexRequest) (*dto.Transaction, error) {
	if api.bitbon.APIStopped() {
		return nil, ErrAPIStopped
	}
	api.bitbon.apiWG.Add(1)
	defer api.bitbon.apiWG.Done()

	block := api.bitbon.GetEth().BlockChain().GetBlockByHash(request.BlockHash)
	if block == nil {
		return nil, bitbonErrors.NewNotFoundError(ethereum.NotFound)
	}

	if request.TxIndex >= uint64(block.Transactions().Len()) {
		return nil, bitbonErrors.NewInvalidRequestError(errors.New("bad request given"))
	}

	transaction := block.Transactions()[request.TxIndex]

	return api.getTransactionResponse(ctx, block, transaction)
}

func (api *API) TransactionByBlockNumberAndIndex(ctx context.Context,
	request *dto.TransactionByBlockNumberAndIndexRequest) (*dto.Transaction, error) {
	if api.bitbon.APIStopped() {
		return nil, ErrAPIStopped
	}
	api.bitbon.apiWG.Add(1)
	defer api.bitbon.apiWG.Done()

	if request.BlockNumber == 0 {
		return nil, bitbonErrors.NewNotFoundError(ethereum.NotFound)
	}

	block := api.bitbon.GetEth().BlockChain().GetBlockByNumber(request.BlockNumber)
	if block == nil {
		return nil, bitbonErrors.NewNotFoundError(ethereum.NotFound)
	}

	if request.TxIndex >= uint64(block.Transactions().Len()) {
		return nil, bitbonErrors.NewInvalidRequestError(errors.New("bad request given"))
	}

	transaction := block.Transactions()[request.TxIndex]

	return api.getTransactionResponse(ctx, block, transaction)
}

func (api *API) TransactionByHash(ctx context.Context, txHash common.Hash) (*dto.Transaction, error) {
	if api.bitbon.APIStopped() {
		return nil, ErrAPIStopped
	}
	api.bitbon.apiWG.Add(1)
	defer api.bitbon.apiWG.Done()

	if txHash == (common.Hash{}) {
		return nil, bitbonErrors.NewInvalidRequestError(errEmptyRequest)
	}

	transaction, blockHash, blockNumber := api.bitbon.GetContractsManager().GetEthAPIWrapper().Transaction(ctx, txHash)
	if transaction == nil {
		return nil, bitbonErrors.NewNotFoundError(errors.New("tx not found"))
	}

	block := api.bitbon.GetEth().BlockChain().GetBlock(blockHash, blockNumber)
	if block == nil {
		return nil, bitbonErrors.NewNotFoundError(errors.New("block not found"))
	}
	return api.getTransactionResponse(ctx, block, transaction)
}

const MaxTxPeriod = 300

func (api *API) TransactionsByTimePeriod(ctx context.Context,
	request *dto.TransactionTimePeriodRequest) ([]*dto.Transaction, error) {
	if api.bitbon.APIStopped() {
		return nil, ErrAPIStopped
	}
	api.bitbon.apiWG.Add(1)
	defer api.bitbon.apiWG.Done()

	if err := validateTransactionsByTimePeriodRequest(request, MaxTxPeriod); err != nil {
		return nil, err
	}

	var transactions []*dto.Transaction
	for i := request.From; i <= request.To; i++ {
		block := api.bitbon.GetEth().BlockChain().GetBlockByTime(i)
		if block == nil {
			api.logger.Error("Getting block by time - block not found", "time", i)
			continue
		}

		for j, transaction := range block.Transactions() {
			clientTransaction, err := api.getTransactionResponse(ctx, block, transaction)
			if err != nil {
				api.logger.Error(" api.getTransactionResponse", "txIndex", j, "time", i, "error", err)
				continue
			}
			transactions = append(transactions, clientTransaction)
		}
	}

	if len(transactions) == 0 {
		return nil, bitbonErrors.NewNotFoundError(ethereum.NotFound)
	}

	return transactions, nil
}

func validateTransactionsByTimePeriodRequest(request *dto.TransactionTimePeriodRequest, maxPeriod uint64) error {
	if request == nil || request.From == 0 || request.To == 0 {
		return bitbonErrors.NewInvalidRequestError(errEmptyRequest)
	}
	if request.From >= request.To || (request.To-request.From) > maxPeriod {
		return bitbonErrors.NewInvalidRequestError(errors.New("bad range given"))
	}

	return nil
}

func (api *API) getTransactionResponse(ctx context.Context,
	block *types.Block, tx *types.Transaction) (*dto.Transaction, error) {
	if tx == nil {
		return nil, bitbonErrors.NewNotFoundError(ethereum.NotFound)
	}

	bbTx := api.bitbon.parser.Parse(ctx, tx)

	return dto.ToTransaction(tx, bbTx, &dto.BlockInfo{Hash: block.Hash(),
		BlockNumber: block.NumberU64(), BlockTime: block.Time()}), nil
}

const MaxCheckTxPeriod = 300

func (api *API) CheckTransactionsByTimePeriod(_ context.Context,
	request *dto.TransactionTimePeriodRequest) (hash common.Hash, err error) {
	if api.bitbon.APIStopped() {
		return common.Hash{}, ErrAPIStopped
	}
	api.bitbon.apiWG.Add(1)
	defer api.bitbon.apiWG.Done()

	if err := validateTransactionsByTimePeriodRequest(request, MaxCheckTxPeriod); err != nil {
		return common.Hash{}, err
	}

	var txHashes []common.Hash
	for i := request.From; i <= request.To; i++ {
		block := api.bitbon.GetEth().BlockChain().GetBlockByTime(i)
		if block == nil {
			api.logger.Warn("Getting block by time - block not found", "time", i)
			continue
		}

		for _, transaction := range block.Transactions() {
			txHashes = append(txHashes, transaction.Hash())
		}
	}

	return common.Keccak256Hashes(txHashes...), nil
}
