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

func TestCreateWPCSafeTransfer(t *testing.T) {
	ctx := context.TODO()

	mockController := gomock.NewController(t)
	defer mockController.Finish()

	t.Run("CreateWPCSafeTransferPositive", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareCreateWPCSafeTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(sourcePassphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(sourceWallet, nil)
		mockContractManager.EXPECT().GetAssetboxBalance(ctx, transfer.From).Return(getSufficientBalance(), nil)
		mockContractManager.EXPECT().TransferExists(ctx, []byte(transfer.TransferID)).Return(false, nil)

		key, err := keystore.DecryptKey(sourceWallet, string(sourcePassphrase))
		if err != nil {
			panic("Error during key decryption")
		}

		blockNum, txHash := createTransferResponse()

		mockContractManager.EXPECT().CreateWPCSafeTransfer(ctx, gomock.Any(), key.PrivateKey, false).
			Return(blockNum, txHash, nil)

		response, err := tm.CreateWPCSafeTransfer(ctx, transfer)

		assert.NoError(t, err, "Unexpected error")
		assert.Equal(t, mapTransferResponse(blockNum, txHash), response, "Response differs from expected")
	})

	t.Run("CreateWPCSafeTransferNegative_InvalidWallet", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareCreateWPCSafeTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(sourcePassphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(invalidWallet, nil)

		response, err := tm.CreateWPCSafeTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrFailedToParseKey.Error(), "Expected InvalidWallet error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("CreateWPCSafeTransferNegative_WrongPassphrase", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareCreateWPCSafeTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(wrongPass, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(sourceWallet, nil)

		response, err := tm.CreateWPCSafeTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrWrongPassphrase.Error(), "Expected WrongPassphrase error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("CreateWPCSafeTransferNegative_TransferExists", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareCreateWPCSafeTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(sourcePassphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(sourceWallet, nil)
		mockContractManager.EXPECT().GetAssetboxBalance(ctx, transfer.From).Return(getSufficientBalance(), nil)
		mockContractManager.EXPECT().TransferExists(ctx, []byte(transfer.TransferID)).Return(true, nil)

		response, err := tm.CreateWPCSafeTransfer(ctx, transfer)

		assert.EqualError(t, err, errors.Errorf("transfer with id %q already exists", transfer.TransferID).
			Error(), "Expected TransferExists error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("CreateWPCSafeTransferNegative_InsufficientBalance", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareCreateWPCSafeTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(sourcePassphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(sourceWallet, nil)
		mockContractManager.EXPECT().GetAssetboxBalance(ctx, transfer.From).Return(getInsufficientBalance(), nil)

		response, err := tm.CreateWPCSafeTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrLowAssetboxBalance.Error(), "Expected LowAssetboxBalance error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("CreateWPCSafeTransferNegative_NilAmount", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareCreateWPCSafeTransferObject()
		transfer.Value = nil

		response, err := tm.CreateWPCSafeTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrValueRequire.Error(), "Expected RequiredValue error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("CreateWPCSafeTransferNegative_ZeroAmount", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareCreateWPCSafeTransferObject()
		transfer.Value = big.NewInt(0)

		response, err := tm.CreateWPCSafeTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrValueRequire.Error(), "Expected RequiredValue error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("CreateWPCSafeTransferNegative_NegativeAmount", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareCreateWPCSafeTransferObject()
		transfer.Value = big.NewInt(-10)

		response, err := tm.CreateWPCSafeTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrValueRequire.Error(), "Expected RequiredValue error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("CreateWPCSafeTransferNegative_AlreadyExpired", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareCreateWPCSafeTransferObject()
		transfer.ExpiresAt = time.Now().Unix()

		response, err := tm.CreateWPCSafeTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrExpireAtPassed.Error(), "Expected AlreadyExpired error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("CreateWPCSafeTransferNegative_MissingTransferID", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareCreateWPCSafeTransferObject()
		transfer.TransferID = ""

		response, err := tm.CreateWPCSafeTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrTransferIDRequire.Error(), "Expected TransferIDRequired error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("CreateWPCSafeTransferNegative_FromAndToAreIdentical", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareCreateWPCSafeTransferObject()
		transfer.To = transfer.From

		response, err := tm.CreateWPCSafeTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrSameAccount.Error(), "Expected SameAccount error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("CreateWPCSafeTransferNegative_MissingTo", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareCreateWPCSafeTransferObject()
		transfer.To = common.Address{}

		response, err := tm.CreateWPCSafeTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrToRequire.Error(), "Expected ToRequired error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("CreateWPCSafeTransferNegative_MissingFrom", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareCreateWPCSafeTransferObject()
		transfer.From = common.Address{}

		response, err := tm.CreateWPCSafeTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrFromRequire.Error(), "Expected FromRequired error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("CreateWPCSafeTransferNegative_ManagerIsNotReady", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)

		transfer := prepareCreateWPCSafeTransferObject()

		response, err := tm.CreateWPCSafeTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrTransferNotReady.Error(), "Expected ManagerIsNotReady error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("CreateWPCSafeTransferNegative_AssetboxIsNotOwnerOfWallet", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareCreateWPCSafeTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(destPassphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(destWallet, nil)

		response, err := tm.CreateWPCSafeTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrAssetboxIsNotOwnerOfWallet.Error(), "Expected AssetboxIsNotOwnerOfWallet error")
		assert.Empty(t, response, "Response wasn't empty")
	})
}

