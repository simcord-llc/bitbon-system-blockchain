package quorum

import (
	"encoding/base64"
	"errors"

	quorum "github.com/simcord-llc/bitbon-system-blockchain/quorum/mock"

	"github.com/golang/mock/gomock"
	"github.com/simcord-llc/bitbon-system-blockchain/common"
	"github.com/simcord-llc/bitbon-system-blockchain/p2p"
	"github.com/simcord-llc/bitbon-system-blockchain/p2p/enode"
	"github.com/simcord-llc/bitbon-system-blockchain/p2p/protocols"
	"github.com/simcord-llc/bitbon-system-blockchain/quorum/snapshot"
	"github.com/simcord-llc/bitbon-system-blockchain/rlp"
)

var (
	ErrOnWriteMessage = errors.New("error during message writing")
	ErrHookValidate   = errors.New("error during hook validation")
	ErrHookApply      = errors.New("error during hook applying")
)

func prepareMockServices(controller *gomock.Controller) (*MockPingPeerService, *MockHandlers,
	*quorum.MockMsgReadWriter, *MockHook) {
	return NewMockPingPeerService(controller),
		NewMockHandlers(controller),
		quorum.NewMockMsgReadWriter(controller),
		NewMockHook(controller)
}

func getNewPeerState(pingTime int64) *peerState {
	return &peerState{
		lastPing:         &pingTime,
		inPingList:       false,
		isMinerCandidate: true,
	}
}

func prepareNodeCandidateListMessage() *NodeCandidateListMsg {
	nodeCandidates := make(map[common.Address]bool)
	epoch := uint64(10)

	addresses := snapshot.GetRandomAddresses(10)

	for _, v := range addresses {
		nodeCandidates[v] = true
	}

	return NewNodeCandidateListMsg(nodeCandidates, epoch)
}

func prepareMinersMessage() *MinersMsg {
	miners := snapshot.GetRandomAddresses(10)
	epoch := uint64(10)

	return NewMinersMsg(epoch, miners, snapshot.Hash)
}

func getNewPeer(peerID enode.ID, service PingPeerService, handlers Handlers, buffer p2p.MsgReadWriter, hook protocols.Hook) *Peer {
	messageBuffer := buffer

	spec := Spec
	if hook != nil {
		spec.Hook = hook
	}

	p := int64(0)
	return &Peer{
		Peer: protocols.NewPeer(
			p2p.NewPeer(peerID, "myPeer", nil),
			messageBuffer,
			spec,
		),
		address: common.Address{},
		peerState: &peerState{
			lastPing: &p,
		},
		PingPeer:      service,
		quorumHandler: handlers,
	}
}

func nodeToString(node *enode.Node) string {
	enc, err := rlp.EncodeToBytes(node.Record())
	if err != nil {
		panic("Error during record encoding")
	}

	b64 := base64.RawURLEncoding.EncodeToString(enc)
	return "enr:" + b64
}
