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

package agent

import (
	"context"
	"errors"

	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/amqp"
)

var (
	errJournalNotDefined = errors.New("journal is not defined")
)

// API provides the bitbon bitbon RPC service
type API struct {
	agent *Agent
}

// NewAPI create agent new RPC bitbon agent service.
func NewAPI(a *Agent) *API {
	return &API{agent: a}
}

// Version returns version.
func (api *API) Version(_ context.Context) string {
	return ProtocolVersionStr
}

func (api *API) GetJournal() ([]*amqp.Publishing, error) {
	j := api.agent.GetJournal()
	if j == nil {
		return nil, errJournalNotDefined
	}
	return j.GetAll(), nil
}
