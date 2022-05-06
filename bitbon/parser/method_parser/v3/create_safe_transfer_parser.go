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

package v3

import (
	"context"
	"math/big"

	"github.com/pkg/errors"
	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/contracts"
	cs "github.com/simcord-llc/bitbon-system-blockchain/bitbon/contracts/contract_snapshot"
	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/dto"
	bitbonErrors "github.com/simcord-llc/bitbon-system-blockchain/bitbon/errors"
	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/parser/interfaces"
	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/parser/method_parser"
	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/parser/storage_contract_snapshot"
	"github.com/simcord-llc/bitbon-system-blockchain/common"
	"github.com/simcord-llc/bitbon-system-blockchain/core/types"
	"github.com/simcord-llc/bitbon-system-blockchain/log"
)

var (
	ErrDuplicatingAssetboxBalanceEventFound = errors.New("found too much balance events for the same assetbox")
)

type CreateSafeTransferParser struct {
	*method_parser.BaseMethodParser
}

func NewCreateSafeTransferParser(ethApiWrapper interfaces.Contract, contractsSnapshot *storage_contract_snapshot.StorageContractSnapshot) *CreateSafeTransferParser {
	var p = &CreateSafeTransferParser{
		BaseMethodParser: method_parser.NewBaseMethodParser(
			ethApiWrapper,
			contractsSnapshot,
			cs.ContractTypeBitbonImpl,
			dto.CreateSafeTransferV3MethodID,
		),
	}

	return p
}

func (p *CreateSafeTransferParser) Parse(ctx context.Context, tx *types.Transaction) (*dto.BitbonTx, error) {
	res := &dto.BitbonTx{ParseStatus: dto.ParseStatusSuccess}
	bst := &dto.SafeTransferArgs{}

	sender, err := p.Sender(tx)
	if err != nil {
		res.ParseStatus = dto.ParseStatusParseError
		res.ParseError = errors.Wrap(err, "unable to get tx sender")
		return res, res.ParseError
	}

	bst.From = sender
	res.Sender = sender

	methodName, err := p.ParseInput(bst, tx)
	if err != nil {
		res.ParseStatus = dto.ParseStatusParseInputError
		res.ParseError = errors.Wrap(err, "unable to parse input")
		return res, res.ParseError
	}

	res.Args = bst
	res.Method = methodName

	receipt, err := p.EthApiWrapper().TransactionReceipt(ctx, tx.Hash())
	if err != nil {
		res.ParseStatus = dto.ParseStatusParseLogsError
		res.ParseError = errors.Wrap(err, "unable to get transaction receipt to parse logs")
		return res, res.ParseError
	}

	res.Status = receipt.Status == types.ReceiptStatusSuccessful
	if receipt.Status != types.ReceiptStatusSuccessful {
		log.Error("Receipt Status is not successful")
		res.TxError = bitbonErrors.NewUnsuccessfullTxError(contracts.ErrorUnsuccessfulTx)
	}

	addr, err := p.ContractsSnapshot().GetContractAddress(cs.ContractTypeTransferImpl)
	if err != nil {
		res.ParseStatus = dto.ParseStatusParseLogsError
		res.ParseError = errors.Wrap(err, "unable to get TransferImpl address from storage snapshot")
		return res, res.ParseError
	}

	balances, err := p.ParseBalances(addr, receipt.Logs)
	if err != nil {
		res.ParseStatus = dto.ParseStatusParseLogsError
		res.ParseError = errors.Wrap(err, "unable to parse balances from logs")
		return res, res.ParseError
	}
	res.Balances = balances

	feeChargedEvents, err := p.ParseFeeChargedLogs(cs.ContractTypeTransferImpl, addr, receipt.Logs)
	if err != nil {
		res.ParseStatus = dto.ParseStatusParseLogsError
		res.ParseError = errors.Wrap(err, "unable to parse logs")
		return res, res.ParseError
	}

	events := make([]dto.CustomEvent, 0, len(feeChargedEvents))
	for idx := range feeChargedEvents {
		events = append(
			events,
			&feeChargedEvents[idx],
		)
	}
	res.Events = events

	return res, nil
}

func (p *CreateSafeTransferParser) ParseBalances(contractAddress common.Address,
	logs []*types.Log) (map[common.Address]*big.Int, error) {
	lockedBalances := make(map[common.Address]*big.Int, len(logs))
	err := p.ParseBalanceLockedLogs(lockedBalances, cs.ContractTypeTransferImpl, contractAddress, logs)
	if err != nil {
		return nil, errors.Wrap(err, "error parsing BalanceLocked logs")
	}

	changedBalances := make(map[common.Address]*big.Int, len(logs))
	err = p.ParseBalanceChangedLogs(changedBalances, cs.ContractTypeTransferImpl, contractAddress, logs)
	if err != nil {
		return nil, errors.Wrap(err, "error parsing BalanceChanged logs")
	}

	balances := make(map[common.Address]*big.Int, len(lockedBalances)+len(changedBalances))
	for assetbox := range lockedBalances {
		balances[assetbox] = lockedBalances[assetbox]
	}
	for assetbox := range changedBalances {
		if _, ok := balances[assetbox]; ok {
			// in the current implementation of smart-contracts,
			// emmited events don't contain duplicates
			// for the same assetbox (every assetbox that address
			// present in emmited events contain only one event of existing types)
			//
			// so, if other where present, we should raise up an error
			// and think about other implementation of parser
			return nil, ErrDuplicatingAssetboxBalanceEventFound
		}

		balances[assetbox] = changedBalances[assetbox]
	}

	return balances, nil
}
