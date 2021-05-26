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

package sqlite

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/storage"
)

// to ensure that SettingsStorage satisfies storage.SettingsStorage interface. If not there will be compile time error
var _ = storage.SettingsStorage(&SettingsStorage{})

const (
	// settings table related queries
	createSettingsStatement = `CREATE TABLE IF NOT EXISTS settings (
		key				VARCHAR PRIMARY KEY,
		value			VARCHAR
	)`
	insertSettingStatement = `INSERT INTO settings (key, value) values (?, ?)`
	updateSettingStatement = `UPDATE settings SET value=? WHERE key=?`
	existSettingStatement  = `SELECT EXISTS(SELECT 1 FROM settings WHERE key=? LIMIT 1)`
	readSettingStatement   = `SELECT value FROM settings WHERE key=? LIMIT 1`
)

type SettingsStorage struct {
	*Store
}

func NewSettingsStorage(store *Store) *SettingsStorage {
	return &SettingsStorage{Store: store}
}

func (s *SettingsStorage) Init() (err error) {
	if _, err = s.db.Exec(createSettingsStatement); err != nil {
		return errors.Wrap(err, "failed to create assetbox table")
	}
	return nil
}

func (s *SettingsStorage) SaveSetting(key, value string) (err error) {
	tx, err := s.db.BeginTxx(context.Background(), nil)
	if err != nil {
		return errors.Wrap(err, "failed to begin setting transaction")
	}
	defer func() {
		if err != nil {
			tx.Rollback() // nolint:errcheck
			return
		}
		tx.Commit() // nolint:errcheck
	}()

	exists, err := s.existSetting(key, tx)
	if err != nil {
		return errors.Wrap(err, "failed to check setting for existence")
	}

	if exists {
		if err = s.updateSetting(key, value, tx); err != nil {
			return errors.Wrap(err, "failed to update setting")
		}
		return nil
	}
	if err = s.insertSetting(key, value, tx); err != nil {
		return errors.Wrap(err, "failed to insert setting")
	}
	return nil
}

func (s *SettingsStorage) GetSetting(key string) (value string, err error) {
	tx, err := s.db.BeginTxx(context.Background(), nil)
	if err != nil {
		return "", errors.Wrap(err, "failed to begin setting transaction")
	}
	defer func() {
		if err != nil {
			tx.Rollback() // nolint:errcheck
			return
		}
		tx.Commit() // nolint:errcheck
	}()
	return s.getSetting(key, tx)
}

func (s *SettingsStorage) insertSetting(key, value string, tx *sqlx.Tx) (err error) {
	_, err = tx.Exec(insertSettingStatement, key, value)
	return
}

func (s *SettingsStorage) updateSetting(key, value string, tx *sqlx.Tx) (err error) {
	_, err = tx.Exec(updateSettingStatement, value, key)
	return
}

func (s *SettingsStorage) existSetting(key string, tx *sqlx.Tx) (bool, error) {
	var exists bool
	err := tx.QueryRow(existSettingStatement, key).Scan(&exists)
	return exists, err
}

func (s *SettingsStorage) getSetting(key string, tx *sqlx.Tx) (value string, err error) {
	err = tx.QueryRow(readSettingStatement, key).Scan(&value)

	if err == sql.ErrNoRows {
		return "", storage.ErrNotFound
	}

	return value, err
}
