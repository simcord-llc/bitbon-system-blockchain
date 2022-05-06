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
	"github.com/simcord-llc/bitbon-system-blockchain/zksnark"
	"math/big"
	"time"

	bb "github.com/simcord-llc/bitbon-system-blockchain/bitbon"

	"github.com/pkg/errors"
	bitbonErrors "github.com/simcord-llc/bitbon-system-blockchain/bitbon/errors"
	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/models"
	"github.com/simcord-llc/bitbon-system-blockchain/common"
)

const (
	BitbonTransferStateUndefined = iota
	BitbonTransferStatePending
	BitbonTransferStateExecuted
	BitbonTransferStateExpired
	BitbonTransferStateCancelled
	BitbonTransferStateFailed
)

func (tm *Manager) CreateSafeTransfer(ctx context.Context,
	transfer *models.CreateTransferObj) (*models.TransferResponseObj, error) {
	return tm.createSafeTransfer(ctx, transfer, false)
}

func (tm *Manager) CreateSafeTransferAsync(ctx context.Context,
	transfer *models.CreateTransferObj) (*models.TransferResponseObj, error) {
	return tm.createSafeTransfer(ctx, transfer, true)
}

func (tm *Manager) CreateFullBalanceSafeTransfer(ctx context.Context,
	transfer *models.CreateTransferObj) (*models.TransferResponseObj, error) {
	return tm.createFullBalanceSafeTransfer(ctx, transfer, false)
}

func (tm *Manager) CreateFullBalanceSafeTransferAsync(ctx context.Context,
	transfer *models.CreateTransferObj) (*models.TransferResponseObj, error) {
	return tm.createFullBalanceSafeTransfer(ctx, transfer, true)
}

func (tm *Manager) ApproveSafeTransfer(ctx context.Context,
	transfer *models.ApproveTransferObj) (*models.TransferResponseObj, error) {
	return tm.approveSafeTransfer(ctx, transfer, false)
}

func (tm *Manager) ApproveSafeTransferAsync(ctx context.Context,
	transfer *models.ApproveTransferObj) (*models.TransferResponseObj, error) {
	return tm.approveSafeTransfer(ctx, transfer, true)
}

func (tm *Manager) ApproveFullBalanceSafeTransfer(ctx context.Context,
	transfer *models.ApproveTransferObj) (*models.TransferResponseObj, error) {
	return tm.approveFullBalanceSafeTransfer(ctx, transfer, false)
}

func (tm *Manager) ApproveFullBalanceSafeTransferAsync(ctx context.Context,
	transfer *models.ApproveTransferObj) (*models.TransferResponseObj, error) {
	return tm.approveFullBalanceSafeTransfer(ctx, transfer, true)
}

func (tm *Manager) CancelSafeTransfer(ctx context.Context,
	transfer *models.CancelTransferObj) (*models.TransferResponseObj, error) {
	return tm.cancelSafeTransfer(ctx, transfer, false)
}

func (tm *Manager) CancelSafeTransferAsync(ctx context.Context,
	transfer *models.CancelTransferObj) (*models.TransferResponseObj, error) {
	return tm.cancelSafeTransfer(ctx, transfer, true)
}

func (tm *Manager) CancelFullBalanceSafeTransfer(ctx context.Context,
	transfer *models.CancelTransferObj) (*models.TransferResponseObj, error) {
	return tm.cancelFullBalanceSafeTransfer(ctx, transfer, false)
}

func (tm *Manager) CancelFullBalanceSafeTransferAsync(ctx context.Context,
	transfer *models.CancelTransferObj) (*models.TransferResponseObj, error) {
	return tm.cancelFullBalanceSafeTransfer(ctx, transfer, true)
}

func (tm *Manager) CreateWPCSafeTransfer(ctx context.Context,
	transfer *models.CreateTransferObj) (*models.TransferResponseObj, error) {
	return tm.createWPCSafeTransfer(ctx, transfer, false)
}

