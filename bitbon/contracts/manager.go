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
package contracts

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"sync"

	"github.com/pkg/errors"
	"github.com/simcord-llc/bitbon-system-blockchain/accounts/abi/bind"
	loggerContext "github.com/simcord-llc/bitbon-system-blockchain/bitbon/context"
	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/contracts/contract_snapshot"
	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/dto"
	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/noncer"
	"github.com/simcord-llc/bitbon-system-blockchain/common"
	"github.com/simcord-llc/bitbon-system-blockchain/core/types"
	"github.com/simcord-llc/bitbon-system-blockchain/crypto"
	"github.com/simcord-llc/bitbon-system-blockchain/eth"
	"github.com/simcord-llc/bitbon-system-blockchain/log"
	"github.com/simcord-llc/bitbon-system-blockchain/rpc"
)

var (
	errNilServicePK   = errors.New("nil servicePK given")
	errAlreadyStopped = errors.New("already stopped")
)

type AddressGetter func() (common.Address, error)
type AddressSetter func(common.Address) error

type Manager struct {
	EthApiWrapper *EthApiWrapper
	get           AddressGetter
	set           AddressSetter

	Noncer *noncer.Noncer

	servicePK      *ecdsa.PrivateKey
	serviceAddress common.Address

	snapshot *contract_snapshot.ContractsSnapshot

	redeployWatchersMux sync.RWMutex             // redeployWatchers
	redeployWatchers    map[rpc.ID]chan struct{} // map of watchers to notify when contracts will be redeployed

	ready chan struct{} // ready channel is closed when Manager is ready to use
	done  chan struct{} // done channel used to notify about quiting

	logger log.Logger
}

var (
	ErrorUnsuccessfulTx          = errors.New("unsuccessful tx receipt status")
	ErrorUnsuccessfulTxEmitEther = errors.New("unsuccessful tx receipt status while ether emission")
	ErrorNotReady                = errors.New("manager is not ready")
)

func NewManager(APIBackend *eth.EthAPIBackend, getter AddressGetter, setter AddressSetter, servicePK *ecdsa.PrivateKey, noncer *noncer.Noncer) (m *Manager, err error) {
	m = &Manager{
		get:              getter,
		set:              setter,
		snapshot:         contract_snapshot.New(common.Address{}, map[contract_snapshot.ContractType]common.Address{}),
		redeployWatchers: make(map[rpc.ID]chan struct{}),
		ready:            make(chan struct{}),
		done:             make(chan struct{}),
		Noncer:           noncer,
		EthApiWrapper:    NewEthApiWrapper(APIBackend),
		logger:           loggerContext.LoggerFromContext(loggerContext.NewLoggerContext(context.Background(), log.New("src", "Manager"))),
	}
	if servicePK == nil {
		err = errNilServicePK
		return
	}
	m.servicePK = servicePK
	m.serviceAddress = crypto.PubkeyToAddress(m.servicePK.PublicKey)

	csa, err := m.get()
	if err != nil {
		err = fmt.Errorf("error getting contract storage address from persistent storage: %v", err)
		return
	}

	m.logger.Warn("Contract storage address obtained from storage", "address", csa.Hex())

	if csa == (common.Address{}) {
		err = m.watchNewContractStorageDeployed(nil) // nolint:staticcheck
		if err != nil {
			err = fmt.Errorf("error watching new contract storage deployed: %v", err)
			return
		}

		m.logger.Warn("Begin watching new contract storage deployed successfully")
	} else {
		err = m.setContractStorageAddress(csa)
		if err != nil {
			err = fmt.Errorf("error setting contract storage address: %v", err)
			return
		}
	}

	err = m.watchContractStorageMoved(nil) // nolint:staticcheck
	if err != nil {
		err = fmt.Errorf("error watching contract storage moved: %v", err)
		return
	}
	m.logger.Warn("Begin watching contract storage moved successfully")

	//Set default value for nonce

	if noncer != nil {
		var nonce uint64
		nonce, err = m.EthApiWrapper.GetNonce(m.serviceAddress)
		if err != nil {
			err = errors.Wrap(err, "unable to get nonce for service assetbox from blockchain")
			return
		}

		err = noncer.SetUpNonce(m.serviceAddress, int64(nonce))
		if err != nil {
			err = errors.Wrap(err, "unable to set nonce for service assetbox to redis")
			return
		}
		m.logger.Warn("Nonce for service account set up to redis successfully")
	}

	return
}

