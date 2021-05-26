package ping

import (
	"fmt"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/simcord-llc/bitbon-system-blockchain/p2p"
	quorum "github.com/simcord-llc/bitbon-system-blockchain/quorum/mock"
	"github.com/simcord-llc/bitbon-system-blockchain/quorum/snapshot"
	"github.com/stretchr/testify/assert"
)

func TestProtocols(t *testing.T) {
	protocol := NewProtocol(time.Second)

	expected := []p2p.Protocol{
		{
			Name:    Spec.Name,
			Version: Spec.Version,
			Length:  Spec.Length(),
			Run:     protocol.runProtocol,
		},
	}
	actual := protocol.Protocols()

	assert.Equal(t, expected[0].Name, actual[0].Name, "Call result differs from expected")
	assert.Equal(t, expected[0].Version, actual[0].Version, "Call result differs from expected")
	assert.Equal(t, expected[0].Length, actual[0].Length, "Call result differs from expected")
}

func TestGetState(t *testing.T) {
	t.Run("PositiveStateNotPresent", func(t *testing.T) {
		protocol := NewProtocol(time.Second)

		peerID := snapshot.GetRandomEnodeID()
		state, found := protocol.GetState(peerID)

		assert.NotEmpty(t, state, "Expected non-empty state")
		assert.False(t, found, "Expected state to be missing")
	})

	t.Run("PositiveStateIsPresent", func(t *testing.T) {
		protocol := NewProtocol(time.Second)

		peerID := snapshot.GetRandomEnodeID()
		protocol.GetState(peerID)

		state, found := protocol.GetState(peerID)

		assert.NotEmpty(t, state, "Expected non-empty state")
		assert.True(t, found, "Expected state to be present")
	})
}

func TestRemoveState(t *testing.T) {
	t.Run("PositiveStateNotPresent", func(t *testing.T) {
		protocol := NewProtocol(time.Second)

		peerID := snapshot.GetRandomEnodeID()
		protocol.RemoveState(peerID)

		assert.Empty(t, protocol.states, "Expected non-empty state")
	})

	t.Run("PositiveStateIsNotPresent", func(t *testing.T) {
		protocol := NewProtocol(time.Second)

		peerID := snapshot.GetRandomEnodeID()
		protocol.GetState(peerID)
		protocol.RemoveState(peerID)

		assert.Empty(t, protocol.states, "Expected non-empty state")
	})
}

func TestPerformHandshake(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	t.Run("Positive", func(t *testing.T) {
		mockBuffer := quorum.NewMockMsgReadWriter(mockController)
		protocol, peer := prepareProtocolAndPeer(1103823438081, mockBuffer)

		mockBuffer.EXPECT().WriteMsg(gomock.Any()).Return(nil)
		mockBuffer.EXPECT().ReadMsg().Return(prepareValidMessage(), nil)

		err := protocol.performHandshake(peer)

		assert.NoError(t, err, "Unexpected error")
	})

	t.Run("NegativeSpecVersionMismatch", func(t *testing.T) {
		mockBuffer := quorum.NewMockMsgReadWriter(mockController)
		protocol, peer := prepareProtocolAndPeer(1, mockBuffer)

		mockBuffer.EXPECT().WriteMsg(gomock.Any()).Return(nil)
		mockBuffer.EXPECT().ReadMsg().Return(prepareValidMessage(), nil)

		err := protocol.performHandshake(peer)

		assert.EqualError(t, err, "message handler: (msg code 0): version mismatch 1103823438081 (!= 1)",
			"Expected VersionMismatch error")
	})

	t.Run("NegativeMessageDecodingError", func(t *testing.T) {
		mockBuffer := quorum.NewMockMsgReadWriter(mockController)
		protocol, peer := prepareProtocolAndPeer(1, mockBuffer)

		mockBuffer.EXPECT().WriteMsg(gomock.Any()).Return(nil)
		mockBuffer.EXPECT().ReadMsg().Return(prepareInvalidMessage(), nil)

		err := protocol.performHandshake(peer)

		assert.EqualError(t, err, "invalid message (RLP error): 0 err=invalid message: "+
			"(code 0) (size 7) rlp: expected input list for protocols.msgWithContext", "Expected RLPDecoding error")
	})
}

func TestCheckHandshake(t *testing.T) {
	t.Run("Positive", func(t *testing.T) {
		protocol := NewProtocol(time.Second)

		err := protocol.checkHandshake(&HandshakeMsg{Version: 1})

		assert.NoError(t, err, "Unexpected error")
	})

	t.Run("NegativeVersionMismatch", func(t *testing.T) {
		protocol := NewProtocol(time.Second)

		err := protocol.checkHandshake(&HandshakeMsg{Version: 2})

		assert.EqualError(t, err, "version mismatch 2 (!= 1)", "Expected VersionMismatch error")
	})

	t.Run("NegativeWrongType", func(t *testing.T) {
		protocol := NewProtocol(time.Second)

		err := protocol.checkHandshake("Whoops")

		assert.EqualError(t, err, "wanted (*HandshakeMsg) got (string)", "Expected WrongType error")
	})
}

