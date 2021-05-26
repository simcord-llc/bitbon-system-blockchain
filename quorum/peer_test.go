package quorum

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestPeerStateToDTO(t *testing.T) {
	ping := int64(10)
	peerState := &peerState{
		lastPing:         &ping,
		inPingList:       true,
		isMinerCandidate: false,
	}

	expected := &PeerStateDTO{
		LastPing:         10,
		LastPingStr:      "10ns",
		InPingList:       true,
		IsMinerCandidate: false,
	}

	assert.Equal(t, expected, peerState.ToDTO(), "Mapped DTO differs from expected")
}

func TestPeerToDTO(t *testing.T) {
	t.Run("PositiveDefaultPeerState", func(t *testing.T) {
		peerID := getRandomEnodeID()
		lastPing := time.Duration(0)
		peer := getNewPeer(peerID, nil, nil, nil, nil)

		expected := &PeerDTO{
			ID:    peerID,
			Enode: nodeToString(peer.Node()),
			PeerStateDTO: &PeerStateDTO{
				LastPing:    lastPing,
				LastPingStr: lastPing.String(),
			},
		}

		assert.Equal(t, expected, peer.ToDTO(), "Mapped DTO differs from expected")
	})

	t.Run("PositiveNotDefaultPeerState", func(t *testing.T) {
		peerID := getRandomEnodeID()
		state := getNewPeerState(10)
		lastPing := time.Duration(10)
		peer := getNewPeer(peerID, nil, nil, nil, nil)
		peer.setState(state)

		expected := &PeerDTO{
			ID:    peerID,
			Enode: nodeToString(peer.Node()),
			PeerStateDTO: &PeerStateDTO{
				LastPing:         lastPing,
				LastPingStr:      lastPing.String(),
				InPingList:       false,
				IsMinerCandidate: true,
			},
		}

		assert.Equal(t, expected, peer.ToDTO(), "Mapped DTO differs from expected")
	})
}

func TestHandleMessage(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	t.Run("PositiveVoteMinerAddMessage", func(t *testing.T) {
		mockPingPeerService, mockHandlers, mockBuffer, _ := prepareMockServices(mockController)
		peer := getNewPeer(getRandomEnodeID(), mockPingPeerService, mockHandlers, mockBuffer, nil)

		waitGroup := sync.WaitGroup{}
		waitGroup.Add(1)
		defer waitGroup.Wait()

		mockHandlers.EXPECT().HandleVoteMinerAddMsg(peer).DoAndReturn(func(peer *Peer) {
			waitGroup.Done()
		})

		err := peer.HandleMsg(context.TODO(), NewVoteMinerAddMsg())

		assert.NoError(t, err, "Unexpected error")
	})

	t.Run("PositiveVoteMinerDelMessage", func(t *testing.T) {
		mockPingPeerService, mockHandlers, mockBuffer, _ := prepareMockServices(mockController)
		peer := getNewPeer(getRandomEnodeID(), mockPingPeerService, mockHandlers, mockBuffer, nil)

		waitGroup := sync.WaitGroup{}
		waitGroup.Add(1)
		defer waitGroup.Wait()

		mockHandlers.EXPECT().HandleVoteMinerDelMsg(peer).DoAndReturn(func(peer *Peer) {
			waitGroup.Done()
		})

		err := peer.HandleMsg(context.TODO(), NewVoteMinerDelMsg())

		assert.NoError(t, err, "Unexpected error")
	})

	t.Run("PositiveNodeCandidateListMessage", func(t *testing.T) {
		mockPingPeerService, mockHandlers, mockBuffer, _ := prepareMockServices(mockController)
		peer := getNewPeer(getRandomEnodeID(), mockPingPeerService, mockHandlers, mockBuffer, nil)

		waitGroup := sync.WaitGroup{}
		waitGroup.Add(1)
		defer waitGroup.Wait()

		mockHandlers.EXPECT().
			HandleNodeCandidateListMsg(gomock.Any(), peer).DoAndReturn(func(msg interface{}, peer *Peer) {
			waitGroup.Done()
		})

		err := peer.HandleMsg(context.TODO(), prepareNodeCandidateListMessage())

		assert.NoError(t, err, "Unexpected error")
	})

	t.Run("PositiveMinersMessage", func(t *testing.T) {
		mockPingPeerService, mockHandlers, mockBuffer, _ := prepareMockServices(mockController)
		peer := getNewPeer(getRandomEnodeID(), mockPingPeerService, mockHandlers, mockBuffer, nil)

		waitGroup := sync.WaitGroup{}
		waitGroup.Add(1)
		defer waitGroup.Wait()

		mockHandlers.EXPECT().HandleMinersMsg(gomock.Any(), peer).DoAndReturn(func(msg interface{}, peer *Peer) {
			waitGroup.Done()
		})

		err := peer.HandleMsg(context.TODO(), prepareMinersMessage())

		assert.NoError(t, err, "Unexpected error")
	})

	t.Run("NegativeUnknownMessageType", func(t *testing.T) {
		mockPingPeerService, mockHandlers, mockBuffer, _ := prepareMockServices(mockController)
		peer := getNewPeer(getRandomEnodeID(), mockPingPeerService, mockHandlers, mockBuffer, nil)

		msg := "Wrong type of message"
		err := peer.HandleMsg(context.TODO(), msg)

		assert.EqualError(t, err,
			fmt.Sprintf("unknown message type: %T", msg), "Expected UnknownMessageType error")
	})
}

