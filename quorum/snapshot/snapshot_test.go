package snapshot

import (
	"testing"

	"github.com/simcord-llc/bitbon-system-blockchain/common"
	"github.com/stretchr/testify/assert"
)

func TestMinerListToDTO(t *testing.T) {
	minerList := &MinerList{
		miners: make([]common.Address, 0),
		hash:   Hash,
	}

	expected := &MinerListDTO{
		Miners: make([]common.Address, 0),
		Hash:   Hash,
	}

	assert.Equal(t, expected, minerList.ToDTO(), "Mapped DTO differs from expected")
}

func TestSnapshotToDTO(t *testing.T) {
	t.Run("PositiveWithoutRemoteMiners", func(t *testing.T) {
		nodeID := GetRandomEnodeID()
		node := NewNode(nodeID, address, true, make(map[common.Address]bool))

		snapshotID := GetRandomEnodeID()
		snapshot := NewSnapshot(epochNumber, snapshotID, []*Node{node}, hasMajorityTrue, logger)

		nodeDTO := &NodeDTO{
			ID:                   nodeID,
			Address:              address,
			IsMinerCandidate:     true,
			RemoteNodeCandidates: make(map[common.Address]bool),
			RemoteMiners:         RemoteMinersDTO{},
		}

		expected := &SnapDTO{
			Epoch:     epochNumber,
			SelfID:    snapshotID,
			Nodes:     []*NodeDTO{nodeDTO},
			MinerList: nil,
		}

		assert.Equal(t, expected, snapshot.ToDTO(), "Mapped DTO differs from expected")
	})

	t.Run("PositiveWithRemoteMiners", func(t *testing.T) {
		nodeID := GetRandomEnodeID()
		node := NewNode(nodeID, address, true, make(map[common.Address]bool))
		node.setRemoteMiners(make([]common.Address, 0), Hash)

		snapshotID := GetRandomEnodeID()
		snapshot := NewSnapshot(epochNumber, snapshotID, []*Node{node}, hasMajorityTrue, logger)

		snapshot.TryApproveMinerList()

		nodeDTO := &NodeDTO{
			ID:                   nodeID,
			Address:              address,
			IsMinerCandidate:     true,
			RemoteNodeCandidates: make(map[common.Address]bool),
			RemoteMiners: RemoteMinersDTO{
				Hash:   Hash,
				Miners: make([]common.Address, 0),
			},
		}

		expected := &SnapDTO{
			Epoch:  epochNumber,
			SelfID: snapshotID,
			Nodes:  []*NodeDTO{nodeDTO},
			MinerList: &MinerListDTO{
				Miners: make([]common.Address, 0),
				Hash:   Hash,
			},
		}

		assert.Equal(t, expected, snapshot.ToDTO(), "Mapped DTO differs from expected")
	})
}

func TestSnapshotToMap(t *testing.T) {
	node := NewNode(GetRandomEnodeID(), address, true, make(map[common.Address]bool))
	snapshot := NewSnapshot(epochNumber, GetRandomEnodeID(), []*Node{node}, hasMajorityTrue, logger)

	expected := getMapFromSnapshot(snapshot)

	assert.Equal(t, expected, snapshot.ToMap(), "Resulting map differs from expected")
}

func TestGetSnapshotNode(t *testing.T) {
	t.Run("PositiveNodeIsPresentInSnapshot", func(t *testing.T) {
		nodeID := GetRandomEnodeID()
		node := NewNode(nodeID, address, true, make(map[common.Address]bool))
		snapshot := NewSnapshot(epochNumber, GetRandomEnodeID(), []*Node{node}, hasMajorityTrue, logger)

		result := snapshot.getSnapshotNode(nodeID)

		assert.Equal(t, node, result, "Resulting node differs from expected")
	})

	t.Run("PositiveNodeIsNotPresentInSnapshot", func(t *testing.T) {
		snapshot := NewSnapshot(epochNumber, GetRandomEnodeID(), []*Node{}, hasMajorityTrue, logger)

		result := snapshot.getSnapshotNode(GetRandomEnodeID())

		assert.Empty(t, result, "Resulting node differs from expected")
	})
}

