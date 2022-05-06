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

	"github.com/pkg/errors"
	"github.com/simcord-llc/bitbon-system-blockchain/accounts/abi/bind"
	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/contracts/contract_snapshot"
	bitbonErrors "github.com/simcord-llc/bitbon-system-blockchain/bitbon/errors"
	"github.com/simcord-llc/bitbon-system-blockchain/common"
)

func (m *Manager) getContractStorageKit(address common.Address, ctx context.Context) (*ContractStorageImpl, *bind.CallOpts, error) {
	contract, err := NewContractStorageImpl(address, m.EthApiWrapper)
	if err != nil {
		return nil, nil, errors.Wrap(err, "error creating instance of contract storage")
	}
	// call non-payed smart contract method, signing transaction with assetbox private key
	opts := &bind.CallOpts{
		Pending: false,
		From:    m.serviceAddress,
		Context: ctx,
	}
	return contract, opts, nil
}

func (m *Manager) getAddress(ctx context.Context, address common.Address, contractType contract_snapshot.ContractType) (addr common.Address, err error) {
	defer func() {
		if err != nil {
			err = bitbonErrors.NewContractCallError(err)
		}
	}()

	contract, opts, err := m.getContractStorageKit(address, ctx)
	if err != nil {
		err = errors.New("error creating instance of contract storage")
		return
	}

	switch contractType {
	case contract_snapshot.ContractTypeAssetboxImpl:
		addr, err = contract.GetAssetbox(opts)
	case contract_snapshot.ContractTypeAssetboxStorageImpl:
		addr, err = contract.GetAssetboxStorage(opts)
	case contract_snapshot.ContractTypeBitbonImpl:
		addr, err = contract.GetBitbon(opts)
	case contract_snapshot.ContractTypeTransferImpl:
		addr, err = contract.GetTransferGate(opts)
	case contract_snapshot.ContractTypeSafeTransferStorageImpl:
		addr, err = contract.GetSafeTransferStorage(opts)
	case contract_snapshot.ContractTypeBitbonSupportImpl:
		addr, err = contract.GetBitbonSupport(opts)
	case contract_snapshot.ContractTypeBitbonStorageImpl:
		addr, err = contract.GetBitbonStorage(opts)
	case contract_snapshot.ContractTypeAccessStorageImpl:
		addr, err = contract.GetAccessStorage(opts)
	case contract_snapshot.ContractTypeReservedAliasStorageImpl:
		addr, err = contract.GetReservedAliasStorage(opts)
	case contract_snapshot.ContractTypeNewContractStorage:
		addr, err = contract.GetNewContactStorageAddress(opts)
	case contract_snapshot.ContractTypeDistributionStorageImpl:
		addr, err = contract.GetDistributionStorage(opts)
	case contract_snapshot.ContractTypeMiningAgentStorageImpl:
		addr, err = contract.GetMiningAgentStorage(opts)
	case contract_snapshot.ContractTypeFeeStorageImpl:
		addr, err = contract.GetFeeStorage(opts)
	default:
		err = errors.New("undefined contract type")
	}

	if err != nil {
		return
	}
	if addr == (common.Address{}) {
		err = errors.New("empty address returned")
		return
	}

	return
}

func (m *Manager) GetAssetboxInfoStorageContractAddress(ctx context.Context, address common.Address) (addr common.Address, err error) {
	return m.getAddress(ctx, address, contract_snapshot.ContractTypeAssetboxStorageImpl)
}

func (m *Manager) GetAssetboxContractAddress(ctx context.Context, address common.Address) (addr common.Address, err error) {
	return m.getAddress(ctx, address, contract_snapshot.ContractTypeAssetboxImpl)
}

func (m *Manager) GetBitbonContractAddress(ctx context.Context, address common.Address) (addr common.Address, err error) {
	return m.getAddress(ctx, address, contract_snapshot.ContractTypeBitbonImpl)
}

func (m *Manager) GetTransferContractAddress(ctx context.Context, address common.Address) (addr common.Address, err error) {
	return m.getAddress(ctx, address, contract_snapshot.ContractTypeTransferImpl)
}

func (m *Manager) GetSafeTransferStorageContractAddress(ctx context.Context, address common.Address) (addr common.Address, err error) {
	return m.getAddress(ctx, address, contract_snapshot.ContractTypeSafeTransferStorageImpl)
}

func (m *Manager) GetBitbonSupportContractAddress(ctx context.Context, address common.Address) (common.Address, error) {
	return m.getAddress(ctx, address, contract_snapshot.ContractTypeBitbonSupportImpl)
}

func (m *Manager) GetBitbonStorageContractAddress(ctx context.Context, address common.Address) (common.Address, error) {
	return m.getAddress(ctx, address, contract_snapshot.ContractTypeBitbonStorageImpl)
}

func (m *Manager) GetAccessStorageContractAddress(ctx context.Context, address common.Address) (common.Address, error) {
	return m.getAddress(ctx, address, contract_snapshot.ContractTypeAccessStorageImpl)
}

