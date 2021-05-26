package snapshot

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSystemToDTO(t *testing.T) {
	t.Run("PositiveWithoutSnapshots", func(t *testing.T) {
		snapshotSystem := NewSystem(snapshotConstructor, logger)

		expected := &SystemDTO{
			SnapshotEpoch: make(map[uint64]*SnapDTO),
		}

		assert.Equal(t, expected, snapshotSystem.ToDTO(), "Mapped DTO differs from expected")
	})

	t.Run("PositiveWithSnapshots", func(t *testing.T) {
		snapshotSystem := NewSystem(snapshotConstructor, logger)
		snapshot := AddSnapshotAtEpoch(snapshotSystem, epochNumber)

		snapshotMap := make(map[uint64]*SnapDTO)
		snapshotMap[epochNumber] = snapshot.ToDTO()

		expected := &SystemDTO{
			SnapshotEpoch: snapshotMap,
		}

		assert.Equal(t, expected, snapshotSystem.ToDTO(), "Mapped DTO differs from expected")
	})
}

func TestEpochs(t *testing.T) {
	t.Run("PositiveNoSnapshots", func(t *testing.T) {
		snapshotSystem := NewSystem(snapshotConstructor, logger)

		assert.Empty(t, snapshotSystem.Epochs(), "Expected epoch list to be empty")
	})

	t.Run("PositiveWithSnapshots", func(t *testing.T) {
		snapshotSystem := NewSystem(snapshotConstructor, logger)
		AddSnapshotAtEpoch(snapshotSystem, epochNumber)
		AddSnapshotAtEpoch(snapshotSystem, epochNumber+1)

		expected := []uint64{epochNumber, epochNumber + 1}

		assert.ElementsMatch(t, expected, snapshotSystem.Epochs(), "Result differs from expected")
	})
}

func TestPrepareSnapshot(t *testing.T) {
	t.Run("PositiveNewEpoch", func(t *testing.T) {
		snapshotSystem := NewSystem(snapshotConstructor, logger)

		snapshot, err := snapshotSystem.PrepareSnapshot(epochNumber)

		assert.NoError(t, err, "Unexpected error")
		assert.NotEmpty(t, snapshot, "Result differs from expected")
		assert.Equal(t, epochNumber, snapshotSystem.RecentEpoch(), "Epoch number differs from expected")
	})

	t.Run("NegativeFailedSnapshotConstructor", func(t *testing.T) {
		snapshotSystem := NewSystem(invalidSnapshotConstructor, logger)
		snapshot, err := snapshotSystem.PrepareSnapshot(epochNumber)

		assert.Empty(t, snapshot, "Expected snapshot to be empty")
		assert.EqualError(t, err, ErrFailedSnapshotCreation.Error(), "Expected SnapshotCreationFail error")
	})

	t.Run("NegativeOldEpoch", func(t *testing.T) {
		snapshotSystem := NewSystem(snapshotConstructor, logger)
		snapshotSystem.recentEpoch = epochNumber
		snapshot, err := snapshotSystem.PrepareSnapshot(epochNumber)

		assert.Empty(t, snapshot, "Expected snapshot to be empty")
		assert.EqualError(t, err, ErrAlreadyExists.Error(), "Expected AlreadyExists error")
	})
}

func TestGetSnapshot(t *testing.T) {
	t.Run("PositiveSnapshotExists", func(t *testing.T) {
		snapshotSystem := NewSystem(snapshotConstructor, logger)

		assert.Empty(t, snapshotSystem.GetSnapshot(epochNumber), "Expected empty result")
	})

	t.Run("PositiveSnapshotDoesNotExist", func(t *testing.T) {
		snapshotSystem := NewSystem(snapshotConstructor, logger)
		snapshot, err := snapshotSystem.PrepareSnapshot(epochNumber)

		assert.NoError(t, err, "Unexpected error")
		assert.Equal(t, snapshot, snapshotSystem.GetSnapshot(epochNumber), "Result differs from expected")
	})
}

func TestDeleteSnapshot(t *testing.T) {
	t.Run("PositiveSnapshotExists", func(t *testing.T) {
		snapshotSystem := NewSystem(snapshotConstructor, logger)
		AddSnapshotAtEpoch(snapshotSystem, epochNumber)

		snapshotSystem.DeleteSnapshot(epochNumber)

		assert.Empty(t, snapshotSystem.GetSnapshot(epochNumber), "Expected result to be empty")
	})

	t.Run("PositiveSnapshotDoesNotExist", func(t *testing.T) {
		snapshotSystem := NewSystem(snapshotConstructor, logger)
		snapshotSystem.DeleteSnapshot(epochNumber)

		assert.Empty(t, snapshotSystem.GetSnapshot(epochNumber), "Expected result to be empty")
	})
}

func TestClearOldSnapshots(t *testing.T) {
	t.Run("PositiveSnapshotCountIsLessThanMax", func(t *testing.T) {
		snapshotSystem := NewSystem(snapshotConstructor, logger)

		before := snapshotSystem.snapshotEpoch
		snapshotSystem.ClearOldSnapshots()

		assert.Equal(t, before, snapshotSystem.snapshotEpoch, "Didn't expect snapshot count to change")
	})

	t.Run("PositiveSnapshotCountIsBiggerThanMax", func(t *testing.T) {
		snapshotSystem := NewSystem(snapshotConstructor, logger)

		for i := 0; i < MaxEpochs+3; i++ {
			AddSnapshotAtEpoch(snapshotSystem, epochNumber+uint64(i))
		}

		snapshotSystem.ClearOldSnapshots()

		expected := TrimOldSnapshots(snapshotSystem.snapshotEpoch)
		assert.Equal(t, expected, snapshotSystem.snapshotEpoch, "Result differs from expected")
	})
}

func TestMinerListAtEpoch(t *testing.T) {
	t.Run("PositiveSnapshotPresentWithMiners", func(t *testing.T) {
		snapshotSystem := NewSystem(snapshotConstructor, logger)
		AddSnapshotAtEpoch(snapshotSystem, epochNumber)

		miners := GetRandomAddresses(3)
		snapshotSystem.snapshotEpoch[epochNumber].minerList = &MinerList{
			miners: miners,
			hash:   Hash,
		}

		assert.Equal(t, miners, snapshotSystem.MinerListAtEpoch(epochNumber), "Result differs from expected")
	})

	t.Run("PositiveSnapshotPresentNoMiners", func(t *testing.T) {
		snapshotSystem := NewSystem(snapshotConstructor, logger)
		AddSnapshotAtEpoch(snapshotSystem, epochNumber)

		assert.Empty(t, snapshotSystem.MinerListAtEpoch(epochNumber), "Expected result to be empty")
	})

	t.Run("PositiveSnapshotNotPresent", func(t *testing.T) {
		snapshotSystem := NewSystem(snapshotConstructor, logger)

		assert.Empty(t, snapshotSystem.MinerListAtEpoch(epochNumber), "Expected result to be empty")
	})
}
