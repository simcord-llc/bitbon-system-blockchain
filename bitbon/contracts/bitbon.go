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

// nolint:nakedret
package contracts

import (
	"context"
	"math/big"
	"time"

	"github.com/simcord-llc/bitbon-system-blockchain/core"
	"github.com/simcord-llc/bitbon-system-blockchain/core/types"

	"crypto/ecdsa"

	"bytes"

	"github.com/pkg/errors"
	"github.com/simcord-llc/bitbon-system-blockchain/accounts/abi/bind"
	bitbonErrors "github.com/simcord-llc/bitbon-system-blockchain/bitbon/errors"
	"github.com/simcord-llc/bitbon-system-blockchain/common"
)

type StatusType int

const (
	SafeTransferStatusUndefined                        StatusType = 0
	SafeTransferStatusSuccess                          StatusType = 1
	SafeTransferStatusWrongProtectionCode              StatusType = 2
	SafeTransferStatusWrongProtectionCodeLimitExceeded StatusType = 3
	SafeTransferStatusExpired                          StatusType = 4
)

const (
	RecoverableErrRetries = 10
	RecoverableErrDelay   = 500 * time.Millisecond
)

func (m *Manager) getBitbonImpl() (*BitbonImpl, error) {
	if !m.Ready() {
		return nil, ErrorNotReady
	}
	contractAddress, err := m.GetContractAddressBitbon()
	if err != nil {
		return nil, err
	}
	contract, err := NewBitbonImpl(contractAddress, m.EthApiWrapper)
	if err != nil {
		return nil, err
	}
	return contract, nil
}

// GetAssetboxBalance gets assetbox balance by assetbox address from blockchain
func (m *Manager) GetAssetboxBalance(ctx context.Context, addr common.Address) (balance *big.Int, err error) {
	m.logger.Debug("contracts manager GetAssetboxBalance called", "address", addr.Hex())

	defer func() {
		if err != nil {
			err = bitbonErrors.NewContractCallError(err)
		}
	}()

	contract, err := m.getBitbonImpl()
	if err != nil {
		return nil, err
	}

	opts := &bind.CallOpts{
		Pending: false,
		From:    addr,
		Context: ctx,
	}

	return contract.BalanceOf(opts, addr)
}

// GetAssetboxBalances gets assetbox balances by assetbox addresses from blockchain
func (m *Manager) GetAssetboxBalances(ctx context.Context, addresses []common.Address) (balances map[common.Address]*big.Int, err error) {
	m.logger.Debug("contracts manager GetAssetboxBalances called", "addresses", addresses)

	defer func() {
		if err != nil {
			err = bitbonErrors.NewContractCallError(err)
		}
	}()

	contract, err := m.getBitbonImpl()
	if err != nil {
		return nil, err
	}

	opts := &bind.CallOpts{
		Pending: false,
		Context: ctx,
	}

	rawBalances, err := contract.GetScopeBalances(opts, addresses)
	if err != nil {
		return nil, err
	}

	balances = make(map[common.Address]*big.Int)

	for index, rawBalance := range rawBalances {
		balances[addresses[index]] = rawBalance
	}

	return
}

// GetAssetboxBalancesSum gets sum of assetbox balances by assetbox addresses
func (m *Manager) GetAssetboxBalancesSum(ctx context.Context, addresses []common.Address) (sum *big.Int, err error) {
	m.logger.Debug("contracts manager GetAssetboxBalancesSum called", "addresses", addresses)

	defer func() {
		if err != nil {
			err = bitbonErrors.NewContractCallError(err)
		}
	}()

	contract, err := m.getBitbonImpl()
	if err != nil {
		return nil, err
	}

	opts := &bind.CallOpts{
		Pending: false,
		Context: ctx,
	}

	return contract.GetScopeBalance(opts, addresses)
}

func (m *Manager) ExpireSafeTransfers(ctx context.Context, ids [][32]byte, key *ecdsa.PrivateKey) (hash common.Hash, err error) {
	m.logger.Debug("contracts manager ExpireSafeTransfers called", "IdsLen", len(ids))
	defer func() {
		if err != nil {
			err = bitbonErrors.NewContractCallError(err)
		}
	}()

	contract, err := m.getBitbonImpl()
	if err != nil {
		return
	}
	opts, err := m.prepareTransactOpts(ctx, key)
	if err != nil {
		return
	}
	tx, err := contract.ExpireSafeTransfers(opts, ids)
	if err != nil {
		return
	}

	hash = tx.Hash()

	m.logger.Debug("Waiting mining tx ExpireSafeTransfers", "ids", ids, "txHash", tx.Hash())
	_, err = m.waitAndCheckTransaction(ctx, tx)

	return
}

