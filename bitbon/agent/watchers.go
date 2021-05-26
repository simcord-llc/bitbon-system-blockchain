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
	"fmt"
	"strings"
	"time"

	"github.com/golang/protobuf/proto"
	bitbonAmqp "github.com/simcord-llc/bitbon-system-blockchain/bitbon/amqp"
	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/contracts"
	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/dto"
	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/dto/external"
	"github.com/simcord-llc/bitbon-system-blockchain/core"
	"github.com/simcord-llc/bitbon-system-blockchain/log"
	"github.com/streadway/amqp"
	"github.com/vidmed/cony"
)

func (a *Agent) publishProto(protoMessage proto.Message, publisher *cony.Publisher, logger log.Logger) {
	body, err := proto.Marshal(protoMessage)
	if err != nil {
		logger.Error("Bitbon Agent proto marshalling message error", "error", err)
		return
	}
	logger.Debug("Bitbon Agent message successfully marshaled in proto")

	now := time.Now()
	err = publisher.Publish(&amqp.Publishing{
		Body: body,
	}, DefaultPublishTimeout)
	if err != nil {
		logger.Error("Bitbon Agent Rabbit error publishing message", "(publish took)", time.Since(now), "error", err)

		// add publishing to journal
		if err := a.journal.Insert(&bitbonAmqp.Publishing{
			Body:       body,
			Exchange:   publisher.Exchange(),
			RoutingKey: publisher.RoutingKey(),
		}); err != nil {
			logger.Error("Bitbon Agent unable to add publishing to journal", "error", err)
		} else {
			logger.Debug("Bitbon Agent successfully added publishing to journal", "error", err)
		}

	} else {
		logger.Debug("Bitbon Agent message successfully published to rabbit", "(publish took)", time.Since(now))
	}
}

// watchTransaction method is blocking.
func (a *Agent) watchTransaction(routingKey string) {
	logger := log.New("exchange", a.cfg.AmqpExchange, "routingKey", routingKey, "component", "(bitbon agent watchTransaction)")
	ctx := context.Background()
	chainEventChannel := make(chan core.ChainEvent, 1)

	sub := a.bitbon.GetContractsManager().GetEthAPIWrapper().SubscribeChainEvent(ctx, chainEventChannel)
	defer func() {
		sub.Unsubscribe()
		log.Warn("watchTransaction unsubscribed.")
	}()

	publisher := bitbonAmqp.NewPublisher(a.amqpClient, a.cfg.AmqpExchange, routingKey, DefaultConfirmTimeout, logger)
	// cancel publisher when processor is done
	defer a.amqpClient.CancelPublisher(publisher)

	log.Debug("watchTransaction successfully started")

	for {
		select {
		case ch, ok := <-chainEventChannel:
			if !ok {
				logger.Warn("watchTransaction subscription channel closed")
				return
			}

			if ch.Block.Transactions().Len() == 0 {
				logger.Trace("watchTransaction subscription got block without transactions. Continue...")
				continue
			}

			var transactions []*external.TransactionObject
			var transactionHashes []string

			for _, tx := range ch.Block.Transactions() {
				var transactionObj = &external.TransactionObject{
					Hash:        strings.ToLower(tx.Hash().Hex()),
					Nonce:       tx.Nonce(),
					BlockHash:   strings.ToLower(ch.Block.Hash().String()),
					BlockNumber: ch.Block.NumberU64(),
					BlockTime:   ch.Block.Time(),
					GasPrice:    tx.GasPrice().Uint64(),
					GasLimit:    tx.Gas(),
					Time:        tx.Time(),
					Size:        tx.Size().String(),
				}

				bitbonTx := a.bitbon.GetParser().Parse(ctx, tx)

				if bitbonTx != nil {
					var method = dto.GetMethod(bitbonTx.Method)
					transactionObj.BitbonTx = &external.BitbonTransaction{
						Status:      bitbonTx.Status,
						ParseStatus: int64(bitbonTx.ParseStatus),
						Method:      method,
						Sender:      strings.ToLower(bitbonTx.Sender.Hex()),
					}

					if bitbonTx.ParseError != nil {
						transactionObj.BitbonTx.ParseError = bitbonTx.ParseError.Error()
					}

					if bitbonTx.Args != nil {
						transactionObj.BitbonTx.Args = bitbonTx.Args.ToExternalDTO()
					}

					if bitbonTx.TxError != nil {
						transactionObj.BitbonTx.TxError = &external.Error{
							Code:      int64(bitbonTx.TxError.ErrorCode()),
							Message:   bitbonTx.TxError.Error(),
							DataError: nil,
						}
					}

					if bitbonTx.Balances != nil {
						transactionObj.BitbonTx.Balances = make([]*external.Balance, 0, len(bitbonTx.Balances))
						for address, value := range bitbonTx.Balances {
							balance := &external.Balance{
								Address: strings.ToLower(address.Hex()),
								Balance: value.String(),
							}
							transactionObj.BitbonTx.Balances = append(transactionObj.BitbonTx.Balances, balance)
						}
					}
				}

				transactions = append(transactions, transactionObj)
				transactionHashes = append(transactionHashes, transactionObj.GetHash())
			}

			if len(transactions) > 0 {
				txNotification := &external.TransactionNotification{
					Transactions: transactions,
				}
				a.publishProto(txNotification, publisher, logger.New("txLen", len(transactions), "txHashes", strings.Join(transactionHashes, ", ")))
			}

		case err, ok := <-sub.Err():
			if !ok {
				log.Warn("watchTransaction subscription ends")
			} else {
				log.Error("watchTransaction subscription error", "err", err)
			}
			return

		case <-a.done:
			log.Warn("watchTransaction loop terminated.")
			return
		}
	}
}

