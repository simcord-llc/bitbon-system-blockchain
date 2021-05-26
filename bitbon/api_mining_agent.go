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

	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/dto"

	"github.com/pkg/errors"
	bitbonErrors "github.com/simcord-llc/bitbon-system-blockchain/bitbon/errors"
	"github.com/simcord-llc/bitbon-system-blockchain/common"
)

func (api *API) GetCurrentDistribution(ctx context.Context) (map[string]uint64, error) {
	if api.bitbon.APIStopped() {
		return nil, ErrAPIStopped
	}
	api.bitbon.apiWG.Add(1)
	defer api.bitbon.apiWG.Done()

	distribution, err := api.bitbon.miningAgentManager.GetCurrentDistribution(ctx)
	if err != nil {
		return nil, bitbonErrors.NewInternalError(err)
	}

	return distribution, nil
}

func (api *API) ProposeDistribution(ctx context.Context, req dto.ProposeDistributionRequest) error {
	if api.bitbon.APIStopped() {
		return ErrAPIStopped
	}
	api.bitbon.apiWG.Add(1)
	defer api.bitbon.apiWG.Done()

	if req.Address == (common.Address{}) {
		return bitbonErrors.NewInvalidRequestError(errors.New("empty address given"))
	}

	if len(req.Distribution) == 0 {
		return bitbonErrors.NewInvalidRequestError(errors.New("empty distribution given"))
	}

	// check if request.Address is authorized mining agent
	isMiningAgent, err := api.bitbon.miningAgentManager.IsMiningAgent(ctx, req.Address)
	if err != nil {
		return bitbonErrors.NewInternalError(err)
	}
	if !isMiningAgent {
		return bitbonErrors.NewInvalidRequestError(errors.New("given address is not mining agent"))
	}

	err = api.bitbon.miningAgentManager.ProposeDistribution(ctx, req)
	if err != nil {
		return bitbonErrors.NewInternalError(err)
	}

	return nil
}

func (api *API) GetAllMiningAgents(ctx context.Context) ([]common.Address, error) {
	if api.bitbon.APIStopped() {
		return nil, ErrAPIStopped
	}
	api.bitbon.apiWG.Add(1)
	defer api.bitbon.apiWG.Done()

	res, err := api.bitbon.miningAgentManager.GetAllMiningAgents(ctx)
	if err != nil {
		return nil, bitbonErrors.NewInternalError(err)
	}

	return res, nil
}

func (api *API) IsMiningAgent(ctx context.Context, address common.Address) (bool, error) {
	if api.bitbon.APIStopped() {
		return false, ErrAPIStopped
	}
	api.bitbon.apiWG.Add(1)
	defer api.bitbon.apiWG.Done()

	res, err := api.bitbon.miningAgentManager.IsMiningAgent(ctx, address)
	if err != nil {
		return false, bitbonErrors.NewInternalError(err)
	}

	return res, nil
}

func (api *API) AddMiningAgent(ctx context.Context, address common.Address) error {
	if api.bitbon.APIStopped() {
		return ErrAPIStopped
	}
	api.bitbon.apiWG.Add(1)
	defer api.bitbon.apiWG.Done()

	err := api.bitbon.miningAgentManager.AddMiningAgent(ctx, address)
	if err != nil {
		return bitbonErrors.NewInternalError(err)
	}

	return nil
}

func (api *API) RemoveMiningAgent(ctx context.Context, address common.Address) error {
	if api.bitbon.APIStopped() {
		return ErrAPIStopped
	}
	api.bitbon.apiWG.Add(1)
	defer api.bitbon.apiWG.Done()

	err := api.bitbon.miningAgentManager.RemoveMiningAgent(ctx, address)
	if err != nil {
		return bitbonErrors.NewInternalError(err)
	}

	return nil
}
