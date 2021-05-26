// Copyright Simcord LLC
// This file is part of Bitbon System Blockchain Node.
// This file has been modified, you can find the original version by following the link
// <https://github.com/ethereum/go-ethereum>
//
// Bitbon System Blockchain Node is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// Bitbon System Blockchain Node is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with Bitbon System Blockchain Node. If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"strconv"

	"github.com/simcord-llc/bitbon-system-blockchain/accounts"
	"github.com/simcord-llc/bitbon-system-blockchain/accounts/abi/bind"
	"github.com/simcord-llc/bitbon-system-blockchain/accounts/external"
	"github.com/simcord-llc/bitbon-system-blockchain/cmd/utils"
	"github.com/simcord-llc/bitbon-system-blockchain/common"
	"github.com/simcord-llc/bitbon-system-blockchain/contracts/checkpointoracle"
	"github.com/simcord-llc/bitbon-system-blockchain/ethclient"
	"github.com/simcord-llc/bitbon-system-blockchain/params"
	"github.com/simcord-llc/bitbon-system-blockchain/rpc"
	"gopkg.in/urfave/cli.v1"
)

// newClient creates a client with specified remote URL.
func newClient(ctx *cli.Context) *ethclient.Client {
	client, err := ethclient.Dial(ctx.GlobalString(nodeURLFlag.Name))
	if err != nil {
		utils.Fatalf("Failed to connect to Ethereum node: %v", err)
	}
	return client
}

// newRPCClient creates a rpc client with specified node URL.
func newRPCClient(url string) *rpc.Client {
	client, err := rpc.Dial(url)
	if err != nil {
		utils.Fatalf("Failed to connect to Ethereum node: %v", err)
	}
	return client
}

// getContractAddr retrieves the register contract address through
// rpc request.
func getContractAddr(client *rpc.Client) common.Address {
	var addr string
	if err := client.Call(&addr, "les_getCheckpointContractAddress"); err != nil {
		utils.Fatalf("Failed to fetch checkpoint oracle address: %v", err)
	}
	return common.HexToAddress(addr)
}

// getCheckpoint retrieves the specified checkpoint or the latest one
// through rpc request.
func getCheckpoint(ctx *cli.Context, client *rpc.Client) *params.TrustedCheckpoint {
	var checkpoint *params.TrustedCheckpoint

	if ctx.GlobalIsSet(indexFlag.Name) {
		var result [3]string
		index := uint64(ctx.GlobalInt64(indexFlag.Name))
		if err := client.Call(&result, "les_getCheckpoint", index); err != nil {
			utils.Fatalf("Failed to get local checkpoint %v, please ensure the les API is exposed", err)
		}
		checkpoint = &params.TrustedCheckpoint{
			SectionIndex: index,
			SectionHead:  common.HexToHash(result[0]),
			CHTRoot:      common.HexToHash(result[1]),
			BloomRoot:    common.HexToHash(result[2]),
		}
	} else {
		var result [4]string
		err := client.Call(&result, "les_latestCheckpoint")
		if err != nil {
			utils.Fatalf("Failed to get local checkpoint %v, please ensure the les API is exposed", err)
		}
		index, err := strconv.ParseUint(result[0], 0, 64)
		if err != nil {
			utils.Fatalf("Failed to parse checkpoint index %v", err)
		}
		checkpoint = &params.TrustedCheckpoint{
			SectionIndex: index,
			SectionHead:  common.HexToHash(result[1]),
			CHTRoot:      common.HexToHash(result[2]),
			BloomRoot:    common.HexToHash(result[3]),
		}
	}
	return checkpoint
}

// newContract creates a registrar contract instance with specified
// contract address or the default contracts for mainnet or testnet.
func newContract(client *rpc.Client) (common.Address, *checkpointoracle.CheckpointOracle) {
	addr := getContractAddr(client)
	if addr == (common.Address{}) {
		utils.Fatalf("No specified registrar contract address")
	}
	contract, err := checkpointoracle.NewCheckpointOracle(addr, ethclient.NewClient(client))
	if err != nil {
		utils.Fatalf("Failed to setup registrar contract %s: %v", addr, err)
	}
	return addr, contract
}

// newClefSigner sets up a clef backend and returns a clef transaction signer.
func newClefSigner(ctx *cli.Context) *bind.TransactOpts {
	clef, err := external.NewExternalSigner(ctx.String(clefURLFlag.Name))
	if err != nil {
		utils.Fatalf("Failed to create clef signer %v", err)
	}
	return bind.NewClefTransactor(clef, accounts.Account{Address: common.HexToAddress(ctx.String(signerFlag.Name))})
}
