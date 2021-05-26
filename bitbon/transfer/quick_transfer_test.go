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

func TestQuickTransfer(t *testing.T) {
	ctx := context.TODO()

	mockController := gomock.NewController(t)
	defer mockController.Finish()

	t.Run("QuickTransferPositive", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareQuickTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)

		transfer := prepareQuickTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(passphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(wallet, nil)
		mockContractManager.EXPECT().GetAssetboxBalance(ctx, transfer.From).Return(getSufficientBalance(), nil)

		key, err := keystore.DecryptKey(wallet, string(passphrase))
		if err != nil {
			panic("Error during key decryption")
		}

		blockNum, txHash := createTransferResponse()

		mockContractManager.EXPECT().
			QuickTransfer(ctx, transfer.To, transfer.Value, transfer.ExtraData, key.PrivateKey, false).
			Return(blockNum, txHash, nil)

		response, err := tm.QuickTransfer(ctx, transfer)

		assert.NoError(t, err, "Unexpected error occurred")
		assert.Equal(t, mapTransferResponse(blockNum, txHash), response, "Response differs from expected")
	})

	t.Run("QuickTransferNegative_InvalidWallet", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareQuickTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)

		transfer := prepareQuickTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(passphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(invalidWallet, nil)

		response, err := tm.QuickTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrFailedToParseKey.Error(), "Expected InvalidWallet error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("QuickTransferNegative_WrongPassphrase", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareQuickTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)

		transfer := prepareQuickTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(wrongPass, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(wallet, nil)

		response, err := tm.QuickTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrWrongPassphrase.Error(), "Expected WrongPassphrase error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("QuickTransferNegative_InsufficientBalance", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareQuickTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)

		transfer := prepareQuickTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(passphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(wallet, nil)
		mockContractManager.EXPECT().GetAssetboxBalance(ctx, transfer.From).Return(getInsufficientBalance(), nil)

		response, err := tm.QuickTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrLowAssetboxBalance.Error(), "Expected LowAssetboxBalance error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("QuickTransferNegative_NilAmount", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareQuickTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)

		transfer := prepareQuickTransferObject()
		transfer.Value = nil

		response, err := tm.QuickTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrValueRequire.Error(), "Expected RequiredValue error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("QuickTransferNegative_NegativeAmount", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareQuickTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)

		transfer := prepareQuickTransferObject()
		transfer.Value = big.NewInt(-10)

		response, err := tm.QuickTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrValueRequire.Error(), "Expected RequiredValue error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("QuickTransferNegative_ZeroAmount", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareQuickTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)

		transfer := prepareQuickTransferObject()
		transfer.Value = big.NewInt(0)

		response, err := tm.QuickTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrValueRequire.Error(), "Expected RequiredValue error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("QuickTransferNegative_EmptyAccountID", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareQuickTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)

		transfer := prepareQuickTransferObject()
		transfer.AccountID = ""

		response, err := tm.QuickTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrAccountIDRequire.Error(), "Expected EmptyAccountId error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("QuickTransferNegative_FromAndToAreIdentical", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareQuickTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)

		transfer := prepareQuickTransferObject()
		transfer.To = transfer.From

		response, err := tm.QuickTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrSameAccount.Error(), "Expected SameAccount error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("QuickTransferNegative_MissingTo", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareQuickTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)

		transfer := prepareQuickTransferObject()
		transfer.To = common.Address{}

		response, err := tm.QuickTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrToRequire.Error(), "Expected ToRequired error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("QuickTransferNegative_MissingFrom", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareQuickTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)

		transfer := prepareQuickTransferObject()
		transfer.From = common.Address{}

		response, err := tm.QuickTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrFromRequire.Error(), "Expected FromRequired error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("QuickTransferNegative_ManagerNotReady", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareQuickTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)

		// Recreating ready channel of transfer manager to make it not ready
		tm.ready = make(chan struct{})

		transfer := prepareQuickTransferObject()

		response, err := tm.QuickTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrTransferNotReady.Error(), "Expected ManagerIsNotReady error")
		assert.Empty(t, response, "Response wasn't empty")
	})
}

