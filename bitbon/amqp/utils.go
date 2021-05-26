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

// nolint:dupl
package amqp

import (
	"time"

	"github.com/simcord-llc/bitbon-system-blockchain/log"
	"github.com/vidmed/cony"
)

func NewClient(connection []string, logger log.Logger) (*cony.Client, error) {
	client := cony.NewClient(
		cony.URLs(connection),
		cony.Backoff(cony.DefaultBackoff),
		cony.Log(logger.(cony.Logger)),
	)
	err := client.Connect()
	if err != nil {
		return nil, err
	}

	// listen amqpPublishClient error channel and blocking channel
	go func() {
		defer logger.Debug("client error listener terminated")

		for client.Loop() {
			// try to exit ASAP
			select {
			case <-client.CloseCh():
				logger.Debug("client closed")
				return
			default:
			}

			select {
			case <-client.CloseCh():
				logger.Debug("client closed")
				return
			case err = <-client.Errors():
				if err == cony.ErrStopped {
					logger.Debug("client closed")
					return
				}
				logger.Error("client error", "err", err)
			case blocked := <-client.Blocking():
				logger.Warn("client is blocked", "reason", blocked.Reason)
			}
		}
	}()

	return client, err
}

func NewConsumer(client *cony.Client, queue *cony.Queue, logger log.Logger) *cony.Consumer {
	// Declare and register a consumer
	consumer := cony.NewConsumer(
		queue,
		cony.Qos(1),
		cony.ConsumerLog(logger.New("component", "consumer").(cony.Logger)),
	)
	client.Consume(consumer)
	// listen consumer error channel
	go func() {
		logger.Debug("Consumer error listener started")
		defer logger.Debug("Consumer error listener terminated")

		errs := cony.NewErrorBatch()
		for !client.Closed() {
			select {
			case err := <-consumer.Errors():
				if err == cony.ErrConsumerStopped {
					logger.Debug("Consumer error stopped.")
					return
				}
				if err == cony.ErrConsumerServingStopped {
					logger.Debug("Consumer error serving stopped received.", "err", err, "errSnaphot", errs)
					errs.Reset()
					continue
				}

				if client.Connected() {
					// there are some recent errors - try to disconnect client and reset errors
					if errs.Len() > 1 {
						logger.Error("Consumer error. Rabbit client connected. Recent errors > 1. Disconnect client...",
							"err", err, "errSnaphot", errs)
						if discErr := client.Disconnect(); discErr != nil {
							logger.Error("Consumer error. Rabbit client connected. Recent errors > 1. Disconnect client error",
								"discErr", discErr, "errSnaphot", errs, "err", err)
						}
						errs.Reset()
					} else { // there are NO recent errors - redeclare Consumer (amqp channel recreate)
						logger.Error("Consumer error. Rabbit client connected. Recent errors <= 1. Recreate Consumer...",
							"errSnaphot", errs, "err", err)
						errs.Add(err)
						client.Consume(consumer)
					}
				} else {
					logger.Error("Consumer error. Rabbit client NOT connected. Wait for connection...",
						"err", err, "errSnaphot", errs)
					// if there is no connection  - new connection will be established - so reset errors
					errs.Reset()
				}
			case <-consumer.StopCh():
				logger.Debug("Consumer error stopped.")
				return
			}
		}
	}()

	return consumer
}

func NewPublisher(client *cony.Client, exchange, routingKey string,
	confirmTimeout time.Duration, logger log.Logger) *cony.Publisher {
	// Declare and register a publisher
	publisher := cony.NewPublisher(
		exchange,
		routingKey,
		cony.PublisherLog(logger.New("component", "publisher").(cony.Logger)),
		cony.PublisherConfirmTimeout(confirmTimeout),
	)
	client.Publish(publisher)
	// listen publisher error channel
	go func() {
		logger.Debug("Publisher error listener started")
		defer logger.Debug("Publisher error listener terminated")

		errs := cony.NewErrorBatch()
		for !client.Closed() {
			select {
			case err := <-publisher.Errors():
				if err == cony.ErrPublisherStopped {
					logger.Debug("Publisher error stopped.")
					return
				}
				if err == cony.ErrPublisherServingStopped {
					logger.Debug("Publisher error serving stopped received.", "err", err, "errSnaphot", errs)
					errs.Reset()
					continue
				}

				if client.Connected() {
					// there are some recent errors - try to disconnect client and reset errors
					if errs.Len() > 1 {
						logger.Error("Publisher error. Rabbit client connected. Recent errors > 1. Disconnect client...",
							"err", err, "errSnaphot", errs)
						if discErr := client.Disconnect(); discErr != nil {
							logger.Error("Publisher error. Rabbit client connected. Recent errors > 1. Disconnect client error",
								"discErr", discErr, "errSnaphot", errs, "err", err)
						}
						errs.Reset()
					} else { // there are NO recent errors - redeclare publisher (amqp channel recreate)
						logger.Error("Publisher error. Rabbit client connected. Recent errors <= 1. Recreate publisher...",
							"errSnaphot", errs, "err", err)
						errs.Add(err)
						client.Publish(publisher)
					}
				} else {
					logger.Error("Publisher error. Rabbit client NOT connected. Wait for connection...",
						"err", err, "errSnaphot", errs)
					// if there is no connection  - new connection will be established - so reset errors
					errs.Reset()
				}
			case <-publisher.StopCh():
				logger.Warn("Publisher error stopped.")
				return
			}
		}
	}()

	return publisher
}
