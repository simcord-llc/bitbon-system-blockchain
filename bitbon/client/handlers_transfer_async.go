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
		AccountID: req.GetAccountId(),
		To:        common.HexToAddress(req.GetTo()),
		Value:     value,
		ExtraData: req.GetExtraData(),
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
		AccountID:  req.GetAccountId(),
		To:         common.HexToAddress(req.GetTo()),
		Value:      value,
		TransferID: req.GetTransferId(),
		ExpiresAt:  req.GetExpiresAt(),
		ExtraData:  req.GetExtraData(),
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
		AccountID:      req.GetAccountId(),
		To:             common.HexToAddress(req.GetTo()),
		Value:          value,
		TransferID:     req.GetTransferId(),
		ProtectionCode: req.GetProtectionCode(),
		Retries:        req.GetRetries(),
		ExpiresAt:      req.GetExpiresAt(),
		ExtraData:      req.GetExtraData(),
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
