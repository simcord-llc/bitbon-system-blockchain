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
	"encoding/json"
	"math/big"
	"strings"

	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
	loggerContext "github.com/simcord-llc/bitbon-system-blockchain/bitbon/context"
	apiDto "github.com/simcord-llc/bitbon-system-blockchain/bitbon/dto"
	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/dto/external"
	"github.com/simcord-llc/bitbon-system-blockchain/common"
	"github.com/simcord-llc/bitbon-system-blockchain/log"
)

func (c *Client) handleQuickTransferAsync(body []byte) (protoMessage protoMessageWithID, err error) {
	req := &external.QuickTransferRequest{}
	err = proto.Unmarshal(body, req)
	if err != nil {
		err = errors.Wrap(err, "failed to unmarshal request")
		return
	}

	response := &external.TransferAsyncResponse{
		Id: req.Id,
	}

	logger := log.New("requestId", req.GetId())
	defer func() {
		if err != nil {
			jsonReq, _ := json.Marshal(req)
			logger.Warn("handleQuickTransferAsync method ended with error", "request", string(jsonReq), "error", err)

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

	hash, err := c.bitbon.GetAPI().QuickTransferAsync(
		loggerContext.NewLoggerContext(context.Background(), logger), apiRequest)
	if err != nil {
		err = errors.Wrap(err, "error calling api method")
		return
	}

	response.TxHash = strings.ToLower(hash.Hex())
	logger.Debug("handleQuickTransferAsync method ended with success", "TxHash", response.TxHash)
	return
}

func (c *Client) handleFrameTransferAsync(body []byte) (protoMessage protoMessageWithID, err error) {
	req := &external.QuickTransferRequest{}
	err = proto.Unmarshal(body, req)
	if err != nil {
		err = errors.Wrap(err, "failed to unmarshal request")
		return
	}

	response := &external.TransferAsyncResponse{
		Id: req.Id,
	}

	logger := log.New("requestId", req.GetId())
	defer func() {
		if err != nil {
			jsonReq, _ := json.Marshal(req)
			logger.Warn("handleFrameTransferAsync method ended with error", "request", string(jsonReq), "error", err)

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

	hash, err := c.bitbon.GetAPI().FrameTransferAsync(
		loggerContext.NewLoggerContext(context.Background(), logger), apiRequest)
	if err != nil {
		err = errors.Wrap(err, "error calling api method")
		return
	}

	response.TxHash = strings.ToLower(hash.Hex())
	logger.Debug("handleFrameTransferAsync method ended with success", "TxHash", response.TxHash)
	return
}

func (c *Client) handleServiceFeeTransferAsync(body []byte) (protoMessage protoMessageWithID, err error) {
	req := &external.ServiceFeeTransferRequest{}
	err = proto.Unmarshal(body, req)
	if err != nil {
		err = errors.Wrap(err, "failed to unmarshal request")
		return
	}

	response := &external.TransferAsyncResponse{
		Id: req.Id,
	}

	logger := log.New("requestId", req.GetId())
	defer func() {
		if err != nil {
			jsonReq, _ := json.Marshal(req)
			logger.Warn("handleServiceFeeTransferAsync method ended with error", "request", string(jsonReq), "error", err)

			response.Error = c.handleError(err)
			err = errors.Wrapf(err, "handle request (%s) error", req.GetId())
		}

		protoMessage = response
	}()

	apiRequest := &apiDto.ServiceFeeTransferRequest{
		From:          common.HexToAddress(req.GetFrom()),
		OperationType: big.NewInt(int64(req.OperationType)),
	}

	hash, err := c.bitbon.GetAPI().FeeTransferAsync(loggerContext.NewLoggerContext(context.Background(), logger), apiRequest)
	if err != nil {
		err = errors.Wrap(err, "error calling api method")
		return
	}

	response.TxHash = strings.ToLower(hash.Hex())
	logger.Warn("handleServiceFeeTransferAsync method ended with success", "TxHash", response.TxHash)
	return
}

func (c *Client) handleCreateWPCSafeTransferAsync(body []byte) (protoMessage protoMessageWithID, err error) {
	req := &external.CreateWPCSafeTransferRequest{}
	err = proto.Unmarshal(body, req)
	if err != nil {
		err = errors.Wrap(err, "failed to unmarshal request")
		return
	}

	response := &external.TransferAsyncResponse{
		Id: req.Id,
	}

	logger := log.New("requestId", req.GetId())
	defer func() {
		if err != nil {
			jsonReq, _ := json.Marshal(req)
			logger.Warn("handleCreateWPCSafeTransferAsync method ended with error", "request", string(jsonReq), "error", err)

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

	hash, err := c.bitbon.GetAPI().CreateWPCSafeTransferAsync(
		loggerContext.NewLoggerContext(context.Background(), logger), apiRequest)
	if err != nil {
		err = errors.Wrap(err, "error calling api method")
		return
	}

	response.TxHash = strings.ToLower(hash.Hex())
	logger.Debug("handleCreateWPCSafeTransferAsync method ended with success", "TxHash", response.TxHash)
	return
}

func (c *Client) handleApproveWPCSafeTransferAsync(body []byte) (protoMessage protoMessageWithID, err error) {
	req := &external.ApproveWPCSafeTransferRequest{}
	err = proto.Unmarshal(body, req)
	if err != nil {
		err = errors.Wrap(err, "failed to unmarshal request")
		return
	}
	response := &external.TransferAsyncResponse{
		Id: req.Id,
	}

	logger := log.New("requestId", req.GetId())
	defer func() {
		if err != nil {
			jsonReq, _ := json.Marshal(req)
			logger.Warn("handleApproveWPCSafeTransferAsync method ended with error", "request", string(jsonReq), "error", err)

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
	hash, err := c.bitbon.GetAPI().ApproveWPCSafeTransferAsync(
		loggerContext.NewLoggerContext(context.Background(), logger), apiRequest)
	if err != nil {
		err = errors.Wrap(err, "error calling api method")
		return
	}

	response.TxHash = strings.ToLower(hash.Hex())
	logger.Debug("handleApproveWPCSafeTransferAsync method ended with success", "TxHash", response.TxHash)
	return
}

func (c *Client) handleCancelWPCSafeTransferAsync(body []byte) (protoMessage protoMessageWithID, err error) {
	req := &external.CancelWPCSafeTransferRequest{}
	err = proto.Unmarshal(body, req)
	if err != nil {
		err = errors.Wrap(err, "failed to unmarshal request")
		return
	}
	response := &external.TransferAsyncResponse{
		Id: req.Id,
	}

	logger := log.New("requestId", req.GetId())
	defer func() {
		if err != nil {
			jsonReq, _ := json.Marshal(req)
			logger.Warn("handleCancelWPCSafeTransferAsync method ended with error", "request", string(jsonReq), "error", err)

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
	hash, err := c.bitbon.GetAPI().CancelWPCSafeTransferAsync(
		loggerContext.NewLoggerContext(context.Background(), logger), apiRequest)
	if err != nil {
		err = errors.Wrap(err, "error calling api method")
		return
	}

	response.TxHash = strings.ToLower(hash.Hex())
	logger.Debug("handleCancelWPCSafeTransferAsync method ended with success", "TxHash", response.TxHash)
	return
}

func (c *Client) handleCreateSafeTransferAsync(body []byte) (protoMessage protoMessageWithID, err error) {
	req := &external.CreateSafeTransferRequest{}
	err = proto.Unmarshal(body, req)
	if err != nil {
		err = errors.Wrap(err, "failed to unmarshal request")
		return
	}
	response := &external.TransferAsyncResponse{
		Id: req.Id,
	}

	logger := log.New("requestId", req.GetId())
	defer func() {
		if err != nil {
			jsonReq, _ := json.Marshal(req)
			logger.Warn("handleCreateSafeTransferAsync method ended with error", "request", string(jsonReq), "error", err)

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

	hash, err := c.bitbon.GetAPI().CreateSafeTransferAsync(
		loggerContext.NewLoggerContext(context.Background(), logger), apiRequest)
	if err != nil {
		err = errors.Wrap(err, "error calling api method")
		return
	}

	response.TxHash = strings.ToLower(hash.Hex())
	logger.Debug("handleCreateSafeTransferAsync method ended with success", "TxHash", response.TxHash)
	return
}

func (c *Client) handleApproveSafeTransferAsync(body []byte) (protoMessage protoMessageWithID, err error) {
	req := &external.ApproveSafeTransferRequest{}
	err = proto.Unmarshal(body, req)
	if err != nil {
		err = errors.Wrap(err, "failed to unmarshal request")
		return
	}
	response := &external.TransferAsyncResponse{
		Id: req.Id,
	}

	logger := log.New("requestId", req.GetId())
	defer func() {
		if err != nil {
			jsonReq, _ := json.Marshal(req)
			logger.Warn("handleApproveSafeTransferAsync method ended with error", "request", string(jsonReq), "error", err)

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
	hash, err := c.bitbon.GetAPI().ApproveSafeTransferAsync(
		loggerContext.NewLoggerContext(context.Background(), logger), apiRequest)
	if err != nil {
		err = errors.Wrap(err, "error calling api method")
		return
	}

	response.TxHash = strings.ToLower(hash.Hex())
	logger.Debug("handleApproveSafeTransferAsync method ended with success", "TxHash", response.TxHash)
	return
}

func (c *Client) handleCancelSafeTransferAsync(body []byte) (protoMessage protoMessageWithID, err error) {
	req := &external.CancelSafeTransferRequest{}
	err = proto.Unmarshal(body, req)
	if err != nil {
		err = errors.Wrap(err, "failed to unmarshal request")
		return
	}
	response := &external.TransferAsyncResponse{
		Id: req.Id,
	}

	logger := log.New("requestId", req.GetId())
	defer func() {
		if err != nil {
			jsonReq, _ := json.Marshal(req)
			logger.Warn("handleCancelSafeTransferAsync method ended with error", "request", string(jsonReq), "error", err)

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
	hash, err := c.bitbon.GetAPI().CancelSafeTransferAsync(loggerContext.NewLoggerContext(context.Background(), logger),
		apiRequest)
	if err != nil {
		err = errors.Wrap(err, "error calling api method")
		return
	}

	response.TxHash = strings.ToLower(hash.Hex())
	logger.Warn("handleCancelSafeTransferAsync method ended with success", "TxHash", response.TxHash)
	return
}

func (c *Client) handleDirectTransferAsync(body []byte) (protoMessage protoMessageWithID, err error) {
	req := &external.DirectTransferRequest{}
	err = proto.Unmarshal(body, req)
	if err != nil {
		err = errors.Wrap(err, "failed to unmarshal request")
		return
	}
	response := &external.TransferAsyncResponse{
		Id: req.Id,
	}

	logger := log.New("requestId", req.GetId())
	defer func() {
		if err != nil {
			jsonReq, _ := json.Marshal(req)
			logger.Warn("handleDirectTransferAsync method ended with error", "request", string(jsonReq), "error", err)

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
	hash, err := c.bitbon.GetAPI().DirectTransferAsync(loggerContext.NewLoggerContext(context.Background(), logger),
		apiRequest)
	if err != nil {
		err = errors.Wrap(err, "error calling api method")
		return
	}

	response.TxHash = strings.ToLower(hash.Hex())
	logger.Debug("handleDirectTransferAsync method ended with success", "TxHash", response.TxHash)
	return
}

func (c *Client) handleFullBalanceQuickTransferAsync(body []byte) (protoMessage protoMessageWithID, err error) {
	req := &external.FullBalanceQuickTransferRequest{}
	err = proto.Unmarshal(body, req)
	if err != nil {
		err = errors.Wrap(err, "failed to unmarshal request")
		return
	}

	response := &external.TransferAsyncResponse{
		Id: req.Id,
	}

	logger := log.New("requestId", req.GetId())
	defer func() {
		if err != nil {
			jsonReq, _ := json.Marshal(req)
			logger.Warn("handleFullBalanceQuickTransferAsync method ended with error", "request", string(jsonReq), "error", err)

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

	hash, err := c.bitbon.GetAPI().FullBalanceQuickTransferAsync(loggerContext.NewLoggerContext(context.Background(), logger), apiRequest)
	if err != nil {
		err = errors.Wrap(err, "error calling api method")
		return
	}

	response.TxHash = strings.ToLower(hash.Hex())
	logger.Warn("handleFullBalanceQuickTransferAsync method ended with success", "TxHash", response.TxHash)
	return
}

func (c *Client) handleCreateFullBalanceWPCSafeTransferAsync(body []byte) (protoMessage protoMessageWithID, err error) {
	req := &external.CreateFullBalanceWPCSafeTransferRequest{}
	err = proto.Unmarshal(body, req)
	if err != nil {
		err = errors.Wrap(err, "failed to unmarshal request")
		return
	}

	response := &external.TransferAsyncResponse{
		Id: req.Id,
	}

	logger := log.New("requestId", req.GetId())
	defer func() {
		if err != nil {
			jsonReq, _ := json.Marshal(req)
			logger.Warn("handleCreateFullBalanceWPCSafeTransferAsync method ended with error", "request", string(jsonReq), "error", err)

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

	hash, err := c.bitbon.GetAPI().CreateFullBalanceWPCSafeTransferAsync(loggerContext.NewLoggerContext(context.Background(), logger), apiRequest)
	if err != nil {
		err = errors.Wrap(err, "error calling api method")
		return
	}

	response.TxHash = strings.ToLower(hash.Hex())
	logger.Warn("handleCreateWPCSafeTransferAsync method ended with success", "TxHash", response.TxHash)
	return
}

func (c *Client) handleApproveFullBalanceWPCSafeTransferAsync(body []byte) (protoMessage protoMessageWithID, err error) {
	req := &external.ApproveFullBalanceWPCSafeTransferRequest{}
	err = proto.Unmarshal(body, req)
	if err != nil {
		err = errors.Wrap(err, "failed to unmarshal request")
		return
	}
	response := &external.TransferAsyncResponse{
		Id: req.Id,
	}

	logger := log.New("requestId", req.GetId())
	defer func() {
		if err != nil {
			jsonReq, _ := json.Marshal(req)
			logger.Warn("handleApproveFullBalanceWPCSafeTransferAsync method ended with error", "request", string(jsonReq), "error", err)

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
	hash, err := c.bitbon.GetAPI().ApproveFullBalanceWPCSafeTransferAsync(loggerContext.NewLoggerContext(context.Background(), logger), apiRequest)
	if err != nil {
		err = errors.Wrap(err, "error calling api method")
		return
	}

	response.TxHash = strings.ToLower(hash.Hex())
	logger.Warn("handleApproveFullBalanceWPCSafeTransferAsync method ended with success", "TxHash", response.TxHash)
	return
}

func (c *Client) handleCancelFullBalanceWPCSafeTransferAsync(body []byte) (protoMessage protoMessageWithID, err error) {
	req := &external.CancelFullBalanceWPCSafeTransferRequest{}
	err = proto.Unmarshal(body, req)
	if err != nil {
		err = errors.Wrap(err, "failed to unmarshal request")
		return
	}
	response := &external.TransferAsyncResponse{
		Id: req.Id,
	}

	logger := log.New("requestId", req.GetId())
	defer func() {
		if err != nil {
			jsonReq, _ := json.Marshal(req)
			logger.Warn("handleCancelFullBalanceWPCSafeTransferAsync method ended with error", "request", string(jsonReq), "error", err)

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
	hash, err := c.bitbon.GetAPI().CancelFullBalanceWPCSafeTransferAsync(loggerContext.NewLoggerContext(context.Background(), logger), apiRequest)
	if err != nil {
		err = errors.Wrap(err, "error calling api method")
		return
	}

	response.TxHash = strings.ToLower(hash.Hex())
	logger.Warn("handleCancelFullBalanceWPCSafeTransferAsync method ended with success", "TxHash", response.TxHash)
	return
}

func (c *Client) handleCreateFullBalanceSafeTransferAsync(body []byte) (protoMessage protoMessageWithID, err error) {
	req := &external.CreateFullBalanceSafeTransferRequest{}
	err = proto.Unmarshal(body, req)
	if err != nil {
		err = errors.Wrap(err, "failed to unmarshal request")
		return
	}
	response := &external.TransferAsyncResponse{
		Id: req.Id,
	}

	logger := log.New("requestId", req.GetId())
	defer func() {
		if err != nil {
			jsonReq, _ := json.Marshal(req)
			logger.Warn("handleCreateFullBalanceSafeTransferAsync method ended with error", "request", string(jsonReq), "error", err)

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

	hash, err := c.bitbon.GetAPI().CreateFullBalanceSafeTransferAsync(loggerContext.NewLoggerContext(context.Background(), logger), apiRequest)
	if err != nil {
		err = errors.Wrap(err, "error calling api method")
		return
	}

	response.TxHash = strings.ToLower(hash.Hex())
	logger.Warn("handleCreateFullBalanceSafeTransferAsync method ended with success", "TxHash", response.TxHash)
	return
}

func (c *Client) handleApproveFullBalanceSafeTransferAsync(body []byte) (protoMessage protoMessageWithID, err error) {
	req := &external.ApproveFullBalanceSafeTransferRequest{}
	err = proto.Unmarshal(body, req)
	if err != nil {
		err = errors.Wrap(err, "failed to unmarshal request")
		return
	}
	response := &external.TransferAsyncResponse{
		Id: req.Id,
	}

	logger := log.New("requestId", req.GetId())
	defer func() {
		if err != nil {
			jsonReq, _ := json.Marshal(req)
			logger.Warn("handleApproveFullBalanceSafeTransferAsync method ended with error", "request", string(jsonReq), "error", err)

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
	hash, err := c.bitbon.GetAPI().ApproveFullBalanceSafeTransferAsync(loggerContext.NewLoggerContext(context.Background(), logger), apiRequest)
	if err != nil {
		err = errors.Wrap(err, "error calling api method")
		return
	}

	response.TxHash = strings.ToLower(hash.Hex())
	logger.Warn("handleApproveFullBalanceSafeTransferAsync method ended with success", "TxHash", response.TxHash)
	return
}

func (c *Client) handleCancelFullBalanceSafeTransferAsync(body []byte) (protoMessage protoMessageWithID, err error) {
	req := &external.CancelFullBalanceSafeTransferRequest{}
	err = proto.Unmarshal(body, req)
	if err != nil {
		err = errors.Wrap(err, "failed to unmarshal request")
		return
	}
	response := &external.TransferAsyncResponse{
		Id: req.Id,
	}

	logger := log.New("requestId", req.GetId())
	defer func() {
		if err != nil {
			jsonReq, _ := json.Marshal(req)
			logger.Warn("handleCancelFullBalanceSafeTransferAsync method ended with error", "request", string(jsonReq), "error", err)

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
	hash, err := c.bitbon.GetAPI().CancelFullBalanceSafeTransferAsync(loggerContext.NewLoggerContext(context.Background(), logger), apiRequest)
	if err != nil {
		err = errors.Wrap(err, "error calling api method")
		return
	}

	response.TxHash = strings.ToLower(hash.Hex())
	logger.Warn("handleCancelFullBalanceSafeTransferAsync method ended with success", "TxHash", response.TxHash)
	return
}

func (c *Client) handleExpireTransfersAsync(body []byte) (protoMessage protoMessageWithID, err error) {
	req := &external.ExpireTransfersRequest{}
	err = proto.Unmarshal(body, req)
	if err != nil {
		err = errors.Wrap(err, "failed to unmarshal request")
		return
	}
	response := &external.TransferAsyncResponse{
		Id: req.Id,
	}

	logger := log.New("requestId", req.GetId())
	defer func() {
		if err != nil {
			jsonReq, _ := json.Marshal(req)
			logger.Warn("handleExpireTransfers method ended with error", "request", string(jsonReq), "error", err)

			response.Error = c.handleError(err)
			err = errors.Wrapf(err, "handle request (%s) error", req.GetId())
		}

		protoMessage = response
	}()

	hash, err := c.bitbon.GetAPI().
		ExpireTransfersAsync(loggerContext.NewLoggerContext(context.Background(), logger), req.TransferIds)
	if err != nil {
		err = errors.Wrap(err, "error calling api method")
		return
	}

	response.TxHash = strings.ToLower(hash.Hex())
	logger.Debug("handleExpireTransfersAsync method ended with success", "TxHash", response.TxHash)
	return
}
