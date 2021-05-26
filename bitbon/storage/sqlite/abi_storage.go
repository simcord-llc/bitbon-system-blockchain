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
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/storage"
	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/storage/models"
	"github.com/simcord-llc/bitbon-system-blockchain/log"
)

var _ = storage.AbiStorage(&AbiStorage{})

var errAssertTransaction = errors.New("failed to assert Transaction")

const (
	createAbiStatement = `CREATE TABLE IF NOT EXISTS abi (
		address	VARCHAR PRIMARY KEY,
    	abi	VARCHAR
	)`
	addVersionFieldStatement      = `ALTER TABLE abi ADD version VARCHAR NULL;`
	addContractTypeFieldStatement = `ALTER TABLE abi ADD contract_type VARCHAR NULL;`

	getAbiAddressesByVersion       = `SELECT address FROM abi WHERE version = ?;`
	getAbiInfosByVersion           = `SELECT abi, address, version, contract_type FROM abi WHERE version = ?;`
	getAbiJSON                     = `SELECT abi FROM abi WHERE address = ?;`
	getAbiVersion                  = `SELECT version FROM abi WHERE address = ?;`
	getFullAbiInfo                 = `SELECT abi, address, version, contract_type FROM abi WHERE address = ?;`
	getFullAbiInfoByVersionAndType = `SELECT abi, address, version, contract_type FROM abi 
										WHERE version=? AND contract_type=?;`
	saveAbiInfo            = `INSERT INTO abi (address, abi, version, contract_type) VALUES(?, ?, ?, ?);`
	editAbiInfo            = `UPDATE abi SET abi=?, version=?, contract_type=? WHERE address=?;`
	deleteAbiInfoByAddress = `DELETE FROM abi WHERE address=?;`
)

const (
	defaultAllocationCapacity = 10
)

type AbiStorage struct {
	*Store
}

func NewAbiStorage(store *Store) *AbiStorage {
	return &AbiStorage{Store: store}
}

func (s *AbiStorage) Close() error {
	return s.Store.Close()
}

func (s *AbiStorage) Init() (err error) {
	if _, err = s.db.Exec(createAbiStatement); err != nil {
		return errors.Wrap(err, "failed to create abi table")
	}

	if _, err = s.db.Exec(addVersionFieldStatement); err != nil {
		log.Warn("failed to add version field", "error", err)
	}

	if _, err = s.db.Exec(addContractTypeFieldStatement); err != nil {
		log.Warn("failed to add contract_type field", "error", err)
	}

	return nil
}

func (s *AbiStorage) SaveAbiInfo(tx storage.Transaction, abiInfo *models.AbiInfo) error {
	txx, ok := tx.(*sqlx.Tx)
	if !ok {
		return errAssertTransaction
	}

	_, err := s.GetFullAbiInfoByVersionAndType(tx, abiInfo.Version, abiInfo.ContractType)
	if err != nil && err != storage.ErrNotFound {
		return errors.Errorf("contract with version %s and type %s already exists or unable to check: %s",
			abiInfo.Version, abiInfo.ContractType, err.Error())
	}

	return s.saveAbiInfo(txx, abiInfo.Address, abiInfo.Abi, abiInfo.Version, abiInfo.ContractType)
}

func (s *AbiStorage) EditAbiInfo(tx storage.Transaction, abiInfo *models.AbiInfo) error {
	txx, ok := tx.(*sqlx.Tx)
	if !ok {
		return errAssertTransaction
	}

	res, err := s.GetFullAbiInfoByVersionAndType(tx, abiInfo.Version, abiInfo.ContractType)
	if err != nil {
		if err == storage.ErrNotFound {
			return errors.Errorf("contract with version %s and type %s already exists with address %s",
				abiInfo.Version, abiInfo.ContractType, res.Address)
		}

		return errors.Errorf("unable to check if contract with version %s and type %s already exists",
			abiInfo.Version, abiInfo.ContractType)
	}

	return s.editAbiInfo(txx, abiInfo)
}

func (s *AbiStorage) DeleteAbiInfoByAddress(tx storage.Transaction, address string) error {
	txx, ok := tx.(*sqlx.Tx)
	if !ok {
		return errAssertTransaction
	}

	return s.deleteAbiInfoByAddress(txx, address)
}

func (s *AbiStorage) GetAbiJSON(tx storage.Transaction, address string) (string, error) {
	txx, ok := tx.(*sqlx.Tx)
	if !ok {
		return "", errAssertTransaction
	}

	return s.getAbiJSON(txx, address)
}

