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

// nolint:nakedret,dupl
package client

import (
	"context"
	"math/big"
	"strings"

	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/dto/external"

	"fmt"

	"encoding/json"

	"github.com/golang/protobuf/proto" // nolint:staticcheck
	"github.com/pkg/errors"
	loggerContext "github.com/simcord-llc/bitbon-system-blockchain/bitbon/context"
	apiDto "github.com/simcord-llc/bitbon-system-blockchain/bitbon/dto"
	"github.com/simcord-llc/bitbon-system-blockchain/common"
	"github.com/simcord-llc/bitbon-system-blockchain/log"
)

func (c *Client) handleQuickTransfer(body []byte) (protoMessage protoMessageWithID, err error) {
	req := &external.QuickTransferRequest{}
	err = proto.Unmarshal(body, req)
	if err != nil {
		err = errors.Wrap(err, "failed to unmarshal request")
		return
	}

	response := &external.TransferResponse{
		Id: req.Id,
	}

	logger := log.New("requestId", req.GetId())
	defer func() {
		if err != nil {
			jsonReq, _ := json.Marshal(req)
			logger.Warn("handleQuickTransfer method ended with error", "request", string(jsonReq), "error", err)

			response.Error = c.handleError(err)
			err = errors.Wrapf(err, "handle request (%s) error", req.GetId())
		}

		protoMessage = response
	}()

	value, ok := new(big.Int).SetString(req.GetValue(), 10)
	if !ok {
		err = errors.New("failed to create big int from string")
		return
	}

	apiRequest := &apiDto.QuickTransferRequest{
		From:      common.HexToAddress(req.GetFrom()),
		To:        common.HexToAddress(req.GetTo()),
		Value:     value,
		ExtraData: req.GetExtraData(),
		CryptoData: apiDto.AssetboxCryptoData{
			Wallet:     req.GetFromCryptoData().GetWallet(),
			Passphrase: req.GetFromCryptoData().GetPassphrase(),
		},
	}

	apiResponse, err := c.bitbon.GetAPI().QuickTransfer(
		loggerContext.NewLoggerContext(context.Background(), logger), apiRequest)
	if err != nil {
		err = errors.Wrap(err, "error calling api method")
		return
	}

	response.TxHash = strings.ToLower(apiResponse.TxHash.Hex())
	response.BlockNumber = fmt.Sprintf("%d", apiResponse.BlockNumber)
	logger.Debug("handleQuickTransfer method ended with success",
		"TxHash", response.TxHash, "BlockNumber", response.BlockNumber)
	return
}

func (c *Client) handleFrameTransfer(body []byte) (protoMessage protoMessageWithID, err error) {
	req := &external.QuickTransferRequest{}
	err = proto.Unmarshal(body, req)
	if err != nil {
		err = errors.Wrap(err, "failed to unmarshal request")
		return
	}

	response := &external.TransferResponse{
		Id: req.Id,
	}

	logger := log.New("requestId", req.GetId())
	defer func() {
		if err != nil {
			jsonReq, _ := json.Marshal(req)
			logger.Warn("handleFrameTransfer method ended with error", "request", string(jsonReq), "error", err)

			response.Error = c.handleError(err)
			err = errors.Wrapf(err, "handle request (%s) error", req.GetId())
		}

		protoMessage = response
	}()

	value, ok := new(big.Int).SetString(req.GetValue(), 10)
	if !ok {
		err = errors.New("failed to create big int from string")
		return
	}

	apiRequest := &apiDto.QuickTransferRequest{
		From:      common.HexToAddress(req.GetFrom()),
		To:        common.HexToAddress(req.GetTo()),
		Value:     value,
		ExtraData: req.GetExtraData(),
		CryptoData: apiDto.AssetboxCryptoData{
			Wallet:     req.GetFromCryptoData().GetWallet(),
			Passphrase: req.GetFromCryptoData().GetPassphrase(),
		},
	}

	apiResponse, err := c.bitbon.GetAPI().FrameTransfer(
		loggerContext.NewLoggerContext(context.Background(), logger), apiRequest)
	if err != nil {
		err = errors.Wrap(err, "error calling api method")
		return
	}

	response.TxHash = strings.ToLower(apiResponse.TxHash.Hex())
	response.BlockNumber = fmt.Sprintf("%d", apiResponse.BlockNumber)
	logger.Debug("handleFrameTransfer method ended with success",
		"TxHash", response.TxHash, "BlockNumber", response.BlockNumber)
	return
}

