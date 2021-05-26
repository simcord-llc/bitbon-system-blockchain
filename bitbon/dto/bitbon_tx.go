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

// nolint:golint,stylecheck
package dto

import (
	"encoding/json"
	"math/big"
	"strings"

	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/dto/external"
	bitbonErrors "github.com/simcord-llc/bitbon-system-blockchain/bitbon/errors"
	"github.com/simcord-llc/bitbon-system-blockchain/common"
	"github.com/simcord-llc/bitbon-system-blockchain/core/types"
)

const (
	DirectTransfer         = "transferFrom"
	QuickTransfer          = "quickTransfer"
	CreateSafeTransfer     = "createSafeTransfer"
	ApproveSafeTransfer    = "approveSafeTransfer"
	CancelSafeTransfer     = "cancelSafeTransfer"
	ExpireSafeTransfer     = "expireSafeTransfers"
	CreateWPCSafeTransfer  = "createWPCSafeTransfer"
	ApproveWPCSafeTransfer = "approveWPCSafeTransfer"
	CancelWPCSafeTransfer  = "cancelWPCSafeTransfer"
)

const (
	// v1
	DirectTransferV1MethodID      = "23b872dd"
	CreateSafeTransferV1MethodID  = "d172317b"
	ApproveSafeTransferV1MethodID = "d1408c8d"
	CancelSafeTransferV1MethodID  = "8cac0630"
	MonetizeEmissionV1MethodID    = "6b88cfc1"
	MonetizeCertificateV1MethodID = "0ceb4c53"

	// v2
	DirectTransferV2MethodID      = "ab67aa58"
	CreateSafeTransferV2MethodID  = "6c9bcf6f"
	ApproveSafeTransferV2MethodID = "fd821f3f"
	CancelSafeTransferV2MethodID  = "a8389ace"
	ExpireSafeTransferV2MethodID  = "a5471f24"
	MonetizeEmissionV2MethodID    = "6b88cfc1"
	MonetizeCertificateV2MethodID = "0ceb4c53"

	// v3
	DirectTransferV3MethodID         = "ab67aa58"
	QuickTransferV3MethodID          = "536684e7"
	CreateSafeTransferV3MethodID     = "6c9bcf6f"
	ApproveSafeTransferV3MethodID    = "fd821f3f"
	CancelSafeTransferV3MethodID     = "a8389ace"
	ExpireSafeTransferV3MethodID     = "a5471f24"
	CreateWPCSafeTransferV3MethodID  = "9ca9dfa7"
	ApproveWPCSafeTransferV3MethodID = "c48bb729"
	CancelWPCSafeTransferV3MethodID  = "28b5f788"
)

type EventType string

const (
	BalanceChangedEventType  EventType = "BalanceChanged"
	BalanceLockedEventType   EventType = "BalanceLocked"
	BalanceUnLockedEventType EventType = "BalanceUnLocked"
	TransferExpiredEventType EventType = "TransferExpired"
)

var eventTypeMap = map[EventType]struct{}{
	BalanceChangedEventType:  {},
	BalanceLockedEventType:   {},
	BalanceUnLockedEventType: {},
	TransferExpiredEventType: {},
}

func EventTypeExistsString(eventType string) bool {
	_, ok := eventTypeMap[EventType(eventType)]

	return ok
}

func EventTypeExists(eventType EventType) bool {
	_, ok := eventTypeMap[eventType]

	return ok
}

type ContractVersion string

const (
	VersionV1 ContractVersion = "v1"
	VersionV2 ContractVersion = "v2"
	VersionV3 ContractVersion = "v3"
)

var contractTypeMap = map[ContractVersion]struct{}{
	VersionV1: {},
	VersionV2: {},
	VersionV3: {},
}

func ContractVersionExistsString(contractVersion string) bool {
	_, ok := contractTypeMap[ContractVersion(contractVersion)]

	return ok
}

func ContractVersionExists(contractVersion ContractVersion) bool {
	_, ok := contractTypeMap[contractVersion]

	return ok
}

type ParseStatus uint8

const (
	ParseStatusSuccess         ParseStatus = 0
	ParseStatusParseError      ParseStatus = 1
	ParseStatusParseInputError ParseStatus = 2
	ParseStatusParseLogsError  ParseStatus = 3
)

