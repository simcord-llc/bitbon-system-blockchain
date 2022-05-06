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
	"fmt"
	"github.com/simcord-llc/bitbon-system-blockchain/accounts/abi/bind"
	loggerContext "github.com/simcord-llc/bitbon-system-blockchain/bitbon/context"
	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/dto"
	bitbonErrors "github.com/simcord-llc/bitbon-system-blockchain/bitbon/errors"
	"github.com/simcord-llc/bitbon-system-blockchain/common"
	"github.com/simcord-llc/bitbon-system-blockchain/log"
	"math/big"
	"time"
)

func (m *Manager) getFeeStorageImpl() (*FeeStorageImpl, error) {
	if !m.Ready() {
		return nil, ErrorNotReady
	}
	contractAddress, err := m.GetContractAddressFeeStorage()
	if err != nil {
		return nil, err
	}
	contract, err := NewFeeStorageImpl(contractAddress, m.EthApiWrapper)
	if err != nil {
		return nil, err
	}
	return contract, nil
}

func (m *Manager) GetFee(ctx context.Context, opType dto.OperationType) (fee *big.Int, err error) {
	logger := loggerContext.LoggerFromContext(ctx)
	logger.Warn("contracts manager GetFee called")

	defer func() {
		if err != nil {
			err = bitbonErrors.NewContractCallError(err)
		}
	}()

	contract, err := m.getFeeStorageImpl()
	if err != nil {
		return nil, err
	}

	opts := &bind.CallOpts{
		Pending: false,
		From:    m.serviceAddress,
		Context: ctx,
	}

	fee, err = contract.Get(opts, big.NewInt(int64(opType)))
	if err != nil {
		return nil, err
	}

	return
}

func (m *Manager) GetFeeDistributionAccounts(ctx context.Context, operationType *big.Int) (accounts []common.Address, err error) {
	logger := loggerContext.LoggerFromContext(ctx)
	logger.Warn("contracts manager GetFeeDistributionAccounts called")

	defer func() {
		if err != nil {
			err = bitbonErrors.NewContractCallError(err)
		}
	}()

	contract, err := m.getFeeStorageImpl()
	if err != nil {
		return nil, err
	}

	opts := &bind.CallOpts{
		Pending: false,
		From:    m.serviceAddress,
		Context: ctx,
	}

	accounts, err = contract.GetFeeDistributionAccounts(opts, operationType)
	if err != nil {
		return nil, err
	}

	return
}

func (m *Manager) GetFeeDistributionAmounts(ctx context.Context, operationType *big.Int) (amounts []*big.Int, err error) {
	logger := loggerContext.LoggerFromContext(ctx)
	logger.Warn("contracts manager GetFeeDistributionAmounts called")

	defer func() {
		if err != nil {
			err = bitbonErrors.NewContractCallError(err)
		}
	}()

	contract, err := m.getFeeStorageImpl()
	if err != nil {
		return nil, err
	}

	opts := &bind.CallOpts{
		Pending: false,
		From:    m.serviceAddress,
		Context: ctx,
	}

	amounts, err = contract.GetFeeDistributionAmounts(opts, operationType)
	if err != nil {
		return nil, err
	}

	return
}

func (m *Manager) GetFeeValueSettings(ctx context.Context) (optTypes []*big.Int, feeValues []*big.Int, err error) {
	logger := loggerContext.LoggerFromContext(ctx)
	logger.Warn("contracts manager GetFeeValueSettings called")

	defer func() {
		if err != nil {
			err = bitbonErrors.NewContractCallError(err)
		}
	}()

	contract, err := m.getFeeStorageImpl()
	if err != nil {
		return nil, nil, err
	}

	opts := &bind.CallOpts{
		Pending: false,
		From:    m.serviceAddress,
		Context: ctx,
	}

	optTypes, feeValues, err = contract.GetFeeValueSettings(opts)
	if err != nil {
		return nil, nil, err
	}

	return
}

func (m *Manager) WatchFeeValueChanged(sink chan<- *FeeValueChanged) (err error) {
	defer func() {
		if err != nil {
			err = bitbonErrors.NewContractWatchError(err)
		}
	}()

	contract, err := m.getFeeStorageImpl()
	if err != nil {
		return err
	}
	opts := &bind.WatchOpts{}

	internalSink := make(chan *FeeStorageImplFeeValueChanged, len(sink))
	sub, err := contract.WatchFeeValueChanged(opts, internalSink)
	if err != nil {
		return err
	}

	redeploySignal, unsubscribe := m.SubscribeRedeploy()

	m.logger.Debug("WatchFeeValueChanged started.")
	go func() {
		defer func() {
			sub.Unsubscribe()
			unsubscribe()
			m.logger.Debug("WatchFeeValueChanged unsubscribed.")
		}()

		for {
			select {
			case <-m.done:
				m.logger.Debug("WatchFeeValueChanged subscription loop terminated.")
				return
			case <-redeploySignal:
				m.logger.Debug("WatchFeeValueChanged got redeploy signal. Restarting...")
				retryTimeout := 10 * time.Second
				for err = m.WatchFeeValueChanged(sink); err != nil; {
					m.logger.Error(fmt.Sprintf("error beginning WatchFeeValueChanged. Retry in %s", retryTimeout))
					time.Sleep(retryTimeout)
				}
				return
			case err, ok := <-sub.Err():
				if !ok {
					m.logger.Debug("WatchFeeValueChanged subscription ends")
				} else {
					log.Error("WatchFeeValueChanged subscription error", "err", err)
				}
				return
			case e := <-internalSink:
				sink <- &FeeValueChanged{
					Type:  e.OpType,
					Value: e.Value,
				}
			}
		}
	}()

	return
}

