package transfer

import (
	"context"
	"testing"

	"github.com/simcord-llc/bitbon-system-blockchain/accounts/keystore"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/simcord-llc/bitbon-system-blockchain/common"
	"github.com/stretchr/testify/assert"
)

func TestCancelWPCSafeTransfer(t *testing.T) {
	ctx := context.TODO()

	mockController := gomock.NewController(t)
	defer mockController.Finish()

	t.Run("CancelWPCSafeTransferPositive", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareCancelSafeTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(passphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(wallet, nil)
		mockContractManager.EXPECT().TransferExists(ctx, []byte(transfer.TransferID)).
			Return(true, nil)

		blockNum, txHash := createTransferResponse()

		key, err := keystore.DecryptKey(wallet, string(passphrase))
		if err != nil {
			panic("Error during key decryption")
		}

		mockContractManager.EXPECT().
			CancelWPCSafeTransfer(ctx, []byte(transfer.TransferID), nil, key.PrivateKey, false).
			Return(blockNum, txHash, nil)

		response, err := tm.CancelWPCSafeTransfer(ctx, transfer)

		assert.NoError(t, err, "Unexpected error")
		assert.Equal(t, mapTransferResponse(blockNum, txHash), response, "Response differs from expected")
	})

	t.Run("CancelWPCSafeTransferNegative_InvalidWallet", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareCancelSafeTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(passphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(invalidWallet, nil)

		response, err := tm.CancelWPCSafeTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrFailedToParseKey.Error(), "Expected InvalidWallet error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("CancelWPCSafeTransferNegative_WrongPassphrase", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareCancelSafeTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(wrongPass, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(wallet, nil)

		response, err := tm.CancelWPCSafeTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrWrongPassphrase.Error(), "Expected WrongPassphrase error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("CancelWPCSafeTransferNegative_TransferMissing", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareCancelSafeTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(passphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(wallet, nil)
		mockContractManager.EXPECT().TransferExists(ctx, []byte(transfer.TransferID)).
			Return(false, nil)

		response, err := tm.CancelWPCSafeTransfer(ctx, transfer)

		assert.EqualError(t, err, errors.Errorf("transfer with id %q does not exist", transfer.TransferID).
			Error(), "Expected TransferDoesntExist error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("CancelWPCSafeTransferNegative_MissingTransferID", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareCancelSafeTransferObject()
		transfer.TransferID = ""

		response, err := tm.CancelWPCSafeTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrTransferIDRequire.Error(), "Expected TransferIDRequired error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("CancelWPCSafeTransferNegative_MissingAddress", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareCancelSafeTransferObject()
		transfer.Address = common.Address{}

		response, err := tm.CancelWPCSafeTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrAddressRequire.Error(), "Expected AddressRequired error")
		assert.Empty(t, response, "Response wasn't empty")
	})
}

func TestCancelWPCSafeTransferAsync(t *testing.T) {
	ctx := context.TODO()

	mockController := gomock.NewController(t)
	defer mockController.Finish()

	t.Run("CancelWPCSafeTransferAsyncPositive", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareCancelSafeTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(passphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(wallet, nil)
		mockContractManager.EXPECT().TransferExists(ctx, []byte(transfer.TransferID)).
			Return(true, nil)

		blockNum, txHash := createTransferResponse()

		key, err := keystore.DecryptKey(wallet, string(passphrase))
		if err != nil {
			panic("Error during key decryption")
		}

		mockContractManager.EXPECT().
			CancelWPCSafeTransfer(ctx, []byte(transfer.TransferID), nil, key.PrivateKey, true).
			Return(blockNum, txHash, nil)

		response, err := tm.CancelWPCSafeTransferAsync(ctx, transfer)

		assert.NoError(t, err, "Unexpected error")
		assert.Equal(t, mapTransferResponse(blockNum, txHash), response, "Response differs from expected")
	})

	t.Run("CancelWPCSafeTransferAsyncNegative_InvalidWallet", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareCancelSafeTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(passphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(invalidWallet, nil)

		response, err := tm.CancelWPCSafeTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrFailedToParseKey.Error(), "Expected InvalidWallet error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("CancelWPCSafeTransferAsyncNegative_WrongPassphrase", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareCancelSafeTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(wrongPass, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(wallet, nil)

		response, err := tm.CancelWPCSafeTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrWrongPassphrase.Error(), "Expected WrongPassphrase error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("CancelWPCSafeTransferAsyncNegative_TransferMissing", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareCancelSafeTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(passphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(wallet, nil)
		mockContractManager.EXPECT().TransferExists(ctx, []byte(transfer.TransferID)).
			Return(false, nil)

		response, err := tm.CancelWPCSafeTransferAsync(ctx, transfer)

		assert.EqualError(t, err, errors.Errorf("transfer with id %q does not exist", transfer.TransferID).
			Error(), "Expected TransferDoesntExist error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("CancelWPCSafeTransferAsyncNegative_MissingTransferID", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareCancelSafeTransferObject()
		transfer.TransferID = ""

		response, err := tm.CancelWPCSafeTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrTransferIDRequire.Error(), "Expected TransferIDRequired error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("CancelWPCSafeTransferAsyncNegative_MissingAddress", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareCancelSafeTransferObject()
		transfer.Address = common.Address{}

		response, err := tm.CancelWPCSafeTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrAddressRequire.Error(), "Expected AddressRequired error")
		assert.Empty(t, response, "Response wasn't empty")
	})
}
