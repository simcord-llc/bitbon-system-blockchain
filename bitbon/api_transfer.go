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

var errEmptyRequest = errors.New("empty request given")

func (api *API) QuickTransfer(ctx context.Context, transfer *dto.QuickTransferRequest) (*dto.TransferResponse, error) {
	if api.bitbon.APIStopped() {
		return nil, ErrAPIStopped
	}
	api.bitbon.apiWG.Add(1)
	defer api.bitbon.apiWG.Done()

	if transfer == nil {
		return nil, bitbonErrors.NewInvalidRequestError(errEmptyRequest)
	}
	resp, err := api.bitbon.GetTransferManager().QuickTransfer(ctx, models.QuickTransferObjFromDTO(transfer))
	if err != nil {
		return nil, err
	}
	return resp.ToDTO(), nil
}

func (api *API) FrameTransfer(ctx context.Context, transfer *dto.QuickTransferRequest) (*dto.TransferResponse, error) {
	if api.bitbon.APIStopped() {
		return nil, ErrAPIStopped
	}
	api.bitbon.apiWG.Add(1)
	defer api.bitbon.apiWG.Done()

	if transfer == nil {
		return nil, bitbonErrors.NewInvalidRequestError(errEmptyRequest)
	}
	resp, err := api.bitbon.GetTransferManager().FrameTransfer(ctx, models.QuickTransferObjFromDTO(transfer))
	if err != nil {
		return nil, err
	}
	return resp.ToDTO(), nil
}

func (api *API) ServiceFeeTransfer(ctx context.Context, transfer *dto.ServiceFeeTransferRequest) (*dto.TransferResponse, error) {
	if api.bitbon.APIStopped() {
		return nil, ErrAPIStopped
	}
	api.bitbon.apiWG.Add(1)
	defer api.bitbon.apiWG.Done()

	if transfer == nil {
		return nil, bitbonErrors.NewInvalidRequestError(errEmptyRequest)
	}
	resp, err := api.bitbon.GetTransferManager().ServiceFeeTransfer(ctx, models.FeeTransferObjFromDTO(transfer))
	if err != nil {
		return nil, err
	}
	return resp.DTO(), nil
}

func (api *API) CreateWPCSafeTransfer(ctx context.Context,
	transfer *dto.CreateWPCSafeTransferRequest) (*dto.TransferResponse, error) {
	if api.bitbon.APIStopped() {
		return nil, ErrAPIStopped
	}
	api.bitbon.apiWG.Add(1)
	defer api.bitbon.apiWG.Done()

	if transfer == nil {
		return nil, bitbonErrors.NewInvalidRequestError(errEmptyRequest)
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
		return nil, bitbonErrors.NewInvalidRequestError(errEmptyRequest)
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
		return nil, bitbonErrors.NewInvalidRequestError(errEmptyRequest)
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
		return nil, bitbonErrors.NewInvalidRequestError(errEmptyRequest)
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
		return nil, bitbonErrors.NewInvalidRequestError(errEmptyRequest)
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
		return nil, bitbonErrors.NewInvalidRequestError(errEmptyRequest)
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
		return nil, bitbonErrors.NewInvalidRequestError(errEmptyRequest)
	}
	resp, err := api.bitbon.GetTransferManager().DirectTransfer(ctx, models.DirectTransferObjFromDTO(transfer))
	if err != nil {
		return nil, err
	}
	return resp.ToDTO(), nil
}

func (api *API) FullBalanceQuickTransfer(ctx context.Context, transfer *dto.FullBalanceQuickTransferRequest) (*dto.TransferResponse, error) {
	if api.bitbon.APIStopped() {
		return nil, ErrAPIStopped
	}
	api.bitbon.apiWG.Add(1)
	defer api.bitbon.apiWG.Done()

	if transfer == nil {
		return nil, bitbonErrors.NewInvalidRequestError(errEmptyRequest)
	}
	resp, err := api.bitbon.GetTransferManager().FullBalanceQuickTransfer(ctx, models.FullBalanceQuickTransferObjFromDTO(transfer))
	if err != nil {
		return nil, err
	}
	return resp.ToDTO(), nil
}