func TestTryGetSelfMinerList(t *testing.T) {
	t.Run("PositiveMinerListIsAlreadySet", func(t *testing.T) {
		snapshotID := GetRandomEnodeID()
		node := NewNode(snapshotID, address, true, make(map[common.Address]bool))
		node.setRemoteMiners(make([]common.Address, 0), Hash)

		snapshot := NewSnapshot(epochNumber, snapshotID, []*Node{node}, hasMajorityTrue, logger)
		snapshot.TryApproveMinerList()

		result := snapshot.TryGetSelfMinerList()

		assert.Equal(t, snapshot.minerList, result, "Result differs from expected")
	})

	t.Run("PositiveMinerListIsAlreadySetMinerSliceIsNil", func(t *testing.T) {
		snapshotID := GetRandomEnodeID()
		node := NewNode(snapshotID, address, true, make(map[common.Address]bool))
		node.setRemoteMiners(nil, Hash)

		snapshot := NewSnapshot(epochNumber, snapshotID, []*Node{node}, hasMajorityTrue, logger)
		snapshot.TryApproveMinerList()

		result := snapshot.TryGetSelfMinerList()

		assert.Empty(t, result, "Result differs from expected")
	})

	t.Run("PositiveMinerListOnlyOneNodeIsInSnapshot", func(t *testing.T) {
		node := NewNode(GetRandomEnodeID(), address, true, make(map[common.Address]bool))
		snapshot := NewSnapshot(epochNumber, GetRandomEnodeID(), []*Node{node}, hasMajorityTrue, logger)

		result := snapshot.TryGetSelfMinerList()

		assert.Empty(t, result, "Result differs from expected")
	})

	t.Run("PositiveMinerListMoreThanOneNodeIsInSnapshotNoMajority", func(t *testing.T) {
		node1 := NewNode(GetRandomEnodeID(), address, true, make(map[common.Address]bool))
		node2 := NewNode(GetRandomEnodeID(), address, true, make(map[common.Address]bool))

		snapshot := NewSnapshot(epochNumber, GetRandomEnodeID(), []*Node{node1, node2}, hasMajorityFalse, logger)

		result := snapshot.TryGetSelfMinerList()

		assert.Empty(t, result, "Result differs from expected")
	})

	t.Run("PositiveMinerListMoreThanOneNodeIsInSnapshotHasMajorityNoCandidates", func(t *testing.T) {
		node1 := NewNode(GetRandomEnodeID(), address, true, make(map[common.Address]bool))
		node2 := NewNode(GetRandomEnodeID(), address, true, make(map[common.Address]bool))
		snapshot := NewSnapshot(epochNumber, GetRandomEnodeID(), []*Node{node1, node2}, hasMajorityTrue, logger)

		result := snapshot.TryGetSelfMinerList()

		assert.Empty(t, result, "Result differs from expected")
	})

	t.Run("PositiveMinerListMoreThanOneNodeIsInSnapshotHasMajorityWithCandidates", func(t *testing.T) {
		numOfMinerCandidates := 5

		node1 := NewNode(GetRandomEnodeID(), address, true, make(map[common.Address]bool))
		node2 := NewNode(GetRandomEnodeID(), address, true, make(map[common.Address]bool))
		snapshot := NewSnapshot(epochNumber, GetRandomEnodeID(), []*Node{node1, node2}, hasMajorityTrue, logger)

		addresses := [][]common.Address{
			GetRandomAddresses(numOfMinerCandidates),
			GetRandomAddresses(numOfMinerCandidates),
		}

		var expectedMiners []common.Address

		for i := range snapshot.nodes {
			for j := 0; j < numOfMinerCandidates; j++ {
				snapshot.nodes[i].remoteNodeCandidates[addresses[i][j]] = true
				expectedMiners = append(expectedMiners, addresses[i][j])
			}
		}

		hash := common.Keccak256Addresses(expectedMiners...)

		expected := &MinerList{
			miners: expectedMiners,
			hash:   hash,
		}
		actual := snapshot.TryGetSelfMinerList()

		assert.Equal(t, expected.hash, actual.hash, "Resulting Hash differs from expected")
		assert.ElementsMatch(t, expected.miners, actual.miners, "Resulting miners differs from expected")
	})
}

