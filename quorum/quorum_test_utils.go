package quorum

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"net"
	"testing"
	"time"

	"github.com/simcord-llc/bitbon-system-blockchain/common"
	"github.com/simcord-llc/bitbon-system-blockchain/crypto"
	"github.com/simcord-llc/bitbon-system-blockchain/log"
	"github.com/simcord-llc/bitbon-system-blockchain/p2p"
	"github.com/simcord-llc/bitbon-system-blockchain/p2p/enode"
)

type nodeParams struct {
	nodeHexKey string
	name       string
	ip         net.IP
	port       int
}

var (
	node1Params = &nodeParams{
		nodeHexKey: "faccfaccfaccfaccfaccfaccfaccfaccfaccfaccfaccfaccfaccfaccfacc0001",
		name:       "node1",
		ip:         net.IPv4(0, 0, 0, 0),
		port:       30300,
	}
	node2Params = &nodeParams{
		nodeHexKey: "faccfaccfaccfaccfaccfaccfaccfaccfaccfaccfaccfaccfaccfaccfacc0002",
		name:       "node2",
		ip:         net.IPv4(0, 0, 0, 0),
		port:       30301,
	}
	node3Params = &nodeParams{
		nodeHexKey: "faccfaccfaccfaccfaccfaccfaccfaccfaccfaccfaccfaccfaccfaccfacc0003",
		name:       "node3",
		ip:         net.IPv4(0, 0, 0, 0),
		port:       30302,
	}

	epochNumber    = uint64(10)
	numOfAddresses = 10

	ErrOnGettingDistribution = errors.New("error during getCurrentDistribution contract call")
	ErrGettingSnapshot       = errors.New("error during snapshot nodes getting")

	logger = log.New(context.TODO())
)

func startTestServer(t *testing.T, params *nodeParams) (*p2p.Server, *Quorum) {
	staticNodes := staticNodes()
	config := p2p.Config{
		Name:         params.name,
		MaxPeers:     10,
		ListenAddr:   fmt.Sprintf("%s:%d", params.ip.String(), params.port),
		PrivateKey:   parseKey(params.nodeHexKey),
		StaticNodes:  staticNodes,
		TrustedNodes: staticNodes,
	}

	server := &p2p.Server{
		Config: config,
	}
	quorumService, err := New(nil, &Config{
		HandshakeTimeout: time.Second * 5,
		PingFrequency:    time.Second * 5,
		QuorumMaxPing:    time.Second * 5,
	})

	if err != nil {
		t.Fatalf("error while creating New() quorum: %s", err)
	}
	server.Protocols = append(server.Protocols, quorumService.Protocols()...)

	if err := server.Start(); err != nil {
		t.Fatalf("Could not start server: %v", err)
	}

	if err := quorumService.Start(server); err != nil {
		t.Fatalf("Could not start quorumService: %v", err)
	}

	return server, quorumService
}

func staticNodes() []*enode.Node {
	nodesParams := []*nodeParams{
		node1Params,
		node2Params,
		node3Params,
	}

	nodes := make([]*enode.Node, 0, len(nodesParams))

	for _, nodeParams := range nodesParams {
		privateKey := parseKey(nodeParams.nodeHexKey)
		node := enode.NewV4(&privateKey.PublicKey, nodeParams.ip, nodeParams.port, nodeParams.port)

		nodes = append(nodes, node)
	}
	return nodes
}

func parseKey(keyHex string) *ecdsa.PrivateKey {
	key, err := crypto.HexToECDSA(keyHex)
	if err != nil {
		panic(fmt.Sprintf("error while generating key: %s", err))
	}

	return key
}

func setupLogger() {
	log.Root().SetHandler(log.LvlFilterHandler(log.LvlDebug, log.StdoutHandler))
}

func nodeListToMap(list []*enode.Node) map[enode.ID]*enode.Node {
	nodeMap := make(map[enode.ID]*enode.Node)

	for _, v := range list {
		nodeMap[v.ID()] = v
	}

	return nodeMap
}

func minerListToMap(list []common.Address, epoch uint64) map[common.Address]uint64 {
	minerMap := make(map[common.Address]uint64)

	for _, v := range list {
		minerMap[v] = epoch
	}

	return minerMap
}
