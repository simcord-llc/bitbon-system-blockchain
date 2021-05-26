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

// nolint:dupl
package bitbon

import (
	"context"
	"fmt"
	"math/big"

	"github.com/pkg/errors"
	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/contracts"
	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/dto"
	bitbonErrors "github.com/simcord-llc/bitbon-system-blockchain/bitbon/errors"
	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/models"
	"github.com/simcord-llc/bitbon-system-blockchain/common"
)

func (api *API) QuickTransfer(ctx context.Context, transfer *dto.QuickTransferRequest) (*dto.TransferResponse, error) {
	if api.bitbon.APIStopped() {
		return nil, ErrAPIStopped
	}
	api.bitbon.apiWG.Add(1)
	defer api.bitbon.apiWG.Done()

	if transfer == nil {
		return nil, bitbonErrors.NewInvalidRequestError(errors.New("empty request given"))
	}
	resp, err := api.bitbon.GetTransferManager().QuickTransfer(ctx, models.QuickTransferObjFromDTO(transfer))
	if err != nil {
		return nil, err
	}
	return resp.ToDTO(), nil
}

func (api *API) CreateWPCSafeTransfer(ctx context.Context,
	transfer *dto.CreateWPCSafeTransferRequest) (*dto.TransferResponse, error) {
	if api.bitbon.APIStopped() {
		return nil, ErrAPIStopped
	}
	api.bitbon.apiWG.Add(1)
	defer api.bitbon.apiWG.Done()

	if transfer == nil {
		return nil, bitbonErrors.NewInvalidRequestError(errors.New("empty request given"))
	}
	resp, err := api.bitbon.GetTransferManager().CreateWPCSafeTransfer(ctx, models.CreateWPCTransferObjFromDTO(transfer))
	if err != nil {
		return nil, err
	}
	return resp.ToDTO(), nil
}

func (api *API) ApproveWPCSafeTransfer(ctx context.Context,
	transfer *dto.ApproveWPCSafeTransferRequest) (*dto.TransferResponse, error) {
	if api.bitbon.APIStopped() {
		return nil, ErrAPIStopped
	}
	api.bitbon.apiWG.Add(1)
	defer api.bitbon.apiWG.Done()

	if transfer == nil {
		return nil, bitbonErrors.NewInvalidRequestError(errors.New("empty request given"))
	}
	resp, err := api.bitbon.GetTransferManager().ApproveWPCSafeTransfer(ctx, models.ApproveWPCTransferObjFromDTO(transfer))
	if err != nil {
		return nil, err
	}
	return resp.ToDTO(), nil
}

func (api *API) CancelWPCSafeTransfer(ctx context.Context,
	transfer *dto.CancelWPCSafeTransferRequest) (*dto.TransferResponse, error) {
	if api.bitbon.APIStopped() {
		return nil, ErrAPIStopped
	}
	api.bitbon.apiWG.Add(1)
	defer api.bitbon.apiWG.Done()

	if transfer == nil {
		return nil, bitbonErrors.NewInvalidRequestError(errors.New("empty request given"))
	}
	resp, err := api.bitbon.GetTransferManager().CancelWPCSafeTransfer(ctx, models.CancelWPCTransferObjFromDTO(transfer))
	if err != nil {
		return nil, err
	}
	return resp.ToDTO(), nil
}

func (api *API) CreateSafeTransfer(ctx context.Context,
	transfer *dto.CreateSafeTransferRequest) (*dto.TransferResponse, error) {
	if api.bitbon.APIStopped() {
		return nil, ErrAPIStopped
	}
	api.bitbon.apiWG.Add(1)
	defer api.bitbon.apiWG.Done()

	if transfer == nil {
		return nil, bitbonErrors.NewInvalidRequestError(errors.New("empty request given"))
	}
	resp, err := api.bitbon.GetTransferManager().CreateSafeTransfer(ctx, models.CreateTransferObjFromDTO(transfer))
	if err != nil {
		return nil, err
	}
	return resp.ToDTO(), nil
}

func (api *API) ApproveSafeTransfer(ctx context.Context,
	transfer *dto.ApproveSafeTransferRequest) (*dto.TransferResponse, error) {
	if api.bitbon.APIStopped() {
		return nil, ErrAPIStopped
	}
	api.bitbon.apiWG.Add(1)
	defer api.bitbon.apiWG.Done()

	if transfer == nil {
		return nil, bitbonErrors.NewInvalidRequestError(errors.New("empty request given"))
	}
	resp, err := api.bitbon.GetTransferManager().ApproveSafeTransfer(ctx, models.ApproveTransferObjFromDTO(transfer))
	if err != nil {
		return nil, err
	}
	return resp.ToDTO(), nil
}