func (s *AbiStorage) GetAbiVersion(tx storage.Transaction, address string) (string, error) {
	txx, ok := tx.(*sqlx.Tx)
	if !ok {
		return "", errAssertTransaction
	}

	return s.getAbiVersion(txx, address)
}

func (s *AbiStorage) GetFullAbiInfoByVersionAndType(tx storage.Transaction,
	version, contractType string) (*models.AbiInfo, error) {
	txx, ok := tx.(*sqlx.Tx)
	if !ok {
		return nil, errAssertTransaction
	}

	return s.getFullAbiInfoByVersionAndType(txx, version, contractType)
}

func (s *AbiStorage) GetFullAbiInfo(tx storage.Transaction, address string) (*models.AbiInfo, error) {
	txx, ok := tx.(*sqlx.Tx)
	if !ok {
		return nil, errAssertTransaction
	}

	return s.getFullAbiInfo(txx, address)
}

func (s *AbiStorage) GetAbiAddressesByVersion(tx storage.Transaction, version string) ([]string, error) {
	txx, ok := tx.(*sqlx.Tx)
	if !ok {
		return nil, errAssertTransaction
	}

	return s.getAbiAddressesByVersion(txx, version)
}

func (s *AbiStorage) GetAbiInfosByVersion(tx storage.Transaction, version string) ([]*models.AbiInfo, error) {
	txx, ok := tx.(*sqlx.Tx)
	if !ok {
		return nil, errAssertTransaction
	}

	return s.getAbiInfosByVersion(txx, version)
}

func (s *AbiStorage) saveAbiInfo(tx *sqlx.Tx, address, abi, version, contractType string) error {
	_, err := tx.Exec(saveAbiInfo, address, abi, version, contractType)
	return err
}

func (s *AbiStorage) getAbiJSON(tx *sqlx.Tx, address string) (string, error) {
	var abiJSON string
	err := tx.QueryRowx(getAbiJSON, address).Scan(&abiJSON)

	if err == sql.ErrNoRows {
		return abiJSON, storage.ErrNotFound
	}

	return abiJSON, err
}

func (s *AbiStorage) getAbiVersion(tx *sqlx.Tx, address string) (string, error) {
	var abiVersion string
	err := tx.QueryRowx(getAbiVersion, address).Scan(&abiVersion)

	if err == sql.ErrNoRows {
		return abiVersion, storage.ErrNotFound
	}

	return abiVersion, err
}

func (s *AbiStorage) getFullAbiInfo(tx *sqlx.Tx, address string) (*models.AbiInfo, error) {
	fullAbiInfo := &models.AbiInfo{}
	err := tx.QueryRowx(getFullAbiInfo, address).StructScan(fullAbiInfo)

	if err == sql.ErrNoRows {
		return fullAbiInfo, storage.ErrNotFound
	}

	return fullAbiInfo, err
}

func (s *AbiStorage) getFullAbiInfoByVersionAndType(tx *sqlx.Tx,
	version, contractType string) (*models.AbiInfo, error) {
	fullAbiInfo := &models.AbiInfo{}
	err := tx.QueryRowx(getFullAbiInfoByVersionAndType, version, contractType).StructScan(fullAbiInfo)

	if err == sql.ErrNoRows {
		return fullAbiInfo, storage.ErrNotFound
	}

	return fullAbiInfo, err
}

func (s *AbiStorage) editAbiInfo(tx *sqlx.Tx, abiInfo *models.AbiInfo) error {
	_, err := tx.Exec(editAbiInfo, abiInfo.Abi, abiInfo.Version, abiInfo.ContractType, abiInfo.Address)
	return err
}

func (s *AbiStorage) deleteAbiInfoByAddress(tx *sqlx.Tx, address string) error {
	_, err := tx.Exec(deleteAbiInfoByAddress, address)
	return err
}

func (s *AbiStorage) getAbiAddressesByVersion(tx *sqlx.Tx, version string) ([]string, error) {
	addresses := make([]string, 0, defaultAllocationCapacity)

	err := tx.Select(&addresses, getAbiAddressesByVersion, version)
	if err == sql.ErrNoRows || len(addresses) == 0 {
		return addresses, storage.ErrNotFound
	}

	return addresses, err
}

func (s *AbiStorage) getAbiInfosByVersion(tx *sqlx.Tx, version string) ([]*models.AbiInfo, error) {
	abiInfos := make([]*models.AbiInfo, 0, defaultAllocationCapacity)

	err := tx.Select(&abiInfos, getAbiInfosByVersion, version)
	if err == sql.ErrNoRows || len(abiInfos) == 0 {
		return abiInfos, storage.ErrNotFound
	}

	return abiInfos, err
}
