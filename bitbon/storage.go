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

// nolint:nakedret,dupl
package bitbon

import (
	"fmt"

	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/contracts/contract_snapshot"
	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/dto"

	"strings"

	_ "github.com/mattn/go-sqlite3" // nolint
	"github.com/pkg/errors"
	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/storage"
	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/storage/sqlite"
	"github.com/simcord-llc/bitbon-system-blockchain/common"
	"github.com/simcord-llc/bitbon-system-blockchain/log"
)

const (
	SettingContractStorageAddress = "contract_storage_addr"
)

type Storage struct {
	ss   storage.SettingsStorage
	abis storage.AbiStorage
}

func NewStorage(dbPath string) (dbStorage *Storage, err error) {
	store, err := sqlite.NewStore(dbPath)
	if err != nil {
		err = errors.Wrap(err, "failed to create sqlite store")
		return
	}

	dbStorage = &Storage{
		ss:   initSettingsStorage(store),
		abis: initAbiStorage(store),
	}
	return
}

func initAbiStorage(store *sqlite.Store) (is *sqlite.AbiStorage) {
	var err error
	is = sqlite.NewAbiStorage(store)
	defer func() {
		if err != nil {
			is.Close()
		}
	}()
	if err = is.Init(); err != nil {
		err = errors.Wrap(err, "failed to init abi storage")
		return
	}

	return
}

func initSettingsStorage(store *sqlite.Store) (ss *sqlite.SettingsStorage) {
	var err error
	ss = sqlite.NewSettingsStorage(store)
	defer func() {
		if err != nil {
			ss.Close()
		}
	}()
	if err = ss.Init(); err != nil {
		err = errors.Wrap(err, "failed to init settings storage")
		return
	}
	return
}

func (s *Storage) Close() error {
	if s.ss != nil {
		if err := s.ss.Close(); err != nil {
			log.Error("Error closing settings storage", "err", err)
		}
	}
	if s.abis != nil {
		if err := s.abis.Close(); err != nil {
			log.Error("Error closing abi storage", "err", err)
		}
	}
	return nil
}

// Settings methods

func (s *Storage) SaveSettingContractStorageAddress(addr common.Address) error {
	if err := s.ss.SaveSetting(SettingContractStorageAddress, strings.ToLower(addr.Hex())); err != nil {
		return errors.Wrap(err, fmt.Sprintf("error saving setting %q", SettingContractStorageAddress))
	}
	return nil
}

func (s *Storage) GetSettingContractStorageAddress() (common.Address, error) {
	value, err := s.ss.GetSetting(SettingContractStorageAddress)
	if err == storage.ErrNotFound {
		return common.Address{}, nil
	}
	if err != nil {
		return common.Address{}, errors.Wrap(err, fmt.Sprintf("error getting setting %q", SettingContractStorageAddress))
	}
	return common.HexToAddress(value), nil
}

// ABI Methods

func (s *Storage) BeginAbiTransaction() (storage.Transaction, error) {
	return s.abis.BeginTx()
}

func (s *Storage) SaveAbiInfo(tx storage.Transaction, abiInfo *dto.AbiInfo) (err error) {
	if tx == nil {
		tx, err = s.BeginAbiTransaction()
		if err != nil {
			return err
		}
		defer func() {
			if err != nil {
				tx.Rollback() // nolint:errcheck
				return
			}
			tx.Commit() // nolint:errcheck
		}()
	}

	if !contract_snapshot.ContractTypeExists(abiInfo.ContractType) {
		return errors.Errorf("unknown contract type %s", abiInfo.ContractType)
	}

	if !dto.ContractVersionExists(abiInfo.Version) {
		return errors.Errorf("unknown contract version %s", abiInfo.Version)
	}

	return s.abis.SaveAbiInfo(tx, abiInfo.ToStorageAbiInfo())
}

func (s *Storage) EditAbiInfo(tx storage.Transaction, abiInfo *dto.AbiInfo) (err error) {
	if tx == nil {
		tx, err = s.BeginAbiTransaction()
		if err != nil {
			return err
		}
		defer func() {
			if err != nil {
				tx.Rollback() // nolint:errcheck
				return
			}
			tx.Commit() // nolint:errcheck
		}()
	}

	if !contract_snapshot.ContractTypeExists(abiInfo.ContractType) {
		return errors.Errorf("unknown contract type %s", abiInfo.ContractType)
	}

	if !dto.ContractVersionExists(abiInfo.Version) {
		return errors.Errorf("unknown contract version %s", abiInfo.Version)
	}

	return s.abis.EditAbiInfo(tx, abiInfo.ToStorageAbiInfo())
}

