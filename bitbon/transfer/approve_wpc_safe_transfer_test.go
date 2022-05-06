package transfer

import (
	"context"
	"testing"

	"github.com/simcord-llc/bitbon-system-blockchain/accounts/keystore"

	"github.com/golang/mock/gomock"
	"github.com/simcord-llc/bitbon-system-blockchain/common"
	"github.com/stretchr/testify/assert"
)

func TestApproveWPCSafeTransfer(t *testing.T) {
	ctx := context.TODO()

	mockController := gomock.NewController(t)
	defer mockController.Finish()

	t.Run("ApproveWPCSafeTransferPositive", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareApproveWPCSafeTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(destPassphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(destWallet, nil)
		mockContractManager.EXPECT().GetTransferState(ctx, []byte(transfer.TransferID)).
			Return(uint8(1), nil)

		blockNum, txHash := createTransferResponse()

		key, err := keystore.DecryptKey(destWallet, string(destPassphrase))
		if err != nil {
			panic("Error during key decryption")
		}

		mockContractManager.EXPECT().
			ApproveWPCSafeTransfer(ctx, []byte(transfer.TransferID), nil, key.PrivateKey, false).
			Return(blockNum, txHash, nil)

		response, err := tm.ApproveWPCSafeTransfer(ctx, transfer)

		assert.NoError(t, err, "Unexpected error")
		assert.Equal(t, mapTransferResponse(blockNum, txHash), response, "Response differs from expected")
	})

	t.Run("ApproveWPCSafeTransferNegative_InvalidWallet", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareApproveWPCSafeTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(destPassphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(invalidWallet, nil)

		response, err := tm.ApproveWPCSafeTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrFailedToParseKey.Error(), "Expected InvalidWallet error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("ApproveWPCSafeTransferNegative_WrongPassphrase", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareApproveWPCSafeTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(wrongPass, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(destWallet, nil)

		response, err := tm.ApproveWPCSafeTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrWrongPassphrase.Error(), "Expected WrongPassphrase error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("ApproveWPCSafeTransferNegative_TransferHasExpired", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareApproveWPCSafeTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(destPassphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(destWallet, nil)
		mockContractManager.EXPECT().GetTransferState(ctx, []byte(transfer.TransferID)).
			Return(uint8(3), nil)

		response, err := tm.ApproveWPCSafeTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrTransferIsNotPending.Error())
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("ApproveWPCSafeTransferNegative_TransferMissing", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareApproveWPCSafeTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(destPassphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(destWallet, nil)
		mockContractManager.EXPECT().GetTransferState(ctx, []byte(transfer.TransferID)).
			Return(uint8(0), nil)

		response, err := tm.ApproveWPCSafeTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrTransferIsNotPending.Error())
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("ApproveWPCSafeTransferNegative_MissingTransferID", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareApproveWPCSafeTransferObject()
		transfer.TransferID = ""

		response, err := tm.ApproveWPCSafeTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrTransferIDRequire.Error(), "Expected TransferIDRequired error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("ApproveWPCSafeTransferNegative_MissingAddress", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareApproveWPCSafeTransferObject()
		transfer.Address = common.Address{}

		response, err := tm.ApproveWPCSafeTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrAddressRequire.Error(), "Expected AddressRequired error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("ApproveWPCSafeTransferNegative_ManagerIsNotReady", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)

		transfer := prepareApproveWPCSafeTransferObject()

		response, err := tm.ApproveWPCSafeTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrTransferNotReady.Error(), "Expected ManagerIsNotReady error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("ApproveWPCSafeTransferNegative_TransferIsApproved", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareApproveWPCSafeTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(destPassphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(destWallet, nil)
		mockContractManager.EXPECT().GetTransferState(ctx, []byte(transfer.TransferID)).
			Return(uint8(2), nil)

		response, err := tm.ApproveWPCSafeTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrTransferIsNotPending.Error())
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("ApproveWPCSafeTransferNegative_AssetboxIsNotOwnerOfWallet", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareApproveWPCSafeTransferObject()
		transfer.Address = sourceAssetbox

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(destPassphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(destWallet, nil)

		response, err := tm.ApproveWPCSafeTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrAssetboxIsNotOwnerOfWallet.Error(), "Expected AssetboxIsNotOwnerOfWallet error")
		assert.Empty(t, response, "Response wasn't empty")
	})
}

func TestApproveFullBalanceWPCSafeTransfer(t *testing.T) {
	ctx := context.TODO()

	mockController := gomock.NewController(t)
	defer mockController.Finish()

	t.Run("ApproveFullBalanceWPCSafeTransferPositive", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareApproveWPCSafeTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(destPassphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(destWallet, nil)
		mockContractManager.EXPECT().GetTransferState(ctx, []byte(transfer.TransferID)).
			Return(uint8(1), nil)

		blockNum, txHash := createTransferResponse()

		key, err := keystore.DecryptKey(destWallet, string(destPassphrase))
		if err != nil {
			panic("Error during key decryption")
		}

		mockContractManager.EXPECT().
			ApproveFullBalanceWPCSafeTransfer(ctx, []byte(transfer.TransferID), nil, key.PrivateKey, false).
			Return(blockNum, txHash, nil)

		response, err := tm.ApproveFullBalanceWPCSafeTransfer(ctx, transfer)

		assert.NoError(t, err, "Unexpected error")
		assert.Equal(t, mapTransferResponse(blockNum, txHash), response, "Response differs from expected")
	})

	t.Run("ApproveFullBalanceWPCSafeTransferNegative_InvalidWallet", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareApproveWPCSafeTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(destPassphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(invalidWallet, nil)

		response, err := tm.ApproveFullBalanceWPCSafeTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrFailedToParseKey.Error(), "Expected InvalidWallet error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("ApproveFullBalanceWPCSafeTransferNegative_WrongPassphrase", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareApproveWPCSafeTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(wrongPass, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(destWallet, nil)

		response, err := tm.ApproveFullBalanceWPCSafeTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrWrongPassphrase.Error(), "Expected WrongPassphrase error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("ApproveFullBalanceWPCSafeTransferNegative_TransferHasExpired", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareApproveWPCSafeTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(destPassphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(destWallet, nil)
		mockContractManager.EXPECT().GetTransferState(ctx, []byte(transfer.TransferID)).
			Return(uint8(3), nil)

		response, err := tm.ApproveFullBalanceWPCSafeTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrTransferIsNotPending.Error())
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("ApproveFullBalanceWPCSafeTransferNegative_TransferMissing", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareApproveWPCSafeTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(destPassphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(destWallet, nil)
		mockContractManager.EXPECT().GetTransferState(ctx, []byte(transfer.TransferID)).
			Return(uint8(0), nil)

		response, err := tm.ApproveFullBalanceWPCSafeTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrTransferIsNotPending.Error())
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("ApproveFullBalanceWPCSafeTransferNegative_MissingTransferID", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareApproveWPCSafeTransferObject()
		transfer.TransferID = ""

		response, err := tm.ApproveFullBalanceWPCSafeTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrTransferIDRequire.Error(), "Expected TransferIDRequired error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("ApproveFullBalanceWPCSafeTransferNegative_MissingAddress", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareApproveWPCSafeTransferObject()
		transfer.Address = common.Address{}

		response, err := tm.ApproveFullBalanceWPCSafeTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrAddressRequire.Error(), "Expected AddressRequired error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("ApproveFullBalanceWPCSafeTransferNegative_ManagerIsNotReady", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)

		transfer := prepareApproveWPCSafeTransferObject()

		response, err := tm.ApproveFullBalanceWPCSafeTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrTransferNotReady.Error(), "Expected ManagerIsNotReady error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("ApproveFullBalanceWPCSafeTransferNegative_TransferIsApproved", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareApproveWPCSafeTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(destPassphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(destWallet, nil)
		mockContractManager.EXPECT().GetTransferState(ctx, []byte(transfer.TransferID)).
			Return(uint8(2), nil)

		response, err := tm.ApproveFullBalanceWPCSafeTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrTransferIsNotPending.Error())
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("ApproveFullBalanceWPCSafeTransferNegative_AssetboxIsNotOwnerOfWallet", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareApproveWPCSafeTransferObject()
		transfer.Address = sourceAssetbox

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(destPassphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(destWallet, nil)

		response, err := tm.ApproveFullBalanceWPCSafeTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrAssetboxIsNotOwnerOfWallet.Error(), "Expected AssetboxIsNotOwnerOfWallet error")
		assert.Empty(t, response, "Response wasn't empty")
	})
}

func TestApproveWPCSafeTransferAsync(t *testing.T) {
	ctx := context.TODO()

	mockController := gomock.NewController(t)
	defer mockController.Finish()

	t.Run("ApproveWPCSafeTransferAsyncPositive", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareApproveWPCSafeTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(destPassphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(destWallet, nil)
		mockContractManager.EXPECT().GetTransferState(ctx, []byte(transfer.TransferID)).
			Return(uint8(1), nil)

		blockNum, txHash := createTransferResponse()

		key, err := keystore.DecryptKey(destWallet, string(destPassphrase))
		if err != nil {
			panic("Error during key decryption")
		}

		mockContractManager.EXPECT().
			ApproveWPCSafeTransfer(ctx, []byte(transfer.TransferID), nil, key.PrivateKey, true).
			Return(blockNum, txHash, nil)

		response, err := tm.ApproveWPCSafeTransferAsync(ctx, transfer)

		assert.NoError(t, err, "Unexpected error")
		assert.Equal(t, mapTransferResponse(blockNum, txHash), response, "Response differs from expected")
	})

	t.Run("ApproveWPCSafeTransferAsyncNegative_InvalidWallet", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareApproveWPCSafeTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(destPassphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(invalidWallet, nil)

		response, err := tm.ApproveWPCSafeTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrFailedToParseKey.Error(), "Expected InvalidWallet error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("ApproveWPCSafeTransferAsyncNegative_WrongPassphrase", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareApproveWPCSafeTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(wrongPass, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(destWallet, nil)

		response, err := tm.ApproveWPCSafeTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrWrongPassphrase.Error(), "Expected WrongPassphrase error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("ApproveWPCSafeTransferAsyncNegative_TransferHasExpired", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareApproveWPCSafeTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(destPassphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(destWallet, nil)
		mockContractManager.EXPECT().GetTransferState(ctx, []byte(transfer.TransferID)).
			Return(uint8(3), nil)

		response, err := tm.ApproveWPCSafeTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrTransferIsNotPending.Error())
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("ApproveWPCSafeTransferAsyncNegative_TransferMissing", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareApproveWPCSafeTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(destPassphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(destWallet, nil)
		mockContractManager.EXPECT().GetTransferState(ctx, []byte(transfer.TransferID)).
			Return(uint8(0), nil)

		response, err := tm.ApproveWPCSafeTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrTransferIsNotPending.Error())
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("ApproveWPCSafeTransferAsyncNegative_MissingTransferID", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareApproveWPCSafeTransferObject()
		transfer.TransferID = ""

		response, err := tm.ApproveWPCSafeTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrTransferIDRequire.Error(), "Expected TransferIDRequired error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("ApproveWPCSafeTransferAsyncNegative_MissingAddress", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareApproveWPCSafeTransferObject()
		transfer.Address = common.Address{}

		response, err := tm.ApproveWPCSafeTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrAddressRequire.Error(), "Expected AddressRequired error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("ApproveWPCSafeTransferAsyncNegative_ManagerIsNotReady", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)

		transfer := prepareApproveWPCSafeTransferObject()

		response, err := tm.ApproveWPCSafeTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrTransferNotReady.Error(), "Expected ManagerIsNotReady error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("ApproveWPCSafeTransferAsyncNegative_TransferIsApproved", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareApproveWPCSafeTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(destPassphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(destWallet, nil)
		mockContractManager.EXPECT().GetTransferState(ctx, []byte(transfer.TransferID)).
			Return(uint8(2), nil)

		response, err := tm.ApproveWPCSafeTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrTransferIsNotPending.Error())
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("ApproveWPCSafeTransferAsyncNegative_AssetboxIsNotOwnerOfWallet", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareApproveWPCSafeTransferObject()
		transfer.Address = sourceAssetbox

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(destPassphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(destWallet, nil)

		response, err := tm.ApproveWPCSafeTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrAssetboxIsNotOwnerOfWallet.Error(), "Expected AssetboxIsNotOwnerOfWallet error")
		assert.Empty(t, response, "Response wasn't empty")
	})
}

func TestApproveFullBalanceWPCSafeTransferAsync(t *testing.T) {
	ctx := context.TODO()

	mockController := gomock.NewController(t)
	defer mockController.Finish()

	t.Run("ApproveFullBalanceWPCSafeTransferPositive", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareApproveWPCSafeTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(destPassphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(destWallet, nil)
		mockContractManager.EXPECT().GetTransferState(ctx, []byte(transfer.TransferID)).
			Return(uint8(1), nil)

		blockNum, txHash := createTransferResponse()

		key, err := keystore.DecryptKey(destWallet, string(destPassphrase))
		if err != nil {
			panic("Error during key decryption")
		}

		mockContractManager.EXPECT().
			ApproveFullBalanceWPCSafeTransfer(ctx, []byte(transfer.TransferID), nil, key.PrivateKey, true).
			Return(blockNum, txHash, nil)

		response, err := tm.ApproveFullBalanceWPCSafeTransferAsync(ctx, transfer)

		assert.NoError(t, err, "Unexpected error")
		assert.Equal(t, mapTransferResponse(blockNum, txHash), response, "Response differs from expected")
	})

	t.Run("ApproveFullBalanceWPCSafeTransferNegative_InvalidWallet", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareApproveWPCSafeTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(destPassphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(invalidWallet, nil)

		response, err := tm.ApproveFullBalanceWPCSafeTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrFailedToParseKey.Error(), "Expected InvalidWallet error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("ApproveFullBalanceWPCSafeTransferNegative_WrongPassphrase", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareApproveWPCSafeTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(wrongPass, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(destWallet, nil)

		response, err := tm.ApproveFullBalanceWPCSafeTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrWrongPassphrase.Error(), "Expected WrongPassphrase error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("ApproveFullBalanceWPCSafeTransferNegative_TransferHasExpired", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareApproveWPCSafeTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(destPassphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(destWallet, nil)
		mockContractManager.EXPECT().GetTransferState(ctx, []byte(transfer.TransferID)).
			Return(uint8(3), nil)

		response, err := tm.ApproveFullBalanceWPCSafeTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrTransferIsNotPending.Error())
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("ApproveFullBalanceWPCSafeTransferNegative_TransferMissing", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareApproveWPCSafeTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(destPassphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(destWallet, nil)
		mockContractManager.EXPECT().GetTransferState(ctx, []byte(transfer.TransferID)).
			Return(uint8(0), nil)

		response, err := tm.ApproveFullBalanceWPCSafeTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrTransferIsNotPending.Error())
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("ApproveFullBalanceWPCSafeTransferNegative_MissingTransferID", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareApproveWPCSafeTransferObject()
		transfer.TransferID = ""

		response, err := tm.ApproveFullBalanceWPCSafeTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrTransferIDRequire.Error(), "Expected TransferIDRequired error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("ApproveFullBalanceWPCSafeTransferNegative_MissingAddress", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareApproveWPCSafeTransferObject()
		transfer.Address = common.Address{}

		response, err := tm.ApproveFullBalanceWPCSafeTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrAddressRequire.Error(), "Expected AddressRequired error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("ApproveFullBalanceWPCSafeTransferNegative_ManagerIsNotReady", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)

		transfer := prepareApproveWPCSafeTransferObject()

		response, err := tm.ApproveFullBalanceWPCSafeTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrTransferNotReady.Error(), "Expected ManagerIsNotReady error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("ApproveFullBalanceWPCSafeTransferNegative_TransferIsApproved", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareApproveWPCSafeTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(destPassphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(destWallet, nil)
		mockContractManager.EXPECT().GetTransferState(ctx, []byte(transfer.TransferID)).
			Return(uint8(2), nil)

		response, err := tm.ApproveFullBalanceWPCSafeTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrTransferIsNotPending.Error())
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("ApproveFullBalanceWPCSafeTransferNegative_AssetboxIsNotOwnerOfWallet", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareApproveWPCSafeTransferObject()
		transfer.Address = sourceAssetbox

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(destPassphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(destWallet, nil)

		response, err := tm.ApproveFullBalanceWPCSafeTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrAssetboxIsNotOwnerOfWallet.Error(), "Expected AssetboxIsNotOwnerOfWallet error")
		assert.Empty(t, response, "Response wasn't empty")
	})
}
