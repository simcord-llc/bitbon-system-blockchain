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
	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/models"

	"math/big"

	"github.com/pkg/errors"
	loggerContext "github.com/simcord-llc/bitbon-system-blockchain/bitbon/context"
	"github.com/simcord-llc/bitbon-system-blockchain/common"
)

var (
	ErrDistributionAccountsAndAmountsLengthMismatch = errors.New("fee distribution accounts and amounts has different size")
)

type Manager struct {
	bitbon *Bitbon
}

func NewFeeManager(b *Bitbon) *Manager {
	return &Manager{bitbon: b}
}

func (f *Manager) GetFee(ctx context.Context, opType dto.OperationType) (*big.Int, error) {
	logger := loggerContext.LoggerFromContext(ctx)
	logger.Warn("FeeManager GetFee called", "opType", opType)

	fee, err := f.bitbon.GetContractsManager().GetFee(ctx, opType)
	if err != nil {
		return fee, errors.Wrap(err, "error setting fee in smart contract")
	}

	return fee, nil
}

func (f *Manager) GetFeeDistributionAccounts(ctx context.Context, operationType *big.Int) ([]common.Address, error) {
	logger := loggerContext.LoggerFromContext(ctx)
	logger.Warn("FeeManager GetFeeDistributionAccounts called", "operationType", operationType)

	fee, err := f.bitbon.GetContractsManager().GetFeeDistributionAccounts(ctx, operationType)
	if err != nil {
		return fee, errors.Wrap(err, "error getting fee distribution accounts from smart contract")
	}

	return fee, nil
}

func (f *Manager) GetFeeDistributionAmounts(ctx context.Context, operationType *big.Int) ([]*big.Int, error) {
	logger := loggerContext.LoggerFromContext(ctx)
	logger.Warn("FeeManager GetFeeDistributionAmounts called", "operationType", operationType)

	fee, err := f.bitbon.GetContractsManager().GetFeeDistributionAmounts(ctx, operationType)
	if err != nil {
		return fee, errors.Wrap(err, "error getting fee distribution amounts from smart contract")
	}

	return fee, nil
}

func (f *Manager) GetFeeSettings(ctx context.Context) (*models.FeeSettingsResponseObj, error) {
	logger := loggerContext.LoggerFromContext(ctx)
	logger.Warn("FeeManager GetFeeValueSettings called")

	optTypes, feeValues, err := f.bitbon.GetContractsManager().GetFeeValueSettings(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "error getting fee value settings from smart contract")
	}

	res := &models.FeeSettingsResponseObj{
		FeeValues:        make(map[uint64]*big.Int, len(optTypes)),
		FeeDistributions: make(map[uint64][]*models.FeeDistribution, len(optTypes)),
	}
	for i := range optTypes {
		optType := optTypes[i].Uint64()

		distributionSettings, err := f.GetFeeDistributionSettings(ctx, optTypes[i])
		if err != nil {
			return nil, errors.Wrap(err, "error getting distribution settings")
		}

		res.FeeValues[optType] = feeValues[i]
		res.FeeDistributions[optType] = distributionSettings
	}

	return res, nil
}

func (f *Manager) GetFeeDistributionSettings(ctx context.Context, optType *big.Int) ([]*models.FeeDistribution, error) {
	logger := loggerContext.LoggerFromContext(ctx)
	logger.Warn("FeeManager GetFeeDistributionSettings called")

	distrAccs, err := f.bitbon.GetContractsManager().GetFeeDistributionAccounts(ctx, optType)
	if err != nil {
		return nil, errors.Wrap(err, "error getting fee distribution accounts from smart contract")
	}
	distrAmounts, err := f.bitbon.GetContractsManager().GetFeeDistributionAmounts(ctx, optType)
	if err != nil {
		return nil, errors.Wrap(err, "error getting fee distribution amounts from smart contract")
	}

	if len(distrAccs) != len(distrAmounts) {
		return nil, ErrDistributionAccountsAndAmountsLengthMismatch
	}

	res := make([]*models.FeeDistribution, len(distrAccs))
	for i := range distrAccs {
		res[i] = &models.FeeDistribution{
			Account:     distrAccs[i],
			SelfPercent: distrAmounts[i],
		}
	}

	return res, nil
}
