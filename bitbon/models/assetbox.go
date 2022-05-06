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
	"crypto/ecdsa"
	"time"

	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/contracts"
	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/dto"
	"github.com/simcord-llc/bitbon-system-blockchain/common"
)

type ShortAssetbox struct {
	Address    common.Address `json:"address"`
	Wallet     []byte         `json:"wallet" `
	PassPhrase []byte         `json:"passPhrase"`
}

type Assetbox struct {
	Address    common.Address `json:"address"`
	Alias      string         `json:"alias"`
	ServiceID  string         `json:"serviceID"`
	CreatedAt  *time.Time     `json:"createdAt"`
	UpdatedAt  *time.Time     `json:"updatedAt"`
	DeletedAt  *time.Time     `json:"deletedAt"`
	IsPublic   bool           `json:"isPublic"`
	IsMining   bool           `json:"isMining"`
	ExtraInfo  string         `json:"extraInfo"`
	Wallet     []byte         `json:"wallet" `
	PassPhrase []byte         `json:"passPhrase"`

	Pk *ecdsa.PrivateKey `json:"-"` // nolint, assetbox private key
}

func (a *Assetbox) PopulateContractData(ca *contracts.Assetbox) {
	a.Alias = ca.Alias
	a.ExtraInfo = ca.ExtraInfo
	a.IsPublic = ca.IsPublic
	a.IsMining = ca.IsMining
}

func (a *Assetbox) ToContractsAssetbox() *contracts.Assetbox {
	ca := &contracts.Assetbox{
		Address:   a.Address,
		Alias:     a.Alias,
		ServiceID: []byte(a.ServiceID),
		IsPublic:  a.IsPublic,
		IsMining:  a.IsMining,
	}
	if ca.IsPublic {
		ca.ExtraInfo = a.ExtraInfo
	}
	return ca
}

func ContractAssetboxToDTO(a *contracts.Assetbox) *dto.PublicAssetboxInfoResponse {
	fa := &dto.PublicAssetboxInfoResponse{
		Address:   a.Address,
		Alias:     a.Alias,
		ServiceID: string(a.ServiceID),
		IsPublic:  a.IsPublic,
		IsMining:  a.IsMining,
		ExtraInfo: a.ExtraInfo,
	}
	return fa
}