func (tm *Manager) CreateWPCSafeTransferAsync(ctx context.Context,
	transfer *models.CreateTransferObj) (*models.TransferResponseObj, error) {
	return tm.createWPCSafeTransfer(ctx, transfer, true)
}

func (tm *Manager) CreateFullBalanceWPCSafeTransfer(ctx context.Context,
	transfer *models.CreateTransferObj) (*models.TransferResponseObj, error) {
	return tm.createFullBalanceWPCSafeTransfer(ctx, transfer, false)
}

func (tm *Manager) CreateFullBalanceWPCSafeTransferAsync(ctx context.Context,
	transfer *models.CreateTransferObj) (*models.TransferResponseObj, error) {
	return tm.createFullBalanceWPCSafeTransfer(ctx, transfer, true)
}

func (tm *Manager) ApproveWPCSafeTransfer(ctx context.Context,
	transfer *models.ApproveTransferObj) (*models.TransferResponseObj, error) {
	return tm.approveWPCSafeTransfer(ctx, transfer, false)
}

func (tm *Manager) ApproveWPCSafeTransferAsync(ctx context.Context,
	transfer *models.ApproveTransferObj) (*models.TransferResponseObj, error) {
	return tm.approveWPCSafeTransfer(ctx, transfer, true)
}

func (tm *Manager) ApproveFullBalanceWPCSafeTransfer(ctx context.Context,
	transfer *models.ApproveTransferObj) (*models.TransferResponseObj, error) {
	return tm.approveFullBalanceWPCSafeTransfer(ctx, transfer, false)
}

func (tm *Manager) ApproveFullBalanceWPCSafeTransferAsync(ctx context.Context,
	transfer *models.ApproveTransferObj) (*models.TransferResponseObj, error) {
	return tm.approveFullBalanceWPCSafeTransfer(ctx, transfer, true)
}

func (tm *Manager) CancelWPCSafeTransfer(ctx context.Context,
	transfer *models.CancelTransferObj) (*models.TransferResponseObj, error) {
	return tm.cancelWPCSafeTransfer(ctx, transfer, false)
}

func (tm *Manager) CancelWPCSafeTransferAsync(ctx context.Context,
	transfer *models.CancelTransferObj) (*models.TransferResponseObj, error) {
	return tm.cancelWPCSafeTransfer(ctx, transfer, true)
}

func (tm *Manager) CancelFullBalanceWPCSafeTransfer(ctx context.Context,
	transfer *models.CancelTransferObj) (*models.TransferResponseObj, error) {
	return tm.cancelFullBalanceWPCSafeTransfer(ctx, transfer, false)
}

func (tm *Manager) CancelFullBalanceWPCSafeTransferAsync(ctx context.Context,
	transfer *models.CancelTransferObj) (*models.TransferResponseObj, error) {
	return tm.cancelFullBalanceWPCSafeTransfer(ctx, transfer, true)
}