// Stops manager
func (m *Manager) Stop() error {
	if m.Stopped() {
		return errAlreadyStopped
	}
	close(m.done)
	return nil
}

func (m *Manager) Ready() bool {
	select {
	case <-m.ready:
		return true
	default:
		return false
	}
}

func (m *Manager) Stopped() bool {
	select {
	case <-m.done:
		return true
	default:
		return false
	}
}

func (m *Manager) WaitReady(exit chan struct{}) bool {
	select {
	case <-exit:
		return false
	case <-m.done:
		return false
	case <-m.ready:
		return true
	}
}

func (m *Manager) makeReady() {
	log.Warn("Bitbon contract manager became ready")
	close(m.ready)
}

// obtainContractAddresses calls ContractStorageImpl methods to get needed contract addresses
func (m *Manager) obtainContractAddresses(contractStorageAddress common.Address) error {
	err := m.snapshot.SetContractStorageAddress(contractStorageAddress)
	if err != nil {
		return errors.Wrap(err, "unable to set contract storage address")
	}

	for contractType, getterWrapper := range addressGetters {
		addressGetter := getterWrapper(m)

		contractAddress, err := addressGetter(context.Background(), contractStorageAddress)
		if err != nil {
			m.logger.Error("Unable to obtain contract address", "contractType", contractType, "err", err)
			return errors.Wrap(err, "unable to obtain contract address")
		}

		err = m.snapshot.SetContractAddress(contractType, contractAddress)
		if err != nil {
			m.logger.Error("Unable to set contract address to snapshot", "contractType", contractType, "contractAddress", contractAddress, "err", err)
			return errors.Wrap(err, "unable to set contract address to snapshot")
		}

		m.logger.Debug("Successfully set contract address to snapshot", "contractType", contractType, "contractAddress", contractAddress)
	}

	return nil
}

func (m *Manager) getContractAddress(contractType contract_snapshot.ContractType) (common.Address, error) {
	return m.snapshot.GetContractAddress(contractType)
}

func (m *Manager) GetContractAddressAssetboxInfo() (common.Address, error) {
	return m.getContractAddress(contract_snapshot.ContractTypeAssetboxStorageImpl)
}

func (m *Manager) GetContractAddressAssetbox() (common.Address, error) {
	return m.getContractAddress(contract_snapshot.ContractTypeAssetboxImpl)
}

func (m *Manager) GetContractAddressBitbon() (common.Address, error) {
	return m.getContractAddress(contract_snapshot.ContractTypeBitbonImpl)
}

func (m *Manager) GetContractAddressTransfer() (common.Address, error) {
	return m.getContractAddress(contract_snapshot.ContractTypeTransferImpl)
}

func (m *Manager) GetContractAddressSafeTransferStorage() (common.Address, error) {
	return m.getContractAddress(contract_snapshot.ContractTypeSafeTransferStorageImpl)
}

func (m *Manager) GetContractAddressBitbonSupport() (common.Address, error) {
	return m.getContractAddress(contract_snapshot.ContractTypeBitbonSupportImpl)
}

func (m *Manager) GetContractAddressBitbonStorage() (common.Address, error) {
	return m.getContractAddress(contract_snapshot.ContractTypeBitbonStorageImpl)
}

func (m *Manager) GetContractAddressAccessStorage() (common.Address, error) {
	return m.getContractAddress(contract_snapshot.ContractTypeAccessStorageImpl)
}

func (m *Manager) GetContractAddressDistributionStorage() (common.Address, error) {
	return m.getContractAddress(contract_snapshot.ContractTypeDistributionStorageImpl)
}

