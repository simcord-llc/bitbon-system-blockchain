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

package dto

import (
	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/dto/external"
)

type FeeDistribution struct {
	Account     string
	SelfPercent int32
}

type FeeSettingsResponse struct {
	ValueSettings        map[int32]string
	DistributionSettings map[int32][]*FeeDistribution
}

func (o *FeeSettingsResponse) DistributionSettingsToExternalDTO() map[int32]*external.FeeDistributionSettings {
	distributionSettings := make(map[int32]*external.FeeDistributionSettings, len(o.DistributionSettings))
	for optType := range o.DistributionSettings {
		distrs := make([]*external.FeeDistribution, len(o.DistributionSettings[optType]))
		for i := range o.DistributionSettings[optType] {
			distr := o.DistributionSettings[optType][i]
			distrs[i] = &external.FeeDistribution{
				Account:     distr.Account,
				SelfPercent: distr.SelfPercent,
			}
		}

		distributionSettings[optType] = &external.FeeDistributionSettings{
			DistributionSettings: distrs,
		}
	}

	return distributionSettings
}