func (api *API) CancelSafeTransfer(ctx context.Context,
	transfer *dto.CancelSafeTransferRequest) (*dto.TransferResponse, error) {
	if api.bitbon.APIStopped() {
		return nil, ErrAPIStopped
	}
	api.bitbon.apiWG.Add(1)
	defer api.bitbon.apiWG.Done()

	if transfer == nil {
		return nil, bitbonErrors.NewInvalidRequestError(errors.New("empty request given"))
	}
	resp, err := api.bitbon.GetTransferManager().CancelSafeTransfer(ctx, models.CancelTransferObjFromDTO(transfer))
	if err != nil {
		return nil, err
	}
	return resp.ToDTO(), nil
}

func (api *API) DirectTransfer(ctx context.Context,
	transfer *dto.DirectTransferRequest) (*dto.TransferResponse, error) {
	if api.bitbon.APIStopped() {
		return nil, ErrAPIStopped
	}
	api.bitbon.apiWG.Add(1)
	defer api.bitbon.apiWG.Done()

	if transfer == nil {
		return nil, bitbonErrors.NewInvalidRequestError(errors.New("empty request given"))
	}
	resp, err := api.bitbon.GetTransferManager().DirectTransfer(ctx, models.DirectTransferObjFromDTO(transfer))
	if err != nil {
		return nil, err
	}
	return resp.ToDTO(), nil
}

// Async Transfers - send txHash to client, without waiting while tx has been mined
func (api *API) QuickTransferAsync(ctx context.Context, transfer *dto.QuickTransferRequest) (*common.Hash, error) {
	if api.bitbon.APIStopped() {
		return nil, ErrAPIStopped
	}
	api.bitbon.apiWG.Add(1)
	defer api.bitbon.apiWG.Done()

	if transfer == nil {
		return nil, bitbonErrors.NewInvalidRequestError(errors.New("empty request given"))
	}

	resp, err := api.bitbon.GetTransferManager().QuickTransferAsync(ctx, models.QuickTransferObjFromDTO(transfer))
	if err != nil {
		return nil, err
	}
	return &resp.TxHash, nil
}

func (api *API) CreateWPCSafeTransferAsync(ctx context.Context,
	transfer *dto.CreateWPCSafeTransferRequest) (*common.Hash, error) {
	if api.bitbon.APIStopped() {
		return nil, ErrAPIStopped
	}
	api.bitbon.apiWG.Add(1)
	defer api.bitbon.apiWG.Done()

	if transfer == nil {
		return nil, bitbonErrors.NewInvalidRequestError(errors.New("empty request given"))
	}

	resp, err := api.bitbon.GetTransferManager().
		CreateWPCSafeTransferAsync(ctx, models.CreateWPCTransferObjFromDTO(transfer))
	if err != nil {
		return nil, err
	}
	return &resp.TxHash, nil
}

func (api *API) ApproveWPCSafeTransferAsync(ctx context.Context,
	transfer *dto.ApproveWPCSafeTransferRequest) (*common.Hash, error) {
	if api.bitbon.APIStopped() {
		return nil, ErrAPIStopped
	}
	api.bitbon.apiWG.Add(1)
	defer api.bitbon.apiWG.Done()

	if transfer == nil {
		return nil, bitbonErrors.NewInvalidRequestError(errors.New("empty request given"))
	}
	resp, err := api.bitbon.GetTransferManager().
		ApproveWPCSafeTransferAsync(ctx, models.ApproveWPCTransferObjFromDTO(transfer))
	if err != nil {
		return nil, err
	}
	return &resp.TxHash, nil
}

func (api *API) CancelWPCSafeTransferAsync(ctx context.Context,
	transfer *dto.CancelWPCSafeTransferRequest) (*common.Hash, error) {
	if api.bitbon.APIStopped() {
		return nil, ErrAPIStopped
	}
	api.bitbon.apiWG.Add(1)
	defer api.bitbon.apiWG.Done()

	if transfer == nil {
		return nil, bitbonErrors.NewInvalidRequestError(errors.New("empty request given"))
	}
	resp, err := api.bitbon.GetTransferManager().
		CancelWPCSafeTransferAsync(ctx, models.CancelWPCTransferObjFromDTO(transfer))
	if err != nil {
		return nil, err
	}

	return &resp.TxHash, nil
}

func (api *API) CreateSafeTransferAsync(ctx context.Context,
	transfer *dto.CreateSafeTransferRequest) (*common.Hash, error) {
	if api.bitbon.APIStopped() {
		return nil, ErrAPIStopped
	}
	api.bitbon.apiWG.Add(1)
	defer api.bitbon.apiWG.Done()

	if transfer == nil {
		return nil, bitbonErrors.NewInvalidRequestError(errors.New("empty request given"))
	}

	resp, err := api.bitbon.GetTransferManager().CreateSafeTransferAsync(ctx, models.CreateTransferObjFromDTO(transfer))
	if err != nil {
		return nil, err
	}
	return &resp.TxHash, nil
}

