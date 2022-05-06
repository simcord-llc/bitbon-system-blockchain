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

package assetbox

import (
	"context"
	"fmt"

	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/dto"

	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/models"

	"github.com/pkg/errors"
	bb "github.com/simcord-llc/bitbon-system-blockchain/bitbon"
	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/contracts"
	bitbonErrors "github.com/simcord-llc/bitbon-system-blockchain/bitbon/errors"
)

func (am *Manager) PrepareAssetbox(_ context.Context) (assetbox *models.Assetbox, err error) {
	// create unbound assetbox without public info
	assetbox, err = am.newUnboundAssetbox()
	if err != nil {
		return nil, errors.Wrap(err, "error instantiating new assetbox")
	}

	// encrypt wallet and passphrase
	w, err := am.encryptor.Encrypt(assetbox.Wallet, []byte(am.bitbon.GetDecryptAssetboxWalletPassword()))
	if err != nil {
		return nil, err
	}
	p, err := am.encryptor.Encrypt(assetbox.PassPhrase, []byte(am.bitbon.GetDecryptAssetboxWalletPassword()))
	if err != nil {
		return nil, err
	}

	assetbox.Wallet = w
	assetbox.PassPhrase = p

	return assetbox, nil
}

func (am *Manager) PrepareAssetboxes(ctx context.Context, count uint64) (assetboxes []*models.Assetbox, err error) {
	am.logger.Debug("PrepareAssetboxes called", "count", count)

	if count > bb.MaxPrepareAssetboxBatchSize {
		return nil, bitbonErrors.NewInvalidParamsError(errors.New(fmt.Sprintf(
			"requested prepared assetboxes count exceeds max batch size of %d", bb.MaxPrepareAssetboxBatchSize)))
	}

	if count == 0 {
		count = bb.DefaultPrepareAssetboxBatchSize
	}
	assetboxes = make([]*models.Assetbox, 0, count)

	var lastError error

	for i := uint64(1); i <= count; i++ {
		for retries := 1; retries <= bb.MaxPrepareAssetboxRetries; retries++ {
			assetbox, err := am.PrepareAssetbox(ctx)
			if err != nil {
				lastError = err
				if retries == bb.MaxPrepareAssetboxRetries {
					// if we tried several times to prepare assetbox and there is no success - check assetboxes len
					// if assetboxes is empty - do not overload the system and return with error
					if len(assetboxes) == 0 {
						return nil, err
					}
					am.logger.Error("number of error retries while preparing assetboxes exceeded max retries",
						"retries", bb.MaxPrepareAssetboxRetries, "error", err)
				}
				continue
			}
			assetboxes = append(assetboxes, assetbox)
			break
		}
	}
	if len(assetboxes) == 0 {
		return nil, errors.Wrap(lastError, "unable to prepare assetboxes")
	}
	return assetboxes, nil
}

func (am *Manager) SetPublicAssetboxInfo(ctx context.Context, req *dto.SetPublicAssetboxInfoRequest) (err error) {
	am.logger.Debug("SetPublicAssetboxInfo called", "address",
		req.Address.Hex(), "alias", req.Alias, "isPublic", req.IsPublic, "extraInfo", req.ExtraInfo, "isMiner", req.IsMining)

	assetbox := &models.Assetbox{
		Address:    req.Address,
		Alias:      req.Alias,
		ServiceID:  req.ServiceID,
		IsPublic:   req.IsPublic,
		ExtraInfo:  req.ExtraInfo,
		IsMining:   req.IsMining,
		Wallet:     req.CryptoData.Wallet,
		PassPhrase: req.CryptoData.Passphrase,
	}

	privateKey, err := bb.DecryptPrivateKeyForAssetbox(assetbox.Address, assetbox.Wallet, assetbox.PassPhrase, am.bitbon.GetDecryptAssetboxWalletPassword(), am.encryptor.Decrypt)
	if err != nil {
		return err
	}
	assetbox.Pk = privateKey

	err = am.bitbon.GetContractsManager().SetAssetboxInfo(ctx, assetbox.ToContractsAssetbox(), assetbox.Pk)
	if err != nil {
		return errors.Wrap(err, "error setting assetbox info to blockchain")
	}

	return nil
}

func (am *Manager) SetPublicAssetboxInfos(ctx context.Context, assetboxes []*contracts.Assetbox) (err error) {
	am.logger.Debug("SetPublicAssetboxInfos called", "address", assetboxes)

	err = am.bitbon.GetContractsManager().SetAssetboxInfos(ctx, assetboxes, am.bitbon.GetServicePk())
	if err != nil {
		return errors.Wrap(err, "error setting assetboxes info to blockchain")
	}

	return nil
}

func (am *Manager) DeleteAssetbox(ctx context.Context, req *dto.DeleteAssetboxRequest) (err error) {
	am.logger.Debug("DeleteAssetbox called", "address", req.Address.Hex())

	privateKey, err := bb.DecryptPrivateKeyForAssetbox(req.Address, req.CryptoData.Wallet, req.CryptoData.Passphrase, am.bitbon.GetDecryptAssetboxWalletPassword(), am.encryptor.Decrypt)
	if err != nil {
		return err
	}

	err = am.deleteAssetbox(ctx, privateKey)
	if err != nil {
		return bitbonErrors.NewDeleteLastAssetboxError(err)
	}

	return nil
}