func TestCreateFullBalanceWPCSafeTransfer(t *testing.T) {
	ctx := context.TODO()

	mockController := gomock.NewController(t)
	defer mockController.Finish()

	t.Run("CreateFullBalanceWPCSafeTransferPositive", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareCreateWPCSafeTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(sourcePassphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(sourceWallet, nil)
		mockContractManager.EXPECT().GetAssetboxBalance(ctx, transfer.From).Return(getSufficientBalance(), nil)
		mockContractManager.EXPECT().TransferExists(ctx, []byte(transfer.TransferID)).Return(false, nil)
		mockContractManager.EXPECT().BalanceOfLocked(ctx, transfer.From).Return(big.NewInt(0), nil)

		key, err := keystore.DecryptKey(sourceWallet, string(sourcePassphrase))
		if err != nil {
			panic("Error during key decryption")
		}

		blockNum, txHash := createTransferResponse()

		mockContractManager.EXPECT().CreateFullBalanceWPCSafeTransfer(ctx, gomock.Any(), key.PrivateKey, false).
			Return(blockNum, txHash, nil)

		response, err := tm.CreateFullBalanceWPCSafeTransfer(ctx, transfer)

		assert.NoError(t, err, "Unexpected error")
		assert.Equal(t, mapTransferResponse(blockNum, txHash), response, "Response differs from expected")
	})

	t.Run("CreateFullBalanceWPCSafeTransferNegative_InvalidWallet", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareCreateWPCSafeTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(sourcePassphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(invalidWallet, nil)

		response, err := tm.CreateFullBalanceWPCSafeTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrFailedToParseKey.Error(), "Expected InvalidWallet error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("CreateFullBalanceWPCSafeTransferNegative_WrongPassphrase", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareCreateWPCSafeTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(wrongPass, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(sourceWallet, nil)

		response, err := tm.CreateFullBalanceWPCSafeTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrWrongPassphrase.Error(), "Expected WrongPassphrase error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("CreateFullBalanceWPCSafeTransferNegative_TransferExists", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareCreateWPCSafeTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(sourcePassphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(sourceWallet, nil)
		mockContractManager.EXPECT().GetAssetboxBalance(ctx, transfer.From).Return(getSufficientBalance(), nil)
		mockContractManager.EXPECT().TransferExists(ctx, []byte(transfer.TransferID)).Return(true, nil)
		mockContractManager.EXPECT().BalanceOfLocked(ctx, transfer.From).Return(big.NewInt(0), nil)

		response, err := tm.CreateFullBalanceWPCSafeTransfer(ctx, transfer)

		assert.EqualError(t, err, errors.Errorf("transfer with id %q already exists", transfer.TransferID).
			Error(), "Expected TransferExists error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("CreateFullBalanceWPCSafeTransferNegative_ZeroBalance", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareCreateWPCSafeTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(sourcePassphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(sourceWallet, nil)
		mockContractManager.EXPECT().GetAssetboxBalance(ctx, transfer.From).Return(big.NewInt(0), nil)

		response, err := tm.CreateFullBalanceWPCSafeTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrLowAssetboxBalance.Error(), "Expected LowAssetboxBalance error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("CreateFullBalanceWPCSafeTransferNegative_AlreadyExpired", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareCreateWPCSafeTransferObject()
		transfer.ExpiresAt = time.Now().Unix()

		response, err := tm.CreateFullBalanceWPCSafeTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrExpireAtPassed.Error(), "Expected AlreadyExpired error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("CreateFullBalanceWPCSafeTransferNegative_MissingTransferID", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareCreateWPCSafeTransferObject()
		transfer.TransferID = ""

		response, err := tm.CreateFullBalanceWPCSafeTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrTransferIDRequire.Error(), "Expected TransferIDRequired error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("CreateFullBalanceWPCSafeTransferNegative_FromAndToAreIdentical", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareCreateWPCSafeTransferObject()
		transfer.To = transfer.From

		response, err := tm.CreateFullBalanceWPCSafeTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrSameAccount.Error(), "Expected SameAccount error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("CreateFullBalanceWPCSafeTransferNegative_MissingTo", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareCreateWPCSafeTransferObject()
		transfer.To = common.Address{}

		response, err := tm.CreateFullBalanceWPCSafeTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrToRequire.Error(), "Expected ToRequired error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("CreateFullBalanceWPCSafeTransferNegative_MissingFrom", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareCreateWPCSafeTransferObject()
		transfer.From = common.Address{}

		response, err := tm.CreateFullBalanceWPCSafeTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrFromRequire.Error(), "Expected FromRequired error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("CreateFullBalanceWPCSafeTransferNegative_ManagerIsNotReady", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)

		transfer := prepareCreateWPCSafeTransferObject()

		response, err := tm.CreateFullBalanceWPCSafeTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrTransferNotReady.Error(), "Expected ManagerIsNotReady error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("CreateFullBalanceWPCSafeTransferNegative_AssetboxIsNotOwnerOfWallet", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareCreateWPCSafeTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(destPassphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(destWallet, nil)

		response, err := tm.CreateFullBalanceWPCSafeTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrAssetboxIsNotOwnerOfWallet.Error(), "Expected AssetboxIsNotOwnerOfWallet error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("CreateFullBalanceWPCSafeTransferNegative_LockedBalance", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareCreateWPCSafeTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(sourcePassphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(sourceWallet, nil)
		mockContractManager.EXPECT().GetAssetboxBalance(ctx, transfer.From).Return(getSufficientBalance(), nil)
		mockContractManager.EXPECT().BalanceOfLocked(ctx, transfer.From).Return(big.NewInt(1), nil)

		response, err := tm.CreateFullBalanceWPCSafeTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrHasLockedBalanceError.Error(), "Expected HasLockedBalanceError error")
		assert.Empty(t, response, "Response wasn't empty")
	})
}

