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

type QuickTransferParser struct {
	*method_parser.BaseMethodParser
}

func NewQuickTransferParser(ethApiWrapper interfaces.Contract, contractsSnapshot *storage_contract_snapshot.StorageContractSnapshot) *QuickTransferParser {
	var p = &QuickTransferParser{
		BaseMethodParser: method_parser.NewBaseMethodParser(
			ethApiWrapper,
			contractsSnapshot,
			cs.ContractTypeBitbonImpl,
			dto.QuickTransferV3MethodID,
		),
	}

	return p
}

func (p *QuickTransferParser) Parse(ctx context.Context, tx *types.Transaction) (*dto.BitbonTx, error) {
	res := &dto.BitbonTx{ParseStatus: dto.ParseStatusSuccess}
	bst := &dto.QuickTransferArgs{}

	sender, err := p.Sender(tx)
	if err != nil {
		res.ParseStatus = dto.ParseStatusParseError
		err = errors.Wrap(err, "unable to get tx sender")
		res.ParseError = err
		return res, err
	}

	bst.From = sender
	res.Sender = sender

	methodName, err := p.ParseInput(bst, tx)
	if err != nil {
		res.ParseStatus = dto.ParseStatusParseInputError
		err = errors.Wrap(err, "unable to parse input")
		res.ParseError = err
		return res, err
	}

	res.Args = bst
	res.Method = methodName

	receipt, err := p.EthApiWrapper().TransactionReceipt(ctx, tx.Hash())
	if err != nil {
		res.ParseStatus = dto.ParseStatusParseLogsError
		err = errors.Wrap(err, "unable to get transaction receipt to parse logs")
		res.ParseError = err
		return res, err
	}

	res.Status = receipt.Status == types.ReceiptStatusSuccessful
	if receipt.Status != types.ReceiptStatusSuccessful {
		log.Error("Receipt Status is not successful")
		res.TxError = bitbonErrors.NewUnsuccessfullTxError(contracts.ErrorUnsuccessfulTx)
	}

	addr, err := p.ContractsSnapshot().GetContractAddress(cs.ContractTypeTransferImpl)
	if err != nil {
		res.ParseStatus = dto.ParseStatusParseLogsError
		err = errors.Wrap(err, "unable to get TransferImpl address from storage snapshot")
		res.ParseError = err
		return res, err
	}

	balances := make(map[common.Address]*big.Int)
	err = p.ParseBalanceChangedLogs(balances, cs.ContractTypeTransferImpl, addr, receipt.Logs)
	if err != nil {
		res.ParseStatus = dto.ParseStatusParseLogsError
		err = errors.Wrap(err, "unable to parse logs")
		res.ParseError = err
		return res, err
	}

	res.Balances = balances

	return res, nil
}
