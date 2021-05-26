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
	"errors"
	"math/big"

	"github.com/simcord-llc/bitbon-system-blockchain/log"

	"github.com/simcord-llc/bitbon-system-blockchain/core"

	"crypto/ecdsa"

	ethereum "github.com/simcord-llc/bitbon-system-blockchain"
	"github.com/simcord-llc/bitbon-system-blockchain/accounts/abi/bind"
	loggerContext "github.com/simcord-llc/bitbon-system-blockchain/bitbon/context"
	"github.com/simcord-llc/bitbon-system-blockchain/common"
	"github.com/simcord-llc/bitbon-system-blockchain/common/hexutil"
	"github.com/simcord-llc/bitbon-system-blockchain/core/rawdb"
	"github.com/simcord-llc/bitbon-system-blockchain/core/types"
	"github.com/simcord-llc/bitbon-system-blockchain/crypto"
	"github.com/simcord-llc/bitbon-system-blockchain/eth"
	"github.com/simcord-llc/bitbon-system-blockchain/eth/filters"
	"github.com/simcord-llc/bitbon-system-blockchain/event"
	"github.com/simcord-llc/bitbon-system-blockchain/internal/ethapi"
	"github.com/simcord-llc/bitbon-system-blockchain/rlp"
	"github.com/simcord-llc/bitbon-system-blockchain/rpc"
)

// to ensure that EthApiWrapper satisfies bind.ContractBackend interface. If not there will be compile time error
var _ = bind.ContractBackend(&EthApiWrapper{})

// Args represents the arguments to sumbit a new transaction into the transaction pool
type Args struct {
	From     common.Address
	FromPK   *ecdsa.PrivateKey
	To       *common.Address
	Gas      uint64
	GasPrice *big.Int
	Value    *big.Int
	Nonce    *uint64
	Input    []byte
}

func (args *Args) checkSender() (err error) {
	from := crypto.PubkeyToAddress(args.FromPK.PublicKey)
	if args.From == (common.Address{}) {
		args.From = from
	} else if args.From != from {
		err = errors.New("tx From vs FromPk mismatch")
	}
	return err
}

// setDefaults is a helper function that fills in default values for unspecified tx fields.
func (args *Args) setDefaults(ctx context.Context, e *EthApiWrapper) error {
	if args.Gas == 0 {
		args.Gas = e.SuggestGasLimit()
	}
	if args.GasPrice == nil {
		args.GasPrice = big.NewInt(defaultGasPrice)
	}
	if args.Value == nil {
		args.Value = new(big.Int)
	}
	if args.Nonce == nil {
		nonce, err := e.backend.GetPoolNonce(ctx, args.From)
		if err != nil {
			return err
		}
		args.Nonce = &nonce
	}
	if args.To == nil {
		// Contract creation
		if len(args.Input) == 0 {
			return errors.New(`contract creation without any data provided`)
		}
	}
	return nil
}

func (args *Args) toTransaction() *types.Transaction {
	if args.To == nil {
		return types.NewContractCreation(*args.Nonce, args.Value, args.Gas, args.GasPrice, args.Input)
	}
	return types.NewTransaction(*args.Nonce, *args.To, args.Value, args.Gas, args.GasPrice, args.Input)
}

// PublicTransactionPoolAPI exposes methods for the RPC interface
type EthApiWrapper struct {
	backend                  *eth.EthAPIBackend
	nonceLock                *ethapi.AddrLocker
	publicBlockChainAPI      *ethapi.PublicBlockChainAPI
	publicTransactionPoolAPI *ethapi.PublicTransactionPoolAPI
	events                   *filters.EventSystem

	logger log.Logger
}

// NewPublicTransactionPoolAPI creates a new RPC service with methods specific for the transaction pool.
func NewEthApiWrapper(b *eth.EthAPIBackend) *EthApiWrapper {
	nonceLock := new(ethapi.AddrLocker)
	return &EthApiWrapper{
		backend:                  b,
		nonceLock:                nonceLock,
		publicBlockChainAPI:      ethapi.NewPublicBlockChainAPI(b),
		publicTransactionPoolAPI: ethapi.NewPublicTransactionPoolAPI(b, nonceLock),
		events:                   filters.NewEventSystem(b, false),

		logger: loggerContext.LoggerFromContext(loggerContext.NewLoggerContext(context.Background(), log.New("src", "EthApiWrapper"))),
	}
}

