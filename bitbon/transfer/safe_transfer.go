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

// nolint:dupl,nakedret
package transfer

import (
	"context"
	"math/big"
	"time"

	bb "github.com/simcord-llc/bitbon-system-blockchain/bitbon"

	"github.com/pkg/errors"
	bitbonErrors "github.com/simcord-llc/bitbon-system-blockchain/bitbon/errors"
	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/models"
	"github.com/simcord-llc/bitbon-system-blockchain/common"
)

func (tm *Manager) CreateSafeTransfer(ctx context.Context,
	transfer *models.CreateTransferObj) (*models.TransferResponseObj, error) {
	return tm.createSafeTransfer(ctx, transfer, false)
}

func (tm *Manager) CreateSafeTransferAsync(ctx context.Context,
	transfer *models.CreateTransferObj) (*models.TransferResponseObj, error) {
	return tm.createSafeTransfer(ctx, transfer, true)
}

func (tm *Manager) ApproveSafeTransfer(ctx context.Context,
	transfer *models.ApproveTransferObj) (*models.TransferResponseObj, error) {
	return tm.approveSafeTransfer(ctx, transfer, false)
}

func (tm *Manager) ApproveSafeTransferAsync(ctx context.Context,
	transfer *models.ApproveTransferObj) (*models.TransferResponseObj, error) {
	return tm.approveSafeTransfer(ctx, transfer, true)
}

func (tm *Manager) CancelSafeTransfer(ctx context.Context,
	transfer *models.CancelTransferObj) (*models.TransferResponseObj, error) {
	return tm.cancelSafeTransfer(ctx, transfer, false)
}

func (tm *Manager) CancelSafeTransferAsync(ctx context.Context,
	transfer *models.CancelTransferObj) (*models.TransferResponseObj, error) {
	return tm.cancelSafeTransfer(ctx, transfer, true)
}

func (tm *Manager) CreateWPCSafeTransfer(ctx context.Context,
	transfer *models.CreateTransferObj) (*models.TransferResponseObj, error) {
	return tm.createWPCSafeTransfer(ctx, transfer, false)
}

func (tm *Manager) CreateWPCSafeTransferAsync(ctx context.Context,
	transfer *models.CreateTransferObj) (*models.TransferResponseObj, error) {
	return tm.createWPCSafeTransfer(ctx, transfer, true)
}

func (tm *Manager) ApproveWPCSafeTransfer(ctx context.Context,
	transfer *models.ApproveTransferObj) (*models.TransferResponseObj, error) {
	return tm.approveWPCSafeTransfer(ctx, transfer, false)
}

func (tm *Manager) ApproveWPCSafeTransferAsync(ctx context.Context,
	transfer *models.ApproveTransferObj) (*models.TransferResponseObj, error) {
	return tm.approveWPCSafeTransfer(ctx, transfer, true)
}

func (tm *Manager) CancelWPCSafeTransfer(ctx context.Context,
	transfer *models.CancelTransferObj) (*models.TransferResponseObj, error) {
	return tm.cancelWPCSafeTransfer(ctx, transfer, false)
}

func (tm *Manager) CancelWPCSafeTransferAsync(ctx context.Context,
	transfer *models.CancelTransferObj) (*models.TransferResponseObj, error) {
	return tm.cancelWPCSafeTransfer(ctx, transfer, true)
}

