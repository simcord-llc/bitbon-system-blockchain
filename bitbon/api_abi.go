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
	"sort"

	"github.com/pkg/errors"
	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/contracts/contract_snapshot"
	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/dto"
	"github.com/simcord-llc/bitbon-system-blockchain/common"
)

func (api *API) SaveAbiInfo(abiInfo *dto.AbiInfo) (err error) {
	if api.bitbon.APIStopped() {
		return ErrAPIStopped
	}
	api.bitbon.apiWG.Add(1)
	defer api.bitbon.apiWG.Done()

	return api.bitbon.GetStorage().SaveAbiInfo(nil, abiInfo)
}

func (api *API) EditAbiInfo(abiInfo *dto.AbiInfo) (err error) {
	if api.bitbon.APIStopped() {
		return ErrAPIStopped
	}
	api.bitbon.apiWG.Add(1)
	defer api.bitbon.apiWG.Done()

	return api.bitbon.GetStorage().EditAbiInfo(nil, abiInfo)
}

func (api *API) DeleteAbiInfoByAddress(address common.Address) (err error) {
	if api.bitbon.APIStopped() {
		return ErrAPIStopped
	}
	api.bitbon.apiWG.Add(1)
	defer api.bitbon.apiWG.Done()

	return api.bitbon.GetStorage().DeleteAbiInfoByAddress(nil, address)
}

func (api *API) GetAbiJSON(address common.Address) (abiJSON string, err error) {
	if api.bitbon.APIStopped() {
		return "", ErrAPIStopped
	}
	api.bitbon.apiWG.Add(1)
	defer api.bitbon.apiWG.Done()

	return api.bitbon.GetStorage().GetAbiJSON(nil, address)
}

func (api *API) GetAbiVersion(address common.Address) (version dto.ContractVersion, err error) {
	if api.bitbon.APIStopped() {
		return "", ErrAPIStopped
	}
	api.bitbon.apiWG.Add(1)
	defer api.bitbon.apiWG.Done()

	return api.bitbon.GetStorage().GetAbiVersion(nil, address)
}

func (api *API) GetFullAbiInfo(address common.Address) (abiInfo *dto.AbiInfo, err error) {
	if api.bitbon.APIStopped() {
		return nil, ErrAPIStopped
	}
	api.bitbon.apiWG.Add(1)
	defer api.bitbon.apiWG.Done()

	return api.bitbon.GetStorage().GetFullAbiInfo(nil, address)
}

func (api *API) GetFullAbiInfoByVersionAndType(version dto.ContractVersion,
	contractType contract_snapshot.ContractType) (abiInfo *dto.AbiInfo, err error) {
	if api.bitbon.APIStopped() {
		return nil, ErrAPIStopped
	}
	api.bitbon.apiWG.Add(1)
	defer api.bitbon.apiWG.Done()

	return api.bitbon.GetStorage().GetFullAbiInfoByVersionAndType(nil, version, contractType)
}

func (api *API) GetAbiAddressesByVersion(version dto.ContractVersion) (addresses []common.Address, err error) {
	if api.bitbon.APIStopped() {
		return nil, ErrAPIStopped
	}
	api.bitbon.apiWG.Add(1)
	defer api.bitbon.apiWG.Done()

	addresses, err = api.bitbon.GetStorage().GetAbiAddressesByVersion(nil, version)
	if err != nil {
		return nil, err
	}
	sort.Sort(common.SortAddresses(addresses))

	return addresses, nil
}

func (api *API) GetAbiInfosByVersion(version dto.ContractVersion) (abiInfos []*dto.AbiInfo, err error) {
	if api.bitbon.APIStopped() {
		return nil, ErrAPIStopped
	}
	api.bitbon.apiWG.Add(1)
	defer api.bitbon.apiWG.Done()

	return api.bitbon.GetStorage().GetAbiInfosByVersion(nil, version)
}

func (api *API) FillCurrentAbiInfo(_ context.Context,
	version dto.ContractVersion) (addresses []common.Address, err error) {
	if api.bitbon.APIStopped() {
		return nil, ErrAPIStopped
	}

	api.bitbon.apiWG.Add(1)
	defer api.bitbon.apiWG.Done()

	abiInfos, err := api.bitbon.GetContractsManager().GetCurrentContractAbiInfo(version)
	if err != nil {
		api.logger.Error("Error getting abi infos", "error", err)
		return nil, errors.Wrap(err, "error getting abi infos")
	}

	addresses = make([]common.Address, len(abiInfos))
	for i, abiInfo := range abiInfos {
		err = api.bitbon.storage.SaveAbiInfo(nil, abiInfo)

		if err != nil {
			api.logger.Error("Error saving abi info to storage", "error", err)
			return nil, errors.Wrap(err, "error saving abi info to storage")
		}
		addresses[i] = abiInfo.Address
	}

	sort.Sort(common.SortAddresses(addresses))

	return addresses, nil
}
