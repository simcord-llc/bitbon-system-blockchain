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
	"errors"
)

// Textual version number of accounting API
const AccountingVersion = "1.0"

var errNoAccountingMetrics = errors.New("accounting metrics not enabled")

// AccountingApi provides an API to access account related information
type AccountingApi struct {
	metrics *AccountingMetrics
}

// NewAccountingApi creates a new AccountingApi
// m will be used to check if accounting metrics are enabled
func NewAccountingApi(m *AccountingMetrics) *AccountingApi {
	return &AccountingApi{m}
}

// Balance returns local node balance (units credited - units debited)
func (a *AccountingApi) Balance() (int64, error) {
	if a.metrics == nil {
		return 0, errNoAccountingMetrics
	}
	balance := mBalanceCredit.Count() - mBalanceDebit.Count()
	return balance, nil
}

// BalanceCredit returns total amount of units credited by local node
func (a *AccountingApi) BalanceCredit() (int64, error) {
	if a.metrics == nil {
		return 0, errNoAccountingMetrics
	}
	return mBalanceCredit.Count(), nil
}

// BalanceCredit returns total amount of units debited by local node
func (a *AccountingApi) BalanceDebit() (int64, error) {
	if a.metrics == nil {
		return 0, errNoAccountingMetrics
	}
	return mBalanceDebit.Count(), nil
}

// BytesCredit returns total amount of bytes credited by local node
func (a *AccountingApi) BytesCredit() (int64, error) {
	if a.metrics == nil {
		return 0, errNoAccountingMetrics
	}
	return mBytesCredit.Count(), nil
}

// BalanceCredit returns total amount of bytes debited by local node
func (a *AccountingApi) BytesDebit() (int64, error) {
	if a.metrics == nil {
		return 0, errNoAccountingMetrics
	}
	return mBytesDebit.Count(), nil
}

// MsgCredit returns total amount of messages credited by local node
func (a *AccountingApi) MsgCredit() (int64, error) {
	if a.metrics == nil {
		return 0, errNoAccountingMetrics
	}
	return mMsgCredit.Count(), nil
}

// MsgDebit returns total amount of messages debited by local node
func (a *AccountingApi) MsgDebit() (int64, error) {
	if a.metrics == nil {
		return 0, errNoAccountingMetrics
	}
	return mMsgDebit.Count(), nil
}

// PeerDrops returns number of times when local node had to drop remote peers
func (a *AccountingApi) PeerDrops() (int64, error) {
	if a.metrics == nil {
		return 0, errNoAccountingMetrics
	}
	return mPeerDrops.Count(), nil
}

// SelfDrops returns number of times when local node was overdrafted and dropped
func (a *AccountingApi) SelfDrops() (int64, error) {
	if a.metrics == nil {
		return 0, errNoAccountingMetrics
	}
	return mSelfDrops.Count(), nil
}
