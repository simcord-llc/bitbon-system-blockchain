// Copyright Simcord LLC
// This file is part of the Bitbon System Blockchain Node library.
//
// The Bitbon System Blockchain Node library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The Bitbon System Blockchain Node library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the Bitbon System Blockchain Node library. If not, see <http://www.gnu.org/licenses/>.

package bitbon

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/contracts"
	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/contracts/contract_snapshot"
	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/dto"
	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/models"
	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/noncer"
	"github.com/simcord-llc/bitbon-system-blockchain/common"
)

type TransferManager interface {
	Stop() error

	QuickTransfer(ctx context.Context, transfer *models.QuickTransferObj) (*models.TransferResponseObj, error)
	QuickTransferAsync(ctx context.Context, transfer *models.QuickTransferObj) (*models.TransferResponseObj, error)
	CreateSafeTransfer(ctx context.Context, transfer *models.CreateTransferObj) (*models.TransferResponseObj, error)
	CreateSafeTransferAsync(ctx context.Context,
		transfer *models.CreateTransferObj) (*models.TransferResponseObj, error)
	ApproveSafeTransfer(ctx context.Context, transfer *models.ApproveTransferObj) (*models.TransferResponseObj, error)
	ApproveSafeTransferAsync(ctx context.Context,
		transfer *models.ApproveTransferObj) (*models.TransferResponseObj, error)
	CancelSafeTransfer(ctx context.Context, transfer *models.CancelTransferObj) (*models.TransferResponseObj, error)
	CancelSafeTransferAsync(ctx context.Context,
		transfer *models.CancelTransferObj) (*models.TransferResponseObj, error)
	CreateWPCSafeTransfer(ctx context.Context,
		transfer *models.CreateTransferObj) (*models.TransferResponseObj, error)
	CreateWPCSafeTransferAsync(ctx context.Context,
		transfer *models.CreateTransferObj) (*models.TransferResponseObj, error)
	ApproveWPCSafeTransfer(ctx context.Context,
		transfer *models.ApproveTransferObj) (*models.TransferResponseObj, error)
	ApproveWPCSafeTransferAsync(ctx context.Context,
		transfer *models.ApproveTransferObj) (*models.TransferResponseObj, error)
	CancelWPCSafeTransfer(ctx context.Context, transfer *models.CancelTransferObj) (*models.TransferResponseObj, error)
	CancelWPCSafeTransferAsync(ctx context.Context,
		transfer *models.CancelTransferObj) (*models.TransferResponseObj, error)
	DirectTransfer(ctx context.Context, transfer *models.DirectTransferObj) (*models.TransferResponseObj, error)
	DirectTransferAsync(ctx context.Context, transfer *models.DirectTransferObj) (*models.TransferResponseObj, error)
	ExpireTransfers(ctx context.Context) (hashes []common.Hash, num int, err error)
	GetTransfer(ctx context.Context, transferID string) (*contracts.ReceiptTransfer, error)
	GetTransferByIndex(ctx context.Context, index *big.Int) (*contracts.ReceiptTransfer, error)
	TransferExists(ctx context.Context, transferID string) (bool, error)
}

type AssetboxManager interface {
	PrepareAssetbox(ctx context.Context) (assetbox *models.Assetbox, err error)
	PrepareAssetboxes(ctx context.Context, count uint64) (assetboxes []*models.Assetbox, err error)
	SetPublicAssetboxInfo(ctx context.Context, req *dto.SetPublicAssetboxInfoRequest) (err error)
	SetPublicAssetboxInfos(ctx context.Context, assetboxes []*contracts.Assetbox) (err error)
	DeleteAssetbox(ctx context.Context, req *dto.DeleteAssetboxRequest) (err error)
}

type MiningAgent interface {
	ProposeDistribution(ctx context.Context,
		req dto.ProposeDistributionRequest) (err error)
	GetCurrentDistribution(ctx context.Context) (distribution map[string]uint64, err error)
	GetAllMiningAgents(ctx context.Context) (res []common.Address, err error)
	IsMiningAgent(ctx context.Context, address common.Address) (res bool, err error)
	AddMiningAgent(ctx context.Context, address common.Address) (err error)
	RemoveMiningAgent(ctx context.Context, address common.Address) (err error)
}

type PkEncryptor interface {
	Encrypt(data, key []byte) ([]byte, error)
	Decrypt(data, key []byte) ([]byte, error)
}