func (m *Manager) GetContractAddressMiningAgentStorage() (common.Address, error) {
	return m.getContractAddress(contract_snapshot.ContractTypeMiningAgentStorageImpl)
}

// GetContractStorageAddress gets ContractStorageAddress from manager field
func (m *Manager) GetContractStorageAddress() (common.Address, error) {
	return m.snapshot.GetContractStorageAddress()
}

func (m *Manager) GetContractAddresses() map[contract_snapshot.ContractType]common.Address {
	return m.snapshot.GetContractAddresses()
}

func (m *Manager) GetCurrentContractAbiInfo(version dto.ContractVersion) ([]*dto.AbiInfo, error) {
	if !dto.ContractVersionExists(version) {
		return nil, errors.Errorf("unknown contract version %s", version)
	}

	addresses := m.snapshot.GetContractAddresses()
	res := make([]*dto.AbiInfo, 0, len(addresses))

	for ct, ca := range addresses {
		abi, err := ContractTypeAbi(ct)
		if err != nil {
			return nil, errors.Wrap(err, "error getting abi json")
		}

		res = append(res, &dto.AbiInfo{
			Address:      ca,
			AbiJSON:      abi,
			Version:      version,
			ContractType: ct,
		})
	}

	return res, nil
}

// setContractStorageAddress sets ContractStorageAddress to manager field. if address has changed - obtain the rest of contract addresses
func (m *Manager) setContractStorageAddress(addr common.Address) error {
	if addr == (common.Address{}) {
		return errors.New("trying to set empty ContractStorageAddress")
	}

	// makeReady is the flag that indicates that contract storage deployed (not moved)
	prevContractStorage, err := m.snapshot.GetContractStorageAddress()
	if err != nil {
		return errors.Wrap(err, "error getting contract storage address")
	}
	makeReady := prevContractStorage == (common.Address{})

	// setting all methods results  with possible error to tmp variables to prevent inconsistent state of manager fields
	err = m.obtainContractAddresses(addr)
	if err != nil {
		return errors.Wrap(err, "error obtaining contract addresses")
	}

	if makeReady {
		m.makeReady()
	} else {
		// notify watchers that contract storage moved
		m.notifyRedeployWatchers()
	}

	m.logger.Info("Bitbon contract manager successfully set contract address", "contract address", addr.Hex())

	return nil
}

// SaveContractStorageAddress calls setContractStorageAddress before saving to persistent storage
func (m *Manager) SaveContractStorageAddress(addr common.Address) error {
	if addr == (common.Address{}) {
		return errors.New("trying to save empty ContractStorageAddress")
	}

	if err := m.setContractStorageAddress(addr); err != nil {
		return errors.Wrap(err, "setting contract storage address to manager error")
	}

	if err := m.set(addr); err != nil {
		return errors.Wrap(err, "setting contract storage address to persistent storage error")
	}

	return nil
}

// suggestGasLimit suggests ether emission
func (m *Manager) suggestEtherEmission() *big.Int {
	return new(big.Int).Mul(new(big.Int).SetUint64(m.EthApiWrapper.SuggestGasLimit()), big.NewInt(defaultGasPrice))
}

// waitAndCheckTransaction waits for transaction will be mined and checks it`s receipt for success
func (m *Manager) waitAndCheckTransaction(ctx context.Context, tx *types.Transaction) (*types.Receipt, error) {
	m.logger.Debug("Waiting for tx receipt (contract call)...", "txHash", tx.Hash().Hex())

	ctx, cancel := context.WithTimeout(ctx, waitForTxReceiptTimeout)
	defer cancel()

	receipt, err := bind.WaitMined(ctx, m.EthApiWrapper, tx)
	if err != nil {
		return nil, err
	}
	if receipt.Status != types.ReceiptStatusSuccessful {
		return receipt, ErrorUnsuccessfulTx
	}

	m.logger.Debug("Smart contract called, receipt obtained", "txHash", tx.Hash().Hex(), "gasUsed", receipt.GasUsed, "success", receipt.Status)
	return receipt, nil
}

