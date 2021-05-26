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

// nolint:dupl,nakedret
package transfer

import (
	"context"
	"sync"
	"time"

	"github.com/pkg/errors"
	bb "github.com/simcord-llc/bitbon-system-blockchain/bitbon"
	loggerContext "github.com/simcord-llc/bitbon-system-blockchain/bitbon/context"
	"github.com/simcord-llc/bitbon-system-blockchain/log"
	"github.com/simcord-llc/bitbon-system-blockchain/zksnark"
)

const (
	// timeout to wait for transaction receipt
	updateKeyPeriod = 5 * time.Minute
	constraintsLen  = 1000

	batchSearchSize = 100
	batchExpireSize = 10
)

type Manager struct {
	bitbon    *bb.Bitbon
	encryptor bb.PkEncryptor
	logger    log.Logger

	keyMux sync.RWMutex // protect pk, vk
	pk, vk []byte       // keys needed fro zk-snark generation

	ready chan struct{} // ready channel is closed when Manager is ready to use
	done  chan struct{} // done channel used to notify about quiting
}

// NewTransferManager create a new TransferManager instance.
func NewTransferManager(b *bb.Bitbon, encryptor bb.PkEncryptor) (*Manager, error) {
	tm := &Manager{
		bitbon:    b,
		encryptor: encryptor,
		logger: loggerContext.LoggerFromContext(loggerContext.NewLoggerContext(context.Background(),
			log.New("src", "TransferManager"))),
		ready: make(chan struct{}),
		done:  make(chan struct{}),
	}
	// run obtain keys
	tm.obtainKeys()
	// check zk-snark works fine
	log.Warn("Checking ZK-snark works...")
	err := tm.checkZKSnark()
	if err == nil {
		log.Warn("ZK-snark works fine. Continue...")
	}
	return tm, err
}

// Stops manager
func (tm *Manager) Stop() error {
	if tm.Stopped() {
		return ErrAlreadyStopped
	}
	close(tm.done)
	return nil
}

func (tm *Manager) Ready() bool {
	select {
	case <-tm.ready:
		return true
	default:
		return false
	}
}

func (tm *Manager) Stopped() bool {
	select {
	case <-tm.done:
		return true
	default:
		return false
	}
}

func (tm *Manager) WaitReady(exit chan struct{}) bool {
	select {
	case <-exit:
		return false
	case <-tm.done:
		return false
	case <-tm.ready:
		return true
	}
}

func (tm *Manager) makeReady() {
	log.Warn("Bitbon transfer manager became ready")
	close(tm.ready)
}

func (tm *Manager) obtainKeys() {
	// generate keys first time
	tm.setKeys(zksnark.GenerateKeys(constraintsLen))
	log.Warn("Transfer manager PK and VK successfully obtained")
	// make ready TransferManager
	tm.makeReady()

	// periodically update zk-snark keys
	go func() {
		ticker := time.NewTicker(updateKeyPeriod)

		for {
			select {
			case tick := <-ticker.C:
				ticker.Stop()

				log.Warn("Update transfer manager keys ticker fired", "period", updateKeyPeriod.String(), "ticker time", tick)
				tm.setKeys(zksnark.GenerateKeys(constraintsLen))

				// reset ticker
				ticker = time.NewTicker(updateKeyPeriod)
			case <-tm.done:
				ticker.Stop()
				log.Warn("Update transfer manager keys loop terminated.")
				return
			}
		}
	}()
}

func (tm *Manager) checkZKSnark() error {
	protectionHash := tm.getProtectionHash("testTransferID", "testProtectionCode")
	log.Warn("protectionHash successfully obtained", "protectionHash", protectionHash)
	var pk, vk []byte
	pk, vk = tm.getKeys()
	log.Warn("zk-snrak keys obtaioned (pk, vk)", "len(pk)", len(pk), "len(vk)", len(vk))
	proof, err := zksnark.Prove(protectionHash, pk, constraintsLen)
	if err != nil {
		return errors.Wrap(err, "error generating ZK-snark params")
	}
	if len(proof) == 0 {
		return ErrZkSnarkProofEmpty
	}
	return nil
}

func (tm *Manager) setKeys(pk, vk []byte) {
	tm.keyMux.Lock()
	defer tm.keyMux.Unlock()
	tm.pk, tm.vk = pk, vk
}
