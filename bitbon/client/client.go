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

package client

import (
	"fmt"
	"github.com/simcord-llc/bitbon-system-blockchain/signals"
	"runtime"
	"sync"
	"time"

	"github.com/golang/protobuf/proto" // nolint:staticcheck
	"github.com/pkg/errors"
	"github.com/simcord-llc/bitbon-system-blockchain/bitbon"
	bitbonAmqp "github.com/simcord-llc/bitbon-system-blockchain/bitbon/amqp"
	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/dto/external"
	bitbonErrors "github.com/simcord-llc/bitbon-system-blockchain/bitbon/errors"
	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/storage"
	"github.com/simcord-llc/bitbon-system-blockchain/log"
	"github.com/simcord-llc/bitbon-system-blockchain/node"
	"github.com/simcord-llc/bitbon-system-blockchain/p2p"
	"github.com/simcord-llc/bitbon-system-blockchain/quorum"
	"github.com/simcord-llc/bitbon-system-blockchain/rpc"
	"github.com/streadway/amqp"
	"github.com/vidmed/cony"
)

var (
	errInvalidClientMode = errors.New("invalid client mode")
)

type protoMessageWithID interface {
	proto.Message
	GetId() string
}

type handleFunc func([]byte) (protoMessageWithID, error)

type workerConfig struct {
	exchange       string
	queueName      string
	reqRoutingKey  string
	respRoutingKey string
	handle         handleFunc
	threads        int
}

type Client struct {
	bitbon            *bitbon.Bitbon
	quorum            *quorum.Quorum
	amqpConsumeClient *cony.Client
	amqpPublishClient *cony.Client
	cfg               *Config

	journal     *bitbonAmqp.PublishJournal
	journalPath string

	amqpWorkerPoolMux sync.RWMutex
	amqpWorkerPool    map[string]*Pool

	amqpWorkerPoolErrorsMux sync.RWMutex
	amqpWorkerPoolErrors    map[string]map[int]*cony.ErrorsBatch

	done chan struct{}
}

func New(ctx *node.ServiceContext, cfg *Config) (*Client, error) {
	var bitbonServ *bitbon.Bitbon
	if err := ctx.Service(&bitbonServ); err != nil {
		return nil, errors.Wrap(err, "error getting bitbon service")
	}

	var quorumServ *quorum.Quorum
	if err := ctx.Service(&quorumServ); err != nil {
		log.Error("error getting quorum service", "err", err)
	}

	journalPath := ctx.ResolvePath(cfg.JournalName)
	if journalPath == "" {
		return nil, fmt.Errorf("client journal name is not specified")
	}

	return &Client{
		bitbon:               bitbonServ,
		quorum:               quorumServ,
		cfg:                  cfg,
		amqpWorkerPool:       make(map[string]*Pool),
		amqpWorkerPoolErrors: make(map[string]map[int]*cony.ErrorsBatch),
		journalPath:          journalPath,
		done:                 make(chan struct{}),
	}, nil
}

func (c *Client) APIs() []rpc.API {
	return []rpc.API{
		{
			Namespace: ProtocolName,
			Version:   ProtocolVersionStr,
			Service:   NewAPI(c),
			Public:    true,
		},
	}
}

func (c *Client) Start(_ *p2p.Server) error {
	// wait until eth core would start
	log.Debug("wait until eth core will start...")
	ethStartedCh := signals.GetEthStartedChannel()
	<-ethStartedCh
	time.Sleep(1 * time.Second)
	log.Debug("eth core started - starting up bitbon client...")

	if !c.cfg.WaitSync {
		log.Debug("starting bitbon client with sync ignore")
		return c.init()
	}

	go func() {
		<-c.done
		log.Error("bitbon client stopped before sync is finished")
	}()
	return nil
}

func (c *Client) init() (err error) {
	defer func() {
		if err != nil {
			err2 := c.Stop()
			log.Error("Bitbon client init failed. So it has been stopped", "error", err, "(stop error)", err2)
		}
	}()

	// ------------------- Publish client
	c.amqpPublishClient, err = bitbonAmqp.NewClient(
		c.cfg.AmqpConnStrs,
		log.New("component", "(bitbon client amqpPublishClient)"),
	)
	if err != nil {
		return errors.Wrap(err, "amqp publish client connect error")
	}

	// ------------------- Consume client
	c.amqpConsumeClient, err = bitbonAmqp.NewClient(
		c.cfg.AmqpConnStrs,
		log.New("component", "(bitbon client amqpConsumeClient)"),
	)
	if err != nil {
		return errors.Wrap(err, "amqp consume client connect error")
	}

	// ------------------- Journal
	c.journal, err = bitbonAmqp.NewPublishJournal(c.journalPath, c.cfg.JournalTimeout,
		c.cfg.AmqpConnStrs, log.New("component", "(bitbon client journal)"))
	if err != nil {
		return errors.Wrap(err, "error creating bitbon client PublishJournal")
	}

	// ------------------- Subscribing
	err = c.parallelSubscribe(c.cfg.Mode)
	if err != nil {
		return errors.Wrapf(err, "failed to subscribe %v", err)
	}

	return nil
}