func (c *Client) handleServiceFeeTransfer(body []byte) (protoMessage protoMessageWithID, err error) {
	req := &external.ServiceFeeTransferRequest{}
	err = proto.Unmarshal(body, req)
	if err != nil {
		err = errors.Wrap(err, "failed to unmarshal request")
		return
	}

	response := &external.TransferResponse{
		Id: req.Id,
	}

	logger := log.New("requestId", req.GetId())
	defer func() {
		if err != nil {
			jsonReq, _ := json.Marshal(req)
			logger.Warn("handleServiceFeeTransfer method ended with error", "request", string(jsonReq), "error", err)

			response.Error = c.handleError(err)
			err = errors.Wrapf(err, "handle request (%s) error", req.GetId())
		}

		protoMessage = response
	}()

	apiRequest := &apiDto.ServiceFeeTransferRequest{
		From:          common.HexToAddress(req.GetFrom()),
		OperationType: big.NewInt(int64(req.OperationType)),
	}

	apiResponse, err := c.bitbon.GetAPI().ServiceFeeTransfer(loggerContext.NewLoggerContext(context.Background(), logger), apiRequest)
	if err != nil {
		err = errors.Wrap(err, "error calling api method")
		return
	}

	response.TxHash = strings.ToLower(apiResponse.TxHash.Hex())
	response.BlockNumber = fmt.Sprintf("%d", apiResponse.BlockNumber)
	logger.Warn("handleServiceFeeTransfer method ended with success", "TxHash", response.TxHash, "BlockNumber", response.BlockNumber)
	return
}

func (c *Client) handleCreateWPCSafeTransfer(body []byte) (protoMessage protoMessageWithID, err error) {
	req := &external.CreateWPCSafeTransferRequest{}
	err = proto.Unmarshal(body, req)
	if err != nil {
		err = errors.Wrap(err, "failed to unmarshal request")
		return
	}

	response := &external.TransferResponse{
		Id: req.Id,
	}

	logger := log.New("requestId", req.GetId())
	defer func() {
		if err != nil {
			jsonReq, _ := json.Marshal(req)
			logger.Warn("handleCreateWPCSafeTransfer method ended with error", "request", string(jsonReq), "error", err)

			response.Error = c.handleError(err)
			err = errors.Wrapf(err, "handle request (%s) error", req.GetId())
		}

		protoMessage = response
	}()

	value, ok := new(big.Int).SetString(req.GetValue(), 10)
	if !ok {
		err = errors.New("failed to create big int from string")
		return
	}

	apiRequest := &apiDto.CreateWPCSafeTransferRequest{
		From:       common.HexToAddress(req.GetFrom()),
		To:         common.HexToAddress(req.GetTo()),
		Value:      value,
		TransferID: req.GetTransferId(),
		ExpiresAt:  req.GetExpiresAt(),
		ExtraData:  req.GetExtraData(),
		CryptoData: apiDto.AssetboxCryptoData{
			Wallet:     req.GetFromCryptoData().GetWallet(),
			Passphrase: req.GetFromCryptoData().GetPassphrase(),
		},
	}
	apiResponse, err := c.bitbon.GetAPI().CreateWPCSafeTransfer(
		loggerContext.NewLoggerContext(context.Background(), logger), apiRequest)
	if err != nil {
		err = errors.Wrap(err, "error calling api method")
		return
	}

	response.TxHash = strings.ToLower(apiResponse.TxHash.Hex())
	response.BlockNumber = fmt.Sprintf("%d", apiResponse.BlockNumber)
	logger.Debug("handleCreateWPCSafeTransfer method ended with success",
		"TxHash", response.TxHash, "BlockNumber", response.BlockNumber)
	return
}