func (tm *Manager) createSafeTransfer(ctx context.Context,
	transfer *models.CreateTransferObj, async bool) (*models.TransferResponseObj, error) {
	assetbox := &models.Assetbox{
		Wallet:     transfer.CryptoData.Wallet,
		PassPhrase: transfer.CryptoData.Passphrase,
	}

	if !tm.Ready() {
		return nil, ErrTransferNotReady
	}

	tm.logger.Debug("TransferManager createSafeTransfer called",
		"from", transfer.From.Hex(), "to", transfer.To.Hex(),
		"value", transfer.Value.String(), "transferID", transfer.TransferID)

	if err := tm.createSafeTransferBaseCheck(transfer); err != nil {
		return nil, err
	}

	if err := bb.DecryptAssetboxWallet(assetbox, tm.bitbon.GetDecryptAssetboxWalletPassword(), tm.encryptor.Decrypt); err != nil {
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

	// check TransferID for existence
	transferExists, err := tm.TransferExists(ctx, transfer.TransferID)
	if err != nil {
		return nil, errors.Wrap(err, "error checking if transfer with current TransferID exists")
	}
	if transferExists {
		return nil, bitbonErrors.NewTransferExistsError(
			errors.Errorf("transfer with id %q already exists", transfer.TransferID))
	}
	tm.logger.Debug("TransferID was checked for existence (not exists)")

	// generate zk-snark params
	err = tm.generateZkSnark(transfer)
	if err != nil {
		return nil, err
	}

	contrTransfer := transfer.ToContractsTransfer()

	blockNum, txHash, err := tm.bitbon.GetContractsManager().
		CreateSafeTransfer(ctx, contrTransfer, assetbox.Pk, async)

	tm.logger.Debug("Contract manager createSafeTransfer results",
		"blockNum", blockNum, "txHash", txHash.Hex(), "error", err)
	if err != nil {
		return nil, errors.Wrap(err, "error creating safe transfer in blockchain")
	}

	return &models.TransferResponseObj{
		BlockNumber: blockNum,
		TxHash:      txHash,
	}, nil
}

func (tm *Manager) approveSafeTransfer(ctx context.Context,
	transfer *models.ApproveTransferObj, async bool) (*models.TransferResponseObj, error) {
	assetbox := &models.Assetbox{
		Wallet:     transfer.CryptoData.Wallet,
		PassPhrase: transfer.CryptoData.Passphrase,
	}

	if !tm.Ready() {
		return nil, ErrTransferNotReady
	}

	tm.logger.Debug("TransferManager approveSafeTransfer called",
		"address", transfer.Address.Hex(), "transferID", transfer.TransferID)

	// basic validation
	if transfer.Address == (common.Address{}) {
		return nil, bitbonErrors.NewInvalidParamsError(ErrAddressRequire)
	}
	if transfer.TransferID == "" {
		return nil, bitbonErrors.NewInvalidParamsError(ErrTransferIDRequire)
	}
	if transfer.ProtectionCode == "" {
		return nil, bitbonErrors.NewInvalidParamsError(ErrProtectionCodeRequire)
	}

	if err := bb.DecryptAssetboxWallet(assetbox, tm.bitbon.GetDecryptAssetboxWalletPassword(), tm.encryptor.Decrypt); err != nil {
		return nil, err
	}

	// check TransferID for existence
	transferExists, err := tm.TransferExists(ctx, transfer.TransferID)
	if err != nil {
		return nil, errors.Wrap(err, "error checking if transfer with current TransferID exists")
	}
	if !transferExists {
		return nil, bitbonErrors.NewTransferIsNotExistError(
			errors.Errorf("transfer with id %q does not exist", transfer.TransferID))
	}

	tm.logger.Debug("TransferID was checked for existence (exists)")

	// generate protectionHash
	protectionHash := tm.getProtectionHash(transfer.TransferID, transfer.ProtectionCode)
	tm.logger.Debug("ProtectionHash successfully generated")

	// call smart contract method
	blockNum, txHash, err := tm.bitbon.GetContractsManager().ApproveSafeTransfer(ctx,
		[]byte(transfer.TransferID), protectionHash, transfer.ExtraData, assetbox.Pk, async)
	if err != nil {
		return nil, errors.Wrap(err, "error approving safe transfer in blockchain")
	}
	tm.logger.Debug("Contract manager approveSafeTransfer results", "blockNum", blockNum, "txHash", txHash.Hex())

	return &models.TransferResponseObj{
		BlockNumber: blockNum,
		TxHash:      txHash,
	}, nil
}

func (tm *Manager) cancelSafeTransfer(ctx context.Context,
	transfer *models.CancelTransferObj, async bool) (*models.TransferResponseObj, error) {
	assetbox := &models.Assetbox{
		Wallet:     transfer.CryptoData.Wallet,
		PassPhrase: transfer.CryptoData.Passphrase,
	}

	tm.logger.Debug("TransferManager cancelSafeTransfer called",
		"address", transfer.Address.Hex(), "transferID", transfer.TransferID)

	// basic validation
	if transfer.Address == (common.Address{}) {
		return nil, bitbonErrors.NewInvalidParamsError(ErrAddressRequire)
	}
	if transfer.TransferID == "" {
		return nil, bitbonErrors.NewInvalidParamsError(ErrTransferIDRequire)
	}

	if err := bb.DecryptAssetboxWallet(assetbox, tm.bitbon.GetDecryptAssetboxWalletPassword(), tm.encryptor.Decrypt); err != nil {
		return nil, err
	}

	// check TransferID for existence
	transferExists, err := tm.TransferExists(ctx, transfer.TransferID)
	if err != nil {
		return nil, errors.Wrap(err, "error checking if transfer with current TransferID exists")
	}
	if !transferExists {
		return nil, bitbonErrors.NewTransferIsNotExistError(
			errors.Errorf("transfer with id %q does not exist", transfer.TransferID))
	}

	tm.logger.Debug("TransferID was checked for existence (exists)")

	// call smart contract method
	blockNum, txHash, err := tm.bitbon.GetContractsManager().CancelSafeTransfer(ctx,
		[]byte(transfer.TransferID), transfer.ExtraData, assetbox.Pk, async)
	if err != nil {
		return nil, errors.Wrap(err, "error canceling safe transfer in blockchain")
	}
	tm.logger.Debug("Contract manager cancelSafeTransfer results", "blockNum", blockNum, "txHash", txHash.Hex())

	return &models.TransferResponseObj{
		BlockNumber: blockNum,
		TxHash:      txHash,
	}, nil
}

func (tm *Manager) createWPCSafeTransfer(ctx context.Context,
	transfer *models.CreateTransferObj, async bool) (*models.TransferResponseObj, error) {
	assetbox := &models.Assetbox{
		Wallet:     transfer.CryptoData.Wallet,
		PassPhrase: transfer.CryptoData.Passphrase,
	}

	if !tm.Ready() {
		return nil, ErrTransferNotReady
	}

	tm.logger.Debug("TransferManager createWPCSafeTransfer called",
		"from", transfer.From.Hex(), "to", transfer.To.Hex(), "value",
		transfer.Value.String(), "transferID", transfer.TransferID)

	if err := tm.createWPCSafeTransferBaseCheck(transfer); err != nil {
		return nil, err
	}

	if err := bb.DecryptAssetboxWallet(assetbox, tm.bitbon.GetDecryptAssetboxWalletPassword(), tm.encryptor.Decrypt); err != nil {
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

	// check TransferID for existence
	transferExists, err := tm.bitbon.GetContractsManager().TransferExists(ctx, []byte(transfer.TransferID))
	if err != nil {
		return nil, errors.Wrap(err, "error checking if transfer with current TransferID exists")
	}
	if transferExists {
		return nil, bitbonErrors.NewTransferExistsError(
			errors.Errorf("transfer with id %q already exists", transfer.TransferID))
	}
	tm.logger.Debug("TransferID was checked for existence (does not exist)")

	contrTransfer := transfer.ToContractsTransfer()

	blockNum, txHash, err := tm.bitbon.GetContractsManager().
		CreateWPCSafeTransfer(ctx, contrTransfer, assetbox.Pk, async)

	tm.logger.Debug("Contract manager createWpcSafeTransfer results",
		"blockNum", blockNum, "txHash", txHash.Hex(), "error", err)
	if err != nil {
		return nil, errors.Wrap(err, "error creating WPC safe transfer in blockchain")
	}

	return &models.TransferResponseObj{
		BlockNumber: blockNum,
		TxHash:      txHash,
	}, nil
}

func (tm *Manager) approveWPCSafeTransfer(ctx context.Context,
	transfer *models.ApproveTransferObj, async bool) (*models.TransferResponseObj, error) {
	assetbox := &models.Assetbox{
		Wallet:     transfer.CryptoData.Wallet,
		PassPhrase: transfer.CryptoData.Passphrase,
	}

	if !tm.Ready() {
		return nil, ErrTransferNotReady
	}

	tm.logger.Debug("TransferManager approveWPCSafeTransfer called",
		"address", transfer.Address.Hex(), "transferID", transfer.TransferID)

	// basic validation
	if transfer.Address == (common.Address{}) {
		return nil, bitbonErrors.NewInvalidParamsError(ErrAddressRequire)
	}
	if transfer.TransferID == "" {
		return nil, bitbonErrors.NewInvalidParamsError(ErrTransferIDRequire)
	}

	if err := bb.DecryptAssetboxWallet(assetbox, tm.bitbon.GetDecryptAssetboxWalletPassword(), tm.encryptor.Decrypt); err != nil {
		return nil, err
	}

	// check TransferID for existence
	transferExists, err := tm.TransferExists(ctx, transfer.TransferID)
	if err != nil {
		return nil, errors.Wrap(err, "error checking if transfer with current TransferID exists")
	}
	if !transferExists {
		return nil, bitbonErrors.NewTransferIsNotExistError(
			errors.Errorf("transfer with id %q does not exist", transfer.TransferID))
	}

	tm.logger.Debug("TransferID was checked for existence (exists)")

	// call smart contract method
	blockNum, txHash, err := tm.bitbon.GetContractsManager().ApproveWPCSafeTransfer(ctx,
		[]byte(transfer.TransferID), transfer.ExtraData, assetbox.Pk, async)
	if err != nil {
		return nil, errors.Wrap(err, "error approving WPC safe transfer in blockchain")
	}
	tm.logger.Debug("Contract manager approveWPCSafeTransfer results", "blockNum", blockNum, "txHash", txHash.Hex())

	return &models.TransferResponseObj{
		BlockNumber: blockNum,
		TxHash:      txHash,
	}, nil
}

func (tm *Manager) cancelWPCSafeTransfer(ctx context.Context,
	transfer *models.CancelTransferObj, async bool) (*models.TransferResponseObj, error) {
	assetbox := &models.Assetbox{
		Wallet:     transfer.CryptoData.Wallet,
		PassPhrase: transfer.CryptoData.Passphrase,
	}

	tm.logger.Debug("TransferManager cancelWPCSafeTransfer called",
		"address", transfer.Address.Hex(), "transferID", transfer.TransferID)

	// basic validation
	if transfer.Address == (common.Address{}) {
		return nil, bitbonErrors.NewInvalidParamsError(ErrAddressRequire)
	}
	if transfer.TransferID == "" {
		return nil, bitbonErrors.NewInvalidParamsError(ErrTransferIDRequire)
	}

	if err := bb.DecryptAssetboxWallet(assetbox, tm.bitbon.GetDecryptAssetboxWalletPassword(), tm.encryptor.Decrypt); err != nil {
		return nil, err
	}

	// check TransferID for existence
	transferExists, err := tm.TransferExists(ctx, transfer.TransferID)
	if err != nil {
		return nil, errors.Wrap(err, "error checking if transfer with current TransferID exists")
	}
	if !transferExists {
		return nil, bitbonErrors.NewTransferIsNotExistError(
			errors.Errorf("transfer with id %q does not exist", transfer.TransferID))
	}

	tm.logger.Debug("TransferID was checked for existence (exists)")

	// call smart contract method
	blockNum, txHash, err := tm.bitbon.GetContractsManager().CancelWPCSafeTransfer(ctx,
		[]byte(transfer.TransferID), transfer.ExtraData, assetbox.Pk, async)
	if err != nil {
		return nil, errors.Wrap(err, "error canceling WPC safe transfer in blockchain")
	}
	tm.logger.Debug("Contract manager cancelWPCSafeTransfer results", "blockNum", blockNum, "txHash", txHash.Hex())

	return &models.TransferResponseObj{
		BlockNumber: blockNum,
		TxHash:      txHash,
	}, nil
}

func (tm *Manager) createSafeTransferBaseCheck(transfer *models.CreateTransferObj) error {
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
	if transfer.TransferID == "" {
		return bitbonErrors.NewInvalidParamsError(ErrTransferIDRequire)
	}
	if transfer.ProtectionCode == "" {
		return bitbonErrors.NewInvalidParamsError(ErrProtectionCodeRequire)
	}
	if transfer.ExpiresAt <= time.Now().Unix() {
		return bitbonErrors.NewInvalidParamsError(ErrExpireAtPassed)
	}
	if transfer.Value == nil || transfer.Value.Cmp(big.NewInt(0)) <= 0 {
		return bitbonErrors.NewInvalidParamsError(ErrValueRequire)
	}

	return nil
}

func (tm *Manager) createWPCSafeTransferBaseCheck(transfer *models.CreateTransferObj) error {
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
	if transfer.TransferID == "" {
		return bitbonErrors.NewInvalidParamsError(ErrTransferIDRequire)
	}
	if transfer.ExpiresAt <= time.Now().Unix() {
		return bitbonErrors.NewInvalidParamsError(ErrExpireAtPassed)
	}
	if transfer.Value == nil || transfer.Value.Cmp(big.NewInt(0)) <= 0 {
		return bitbonErrors.NewInvalidParamsError(ErrValueRequire)
	}

	return nil
}