func (m *Manager) QuickTransfer(ctx context.Context, to common.Address, value *big.Int, extraData []byte, key *ecdsa.PrivateKey, async bool) (blockNum uint64, txHash common.Hash, err error) {
	m.logger.Debug("contracts manager QuickTransfer called", "to", to.Hex(), "value", value.String(), "async", async)

	defer func() {
		if err != nil {
			err = bitbonErrors.NewContractCallError(err)
		}
	}()

	contract, err := m.getBitbonImpl()
	if err != nil {
		return
	}

	opts, err := m.prepareTransactOpts(ctx, key)
	if err != nil {
		return
	}

	tx, err := contract.QuickTransfer(opts, to, value, extraData)
	if err != nil {
		return
	}

	if async {
		return 0, tx.Hash(), nil
	}

	m.logger.Debug("Waiting mining tx QuickTransfer", "to", to.Hex(), "value", value.String(), "async", async)
	receipt, err := m.waitAndCheckTransaction(ctx, tx)
	if err != nil {
		return
	}

	return m.getReceiptParams(receipt)
}

// CreateSafeTransfer creates safe transfer and wait processing
func (m *Manager) CreateSafeTransfer(ctx context.Context, t *Transfer, key *ecdsa.PrivateKey, async bool) (blockNum uint64, txHash common.Hash, err error) {
	m.logger.Debug("contracts manager CreateSafeTransfer called", "TransferID", string(t.TransferID))
	defer func() {
		if err != nil {
			err = bitbonErrors.NewContractCallError(err)
		}
	}()

	contract, err := m.getBitbonImpl()
	if err != nil {
		return
	}

	opts, err := m.prepareTransactOpts(ctx, key)
	if err != nil {
		return
	}

	tx, err := m.retryTransaction(func() (*types.Transaction, error) {
		return contract.CreateSafeTransfer(opts, t.To, t.Value, t.Proof, t.VK, t.TransferID, t.Retries, t.ExpiresAt, t.ExtraData)
	})
	if err != nil {
		return
	}

	if async {
		return 0, tx.Hash(), nil
	}

	m.logger.Debug("Waiting mining for CreateSafeTransfer tx", "TransferID", string(t.TransferID))
	receipt, err := m.waitAndCheckTransaction(ctx, tx)
	if err != nil {
		return
	}

	return m.getReceiptParams(receipt)
}

// CreateWPCSafeTransfer creates safe transfer without protection code and wait processing
func (m *Manager) CreateWPCSafeTransfer(ctx context.Context, t *Transfer, key *ecdsa.PrivateKey, async bool) (blockNum uint64, txHash common.Hash, err error) {
	m.logger.Debug("contracts manager CreateWPCSafeTransfer called", "TransferID", string(t.TransferID))
	defer func() {
		if err != nil {
			err = bitbonErrors.NewContractCallError(err)
		}
	}()

	contract, err := m.getBitbonImpl()
	if err != nil {
		return
	}

	opts, err := m.prepareTransactOpts(ctx, key)
	if err != nil {
		return
	}

	tx, err := m.retryTransaction(func() (*types.Transaction, error) {
		return contract.CreateWPCSafeTransfer(opts, t.To, t.Value, t.TransferID, t.ExpiresAt, t.ExtraData)
	})
	if err != nil {
		return
	}

	if async {
		return 0, tx.Hash(), nil
	}

	m.logger.Debug("Waiting mining for CreateWPCSafeTransfer tx", "TransferID", string(t.TransferID))
	receipt, err := m.waitAndCheckTransaction(ctx, tx)
	if err != nil {
		return
	}

	return m.getReceiptParams(receipt)
}

