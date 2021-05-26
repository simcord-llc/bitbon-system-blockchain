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

	"github.com/golang/protobuf/proto" // nolint:staticcheck
	"github.com/pkg/errors"
	loggerContext "github.com/simcord-llc/bitbon-system-blockchain/bitbon/context"
	clientDto "github.com/simcord-llc/bitbon-system-blockchain/bitbon/dto"
	"github.com/simcord-llc/bitbon-system-blockchain/common"
	"github.com/simcord-llc/bitbon-system-blockchain/log"
)

func (c *Client) handleBlockByHash(body []byte) (protoMessage protoMessageWithID, err error) {
	req := &external.BlockByHashRequest{}
	err = proto.Unmarshal(body, req)
	if err != nil {
		err = errors.Wrap(err, "failed to unmarshal request")
		return
	}
	response := &external.BlockInfoResponse{
		Id: req.Id,
	}

	logger := log.New("requestId", req.GetId())
	defer func() {
		if err != nil {
			jsonReq, _ := json.Marshal(req)
			logger.Warn("handleBlockByHash method ended with error", "request", string(jsonReq), "error", err)

			response.Error = c.handleError(err)
			err = errors.Wrapf(err, "handle request (%s) error", req.GetId())
		}

		protoMessage = response
	}()

	apiResponse, err := c.bitbon.GetAPI().BlockByHash(
		loggerContext.NewLoggerContext(context.Background(), logger), common.HexToHash(req.GetBlockHash()))
	if err != nil {
		err = errors.Wrap(err, "error calling api method")
		return
	}

	response.Block = clientDto.BitbonBlockToDto([]*clientDto.Block{apiResponse}, req.GetReturnFullTransaction())
	logger.Debug("handleBlockByHash method ended with success")
	return
}

func (c *Client) handleBlockByNumber(body []byte) (protoMessage protoMessageWithID, err error) {
	req := &external.BlockByNumberRequest{}
	err = proto.Unmarshal(body, req)
	if err != nil {
		err = errors.Wrap(err, "failed to unmarshal request")
		return
	}

	response := &external.BlockInfoResponse{
		Id: req.Id,
	}

	logger := log.New("requestId", req.GetId())
	defer func() {
		if err != nil {
			jsonReq, _ := json.Marshal(req)
			logger.Warn("handleBlockByHash method ended with error", "request", string(jsonReq), "error", err)

			response.Error = c.handleError(err)
			err = errors.Wrapf(err, "handle request (%s) error", req.GetId())
		}

		protoMessage = response
	}()

	apiResponse, err := c.bitbon.GetAPI().BlockByNumber(
		loggerContext.NewLoggerContext(context.Background(), logger), req.GetBlockNumber())
	if err != nil {
		err = errors.Wrap(err, "error calling api method")
		return
	}

	response.Block = clientDto.BitbonBlockToDto([]*clientDto.Block{apiResponse}, req.GetReturnFullTransaction())
	logger.Debug("handleBlockByNumber method ended with success")
	return
}

func (c *Client) handleRangeBlocksByNumber(body []byte) (protoMessage protoMessageWithID, err error) {
	req := &external.RangeBlocksByNumberRequest{}
	err = proto.Unmarshal(body, req)
	if err != nil {
		err = errors.Wrap(err, "failed to unmarshal request")
		return
	}
	response := &external.BlockInfoResponse{
		Id: req.Id,
	}

	logger := log.New("requestId", req.GetId())
	defer func() {
		if err != nil {
			jsonReq, _ := json.Marshal(req)
			logger.Warn("handleBlockByNumber method ended with error", "request", string(jsonReq), "error", err)

			response.Error = c.handleError(err)
			err = errors.Wrapf(err, "handle request (%s) error", req.GetId())
		}

		protoMessage = response
	}()

	apiRequest := &clientDto.RangeBlocksByNumberRequest{
		BlockNumberFrom:       req.GetBlockNumberFrom(),
		BlockNumberTo:         req.GetBlockNumberTo(),
		ReturnFullTransaction: req.GetReturnFullTransaction(),
	}

	apiResponse, err := c.bitbon.GetAPI().RangeBlockByNumber(
		loggerContext.NewLoggerContext(context.Background(), logger), apiRequest)
	if err != nil {
		err = errors.Wrap(err, "error calling api method")
		return
	}

	response.Block = clientDto.BitbonBlockToDto(apiResponse, apiRequest.ReturnFullTransaction)
	logger.Debug("handleRangeBlocksByNumber method ended with success")
	return
}

func (c *Client) handleBlocksByTimePeriod(body []byte) (protoMessage protoMessageWithID, err error) {
	req := &external.BlocksByTimePeriodRequest{}
	err = proto.Unmarshal(body, req)
	if err != nil {
		err = errors.Wrap(err, "failed to unmarshal request")
		return
	}
	response := &external.BlockInfoResponse{
		Id: req.Id,
	}

	logger := log.New("requestId", req.GetId())
	defer func() {
		if err != nil {
			jsonReq, _ := json.Marshal(req)
			logger.Warn("handleBlockByNumber method ended with error", "request", string(jsonReq), "error", err)

			response.Error = c.handleError(err)
			err = errors.Wrapf(err, "handle request (%s) error", req.GetId())
		}

		protoMessage = response
	}()

	apiRequest := &clientDto.BlocksTimePeriodRequest{
		From: req.GetTimestampFrom(),
		To:   req.GetTimestampTo(),
	}

	apiResponse, err := c.bitbon.GetAPI().BlocksByTimePeriod(
		loggerContext.NewLoggerContext(context.Background(), logger), apiRequest)
	if err != nil {
		err = errors.Wrap(err, "error calling api method")
		return
	}

	response.Block = clientDto.BitbonBlockToDto(apiResponse, true)
	logger.Debug("handleBlocksByTimePeriod method ended with success")
	return
}
