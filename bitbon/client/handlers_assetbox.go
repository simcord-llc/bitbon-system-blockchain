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
	"fmt"
	"strings"

	"github.com/golang/protobuf/proto" // nolint:staticcheck
	"github.com/pkg/errors"
	loggerContext "github.com/simcord-llc/bitbon-system-blockchain/bitbon/context"
	apiDto "github.com/simcord-llc/bitbon-system-blockchain/bitbon/dto"
	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/dto/external"
	"github.com/simcord-llc/bitbon-system-blockchain/common"
	"github.com/simcord-llc/bitbon-system-blockchain/log"
)

const currency = "bitbon"

func (c *Client) handlePrepareAssetbox(body []byte) (protoMessage protoMessageWithID, err error) {
	req := &external.PrepareAssetboxRequest{}
	if err = unmarshalRequest(body, req); err != nil {
		return
	}
	response := &external.PrepareAssetboxResponse{
		Id: req.Id,
	}

	logger := log.New("requestId", req.GetId())
	defer func() {
		if err != nil {
			err = deferFunc(err, req, req.GetId(), "handlePrepareAssetbox", logger)
			response.Error = c.handleError(err)
		}

		protoMessage = response
	}()

	apiRequest := &apiDto.PrepareAssetboxesRequest{
		Count: req.GetCount(),
	}

	assetboxes, err := c.bitbon.GetAPI().PrepareAssetboxes(
		loggerContext.NewLoggerContext(context.Background(), logger), apiRequest)
	if err != nil {
		err = errors.Wrap(err, "error calling api method")
		return
	}
	assetboxData := make(map[string]*external.AssetboxCryptoData)
	for idx := range assetboxes {
		assetboxData[assetboxes[idx].Address.Hex()] = &external.AssetboxCryptoData{
			Wallet:     assetboxes[idx].Wallet,
			Passphrase: assetboxes[idx].PassPhrase,
		}
	}

	response.AddressData = assetboxData
	logger.Debug("handlePrepareAssetbox method ended with success")
	return
}

func (c *Client) handleDeleteAssetbox(body []byte) (protoMessage protoMessageWithID, err error) {
	req := &external.DeleteAssetboxRequest{}
	if err = unmarshalRequest(body, req); err != nil {
		return
	}
	response := &external.DeleteAssetboxResponse{
		Id: req.Id,
	}

	logger := log.New("requestId", req.GetId())
	defer func() {
		if err != nil {
			err = deferFunc(err, req, req.GetId(), "handleDeleteAssetbox", logger)
			response.Error = c.handleError(err)
		}
		protoMessage = response
	}()

	apiRequest := &apiDto.DeleteAssetboxRequest{
		Address: common.HexToAddress(req.GetAddress()),
		CryptoData: apiDto.AssetboxCryptoData{
			Wallet:     req.GetCryptoData().GetWallet(),
			Passphrase: req.GetCryptoData().GetPassphrase(),
		},
	}

	err = c.bitbon.GetAPI().DeleteAssetbox(loggerContext.NewLoggerContext(context.Background(), logger), apiRequest)
	if err != nil {
		err = errors.Wrap(err, "error calling api method")
		return
	}

	logger.Debug("handleDeleteAssetbox method ended with success")
	return
}

func (c *Client) handleSetPublicAssetboxInfo(body []byte) (protoMessage protoMessageWithID, err error) {
	req := &external.SetPublicAssetboxInfoRequest{}
	if err = unmarshalRequest(body, req); err != nil {
		return
	}
	response := &external.SetPublicAssetboxInfoResponse{
		Id: req.Id,
	}

	logger := log.New("requestId", req.GetId())
	logger.Debug("handleSetPublicAssetboxInfo new message")
	defer func() {
		if err != nil {
			err = deferFunc(err, req, req.GetId(), "handleSetPublicAssetboxInfo", logger)
			response.Error = c.handleError(err)
		}

		protoMessage = response
	}()

	apiRequest := &apiDto.SetPublicAssetboxInfoRequest{
		Address:   common.HexToAddress(req.GetAddress()),
		Alias:     req.GetAlias(),
		IsPublic:  req.GetIsPublic(),
		ExtraInfo: req.GetExtractInfo(),
		IsMining:  req.GetIsMining(),
		ServiceID: req.GetServiceId(),
		CryptoData: apiDto.AssetboxCryptoData{
			Wallet:     req.GetCryptoData().GetWallet(),
			Passphrase: req.GetCryptoData().GetPassphrase(),
		},
	}

	err = c.bitbon.GetAPI().SetPublicAssetboxInfo(loggerContext.NewLoggerContext(context.Background(), logger), apiRequest)
	if err != nil {
		err = errors.Wrap(err, "error calling api method")
		return
	}

	logger.Debug("handleSetPublicAssetboxInfo method ended with success")
	return
}

func (c *Client) handleAssetboxBalances(body []byte) (protoMessage protoMessageWithID, err error) {
	req := &external.AssetboxBalancesRequest{}
	if err = unmarshalRequest(body, req); err != nil {
		return
	}
	response := &external.AssetboxBalancesResponse{
		Id: req.Id,
	}

	logger := log.New("requestId", req.GetId())
	defer func() {
		if err != nil {
			err = deferFunc(err, req, req.GetId(), "handleAssetboxBalances", logger)
			response.Error = c.handleError(err)
		}

		protoMessage = response
	}()

	if req.GetAddresses() == nil {
		err = errors.New("addresses field in request is not set")
		return
	}

	addresses := make([]common.Address, len(req.GetAddresses()))
	for k, v := range req.GetAddresses() {
		addresses[k] = common.HexToAddress(v)
	}

	apiResponse, err := c.bitbon.GetAPI().GetAssetboxBalances(
		loggerContext.NewLoggerContext(context.Background(), logger), addresses)
	if err != nil {
		err = errors.Wrap(err, "error calling api method")
		return
	}

	response.BalanceResponse = make([]*external.AssetboxBalanceResponse, 0, len(apiResponse))
	for addr, bal := range apiResponse {
		balanceResponse := &external.AssetboxBalanceResponse{
			Address: strings.ToLower(addr.Hex()),
			Balances: []*external.AssetboxBalance{
				{
					Currency: currency,
					Amount:   bal.String(),
				},
			},
		}
		response.BalanceResponse = append(response.BalanceResponse, balanceResponse)
	}

	logger.Debug("handleAssetboxBalances method ended with success")
	return
}

func deferFunc(err error, req interface{}, id, methodName string, logger log.Logger) error {
	if err != nil {
		jsonReq, _ := json.Marshal(req)
		logger.Warn(fmt.Sprintf("%s method ended with error", methodName), "request", string(jsonReq), "error", err)
		err = errors.Wrapf(err, "handle request (%s) error", id)
	}
	return err
}

func unmarshalRequest(body []byte, req proto.Message) error {
	err := proto.Unmarshal(body, req)
	if err != nil {
		err = errors.Wrap(err, "failed to unmarshal request")
		return err
	}
	return nil
}