func TestQuickTransferAsync(t *testing.T) {
	ctx := context.TODO()

	mockController := gomock.NewController(t)
	defer mockController.Finish()

	t.Run("QuickTransferAsyncPositive", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareQuickTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)

		transfer := prepareQuickTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(passphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(wallet, nil)
		mockContractManager.EXPECT().GetAssetboxBalance(ctx, transfer.From).Return(getSufficientBalance(), nil)

		key, err := keystore.DecryptKey(wallet, string(passphrase))
		if err != nil {
			panic("Error during key decryption")
		}

		blockNum, txHash := createTransferResponse()
		blockNum = uint64(0)

		mockContractManager.EXPECT().
			QuickTransfer(ctx, transfer.To, transfer.Value, transfer.ExtraData, key.PrivateKey, true).
			Return(blockNum, txHash, nil)

		response, err := tm.QuickTransferAsync(ctx, transfer)

		assert.NoError(t, err, "Unexpected error occurred")
		assert.Equal(t, mapTransferResponse(blockNum, txHash), response, "Response differs from expected")
	})

	t.Run("QuickTransferAsyncNegative_InvalidWallet", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareQuickTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)

		transfer := prepareQuickTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(passphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(invalidWallet, nil)

		response, err := tm.QuickTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrFailedToParseKey.Error(), "Expected InvalidWallet error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("QuickTransferAsyncNegative_WrongPassphrase", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareQuickTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)

		transfer := prepareQuickTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(wrongPass, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(wallet, nil)

		response, err := tm.QuickTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrWrongPassphrase.Error(), "Expected WrongPassphrase error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("QuickTransferAsyncNegative_InsufficientBalance", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareQuickTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)

		transfer := prepareQuickTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(passphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(wallet, nil)
		mockContractManager.EXPECT().GetAssetboxBalance(ctx, transfer.From).Return(getInsufficientBalance(), nil)

		response, err := tm.QuickTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrLowAssetboxBalance.Error(), "Expected LowAssetboxBalance error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("QuickTransferAsyncNegative_NilAmount", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareQuickTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)

		transfer := prepareQuickTransferObject()
		transfer.Value = nil

		response, err := tm.QuickTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrValueRequire.Error(), "Expected RequiredValue error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("QuickTransferAsyncNegative_NegativeAmount", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareQuickTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)

		transfer := prepareQuickTransferObject()
		transfer.Value = big.NewInt(-10)

		response, err := tm.QuickTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrValueRequire.Error(), "Expected RequiredValue error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("QuickTransferAsyncNegative_ZeroAmount", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareQuickTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)

		transfer := prepareQuickTransferObject()
		transfer.Value = big.NewInt(0)

		response, err := tm.QuickTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrValueRequire.Error(), "Expected RequiredValue error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("QuickTransferAsyncNegative_EmptyAccountID", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareQuickTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)

		transfer := prepareQuickTransferObject()
		transfer.AccountID = ""

		response, err := tm.QuickTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrAccountIDRequire.Error(), "Expected EmptyAccountId error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("QuickTransferAsyncNegative_FromAndToAreIdentical", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareQuickTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)

		transfer := prepareQuickTransferObject()
		transfer.To = transfer.From

		response, err := tm.QuickTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrSameAccount.Error(), "Expected SameAccount error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("QuickTransferAsyncNegative_MissingTo", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareQuickTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)

		transfer := prepareQuickTransferObject()
		transfer.To = common.Address{}

		response, err := tm.QuickTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrToRequire.Error(), "Expected ToRequired error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("QuickTransferAsyncNegative_MissingFrom", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareQuickTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)

		transfer := prepareQuickTransferObject()
		transfer.From = common.Address{}

		response, err := tm.QuickTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrFromRequire.Error(), "Expected FromRequired error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("QuickTransferAsyncNegative_ManagerNotReady", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareQuickTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)

		// Recreating ready channel of transfer manager to make it not ready
		tm.ready = make(chan struct{})

		transfer := prepareQuickTransferObject()

		response, err := tm.QuickTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrTransferNotReady.Error(), "Expected ManagerIsNotReady error")
		assert.Empty(t, response, "Response wasn't empty")
	})
}