func TestTryApproveMinerList(t *testing.T) {
	t.Run("PositiveMinerListAlreadySet", func(t *testing.T) {
		node := NewNode(GetRandomEnodeID(), address, true, make(map[common.Address]bool))
		snapshot := NewSnapshot(epochNumber, GetRandomEnodeID(), []*Node{node}, hasMajorityTrue, logger)
		snapshot.minerList = &MinerList{
			miners: GetRandomAddresses(3),
			hash:   Hash,
		}

		before := snapshot.minerList

		snapshot.TryApproveMinerList()

		after := snapshot.minerList

		assert.Equal(t, before, after, "Unexpected miner list change")
	})

	t.Run("PositiveRemoteMinerMajority", func(t *testing.T) {
		node := NewNode(GetRandomEnodeID(), address, true, make(map[common.Address]bool))
		miners := GetRandomAddresses(3)
		node.remoteMiners = remoteMiners{
			miners: miners,
			hash:   Hash,
		}

		snapshot := NewSnapshot(epochNumber, GetRandomEnodeID(), []*Node{node}, hasMajorityTrue, logger)

		snapshot.TryApproveMinerList()

		expected := &MinerList{
			miners: miners,
			hash:   Hash,
		}

		assert.Equal(t, expected, snapshot.minerList, "Miner list differs from expected")
	})

	t.Run("PositiveRemoteMinerNoMajority", func(t *testing.T) {
		node := NewNode(GetRandomEnodeID(), address, true, make(map[common.Address]bool))
		miners := GetRandomAddresses(3)
		node.remoteMiners = remoteMiners{
			miners: miners,
			hash:   Hash,
		}

		snapshot := NewSnapshot(epochNumber, GetRandomEnodeID(), []*Node{node}, hasMajorityFalse, logger)

		snapshot.TryApproveMinerList()

		assert.Empty(t, snapshot.minerList, "Miner list differs from expected")
	})
}

func TestSetRemotePeerMiners(t *testing.T) {
	t.Run("PositiveMinerListAlreadySet", func(t *testing.T) {
		nodeID := GetRandomEnodeID()
		node := NewNode(nodeID, address, true, make(map[common.Address]bool))

		snapshot := NewSnapshot(epochNumber, GetRandomEnodeID(), []*Node{node}, hasMajorityTrue, logger)

		snapshot.minerList = &MinerList{
			miners: GetRandomAddresses(3),
			hash:   Hash,
		}

		beforeMiners, beforeHash := snapshot.getSnapshotNode(nodeID).getRemoteMiners()

		snapshot.SetRemotePeerMiners(nodeID, GetRandomAddresses(3), Hash)

		afterMiners, afterHash := snapshot.getSnapshotNode(nodeID).getRemoteMiners()

		assert.Equal(t, beforeHash, afterHash, "Unexpected Hash change")
		assert.ElementsMatch(t, beforeMiners, afterMiners, "Unexpected remote miners list change")
	})

	t.Run("PositiveSnapshotNodeHasNoCandidates", func(t *testing.T) {
		nodeID := GetRandomEnodeID()
		node := NewNode(nodeID, address, true, make(map[common.Address]bool))

		snapshot := NewSnapshot(epochNumber, GetRandomEnodeID(), []*Node{node}, hasMajorityTrue, logger)

		nodeMiners, nodeHash := snapshot.getSnapshotNode(nodeID).getRemoteMiners()

		snapshot.SetRemotePeerMiners(nodeID, GetRandomAddresses(3), Hash)

		assert.Empty(t, nodeMiners, "Expected remote miners list to be empty")
		assert.Equal(t, nodeHash, common.Hash{}, "Unexpected Hash change")
	})

	t.Run("PositiveSnapshotNodeIsValid", func(t *testing.T) {
		nodeID := GetRandomEnodeID()
		node := NewNode(nodeID, address, true, make(map[common.Address]bool))

		addresses := GetRandomAddresses(2)

		for _, v := range addresses {
			node.remoteNodeCandidates[v] = true
		}

		snapshot := NewSnapshot(epochNumber, GetRandomEnodeID(), []*Node{node}, hasMajorityTrue, logger)

		expectedMiners := GetRandomAddresses(3)

		snapshot.SetRemotePeerMiners(nodeID, expectedMiners, Hash)

		actualMiners, actualHash := snapshot.getSnapshotNode(nodeID).getRemoteMiners()

		assert.Equal(t, Hash, actualHash, "Unexpected Hash change")
		assert.ElementsMatch(t, expectedMiners, actualMiners, "Unexpected remote miners list change")
	})
}
