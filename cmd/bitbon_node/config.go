// Copyright Simcord LLC
// This file is part of Bitbon System Blockchain Node.
// This file has been modified, you can find the original version by following the link
// <https://github.com/ethereum/go-ethereum>
//
// Bitbon System Blockchain Node is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// Bitbon System Blockchain Node is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with Bitbon System Blockchain Node. If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"reflect"
	"unicode"

	"github.com/naoina/toml"
	"github.com/simcord-llc/bitbon-system-blockchain/bitbon"
	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/agent"
	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/client"
	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/noncer"
	"github.com/simcord-llc/bitbon-system-blockchain/cmd/utils"
	"github.com/simcord-llc/bitbon-system-blockchain/eth"
	"github.com/simcord-llc/bitbon-system-blockchain/node"
	"github.com/simcord-llc/bitbon-system-blockchain/params"
	"github.com/simcord-llc/bitbon-system-blockchain/quorum"
	"gopkg.in/urfave/cli.v1"
)

var (
	dumpConfigCommand = cli.Command{
		Action:      utils.MigrateFlags(dumpConfig),
		Name:        "dumpconfig",
		Usage:       "Show configuration values",
		ArgsUsage:   "",
		Flags:       append(nodeFlags, rpcFlags...),
		Category:    "MISCELLANEOUS COMMANDS",
		Description: `The dumpconfig command shows configuration values.`,
	}

	configFileFlag = cli.StringFlag{
		Name:  "config",
		Usage: "TOML configuration file",
	}
)

// These settings ensure that TOML keys use the same names as Go struct fields.
var tomlSettings = toml.Config{
	NormFieldName: func(rt reflect.Type, key string) string {
		return key
	},
	FieldToKey: func(rt reflect.Type, field string) string {
		return field
	},
	MissingField: func(rt reflect.Type, field string) error {
		link := ""
		if unicode.IsUpper(rune(rt.Name()[0])) && rt.PkgPath() != "main" {
			link = fmt.Sprintf(", see https://godoc.org/%s#%s for available fields", rt.PkgPath(), rt.Name())
		}
		return fmt.Errorf("field '%s' is not defined in %s%s", field, rt.String(), link)
	},
}

type gethConfig struct {
	Eth          eth.Config
	Node         node.Config
	Quorum       quorum.Config
	Bitbon       bitbon.Config
	BitbonAgent  agent.Config
	BitbonClient client.Config
	BitbonNoncer noncer.Config
}

func loadConfig(file string, cfg *gethConfig) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	err = tomlSettings.NewDecoder(bufio.NewReader(f)).Decode(cfg)
	// Add file name to errors that have a line number.
	if _, ok := err.(*toml.LineError); ok {
		err = errors.New(file + ", " + err.Error())
	}
	return err
}

func defaultNodeConfig() node.Config {
	cfg := node.DefaultConfig
	cfg.Name = clientIdentifier
	cfg.Version = params.VersionWithCommit(gitCommit, gitDate)
	cfg.HTTPModules = append(cfg.HTTPModules, "eth")
	cfg.WSModules = append(cfg.WSModules, "eth")
	cfg.IPCPath = "geth.ipc"
	return cfg
}

func makeConfigNode(ctx *cli.Context) (*node.Node, gethConfig) {
	// Load defaults.
	cfg := gethConfig{
		Eth:    eth.DefaultConfig,
		Node:   defaultNodeConfig(),
		Quorum: quorum.DefaultConfig,
		Bitbon: bitbon.DefaultConfig,
	}

	// Load config file.
	if file := ctx.GlobalString(configFileFlag.Name); file != "" {
		if err := loadConfig(file, &cfg); err != nil {
			utils.Fatalf("%v", err)
		}
	}

	// Apply flags.
	utils.SetNodeConfig(ctx, &cfg.Node)
	stack, err := node.New(&cfg.Node)
	if err != nil {
		utils.Fatalf("Failed to create the protocol stack: %v", err)
	}
	utils.SetEthConfig(ctx, stack, &cfg.Eth)

	// quorum
	utils.SetQuorumConfig(ctx, &cfg.Quorum)

	// bitbon
	utils.SetBitbonConfig(ctx, &cfg.Bitbon)
	utils.SetBitbonClientConfig(ctx, &cfg.BitbonClient)
	utils.SetBitbonAgentConfig(ctx, &cfg.BitbonAgent)

	// noncer
	utils.SetNoncerConfig(ctx, &cfg.BitbonNoncer)

	return stack, cfg
}

func makeFullNode(ctx *cli.Context) *node.Node {
	stack, cfg := makeConfigNode(ctx)

	// Register quorum
	utils.RegisterQuorumService(stack, &cfg.Quorum)

	utils.RegisterEthService(stack, &cfg.Eth)

	//Register Bitbon Noncer if requested
	if noncer.IsAllowed() && cfg.BitbonNoncer.Enabled {
		utils.RegisterBitbonNoncerService(stack, &cfg.BitbonNoncer)
	}

	// Register bitbon
	utils.RegisterBitbonService(stack, &cfg.Bitbon)

	// Register bitbon agent if requested
	if agent.IsAllowed() && cfg.BitbonAgent.Enabled {
		utils.RegisterBitbonAgentService(stack, &cfg.BitbonAgent)
	}

	if client.IsAllowed() && cfg.BitbonClient.Enabled {
		utils.RegisterBitbonClientService(stack, &cfg.BitbonClient)
	}

	return stack
}

// dumpConfig is the dumpconfig command.
func dumpConfig(ctx *cli.Context) error {
	_, cfg := makeConfigNode(ctx)
	comment := ""

	if cfg.Eth.Genesis != nil {
		cfg.Eth.Genesis = nil
		comment += "# Note: this config doesn't contain the genesis block.\n\n"
	}

	out, err := tomlSettings.Marshal(&cfg)
	if err != nil {
		return err
	}

	dump := os.Stdout
	if ctx.NArg() > 0 {
		dump, err = os.OpenFile(ctx.Args().Get(0), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			return err
		}
		defer dump.Close()
	}
	dump.WriteString(comment)
	dump.Write(out)

	return nil
}