// ApproveSafeTransfer approves safe transfer
func (m *Manager) ApproveSafeTransfer(ctx context.Context, transferId []byte, protectionCode []byte, extraData []byte, key *ecdsa.PrivateKey, async bool) (blockNum uint64, txHash common.Hash, err error) {
	m.logger.Debug("contracts manager ApproveSafeTransfer called", "TransferID", string(transferId))

	contract, err := m.getBitbonImpl()
	if err != nil {
		err = bitbonErrors.NewContractCallError(err)
		return
	}

	transferContract, err := m.getTransferImpl()
	if err != nil {
		err = bitbonErrors.NewContractCallError(err)
		return
	}

	opts, err := m.prepareTransactOpts(ctx, key)
	if err != nil {
		err = bitbonErrors.NewContractCallError(err)
		return
	}

	tx, err := m.retryTransaction(func() (*types.Transaction, error) {
		return contract.ApproveSafeTransfer(opts, transferId, protectionCode, extraData)
	})
	if err != nil {
		err = bitbonErrors.NewContractCallError(err)
		return
	}

	if async {
		return 0, tx.Hash(), nil
	}

	m.logger.Debug("Waiting mining for ApproveSafeTransfer tx", "TransferID", string(transferId))
	receipt, err := m.waitAndCheckTransaction(ctx, tx)
	if err != nil {
		err = bitbonErrors.NewContractCallError(err)
		return
	}

	// in ApproveSafeTransfer we need to save the state to decrease retry limit
	// so method in solidity do not use assert or require
	// Thus we must look at logs, catch SafeTransferApprovalResult event and look at status of this event to understand the result of approve method
	blockNum, txHash, err = m.getReceiptParams(receipt)
	if err != nil {
		err = bitbonErrors.NewContractCallError(err)
		return
	}

	it, err := transferContract.FilterSafeTransferApprovalResult(&bind.FilterOpts{Start: blockNum, End: &blockNum})
	if err != nil {
		err = bitbonErrors.NewContractCallError(err)
		return
	}
	defer it.Close()

	for it.Next() {
		// if event exists with our transferId and status is Success - everything is ok and transfer is approved
		if bytes.Equal(it.Event.TransferId, transferId) {
			switch StatusType(it.Event.Status) {
			case SafeTransferStatusSuccess:
				m.logger.Debug("obtained 'success' status in SafeTransferApprovalResult event")
				err = nil
				return
			case SafeTransferStatusUndefined:
				err = bitbonErrors.NewSafeTransferUndefinedError(errors.New("transfer approve failed. Status is 'undefined'"))
				return
			case SafeTransferStatusWrongProtectionCode:
				err = bitbonErrors.NewSafeTransferWrongProtectionCodeError(errors.New("transfer approve failed. Status is 'wrong protection code'"))
				return
			case SafeTransferStatusWrongProtectionCodeLimitExceeded:
				err = bitbonErrors.NewSafeTransferWrongProtectionCodeLimitError(errors.New("transfer approve failed. Status is 'wrong protection code limit exceeded'"))
				return
			case SafeTransferStatusExpired:
				err = bitbonErrors.NewSafeTransferExpiredError(errors.New("transfer approve failed. Status is 'expired'"))
				return
			}
		}
	}
	if it.Error() != nil {
		m.logger.Error(it.Error().Error())
		err = bitbonErrors.NewContractCallError(it.Error())
		return
	}

	err = bitbonErrors.NewSafeTransferUnknownReasonError(errors.New("transfer approve failed for unknown reason"))
	return
}

// ApproveWPCSafeTransfer approves safe transfer without protection code
func (m *Manager) ApproveWPCSafeTransfer(ctx context.Context, transferId []byte, extraData []byte, key *ecdsa.PrivateKey, async bool) (blockNum uint64, txHash common.Hash, err error) {
	m.logger.Debug("contracts manager ApproveWPCSafeTransfer called", "TransferID", string(transferId))

	contract, err := m.getBitbonImpl()
	if err != nil {
		err = bitbonErrors.NewContractCallError(err)
		return
	}

	transferContract, err := m.getTransferImpl()
	if err != nil {
		err = bitbonErrors.NewContractCallError(err)
		return
	}

	opts, err := m.prepareTransactOpts(ctx, key)
	if err != nil {
		err = bitbonErrors.NewContractCallError(err)
		return
	}

	tx, err := m.retryTransaction(func() (*types.Transaction, error) {
		return contract.ApproveWPCSafeTransfer(opts, transferId, extraData)
	})
	if err != nil {
		err = bitbonErrors.NewContractCallError(err)
		return
	}

	if async {
		return 0, tx.Hash(), nil
	}

	m.logger.Debug("Waiting mining for ApproveSafeTransfer tx", "TransferID", string(transferId))
	receipt, err := m.waitAndCheckTransaction(ctx, tx)
	if err != nil {
		err = bitbonErrors.NewContractCallError(err)
		return
	}

	// in ApproveSafeTransfer we need to save the state to decrease retry limit
	// so method in solidity do not use assert or require
	// Thus we must look at logs, catch SafeTransferApprovalResult event and look at status of this event to understand the result of approve method
	blockNum, txHash, err = m.getReceiptParams(receipt)
	if err != nil {
		err = bitbonErrors.NewContractCallError(err)
		return
	}

	it, err := transferContract.FilterSafeTransferApprovalResult(&bind.FilterOpts{Start: blockNum, End: &blockNum})
	if err != nil {
		err = bitbonErrors.NewContractCallError(err)
		return
	}
	defer it.Close()

	for it.Next() {
		// if event exists with our transferId and status is Success - everything is ok and transfer is approved
		if bytes.Equal(it.Event.TransferId, transferId) {
			switch StatusType(it.Event.Status) {
			case SafeTransferStatusSuccess:
				m.logger.Debug("obtained 'success' status in SafeTransferApprovalResult event")
				err = nil
				return
			case SafeTransferStatusUndefined:
				err = bitbonErrors.NewSafeTransferUndefinedError(errors.New("transfer approve failed. Status is 'undefined'"))
				return
			case SafeTransferStatusWrongProtectionCode:
				err = bitbonErrors.NewSafeTransferWrongProtectionCodeError(errors.New("transfer approve failed. Status is 'wrong protection code'"))
				return
			case SafeTransferStatusWrongProtectionCodeLimitExceeded:
				err = bitbonErrors.NewSafeTransferWrongProtectionCodeLimitError(errors.New("transfer approve failed. Status is 'wrong protection code limit exceeded'"))
				return
			case SafeTransferStatusExpired:
				err = bitbonErrors.NewSafeTransferExpiredError(errors.New("transfer approve failed. Status is 'expired'"))
				return
			}
		}
	}
	if it.Error() != nil {
		m.logger.Error(it.Error().Error())
		err = bitbonErrors.NewContractCallError(it.Error())
		return
	}

	err = bitbonErrors.NewSafeTransferUnknownReasonError(errors.New("transfer approve failed for unknown reason"))
	return
}

