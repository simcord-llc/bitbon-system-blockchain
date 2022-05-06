package transfer

import (
	"context"
	"github.com/simcord-llc/bitbon-system-blockchain/crypto"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestExpireTransferAsync(t *testing.T) {
	ctx := context.TODO()

	mockController := gomock.NewController(t)
	defer mockController.Finish()

	t.Run("ExpireTransferAsyncPositive", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, _ := prepareMockServices(mockController)
		tm := prepareDirectTransferManager(mockAssetboxManager, mockContractManager)

		blockNum, txHash := createTransferResponse()
		blockNum = uint64(0)

		transferIds := prepareTestIDSequence(20)
		transferIdsBytes := prepareTransferIdsBytes(transferIds)

		mockContractManager.EXPECT().ExpireSafeTransfers(ctx, transferIdsBytes, tm.bitbon.GetServicePk()).Return(blockNum, txHash, nil)

		response, err := tm.ExpireTransfersAsync(ctx, transferIds)

		assert.NoError(t, err, "Unexpected error occurred")
		assert.Equal(t, mapTransferResponse(blockNum, txHash), response, "Response differs from expected")
	})

	t.Run("ExpireTransferAsyncNegative_MaxBatchSizeExceeded", func(t *testing.T) {
		mockAssetboxManager, mockContractManager, _ := prepareMockServices(mockController)
		tm := prepareDirectTransferManager(mockAssetboxManager, mockContractManager)

		transferIds := prepareTestIDSequence(21)

		response, err := tm.ExpireTransfersAsync(ctx, transferIds)

		assert.EqualError(t, err, "max batch size of 20 exceeded", "expected MaxBatchExceeded error")
		assert.Empty(t, response, "Response wasn't empty")
	})
}

func prepareTransferIdsBytes(transferIds []string) [][32]byte {
	transferIdsBytes := make([][32]byte, len(transferIds))
	for i := range transferIds {
		copy(transferIdsBytes[i][:], crypto.Keccak256([]byte(transferIds[i])))
	}

	return transferIdsBytes
}