func (c *Client) handleApproveWPCSafeTransfer(body []byte) (protoMessage protoMessageWithID, err error) {
	req := &external.ApproveWPCSafeTransferRequest{}
	err = proto.Unmarshal(body, req)
	if err != nil {
		err = errors.Wrap(err, "failed to unmarshal request")
		return
	}
	response := &external.TransferResponse{
		Id: req.Id,
	}

	logger := log.New("requestId", req.GetId())
	defer func() {
		if err != nil {
			jsonReq, _ := json.Marshal(req)
			logger.Warn("handleApproveWPCSafeTransfer method ended with error", "request", string(jsonReq), "error", err)

			response.Error = c.handleError(err)
			err = errors.Wrapf(err, "handle request (%s) error", req.GetId())
		}

		protoMessage = response
	}()

	apiRequest := &apiDto.ApproveWPCSafeTransferRequest{
		Address:    common.HexToAddress(req.GetAddress()),
		TransferID: req.GetTransferID(),
		ExtraData:  req.GetExtraData(),
		CryptoData: apiDto.AssetboxCryptoData{
			Wallet:     req.GetCryptoData().GetWallet(),
			Passphrase: req.GetCryptoData().GetPassphrase(),
		},
	}
	apiResponse, err := c.bitbon.GetAPI().ApproveWPCSafeTransfer(
		loggerContext.NewLoggerContext(context.Background(), logger), apiRequest)
	if err != nil {
		err = errors.Wrap(err, "error calling api method")
		return
	}

	response.TxHash = strings.ToLower(apiResponse.TxHash.Hex())
	response.BlockNumber = fmt.Sprintf("%d", apiResponse.BlockNumber)
	logger.Debug("handleApproveWPCSafeTransfer method ended with success",
		"TxHash", response.TxHash, "BlockNumber", response.BlockNumber)
	return
}

func (c *Client) handleCancelWPCSafeTransfer(body []byte) (protoMessage protoMessageWithID, err error) {
	req := &external.CancelWPCSafeTransferRequest{}
	err = proto.Unmarshal(body, req)
	if err != nil {
		err = errors.Wrap(err, "failed to unmarshal request")
		return
	}
	response := &external.TransferResponse{
		Id: req.Id,
	}

	logger := log.New("requestId", req.GetId())
	defer func() {
		if err != nil {
			jsonReq, _ := json.Marshal(req)
			logger.Warn("handleCancelWPCSafeTransfer method ended with error", "request", string(jsonReq), "error", err)

			response.Error = c.handleError(err)
			err = errors.Wrapf(err, "handle request (%s) error", req.GetId())
		}

		protoMessage = response
	}()

	apiRequest := &apiDto.CancelWPCSafeTransferRequest{
		Address:    common.HexToAddress(req.GetAddress()),
		TransferID: req.GetTransferID(),
		ExtraData:  req.GetExtraData(),
		CryptoData: apiDto.AssetboxCryptoData{
			Wallet:     req.GetCryptoData().GetWallet(),
			Passphrase: req.GetCryptoData().GetPassphrase(),
		},
	}
	apiResponse, err := c.bitbon.GetAPI().CancelWPCSafeTransfer(
		loggerContext.NewLoggerContext(context.Background(), logger), apiRequest)
	if err != nil {
		err = errors.Wrap(err, "error calling api method")
		return
	}

	response.TxHash = strings.ToLower(apiResponse.TxHash.Hex())
	response.BlockNumber = fmt.Sprintf("%d", apiResponse.BlockNumber)
	logger.Debug("handleCancelWPCSafeTransfer method ended with success",
		"TxHash", response.TxHash, "BlockNumber", response.BlockNumber)
	return
}