func TestSendVoteMinerAddMessage(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	t.Run("PositiveNoError", func(t *testing.T) {
		mockPingPeerService, mockHandlers, mockBuffer, _ := prepareMockServices(mockController)
		peer := getNewPeer(getRandomEnodeID(), mockPingPeerService, mockHandlers, mockBuffer, nil)

		mockBuffer.EXPECT().WriteMsg(gomock.Any()).Return(nil)

		err := peer.SendVoteMinerAddMsg(context.TODO())

		assert.NoError(t, err, "Unexpected error")
	})

	t.Run("NegativeErrorOccurred", func(t *testing.T) {
		mockPingPeerService, mockHandlers, mockBuffer, _ := prepareMockServices(mockController)
		peer := getNewPeer(getRandomEnodeID(), mockPingPeerService, mockHandlers, mockBuffer, nil)

		mockBuffer.EXPECT().WriteMsg(gomock.Any()).Return(ErrOnWriteMessage)

		err := peer.SendVoteMinerAddMsg(context.TODO())

		assert.EqualError(t, err, ErrOnWriteMessage.Error(), "Expected WriteMessage error")
	})
}

func TestSendVoteMinerDelMessage(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	t.Run("PositiveNoError", func(t *testing.T) {
		mockPingPeerService, mockHandlers, mockBuffer, _ := prepareMockServices(mockController)
		peer := getNewPeer(getRandomEnodeID(), mockPingPeerService, mockHandlers, mockBuffer, nil)

		mockBuffer.EXPECT().WriteMsg(gomock.Any()).Return(nil)

		err := peer.SendVoteMinerDelMsg(context.TODO())

		assert.NoError(t, err, "Unexpected error")
	})

	t.Run("NegativeErrorOccurred", func(t *testing.T) {
		mockPingPeerService, mockHandlers, mockBuffer, _ := prepareMockServices(mockController)
		peer := getNewPeer(getRandomEnodeID(), mockPingPeerService, mockHandlers, mockBuffer, nil)

		mockBuffer.EXPECT().WriteMsg(gomock.Any()).Return(ErrOnWriteMessage)

		err := peer.SendVoteMinerDelMsg(context.TODO())

		assert.EqualError(t, err, ErrOnWriteMessage.Error(), "Expected WriteMessage error")
	})
}

func TestSendNodeCandidateListMessage(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	t.Run("PositiveNoError", func(t *testing.T) {
		mockPingPeerService, mockHandlers, mockBuffer, _ := prepareMockServices(mockController)
		peer := getNewPeer(getRandomEnodeID(), mockPingPeerService, mockHandlers, mockBuffer, nil)

		mockBuffer.EXPECT().WriteMsg(gomock.Any()).Return(nil)

		err := peer.SendNodeCandidateList(context.TODO(), prepareNodeCandidateListMessage())

		assert.NoError(t, err, "Unexpected error")
	})

	t.Run("NegativeErrorOccurred", func(t *testing.T) {
		mockPingPeerService, mockHandlers, mockBuffer, _ := prepareMockServices(mockController)
		peer := getNewPeer(getRandomEnodeID(), mockPingPeerService, mockHandlers, mockBuffer, nil)

		mockBuffer.EXPECT().WriteMsg(gomock.Any()).Return(ErrOnWriteMessage)

		err := peer.SendNodeCandidateList(context.TODO(), prepareNodeCandidateListMessage())

		assert.EqualError(t, err, ErrOnWriteMessage.Error(), "Expected WriteMessage error")
	})
}

