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
	"crypto/ecdsa"
	"math/big"

	bitbonErrors "github.com/simcord-llc/bitbon-system-blockchain/bitbon/errors"
	"github.com/simcord-llc/bitbon-system-blockchain/common"
)

func (m *Manager) getBitbonSupportImpl() (*BitbonSupportImpl, error) {
	if !m.Ready() {
		return nil, ErrorNotReady
	}
	contractAddress, err := m.GetContractAddressBitbonSupport()
	if err != nil {
		return nil, err
	}

	contract, err := NewBitbonSupportImpl(contractAddress, m.EthApiWrapper)
	if err != nil {
		return nil, err
	}

	return contract, nil
}

// DirectTransfer transfers bitbon unsafely (directly from one assetbox to another)
func (m *Manager) DirectTransfer(ctx context.Context, from common.Address, to common.Address, value *big.Int, extraData []byte, key *ecdsa.PrivateKey, async bool) (blockNum uint64, txHash common.Hash, err error) {
	m.logger.Debug("contracts manager DirectTransfer called", "from", from.Hex(), "to", to.Hex(), "value", value.String())

	defer func() {
		if err != nil {
			err = bitbonErrors.NewContractCallError(err)
		}
	}()

	contract, err := m.getBitbonSupportImpl()
	if err != nil {
		return
	}

	opts, err := m.prepareTransactOpts(ctx, key)
	if err != nil {
		return
	}

	tx, err := contract.TransferFrom(opts, from, to, value, extraData)
	if err != nil {
		return
	}

	if async {
		return 0, tx.Hash(), nil
	}

	receipt, err := m.waitAndCheckTransaction(ctx, tx)
	if err != nil {
		return
	}

	return m.getReceiptParams(receipt)
}
