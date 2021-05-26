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

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/storage"
)

type Store struct {
	db *sqlx.DB
}

func NewStore(dbPath string) (*Store, error) {
	connStr := dbPath + "?cache=shared&mode=rwc&_busy_timeout=5000&_journal_mode=WAL"
	d, err := sqlx.Open("sqlite3", connStr)
	if err != nil {
		return nil, errors.Wrap(err, "failed to open storage new database")
	}
	if err = d.Ping(); err != nil {
		return nil, errors.Wrap(err, "failed to ping storage database")
	}
	d.SetMaxOpenConns(1)
	return &Store{db: d}, nil
}

func (s *Store) GetOriginStorage() interface{} {
	return s.db
}
func (s *Store) BeginTx() (storage.Transaction, error) {
	return s.db.BeginTxx(context.Background(), nil)
}

func (s *Store) Close() error {
	return s.db.Close()
}
