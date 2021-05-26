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
	bitbonErrors "github.com/simcord-llc/bitbon-system-blockchain/bitbon/errors"
	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/models"
	"github.com/simcord-llc/bitbon-system-blockchain/common"
)

func (tm *Manager) DirectTransfer(ctx context.Context,
	transfer *models.DirectTransferObj) (*models.TransferResponseObj, error) {
	return tm.directTransfer(ctx, transfer, false)
}

func (tm *Manager) DirectTransferAsync(ctx context.Context,
	transfer *models.DirectTransferObj) (*models.TransferResponseObj, error) {
	return tm.directTransfer(ctx, transfer, true)
}

func (tm *Manager) directTransfer(ctx context.Context,
	transfer *models.DirectTransferObj, async bool) (*models.TransferResponseObj, error) {
	tm.logger.Debug("TransferManager directTransfer called",
		"from", transfer.From.Hex(), "to", transfer.To.Hex(), "value", transfer.Value.String())

	if err := tm.directTransferBaseCheck(transfer); err != nil {
		return nil, err
	}

	// check balance
	currentBalance, err := tm.bitbon.GetContractsManager().GetAssetboxBalance(ctx, transfer.From)
	if err != nil {
		return nil, errors.Wrap(err, "error getting current assetbox balance")
	}

	if currentBalance.Cmp(transfer.Value) < 0 {
		return nil, bitbonErrors.NewBalanceTooLowError(ErrLowAssetboxBalance)
	}

	// call smart contract method
	blockNum, txHash, err := tm.bitbon.GetContractsManager().DirectTransfer(ctx, transfer.From, transfer.To,
		transfer.Value, transfer.ExtraData, tm.bitbon.GetServicePk(), async)
	if err != nil {
		return nil, errors.Wrap(err, "error doing direct transfer in blockchain")
	}
	tm.logger.Debug("Contract manager DirectTransfer results", "blockNum", blockNum, "txHash", txHash.Hex())

	return &models.TransferResponseObj{
		BlockNumber: blockNum,
		TxHash:      txHash,
	}, nil
}

func (tm *Manager) directTransferBaseCheck(transfer *models.DirectTransferObj) error {
	if transfer.From == (common.Address{}) {
		return bitbonErrors.NewInvalidParamsError(ErrFromRequire)
	}
	if transfer.To == (common.Address{}) {
		return bitbonErrors.NewInvalidParamsError(ErrToRequire)
	}
	if transfer.From == transfer.To {
		return bitbonErrors.NewInvalidParamsError(ErrSameAccount)
	}
	if transfer.Value == nil || transfer.Value.Cmp(big.NewInt(0)) <= 0 {
		return bitbonErrors.NewInvalidParamsError(ErrValueRequire)
	}

	return nil
}