func (c *Client) handleCreateSafeTransfer(body []byte) (protoMessage protoMessageWithID, err error) {
	req := &external.CreateSafeTransferRequest{}
	err = proto.Unmarshal(body, req)
	if err != nil {
		err = errors.Wrap(err, "failed to unmarshal request")
		return
	}
	response := &external.TransferResponse{
		Id: req.Id,
	}

	logger := log.New("requestId", req.GetId())
	defer func() {
		if err != nil {
			jsonReq, _ := json.Marshal(req)
			logger.Warn("handleCreateSafeTransfer method ended with error", "request", string(jsonReq), "error", err)

			response.Error = c.handleError(err)
			err = errors.Wrapf(err, "handle request (%s) error", req.GetId())
		}

		protoMessage = response
	}()

	value, ok := new(big.Int).SetString(req.GetValue(), 10)
	if !ok {
		err = errors.New("failed to create big int from string")
		return
	}

	apiRequest := &apiDto.CreateSafeTransferRequest{
		From:           common.HexToAddress(req.GetFrom()),
		To:             common.HexToAddress(req.GetTo()),
		Value:          value,
		TransferID:     req.GetTransferId(),
		ProtectionCode: req.GetProtectionCode(),
		Retries:        req.GetRetries(),
		ExpiresAt:      req.GetExpiresAt(),
		ExtraData:      req.GetExtraData(),
		CryptoData: apiDto.AssetboxCryptoData{
			Wallet:     req.GetFromCryptoData().GetWallet(),
			Passphrase: req.GetFromCryptoData().GetPassphrase(),
		},
	}

	apiResponse, err := c.bitbon.GetAPI().CreateSafeTransfer(
		loggerContext.NewLoggerContext(context.Background(), logger), apiRequest)
	if err != nil {
		err = errors.Wrap(err, "error calling api method")
		return
	}

	response.TxHash = strings.ToLower(apiResponse.TxHash.Hex())
	response.BlockNumber = fmt.Sprintf("%d", apiResponse.BlockNumber)
	logger.Debug("handleCreateSafeTransfer method ended with success",
		"TxHash", response.TxHash, "BlockNumber", response.BlockNumber)
	return
}

func (c *Client) handleApproveSafeTransfer(body []byte) (protoMessage protoMessageWithID, err error) {
	req := &external.ApproveSafeTransferRequest{}
	err = proto.Unmarshal(body, req)
	if err != nil {
		err = errors.Wrap(err, "failed to unmarshal request")
		return
	}
	response := &external.TransferResponse{
		Id: req.Id,
	}

	logger := log.New("requestId", req.GetId())
	defer func() {
		if err != nil {
			jsonReq, _ := json.Marshal(req)
			logger.Warn("handleApproveSafeTransfer method ended with error", "request", string(jsonReq), "error", err)

			response.Error = c.handleError(err)
			err = errors.Wrapf(err, "handle request (%s) error", req.GetId())
		}

		protoMessage = response
	}()

	apiRequest := &apiDto.ApproveSafeTransferRequest{
		Address:        common.HexToAddress(req.GetAddress()),
		TransferID:     req.GetTransferID(),
		ProtectionCode: req.GetProtectionCode(),
		ExtraData:      req.GetExtraData(),
		CryptoData: apiDto.AssetboxCryptoData{
			Wallet:     req.GetCryptoData().GetWallet(),
			Passphrase: req.GetCryptoData().GetPassphrase(),
		},
	}
	apiResponse, err := c.bitbon.GetAPI().ApproveSafeTransfer(
		loggerContext.NewLoggerContext(context.Background(), logger), apiRequest)
	if err != nil {
		err = errors.Wrap(err, "error calling api method")
		return
	}

	response.TxHash = strings.ToLower(apiResponse.TxHash.Hex())
	response.BlockNumber = fmt.Sprintf("%d", apiResponse.BlockNumber)
	logger.Debug("handleApproveSafeTransfer method ended with success",
		"TxHash", response.TxHash, "BlockNumber", response.BlockNumber)
	return
}

