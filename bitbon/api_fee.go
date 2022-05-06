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
	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/dto"
	bitbonErrors "github.com/simcord-llc/bitbon-system-blockchain/bitbon/errors"
	"github.com/simcord-llc/bitbon-system-blockchain/common"
)

var (
	errInvalidOperationType = errors.New("invalid operation type")
)

func (api *API) GetFee(ctx context.Context, opType dto.OperationType) (*big.Int, error) {
	if api.bitbon.APIStopped() {
		return nil, ErrAPIStopped
	}

	api.bitbon.apiWG.Add(1)
	defer api.bitbon.apiWG.Done()

	res, err := api.bitbon.GetFeeManager().GetFee(ctx, opType)
	if err != nil {
		return nil, bitbonErrors.NewInternalError(err)
	}

	return res, nil
}

func (api *API) GetFeeForOperationType(ctx context.Context, operationType dto.OperationType) (*big.Int, error) {
	if api.bitbon.APIStopped() {
		return nil, ErrAPIStopped
	}

	api.bitbon.apiWG.Add(1)
	defer api.bitbon.apiWG.Done()

	if operationType.IsAOperationType() {
		res, err := api.bitbon.GetFeeManager().GetFee(ctx, operationType)
		if err != nil {
			return nil, bitbonErrors.NewInternalError(err)
		}

		return res, nil
	}
	return nil, bitbonErrors.NewInternalError(errInvalidOperationType)
}

func (api *API) GetFeeDistributionAccounts(ctx context.Context, operationType *big.Int) ([]common.Address, error) {
	if api.bitbon.APIStopped() {
		return nil, ErrAPIStopped
	}

	api.bitbon.apiWG.Add(1)
	defer api.bitbon.apiWG.Done()

	res, err := api.bitbon.GetFeeManager().GetFeeDistributionAccounts(ctx, operationType)
	if err != nil {
		return nil, bitbonErrors.NewInternalError(err)
	}

	return res, nil
}

func (api *API) GetFeeDistributionAmounts(ctx context.Context, operationType *big.Int) ([]*big.Int, error) {
	if api.bitbon.APIStopped() {
		return nil, ErrAPIStopped
	}

	api.bitbon.apiWG.Add(1)
	defer api.bitbon.apiWG.Done()

	res, err := api.bitbon.GetFeeManager().GetFeeDistributionAmounts(ctx, operationType)
	if err != nil {
		return nil, bitbonErrors.NewInternalError(err)
	}

	return res, nil
}

func (api *API) GetFeeSettings(ctx context.Context) (*dto.FeeSettingsResponse, error) {
	if api.bitbon.APIStopped() {
		return nil, ErrAPIStopped
	}

	api.bitbon.apiWG.Add(1)
	defer api.bitbon.apiWG.Done()

	feeSettings, err := api.bitbon.GetFeeManager().GetFeeSettings(ctx)
	if err != nil {
		return nil, bitbonErrors.NewInternalError(err)
	}

	return feeSettings.ToDTO(), nil
}
