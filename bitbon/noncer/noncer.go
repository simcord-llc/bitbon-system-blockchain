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

package noncer

import (
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/go-redis/redis"
	"github.com/pkg/errors"
	"github.com/simcord-llc/bitbon-system-blockchain/common"
	"github.com/simcord-llc/bitbon-system-blockchain/log"
	"github.com/simcord-llc/bitbon-system-blockchain/node"
	"github.com/simcord-llc/bitbon-system-blockchain/p2p"
	"github.com/simcord-llc/bitbon-system-blockchain/rpc"
)

type Noncer struct {
	cfg           *Config
	clusterClient *redis.ClusterClient
	done          chan struct{}
}

func New(_ *node.ServiceContext, cfg *Config) (b *Noncer, err error) {
	var addresses = make([]string, 0, len(cfg.RedisConnStrs))
	var password string
	for _, connStr := range cfg.RedisConnStrs {
		u, err := url.Parse(connStr)
		if err != nil {
			return nil, errors.Wrap(err, "unable to parse noncer Redis connection string")
		}

		addresses = append(addresses, u.Host)
		password, _ = u.User.Password()
	}

	clusterClient := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:       addresses,
		DialTimeout: time.Duration(cfg.DialTimeout) * time.Millisecond,
		MaxRetries:  int(cfg.MaxRetries),
		Password:    password,
	})

	log.Warn(fmt.Sprintf("Bitbon noncer started, redis host %s, redis key %s", cfg.RedisConnStrs, cfg.Key))

	return &Noncer{
		cfg:           cfg,
		clusterClient: clusterClient,
		done:          make(chan struct{}),
	}, nil
}

// returns the current amount of transactions (for next transaction it should be incremented)
func (n *Noncer) GetNonce(assetbox common.Address) (int64, error) {
	unMarshallNonce, err := n.clusterClient.HGet(n.cfg.Key, assetbox.Hex()).Result()
	if err != nil {
		return 0, errors.Wrap(err, "noncer: unable get nonce string from redis")
	}

	log.Warn("noncer: raw redis GetNonce result", "result", unMarshallNonce)
	nonce, err := strconv.ParseInt(unMarshallNonce, 10, 64)
	if err != nil {
		return 0, errors.Wrap(err, "unable to parse nonce string from redis")
	}

	return nonce, nil
}

//  increments nonce and returns the incremented value
func (n *Noncer) IncrementAndGetNonce(assetbox common.Address) (int64, error) {
	nonce, err := n.clusterClient.HIncrBy(n.cfg.Key, assetbox.Hex(), 1).Result()
	log.Warn("noncer: IncrementAndGetNonce", "(nonce for use)", nonce, "assetbox", assetbox.Hex())
	if err != nil {
		return 0, errors.Wrap(err, "unable get nonce after INCRBY from redis")
	}

	return nonce, err
}

func (n *Noncer) SetUpNonce(assetbox common.Address, nonce int64) error {
	log.Warn("noncer: Try to set nonce", "key", n.cfg.Key, "assetbox", assetbox.Hex(), "(nonce-1)", nonce-1)
	// we set in redis (nonce-1) because after that we call IncrementAndGetNonce which increments nonce and returns the incremented value
	// so in redis we store the current amount of transactions (blockchain nonce -1 )
	result, err := n.clusterClient.HSetNX(n.cfg.Key, assetbox.Hex(), nonce-1).Result()
	if err != nil {
		return errors.Wrap(err, "error setting nonce to redis")
	}

	log.Warn("noncer: Setting nonce in Redis", "result(was set)", result)
	return nil
}

// directly sets nonce for given address
// use very carefully
func (n *Noncer) ForceNonce(assetbox common.Address, nonce int64) error {
	log.Warn("noncer: TRY TO FORCE NONCE", "KEY", n.cfg.Key, "ASSETBOX", assetbox.Hex(), "(NONCE-1)", nonce-1)
	result, err := n.clusterClient.HSet(n.cfg.Key, assetbox.Hex(), nonce-1).Result()
	if err != nil {
		return errors.Wrap(err, "ERROR FORCING NONCE IN REDIS")
	}

	log.Warn("noncer: FORCING NONCE IN REDIS", "RESULT(TRUE IF SET, FALSE IF UPDATED)", result)
	return nil
}

func (n *Noncer) CheckNoncerEligibility(assetbox common.Address) (bool, error) {
	result, err := n.clusterClient.HExists(n.cfg.Key, assetbox.Hex()).Result()
	log.Warn("noncer: noncer eligibility", "assetbox", assetbox.Hex(), "eligibility", result)

	if err != nil {
		return false, errors.Wrap(err, "error checking noncer eligibility")
	}

	return result, nil
}

func (n *Noncer) GetNoncerAssetboxes() ([]common.Address, error) {
	noncerAssetboxes, err := n.clusterClient.HKeys(n.cfg.Key).Result()
	if err != nil {
		return nil, errors.Wrap(err, "error getting assetboxes eligible for noncer from redis")
	}

	addresses := make([]common.Address, len(noncerAssetboxes))
	for idx := range noncerAssetboxes {
		addresses[idx] = common.HexToAddress(noncerAssetboxes[idx])
	}
	return addresses, nil
}

func (n *Noncer) RemoveAssetboxesFromNoncer(assetboxes []common.Address) error {
	assetboxesStr := make([]string, len(assetboxes))
	for idx := range assetboxes {
		assetboxesStr = append(assetboxesStr, assetboxes[idx].Hex())
	}

	_, err := n.clusterClient.HDel(n.cfg.Key, assetboxesStr...).Result()
	if err != nil {
		return errors.Wrap(err, "error deleting assetboxes for noncer redis")
	}
	return nil
}

// APIs is a meaningless implementation of node.Service.
func (n *Noncer) APIs() []rpc.API {
	return []rpc.API{
		{
			Namespace: ProtocolName,
			Version:   ProtocolVersionStr,
			Service:   NewAPI(n),
			Public:    true,
		},
	}
}

// Protocols is a meaningless implementation of node.Service.
func (n *Noncer) Protocols() []p2p.Protocol {
	return nil
}

// Start implements node.Service
func (n *Noncer) Start(srvr *p2p.Server) (err error) {
	return nil
}

// Stop implements node.Service
func (n *Noncer) Stop() error {
	if n.Stopped() {
		return errors.New("noncer already stopped")
	}

	close(n.done)

	err := n.clusterClient.Close()
	if err != nil {
		return errors.Wrap(err, "unable to close redis connection")
	}

	return err
}

func (n *Noncer) Stopped() bool {
	select {
	case <-n.done:
		return true
	default:
		return false
	}
}