func (c *Client) handleCancelSafeTransfer(body []byte) (protoMessage protoMessageWithID, err error) {
	req := &external.CancelSafeTransferRequest{}
	err = proto.Unmarshal(body, req)
	if err != nil {
		err = errors.Wrap(err, "failed to unmarshal request")
		return
	}
	response := &external.TransferResponse{
		Id: req.Id,
	}

	logger := log.New("requestId", req.GetId())
	defer func() {
		if err != nil {
			jsonReq, _ := json.Marshal(req)
			logger.Warn("handleCancelSafeTransfer method ended with error", "request", string(jsonReq), "error", err)

			response.Error = c.handleError(err)
			err = errors.Wrapf(err, "handle request (%s) error", req.GetId())
		}

		protoMessage = response
	}()

	apiRequest := &apiDto.CancelSafeTransferRequest{
		Address:    common.HexToAddress(req.GetAddress()),
		TransferID: req.GetTransferID(),
		ExtraData:  req.GetExtraData(),
		CryptoData: apiDto.AssetboxCryptoData{
			Wallet:     req.GetCryptoData().GetWallet(),
			Passphrase: req.GetCryptoData().GetPassphrase(),
		},
	}
	apiResponse, err := c.bitbon.GetAPI().CancelSafeTransfer(
		loggerContext.NewLoggerContext(context.Background(), logger), apiRequest)
	if err != nil {
		err = errors.Wrap(err, "error calling api method")
		return
	}

	response.TxHash = strings.ToLower(apiResponse.TxHash.Hex())
	response.BlockNumber = fmt.Sprintf("%d", apiResponse.BlockNumber)
	logger.Debug("handleCancelSafeTransfer method ended with success",
		"TxHash", response.TxHash, "BlockNumber", response.BlockNumber)
	return
}

func (c *Client) handleDirectTransfer(body []byte) (protoMessage protoMessageWithID, err error) {
	req := &external.DirectTransferRequest{}
	err = proto.Unmarshal(body, req)
	if err != nil {
		err = errors.Wrap(err, "failed to unmarshal request")
		return
	}
	response := &external.TransferResponse{
		Id: req.Id,
	}

	logger := log.New("requestId", req.GetId())
	defer func() {
		if err != nil {
			jsonReq, _ := json.Marshal(req)
			logger.Warn("handleDirectTransfer method ended with error", "request", string(jsonReq), "error", err)

			response.Error = c.handleError(err)
			err = errors.Wrapf(err, "handle request (%s) error", req.GetId())
		}

		protoMessage = response
	}()

	value, ok := new(big.Int).SetString(req.GetValue(), 10)
	if !ok {
		err = errors.New("failed to create big int from string")
		return
	}

	apiRequest := &apiDto.DirectTransferRequest{
		From:      common.HexToAddress(req.GetFrom()),
		To:        common.HexToAddress(req.GetTo()),
		Value:     value,
		ExtraData: req.GetExtraData(),
	}
	apiResponse, err := c.bitbon.GetAPI().DirectTransfer(
		loggerContext.NewLoggerContext(context.Background(), logger), apiRequest)
	if err != nil {
		err = errors.Wrap(err, "error calling api method")
		return
	}

	response.TxHash = strings.ToLower(apiResponse.TxHash.Hex())
	response.BlockNumber = fmt.Sprintf("%d", apiResponse.BlockNumber)
	logger.Debug("handleDirectTransfer method ended with success",
		"TxHash", response.TxHash, "BlockNumber", response.BlockNumber)
	return
}

func (c *Client) handleFullBalanceQuickTransfer(body []byte) (protoMessage protoMessageWithID, err error) {
	req := &external.FullBalanceQuickTransferRequest{}
	err = proto.Unmarshal(body, req)
	if err != nil {
		err = errors.Wrap(err, "failed to unmarshal request")
		return
	}

	response := &external.TransferResponse{
		Id: req.Id,
	}

	logger := log.New("requestId", req.GetId())
	defer func() {
		if err != nil {
			jsonReq, _ := json.Marshal(req)
			logger.Warn("handleFullBalanceQuickTransfer method ended with error", "request", string(jsonReq), "error", err)

			response.Error = c.handleError(err)
			err = errors.Wrapf(err, "handle request (%s) error", req.GetId())
		}

		protoMessage = response
	}()

	apiRequest := &apiDto.FullBalanceQuickTransferRequest{
		From:      common.HexToAddress(req.GetFrom()),
		To:        common.HexToAddress(req.GetTo()),
		ExtraData: req.GetExtraData(),
		CryptoData: apiDto.AssetboxCryptoData{
			Wallet:     req.GetFromCryptoData().GetWallet(),
			Passphrase: req.GetFromCryptoData().GetPassphrase(),
		},
	}

	apiResponse, err := c.bitbon.GetAPI().FullBalanceQuickTransfer(loggerContext.NewLoggerContext(context.Background(), logger), apiRequest)
	if err != nil {
		err = errors.Wrap(err, "error calling api method")
		return
	}

	response.TxHash = strings.ToLower(apiResponse.TxHash.Hex())
	response.BlockNumber = fmt.Sprintf("%d", apiResponse.BlockNumber)
	logger.Warn("handleFullBalanceQuickTransfer method ended with success", "TxHash", response.TxHash, "BlockNumber", response.BlockNumber)
	return
}

