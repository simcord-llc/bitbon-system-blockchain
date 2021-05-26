package snapshot

import (
	"context"
	"crypto/rand"
	"errors"
	"math/big"
	"sort"

	"github.com/simcord-llc/bitbon-system-blockchain/log"
	sortUtils "github.com/simcord-llc/bitbon-system-blockchain/quorum/utils/sort"

	"github.com/simcord-llc/bitbon-system-blockchain/common"
	"github.com/simcord-llc/bitbon-system-blockchain/p2p/enode"
)

var (
	address            = common.BytesToAddress([]byte("798e8c22eb5a46a5a401a4a123e4fe768e92abbe"))
	remoteMinerAddress = common.BytesToAddress([]byte("44811112eb1233a5a402546243e4fe763453778e"))

	Hash      = common.BytesToHash([]byte("221e8c22eb5a46a5a401a4a123e4fe768e92a111"))
	wrongHash = common.BytesToHash([]byte("33338c22eb5a46a5a401a4a123e4fe768e92a222"))

	epochNumber = uint64(10)

	addressesCount = 10

	ErrFailedSnapshotCreation = errors.New("error during construction of snapshot")

	logger = log.New(context.TODO())
)

func getNodeDTO(id enode.ID, hasRemoteMiners bool) *NodeDTO {
	nodeDTO := &NodeDTO{
		ID:                   id,
		Address:              address,
		IsMinerCandidate:     true,
		RemoteNodeCandidates: make(map[common.Address]bool),
	}

	if hasRemoteMiners {
		nodeDTO.RemoteMiners = RemoteMinersDTO{
			Hash:   Hash,
			Miners: make([]common.Address, 0),
		}
	}

	return nodeDTO
}

func GetRandomAddresses(count int) []common.Address {
	addresses := make([]common.Address, 0, count)

	for i := 0; i < count; i++ {
		addresses = append(addresses, common.BytesToAddress(getRandomByteArray()))
	}

	return addresses
}

func getRandomByteArray() []byte {
	token := make([]byte, 20)

	if _, err := rand.Read(token); err != nil {
		panic("error while generating random byte array")
	}

	return token
}

func GetRandomEnodeID() enode.ID {
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

func snapshotConstructor(epoch uint64, _ log.Logger) (*Snapshot, error) {
	node := NewNode(GetRandomEnodeID(), address, true, make(map[common.Address]bool))
	return NewSnapshot(epoch, GetRandomEnodeID(), []*Node{node}, hasMajorityTrue, logger), nil
}

func invalidSnapshotConstructor(_ uint64, _ log.Logger) (*Snapshot, error) {
	return nil, ErrFailedSnapshotCreation
}

func getMapFromSnapshot(snapshot *Snapshot) map[common.Address]bool {
	mapOfNodes := make(map[common.Address]bool)

	for _, v := range snapshot.nodes {
		mapOfNodes[v.getAddress()] = v.getIsMinerCandidate()
	}

	return mapOfNodes
}

func AddSnapshotAtEpoch(snapshotSystem *System, epochNumber uint64) *Snapshot {
	snapshot, err := snapshotSystem.PrepareSnapshot(epochNumber)
	if err != nil {
		panic("Error during snapshot preparation")
	}

	return snapshot
}

func TrimOldSnapshots(snapshots map[uint64]*Snapshot) map[uint64]*Snapshot {
	var keys sortUtils.Range

	for k := range snapshots {
		keys = append(keys, k)
	}

	sort.Sort(keys)

	for _, i := range keys[:len(keys)-MaxEpochs] {
		delete(snapshots, i)
	}

	return snapshots
}

func hasMajorityTrue(int, bool) bool {
	return true
}

func hasMajorityFalse(int, bool) bool {
	return false
}
