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
package agent

import "time"

// Bitbon constants
const (
	ProtocolVersionStr = "1.0.0"       // The same, as agent string
	ProtocolName       = "bitbonagent" // Nickname of the protocol in geth

	DefaultJournalName    = ".bitbon.agent.journal.LDB"
	DefaultJournalTimeout = 10 * time.Minute
	DefaultConfirmTimeout = 10 * time.Second
	DefaultPublishTimeout = 3 * time.Second

	// Rabbit routing keys
	assetboxInfoChangedRK            = "r.events.bitbon-node.AssetboxInfoChanged"
	assetboxInfoDeletedRK            = "r.events.bitbon-node.AssetboxInfoDeleted"
	assetboxBalanceChangedRK         = "r.events.bitbon-node.AssetboxBalanceChanged"
	transactionsChangedRK            = "r.events.bitbon-node.TransactionsMined"
	transferExpiredRK                = "r.events.bitbon-node.TransferExpired"
	blocksMinedRK                    = "r.events.bitbon-node.BlocksMined"
	feeValueChangedRK                = "r.events.bitbon-node.FeeValueChanged"
	feeDistributionSettingsChangedRK = "r.events.bitbon-node.FeeDistributionSettingsChanged"
	exceptionalAccountsChangedRK     = "r.events.bitbon-node.ExceptionalAccountsChanged"
)

func IsAllowed() bool {
	return true
}