type BitbonArgs interface {
	ToExternalDTO() *external.BitbonArgs
}

type BitbonTx struct {
	Status      bool                        `json:"status"`
	ParseStatus ParseStatus                 `json:"parseStatus"`
	ParseError  error                       `json:"-"`
	TxError     bitbonErrors.Error          `json:"-"`
	Method      string                      `json:"method"`
	Args        BitbonArgs                  `json:"args"`
	Balances    map[common.Address]*big.Int `json:"balances"`
	Sender      common.Address              `json:"sender"`
	Events      []CustomEvent               `json:"events"`
}

func (t *BitbonTx) MarshalJSON() ([]byte, error) {
	var parseErrStr string
	if t.ParseError != nil {
		parseErrStr = t.ParseError.Error()
	}

	var txErr interface{}
	if t.TxError != nil {
		txErr = struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		}{
			Code:    t.TxError.ErrorCode(),
			Message: t.TxError.Error(),
		}
	}

	type Alias BitbonTx
	return json.Marshal(&struct {
		ParseError string      `json:"parseError"`
		TxError    interface{} `json:"txError"`
		*Alias
	}{
		ParseError: parseErrStr,
		TxError:    txErr,
		Alias:      (*Alias)(t),
	})
}

// ApproveTransferEvent Duplicated code for reflection parse input in transactions
type ApproveTransferEvent struct {
	TransferId []byte
	Status     uint8
}

func (arg *ApproveTransferEvent) MarshalJSON() ([]byte, error) {
	type Alias ApproveTransferEvent

	return json.Marshal(&struct {
		TransferId string `json:"transferId"`
		*Alias
	}{
		TransferId: string(arg.TransferId),
		Alias:      (*Alias)(arg),
	})
}

// CancelApprovalArgs For cancel and approve tx
type CancelApprovalArgs struct {
	Address    common.Address `json:"address"`
	TransferId []byte         `json:"transferId"`
	ExtraData  []byte         `json:"extraData"`
}

func (arg *CancelApprovalArgs) MarshalJSON() ([]byte, error) {
	type Alias CancelApprovalArgs

	return json.Marshal(&struct {
		TransferId string `json:"transferId"`
		*Alias
	}{
		TransferId: string(arg.TransferId),
		Alias:      (*Alias)(arg),
	})
}

func (arg *CancelApprovalArgs) ToExternalDTO() *external.BitbonArgs {
	return &external.BitbonArgs{
		Args: &external.BitbonArgs_CancelApprovalArgs{
			CancelApprovalArgs: &external.CancelApprovalArgs{
				TransferId: arg.TransferId,
				Address:    strings.ToLower(arg.Address.Hex()),
				ExtraData:  arg.ExtraData,
			},
		},
	}
}

// For Expire Safe Transfers
type ExpireSafeTransfersArgs struct {
	Ids [][32]byte `json:"ids"`
}

func (arg *ExpireSafeTransfersArgs) MarshalJSON() ([]byte, error) {
	type Alias ExpireSafeTransfersArgs

	res := make([]string, len(arg.Ids))
	for i, tr := range arg.Ids {
		res[i] = common.Bytes2Hex(tr[:])
	}

	return json.Marshal(&struct {
		Ids []string `json:"ids"`
		*Alias
	}{
		Ids:   res,
		Alias: (*Alias)(arg),
	})
}

func (arg *ExpireSafeTransfersArgs) ToExternalDTO() *external.BitbonArgs {
	res := make([]string, len(arg.Ids))
	for i, tr := range arg.Ids {
		res[i] = strings.ToLower(common.Bytes2Hex(tr[:]))
	}

	return &external.BitbonArgs{
		Args: &external.BitbonArgs_ExpireSafeTransferArgs{
			ExpireSafeTransferArgs: &external.ExpireSafeTransferArgs{
				TransferIds: res,
			},
		},
	}
}

// For SafeTransfer
type SafeTransferArgs struct {
	From       common.Address `json:"from"`
	To         common.Address `json:"to"`
	TransferId []byte         `json:"transferId"`
	Value      *big.Int       `json:"value"`
	ExtraData  []byte         `json:"extraData"`
}

