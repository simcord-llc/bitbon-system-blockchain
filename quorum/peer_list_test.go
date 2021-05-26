package quorum

import (
	"testing"

	"github.com/simcord-llc/bitbon-system-blockchain/p2p/enode"
	"github.com/stretchr/testify/assert"
)

func TestPeerListToDTO(t *testing.T) {
	t.Run("PositiveWithoutPeersOrRecentPeers", func(t *testing.T) {
		peerList := getNewPeerList(hasMajorityTrue)

		expected := &PeerListDTO{
			SelfID:               peerList.self.ID(),
			SelfEnode:            peerList.self.String(),
			SelfIsMinerCandidate: peerList.selfState.isMinerCandidate,
			Peers:                make([]*PeerDTO, 0, len(peerList.peers)),
			RecentPeers:          make(map[enode.ID]*PeerStateDTO, len(peerList.recentPeers)),
		}

		assert.Equal(t, expected, peerList.ToDTO(), "Mapped DTO differs from expected")
	})

	t.Run("PositiveWithPeers", func(t *testing.T) {
		peerList := getNewPeerList(hasMajorityTrue)
		peer := createNewPeerStruct(getRandomEnodeID(), "myPeer", true, true)
		peerList.AddPeer(peer)

		expected := &PeerListDTO{
			SelfID:               peerList.self.ID(),
			SelfEnode:            peerList.self.String(),
			SelfIsMinerCandidate: peerList.selfState.isMinerCandidate,
			Peers:                []*PeerDTO{peer.ToDTO()},
			RecentPeers:          make(map[enode.ID]*PeerStateDTO, len(peerList.recentPeers)),
		}

		assert.Equal(t, expected, peerList.ToDTO(), "Mapped DTO differs from expected")
	})

	t.Run("PositiveWithPeersAndRecentPeers", func(t *testing.T) {
		peerList := getNewPeerList(hasMajorityTrue)
		peer1 := createNewPeerStruct(getRandomEnodeID(), "myPeer", true, true)
		peer2ID := getRandomEnodeID()
		peer2 := createNewPeerStruct(peer2ID, "myPeer2", true, true)

		peerList.AddPeer(peer1)
		peerList.AddPeer(peer2)
		peerList.RemovePeer(peer2ID)

		expectedRecentPeers := make(map[enode.ID]*PeerStateDTO)
		expectedRecentPeers[peer2ID] = peer2.GetState().ToDTO()
		expected := &PeerListDTO{
			SelfID:               peerList.self.ID(),
			SelfEnode:            peerList.self.String(),
			SelfIsMinerCandidate: peerList.selfState.isMinerCandidate,
			Peers:                []*PeerDTO{peer1.ToDTO()},
			RecentPeers:          expectedRecentPeers,
		}

		assert.Equal(t, expected, peerList.ToDTO(), "Mapped DTO differs from expected")
	})
}

func TestGetPeer(t *testing.T) {
	t.Run("PositivePeerWasAdded", func(t *testing.T) {
		peerList := getNewPeerList(hasMajorityTrue)
		peerID := getRandomEnodeID()
		peer := createNewPeerStruct(peerID, "myPeer", true, true)

		peerList.AddPeer(peer)

		result := peerList.GetPeer(peerID)

		assert.NotEmpty(t, result, "Expected non-empty result")
	})

	t.Run("PositivePeerWasNotAdded", func(t *testing.T) {
		peerList := getNewPeerList(hasMajorityTrue)

		result := peerList.GetPeer(getRandomEnodeID())

		assert.Empty(t, result, "Expected empty result")
	})

	t.Run("PositivePeerWasRemoved", func(t *testing.T) {
		peerList := getNewPeerList(hasMajorityTrue)
		peerID := getRandomEnodeID()
		peer := createNewPeerStruct(peerID, "myPeer", true, true)

		peerList.AddPeer(peer)
		peerList.RemovePeer(peerID)

		result := peerList.GetPeer(peerID)

		assert.Empty(t, result, "Expected empty result")
	})
}

