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

package models

import (
	"math/big"

	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/contracts"
	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/dto"
	"github.com/simcord-llc/bitbon-system-blockchain/common"
)

type EmitObj struct {
	Address common.Address
	Value   *big.Int
}

type DirectTransferObj struct {
	From      common.Address
	To        common.Address
	Value     *big.Int
	ExtraData []byte
}

type QuickTransferObj struct {
	From       common.Address
	AccountID  string
	To         common.Address
	Value      *big.Int
	ExtraData  []byte
	CryptoData dto.AssetboxCryptoData
}

type CreateTransferObj struct {
	From           common.Address
	AccountID      string
	To             common.Address
	Value          *big.Int
	TransferID     string
	ProtectionCode string
	Retries        uint64
	ExpiresAt      int64
	Proof          []byte
	VK             []byte
	ExtraData      []byte
	CryptoData     dto.AssetboxCryptoData
}

type ApproveTransferObj struct {
	Address        common.Address
	TransferID     string
	ProtectionCode string
	ExtraData      []byte
	CryptoData     dto.AssetboxCryptoData
}

type CancelTransferObj struct {
	Address    common.Address
	TransferID string
	ExtraData  []byte
	CryptoData dto.AssetboxCryptoData
}

type TransferResponseObj struct {
	BlockNumber uint64      `json:"blockNumber"`
	TxHash      common.Hash `json:"txHash"`
	ExtraData   []byte      `json:"extraData"`
}

func (t *CreateTransferObj) ToContractsTransfer() *contracts.Transfer {
	ct := &contracts.Transfer{
		To:         t.To,
		Value:      t.Value,
		TransferID: []byte(t.TransferID),
		Retries:    new(big.Int).SetUint64(t.Retries),
		ExpiresAt:  new(big.Int).SetInt64(t.ExpiresAt),
		Proof:      t.Proof,
		VK:         t.VK,
		ExtraData:  t.ExtraData,
	}
	return ct
}

func (tr *TransferResponseObj) ToDTO() *dto.TransferResponse {
	o := &dto.TransferResponse{
		BlockNumber: tr.BlockNumber,
		TxHash:      tr.TxHash,
	}
	return o
}

func QuickTransferObjFromDTO(cstor *dto.QuickTransferRequest) *QuickTransferObj {
	o := &QuickTransferObj{
		From:       cstor.From,
		AccountID:  cstor.AccountID,
		To:         cstor.To,
		Value:      cstor.Value,
		ExtraData:  cstor.ExtraData,
		CryptoData: cstor.CryptoData,
	}
	return o
}

func CreateWPCTransferObjFromDTO(cstor *dto.CreateWPCSafeTransferRequest) *CreateTransferObj {
	o := &CreateTransferObj{
		From:       cstor.From,
		AccountID:  cstor.AccountID,
		To:         cstor.To,
		Value:      cstor.Value,
		TransferID: cstor.TransferID,
		ExpiresAt:  cstor.ExpiresAt,
		ExtraData:  cstor.ExtraData,
		CryptoData: cstor.CryptoData,
	}
	return o
}

func ApproveTransferObjFromDTO(astor *dto.ApproveSafeTransferRequest) *ApproveTransferObj {
	o := &ApproveTransferObj{
		Address:        astor.Address,
		TransferID:     astor.TransferID,
		ProtectionCode: astor.ProtectionCode,
		ExtraData:      astor.ExtraData,
		CryptoData:     astor.CryptoData,
	}
	return o
}

func CancelTransferObjFromDTO(cstor *dto.CancelSafeTransferRequest) *CancelTransferObj {
	o := &CancelTransferObj{
		Address:    cstor.Address,
		TransferID: cstor.TransferID,
		ExtraData:  cstor.ExtraData,
		CryptoData: cstor.CryptoData,
	}
	return o
}

func CreateTransferObjFromDTO(cstor *dto.CreateSafeTransferRequest) *CreateTransferObj {
	o := &CreateTransferObj{
		From:           cstor.From,
		AccountID:      cstor.AccountID,
		To:             cstor.To,
		Value:          cstor.Value,
		TransferID:     cstor.TransferID,
		ProtectionCode: cstor.ProtectionCode,
		Retries:        cstor.Retries,
		ExpiresAt:      cstor.ExpiresAt,
		ExtraData:      cstor.ExtraData,
		CryptoData:     cstor.CryptoData,
	}
	return o
}

func ApproveWPCTransferObjFromDTO(astor *dto.ApproveWPCSafeTransferRequest) *ApproveTransferObj {
	o := &ApproveTransferObj{
		Address:    astor.Address,
		TransferID: astor.TransferID,
		ExtraData:  astor.ExtraData,
		CryptoData: astor.CryptoData,
	}
	return o
}

func CancelWPCTransferObjFromDTO(cstor *dto.CancelWPCSafeTransferRequest) *CancelTransferObj {
	o := &CancelTransferObj{
		Address:    cstor.Address,
		TransferID: cstor.TransferID,
		ExtraData:  cstor.ExtraData,
		CryptoData: cstor.CryptoData,
	}
	return o
}

func DirectTransferObjFromDTO(dtr *dto.DirectTransferRequest) *DirectTransferObj {
	o := &DirectTransferObj{
		From:      dtr.From,
		To:        dtr.To,
		Value:     dtr.Value,
		ExtraData: dtr.ExtraData,
	}
	return o
}
