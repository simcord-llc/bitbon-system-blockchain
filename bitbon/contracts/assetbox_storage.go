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

func (m *Manager) getAssetboxStorageImpl() (*AssetboxStorageImpl, error) {
	if !m.Ready() {
		return nil, ErrorNotReady
	}
	contractAddress, err := m.GetContractAddressAssetboxInfo()
	if err != nil {
		return nil, err
	}
	contract, err := NewAssetboxStorageImpl(contractAddress, m.EthApiWrapper)
	if err != nil {
		return nil, err
	}
	return contract, nil
}

// WatchAssetboxInfoSet watches public assetbox info blockchain events(logs)
func (m *Manager) WatchAssetboxInfoSet(sink chan<- *Assetbox) (err error) {
	defer func() {
		if err != nil {
			err = bitbonErrors.NewContractWatchError(err)
		}
	}()

	contract, err := m.getAssetboxStorageImpl()
	if err != nil {
		return err
	}
	opts := &bind.WatchOpts{}

	internalSink := make(chan *AssetboxStorageImplAssetboxInfoSet, len(sink))
	sub, err := contract.WatchAssetboxInfoSet(opts, internalSink)
	if err != nil {
		return err
	}

	redeploySignal, unsubscribe := m.SubscribeRedeploy()

	m.logger.Debug("WatchAssetboxInfoSet started.")
	go func() {
		defer func() {
			sub.Unsubscribe()
			unsubscribe()
			m.logger.Debug("WatchAssetboxInfoSet unsubscribed.")
		}()

		for {
			select {
			case <-redeploySignal:
				m.logger.Debug("WatchAssetboxInfoSet got redeploy signal. Restarting...")
				retryTimeout := 10 * time.Second
				for err = m.WatchAssetboxInfoSet(sink); err != nil; {
					m.logger.Error(fmt.Sprintf("error beginning WatchAssetboxInfoSet. Retry in %s", retryTimeout))
					time.Sleep(retryTimeout)
				}
				return
			case <-m.done:
				m.logger.Debug("WatchAssetboxInfoSet subscription loop terminated.")
				return
			case err, ok := <-sub.Err():
				if !ok {
					m.logger.Debug("WatchAssetboxInfoSet subscription ends")
				} else {
					m.logger.Debug("WatchAssetboxInfoSet subscription error", "err", err)
				}
				return
			case info := <-internalSink:
				ass, err := ContractAssetbox(info.Assetbox, info.AliasString, info.ServiceId, info.OtherInfo, info.IsMining)
				if err != nil {
					m.logger.Error("WatchAssetboxInfoSet error decoding assetbox other info", "error", err)
					continue
				}

				sink <- ass
			}
		}
	}()
	return nil
}

// WatchAssetboxInfoDeleted watches public assetbox info deletion blockchain events(logs)
func (m *Manager) WatchAssetboxInfoDeleted(sink chan<- *Assetbox) (err error) {
	defer func() {
		if err != nil {
			err = bitbonErrors.NewContractWatchError(err)
		}
	}()

	contract, err := m.getAssetboxStorageImpl()
	if err != nil {
		return err
	}
	opts := &bind.WatchOpts{}

	internalSink := make(chan *AssetboxStorageImplAssetboxInfoDeleted, len(sink))
	sub, err := contract.WatchAssetboxInfoDeleted(opts, internalSink)
	if err != nil {
		return err
	}

	redeploySignal, unsubscribe := m.SubscribeRedeploy()

	log.Warn("WatchAssetboxInfoDeleted started.")
	go func() {
		defer func() {
			sub.Unsubscribe()
			unsubscribe()
			log.Warn("WatchAssetboxInfoDeleted unsubscribed.")
		}()

		for {
			select {
			case <-redeploySignal:
				log.Warn("WatchAssetboxInfoDeleted got redeploy signal. Restarting...")
				retryTimeout := 10 * time.Second
				for err = m.WatchAssetboxInfoDeleted(sink); err != nil; {
					log.Error(fmt.Sprintf("error beginning WatchAssetboxInfoDeleted. Retry in %s", retryTimeout))
					time.Sleep(retryTimeout)
				}
				return
			case <-m.done:
				log.Warn("WatchAssetboxInfoDeleted subscription loop terminated.")
				return
			case err, ok := <-sub.Err():
				if !ok {
					log.Warn("WatchAssetboxInfoDeleted subscription ends")
				} else {
					log.Error("WatchAssetboxInfoDeleted subscription error", "err", err)
				}
				return
			case info := <-internalSink:
				ass, err := ContractAssetbox(info.Assetbox, info.AliasString, info.ServiceId, info.OtherInfo, info.IsMining)
				if err != nil {
					log.Error("WatchAssetboxInfoDeleted error decoding assetbox other info", "error", err)
					continue
				}

				sink <- ass
			}
		}
	}()
	return nil
}
