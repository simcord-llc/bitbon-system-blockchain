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

// nolint:nakedret
package client

import (
	"context"
	"encoding/json"

	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/dto/external"

	"github.com/simcord-llc/bitbon-system-blockchain/common"

	"github.com/golang/protobuf/proto" // nolint:staticcheck
	"github.com/pkg/errors"
	loggerContext "github.com/simcord-llc/bitbon-system-blockchain/bitbon/context"
	clientDto "github.com/simcord-llc/bitbon-system-blockchain/bitbon/dto"
	"github.com/simcord-llc/bitbon-system-blockchain/log"
)

func (c *Client) handleBlockTransactionCountByBlockHash(body []byte) (protoMessage protoMessageWithID, err error) {
	req := &external.BlockTransactionCountByHashRequest{}
	err = proto.Unmarshal(body, req)
	if err != nil {
		err = errors.Wrap(err, "failed to unmarshal request")
		return
	}
	response := &external.TransactionCountResponse{
		Id: req.Id,
	}

	logger := log.New("requestId", req.GetId())
	defer func() {
		if err != nil {
			jsonReq, _ := json.Marshal(req)
			logger.Warn("handleBlockTransactionCountByBlockHash method ended with error",
				"request", string(jsonReq), "error", err)

			response.Error = c.handleError(err)
			err = errors.Wrapf(err, "handle request (%s) error", req.GetId())
		}

		protoMessage = response
	}()

	resp, err := c.bitbon.GetAPI().BlockTransactionCountByBlockHash(
		loggerContext.NewLoggerContext(context.Background(), logger), common.HexToHash(req.GetBlockHash()))
	if err != nil {
		err = errors.Wrap(err, "error calling api method")
		return
	}
	response.Count = resp

	logger.Debug("handleBlockTransactionCountByBlockHash method ended with success")
	return
}

func (c *Client) handleBlockTransactionCountByNumber(body []byte) (protoMessage protoMessageWithID, err error) {
	req := &external.BlockTransactionCountByNumberRequest{}
	err = proto.Unmarshal(body, req)
	if err != nil {
		err = errors.Wrap(err, "failed to unmarshal request")
		return
	}
	response := &external.TransactionCountResponse{
		Id: req.Id,
	}

	logger := log.New("requestId", req.GetId())
	defer func() {
		if err != nil {
			jsonReq, _ := json.Marshal(req)
			logger.Warn("handleBlockTransactionCountByNumber method ended with error", "request", string(jsonReq), "error", err)

			response.Error = c.handleError(err)
			err = errors.Wrapf(err, "handle request (%s) error", req.GetId())
		}

		protoMessage = response
	}()

	resp, err := c.bitbon.GetAPI().BlockTransactionCountByBlockNumber(
		loggerContext.NewLoggerContext(context.Background(), logger), req.GetBlockNumber())
	if err != nil {
		err = errors.Wrap(err, "error calling api method")
		return
	}
	response.Count = resp
	logger.Debug("handleBlockTransactionCountByNumber method ended with success")

	return
}

func (c *Client) handleTransactionByBlockHashAndIndex(body []byte) (protoMessage protoMessageWithID, err error) {
	req := &external.TransactionByBlockHashAndIndexRequest{}
	err = proto.Unmarshal(body, req)
	if err != nil {
		err = errors.Wrap(err, "failed to unmarshal request")
		return
	}
	response := &external.TransactionResponse{
		Id: req.Id,
	}

	logger := log.New("requestId", req.GetId())
	defer func() {
		if err != nil {
			jsonReq, _ := json.Marshal(req)
			logger.Warn("handleTransactionByBlockHashAndIndex method ended with error", "request", string(jsonReq), "error", err)

			response.Error = c.handleError(err)
			err = errors.Wrapf(err, "handle request (%s) error", req.GetId())
		}

		protoMessage = response
	}()

	apiRequest := &clientDto.TransactionByBlockHashAndIndexRequest{
		BlockHash: common.HexToHash(req.GetBlockHash()),
		TxIndex:   req.GetTxIndex(),
	}

	resp, err := c.bitbon.GetAPI().TransactionByBlockHashAndIndex(
		loggerContext.NewLoggerContext(context.Background(), logger), apiRequest)
	if err != nil {
		err = errors.Wrap(err, "error calling api method")
		return
	}
	response.TransactionResult = clientDto.BitbonTransactionToDto([]*clientDto.Transaction{resp})

	logger.Debug("handleTransactionByBlockHashAndIndex method ended with success")
	return
}