func (s *Storage) DeleteAbiInfoByAddress(tx storage.Transaction, address common.Address) (err error) {
	if tx == nil {
		tx, err = s.BeginAbiTransaction()
		if err != nil {
			return err
		}
		defer func() {
			if err != nil {
				tx.Rollback() // nolint:errcheck
				return
			}
			tx.Commit() // nolint:errcheck
		}()
	}

	return s.abis.DeleteAbiInfoByAddress(tx, dto.AbiInfoAddressToStorage(address))
}

func (s *Storage) GetAbiJSON(tx storage.Transaction, address common.Address) (abiJSON string, err error) {
	if tx == nil {
		tx, err = s.BeginAbiTransaction()
		if err != nil {
			return "", err
		}
		defer func() {
			if err != nil {
				tx.Rollback() // nolint:errcheck
				return
			}
			tx.Commit() // nolint:errcheck
		}()
	}

	return s.abis.GetAbiJSON(tx, dto.AbiInfoAddressToStorage(address))
}

func (s *Storage) GetAbiVersion(tx storage.Transaction,
	address common.Address) (version dto.ContractVersion, err error) {
	if tx == nil {
		tx, err = s.BeginAbiTransaction()
		if err != nil {
			return "", err
		}
		defer func() {
			if err != nil {
				tx.Rollback() // nolint:errcheck
				return
			}
			tx.Commit() // nolint:errcheck
		}()
	}

	v, err := s.abis.GetAbiVersion(tx, dto.AbiInfoAddressToStorage(address))
	if err != nil {
		return "", errors.Wrap(err, "unable to get version")
	}

	return dto.ContractVersion(v), err
}

func (s *Storage) GetFullAbiInfo(tx storage.Transaction, address common.Address) (abiInfo *dto.AbiInfo, err error) {
	if tx == nil {
		tx, err = s.BeginAbiTransaction()
		if err != nil {
			return nil, err
		}
		defer func() {
			if err != nil {
				tx.Rollback() // nolint:errcheck
				return
			}
			tx.Commit() // nolint:errcheck
		}()
	}

	res, err := s.abis.GetFullAbiInfo(tx, dto.AbiInfoAddressToStorage(address))
	if err != nil {
		return nil, errors.Wrap(err, "unable to get full abi info")
	}

	return dto.AbiInfoFromStorage(res), err
}
func (s *Storage) GetFullAbiInfoByVersionAndType(tx storage.Transaction, version dto.ContractVersion,
	contractType contract_snapshot.ContractType) (abiInfo *dto.AbiInfo, err error) {
	if tx == nil {
		tx, err = s.BeginAbiTransaction()
		if err != nil {
			return nil, err
		}
		defer func() {
			if err != nil {
				tx.Rollback() // nolint:errcheck
				return
			}
			tx.Commit() // nolint:errcheck
		}()
	}

	if !dto.ContractVersionExists(version) {
		return nil, errors.Errorf("unknown contract version %s", version)
	}

	res, err := s.abis.GetFullAbiInfoByVersionAndType(tx, string(version), string(contractType))
	if err != nil {
		return nil, errors.Wrap(err, "unable to get full abi info")
	}

	return dto.AbiInfoFromStorage(res), err
}

func (s *Storage) GetAbiAddressesByVersion(tx storage.Transaction,
	version dto.ContractVersion) (addresses []common.Address, err error) {
	if tx == nil {
		tx, err = s.BeginAbiTransaction()
		if err != nil {
			return nil, err
		}
		defer func() {
			if err != nil {
				tx.Rollback() // nolint:errcheck
				return
			}
			tx.Commit() // nolint:errcheck
		}()
	}

	as, err := s.abis.GetAbiAddressesByVersion(tx, string(version))
	if err != nil {
		return nil, errors.Wrap(err, "unable to get addresses by version")
	}

	res := make([]common.Address, len(as))
	for i, a := range as {
		res[i] = dto.StorageToAbiInfoAddress(a)
	}

	return res, nil
}

func (s *Storage) GetAbiInfosByVersion(tx storage.Transaction,
	version dto.ContractVersion) (abiInfos []*dto.AbiInfo, err error) {
	if tx == nil {
		tx, err = s.BeginAbiTransaction()
		if err != nil {
			return nil, err
		}
		defer func() {
			if err != nil {
				tx.Rollback() // nolint:errcheck
				return
			}
			tx.Commit() // nolint:errcheck
		}()
	}

	ais, err := s.abis.GetAbiInfosByVersion(tx, string(version))
	if err != nil {
		return nil, errors.Wrap(err, "unable to get full abi info")
	}

	abiInfos = make([]*dto.AbiInfo, len(ais))
	for i, ai := range ais {
		abiInfos[i] = dto.AbiInfoFromStorage(ai)
	}

	return abiInfos, nil
}
