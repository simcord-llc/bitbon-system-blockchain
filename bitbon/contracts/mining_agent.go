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

package contracts

import (
	"context"
	"crypto/ecdsa"

	"github.com/pkg/errors"
	"github.com/simcord-llc/bitbon-system-blockchain/accounts/abi/bind"
	"github.com/simcord-llc/bitbon-system-blockchain/common"

	bitbonErrors "github.com/simcord-llc/bitbon-system-blockchain/bitbon/errors"
)

func (m *Manager) getMiningAgentStorageImpl() (*MiningAgentStorageImpl, error) {
	if !m.Ready() {
		return nil, ErrorNotReady
	}
	contractAddress, err := m.GetContractAddressMiningAgentStorage()
	if err != nil {
		return nil, err
	}
	contract, err := NewMiningAgentStorageImpl(contractAddress, m.EthApiWrapper)
	if err != nil {
		return nil, err
	}
	return contract, nil
}

// GetAllMiningAgents returns all mining agents
func (m *Manager) GetAllMiningAgents(ctx context.Context) (res []common.Address, err error) {
	m.logger.Debug("contracts manager GetAllMiningAgents called")

	defer func() {
		if err != nil {
			err = bitbonErrors.NewContractCallError(err)
		}
	}()

	contract, err := m.getMiningAgentStorageImpl()
	if err != nil {
		return nil, err
	}

	opts := &bind.CallOpts{
		Pending: false,
		From:    m.serviceAddress,
		Context: ctx,
	}

	res, err = contract.GetAll(opts)
	if err != nil {
		return nil, err
	}

	if len(res) == 0 {
		return nil, errors.Wrap(err, "empty mining agents in blockchain")
	}

	return res, nil
}

// IsMiningAgent checks if this address is a mining agent address
func (m *Manager) IsMiningAgent(ctx context.Context, address common.Address) (res bool, err error) {
	m.logger.Debug("contracts manager IsMiningAgent called")

	defer func() {
		if err != nil {
			err = bitbonErrors.NewContractCallError(err)
		}
	}()

	contract, err := m.getMiningAgentStorageImpl()
	if err != nil {
		return false, err
	}

	opts := &bind.CallOpts{
		Pending: false,
		From:    m.serviceAddress,
		Context: ctx,
	}

	return contract.IsMiningAgent(opts, address)
}

// AddMiningAgent adds a mining agent to blockchain
func (m *Manager) AddMiningAgent(ctx context.Context, address common.Address, key *ecdsa.PrivateKey) (err error) {
	m.logger.Debug("contracts manager AddMiningAgent called", "address", address.Hex())

	defer func() {
		if err != nil {
			err = bitbonErrors.NewContractCallError(err)
		}
	}()

	contract, err := m.getMiningAgentStorageImpl()
	if err != nil {
		return err
	}

	opts, err := m.prepareTransactOpts(ctx, key)
	if err != nil {
		return err
	}

	tx, err := contract.Add(opts, address)
	if err != nil {
		return err
	}

	_, err = m.waitAndCheckTransaction(ctx, tx)
	return
}

// RemoveMiningAgent removes a mining agent to blockchain
func (m *Manager) RemoveMiningAgent(ctx context.Context, address common.Address, key *ecdsa.PrivateKey) (err error) {
	m.logger.Debug("contracts manager AddMiningAgent called", "address", address.Hex())

	defer func() {
		if err != nil {
			err = bitbonErrors.NewContractCallError(err)
		}
	}()

	contract, err := m.getMiningAgentStorageImpl()
	if err != nil {
		return err
	}

	opts, err := m.prepareTransactOpts(ctx, key)
	if err != nil {
		return err
	}

	tx, err := contract.Remove(opts, address)
	if err != nil {
		return err
	}

	_, err = m.waitAndCheckTransaction(ctx, tx)
	return
}
