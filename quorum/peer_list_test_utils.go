package quorum

import (
	"crypto/ecdsa"
	"crypto/rand"
	"math/big"

	"github.com/simcord-llc/bitbon-system-blockchain/common"
	"github.com/simcord-llc/bitbon-system-blockchain/crypto"
	"github.com/simcord-llc/bitbon-system-blockchain/p2p"
	"github.com/simcord-llc/bitbon-system-blockchain/p2p/enode"
	"github.com/simcord-llc/bitbon-system-blockchain/p2p/enr"
	"github.com/simcord-llc/bitbon-system-blockchain/p2p/protocols"
	"github.com/simcord-llc/bitbon-system-blockchain/quorum/snapshot"
)

func getNewPeerList(hasMajority majorityChecker) *PeerList {
	return newPeerList(getNewNode(), hasMajority)
}

func getNewNode() *enode.Node {
	record := new(enr.Record)

	key, err := ecdsa.GenerateKey(crypto.S256(), rand.Reader)
	if err != nil {
		panic("Error during key generation")
	}

	err = enode.SignV4(record, key)
	if err != nil {
		panic("Error during record signing")
	}

	node, err := enode.New(enode.ValidSchemes, record)
	if err != nil {
		panic("Error during Node creation")
	}

	return node
}

func createNewPeerStruct(id enode.ID, name string, inPingList, isMinerCandidate bool) *Peer {
	p := int64(0)
	return &Peer{
		Peer:    protocols.NewPeer(p2p.NewPeer(id, name, nil), nil, nil),
		address: common.Address{},
		peerState: &peerState{
			inPingList:       inPingList,
			isMinerCandidate: isMinerCandidate,
			lastPing:         &p,
		},
	}
}

func getRandomEnodeID() enode.ID {
	letters := []rune("0123456789abcdef")

	s := make([]rune, 66)
	s[0] = '0'
	s[1] = 'x'
	for i := 2; i < 66; i++ {
		index, _ := rand.Int(rand.Reader, big.NewInt(16))
		s[i] = letters[index.Uint64()]
	}
	return enode.HexID(string(s))
}

func prepareExpectedNodesForPeerList(peerList *PeerList) []*snapshot.Node {
	selfRemoteNodes := make(map[common.Address]bool)
	selfRemoteNodes[peerList.selfAddress] = peerList.SelfIsMinerCandidate()

	nodes := make([]*snapshot.Node, 0, len(peerList.peers))

	for _, v := range peerList.peers {
		if v.inPingList {
			selfRemoteNodes[v.Address()] = v.IsMinerCandidate()
			nodes = append(nodes, snapshot.NewNode(v.ID(), v.Address(), v.IsMinerCandidate(), nil))
		}
	}

	nodes = append(nodes,
		snapshot.NewNode(peerList.self.ID(), peerList.selfAddress, peerList.SelfIsMinerCandidate(), selfRemoteNodes))

	return nodes
}

func hasMajorityTrue(int, bool) bool {
	return true
}

func hasMajorityFalse(int, bool) bool {
	return false
}
