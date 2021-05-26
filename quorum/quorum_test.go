package quorum

import (
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/simcord-llc/bitbon-system-blockchain/p2p"
	"github.com/simcord-llc/bitbon-system-blockchain/p2p/enode"
	"github.com/simcord-llc/bitbon-system-blockchain/quorum/snapshot"
	"github.com/simcord-llc/bitbon-system-blockchain/rpc"
	"github.com/stretchr/testify/assert"

	"github.com/davecgh/go-spew/spew"

	"github.com/simcord-llc/bitbon-system-blockchain/log"
)

func TestIntegration(t *testing.T) {
	setupLogger()

	srv1, quorumService1 := startTestServer(t, node1Params)
	srv2, quorumService2 := startTestServer(t, node2Params)
	srv3, quorumService3 := startTestServer(t, node3Params)

	defer srv1.Stop()
	defer srv2.Stop()
	defer srv3.Stop()

	time.Sleep(10 * time.Second)

	log.Info("------------------------------------srv1.PeersInfo()------------------")
	spew.Dump(srv1.PeersInfo())
	log.Info("------------------------------------srv2.PeersInfo()------------------")
	spew.Dump(srv2.PeersInfo())
	log.Info("------------------------------------srv3.PeersInfo()------------------")
	spew.Dump(srv3.PeersInfo())

	log.Info("------------------------------------quorum.NewAPI(quorumService1).PeerList(context.TODO())------------------")
	spew.Dump(NewAPI(quorumService1).PeerList(context.TODO()))
	log.Info("------------------------------------quorum.NewAPI(quorumService2).PeerList(context.TODO())------------------")
	spew.Dump(NewAPI(quorumService2).PeerList(context.TODO()))
	log.Info("------------------------------------quorum.NewAPI(quorumService3).PeerList(context.TODO())------------------")
	spew.Dump(NewAPI(quorumService3).PeerList(context.TODO()))

	time.Sleep(20 * time.Second)
}

func TestAPIs(t *testing.T) {
	srv, quorum := startTestServer(t, node1Params)
	defer srv.Stop()

	expected := []rpc.API{
		{
			Namespace: ProtocolName,
			Version:   ProtocolVersionStr,
			Service:   NewAPI(quorum),
			Public:    true,
		},
	}

	assert.Equal(t, expected, quorum.APIs(), "Result differs from expected")
}

func TestProtocols(t *testing.T) {
	srv, quorum := startTestServer(t, node1Params)
	defer srv.Stop()

	expected := quorum.ping.Protocols()
	expected = append(expected, p2p.Protocol{
		Name:    Spec.Name,
		Version: Spec.Version,
		Length:  Spec.Length(),
		Run:     quorum.runProtocol,
	})

	actual := quorum.Protocols()

	for i, v := range expected {
		assert.Equal(t, v.Name, actual[i].Name, "Name differs from expected")
		assert.Equal(t, v.Version, actual[i].Version, "Version differs from expected")
		assert.Equal(t, v.Length, actual[i].Length, "Length differs from expected")
	}
}

func TestStop(t *testing.T) {
	srv, quorum := startTestServer(t, node2Params)
	defer srv.Stop()

	err := quorum.Stop()

	_, ok := <-quorum.done

	assert.NoError(t, err, "Unexpected error")
	assert.False(t, ok, "Expected channel to be closed")
	assert.True(t, quorum.Stopped(), "Expected quorum to be stopped")
}

func TestStopped(t *testing.T) {
	t.Run("PositiveStopped", func(t *testing.T) {
		srv, quorum := startTestServer(t, node1Params)
		defer srv.Stop()

		assert.False(t, quorum.Stopped(), "Expected quorum to be running")
	})

	t.Run("PositiveRunning", func(t *testing.T) {
		srv, quorum := startTestServer(t, node2Params)
		defer srv.Stop()

		err := quorum.Stop()
		if err != nil {
			panic("Error while stopping quorum")
		}

		assert.True(t, quorum.Stopped(), "Expected quorum to be stopped")
	})
}