func (c *Client) parallelSubscribe(mode BitbonClientMode) error {
	configMap := c.processConfigs(mode)
	if configMap == nil {
		return errInvalidClientMode
	}
	dataLen := len(configMap)
	cpuNumber := runtime.NumCPU()
	partSize := dataLen/(cpuNumber*3) + 1

	runtime.GOMAXPROCS(cpuNumber)
	wg := sync.WaitGroup{}

	dataKeySlice := make([]string, 0, len(configMap))
	for key := range configMap {
		dataKeySlice = append(dataKeySlice, key)
	}

	errCh := make(chan error, dataLen)
	defer close(errCh)

	for startIndex := 0; startIndex < dataLen; startIndex += partSize {
		endIndex := partSize + startIndex
		if endIndex > dataLen {
			endIndex = dataLen
		}

		wg.Add(1)
		go func(data map[string]workerConfig, keys []string, errorCh chan error) {
			defer wg.Done()

			for idx := range keys {
				tmpConfig := data[keys[idx]]
				err := c.subscribe(keys[idx], &tmpConfig)
				if err != nil {
					errorCh <- errors.Wrapf(err, "failed to subscribe %s", keys[idx])
					return
				}
			}
		}(configMap, dataKeySlice[startIndex:endIndex], errCh)
	}

	wg.Wait()

	if len(errCh) > 0 {
		return <-errCh
	}

	return nil
}

func (c *Client) Stop() error {
	log.Info("Bitbon Client service is waiting for all tasks are finished...")

	// at first notify all client workers to terminate
	if c.Stopped() {
		log.Warn("Bitbon Client already stopped")
	} else {
		close(c.done)
	}

	// ------------------- Stop workers
	log.Debug("Bitbon Client is waiting for workers...")
	c.amqpWorkerPoolMux.Lock()
	for key, amqpWorkerPool := range c.amqpWorkerPool {
		log.Debug("Bitbon Client stopping workers...", "key", key)
		amqpWorkerPool.Close()
		log.Debug("Bitbon Client workers stopped", "key", key)
	}
	c.amqpWorkerPoolMux.Unlock()
	log.Debug("Bitbon Client workers done")

	// ------------------- Stop journal
	if c.journal != nil {
		c.journal.Stop()
		log.Debug("Bitbon Client journal stopped")
	}

	// ------------------- Stop amqpConsumeClient
	if c.amqpConsumeClient != nil {
		c.amqpConsumeClient.Close()
		log.Debug("Bitbon Client amqpConsumeClient closed")
	}

	// ------------------- Stop amqpPublishClient
	if c.amqpPublishClient != nil {
		c.amqpPublishClient.Close()
		log.Debug("Bitbon Client amqpPublishClient closed")
	}

	log.Info("Bitbon Client service stopped")
	return nil
}

func (c *Client) Stopped() bool {
	select {
	case <-c.done:
		return true
	default:
		return false
	}
}

func (c *Client) handleError(err error) *external.Error {
	if err != nil {
		var (
			bErr bitbonErrors.Error
			ok   bool
		)
		bErr, ok = err.(bitbonErrors.Error)
		if !ok {
			errCause := errors.Cause(err)
			bErr, ok = errCause.(bitbonErrors.Error)
			if !ok {
				switch errCause {
				case storage.ErrNotFound:
					bErr = bitbonErrors.NewNotFoundError(err)
				default:
					log.Debug("unable assert error to bitbon Error")
					bErr = bitbonErrors.NewInternalError(err)
				}
			}
		}

		errorCode := int64(bErr.ErrorCode())
		return &external.Error{
			Code:      errorCode,
			Message:   bErr.Error(),
			DataError: nil,
		}
	}

	return nil
}