func (tm *Manager) createSafeTransfer(ctx context.Context,
	transfer *models.CreateTransferObj, async bool) (*models.TransferResponseObj, error) {
	if !tm.Ready() {
		return nil, ErrTransferNotReady
	}

	tm.logger.Debug("TransferManager createSafeTransfer called",
		"from", transfer.From.Hex(), "to", transfer.To.Hex(),
		"value", transfer.Value.String(), "transferID", transfer.TransferID)

	if err := tm.createSafeTransferBaseCheck(transfer); err != nil {
		return nil, err
	}

	privateKey, err := bb.DecryptPrivateKeyForAssetbox(transfer.From, transfer.CryptoData.Wallet, transfer.CryptoData.Passphrase, tm.bitbon.GetDecryptAssetboxWalletPassword(), tm.encryptor.Decrypt)
	if err != nil {
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
		CreateSafeTransfer(ctx, contrTransfer, privateKey, async)

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

	privateKey, err := bb.DecryptPrivateKeyForAssetbox(transfer.Address, transfer.CryptoData.Wallet, transfer.CryptoData.Passphrase, tm.bitbon.GetDecryptAssetboxWalletPassword(), tm.encryptor.Decrypt)
	if err != nil {
		return nil, err
	}

	// check that transfer exists and is pending
	transferState, err := tm.GetTransferState(ctx, transfer.TransferID)
	if err != nil {
		return nil, errors.Wrap(err, "error getting transfer state for current TransferID")
	}
	if transferState != BitbonTransferStatePending {
		return nil, bitbonErrors.NewTransferIsNotExistError(
			errors.Errorf("transfer is not pending for approval"))
	}
	tm.logger.Debug("TransferID was checked for pending (is pending)")

	// generate protectionHash
	protectionHash := tm.getProtectionHash(transfer.TransferID, transfer.ProtectionCode)
	tm.logger.Debug("ProtectionHash successfully generated")

	// call smart contract method
	blockNum, txHash, err := tm.bitbon.GetContractsManager().ApproveSafeTransfer(ctx,
		[]byte(transfer.TransferID), protectionHash, transfer.ExtraData, privateKey, async)
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
	tm.logger.Debug("TransferManager cancelSafeTransfer called",
		"address", transfer.Address.Hex(), "transferID", transfer.TransferID)

	// basic validation
	if transfer.Address == (common.Address{}) {
		return nil, bitbonErrors.NewInvalidParamsError(ErrAddressRequire)
	}
	if transfer.TransferID == "" {
		return nil, bitbonErrors.NewInvalidParamsError(ErrTransferIDRequire)
	}

	privateKey, err := bb.DecryptPrivateKeyForAssetbox(transfer.Address, transfer.CryptoData.Wallet, transfer.CryptoData.Passphrase, tm.bitbon.GetDecryptAssetboxWalletPassword(), tm.encryptor.Decrypt)
	if err != nil {
		return nil, err
	}

	// check that transfer exists and is pending
	transferState, err := tm.GetTransferState(ctx, transfer.TransferID)
	if err != nil {
		return nil, errors.Wrap(err, "error getting transfer state for current TransferID")
	}
	if transferState != BitbonTransferStatePending {
		return nil, bitbonErrors.NewTransferIsNotExistError(
			errors.Errorf("transfer is not pending for approval"))
	}
	tm.logger.Debug("TransferID was checked for pending (is pending)")

	// call smart contract method
	blockNum, txHash, err := tm.bitbon.GetContractsManager().CancelSafeTransfer(ctx,
		[]byte(transfer.TransferID), transfer.ExtraData, privateKey, async)
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
	if !tm.Ready() {
		return nil, ErrTransferNotReady
	}

	tm.logger.Debug("TransferManager createWPCSafeTransfer called",
		"from", transfer.From.Hex(), "to", transfer.To.Hex(), "value",
		transfer.Value.String(), "transferID", transfer.TransferID)

	if err := tm.createWPCSafeTransferBaseCheck(transfer); err != nil {
		return nil, err
	}

	privateKey, err := bb.DecryptPrivateKeyForAssetbox(transfer.From, transfer.CryptoData.Wallet, transfer.CryptoData.Passphrase, tm.bitbon.GetDecryptAssetboxWalletPassword(), tm.encryptor.Decrypt)
	if err != nil {
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
		CreateWPCSafeTransfer(ctx, contrTransfer, privateKey, async)

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

	privateKey, err := bb.DecryptPrivateKeyForAssetbox(transfer.Address, transfer.CryptoData.Wallet, transfer.CryptoData.Passphrase, tm.bitbon.GetDecryptAssetboxWalletPassword(), tm.encryptor.Decrypt)
	if err != nil {
		return nil, err
	}

	// check that transfer exists and pending
	transferState, err := tm.GetTransferState(ctx, transfer.TransferID)
	if err != nil {
		return nil, errors.Wrap(err, "error getting transfer state for current TransferID")
	}
	if transferState != BitbonTransferStatePending {
		return nil, bitbonErrors.NewTransferIsNotExistError(
			errors.Errorf("transfer is not pending for approval"))
	}

	tm.logger.Debug("TransferID was checked for pending (is pending)")

	// call smart contract method
	blockNum, txHash, err := tm.bitbon.GetContractsManager().ApproveWPCSafeTransfer(ctx,
		[]byte(transfer.TransferID), transfer.ExtraData, privateKey, async)
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
	tm.logger.Debug("TransferManager cancelWPCSafeTransfer called",
		"address", transfer.Address.Hex(), "transferID", transfer.TransferID)

	// basic validation
	if transfer.Address == (common.Address{}) {
		return nil, bitbonErrors.NewInvalidParamsError(ErrAddressRequire)
	}
	if transfer.TransferID == "" {
		return nil, bitbonErrors.NewInvalidParamsError(ErrTransferIDRequire)
	}

	privateKey, err := bb.DecryptPrivateKeyForAssetbox(transfer.Address, transfer.CryptoData.Wallet, transfer.CryptoData.Passphrase, tm.bitbon.GetDecryptAssetboxWalletPassword(), tm.encryptor.Decrypt)
	if err != nil {
		return nil, err
	}

	// check that transfer exists and is pending
	transferState, err := tm.GetTransferState(ctx, transfer.TransferID)
	if err != nil {
		return nil, errors.Wrap(err, "error getting transfer state for current TransferID")
	}
	if transferState != BitbonTransferStatePending {
		return nil, bitbonErrors.NewTransferIsNotExistError(
			errors.Errorf("transfer is not pending for approval"))
	}
	tm.logger.Debug("TransferID was checked for pending (is pending)")

	// call smart contract method
	blockNum, txHash, err := tm.bitbon.GetContractsManager().CancelWPCSafeTransfer(ctx,
		[]byte(transfer.TransferID), transfer.ExtraData, privateKey, async)
	if err != nil {
		return nil, errors.Wrap(err, "error canceling WPC safe transfer in blockchain")
	}
	tm.logger.Debug("Contract manager cancelWPCSafeTransfer results", "blockNum", blockNum, "txHash", txHash.Hex())

	return &models.TransferResponseObj{
		BlockNumber: blockNum,
		TxHash:      txHash,
	}, nil
}

func (tm *Manager) createFullBalanceSafeTransfer(ctx context.Context, transfer *models.CreateTransferObj, async bool) (*models.TransferResponseObj, error) {
	if !tm.Ready() {
		return nil, ErrTransferNotReady
	}
	tm.logger.Warn("TransferManager createFullBalanceSafeTransfer called", "from", transfer.From.Hex(), "to", transfer.To.Hex(), "value", transfer.Value.String(), "transferID", transfer.TransferID)

	// basic validation
	if err := tm.createFullBalanceSafeTransferBaseCheck(transfer); err != nil {
		return nil, err
	}

	// getting decrypted assetbox
	privateKey, err := bb.DecryptPrivateKeyForAssetbox(transfer.From, transfer.CryptoData.Wallet, transfer.CryptoData.Passphrase, tm.bitbon.GetDecryptAssetboxWalletPassword(), tm.encryptor.Decrypt)
	if err != nil {
		return nil, err
	}

	tm.logger.Debug("BTSC successfully validated")

	// check locked and available balance
	currentBalance, err := tm.bitbon.GetContractsManager().GetAssetboxBalance(ctx, transfer.From)
	if err != nil {
		return nil, errors.Wrap(err, "error getting current assetbox balance")
	}

	if currentBalance.Cmp(big.NewInt(0)) <= 0 {
		return nil, bitbonErrors.NewBalanceTooLowError(errors.New("assetbox balance to low to perform transfer"))
	}

	lockedBalance, err := tm.bitbon.GetContractsManager().BalanceOfLocked(ctx, transfer.From)
	if err != nil {
		return nil, errors.Wrap(err, "error getting locked assetbox balance")
	}

	if lockedBalance.Cmp(big.NewInt(0)) > 0 {
		return nil, bitbonErrors.HasLockedBalanceError(errors.New("assetbox has locked balance"))
	}

	// check TransferID for existence
	transferExists, err := tm.TransferExists(ctx, transfer.TransferID)
	if err != nil {
		return nil, errors.Wrap(err, "error checking if transfer with current TransferID exists")
	}
	if transferExists {
		return nil, bitbonErrors.NewTransferExistsError(errors.Errorf("transfer with id %q already exists", transfer.TransferID))
	}
	tm.logger.Debug("TransferID was checked for existence (not exists)")

	// generate zk-snark params
	protectionHash := tm.getProtectionHash(transfer.TransferID, transfer.ProtectionCode)
	tm.logger.Debug("protectionHash successfully obtained", "protectionHash", protectionHash)
	var pk []byte
	pk, transfer.VK = tm.getKeys()
	tm.logger.Debug("zk-snrak keys obtaioned (pk, vk)", "len(pk)", len(pk), "len(transfer.VK)", len(transfer.VK))
	transfer.Proof, err = zksnark.Prove(protectionHash, pk, constraintsLen)
	if err != nil {
		return nil, errors.Wrap(err, "error generating ZK-snark params")
	}
	tm.logger.Debug("ZK-snark params successfully generated")

	contrTransfer := transfer.ToContractsTransfer()

	blockNum, txHash, err := tm.bitbon.GetContractsManager().CreateFullBalanceSafeTransfer(ctx, contrTransfer, privateKey, async)

	tm.logger.Info("Contract manager CreateFullBalanceSafeTransfer results", "blockNum", blockNum, "txHash", txHash.Hex(), "error", err)
	if err != nil {
		return nil, errors.Wrap(err, "error creating safe transfer in blockchain")
	}

	return &models.TransferResponseObj{
		BlockNumber: blockNum,
		TxHash:      txHash,
	}, nil
}

func (tm *Manager) approveFullBalanceSafeTransfer(ctx context.Context,
	transfer *models.ApproveTransferObj, async bool) (*models.TransferResponseObj, error) {
	if !tm.Ready() {
		return nil, ErrTransferNotReady
	}

	tm.logger.Debug("TransferManager approveFullBalanceSafeTransfer called",
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

	privateKey, err := bb.DecryptPrivateKeyForAssetbox(transfer.Address, transfer.CryptoData.Wallet, transfer.CryptoData.Passphrase, tm.bitbon.GetDecryptAssetboxWalletPassword(), tm.encryptor.Decrypt)
	if err != nil {
		return nil, err
	}

	// check that transfer exists and is pending
	transferState, err := tm.GetTransferState(ctx, transfer.TransferID)
	if err != nil {
		return nil, errors.Wrap(err, "error getting transfer state for current TransferID")
	}
	if transferState != BitbonTransferStatePending {
		return nil, bitbonErrors.NewTransferIsNotExistError(
			errors.Errorf("transfer is not pending for approval"))
	}

	tm.logger.Debug("TransferID was checked for pending (is pending)")

	// generate protectionHash
	protectionHash := tm.getProtectionHash(transfer.TransferID, transfer.ProtectionCode)
	tm.logger.Debug("ProtectionHash successfully generated")

	// call smart contract method
	blockNum, txHash, err := tm.bitbon.GetContractsManager().ApproveFullBalanceSafeTransfer(ctx,
		[]byte(transfer.TransferID), protectionHash, transfer.ExtraData, privateKey, async)
	if err != nil {
		return nil, errors.Wrap(err, "error approving full balance safe transfer in blockchain")
	}
	tm.logger.Debug("Contract manager approveFullBalanceSafeTransfer results", "blockNum", blockNum, "txHash", txHash.Hex())

	return &models.TransferResponseObj{
		BlockNumber: blockNum,
		TxHash:      txHash,
	}, nil
}

func (tm *Manager) cancelFullBalanceSafeTransfer(ctx context.Context, transfer *models.CancelTransferObj, async bool) (*models.TransferResponseObj, error) {
	tm.logger.Warn("TransferManager cancelFullBalanceSafeTransfer called", "address", transfer.Address.Hex(), "transferID", transfer.TransferID)

	// basic validation
	if transfer.Address == (common.Address{}) {
		return nil, bitbonErrors.NewInvalidParamsError(errors.New("address is required"))
	}
	if transfer.TransferID == "" {
		return nil, bitbonErrors.NewInvalidParamsError(errors.New("transferID is required"))
	}

	// getting decrypted assetbox
	privateKey, err := bb.DecryptPrivateKeyForAssetbox(transfer.Address, transfer.CryptoData.Wallet, transfer.CryptoData.Passphrase, tm.bitbon.GetDecryptAssetboxWalletPassword(), tm.encryptor.Decrypt)
	if err != nil {
		return nil, err
	}

	// check that transfer exists and is pending
	transferState, err := tm.GetTransferState(ctx, transfer.TransferID)
	if err != nil {
		return nil, errors.Wrap(err, "error getting transfer state for current TransferID")
	}
	if transferState != BitbonTransferStatePending {
		return nil, bitbonErrors.NewTransferIsNotExistError(
			errors.Errorf("transfer is not pending for approval"))
	}
	tm.logger.Debug("TransferID was checked for pending (is pending)")

	// call smart contract method
	blockNum, txHash, err := tm.bitbon.GetContractsManager().CancelFullBalanceSafeTransfer(ctx, []byte(transfer.TransferID), transfer.ExtraData, privateKey, async)
	if err != nil {
		return nil, errors.Wrap(err, "error canceling safe transfer in blockchain")
	}
	tm.logger.Info("Contract manager cancelFullBalanceSafeTransfer results", "blockNum", blockNum, "txHash", txHash.Hex())

	return &models.TransferResponseObj{
		BlockNumber: blockNum,
		TxHash:      txHash,
	}, nil
}

func (tm *Manager) createFullBalanceWPCSafeTransfer(ctx context.Context, transfer *models.CreateTransferObj, async bool) (*models.TransferResponseObj, error) {
	if !tm.Ready() {
		return nil, ErrTransferNotReady
	}
	tm.logger.Warn("TransferManager createFullBalanceWPCSafeTransfer called", "from", transfer.From.Hex(), "to", transfer.To.Hex(), "value", transfer.Value.String(), "transferID", transfer.TransferID)

	// basic validation
	if err := tm.createFullBalanceWPCSafeTransferBaseCheck(transfer); err != nil {
		return nil, err
	}

	// getting decrypted assetbox
	privateKey, err := bb.DecryptPrivateKeyForAssetbox(transfer.From, transfer.CryptoData.Wallet, transfer.CryptoData.Passphrase, tm.bitbon.GetDecryptAssetboxWalletPassword(), tm.encryptor.Decrypt)
	if err != nil {
		return nil, err
	}

	// check locked and available balance
	currentBalance, err := tm.bitbon.GetContractsManager().GetAssetboxBalance(ctx, transfer.From)
	if err != nil {
		return nil, errors.Wrap(err, "error getting current assetbox balance")
	}

	if currentBalance.Cmp(big.NewInt(0)) <= 0 {
		return nil, bitbonErrors.NewBalanceTooLowError(errors.New("assetbox balance to low to perform transfer"))
	}

	lockedBalance, err := tm.bitbon.GetContractsManager().BalanceOfLocked(ctx, transfer.From)
	if err != nil {
		return nil, errors.Wrap(err, "error getting locked assetbox balance")
	}

	if lockedBalance.Cmp(big.NewInt(0)) > 0 {
		return nil, bitbonErrors.HasLockedBalanceError(errors.New("assetbox has locked balance"))
	}

	// check TransferID for existence
	transferExists, err := tm.bitbon.GetContractsManager().TransferExists(ctx, []byte(transfer.TransferID))
	if err != nil {
		return nil, errors.Wrap(err, "error checking if transfer with current TransferID exists")
	}
	if transferExists {
		return nil, bitbonErrors.NewTransferExistsError(errors.Errorf("transfer with id %q already exists", transfer.TransferID))
	}
	tm.logger.Debug("TransferID was checked for existence (does not exist)")

	if err != nil {
		return nil, errors.Wrap(err, "error preparing safe transfer")
	}

	contrTransfer := transfer.ToContractsTransfer()

	blockNum, txHash, err := tm.bitbon.GetContractsManager().CreateFullBalanceWPCSafeTransfer(ctx, contrTransfer, privateKey, async)

	tm.logger.Info("Contract manager createWpcSafeTransfer results", "blockNum", blockNum, "txHash", txHash.Hex(), "error", err)
	if err != nil {
		return nil, errors.Wrap(err, "error creating WPC safe transfer in blockchain")
	}

	return &models.TransferResponseObj{
		BlockNumber: blockNum,
		TxHash:      txHash,
	}, nil
}

func (tm *Manager) approveFullBalanceWPCSafeTransfer(ctx context.Context, transfer *models.ApproveTransferObj, async bool) (*models.TransferResponseObj, error) {
	if !tm.Ready() {
		return nil, ErrTransferNotReady
	}

	tm.logger.Debug("TransferManager approveFullBalanceWPCSafeTransfer called",
		"address", transfer.Address.Hex(), "transferID", transfer.TransferID)

	// basic validation
	if transfer.Address == (common.Address{}) {
		return nil, bitbonErrors.NewInvalidParamsError(ErrAddressRequire)
	}
	if transfer.TransferID == "" {
		return nil, bitbonErrors.NewInvalidParamsError(ErrTransferIDRequire)
	}

	privateKey, err := bb.DecryptPrivateKeyForAssetbox(transfer.Address, transfer.CryptoData.Wallet, transfer.CryptoData.Passphrase, tm.bitbon.GetDecryptAssetboxWalletPassword(), tm.encryptor.Decrypt)
	if err != nil {
		return nil, err
	}

	// check that transfer exists and is pending
	transferState, err := tm.GetTransferState(ctx, transfer.TransferID)
	if err != nil {
		return nil, errors.Wrap(err, "error getting transfer state for current TransferID")
	}
	if transferState != BitbonTransferStatePending {
		return nil, bitbonErrors.NewTransferIsNotExistError(
			errors.Errorf("transfer is not pending for approval"))
	}
	tm.logger.Debug("TransferID was checked for pending (is pending)")

	// call smart contract method
	blockNum, txHash, err := tm.bitbon.GetContractsManager().ApproveFullBalanceWPCSafeTransfer(ctx,
		[]byte(transfer.TransferID), transfer.ExtraData, privateKey, async)
	if err != nil {
		return nil, errors.Wrap(err, "error approving full balance WPC safe transfer in blockchain")
	}
	tm.logger.Debug("Contract manager approveFullBalanceWPCSafeTransfer results", "blockNum", blockNum, "txHash", txHash.Hex())

	return &models.TransferResponseObj{
		BlockNumber: blockNum,
		TxHash:      txHash,
	}, nil
}

func (tm *Manager) cancelFullBalanceWPCSafeTransfer(ctx context.Context, transfer *models.CancelTransferObj, async bool) (*models.TransferResponseObj, error) {
	if !tm.Ready() {
		return nil, ErrTransferNotReady
	}
	tm.logger.Warn("TransferManager cancelFullBalanceWPCSafeTransfer called", "address", transfer.Address.Hex(), "transferID", transfer.TransferID)

	// basic validation
	if transfer.Address == (common.Address{}) {
		return nil, bitbonErrors.NewInvalidParamsError(errors.New("address is required"))
	}
	if transfer.TransferID == "" {
		return nil, bitbonErrors.NewInvalidParamsError(errors.New("transferID is required"))
	}

	// getting decrypted assetbox
	privateKey, err := bb.DecryptPrivateKeyForAssetbox(transfer.Address, transfer.CryptoData.Wallet, transfer.CryptoData.Passphrase, tm.bitbon.GetDecryptAssetboxWalletPassword(), tm.encryptor.Decrypt)
	if err != nil {
		return nil, err
	}

	// check that transfer exists and is pending
	transferState, err := tm.GetTransferState(ctx, transfer.TransferID)
	if err != nil {
		return nil, errors.Wrap(err, "error getting transfer state for current TransferID")
	}
	if transferState != BitbonTransferStatePending {
		return nil, bitbonErrors.NewTransferIsNotExistError(
			errors.Errorf("transfer is not pending for approval"))
	}
	tm.logger.Debug("TransferID was checked for pending (is pending)")

	// call smart contract method
	blockNum, txHash, err := tm.bitbon.GetContractsManager().CancelFullBalanceWPCSafeTransfer(ctx, []byte(transfer.TransferID), transfer.ExtraData, privateKey, async)
	if err != nil {
		return nil, errors.Wrap(err, "error canceling WPC safe transfer in blockchain")
	}
	tm.logger.Info("Contract manager cancelFullBalanceWPCSafeTransfer results", "blockNum", blockNum, "txHash", txHash.Hex())

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
	if transfer.TransferID == "" {
		return bitbonErrors.NewInvalidParamsError(ErrTransferIDRequire)
	}
	if transfer.ProtectionCode == "" {
		return bitbonErrors.NewInvalidParamsError(ErrProtectionCodeRequire)
	}
	if transfer.Retries == 0 {
		return bitbonErrors.NewInvalidParamsError(ErrRetriesRequired)
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

func (tm *Manager) createFullBalanceSafeTransferBaseCheck(transfer *models.CreateTransferObj) error {
	if transfer.From == (common.Address{}) {
		return bitbonErrors.NewInvalidParamsError(ErrFromRequire)
	}
	if transfer.To == (common.Address{}) {
		return bitbonErrors.NewInvalidParamsError(ErrToRequire)
	}
	if transfer.From == transfer.To {
		return bitbonErrors.NewInvalidParamsError(ErrSameAccount)
	}
	if transfer.TransferID == "" {
		return bitbonErrors.NewInvalidParamsError(ErrTransferIDRequire)
	}
	if transfer.ProtectionCode == "" {
		return bitbonErrors.NewInvalidParamsError(ErrProtectionCodeRequire)
	}
	if transfer.Retries == 0 {
		return bitbonErrors.NewInvalidParamsError(ErrRetriesRequired)
	}
	if transfer.ExpiresAt <= time.Now().Unix() {
		return bitbonErrors.NewInvalidParamsError(ErrExpireAtPassed)
	}

	return nil
}

func (tm *Manager) createFullBalanceWPCSafeTransferBaseCheck(transfer *models.CreateTransferObj) error {
	if transfer.From == (common.Address{}) {
		return bitbonErrors.NewInvalidParamsError(ErrFromRequire)
	}
	if transfer.To == (common.Address{}) {
		return bitbonErrors.NewInvalidParamsError(ErrToRequire)
	}
	if transfer.From == transfer.To {
		return bitbonErrors.NewInvalidParamsError(ErrSameAccount)
	}
	if transfer.TransferID == "" {
		return bitbonErrors.NewInvalidParamsError(ErrTransferIDRequire)
	}
	if transfer.ExpiresAt <= time.Now().Unix() {
		return bitbonErrors.NewInvalidParamsError(ErrExpireAtPassed)
	}

	return nil
}
