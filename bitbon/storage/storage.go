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

package storage

import (
	"errors"

	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/storage/models"
)

var ErrNotFound = errors.New("not found")
var ErrAlreadyExists = errors.New("already exists")

type Transaction interface {
	Commit() error
	Rollback() error
}

type Storage interface {
	Init() error
	BeginTx() (Transaction, error)
	GetOriginStorage() interface{}
	Close() error
}

type SettingsStorage interface {
	Storage
	SaveSetting(key, value string) error
	GetSetting(key string) (string, error)
}

type AbiStorage interface {
	Storage
	SaveAbiInfo(tx Transaction, abiInfo *models.AbiInfo) error
	EditAbiInfo(tx Transaction, abiInfo *models.AbiInfo) error
	DeleteAbiInfoByAddress(tx Transaction, address string) error
	GetAbiJSON(tx Transaction, address string) (string, error)
	GetAbiVersion(tx Transaction, address string) (string, error)
	GetFullAbiInfoByVersionAndType(tx Transaction, version string, contractType string) (*models.AbiInfo, error)
	GetFullAbiInfo(tx Transaction, address string) (*models.AbiInfo, error)
	GetAbiAddressesByVersion(tx Transaction, version string) ([]string, error)
	GetAbiInfosByVersion(tx Transaction, version string) ([]*models.AbiInfo, error)
}
