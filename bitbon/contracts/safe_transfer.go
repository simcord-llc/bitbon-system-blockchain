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
	"fmt"
	"math/big"
	"time"

	"github.com/simcord-llc/bitbon-system-blockchain/accounts/abi/bind"
	"github.com/simcord-llc/bitbon-system-blockchain/log"

	bitbonErrors "github.com/simcord-llc/bitbon-system-blockchain/bitbon/errors"
)

func (m *Manager) getSafeTransferStorageImpl() (*SafeTransferStorageImpl, error) {
	if !m.Ready() {
		return nil, ErrorNotReady
	}
	contractAddress, err := m.GetContractAddressSafeTransferStorage()
	if err != nil {
		return nil, err
	}
	contract, err := NewSafeTransferStorageImpl(contractAddress, m.EthApiWrapper)
	if err != nil {
		return nil, err
	}
	return contract, nil
}

func (m *Manager) GetTransferLength(ctx context.Context) (transferLength *big.Int, err error) {
	m.logger.Debug("contracts manager getTransferLength called")

	defer func() {
		if err != nil {
			err = bitbonErrors.NewContractCallError(err)
		}
	}()

	contract, err := m.getSafeTransferStorageImpl()
	if err != nil {
		return
	}

	opts := &bind.CallOpts{
		Pending: false,
		Context: ctx,
	}
	return contract.GetTransferLength(opts)
}

func (m *Manager) GetOldestPending(ctx context.Context) (oldestPendingId *big.Int, err error) {
	m.logger.Debug("contracts manager getOldestPending called")

	defer func() {
		if err != nil {
			err = bitbonErrors.NewContractCallError(err)
		}
	}()

	contract, err := m.getSafeTransferStorageImpl()
	if err != nil {
		return
	}

	opts := &bind.CallOpts{
		Pending: false,
		Context: ctx,
	}
	return contract.GetOldestPending(opts)
}

func (m *Manager) SearchOldestPending(ctx context.Context, firstTransfer *big.Int, lastTransfer *big.Int) (index *big.Int, present bool, err error) {
	m.logger.Debug("contracts manager SearchOldestPending called", "first transfer num", firstTransfer, "last transfer num", lastTransfer)

	defer func() {
		if err != nil {
			err = bitbonErrors.NewContractCallError(err)
		}
	}()

	contract, err := m.getSafeTransferStorageImpl()
	if err != nil {
		return nil, false, err
	}

	opts := &bind.CallOpts{
		Pending: false,
		Context: ctx,
	}
	res, err := contract.SearchOldestPending(opts, firstTransfer, lastTransfer)
	if err != nil {
		return nil, false, err
	}

	return res.Index, res.IsPresent, nil
}

func (m *Manager) GetExpiredTransfers(ctx context.Context, firstTransfer *big.Int, lastTransfer *big.Int) (expireTransferIds [][32]byte, err error) {
	m.logger.Debug("contracts manager getExpiredTransfers called", "first transfer num", firstTransfer, "last transfer num", lastTransfer)

	defer func() {
		if err != nil {
			err = bitbonErrors.NewContractCallError(err)
		}
	}()

	contract, err := m.getSafeTransferStorageImpl()
	if err != nil {

		return nil, err
	}

	opts := &bind.CallOpts{
		Pending: false,
		Context: ctx,
	}
	return contract.GetExpiredTransfers(opts, firstTransfer, lastTransfer)
}

func (m *Manager) SetOldestPending(ctx context.Context, pendingIndex *big.Int, key *ecdsa.PrivateKey) (err error) {
	m.logger.Debug("contracts manager SetOldestPending called", "pendingIndex", pendingIndex)
	defer func() {
		if err != nil {
			err = bitbonErrors.NewContractCallError(err)
		}
	}()

	contract, err := m.getSafeTransferStorageImpl()
	if err != nil {
		return
	}

	opts, err := m.prepareTransactOpts(ctx, key)
	if err != nil {
		return
	}

	tx, err := contract.SetOldestPending(opts, pendingIndex)
	if err != nil {
		return
	}

	m.logger.Debug("Waiting mining tx SetOldestPending", "oldestPendingIndex", pendingIndex, "txHash", tx.Hash())
	_, err = m.waitAndCheckTransaction(ctx, tx)

	return
}

func (m *Manager) GetTransferIdByHash(ctx context.Context, transferIdHash [32]byte) (transferIds []byte, err error) {
	m.logger.Debug("contracts manager GetTransferIdByHash called", "transferIdHAsh", transferIdHash)
	defer func() {
		if err != nil {
			err = bitbonErrors.NewContractCallError(err)
		}
	}()

	contract, err := m.getSafeTransferStorageImpl()
	if err != nil {
		return
	}

	opts := &bind.CallOpts{
		Pending: false,
		Context: ctx,
	}
	res, err := contract.GetTransferIdByHash(opts, transferIdHash)
	if err != nil {
		return nil, err
	}
	return res, nil
}

