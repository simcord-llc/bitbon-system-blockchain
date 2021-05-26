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

package method_parser

import (
	"math/big"

	"github.com/davecgh/go-spew/spew"
	"github.com/pkg/errors"
	"github.com/simcord-llc/bitbon-system-blockchain/accounts/abi/bind"
	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/contracts"
	cs "github.com/simcord-llc/bitbon-system-blockchain/bitbon/contracts/contract_snapshot"
	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/dto"
	bitbonErrors "github.com/simcord-llc/bitbon-system-blockchain/bitbon/errors"
	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/parser/interfaces"
	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/parser/storage_contract_snapshot"
	"github.com/simcord-llc/bitbon-system-blockchain/common"
	"github.com/simcord-llc/bitbon-system-blockchain/core/types"
)

type BaseMethodParser struct {
	ethApiWrapper     interfaces.Contract
	contractsSnapshot *storage_contract_snapshot.StorageContractSnapshot
	ContractType      cs.ContractType
	MethodId          string
}

func NewBaseMethodParser(ethApiWrapper interfaces.Contract, contractsSnapshot *storage_contract_snapshot.StorageContractSnapshot, contractType cs.ContractType, methodId string) *BaseMethodParser {
	p := &BaseMethodParser{
		ethApiWrapper:     ethApiWrapper,
		contractsSnapshot: contractsSnapshot,
		ContractType:      contractType,
		MethodId:          methodId,
	}

	return p
}

func (p *BaseMethodParser) ParseInput(inputOut interface{}, tx *types.Transaction) (string, error) {
	contract, err := p.contractsSnapshot.GetAbi(p.ContractType)
	if err != nil {
		return "", errors.Wrap(err, "unable to get safe transfer storage v1 address")
	}

	method, err := contract.MethodById(common.Hex2Bytes(p.MethodId))
	if err != nil {
		return "", errors.Errorf("unable to get %s method in contract %s of version %s: %s", p.MethodId, p.ContractType, p.contractsSnapshot.Version, err.Error())
	}

	err = method.Inputs.Unpack(inputOut, tx.Data()[4:])
	if err != nil {
		return "", errors.Wrap(err, "unable to parse input")
	}

	return method.Name, nil
}

func (p *BaseMethodParser) ParseLogs(contractType cs.ContractType, contractAddress common.Address, eventName dto.EventType, outConstructor func() interface{}, logs []*types.Log) ([]interface{}, error) {
	if len(logs) == 0 {
		return make([]interface{}, 0), nil
	}

	contract, err := p.contractsSnapshot.GetAbi(contractType)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get safe transfer storage v1 address")
	}

	boundContract := bind.NewBoundContract(contractAddress, contract, p.ethApiWrapper, p.ethApiWrapper, p.ethApiWrapper)
	filteredLogs, err := boundContract.FilterGivenLogs(logs, string(eventName))
	if err != nil {
		return nil, errors.Wrap(err, "unable to filter logs")
	}

	res := make([]interface{}, 0, 5)
	for _, l := range filteredLogs {
		out := outConstructor()
		err = boundContract.UnpackLog(out, string(eventName), l)
		if err != nil {
			continue
		}

		res = append(res, out)
	}

	return res, nil
}

func (p *BaseMethodParser) EthApiWrapper() interfaces.Contract {
	return p.ethApiWrapper
}

func (p *BaseMethodParser) ContractsSnapshot() *storage_contract_snapshot.StorageContractSnapshot {
	return p.contractsSnapshot
}

func (p *BaseMethodParser) Sender(tx *types.Transaction) (common.Address, error) {
	var signer types.Signer = types.FrontierSigner{}
	if tx.Protected() {
		signer = types.NewEIP155Signer(tx.ChainID())
	}

	return types.Sender(signer, tx)
}

func (p *BaseMethodParser) ParseBalanceChangedLogs(balances map[common.Address]*big.Int, contractType cs.ContractType, contractAddress common.Address, rawLogs []*types.Log) error {
	outConstructor := func() interface{} {
		return &dto.BitbonBalanceChanged{}
	}

	logs, err := p.ParseLogs(contractType, contractAddress, dto.BalanceChangedEventType, outConstructor, rawLogs)
	if err != nil {
		return err
	}

	for _, v := range logs {
		l, ok := v.(*dto.BitbonBalanceChanged)
		if !ok {
			return errors.New("unable to cast log to BitbonBalanceChanged")
		}

		balances[l.Assetbox] = l.Balance
	}

	return nil
}

