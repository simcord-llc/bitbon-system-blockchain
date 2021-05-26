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

package v1

import (
	"context"
	"math/big"

	"github.com/pkg/errors"
	cs "github.com/simcord-llc/bitbon-system-blockchain/bitbon/contracts/contract_snapshot"
	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/dto"
	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/parser/interfaces"
	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/parser/method_parser"
	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/parser/storage_contract_snapshot"
	"github.com/simcord-llc/bitbon-system-blockchain/common"
	"github.com/simcord-llc/bitbon-system-blockchain/core/types"
)

type MonetizeCertificateParser struct {
	*method_parser.BaseMethodParser
}

func NewMonetizeCertificateParser(ethApiWrapper interfaces.Contract, contractsSnapshot *storage_contract_snapshot.StorageContractSnapshot) *MonetizeCertificateParser {
	var p = &MonetizeCertificateParser{
		BaseMethodParser: method_parser.NewBaseMethodParser(
			ethApiWrapper,
			contractsSnapshot,
			cs.ContractTypeBitbonSupportImpl,
			dto.MonetizeCertificateV1MethodID,
		),
	}

	return p
}

func (p *MonetizeCertificateParser) Parse(ctx context.Context, tx *types.Transaction) (*dto.BitbonTx, error) {
	res := &dto.BitbonTx{ParseStatus: dto.ParseStatusSuccess}
	bst := &dto.BitbonMonetizeArgs{}

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

	balances := make(map[common.Address]*big.Int)

	err = p.ParseBalanceChangedLogs(balances, p.ContractType, *tx.To(), receipt.Logs)
	if err != nil {
		res.ParseStatus = dto.ParseStatusParseLogsError
		err = errors.Wrap(err, "unable to parse balance changed logs")
		res.ParseError = err
		return res, err
	}

	res.Balances = balances

	return res, nil
}
