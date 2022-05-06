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
	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/models"
	"github.com/simcord-llc/bitbon-system-blockchain/crypto"
	"math/big"

	"github.com/pkg/errors"
	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/contracts"
)

const (
	expireMaxBatchSize = 20
)

var (
	errMaxBatchSizeExceeded = errors.Errorf("max batch size of %d exceeded", expireMaxBatchSize)
)

// nolint:godox
func (tm *Manager) ExpireTransfersAsync(ctx context.Context, ids []string) (*models.TransferResponseObj, error) {
	tm.logger.Debug("ExpireTransfersAsync called")

	if len(ids) > expireMaxBatchSize {
		return nil, errMaxBatchSizeExceeded
	}

	// todo doesn't look good, maybe can be done in a nice way?
	idsBytes := make([][32]byte, len(ids))
	for i := range ids {
		var idBytes [32]byte
		copy(idBytes[:], crypto.Keccak256([]byte(ids[i])))

		idsBytes[i] = idBytes
	}

	blockNumber, txHash, err := tm.bitbon.GetContractsManager().ExpireSafeTransfers(ctx, idsBytes, tm.bitbon.GetServicePk())
	if err != nil {
		return nil, errors.Wrap(err, "error expiring safe transfers")
	}

	return &models.TransferResponseObj{
		BlockNumber: blockNumber,
		TxHash:      txHash,
	}, nil
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

func (tm *Manager) GetTransferState(ctx context.Context, transferID string) (uint8, error) {
	tm.logger.Trace("TransferExists called")

	return tm.bitbon.GetContractsManager().GetTransferState(ctx, []byte(transferID))
}
