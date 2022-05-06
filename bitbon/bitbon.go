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
package bitbon

import (
	"crypto/ecdsa"
	"fmt"
	"sync"

	"github.com/pkg/errors"
	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/contracts"
	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/parser"
	"github.com/simcord-llc/bitbon-system-blockchain/common"
	"github.com/simcord-llc/bitbon-system-blockchain/crypto"
	"github.com/simcord-llc/bitbon-system-blockchain/eth"
	"github.com/simcord-llc/bitbon-system-blockchain/log"
	"github.com/simcord-llc/bitbon-system-blockchain/node"
	"github.com/simcord-llc/bitbon-system-blockchain/p2p"
	"github.com/simcord-llc/bitbon-system-blockchain/rpc"
)

type Bitbon struct {
	config  *Config
	storage *Storage

	servicePK      *ecdsa.PrivateKey
	serviceAddress common.Address

	apiWG sync.WaitGroup
	api   *API

	contractsManager   ContractManager
	assetboxManager    AssetboxManager
	transferManager    TransferManager
	miningAgentManager MiningAgent
	feeManager         FeeManager

	parser *parser.Parser
	eth    *eth.Ethereum

	stopAPI chan struct{}
	done    chan struct{}
}

var ErrAPIStopped = errors.New("bitbon API has stopped")

// New creates a new Bitbon object (including the
// initialisation of the common Bitbon object)
func New(ctx *node.ServiceContext, config *Config) (b *Bitbon, err error) {
	var ethServ *eth.Ethereum

	err = ctx.Service(&ethServ)
	if err != nil {
		return b, errors.Wrap(err, "error getting ethereum service from context")
	}

	b = InitDefaultBitbonService(config, ethServ)

	err = InitServicePart(b, config)
	if err != nil {
		return
	}

	err = InitStorages(ctx, b)
	if err != nil {
		return
	}

	// Initialize api interfaces
	b.api = NewAPI(b)

	log.Info("Bitbon service created")

	return
}

func InitDefaultBitbonService(cfg *Config, ethServ *eth.Ethereum) *Bitbon {
	return &Bitbon{
		config:  cfg,
		eth:     ethServ,
		stopAPI: make(chan struct{}),
		done:    make(chan struct{}),
	}
}

func (b *Bitbon) InitBitbonParser(apiWrapper *contracts.EthApiWrapper, storage *Storage) { // nolint
	b.parser = parser.NewParser(apiWrapper, storage)
}

func InitServicePart(b *Bitbon, config *Config) (err error) {
	// Service account
	b.servicePK, err = crypto.HexToECDSA(config.ServicePrivateKey)
	if err != nil {
		err = fmt.Errorf("error while parsing service PK: %v", err)
		return err
	}
	b.serviceAddress = crypto.PubkeyToAddress(b.servicePK.PublicKey)
	log.Debug("Service account inited", "address", b.serviceAddress.Hex())

	return nil
}

func InitStorages(ctx *node.ServiceContext, b *Bitbon) (err error) {
	// Storage
	storagePath := ctx.ResolvePath(PrivateStorageFileName)
	if b.storage, err = NewStorage(storagePath); err != nil {
		err = fmt.Errorf("error while creating Bitbon private storage: %v", err)
		return
	}
	defer func() {
		if err != nil {
			b.storage.Close()
		}
	}()

	return
}

func (b *Bitbon) InjectContractManager(cm ContractManager) {
	// make it settable only once
	if b.contractsManager == nil {
		b.contractsManager = cm
	}
}

func (b *Bitbon) InjectAssetboxManager(am AssetboxManager) {
	// make it settable only once
	if b.assetboxManager == nil {
		b.assetboxManager = am
	}
}

func (b *Bitbon) InjectTransferManager(tm TransferManager) {
	// make it settable only once
	if b.transferManager == nil {
		b.transferManager = tm
	}
}

func (b *Bitbon) InjectMiningAgent(ma MiningAgent) {
	// make it settable only once
	if b.miningAgentManager == nil {
		b.miningAgentManager = ma
	}
}

func (b *Bitbon) InjectFeeManager(fm FeeManager) {
	// make it settable only once
	if b.feeManager == nil {
		b.feeManager = fm
	}
}

func (b *Bitbon) GetServiceAddress() common.Address {
	return b.serviceAddress
}

func (b *Bitbon) GetServicePk() *ecdsa.PrivateKey {
	return b.servicePK
}

func (b *Bitbon) GetServiceID() string {
	return b.config.ServiceID
}

func (b *Bitbon) GetDecryptAssetboxWalletPassword() string {
	return b.config.DecryptAssetboxWalletPassword
}

func (b *Bitbon) SetDecryptAssetboxWalletPassword(value string) {
	if b.config == nil {
		b.config = &Config{}
	}
	b.config.DecryptAssetboxWalletPassword = value
}

// APIs is a meaningless implementation of node.Service.
func (b *Bitbon) APIs() []rpc.API {
	return []rpc.API{
		{
			Namespace: ProtocolName,
			Version:   ProtocolVersionStr,
			Service:   b.GetAPI(),
			Public:    true,
		},
		{
			Namespace: ProtocolName,
			Version:   ProtocolVersionStr,
			Service:   NewAPINonce(b),
			Public:    true,
		},
	}
}

// Protocols is a meaningless implementation of node.Service.
func (b *Bitbon) Protocols() []p2p.Protocol {
	return nil
}

// Start implements node.Service
func (b *Bitbon) Start(_ *p2p.Server) (err error) {
	log.Info("Bitbon service started")
	return nil
}

// Stop implements node.Service
func (b *Bitbon) Stop() error {
	if b.Stopped() {
		return errors.New("already stopped")
	}

	if !b.APIStopped() {
		close(b.stopAPI)
	}

	// wait for all api methods
	log.Warn("Waiting for all api methods to stop...")
	b.apiWG.Wait()
	log.Info("All api methods are stopped")

	b.stopAllService()

	return nil
}

func (b *Bitbon) stopAllService() {
	if b.contractsManager != nil {
		if err := b.contractsManager.Stop(); err != nil {
			log.Error("Failed to close Bitbon contracts manager", "err", err)
		}
		log.Info("Bitbon contracts manager stopped")
	}

	if b.transferManager != nil {
		if err := b.transferManager.Stop(); err != nil {
			log.Error("Failed to close Bitbon transfer manager", "err", err)
		}
		log.Info("Bitbon transfer manager stopped")
	}

	if b.storage != nil {
		if err := b.storage.Close(); err != nil {
			log.Error("Failed to close Bitbon private storage", "err", err)
		}
		log.Info("Bitbon private storage closed")
	}

	close(b.done)
}

func (b *Bitbon) Stopped() bool {
	select {
	case <-b.done:
		return true
	default:
		return false
	}
}

func (b *Bitbon) APIStopped() bool {
	select {
	case <-b.stopAPI:
		return true
	default:
		return false
	}
}

func (b *Bitbon) GetAPI() *API {
	return b.api
}

func (b *Bitbon) GetEth() *eth.Ethereum {
	return b.eth
}

func (b *Bitbon) GetStorage() *Storage {
	return b.storage
}

func (b *Bitbon) GetContractsManager() ContractManager {
	return b.contractsManager
}

func (b *Bitbon) GetAssetboxManager() AssetboxManager {
	return b.assetboxManager
}

func (b *Bitbon) GetTransferManager() TransferManager {
	return b.transferManager
}

func (b *Bitbon) GetFeeManager() FeeManager {
	return b.feeManager
}

func (b *Bitbon) GetParser() *parser.Parser {
	return b.parser
}