func (p *BaseMethodParser) ParseBalanceLockedLogs(balances map[common.Address]*big.Int, contractType cs.ContractType, contractAddress common.Address, rawLogs []*types.Log) error {
	outConstructor := func() interface{} {
		return &dto.BitbonBalanceLocked{}
	}

	logs, err := p.ParseLogs(contractType, contractAddress, dto.BalanceLockedEventType, outConstructor, rawLogs)
	if err != nil {
		return err
	}

	for _, v := range logs {
		l, ok := v.(*dto.BitbonBalanceLocked)
		if !ok {
			return errors.New("unable to cast log to BitbonBalanceLocked")
		}

		balances[l.Assetbox] = l.BalanceAviable
	}

	return nil
}

func (p *BaseMethodParser) ParseBalanceUnLockedLogs(balances map[common.Address]*big.Int, contractType cs.ContractType, contractAddress common.Address, rawLogs []*types.Log) error {
	outConstructor := func() interface{} {
		return &dto.BitbonBalanceUnlocked{}
	}

	logs, err := p.ParseLogs(contractType, contractAddress, dto.BalanceUnLockedEventType, outConstructor, rawLogs)
	if err != nil {
		return err
	}

	for _, v := range logs {
		l, ok := v.(*dto.BitbonBalanceUnlocked)
		if !ok {
			return errors.New("unable to cast log to BitbonBalanceUnlocked")
		}

		balances[l.Assetbox] = l.BalanceAviable
	}

	return nil
}

func (p *BaseMethodParser) ParseTransferExpiredLogs(contractType cs.ContractType, contractAddress common.Address, rawLogs []*types.Log) (res [][]byte, err error) {
	outConstructor := func() interface{} {
		return &dto.TransferExpiredEvent{}
	}

	logs, err := p.ParseLogs(contractType, contractAddress, dto.TransferExpiredEventType, outConstructor, rawLogs)
	if err != nil {
		return nil, err
	}

	res = make([][]byte, 0, len(logs))

	for _, v := range logs {
		l, ok := v.(*dto.TransferExpiredEvent)
		if !ok {
			return nil, errors.New("unable to cast log to TransferExpiredEvent")
		}

		res = append(res, l.TransferId)
	}

	return res, nil
}

func (p *BaseMethodParser) ParseSafeTransferApprovalResultLogs(contractType cs.ContractType, contractAddress common.Address, rawLogs []*types.Log) (bitbonErrors.Error, error) {
	outConstructor := func() interface{} {
		return &dto.ApproveTransferEvent{}
	}

	logs, err := p.ParseLogs(contractType, contractAddress, "SafeTransferApprovalResult", outConstructor, rawLogs)
	if err != nil {
		return nil, err
	}

	for _, v := range logs {
		l, ok := v.(*dto.ApproveTransferEvent)
		if !ok {
			return nil, errors.New("unable to cast log to ApproveTransferEvent")
		}

		spew.Dump(l)

		switch contracts.StatusType(l.Status) {
		case contracts.SafeTransferStatusSuccess:
			return nil, nil
		case contracts.SafeTransferStatusUndefined:
			return bitbonErrors.NewSafeTransferUndefinedError(errors.New("transfer approve failed - undefined status")), nil
		case contracts.SafeTransferStatusWrongProtectionCode:
			return bitbonErrors.NewSafeTransferWrongProtectionCodeError(errors.New("transfer approve failed - wrong protection code")), nil
		case contracts.SafeTransferStatusWrongProtectionCodeLimitExceeded:
			return bitbonErrors.NewSafeTransferWrongProtectionCodeLimitError(errors.New("transfer approve failed - wrong protection code limit exceeded")), nil
		case contracts.SafeTransferStatusExpired:
			return bitbonErrors.NewSafeTransferExpiredError(errors.New("transfer approve failed - expired")), nil
		}
	}

	return nil, errors.New("logs not found")
}