func (arg *SafeTransferArgs) MarshalJSON() ([]byte, error) {
	type Alias SafeTransferArgs

	return json.Marshal(&struct {
		TransferId string `json:"transferId"`
		*Alias
	}{
		TransferId: string(arg.TransferId),
		Alias:      (*Alias)(arg),
	})
}

func (arg *SafeTransferArgs) ToExternalDTO() *external.BitbonArgs {
	return &external.BitbonArgs{
		Args: &external.BitbonArgs_SafeTransferArgs{
			SafeTransferArgs: &external.SafeTransferArgs{
				From:       strings.ToLower(arg.From.Hex()),
				To:         strings.ToLower(arg.To.Hex()),
				Value:      arg.Value.String(),
				TransferId: arg.TransferId,
				ExtraData:  arg.ExtraData,
			},
		},
	}
}

// For DirectTransfer
type DirectTransferArgs struct {
	From      common.Address `json:"from"`
	To        common.Address `json:"to"`
	Value     *big.Int       `json:"value"`
	ExtraData []byte         `json:"extraData"`
}

func (arg *DirectTransferArgs) ToExternalDTO() *external.BitbonArgs {
	return &external.BitbonArgs{
		Args: &external.BitbonArgs_DirectTransferArgs{
			DirectTransferArgs: &external.DirectTransferArgs{
				From:      strings.ToLower(arg.From.Hex()),
				To:        strings.ToLower(arg.To.Hex()),
				Value:     arg.Value.String(),
				ExtraData: arg.ExtraData,
			},
		},
	}
}

// For QuickTransfer
type QuickTransferArgs struct {
	From      common.Address `json:"from"`
	To        common.Address `json:"to"`
	Value     *big.Int       `json:"value"`
	ExtraData []byte         `json:"extraData"`
}

func (arg *QuickTransferArgs) ToExternalDTO() *external.BitbonArgs {
	return &external.BitbonArgs{
		Args: &external.BitbonArgs_QuickTransferArgs{
			QuickTransferArgs: &external.QuickTransferArgs{
				From:      strings.ToLower(arg.From.Hex()),
				To:        strings.ToLower(arg.To.Hex()),
				Value:     arg.Value.String(),
				ExtraData: arg.ExtraData,
			},
		},
	}
}

// For monetize txs
type BitbonMonetizeArgs struct {
	Assetbox common.Address `json:"assetbox"`
	Amount   *big.Int       `json:"amount"`
}

func (arg *BitbonMonetizeArgs) ToExternalDTO() *external.BitbonArgs {
	return &external.BitbonArgs{}
}

type BitbonBalanceChanged struct {
	Assetbox     common.Address `json:"assetbox"`
	Balance      *big.Int       `json:"balance"`
	AssetboxType *big.Int       `json:"assetboxType"`
	Raw          types.Log      `json:"raw"`
}

type BitbonBalanceLocked struct {
	Assetbox       common.Address `json:"assetbox"`
	BalanceAviable *big.Int       `json:"balanceAviable"`
	BalanceLocked  *big.Int       `json:"balanceLocked"`
	AssetboxType   *big.Int       `json:"assetboxType"`
	Raw            types.Log      `json:"raw"`
}

type BitbonBalanceUnlocked struct {
	Assetbox       common.Address `json:"assetbox"`
	BalanceAviable *big.Int       `json:"balanceAviable"`
	BalanceLocked  *big.Int       `json:"balanceLocked"`
	AssetboxType   *big.Int       `json:"assetboxType"`
	Raw            types.Log      `json:"raw"`
}

type CustomEvent interface {
	GetType() EventType
	MarshalJSON() ([]byte, error)
}

type TransferExpiredEvent struct {
	TransferId []byte `json:"transferId"`
}

func (arg *TransferExpiredEvent) GetType() EventType {
	return TransferExpiredEventType
}

func (arg *TransferExpiredEvent) MarshalJSON() ([]byte, error) {
	type Alias TransferExpiredEvent

	return json.Marshal(&struct {
		TransferId string `json:"transferId"`
		EventType  string `json:"eventType"`
		*Alias
	}{
		TransferId: string(arg.TransferId),
		EventType:  string(arg.GetType()),
		Alias:      (*Alias)(arg),
	})
}