func TestSendMinersMessage(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	t.Run("PositiveNoError", func(t *testing.T) {
		mockPingPeerService, mockHandlers, mockBuffer, _ := prepareMockServices(mockController)
		peer := getNewPeer(getRandomEnodeID(), mockPingPeerService, mockHandlers, mockBuffer, nil)

		mockBuffer.EXPECT().WriteMsg(gomock.Any()).Return(nil)

		err := peer.SendMiners(context.TODO(), prepareMinersMessage())

		assert.NoError(t, err, "Unexpected error")
	})

	t.Run("NegativeErrorOccurred", func(t *testing.T) {
		mockPingPeerService, mockHandlers, mockBuffer, _ := prepareMockServices(mockController)
		peer := getNewPeer(getRandomEnodeID(), mockPingPeerService, mockHandlers, mockBuffer, nil)

		mockBuffer.EXPECT().WriteMsg(gomock.Any()).Return(ErrOnWriteMessage)

		err := peer.SendMiners(context.TODO(), prepareMinersMessage())

		assert.EqualError(t, err, ErrOnWriteMessage.Error(), "Expected WriteMessage error")
	})
}

func TestSendVoteMinerAddMessageWithHook(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	t.Run("NegativeErrorOccurredDuringHookValidation", func(t *testing.T) {
		mockPingPeerService, mockHandlers, mockBuffer, mockHook := prepareMockServices(mockController)
		peer := getNewPeer(getRandomEnodeID(), mockPingPeerService, mockHandlers, mockBuffer, mockHook)

		mockHook.EXPECT().
			Validate(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(int64(0), ErrHookValidate)

		err := peer.SendVoteMinerAddMsg(context.TODO())

		assert.EqualError(t, err, ErrHookValidate.Error(), "Expected HookValidate error")
	})

	t.Run("NegativeErrorOccurredDuringHookSending", func(t *testing.T) {
		mockPingPeerService, mockHandlers, mockBuffer, mockHook := prepareMockServices(mockController)
		peer := getNewPeer(getRandomEnodeID(), mockPingPeerService, mockHandlers, mockBuffer, mockHook)

		mockHook.EXPECT().Validate(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(int64(0), nil)
		mockBuffer.EXPECT().WriteMsg(gomock.Any()).Return(ErrOnWriteMessage)

		err := peer.SendVoteMinerAddMsg(context.TODO())

		assert.EqualError(t, err, ErrOnWriteMessage.Error(), "Expected WriteMessage error")
	})

	t.Run("NegativeErrorOccurredDuringHookApplying", func(t *testing.T) {
		mockPingPeerService, mockHandlers, mockBuffer, mockHook := prepareMockServices(mockController)
		peer := getNewPeer(getRandomEnodeID(), mockPingPeerService, mockHandlers, mockBuffer, mockHook)

		mockHook.EXPECT().Validate(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(int64(0), nil)
		mockBuffer.EXPECT().WriteMsg(gomock.Any()).Return(nil)
		mockHook.EXPECT().Apply(gomock.Any(), gomock.Any(), gomock.Any()).Return(ErrHookApply)

		err := peer.SendVoteMinerAddMsg(context.TODO())

		assert.EqualError(t, err, ErrHookApply.Error(), "Expected HookApply error")
	})
}

func TestSendVoteMinerDelMessageWithHook(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	t.Run("NegativeErrorOccurredDuringHookValidation", func(t *testing.T) {
		mockPingPeerService, mockHandlers, mockBuffer, mockHook := prepareMockServices(mockController)
		peer := getNewPeer(getRandomEnodeID(), mockPingPeerService, mockHandlers, mockBuffer, mockHook)

		mockHook.EXPECT().
			Validate(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(int64(0), ErrHookValidate)

		err := peer.SendVoteMinerDelMsg(context.TODO())

		assert.EqualError(t, err, ErrHookValidate.Error(), "Expected HookValidate error")
	})

	t.Run("NegativeErrorOccurredDuringHookApplying", func(t *testing.T) {
		mockPingPeerService, mockHandlers, mockBuffer, mockHook := prepareMockServices(mockController)
		peer := getNewPeer(getRandomEnodeID(), mockPingPeerService, mockHandlers, mockBuffer, mockHook)

		mockHook.EXPECT().Validate(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(int64(0), nil)
		mockBuffer.EXPECT().WriteMsg(gomock.Any()).Return(ErrOnWriteMessage)

		err := peer.SendVoteMinerDelMsg(context.TODO())

		assert.EqualError(t, err, ErrOnWriteMessage.Error(), "Expected WriteMessage error")
	})

	t.Run("NegativeErrorOccurredDuringHookApplying", func(t *testing.T) {
		mockPingPeerService, mockHandlers, mockBuffer, mockHook := prepareMockServices(mockController)
		peer := getNewPeer(getRandomEnodeID(), mockPingPeerService, mockHandlers, mockBuffer, mockHook)

		mockHook.EXPECT().Validate(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(int64(0), nil)
		mockBuffer.EXPECT().WriteMsg(gomock.Any()).Return(nil)
		mockHook.EXPECT().Apply(gomock.Any(), gomock.Any(), gomock.Any()).Return(ErrHookApply)

		err := peer.SendVoteMinerDelMsg(context.TODO())

		assert.EqualError(t, err, ErrHookApply.Error(), "Expected HookApply error")
	})
}

func TestSendNodeCandidateListMessageWithHook(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	t.Run("NegativeErrorOccurredDuringHookValidation", func(t *testing.T) {
		mockPingPeerService, mockHandlers, mockBuffer, mockHook := prepareMockServices(mockController)
		peer := getNewPeer(getRandomEnodeID(), mockPingPeerService, mockHandlers, mockBuffer, mockHook)

		mockHook.EXPECT().
			Validate(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(int64(0), ErrHookValidate)

		err := peer.SendNodeCandidateList(context.TODO(), prepareNodeCandidateListMessage())

		assert.EqualError(t, err, ErrHookValidate.Error(), "Expected HookValidate error")
	})

	t.Run("NegativeErrorOccurredDuringHookSending", func(t *testing.T) {
		mockPingPeerService, mockHandlers, mockBuffer, mockHook := prepareMockServices(mockController)
		peer := getNewPeer(getRandomEnodeID(), mockPingPeerService, mockHandlers, mockBuffer, mockHook)

		mockHook.EXPECT().Validate(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(int64(0), nil)
		mockBuffer.EXPECT().WriteMsg(gomock.Any()).Return(ErrOnWriteMessage)

		err := peer.SendNodeCandidateList(context.TODO(), prepareNodeCandidateListMessage())

		assert.EqualError(t, err, ErrOnWriteMessage.Error(), "Expected WriteMessage error")
	})

	t.Run("NegativeErrorOccurredDuringHookApplying", func(t *testing.T) {
		mockPingPeerService, mockHandlers, mockBuffer, mockHook := prepareMockServices(mockController)
		peer := getNewPeer(getRandomEnodeID(), mockPingPeerService, mockHandlers, mockBuffer, mockHook)

		mockHook.EXPECT().Validate(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(int64(0), nil)
		mockBuffer.EXPECT().WriteMsg(gomock.Any()).Return(nil)
		mockHook.EXPECT().Apply(gomock.Any(), gomock.Any(), gomock.Any()).Return(ErrHookApply)

		err := peer.SendNodeCandidateList(context.TODO(), prepareNodeCandidateListMessage())

		assert.EqualError(t, err, ErrHookApply.Error(), "Expected HookApply error")
	})
}

func TestSendMinersMessageWithHook(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	t.Run("NegativeErrorOccurredDuringHookValidation", func(t *testing.T) {
		mockPingPeerService, mockHandlers, mockBuffer, mockHook := prepareMockServices(mockController)
		peer := getNewPeer(getRandomEnodeID(), mockPingPeerService, mockHandlers, mockBuffer, mockHook)

		mockHook.EXPECT().
			Validate(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(int64(0), ErrHookValidate)

		err := peer.SendMiners(context.TODO(), prepareMinersMessage())

		assert.EqualError(t, err, ErrHookValidate.Error(), "Expected HookValidate error")
	})

	t.Run("NegativeErrorOccurredDuringHookSending", func(t *testing.T) {
		mockPingPeerService, mockHandlers, mockBuffer, mockHook := prepareMockServices(mockController)
		peer := getNewPeer(getRandomEnodeID(), mockPingPeerService, mockHandlers, mockBuffer, mockHook)

		mockHook.EXPECT().Validate(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(int64(0), nil)
		mockBuffer.EXPECT().WriteMsg(gomock.Any()).Return(ErrOnWriteMessage)

		err := peer.SendMiners(context.TODO(), prepareMinersMessage())

		assert.EqualError(t, err, ErrOnWriteMessage.Error(), "Expected WriteMessage error")
	})

	t.Run("NegativeErrorOccurredDuringHookApplying", func(t *testing.T) {
		mockPingPeerService, mockHandlers, mockBuffer, mockHook := prepareMockServices(mockController)
		peer := getNewPeer(getRandomEnodeID(), mockPingPeerService, mockHandlers, mockBuffer, mockHook)

		mockHook.EXPECT().Validate(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(int64(0), nil)
		mockBuffer.EXPECT().WriteMsg(gomock.Any()).Return(nil)
		mockHook.EXPECT().Apply(gomock.Any(), gomock.Any(), gomock.Any()).Return(ErrHookApply)

		err := peer.SendMiners(context.TODO(), prepareMinersMessage())

		assert.EqualError(t, err, ErrHookApply.Error(), "Expected HookApply error")
	})
}
