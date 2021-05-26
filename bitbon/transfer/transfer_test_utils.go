package transfer

import (
	"context"
	"crypto/rand"
	"errors"
	"math/big"
	"time"

	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/dto"

	"github.com/golang/mock/gomock"

	bb "github.com/simcord-llc/bitbon-system-blockchain/bitbon"
	loggerContext "github.com/simcord-llc/bitbon-system-blockchain/bitbon/context"
	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/models"
	"github.com/simcord-llc/bitbon-system-blockchain/common"
	"github.com/simcord-llc/bitbon-system-blockchain/crypto"
	"github.com/simcord-llc/bitbon-system-blockchain/log"
)

var (
	sourceAssetbox      = common.BytesToAddress([]byte("798e8c22eb5a46a5a401a4a123e4fe768e92abbe"))
	destinationAssetbox = common.BytesToAddress([]byte("628e8c09eb5346a5a451a4654634fe743e92acce"))

	accountID = "exampleAccountId"

	servicePk = "fffffffffffaffaffafffffffffffffebaaedce6af48a03bbfd25e8cd0364141"

	transferAmount int64 = 10

	ErrTransferHasExpired          = errors.New("transfer approve failed. Status is 'expired'")
	ErrWrongProtectionCode         = errors.New("transfer approve failed. Status is 'wrong protection code'")
	ErrProtectionCodeLimitExceeded = errors.
					New("transfer approve failed. Status is 'wrong protection code limit exceeded'")
	ErrFailedToParseKey = errors.
				New("error decrypting assetbox wallet: invalid character 'S' looking for beginning of value")
	ErrWrongPassphrase = errors.
				New("error decrypting assetbox wallet: could not decrypt key with given password")

	protectionCode = "123456"

	transferID = "0x0000000f"

	encryptionKey = []byte("0123456789abcdef")
	passphrase    = []byte("testpass")
	wallet        = []byte("{\"address\":\"6855d923dd11e7af487a8fa9c6e05bb9804f6036\",\"crypto\":{\"cipher\":\"aes-128-ctr\",\"ciphertext\":\"b47f5125f44de17826716c38740bc9350f272c2fc2a861fbb440d7f7a2903a3c\",\"cipherparams\":{\"iv\":\"54478ea716f62cc5bac434298ccc8080\"},\"kdf\":\"scrypt\",\"kdfparams\":{\"dklen\":32,\"n\":262144,\"p\":1,\"r\":8,\"salt\":\"fd1502b4efc63dc77671aba14dbf478a3b180b8de898e377d6a74338a07bf4b8\"},\"mac\":\"7e1d1ab49f3ef3e9c8c651eea6816439f84e00bb54274a3ce5f804f8b8cbd660\"},\"id\":\"0834b48c-d766-4cd0-80d6-7429524e549a\",\"version\":3}")
	wrongPass     = []byte("Wrong")
	invalidWallet = []byte("SampleText")
)

func prepareQuickTransferManager(am bb.AssetboxManager, cm bb.ContractManager, enc bb.PkEncryptor) *Manager {
	tm := &Manager{
		bitbon:    &bb.Bitbon{},
		encryptor: enc,
		logger: loggerContext.LoggerFromContext(
			loggerContext.NewLoggerContext(
				context.Background(),
				log.New("src", "TransferManager"),
			),
		),
		ready: make(chan struct{}),
		done:  make(chan struct{}),
	}

	tm.bitbon.SetDecryptAssetboxWalletPassword(string(encryptionKey))
	tm.bitbon.InjectAssetboxManager(am)
	tm.bitbon.InjectContractManager(cm)

	tm.makeReady()

	return tm
}

func prepareSafeTransferManager(am bb.AssetboxManager, cm bb.ContractManager, enc bb.PkEncryptor) *Manager {
	tm := &Manager{
		bitbon:    &bb.Bitbon{},
		encryptor: enc,
		logger: loggerContext.LoggerFromContext(
			loggerContext.NewLoggerContext(
				context.Background(),
				log.New("src", "TransferManager"),
			),
		),
		ready: make(chan struct{}),
		done:  make(chan struct{}),
	}

	tm.bitbon.SetDecryptAssetboxWalletPassword(string(encryptionKey))
	tm.bitbon.InjectAssetboxManager(am)
	tm.bitbon.InjectContractManager(cm)

	return tm
}

func prepareDirectTransferManager(am bb.AssetboxManager, cm bb.ContractManager) *Manager {
	bitbon := bb.InitDefaultBitbonService(&bb.Config{}, nil)

	err := bb.InitServicePart(bitbon, &bb.Config{
		ServicePrivateKey: servicePk,
	})
	if err != nil {
		panic("error during TransferManager preparation")
	}

	tm := &Manager{
		bitbon: bitbon,
		logger: loggerContext.LoggerFromContext(
			loggerContext.NewLoggerContext(
				context.Background(),
				log.New("src", "TransferManager"),
			),
		),
		ready: make(chan struct{}),
		done:  make(chan struct{}),
	}

	tm.bitbon.InjectAssetboxManager(am)
	tm.bitbon.InjectContractManager(cm)

	return tm
}

