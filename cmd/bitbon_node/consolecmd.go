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
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/simcord-llc/bitbon-system-blockchain/cmd/utils"
	"github.com/simcord-llc/bitbon-system-blockchain/console"
	"github.com/simcord-llc/bitbon-system-blockchain/node"
	"github.com/simcord-llc/bitbon-system-blockchain/rpc"
	"gopkg.in/urfave/cli.v1"
)

var (
	consoleFlags = []cli.Flag{utils.JSpathFlag, utils.ExecFlag, utils.PreloadJSFlag}

	consoleCommand = cli.Command{
		Action:   utils.MigrateFlags(localConsole),
		Name:     "console",
		Usage:    "Start an interactive JavaScript environment",
		Flags:    append(append(append(append(append(nodeFlags, rpcFlags...), consoleFlags...)), bitbonFlags...), quorumFlags...),
		Category: "CONSOLE COMMANDS",
		Description: `
The Geth console is an interactive shell for the JavaScript runtime environment
which exposes a node admin interface as well as the Ðapp JavaScript API.
See https://github.com/simcord-llc/bitbon-system-blockchain/wiki/JavaScript-Console.`,
	}

	attachCommand = cli.Command{
		Action:    utils.MigrateFlags(remoteConsole),
		Name:      "attach",
		Usage:     "Start an interactive JavaScript environment (connect to node)",
		ArgsUsage: "[endpoint]",
		Flags:     append(consoleFlags, utils.DataDirFlag),
		Category:  "CONSOLE COMMANDS",
		Description: `
The Geth console is an interactive shell for the JavaScript runtime environment
which exposes a node admin interface as well as the Ðapp JavaScript API.
See https://github.com/simcord-llc/bitbon-system-blockchain/wiki/JavaScript-Console.
This command allows to open a console on a running geth node.`,
	}

	javascriptCommand = cli.Command{
		Action:    utils.MigrateFlags(ephemeralConsole),
		Name:      "js",
		Usage:     "Execute the specified JavaScript files",
		ArgsUsage: "<jsfile> [jsfile...]",
		Flags:     append(nodeFlags, consoleFlags...),
		Category:  "CONSOLE COMMANDS",
		Description: `
The JavaScript VM exposes a node admin interface as well as the Ðapp
JavaScript API. See https://github.com/simcord-llc/bitbon-system-blockchain/wiki/JavaScript-Console`,
	}
)

// localConsole starts a new geth node, attaching a JavaScript console to it at the
// same time.
func localConsole(ctx *cli.Context) error {
	// Create and start the node based on the CLI flags
	prepare(ctx)
	fullNode := makeFullNode(ctx)
	startNode(ctx, fullNode)
	defer fullNode.Close()

	// Attach to the newly started node and start the JavaScript console
	client, err := fullNode.Attach()
	if err != nil {
		utils.Fatalf("Failed to attach to the inproc geth: %v", err)
	}
	config := console.Config{
		DataDir: utils.MakeDataDir(ctx),
		DocRoot: ctx.GlobalString(utils.JSpathFlag.Name),
		Client:  client,
		Preload: utils.MakeConsolePreloads(ctx),
	}

	consoleObj, err := console.New(config)
	if err != nil {
		utils.Fatalf("Failed to start the JavaScript console: %v", err)
	}
	defer consoleObj.Stop(false)

	// If only a short execution was requested, evaluate and return
	if script := ctx.GlobalString(utils.ExecFlag.Name); script != "" {
		consoleObj.Evaluate(script)
		return nil
	}
	// Otherwise print the welcome screen and enter interactive mode
	consoleObj.Welcome()
	consoleObj.Interactive()

	return nil
}