//TransferExists checks if the transfer with id transferId exist in blockchain
func (m *Manager) TransferExists(ctx context.Context, transferId []byte) (exists bool, err error) {
	m.logger.Debug("contracts manager TransferExists called", "transferId", string(transferId))
	defer func() {
		if err != nil {
			err = bitbonErrors.NewContractCallError(err)
		}
	}()

	contract, err := m.getSafeTransferStorageImpl()
	if err != nil {
		return false, err
	}

	opts := &bind.CallOpts{
		Pending: false,
		Context: ctx,
	}

	return contract.TransferExists(opts, transferId)
}

func (m *Manager) GetTransfer(ctx context.Context, transferId []byte) (transfer *ReceiptTransfer, err error) {
	m.logger.Debug("contracts manager getTransfer called", "transferId", string(transferId))

	defer func() {
		if err != nil {
			err = bitbonErrors.NewContractCallError(err)
		}
	}()

	contract, err := m.getSafeTransferStorageImpl()
	if err != nil {
		return
	}
	opts := &bind.CallOpts{
		Pending: false,
		Context: ctx,
	}
	res, err := contract.GetTransfer(opts, transferId)
	if err != nil {
		return nil, err
	}
	return &ReceiptTransfer{
		State:          res.State,
		TransferIdHash: res.TransferIdHash,
		Source:         res.Source,
		Dest:           res.Dest,
		Amount:         res.Amount,
		Proof:          res.Proof,
		Vk:             res.Vk,
		CreatedAt:      res.CreatedAt,
		ExpiresAt:      res.ExpiresAt,
		CurrentRetries: res.CurrentRetries,
		MaxRetries:     res.MaxRetries,
		Canceller:      res.Canceller,
	}, nil
}

func (m *Manager) GetTransferByIndex(ctx context.Context, index *big.Int) (transfer *ReceiptTransfer, err error) {
	m.logger.Debug("contracts manager GetTransferByIndex called", "index", index)

	defer func() {
		if err != nil {
			err = bitbonErrors.NewContractCallError(err)
		}
	}()

	contract, err := m.getSafeTransferStorageImpl()
	if err != nil {
		return
	}
	opts := &bind.CallOpts{
		Pending: false,
		Context: ctx,
	}
	hash, err := contract.TransferIds(opts, index)
	if err != nil {
		return nil, err
	}
	m.logger.Debug("contracts manager GetTransferByIndex got contract hash by index", "index", index, "hash", hash)

	res, err := contract.GetTransferByHash(opts, hash)
	if err != nil {
		return nil, err
	}
	return &ReceiptTransfer{
		State:          res.State,
		TransferIdHash: res.TransferIdHash,
		Source:         res.Source,
		Dest:           res.Dest,
		Amount:         res.Amount,
		Proof:          res.Proof,
		Vk:             res.Vk,
		CreatedAt:      res.CreatedAt,
		ExpiresAt:      res.ExpiresAt,
		CurrentRetries: res.CurrentRetries,
		MaxRetries:     res.MaxRetries,
		Canceller:      res.Canceller,
	}, nil
}

// WatchTransferExpired watches TransferExpired(bytes) blockchain events(logs)
func (m *Manager) WatchTransferExpired(sink chan<- *TransferExpired) (err error) {
	defer func() {
		if err != nil {
			err = bitbonErrors.NewContractWatchError(err)
		}
	}()

	contract, err := m.getSafeTransferStorageImpl()
	if err != nil {
		return err
	}
	opts := &bind.WatchOpts{}

	internalSink := make(chan *SafeTransferStorageImplTransferExpired, len(sink))
	sub, err := contract.WatchTransferExpired(opts, internalSink)
	if err != nil {
		return err
	}

	redeploySignal, unsubscribe := m.SubscribeRedeploy()

	m.logger.Debug("WatchTransferExpired started.")
	go func() {
		defer func() {
			sub.Unsubscribe()
			unsubscribe()
			m.logger.Debug("WatchTransferExpired unsubscribed.")
		}()

		for {
			select {
			case <-redeploySignal:
				m.logger.Debug("WatchTransferExpired got redeploy signal. Restarting...")
				retryTimeout := 10 * time.Second
				for err = m.WatchTransferExpired(sink); err != nil; {
					log.Error(fmt.Sprintf("error beginning WatchTransferExpired. Retry in %s", retryTimeout))
					time.Sleep(retryTimeout)
				}
				return
			case <-m.done:
				m.logger.Debug("WatchTransferExpired subscription loop terminated.")
				return
			case err, ok := <-sub.Err():
				if !ok {
					m.logger.Debug("WatchTransferExpired subscription ends")
				} else {
					log.Error("WatchTransferExpired subscription error", "err", err)
				}
				return
			case transferExpired := <-internalSink:
				sink <- &TransferExpired{
					TransferID: transferExpired.TransferId,
				}
			}
		}
	}()
	return nil
}
