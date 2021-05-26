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
		log.Debug("Redis Params: ", "host", u.Host, "pass", password)
	}

	clusterClient := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:       addresses,
		DialTimeout: time.Duration(cfg.DialTimeout) * time.Millisecond,
		MaxRetries:  int(cfg.MaxRetries),
		Password:    password,
	})

	log.Info("Bitbon noncer started")

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
		return 0, errors.Wrap(err, "unable get nonce string from redis")
	}

	log.Debug("raw redis GetNonce result", "result", unMarshallNonce)
	nonce, err := strconv.ParseInt(unMarshallNonce, 10, 64)
	if err != nil {
		return 0, errors.Wrap(err, "unable to parse nonce string from redis")
	}

	return nonce, nil
}

//  increments nonce and returns the incremented value
func (n *Noncer) IncrementAndGetNonce(assetbox common.Address) (int64, error) {
	nonce, err := n.clusterClient.HIncrBy(n.cfg.Key, assetbox.Hex(), 1).Result()
	log.Debug("IncrementAndGetNonce", "(nonce for use)", nonce)
	if err != nil {
		return 0, errors.Wrap(err, "unable get nonce after INCRBY from redis")
	}

	return nonce, err
}

func (n *Noncer) SetUpNonce(assetbox common.Address, nonce int64) error {
	log.Trace("Try to set nonce", "key", n.cfg.Key, "assetbox", assetbox.Hex(), "(nonce-1)", nonce-1)
	// we set in redis (nonce-1) because after that we call IncrementAndGetNonce
	// which increments nonce and returns the incremented value
	// so in redis we store the current amount of transactions (blockchain nonce -1 )
	result, err := n.clusterClient.HSetNX(n.cfg.Key, assetbox.Hex(), nonce-1).Result()
	if err != nil {
		return errors.Wrap(err, "error setting nonce to redis")
	}

	log.Debug("Setting nonce in Redis", "result(was set)", result)
	return nil
}

// directly sets nonce for given address
// use very carefully
func (n *Noncer) ForceNonce(assetbox common.Address, nonce int64) error {
	log.Trace("try to force nonce", "key", n.cfg.Key, "assetbox", assetbox.Hex(), "(nonce-1)", nonce-1)
	result, err := n.clusterClient.HSet(n.cfg.Key, assetbox.Hex(), nonce-1).Result()
	if err != nil {
		return errors.Wrap(err, "ERROR FORCING NONCE IN REDIS")
	}

	log.Debug("forcing nonce in redis", "result(true if set, false if updated)", result)
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
func (n *Noncer) Start(_ *p2p.Server) (err error) {
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