func (a *Agent) watchBlocks(routingKey string) {
	logger := log.New("exchange", a.cfg.AmqpExchange, "routingKey", routingKey, "component", "(bitbon agent watchBlocks)")
	ctx := context.Background()
	chainEventChannel := make(chan core.ChainEvent, 1)

	sub := a.bitbon.GetContractsManager().GetEthAPIWrapper().SubscribeChainEvent(ctx, chainEventChannel)
	defer func() {
		sub.Unsubscribe()
		log.Warn("watchBlocks unsubscribed.")
	}()

	publisher := bitbonAmqp.NewPublisher(a.amqpClient, a.cfg.AmqpExchange, routingKey, DefaultConfirmTimeout, logger)
	// cancel publisher when processor is done
	defer a.amqpClient.CancelPublisher(publisher)

	log.Debug("watchBlocks successfully started")

	for {
		select {
		case ch, ok := <-chainEventChannel:
			if !ok {
				logger.Warn("watchBlocks subscription channel closed")
				return
			}

			blockNotification := &external.BlockNotification{
				BlockInfo: &external.BlockMinedInfo{
					Number:            ch.Block.NumberU64(),
					Timestamp:         ch.Block.Time(),
					TransactionNumber: uint64(ch.Block.Transactions().Len()),
					Hash:              ch.Block.Hash().Hex(),
					ParentHash:        ch.Block.ParentHash().Hex(),
					Size:              ch.Block.Size().String(),
				},
			}

			logger2 := logger.New("block number", blockNotification.BlockInfo.Number, "block hash", blockNotification.BlockInfo.Hash, "tx number", blockNotification.BlockInfo.TransactionNumber)
			a.publishProto(blockNotification, publisher, logger2)

		case err, ok := <-sub.Err():
			if !ok {
				log.Warn("watchBlocks subscription ends")
			} else {
				log.Error("watchBlocks subscription error", "err", err)
			}
			return

		case <-a.done:
			log.Warn("watchBlocks loop terminated.")
			return
		}
	}
}

// watchAssetboxInfo method is blocking.
func (a *Agent) watchAssetboxInfo(routingKey string) {
	logger := log.New("exchange", a.cfg.AmqpExchange, "routingKey", routingKey, "component", "(bitbon agent watchAssetboxInfo)")

	if !a.bitbon.GetContractsManager().WaitReady(a.done) {
		logger.Warn("watchAssetboxInfo terminated as bitbon contracts manager is not ready")
		return
	}
	retryTimeout := 10 * time.Second
	sink := make(chan *contracts.Assetbox, 128)

	for err := a.bitbon.GetContractsManager().WatchAssetboxInfoSet(sink); err != nil; {
		logger.Error(fmt.Sprintf("error beginning watching assetbox info. Retry in %s", retryTimeout))
		time.Sleep(retryTimeout)
	}

	log.Info("watchAssetboxInfo successfully started")

	publisher := bitbonAmqp.NewPublisher(a.amqpClient, a.cfg.AmqpExchange, routingKey, DefaultConfirmTimeout, logger)
	// cancel publisher when processor is done
	defer a.amqpClient.CancelPublisher(publisher)

	for {
		select {
		case event, ok := <-sink:
			if !ok {
				logger.Warn("watchAssetboxInfo sink channel closed")
				return
			}

			send := &external.AssetboxInfoChanged{
				Address:   strings.ToLower(event.Address.Hex()),
				IsPublic:  event.IsPublic,
				Alias:     event.Alias,
				ServiceId: string(event.ServiceID),
				ExtraInfo: event.ExtraInfo,
				IsMining:  event.IsMining,
			}
			logger2 := logger.New("assetbox", event.Address.Hex())
			a.publishProto(send, publisher, logger2)

		case <-a.done:
			logger.Warn("watchAssetboxInfo loop terminated.")
			return
		}
	}
}

