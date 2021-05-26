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
	"time"

	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/dto"
	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/dto/external"

	"github.com/golang/protobuf/proto" // nolint:staticcheck
	"github.com/pkg/errors"
	"github.com/simcord-llc/bitbon-system-blockchain/common"
	"github.com/simcord-llc/bitbon-system-blockchain/log"

	loggerContext "github.com/simcord-llc/bitbon-system-blockchain/bitbon/context"
)

func (c *Client) handleGetMinerNodes(body []byte) (protoMessage protoMessageWithID, err error) {
	req := &external.MinerNodesRequest{}
	err = proto.Unmarshal(body, req)
	if err != nil {
		err = errors.Wrap(err, "failed to unmarshal request")
		return
	}
	response := &external.MinerNodesResponse{
		Id: req.Id,
	}

	logger := log.New("requestId", req.GetId())
	defer func() {
		if err != nil {
			jsonReq, _ := json.Marshal(req)
			logger.Warn("handleGetMinerNodes method ended with error", "request", string(jsonReq), "error", err)

			response.Error = c.handleError(err)
			err = errors.Wrapf(err, "handle request (%s) error", req.GetId())
		}

		protoMessage = response
	}()

	if c.quorum == nil {
		return nil, errors.New("node is not in quorum mode")
	}

	resp := c.quorum.MinerListAtRecentEpoch()

	nodes := make([]string, len(resp))
	for i, addr := range resp {
		nodes[i] = addr.Hex()
	}

	response.Nodes = nodes
	response.ApplicationTimestamp = time.Now().Unix()

	logger.Debug("handleGetMinerNodes method ended with success")
	return
}

func (c *Client) handleProposeDistribution(body []byte) (protoMessage protoMessageWithID, err error) {
	req := &external.ProposeDistributionRequest{}
	err = proto.Unmarshal(body, req)
	if err != nil {
		err = errors.Wrap(err, "failed to unmarshal request")
		return
	}
	response := &external.ProposeDistributionResponse{
		Id: req.Id,
	}

	logger := log.New("requestId", req.GetId())
	defer func() {
		if err != nil {
			jsonReq, _ := json.Marshal(req)
			logger.Warn("handleProposeDistribution method ended with error", "request", string(jsonReq), "error", err)

			response.Error = c.handleError(err)
			err = errors.Wrapf(err, "handle request (%s) error", req.GetId())
		}

		protoMessage = response
	}()

	err = c.bitbon.GetAPI().ProposeDistribution(loggerContext.NewLoggerContext(
		context.Background(), logger), dto.ProposeDistributionRequest{
		Address: common.HexToAddress(req.Address),
		CryptoData: dto.AssetboxCryptoData{
			Wallet:     req.GetCryptoData().GetWallet(),
			Passphrase: req.GetCryptoData().GetPassphrase(),
		},
		Distribution: req.Distribution,
	})
	if err != nil {
		err = errors.Wrap(err, "error calling api method")
		return
	}

	response.ApplicationTimestamp = time.Now().Unix()

	logger.Debug("handleProposeDistribution method ended with success")
	return
}

func (c *Client) handleGetCurrentDistribution(body []byte) (protoMessage protoMessageWithID, err error) {
	req := &external.MinerNodesRequest{}
	err = proto.Unmarshal(body, req)
	if err != nil {
		err = errors.Wrap(err, "failed to unmarshal request")
		return
	}
	response := &external.GetCurrentDistributionResponse{
		Id: req.Id,
	}

	logger := log.New("requestId", req.GetId())
	defer func() {
		if err != nil {
			jsonReq, _ := json.Marshal(req)
			logger.Warn("handleGetCurrentDistribution method ended with error", "request", string(jsonReq), "error", err)

			response.Error = c.handleError(err)
			err = errors.Wrapf(err, "handle request (%s) error", req.GetId())
		}

		protoMessage = response
	}()

	resp, err := c.bitbon.GetAPI().GetCurrentDistribution(loggerContext.NewLoggerContext(context.Background(), logger))
	if err != nil {
		err = errors.Wrap(err, "error calling api method")
		return
	}

	response.Distribution = resp
	response.ApplicationTimestamp = time.Now().Unix()

	logger.Debug("handleGetCurrentDistribution method ended with success")
	return
}