func (api *API) CreateFullBalanceWPCSafeTransfer(ctx context.Context, transfer *dto.CreateFullBalanceWPCSafeTransferRequest) (*dto.TransferResponse, error) {
	if api.bitbon.APIStopped() {
		return nil, ErrAPIStopped
	}
	api.bitbon.apiWG.Add(1)
	defer api.bitbon.apiWG.Done()

	if transfer == nil {
		return nil, bitbonErrors.NewInvalidRequestError(errEmptyRequest)
	}
	resp, err := api.bitbon.GetTransferManager().CreateFullBalanceWPCSafeTransfer(ctx, models.CreateFullBalanceWPCTransferObjFromDTO(transfer))
	if err != nil {
		return nil, err
	}
	return resp.ToDTO(), nil
}

func (api *API) ApproveFullBalanceWPCSafeTransfer(ctx context.Context, transfer *dto.ApproveFullBalanceWPCSafeTransferRequest) (*dto.TransferResponse, error) {
	if api.bitbon.APIStopped() {
		return nil, ErrAPIStopped
	}
	api.bitbon.apiWG.Add(1)
	defer api.bitbon.apiWG.Done()

	if transfer == nil {
		return nil, bitbonErrors.NewInvalidRequestError(errEmptyRequest)
	}
	resp, err := api.bitbon.GetTransferManager().ApproveFullBalanceWPCSafeTransfer(ctx, models.ApproveFullBalanceWPCTransferObjFromDTO(transfer))
	if err != nil {
		return nil, err
	}
	return resp.ToDTO(), nil
}

func (api *API) CancelFullBalanceWPCSafeTransfer(ctx context.Context, transfer *dto.CancelFullBalanceWPCSafeTransferRequest) (*dto.TransferResponse, error) {
	if api.bitbon.APIStopped() {
		return nil, ErrAPIStopped
	}
	api.bitbon.apiWG.Add(1)
	defer api.bitbon.apiWG.Done()

	if transfer == nil {
		return nil, bitbonErrors.NewInvalidRequestError(errEmptyRequest)
	}
	resp, err := api.bitbon.GetTransferManager().CancelFullBalanceWPCSafeTransfer(ctx, models.CancelFullBalanceWPCTransferObjFromDTO(transfer))
	if err != nil {
		return nil, err
	}
	return resp.ToDTO(), nil
}

func (api *API) CreateFullBalanceSafeTransfer(ctx context.Context, transfer *dto.CreateFullBalanceSafeTransferRequest) (*dto.TransferResponse, error) {
	if api.bitbon.APIStopped() {
		return nil, ErrAPIStopped
	}
	api.bitbon.apiWG.Add(1)
	defer api.bitbon.apiWG.Done()

	if transfer == nil {
		return nil, bitbonErrors.NewInvalidRequestError(errEmptyRequest)
	}
	resp, err := api.bitbon.GetTransferManager().CreateFullBalanceSafeTransfer(ctx, models.CreateFullBalanceTransferObjFromDTO(transfer))
	if err != nil {
		return nil, err
	}
	return resp.ToDTO(), nil
}

func (api *API) ApproveFullBalanceSafeTransfer(ctx context.Context, transfer *dto.ApproveFullBalanceSafeTransferRequest) (*dto.TransferResponse, error) {
	if api.bitbon.APIStopped() {
		return nil, ErrAPIStopped
	}
	api.bitbon.apiWG.Add(1)
	defer api.bitbon.apiWG.Done()

	if transfer == nil {
		return nil, bitbonErrors.NewInvalidRequestError(errEmptyRequest)
	}
	resp, err := api.bitbon.GetTransferManager().ApproveFullBalanceSafeTransfer(ctx, models.ApproveFullBalanceTransferObjFromDTO(transfer))
	if err != nil {
		return nil, err
	}
	return resp.ToDTO(), nil
}

func (api *API) CancelFullBalanceSafeTransfer(ctx context.Context, transfer *dto.CancelFullBalanceSafeTransferRequest) (*dto.TransferResponse, error) {
	if api.bitbon.APIStopped() {
		return nil, ErrAPIStopped
	}
	api.bitbon.apiWG.Add(1)
	defer api.bitbon.apiWG.Done()

	if transfer == nil {
		return nil, bitbonErrors.NewInvalidRequestError(errEmptyRequest)
	}
	resp, err := api.bitbon.GetTransferManager().CancelFullBalanceSafeTransfer(ctx, models.CancelFullBalanceTransferObjFromDTO(transfer))
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
		return nil, bitbonErrors.NewInvalidRequestError(errEmptyRequest)
	}

	resp, err := api.bitbon.GetTransferManager().QuickTransferAsync(ctx, models.QuickTransferObjFromDTO(transfer))
	if err != nil {
		return nil, err
	}
	return &resp.TxHash, nil
}