func generateServiceAssetboxAddress() common.Address {
	privateKey, err := crypto.HexToECDSA(servicePk)
	if err != nil {
		panic("error during service assetbox address generation")
	}

	return crypto.PubkeyToAddress(privateKey.PublicKey)
}

func getSufficientBalance() *big.Int {
	return big.NewInt(transferAmount)
}

func getInsufficientBalance() *big.Int {
	return big.NewInt(transferAmount - 1)
}

func createTransferResponse() (blockNum uint64, txHash common.Hash) {
	random, _ := rand.Int(rand.Reader, big.NewInt(999999999))
	blockNum = random.Uint64()

	token := make([]byte, 32)

	if _, err := rand.Read(token); err != nil {
		panic("error while generating txHash")
	}

	txHash = common.BytesToHash(token)

	return blockNum, txHash
}

func mapTransferResponse(blockNum uint64, txHash common.Hash) *models.TransferResponseObj {
	return &models.TransferResponseObj{
		BlockNumber: blockNum,
		TxHash:      txHash,
	}
}

func prepareMockServices(mockController *gomock.Controller) (*MockAssetboxManager,
	*MockContractManager, *MockPkEncryptor) {
	return NewMockAssetboxManager(mockController),
		NewMockContractManager(mockController),
		NewMockPkEncryptor(mockController)
}

func prepareCancelSafeTransferObject() *models.CancelTransferObj {
	return &models.CancelTransferObj{
		Address:    destinationAssetbox,
		TransferID: transferID,
		CryptoData: dto.AssetboxCryptoData{
			Wallet:     []byte("abcd4561abcd4561abcd4561abcd4561abcd4561abcd4561abcd4561abcd4561"),
			Passphrase: []byte("abcd4561abcd4561"),
		},
	}
}

func prepareCreateSafeTransferObject() *models.CreateTransferObj {
	return &models.CreateTransferObj{
		From:           sourceAssetbox,
		AccountID:      accountID,
		To:             destinationAssetbox,
		Value:          big.NewInt(transferAmount),
		TransferID:     transferID,
		ProtectionCode: protectionCode,
		Retries:        3,
		ExpiresAt:      time.Now().Add(time.Second * 10000).Unix(),
		CryptoData: dto.AssetboxCryptoData{
			Wallet:     []byte("abcd4561abcd4561abcd4561abcd4561abcd4561abcd4561abcd4561abcd4561"),
			Passphrase: []byte("abcd4561abcd4561"),
		},
	}
}

func prepareCreateWPCSafeTransferObject() *models.CreateTransferObj {
	return &models.CreateTransferObj{
		From:       sourceAssetbox,
		AccountID:  accountID,
		To:         destinationAssetbox,
		Value:      big.NewInt(transferAmount),
		TransferID: transferID,
		Retries:    3,
		ExpiresAt:  time.Now().Add(time.Second * 10000).Unix(),
		CryptoData: dto.AssetboxCryptoData{
			Wallet:     []byte("abcd4561abcd4561abcd4561abcd4561abcd4561abcd4561abcd4561abcd4561"),
			Passphrase: []byte("abcd4561abcd4561"),
		},
	}
}

func prepareDirectTransferObject() *models.DirectTransferObj {
	return &models.DirectTransferObj{
		From:  sourceAssetbox,
		To:    destinationAssetbox,
		Value: big.NewInt(transferAmount),
	}
}

func prepareQuickTransferObject() *models.QuickTransferObj {
	return &models.QuickTransferObj{
		From:      sourceAssetbox,
		AccountID: accountID,
		To:        destinationAssetbox,
		Value:     big.NewInt(transferAmount),
		CryptoData: dto.AssetboxCryptoData{
			Wallet:     []byte("abcd4561abcd4561abcd4561abcd4561abcd4561abcd4561abcd4561abcd4561"),
			Passphrase: []byte("abcd4561abcd4561"),
		},
	}
}

func prepareApproveSafeTransferObject() *models.ApproveTransferObj {
	return &models.ApproveTransferObj{
		Address:        destinationAssetbox,
		TransferID:     transferID,
		ProtectionCode: protectionCode,
		CryptoData: dto.AssetboxCryptoData{
			Wallet:     []byte("abcd4561abcd4561abcd4561abcd4561abcd4561abcd4561abcd4561abcd4561"),
			Passphrase: []byte("abcd4561abcd4561"),
		},
	}
}

func prepareApproveWPCSafeTransferObject() *models.ApproveTransferObj {
	return &models.ApproveTransferObj{
		Address:    destinationAssetbox,
		TransferID: transferID,
		CryptoData: dto.AssetboxCryptoData{
			Wallet:     []byte("abcd4561abcd4561abcd4561abcd4561abcd4561abcd4561abcd4561abcd4561"),
			Passphrase: []byte("abcd4561abcd4561"),
		},
	}
}
