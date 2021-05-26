// Copyright 2017 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package quorum

import (
	"math/big"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/simcord-llc/bitbon-system-blockchain/common"
	"github.com/simcord-llc/bitbon-system-blockchain/core/types"
	"github.com/simcord-llc/bitbon-system-blockchain/params"
)

func TestInturn(t *testing.T) {
	firstSigner := common.HexToAddress("0x66acc38944b6f36284dde66f9d6ae7cbb16f5e74")
	secondSigner := common.HexToAddress("0xc46eae38e2ab01f91de7afff3097596dfff0c493")
	now := time.Now().Unix()
	config := &params.QuorumEngineConfig{Period: 2, Epoch: 100, Start: uint64(now), InitialSigners: []common.Address{firstSigner, secondSigner}}
	type TestCase struct {
		signer    common.Address
		blockTime *big.Int
		expected  bool
	}
	cases := []TestCase{
		{
			signer:    firstSigner,
			blockTime: big.NewInt(now),
			expected:  true,
		},
		{
			signer:    firstSigner,
			blockTime: big.NewInt(now + 1),
			expected:  true,
		},
		{
			signer:    firstSigner,
			blockTime: big.NewInt(now + 2),
			expected:  false,
		},
		{
			signer:    secondSigner,
			blockTime: big.NewInt(now + 2),
			expected:  true,
		},
	}
	signers := map[common.Address]BlockProducer{
		firstSigner: {
			SequenceNumber: 0,
		},
		secondSigner: {
			SequenceNumber: 1,
		},
	}

	snapshot := newSnapshot(config, nil, signers, 0)
	for _, testCase := range cases {
		header := &types.Header{
			Number: big.NewInt(1),
			Time:   testCase.blockTime.Uint64(),
		}
		inturn := snapshot.inturn(header, testCase.signer)

		assert.Equal(t, testCase.expected, inturn)
	}
}
