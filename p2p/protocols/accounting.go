// Copyright Simcord LLC
// This file is part of the Bitbon System Blockchain Node library.
// This file has been modified, you can find the original version by following the link
// <https://github.com/ethereum/go-ethereum>
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

package protocols

import (
	"time"

	"github.com/simcord-llc/bitbon-system-blockchain/metrics"
)

// define some metrics
var (
	// All metrics are cumulative

	// total amount of units credited
	mBalanceCredit = metrics.NewRegisteredCounterForced("account/balance/credit", metrics.AccountingRegistry)
	// total amount of units debited
	mBalanceDebit = metrics.NewRegisteredCounterForced("account/balance/debit", metrics.AccountingRegistry)
	// total amount of bytes credited
	mBytesCredit = metrics.NewRegisteredCounterForced("account/bytes/credit", metrics.AccountingRegistry)
	// total amount of bytes debited
	mBytesDebit = metrics.NewRegisteredCounterForced("account/bytes/debit", metrics.AccountingRegistry)
	// total amount of credited messages
	mMsgCredit = metrics.NewRegisteredCounterForced("account/msg/credit", metrics.AccountingRegistry)
	// total amount of debited messages
	mMsgDebit = metrics.NewRegisteredCounterForced("account/msg/debit", metrics.AccountingRegistry)
	// how many times local node had to drop remote peers
	mPeerDrops = metrics.NewRegisteredCounterForced("account/peerdrops", metrics.AccountingRegistry)
	// how many times local node overdrafted and dropped
	mSelfDrops = metrics.NewRegisteredCounterForced("account/selfdrops", metrics.AccountingRegistry)
)

// PricedMessage defines how a message type identifies itself as to be accounted
type PricedMessage interface {
	// Return the Price for a message
	Price() *Price
}

// Payer is the base type to define who pays in an exchange between peers
type Payer bool

const (
	// Sender declares that a message needs to be paid by the sender of the message
	Sender = Payer(true)
	// Receiver declares that a message needs to be paid by the receiver of the message
	Receiver = Payer(false)
)

// Price represents the costs of a message
type Price struct {
	Value   uint64
	PerByte bool // True if the price is per byte or for unit
	Payer   Payer
}

// For gives back the price for a message
// A protocol provides the message price in absolute value
// This method then returns the correct signed amount,
// depending on who pays, which is identified by the `payer` argument:
// Sending will pass a `Sender` payer, receiving will pass the `Receiver` argument.
// Thus: If Sending and sender pays, amount negative, otherwise positive
// If Receiving, and receiver pays, amount negative, otherwise positive
func (p *Price) For(payer Payer, size uint32) int64 {
	price := p.Value
	if p.PerByte {
		price *= uint64(size)
	}
	if p.Payer == payer {
		return 0 - int64(price)
	}
	return int64(price)
}

// Balance is the actual accounting instance
// Balance defines the operations needed for accounting
// Implementations internally maintain the balance for every peer
type Balance interface {
	// Adds amount to the local balance with remote node `peer`;
	// positive amount = credit local node
	// negative amount = debit local node
	Add(amount int64, peer *Peer) error
	// Check is a dry-run for the Add operation:
	// As the accounting takes place **after** the actual send/receive operation happens,
	// we want to make sure that that operation would not result in any problem
	Check(amount int64, peer *Peer) error
}

// Accounting implements the Hook interface
// It interfaces to the balances through the Balance interface
type Accounting struct {
	Balance // interface to accounting logic
}

// NewAccounting creates a new instance of Accounting
func NewAccounting(balance Balance) *Accounting {
	ah := &Accounting{
		Balance: balance,
	}
	return ah
}

// SetupAccountingMetrics uses a separate registry for p2p accounting metrics;
// this registry should be independent of any other metrics as it persists at different endpoints.
// It also starts the persisting go-routine which
// at the passed interval writes the metrics to a LevelDB
func SetupAccountingMetrics(reportInterval time.Duration, path string) *AccountingMetrics {
	// create the DB and start persisting
	return NewAccountingMetrics(metrics.AccountingRegistry, reportInterval, path)
}

// Apply takes a peer, the signed cost for the local node and the msg size and credits/debits local node using balance interface
func (ah *Accounting) Apply(peer *Peer, costToLocalNode int64, size uint32) error {
	// do the accounting
	err := ah.Add(costToLocalNode, peer)
	// record metrics: just increase counters for user-facing metrics
	ah.doMetrics(costToLocalNode, size, err)
	return err
}

// Validate calculates the cost for the local node sending or receiving a msg to/from a peer querying the message for its price.
// It returns either the signed cost for the local node as int64 or an error, signaling that the accounting operation would fail
// (no change has been applied at this point)
func (ah *Accounting) Validate(peer *Peer, size uint32, msg interface{}, payer Payer) (int64, error) {
	// get the price for a message (by querying the message type via the PricedMessage interface)
	var pricedMessage PricedMessage
	var ok bool
	// if the msg implements `Price`, it is an accounted message
	if pricedMessage, ok = msg.(PricedMessage); !ok {
		return 0, nil
	}
	// evaluate the price for receiving messages
	costToLocalNode := pricedMessage.Price().For(payer, size)
	// check that the operation would perform correctly
	err := ah.Check(costToLocalNode, peer)
	if err != nil {
		// signal to caller that the operation would fail
		return 0, err
	}
	return costToLocalNode, nil
}

// record some metrics
// this is not an error handling. `err` is returned by both `Send` and `Receive`
// `err` will only be non-nil if a limit has been violated (overdraft), in which case the peer has been dropped.
// if the limit has been violated and `err` is thus not nil:
//   * if the price is positive, local node has been credited; thus `err` implicitly signals the REMOTE has been dropped
//   * if the price is negative, local node has been debited, thus `err` implicitly signals LOCAL node "overdraft"
func (ah *Accounting) doMetrics(price int64, size uint32, err error) {
	if price > 0 {
		mBalanceCredit.Inc(price)
		mBytesCredit.Inc(int64(size))
		mMsgCredit.Inc(1)
		if err != nil {
			// increase the number of times a remote node has been dropped due to "overdraft"
			mPeerDrops.Inc(1)
		}
	} else {
		mBalanceDebit.Inc(price)
		mBytesDebit.Inc(int64(size))
		mMsgDebit.Inc(1)
		if err != nil {
			// increase the number of times the local node has done an "overdraft" in respect to other nodes
			mSelfDrops.Inc(1)
		}
	}
}
