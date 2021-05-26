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

// nolint:nakedret
package contracts

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"

	"github.com/pkg/errors"
	"github.com/simcord-llc/bitbon-system-blockchain/accounts/abi/bind"
	bitbonErrors "github.com/simcord-llc/bitbon-system-blockchain/bitbon/errors"
)

func (m *Manager) getDistributionStorageImpl() (*DistributionStorageImpl, error) {
	if !m.Ready() {
		return nil, ErrorNotReady
	}
	contractAddress, err := m.GetContractAddressDistributionStorage()
	if err != nil {
		return nil, err
	}
	contract, err := NewDistributionStorageImpl(contractAddress, m.EthApiWrapper)
	if err != nil {
		return nil, err
	}
	return contract, nil
}

// ProposeDistribution adds a miner distribution proposal to blockchain
//  distribution map[string]uint - stringified address value to its value (cost)
func (m *Manager) ProposeDistribution(ctx context.Context, distribution map[string]uint64, key *ecdsa.PrivateKey) (err error) {
	m.logger.Debug("contracts manager ProposeDistribution called", "distribution", distribution)

	if distribution == nil {
		return errors.New("empty distribution given")
	}

	defer func() {
		if err != nil {
			err = bitbonErrors.NewContractCallError(err)
		}
	}()

	contract, err := m.getDistributionStorageImpl()
	if err != nil {
		return err
	}

	// set json to blockchain
	data, err := json.Marshal(distribution)
	if err != nil {
		return errors.Wrap(err, "error marshalling distribution to json")
	}

	opts, err := m.prepareTransactOpts(ctx, key)
	if err != nil {
		return err
	}

	tx, err := contract.Add(opts, data)
	if err != nil {
		return err
	}

	_, err = m.waitAndCheckTransaction(ctx, tx)
	return
}

// GetCurrentDistribution returns the current miner distribution from blockchain
//  returns distribution map[string]uint - stringified address value to its value (cost)
func (m *Manager) GetCurrentDistribution(ctx context.Context) (distribution map[string]uint64, err error) {
	m.logger.Debug("contracts manager GetCurrentDistribution called")

	defer func() {
		if err != nil {
			err = bitbonErrors.NewContractCallError(err)
		}
	}()

	contract, err := m.getDistributionStorageImpl()
	if err != nil {
		return nil, err
	}

	opts := &bind.CallOpts{
		Pending: false,
		From:    m.serviceAddress,
		Context: ctx,
	}

	res, err := contract.GetCurrent(opts)
	if err != nil {
		return nil, err
	}

	if len(res) == 0 {
		return nil, errors.Wrap(err, "empty distribution in blockchain")
	}

	err = json.Unmarshal(res, &distribution)
	if err != nil {
		return nil, errors.Wrap(err, "error unmarshalling distribution from json")
	}

	return
}