func (m *Manager) GetReservedAliasStorageContractAddress(ctx context.Context, address common.Address) (common.Address, error) {
	return m.getAddress(ctx, address, contract_snapshot.ContractTypeReservedAliasStorageImpl)
}

func (m *Manager) GetFeeStorageContractAddress(ctx context.Context, address common.Address) (common.Address, error) {
	return m.getAddress(ctx, address, contract_snapshot.ContractTypeFeeStorageImpl)
}

func (m *Manager) GetNewContactStorageAddress(ctx context.Context, address common.Address) (common.Address, error) {
	return m.getAddress(ctx, address, contract_snapshot.ContractTypeNewContractStorage)
}

func (m *Manager) GetDistributionStorageContractAddress(ctx context.Context, address common.Address) (common.Address, error) {
	return m.getAddress(ctx, address, contract_snapshot.ContractTypeDistributionStorageImpl)
}

func (m *Manager) GetMiningAgentStorageContractAddress(ctx context.Context, address common.Address) (common.Address, error) {
	return m.getAddress(ctx, address, contract_snapshot.ContractTypeMiningAgentStorageImpl)
}

func (m *Manager) watchNewContractStorageDeployed(ctx context.Context) (err error) {
	defer func() {
		if err != nil {
			err = bitbonErrors.NewContractWatchError(err)
		}
	}()

	contract, err := NewContractStorageImpl(common.Address{}, m.EthApiWrapper)
	if err != nil {
		return errors.Wrap(err, "error creating instance of contract storage")
	}
	opts := &bind.WatchOpts{
		Start:   nil,
		Context: ctx,
	}
	sink := make(chan *ContractStorageImplNewContractStorageDeployed, 1)
	sub, err := contract.WatchNewContractStorageDeployed(opts, sink)
	if err != nil {
		return errors.Wrap(err, "WatchNewContractStorageDeployed error")
	}

	go func() {
		defer sub.Unsubscribe()

		select {
		case <-m.ready:
			m.logger.Debug("Manager became ready externally. ContractStorageDeployed subscription loop terminated.")
		case <-m.done:
			m.logger.Debug("ContractStorageDeployed subscription loop terminated.")
		case event := <-sink:
			m.logger.Debug("ContractStorageDeployed event received", "contract address", event.NewAddress)
			if err = m.SaveContractStorageAddress(event.NewAddress); err != nil {
				m.logger.Crit("Error setting ContactStorageAddress from (ContractStorageDeployed event)", "err", err, "new address", event.NewAddress.Hex())
			}
		case err, ok := <-sub.Err(): // node stopped
			if !ok {
				m.logger.Debug("ContractStorageDeployed subscription ends")
			} else {
				m.logger.Error("ContractStorageDeployed subscription error", "err", err)
			}
		}
	}()
	return nil
}

func (m *Manager) watchContractStorageMoved(ctx context.Context) (err error) {
	defer func() {
		if err != nil {
			err = bitbonErrors.NewContractWatchError(err)
		}
	}()

	contract, err := NewContractStorageImpl(common.Address{}, m.EthApiWrapper)
	if err != nil {
		return errors.Wrap(err, "error creating instance of contract storage")
	}

	opts := &bind.WatchOpts{
		Start:   nil,
		Context: ctx,
	}
	sink := make(chan *ContractStorageImplContractStorageMoved, 1)
	sub, err := contract.WatchContractStorageMoved(opts, sink)
	if err != nil {
		return errors.Wrap(err, "WatchContractStorageMoved error")
	}

	go func() {
		defer sub.Unsubscribe()

		for {
			select {
			case <-m.done:
				m.logger.Debug("ContractStorageMoved subscription loop terminated.")
				return
			case event := <-sink:
				m.logger.Debug("ContractStorageMoved event received", "contract address", event.NewAddress)
				if !m.Ready() {
					m.logger.Error("Manager got ContractStorageMoved event when it is not ready")
					continue
				}

				address, err := m.GetContractStorageAddress()
				if err != nil {
					m.logger.Error("Unable to get contract storage address", "error", err)
					continue
				}

				newAddress, err := m.GetNewContactStorageAddress(nil, address) // nolint:staticcheck
				if err != nil {
					m.logger.Error("Error getting NewContactStorageAddress", "err", err)
					continue
				}
				if newAddress != event.NewAddress {
					m.logger.Error("New address of ContractStorageMoved event is not equal NewContactStorageAddress method returned", "contract new address", newAddress.Hex(), "event new address", event.NewAddress.Hex())
					continue
				}
				m.logger.Debug("ContractStorageMoved new address approved by ContractStorageImpl")
				if err = m.SaveContractStorageAddress(newAddress); err != nil {
					m.logger.Crit("Error saving NewContactStorageAddress", "err", err, "new address", newAddress.Hex())
				}

			case err, ok := <-sub.Err(): // node stopped
				if !ok {
					m.logger.Debug("ContractStorageMoved subscription ends")
				} else {
					m.logger.Debug("ContractStorageMoved subscription error", "err", err)
				}
				return
			}
		}
	}()
	return nil
}
