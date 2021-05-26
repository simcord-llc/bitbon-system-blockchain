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

package mining_agent

import (
	"context"

	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/dto"

	"github.com/pkg/errors"
	bb "github.com/simcord-llc/bitbon-system-blockchain/bitbon"
	loggerContext "github.com/simcord-llc/bitbon-system-blockchain/bitbon/context"
	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/models"
	"github.com/simcord-llc/bitbon-system-blockchain/common"
	"github.com/simcord-llc/bitbon-system-blockchain/log"
)

type Manager struct {
	bitbon    *bb.Bitbon
	encryptor bb.PkEncryptor
	logger    log.Logger
}

// MiningAgentManager create a new MiningAgentManager instance.
func NewMiningAgentManager(b *bb.Bitbon, encryptor bb.PkEncryptor) *Manager {
	return &Manager{
		bitbon:    b,
		encryptor: encryptor,
		logger: loggerContext.LoggerFromContext(loggerContext.NewLoggerContext(context.Background(),
			log.New("src", "MiningAgentManager"))),
	}
}

func (mam *Manager) ProposeDistribution(ctx context.Context,
	req dto.ProposeDistributionRequest) (err error) {
	mam.logger.Debug("MiningAgentManager ProposeDistribution called",
		"address", req.Address.Hex(), "distribution", req.Distribution, "isPublic")

	assetbox := &models.Assetbox{
		Wallet:     req.CryptoData.Wallet,
		PassPhrase: req.CryptoData.Passphrase,
	}

	if err := bb.DecryptAssetboxWallet(assetbox, mam.bitbon.GetDecryptAssetboxWalletPassword(), mam.encryptor.Decrypt); err != nil {
		return err
	}

	err = mam.bitbon.GetContractsManager().ProposeDistribution(ctx, req.Distribution, assetbox.Pk)
	if err != nil {
		return errors.Wrap(err, "error adding distribution proposal to blockchain")
	}

	return nil
}

func (mam *Manager) GetCurrentDistribution(ctx context.Context) (distribution map[string]uint64, err error) {
	mam.logger.Debug("MiningAgentManager GetCurrentDistribution called")

	distribution, err = mam.bitbon.GetContractsManager().GetCurrentDistribution(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "error getting current distribution from blockchain")
	}

	return distribution, nil
}

func (mam *Manager) GetAllMiningAgents(ctx context.Context) (res []common.Address, err error) {
	mam.logger.Debug("MiningAgentManager GetAllMiningAgents called")
	res, err = mam.bitbon.GetContractsManager().GetAllMiningAgents(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "error getting all mining agents from blockchain")
	}

	return res, nil
}

func (mam *Manager) IsMiningAgent(ctx context.Context, address common.Address) (res bool, err error) {
	mam.logger.Debug("MiningAgentManager IsMiningAgent called")

	res, err = mam.bitbon.GetContractsManager().IsMiningAgent(ctx, address)
	if err != nil {
		return false, errors.Wrap(err, "error checking mining agents in blockchain")
	}

	return res, nil
}

func (mam *Manager) AddMiningAgent(ctx context.Context, address common.Address) (err error) {
	mam.logger.Debug("MiningAgentManager AddMiningAgent called")

	err = mam.bitbon.GetContractsManager().AddMiningAgent(ctx, address, mam.bitbon.GetServicePk())
	if err != nil {
		return errors.Wrap(err, "error adding mining agents to blockchain")
	}

	return nil
}

func (mam *Manager) RemoveMiningAgent(ctx context.Context, address common.Address) (err error) {
	mam.logger.Debug("MiningAgentManager RemoveMiningAgent called")

	err = mam.bitbon.GetContractsManager().RemoveMiningAgent(ctx, address, mam.bitbon.GetServicePk())
	if err != nil {
		return errors.Wrap(err, "error removing mining agents from blockchain")
	}

	return nil
}
