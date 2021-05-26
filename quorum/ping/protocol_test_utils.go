package ping

import (
	"time"

	"github.com/simcord-llc/bitbon-system-blockchain/p2p"
	"github.com/simcord-llc/bitbon-system-blockchain/p2p/protocols"
	"github.com/simcord-llc/bitbon-system-blockchain/quorum/snapshot"
	"github.com/simcord-llc/bitbon-system-blockchain/rlp"
)

func prepareProtocolAndPeer(specVersion uint, buffer p2p.MsgReadWriter) (*Protocol, *Peer) {
	Spec.Version = specVersion
	return NewProtocol(time.Second),
		NewPeer(protocols.NewPeer(p2p.NewPeer(snapshot.GetRandomEnodeID(), "newPeer", nil), buffer, Spec))
}

func prepareValidMessage() p2p.Msg {
	size, payload, _ := rlp.EncodeToReader([]interface{}{"", []byte{0xC7, 0x86, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01}})
	return p2p.Msg{
		Size:    uint32(size),
		Payload: payload,
	}
}

func prepareInvalidMessage() p2p.Msg {
	size, payload, _ := rlp.EncodeToReader("Whoops")
	return p2p.Msg{
		Size:    uint32(size),
		Payload: payload,
	}
}