func (api *API) FrameTransferAsync(ctx context.Context, transfer *dto.QuickTransferRequest) (*common.Hash, error) {
	if api.bitbon.APIStopped() {
		return nil, ErrAPIStopped
	}
	api.bitbon.apiWG.Add(1)
	defer api.bitbon.apiWG.Done()

	if transfer == nil {
		return nil, bitbonErrors.NewInvalidRequestError(errEmptyRequest)
	}

	resp, err := api.bitbon.GetTransferManager().FrameTransferAsync(ctx, models.QuickTransferObjFromDTO(transfer))
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
		return nil, bitbonErrors.NewInvalidRequestError(errEmptyRequest)
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
		return nil, bitbonErrors.NewInvalidRequestError(errEmptyRequest)
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
		return nil, bitbonErrors.NewInvalidRequestError(errEmptyRequest)
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
		return nil, bitbonErrors.NewInvalidRequestError(errEmptyRequest)
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
		return nil, bitbonErrors.NewInvalidRequestError(errEmptyRequest)
	}
	resp, err := api.bitbon.GetTransferManager().ApproveSafeTransferAsync(ctx, models.ApproveTransferObjFromDTO(transfer))
	if err != nil {
		return nil, err
	}
	return &resp.TxHash, nil
}

func (api *API) FeeTransferAsync(ctx context.Context, transfer *dto.ServiceFeeTransferRequest) (*common.Hash, error) {
	if api.bitbon.APIStopped() {
		return nil, ErrAPIStopped
	}
	api.bitbon.apiWG.Add(1)
	defer api.bitbon.apiWG.Done()

	if transfer == nil {
		return nil, bitbonErrors.NewInvalidRequestError(errEmptyRequest)
	}

	resp, err := api.bitbon.GetTransferManager().ServiceFeeTransferAsync(ctx, models.FeeTransferObjFromDTO(transfer))
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
		return nil, bitbonErrors.NewInvalidRequestError(errEmptyRequest)
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
		return nil, bitbonErrors.NewInvalidRequestError(errEmptyRequest)
	}
	resp, err := api.bitbon.GetTransferManager().DirectTransferAsync(ctx, models.DirectTransferObjFromDTO(transfer))
	if err != nil {
		return nil, err
	}
	return &resp.TxHash, nil
}

func (api *API) FullBalanceQuickTransferAsync(ctx context.Context, transfer *dto.FullBalanceQuickTransferRequest) (*common.Hash, error) {
	if api.bitbon.APIStopped() {
		return nil, ErrAPIStopped
	}
	api.bitbon.apiWG.Add(1)
	defer api.bitbon.apiWG.Done()

	if transfer == nil {
		return nil, bitbonErrors.NewInvalidRequestError(errEmptyRequest)
	}

	resp, err := api.bitbon.GetTransferManager().FullBalanceQuickTransferAsync(ctx, models.FullBalanceQuickTransferObjFromDTO(transfer))
	if err != nil {
		return nil, err
	}
	return &resp.TxHash, nil
}

func (api *API) CreateFullBalanceWPCSafeTransferAsync(ctx context.Context, transfer *dto.CreateFullBalanceWPCSafeTransferRequest) (*common.Hash, error) {
	if api.bitbon.APIStopped() {
		return nil, ErrAPIStopped
	}
	api.bitbon.apiWG.Add(1)
	defer api.bitbon.apiWG.Done()

	if transfer == nil {
		return nil, bitbonErrors.NewInvalidRequestError(errEmptyRequest)
	}

	resp, err := api.bitbon.GetTransferManager().CreateFullBalanceWPCSafeTransferAsync(ctx, models.CreateFullBalanceWPCTransferObjFromDTO(transfer))
	if err != nil {
		return nil, err
	}
	return &resp.TxHash, nil
}

func (api *API) ApproveFullBalanceWPCSafeTransferAsync(ctx context.Context, transfer *dto.ApproveFullBalanceWPCSafeTransferRequest) (*common.Hash, error) {
	if api.bitbon.APIStopped() {
		return nil, ErrAPIStopped
	}
	api.bitbon.apiWG.Add(1)
	defer api.bitbon.apiWG.Done()

	if transfer == nil {
		return nil, bitbonErrors.NewInvalidRequestError(errEmptyRequest)
	}
	resp, err := api.bitbon.GetTransferManager().ApproveFullBalanceWPCSafeTransferAsync(ctx, models.ApproveFullBalanceWPCTransferObjFromDTO(transfer))
	if err != nil {
		return nil, err
	}
	return &resp.TxHash, nil
}