func TestCreateWPCSafeTransferAsync(t *testing.T) {
	ctx := context.TODO()

	mockController := gomock.NewController(t)
	defer mockController.Finish()

	t.Run("CreateWPCSafeTransferAsyncPositive", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareCreateWPCSafeTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(sourcePassphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(sourceWallet, nil)
		mockContractManager.EXPECT().GetAssetboxBalance(ctx, transfer.From).Return(getSufficientBalance(), nil)
		mockContractManager.EXPECT().TransferExists(ctx, []byte(transfer.TransferID)).Return(false, nil)

		key, err := keystore.DecryptKey(sourceWallet, string(sourcePassphrase))
		if err != nil {
			panic("Error during key decryption")
		}

		blockNum, txHash := createTransferResponse()

		mockContractManager.EXPECT().CreateWPCSafeTransfer(ctx, gomock.Any(), key.PrivateKey, true).
			Return(blockNum, txHash, nil)

		response, err := tm.CreateWPCSafeTransferAsync(ctx, transfer)

		assert.NoError(t, err, "Unexpected error")
		assert.Equal(t, mapTransferResponse(blockNum, txHash), response, "Response differs from expected")
	})

	t.Run("CreateWPCSafeTransferAsyncNegative_InvalidWallet", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareCreateWPCSafeTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(sourcePassphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(invalidWallet, nil)

		response, err := tm.CreateWPCSafeTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrFailedToParseKey.Error(), "Expected InvalidWallet error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("CreateWPCSafeTransferAsyncNegative_WrongPassphrase", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareCreateWPCSafeTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(wrongPass, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(sourceWallet, nil)

		response, err := tm.CreateWPCSafeTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrWrongPassphrase.Error(), "Expected WrongPassphrase error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("CreateWPCSafeTransferAsyncNegative_TransferExists", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareCreateWPCSafeTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(sourcePassphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(sourceWallet, nil)
		mockContractManager.EXPECT().GetAssetboxBalance(ctx, transfer.From).Return(getSufficientBalance(), nil)
		mockContractManager.EXPECT().TransferExists(ctx, []byte(transfer.TransferID)).Return(true, nil)

		response, err := tm.CreateWPCSafeTransferAsync(ctx, transfer)

		assert.EqualError(t, err, errors.Errorf("transfer with id %q already exists", transfer.TransferID).
			Error(), "Expected TransferExists error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("CreateWPCSafeTransferAsyncNegative_InsufficientBalance", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareCreateWPCSafeTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(sourcePassphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(sourceWallet, nil)
		mockContractManager.EXPECT().GetAssetboxBalance(ctx, transfer.From).Return(getInsufficientBalance(), nil)

		response, err := tm.CreateWPCSafeTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrLowAssetboxBalance.Error(), "Expected LowAssetboxBalance error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("CreateWPCSafeTransferAsyncNegative_NilAmount", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareCreateWPCSafeTransferObject()
		transfer.Value = nil

		response, err := tm.CreateWPCSafeTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrValueRequire.Error(), "Expected RequiredValue error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("CreateWPCSafeTransferAsyncNegative_ZeroAmount", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareCreateWPCSafeTransferObject()
		transfer.Value = big.NewInt(0)

		response, err := tm.CreateWPCSafeTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrValueRequire.Error(), "Expected RequiredValue error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("CreateWPCSafeTransferAsyncNegative_NegativeAmount", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareCreateWPCSafeTransferObject()
		transfer.Value = big.NewInt(-10)

		response, err := tm.CreateWPCSafeTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrValueRequire.Error(), "Expected RequiredValue error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("CreateWPCSafeTransferAsyncNegative_AlreadyExpired", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareCreateWPCSafeTransferObject()
		transfer.ExpiresAt = time.Now().Unix()

		response, err := tm.CreateWPCSafeTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrExpireAtPassed.Error(), "Expected AlreadyExpired error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("CreateWPCSafeTransferAsyncNegative_MissingTransferID", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareCreateWPCSafeTransferObject()
		transfer.TransferID = ""

		response, err := tm.CreateWPCSafeTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrTransferIDRequire.Error(), "Expected TransferIDRequired error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("CreateWPCSafeTransferAsyncNegative_FromAndToAreIdentical", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareCreateWPCSafeTransferObject()
		transfer.To = transfer.From

		response, err := tm.CreateWPCSafeTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrSameAccount.Error(), "Expected SameAccount error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("CreateWPCSafeTransferAsyncNegative_MissingTo", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareCreateWPCSafeTransferObject()
		transfer.To = common.Address{}

		response, err := tm.CreateWPCSafeTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrToRequire.Error(), "Expected ToRequired error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("CreateWPCSafeTransferAsyncNegative_MissingFrom", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareCreateWPCSafeTransferObject()
		transfer.From = common.Address{}

		response, err := tm.CreateWPCSafeTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrFromRequire.Error(), "Expected FromRequired error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("CreateWPCSafeTransferAsyncNegative_ManagerIsNotReady", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)

		transfer := prepareCreateWPCSafeTransferObject()

		response, err := tm.CreateWPCSafeTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrTransferNotReady.Error(), "Expected ManagerIsNotReady error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("CreateWPCSafeTransferAsyncNegative_AssetboxIsNotOwnerOfWallet", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareCreateWPCSafeTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(destPassphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(destWallet, nil)

		response, err := tm.CreateWPCSafeTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrAssetboxIsNotOwnerOfWallet.Error(), "Expected AssetboxIsNotOwnerOfWallet error")
		assert.Empty(t, response, "Response wasn't empty")
	})
}