func TestMarkPeersGathered(t *testing.T) {
	t.Run("PositiveSetTrue", func(t *testing.T) {
		srv, quorum := startTestServer(t, node2Params)
		defer srv.Stop()

		quorum.markPeersGathered(true)
		assert.Equal(t, int32(1), *quorum.peersGathered, "Wrong value for peersGathered field")
	})

	t.Run("PositiveSetFalse", func(t *testing.T) {
		srv, quorum := startTestServer(t, node2Params)
		defer srv.Stop()

		quorum.markPeersGathered(false)
		assert.Equal(t, int32(0), *quorum.peersGathered, "Wrong value for peersGathered field")
	})
}

func TestPrepared(t *testing.T) {
	t.Run("PositivePeersGathered", func(t *testing.T) {
		srv, quorum := startTestServer(t, node2Params)
		defer srv.Stop()

		quorum.markPeersGathered(true)
		assert.True(t, quorum.Prepared(), "Expected quorum to be prepared")
	})

	t.Run("PositivePeersNotGathered", func(t *testing.T) {
		srv, quorum := startTestServer(t, node2Params)
		defer srv.Stop()

		assert.False(t, quorum.Prepared(), "Expected quorum to be unprepared")
	})
}

func TestSetAllNodes(t *testing.T) {
	t.Run("PositiveSetList", func(t *testing.T) {
		srv, quorum := startTestServer(t, node2Params)
		defer srv.Stop()

		nodeList := []*enode.Node{
			getNewNode(),
			getNewNode(),
			getNewNode(),
		}

		expected := nodeListToMap(nodeList)
		for _, v := range quorum.allNodes {
			expected[v.ID()] = v
		}

		quorum.setAllNodes(nodeList)

		assert.Equal(t, expected, quorum.allNodes, "Node list in quorum differs from expected")
	})

	t.Run("PositiveSetListWithSelfNode", func(t *testing.T) {
		srv, quorum := startTestServer(t, node2Params)
		defer srv.Stop()

		nodeList := []*enode.Node{
			getNewNode(),
			getNewNode(),
			quorum.srv.Self(),
		}

		expected := nodeListToMap(nodeList)
		for _, v := range quorum.allNodes {
			expected[v.ID()] = v
		}

		delete(expected, quorum.srv.Self().ID())

		quorum.setAllNodes(nodeList)

		assert.Equal(t, expected, quorum.allNodes, "Node list in quorum differs from expected")
	})

	t.Run("PositiveSetListSelfNodeAlreadySet", func(t *testing.T) {
		srv, quorum := startTestServer(t, node2Params)
		defer srv.Stop()

		var nodeList []*enode.Node

		expected := nodeListToMap(nodeList)
		for _, v := range quorum.allNodes {
			expected[v.ID()] = v
		}

		quorum.allNodes[srv.Self().ID()] = getNewNode()
		quorum.setAllNodes(nodeList)

		assert.Equal(t, expected, quorum.allNodes, "Node list in quorum differs from expected")
	})
}

func TestHasMajority(t *testing.T) {
	t.Run("PositiveHasMajority", func(t *testing.T) {
		srv, quorum := startTestServer(t, node2Params)
		defer srv.Stop()

		assert.True(t, quorum.hasMajority(2, false), "Expected quorum to have majority with 2 nodes")
	})

	t.Run("PositiveHasMajorityWithSelf", func(t *testing.T) {
		srv, quorum := startTestServer(t, node2Params)
		defer srv.Stop()

		assert.True(t, quorum.hasMajority(2, true), "Expected quorum to have majority with 3 nodes")
	})

	t.Run("PositiveNoMajority", func(t *testing.T) {
		srv, quorum := startTestServer(t, node2Params)
		defer srv.Stop()

		assert.False(t, quorum.hasMajority(1, false), "Expected quorum to have no majority with 1 node")
	})

	t.Run("PositiveNoMajorityWithSelf", func(t *testing.T) {
		srv, quorum := startTestServer(t, node2Params)
		defer srv.Stop()

		assert.False(t, quorum.hasMajority(0, true), "Expected quorum to have no majority with 1 node")
	})
}