func (api *API) CancelFullBalanceWPCSafeTransferAsync(ctx context.Context, transfer *dto.CancelFullBalanceWPCSafeTransferRequest) (*common.Hash, error) {
	if api.bitbon.APIStopped() {
		return nil, ErrAPIStopped
	}
	api.bitbon.apiWG.Add(1)
	defer api.bitbon.apiWG.Done()

	if transfer == nil {
		return nil, bitbonErrors.NewInvalidRequestError(errEmptyRequest)
	}
	resp, err := api.bitbon.GetTransferManager().CancelFullBalanceWPCSafeTransferAsync(ctx, models.CancelFullBalanceWPCTransferObjFromDTO(transfer))
	if err != nil {
		return nil, err
	}

	return &resp.TxHash, nil
}

func (api *API) CreateFullBalanceSafeTransferAsync(ctx context.Context, transfer *dto.CreateFullBalanceSafeTransferRequest) (*common.Hash, error) {
	if api.bitbon.APIStopped() {
		return nil, ErrAPIStopped
	}
	api.bitbon.apiWG.Add(1)
	defer api.bitbon.apiWG.Done()

	if transfer == nil {
		return nil, bitbonErrors.NewInvalidRequestError(errEmptyRequest)
	}

	resp, err := api.bitbon.GetTransferManager().CreateFullBalanceSafeTransferAsync(ctx, models.CreateFullBalanceTransferObjFromDTO(transfer))
	if err != nil {
		return nil, err
	}
	return &resp.TxHash, nil
}

func (api *API) ApproveFullBalanceSafeTransferAsync(ctx context.Context, transfer *dto.ApproveFullBalanceSafeTransferRequest) (*common.Hash, error) {
	if api.bitbon.APIStopped() {
		return nil, ErrAPIStopped
	}
	api.bitbon.apiWG.Add(1)
	defer api.bitbon.apiWG.Done()

	if transfer == nil {
		return nil, bitbonErrors.NewInvalidRequestError(errEmptyRequest)
	}
	resp, err := api.bitbon.GetTransferManager().ApproveFullBalanceSafeTransferAsync(ctx, models.ApproveFullBalanceTransferObjFromDTO(transfer))
	if err != nil {
		return nil, err
	}
	return &resp.TxHash, nil
}

func (api *API) CancelFullBalanceSafeTransferAsync(ctx context.Context, transfer *dto.CancelFullBalanceSafeTransferRequest) (*common.Hash, error) {
	if api.bitbon.APIStopped() {
		return nil, ErrAPIStopped
	}
	api.bitbon.apiWG.Add(1)
	defer api.bitbon.apiWG.Done()

	if transfer == nil {
		return nil, bitbonErrors.NewInvalidRequestError(errEmptyRequest)
	}
	resp, err := api.bitbon.GetTransferManager().CancelFullBalanceSafeTransferAsync(ctx, models.CancelFullBalanceTransferObjFromDTO(transfer))
	if err != nil {
		return nil, err
	}

	return &resp.TxHash, nil
}

func (api *API) ExpireTransfersAsync(ctx context.Context, ids []string) (*common.Hash, error) {
	if api.bitbon.APIStopped() {
		return nil, ErrAPIStopped
	}
	api.bitbon.apiWG.Add(1)
	defer api.bitbon.apiWG.Done()

	resp, err := api.bitbon.GetTransferManager().ExpireTransfersAsync(ctx, ids)
	if err != nil {
		return nil, bitbonErrors.NewInternalError(err)
	}

	return &resp.TxHash, nil
}

// GetTransfer gets transfer data from blockchain
func (api *API) GetTransfer(ctx context.Context, transferID string) (*contracts.ReceiptTransfer, error) {
	if api.bitbon.APIStopped() {
		return nil, ErrAPIStopped
	}
	api.bitbon.apiWG.Add(1)
	defer api.bitbon.apiWG.Done()

	if transferID == "" {
		return nil, bitbonErrors.NewInvalidRequestError(errEmptyRequest)
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
		return nil, bitbonErrors.NewInvalidRequestError(errEmptyRequest)
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
		return false, bitbonErrors.NewInvalidRequestError(errEmptyRequest)
	}

	return api.bitbon.GetTransferManager().TransferExists(ctx, transferID)
}
