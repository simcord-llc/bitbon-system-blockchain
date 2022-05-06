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

	"github.com/pkg/errors"
	loggerContext "github.com/simcord-llc/bitbon-system-blockchain/bitbon/context"
	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/dto"
	bitbonErrors "github.com/simcord-llc/bitbon-system-blockchain/bitbon/errors"
	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/models"
	"github.com/simcord-llc/bitbon-system-blockchain/common"
)

func (tm *Manager) serviceFeeTransfer(ctx context.Context, transfer *models.ServiceFeeTransferObj, async bool) (*models.TransferResponseObj, error) {
	if !tm.Ready() {
		return nil, ErrTransferNotReady
	}
	logger := loggerContext.LoggerFromContext(ctx)

	logger.Warn("TransferManager serviceFeeTransfer called", "from", transfer.From.Hex(), "operationType", transfer.OperationType.String())

	// basic validation
	if transfer.From == (common.Address{}) {
		return nil, bitbonErrors.NewInvalidParamsError(ErrFromRequire)
	}

	// check balance
	currentBalance, err := tm.bitbon.GetContractsManager().GetAssetboxBalance(ctx, transfer.From)
	if err != nil {
		return nil, errors.Wrap(err, "error getting current assetbox balance")
	}

	operationFee, err := tm.bitbon.GetFeeManager().GetFee(ctx, dto.OperationType(transfer.OperationType.Int64())) // nolint
	if currentBalance.Cmp(operationFee) < 0 {
		return nil, bitbonErrors.NewBalanceTooLowError(ErrLowAssetboxBalance)
	}

	params := dto.ServiceFeeTransferParams{
		OperationType: transfer.OperationType,
		From:          transfer.From,
		Key:           tm.bitbon.GetServicePk(),
	}
	blockNum, txHash, err := tm.bitbon.GetContractsManager().ServiceFeeTransfer(ctx, params, async)

	logger.Info("Contract manager serviceFeeTransfer results", "blockNum", blockNum, "txHash", txHash.Hex(), "error", err)
	if err != nil {
		return nil, errors.Wrap(err, "error creating quick transfer in blockchain")
	}

	return &models.TransferResponseObj{
		BlockNumber: blockNum,
		TxHash:      txHash,
	}, nil
}

// Public methods for different type of transfers

func (tm *Manager) ServiceFeeTransfer(ctx context.Context, transfer *models.ServiceFeeTransferObj) (*models.TransferResponseObj, error) {
	return tm.serviceFeeTransfer(ctx, transfer, false)
}

func (tm *Manager) ServiceFeeTransferAsync(ctx context.Context, transfer *models.ServiceFeeTransferObj) (*models.TransferResponseObj, error) {
	return tm.serviceFeeTransfer(ctx, transfer, true)
}