func TestGetRecentPeer(t *testing.T) {
	t.Run("PositiveAfterPeerRemoval", func(t *testing.T) {
		peerList := getNewPeerList(hasMajorityTrue)
		peerID := getRandomEnodeID()
		peer := createNewPeerStruct(peerID, "myPeer", true, true)

		peerList.AddPeer(peer)
		peerList.RemovePeer(peerID)

		result := peerList.GetPeer(peerID)
		expectedRecentPeers := peerList.recentPeers[peerID]

		assert.Empty(t, result, "Expected empty result")
		assert.Equal(t, expectedRecentPeers, peer.peerState, "Expected removed peer to appear in recentPeers list")
	})

	t.Run("PositivePeerWasNotRemoved", func(t *testing.T) {
		peerList := getNewPeerList(hasMajorityTrue)
		peerID := getRandomEnodeID()
		peer := createNewPeerStruct(peerID, "myPeer", true, true)

		peerList.AddPeer(peer)

		assert.Empty(t, peerList.recentPeers, "Did not expect recentPeers to contain elements")
	})
}

func TestPingList(t *testing.T) {
	t.Run("PositivePingListIsEmpty", func(t *testing.T) {
		peerList := getNewPeerList(hasMajorityTrue)

		result := peerList.PingList()

		assert.Empty(t, result, "Result differs from expected")
	})

	t.Run("PositiveSomePeersAreInPingList", func(t *testing.T) {
		peerList := getNewPeerList(hasMajorityTrue)
		peerID := getRandomEnodeID()
		peer := createNewPeerStruct(peerID, "myPeer", true, true)
		peer2 := createNewPeerStruct(getRandomEnodeID(), "myPeer2", false, true)

		peerList.AddPeer(peer)
		peerList.AddPeer(peer2)

		expected := []enode.ID{peerID}
		actual := peerList.PingList()

		assert.ElementsMatch(t, expected, actual, "Result differs from expected")
	})

	t.Run("PositiveAllPeersAreInPingList", func(t *testing.T) {
		peerList := getNewPeerList(hasMajorityTrue)
		peerID := getRandomEnodeID()
		peer2ID := getRandomEnodeID()
		peer := createNewPeerStruct(peerID, "myPeer", true, true)
		peer2 := createNewPeerStruct(peer2ID, "myPeer2", true, true)

		peerList.AddPeer(peer)
		peerList.AddPeer(peer2)

		expected := []enode.ID{peerID, peer2ID}
		actual := peerList.PingList()

		assert.ElementsMatch(t, expected, actual, "Result differs from expected")
	})
}

func TestPingListPeers(t *testing.T) {
	t.Run("PositivePingListIsEmpty", func(t *testing.T) {
		peerList := getNewPeerList(hasMajorityTrue)

		result := peerList.PingListPeers()

		assert.Empty(t, result, "Result differs from expected")
	})

	t.Run("PositiveSomePeersAreInPingList", func(t *testing.T) {
		peerList := getNewPeerList(hasMajorityTrue)
		peerID := getRandomEnodeID()
		peer := createNewPeerStruct(peerID, "myPeer", true, true)
		peer2 := createNewPeerStruct(getRandomEnodeID(), "myPeer2", false, false)

		peerList.AddPeer(peer)
		peerList.AddPeer(peer2)

		expected := []*Peer{peer}
		actual := peerList.PingListPeers()

		assert.ElementsMatch(t, expected, actual, "Result differs from expected")
	})

	t.Run("PositiveAllPeersAreInPingList", func(t *testing.T) {
		peerList := getNewPeerList(hasMajorityTrue)
		peerID := getRandomEnodeID()
		peer2ID := getRandomEnodeID()
		peer := createNewPeerStruct(peerID, "myPeer", true, true)
		peer2 := createNewPeerStruct(peer2ID, "myPeer2", true, true)

		peerList.AddPeer(peer)
		peerList.AddPeer(peer2)

		expected := []*Peer{peer, peer2}
		actual := peerList.PingListPeers()

		assert.ElementsMatch(t, expected, actual, "Result differs from expected")
	})
}

