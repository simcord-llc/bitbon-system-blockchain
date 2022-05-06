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
	To         common.Address
	Value      *big.Int
	ExtraData  []byte
	CryptoData dto.AssetboxCryptoData
}

type CreateTransferObj struct {
	From           common.Address
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

type ServiceFeeTransferObj struct {
	From          common.Address
	OperationType *big.Int
}

func FeeTransferObjFromDTO(cstor *dto.ServiceFeeTransferRequest) *ServiceFeeTransferObj {
	return &ServiceFeeTransferObj{
		From:          cstor.From,
		OperationType: cstor.OperationType,
	}
}

func (t *CreateTransferObj) ToContractsTransfer() *contracts.Transfer {
	return &contracts.Transfer{
		To:         t.To,
		Value:      t.Value,
		TransferID: []byte(t.TransferID),
		Retries:    new(big.Int).SetUint64(t.Retries),
		ExpiresAt:  new(big.Int).SetInt64(t.ExpiresAt),
		Proof:      t.Proof,
		VK:         t.VK,
		ExtraData:  t.ExtraData,
	}
}

func (tr *TransferResponseObj) ToDTO() *dto.TransferResponse {
	return &dto.TransferResponse{
		BlockNumber: tr.BlockNumber,
		TxHash:      tr.TxHash,
	}
}

func QuickTransferObjFromDTO(cstor *dto.QuickTransferRequest) *QuickTransferObj {
	return &QuickTransferObj{
		From:       cstor.From,
		To:         cstor.To,
		Value:      cstor.Value,
		ExtraData:  cstor.ExtraData,
		CryptoData: cstor.CryptoData,
	}
}

func CreateWPCTransferObjFromDTO(cstor *dto.CreateWPCSafeTransferRequest) *CreateTransferObj {
	return &CreateTransferObj{
		From:       cstor.From,
		To:         cstor.To,
		Value:      cstor.Value,
		TransferID: cstor.TransferID,
		ExpiresAt:  cstor.ExpiresAt,
		ExtraData:  cstor.ExtraData,
		CryptoData: cstor.CryptoData,
	}
}

func ApproveTransferObjFromDTO(astor *dto.ApproveSafeTransferRequest) *ApproveTransferObj {
	return &ApproveTransferObj{
		Address:        astor.Address,
		TransferID:     astor.TransferID,
		ProtectionCode: astor.ProtectionCode,
		ExtraData:      astor.ExtraData,
		CryptoData:     astor.CryptoData,
	}
}

func CancelTransferObjFromDTO(cstor *dto.CancelSafeTransferRequest) *CancelTransferObj {
	return &CancelTransferObj{
		Address:    cstor.Address,
		TransferID: cstor.TransferID,
		ExtraData:  cstor.ExtraData,
		CryptoData: cstor.CryptoData,
	}
}

func CreateTransferObjFromDTO(cstor *dto.CreateSafeTransferRequest) *CreateTransferObj {
	return &CreateTransferObj{
		From:           cstor.From,
		To:             cstor.To,
		Value:          cstor.Value,
		TransferID:     cstor.TransferID,
		ProtectionCode: cstor.ProtectionCode,
		Retries:        cstor.Retries,
		ExpiresAt:      cstor.ExpiresAt,
		ExtraData:      cstor.ExtraData,
		CryptoData:     cstor.CryptoData,
	}
}

func ApproveWPCTransferObjFromDTO(astor *dto.ApproveWPCSafeTransferRequest) *ApproveTransferObj {
	return &ApproveTransferObj{
		Address:    astor.Address,
		TransferID: astor.TransferID,
		ExtraData:  astor.ExtraData,
		CryptoData: astor.CryptoData,
	}
}

func CancelWPCTransferObjFromDTO(cstor *dto.CancelWPCSafeTransferRequest) *CancelTransferObj {
	return &CancelTransferObj{
		Address:    cstor.Address,
		TransferID: cstor.TransferID,
		ExtraData:  cstor.ExtraData,
		CryptoData: cstor.CryptoData,
	}
}

func DirectTransferObjFromDTO(dtr *dto.DirectTransferRequest) *DirectTransferObj {
	return &DirectTransferObj{
		From:      dtr.From,
		To:        dtr.To,
		Value:     dtr.Value,
		ExtraData: dtr.ExtraData,
	}
}

func FullBalanceQuickTransferObjFromDTO(cstor *dto.FullBalanceQuickTransferRequest) *QuickTransferObj {
	return &QuickTransferObj{
		From:       cstor.From,
		To:         cstor.To,
		ExtraData:  cstor.ExtraData,
		CryptoData: cstor.CryptoData,
	}
}

func CreateFullBalanceWPCTransferObjFromDTO(cstor *dto.CreateFullBalanceWPCSafeTransferRequest) *CreateTransferObj {
	return &CreateTransferObj{
		From:       cstor.From,
		To:         cstor.To,
		TransferID: cstor.TransferID,
		ExpiresAt:  cstor.ExpiresAt,
		ExtraData:  cstor.ExtraData,
		CryptoData: cstor.CryptoData,
	}
}

func ApproveFullBalanceTransferObjFromDTO(astor *dto.ApproveFullBalanceSafeTransferRequest) *ApproveTransferObj {
	return &ApproveTransferObj{
		Address:        astor.Address,
		TransferID:     astor.TransferID,
		ProtectionCode: astor.ProtectionCode,
		ExtraData:      astor.ExtraData,
		CryptoData:     astor.CryptoData,
	}
}

func CancelFullBalanceTransferObjFromDTO(cstor *dto.CancelFullBalanceSafeTransferRequest) *CancelTransferObj {
	return &CancelTransferObj{
		Address:    cstor.Address,
		TransferID: cstor.TransferID,
		ExtraData:  cstor.ExtraData,
		CryptoData: cstor.CryptoData,
	}
}

func CreateFullBalanceTransferObjFromDTO(cstor *dto.CreateFullBalanceSafeTransferRequest) *CreateTransferObj {
	return &CreateTransferObj{
		From:           cstor.From,
		To:             cstor.To,
		TransferID:     cstor.TransferID,
		ProtectionCode: cstor.ProtectionCode,
		Retries:        cstor.Retries,
		ExpiresAt:      cstor.ExpiresAt,
		ExtraData:      cstor.ExtraData,
		CryptoData:     cstor.CryptoData,
	}
}

func ApproveFullBalanceWPCTransferObjFromDTO(astor *dto.ApproveFullBalanceWPCSafeTransferRequest) *ApproveTransferObj {
	return &ApproveTransferObj{
		Address:    astor.Address,
		TransferID: astor.TransferID,
		ExtraData:  astor.ExtraData,
		CryptoData: astor.CryptoData,
	}
}

func CancelFullBalanceWPCTransferObjFromDTO(cstor *dto.CancelFullBalanceWPCSafeTransferRequest) *CancelTransferObj {
	return &CancelTransferObj{
		Address:    cstor.Address,
		TransferID: cstor.TransferID,
		ExtraData:  cstor.ExtraData,
		CryptoData: cstor.CryptoData,
	}
}

func (tr *TransferResponseObj) DTO() *dto.TransferResponse {
	return &dto.TransferResponse{
		BlockNumber: tr.BlockNumber,
		TxHash:      tr.TxHash,
	}
}