func (c *Client) handleCreateFullBalanceWPCSafeTransfer(body []byte) (protoMessage protoMessageWithID, err error) {
	req := &external.CreateFullBalanceWPCSafeTransferRequest{}
	err = proto.Unmarshal(body, req)
	if err != nil {
		err = errors.Wrap(err, "failed to unmarshal request")
		return
	}

	response := &external.TransferResponse{
		Id: req.Id,
	}

	logger := log.New("requestId", req.GetId())
	defer func() {
		if err != nil {
			jsonReq, _ := json.Marshal(req)
			logger.Warn("handleCreateFullBalanceWPCSafeTransfer method ended with error", "request", string(jsonReq), "error", err)

			response.Error = c.handleError(err)
			err = errors.Wrapf(err, "handle request (%s) error", req.GetId())
		}

		protoMessage = response
	}()

	apiRequest := &apiDto.CreateFullBalanceWPCSafeTransferRequest{
		From:       common.HexToAddress(req.GetFrom()),
		To:         common.HexToAddress(req.GetTo()),
		TransferID: req.GetTransferId(),
		ExpiresAt:  req.GetExpiresAt(),
		ExtraData:  req.GetExtraData(),
		CryptoData: apiDto.AssetboxCryptoData{
			Wallet:     req.GetFromCryptoData().GetWallet(),
			Passphrase: req.GetFromCryptoData().GetPassphrase(),
		},
	}
	apiResponse, err := c.bitbon.GetAPI().CreateFullBalanceWPCSafeTransfer(loggerContext.NewLoggerContext(context.Background(), logger), apiRequest)
	if err != nil {
		err = errors.Wrap(err, "error calling api method")
		return
	}

	response.TxHash = strings.ToLower(apiResponse.TxHash.Hex())
	response.BlockNumber = fmt.Sprintf("%d", apiResponse.BlockNumber)
	logger.Warn("handleCreateFullBalanceWPCSafeTransfer method ended with success", "TxHash", response.TxHash, "BlockNumber", response.BlockNumber)
	return
}

func (c *Client) handleApproveFullBalanceWPCSafeTransfer(body []byte) (protoMessage protoMessageWithID, err error) {
	req := &external.ApproveFullBalanceWPCSafeTransferRequest{}
	err = proto.Unmarshal(body, req)
	if err != nil {
		err = errors.Wrap(err, "failed to unmarshal request")
		return
	}
	response := &external.TransferResponse{
		Id: req.Id,
	}

	logger := log.New("requestId", req.GetId())
	defer func() {
		if err != nil {
			jsonReq, _ := json.Marshal(req)
			logger.Warn("handleApproveFullBalanceWPCSafeTransfer method ended with error", "request", string(jsonReq), "error", err)

			response.Error = c.handleError(err)
			err = errors.Wrapf(err, "handle request (%s) error", req.GetId())
		}

		protoMessage = response
	}()

	apiRequest := &apiDto.ApproveFullBalanceWPCSafeTransferRequest{
		Address:    common.HexToAddress(req.GetAddress()),
		TransferID: req.GetTransferID(),
		ExtraData:  req.GetExtraData(),
		CryptoData: apiDto.AssetboxCryptoData{
			Wallet:     req.GetCryptoData().GetWallet(),
			Passphrase: req.GetCryptoData().GetPassphrase(),
		},
	}
	apiResponse, err := c.bitbon.GetAPI().ApproveFullBalanceWPCSafeTransfer(loggerContext.NewLoggerContext(context.Background(), logger), apiRequest)
	if err != nil {
		err = errors.Wrap(err, "error calling api method")
		return
	}

	response.TxHash = strings.ToLower(apiResponse.TxHash.Hex())
	response.BlockNumber = fmt.Sprintf("%d", apiResponse.BlockNumber)
	logger.Warn("handleApproveFullBalanceWPCSafeTransfer method ended with success", "TxHash", response.TxHash, "BlockNumber", response.BlockNumber)
	return
}

