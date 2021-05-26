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

package quorum

import (
	"context"

	"github.com/simcord-llc/bitbon-system-blockchain/common"
	"github.com/simcord-llc/bitbon-system-blockchain/quorum/snapshot"
)

// API provides the quorum RPC service
type API struct {
	q *Quorum
}

// NewAPI create a new RPC quorum service.
func NewAPI(q *Quorum) *API {
	return &API{q: q}
}

// Version returns the quorum sub-protocol version.
func (api *API) Version(_ context.Context) string {
	return ProtocolVersionStr
}

// AllNodes returns all nodes known by quorum protocol (static & trusted nodes)
func (api *API) AllNodes(_ context.Context) map[string]string {
	m := make(map[string]string, len(api.q.allNodes))
	for _, peer := range api.q.allNodes {
		m[peer.ID().String()] = peer.String()
	}
	return m
}

// PingList returns peer list of quorum protocol (with satisfying ping duration)
func (api *API) PingList(_ context.Context) map[string]string {
	m := make(map[string]string, api.q.peerList.PingListLen())
	for _, peer := range api.q.peerList.PingListPeers() {
		m[peer.ID().String()] = peer.Node().String()
	}
	return m
}

// MinerListAtRecentEpoch returns miner list at recent epoch
func (api *API) MinerListAtRecentEpoch(_ context.Context) []common.Address {
	return api.q.MinerListAtRecentEpoch()
}

// PeerList returns list of all connected peers
func (api *API) PeerList(_ context.Context) *PeerListDTO {
	return api.q.peerList.ToDTO()
}

func (api *API) Snapshot(_ context.Context, epoch uint64) *snapshot.SnapDTO {
	s := api.q.snapshotSystem.GetSnapshot(epoch)
	if s == nil {
		return nil
	}
	return s.ToDTO()
}

func (api *API) SnapshotSystem(_ context.Context) *snapshot.SystemDTO {
	return api.q.snapshotSystem.ToDTO()
}

func (api *API) SnapshotSystemEpochs(_ context.Context) []uint64 {
	return api.q.snapshotSystem.Epochs()
}

func (api *API) SnapshotSystemRecentEpoch(_ context.Context) uint64 {
	return api.q.snapshotSystem.RecentEpoch()
}