// SendBitbonTransaction is method for sending transaction, using Args.
// The main difference from SendTransaction (from publicTransactionPoolAPI) is that we use sender PrivateKey
// to sign the transaction without calling AccountManager methods.
// Also we populate From field using sender PrivateKey if it is not set.
func (e *EthApiWrapper) SendBitbonTransaction(ctx context.Context, args Args) (*types.Transaction, error) {
	// check from and terminate on failure
	if err := args.checkSender(); err != nil {
		return nil, err
	}

	if args.Nonce == nil {
		// Hold the addresse'e mutex around signing to prevent concurrent assignment of
		// the same nonce to multiple accounts.
		e.nonceLock.LockAddr(args.From)
		defer e.nonceLock.UnlockAddr(args.From)
	}

	// Set some sanity defaults and terminate on failure
	if err := args.setDefaults(ctx, e); err != nil {
		return nil, err
	}
	e.logger.Debug("Tx args are ready", "from", args.From, "to", args.To, "value", args.Value, "gas", args.Gas, "gasPrice", args.GasPrice, "nonce", *args.Nonce)
	// Assemble the transaction and sign with the wallet
	tx := args.toTransaction()

	signer := types.MakeSigner(e.backend.ChainConfig(), e.backend.CurrentBlock().Number())

	tx, err := types.SignTx(tx, signer, args.FromPK)
	if err != nil {
		return nil, err
	}

	if err := e.backend.SendTx(ctx, tx); err != nil {
		return nil, err
	}
	if tx.To() == nil {
		addr := crypto.CreateAddress(args.From, tx.Nonce())
		e.logger.Debug("Submitted contract creation", "fullhash", tx.Hash().Hex(), "contract", addr.Hex())
	} else {
		e.logger.Debug("Submitted transaction", "fullhash", tx.Hash().Hex(), "recipient", tx.To())
	}
	return tx, nil
}

// converts *big.Int to rpc.BlockNumber
func toBlockNumArg(number *big.Int) rpc.BlockNumber {
	if number == nil {
		return rpc.LatestBlockNumber
	}
	return rpc.BlockNumber(number.Uint64())
}

func toCallArg(msg ethereum.CallMsg) ethapi.CallArgs {
	args := ethapi.CallArgs{
		From: &msg.From,
		To:   msg.To,
	}

	if len(msg.Data) > 0 {
		args.Data = (*hexutil.Bytes)(&msg.Data)
	}
	if msg.Value != nil {
		args.Value = (*hexutil.Big)(msg.Value)
	}
	if msg.Gas != 0 {
		args.Gas = (*hexutil.Uint64)(&msg.Gas)
	}
	if msg.GasPrice != nil {
		args.GasPrice = (*hexutil.Big)(msg.GasPrice)
	}
	return args
}

// CodeAt returns the code of the given account. This is needed to differentiate
// between contract internal errors and the local chain being out of sync. (ContractCaller implementation)
func (e *EthApiWrapper) CodeAt(ctx context.Context, contract common.Address, blockNumber *big.Int) ([]byte, error) {
	return e.publicBlockChainAPI.GetCode(ctx, contract, rpc.BlockNumberOrHashWithNumber(toBlockNumArg(blockNumber)))
}

// GetNonce returns the amount transaction by account.
func (e *EthApiWrapper) GetNonce(acc common.Address) (uint64, error) {
	return e.backend.GetPoolNonce(context.Background(), acc)
}

// ContractCall executes an Ethereum contract call with the specified data as the input. (ContractCaller implementation)
func (e *EthApiWrapper) CallContract(ctx context.Context, call ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	return e.publicBlockChainAPI.Call(ctx, toCallArg(call), rpc.BlockNumberOrHashWithNumber(toBlockNumArg(blockNumber)), nil)
}

// PendingCodeAt returns the code of the given account in the pending state. (PendingContractCaller and ContractTransactor implementation)
func (e *EthApiWrapper) PendingCodeAt(ctx context.Context, contract common.Address) ([]byte, error) {
	return e.publicBlockChainAPI.GetCode(ctx, contract, rpc.BlockNumberOrHashWithNumber(rpc.PendingBlockNumber))
}

// PendingCallContract executes an Ethereum contract call against the pending state. (PendingContractCaller implementation)
func (e *EthApiWrapper) PendingCallContract(ctx context.Context, call ethereum.CallMsg) ([]byte, error) {
	return e.publicBlockChainAPI.Call(ctx, toCallArg(call), rpc.BlockNumberOrHashWithNumber(rpc.PendingBlockNumber), nil)
}

// GetBalance returns the amount of wei for the given address in the state of the
// given block number. The rpc.LatestBlockNumber and rpc.PendingBlockNumber meta
// block numbers are also allowed.
func (e *EthApiWrapper) GetBalance(ctx context.Context, address common.Address) (*hexutil.Big, error) {
	return e.publicBlockChainAPI.GetBalance(ctx, address, rpc.BlockNumberOrHashWithNumber(rpc.PendingBlockNumber))
}

// PendingNonceAt returns the account nonce of the given account in the pending state.
// This is the nonce that should be used for the next transaction. (ContractTransactor implementation)
func (e *EthApiWrapper) PendingNonceAt(ctx context.Context, account common.Address) (uint64, error) {
	result, err := e.publicTransactionPoolAPI.GetTransactionCount(ctx, account, rpc.BlockNumberOrHashWithNumber(rpc.PendingBlockNumber))
	if result == nil {
		return 0, err
	}
	return uint64(*result), err
}

// SuggestGasPrice retrieves the currently suggested gas price to allow a timely
// execution of a transaction. (ContractTransactor implementation)
func (e *EthApiWrapper) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	return e.backend.SuggestPrice(ctx)
}

