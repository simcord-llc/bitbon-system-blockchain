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

// Contains the Bibon constant definitions
package bitbon

import (
	"crypto/ecdsa"

	"github.com/pkg/errors"
	"github.com/simcord-llc/bitbon-system-blockchain/accounts/keystore"
	"github.com/simcord-llc/bitbon-system-blockchain/common"
)

// Bitbon constants
const (
	ProtocolVersionStr     = "1.0.0"                  // The same, as a string
	ProtocolName           = "bitbon"                 // Nickname of the protocol in geth
	PrivateStorageFileName = ".bitbon.storage.sqlite" // File name for storage

	MaxPrepareAssetboxBatchSize     = 100
	DefaultPrepareAssetboxBatchSize = 10
	MaxPrepareAssetboxRetries       = 2
)

func DecryptAssetboxWallet(walletEncrypted, passPhraseEncrypted []byte, password string, decryptFn func(data, key []byte) ([]byte, error)) (*keystore.Key, error) {
	if walletEncrypted == nil {
		return nil, errors.New("empty wallet")
	}
	if len(passPhraseEncrypted) == 0 {
		return nil, errors.New("empty pass phrase")
	}

	passPhrase, err := decryptFn(passPhraseEncrypted, []byte(password))
	if err != nil {
		return nil, err
	}

	wallet, err := decryptFn(walletEncrypted, []byte(password))
	if err != nil {
		return nil, err
	}

	key, err := keystore.DecryptKey(wallet, string(passPhrase))
	if err != nil {
		return nil, errors.Wrap(err, "error decrypting assetbox wallet")
	}

	return key, nil
}

func DecryptPrivateKeyForAssetbox(address common.Address, walletEncrypted, passPhraseEncrypted []byte, password string, decryptFn func(data, key []byte) ([]byte, error)) (*ecdsa.PrivateKey, error) {
	key, err := DecryptAssetboxWallet(walletEncrypted, passPhraseEncrypted, password, decryptFn)
	if err != nil {
		return nil, err
	}

	if address != key.Address {
		return nil, errors.New("assetbox is not owner of wallet")
	}

	return key.PrivateKey, nil
}
