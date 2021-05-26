package transfer

import (
	"context"
	"math/big"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/simcord-llc/bitbon-system-blockchain/common"
	"github.com/stretchr/testify/assert"
)

func TestDirectTransfer(t *testing.T) {
	ctx := context.TODO()

	mockController := gomock.NewController(t)
	defer mockController.Finish()

	t.Run("DirectTransferPositive", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, _ := prepareMockServices(mockController)
		tm := prepareDirectTransferManager(mockAssetboxManager, mockContractManager)

		transfer := prepareDirectTransferObject()

		mockContractManager.EXPECT().GetAssetboxBalance(ctx, transfer.From).
			Return(getSufficientBalance(), nil)

		blockNum, txHash := createTransferResponse()

		mockContractManager.EXPECT().
			DirectTransfer(ctx, transfer.From, transfer.To, transfer.Value,
				transfer.ExtraData, tm.bitbon.GetServicePk(), false).
			Return(blockNum, txHash, nil)

		response, err := tm.DirectTransfer(ctx, transfer)

		assert.NoError(t, err, "Unexpected error")
		assert.Equal(t, mapTransferResponse(blockNum, txHash), response, "Response differs from expected")
	})

	t.Run("DirectTransferPositive_DestinationIsServiceAccount", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, _ := prepareMockServices(mockController)
		tm := prepareDirectTransferManager(mockAssetboxManager, mockContractManager)

		transfer := prepareDirectTransferObject()
		transfer.To = generateServiceAssetboxAddress()

		mockContractManager.EXPECT().GetAssetboxBalance(ctx, transfer.From).
			Return(getSufficientBalance(), nil)

		blockNum, txHash := createTransferResponse()

		mockContractManager.EXPECT().
			DirectTransfer(ctx, transfer.From, transfer.To, transfer.Value,
				transfer.ExtraData, tm.bitbon.GetServicePk(), false).
			Return(blockNum, txHash, nil)

		response, err := tm.DirectTransfer(ctx, transfer)

		assert.NoError(t, err, "Unexpected error")
		assert.Equal(t, mapTransferResponse(blockNum, txHash), response, "Response differs from expected")
	})

	t.Run("DirectTransferPositive_SourceIsServiceAccount", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, _ := prepareMockServices(mockController)
		tm := prepareDirectTransferManager(mockAssetboxManager, mockContractManager)

		transfer := prepareDirectTransferObject()
		transfer.From = generateServiceAssetboxAddress()

		mockContractManager.EXPECT().GetAssetboxBalance(ctx, transfer.From).
			Return(getSufficientBalance(), nil)

		blockNum, txHash := createTransferResponse()

		mockContractManager.EXPECT().
			DirectTransfer(ctx, transfer.From, transfer.To, transfer.Value,
				transfer.ExtraData, tm.bitbon.GetServicePk(), false).
			Return(blockNum, txHash, nil)

		response, err := tm.DirectTransfer(ctx, transfer)

		assert.NoError(t, err, "Unexpected error")
		assert.Equal(t, mapTransferResponse(blockNum, txHash), response, "Response differs from expected")
	})

	t.Run("DirectTransferNegative_InsufficientBalance", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, _ := prepareMockServices(mockController)
		tm := prepareDirectTransferManager(mockAssetboxManager, mockContractManager)

		transfer := prepareDirectTransferObject()

		mockContractManager.EXPECT().GetAssetboxBalance(ctx, transfer.From).
			Return(getInsufficientBalance(), nil)

		response, err := tm.DirectTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrLowAssetboxBalance.Error(), "Expected LowAssetboxBalance error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("DirectTransferNegative_NilAmount", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, _ := prepareMockServices(mockController)
		tm := prepareDirectTransferManager(mockAssetboxManager, mockContractManager)

		transfer := prepareDirectTransferObject()
		transfer.Value = nil

		response, err := tm.DirectTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrValueRequire.Error(), "Expected RequiredValue error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("DirectTransferNegative_NegativeAmount", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, _ := prepareMockServices(mockController)
		tm := prepareDirectTransferManager(mockAssetboxManager, mockContractManager)

		transfer := prepareDirectTransferObject()
		transfer.Value = big.NewInt(-10)

		response, err := tm.DirectTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrValueRequire.Error(), "Expected RequiredValue error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("DirectTransferNegative_ZeroAmount", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, _ := prepareMockServices(mockController)
		tm := prepareDirectTransferManager(mockAssetboxManager, mockContractManager)

		transfer := prepareDirectTransferObject()
		transfer.Value = big.NewInt(0)

		response, err := tm.DirectTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrValueRequire.Error(), "Expected RequiredValue error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("DirectTransferNegative_FromAndToAreIdentical", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, _ := prepareMockServices(mockController)
		tm := prepareDirectTransferManager(mockAssetboxManager, mockContractManager)

		transfer := prepareDirectTransferObject()
		transfer.To = transfer.From

		response, err := tm.DirectTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrSameAccount.Error(), "Expected SameAccount error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("DirectTransferNegative_MissingTo", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, _ := prepareMockServices(mockController)
		tm := prepareDirectTransferManager(mockAssetboxManager, mockContractManager)

		transfer := prepareDirectTransferObject()
		transfer.To = common.Address{}

		response, err := tm.DirectTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrToRequire.Error(), "Expected RequiredTo error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("DirectTransferNegative_MissingFrom", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, _ := prepareMockServices(mockController)
		tm := prepareDirectTransferManager(mockAssetboxManager, mockContractManager)

		transfer := prepareDirectTransferObject()
		transfer.From = common.Address{}

		response, err := tm.DirectTransfer(ctx, transfer)

		assert.EqualError(t, err, ErrFromRequire.Error(), "Expected RequiredFrom error")
		assert.Empty(t, response, "Response wasn't empty")
	})
}

