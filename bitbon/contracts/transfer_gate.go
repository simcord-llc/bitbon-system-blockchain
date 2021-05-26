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
	"fmt"
	"time"

	"github.com/simcord-llc/bitbon-system-blockchain/accounts/abi/bind"
	bitbonErrors "github.com/simcord-llc/bitbon-system-blockchain/bitbon/errors"
	"github.com/simcord-llc/bitbon-system-blockchain/log"
)

func (m *Manager) getTransferImpl() (*TransferImpl, error) {
	if !m.Ready() {
		return nil, ErrorNotReady
	}
	contractAddress, err := m.GetContractAddressTransfer()
	if err != nil {
		return nil, err
	}
	contract, err := NewTransferImpl(contractAddress, m.EthApiWrapper)
	if err != nil {
		return nil, err
	}
	return contract, nil
}

// WatchBitbonBalanceChanged watches balance changes blockchain events(logs)
func (m *Manager) WatchBitbonBalanceChanged(sink chan<- *Balance) (err error) {
	defer func() {
		if err != nil {
			err = bitbonErrors.NewContractWatchError(err)
		}
	}()

	contract, err := m.getTransferImpl()
	if err != nil {
		return err
	}
	opts := &bind.WatchOpts{}

	internalSink := make(chan *TransferImplBalanceChanged, len(sink))
	sub, err := contract.WatchBalanceChanged(opts, internalSink)
	if err != nil {
		return err
	}

	redeploySignal, unsubscribe := m.SubscribeRedeploy()

	m.logger.Debug("WatchBitbonBalanceChanged started.")
	go func() {
		defer func() {
			sub.Unsubscribe()
			unsubscribe()
			m.logger.Debug("WatchBitbonBalanceChanged unsubscribed.")
		}()

		for {
			select {
			case <-redeploySignal:
				m.logger.Debug("WatchBitbonBalanceChanged got redeploy signal. Restarting...")
				retryTimeout := 10 * time.Second
				for err = m.WatchBitbonBalanceChanged(sink); err != nil; {
					m.logger.Error(fmt.Sprintf("error beginning WatchBitbonBalanceChanged. Retry in %s", retryTimeout))
					time.Sleep(retryTimeout)
				}
				return
			case <-m.done:
				m.logger.Debug("WatchBitbonBalanceChanged subscription loop terminated.")
				return
			case err, ok := <-sub.Err():
				if !ok {
					m.logger.Debug("WatchBitbonBalanceChanged subscription ends")
				} else {
					m.logger.Error("WatchBitbonBalanceChanged subscription error", "err", err)
				}
				return
			case balance := <-internalSink:
				sink <- &Balance{
					Address:      balance.Assetbox,
					Balance:      balance.Balance,
					AssetboxType: balance.AssetboxType,
				}
			}
		}
	}()
	return nil
}

// WatchBitbonBalanceLocked watches balance locked blockchain events(logs)
func (m *Manager) WatchBitbonBalanceLocked(sink chan<- *Balance) (err error) {
	defer func() {
		if err != nil {
			err = bitbonErrors.NewContractWatchError(err)
		}
	}()

	contract, err := m.getTransferImpl()
	if err != nil {
		return err
	}
	opts := &bind.WatchOpts{}

	internalSink := make(chan *TransferImplBalanceLocked, len(sink))
	sub, err := contract.WatchBalanceLocked(opts, internalSink)
	if err != nil {
		return err
	}

	redeploySignal, unsubscribe := m.SubscribeRedeploy()

	m.logger.Debug("WatchBitbonBalanceLocked started.")
	go func() {
		defer func() {
			sub.Unsubscribe()
			unsubscribe()
			m.logger.Debug("WatchBitbonBalanceLocked unsubscribed.")
		}()

		for {
			select {
			case <-redeploySignal:
				m.logger.Debug("WatchBitbonBalanceLocked got redeploy signal. Restarting...")
				retryTimeout := 10 * time.Second
				for err = m.WatchBitbonBalanceLocked(sink); err != nil; {
					m.logger.Error(fmt.Sprintf("error beginning WatchBitbonBalanceLocked. Retry in %s", retryTimeout))
					time.Sleep(retryTimeout)
				}
				return
			case <-m.done:
				m.logger.Debug("WatchBitbonBalanceLocked subscription loop terminated.")
				return
			case err, ok := <-sub.Err():
				if !ok {
					m.logger.Debug("WatchBitbonBalanceLocked subscription ends")
				} else {
					log.Error("WatchBitbonBalanceLocked subscription error", "err", err)
				}
				return
			case balance := <-internalSink:
				sink <- &Balance{
					Address:      balance.Assetbox,
					Balance:      balance.BalanceAviable,
					AssetboxType: balance.AssetboxType,
				}
			}
		}
	}()
	return nil
}

// WatchBitbonBalanceUnLocked watches balance unlocked blockchain events(logs)
func (m *Manager) WatchBitbonBalanceUnLocked(sink chan<- *Balance) (err error) {
	defer func() {
		if err != nil {
			err = bitbonErrors.NewContractWatchError(err)
		}
	}()

	contract, err := m.getTransferImpl()
	if err != nil {
		return err
	}
	opts := &bind.WatchOpts{}

	internalSink := make(chan *TransferImplBalanceUnLocked, len(sink))
	sub, err := contract.WatchBalanceUnLocked(opts, internalSink)
	if err != nil {
		return err
	}

	redeploySignal, unsubscribe := m.SubscribeRedeploy()

	m.logger.Debug("WatchBitbonBalanceUnLocked started.")
	go func() {
		defer func() {
			sub.Unsubscribe()
			unsubscribe()
			m.logger.Debug("WatchBitbonBalanceUnLocked unsubscribed.")
		}()

		for {
			select {
			case <-redeploySignal:
				m.logger.Debug("WatchBitbonBalanceUnLocked got redeploy signal. Restarting...")
				retryTimeout := 10 * time.Second
				for err = m.WatchBitbonBalanceUnLocked(sink); err != nil; {
					m.logger.Error(fmt.Sprintf("error beginning WatchBitbonBalanceUnLocked. Retry in %s", retryTimeout))
					time.Sleep(retryTimeout)
				}
				return
			case <-m.done:
				m.logger.Debug("WatchBitbonBalanceUnLocked subscription loop terminated.")
				return
			case err, ok := <-sub.Err():
				if !ok {
					m.logger.Debug("WatchBitbonBalanceUnLocked subscription ends")
				} else {
					log.Error("WatchBitbonBalanceUnLocked subscription error", "err", err)
				}
				return
			case balance := <-internalSink:
				sink <- &Balance{
					Address:      balance.Assetbox,
					Balance:      balance.BalanceAviable,
					AssetboxType: balance.AssetboxType,
				}
			}
		}
	}()
	return nil
}
