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
	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
	loggerContext "github.com/simcord-llc/bitbon-system-blockchain/bitbon/context"
	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/dto/external"
	"github.com/simcord-llc/bitbon-system-blockchain/log"
)

func (c *Client) handleFeeSettingsRequest(body []byte) (protoMessage protoMessageWithID, err error) {
	req := &external.FeeSettingsRequest{}
	err = proto.Unmarshal(body, req)
	if err != nil {
		err = errors.Wrap(err, "failed to unmarshal request")
		return
	}
	response := &external.FeeSettingsResponse{
		Id: req.Id,
	}

	logger := log.New("requestId", req.GetId())
	defer func() {
		if err != nil {
			jsonReq, _ := json.Marshal(req)
			logger.Warn("handleFeeSettingsRequest method ended with error",
				"request", string(jsonReq), "error", err)

			response.Error = c.handleError(err)
			err = errors.Wrapf(err, "handle request (%s) error", req.GetId())
		}

		protoMessage = response
	}()

	apiResponse, err := c.bitbon.GetAPI().GetFeeSettings(loggerContext.NewLoggerContext(context.Background(), logger))
	if err != nil {
		err = errors.Wrap(err, "error calling api method")
		return
	}

	response.ValueSettings = apiResponse.ValueSettings
	response.DistributionSettings = apiResponse.DistributionSettingsToExternalDTO()

	logger.Debug("handleFeeSettingsRequest method ended with success")
	return
}
