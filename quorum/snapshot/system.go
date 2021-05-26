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

package snapshot

import (
	"sort"
	"sync"

	"github.com/pkg/errors"
	"github.com/simcord-llc/bitbon-system-blockchain/common"
	"github.com/simcord-llc/bitbon-system-blockchain/log"
	sortUtils "github.com/simcord-llc/bitbon-system-blockchain/quorum/utils/sort"
)

const MaxEpochs = 20

var (
	ErrAlreadyExists = errors.New("snapshot already exists")
)

type Constructor func(epoch uint64, logger log.Logger) (*Snapshot, error)

type System struct {
	mu sync.RWMutex // Mutex protects all fields

	recentEpoch uint64
	snapshot    Constructor
	logger      log.Logger

	snapshotEpoch map[uint64]*Snapshot // map of epoch to snapshot
}

type SystemDTO struct {
	SnapshotEpoch map[uint64]*SnapDTO
}

// NewSystem - System constructor
func NewSystem(constructor Constructor, logger log.Logger) *System {
	if logger == nil {
		logger = log.New("component", "SnapshotSystem")
	}
	return &System{
		snapshot:      constructor,
		logger:        logger,
		snapshotEpoch: make(map[uint64]*Snapshot),
	}
}

func (s *System) ToDTO() *SystemDTO {
	s.mu.RLock()
	defer s.mu.RUnlock()

	dto := make(map[uint64]*SnapDTO, len(s.snapshotEpoch))
	for epoch, s := range s.snapshotEpoch {
		dto[epoch] = s.ToDTO()
	}

	return &SystemDTO{
		SnapshotEpoch: dto,
	}
}

func (s *System) RecentEpoch() uint64 {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.recentEpoch
}

func (s *System) Epochs() []uint64 {
	s.mu.RLock()
	defer s.mu.RUnlock()

	epochs := make([]uint64, 0, len(s.snapshotEpoch))
	for epoch := range s.snapshotEpoch {
		epochs = append(epochs, epoch)
	}

	return epochs
}

func (s *System) PrepareSnapshot(epoch uint64) (*Snapshot, error) {
	s.mu.Lock()

	if s.recentEpoch >= epoch {
		s.mu.Unlock()
		return nil, ErrAlreadyExists
	}
	snapshot, err := s.snapshot(epoch, s.logger.New("subcomponent", "Snapshot", "epoch", epoch))
	if err != nil {
		s.mu.Unlock()
		return nil, err
	}

	s.recentEpoch = epoch
	s.snapshotEpoch[epoch] = snapshot

	s.mu.Unlock()

	s.ClearOldSnapshots()

	return snapshot, nil
}

// GetSnapshot returns snapshot at epoch
func (s *System) GetSnapshot(epoch uint64) *Snapshot {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.snapshotEpoch[epoch]
}

// DeleteSnapshot deletes old epoch
func (s *System) DeleteSnapshot(epoch uint64) {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.snapshotEpoch, epoch)
}

// ClearOldSnapshots trims old snapshots that exceed MaxEpochs number
func (s *System) ClearOldSnapshots() {
	s.mu.RLock()

	if len(s.snapshotEpoch) <= MaxEpochs {
		s.mu.RUnlock()
		return
	}

	var keys sortUtils.Range
	for k := range s.snapshotEpoch {
		keys = append(keys, k)
	}

	s.mu.RUnlock()

	sort.Sort(keys)
	for _, i := range keys[:len(keys)-MaxEpochs] {
		s.DeleteSnapshot(i)
	}
}

func (s *System) MinerListAtEpoch(epoch uint64) []common.Address {
	snap := s.GetSnapshot(epoch)
	if snap == nil {
		s.logger.Debug("Snapshot not found in snapshot system", "epoch", epoch)
		return nil
	}

	minerList := snap.MinerList()
	if minerList == nil || len(minerList.Miners()) == 0 {
		s.logger.Debug("Snapshot returned empty miner list", "epoch", epoch)
		return nil
	}

	return minerList.Miners()
}
