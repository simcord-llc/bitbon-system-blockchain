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

	bb "github.com/simcord-llc/bitbon-system-blockchain/bitbon"

	"github.com/pkg/errors"
	bitbonErrors "github.com/simcord-llc/bitbon-system-blockchain/bitbon/errors"
	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/models"
	"github.com/simcord-llc/bitbon-system-blockchain/common"
)

func (tm *Manager) quickTransfer(ctx context.Context, transfer *models.QuickTransferObj,
	async bool) (*models.TransferResponseObj, error) {
	assetbox := &models.Assetbox{
		Wallet:     transfer.CryptoData.Wallet,
		PassPhrase: transfer.CryptoData.Passphrase,
	}

	if !tm.Ready() {
		return nil, ErrTransferNotReady
	}

	tm.logger.Debug("TransferManager quickTransfer called", "from", transfer.From.Hex(),
		"to", transfer.To.Hex(), "value", transfer.Value.String())

	// basic validation
	if err := tm.quickTransferBaseCheck(transfer); err != nil {
		return nil, err
	}

	if err := bb.DecryptAssetboxWallet(assetbox, tm.bitbon.GetDecryptAssetboxWalletPassword(), tm.encryptor.Decrypt); err != nil {
		return nil, err
	}

	// check balance
	currentBalance, err := tm.bitbon.GetContractsManager().GetAssetboxBalance(ctx, transfer.From)
	if err != nil {
		return nil, errors.Wrap(err, "error getting current fromAssetbox balance")
	}

	if currentBalance.Cmp(transfer.Value) < 0 {
		return nil, bitbonErrors.NewBalanceTooLowError(ErrLowAssetboxBalance)
	}

	blockNum, txHash, err := tm.bitbon.GetContractsManager().QuickTransfer(ctx, transfer.To,
		transfer.Value, transfer.ExtraData, assetbox.Pk, async)

	tm.logger.Debug("Contract manager quickTransfer results", "blockNum", blockNum, "txHash", txHash.Hex(), "error", err)
	if err != nil {
		return nil, errors.Wrap(err, "error creating quick transfer in blockchain")
	}

	return &models.TransferResponseObj{
		BlockNumber: blockNum,
		TxHash:      txHash,
	}, nil
}

// Public methods for different type of transfers
func (tm *Manager) QuickTransfer(ctx context.Context,
	transfer *models.QuickTransferObj) (*models.TransferResponseObj, error) {
	return tm.quickTransfer(ctx, transfer, false)
}

func (tm *Manager) QuickTransferAsync(ctx context.Context,
	transfer *models.QuickTransferObj) (*models.TransferResponseObj, error) {
	return tm.quickTransfer(ctx, transfer, true)
}

func (tm *Manager) quickTransferBaseCheck(transfer *models.QuickTransferObj) error {
	// basic validation
	if transfer.From == (common.Address{}) {
		return bitbonErrors.NewInvalidParamsError(ErrFromRequire)
	}
	if transfer.To == (common.Address{}) {
		return bitbonErrors.NewInvalidParamsError(ErrToRequire)
	}
	if transfer.From == transfer.To {
		return bitbonErrors.NewInvalidParamsError(ErrSameAccount)
	}
	if transfer.AccountID == "" {
		return bitbonErrors.NewInvalidParamsError(ErrAccountIDRequire)
	}
	if transfer.Value == nil || transfer.Value.Cmp(big.NewInt(0)) <= 0 {
		return bitbonErrors.NewInvalidParamsError(ErrValueRequire)
	}

	return nil
}
