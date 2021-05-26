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

package dto

import (
	"math/big"

	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/dto/external"
	"github.com/simcord-llc/bitbon-system-blockchain/core/types"

	"github.com/simcord-llc/bitbon-system-blockchain/common"
)

type TransactionByBlockNumberAndIndexRequest struct {
	BlockNumber uint64 `json:"blockNumber"`
	TxIndex     uint64 `json:"txIndex"`
}

type TransactionByBlockHashAndIndexRequest struct {
	BlockHash common.Hash `json:"blockHash"`
	TxIndex   uint64      `json:"txIndex"`
}

type TransactionTimePeriodRequest struct {
	From uint64 `json:"from"`
	To   uint64 `json:"to"`
}

type BlockInfo struct {
	Hash        common.Hash
	BlockNumber uint64
	BlockTime   uint64
}

type Transaction struct {
	Hash        common.Hash `json:"hash"`
	Nonce       uint64      `json:"nonce"`
	BlockHash   common.Hash `json:"blockHash"`
	BlockNumber uint64      `json:"blockNumber"`
	BlockTime   uint64      `json:"blockTime"`
	GasPrice    *big.Int    `json:"gasPrice"`
	Gas         uint64      `json:"gasLimit"`
	Time        uint64      `json:"creates"`
	Size        string      `json:"size"`
	BitbonTx    *BitbonTx   `json:"bitbonTx"`
}

// BitbonTransactionToDto converts []clientTransaction to []*dto.TransactionResult (proto).
// Contains all fields dto.TransactionResult.Transaction
func BitbonTransactionToDto(t []*Transaction) []*external.TransactionResult {
	transactions := make([]*external.TransactionResult, len(t))

	for index, transaction := range t {
		var transactionObj = &external.TransactionObject{
			Hash:        transaction.Hash.Hex(),
			Nonce:       transaction.Nonce,
			BlockHash:   transaction.BlockHash.String(),
			BlockNumber: transaction.BlockNumber,
			BlockTime:   transaction.BlockTime,
			GasPrice:    transaction.GasPrice.Uint64(),
			GasLimit:    transaction.Gas,
			Time:        transaction.Time,
			Size:        transaction.Size,
		}

		bitbonTx := transaction.BitbonTx

		if bitbonTx != nil {
			var method = GetMethod(bitbonTx.Method)
			var parseStatus = int64(bitbonTx.ParseStatus)

			transactionObj.BitbonTx = &external.BitbonTransaction{
				Status:      bitbonTx.Status,
				ParseStatus: parseStatus,
				Method:      method,
				Sender:      bitbonTx.Sender.Hex(),
			}

			if bitbonTx.ParseError != nil {
				transactionObj.BitbonTx.ParseError = bitbonTx.ParseError.Error()
			}

			if bitbonTx.Args != nil {
				transactionObj.BitbonTx.Args = bitbonTx.Args.ToExternalDTO()
			}

			if bitbonTx.TxError != nil {
				errorCode := int64(bitbonTx.TxError.ErrorCode())

				transactionObj.BitbonTx.TxError = &external.Error{
					Code:      errorCode,
					Message:   bitbonTx.TxError.Error(),
					DataError: nil,
				}
			}

			if bitbonTx.Balances != nil {
				transactionObj.BitbonTx.Balances = make([]*external.Balance, 0, len(bitbonTx.Balances))
				for address, value := range bitbonTx.Balances {
					balance := &external.Balance{
						Address: address.Hex(),
						Balance: value.String(),
					}
					transactionObj.BitbonTx.Balances = append(transactionObj.BitbonTx.Balances, balance)
				}
			}
		}

		transactions[index] = &external.TransactionResult{
			Transaction: &external.TransactionResult_TxObj{
				TxObj: transactionObj,
			},
		}
	}
	return transactions
}

func GetMethod(methodName string) external.Method {
	switch methodName {
	case CreateSafeTransfer:
		return external.Method_CREATE_SAFE_TRANSFER
	case ApproveSafeTransfer:
		return external.Method_APPROVE_SAFE_TRANSFER
	case CancelSafeTransfer:
		return external.Method_CANCEL_SAFE_TRANSFER

	case DirectTransfer:
		return external.Method_DIRECT_TRANSFER
	case QuickTransfer:
		return external.Method_QUICK_TRANSFER

	case CreateWPCSafeTransfer:
		return external.Method_CREATE_WPC_SAFE_TRANSFER
	case ApproveWPCSafeTransfer:
		return external.Method_APPROVE_WPC_SAFE_TRANSFER
	case CancelWPCSafeTransfer:
		return external.Method_CANCEL_WPC_SAFE_TRANSFER

	case ExpireSafeTransfer:
		return external.Method_EXPIRE_TRANSFER

	default:
		return external.Method_UNDEFINED
	}
}

// BitbonTransactionToHashesDto converts []clientTransaction to []*dto.TransactionResult (proto).
// Contains only txHashes
func BitbonTransactionToHashesDto(t []*Transaction) []*external.TransactionResult {
	transactions := make([]*external.TransactionResult, len(t))

	for index, transaction := range t {
		transactions[index] = &external.TransactionResult{
			Transaction: &external.TransactionResult_Hash{
				Hash: transaction.Hash.Hex(),
			},
		}
	}

	return transactions
}

func ToTransaction(t *types.Transaction, bb *BitbonTx, bi *BlockInfo) *Transaction {
	var clientTx = Transaction{
		Hash:        t.Hash(),
		Nonce:       t.Nonce(),
		BlockHash:   bi.Hash,
		BlockNumber: bi.BlockNumber,
		BlockTime:   bi.BlockTime,
		GasPrice:    t.GasPrice(),
		Gas:         t.Gas(),
		Time:        t.Time(),
		Size:        t.Size().String(),
	}
	if bb == nil {
		return &clientTx
	}

	clientTx.BitbonTx = bb

	return &clientTx
}