// CancelSafeTransfer cancels safe transfer
func (m *Manager) CancelSafeTransfer(ctx context.Context, transferId []byte, extraData []byte, key *ecdsa.PrivateKey, async bool) (blockNum uint64, txHash common.Hash, err error) {
	m.logger.Debug("contracts manager CancelSafeTransfer called", "TransferID", string(transferId))
	defer func() {
		if err != nil {
			err = bitbonErrors.NewContractCallError(err)
		}
	}()

	contract, err := m.getBitbonImpl()
	if err != nil {
		return
	}

	opts, err := m.prepareTransactOpts(ctx, key)
	if err != nil {
		return
	}

	tx, err := m.retryTransaction(func() (*types.Transaction, error) {
		return contract.CancelSafeTransfer(opts, transferId, extraData)
	})
	if err != nil {
		return
	}

	txHash = tx.Hash()

	if async {
		return
	}

	receipt, err := m.waitAndCheckTransaction(ctx, tx)
	if err != nil {
		return
	}

	return m.getReceiptParams(receipt)
}

// CancelWPCSafeTransfer cancels safe transfer without protection code
func (m *Manager) CancelWPCSafeTransfer(ctx context.Context, transferId []byte, extraData []byte, key *ecdsa.PrivateKey, async bool) (blockNum uint64, txHash common.Hash, err error) {
	m.logger.Debug("contracts manager CancelWPCSafeTransfer called", "TransferID", string(transferId))
	defer func() {
		if err != nil {
			err = bitbonErrors.NewContractCallError(err)
		}
	}()

	contract, err := m.getBitbonImpl()
	if err != nil {
		return
	}

	opts, err := m.prepareTransactOpts(ctx, key)
	if err != nil {
		return
	}

	tx, err := m.retryTransaction(func() (*types.Transaction, error) {
		return contract.CancelWPCSafeTransfer(opts, transferId, extraData)
	})
	if err != nil {
		return
	}

	txHash = tx.Hash()

	if async {
		return
	}

	receipt, err := m.waitAndCheckTransaction(ctx, tx)
	if err != nil {
		return
	}

	return m.getReceiptParams(receipt)
}

func (m *Manager) retryTransaction(transactionFunc func() (*types.Transaction, error)) (tx *types.Transaction, err error) {
	// This "for" statement prevents improper reaction to ErrInsufficientFunds
	// As we replenish eth balance of the t.From before we call the contract we can't be sure that the eth replenishment
	// transaction will succeed (due to the GHOST algorithm)
	// The statement retries contract call several times if the call ends with ErrInsufficientFunds
	// (so it basically "waits" for the eth replenishment tx to be re-mined in the canonical chain)
	// Otherwise it proceeds as normal
	for retries := 0; retries < RecoverableErrRetries; retries++ {
		tx, err = transactionFunc()
		if err != core.ErrInsufficientFunds && err != core.ErrNonceTooLow {
			break
		}
		time.Sleep(RecoverableErrDelay)
	}
	return
}