func TestMinerListAtEpoch(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	t.Run("PositiveMinerListIsPresent", func(t *testing.T) {
		srv, quorum := startTestServer(t, node2Params)
		defer srv.Stop()

		mockSnapshotSystem := NewMockSnapshotService(mockController)
		mockContractService := NewMockContractService(mockController)
		quorum.snapshotSystem = mockSnapshotSystem
		quorum.SetContractService(mockContractService)

		minerList := snapshot.GetRandomAddresses(numOfAddresses)

		mockSnapshotSystem.EXPECT().MinerListAtEpoch(epochNumber).Return(minerList)
		mockContractService.EXPECT().GetCurrentDistribution(context.TODO()).Return(minerListToMap(minerList, 10), nil)

		assert.ElementsMatch(t, minerList, quorum.MinerListAtEpoch(epochNumber), "Resulting list differs from expected")
	})

	t.Run("PositiveMinerListIsPartiallyPresent", func(t *testing.T) {
		srv, quorum := startTestServer(t, node2Params)
		defer srv.Stop()

		mockSnapshotSystem := NewMockSnapshotService(mockController)
		mockContractService := NewMockContractService(mockController)
		quorum.snapshotSystem = mockSnapshotSystem
		quorum.SetContractService(mockContractService)

		minerList := snapshot.GetRandomAddresses(numOfAddresses)

		mockSnapshotSystem.EXPECT().MinerListAtEpoch(epochNumber).Return(minerList)
		mockContractService.EXPECT().GetCurrentDistribution(context.TODO()).
			Return(minerListToMap(minerList[:numOfAddresses-5], epochNumber), nil)

		assert.ElementsMatch(t, minerList[:numOfAddresses-5], quorum.MinerListAtEpoch(epochNumber), "Resulting list differs from expected")
	})

	t.Run("PositiveMinerListIsEmpty", func(t *testing.T) {
		srv, quorum := startTestServer(t, node2Params)
		defer srv.Stop()

		mockSnapshotSystem := NewMockSnapshotService(mockController)
		quorum.snapshotSystem = mockSnapshotSystem

		mockSnapshotSystem.EXPECT().MinerListAtEpoch(epochNumber).Return(nil)

		assert.Empty(t, quorum.MinerListAtEpoch(epochNumber), "Expected resulting list to be empty")
	})

	t.Run("PositiveContractServiceNotSet", func(t *testing.T) {
		srv, quorum := startTestServer(t, node2Params)
		defer srv.Stop()

		mockSnapshotSystem := NewMockSnapshotService(mockController)
		quorum.snapshotSystem = mockSnapshotSystem

		minerList := snapshot.GetRandomAddresses(numOfAddresses)

		mockSnapshotSystem.EXPECT().MinerListAtEpoch(epochNumber).Return(minerList)

		assert.ElementsMatch(t, minerList, quorum.MinerListAtEpoch(epochNumber), "Resulting list differs from expected")
	})

	t.Run("PositiveContractReturnedError", func(t *testing.T) {
		srv, quorum := startTestServer(t, node2Params)
		defer srv.Stop()

		mockSnapshotSystem := NewMockSnapshotService(mockController)
		mockContractService := NewMockContractService(mockController)
		quorum.snapshotSystem = mockSnapshotSystem
		quorum.SetContractService(mockContractService)

		minerList := snapshot.GetRandomAddresses(numOfAddresses)

		mockSnapshotSystem.EXPECT().MinerListAtEpoch(epochNumber).Return(minerList)
		mockContractService.EXPECT().GetCurrentDistribution(context.TODO()).Return(nil, ErrOnGettingDistribution)

		assert.ElementsMatch(t, minerList, quorum.MinerListAtEpoch(epochNumber), "Resulting list differs from expected")
	})

	t.Run("PositiveContractReturnedEmptyList", func(t *testing.T) {
		srv, quorum := startTestServer(t, node2Params)
		defer srv.Stop()

		mockSnapshotSystem := NewMockSnapshotService(mockController)
		mockContractService := NewMockContractService(mockController)
		quorum.snapshotSystem = mockSnapshotSystem
		quorum.SetContractService(mockContractService)

		minerList := snapshot.GetRandomAddresses(numOfAddresses)

		mockSnapshotSystem.EXPECT().MinerListAtEpoch(epochNumber).Return(minerList)
		mockContractService.EXPECT().GetCurrentDistribution(context.TODO()).Return(nil, nil)

		assert.ElementsMatch(t, minerList, quorum.MinerListAtEpoch(epochNumber), "Resulting list differs from expected")
	})

	t.Run("PositiveDistributionDiffersFromSnapshot", func(t *testing.T) {
		srv, quorum := startTestServer(t, node2Params)
		defer srv.Stop()

		mockSnapshotSystem := NewMockSnapshotService(mockController)
		mockContractService := NewMockContractService(mockController)
		quorum.snapshotSystem = mockSnapshotSystem
		quorum.SetContractService(mockContractService)

		minerList := snapshot.GetRandomAddresses(numOfAddresses)

		mockSnapshotSystem.EXPECT().MinerListAtEpoch(epochNumber).Return(minerList)
		mockContractService.EXPECT().GetCurrentDistribution(context.TODO()).
			Return(minerListToMap(snapshot.GetRandomAddresses(numOfAddresses), epochNumber), nil)

		assert.Empty(t, quorum.MinerListAtEpoch(epochNumber), "Expected resulting list to be empty")
	})
}

