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

// nolint:dupl
package bitbon

import (
	"context"
	"math/big"

	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/models"

	"github.com/pkg/errors"
	loggerContext "github.com/simcord-llc/bitbon-system-blockchain/bitbon/context"
	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/dto"
	bitbonErrors "github.com/simcord-llc/bitbon-system-blockchain/bitbon/errors"
	"github.com/simcord-llc/bitbon-system-blockchain/common"
	"github.com/simcord-llc/bitbon-system-blockchain/log"
)

// API provides the bitbon RPC service
type API struct {
	bitbon *Bitbon
	logger log.Logger
}

// NewAPI create a new RPC bitbon service.
func NewAPI(b *Bitbon) *API {
	return &API{
		bitbon: b,
		logger: loggerContext.LoggerFromContext(loggerContext.NewLoggerContext(context.Background(), log.New("src", "API"))),
	}
}

// Version returns the Whisper sub-protocol version.
func (api *API) Version(_ context.Context) string {
	return ProtocolVersionStr
}

func (api *API) SetContractStorageAddress(address common.Address) error {
	if api.bitbon.APIStopped() {
		return ErrAPIStopped
	}
	api.bitbon.apiWG.Add(1)
	defer api.bitbon.apiWG.Done()

	if err := api.bitbon.GetContractsManager().SaveContractStorageAddress(address); err != nil {
		return err
	}
	return nil
}

func (api *API) GetContractAddresses() (map[string]common.Address, error) {
	if api.bitbon.APIStopped() {
		return nil, ErrAPIStopped
	}
	api.bitbon.apiWG.Add(1)
	defer api.bitbon.apiWG.Done()

	addresses := api.bitbon.GetContractsManager().GetContractAddresses()
	ret := make(map[string]common.Address, len(addresses))
	for t, a := range addresses {
		ret[string(t)] = a
	}

	return ret, nil
}

func (api *API) GetServiceAddress() common.Address {
	return api.bitbon.serviceAddress
}

func (api *API) PrepareAssetboxes(ctx context.Context, req *dto.PrepareAssetboxesRequest) ([]*models.ShortAssetbox, error) {
	if api.bitbon.APIStopped() {
		return nil, ErrAPIStopped
	}
	api.bitbon.apiWG.Add(1)
	defer api.bitbon.apiWG.Done()

	if req == nil {
		return nil, bitbonErrors.NewInvalidRequestError(errors.New("empty request given"))
	}
	assetboxes, err := api.bitbon.GetAssetboxManager().PrepareAssetboxes(ctx, req.Count)
	if err != nil {
		return nil, err
	}

	res := make([]*models.ShortAssetbox, len(assetboxes))
	for idx := range assetboxes {
		res[idx] = &models.ShortAssetbox{
			Address:    assetboxes[idx].Address,
			Wallet:     assetboxes[idx].Wallet,
			PassPhrase: assetboxes[idx].PassPhrase,
		}
	}

	return res, nil
}

func (api *API) DeleteAssetbox(ctx context.Context, req *dto.DeleteAssetboxRequest) error {
	if api.bitbon.APIStopped() {
		return ErrAPIStopped
	}
	api.bitbon.apiWG.Add(1)
	defer api.bitbon.apiWG.Done()

	if req == nil {
		return bitbonErrors.NewInvalidRequestError(errors.New("empty request given"))
	}
	err := api.bitbon.GetAssetboxManager().DeleteAssetbox(ctx, req)
	if err != nil {
		return err
	}
	return nil
}

func (api *API) SetPublicAssetboxInfo(ctx context.Context, req *dto.SetPublicAssetboxInfoRequest) error {
	if api.bitbon.APIStopped() {
		return ErrAPIStopped
	}
	api.bitbon.apiWG.Add(1)
	defer api.bitbon.apiWG.Done()

	if req == nil {
		return bitbonErrors.NewInvalidRequestError(errors.New("empty request given"))
	}
	if err := api.bitbon.GetAssetboxManager().SetPublicAssetboxInfo(ctx, req); err != nil {
		return err
	}
	return nil
}

// GetPublicAssetboxInfo returns public assetbox info from blockchain via calling smart contract method
func (api *API) GetPublicAssetboxInfo(ctx context.Context,
	address common.Address) (*dto.PublicAssetboxInfoResponse, error) {
	if api.bitbon.APIStopped() {
		return nil, ErrAPIStopped
	}
	api.bitbon.apiWG.Add(1)
	defer api.bitbon.apiWG.Done()

	if address == (common.Address{}) {
		return nil, bitbonErrors.NewInvalidRequestError(errors.New("empty address given"))
	}
	contractAssetbox, err := api.bitbon.GetContractsManager().GetAssetboxInfo(ctx, address)
	if err != nil {
		return nil, err
	}
	return models.ContractAssetboxToDTO(contractAssetbox), nil
}

func (api *API) GetAssetboxBalance(ctx context.Context, address common.Address) (*big.Int, error) {
	if api.bitbon.APIStopped() {
		return nil, ErrAPIStopped
	}
	api.bitbon.apiWG.Add(1)
	defer api.bitbon.apiWG.Done()

	if address == (common.Address{}) {
		return nil, bitbonErrors.NewInvalidRequestError(errors.New("empty address given"))
	}
	balance, err := api.bitbon.GetContractsManager().GetAssetboxBalance(ctx, address)
	if err != nil {
		return nil, err
	}
	return balance, nil
}

func (api *API) GetAssetboxBalances(ctx context.Context,
	addresses []common.Address) (map[common.Address]*big.Int, error) {
	if api.bitbon.APIStopped() {
		return nil, ErrAPIStopped
	}
	api.bitbon.apiWG.Add(1)
	defer api.bitbon.apiWG.Done()

	if addresses == nil {
		return nil, bitbonErrors.NewInvalidRequestError(errors.New("empty addresses given"))
	}
	balances, err := api.getAssetboxBalances(ctx, addresses)
	if err != nil {
		return nil, err
	}
	return balances, nil
}

func (api *API) getAssetboxBalances(ctx context.Context,
	addresses []common.Address) (map[common.Address]*big.Int, error) {
	if addresses == nil {
		return nil, bitbonErrors.NewInvalidRequestError(errors.New("empty addresses given"))
	}
	balances, err := api.bitbon.GetContractsManager().GetAssetboxBalances(ctx, addresses)
	if err != nil {
		return nil, err
	}
	return balances, nil
}

func (api *API) GetAsseboxesMiningState(ctx context.Context, addresses []common.Address) (map[common.Address]bool, error) {
	if addresses == nil {
		return nil, bitbonErrors.NewInvalidRequestError(errors.New("empty addresses given"))
	}
	isMinings := make(map[common.Address]bool)

	for _, addr := range addresses {
		res, err := api.bitbon.GetContractsManager().GetAsseboxesMiningState(ctx, addr)
		if err != nil {
			return nil, err
		}
		isMinings[addr] = res
	}
	return isMinings, nil
}
func (api *API) GetAssetboxBalancesSum(ctx context.Context, addresses []common.Address) (*big.Int, error) {
	if api.bitbon.APIStopped() {
		return nil, ErrAPIStopped
	}
	api.bitbon.apiWG.Add(1)
	defer api.bitbon.apiWG.Done()

	if addresses == nil {
		return nil, bitbonErrors.NewInvalidRequestError(errors.New("empty addresses given"))
	}
	balance, err := api.bitbon.GetContractsManager().GetAssetboxBalancesSum(ctx, addresses)
	if err != nil {
		return nil, err
	}
	return balance, nil
}
