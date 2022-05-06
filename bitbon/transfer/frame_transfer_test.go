package transfer

import (
	"context"
	"math/big"
	"testing"

	"github.com/simcord-llc/bitbon-system-blockchain/accounts/keystore"

	"github.com/golang/mock/gomock"
	"github.com/simcord-llc/bitbon-system-blockchain/common"
	"github.com/stretchr/testify/assert"
)

func TestFrameTransfer(t *testing.T) {
	ctx := context.TODO()

	mockController := gomock.NewController(t)
	defer mockController.Finish()

	t.Run("FrameTransferPositive", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareFrameTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)

		transfer := prepareFrameTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(sourcePassphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(sourceWallet, nil)
		mockContractManager.EXPECT().GetAssetboxBalance(ctx, transfer.From).Return(getSufficientBalance(), nil)

		key, err := keystore.DecryptKey(sourceWallet, string(sourcePassphrase))
		if err != nil {
			panic("Error during key decryption")
		}

		blockNum, txHash := createTransferResponse()

		mockContractManager.EXPECT().
			FrameTransfer(ctx, transfer.To, transfer.Value, transfer.ExtraData, key.PrivateKey, false).
			Return(blockNum, txHash, nil)

		response, err := tm.FrameTransfer(ctx, transfer)

		assert.NoError(t, err, "Unexpected error occurred")
		assert.Equal(t, mapTransferResponse(blockNum, txHash), response, "Response differs from expected")
	})

	t.Run("FrameTransferNegative_InvalidWallet", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareFrameTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)

		transfer := prepareFrameTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(sourcePassphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(invalidWallet, nil)

		response, err := tm.FrameTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrFailedToParseKey.Error(), "Expected InvalidWallet error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("FrameTransferNegative_WrongPassphrase", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareFrameTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)

		transfer := prepareFrameTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(wrongPass, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(sourceWallet, nil)

		response, err := tm.FrameTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrWrongPassphrase.Error(), "Expected WrongPassphrase error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("FrameTransferNegative_InsufficientBalance", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareFrameTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)

		transfer := prepareFrameTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(sourcePassphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(sourceWallet, nil)
		mockContractManager.EXPECT().GetAssetboxBalance(ctx, transfer.From).Return(getInsufficientBalance(), nil)

		response, err := tm.FrameTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrLowAssetboxBalance.Error(), "Expected LowAssetboxBalance error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("FrameTransferNegative_NilAmount", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareFrameTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)

		transfer := prepareFrameTransferObject()
		transfer.Value = nil

		response, err := tm.FrameTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrValueRequire.Error(), "Expected RequiredValue error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("FrameTransferNegative_NegativeAmount", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareFrameTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)

		transfer := prepareFrameTransferObject()
		transfer.Value = big.NewInt(-10)

		response, err := tm.FrameTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrValueRequire.Error(), "Expected RequiredValue error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("FrameTransferNegative_ZeroAmount", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareFrameTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)

		transfer := prepareFrameTransferObject()
		transfer.Value = big.NewInt(0)

		response, err := tm.FrameTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrValueRequire.Error(), "Expected RequiredValue error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("FrameTransferNegative_FromAndToAreIdentical", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareFrameTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)

		transfer := prepareFrameTransferObject()
		transfer.To = transfer.From

		response, err := tm.FrameTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrSameAccount.Error(), "Expected SameAccount error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("FrameTransferNegative_MissingTo", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareFrameTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)

		transfer := prepareFrameTransferObject()
		transfer.To = common.Address{}

		response, err := tm.FrameTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrToRequire.Error(), "Expected ToRequired error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("FrameTransferNegative_MissingFrom", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareFrameTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)

		transfer := prepareFrameTransferObject()
		transfer.From = common.Address{}

		response, err := tm.FrameTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrFromRequire.Error(), "Expected FromRequired error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("FrameTransferNegative_ManagerNotReady", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareFrameTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)

		// Recreating ready channel of transfer manager to make it not ready
		tm.ready = make(chan struct{})

		transfer := prepareFrameTransferObject()

		response, err := tm.FrameTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrTransferNotReady.Error(), "Expected ManagerIsNotReady error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("FrameTransferNegative_AssetboxIsNotOwnerOfWallet", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareFrameTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)

		transfer := prepareFrameTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(destPassphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(destWallet, nil)

		response, err := tm.FrameTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrAssetboxIsNotOwnerOfWallet.Error(), "Expected AssetboxIsNotOwnerOfWallet error")
		assert.Empty(t, response, "Response wasn't empty")
	})
}