func (c *Client) setPoolSize(key string, n int) error {
	if n <= 0 {
		return errors.New("unpositive pool size")
	}
	c.amqpWorkerPoolMux.Lock()
	defer c.amqpWorkerPoolMux.Unlock()
	amqpWorkerPool, ok := c.amqpWorkerPool[key]
	if !ok {
		return errors.Errorf("worker with key %q is not registered", key)
	}
	// change pool size
	amqpWorkerPool.SetSize(n)

	return nil
}

func (c *Client) getPoolSize(key string) (int, error) {
	c.amqpWorkerPoolMux.RLock()
	defer c.amqpWorkerPoolMux.RUnlock()
	amqpWorkerPool, ok := c.amqpWorkerPool[key]
	if !ok {
		return 0, errors.Errorf("worker with key %q is not registered", key)
	}

	return amqpWorkerPool.GetSize(), nil
}

func (c *Client) addAmqpWorkerPoolErrors(key string, index int, eb *cony.ErrorsBatch) {
	c.amqpWorkerPoolErrorsMux.Lock()
	defer c.amqpWorkerPoolErrorsMux.Unlock()
	if _, ok := c.amqpWorkerPoolErrors[key]; !ok {
		c.amqpWorkerPoolErrors[key] = make(map[int]*cony.ErrorsBatch)
	}
	c.amqpWorkerPoolErrors[key][index] = eb
}

func (c *Client) deleteAmqpWorkerPoolErrors(key string, index int) {
	c.amqpWorkerPoolErrorsMux.Lock()
	defer c.amqpWorkerPoolErrorsMux.Unlock()

	delete(c.amqpWorkerPoolErrors[key], index)
}

func (c *Client) getAmqpWorkerPoolErrors(key string) (map[int]*cony.ErrorsBatch, error) {
	c.amqpWorkerPoolErrorsMux.RLock()
	defer c.amqpWorkerPoolErrorsMux.RUnlock()
	errs, ok := c.amqpWorkerPoolErrors[key]
	if !ok {
		return nil, errors.Errorf("worker with key %q is not registered", key)
	}
	return errs, nil
}

// subscribe
func (c *Client) subscribe(key string, cfg *workerConfig) error {
	// Declarations
	exc := &cony.Exchange{
		Name:       cfg.exchange,
		Kind:       "topic",
		Durable:    true,
		AutoDelete: false,
		Internal:   false,
		NoWait:     false,
		Args:       nil,
	}
	que := &cony.Queue{
		Name:       cfg.queueName,
		Durable:    true,
		AutoDelete: false,
		Exclusive:  false,
		NoWait:     false,
		Args:       nil,
	}
	bnd := &cony.Binding{
		Queue:    que,
		Exchange: exc,
		Key:      cfg.reqRoutingKey,
		NoWait:   false,
		Args:     nil,
	}

	err := c.amqpConsumeClient.Declare([]cony.Declaration{
		cony.DeclareQueue(que),
		cony.DeclareExchange(exc),
		cony.DeclareBinding(bnd),
	})
	if err != nil {
		return errors.Wrapf(err, "error while amqpConsumeClient declaration: %q, exchange: %q, routing key: %q",
			que.Name, cfg.exchange, cfg.reqRoutingKey)
	}

	logger := log.New("exchange", cfg.exchange, "reqRoutingKey", cfg.reqRoutingKey,
		"respRoutingKey", cfg.respRoutingKey, "queue", cfg.queueName, "threads", cfg.threads)

	// Create worker pool
	pool := c.getPool(que, cfg, key, logger)

	logger.Debug("Bitbon Client starting worker pool...")
	pool.Process(cfg.threads)

	c.amqpWorkerPoolMux.Lock()
	defer c.amqpWorkerPoolMux.Unlock()
	c.amqpWorkerPool[key] = pool

	return nil
}

func (c *Client) getPool(que *cony.Queue, cfg *workerConfig, key string, logger log.Logger) *Pool {
	return NewPool(func(num int, closeChan <-chan struct{}) {
		logger2 := logger.New("workerNum", num)

		consumer := bitbonAmqp.NewConsumer(c.amqpConsumeClient, que, logger2)
		// stop consumer when closeChan closed
		go func() {
			select { // nolint:gosimple
			case <-closeChan:
				c.amqpConsumeClient.CancelConsumer(consumer)
			}
		}()

		publisher := bitbonAmqp.NewPublisher(c.amqpPublishClient, cfg.exchange,
			cfg.respRoutingKey, c.cfg.AmqpConfirmTimeout, logger2)
		// cancel publisher when processor is done
		defer c.amqpPublishClient.CancelPublisher(publisher)

		// publishing errors
		errs := cony.NewErrorBatch()
		c.addAmqpWorkerPoolErrors(key, num, errs)
		defer c.deleteAmqpWorkerPoolErrors(key, num)

		// this func will return when consumer will be closed
		c.process(consumer, publisher, cfg.handle, errs, logger2)
	})
}

