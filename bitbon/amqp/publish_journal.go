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

package amqp

import (
	"sync"
	"time"

	"github.com/streadway/amqp"

	"github.com/simcord-llc/bitbon-system-blockchain/log"
	"github.com/vidmed/cony"

	"github.com/pborman/uuid"
	"github.com/pkg/errors"
	"github.com/simcord-llc/bitbon-system-blockchain/rlp"

	"github.com/syndtr/goleveldb/leveldb"
)

var (
	errEmptyPublishing = errors.New("empty publishing given")
	errStopped         = errors.New("stopped")
)

const (
	publishTimeout = 3 * time.Second
	confirmTimeout = 10 * time.Second
)

// Publishing represents an unpublished messages to RabbitMQ
type Publishing struct {
	Exchange   string
	RoutingKey string
	Body       []byte
	RequestID  string
}

type replayFunc func(*Publishing) error

// PublishJournal is log of unpublished messages to RabbitMQ
// with the aim of storing locally Publishings that cannot be published at the time
type PublishJournal struct {
	db            *leveldb.DB
	replayTimeout time.Duration
	client        *cony.Client
	logger        log.Logger

	wg   sync.WaitGroup
	done chan struct{}
}

func NewPublishJournal(path string, replayTimeout time.Duration,
	amqpConnStrs []string, logger log.Logger) (*PublishJournal, error) {
	db, err := leveldb.OpenFile(path, nil)
	if err != nil {
		logger.Error("error while opening leveldb file", "err", err)
		return nil, errors.Wrap(err, "publish journal unable to open journal file")
	}
	client, err := NewClient(
		amqpConnStrs,
		log.New("component", "(publish journal)"),
	)
	if err != nil {
		return nil, errors.Wrap(err, "publish journal amqp connect error")
	}

	j := &PublishJournal{
		db:            db,
		replayTimeout: replayTimeout,
		client:        client,
		logger:        logger,
		done:          make(chan struct{}),
	}

	go j.run()

	return j, nil
}

// Stop stops PublishJournal
// Stop waits until all its tasks finished
func (j *PublishJournal) Stop() {
	if j.Stopped() {
		j.logger.Error("PublishJournal already stopped")
		return
	}
	close(j.done)

	j.logger.Debug("PublishJournal waiting all tasks finished...")
	j.wg.Wait()
	j.logger.Debug("PublishJournal all tasks done...")
}

func (j *PublishJournal) Stopped() bool {
	select {
	case <-j.done:
		return true
	default:
		return false
	}
}

func (j *PublishJournal) run() {
	j.logger.Debug("PublishJournal worker started")
	defer j.logger.Debug("PublishJournal worker stopped")

	j.wg.Add(1)
	defer j.wg.Done()

	logger := j.logger.New("component", "PublishJournal publisher")

	publisher := NewPublisher(j.client, "", "", confirmTimeout, logger)
	defer j.client.CancelPublisher(publisher)

	replayTicker := time.NewTicker(j.replayTimeout)
	defer replayTicker.Stop()

	replay := func(p *Publishing) (err error) {
		publishing := &amqp.Publishing{
			Body: p.Body,
		}
		now := time.Now()
		defer func() {
			j.logger.Debug("publish called", "(publish took)", time.Since(now), "error", err)
		}()
		return publisher.PublishWithParams(publishing, p.Exchange, p.RoutingKey, publishTimeout)
	}

	if err := j.replay(replay); err != nil {
		j.logger.Error("Replay error", "error", err)
	}

	for {
		// to exit as soon as possible
		select {
		case <-j.done:
			return
		default:
		}

		select {
		case <-j.done:
			return
		case <-replayTicker.C:
			if err := j.replay(replay); err != nil {
				j.logger.Error("Replay error", "error", err)
			}
		}
	}
}

// Insert inserts Publishing to journal
func (j *PublishJournal) Insert(p *Publishing) (err error) {
	if p == nil {
		return errEmptyPublishing
	}

	defer func() {
		if err2 := recover(); err2 != nil {
			err = errors.Wrap(err, "unable to generate key")
			return
		}
	}()
	// this could panic
	key := uuid.NewRandom()

	value, err := rlp.EncodeToBytes(p)
	if err != nil {
		return errors.Wrap(err, "unable to rlp encode value")
	}

	err = j.db.Put(key, value, nil)
	if err != nil {
		return errors.Wrap(err, "unable to save publishing")
	}

	return nil
}

func (j *PublishJournal) GetAll() []*Publishing {
	iter := j.db.NewIterator(nil, nil)
	defer iter.Release()

	result := make([]*Publishing, 0)

	for iter.Next() {
		var (
			p       = &Publishing{}
			loopErr error
		)

		if loopErr = rlp.DecodeBytes(iter.Value(), p); loopErr != nil {
			j.logger.Error("PublishJournal unable to decode rlp bytes while replay. Continue..",
				"key", string(iter.Key()), "value", string(iter.Value()), "error", loopErr)
			continue
		}

		result = append(result, p)
	}

	return result
}

func (j *PublishJournal) replay(replayF replayFunc) (err error) {
	iter := j.db.NewIterator(nil, nil)
	successes, failed := 0, 0

	defer func() {
		iter.Release()
		j.logger.Debug("PublishJournal replay finished",
			"total", successes+failed, "successes", successes, "failed", failed, "error", err)
	}()

	for iter.Next() {
		if j.Stopped() {
			return errStopped
		}
		var (
			p       = &Publishing{}
			loopErr error
		)

		if loopErr = rlp.DecodeBytes(iter.Value(), p); loopErr != nil {
			failed++
			j.logger.Error("PublishJournal unable to decode rlp bytes while replay. Continue..",
				"key", string(iter.Key()), "value", string(iter.Value()), "error", loopErr)
			continue
		}

		if loopErr = replayF(p); loopErr != nil {
			failed++
			j.logger.Error("PublishJournal unable to replay publishing. Continue..",
				"key", string(iter.Key()), "value", string(iter.Value()), "error", loopErr)
			continue
		}

		if loopErr = j.db.Delete(iter.Key(), nil); loopErr != nil {
			failed++
			j.logger.Error("PublishJournal unable to delete key after success replay. Continue..",
				"key", string(iter.Key()), "value", string(iter.Value()), "error", loopErr)
			continue
		}
		successes++
	}

	return iter.Error()
}