func TestDirectTransferAsync(t *testing.T) {
	ctx := context.TODO()

	mockController := gomock.NewController(t)
	defer mockController.Finish()

	t.Run("DirectTransferAsyncPositive", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, _ := prepareMockServices(mockController)
		tm := prepareDirectTransferManager(mockAssetboxManager, mockContractManager)

		transfer := prepareDirectTransferObject()

		mockContractManager.EXPECT().GetAssetboxBalance(ctx, transfer.From).
			Return(getSufficientBalance(), nil)

		blockNum, txHash := createTransferResponse()
		blockNum = uint64(0)

		mockContractManager.EXPECT().
			DirectTransfer(ctx, transfer.From, transfer.To, transfer.Value,
				transfer.ExtraData, tm.bitbon.GetServicePk(), true).
			Return(blockNum, txHash, nil)

		response, err := tm.DirectTransferAsync(ctx, transfer)

		assert.NoError(t, err, "Unexpected error")
		assert.Equal(t, mapTransferResponse(blockNum, txHash), response, "Response differs from expected")
	})

	t.Run("DirectTransferAsyncPositive_DestinationIsServiceAccount", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, _ := prepareMockServices(mockController)
		tm := prepareDirectTransferManager(mockAssetboxManager, mockContractManager)

		transfer := prepareDirectTransferObject()
		transfer.To = generateServiceAssetboxAddress()

		mockContractManager.EXPECT().GetAssetboxBalance(ctx, transfer.From).
			Return(getSufficientBalance(), nil)

		blockNum, txHash := createTransferResponse()
		blockNum = uint64(0)

		mockContractManager.EXPECT().
			DirectTransfer(ctx, transfer.From, transfer.To, transfer.Value,
				transfer.ExtraData, tm.bitbon.GetServicePk(), true).
			Return(blockNum, txHash, nil)

		response, err := tm.DirectTransferAsync(ctx, transfer)

		assert.NoError(t, err, "Unexpected error")
		assert.Equal(t, mapTransferResponse(blockNum, txHash), response, "Response differs from expected")
	})

	t.Run("DirectTransferAsyncPositive_SourceIsServiceAccount", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, _ := prepareMockServices(mockController)
		tm := prepareDirectTransferManager(mockAssetboxManager, mockContractManager)

		transfer := prepareDirectTransferObject()
		transfer.From = generateServiceAssetboxAddress()

		mockContractManager.EXPECT().GetAssetboxBalance(ctx, transfer.From).
			Return(getSufficientBalance(), nil)

		blockNum, txHash := createTransferResponse()
		blockNum = uint64(0)

		mockContractManager.EXPECT().
			DirectTransfer(ctx, transfer.From, transfer.To, transfer.Value,
				transfer.ExtraData, tm.bitbon.GetServicePk(), true).
			Return(blockNum, txHash, nil)

		response, err := tm.DirectTransferAsync(ctx, transfer)

		assert.NoError(t, err, "Unexpected error")
		assert.Equal(t, mapTransferResponse(blockNum, txHash), response, "Response differs from expected")
	})

	t.Run("DirectTransferAsyncNegative_InsufficientBalance", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, _ := prepareMockServices(mockController)
		tm := prepareDirectTransferManager(mockAssetboxManager, mockContractManager)

		transfer := prepareDirectTransferObject()

		mockContractManager.EXPECT().GetAssetboxBalance(ctx, transfer.From).
			Return(getInsufficientBalance(), nil)

		response, err := tm.DirectTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrLowAssetboxBalance.Error(), "Expected LowAssetboxBalance error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("DirectTransferAsyncNegative_NilAmount", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, _ := prepareMockServices(mockController)
		tm := prepareDirectTransferManager(mockAssetboxManager, mockContractManager)

		transfer := prepareDirectTransferObject()
		transfer.Value = nil

		response, err := tm.DirectTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrValueRequire.Error(), "Expected RequiredValue error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("DirectTransferAsyncNegative_NegativeAmount", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, _ := prepareMockServices(mockController)
		tm := prepareDirectTransferManager(mockAssetboxManager, mockContractManager)

		transfer := prepareDirectTransferObject()
		transfer.Value = big.NewInt(-10)

		response, err := tm.DirectTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrValueRequire.Error(), "Expected RequiredValue error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("DirectTransferAsyncNegative_ZeroAmount", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, _ := prepareMockServices(mockController)
		tm := prepareDirectTransferManager(mockAssetboxManager, mockContractManager)

		transfer := prepareDirectTransferObject()
		transfer.Value = big.NewInt(0)

		response, err := tm.DirectTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrValueRequire.Error(), "Expected RequiredValue error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("DirectTransferAsyncNegative_FromAndToAreIdentical", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, _ := prepareMockServices(mockController)
		tm := prepareDirectTransferManager(mockAssetboxManager, mockContractManager)

		transfer := prepareDirectTransferObject()
		transfer.To = transfer.From

		response, err := tm.DirectTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrSameAccount.Error(), "Expected SameAccount error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("DirectTransferAsyncNegative_MissingTo", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, _ := prepareMockServices(mockController)
		tm := prepareDirectTransferManager(mockAssetboxManager, mockContractManager)

		transfer := prepareDirectTransferObject()
		transfer.To = common.Address{}

		response, err := tm.DirectTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrToRequire.Error(), "Expected RequiredTo error")
		assert.Empty(t, response, "Response wasn't empty")
	})

	t.Run("DirectTransferAsyncNegative_MissingFrom", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, _ := prepareMockServices(mockController)
		tm := prepareDirectTransferManager(mockAssetboxManager, mockContractManager)

		transfer := prepareDirectTransferObject()
		transfer.From = common.Address{}

		response, err := tm.DirectTransferAsync(ctx, transfer)

		assert.EqualError(t, err, ErrFromRequire.Error(), "Expected RequiredFrom error")
		assert.Empty(t, response, "Response wasn't empty")
	})
}
