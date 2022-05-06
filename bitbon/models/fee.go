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

package models

import (
	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/dto"
	"github.com/simcord-llc/bitbon-system-blockchain/common"
	"math/big"
)

type FeeDistribution struct {
	Account     common.Address
	SelfPercent *big.Int
}

type FeeSettingsResponseObj struct {
	FeeValues        map[uint64]*big.Int
	FeeDistributions map[uint64][]*FeeDistribution
}

func (o *FeeSettingsResponseObj) ToDTO() *dto.FeeSettingsResponse {
	valueSettings := make(map[int32]string, len(o.FeeValues))
	for optType := range o.FeeValues {
		valueSettings[int32(optType)] = o.FeeValues[optType].String()
	}

	distributionSettings := make(map[int32][]*dto.FeeDistribution, len(o.FeeDistributions))
	for optType := range o.FeeDistributions {
		feeDistribution := make([]*dto.FeeDistribution, len(o.FeeDistributions[optType]))
		for i := range o.FeeDistributions[optType] {
			feeDistribution[i] = o.FeeDistributions[optType][i].ToDTO()
		}

		distributionSettings[int32(optType)] = feeDistribution
	}

	return &dto.FeeSettingsResponse{
		ValueSettings:        valueSettings,
		DistributionSettings: distributionSettings,
	}
}

func (o *FeeDistribution) ToDTO() *dto.FeeDistribution {
	return &dto.FeeDistribution{
		Account:     o.Account.Hex(),
		SelfPercent: int32(o.SelfPercent.Int64()),
	}
}