func TestPingListLen(t *testing.T) {
	t.Run("PositivePingListIsEmpty", func(t *testing.T) {
		peerList := getNewPeerList(hasMajorityTrue)

		actual := peerList.PingListLen()

		assert.Equal(t, 0, actual, "Count of peers in ping list differs from expected")
	})

	t.Run("PositiveSomePeersAreInPingList", func(t *testing.T) {
		peerList := getNewPeerList(hasMajorityTrue)
		peerID := getRandomEnodeID()
		peer2ID := getRandomEnodeID()
		peer := createNewPeerStruct(peerID, "myPeer", true, true)
		peer2 := createNewPeerStruct(peer2ID, "myPeer2", false, true)

		peerList.AddPeer(peer)
		peerList.AddPeer(peer2)

		actual := peerList.PingListLen()

		assert.Equal(t, 1, actual, "Count of peers in ping list differs from expected")
	})

	t.Run("PositiveAllPeersAreInPingList", func(t *testing.T) {
		peerList := getNewPeerList(hasMajorityTrue)
		peerID := getRandomEnodeID()
		peer2ID := getRandomEnodeID()
		peer := createNewPeerStruct(peerID, "myPeer", true, true)
		peer2 := createNewPeerStruct(peer2ID, "myPeer2", true, true)

		peerList.AddPeer(peer)
		peerList.AddPeer(peer2)

		actual := peerList.PingListLen()

		assert.Equal(t, 2, actual, "Count of peers in ping list differs from expected")
	})
}

func TestLen(t *testing.T) {
	t.Run("PositivePeerListIsEmpty", func(t *testing.T) {
		peerList := getNewPeerList(hasMajorityTrue)

		actual := peerList.Len()

		assert.Equal(t, 0, actual, "Count of peers in ping list differs from expected")
	})

	t.Run("PositiveSomePeersArePresent", func(t *testing.T) {
		peerList := getNewPeerList(hasMajorityTrue)
		peer := createNewPeerStruct(getRandomEnodeID(), "myPeer", true, true)
		peer2 := createNewPeerStruct(getRandomEnodeID(), "myPeer2", true, true)

		peerList.AddPeer(peer)
		peerList.AddPeer(peer2)

		actual := peerList.Len()

		assert.Equal(t, 2, actual, "Count of peers in ping list differs from expected")
	})
}

func TestSelfIsMinerCandidate(t *testing.T) {
	t.Run("PositiveMarkedAsMinerCandidate", func(t *testing.T) {
		peerList := getNewPeerList(hasMajorityTrue)
		peerList.MarkSelfIsMinerCandidate(true)

		actual := peerList.SelfIsMinerCandidate()

		assert.Equal(t, true, actual, "Expected self node to be a miner candidate")
	})

	t.Run("PositiveNotMarkedAsMinerCandidate", func(t *testing.T) {
		peerList := getNewPeerList(hasMajorityTrue)

		actual := peerList.SelfIsMinerCandidate()

		assert.Equal(t, false, actual, "Expected self node to be a miner candidate")
	})
}

func TestPeers(t *testing.T) {
	t.Run("PositivePeerListIsEmpty", func(t *testing.T) {
		peerList := getNewPeerList(hasMajorityTrue)

		var expected []*Peer
		actual := peerList.Peers()

		assert.ElementsMatch(t, expected, actual, "Result differs from expected")
	})

	t.Run("PositiveSomePeersArePresent", func(t *testing.T) {
		peerList := getNewPeerList(hasMajorityTrue)
		peer := createNewPeerStruct(getRandomEnodeID(), "myPeer", true, true)
		peer2 := createNewPeerStruct(getRandomEnodeID(), "myPeer2", true, true)

		peerList.AddPeer(peer)
		peerList.AddPeer(peer2)

		expected := []*Peer{peer, peer2}
		actual := peerList.Peers()

		assert.ElementsMatch(t, expected, actual, "Result differs from expected")
	})
}

