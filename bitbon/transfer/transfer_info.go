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

package transfer

import (
	"context"
	"math/big"

	"github.com/pkg/errors"
	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/contracts"
	"github.com/simcord-llc/bitbon-system-blockchain/common"
)

// Returns
// the number of transfers expired (num int)
// the ethereum transaction hashes - contract calls of method ExpireSafeTransfers (hashes []common.Hash)
func (tm *Manager) ExpireTransfers(ctx context.Context) (hashes []common.Hash, num int, err error) {
	tm.logger.Debug("ExpireTransfers called")
	defer func() {
		if err != nil {
			tm.logger.Error("ExpireTransfers ended with error", "error", err)
		}
	}()

	oldestPendingIndex, err := tm.bitbon.GetContractsManager().GetOldestPending(ctx)
	if err != nil {
		return nil, num, errors.Wrap(err, "error GetOldestPending() to safe transfer storage contract")
	}
	tm.logger.Debug("ExpireTransfers", "oldestPendingIndex", oldestPendingIndex)

	transferLength, err := tm.bitbon.GetContractsManager().GetTransferLength(ctx)
	if err != nil {
		return nil, num, errors.Wrap(err, "error GetTransferLength() to safe transfer storage contract")
	}
	tm.logger.Debug("ExpireTransfers", "transferLength", transferLength)

	lastIndex := new(big.Int).Sub(transferLength, big.NewInt(1))
	if oldestPendingIndex.Cmp(lastIndex) == 0 {
		tm.logger.Debug("not expire safe transfer")
		return nil, num, nil
	}

	batchSize := big.NewInt(batchSearchSize)
	pendingTransferExists, err := tm.processPending(ctx, oldestPendingIndex, lastIndex, batchSize)

	if !pendingTransferExists {
		tm.logger.Info("no safe transfers to expire")
		return nil, num, nil
	}

	batchExpireSize := big.NewInt(batchExpireSize)
	for startSearch := new(big.Int).Set(oldestPendingIndex); startSearch.Cmp(lastIndex) <= 0; startSearch.Add(startSearch, batchExpireSize) { // nolint:lll
		endSearch := new(big.Int).Add(startSearch, batchExpireSize)
		ids, err := tm.bitbon.GetContractsManager().GetExpiredTransfers(ctx, startSearch, endSearch)
		tm.logger.Info("GetExpiredTransfers returned", "len(ids)", len(ids))
		if err != nil {
			return nil, num, errors.Wrap(err, "error GetExpiredTransfers() to safe transfer storage contract")
		}
		if len(ids) > 0 {
			hash, err := tm.bitbon.GetContractsManager().ExpireSafeTransfers(ctx, ids, tm.bitbon.GetServicePk())
			if err != nil {
				return hashes, num, errors.Wrap(err, "error ExpireSafeTransfers() to safe transfer storage contract")
			}
			num += len(ids)
			hashes = append(hashes, hash)
		}
	}
	return hashes, num, nil
}

func (tm *Manager) GetTransfer(ctx context.Context, transferID string) (*contracts.ReceiptTransfer, error) {
	tm.logger.Trace("GetTransfer called")

	return tm.bitbon.GetContractsManager().GetTransfer(ctx, []byte(transferID))
}

func (tm *Manager) GetTransferByIndex(ctx context.Context, index *big.Int) (*contracts.ReceiptTransfer, error) {
	tm.logger.Trace("GetTransferByIndex called")

	return tm.bitbon.GetContractsManager().GetTransferByIndex(ctx, index)
}

func (tm *Manager) TransferExists(ctx context.Context, transferID string) (bool, error) {
	tm.logger.Trace("TransferExists called")

	return tm.bitbon.GetContractsManager().TransferExists(ctx, []byte(transferID))
}
