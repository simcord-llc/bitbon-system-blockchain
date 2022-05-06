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

func TestApproveSafeTransfer(t *testing.T) {
	ctx := context.TODO()

	mockController := gomock.NewController(t)
	defer mockController.Finish()

	t.Run("ApproveSafeTransferPositive", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareApproveSafeTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(destPassphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(destWallet, nil)
		mockContractManager.EXPECT().GetTransferState(ctx, []byte(transfer.TransferID)).
			Return(uint8(1), nil)

		protectionHash := tm.getProtectionHash(transfer.TransferID, transfer.ProtectionCode)

		key, err := keystore.DecryptKey(destWallet, string(destPassphrase))
		if err != nil {
			panic("Error during key decryption")
		}

		blockNum, txHash := createTransferResponse()

		mockContractManager.EXPECT().
			ApproveSafeTransfer(ctx, []byte(transfer.TransferID), protectionHash, nil, key.PrivateKey, false).
			Return(blockNum, txHash, nil)

		response, err := tm.ApproveSafeTransfer(ctx, transfer)

		assert.NoError(t, err, "Unexpected error")
		assert.Equal(t, mapTransferResponse(blockNum, txHash), response, "Response differs from expected")
	})

	t.Run("ApproveSafeTransferNegative_InvalidWallet", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareApproveSafeTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(destPassphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(invalidWallet, nil)

		response, err := tm.ApproveSafeTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrFailedToParseKey.Error(), "Expected InvalidWallet error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("ApproveSafeTransferNegative_WrongPassphrase", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareApproveSafeTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(wrongPass, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(destWallet, nil)

		response, err := tm.ApproveSafeTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrWrongPassphrase.Error(), "Expected WrongPassphrase error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("ApproveSafeTransferNegative_TransferHasExpired", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareApproveSafeTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(destPassphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(destWallet, nil)
		mockContractManager.EXPECT().GetTransferState(ctx, []byte(transfer.TransferID)).
			Return(uint8(3), nil)

		response, err := tm.ApproveSafeTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrTransferIsNotPending.Error())
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("ApproveSafeTransferNegative_WrongProtectionCodeLimitExceeded", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareApproveSafeTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(destPassphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(destWallet, nil)
		mockContractManager.EXPECT().GetTransferState(ctx, []byte(transfer.TransferID)).
			Return(uint8(1), nil)

		protectionHash := tm.getProtectionHash(transfer.TransferID, transfer.ProtectionCode)

		key, err := keystore.DecryptKey(destWallet, string(destPassphrase))
		if err != nil {
			panic("Error during key decryption")
		}

		mockContractManager.EXPECT().
			ApproveSafeTransfer(ctx, []byte(transfer.TransferID), protectionHash, nil, key.PrivateKey, false).
			Return(uint64(0), [32]byte{}, ErrProtectionCodeLimitExceeded)

		response, err := tm.ApproveSafeTransfer(ctx, transfer)

		assert.EqualError(t, err, errors.Wrap(ErrProtectionCodeLimitExceeded,
			"error approving safe transfer in blockchain").Error(), "Expected TransferHasExpired error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("ApproveSafeTransferNegative_WrongProtectionCode", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareApproveSafeTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(destPassphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(destWallet, nil)
		mockContractManager.EXPECT().GetTransferState(ctx, []byte(transfer.TransferID)).
			Return(uint8(1), nil)

		protectionHash := tm.getProtectionHash(transfer.TransferID, transfer.ProtectionCode)

		key, err := keystore.DecryptKey(destWallet, string(destPassphrase))
		if err != nil {
			panic("Error during key decryption")
		}

		mockContractManager.EXPECT().
			ApproveSafeTransfer(ctx, []byte(transfer.TransferID), protectionHash, nil, key.PrivateKey, false).
			Return(uint64(0), [32]byte{}, ErrWrongProtectionCode)

		response, err := tm.ApproveSafeTransfer(ctx, transfer)

		assert.EqualError(t, err, errors.Wrap(ErrWrongProtectionCode,
			"error approving safe transfer in blockchain").Error(), "Expected TransferHasExpired error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("ApproveSafeTransferNegative_TransferMissing", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareApproveSafeTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(destPassphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(destWallet, nil)
		mockContractManager.EXPECT().GetTransferState(ctx, []byte(transfer.TransferID)).
			Return(uint8(0), nil)

		response, err := tm.ApproveSafeTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrTransferIsNotPending.Error())
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("ApproveSafeTransferNegative_MissingProtectionCode", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareApproveSafeTransferObject()
		transfer.ProtectionCode = ""

		response, err := tm.ApproveSafeTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrProtectionCodeRequire.Error(), "Expected ProtectionCodeRequired error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("ApproveSafeTransferNegative_MissingTransferID", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareApproveSafeTransferObject()
		transfer.TransferID = ""

		response, err := tm.ApproveSafeTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrTransferIDRequire.Error(), "Expected TransferIDRequired error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("ApproveSafeTransferNegative_MissingAddress", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareApproveSafeTransferObject()
		transfer.Address = common.Address{}

		response, err := tm.ApproveSafeTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrAddressRequire.Error(), "Expected AddressRequired error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("ApproveSafeTransferNegative_ManagerIsNotReady", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)

		transfer := prepareApproveSafeTransferObject()

		response, err := tm.ApproveSafeTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrTransferNotReady.Error(), "Expected ManagerIsNotReady error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("ApproveSafeTransferNegative_TransferMissing", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareApproveSafeTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(destPassphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(destWallet, nil)
		mockContractManager.EXPECT().GetTransferState(ctx, []byte(transfer.TransferID)).
			Return(uint8(2), nil)

		response, err := tm.ApproveSafeTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrTransferIsNotPending.Error())
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("ApproveSafeTransferNegative_AssetboxIsNotOwnerOfWallet", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareApproveSafeTransferObject()
		transfer.Address = sourceAssetbox

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(destPassphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(destWallet, nil)

		response, err := tm.ApproveSafeTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrAssetboxIsNotOwnerOfWallet.Error(), "Expected AssetboxIsNotOwnerOfWallet error")
		assert.Empty(t, response, "Response wasn't empty")
	})
}

func TestApproveFullBalanceSafeTransfer(t *testing.T) {
	ctx := context.TODO()

	mockController := gomock.NewController(t)
	defer mockController.Finish()

	t.Run("ApproveFullBalanceSafeTransferPositive", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareApproveSafeTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(destPassphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(destWallet, nil)
		mockContractManager.EXPECT().GetTransferState(ctx, []byte(transfer.TransferID)).
			Return(uint8(1), nil)

		protectionHash := tm.getProtectionHash(transfer.TransferID, transfer.ProtectionCode)

		key, err := keystore.DecryptKey(destWallet, string(destPassphrase))
		if err != nil {
			panic("Error during key decryption")
		}

		blockNum, txHash := createTransferResponse()

		mockContractManager.EXPECT().
			ApproveFullBalanceSafeTransfer(ctx, []byte(transfer.TransferID), protectionHash, nil, key.PrivateKey, false).
			Return(blockNum, txHash, nil)

		response, err := tm.ApproveFullBalanceSafeTransfer(ctx, transfer)

		assert.NoError(t, err, "Unexpected error")
		assert.Equal(t, mapTransferResponse(blockNum, txHash), response, "Response differs from expected")
	})

	t.Run("ApproveFullBalanceSafeTransferNegative_InvalidWallet", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareApproveSafeTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(destPassphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(invalidWallet, nil)

		response, err := tm.ApproveFullBalanceSafeTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrFailedToParseKey.Error(), "Expected InvalidWallet error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("ApproveFullBalanceSafeTransferNegative_WrongPassphrase", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareApproveSafeTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(wrongPass, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(destWallet, nil)

		response, err := tm.ApproveFullBalanceSafeTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrWrongPassphrase.Error(), "Expected WrongPassphrase error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("ApproveFullBalanceSafeTransferNegative_TransferHasExpired", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareApproveSafeTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(destPassphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(destWallet, nil)
		mockContractManager.EXPECT().GetTransferState(ctx, []byte(transfer.TransferID)).
			Return(uint8(3), nil)

		response, err := tm.ApproveFullBalanceSafeTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrTransferIsNotPending.Error())
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("ApproveFullBalanceSafeTransferNegative_WrongProtectionCodeLimitExceeded", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareApproveSafeTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(destPassphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(destWallet, nil)
		mockContractManager.EXPECT().GetTransferState(ctx, []byte(transfer.TransferID)).
			Return(uint8(1), nil)

		protectionHash := tm.getProtectionHash(transfer.TransferID, transfer.ProtectionCode)

		key, err := keystore.DecryptKey(destWallet, string(destPassphrase))
		if err != nil {
			panic("Error during key decryption")
		}

		mockContractManager.EXPECT().
			ApproveFullBalanceSafeTransfer(ctx, []byte(transfer.TransferID), protectionHash, nil, key.PrivateKey, false).
			Return(uint64(0), [32]byte{}, ErrProtectionCodeLimitExceeded)

		response, err := tm.ApproveFullBalanceSafeTransfer(ctx, transfer)

		assert.EqualError(t, err, errors.Wrap(ErrProtectionCodeLimitExceeded,
			"error approving full balance safe transfer in blockchain").Error(), "Expected TransferHasExpired error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("ApproveFullBalanceSafeTransferNegative_WrongProtectionCode", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareApproveSafeTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(destPassphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(destWallet, nil)
		mockContractManager.EXPECT().GetTransferState(ctx, []byte(transfer.TransferID)).
			Return(uint8(1), nil)

		protectionHash := tm.getProtectionHash(transfer.TransferID, transfer.ProtectionCode)

		key, err := keystore.DecryptKey(destWallet, string(destPassphrase))
		if err != nil {
			panic("Error during key decryption")
		}

		mockContractManager.EXPECT().
			ApproveFullBalanceSafeTransfer(ctx, []byte(transfer.TransferID), protectionHash, nil, key.PrivateKey, false).
			Return(uint64(0), [32]byte{}, ErrWrongProtectionCode)

		response, err := tm.ApproveFullBalanceSafeTransfer(ctx, transfer)

		assert.EqualError(t, err, errors.Wrap(ErrWrongProtectionCode,
			"error approving full balance safe transfer in blockchain").Error(), "Expected TransferHasExpired error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("ApproveFullBalanceSafeTransferNegative_TransferMissing", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareApproveSafeTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(destPassphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(destWallet, nil)
		mockContractManager.EXPECT().GetTransferState(ctx, []byte(transfer.TransferID)).
			Return(uint8(0), nil)

		response, err := tm.ApproveFullBalanceSafeTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrTransferIsNotPending.Error())
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("ApproveFullBalanceSafeTransferNegative_MissingProtectionCode", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareApproveSafeTransferObject()
		transfer.ProtectionCode = ""

		response, err := tm.ApproveFullBalanceSafeTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrProtectionCodeRequire.Error(), "Expected ProtectionCodeRequired error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("ApproveFullBalanceSafeTransferNegative_MissingTransferID", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareApproveSafeTransferObject()
		transfer.TransferID = ""

		response, err := tm.ApproveFullBalanceSafeTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrTransferIDRequire.Error(), "Expected TransferIDRequired error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("ApproveFullBalanceSafeTransferNegative_MissingAddress", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareApproveSafeTransferObject()
		transfer.Address = common.Address{}

		response, err := tm.ApproveFullBalanceSafeTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrAddressRequire.Error(), "Expected AddressRequired error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("ApproveFullBalanceSafeTransferNegative_ManagerIsNotReady", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)

		transfer := prepareApproveSafeTransferObject()

		response, err := tm.ApproveFullBalanceSafeTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrTransferNotReady.Error(), "Expected ManagerIsNotReady error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("ApproveFullBalanceSafeTransferNegative_TransferMissing", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareApproveSafeTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(destPassphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(destWallet, nil)
		mockContractManager.EXPECT().GetTransferState(ctx, []byte(transfer.TransferID)).
			Return(uint8(2), nil)

		response, err := tm.ApproveFullBalanceSafeTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrTransferIsNotPending.Error())
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("ApproveFullBalanceSafeTransferNegative_AssetboxIsNotOwnerOfWallet", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareApproveSafeTransferObject()
		transfer.Address = sourceAssetbox

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(destPassphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(destWallet, nil)

		response, err := tm.ApproveFullBalanceSafeTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrAssetboxIsNotOwnerOfWallet.Error(), "Expected AssetboxIsNotOwnerOfWallet error")
		assert.Empty(t, response, "Response wasn't empty")
	})
}

func TestApproveSafeTransferAsync(t *testing.T) {
	ctx := context.TODO()

	mockController := gomock.NewController(t)
	defer mockController.Finish()

	t.Run("ApproveSafeTransferAsyncPositive", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareApproveSafeTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(destPassphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(destWallet, nil)
		mockContractManager.EXPECT().GetTransferState(ctx, []byte(transfer.TransferID)).
			Return(uint8(1), nil)

		protectionHash := tm.getProtectionHash(transfer.TransferID, transfer.ProtectionCode)

		key, err := keystore.DecryptKey(destWallet, string(destPassphrase))
		if err != nil {
			panic("Error during key decryption")
		}

		blockNum, txHash := createTransferResponse()

		mockContractManager.EXPECT().
			ApproveSafeTransfer(ctx, []byte(transfer.TransferID), protectionHash, nil, key.PrivateKey, true).
			Return(blockNum, txHash, nil)

		response, err := tm.ApproveSafeTransferAsync(ctx, transfer)

		assert.NoError(t, err, "Unexpected error")
		assert.Equal(t, mapTransferResponse(blockNum, txHash), response, "Response differs from expected")
	})

	t.Run("ApproveSafeTransferAsyncNegative_InvalidWallet", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareApproveSafeTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(destPassphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(invalidWallet, nil)

		response, err := tm.ApproveSafeTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrFailedToParseKey.Error(), "Expected InvalidWallet error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("ApproveSafeTransferAsyncNegative_WrongPassphrase", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareApproveSafeTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(wrongPass, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(destWallet, nil)

		response, err := tm.ApproveSafeTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrWrongPassphrase.Error(), "Expected WrongPassphrase error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("ApproveSafeTransferAsyncNegative_TransferHasExpired", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareApproveSafeTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(destPassphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(destWallet, nil)
		mockContractManager.EXPECT().GetTransferState(ctx, []byte(transfer.TransferID)).
			Return(uint8(3), nil)

		response, err := tm.ApproveSafeTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrTransferIsNotPending.Error())
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("ApproveSafeTransferAsyncNegative_WrongProtectionCodeLimitExceeded", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareApproveSafeTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(destPassphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(destWallet, nil)
		mockContractManager.EXPECT().GetTransferState(ctx, []byte(transfer.TransferID)).
			Return(uint8(1), nil)

		protectionHash := tm.getProtectionHash(transfer.TransferID, transfer.ProtectionCode)

		key, err := keystore.DecryptKey(destWallet, string(destPassphrase))
		if err != nil {
			panic("Error during key decryption")
		}

		mockContractManager.EXPECT().
			ApproveSafeTransfer(ctx, []byte(transfer.TransferID), protectionHash, nil, key.PrivateKey, true).
			Return(uint64(0), [32]byte{}, ErrProtectionCodeLimitExceeded)

		response, err := tm.ApproveSafeTransferAsync(ctx, transfer)

		assert.EqualError(t, err, errors.Wrap(ErrProtectionCodeLimitExceeded,
			"error approving safe transfer in blockchain").Error(), "Expected TransferHasExpired error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("ApproveSafeTransferAsyncNegative_WrongProtectionCode", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareApproveSafeTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(destPassphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(destWallet, nil)
		mockContractManager.EXPECT().GetTransferState(ctx, []byte(transfer.TransferID)).
			Return(uint8(1), nil)

		protectionHash := tm.getProtectionHash(transfer.TransferID, transfer.ProtectionCode)

		key, err := keystore.DecryptKey(destWallet, string(destPassphrase))
		if err != nil {
			panic("Error during key decryption")
		}

		mockContractManager.EXPECT().
			ApproveSafeTransfer(ctx, []byte(transfer.TransferID), protectionHash, nil, key.PrivateKey, true).
			Return(uint64(0), [32]byte{}, ErrWrongProtectionCode)

		response, err := tm.ApproveSafeTransferAsync(ctx, transfer)

		assert.EqualError(t, err, errors.Wrap(ErrWrongProtectionCode,
			"error approving safe transfer in blockchain").Error(), "Expected TransferHasExpired error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("ApproveSafeTransferAsyncNegative_TransferMissing", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareApproveSafeTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(destPassphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(destWallet, nil)
		mockContractManager.EXPECT().GetTransferState(ctx, []byte(transfer.TransferID)).
			Return(uint8(0), nil)

		response, err := tm.ApproveSafeTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrTransferIsNotPending.Error())
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("ApproveSafeTransferAsyncNegative_MissingProtectionCode", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareApproveSafeTransferObject()
		transfer.ProtectionCode = ""

		response, err := tm.ApproveSafeTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrProtectionCodeRequire.Error(), "Expected ProtectionCodeRequired error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("ApproveSafeTransferAsyncNegative_MissingTransferID", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareApproveSafeTransferObject()
		transfer.TransferID = ""

		response, err := tm.ApproveSafeTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrTransferIDRequire.Error(), "Expected TransferIDRequired error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("ApproveSafeTransferAsyncNegative_MissingAddress", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareApproveSafeTransferObject()
		transfer.Address = common.Address{}

		response, err := tm.ApproveSafeTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrAddressRequire.Error(), "Expected AddressRequired error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("ApproveSafeTransferAsyncNegative_ManagerIsNotReady", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)

		transfer := prepareApproveSafeTransferObject()

		response, err := tm.ApproveSafeTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrTransferNotReady.Error(), "Expected ManagerIsNotReady error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("ApproveSafeTransferAsyncNegative_TransferMissing", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareApproveSafeTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(destPassphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(destWallet, nil)
		mockContractManager.EXPECT().GetTransferState(ctx, []byte(transfer.TransferID)).
			Return(uint8(2), nil)

		response, err := tm.ApproveSafeTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrTransferIsNotPending.Error())
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("ApproveSafeTransferAsyncNegative_AssetboxIsNotOwnerOfWallet", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareApproveSafeTransferObject()
		transfer.Address = sourceAssetbox

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(destPassphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(destWallet, nil)

		response, err := tm.ApproveSafeTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrAssetboxIsNotOwnerOfWallet.Error(), "Expected AssetboxIsNotOwnerOfWallet error")
		assert.Empty(t, response, "Response wasn't empty")
	})
}

func TestApproveFullBalanceSafeTransferAsync(t *testing.T) {
	ctx := context.TODO()

	mockController := gomock.NewController(t)
	defer mockController.Finish()

	t.Run("ApproveFullBalanceSafeTransferAsyncPositive", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareApproveSafeTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(destPassphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(destWallet, nil)
		mockContractManager.EXPECT().GetTransferState(ctx, []byte(transfer.TransferID)).
			Return(uint8(1), nil)

		protectionHash := tm.getProtectionHash(transfer.TransferID, transfer.ProtectionCode)

		key, err := keystore.DecryptKey(destWallet, string(destPassphrase))
		if err != nil {
			panic("Error during key decryption")
		}

		blockNum, txHash := createTransferResponse()

		mockContractManager.EXPECT().
			ApproveFullBalanceSafeTransfer(ctx, []byte(transfer.TransferID), protectionHash, nil, key.PrivateKey, true).
			Return(blockNum, txHash, nil)

		response, err := tm.ApproveFullBalanceSafeTransferAsync(ctx, transfer)

		assert.NoError(t, err, "Unexpected error")
		assert.Equal(t, mapTransferResponse(blockNum, txHash), response, "Response differs from expected")
	})

	t.Run("ApproveFullBalanceSafeTransferAsyncNegative_InvalidWallet", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareApproveSafeTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(destPassphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(invalidWallet, nil)

		response, err := tm.ApproveFullBalanceSafeTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrFailedToParseKey.Error(), "Expected InvalidWallet error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("ApproveFullBalanceSafeTransferAsyncNegative_WrongPassphrase", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareApproveSafeTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(wrongPass, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(destWallet, nil)

		response, err := tm.ApproveFullBalanceSafeTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrWrongPassphrase.Error(), "Expected WrongPassphrase error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("ApproveFullBalanceSafeTransferAsyncNegative_TransferHasExpired", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareApproveSafeTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(destPassphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(destWallet, nil)
		mockContractManager.EXPECT().GetTransferState(ctx, []byte(transfer.TransferID)).
			Return(uint8(3), nil)

		response, err := tm.ApproveFullBalanceSafeTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrTransferIsNotPending.Error())
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("ApproveFullBalanceSafeTransferAsyncNegative_WrongProtectionCodeLimitExceeded", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareApproveSafeTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(destPassphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(destWallet, nil)
		mockContractManager.EXPECT().GetTransferState(ctx, []byte(transfer.TransferID)).
			Return(uint8(1), nil)

		protectionHash := tm.getProtectionHash(transfer.TransferID, transfer.ProtectionCode)

		key, err := keystore.DecryptKey(destWallet, string(destPassphrase))
		if err != nil {
			panic("Error during key decryption")
		}

		mockContractManager.EXPECT().
			ApproveFullBalanceSafeTransfer(ctx, []byte(transfer.TransferID), protectionHash, nil, key.PrivateKey, true).
			Return(uint64(0), [32]byte{}, ErrProtectionCodeLimitExceeded)

		response, err := tm.ApproveFullBalanceSafeTransferAsync(ctx, transfer)

		assert.EqualError(t, err, errors.Wrap(ErrProtectionCodeLimitExceeded,
			"error approving full balance safe transfer in blockchain").Error(), "Expected TransferHasExpired error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("ApproveFullBalanceSafeTransferAsyncNegative_WrongProtectionCode", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareApproveSafeTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(destPassphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(destWallet, nil)
		mockContractManager.EXPECT().GetTransferState(ctx, []byte(transfer.TransferID)).
			Return(uint8(1), nil)

		protectionHash := tm.getProtectionHash(transfer.TransferID, transfer.ProtectionCode)

		key, err := keystore.DecryptKey(destWallet, string(destPassphrase))
		if err != nil {
			panic("Error during key decryption")
		}

		mockContractManager.EXPECT().
			ApproveFullBalanceSafeTransfer(ctx, []byte(transfer.TransferID), protectionHash, nil, key.PrivateKey, true).
			Return(uint64(0), [32]byte{}, ErrWrongProtectionCode)

		response, err := tm.ApproveFullBalanceSafeTransferAsync(ctx, transfer)

		assert.EqualError(t, err, errors.Wrap(ErrWrongProtectionCode,
			"error approving full balance safe transfer in blockchain").Error(), "Expected TransferHasExpired error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("ApproveFullBalanceSafeTransferAsyncNegative_TransferMissing", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareApproveSafeTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(destPassphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(destWallet, nil)
		mockContractManager.EXPECT().GetTransferState(ctx, []byte(transfer.TransferID)).
			Return(uint8(0), nil)

		response, err := tm.ApproveFullBalanceSafeTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrTransferIsNotPending.Error())
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("ApproveFullBalanceSafeTransferAsyncNegative_MissingProtectionCode", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareApproveSafeTransferObject()
		transfer.ProtectionCode = ""

		response, err := tm.ApproveFullBalanceSafeTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrProtectionCodeRequire.Error(), "Expected ProtectionCodeRequired error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("ApproveFullBalanceSafeTransferAsyncNegative_MissingTransferID", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareApproveSafeTransferObject()
		transfer.TransferID = ""

		response, err := tm.ApproveFullBalanceSafeTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrTransferIDRequire.Error(), "Expected TransferIDRequired error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("ApproveFullBalanceSafeTransferAsyncNegative_MissingAddress", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareApproveSafeTransferObject()
		transfer.Address = common.Address{}

		response, err := tm.ApproveFullBalanceSafeTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrAddressRequire.Error(), "Expected AddressRequired error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("ApproveFullBalanceSafeTransferAsyncNegative_ManagerIsNotReady", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)

		transfer := prepareApproveSafeTransferObject()

		response, err := tm.ApproveFullBalanceSafeTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrTransferNotReady.Error(), "Expected ManagerIsNotReady error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("ApproveFullBalanceSafeTransferAsyncNegative_TransferMissing", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareApproveSafeTransferObject()

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(destPassphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(destWallet, nil)
		mockContractManager.EXPECT().GetTransferState(ctx, []byte(transfer.TransferID)).
			Return(uint8(2), nil)

		response, err := tm.ApproveFullBalanceSafeTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrTransferIsNotPending.Error())
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("ApproveFullBalanceSafeTransferAsyncNegative_AssetboxIsNotOwnerOfWallet", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, mockEncryptor := prepareMockServices(mockController)
		tm := prepareSafeTransferManager(mockAssetboxManager, mockContractManager, mockEncryptor)
		tm.obtainKeys()

		transfer := prepareApproveSafeTransferObject()
		transfer.Address = sourceAssetbox

		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Passphrase, encryptionKey).Return(destPassphrase, nil)
		mockEncryptor.EXPECT().Decrypt(transfer.CryptoData.Wallet, encryptionKey).Return(destWallet, nil)

		response, err := tm.ApproveFullBalanceSafeTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrAssetboxIsNotOwnerOfWallet.Error(), "Expected AssetboxIsNotOwnerOfWallet error")
		assert.Empty(t, response, "Response wasn't empty")
	})
}