func (api *API) ApproveSafeTransferAsync(ctx context.Context,
	transfer *dto.ApproveSafeTransferRequest) (*common.Hash, error) {
	if api.bitbon.APIStopped() {
		return nil, ErrAPIStopped
	}
	api.bitbon.apiWG.Add(1)
	defer api.bitbon.apiWG.Done()

	if transfer == nil {
		return nil, bitbonErrors.NewInvalidRequestError(errors.New("empty request given"))
	}
	resp, err := api.bitbon.GetTransferManager().ApproveSafeTransferAsync(ctx, models.ApproveTransferObjFromDTO(transfer))
	if err != nil {
		return nil, err
	}
	return &resp.TxHash, nil
}

func (api *API) CancelSafeTransferAsync(ctx context.Context,
	transfer *dto.CancelSafeTransferRequest) (*common.Hash, error) {
	if api.bitbon.APIStopped() {
		return nil, ErrAPIStopped
	}
	api.bitbon.apiWG.Add(1)
	defer api.bitbon.apiWG.Done()

	if transfer == nil {
		return nil, bitbonErrors.NewInvalidRequestError(errors.New("empty request given"))
	}
	resp, err := api.bitbon.GetTransferManager().CancelSafeTransferAsync(ctx, models.CancelTransferObjFromDTO(transfer))
	if err != nil {
		return nil, err
	}

	return &resp.TxHash, nil
}

func (api *API) DirectTransferAsync(ctx context.Context, transfer *dto.DirectTransferRequest) (*common.Hash, error) {
	if api.bitbon.APIStopped() {
		return nil, ErrAPIStopped
	}
	api.bitbon.apiWG.Add(1)
	defer api.bitbon.apiWG.Done()

	if transfer == nil {
		return nil, bitbonErrors.NewInvalidRequestError(errors.New("empty request given"))
	}
	resp, err := api.bitbon.GetTransferManager().DirectTransferAsync(ctx, models.DirectTransferObjFromDTO(transfer))
	if err != nil {
		return nil, err
	}
	return &resp.TxHash, nil
}

func (api *API) ExpireTransfers(ctx context.Context) (*dto.ExpireTransferResponse, error) {
	if api.bitbon.APIStopped() {
		return nil, ErrAPIStopped
	}
	api.bitbon.apiWG.Add(1)
	defer api.bitbon.apiWG.Done()

	hashes, num, err := api.bitbon.GetTransferManager().ExpireTransfers(ctx)
	if err != nil {
		return nil, bitbonErrors.NewInternalError(err)
	}

	return &dto.ExpireTransferResponse{
		TxHashes:   hashes,
		ExpiredNum: num,
	}, nil
}

// GetTransfer gets transfer data from blockchain
func (api *API) GetTransfer(ctx context.Context, transferID string) (*contracts.ReceiptTransfer, error) {
	if api.bitbon.APIStopped() {
		return nil, ErrAPIStopped
	}
	api.bitbon.apiWG.Add(1)
	defer api.bitbon.apiWG.Done()

	if transferID == "" {
		return nil, bitbonErrors.NewInvalidRequestError(errors.New("empty request given"))
	}

	return api.bitbon.GetTransferManager().GetTransfer(ctx, transferID)
}

// GetTransferByIndex gets transfer data from blockchain
func (api *API) GetTransferByIndex(ctx context.Context, index *big.Int) (*contracts.ReceiptTransfer, error) {
	if api.bitbon.APIStopped() {
		return nil, ErrAPIStopped
	}
	api.bitbon.apiWG.Add(1)
	defer api.bitbon.apiWG.Done()

	if index == nil {
		return nil, bitbonErrors.NewInvalidRequestError(errors.New("empty request given"))
	}
	fmt.Println(index.String())

	return api.bitbon.GetTransferManager().GetTransferByIndex(ctx, index)
}

// TransferExists checks if transfer exists blockchain
func (api *API) TransferExists(ctx context.Context, transferID string) (bool, error) {
	if api.bitbon.APIStopped() {
		return false, ErrAPIStopped
	}
	api.bitbon.apiWG.Add(1)
	defer api.bitbon.apiWG.Done()

	if transferID == "" {
		return false, bitbonErrors.NewInvalidRequestError(errors.New("empty request given"))
	}

	return api.bitbon.GetTransferManager().TransferExists(ctx, transferID)
}
