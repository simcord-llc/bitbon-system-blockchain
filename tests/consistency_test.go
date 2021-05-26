package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"testing"

	"github.com/simcord-llc/bitbon-system-blockchain/core/types"
	"github.com/simcord-llc/bitbon-system-blockchain/ethclient"
)

var (
	consistencyTestConfig = ConsistencyTestConfig{
		startBlock:     big.NewInt(1), // set startBlock to nil for 'latest'
		endBlock:       nil,           // set endBlock to nil for 'latest'
		checkBlockTime: true,
		blockTime:      1, // seconds
		verbose:        true,
		nodes: []string{
			"ws://172.16.28.46:10555",
			"ws://172.16.28.46:10556",
			"ws://172.16.28.45:10555",
			"ws://172.16.28.45:10556",
		},
	}
)

type ConsistencyTestConfig struct {
	startBlock     *big.Int
	endBlock       *big.Int
	nodes          []string
	checkBlockTime bool
	blockTime      int
	verbose        bool
}

type NodeHeader struct {
	node   string
	header *types.Header
}

func compareHeaders(headers []*NodeHeader) (err error) {
	newError := func(msg string) error {
		return fmt.Errorf("different header %s", msg)
	}
	if len(headers) < 2 {
		return errors.New("too few headers to compare")
	}
	controlHeader := headers[0]
	for i := 1; i < len(headers); i++ {
		currentHeader := headers[i]
		switch {
		case currentHeader.header.Number.Int64() != controlHeader.header.Number.Int64():
			return newError("numbers")
		case currentHeader.header.Hash() != controlHeader.header.Hash():
			return newError("hashes")
		case currentHeader.header.Root != controlHeader.header.Root:
			return newError("state roots")
		case currentHeader.header.Time != controlHeader.header.Time:
			return newError("timestamps")
		case currentHeader.header.Coinbase != controlHeader.header.Coinbase:
			return newError("miners")
		case currentHeader.header.ParentHash != controlHeader.header.ParentHash:
			return newError("parents")
		case !bytes.Equal(currentHeader.header.Extra, controlHeader.header.Extra):
			return newError("extras")
		case currentHeader.header.TxHash != controlHeader.header.TxHash:
			return newError("transaction roots")
		case currentHeader.header.Difficulty.Uint64() != controlHeader.header.Difficulty.Uint64():
			return newError("difficulty")
		case currentHeader.header.Nonce.Uint64() != controlHeader.header.Nonce.Uint64():
			return newError("nonces")
		case currentHeader.header.ReceiptHash != controlHeader.header.ReceiptHash:
			return newError("receipt hashes")
		case currentHeader.header.GasUsed != controlHeader.header.GasUsed:
			return newError("amounts of gas used")
		}
	}
	return nil
}

func checkBlockTimes(current *types.Header, parent *types.Header) (delta int64, err error) {
	return 1, nil
	//delta = current.Time.Sub(parent.Time, ).Int64()
	//if delta != int64(consistencyTestConfig.blockTime) {
	//	return delta, errors.New("wrong block time")
	//}
	//return delta, nil
}

type ConsistencyNode struct {
	uri    string
	client *ethclient.Client
}

func TestConsistency(t *testing.T) {
	nodes := make([]*ConsistencyNode, len(consistencyTestConfig.nodes))
	for index, nodeUri := range consistencyTestConfig.nodes {
		client, err := ethclient.Dial(nodeUri)
		if err != nil {
			fmt.Printf("Unable to connect to %s: %s\n", nodeUri, err.Error())
			t.Fail()
		}
		nodes[index] = &ConsistencyNode{
			uri:    nodeUri,
			client: client,
		}
	}

	controlNode := nodes[0]
	start := consistencyTestConfig.startBlock
	if start == nil {
		h, err := controlNode.client.HeaderByNumber(context.TODO(), start)
		if err != nil || h == nil {
			fmt.Printf("Unable to resolve 'latest' for startBlock\n")
			t.Fail()
			return
		}
		start = h.Number
	}

	end := consistencyTestConfig.endBlock
	if end == nil {
		h, err := controlNode.client.HeaderByNumber(context.TODO(), end)
		if err != nil || h == nil {
			fmt.Printf("Unable to resolve 'latest' for endBlock\n")
			t.Fail()
			return
		}
		end = h.Number
	}

	var prevHeader *types.Header

	for i := start.Int64(); i < end.Int64(); i++ {
		nodeHeaders := make([]*NodeHeader, len(nodes))
		for index, node := range nodes {
			header, err := node.client.HeaderByNumber(context.TODO(), big.NewInt(i))
			if err != nil {
				fmt.Printf("Unable to retreive block %d from node %s: %s\n", i, node.uri, err.Error())
				t.Fail()
				return
			}
			nodeHeaders[index] = &NodeHeader{node: node.uri, header: header}
		}

		err := compareHeaders(nodeHeaders)
		if err != nil {
			j, e := json.Marshal(nodeHeaders)
			if e != nil {
				fmt.Printf("Unable to marshall nodeHeaders to json: %s\n", err.Error())
				t.Fail()
				return
			}
			fmt.Printf("Inconsistency found on block %d:\n%s\n\n%v\n", i, err.Error(), string(j))
			t.Fail()
			return
		}

		delta := int64(0)
		currHeader := nodeHeaders[0]
		if prevHeader != nil {
			delta, err = checkBlockTimes(prevHeader, currHeader.header)
			if err != nil && consistencyTestConfig.checkBlockTime {
				t.Errorf("Blocks %d and %d have incorrect time delta (%d s)", prevHeader.Number.Int64(), currHeader.header.Number.Int64(), delta)
			}
		}

		prevHeader = nodeHeaders[0].header

		if consistencyTestConfig.verbose {
			fmt.Printf("# %d H %s ⛏ %s ⏳ %d (+%ds)\n", currHeader.header.Number.Int64(), currHeader.header.Hash().Hex(), currHeader.header.Coinbase.Hex(), currHeader.header.Time, delta)
		}
	}

}