func (c *Client) handleTransactionByBlockNumberAndIndex(body []byte) (protoMessage protoMessageWithID, err error) {
	req := &external.TransactionByBlockNumberAndIndexRequest{}
	err = proto.Unmarshal(body, req)
	if err != nil {
		err = errors.Wrap(err, "failed to unmarshal request")
		return
	}
	response := &external.TransactionResponse{
		Id: req.Id,
	}

	logger := log.New("requestId", req.GetId())
	defer func() {
		if err != nil {
			jsonReq, _ := json.Marshal(req)
			logger.Warn("handleTransactionByBlockNumberAndIndex method ended with error",
				"request", string(jsonReq), "error", err)

			response.Error = c.handleError(err)
			err = errors.Wrapf(err, "handle request (%s) error", req.GetId())
		}

		protoMessage = response
	}()

	apiRequest := &clientDto.TransactionByBlockNumberAndIndexRequest{
		BlockNumber: req.GetBlockNumber(),
		TxIndex:     req.GetTxIndex(),
	}

	resp, err := c.bitbon.GetAPI().TransactionByBlockNumberAndIndex(
		loggerContext.NewLoggerContext(context.Background(), logger), apiRequest)
	if err != nil {
		err = errors.Wrap(err, "error calling api method")
		return
	}
	response.TransactionResult = clientDto.BitbonTransactionToDto([]*clientDto.Transaction{resp})

	logger.Debug("handleTransactionByBlockNumberAndIndex method ended with success")
	return
}

func (c *Client) handleTransaction(body []byte) (protoMessage protoMessageWithID, err error) {
	req := &external.TransactionRequest{}
	err = proto.Unmarshal(body, req)
	if err != nil {
		err = errors.Wrap(err, "failed to unmarshal request")
		return
	}
	response := &external.TransactionResponse{
		Id: req.Id,
	}

	logger := log.New("requestId", req.GetId())
	defer func() {
		if err != nil {
			jsonReq, _ := json.Marshal(req)
			logger.Warn("handleTransaction method ended with error", "request", string(jsonReq), "error", err)

			response.Error = c.handleError(err)
			err = errors.Wrapf(err, "handle request (%s) error", req.GetId())
		}

		protoMessage = response
	}()

	resp, err := c.bitbon.GetAPI().TransactionByHash(loggerContext.NewLoggerContext(context.Background(), logger),
		common.HexToHash(req.GetTxHash()))
	if err != nil {
		err = errors.Wrap(err, "error calling api method")
		return
	}
	response.TransactionResult = clientDto.BitbonTransactionToDto([]*clientDto.Transaction{resp})

	logger.Debug("handleTransaction method ended with success")
	return
}

func (c *Client) handleTransactionsByPeriod(body []byte) (protoMessage protoMessageWithID, err error) {
	req := &external.TransactionsByTimePeriodRequest{}
	err = proto.Unmarshal(body, req)
	if err != nil {
		err = errors.Wrap(err, "failed to unmarshal request")
		return
	}
	response := &external.TransactionResponse{
		Id: req.Id,
	}

	logger := log.New("requestId", req.GetId())
	defer func() {
		if err != nil {
			jsonReq, _ := json.Marshal(req)
			logger.Warn("handleTransactionsByPeriod method ended with error", "request", string(jsonReq), "error", err)

			response.Error = c.handleError(err)
			err = errors.Wrapf(err, "handle request (%s) error", req.GetId())
		}

		protoMessage = response
	}()

	apiRequest := &clientDto.TransactionTimePeriodRequest{
		To:   req.GetTimestampTo(),
		From: req.GetTimestampFrom(),
	}

	resp, err := c.bitbon.GetAPI().TransactionsByTimePeriod(
		loggerContext.NewLoggerContext(context.Background(), logger), apiRequest)
	if err != nil {
		err = errors.Wrap(err, "error calling api method")
		return
	}
	response.TransactionResult = clientDto.BitbonTransactionToDto(resp)

	logger.Debug("handleTransactionsByPeriod method ended with success")
	return
}

func (c *Client) handleCheckTransactionsByPeriod(body []byte) (protoMessage protoMessageWithID, err error) {
	req := &external.TransactionsByTimePeriodRequest{}
	err = proto.Unmarshal(body, req)
	if err != nil {
		err = errors.Wrap(err, "failed to unmarshal request")
		return
	}
	response := &external.CheckTransactionsByTimePeriodResponse{
		Id: req.Id,
	}

	logger := log.New("requestId", req.GetId())
	defer func() {
		if err != nil {
			jsonReq, _ := json.Marshal(req)
			logger.Warn("handleCheckTransactionsByPeriod method ended with error", "request", string(jsonReq), "error", err)

			response.Error = c.handleError(err)
			err = errors.Wrapf(err, "handle request (%s) error", req.GetId())
		}

		protoMessage = response
	}()

	apiRequest := &clientDto.TransactionTimePeriodRequest{
		To:   req.GetTimestampTo(),
		From: req.GetTimestampFrom(),
	}

	resp, err := c.bitbon.GetAPI().CheckTransactionsByTimePeriod(
		loggerContext.NewLoggerContext(context.Background(), logger), apiRequest)
	if err != nil {
		err = errors.Wrap(err, "error calling api method")
		return
	}
	response.Hash = resp.Hex()

	logger.Debug("handleCheckTransactionsByPeriod method ended with success")
	return
}