func TestCreateQuorumSnapshot(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	t.Run("Positive", func(t *testing.T) {
		srv, quorum := startTestServer(t, node2Params)
		defer srv.Stop()

		mockPeerListService := NewMockPeerListService(mockController)
		quorum.peerList = mockPeerListService

		peerList := getNewPeerList(quorum.hasMajority)
		peerList.AddPeer(createNewPeerStruct(getRandomEnodeID(), "myPeer", true, true))
		peerList.AddPeer(createNewPeerStruct(getRandomEnodeID(), "myPeer2", true, true))

		snapshotNodes := prepareExpectedNodesForPeerList(peerList)
		mockPeerListService.EXPECT().GetSnapshotNodes().Return(snapshotNodes, nil)
		mockPeerListService.EXPECT().PingList().Return(nil).AnyTimes()

		expected := snapshot.
			NewSnapshot(epochNumber, srv.Self().ID(), snapshotNodes, quorum.hasMajority, logger)

		actual, err := quorum.CreateQuorumSnapshot(epochNumber, logger)

		assert.NoError(t, err, "Unexpected error")
		assert.Equal(t, expected.ToDTO(), actual.ToDTO(), "Expected snapshot to be not empty")
	})

	t.Run("NegativeErrorOnGettingSnapshot", func(t *testing.T) {
		srv, quorum := startTestServer(t, node2Params)
		defer srv.Stop()

		mockPeerListService := NewMockPeerListService(mockController)
		quorum.peerList = mockPeerListService

		mockPeerListService.EXPECT().GetSnapshotNodes().Return(nil, ErrGettingSnapshot)
		mockPeerListService.EXPECT().PingList().Return(nil).AnyTimes()

		quorumSnapshot, err := quorum.CreateQuorumSnapshot(epochNumber, logger)

		assert.EqualError(t, err, ErrGettingSnapshot.Error(), "Expected GettingSnapshot error")
		assert.Empty(t, quorumSnapshot, "Expected snapshot to be not empty")
	})
}

func TestMinerListAtRecentEpoch(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	t.Run("PositiveMinerListIsPresent", func(t *testing.T) {
		srv, quorum := startTestServer(t, node2Params)
		defer srv.Stop()

		mockSnapshotSystem := NewMockSnapshotService(mockController)
		quorum.snapshotSystem = mockSnapshotSystem

		minerList := snapshot.GetRandomAddresses(numOfAddresses)

		mockSnapshotSystem.EXPECT().RecentEpoch().Return(epochNumber)
		mockSnapshotSystem.EXPECT().MinerListAtEpoch(epochNumber).Return(minerList)

		assert.ElementsMatch(t, minerList, quorum.MinerListAtRecentEpoch(), "Resulting list differs from expected")
	})

	t.Run("PositiveMinerListIsPresentAtPreviousEpoch", func(t *testing.T) {
		srv, quorum := startTestServer(t, node2Params)
		defer srv.Stop()

		mockSnapshotSystem := NewMockSnapshotService(mockController)
		quorum.snapshotSystem = mockSnapshotSystem

		minerList := snapshot.GetRandomAddresses(numOfAddresses)

		mockSnapshotSystem.EXPECT().RecentEpoch().Return(epochNumber)
		mockSnapshotSystem.EXPECT().MinerListAtEpoch(epochNumber).Return(nil)
		mockSnapshotSystem.EXPECT().MinerListAtEpoch(epochNumber - 1).Return(minerList)

		assert.ElementsMatch(t, minerList, quorum.MinerListAtRecentEpoch(), "Resulting list differs from expected")
	})
}
