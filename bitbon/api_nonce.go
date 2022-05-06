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
	"math/big"

	"github.com/pkg/errors"
	loggerContext "github.com/simcord-llc/bitbon-system-blockchain/bitbon/context"
	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/dto"
	"github.com/simcord-llc/bitbon-system-blockchain/common"
	"github.com/simcord-llc/bitbon-system-blockchain/core/types"
	"github.com/simcord-llc/bitbon-system-blockchain/log"
	"github.com/simcord-llc/bitbon-system-blockchain/params"
)

// API provides the bitbon RPC service
type APINonce struct {
	bitbon *Bitbon
	logger log.Logger
}

// NewAPI create a new RPC bitbon service.
func NewAPINonce(b *Bitbon) *APINonce {
	return &APINonce{
		bitbon: b,
		logger: loggerContext.LoggerFromContext(loggerContext.NewLoggerContext(context.Background(),
			log.New("src", "APINonce"))),
	}
}

var (
	errNoncerDisabled           = errors.New("noncer is disabled")
	errNonceMismatchValue       = errors.New("fromNonce can`t be grater(>=) than toNonce")
	errNoncerGreater            = errors.New("toNonce can`t be grater(>) than (nonce + 1) in redis")
	errNoncerLess               = errors.New("fromNonce can`t be less(<) than nonce in blockchain")
	errNoncerLessThanRedisValue = errors.New("nonce can`t be less(<=) than nonce + 1 in redis")
)

// GetNonces returns data according to nonces
// Fields in *dto.Nonces:
// dto.Nonces.Redis - nonce in noncer redis (it should be (nonce in blockchain - 1))
// dto.Nonces.Blockchain - nonce in blockchain
// dto.Nonces.TxPoolQueued - min nonce in tx pool queued transactions
// dto.Nonces.TxPoolPending - min nonce in tx pool pending transactions
// dto.Nonces.Broken - flag that indicates if nonce is broken (there are transactions in tx pool queued transactions)
func (api *APINonce) GetNonces(ctx context.Context, address common.Address) *dto.Nonces {
	if api.bitbon.APIStopped() {
		return nil
	}
	api.bitbon.apiWG.Add(1)
	defer api.bitbon.apiWG.Done()

	var (
		res = &dto.Nonces{}
		err error
	)

	if api.bitbon.contractsManager.GetNoncer() == nil {
		res.Redis = "Noncer is disabled"
	} else {
		res.Redis, err = api.bitbon.contractsManager.GetNoncer().GetNonce(address)
		if err != nil {
			api.logger.Error("bitbon_getNonces was unable to get nonce from redis", "error", err)
			res.Redis = err.Error()
		}
	}

	res.Blockchain, err = api.bitbon.contractsManager.GetEthAPIWrapper().GetNonce(address)
	if err != nil {
		api.logger.Error("bitbon_getNonces was unable to get nonce from blockchain", "error", err)
		res.Blockchain = err.Error()
	}

	handleTxMap := func(txMap map[common.Address]types.Transactions) (interface{}, bool) {
		txs := txMap[address]
		if len(txs) == 0 {
			return "no transactions from given address", false
		}
		if txs[0] == nil {
			return "error getting first tx", false
		}

		minTxPoolNonce := txs[0].Nonce()
		for _, tx := range txs {
			if tx != nil && minTxPoolNonce > tx.Nonce() {
				minTxPoolNonce = tx.Nonce()
			}
		}
		return minTxPoolNonce, true
	}

	pending, queued := api.bitbon.eth.TxPool().Content()
	res.TxPoolQueued, res.Broken = handleTxMap(queued)
	res.TxPoolPending, _ = handleTxMap(pending)

	return res
}

