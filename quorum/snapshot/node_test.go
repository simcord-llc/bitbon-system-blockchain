package snapshot

import (
	"testing"

	"github.com/simcord-llc/bitbon-system-blockchain/common"
	"github.com/stretchr/testify/assert"
)

func TestNodeToDTO(t *testing.T) {
	t.Run("Positive_WithRemoteMiners", func(t *testing.T) {
		nodeID := GetRandomEnodeID()
		node := NewNode(nodeID, address, true, make(map[common.Address]bool))

		expected := getNodeDTO(nodeID, false)

		assert.Equal(t, expected, node.ToDTO(), "Mapped DTO differs from expected")
	})

	t.Run("Positive_WithoutRemoteMiners", func(t *testing.T) {
		nodeID := GetRandomEnodeID()
		node := NewNode(nodeID, address, true, make(map[common.Address]bool))
		node.setRemoteMiners(make([]common.Address, 0), Hash)

		expected := getNodeDTO(nodeID, true)

		assert.Equal(t, expected, node.ToDTO(), "Mapped DTO differs from expected")
	})
}

func TestIsSetRemoteNodeCandidates(t *testing.T) {
	t.Run("Positive_RemoteCandidatesAreNotSet", func(t *testing.T) {
		node := NewNode(GetRandomEnodeID(), address, true, make(map[common.Address]bool))

		result := node.isSetRemoteNodeCandidates()

		assert.False(t, result, "Expected node to have miner candidates unset")
	})

	t.Run("Positive_RemoteCandidatesAreSetButNotAsMiners", func(t *testing.T) {
		node := NewNode(GetRandomEnodeID(), address, true, make(map[common.Address]bool))
		node.remoteNodeCandidates[remoteMinerAddress] = false

		result := node.isSetRemoteNodeCandidates()

		assert.True(t, result, "Expected node to have miner candidates set")
	})

	t.Run("Positive_RemoteCandidatesAreSetAndAreMiners", func(t *testing.T) {
		node := NewNode(GetRandomEnodeID(), address, true, make(map[common.Address]bool))
		node.remoteNodeCandidates[remoteMinerAddress] = true

		result := node.isSetRemoteNodeCandidates()

		assert.True(t, result, "Expected node to have miner candidates set")
	})
}

func TestGetRemoteMinerCandidates(t *testing.T) {
	t.Run("Positive_RemoteCandidatesAreAllMiners", func(t *testing.T) {
		node := NewNode(GetRandomEnodeID(), address, true, make(map[common.Address]bool))
		addresses := GetRandomAddresses(addressesCount)

		var expected []common.Address

		for _, v := range addresses {
			node.remoteNodeCandidates[v] = true
			expected = append(expected, v)
		}

		actual := node.getRemoteMinerCandidates()

		assert.ElementsMatch(t, expected, actual, "Result differs from expected")
	})

	t.Run("Positive_SomeRemoteCandidatesAreMiners", func(t *testing.T) {
		node := NewNode(GetRandomEnodeID(), address, true, make(map[common.Address]bool))
		addresses := GetRandomAddresses(addressesCount)

		var expected []common.Address

		for i, v := range addresses {
			if i%2 == 0 {
				node.remoteNodeCandidates[v] = false
			} else {
				node.remoteNodeCandidates[v] = true
				expected = append(expected, v)
			}
		}

		actual := node.getRemoteMinerCandidates()

		assert.ElementsMatch(t, expected, actual, "Result differs from expected")
	})

	t.Run("Positive_RemoteCandidatesAreAllNotMiners", func(t *testing.T) {
		node := NewNode(GetRandomEnodeID(), address, true, make(map[common.Address]bool))
		addresses := GetRandomAddresses(addressesCount)

		var expected []common.Address

		for _, v := range addresses {
			node.remoteNodeCandidates[v] = false
		}

		actual := node.getRemoteMinerCandidates()

		assert.ElementsMatch(t, expected, actual, "Result differs from expected")
	})
}

func TestHasRemoteMinerCandidate(t *testing.T) {
	t.Run("Positive_RemoteIsPresentAndMiner", func(t *testing.T) {
		node := NewNode(GetRandomEnodeID(), address, true, make(map[common.Address]bool))
		node.remoteNodeCandidates[remoteMinerAddress] = true

		result := node.hasRemoteMinerCandidate(remoteMinerAddress, nil)

		assert.True(t, result, "Expected node to have given address as a miner candidate")
	})

	t.Run("Positive_RemoteIsPresentButNotMiner", func(t *testing.T) {
		node := NewNode(GetRandomEnodeID(), address, true, make(map[common.Address]bool))
		node.remoteNodeCandidates[remoteMinerAddress] = false

		result := node.hasRemoteMinerCandidate(remoteMinerAddress, nil)

		assert.False(t, result, "Expected node to have given address as not a miner candidate")
	})

	t.Run("Positive_RemoteIsNotPresent", func(t *testing.T) {
		node := NewNode(GetRandomEnodeID(), address, true, make(map[common.Address]bool))
		result := node.hasRemoteMinerCandidate(remoteMinerAddress, nil)

		assert.False(t, result, "Expected node to have given address missing")
	})
}

func TestEqualRemoteMinerHash(t *testing.T) {
	t.Run("Positive_EqualHash", func(t *testing.T) {
		node := NewNode(GetRandomEnodeID(), address, true, make(map[common.Address]bool))
		node.setRemoteMiners(make([]common.Address, 0), Hash)

		result := node.equalRemoteMinersHash(Hash)

		assert.True(t, result, "Expected equal Hash for remote miners")
	})

	t.Run("Positive_DifferentHash", func(t *testing.T) {
		node := NewNode(GetRandomEnodeID(), address, true, make(map[common.Address]bool))
		node.setRemoteMiners(make([]common.Address, 0), Hash)

		result := node.equalRemoteMinersHash(wrongHash)

		assert.False(t, result, "Expected not equal Hash for remote miners")
	})
}
