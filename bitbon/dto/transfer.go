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

	"encoding/json"

	"github.com/simcord-llc/bitbon-system-blockchain/common"
	"github.com/simcord-llc/bitbon-system-blockchain/common/hexutil"
)

type QuickTransferRequest struct {
	From       common.Address     `json:"from"`
	To         common.Address     `json:"to"`
	Value      *big.Int           `json:"value"`
	ExtraData  []byte             `json:"extraData"`
	CryptoData AssetboxCryptoData `json:"crypto_data"`
}

func (r *QuickTransferRequest) UnmarshalJSON(data []byte) error {
	type Alias QuickTransferRequest
	aux := &struct {
		Value *hexutil.Big `json:"value"`
		*Alias
	}{
		Alias: (*Alias)(r),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	r.Value = aux.Value.ToInt()
	return nil
}

type FullBalanceQuickTransferRequest struct {
	From       common.Address     `json:"from"`
	To         common.Address     `json:"to"`
	BTSC       string             `json:"btsc"`
	ExtraData  []byte             `json:"extraData"`
	CryptoData AssetboxCryptoData `json:"crypto_data"`
}

type CreateWPCSafeTransferRequest struct {
	From       common.Address     `json:"from"`
	To         common.Address     `json:"to"`
	Value      *big.Int           `json:"value"`
	TransferID string             `json:"transferId"`
	ExpiresAt  int64              `json:"expiresAt"`
	ExtraData  []byte             `json:"extraData"`
	CryptoData AssetboxCryptoData `json:"crypto_data"`
}

func (r *CreateWPCSafeTransferRequest) UnmarshalJSON(data []byte) error {
	type Alias CreateWPCSafeTransferRequest
	aux := &struct {
		Value *hexutil.Big `json:"value"`
		*Alias
	}{
		Alias: (*Alias)(r),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	r.Value = aux.Value.ToInt()
	return nil
}

type ApproveWPCSafeTransferRequest struct {
	Address    common.Address     `json:"address"`
	TransferID string             `json:"transferId"`
	ExtraData  []byte             `json:"extraData"`
	CryptoData AssetboxCryptoData `json:"crypto_data"`
}

type CancelWPCSafeTransferRequest struct {
	Address    common.Address     `json:"address"`
	TransferID string             `json:"transferId"`
	ExtraData  []byte             `json:"extraData"`
	CryptoData AssetboxCryptoData `json:"crypto_data"`
}

type CreateSafeTransferRequest struct {
	From           common.Address     `json:"from"`
	To             common.Address     `json:"to"`
	Value          *big.Int           `json:"value"`
	TransferID     string             `json:"transferId"`
	ProtectionCode string             `json:"protectionCode"`
	Retries        uint64             `json:"retries"`
	ExpiresAt      int64              `json:"expiresAt"`
	ExtraData      []byte             `json:"extraData"`
	CryptoData     AssetboxCryptoData `json:"crypto_data"`
}

func (r *CreateSafeTransferRequest) UnmarshalJSON(data []byte) error {
	type Alias CreateSafeTransferRequest
	aux := &struct {
		Value *hexutil.Big `json:"value"`
		*Alias
	}{
		Alias: (*Alias)(r),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	r.Value = aux.Value.ToInt()
	return nil
}

type ApproveSafeTransferRequest struct {
	Address        common.Address     `json:"address"`
	TransferID     string             `json:"transferId"`
	ProtectionCode string             `json:"protectionCode"`
	ExtraData      []byte             `json:"extraData"`
	CryptoData     AssetboxCryptoData `json:"crypto_data"`
}

type CancelSafeTransferRequest struct {
	Address    common.Address     `json:"address"`
	TransferID string             `json:"transferId"`
	ExtraData  []byte             `json:"extraData"`
	CryptoData AssetboxCryptoData `json:"crypto_data"`
}

type CreateFullBalanceWPCSafeTransferRequest struct {
	From       common.Address     `json:"from"`
	To         common.Address     `json:"to"`
	TransferID string             `json:"transferId"`
	BTSC       string             `json:"btsc"`
	ExpiresAt  int64              `json:"expiresAt"`
	ExtraData  []byte             `json:"extraData"`
	CryptoData AssetboxCryptoData `json:"crypto_data"`
}

func (r *CreateFullBalanceWPCSafeTransferRequest) UnmarshalJSON(data []byte) error {
	type Alias CreateFullBalanceWPCSafeTransferRequest
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(r),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}

type ApproveFullBalanceWPCSafeTransferRequest struct {
	Address    common.Address     `json:"address"`
	TransferID string             `json:"transferId"`
	ExtraData  []byte             `json:"extraData"`
	CryptoData AssetboxCryptoData `json:"crypto_data"`
}

type CancelFullBalanceWPCSafeTransferRequest struct {
	Address    common.Address     `json:"address"`
	TransferID string             `json:"transferId"`
	ExtraData  []byte             `json:"extraData"`
	CryptoData AssetboxCryptoData `json:"crypto_data"`
}

type CreateFullBalanceSafeTransferRequest struct {
	From           common.Address     `json:"from"`
	To             common.Address     `json:"to"`
	TransferID     string             `json:"transferId"`
	ProtectionCode string             `json:"protectionCode"`
	BTSC           string             `json:"btsc"`
	Retries        uint64             `json:"retries"`
	ExpiresAt      int64              `json:"expiresAt"`
	ExtraData      []byte             `json:"extraData"`
	CryptoData     AssetboxCryptoData `json:"crypto_data"`
}

func (r *CreateFullBalanceSafeTransferRequest) UnmarshalJSON(data []byte) error {
	type Alias CreateFullBalanceSafeTransferRequest
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(r),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}

type ApproveFullBalanceSafeTransferRequest struct {
	Address        common.Address     `json:"address"`
	TransferID     string             `json:"transferId"`
	ProtectionCode string             `json:"protectionCode"`
	ExtraData      []byte             `json:"extraData"`
	CryptoData     AssetboxCryptoData `json:"crypto_data"`
}

type CancelFullBalanceSafeTransferRequest struct {
	Address    common.Address     `json:"address"`
	TransferID string             `json:"transferId"`
	ExtraData  []byte             `json:"extraData"`
	CryptoData AssetboxCryptoData `json:"crypto_data"`
}

type DirectTransferRequest struct {
	From      common.Address `json:"from"`
	To        common.Address `json:"to"`
	Value     *big.Int       `json:"value"`
	ExtraData []byte         `json:"extraData"`
}

func (r *DirectTransferRequest) UnmarshalJSON(data []byte) error {
	type Alias DirectTransferRequest
	aux := &struct {
		Value *hexutil.Big `json:"value"`
		*Alias
	}{
		Alias: (*Alias)(r),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	r.Value = aux.Value.ToInt()
	return nil
}

type MonetizeRequest struct {
	Address common.Address `json:"address"`
	Value   *big.Int       `json:"value"`
}

func (r *MonetizeRequest) UnmarshalJSON(data []byte) error {
	type Alias MonetizeRequest
	aux := &struct {
		Value *hexutil.Big `json:"value"`
		*Alias
	}{
		Alias: (*Alias)(r),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	r.Value = aux.Value.ToInt()
	return nil
}

type EmitBitbonRequest struct {
	Address common.Address `json:"address"`
	Value   *big.Int       `json:"value"`
}

func (r *EmitBitbonRequest) UnmarshalJSON(data []byte) error {
	type Alias EmitBitbonRequest
	aux := &struct {
		Value *hexutil.Big `json:"value"`
		*Alias
	}{
		Alias: (*Alias)(r),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	r.Value = aux.Value.ToInt()
	return nil
}

type TransferResponse struct {
	BlockNumber uint64      `json:"blockNumber"`
	TxHash      common.Hash `json:"txHash"`
}

type ExpireTransferResponse struct {
	TxHashes   []common.Hash `json:"txHashes"`
	ExpiredNum int           `json:"expiredNum"`
}

type ServiceFeeTransferRequest struct {
	From          common.Address `json:"from"`
	OperationType *big.Int       `json:"operationType"`
}