func (m *Manager) WatchExceptionalAccountsChanged(sink chan<- *ExceptionalAccountsChanged) (err error) {
	defer func() {
		if err != nil {
			err = bitbonErrors.NewContractWatchError(err)
		}
	}()

	contract, err := m.getFeeStorageImpl()
	if err != nil {
		return err
	}
	opts := &bind.WatchOpts{}

	internalSink := make(chan *FeeStorageImplExceptionalAccountsChanged, len(sink))
	sub, err := contract.WatchExceptionalAccountsChanged(opts, internalSink)
	if err != nil {
		return err
	}

	redeploySignal, unsubscribe := m.SubscribeRedeploy()

	m.logger.Debug("WatchExceptionalAccountsChanged started.")
	go func() {
		defer func() {
			sub.Unsubscribe()
			unsubscribe()
			m.logger.Debug("WatchExceptionalAccountsChanged unsubscribed.")
		}()

		for {
			select {
			case <-m.done:
				m.logger.Debug("WatchExceptionalAccountsChanged subscription loop terminated.")
				return
			case <-redeploySignal:
				m.logger.Debug("WatchExceptionalAccountsChanged got redeploy signal. Restarting...")
				retryTimeout := 10 * time.Second
				for err = m.WatchExceptionalAccountsChanged(sink); err != nil; {
					m.logger.Error(fmt.Sprintf("error beginning WatchExceptionalAccountsChanged. Retry in %s", retryTimeout))
					time.Sleep(retryTimeout)
				}
				return
			case err, ok := <-sub.Err():
				if !ok {
					m.logger.Debug("WatchExceptionalAccountsChanged subscription ends")
				} else {
					log.Error("WatchExceptionalAccountsChanged subscription error", "err", err)
				}
				return
			case e := <-internalSink:
				sink <- &ExceptionalAccountsChanged{
					Accounts: e.Accounts,
				}
			}
		}
	}()

	return
}

func (m *Manager) WatchFeeDistributionSettingsChanged(sink chan<- *DistributionSettingsChanged) (err error) {
	defer func() {
		if err != nil {
			err = bitbonErrors.NewContractWatchError(err)
		}
	}()

	contract, err := m.getFeeStorageImpl()
	if err != nil {
		return err
	}
	opts := &bind.WatchOpts{}

	internalSink := make(chan *FeeStorageImplFeeDistributionChanged, len(sink))
	sub, err := contract.WatchFeeDistributionChanged(opts, internalSink)
	if err != nil {
		return err
	}

	redeploySignal, unsubscribe := m.SubscribeRedeploy()

	m.logger.Debug("WatchFeeDistributionSettingsChanged started.")
	go func() {
		defer func() {
			sub.Unsubscribe()
			unsubscribe()
			m.logger.Debug("WatchFeeDistributionSettingsChanged unsubscribed.")
		}()

		for {
			select {
			case <-m.done:
				m.logger.Debug("WatchFeeDistributionSettingsChanged subscription loop terminated.")
				return
			case <-redeploySignal:
				m.logger.Debug("WatchFeeDistributionSettingsChanged got redeploy signal. Restarting...")
				retryTimeout := 10 * time.Second
				for err = m.WatchFeeDistributionSettingsChanged(sink); err != nil; {
					m.logger.Error(fmt.Sprintf("error beginning DistributionSettingsChanged. Retry in %s", retryTimeout))
					time.Sleep(retryTimeout)
				}
				return
			case err, ok := <-sub.Err():
				if !ok {
					m.logger.Debug("WatchFeeDistributionSettingsChanged subscription ends")
				} else {
					log.Error("WatchFeeDistributionSettingsChanged subscription error", "err", err)
				}
				return
			case e := <-internalSink:
				sink <- &DistributionSettingsChanged{
					Type:     e.OpType,
					Accounts: e.DistributionAccounts,
					Amounts:  e.DistributionAmounts,
				}
			}
		}
	}()

	return
}