func (c *Client) handleCancelFullBalanceWPCSafeTransfer(body []byte) (protoMessage protoMessageWithID, err error) {
	req := &external.CancelFullBalanceWPCSafeTransferRequest{}
	err = proto.Unmarshal(body, req)
	if err != nil {
		err = errors.Wrap(err, "failed to unmarshal request")
		return
	}
	response := &external.TransferResponse{
		Id: req.Id,
	}

	logger := log.New("requestId", req.GetId())
	defer func() {
		if err != nil {
			jsonReq, _ := json.Marshal(req)
			logger.Warn("handleCancelFullBalanceWPCSafeTransfer method ended with error", "request", string(jsonReq), "error", err)

			response.Error = c.handleError(err)
			err = errors.Wrapf(err, "handle request (%s) error", req.GetId())
		}

		protoMessage = response
	}()

	apiRequest := &apiDto.CancelFullBalanceWPCSafeTransferRequest{
		Address:    common.HexToAddress(req.GetAddress()),
		TransferID: req.GetTransferID(),
		ExtraData:  req.GetExtraData(),
		CryptoData: apiDto.AssetboxCryptoData{
			Wallet:     req.GetCryptoData().GetWallet(),
			Passphrase: req.GetCryptoData().GetPassphrase(),
		},
	}
	apiResponse, err := c.bitbon.GetAPI().CancelFullBalanceWPCSafeTransfer(loggerContext.NewLoggerContext(context.Background(), logger), apiRequest)
	if err != nil {
		err = errors.Wrap(err, "error calling api method")
		return
	}

	response.TxHash = strings.ToLower(apiResponse.TxHash.Hex())
	response.BlockNumber = fmt.Sprintf("%d", apiResponse.BlockNumber)
	logger.Warn("handleCancelFullBalanceWPCSafeTransfer method ended with success", "TxHash", response.TxHash, "BlockNumber", response.BlockNumber)
	return
}

func (c *Client) handleCreateFullBalanceSafeTransfer(body []byte) (protoMessage protoMessageWithID, err error) {
	req := &external.CreateFullBalanceSafeTransferRequest{}
	err = proto.Unmarshal(body, req)
	if err != nil {
		err = errors.Wrap(err, "failed to unmarshal request")
		return
	}
	response := &external.TransferResponse{
		Id: req.Id,
	}

	logger := log.New("requestId", req.GetId())
	defer func() {
		if err != nil {
			jsonReq, _ := json.Marshal(req)
			logger.Warn("handleCreateFullBalanceSafeTransfer method ended with error", "request", string(jsonReq), "error", err)

			response.Error = c.handleError(err)
			err = errors.Wrapf(err, "handle request (%s) error", req.GetId())
		}

		protoMessage = response
	}()

	apiRequest := &apiDto.CreateFullBalanceSafeTransferRequest{
		From:           common.HexToAddress(req.GetFrom()),
		To:             common.HexToAddress(req.GetTo()),
		TransferID:     req.GetTransferId(),
		ProtectionCode: req.GetProtectionCode(),
		Retries:        req.GetRetries(),
		ExpiresAt:      req.GetExpiresAt(),
		ExtraData:      req.GetExtraData(),
		CryptoData: apiDto.AssetboxCryptoData{
			Wallet:     req.GetFromCryptoData().GetWallet(),
			Passphrase: req.GetFromCryptoData().GetPassphrase(),
		},
	}

	apiResponse, err := c.bitbon.GetAPI().CreateFullBalanceSafeTransfer(loggerContext.NewLoggerContext(context.Background(), logger), apiRequest)
	if err != nil {
		err = errors.Wrap(err, "error calling api method")
		return
	}

	response.TxHash = strings.ToLower(apiResponse.TxHash.Hex())
	response.BlockNumber = fmt.Sprintf("%d", apiResponse.BlockNumber)
	logger.Warn("handleCreateFullBalanceSafeTransfer method ended with success", "TxHash", response.TxHash, "BlockNumber", response.BlockNumber)
	return
}

