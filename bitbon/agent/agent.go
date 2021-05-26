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
	"fmt"

	"github.com/pkg/errors"
	_ "github.com/simcord-llc/bitbon-system-blockchain"
	bitbonPkg "github.com/simcord-llc/bitbon-system-blockchain/bitbon"
	bitbonAmqp "github.com/simcord-llc/bitbon-system-blockchain/bitbon/amqp"
	"github.com/simcord-llc/bitbon-system-blockchain/log"
	"github.com/simcord-llc/bitbon-system-blockchain/node"
	"github.com/simcord-llc/bitbon-system-blockchain/p2p"
	"github.com/simcord-llc/bitbon-system-blockchain/rpc"
	"github.com/vidmed/cony"
)

type Agent struct {
	bitbon     *bitbonPkg.Bitbon
	cfg        *Config
	amqpClient *cony.Client

	journal     *bitbonAmqp.PublishJournal
	journalPath string

	done chan struct{}
}

func New(ctx *node.ServiceContext, cfg *Config) (*Agent, error) {
	var bitbonService *bitbonPkg.Bitbon
	if err := ctx.Service(&bitbonService); err != nil {
		return nil, errors.Wrap(err, "error getting bitbon service")
	}

	journalPath := ctx.ResolvePath(cfg.JournalName)
	if journalPath == "" {
		return nil, fmt.Errorf("journal name is not specified")
	}

	return &Agent{
		bitbon:      bitbonService,
		cfg:         cfg,
		journalPath: journalPath,
		done:        make(chan struct{}),
	}, nil
}

func (a *Agent) Start(_ *p2p.Server) (err error) {
	a.amqpClient, err = bitbonAmqp.NewClient(
		a.cfg.AmqpConnStrs,
		log.New("component", "bitbon agent amqpClient"),
	)
	if err != nil {
		return errors.Wrap(err, "error creating amqp client ina bitbon agent")
	}

	// Declarations
	exc := &cony.Exchange{
		Name:       a.cfg.AmqpExchange,
		Kind:       "topic",
		Durable:    true,
		AutoDelete: false,
		Internal:   false,
		NoWait:     false,
		Args:       nil,
	}
	err = a.amqpClient.Declare([]cony.Declaration{
		cony.DeclareExchange(exc),
	})

	if err != nil {
		return errors.Wrap(err, "error declaring bitbon agent exchange")
	}

	// ------------------- Journal
	a.journal, err = bitbonAmqp.NewPublishJournal(a.journalPath,
		a.cfg.JournalTimeout, a.cfg.AmqpConnStrs, log.New("component", "bitbon agent journal"))
	if err != nil {
		return errors.Wrap(err, "error creating bitbon agent PublishJournal")
	}

	go a.watchAssetboxInfo(assetboxInfoChangedRK)
	go a.watchAssetboxBalance(assetboxBalanceChangedRK)
	go a.watchTransaction(transactionsChangedRK)
	go a.watchExpiredTransfers(transferExpiredRK)
	go a.watchBlocks(blocksMinedRK)

	log.Info("Bitbon agent service started")
	return nil
}

func (a *Agent) Stop() error {
	log.Info("Bitbon Agent service is waiting for all tasks are finished...")

	// at first notify all client workers to terminate
	if a.Stopped() {
		log.Error("Bitbon Agent already stopped")
	} else {
		close(a.done)
	}

	// ------------------- Stop journal
	if a.journal != nil {
		a.journal.Stop()
		log.Info("Bitbon Agent journal stopped")
	}
	// ------------------- Stop amqpConsumeClient
	if a.amqpClient != nil {
		a.amqpClient.Close()
		log.Info("Bitbon Agent amqpConsumeClient closed")
	}

	log.Info("Bitbon Agent service stopped")
	return nil
}

func (a *Agent) Stopped() bool {
	select {
	case <-a.done:
		return true
	default:
		return false
	}
}

func (a *Agent) GetJournal() *bitbonAmqp.PublishJournal {
	return a.journal
}

// APIs is agent meaningless implementation of node.Service.
func (a *Agent) APIs() []rpc.API {
	return []rpc.API{
		{
			Namespace: ProtocolName,
			Version:   ProtocolVersionStr,
			Service:   NewAPI(a),
			Public:    true,
		},
	}
}

// Protocols is agent meaningless implementation of node.Service.
func (a *Agent) Protocols() []p2p.Protocol {
	return nil
}