func TestFrameTransferAsync(t *testing.T) {
	ctx := context.TODO()

	mockController := gomock.NewController(t)
	defer mockController.Finish()

	t.Run("FrameTransferAsyncPositive", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareFrameTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)

		transfer := prepareFrameTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(sourcePassphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(sourceWallet, nil)
		mockContractManager.EXPECT().GetAssetboxBalance(ctx, transfer.From).Return(getSufficientBalance(), nil)

		key, err := keystore.DecryptKey(sourceWallet, string(sourcePassphrase))
		if err != nil {
			panic("Error during key decryption")
		}

		blockNum, txHash := createTransferResponse()
		blockNum = uint64(0)

		mockContractManager.EXPECT().
			FrameTransfer(ctx, transfer.To, transfer.Value, transfer.ExtraData, key.PrivateKey, true).
			Return(blockNum, txHash, nil)

		response, err := tm.FrameTransferAsync(ctx, transfer)

		assert.NoError(t, err, "Unexpected error occurred")
		assert.Equal(t, mapTransferResponse(blockNum, txHash), response, "Response differs from expected")
	})

	t.Run("FrameTransferAsyncNegative_InvalidWallet", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareFrameTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)

		transfer := prepareFrameTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(sourcePassphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(invalidWallet, nil)

		response, err := tm.FrameTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrFailedToParseKey.Error(), "Expected InvalidWallet error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("FrameTransferAsyncNegative_WrongPassphrase", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareFrameTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)

		transfer := prepareFrameTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(wrongPass, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(sourceWallet, nil)

		response, err := tm.FrameTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrWrongPassphrase.Error(), "Expected WrongPassphrase error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("FrameTransferAsyncNegative_InsufficientBalance", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareFrameTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)

		transfer := prepareFrameTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(sourcePassphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(sourceWallet, nil)
		mockContractManager.EXPECT().GetAssetboxBalance(ctx, transfer.From).Return(getInsufficientBalance(), nil)

		response, err := tm.FrameTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrLowAssetboxBalance.Error(), "Expected LowAssetboxBalance error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("FrameTransferAsyncNegative_NilAmount", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareFrameTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)

		transfer := prepareFrameTransferObject()
		transfer.Value = nil

		response, err := tm.FrameTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrValueRequire.Error(), "Expected RequiredValue error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("FrameTransferAsyncNegative_NegativeAmount", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareFrameTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)

		transfer := prepareFrameTransferObject()
		transfer.Value = big.NewInt(-10)

		response, err := tm.FrameTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrValueRequire.Error(), "Expected RequiredValue error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("FrameTransferAsyncNegative_ZeroAmount", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareFrameTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)

		transfer := prepareFrameTransferObject()
		transfer.Value = big.NewInt(0)

		response, err := tm.FrameTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrValueRequire.Error(), "Expected RequiredValue error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("FrameTransferAsyncNegative_FromAndToAreIdentical", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareFrameTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)

		transfer := prepareFrameTransferObject()
		transfer.To = transfer.From

		response, err := tm.FrameTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrSameAccount.Error(), "Expected SameAccount error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("FrameTransferAsyncNegative_MissingTo", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareFrameTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)

		transfer := prepareFrameTransferObject()
		transfer.To = common.Address{}

		response, err := tm.FrameTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrToRequire.Error(), "Expected ToRequired error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("FrameTransferAsyncNegative_MissingFrom", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareFrameTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)

		transfer := prepareFrameTransferObject()
		transfer.From = common.Address{}

		response, err := tm.FrameTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrFromRequire.Error(), "Expected FromRequired error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("FrameTransferAsyncNegative_ManagerNotReady", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareFrameTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)

		// Recreating ready channel of transfer manager to make it not ready
		tm.ready = make(chan struct{})

		transfer := prepareFrameTransferObject()

		response, err := tm.FrameTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrTransferNotReady.Error(), "Expected ManagerIsNotReady error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("FrameTransferAsyncNegative_AssetboxIsNotOwnerOfWallet", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareFrameTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)

		transfer := prepareFrameTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(destPassphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(destWallet, nil)

		response, err := tm.FrameTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrAssetboxIsNotOwnerOfWallet.Error(), "Expected AssetboxIsNotOwnerOfWallet error")
		assert.Empty(t, response, "Response wasn't empty")
	})
}
