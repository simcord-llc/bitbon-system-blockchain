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

	"math/big"

	"github.com/simcord-llc/bitbon-system-blockchain/accounts/abi/bind"
	bitbonErrors "github.com/simcord-llc/bitbon-system-blockchain/bitbon/errors"
	"github.com/simcord-llc/bitbon-system-blockchain/common"
	"github.com/simcord-llc/bitbon-system-blockchain/core/types"
)

// EmitEther emits address with assetboxEtherEmission value using service PK
func (m *Manager) EmitEther(ctx context.Context, to common.Address, value *big.Int, nonce *uint64) (err error) {
	m.logger.Debug("EmitEther called", "to", to.Hex(), "value", value.String())

	defer func() {
		if err != nil {
			err = bitbonErrors.NewEmitEtherError(err)
		}
	}()

	if value == nil {
		value = m.suggestEtherEmission()
	}

	args := Args{
		FromPK: m.servicePK,
		To:     &to,
		Value:  value,
	}

	if nonce != nil {
		args.Nonce = nonce
	} else if m.Noncer != nil {
		redisNonce, err := m.Noncer.IncrementAndGetNonce(m.serviceAddress) // nolint:govet
		if err == nil {
			tmp := uint64(redisNonce)
			args.Nonce = &tmp
		} else {
			m.logger.Error("Error getting noncer nonce", "error", err)
		}
	} else {
		m.logger.Debug("Noncer is disabled. Continue with default behaviour")
	}

	tx, err := m.EthApiWrapper.SendBitbonTransaction(ctx, args)
	if err != nil {
		return err
	}

	m.logger.Debug("Waiting for tx receipt (emit ether)...", "txHash", tx.Hash().Hex())

	ctx, cancel := context.WithTimeout(ctx, waitForTxReceiptTimeout)
	defer cancel()

	receipt, err := bind.WaitMined(ctx, m.EthApiWrapper, tx)
	if err != nil {
		return err
	}
	if receipt.Status != types.ReceiptStatusSuccessful {
		return ErrorUnsuccessfulTxEmitEther
	}

	m.logger.Debug("Assetbox successfully emitted with ether, receipt obtained", "assetbox", to.Hex(), "value", value.String(), "txHash", tx.Hash().Hex(), "gasUsed", receipt.GasUsed, "success", receipt.Status)
	return nil
}