func (c *Client) handleApproveFullBalanceSafeTransfer(body []byte) (protoMessage protoMessageWithID, err error) {
	req := &external.ApproveFullBalanceSafeTransferRequest{}
	err = proto.Unmarshal(body, req)
	if err != nil {
		err = errors.Wrap(err, "failed to unmarshal request")
		return
	}
	response := &external.TransferResponse{
		Id: req.Id,
	}

	logger := log.New("requestId", req.GetId())
	defer func() {
		if err != nil {
			jsonReq, _ := json.Marshal(req)
			logger.Warn("handleApproveFullBalanceSafeTransfer method ended with error", "request", string(jsonReq), "error", err)

			response.Error = c.handleError(err)
			err = errors.Wrapf(err, "handle request (%s) error", req.GetId())
		}

		protoMessage = response
	}()

	apiRequest := &apiDto.ApproveFullBalanceSafeTransferRequest{
		Address:        common.HexToAddress(req.GetAddress()),
		TransferID:     req.GetTransferID(),
		ProtectionCode: req.GetProtectionCode(),
		ExtraData:      req.GetExtraData(),
		CryptoData: apiDto.AssetboxCryptoData{
			Wallet:     req.GetCryptoData().GetWallet(),
			Passphrase: req.GetCryptoData().GetPassphrase(),
		},
	}
	apiResponse, err := c.bitbon.GetAPI().ApproveFullBalanceSafeTransfer(loggerContext.NewLoggerContext(context.Background(), logger), apiRequest)
	if err != nil {
		err = errors.Wrap(err, "error calling api method")
		return
	}

	response.TxHash = strings.ToLower(apiResponse.TxHash.Hex())
	response.BlockNumber = fmt.Sprintf("%d", apiResponse.BlockNumber)
	logger.Warn("handleApproveFullBalanceSafeTransfer method ended with success", "TxHash", response.TxHash, "BlockNumber", response.BlockNumber)
	return
}

func (c *Client) handleCancelFullBalanceSafeTransfer(body []byte) (protoMessage protoMessageWithID, err error) {
	req := &external.CancelFullBalanceSafeTransferRequest{}
	err = proto.Unmarshal(body, req)
	if err != nil {
		err = errors.Wrap(err, "failed to unmarshal request")
		return
	}
	response := &external.TransferResponse{
		Id: req.Id,
	}

	logger := log.New("requestId", req.GetId())
	defer func() {
		if err != nil {
			jsonReq, _ := json.Marshal(req)
			logger.Warn("handleCancelFullBalanceSafeTransfer method ended with error", "request", string(jsonReq), "error", err)

			response.Error = c.handleError(err)
			err = errors.Wrapf(err, "handle request (%s) error", req.GetId())
		}

		protoMessage = response
	}()

	apiRequest := &apiDto.CancelFullBalanceSafeTransferRequest{
		Address:    common.HexToAddress(req.GetAddress()),
		TransferID: req.GetTransferID(),
		ExtraData:  req.GetExtraData(),
		CryptoData: apiDto.AssetboxCryptoData{
			Wallet:     req.GetCryptoData().GetWallet(),
			Passphrase: req.GetCryptoData().GetPassphrase(),
		},
	}
	apiResponse, err := c.bitbon.GetAPI().CancelFullBalanceSafeTransfer(loggerContext.NewLoggerContext(context.Background(), logger), apiRequest)
	if err != nil {
		err = errors.Wrap(err, "error calling api method")
		return
	}

	response.TxHash = strings.ToLower(apiResponse.TxHash.Hex())
	response.BlockNumber = fmt.Sprintf("%d", apiResponse.BlockNumber)
	logger.Warn("handleCancelFullBalanceSafeTransfer method ended with success", "TxHash", response.TxHash, "BlockNumber", response.BlockNumber)
	return
}