func TestCreateFullBalanceWPCSafeTransferAsync(t *testing.T) {
	ctx := context.TODO()

	mockController := gomock.NewController(t)
	defer mockController.Finish()

	t.Run("CreateFullBalanceWPCSafeTransferAsyncPositive", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareCreateWPCSafeTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(sourcePassphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(sourceWallet, nil)
		mockContractManager.EXPECT().GetAssetboxBalance(ctx, transfer.From).Return(getSufficientBalance(), nil)
		mockContractManager.EXPECT().TransferExists(ctx, []byte(transfer.TransferID)).Return(false, nil)
		mockContractManager.EXPECT().BalanceOfLocked(ctx, transfer.From).Return(big.NewInt(0), nil)

		key, err := keystore.DecryptKey(sourceWallet, string(sourcePassphrase))
		if err != nil {
			panic("Error during key decryption")
		}

		blockNum, txHash := createTransferResponse()

		mockContractManager.EXPECT().CreateFullBalanceWPCSafeTransfer(ctx, gomock.Any(), key.PrivateKey, true).
			Return(blockNum, txHash, nil)

		response, err := tm.CreateFullBalanceWPCSafeTransferAsync(ctx, transfer)

		assert.NoError(t, err, "Unexpected error")
		assert.Equal(t, mapTransferResponse(blockNum, txHash), response, "Response differs from expected")
	})

	t.Run("CreateFullBalanceWPCSafeTransferAsyncNegative_InvalidWallet", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareCreateWPCSafeTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(sourcePassphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(invalidWallet, nil)

		response, err := tm.CreateFullBalanceWPCSafeTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrFailedToParseKey.Error(), "Expected InvalidWallet error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("CreateFullBalanceWPCSafeTransferAsyncNegative_WrongPassphrase", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareCreateWPCSafeTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(wrongPass, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(sourceWallet, nil)

		response, err := tm.CreateFullBalanceWPCSafeTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrWrongPassphrase.Error(), "Expected WrongPassphrase error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("CreateFullBalanceWPCSafeTransferAsyncNegative_TransferExists", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareCreateWPCSafeTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(sourcePassphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(sourceWallet, nil)
		mockContractManager.EXPECT().GetAssetboxBalance(ctx, transfer.From).Return(getSufficientBalance(), nil)
		mockContractManager.EXPECT().TransferExists(ctx, []byte(transfer.TransferID)).Return(true, nil)
		mockContractManager.EXPECT().BalanceOfLocked(ctx, transfer.From).Return(big.NewInt(0), nil)

		response, err := tm.CreateFullBalanceWPCSafeTransferAsync(ctx, transfer)

		assert.EqualError(t, err, errors.Errorf("transfer with id %q already exists", transfer.TransferID).
			Error(), "Expected TransferExists error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("CreateFullBalanceWPCSafeTransferAsyncNegative_ZeroBalance", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareCreateWPCSafeTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(sourcePassphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(sourceWallet, nil)
		mockContractManager.EXPECT().GetAssetboxBalance(ctx, transfer.From).Return(big.NewInt(0), nil)

		response, err := tm.CreateFullBalanceWPCSafeTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrLowAssetboxBalance.Error(), "Expected LowAssetboxBalance error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("CreateFullBalanceWPCSafeTransferAsyncNegative_AlreadyExpired", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareCreateWPCSafeTransferObject()
		transfer.ExpiresAt = time.Now().Unix()

		response, err := tm.CreateFullBalanceWPCSafeTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrExpireAtPassed.Error(), "Expected AlreadyExpired error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("CreateFullBalanceWPCSafeTransferAsyncNegative_MissingTransferID", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareCreateWPCSafeTransferObject()
		transfer.TransferID = ""

		response, err := tm.CreateFullBalanceWPCSafeTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrTransferIDRequire.Error(), "Expected TransferIDRequired error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("CreateFullBalanceWPCSafeTransferAsyncNegative_FromAndToAreIdentical", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareCreateWPCSafeTransferObject()
		transfer.To = transfer.From

		response, err := tm.CreateFullBalanceWPCSafeTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrSameAccount.Error(), "Expected SameAccount error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("CreateFullBalanceWPCSafeTransferAsyncNegative_MissingTo", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareCreateWPCSafeTransferObject()
		transfer.To = common.Address{}

		response, err := tm.CreateFullBalanceWPCSafeTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrToRequire.Error(), "Expected ToRequired error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("CreateFullBalanceWPCSafeTransferAsyncNegative_MissingFrom", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareCreateWPCSafeTransferObject()
		transfer.From = common.Address{}

		response, err := tm.CreateFullBalanceWPCSafeTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrFromRequire.Error(), "Expected FromRequired error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("CreateFullBalanceWPCSafeTransferAsyncNegative_ManagerIsNotReady", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)

		transfer := prepareCreateWPCSafeTransferObject()

		response, err := tm.CreateFullBalanceWPCSafeTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrTransferNotReady.Error(), "Expected ManagerIsNotReady error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("CreateFullBalanceWPCSafeTransferAsyncNegative_AssetboxIsNotOwnerOfWallet", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareCreateWPCSafeTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(destPassphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(destWallet, nil)

		response, err := tm.CreateFullBalanceWPCSafeTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrAssetboxIsNotOwnerOfWallet.Error(), "Expected AssetboxIsNotOwnerOfWallet error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("CreateFullBalanceWPCSafeTransferAsyncNegative_LockedBalance", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareCreateWPCSafeTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(sourcePassphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(sourceWallet, nil)
		mockContractManager.EXPECT().GetAssetboxBalance(ctx, transfer.From).Return(getSufficientBalance(), nil)
		mockContractManager.EXPECT().BalanceOfLocked(ctx, transfer.From).Return(big.NewInt(1), nil)

		response, err := tm.CreateFullBalanceWPCSafeTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrHasLockedBalanceError.Error(), "Expected HasLockedBalanceError error")
		assert.Empty(t, response, "Response wasn't empty")
	})
}