func (c *Client) process(consumer *cony.Consumer, publisher *cony.Publisher,
	handle handleFunc, errs *cony.ErrorsBatch, logger log.Logger) {
	logger.Info("Bitbon Client Rabbit worker started")
	defer logger.Info("Bitbon Client Rabbit worker stopped")

	// when consumer will be stopped - consumer.Deliveries() channel will be closed
	for { // nolint:gosimple
		select {
		case delivery, ok := <-consumer.Deliveries():
			if !ok {
				logger.Warn("Bitbon Client Rabbit worker deliveries channel closed")
				return
			}
			logger.Warn("Bitbon Client Rabbit worker new delivery")

			// ack message as soon as possible. If there is error while acking delivery - do not process delivery.
			// Because rabbit broker do not know about our ack and will reprocess it via another node
			ackErr := delivery.Ack(false)
			if ackErr != nil {
				logger.Error("Bitbon Client Rabbit worker failed to ack delivery", "error", ackErr)
				continue
			}
			logger.Debug("Bitbon Client Rabbit worker message successfully acked")

			if delivery.Body == nil {
				logger.Error("Bitbon Client Rabbit worker Received client message with empty body.")
				continue
			}

			protoMessage, err := handle(delivery.Body)
			if protoMessage == nil {
				logger.Error("Bitbon Client Rabbit worker Handler returned empty response proto.")
				continue
			}
			if err != nil {
				logger.Error("Bitbon Client Rabbit worker Handler returned error.",
					"requestId", protoMessage.GetId(), "error", err)
			}

			// marshal proto response
			responseBytes, err := proto.Marshal(protoMessage)
			if err != nil {
				logger.Error("Bitbon Client Rabbit worker error marshaling handler response.",
					"requestId", protoMessage.GetId(), "error", err)
				continue
			}
			logger.Debug("Bitbon Client Rabbit worker handler response successfully marshaled in proto",
				"requestId", protoMessage.GetId())

			c.publish(publisher, responseBytes, protoMessage, errs, logger)
		}
	}
}

func (c *Client) publish(publisher *cony.Publisher, responseBytes []byte,
	protoMessage protoMessageWithID, errs *cony.ErrorsBatch, logger log.Logger) {
	now := time.Now()
	err := publisher.Publish(&amqp.Publishing{
		Body: responseBytes,
	}, c.cfg.AmqpPublishTimeout)
	if err != nil {
		logger.Error("Bitbon Client Rabbit worker response publishing error",
			"requestId", protoMessage.GetId(), "(publish took)", time.Since(now), "error", err)

		errs.Add(err)
		logger.Debug("Bitbon Client Rabbit worker response publishing error added to errors batch",
			"requestId", protoMessage.GetId())

		// add publishing to journal
		if err := c.journal.Insert(&bitbonAmqp.Publishing{
			Body:       responseBytes,
			Exchange:   publisher.Exchange(),
			RoutingKey: publisher.RoutingKey(),
			RequestID:  protoMessage.GetId(),
		}); err != nil {
			logger.Error("Bitbon Client Rabbit worker unable to add publishing to journal",
				"requestId", protoMessage.GetId(), "error", err)
		} else {
			logger.Debug("Bitbon Client Rabbit worker successfully added publishing to journal",
				"requestId", protoMessage.GetId(), "error", err)
		}
	} else {
		// if there is some errors in errs and last error happened more than LastErrorCleanInterval ago - reset errors
		if errs.Len() > 0 && (time.Now().UnixNano()-errs.Last().UnixNano()) > int64(LastErrorCleanInterval) {
			errs.Reset()
			logger.Debug("Bitbon Client Rabbit worker errors batch reset", "requestId", protoMessage.GetId())
		}

		logger.Debug("Bitbon Client Rabbit worker response publishing success",
			"(publish took)", time.Since(now), "requestId", protoMessage.GetId())
	}
}

// Protocols is a meaningless implementation of node.Service.
func (c *Client) Protocols() []p2p.Protocol {
	return nil
}

func (c *Client) GetJournal() *bitbonAmqp.PublishJournal {
	return c.journal
}
