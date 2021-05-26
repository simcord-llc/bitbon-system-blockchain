// Copyright Simcord LLC
// This file is part of the Bitbon System Blockchain Node library.
//
// The Bitbon System Blockchain Node library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The Bitbon System Blockchain Node library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the Bitbon System Blockchain Node library. If not, see <http://www.gnu.org/licenses/>.

package quorum

import (
	"encoding/json"
	"errors"
	"math/big"

	lru "github.com/hashicorp/golang-lru"
	"github.com/simcord-llc/bitbon-system-blockchain/common"
	"github.com/simcord-llc/bitbon-system-blockchain/core/rawdb"
	"github.com/simcord-llc/bitbon-system-blockchain/core/types"
	"github.com/simcord-llc/bitbon-system-blockchain/crypto"
	"github.com/simcord-llc/bitbon-system-blockchain/ethdb"
	"github.com/simcord-llc/bitbon-system-blockchain/log"
	"github.com/simcord-llc/bitbon-system-blockchain/params"
)

var (
	errSnapshotNotFound = errors.New("cannot find snapshot in storage ")
)

type TimeFrame struct {
	Start uint64 `json:"start"`
	End   uint64 `json:"end"`
}

type BlockProducer struct {
	SequenceNumber uint64 `json:"sequenceNumber"`
}

// Snapshot is the state of the authorization voting at a given point in time.
type Snapshot struct {
	config   *params.QuorumEngineConfig // Consensus engine parameters to fine tune behavior
	sigcache *lru.ARCCache              // Cache of recent block signatures to speed up ecrecover

	Epoch   uint64                           `json:"epoch"`
	Signers map[common.Address]BlockProducer `json:"signers"` // Set of authorized signers at this moment
	Recents map[uint64]common.Address        `json:"recents"`
	// Set of recent signers for spam protections
}

// newSnapshot creates a new snapshot with the specified startup parameters. This
// method does not initialize the set of recent signers, so only ever use if for
// the genesis block.
func newSnapshot(config *params.QuorumEngineConfig, sigcache *lru.ARCCache,
	signers map[common.Address]BlockProducer, epoch uint64) *Snapshot {
	snap := &Snapshot{
		config:   config,
		sigcache: sigcache,
		Signers:  make(map[common.Address]BlockProducer),
		Recents:  make(map[uint64]common.Address),
		Epoch:    epoch,
	}
	snap.Signers = signers
	return snap
}

// nolint:interfacer
// loadSnapshot loads an existing snapshot from the database.
func loadSnapshot(config *params.QuorumEngineConfig, sigcache *lru.ARCCache, db ethdb.Database,
	hash common.Hash, number uint64) (*Snapshot, error) {
	blob := rawdb.ReadQuorumRLP(db, hash, number)
	if blob == nil {
		return nil, errSnapshotNotFound
	}

	snap := new(Snapshot)
	if err := json.Unmarshal(blob, snap); err != nil {
		log.Error("loadSnapshot json.Unmarshal", "snap", snap, "err", err)
		return nil, err
	}
	snap.config = config
	snap.sigcache = sigcache

	return snap, nil
}

// store inserts the snapshot into the database.
func (s *Snapshot) store(db ethdb.Database) error {
	blob, err := json.Marshal(s)
	if err != nil {
		return err
	}
	return db.Put(rawdb.QuorumKey(s.Hash()), blob)
}

func (s *Snapshot) Hash() (h common.Hash) {
	hash := make([]byte, 0, len(s.Signers)*common.AddressLength)
	for _, address := range s.signers() {
		hash = append(hash, address[:]...)
	}

	return common.BytesToHash(crypto.Keccak256(new(big.Int).SetUint64(s.Epoch).Bytes(), hash))
}

// signers retrieves the list of authorized signers in ascending order.
func (s *Snapshot) signers() []common.Address {
	signatures := make([]common.Address, len(s.Signers))
	for sig, bp := range s.Signers {
		signatures[bp.SequenceNumber] = sig
	}
	return signatures
}

// signersFull retrieves the map of authorized signers.
func (s *Snapshot) signersFull() map[common.Address]BlockProducer {
	return s.Signers
}

// inturn returns if a signer at a given time is in-turn or not.
// inturn = block.timestamp in [startFrame, endFrame) - startFrame included; endFrame - not
func (s *Snapshot) inturn(header *types.Header, signer common.Address) bool {
	bp, ok := s.Signers[signer]
	if !ok {
		log.Error("unknown signer", "signerAddress", signer.Hex())
		return false
	}

	// startFrame - time which indicated start of the time frame of a block producer that produced received block
	// signersLen - number of signers in the epoch
	// blockTime - precise time from received block header
	// elapsed - number of second between the start of the epoch and the startFrame
	// round - sequential number of the round from the epoch start
	signersLen := uint64(len(s.Signers))
	blockTime := header.Time
	epochStart := s.config.Start + s.Epoch*s.config.Epoch
	epochEnd := epochStart + s.config.Epoch
	if epochStart > blockTime || epochEnd < blockTime {
		log.Error("Block timestamp is out of epoch bounds")
		return false
	}
	startFrame := blockTime - blockTime%s.config.Period
	elapsed := startFrame - epochStart
	round := elapsed / (s.config.Period * signersLen)
	roundStart := epochStart + (s.config.Period*signersLen)*round

	inturnSigner := (startFrame - roundStart) / (s.config.Period)

	// check if sequence number of block signer equals calculated inturnSigner
	return bp.SequenceNumber == inturnSigner
}

func (s *Snapshot) signersElapsed(header *types.Header) uint64 {
	blockTime := header.Time
	epochStart := s.config.Start + s.Epoch*s.config.Epoch
	elapsed := blockTime - epochStart
	return elapsed / s.config.Period
}

// checkSigners checks snapshot signers to be present in given candidates
func (s *Snapshot) checkSigners(candidates []common.Address) bool {
	if len(candidates) == 0 {
		return false
	}
	if len(s.Signers) == 0 {
		return false
	}

	candidatesMap := make(map[common.Address]struct{}, len(candidates))
	for _, c := range candidates {
		candidatesMap[c] = struct{}{}
	}

	for addr := range s.Signers {
		if _, ok := candidatesMap[addr]; !ok {
			return false
		}
	}

	return true
}
