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

	ethereum "github.com/simcord-llc/bitbon-system-blockchain"
	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/dto"
	bitbonErrors "github.com/simcord-llc/bitbon-system-blockchain/bitbon/errors"
	"github.com/simcord-llc/bitbon-system-blockchain/common"
	"github.com/simcord-llc/bitbon-system-blockchain/core/types"

	"github.com/pkg/errors"
)

func (api *API) BlockByHash(ctx context.Context, hash common.Hash) (*dto.Block, error) {
	if api.bitbon.APIStopped() {
		return nil, ErrAPIStopped
	}
	api.bitbon.apiWG.Add(1)
	defer api.bitbon.apiWG.Done()

	return api.blockByHash(ctx, hash)
}
func (api *API) blockByHash(ctx context.Context, hash common.Hash) (*dto.Block, error) {
	if hash == (common.Hash{}) {
		return nil, bitbonErrors.NewInvalidRequestError(errors.New("empty request given"))
	}

	resp := api.bitbon.GetEth().BlockChain().GetBlockByHash(hash)
	if resp == nil {
		return nil, bitbonErrors.NewNotFoundError(ethereum.NotFound)
	}
	block := api.getBlockResponse(ctx, resp)

	return block, nil
}

func (api *API) BlockByNumber(ctx context.Context, number uint64) (*dto.Block, error) {
	if api.bitbon.APIStopped() {
		return nil, ErrAPIStopped
	}
	api.bitbon.apiWG.Add(1)
	defer api.bitbon.apiWG.Done()

	return api.blockByNumber(ctx, number)
}

func (api *API) blockByNumber(ctx context.Context, number uint64) (*dto.Block, error) {
	resp := api.bitbon.GetEth().BlockChain().GetBlockByNumber(number)
	if resp == nil {
		return nil, bitbonErrors.NewNotFoundError(ethereum.NotFound)
	}
	block := api.getBlockResponse(ctx, resp)

	return block, nil
}

func (api *API) BlockByTime(ctx context.Context, time uint64) (*dto.Block, error) {
	if api.bitbon.APIStopped() {
		return nil, ErrAPIStopped
	}
	api.bitbon.apiWG.Add(1)
	defer api.bitbon.apiWG.Done()

	return api.blockByTime(ctx, time)
}

func (api *API) blockByTime(ctx context.Context, time uint64) (*dto.Block, error) {
	resp := api.bitbon.GetEth().BlockChain().GetBlockByTime(time)
	if resp == nil {
		return nil, bitbonErrors.NewNotFoundError(ethereum.NotFound)
	}
	block := api.getBlockResponse(ctx, resp)

	return block, nil
}

const MaxBlockRange = 100

func (api *API) RangeBlockByNumber(ctx context.Context, request *dto.RangeBlocksByNumberRequest) ([]*dto.Block, error) {
	if api.bitbon.APIStopped() {
		return nil, ErrAPIStopped
	}
	api.bitbon.apiWG.Add(1)
	defer api.bitbon.apiWG.Done()

	return api.rangeBlockByNumber(ctx, request)
}

func (api *API) rangeBlockByNumber(ctx context.Context, request *dto.RangeBlocksByNumberRequest) ([]*dto.Block, error) {
	if request == nil || request.BlockNumberTo == 0 {
		return nil, bitbonErrors.NewInvalidRequestError(errors.New("empty request given"))
	}

	if request.BlockNumberFrom >= request.BlockNumberTo ||
		(request.BlockNumberTo-request.BlockNumberFrom) > MaxBlockRange {
		return nil, bitbonErrors.NewInvalidRequestError(errors.New("bad range given"))
	}

	blocks := make([]*dto.Block, 0, request.BlockNumberTo-request.BlockNumberFrom)
	for i := request.BlockNumberFrom; i <= request.BlockNumberTo; i++ {
		block, err := api.blockByNumber(ctx, i)
		if err != nil {
			api.logger.Error("Error getting block by number", "number", i, "error", err)
			continue
		}
		blocks = append(blocks, block)
	}

	if len(blocks) == 0 {
		return nil, bitbonErrors.NewNotFoundError(ethereum.NotFound)
	}

	return blocks, nil
}

const MaxBlockPeriod = 300

func (api *API) BlocksByTimePeriod(ctx context.Context, request *dto.BlocksTimePeriodRequest) ([]*dto.Block, error) {
	if api.bitbon.APIStopped() {
		return nil, ErrAPIStopped
	}
	api.bitbon.apiWG.Add(1)
	defer api.bitbon.apiWG.Done()

	return api.blocksByTimePeriod(ctx, request)
}

func (api *API) blocksByTimePeriod(ctx context.Context, request *dto.BlocksTimePeriodRequest) ([]*dto.Block, error) {
	if request == nil || request.From == 0 || request.To == 0 {
		return nil, bitbonErrors.NewInvalidRequestError(errors.New("empty request given"))
	}

	if request.From >= request.To || (request.To-request.From) > MaxBlockPeriod {
		return nil, bitbonErrors.NewInvalidRequestError(errors.New("bad range given"))
	}

	blocks := make([]*dto.Block, 0)
	for i := request.From; i <= request.To; i++ {
		block, err := api.blockByTime(ctx, i)
		if err != nil {
			api.logger.Error("Error getting block by time", "time", i, "error", err)
			continue
		}
		blocks = append(blocks, block)
	}

	if len(blocks) == 0 {
		return nil, bitbonErrors.NewNotFoundError(ethereum.NotFound)
	}

	return blocks, nil
}

func (api *API) getBlockResponse(ctx context.Context, resp *types.Block) *dto.Block {
	clientTxs := make([]*dto.Transaction, 0, resp.Transactions().Len())
	for _, tx := range resp.Transactions() {
		bbTx := api.bitbon.parser.Parse(ctx, tx)

		clientTx := dto.ToTransaction(tx, bbTx, &dto.BlockInfo{Hash: resp.Hash(),
			BlockNumber: resp.Number().Uint64(), BlockTime: resp.Time()})
		if clientTx != nil {
			clientTxs = append(clientTxs, clientTx)
		}
	}
	return dto.ToBitbonBlock(resp, clientTxs)
}