type ContractManager interface {
	// system
	Stop() error
	Ready() bool
	Stopped() bool
	WaitReady(exit chan struct{}) bool
	GetNoncer() *noncer.Noncer
	// contracts

	GetContractAddressAssetboxInfo() (common.Address, error)
	GetContractAddressAssetbox() (common.Address, error)
	GetContractAddressBitbon() (common.Address, error)
	GetContractAddressTransfer() (common.Address, error)
	GetContractAddressSafeTransferStorage() (common.Address, error)
	GetContractAddressBitbonSupport() (common.Address, error)
	GetContractAddressBitbonStorage() (common.Address, error)
	GetContractAddressAccessStorage() (common.Address, error)
	GetContractAddressDistributionStorage() (common.Address, error)
	GetContractAddressMiningAgentStorage() (common.Address, error)
	GetContractStorageAddress() (common.Address, error)
	GetContractAddresses() map[contract_snapshot.ContractType]common.Address
	GetCurrentContractAbiInfo(version dto.ContractVersion) ([]*dto.AbiInfo, error)
	SaveContractStorageAddress(addr common.Address) error
	GetEthAPIWrapper() *contracts.EthApiWrapper

	// assetboxes
	GetAssetboxInfo(ctx context.Context, addr common.Address) (a *contracts.Assetbox, err error)
	GetAssetboxBalance(ctx context.Context, addr common.Address) (balance *big.Int, err error)
	GetAssetboxBalances(ctx context.Context,
		addresses []common.Address) (balances map[common.Address]*big.Int, err error)
	GetAsseboxesMiningState(ctx context.Context, address common.Address) (bool, error)
	GetAssetboxBalancesSum(ctx context.Context, addresses []common.Address) (sum *big.Int, err error)
	SetAssetboxInfo(ctx context.Context, a *contracts.Assetbox, key *ecdsa.PrivateKey) (err error)
	SetAssetboxInfos(ctx context.Context, a []*contracts.Assetbox, key *ecdsa.PrivateKey) (err error)
	GetAssetboxInfoByAlias(ctx context.Context, alias string) (a *contracts.Assetbox, err error)
	DeleteAssetboxInfo(ctx context.Context, key *ecdsa.PrivateKey) (err error)

	// transactions
	EmitEther(ctx context.Context, to common.Address, value *big.Int, nonce *uint64) (err error)
	ApproveWPCSafeTransfer(ctx context.Context, transferID, extraData []byte,
		key *ecdsa.PrivateKey, async bool) (blockNum uint64, txHash common.Hash, err error)
	CreateWPCSafeTransfer(ctx context.Context, t *contracts.Transfer,
		key *ecdsa.PrivateKey, async bool) (blockNum uint64, txHash common.Hash, err error)
	CancelSafeTransfer(ctx context.Context, transferID, extraData []byte,
		key *ecdsa.PrivateKey, async bool) (blockNum uint64, txHash common.Hash, err error)
	ApproveSafeTransfer(ctx context.Context, transferID, protectionCode,
		extraData []byte, key *ecdsa.PrivateKey, async bool) (blockNum uint64, txHash common.Hash, err error)
	WatchBitbonBalanceUnLocked(sink chan<- *contracts.Balance) (err error)
	WatchBitbonBalanceLocked(sink chan<- *contracts.Balance) (err error)
	WatchBitbonBalanceChanged(sink chan<- *contracts.Balance) (err error)
	WatchAssetboxInfoSet(sink chan<- *contracts.Assetbox) (err error)
	WatchTransferExpired(sink chan<- *contracts.TransferExpired) (err error)
	GetExpiredTransfers(ctx context.Context, firstTransfer,
		lastTransfer *big.Int) (expireTransferIds [][32]byte, err error)
	GetTransferLength(ctx context.Context) (transferLength *big.Int, err error)
	GetOldestPending(ctx context.Context) (oldestPendingID *big.Int, err error)
	TransferExists(ctx context.Context, transferID []byte) (exists bool, err error)
	QuickTransfer(ctx context.Context, to common.Address, value *big.Int, extraData []byte,
		key *ecdsa.PrivateKey, async bool) (blockNum uint64, txHash common.Hash, err error)
	GetTransferByIndex(ctx context.Context, index *big.Int) (transfer *contracts.ReceiptTransfer, err error)
	GetTransfer(ctx context.Context, transferID []byte) (transfer *contracts.ReceiptTransfer, err error)
	CancelWPCSafeTransfer(ctx context.Context, transferID, extraData []byte,
		key *ecdsa.PrivateKey, async bool) (blockNum uint64, txHash common.Hash, err error)
	SetOldestPending(ctx context.Context, pendingIndex *big.Int, key *ecdsa.PrivateKey) (err error)
	RemoveMiningAgent(ctx context.Context, address common.Address, key *ecdsa.PrivateKey) (err error)
	ExpireSafeTransfers(ctx context.Context, ids [][32]byte, key *ecdsa.PrivateKey) (hash common.Hash, err error)
	CreateSafeTransfer(ctx context.Context, t *contracts.Transfer,
		key *ecdsa.PrivateKey, async bool) (blockNum uint64, txHash common.Hash, err error)
	SearchOldestPending(ctx context.Context, firstTransfer,
		lastTransfer *big.Int) (index *big.Int, present bool, err error)
	IsMiningAgent(ctx context.Context, address common.Address) (res bool, err error)
	DirectTransfer(ctx context.Context, from, to common.Address,
		value *big.Int, extraData []byte,
		key *ecdsa.PrivateKey, async bool) (blockNum uint64, txHash common.Hash, err error)
	AddMiningAgent(ctx context.Context, address common.Address, key *ecdsa.PrivateKey) (err error)
	GetCurrentDistribution(ctx context.Context) (distribution map[string]uint64, err error)
	ProposeDistribution(ctx context.Context, distribution map[string]uint64, key *ecdsa.PrivateKey) (err error)
	GetAllMiningAgents(ctx context.Context) (res []common.Address, err error)
}
