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
	"github.com/simcord-llc/bitbon-system-blockchain/common"
	"github.com/simcord-llc/bitbon-system-blockchain/core/types"
)

type RangeBlocksByNumberRequest struct {
	BlockNumberFrom       uint64 `json:"blockNumberFrom"`
	BlockNumberTo         uint64 `json:"blockNumberTo"`
	ReturnFullTransaction bool   `json:"returnFullTransaction"`
}

type BlocksTimePeriodRequest struct {
	From uint64 `json:"from"`
	To   uint64 `json:"to"`
}

type Block struct {
	Hash         common.Hash
	Number       *big.Int       `json:"number"`
	ParentHash   common.Hash    `json:"parentHash"`
	Nonce        uint64         `json:"nonce"`
	Miner        common.Address `json:"miner"`
	MixHash      common.Hash    `json:"mixHash"`
	Difficulty   *big.Int       `json:"difficulty"`
	ExtraData    string         `json:"extraData"`
	Size         string         `json:"size"`
	GasLimit     uint64         `json:"gasLimit"`
	GasUsed      uint64         `json:"gasUsed"`
	Timestamp    uint64         `json:"timestamp"`
	Transactions []*Transaction `json:"transactions"`
}

// BitbonBlockToDto converts clientBlock to dto.Block (proto).
// If txFull is true, dto.Block will contain dto.TransactionObject, else only list transaction hashes.
// If b is nil, return nil
func BitbonBlockToDto(b []*Block, txFull bool) []*external.Block {
	if b == nil {
		return nil
	}

	blocks := make([]*external.Block, len(b))

	for index, block := range b {
		var formatTx = BitbonTransactionToHashesDto
		if txFull {
			formatTx = BitbonTransactionToDto
		}

		block := external.Block{
			Hash:              block.Hash.Hex(),
			Number:            block.Number.Uint64(),
			ParentHash:        block.ParentHash.String(),
			Nonce:             block.Nonce,
			Miner:             block.Miner.String(),
			ExtraData:         block.ExtraData,
			Difficulty:        block.Difficulty.String(),
			Size:              block.Size,
			GasLimit:          block.GasLimit,
			GasUsed:           block.GasUsed,
			Timestamp:         block.Timestamp,
			TransactionResult: formatTx(block.Transactions),
		}

		blocks[index] = &block
	}

	return blocks
}

// ToBitbonBlock converts types.Block (core ethereum) to Block.
// If b is nil, return empty struct
func ToBitbonBlock(b *types.Block, txs []*Transaction) *Block {
	if b == nil {
		return nil
	}

	block := &Block{
		Hash:         b.Hash(),
		Number:       b.Number(),
		ParentHash:   b.ParentHash(),
		Nonce:        b.Nonce(),
		Miner:        b.Coinbase(),
		MixHash:      b.MixDigest(),
		Difficulty:   b.Difficulty(),
		ExtraData:    common.Bytes2Hex(b.Extra()),
		Size:         b.Size().String(),
		GasLimit:     b.GasLimit(),
		GasUsed:      b.GasUsed(),
		Timestamp:    b.Time(),
		Transactions: txs,
	}

	return block
}