func (m *Manager) getReceiptParams(receipt *types.Receipt) (blockNum uint64, txHash common.Hash, err error) {
	txHash = receipt.TxHash
	_, _, blockNum = m.EthApiWrapper.Transaction(context.Background(), txHash)
	if blockNum != 0 {
		return
	}

	if len(receipt.Logs) == 0 {
		err = errors.New("can`t getReceiptParams - empty logs")
		return
	}
	blockNum = receipt.Logs[0].BlockNumber
	if blockNum == 0 {
		err = errors.New("can`t getReceiptParams - empty block number")
		return
	}
	return
}

// prepareTransactOpts creates TransactOpts based on PrivateKey and emit ether on that PrivateKey if needed
func (m *Manager) prepareTransactOpts(ctx context.Context, key *ecdsa.PrivateKey) (opts *bind.TransactOpts, err error) {
	// call payed smart contract method, signing transaction with assetbox private key
	opts = bind.NewKeyedTransactor(key)
	opts.GasLimit = m.EthApiWrapper.SuggestGasLimit()
	opts.GasPrice = big.NewInt(defaultGasPrice)
	opts.Context = ctx

	defer func() {
		if err == nil && opts.From == m.serviceAddress {
			if m.Noncer != nil {
				nonce, err := m.Noncer.IncrementAndGetNonce(m.serviceAddress) // nolint:govet

				if err == nil {
					opts.Nonce = new(big.Int).SetInt64(nonce)
				} else {
					m.logger.Error("Error getting noncer nonce", "error", err)
				}
			} else {
				m.logger.Debug("Noncer is disabled. Continue with default behaviour")
			}
		}
	}()

	// emit ether once to allow user to perform operationsPerEmitEther operations
	balance, err := m.EthApiWrapper.GetBalance(ctx, opts.From)
	if err != nil {
		err = errors.Wrap(err, "error getting assetbox ether balance")
		return
	}
	bigIntBalance := balance.ToInt()
	// balance that we want to have to perform one operation
	wantedBalance := new(big.Int).Mul(new(big.Int).SetUint64(opts.GasLimit), opts.GasPrice)

	// check if there is enough ether on assetbox to perform operation
	// if not - emit with operationsPerEmitEther
	if bigIntBalance.Cmp(wantedBalance) >= 0 {
		m.logger.Trace("no need to emit ether to assetbox", "assetbox", opts.From.Hex(), "balance", bigIntBalance.String())
		return
	}

	err = m.EmitEther(ctx, opts.From, new(big.Int).Mul(new(big.Int).SetUint64(operationsPerEmitEther), wantedBalance), nil)
	if err != nil {
		err = errors.Wrap(err, "error emitting assetbox with ether")
	}

	return
}

type (
	Signal      <-chan struct{}
	Unsubscribe func()
)

// SubscribeRedeploy subscribes on redeploy contracts
// Returns:
// - signal channel which will be closed when contracts are redeployed
// - unsubscribe function, that cancels subscription
func (m *Manager) SubscribeRedeploy() (Signal, Unsubscribe) {
	id := rpc.NewID()
	signal := make(chan struct{})
	unsubscribe := func() {
		m.redeployWatchersMux.RLock()
		defer m.redeployWatchersMux.RUnlock()
		delete(m.redeployWatchers, id)
	}

	m.redeployWatchersMux.Lock()
	defer m.redeployWatchersMux.Unlock()
	m.redeployWatchers[id] = signal

	return signal, unsubscribe
}

// notifyRedeployWatchers notifies all redeploy watchers that contracts are redeployed
func (m *Manager) notifyRedeployWatchers() {
	m.redeployWatchersMux.Lock()
	defer m.redeployWatchersMux.Unlock()
	for id, signal := range m.redeployWatchers {
		close(signal)
		delete(m.redeployWatchers, id)
	}
}

func (m *Manager) GetEthAPIWrapper() *EthApiWrapper {
	return m.EthApiWrapper
}

func (m *Manager) GetNoncer() *noncer.Noncer {
	return m.Noncer
}