// suggestGasLimit suggests GasLimit
func (e *EthApiWrapper) SuggestGasLimit() uint64 {
	return defaultGasLimit
	//return e.backend.CurrentBlock().GasLimit()
}

// EstimateGas tries to estimate the gas needed to execute a specific transaction based on
// the current pending state of the backend blockchain. There is no guarantee that this is
// the true gas limit requirement as other transactions may be added or removed by miners,
// but it should provide a basis for setting a reasonable default. (ContractTransactor implementation)
func (e *EthApiWrapper) EstimateGas(ctx context.Context, msg ethereum.CallMsg) (uint64, error) {
	result, err := e.publicBlockChainAPI.EstimateGas(ctx, toCallArg(msg))
	return uint64(result), err
}

// SendTransaction injects a signed transaction into the pending pool for execution.
//
// If the transaction was a contract creation use the TransactionReceipt method to get the
// contract address after the transaction has been mined. (ContractTransactor implementation)
func (e *EthApiWrapper) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	data, err := rlp.EncodeToBytes(tx)
	if err != nil {
		return err
	}
	_, err = e.publicTransactionPoolAPI.SendRawTransaction(ctx, hexutil.Bytes(data))
	return err
}

// TransactionReceipt returns the receipt of a transaction by transaction hash.
// Note that the receipt is not available for pending transactions.
func (e *EthApiWrapper) transactionReceipt(ctx context.Context, hash common.Hash) (*types.Receipt, error) {
	tx, blockHash, _, index := rawdb.ReadTransaction(e.backend.ChainDb(), hash)
	if tx == nil {
		return nil, nil
	}
	receipts, err := e.backend.GetReceipts(ctx, blockHash)
	if err != nil {
		return nil, err
	}
	if len(receipts) <= int(index) {
		return nil, nil
	}
	return receipts[index], nil
}

// Transaction returns the transaction by transaction hash.
// Note that the data is not available for pending transactions.
func (e *EthApiWrapper) Transaction(ctx context.Context, hash common.Hash) (*types.Transaction, common.Hash, uint64) {
	tx, blockHash, blockNumber, _ := rawdb.ReadTransaction(e.backend.ChainDb(), hash)
	if tx == nil {
		return nil, common.Hash{}, 0
	}
	return tx, blockHash, blockNumber
}

// TransactionReceipt returns the receipt of a transaction by transaction hash.
// Note that the receipt is not available for pending transactions.
func (e *EthApiWrapper) TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	//r, err := e.publicTransactionPoolAPI.GetTransactionReceipt(ctx, txHash)
	r, err := e.transactionReceipt(ctx, txHash)
	if err == nil {
		if r == nil {
			return nil, ethereum.NotFound
		}
	}
	return r, err
}

// FilterLogs executes a log filter operation, blocking during execution and
// returning all the results in one batch.
func (e *EthApiWrapper) FilterLogs(ctx context.Context, query ethereum.FilterQuery) ([]types.Log, error) {
	var filter *filters.Filter
	if len(query.Logs) > 0 {
		// Client logs filter requested, construct a logst filter
		filter = filters.NewLogsFilter(e.backend, query.Logs, query.Addresses, query.Topics)
	} else if query.BlockHash != nil {
		// ClientBlock filter requested, construct a single-shot filter
		filter = filters.NewBlockFilter(e.backend, *query.BlockHash, query.Addresses, query.Topics)
	} else {
		// Initialize unset filter boundaried to run from genesis to chain head
		from := int64(0)
		if query.FromBlock != nil {
			from = query.FromBlock.Int64()
		}
		to := int64(-1)
		if query.ToBlock != nil {
			to = query.ToBlock.Int64()
		}
		// Construct the range filter
		filter = filters.NewRangeFilter(e.backend, from, to, query.Addresses, query.Topics)
	}
	// Run the filter and return all the logs
	logs, err := filter.Logs(ctx)
	if err != nil {
		return nil, err
	}
	res := make([]types.Log, len(logs))
	for i, l := range logs {
		res[i] = *l
	}
	return res, nil
}

func emptyAddresses(addrs []common.Address) bool {
	for _, addr := range addrs {
		if addr != (common.Address{}) {
			return false
		}
	}

	return true
}

func (e *EthApiWrapper) SubscribeChainEvent(ctx context.Context, ch chan<- core.ChainEvent) event.Subscription {
	return e.backend.SubscribeChainEvent(ch)
}

// SubscribeFilterLogs creates a background log filtering operation, returning a
// subscription immediately, which can be used to stream the found events.
func (e *EthApiWrapper) SubscribeFilterLogs(ctx context.Context, query ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error) {
	// Subscribe to contract events
	sink := make(chan []*types.Log)

	if emptyAddresses(query.Addresses) {
		query.Addresses = nil
	}

	sub, err := e.events.SubscribeLogs(query, sink)
	if err != nil {
		return nil, err
	}
	// Since we're getting logs in batches, we need to flatten them into a plain stream
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()

		for {
			select {
			case logs := <-sink:
				for _, l := range logs {
					select {
					case ch <- *l:
					case err := <-sub.Err():
						return err
					case <-quit:
						return nil
					}
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}