func TestIds(t *testing.T) {
	t.Run("PositivePeerListIsEmpty", func(t *testing.T) {
		peerList := getNewPeerList(hasMajorityTrue)

		var expected []enode.ID
		actual := peerList.Ids()

		assert.ElementsMatch(t, expected, actual, "Resulting list of IDs differs from expected")
	})

	t.Run("PositiveSomePeersArePresent", func(t *testing.T) {
		peerList := getNewPeerList(hasMajorityTrue)
		peerID := getRandomEnodeID()
		peer2ID := getRandomEnodeID()
		peer := createNewPeerStruct(peerID, "myPeer", true, true)
		peer2 := createNewPeerStruct(peer2ID, "myPeer2", true, true)

		peerList.AddPeer(peer)
		peerList.AddPeer(peer2)

		expected := []enode.ID{peerID, peer2ID}
		actual := peerList.Ids()

		assert.ElementsMatch(t, expected, actual, "Resulting list of IDs differs from expected")
	})
}

func TestGetSnapshotNodes(t *testing.T) {
	t.Run("PositiveHasMajorityNoneInPingList", func(t *testing.T) {
		peerList := getNewPeerList(hasMajorityTrue)
		peer := createNewPeerStruct(getRandomEnodeID(), "myPeer", false, true)
		peer2 := createNewPeerStruct(getRandomEnodeID(), "myPeer2", false, true)

		peerList.AddPeer(peer)
		peerList.AddPeer(peer2)

		expected := prepareExpectedNodesForPeerList(peerList)
		actual, err := peerList.GetSnapshotNodes()

		assert.NoError(t, err, "Unexpected error")
		assert.ElementsMatch(t, expected, actual, "Result differs from expected")
	})

	t.Run("PositiveHasMajoritySomeInPingList", func(t *testing.T) {
		peerList := getNewPeerList(hasMajorityTrue)
		peer := createNewPeerStruct(getRandomEnodeID(), "myPeer", true, true)
		peer2 := createNewPeerStruct(getRandomEnodeID(), "myPeer2", false, true)

		peerList.AddPeer(peer)
		peerList.AddPeer(peer2)

		expected := prepareExpectedNodesForPeerList(peerList)
		actual, err := peerList.GetSnapshotNodes()
		assert.NoError(t, err, "Unexpected error")
		assert.ElementsMatch(t, expected, actual, "Result differs from expected")
	})

	t.Run("PositiveHasMajorityAllInPingList", func(t *testing.T) {
		peerList := getNewPeerList(hasMajorityTrue)
		peer := createNewPeerStruct(getRandomEnodeID(), "myPeer", true, true)
		peer2 := createNewPeerStruct(getRandomEnodeID(), "myPeer2", true, true)

		peerList.AddPeer(peer)
		peerList.AddPeer(peer2)

		expected := prepareExpectedNodesForPeerList(peerList)
		actual, err := peerList.GetSnapshotNodes()

		assert.NoError(t, err, "Unexpected error")
		assert.ElementsMatch(t, expected, actual, "Result differs from expected")
	})

	t.Run("NegativeDoesNotHaveMajority", func(t *testing.T) {
		peerList := getNewPeerList(hasMajorityFalse)

		result, err := peerList.GetSnapshotNodes()

		assert.EqualError(t, err, errNoMajorityOfNodes.Error(), "Expected NoMajorityOfNodes error")
		assert.Empty(t, result, "Expected result to be empty")
	})
}