func TestStart(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	t.Run("Positive", func(t *testing.T) {
		mockBuffer := quorum.NewMockMsgReadWriter(mockController)
		protocol, peer := prepareProtocolAndPeer(1103823438081, mockBuffer)
		state := NewState()

		mockBuffer.EXPECT().WriteMsg(gomock.Any()).Return(nil)
		mockBuffer.EXPECT().ReadMsg().Return(prepareValidMessage(), nil)

		err := protocol.start(peer, state)

		_, ok := <-state.done
		stateErr := state.err

		assert.NoError(t, err, "Unexpected error")
		assert.Empty(t, stateErr, "Expected state error to be empty")
		assert.False(t, ok, "Expected channel to be closed")
	})

	t.Run("NegativeVersionMismatch", func(t *testing.T) {
		mockBuffer := quorum.NewMockMsgReadWriter(mockController)

		protocol, peer := prepareProtocolAndPeer(1, mockBuffer)
		state := NewState()

		mockBuffer.EXPECT().WriteMsg(gomock.Any()).Return(nil)
		mockBuffer.EXPECT().ReadMsg().Return(prepareValidMessage(), nil)

		err := protocol.start(peer, state)

		_, ok := <-state.done
		stateErr := state.err

		assert.EqualError(t, err,
			fmt.Sprintf("%s: handshake failed with remote peer %s: message handler: "+
				"(msg code 0): version mismatch 1103823438081 (!= 1)", SpecString(Spec), peer.ID().String()),
			"Expected VersionMismatch error")
		assert.EqualError(t, err, stateErr.Error(), "State has a different error")
		assert.False(t, ok, "Expected channel to be closed")
	})

	t.Run("NegativeMessageDecodingError", func(t *testing.T) {
		mockBuffer := quorum.NewMockMsgReadWriter(mockController)

		protocol, peer := prepareProtocolAndPeer(1, mockBuffer)
		state := NewState()

		mockBuffer.EXPECT().WriteMsg(gomock.Any()).Return(nil)
		mockBuffer.EXPECT().ReadMsg().Return(prepareInvalidMessage(), nil)

		err := protocol.start(peer, state)

		_, ok := <-state.done
		stateErr := state.err

		assert.EqualError(t, err,
			fmt.Sprintf("%s: handshake failed with remote peer %s: "+
				"invalid message (RLP error): 0 err=invalid message: (code 0) (size 7) rlp: "+
				"expected input list for protocols.msgWithContext", SpecString(Spec), peer.ID().String()),
			"Expected RLPDecode error")
		assert.EqualError(t, err, stateErr.Error(), "State has a different error")
		assert.False(t, ok, "Expected channel to be closed")
	})
}

func TestRunProtocol(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	t.Run("NegativeChannelAlreadyInitialized", func(t *testing.T) {
		protocol, peer := prepareProtocolAndPeer(1103823438081, nil)
		protocol.addPeer(peer)

		state, _ := protocol.GetState(peer.ID())
		state.init = make(chan bool, 1)
		state.init <- false

		err := protocol.runProtocol(peer.Peer.Peer, nil)

		assert.EqualError(t, err,
			fmt.Sprintf("%s: handshake already started on peer %s", SpecString(Spec), peer.ID().String()))
	})

	t.Run("NegativeVersionMismatch", func(t *testing.T) {
		mockBuffer := quorum.NewMockMsgReadWriter(mockController)
		protocol, peer := prepareProtocolAndPeer(1, mockBuffer)
		state, _ := protocol.GetState(peer.ID())

		mockBuffer.EXPECT().WriteMsg(gomock.Any()).Return(nil)
		mockBuffer.EXPECT().ReadMsg().Return(prepareValidMessage(), nil)

		err := protocol.runProtocol(peer.Peer.Peer, mockBuffer)

		_, ok := <-state.done

		assert.EqualError(t, err,
			fmt.Sprintf("%s: handshake failed with remote peer %s: message handler: "+
				"(msg code 0): version mismatch 1103823438081 (!= 1)", SpecString(Spec), peer.ID().String()),
			"Expected VersionMismatch error")
		assert.EqualError(t, err, state.Error().Error(), "State has a different error")
		assert.False(t, ok, "Expected channel to be closed")
	})

	t.Run("NegativeMessageDecodingError", func(t *testing.T) {
		mockBuffer := quorum.NewMockMsgReadWriter(mockController)
		protocol, peer := prepareProtocolAndPeer(1, mockBuffer)
		state, _ := protocol.GetState(peer.ID())

		mockBuffer.EXPECT().WriteMsg(gomock.Any()).Return(nil)
		mockBuffer.EXPECT().ReadMsg().Return(prepareInvalidMessage(), nil)

		err := protocol.runProtocol(peer.Peer.Peer, mockBuffer)

		_, ok := <-state.done

		assert.EqualError(t, err,
			fmt.Sprintf("%s: handshake failed with remote peer %s: "+
				"invalid message (RLP error): 0 err=invalid message: (code 0) (size 7) rlp: "+
				"expected input list for protocols.msgWithContext", SpecString(Spec), peer.ID().String()),
			"Expected RLPDecode error")
		assert.EqualError(t, err, state.Error().Error(), "State has a different error")
		assert.False(t, ok, "Expected channel to be closed")
	})
}