// remoteConsole will connect to a remote geth instance, attaching a JavaScript
// console to it.
func remoteConsole(ctx *cli.Context) error {
	// Attach to a remotely running geth instance and start the JavaScript console
	endpoint := ctx.Args().First()
	if endpoint == "" {
		path := node.DefaultDataDir()
		if ctx.GlobalIsSet(utils.DataDirFlag.Name) {
			path = ctx.GlobalString(utils.DataDirFlag.Name)
		}
		if path != "" {
			if ctx.GlobalBool(utils.LegacyTestnetFlag.Name) || ctx.GlobalBool(utils.RopstenFlag.Name) { // nolint
				// Maintain compatibility with older Geth configurations storing the
				// Ropsten database in `testnet` instead of `ropsten`.
				legacyPath := filepath.Join(path, "testnet")
				if _, err := os.Stat(legacyPath); !os.IsNotExist(err) {
					path = legacyPath
				} else {
					path = filepath.Join(path, "ropsten")
				}
			} else if ctx.GlobalBool(utils.RinkebyFlag.Name) {
				path = filepath.Join(path, "rinkeby")
			} else if ctx.GlobalBool(utils.GoerliFlag.Name) {
				path = filepath.Join(path, "goerli")
			} else if ctx.GlobalBool(utils.YoloV1Flag.Name) {
				path = filepath.Join(path, "yolo-v1")
			}
		}
		endpoint = fmt.Sprintf("%s/geth.ipc", path)
	}
	client, err := dialRPC(endpoint)
	if err != nil {
		utils.Fatalf("Unable to attach to remote geth: %v", err)
	}
	config := console.Config{
		DataDir: utils.MakeDataDir(ctx),
		DocRoot: ctx.GlobalString(utils.JSpathFlag.Name),
		Client:  client,
		Preload: utils.MakeConsolePreloads(ctx),
	}

	consoleObj, err := console.New(config)
	if err != nil {
		utils.Fatalf("Failed to start the JavaScript console: %v", err)
	}
	defer consoleObj.Stop(false)

	if script := ctx.GlobalString(utils.ExecFlag.Name); script != "" {
		consoleObj.Evaluate(script)
		return nil
	}

	// Otherwise print the welcome screen and enter interactive mode
	consoleObj.Welcome()
	consoleObj.Interactive()

	return nil
}

// dialRPC returns a RPC client which connects to the given endpoint.
// The check for empty endpoint implements the defaulting logic
// for "geth attach" and "geth monitor" with no argument.
func dialRPC(endpoint string) (*rpc.Client, error) {
	if endpoint == "" {
		endpoint = node.DefaultIPCEndpoint(clientIdentifier)
	} else if strings.HasPrefix(endpoint, "rpc:") || strings.HasPrefix(endpoint, "ipc:") {
		// Backwards compatibility with geth < 1.5 which required
		// these prefixes.
		endpoint = endpoint[4:]
	}
	return rpc.Dial(endpoint)
}

// ephemeralConsole starts a new geth node, attaches an ephemeral JavaScript
// console to it, executes each of the files specified as arguments and tears
// everything down.
func ephemeralConsole(ctx *cli.Context) error {
	// Create and start the node based on the CLI flags
	fullNode := makeFullNode(ctx)
	startNode(ctx, fullNode)
	defer fullNode.Close()

	// Attach to the newly started node and start the JavaScript console
	client, err := fullNode.Attach()
	if err != nil {
		utils.Fatalf("Failed to attach to the inproc geth: %v", err)
	}
	config := console.Config{
		DataDir: utils.MakeDataDir(ctx),
		DocRoot: ctx.GlobalString(utils.JSpathFlag.Name),
		Client:  client,
		Preload: utils.MakeConsolePreloads(ctx),
	}

	consoleObj, err := console.New(config)
	if err != nil {
		utils.Fatalf("Failed to start the JavaScript console: %v", err)
	}
	defer consoleObj.Stop(false)

	// Evaluate each of the specified JavaScript files
	for _, file := range ctx.Args() {
		if err = consoleObj.Execute(file); err != nil {
			utils.Fatalf("Failed to execute %s: %v", file, err)
		}
	}
	// Wait for pending callbacks, but stop for Ctrl-C.
	abort := make(chan os.Signal, 1)
	signal.Notify(abort, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-abort
		os.Exit(0)
	}()
	consoleObj.Stop(true)

	return nil
}
