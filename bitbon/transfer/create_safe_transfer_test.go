package transfer

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/simcord-llc/bitbon-system-blockchain/accounts/keystore"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/simcord-llc/bitbon-system-blockchain/common"
	"github.com/stretchr/testify/assert"
)

func TestCreateSafeTransfer(t *testing.T) {
	ctx := context.TODO()

	mockController := gomock.NewController(t)
	defer mockController.Finish()

	t.Run("CreateSafeTransferPositive", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareCreateSafeTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(passphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(wallet, nil)
		mockContractManager.EXPECT().GetAssetboxBalance(ctx, transfer.From).Return(getSufficientBalance(), nil)
		mockContractManager.EXPECT().TransferExists(ctx, []byte(transfer.TransferID)).Return(false, nil)

		key, err := keystore.DecryptKey(wallet, string(passphrase))
		if err != nil {
			panic("Error during key decryption")
		}

		blockNum, txHash := createTransferResponse()

		mockContractManager.EXPECT().CreateSafeTransfer(ctx, gomock.Any(), key.PrivateKey, false).
			Return(blockNum, txHash, nil)

		response, err := tm.CreateSafeTransfer(ctx, transfer)

		assert.NoError(t, err, "Unexpected error")
		assert.Equal(t, mapTransferResponse(blockNum, txHash), response, "Response differs from expected")
	})

	t.Run("CreateSafeTransferNegative_InvalidWallet", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareCreateSafeTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(passphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(invalidWallet, nil)

		response, err := tm.CreateSafeTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrFailedToParseKey.Error(), "Expected InvalidWallet error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("CreateSafeTransferNegative_WrongPassphrase", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareCreateSafeTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(wrongPass, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(wallet, nil)

		response, err := tm.CreateSafeTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrWrongPassphrase.Error(), "Expected WrongPassphrase error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("CreateSafeTransferNegative_TransferExists", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareCreateSafeTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(passphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(wallet, nil)
		mockContractManager.EXPECT().GetAssetboxBalance(ctx, transfer.From).Return(getSufficientBalance(), nil)
		mockContractManager.EXPECT().TransferExists(ctx, []byte(transfer.TransferID)).Return(true, nil)

		response, err := tm.CreateSafeTransfer(ctx, transfer)

		assert.EqualError(t, err, errors.Errorf("transfer with id %q already exists", transfer.TransferID).
			Error(), "Expected TransferExists error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("CreateSafeTransferNegative_InsufficientBalance", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareCreateSafeTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(passphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(wallet, nil)
		mockContractManager.EXPECT().GetAssetboxBalance(ctx, transfer.From).Return(getInsufficientBalance(), nil)

		response, err := tm.CreateSafeTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrLowAssetboxBalance.Error(), "Expected LowAssetboxBalance error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("CreateSafeTransferNegative_NilAmount", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareCreateSafeTransferObject()
		transfer.Value = nil

		response, err := tm.CreateSafeTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrValueRequire.Error(), "Expected RequiredValue error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("CreateSafeTransferNegative_ZeroAmount", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareCreateSafeTransferObject()
		transfer.Value = big.NewInt(0)

		response, err := tm.CreateSafeTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrValueRequire.Error(), "Expected RequiredValue error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("CreateSafeTransferNegative_NegativeAmount", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareCreateSafeTransferObject()
		transfer.Value = big.NewInt(-10)

		response, err := tm.CreateSafeTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrValueRequire.Error(), "Expected RequiredValue error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("CreateSafeTransferNegative_AlreadyExpired", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareCreateSafeTransferObject()
		transfer.ExpiresAt = time.Now().Unix()

		response, err := tm.CreateSafeTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrExpireAtPassed.Error(), "Expected AlreadyExpired error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("CreateSafeTransferNegative_MissingProtectionCode", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareCreateSafeTransferObject()
		transfer.ProtectionCode = ""

		response, err := tm.CreateSafeTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrProtectionCodeRequire.Error(), "Expected ProtectionCodeRequired error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("CreateSafeTransferNegative_MissingTransferID", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareCreateSafeTransferObject()
		transfer.TransferID = ""

		response, err := tm.CreateSafeTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrTransferIDRequire.Error(), "Expected TransferIDRequired error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("CreateSafeTransferNegative_FromAndToAreIdentical", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareCreateSafeTransferObject()
		transfer.To = transfer.From

		response, err := tm.CreateSafeTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrSameAccount.Error(), "Expected SameAccount error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("CreateSafeTransferNegative_MissingTo", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareCreateSafeTransferObject()
		transfer.To = common.Address{}

		response, err := tm.CreateSafeTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrToRequire.Error(), "Expected ToRequired error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("CreateSafeTransferNegative_MissingFrom", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareCreateSafeTransferObject()
		transfer.From = common.Address{}

		response, err := tm.CreateSafeTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrFromRequire.Error(), "Expected FromRequired error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("CreateSafeTransferNegative_ManagerIsNotReady", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)

		transfer := prepareCreateSafeTransferObject()

		response, err := tm.CreateSafeTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrTransferNotReady.Error(), "Expected ManagerIsNotReady error")
		assert.Empty(t, response, "Response wasn't empty")
	})
}

