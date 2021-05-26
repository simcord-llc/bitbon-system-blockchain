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

package assetbox

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"time"

	"github.com/pborman/uuid"
	"github.com/pkg/errors"
	"github.com/simcord-llc/bitbon-system-blockchain/accounts/keystore"
	bb "github.com/simcord-llc/bitbon-system-blockchain/bitbon"
	loggerContext "github.com/simcord-llc/bitbon-system-blockchain/bitbon/context"
	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/models"
	"github.com/simcord-llc/bitbon-system-blockchain/crypto"
	"github.com/simcord-llc/bitbon-system-blockchain/log"
)

type Manager struct {
	bitbon    *bb.Bitbon
	encryptor bb.PkEncryptor
	logger    log.Logger
}

// AssetboxManager create a new AssetboxManager instance.
func NewAssetboxManager(b *bb.Bitbon, encryptor bb.PkEncryptor) *Manager {
	return &Manager{
		bitbon:    b,
		encryptor: encryptor,
		logger: loggerContext.LoggerFromContext(loggerContext.NewLoggerContext(context.Background(),
			log.New("src", "Manager"))),
	}
}

func (am *Manager) newUnboundAssetbox() (assetbox *models.Assetbox, err error) {
	defer func() {
		if p := recover(); p != nil {
			err = fmt.Errorf("(newUnboundAssetbox) panic recovered : %v", p)
		}
	}()

	// If not loaded, generate random.
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate assetbox private key")
	}

	// Create the keyfile object with am random UUID.
	id := uuid.NewRandom()
	address := crypto.PubkeyToAddress(privateKey.PublicKey)
	key := &keystore.Key{
		ID:         id,
		Address:    address,
		PrivateKey: privateKey,
	}

	// Encrypt key with passphrase.
	passphrase := uuid.NewRandom().String()
	keyjson, err := keystore.EncryptKey(key, passphrase, keystore.LightScryptN, keystore.LightScryptP)
	if err != nil {
		return nil, errors.Wrap(err, "error encrypting key")
	}

	createdTime := time.Now()

	assetbox = &models.Assetbox{
		ServiceID:  am.bitbon.GetServiceID(),
		Wallet:     keyjson,
		PassPhrase: []byte(passphrase),
		Address:    address,
		CreatedAt:  &createdTime,
		UpdatedAt:  &createdTime,
		Pk:         privateKey,
	}
	return assetbox, nil
}

func (am *Manager) deleteAssetbox(ctx context.Context, pk *ecdsa.PrivateKey) error {
	// delete assetbox in blockchain
	err := am.bitbon.GetContractsManager().DeleteAssetboxInfo(ctx, pk)
	if err != nil {
		return errors.Wrap(err, "error deleting assetbox info from blockchain")
	}
	am.logger.Debug("Deleted assetbox from blockchain")
	return nil
}
