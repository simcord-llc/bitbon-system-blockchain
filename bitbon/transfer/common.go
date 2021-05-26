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
package transfer

import (
	"context"
	"math/big"

	"github.com/pkg/errors"
	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/models"
	"github.com/simcord-llc/bitbon-system-blockchain/crypto"
	"github.com/simcord-llc/bitbon-system-blockchain/zksnark"
)

var (
	ErrAlreadyStopped        = errors.New("already stopped")
	ErrZkSnarkProofEmpty     = errors.New("empty  ZK-snark proof")
	ErrTransferNotReady      = errors.New("transfer manager is not ready")
	ErrFromRequire           = errors.New("from is required")
	ErrToRequire             = errors.New("to is required")
	ErrSameAccount           = errors.New("transfer to the same account (From equals To)")
	ErrAccountIDRequire      = errors.New("accountID is required")
	ErrValueRequire          = errors.New("value is required")
	ErrAssetboxNotAssigned   = errors.New("assetbox is not assigned to any account")
	ErrAccountIDMismatch     = errors.New("transfer and assetbox AccountID mismatch")
	ErrBTSCNotValid          = errors.New("btsc is not valid")
	ErrLowAssetboxBalance    = errors.New("assetbox balance to low to perform transfer")
	ErrTransferIDRequire     = errors.New("transferID is required")
	ErrProtectionCodeRequire = errors.New("protectionCode is required")
	ErrExpireAtPassed        = errors.New("expiresAt has already passed")
	ErrAddressRequire        = errors.New("address is required")
	ErrFromNotAssigned       = errors.New("assetbox From is not assigned to any account")
	ErrToNotAssigned         = errors.New("assetbox To is not assign to any account")

	ErrMsgGettingAssetbox            = "error getting decrypted assetbox"
	ErrMsgGettingSourceAssetbox      = "error getting decrypted assetbox From"
	ErrMsgGettingDestinationAssetbox = "error getting decrypted assetbox To"

	SourceAssetboxTag      = "From"
	DestinationAssetboxTag = "To"
	EmptyAssetboxTag       = ""
)

func (tm *Manager) generateZkSnark(transfer *models.CreateTransferObj) (err error) {
	protectionHash := tm.getProtectionHash(transfer.TransferID, transfer.ProtectionCode)
	tm.logger.Debug("protectionHash successfully obtained", "protectionHash", protectionHash)
	var pk []byte
	pk, transfer.VK = tm.getKeys()
	tm.logger.Debug("zk-snrak keys obtaioned (pk, vk)", "len(pk)", len(pk), "len(transfer.VK)", len(transfer.VK))
	transfer.Proof, err = zksnark.Prove(protectionHash, pk, constraintsLen)
	if err != nil {
		return errors.Wrap(err, "error generating ZK-snark params")
	}
	tm.logger.Debug("ZK-snark params successfully generated")

	return nil
}

func (tm *Manager) processPending(ctx context.Context,
	oldestPendingIndex, lastIndex, batchSize *big.Int) (pendingTransferExists bool, err error) {
	for startSearch := new(big.Int).Set(oldestPendingIndex); startSearch.Cmp(lastIndex) <= 0; startSearch.Add(startSearch, batchSize) { // nolint:lll
		newOldestPendingIndex, present, err := tm.bitbon.GetContractsManager().
			SearchOldestPending(ctx, startSearch, new(big.Int).Add(startSearch, batchSize))
		tm.logger.Debug("ExpireTransfers SearchOldestPending called",
			"newOldestPendingIndex", newOldestPendingIndex, "present", present,
			"error", err, "oldestPendingIndex", oldestPendingIndex)

		if err != nil {
			return pendingTransferExists, errors.Wrap(err, "error SearchOldestPending() to safe transfer storage contract")
		}
		if !present {
			continue
		}
		pendingTransferExists = true
		if oldestPendingIndex.Cmp(newOldestPendingIndex) != 0 {
			tm.logger.Debug("ExpireTransfers newOldestPendingIndex found",
				"newOldestPendingIndex", newOldestPendingIndex, "oldestPendingIndex", oldestPendingIndex)

			oldestPendingIndex = newOldestPendingIndex
			err := tm.bitbon.GetContractsManager().SetOldestPending(ctx, oldestPendingIndex, tm.bitbon.GetServicePk())
			if err != nil {
				return pendingTransferExists, errors.Wrap(err, "error SetOldestPending() to safe transfer storage contract")
			}
			tm.logger.Debug("ExpireTransfers newOldestPendingIndex successfully set",
				"newOldestPendingIndex", oldestPendingIndex)
		} else {
			tm.logger.Debug("ExpireTransfers newOldestPendingIndex not found")
		}
		break
	}

	return
}

// nolint:gocritic
func (tm *Manager) getKeys() ([]byte, []byte) {
	tm.keyMux.RLock()
	defer tm.keyMux.RUnlock()
	return tm.pk, tm.vk
}

func (tm *Manager) getProtectionHash(transferID, protectionCode string) []byte {
	return crypto.Keccak256([]byte(transferID), []byte(protectionCode))
}
