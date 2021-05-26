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

package contracts

import (
	"context"
	"crypto/ecdsa"

	"github.com/pkg/errors"
	"github.com/simcord-llc/bitbon-system-blockchain/accounts/abi/bind"
	bitbonErrors "github.com/simcord-llc/bitbon-system-blockchain/bitbon/errors"
	"github.com/simcord-llc/bitbon-system-blockchain/common"
	"github.com/simcord-llc/bitbon-system-blockchain/rlp"
)

const assetboxPublic = 1

type assetboxOtherInfo struct {
	IsPublic  uint8
	ExtraInfo string
}

func (m *Manager) getAssetboxImpl() (*AssetboxImpl, error) {
	if !m.Ready() {
		return nil, ErrorNotReady
	}
	contractAddress, err := m.GetContractAddressAssetbox()
	if err != nil {
		return nil, err
	}
	contract, err := NewAssetboxImpl(contractAddress, m.EthApiWrapper)
	if err != nil {
		return nil, err
	}
	return contract, nil
}

// SetAssetboxInfo sets public assetbox info to blockchain
func (m *Manager) SetAssetboxInfo(ctx context.Context, assetbox *Assetbox, key *ecdsa.PrivateKey) (err error) {
	m.logger.Debug("contracts manager SetAssetboxInfo called", "address", assetbox.Address.Hex())

	defer func() {
		if err != nil {
			err = bitbonErrors.NewContractCallError(err)
		}
	}()

	contract, err := m.getAssetboxImpl()
	if err != nil {
		return err
	}

	aoi := &assetboxOtherInfo{
		ExtraInfo: assetbox.ExtraInfo,
	}
	if assetbox.IsPublic {
		aoi.IsPublic = assetboxPublic
	}
	otherInfo, err := rlp.EncodeToBytes(aoi)
	if err != nil {
		return errors.Wrap(err, "error rlp encoding assetbox other info")
	}

	opts, err := m.prepareTransactOpts(ctx, key)
	if err != nil {
		return err
	}

	tx, err := contract.SetAssetboxInfo(opts, assetbox.Alias, assetbox.ServiceID, assetbox.IsMining, otherInfo)
	if err != nil {
		return err
	}

	_, err = m.waitAndCheckTransaction(ctx, tx)
	return
}

// DeleteAssetboxInfo deletes assetbox data from blockchain
func (m *Manager) DeleteAssetboxInfo(ctx context.Context, key *ecdsa.PrivateKey) (err error) {
	m.logger.Debug("contracts manager DeleteAssetboxInfo called")

	defer func() {
		if err != nil {
			err = bitbonErrors.NewContractCallError(err)
		}
	}()

	contract, err := m.getAssetboxImpl()
	if err != nil {
		return err
	}

	opts, err := m.prepareTransactOpts(ctx, key)
	if err != nil {
		return err
	}

	tx, err := contract.DeleteAssetboxInfo(opts)
	if err != nil {
		return err
	}

	_, err = m.waitAndCheckTransaction(ctx, tx)
	return
}

// SetAssetboxInfos sets multiple public assetbox infos to blockchain
func (m *Manager) SetAssetboxInfos(ctx context.Context, a []*Assetbox, key *ecdsa.PrivateKey) (err error) {
	m.logger.Debug("contracts manager SetAssetboxInfos called")

	defer func() {
		if err != nil {
			err = bitbonErrors.NewContractCallError(err)
		}
	}()

	contract, err := m.getAssetboxImpl()
	if err != nil {
		return err
	}

	assetboxes := make([]common.Address, 0, len(a))
	aliases := make([]string, 0, len(a))
	serviceIds := make([][]byte, 0, len(a))
	otherInfos := make([][]byte, 0, len(a))
	isMinings := make([]bool, 0, len(a))

	for idx := range a {
		assetboxes = append(assetboxes, a[idx].Address)
		aliases = append(aliases, a[idx].Alias)
		serviceIds = append(serviceIds, a[idx].ServiceID)

		aoi := &assetboxOtherInfo{
			ExtraInfo: a[idx].ExtraInfo,
		}
		if a[idx].IsPublic {
			aoi.IsPublic = assetboxPublic
		}
		otherInfo, err := rlp.EncodeToBytes(aoi) // nolint:govet
		if err != nil {
			return errors.Wrap(err, "error rlp encoding assetbox other info")
		}

		otherInfos = append(otherInfos, otherInfo)
		isMinings = append(isMinings, a[idx].IsMining)
	}

	opts, err := m.prepareTransactOpts(ctx, key)
	if err != nil {
		return err
	}

	tx, err := contract.SetAssetboxInfosAdmin(opts, assetboxes, aliases, serviceIds, isMinings, otherInfos)
	if err != nil {
		return err
	}

	_, err = m.waitAndCheckTransaction(ctx, tx)
	return
}