// RestoreServiceAccountNonce restores service account nonce in blockchain.
// Method creates fake transactions (EmitEther to itself) beginning from fromNonce and till toNonce
// (fromNonce <= nonce < toNonce)
// Also it does some checks, to disable checks set force flag to true
// the main aim of checks is blockchainNonce == noncerNonce + 1
func (api *APINonce) RestoreServiceAccountNonce(ctx context.Context, fromNonce, toNonce uint64, force bool) error {
	if api.bitbon.APIStopped() {
		return ErrAPIStopped
	}
	api.bitbon.apiWG.Add(1)
	defer api.bitbon.apiWG.Done()

	if fromNonce >= toNonce {
		return errNonceMismatchValue
	}

	if api.bitbon.contractsManager.GetNoncer() == nil {
		return errNoncerDisabled
	}

	if force {
		loggerContext.LoggerFromContext(ctx).Warn("force flag provided - skip all checks")
	} else {
		redisNonce, err := api.bitbon.contractsManager.GetNoncer().GetNonce(api.bitbon.serviceAddress)
		if err != nil {
			return errors.Wrap(err, "failed to get service account nonce from redis")
		}
		if int64(toNonce) > redisNonce+1 {
			return errNoncerGreater
		}

		blockchainNonce, err := api.bitbon.contractsManager.GetEthAPIWrapper().GetNonce(api.bitbon.serviceAddress)
		if err != nil {
			return errors.Wrap(err, "failed to get service account nonce from blockchain")
		}
		if fromNonce < blockchainNonce {
			return errNoncerLess
		}
	}

	for nonce := fromNonce; nonce < toNonce; nonce++ {
		err := api.bitbon.GetContractsManager().EmitEther(ctx, api.bitbon.serviceAddress, big.NewInt(1*params.Wei), &nonce) // nolint
		if err != nil {
			return errors.Wrapf(err, "failed to send restoring tx from service assetbox "+
				"(emitting ether to itself) with nonce %d", nonce)
		}
	}
	return nil
}

// ForceNoncerAccountNonce method sets nonce value for address in redis (noncer)
// Also it does some checks, to disable checks set force flag to true
func (api *APINonce) ForceNoncerAccountNonce(ctx context.Context, address common.Address, nonce uint64, force bool) error {
	if api.bitbon.APIStopped() {
		return ErrAPIStopped
	}
	api.bitbon.apiWG.Add(1)
	defer api.bitbon.apiWG.Done()

	if api.bitbon.contractsManager.GetNoncer() == nil {
		return errNoncerDisabled
	}

	if !force {
		redisNonce, err := api.bitbon.contractsManager.GetNoncer().GetNonce(address)
		if err != nil {
			return errors.Wrap(err, "failed to get account nonce from redis")
		}
		if int64(nonce) <= redisNonce+1 {
			return errNoncerLessThanRedisValue
		}

		blockchainNonce, err := api.bitbon.contractsManager.GetEthAPIWrapper().GetNonce(api.bitbon.serviceAddress)
		if err != nil {
			return errors.Wrap(err, "failed to get service account nonce from blockchain")
		}
		if nonce < blockchainNonce {
			return errNoncerLessThanRedisValue
		}
	}

	return api.bitbon.GetContractsManager().GetNoncer().ForceNonce(api.bitbon.serviceAddress, int64(nonce))
}

// ForceNoncerNonceFromBlockChain takes the nonce for address from blockchain and sets it to redis noncer
func (api *APINonce) ForceNoncerNonceFromBlockChain(ctx context.Context, address common.Address) error {
	if api.bitbon.APIStopped() {
		return ErrAPIStopped
	}
	api.bitbon.apiWG.Add(1)
	defer api.bitbon.apiWG.Done()

	if api.bitbon.contractsManager.GetNoncer() == nil {
		return errNoncerDisabled
	}

	blockchainNonce, err := api.bitbon.contractsManager.GetEthAPIWrapper().GetNonce(address)
	if err != nil {
		return errors.Wrap(err, "failed to get service account nonce from blockchain")
	}
	loggerContext.LoggerFromContext(ctx).Warn("Blockchain nonce for address", "address", address, "nonce", blockchainNonce)

	return api.bitbon.GetContractsManager().GetNoncer().ForceNonce(address, int64(blockchainNonce))
}

func (api *APINonce) RemoveAssetboxesFromNoncer(ctx context.Context, addresses []common.Address) error {
	if api.bitbon.APIStopped() {
		return ErrAPIStopped
	}
	api.bitbon.apiWG.Add(1)
	defer api.bitbon.apiWG.Done()

	if api.bitbon.contractsManager.GetNoncer() == nil {
		return errNoncerDisabled
	}

	return api.bitbon.GetContractsManager().GetNoncer().RemoveAssetboxesFromNoncer(addresses)
}

func (api *APINonce) GetNoncerAssetboxes(ctx context.Context) ([]common.Address, error) {
	if api.bitbon.APIStopped() {
		return nil, ErrAPIStopped
	}
	api.bitbon.apiWG.Add(1)
	defer api.bitbon.apiWG.Done()

	if api.bitbon.contractsManager.GetNoncer() == nil {
		return nil, errNoncerDisabled
	}

	return api.bitbon.GetContractsManager().GetNoncer().GetNoncerAssetboxes()
}
