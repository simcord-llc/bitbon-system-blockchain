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
	"github.com/simcord-llc/bitbon-system-blockchain/common"
)

type AssetboxCryptoData struct {
	Wallet     []byte `json:"wallet"`
	Passphrase []byte `json:"passphrase"`
}

type PrepareAssetboxesRequest struct {
	Count uint64 `json:"count"`
}

type DeleteAssetboxRequest struct {
	Address    common.Address     `json:"address"` // assetbox address to delete
	CryptoData AssetboxCryptoData `json:"crypto_data"`
}

type SetPublicAssetboxInfoRequest struct {
	Address    common.Address     `json:"address"`
	Alias      string             `json:"alias"`
	IsPublic   bool               `json:"isPublic"`
	ExtraInfo  string             `json:"extraInfo"`
	IsMining   bool               `json:"isMining"`
	ServiceID  string             `json:"service_id"`
	CryptoData AssetboxCryptoData `json:"crypto_data"`
}

type PublicAssetboxInfoResponse struct {
	Address   common.Address `json:"address"`
	Alias     string         `json:"alias"`
	ServiceID string         `json:"serviceId"`
	IsPublic  bool           `json:"isPublic"`
	ExtraInfo string         `json:"extraInfo"`
	IsMining  bool           `json:"isMining"`
}

type ProposeDistributionRequest struct {
	Address      common.Address     `json:"address"`
	CryptoData   AssetboxCryptoData `json:"crypto_data"`
	Distribution map[string]uint64  `json:"distribution"`
}