// watchExpiredTransfers method is blocking.
func (a *Agent) watchExpiredTransfers(routingKey string) {
	logger := log.New("exchange", a.cfg.AmqpExchange, "routingKey", routingKey, "component", "(bitbon agent watchAssetboxBalance)")

	if !a.bitbon.GetContractsManager().WaitReady(a.done) {
		logger.Warn("watchExpiredTransfers terminated as bitbon contracts manager is not ready")
		return
	}

	retryTimeout := 10 * time.Second
	sink := make(chan *contracts.TransferExpired, 128)

	for err := a.bitbon.GetContractsManager().WatchTransferExpired(sink); err != nil; {
		logger.Error(fmt.Sprintf("error beginning watching WatchTransferExpired. Retry in %s", retryTimeout))
		time.Sleep(retryTimeout)
	}

	log.Info("watchExpiredTransfers successfully started")

	publisher := bitbonAmqp.NewPublisher(a.amqpClient, a.cfg.AmqpExchange, routingKey, DefaultConfirmTimeout, logger)
	// cancel publisher when processor is done
	defer a.amqpClient.CancelPublisher(publisher)

	for {
		select {
		case event, ok := <-sink:
			if !ok {
				logger.Warn("watchExpiredTransfers sink channel closed")
				return
			}

			send := &external.TransferExpiredEvent{
				TransferId: string(event.TransferID),
			}
			logger2 := logger.New("Expired TransferId", string(event.TransferID))
			a.publishProto(send, publisher, logger2)

		case <-a.done:
			log.Warn("watchExpiredTransfers loop terminated.")
			return
		}
	}
}

// watchAssetboxBalance method is blocking.
func (a *Agent) watchAssetboxBalance(routingKey string) {
	logger := log.New("exchange", a.cfg.AmqpExchange, "routingKey", routingKey, "component", "(bitbon agent watchAssetboxBalance)")

	if !a.bitbon.GetContractsManager().WaitReady(a.done) {
		logger.Warn("watchAssetboxBalance terminated as bitbon contracts manager is not ready")
		return
	}
	retryTimeout := 10 * time.Second
	sink := make(chan *contracts.Balance, 128)

	for err := a.bitbon.GetContractsManager().WatchBitbonBalanceChanged(sink); err != nil; {
		logger.Error(fmt.Sprintf("error beginning watching BitbonBalanceChanged. Retry in %s", retryTimeout))
		time.Sleep(retryTimeout)
	}

	for err := a.bitbon.GetContractsManager().WatchBitbonBalanceLocked(sink); err != nil; {
		logger.Error(fmt.Sprintf("error beginning watching WatchBitbonBalanceLocked. Retry in %s", retryTimeout))
		time.Sleep(retryTimeout)
	}

	for err := a.bitbon.GetContractsManager().WatchBitbonBalanceUnLocked(sink); err != nil; {
		logger.Error(fmt.Sprintf("error beginning watching WatchBitbonBalanceLocked. Retry in %s", retryTimeout))
		time.Sleep(retryTimeout)
	}

	log.Info("watchAssetboxBalance successfully started")

	publisher := bitbonAmqp.NewPublisher(a.amqpClient, a.cfg.AmqpExchange, routingKey, DefaultConfirmTimeout, logger)
	// cancel publisher when processor is done
	defer a.amqpClient.CancelPublisher(publisher)

	for {
		select {
		case event, ok := <-sink:
			if !ok {
				logger.Warn("watchAssetboxBalance sink channel closed")
				return
			}

			send := &external.AssetboxBalanceChanged{
				Address: strings.ToLower(event.Address.Hex()),
				Balance: event.Balance.String(),
			}
			logger2 := logger.New("assetbox", event.Address.Hex(), "balance", event.Balance.String())
			a.publishProto(send, publisher, logger2)

		case <-a.done:
			log.Warn("watchAssetboxBalance loop terminated.")
			return
		}
	}
}