func TestCreateSafeAsyncTransfer(t *testing.T) {
	ctx := context.TODO()

	mockController := gomock.NewController(t)
	defer mockController.Finish()

	t.Run("CreateSafeTransferPositive", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareCreateSafeTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(passphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(wallet, nil)
		mockContractManager.EXPECT().GetAssetboxBalance(ctx, transfer.From).Return(getSufficientBalance(), nil)
		mockContractManager.EXPECT().TransferExists(ctx, []byte(transfer.TransferID)).Return(false, nil)

		key, err := keystore.DecryptKey(wallet, string(passphrase))
		if err != nil {
			panic("Error during key decryption")
		}

		blockNum, txHash := createTransferResponse()
		blockNum = uint64(0)

		mockContractManager.EXPECT().CreateSafeTransfer(ctx, gomock.Any(), key.PrivateKey, true).
			Return(blockNum, txHash, nil)

		response, err := tm.CreateSafeTransferAsync(ctx, transfer)

		assert.NoError(t, err, "Unexpected error")
		assert.Equal(t, mapTransferResponse(blockNum, txHash), response, "Response differs from expected")
	})

	t.Run("CreateSafeTransferAsyncNegative_InvalidWallet", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareCreateSafeTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(passphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(invalidWallet, nil)

		response, err := tm.CreateSafeTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrFailedToParseKey.Error(), "Expected InvalidWallet error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("CreateSafeTransferAsyncNegative_WrongPassphrase", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareCreateSafeTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(wrongPass, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(wallet, nil)

		response, err := tm.CreateSafeTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrWrongPassphrase.Error(), "Expected WrongPassphrase error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("CreateSafeTransferNegative_TransferExists", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareCreateSafeTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(passphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(wallet, nil)
		mockContractManager.EXPECT().GetAssetboxBalance(ctx, transfer.From).Return(getSufficientBalance(), nil)
		mockContractManager.EXPECT().TransferExists(ctx, []byte(transfer.TransferID)).Return(true, nil)

		response, err := tm.CreateSafeTransferAsync(ctx, transfer)

		assert.EqualError(t, err, errors.Errorf("transfer with id %q already exists", transfer.TransferID).
			Error(), "Expected TransferExists error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("CreateSafeTransferNegative_InsufficientBalance", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareCreateSafeTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(passphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(wallet, nil)
		mockContractManager.EXPECT().GetAssetboxBalance(ctx, transfer.From).Return(getInsufficientBalance(), nil)

		response, err := tm.CreateSafeTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrLowAssetboxBalance.Error(), "Expected LowAssetboxBalance error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("CreateSafeTransferNegative_NilAmount", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareCreateSafeTransferObject()
		transfer.Value = nil

		response, err := tm.CreateSafeTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrValueRequire.Error(), "Expected RequiredValue error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("CreateSafeTransferNegative_ZeroAmount", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareCreateSafeTransferObject()
		transfer.Value = big.NewInt(0)

		response, err := tm.CreateSafeTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrValueRequire.Error(), "Expected RequiredValue error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("CreateSafeTransferNegative_NegativeAmount", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareCreateSafeTransferObject()
		transfer.Value = big.NewInt(-10)

		response, err := tm.CreateSafeTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrValueRequire.Error(), "Expected RequiredValue error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("CreateSafeTransferNegative_AlreadyExpired", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareCreateSafeTransferObject()
		transfer.ExpiresAt = time.Now().Unix()

		response, err := tm.CreateSafeTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrExpireAtPassed.Error(), "Expected AlreadyExpired error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("CreateSafeTransferNegative_MissingProtectionCode", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareCreateSafeTransferObject()
		transfer.ProtectionCode = ""

		response, err := tm.CreateSafeTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrProtectionCodeRequire.Error(), "Expected ProtectionCodeRequired error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("CreateSafeTransferNegative_MissingTransferID", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareCreateSafeTransferObject()
		transfer.TransferID = ""

		response, err := tm.CreateSafeTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrTransferIDRequire.Error(), "Expected TransferIDRequired error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("CreateSafeTransferNegative_FromAndToAreIdentical", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareCreateSafeTransferObject()
		transfer.To = transfer.From

		response, err := tm.CreateSafeTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrSameAccount.Error(), "Expected SameAccount error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("CreateSafeTransferNegative_MissingTo", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareCreateSafeTransferObject()
		transfer.To = common.Address{}

		response, err := tm.CreateSafeTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrToRequire.Error(), "Expected ToRequired error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("CreateSafeTransferNegative_MissingFrom", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareCreateSafeTransferObject()
		transfer.From = common.Address{}

		response, err := tm.CreateSafeTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrFromRequire.Error(), "Expected FromRequired error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("CreateSafeTransferNegative_ManagerIsNotReady", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)

		transfer := prepareCreateSafeTransferObject()

		response, err := tm.CreateSafeTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrTransferNotReady.Error(), "Expected ManagerIsNotReady error")
		assert.Empty(t, response, "Response wasn't empty")
	})
}
