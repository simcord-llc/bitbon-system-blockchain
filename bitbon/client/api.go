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

package client

import (
	"context"

	"github.com/pkg/errors"
	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/amqp"
)

// API provides the bitbon bitbon RPC service
type API struct {
	client *Client
}

// NewAPI create a new RPC bitbon agent service.
func NewAPI(client *Client) *API {
	return &API{client: client}
}

// Version returns  version.
func (api *API) Version(_ context.Context) string {
	return ProtocolVersionStr
}

func (api *API) GetHandlerKeys() []string {
	return []string{
		prepareAssetboxesKey,
		deleteAssetboxKey,
		setPublicAssetboxInfoKey,
		assetboxBalancesKey,
		createSafeTransferKey,
		approveSafeTransferKey,
		cancelSafeTransferKey,
		directTransferKey,
		monetizeEmissionKey,
		monetizeCertificateKey,
		createSafeTransferAsyncKey,
		approveSafeTransferAsyncKey,
		cancelSafeTransferAsyncKey,
		directTransferAsyncKey,
		blockByHashKey,
		blockByNumberKey,
		blocksByTimePeriodKey,
		rangeBlocksByNumberKey,
		transactionKey,
		transactionsByTimePeriodKey,
		blockTransactionCountByHashKey,
		blockTransactionCountByNumberKey,
		transactionByBlockHashAndIndexKey,
		transactionByBlockNumberAndIndexKey,
		checkTransactionsByTimePeriodKey,
		expireTransactionsKey,
		proposeDistributionKey,
		getCurrentDistributionKey,
		getMinerNodesKey,
	}
}

func (api *API) GetPoolSizes() (map[string]int, error) {
	result := make(map[string]int)
	for _, key := range api.GetHandlerKeys() {
		size, err := api.client.getPoolSize(key)
		if err != nil {
			return nil, err
		}
		result[key] = size
	}
	return result, nil
}

func (api *API) GetPoolSize(key string) (int, error) {
	return api.client.getPoolSize(key)
}

func (api *API) SetPoolSize(key string, n int) error {
	return api.client.setPoolSize(key, n)
}

func (api *API) GetJournal() ([]*amqp.Publishing, error) {
	j := api.client.GetJournal()
	if j == nil {
		return nil, errors.New("journal is not defined")
	}
	return j.GetAll(), nil
}

func (api *API) GetPoolErrors() (map[string]map[int]map[string]interface{}, error) {
	result := make(map[string]map[int]map[string]interface{})
	for _, key := range api.GetHandlerKeys() {
		errs, err := api.GetPoolError(key)
		if err != nil {
			return nil, err
		}
		result[key] = errs
	}
	return result, nil
}

func (api *API) GetPoolError(key string) (map[int]map[string]interface{}, error) {
	errs, err := api.client.getAmqpWorkerPoolErrors(key)
	if err != nil {
		return nil, err
	}

	result := make(map[int]map[string]interface{}, len(errs))
	for key, item := range errs {
		snapshot := item.Snapshot()
		resItem := make(map[string]interface{})
		resItem["First"] = snapshot.First
		resItem["Last"] = snapshot.Last
		resItem["LastReset"] = snapshot.LastReset
		resItem["ErrorsLen"] = snapshot.ErrorsLen
		tmp := make([]string, 0, len(item.Errors()))
		for _, e := range item.Errors() {
			tmp = append(tmp, e.Error())
		}
		resItem["Errors"] = tmp
		result[key] = resItem
	}
	return result, nil
}

func (api *API) GetClientModes() map[string]BitbonClientMode {
	return api.client.getAllModes()
}

func (api *API) GetCurrentMode() BitbonClientMode {
	return api.client.cfg.Mode
}
