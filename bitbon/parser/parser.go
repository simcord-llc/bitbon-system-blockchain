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

package parser

import (
	"context"

	"github.com/davecgh/go-spew/spew"
	"github.com/simcord-llc/bitbon-system-blockchain/log"

	loggerContext "github.com/simcord-llc/bitbon-system-blockchain/bitbon/context"
	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/dto"
	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/parser/interfaces"
	v1 "github.com/simcord-llc/bitbon-system-blockchain/bitbon/parser/method_parser/v1"
	v2 "github.com/simcord-llc/bitbon-system-blockchain/bitbon/parser/method_parser/v2"
	v3 "github.com/simcord-llc/bitbon-system-blockchain/bitbon/parser/method_parser/v3"
	scs "github.com/simcord-llc/bitbon-system-blockchain/bitbon/parser/storage_contract_snapshot"
	"github.com/simcord-llc/bitbon-system-blockchain/common"
	"github.com/simcord-llc/bitbon-system-blockchain/core/types"
)

type MethodParserConstructor func(*scs.StorageContractSnapshot) interfaces.MethodParser
type Parser struct {
	storage       interfaces.BitbonStorage
	methodParsers map[dto.ContractVersion]map[string]MethodParserConstructor

	logger log.Logger
}

// nolint:funlen
func NewParser(ethAPIWrapper interfaces.Contract, storage interfaces.BitbonStorage) *Parser {
	return &Parser{
		logger: loggerContext.LoggerFromContext(loggerContext.NewLoggerContext(context.Background(),
			log.New("src", "Manager"))),

		storage: storage,
		methodParsers: map[dto.ContractVersion]map[string]MethodParserConstructor{
			dto.VersionV1: {
				dto.DirectTransferV1MethodID: func(c *scs.StorageContractSnapshot) interfaces.MethodParser {
					return v1.NewDirectTransferParser(ethAPIWrapper, c)
				},
				dto.CreateSafeTransferV1MethodID: func(c *scs.StorageContractSnapshot) interfaces.MethodParser {
					return v1.NewCreateSafeTransferParser(ethAPIWrapper, c)
				},
				dto.ApproveSafeTransferV1MethodID: func(c *scs.StorageContractSnapshot) interfaces.MethodParser {
					return v1.NewApproveSafeTransferParser(ethAPIWrapper, c)
				},
				dto.CancelSafeTransferV1MethodID: func(c *scs.StorageContractSnapshot) interfaces.MethodParser {
					return v1.NewCancelSafeTransferParser(ethAPIWrapper, c)
				},
				dto.MonetizeCertificateV1MethodID: func(c *scs.StorageContractSnapshot) interfaces.MethodParser {
					return v1.NewMonetizeCertificateParser(ethAPIWrapper, c)
				},
				dto.MonetizeEmissionV1MethodID: func(c *scs.StorageContractSnapshot) interfaces.MethodParser {
					return v1.NewMonetizeEmissionParser(ethAPIWrapper, c)
				},
			},
			dto.VersionV2: {
				dto.DirectTransferV2MethodID: func(c *scs.StorageContractSnapshot) interfaces.MethodParser {
					return v2.NewDirectTransferParser(ethAPIWrapper, c)
				},
				dto.CreateSafeTransferV2MethodID: func(c *scs.StorageContractSnapshot) interfaces.MethodParser {
					return v2.NewCreateSafeTransferParser(ethAPIWrapper, c)
				},
				dto.ApproveSafeTransferV2MethodID: func(c *scs.StorageContractSnapshot) interfaces.MethodParser {
					return v2.NewApproveSafeTransferParser(ethAPIWrapper, c)
				},
				dto.CancelSafeTransferV2MethodID: func(c *scs.StorageContractSnapshot) interfaces.MethodParser {
					return v2.NewCancelSafeTransferParser(ethAPIWrapper, c)
				},
				dto.ExpireSafeTransferV2MethodID: func(c *scs.StorageContractSnapshot) interfaces.MethodParser {
					return v2.NewExpireSafeTransferParser(ethAPIWrapper, c)
				},
				dto.MonetizeCertificateV2MethodID: func(c *scs.StorageContractSnapshot) interfaces.MethodParser {
					return v2.NewMonetizeCertificateParser(ethAPIWrapper, c)
				},
				dto.MonetizeEmissionV2MethodID: func(c *scs.StorageContractSnapshot) interfaces.MethodParser {
					return v2.NewMonetizeEmissionParser(ethAPIWrapper, c)
				},
			},
			dto.VersionV3: {
				dto.DirectTransferV3MethodID: func(c *scs.StorageContractSnapshot) interfaces.MethodParser {
					return v3.NewDirectTransferParser(ethAPIWrapper, c)
				},
				dto.CreateSafeTransferV3MethodID: func(c *scs.StorageContractSnapshot) interfaces.MethodParser {
					return v3.NewCreateSafeTransferParser(ethAPIWrapper, c)
				},
				dto.CreateFullBalanceSafeTransferV3MethodID: func(c *scs.StorageContractSnapshot) interfaces.MethodParser {
					return v3.NewCreateFullBalanceSafeTransferParser(ethAPIWrapper, c)
				},
				dto.ApproveSafeTransferV3MethodID: func(c *scs.StorageContractSnapshot) interfaces.MethodParser {
					return v3.NewApproveSafeTransferParser(ethAPIWrapper, c)
				},
				dto.ApproveFullBalanceSafeTransferV3MethodID: func(c *scs.StorageContractSnapshot) interfaces.MethodParser {
					return v3.NewApproveFullBalanceSafeTransferParser(ethAPIWrapper, c)
				},
				dto.CancelSafeTransferV3MethodID: func(c *scs.StorageContractSnapshot) interfaces.MethodParser {
					return v3.NewCancelSafeTransferParser(ethAPIWrapper, c)
				},
				dto.CancelFullBalanceSafeTransferV3MethodID: func(c *scs.StorageContractSnapshot) interfaces.MethodParser {
					return v3.NewCancelFullBalanceSafeTransferParser(ethAPIWrapper, c)
				},
				dto.ExpireSafeTransferV3MethodID: func(c *scs.StorageContractSnapshot) interfaces.MethodParser {
					return v3.NewExpireSafeTransferParser(ethAPIWrapper, c)
				},
				dto.QuickTransferV3MethodID: func(c *scs.StorageContractSnapshot) interfaces.MethodParser {
					return v3.NewQuickTransferParser(ethAPIWrapper, c)
				},
				dto.FullBalanceQuickTransferV3MethodId: func(c *scs.StorageContractSnapshot) interfaces.MethodParser {
					return v3.NewFullBalanceQuickTransferParser(ethAPIWrapper, c)
				},
				dto.FrameTransferV3MethodID: func(c *scs.StorageContractSnapshot) interfaces.MethodParser {
					return v3.NewFrameTransferParser(ethAPIWrapper, c)
				},
				dto.CreateWPCSafeTransferV3MethodID: func(c *scs.StorageContractSnapshot) interfaces.MethodParser {
					return v3.NewCreateWPCSafeTransferParser(ethAPIWrapper, c)
				},
				dto.CreateFullBalanceWPCSafeTransferV3MethodId: func(c *scs.StorageContractSnapshot) interfaces.MethodParser {
					return v3.NewCreateFullBalanceWpcSafeTransferParser(ethAPIWrapper, c)
				},
				dto.ApproveWPCSafeTransferV3MethodID: func(c *scs.StorageContractSnapshot) interfaces.MethodParser {
					return v3.NewApproveWPCSafeTransferParser(ethAPIWrapper, c)
				},
				dto.ApproveFullBalanceWPCSafeTransferV3MethodID: func(c *scs.StorageContractSnapshot) interfaces.MethodParser {
					return v3.NewApproveFullBalanceWpcSafeTransferParser(ethAPIWrapper, c)
				},
				dto.CancelWPCSafeTransferV3MethodID: func(c *scs.StorageContractSnapshot) interfaces.MethodParser {
					return v3.NewCancelWPCSafeTransferParser(ethAPIWrapper, c)
				},
				dto.CancelFullBalanceWPCSafeTransferV3MethodID: func(c *scs.StorageContractSnapshot) interfaces.MethodParser {
					return v3.NewCancelFullBalanceWpcSafeTransferParser(ethAPIWrapper, c)
				},
				dto.ServiceFeeTransferV3MethodID: func(c *scs.StorageContractSnapshot) interfaces.MethodParser {
					return v3.NewServiceFeeTransferParser(ethAPIWrapper, c)
				},
			},
		},
	}
}

func (p *Parser) Parse(ctx context.Context, tx *types.Transaction) *dto.BitbonTx {
	data := tx.Data()
	if len(data) == 0 {
		return nil
	}

	to := tx.To()
	if to == nil {
		return nil
	}

	storage, err := scs.NewStorageSnapshotByContractAddress(p.storage, *to)
	if err != nil {
		p.logger.Warn("unable to get snapshot from storage")
		return nil
	}

	methodID := common.Bytes2Hex(data[:4])

	methodParserConstructor, ok := p.methodParsers[storage.Version][methodID]
	if !ok {
		p.logger.Warn("no parser found for transaction", "version", storage.Version, "txHash", tx.Hash(), "methodID", methodID)
		return nil
	}

	res, err := methodParserConstructor(storage).Parse(ctx, tx)
	if err != nil {
		p.logger.Warn("error parsing method", "version", storage.Version, "error", err, "res", spew.Sdump(res))
		return nil
	}

	return res
}