// GetAssetboxInfo gets public assetbox info from blockchain
func (m *Manager) GetAssetboxInfo(ctx context.Context, addr common.Address) (a *Assetbox, err error) {
	m.logger.Debug("contracts manager GetAssetboxInfo called", "address", addr.Hex())

	defer func() {
		if err != nil {
			err = bitbonErrors.NewContractCallError(err)
		}
	}()

	contract, err := m.getAssetboxImpl()
	if err != nil {
		return nil, err
	}

	opts := &bind.CallOpts{
		Pending: false,
		From:    addr,
		Context: ctx,
	}

	info, err := contract.GetMyAssetboxInfo(opts)
	if err != nil {
		return nil, err
	}
	ass, err := ContractAssetbox(info.AssetboxAddress, info.AssetboxAlias, info.ServiceId, info.OtherInfo, info.IsMining)
	if err != nil {
		return nil, err
	}

	return ass, nil
}

// GetAssetboxInfoByAlias gets public assetbox info from blockchain by assetbox's alias
func (m *Manager) GetAssetboxInfoByAlias(ctx context.Context, alias string) (a *Assetbox, err error) {
	m.logger.Debug("contracts manager GetAssetboxInfoByAlias called", "alias", alias)
	defer func() {
		if err != nil {
			err = bitbonErrors.NewContractCallError(err)
		}
	}()

	contract, err := m.getAssetboxImpl()
	if err != nil {
		return nil, err
	}

	opts := &bind.CallOpts{
		Pending: false,
		From:    m.serviceAddress,
		Context: ctx,
	}

	info, err := contract.GetAssetboxInfoByAlias(opts, alias)
	if err != nil {
		return nil, err
	}

	ass, err := ContractAssetbox(info.AssetboxAddress, info.AssetboxAlias, info.ServiceId, info.OtherInfo, info.IsMining)
	if err != nil {
		return nil, err
	}

	return ass, nil
}

func (m *Manager) GetAsseboxesMiningState(ctx context.Context, address common.Address) (bool, error) {
	m.logger.Debug("contracts manager GetAsseboxesMiningState called", "address", address.Hex())

	var err error

	defer func() {
		if err != nil {
			err = bitbonErrors.NewContractCallError(err)
		}
	}()

	contract, err := m.getAssetboxImpl()
	if err != nil {
		return false, err
	}

	opts := &bind.CallOpts{
		Pending: false,
		Context: ctx,
	}

	isMining, err := contract.IsAssetboxMining(opts, address)
	if err != nil {
		return false, err
	}
	return isMining, nil
}

func decodeOtherInfo(info []byte) (*assetboxOtherInfo, error) {
	otherInfo := new(assetboxOtherInfo)
	if len(info) == 0 {
		return otherInfo, nil
	}
	if err := rlp.DecodeBytes(info, otherInfo); err != nil {
		return nil, errors.Wrap(err, "error rlp decoding assetbox other info")
	}
	return otherInfo, nil
}

func ContractAssetbox(address common.Address, alias string, serviceId []byte, otherInfo []byte, isMining bool) (*Assetbox, error) {
	parsedOtherInfo, err := decodeOtherInfo(otherInfo)
	if err != nil {
		return nil, err
	}

	return &Assetbox{
		Address:   address,
		Alias:     alias,
		ServiceID: serviceId,
		IsPublic:  parsedOtherInfo.IsPublic == assetboxPublic,
		ExtraInfo: parsedOtherInfo.ExtraInfo,
		IsMining:  isMining,
	}, nil
}
