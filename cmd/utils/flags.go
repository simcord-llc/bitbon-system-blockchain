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

// Package utils contains internal helper functions for bitbon-system-blockchain commands.
package utils

import (
	"context"
	"crypto/ecdsa"
	"fmt"

	"github.com/simcord-llc/bitbon-system-blockchain/crypto/aes"

	"io"
	"io/ioutil"
	"math/big"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"text/tabwriter"
	"text/template"
	"time"

	pcsclite "github.com/gballet/go-libpcsclite"
	"github.com/pkg/errors"
	"github.com/simcord-llc/bitbon-system-blockchain/accounts"
	"github.com/simcord-llc/bitbon-system-blockchain/accounts/keystore"
	"github.com/simcord-llc/bitbon-system-blockchain/bitbon"
	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/agent"
	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/assetbox"
	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/client"
	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/contracts"
	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/mining_agent"
	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/noncer"
	"github.com/simcord-llc/bitbon-system-blockchain/bitbon/transfer"
	"github.com/simcord-llc/bitbon-system-blockchain/common"
	"github.com/simcord-llc/bitbon-system-blockchain/common/fdlimit"
	"github.com/simcord-llc/bitbon-system-blockchain/consensus"
	"github.com/simcord-llc/bitbon-system-blockchain/consensus/clique"
	"github.com/simcord-llc/bitbon-system-blockchain/consensus/ethash"
	quorumEngine "github.com/simcord-llc/bitbon-system-blockchain/consensus/quorum"
	"github.com/simcord-llc/bitbon-system-blockchain/core"
	"github.com/simcord-llc/bitbon-system-blockchain/core/vm"
	"github.com/simcord-llc/bitbon-system-blockchain/crypto"
	"github.com/simcord-llc/bitbon-system-blockchain/eth"
	"github.com/simcord-llc/bitbon-system-blockchain/eth/downloader"
	"github.com/simcord-llc/bitbon-system-blockchain/eth/gasprice"
	"github.com/simcord-llc/bitbon-system-blockchain/ethdb"
	"github.com/simcord-llc/bitbon-system-blockchain/les"
	"github.com/simcord-llc/bitbon-system-blockchain/log"
	"github.com/simcord-llc/bitbon-system-blockchain/metrics"
	"github.com/simcord-llc/bitbon-system-blockchain/metrics/influxdb"
	"github.com/simcord-llc/bitbon-system-blockchain/miner"
	"github.com/simcord-llc/bitbon-system-blockchain/node"
	"github.com/simcord-llc/bitbon-system-blockchain/p2p"
	"github.com/simcord-llc/bitbon-system-blockchain/p2p/discv5"
	"github.com/simcord-llc/bitbon-system-blockchain/p2p/enode"
	"github.com/simcord-llc/bitbon-system-blockchain/p2p/nat"
	"github.com/simcord-llc/bitbon-system-blockchain/p2p/netutil"
	"github.com/simcord-llc/bitbon-system-blockchain/params"
	"github.com/simcord-llc/bitbon-system-blockchain/quorum"
	"gopkg.in/urfave/cli.v1"
)

var (
	CommandHelpTemplate = `{{.cmd.Name}}{{if .cmd.Subcommands}} command{{end}}{{if .cmd.Flags}} [command options]{{end}} [arguments...]
{{if .cmd.Description}}{{.cmd.Description}}
{{end}}{{if .cmd.Subcommands}}
SUBCOMMANDS:
	{{range .cmd.Subcommands}}{{.Name}}{{with .ShortName}}, {{.}}{{end}}{{ "\t" }}{{.Usage}}
	{{end}}{{end}}{{if .categorizedFlags}}
{{range $idx, $categorized := .categorizedFlags}}{{$categorized.Name}} OPTIONS:
{{range $categorized.Flags}}{{"\t"}}{{.}}
{{end}}
{{end}}{{end}}`

	OriginCommandHelpTemplate = `{{.Name}}{{if .Subcommands}} command{{end}}{{if .Flags}} [command options]{{end}} [arguments...]
{{if .Description}}{{.Description}}
{{end}}{{if .Subcommands}}
SUBCOMMANDS:
	{{range .Subcommands}}{{.Name}}{{with .ShortName}}, {{.}}{{end}}{{ "\t" }}{{.Usage}}
	{{end}}{{end}}{{if .Flags}}
OPTIONS:
{{range $.Flags}}{{"\t"}}{{.}}
{{end}}
{{end}}`
)

func init() {
	cli.AppHelpTemplate = `{{.Name}} {{if .Flags}}[global options] {{end}}command{{if .Flags}} [command options]{{end}} [arguments...]

VERSION:
   {{.Version}}

COMMANDS:
   {{range .Commands}}{{.Name}}{{with .ShortName}}, {{.}}{{end}}{{ "\t" }}{{.Usage}}
   {{end}}{{if .Flags}}
GLOBAL OPTIONS:
   {{range .Flags}}{{.}}
   {{end}}{{end}}
`
	cli.CommandHelpTemplate = CommandHelpTemplate
	cli.HelpPrinter = printHelp
}

// NewApp creates an app with sane defaults.
func NewApp(gitCommit, gitDate, usage string) *cli.App {
	app := cli.NewApp()
	app.Name = filepath.Base(os.Args[0])
	app.Author = ""
	app.Email = ""
	app.Version = params.VersionWithCommit(gitCommit, gitDate)
	app.Usage = usage
	return app
}

func printHelp(out io.Writer, templ string, data interface{}) {
	funcMap := template.FuncMap{"join": strings.Join}
	t := template.Must(template.New("help").Funcs(funcMap).Parse(templ))
	w := tabwriter.NewWriter(out, 38, 8, 2, ' ', 0)
	err := t.Execute(w, data)
	if err != nil {
		panic(err)
	}
	w.Flush()
}

// These are all the command line flags we support.
// If you add to this list, please remember to include the
// flag in the appropriate command definition.
//
// The flags are defined here so their names and help texts
// are the same for all commands.

var (
	// General settings
	DataDirFlag = DirectoryFlag{
		Name:  "datadir",
		Usage: "Data directory for the databases and keystore",
		Value: DirectoryString(node.DefaultDataDir()),
	}
	AncientFlag = DirectoryFlag{
		Name:  "datadir.ancient",
		Usage: "Data directory for ancient chain segments (default = inside chaindata)",
	}
	KeyStoreDirFlag = DirectoryFlag{
		Name:  "keystore",
		Usage: "Directory for the keystore (default = inside the datadir)",
	}
	NoUSBFlag = cli.BoolFlag{
		Name:  "nousb",
		Usage: "Disables monitoring for and managing USB hardware wallets",
	}
	SmartCardDaemonPathFlag = cli.StringFlag{
		Name:  "pcscdpath",
		Usage: "Path to the smartcard daemon (pcscd) socket file",
		Value: pcsclite.PCSCDSockName,
	}
	NetworkIdFlag = cli.Uint64Flag{
		Name:  "networkid",
		Usage: "Network identifier (integer, 1=Frontier, 3=Ropsten, 4=Rinkeby, 5=Görli)",
		Value: eth.DefaultConfig.NetworkId,
	}
	GoerliFlag = cli.BoolFlag{
		Name:  "goerli",
		Usage: "Görli network: pre-configured proof-of-authority test network",
	}
	YoloV1Flag = cli.BoolFlag{
		Name:  "yolov1",
		Usage: "YOLOv1 network: pre-configured proof-of-authority shortlived test network.",
	}
	RinkebyFlag = cli.BoolFlag{
		Name:  "rinkeby",
		Usage: "Rinkeby network: pre-configured proof-of-authority test network",
	}
	RopstenFlag = cli.BoolFlag{
		Name:  "ropsten",
		Usage: "Ropsten network: pre-configured proof-of-work test network",
	}
	DeveloperFlag = cli.BoolFlag{
		Name:  "dev",
		Usage: "Ephemeral proof-of-authority network with a pre-funded developer account, mining enabled",
	}
	DeveloperPeriodFlag = cli.IntFlag{
		Name:  "dev.period",
		Usage: "Block period to use in developer mode (0 = mine only if transaction pending)",
	}
	IdentityFlag = cli.StringFlag{
		Name:  "identity",
		Usage: "Custom node name",
	}
	DocRootFlag = DirectoryFlag{
		Name:  "docroot",
		Usage: "Document Root for HTTPClient file scheme",
		Value: DirectoryString(homeDir()),
	}
	ExitWhenSyncedFlag = cli.BoolFlag{
		Name:  "exitwhensynced",
		Usage: "Exits after block synchronisation completes",
	}
	IterativeOutputFlag = cli.BoolFlag{
		Name:  "iterative",
		Usage: "Print streaming JSON iteratively, delimited by newlines",
	}
	ExcludeStorageFlag = cli.BoolFlag{
		Name:  "nostorage",
		Usage: "Exclude storage entries (save db lookups)",
	}
	IncludeIncompletesFlag = cli.BoolFlag{
		Name:  "incompletes",
		Usage: "Include accounts for which we don't have the address (missing preimage)",
	}
	ExcludeCodeFlag = cli.BoolFlag{
		Name:  "nocode",
		Usage: "Exclude contract code (save db lookups)",
	}
	defaultSyncMode = eth.DefaultConfig.SyncMode
	SyncModeFlag    = TextMarshalerFlag{
		Name:  "syncmode",
		Usage: `Blockchain sync mode ("fast", "full", or "light")`,
		Value: &defaultSyncMode,
	}
	GCModeFlag = cli.StringFlag{
		Name:  "gcmode",
		Usage: `Blockchain garbage collection mode ("full", "archive")`,
		Value: "archive",
	}
	SnapshotFlag = cli.BoolFlag{
		Name:  "snapshot",
		Usage: `Enables snapshot-database mode -- experimental work in progress feature`,
	}
	TxLookupLimitFlag = cli.Int64Flag{
		Name:  "txlookuplimit",
		Usage: "Number of recent blocks to maintain transactions index by-hash for (default = index all blocks)",
		Value: 0,
	}
	LightKDFFlag = cli.BoolFlag{
		Name:  "lightkdf",
		Usage: "Reduce key-derivation RAM & CPU usage at some expense of KDF strength",
	}
	WhitelistFlag = cli.StringFlag{
		Name:  "whitelist",
		Usage: "Comma separated block number-to-hash mappings to enforce (<number>=<hash>)",
	}
	// Light server and client settings
	LightServeFlag = cli.IntFlag{
		Name:  "light.serve",
		Usage: "Maximum percentage of time allowed for serving LES requests (multi-threaded processing allows values over 100)",
		Value: eth.DefaultConfig.LightServ,
	}
	LightIngressFlag = cli.IntFlag{
		Name:  "light.ingress",
		Usage: "Incoming bandwidth limit for serving light clients (kilobytes/sec, 0 = unlimited)",
		Value: eth.DefaultConfig.LightIngress,
	}
	LightEgressFlag = cli.IntFlag{
		Name:  "light.egress",
		Usage: "Outgoing bandwidth limit for serving light clients (kilobytes/sec, 0 = unlimited)",
		Value: eth.DefaultConfig.LightEgress,
	}
	LightMaxPeersFlag = cli.IntFlag{
		Name:  "light.maxpeers",
		Usage: "Maximum number of light clients to serve, or light servers to attach to",
		Value: eth.DefaultConfig.LightPeers,
	}
	UltraLightServersFlag = cli.StringFlag{
		Name:  "ulc.servers",
		Usage: "List of trusted ultra-light servers",
		Value: strings.Join(eth.DefaultConfig.UltraLightServers, ","),
	}
	UltraLightFractionFlag = cli.IntFlag{
		Name:  "ulc.fraction",
		Usage: "Minimum % of trusted ultra-light servers required to announce a new head",
		Value: eth.DefaultConfig.UltraLightFraction,
	}
	UltraLightOnlyAnnounceFlag = cli.BoolFlag{
		Name:  "ulc.onlyannounce",
		Usage: "Ultra light server sends announcements only",
	}
	// Ethash settings
	EthashCacheDirFlag = DirectoryFlag{
		Name:  "ethash.cachedir",
		Usage: "Directory to store the ethash verification caches (default = inside the datadir)",
	}
	EthashCachesInMemoryFlag = cli.IntFlag{
		Name:  "ethash.cachesinmem",
		Usage: "Number of recent ethash caches to keep in memory (16MB each)",
		Value: eth.DefaultConfig.Ethash.CachesInMem,
	}
	EthashCachesOnDiskFlag = cli.IntFlag{
		Name:  "ethash.cachesondisk",
		Usage: "Number of recent ethash caches to keep on disk (16MB each)",
		Value: eth.DefaultConfig.Ethash.CachesOnDisk,
	}
	EthashCachesLockMmapFlag = cli.BoolFlag{
		Name:  "ethash.cacheslockmmap",
		Usage: "Lock memory maps of recent ethash caches",
	}
	EthashDatasetDirFlag = DirectoryFlag{
		Name:  "ethash.dagdir",
		Usage: "Directory to store the ethash mining DAGs",
		Value: DirectoryString(eth.DefaultConfig.Ethash.DatasetDir),
	}
	EthashDatasetsInMemoryFlag = cli.IntFlag{
		Name:  "ethash.dagsinmem",
		Usage: "Number of recent ethash mining DAGs to keep in memory (1+GB each)",
		Value: eth.DefaultConfig.Ethash.DatasetsInMem,
	}
	EthashDatasetsOnDiskFlag = cli.IntFlag{
		Name:  "ethash.dagsondisk",
		Usage: "Number of recent ethash mining DAGs to keep on disk (1+GB each)",
		Value: eth.DefaultConfig.Ethash.DatasetsOnDisk,
	}
	EthashDatasetsLockMmapFlag = cli.BoolFlag{
		Name:  "ethash.dagslockmmap",
		Usage: "Lock memory maps for recent ethash mining DAGs",
	}
	// Transaction pool settings
	TxPoolLocalsFlag = cli.StringFlag{
		Name:  "txpool.locals",
		Usage: "Comma separated accounts to treat as locals (no flush, priority inclusion)",
	}
	TxPoolNoLocalsFlag = cli.BoolFlag{
		Name:  "txpool.nolocals",
		Usage: "Disables price exemptions for locally submitted transactions",
	}
	TxPoolJournalFlag = cli.StringFlag{
		Name:  "txpool.journal",
		Usage: "Disk journal for local transaction to survive node restarts",
		Value: core.DefaultTxPoolConfig.Journal,
	}
	TxPoolRejournalFlag = cli.DurationFlag{
		Name:  "txpool.rejournal",
		Usage: "Time interval to regenerate the local transaction journal",
		Value: core.DefaultTxPoolConfig.Rejournal,
	}
	TxPoolPriceLimitFlag = cli.Uint64Flag{
		Name:  "txpool.pricelimit",
		Usage: "Minimum gas price limit to enforce for acceptance into the pool",
		Value: eth.DefaultConfig.TxPool.PriceLimit,
	}
	TxPoolPriceBumpFlag = cli.Uint64Flag{
		Name:  "txpool.pricebump",
		Usage: "Price bump percentage to replace an already existing transaction",
		Value: eth.DefaultConfig.TxPool.PriceBump,
	}
	TxPoolAccountSlotsFlag = cli.Uint64Flag{
		Name:  "txpool.accountslots",
		Usage: "Minimum number of executable transaction slots guaranteed per account",
		Value: eth.DefaultConfig.TxPool.AccountSlots,
	}
	TxPoolGlobalSlotsFlag = cli.Uint64Flag{
		Name:  "txpool.globalslots",
		Usage: "Maximum number of executable transaction slots for all accounts",
		Value: eth.DefaultConfig.TxPool.GlobalSlots,
	}
	TxPoolAccountQueueFlag = cli.Uint64Flag{
		Name:  "txpool.accountqueue",
		Usage: "Maximum number of non-executable transaction slots permitted per account",
		Value: eth.DefaultConfig.TxPool.AccountQueue,
	}
	TxPoolGlobalQueueFlag = cli.Uint64Flag{
		Name:  "txpool.globalqueue",
		Usage: "Maximum number of non-executable transaction slots for all accounts",
		Value: eth.DefaultConfig.TxPool.GlobalQueue,
	}
	TxPoolLifetimeFlag = cli.DurationFlag{
		Name:  "txpool.lifetime",
		Usage: "Maximum amount of time non-executable transaction are queued",
		Value: eth.DefaultConfig.TxPool.Lifetime,
	}
	// Performance tuning settings
	CacheFlag = cli.IntFlag{
		Name:  "cache",
		Usage: "Megabytes of memory allocated to internal caching (default = 4096 mainnet full node, 128 light mode)",
		Value: 1024,
	}
	CacheDatabaseFlag = cli.IntFlag{
		Name:  "cache.database",
		Usage: "Percentage of cache memory allowance to use for database io",
		Value: 50,
	}
	CacheTrieFlag = cli.IntFlag{
		Name:  "cache.trie",
		Usage: "Percentage of cache memory allowance to use for trie caching (default = 15% full mode, 30% archive mode)",
		Value: 15,
	}
	CacheGCFlag = cli.IntFlag{
		Name:  "cache.gc",
		Usage: "Percentage of cache memory allowance to use for trie pruning (default = 25% full mode, 0% archive mode)",
		Value: 25,
	}
	CacheSnapshotFlag = cli.IntFlag{
		Name:  "cache.snapshot",
		Usage: "Percentage of cache memory allowance to use for snapshot caching (default = 10% full mode, 20% archive mode)",
		Value: 10,
	}
	CacheNoPrefetchFlag = cli.BoolFlag{
		Name:  "cache.noprefetch",
		Usage: "Disable heuristic state prefetch during block import (less CPU and disk IO, more time waiting for data)",
	}
	// Miner settings
	MiningEnabledFlag = cli.BoolFlag{
		Name:  "mine",
		Usage: "Enable mining",
	}
	MinerThreadsFlag = cli.IntFlag{
		Name:  "miner.threads",
		Usage: "Number of CPU threads to use for mining",
		Value: 0,
	}
	MinerNotifyFlag = cli.StringFlag{
		Name:  "miner.notify",
		Usage: "Comma separated HTTP URL list to notify of new work packages",
	}
	MinerGasTargetFlag = cli.Uint64Flag{
		Name:  "miner.gastarget",
		Usage: "Target gas floor for mined blocks",
		Value: eth.DefaultConfig.Miner.GasFloor,
	}
	MinerGasLimitFlag = cli.Uint64Flag{
		Name:  "miner.gaslimit",
		Usage: "Target gas ceiling for mined blocks",
		Value: eth.DefaultConfig.Miner.GasCeil,
	}
	MinerGasPriceFlag = BigFlag{
		Name:  "miner.gasprice",
		Usage: "Minimum gas price for mining a transaction",
		Value: eth.DefaultConfig.Miner.GasPrice,
	}
	MinerEtherbaseFlag = cli.StringFlag{
		Name:  "miner.etherbase",
		Usage: "Public address for block mining rewards (default = first account)",
		Value: "0",
	}
	MinerExtraDataFlag = cli.StringFlag{
		Name:  "miner.extradata",
		Usage: "Block extra data set by the miner (default = client version)",
	}
	MinerRecommitIntervalFlag = cli.DurationFlag{
		Name:  "miner.recommit",
		Usage: "Time interval to recreate the block being mined",
		Value: eth.DefaultConfig.Miner.Recommit,
	}
	MinerNoVerfiyFlag = cli.BoolFlag{
		Name:  "miner.noverify",
		Usage: "Disable remote sealing verification",
	}
	// Account settings
	UnlockedAccountFlag = cli.StringFlag{
		Name:  "unlock",
		Usage: "Comma separated list of accounts to unlock",
		Value: "",
	}
	PasswordFileFlag = cli.StringFlag{
		Name:  "password",
		Usage: "Password file to use for non-interactive password input",
		Value: "",
	}
	ExternalSignerFlag = cli.StringFlag{
		Name:  "signer",
		Usage: "External signer (url or path to ipc file)",
		Value: "",
	}
	VMEnableDebugFlag = cli.BoolFlag{
		Name:  "vmdebug",
		Usage: "Record information useful for VM and contract debugging",
	}
	InsecureUnlockAllowedFlag = cli.BoolFlag{
		Name:  "allow-insecure-unlock",
		Usage: "Allow insecure account unlocking when account-related RPCs are exposed by http",
	}
	RPCGlobalGasCap = cli.Uint64Flag{
		Name:  "rpc.gascap",
		Usage: "Sets a cap on gas that can be used in eth_call/estimateGas",
	}
	FakePoWFlag = cli.BoolFlag{
		Name:  "fakepow",
		Usage: "Disables proof-of-work verification",
	}
	NoCompactionFlag = cli.BoolFlag{
		Name:  "nocompaction",
		Usage: "Disables db compaction after import",
	}
	// RPC settings
	IPCDisabledFlag = cli.BoolFlag{
		Name:  "ipcdisable",
		Usage: "Disable the IPC-RPC server",
	}
	IPCPathFlag = DirectoryFlag{
		Name:  "ipcpath",
		Usage: "Filename for IPC socket/pipe within the datadir (explicit paths escape it)",
	}
	HTTPEnabledFlag = cli.BoolFlag{
		Name:  "http",
		Usage: "Enable the HTTP-RPC server",
	}
	HTTPListenAddrFlag = cli.StringFlag{
		Name:  "http.addr",
		Usage: "HTTP-RPC server listening interface",
		Value: node.DefaultHTTPHost,
	}
	HTTPPortFlag = cli.IntFlag{
		Name:  "http.port",
		Usage: "HTTP-RPC server listening port",
		Value: node.DefaultHTTPPort,
	}
	HTTPCORSDomainFlag = cli.StringFlag{
		Name:  "http.corsdomain",
		Usage: "Comma separated list of domains from which to accept cross origin requests (browser enforced)",
		Value: "",
	}
	HTTPVirtualHostsFlag = cli.StringFlag{
		Name:  "http.vhosts",
		Usage: "Comma separated list of virtual hostnames from which to accept requests (server enforced). Accepts '*' wildcard.",
		Value: strings.Join(node.DefaultConfig.HTTPVirtualHosts, ","),
	}
	HTTPApiFlag = cli.StringFlag{
		Name:  "http.api",
		Usage: "API's offered over the HTTP-RPC interface",
		Value: "",
	}
	WSEnabledFlag = cli.BoolFlag{
		Name:  "ws",
		Usage: "Enable the WS-RPC server",
	}
	WSListenAddrFlag = cli.StringFlag{
		Name:  "ws.addr",
		Usage: "WS-RPC server listening interface",
		Value: node.DefaultWSHost,
	}
	WSPortFlag = cli.IntFlag{
		Name:  "ws.port",
		Usage: "WS-RPC server listening port",
		Value: node.DefaultWSPort,
	}
	WSApiFlag = cli.StringFlag{
		Name:  "ws.api",
		Usage: "API's offered over the WS-RPC interface",
		Value: "",
	}
	WSAllowedOriginsFlag = cli.StringFlag{
		Name:  "ws.origins",
		Usage: "Origins from which to accept websockets requests",
		Value: "",
	}
	ExecFlag = cli.StringFlag{
		Name:  "exec",
		Usage: "Execute JavaScript statement",
	}
	PreloadJSFlag = cli.StringFlag{
		Name:  "preload",
		Usage: "Comma separated list of JavaScript files to preload into the console",
	}

	// Network Settings
	MaxPeersFlag = cli.IntFlag{
		Name:  "maxpeers",
		Usage: "Maximum number of network peers (network disabled if set to 0)",
		Value: node.DefaultConfig.P2P.MaxPeers,
	}
	MaxPendingPeersFlag = cli.IntFlag{
		Name:  "maxpendpeers",
		Usage: "Maximum number of pending connection attempts (defaults used if set to 0)",
		Value: node.DefaultConfig.P2P.MaxPendingPeers,
	}
	ListenPortFlag = cli.IntFlag{
		Name:  "port",
		Usage: "Network listening port",
		Value: 30303,
	}
	BootnodesFlag = cli.StringFlag{
		Name:  "bootnodes",
		Usage: "Comma separated enode URLs for P2P discovery bootstrap",
		Value: "",
	}
	NodeKeyFileFlag = cli.StringFlag{
		Name:  "nodekey",
		Usage: "P2P node key file",
	}
	NodeKeyHexFlag = cli.StringFlag{
		Name:  "nodekeyhex",
		Usage: "P2P node key as hex (for testing)",
	}
	NATFlag = cli.StringFlag{
		Name:  "nat",
		Usage: "NAT port mapping mechanism (any|none|upnp|pmp|extip:<IP>)",
		Value: "any",
	}
	NoDiscoverFlag = cli.BoolFlag{
		Name:  "nodiscover",
		Usage: "Disables the peer discovery mechanism (manual peer addition)",
	}
	DiscoveryV5Flag = cli.BoolFlag{
		Name:  "v5disc",
		Usage: "Enables the experimental RLPx V5 (Topic Discovery) mechanism",
	}
	NetrestrictFlag = cli.StringFlag{
		Name:  "netrestrict",
		Usage: "Restricts network communication to the given IP networks (CIDR masks)",
	}
	DNSDiscoveryFlag = cli.StringFlag{
		Name:  "discovery.dns",
		Usage: "Sets DNS discovery entry points (use \"\" to disable DNS)",
	}

	// ATM the url is left to the user and deployment to
	JSpathFlag = cli.StringFlag{
		Name:  "jspath",
		Usage: "JavaScript root path for `loadScript`",
		Value: ".",
	}

	// Gas price oracle settings
	GpoBlocksFlag = cli.IntFlag{
		Name:  "gpo.blocks",
		Usage: "Number of recent blocks to check for gas prices",
		Value: eth.DefaultConfig.GPO.Blocks,
	}
	GpoPercentileFlag = cli.IntFlag{
		Name:  "gpo.percentile",
		Usage: "Suggested gas price is the given percentile of a set of recent transaction gas prices",
		Value: eth.DefaultConfig.GPO.Percentile,
	}

	// bitbon
	BitbonDecryptAssetboxWalletPassword = cli.StringFlag{
		Name:  "bitbon.assetbox.decrypt.password",
		Usage: "password for decrypting assetbox wallet",
	}
	BitbonServicePrivateKey = cli.StringFlag{
		Name:  "bitbon.service.pk",
		Usage: "hex of service private key (ether account)",
	}
	BitbonServiceID = cli.StringFlag{
		Name:  "bitbon.service.id",
		Usage: "ID of service",
	}
	BitbonAgentFlag = cli.BoolFlag{
		Name:  "bitbon.agent",
		Usage: "Enable bitbon agent service",
	}
	BitbonAgentAmqpConnectionFlag = cli.StringFlag{
		Name:  "bitbon.agent.amqp.connection",
		Usage: "RabbitMQ connection string",
		Value: "amqp://guest:guest@localhost:5672/",
	}
	BitbonAgentAmqpExchangeFlag = cli.StringFlag{
		Name:  "bitbon.agent.amqp.exchange",
		Usage: "RabbitMQ agent exchange",
		Value: "e.events.forward",
	}
	BitbonAgentJournalFlag = cli.StringFlag{
		Name:  "bitbon.agent.journal.name",
		Usage: "Journal file",
		Value: agent.DefaultJournalName,
	}
	BitbonAgentJournalTimeout = cli.DurationFlag{
		Name:  "bitbon.agent.journal.timeout",
		Usage: "Journal timeout",
		Value: agent.DefaultJournalTimeout,
	}

	BitbonClientModeFlag = cli.IntFlag{
		Name:  "bitbon.client.mode",
		Usage: "Set client work mode",
		Value: int(client.BitbonClientAllMode),
	}
	BitbonClientFlag = cli.BoolFlag{
		Name:  "bitbon.client",
		Usage: "Enable bitbon client service",
	}
	BitbonClientWaitSyncFlag = cli.BoolFlag{
		Name:  "bitbon.client.waitsync",
		Usage: "If this flag is true, bitbon client service will be waiting for sync to be finished before start. Otherwise - it won`t",
	}
	BitbonClientAmqpConnectionFlag = cli.StringFlag{
		Name:  "bitbon.client.amqp.connection",
		Usage: "RabbitMQ connection string",
		Value: "amqp://guest:guest@localhost:5672/",
	}
	BitbonClientAmqpAssetboxExchange = cli.StringFlag{
		Name:  "bitbon.client.amqp.assetboxexchange",
		Usage: "RabbitMQ assetbox exchange",
		Value: client.DefaultAssetboxExchange,
	}
	BitbonClientAmqpTransferExchange = cli.StringFlag{
		Name:  "bitbon.client.amqp.transferexchange",
		Usage: "RabbitMQ transfer exchange",
		Value: client.DefaultTransferExchange,
	}
	BitbonClientAmqpBlockExchange = cli.StringFlag{
		Name:  "bitbon.client.amqp.blockexchange",
		Usage: "RabbitMQ block exchange",
		Value: client.DefaultBlockExchange,
	}
	BitbonClientAmqpTransactionExchange = cli.StringFlag{
		Name:  "bitbon.client.amqp.transactionexchange",
		Usage: "RabbitMQ transaction exchange",
		Value: client.DefaultTransactionExchange,
	}
	BitbonClientAmqpMiningExchange = cli.StringFlag{
		Name:  "bitbon.client.amqp.miningexchange",
		Usage: "RabbitMQ mining exchange",
		Value: client.DefaultMiningExchange,
	}
	BitbonClientAmqpConfirmTimeout = cli.DurationFlag{
		Name:  "bitbon.client.amqp.confirmtimeout",
		Usage: "RabbitMQ confirm timeout while publishing",
		Value: client.DefaultConfirmTimeout,
	}
	BitbonClientAmqpPublishTimeout = cli.DurationFlag{
		Name:  "bitbon.client.amqp.publishtimeout",
		Usage: "RabbitMQ publish timeout",
		Value: client.DefaultPublishTimeout,
	}
	BitbonClientJournalTimeout = cli.DurationFlag{
		Name:  "bitbon.client.journal.timeout",
		Usage: "Journal timeout",
		Value: client.DefaultJournalTimeout,
	}
	BitbonClientJournalName = cli.StringFlag{
		Name:  "bitbon.client.journal.name",
		Usage: "Journal name",
		Value: client.DefaultJournalName,
	}

	//noncer
	BitbonNoncerFlag = cli.BoolFlag{
		Name:  "bitbon.noncer",
		Usage: "Enable bitbon noncer service",
	}

	BitbonNoncerKey = cli.StringFlag{
		Name:  "bitbon.noncer.redis.key",
		Usage: "Redis nonce key",
		Value: "nonces",
	}

	BitbonNoncerRedis = cli.StringFlag{
		Name:  "bitbon.noncer.redis.connection",
		Usage: "Comma separated list of redis connection strings",
	}

	BitbonNoncerRedisDialTimeout = cli.Uint64Flag{
		Name:  "bitbon.noncer.redis.dial.timeout",
		Usage: "Dial timeout for Redis connection",
		Value: noncer.DefaultDialTimeout,
	}

	BitbonNoncerRedisMaxRetries = cli.Uint64Flag{
		Name:  "bitbon.noncer.redis.dial.retries",
		Usage: "Max number of Redis dial retries",
		Value: noncer.DefaultMaxRetries,
	}

	QuorumPrepareMinersTimeoutFlag = cli.DurationFlag{
		Name:  "quorum.miners.timeout",
		Usage: "Timeout for quorum prepare miner list",
		Value: quorum.DefaultPrepareMinersTimeout,
	}

	QuorumPingFrequencyFlag = cli.DurationFlag{
		Name:  "quorum.ping.frequency",
		Usage: "Period to ping quorum nodes to establish their ping value",
		Value: quorum.DefaultPingFrequency,
	}

	QuorumMaxPingFlag = cli.DurationFlag{
		Name:  "quorum.ping.max",
		Usage: "Max value for ping to consider if the remote node in quorum or not",
		Value: quorum.DefaultQuorumMaxPing,
	}

	QuorumHandshakeTimeoutFlag = cli.DurationFlag{
		Name:  "quorum.handshake.timeout",
		Usage: "Timeout for quorum protocol to handshake with another peer",
		Value: quorum.DefaultHandshakeTimeout,
	}

	QuorumTestFlag = cli.BoolFlag{
		Name:  "quorum.test",
		Usage: "asdads",
	}

	// Metrics flags
	MetricsEnabledFlag = cli.BoolFlag{
		Name:  "metrics",
		Usage: "Enable metrics collection and reporting",
	}
	MetricsEnabledExpensiveFlag = cli.BoolFlag{
		Name:  "metrics.expensive",
		Usage: "Enable expensive metrics collection and reporting",
	}
	MetricsEnableInfluxDBFlag = cli.BoolFlag{
		Name:  "metrics.influxdb",
		Usage: "Enable metrics export/push to an external InfluxDB database",
	}
	MetricsInfluxDBEndpointFlag = cli.StringFlag{
		Name:  "metrics.influxdb.endpoint",
		Usage: "InfluxDB API endpoint to report metrics to",
		Value: "http://localhost:8086",
	}
	MetricsInfluxDBDatabaseFlag = cli.StringFlag{
		Name:  "metrics.influxdb.database",
		Usage: "InfluxDB database name to push reported metrics to",
		Value: "geth",
	}
	MetricsInfluxDBUsernameFlag = cli.StringFlag{
		Name:  "metrics.influxdb.username",
		Usage: "Username to authorize access to the database",
		Value: "test",
	}
	MetricsInfluxDBPasswordFlag = cli.StringFlag{
		Name:  "metrics.influxdb.password",
		Usage: "Password to authorize access to the database",
		Value: "test",
	}
	// Tags are part of every measurement sent to InfluxDB. Queries on tags are faster in InfluxDB.
	// For example `host` tag could be used so that we can group all nodes and average a measurement
	// across all of them, but also so that we can select a specific node and inspect its measurements.
	// https://docs.influxdata.com/influxdb/v1.4/concepts/key_concepts/#tag-key
	MetricsInfluxDBTagsFlag = cli.StringFlag{
		Name:  "metrics.influxdb.tags",
		Usage: "Comma-separated InfluxDB tags (key/values) attached to all measurements",
		Value: "host=localhost",
	}
	EWASMInterpreterFlag = cli.StringFlag{
		Name:  "vm.ewasm",
		Usage: "External ewasm configuration (default = built-in interpreter)",
		Value: "",
	}
	EVMInterpreterFlag = cli.StringFlag{
		Name:  "vm.evm",
		Usage: "External EVM configuration (default = built-in interpreter)",
		Value: "",
	}
)

// MakeDataDir retrieves the currently requested data directory, terminating
// if none (or the empty string) is specified. If the node is starting a testnet,
// then a subdirectory of the specified datadir will be used.
func MakeDataDir(ctx *cli.Context) string {
	if path := ctx.GlobalString(DataDirFlag.Name); path != "" {
		if ctx.GlobalBool(LegacyTestnetFlag.Name) || ctx.GlobalBool(RopstenFlag.Name) {
			// Maintain compatibility with older Geth configurations storing the
			// Ropsten database in `testnet` instead of `ropsten`.
			legacyPath := filepath.Join(path, "testnet")
			if _, err := os.Stat(legacyPath); !os.IsNotExist(err) {
				return legacyPath
			}
			return filepath.Join(path, "ropsten")
		}
		if ctx.GlobalBool(RinkebyFlag.Name) {
			return filepath.Join(path, "rinkeby")
		}
		if ctx.GlobalBool(GoerliFlag.Name) {
			return filepath.Join(path, "goerli")
		}
		if ctx.GlobalBool(YoloV1Flag.Name) {
			return filepath.Join(path, "yolo-v1")
		}
		return path
	}
	Fatalf("Cannot determine default data directory, please set manually (--datadir)")
	return ""
}

// setNodeKey creates a node key from set command line flags, either loading it
// from a file or as a specified hex value. If neither flags were provided, this
// method returns nil and an emphemeral key is to be generated.
func setNodeKey(ctx *cli.Context, cfg *p2p.Config) {
	var (
		hex  = ctx.GlobalString(NodeKeyHexFlag.Name)
		file = ctx.GlobalString(NodeKeyFileFlag.Name)
		key  *ecdsa.PrivateKey
		err  error
	)
	switch {
	case file != "" && hex != "":
		Fatalf("Options %q and %q are mutually exclusive", NodeKeyFileFlag.Name, NodeKeyHexFlag.Name)
	case file != "":
		if key, err = crypto.LoadECDSA(file); err != nil {
			Fatalf("Option %q: %v", NodeKeyFileFlag.Name, err)
		}
		cfg.PrivateKey = key
	case hex != "":
		if key, err = crypto.HexToECDSA(hex); err != nil {
			Fatalf("Option %q: %v", NodeKeyHexFlag.Name, err)
		}
		cfg.PrivateKey = key
	}
}

// setNodeUserIdent creates the user identifier from CLI flags.
func setNodeUserIdent(ctx *cli.Context, cfg *node.Config) {
	if identity := ctx.GlobalString(IdentityFlag.Name); len(identity) > 0 {
		cfg.UserIdent = identity
	}
}

// setBootstrapNodes creates a list of bootstrap nodes from the command line
// flags, reverting to pre-configured ones if none have been specified.
func setBootstrapNodes(ctx *cli.Context, cfg *p2p.Config) {
	//urls := params.MainnetBootnodes
	urls := params.BitbonBootnodes
	switch {
	case ctx.GlobalIsSet(BootnodesFlag.Name) || ctx.GlobalIsSet(LegacyBootnodesV4Flag.Name):
		if ctx.GlobalIsSet(LegacyBootnodesV4Flag.Name) {
			urls = splitAndTrim(ctx.GlobalString(LegacyBootnodesV4Flag.Name))
		} else {
			urls = splitAndTrim(ctx.GlobalString(BootnodesFlag.Name))
		}
	case ctx.GlobalBool(LegacyTestnetFlag.Name) || ctx.GlobalBool(RopstenFlag.Name):
		urls = params.RopstenBootnodes
	case ctx.GlobalBool(RinkebyFlag.Name):
		urls = params.RinkebyBootnodes
	case ctx.GlobalBool(GoerliFlag.Name):
		urls = params.GoerliBootnodes
	case ctx.GlobalBool(YoloV1Flag.Name):
		urls = params.YoloV1Bootnodes
	case cfg.BootstrapNodes != nil:
		return // already set, don't apply defaults.
	}

	cfg.BootstrapNodes = make([]*enode.Node, 0, len(urls))
	for _, url := range urls {
		if url != "" {
			n, err := enode.ParseV4(url)
			if err != nil {
				log.Crit("Bootstrap URL invalid", "enode", url, "err", err)
				continue
			}
			cfg.BootstrapNodes = append(cfg.BootstrapNodes, n)
		}
	}
}

// setBootstrapNodesV5 creates a list of bootstrap nodes from the command line
// flags, reverting to pre-configured ones if none have been specified.
func setBootstrapNodesV5(ctx *cli.Context, cfg *p2p.Config) {
	urls := params.MainnetBootnodes
	switch {
	case ctx.GlobalIsSet(BootnodesFlag.Name) || ctx.GlobalIsSet(LegacyBootnodesV5Flag.Name):
		if ctx.GlobalIsSet(LegacyBootnodesV5Flag.Name) {
			urls = splitAndTrim(ctx.GlobalString(LegacyBootnodesV5Flag.Name))
		} else {
			urls = splitAndTrim(ctx.GlobalString(BootnodesFlag.Name))
		}
	case ctx.GlobalBool(RopstenFlag.Name):
		urls = params.RopstenBootnodes
	case ctx.GlobalBool(RinkebyFlag.Name):
		urls = params.RinkebyBootnodes
	case ctx.GlobalBool(GoerliFlag.Name):
		urls = params.GoerliBootnodes
	case ctx.GlobalBool(YoloV1Flag.Name):
		urls = params.YoloV1Bootnodes
	case cfg.BootstrapNodesV5 != nil:
		return // already set, don't apply defaults.
	}

	cfg.BootstrapNodesV5 = make([]*discv5.Node, 0, len(urls))
	for _, url := range urls {
		if url != "" {
			n, err := discv5.ParseNode(url)
			if err != nil {
				log.Error("Bootstrap URL invalid", "enode", url, "err", err)
				continue
			}
			cfg.BootstrapNodesV5 = append(cfg.BootstrapNodesV5, n)
		}
	}
}

// setListenAddress creates a TCP listening address string from set command
// line flags.
func setListenAddress(ctx *cli.Context, cfg *p2p.Config) {
	if ctx.GlobalIsSet(ListenPortFlag.Name) {
		cfg.ListenAddr = fmt.Sprintf(":%d", ctx.GlobalInt(ListenPortFlag.Name))
	}
}

// setNAT creates a port mapper from command line flags.
func setNAT(ctx *cli.Context, cfg *p2p.Config) {
	if ctx.GlobalIsSet(NATFlag.Name) {
		natif, err := nat.Parse(ctx.GlobalString(NATFlag.Name))
		if err != nil {
			Fatalf("Option %s: %v", NATFlag.Name, err)
		}
		cfg.NAT = natif
	}
}

// splitAndTrim splits input separated by a comma
// and trims excessive white space from the substrings.
func splitAndTrim(input string) []string {
	result := strings.Split(input, ",")
	for i, r := range result {
		result[i] = strings.TrimSpace(r)
	}
	return result
}

// setHTTP creates the HTTP RPC listener interface string from the set
// command line flags, returning empty if the HTTP endpoint is disabled.
func setHTTP(ctx *cli.Context, cfg *node.Config) {
	if ctx.GlobalBool(LegacyRPCEnabledFlag.Name) && cfg.HTTPHost == "" {
		log.Warn("The flag --rpc is deprecated and will be removed in the future, please use --http")
		cfg.HTTPHost = "127.0.0.1"
		if ctx.GlobalIsSet(LegacyRPCListenAddrFlag.Name) {
			cfg.HTTPHost = ctx.GlobalString(LegacyRPCListenAddrFlag.Name)
			log.Warn("The flag --rpcaddr is deprecated and will be removed in the future, please use --http.addr")
		}
	}
	if ctx.GlobalBool(HTTPEnabledFlag.Name) && cfg.HTTPHost == "" {
		cfg.HTTPHost = "127.0.0.1"
		if ctx.GlobalIsSet(HTTPListenAddrFlag.Name) {
			cfg.HTTPHost = ctx.GlobalString(HTTPListenAddrFlag.Name)
		}
	}

	if ctx.GlobalIsSet(LegacyRPCPortFlag.Name) {
		cfg.HTTPPort = ctx.GlobalInt(LegacyRPCPortFlag.Name)
		log.Warn("The flag --rpcport is deprecated and will be removed in the future, please use --http.port")
	}
	if ctx.GlobalIsSet(HTTPPortFlag.Name) {
		cfg.HTTPPort = ctx.GlobalInt(HTTPPortFlag.Name)
	}

	if ctx.GlobalIsSet(LegacyRPCCORSDomainFlag.Name) {
		cfg.HTTPCors = splitAndTrim(ctx.GlobalString(LegacyRPCCORSDomainFlag.Name))
		log.Warn("The flag --rpccorsdomain is deprecated and will be removed in the future, please use --http.corsdomain")
	}
	if ctx.GlobalIsSet(HTTPCORSDomainFlag.Name) {
		cfg.HTTPCors = splitAndTrim(ctx.GlobalString(HTTPCORSDomainFlag.Name))
	}

	if ctx.GlobalIsSet(LegacyRPCApiFlag.Name) {
		cfg.HTTPModules = splitAndTrim(ctx.GlobalString(LegacyRPCApiFlag.Name))
		log.Warn("The flag --rpcapi is deprecated and will be removed in the future, please use --http.api")
	}
	if ctx.GlobalIsSet(HTTPApiFlag.Name) {
		cfg.HTTPModules = splitAndTrim(ctx.GlobalString(HTTPApiFlag.Name))
	}

	if ctx.GlobalIsSet(LegacyRPCVirtualHostsFlag.Name) {
		cfg.HTTPVirtualHosts = splitAndTrim(ctx.GlobalString(LegacyRPCVirtualHostsFlag.Name))
		log.Warn("The flag --rpcvhosts is deprecated and will be removed in the future, please use --http.vhosts")
	}
	if ctx.GlobalIsSet(HTTPVirtualHostsFlag.Name) {
		cfg.HTTPVirtualHosts = splitAndTrim(ctx.GlobalString(HTTPVirtualHostsFlag.Name))
	}
}

// setWS creates the WebSocket RPC listener interface string from the set
// command line flags, returning empty if the HTTP endpoint is disabled.
func setWS(ctx *cli.Context, cfg *node.Config) {
	if ctx.GlobalBool(WSEnabledFlag.Name) && cfg.WSHost == "" {
		cfg.WSHost = "127.0.0.1"
		if ctx.GlobalIsSet(LegacyWSListenAddrFlag.Name) {
			cfg.WSHost = ctx.GlobalString(LegacyWSListenAddrFlag.Name)
			log.Warn("The flag --wsaddr is deprecated and will be removed in the future, please use --ws.addr")
		}
		if ctx.GlobalIsSet(WSListenAddrFlag.Name) {
			cfg.WSHost = ctx.GlobalString(WSListenAddrFlag.Name)
		}
	}
	if ctx.GlobalIsSet(LegacyWSPortFlag.Name) {
		cfg.WSPort = ctx.GlobalInt(LegacyWSPortFlag.Name)
		log.Warn("The flag --wsport is deprecated and will be removed in the future, please use --ws.port")
	}
	if ctx.GlobalIsSet(WSPortFlag.Name) {
		cfg.WSPort = ctx.GlobalInt(WSPortFlag.Name)
	}

	if ctx.GlobalIsSet(LegacyWSAllowedOriginsFlag.Name) {
		cfg.WSOrigins = splitAndTrim(ctx.GlobalString(LegacyWSAllowedOriginsFlag.Name))
		log.Warn("The flag --wsorigins is deprecated and will be removed in the future, please use --ws.origins")
	}
	if ctx.GlobalIsSet(WSAllowedOriginsFlag.Name) {
		cfg.WSOrigins = splitAndTrim(ctx.GlobalString(WSAllowedOriginsFlag.Name))
	}

	if ctx.GlobalIsSet(LegacyWSApiFlag.Name) {
		cfg.WSModules = splitAndTrim(ctx.GlobalString(LegacyWSApiFlag.Name))
		log.Warn("The flag --wsapi is deprecated and will be removed in the future, please use --ws.api")
	}
	if ctx.GlobalIsSet(WSApiFlag.Name) {
		cfg.WSModules = splitAndTrim(ctx.GlobalString(WSApiFlag.Name))
	}
}

// setIPC creates an IPC path configuration from the set command line flags,
// returning an empty string if IPC was explicitly disabled, or the set path.
func setIPC(ctx *cli.Context, cfg *node.Config) {
	CheckExclusive(ctx, IPCDisabledFlag, IPCPathFlag)
	switch {
	case ctx.GlobalBool(IPCDisabledFlag.Name):
		cfg.IPCPath = ""
	case ctx.GlobalIsSet(IPCPathFlag.Name):
		cfg.IPCPath = ctx.GlobalString(IPCPathFlag.Name)
	}
}

// setLes configures the les server and ultra light client settings from the command line flags.
func setLes(ctx *cli.Context, cfg *eth.Config) {
	if ctx.GlobalIsSet(LegacyLightServFlag.Name) {
		cfg.LightServ = ctx.GlobalInt(LegacyLightServFlag.Name)
		log.Warn("The flag --lightserv is deprecated and will be removed in the future, please use --light.serve")
	}
	if ctx.GlobalIsSet(LightServeFlag.Name) {
		cfg.LightServ = ctx.GlobalInt(LightServeFlag.Name)
	}
	if ctx.GlobalIsSet(LightIngressFlag.Name) {
		cfg.LightIngress = ctx.GlobalInt(LightIngressFlag.Name)
	}
	if ctx.GlobalIsSet(LightEgressFlag.Name) {
		cfg.LightEgress = ctx.GlobalInt(LightEgressFlag.Name)
	}
	if ctx.GlobalIsSet(LegacyLightPeersFlag.Name) {
		cfg.LightPeers = ctx.GlobalInt(LegacyLightPeersFlag.Name)
		log.Warn("The flag --lightpeers is deprecated and will be removed in the future, please use --light.maxpeers")
	}
	if ctx.GlobalIsSet(LightMaxPeersFlag.Name) {
		cfg.LightPeers = ctx.GlobalInt(LightMaxPeersFlag.Name)
	}
	if ctx.GlobalIsSet(UltraLightServersFlag.Name) {
		cfg.UltraLightServers = strings.Split(ctx.GlobalString(UltraLightServersFlag.Name), ",")
	}
	if ctx.GlobalIsSet(UltraLightFractionFlag.Name) {
		cfg.UltraLightFraction = ctx.GlobalInt(UltraLightFractionFlag.Name)
	}
	if cfg.UltraLightFraction <= 0 && cfg.UltraLightFraction > 100 {
		log.Error("Ultra light fraction is invalid", "had", cfg.UltraLightFraction, "updated", eth.DefaultConfig.UltraLightFraction)
		cfg.UltraLightFraction = eth.DefaultConfig.UltraLightFraction
	}
	if ctx.GlobalIsSet(UltraLightOnlyAnnounceFlag.Name) {
		cfg.UltraLightOnlyAnnounce = ctx.GlobalBool(UltraLightOnlyAnnounceFlag.Name)
	}
}

// makeDatabaseHandles raises out the number of allowed file handles per process
// for Geth and returns half of the allowance to assign to the database.
func makeDatabaseHandles() int {
	limit, err := fdlimit.Maximum()
	if err != nil {
		Fatalf("Failed to retrieve file descriptor allowance: %v", err)
	}
	raised, err := fdlimit.Raise(uint64(limit))
	if err != nil {
		Fatalf("Failed to raise file descriptor allowance: %v", err)
	}
	return int(raised / 2) // Leave half for networking and other stuff
}

// MakeAddress converts an account specified directly as a hex encoded string or
// a key index in the key store to an internal account representation.
func MakeAddress(ks *keystore.KeyStore, account string) (accounts.Account, error) {
	// If the specified account is a valid address, return it
	if common.IsHexAddress(account) {
		return accounts.Account{Address: common.HexToAddress(account)}, nil
	}
	// Otherwise try to interpret the account as a keystore index
	index, err := strconv.Atoi(account)
	if err != nil || index < 0 {
		return accounts.Account{}, fmt.Errorf("invalid account address or index %q", account)
	}
	log.Warn("-------------------------------------------------------------------")
	log.Warn("Referring to accounts by order in the keystore folder is dangerous!")
	log.Warn("This functionality is deprecated and will be removed in the future!")
	log.Warn("Please use explicit addresses! (can search via `geth account list`)")
	log.Warn("-------------------------------------------------------------------")

	accs := ks.Accounts()
	if len(accs) <= index {
		return accounts.Account{}, fmt.Errorf("index %d higher than number of accounts %d", index, len(accs))
	}
	return accs[index], nil
}

// setEtherbase retrieves the etherbase either from the directly specified
// command line flags or from the keystore if CLI indexed.
func setEtherbase(ctx *cli.Context, ks *keystore.KeyStore, cfg *eth.Config) {
	// Extract the current etherbase, new flag overriding legacy one
	var etherbase string
	if ctx.GlobalIsSet(LegacyMinerEtherbaseFlag.Name) {
		etherbase = ctx.GlobalString(LegacyMinerEtherbaseFlag.Name)
		log.Warn("The flag --etherbase is deprecated and will be removed in the future, please use --miner.etherbase")
	}
	if ctx.GlobalIsSet(MinerEtherbaseFlag.Name) {
		etherbase = ctx.GlobalString(MinerEtherbaseFlag.Name)
	}
	// Convert the etherbase into an address and configure it
	if etherbase != "" {
		if ks != nil {
			account, err := MakeAddress(ks, etherbase)
			if err != nil {
				Fatalf("Invalid miner etherbase: %v", err)
			}
			cfg.Miner.Etherbase = account.Address
		} else {
			Fatalf("No etherbase configured")
		}
	}
}

// MakePasswordList reads password lines from the file specified by the global --password flag.
func MakePasswordList(ctx *cli.Context) []string {
	path := ctx.GlobalString(PasswordFileFlag.Name)
	if path == "" {
		return nil
	}
	text, err := ioutil.ReadFile(path)
	if err != nil {
		Fatalf("Failed to read password file: %v", err)
	}
	lines := strings.Split(string(text), "\n")
	// Sanitise DOS line endings.
	for i := range lines {
		lines[i] = strings.TrimRight(lines[i], "\r")
	}
	return lines
}

func SetP2PConfig(ctx *cli.Context, cfg *p2p.Config) { // nolint:gocyclo
	setNodeKey(ctx, cfg)
	setNAT(ctx, cfg)
	setListenAddress(ctx, cfg)
	setBootstrapNodes(ctx, cfg)
	setBootstrapNodesV5(ctx, cfg)

	lightClient := ctx.GlobalString(SyncModeFlag.Name) == "light"
	lightServer := (ctx.GlobalInt(LegacyLightServFlag.Name) != 0 || ctx.GlobalInt(LightServeFlag.Name) != 0)

	lightPeers := ctx.GlobalInt(LegacyLightPeersFlag.Name)
	if ctx.GlobalIsSet(LightMaxPeersFlag.Name) {
		lightPeers = ctx.GlobalInt(LightMaxPeersFlag.Name)
	}
	if lightClient && !ctx.GlobalIsSet(LegacyLightPeersFlag.Name) && !ctx.GlobalIsSet(LightMaxPeersFlag.Name) {
		// dynamic default - for clients we use 1/10th of the default for servers
		lightPeers /= 10
	}

	if ctx.GlobalIsSet(MaxPeersFlag.Name) {
		cfg.MaxPeers = ctx.GlobalInt(MaxPeersFlag.Name)
		if lightServer && !ctx.GlobalIsSet(LegacyLightPeersFlag.Name) && !ctx.GlobalIsSet(LightMaxPeersFlag.Name) {
			cfg.MaxPeers += lightPeers
		}
	} else {
		if lightServer {
			cfg.MaxPeers += lightPeers
		}
		if lightClient && (ctx.GlobalIsSet(LegacyLightPeersFlag.Name) || ctx.GlobalIsSet(LightMaxPeersFlag.Name)) && cfg.MaxPeers < lightPeers {
			cfg.MaxPeers = lightPeers
		}
	}
	if !(lightClient || lightServer) {
		lightPeers = 0
	}
	ethPeers := cfg.MaxPeers - lightPeers
	if lightClient {
		ethPeers = 0
	}
	log.Info("Maximum peer count", "ETH", ethPeers, "LES", lightPeers, "total", cfg.MaxPeers)

	if ctx.GlobalIsSet(MaxPendingPeersFlag.Name) {
		cfg.MaxPendingPeers = ctx.GlobalInt(MaxPendingPeersFlag.Name)
	}
	if ctx.GlobalIsSet(NoDiscoverFlag.Name) || lightClient {
		cfg.NoDiscovery = true
	}

	// if we're running a light client or server, force enable the v5 peer discovery
	// unless it is explicitly disabled with --nodiscover note that explicitly specifying
	// --v5disc overrides --nodiscover, in which case the later only disables v4 discovery
	forceV5Discovery := (lightClient || lightServer) && !ctx.GlobalBool(NoDiscoverFlag.Name)
	if ctx.GlobalIsSet(DiscoveryV5Flag.Name) {
		cfg.DiscoveryV5 = ctx.GlobalBool(DiscoveryV5Flag.Name)
	} else if forceV5Discovery {
		cfg.DiscoveryV5 = true
	}

	if netrestrict := ctx.GlobalString(NetrestrictFlag.Name); netrestrict != "" {
		list, err := netutil.ParseNetlist(netrestrict)
		if err != nil {
			Fatalf("Option %q: %v", NetrestrictFlag.Name, err)
		}
		cfg.NetRestrict = list
	}

	if ctx.GlobalBool(DeveloperFlag.Name) {
		// --dev mode can't use p2p networking.
		cfg.MaxPeers = 0
		cfg.ListenAddr = ":0"
		cfg.NoDiscovery = true
		cfg.DiscoveryV5 = false
	}
}

// SetNodeConfig applies node-related command line flags to the config.
func SetNodeConfig(ctx *cli.Context, cfg *node.Config) {
	SetP2PConfig(ctx, &cfg.P2P)
	setIPC(ctx, cfg)
	setHTTP(ctx, cfg)
	setWS(ctx, cfg)
	setNodeUserIdent(ctx, cfg)
	setDataDir(ctx, cfg)
	setSmartCard(ctx, cfg)

	if ctx.GlobalIsSet(ExternalSignerFlag.Name) {
		cfg.ExternalSigner = ctx.GlobalString(ExternalSignerFlag.Name)
	}

	if ctx.GlobalIsSet(KeyStoreDirFlag.Name) {
		cfg.KeyStoreDir = ctx.GlobalString(KeyStoreDirFlag.Name)
	}
	if ctx.GlobalIsSet(LightKDFFlag.Name) {
		cfg.UseLightweightKDF = ctx.GlobalBool(LightKDFFlag.Name)
	}
	if ctx.GlobalIsSet(NoUSBFlag.Name) {
		cfg.NoUSB = ctx.GlobalBool(NoUSBFlag.Name)
	}
	if ctx.GlobalIsSet(InsecureUnlockAllowedFlag.Name) {
		cfg.InsecureUnlockAllowed = ctx.GlobalBool(InsecureUnlockAllowedFlag.Name)
	}
}

func setSmartCard(ctx *cli.Context, cfg *node.Config) {
	// Skip enabling smartcards if no path is set
	path := ctx.GlobalString(SmartCardDaemonPathFlag.Name)
	if path == "" {
		return
	}
	// Sanity check that the smartcard path is valid
	fi, err := os.Stat(path)
	if err != nil {
		log.Info("Smartcard socket not found, disabling", "err", err)
		return
	}
	if fi.Mode()&os.ModeType != os.ModeSocket {
		log.Error("Invalid smartcard daemon path", "path", path, "type", fi.Mode().String())
		return
	}
	// Smartcard daemon path exists and is a socket, enable it
	cfg.SmartCardDaemonPath = path
}

func setDataDir(ctx *cli.Context, cfg *node.Config) {
	switch {
	case ctx.GlobalIsSet(DataDirFlag.Name):
		cfg.DataDir = ctx.GlobalString(DataDirFlag.Name)
	case ctx.GlobalBool(DeveloperFlag.Name):
		cfg.DataDir = "" // unless explicitly requested, use memory databases
	case (ctx.GlobalBool(LegacyTestnetFlag.Name) || ctx.GlobalBool(RopstenFlag.Name)) && cfg.DataDir == node.DefaultDataDir():
		// Maintain compatibility with older Geth configurations storing the
		// Ropsten database in `testnet` instead of `ropsten`.
		legacyPath := filepath.Join(node.DefaultDataDir(), "testnet")
		if _, err := os.Stat(legacyPath); !os.IsNotExist(err) {
			log.Warn("Using the deprecated `testnet` datadir. Future versions will store the Ropsten chain in `ropsten`.")
			cfg.DataDir = legacyPath
		} else {
			cfg.DataDir = filepath.Join(node.DefaultDataDir(), "ropsten")
		}
	case ctx.GlobalBool(RinkebyFlag.Name) && cfg.DataDir == node.DefaultDataDir():
		cfg.DataDir = filepath.Join(node.DefaultDataDir(), "rinkeby")
	case ctx.GlobalBool(GoerliFlag.Name) && cfg.DataDir == node.DefaultDataDir():
		cfg.DataDir = filepath.Join(node.DefaultDataDir(), "goerli")
	case ctx.GlobalBool(YoloV1Flag.Name) && cfg.DataDir == node.DefaultDataDir():
		cfg.DataDir = filepath.Join(node.DefaultDataDir(), "yolo-v1")
	}
}

func setGPO(ctx *cli.Context, cfg *gasprice.Config) {
	if ctx.GlobalIsSet(LegacyGpoBlocksFlag.Name) {
		cfg.Blocks = ctx.GlobalInt(LegacyGpoBlocksFlag.Name)
		log.Warn("The flag --gpoblocks is deprecated and will be removed in the future, please use --gpo.blocks")
	}
	if ctx.GlobalIsSet(GpoBlocksFlag.Name) {
		cfg.Blocks = ctx.GlobalInt(GpoBlocksFlag.Name)
	}

	if ctx.GlobalIsSet(LegacyGpoPercentileFlag.Name) {
		cfg.Percentile = ctx.GlobalInt(LegacyGpoPercentileFlag.Name)
		log.Warn("The flag --gpopercentile is deprecated and will be removed in the future, please use --gpo.percentile")
	}
	if ctx.GlobalIsSet(GpoPercentileFlag.Name) {
		cfg.Percentile = ctx.GlobalInt(GpoPercentileFlag.Name)
	}
}

func setTxPool(ctx *cli.Context, cfg *core.TxPoolConfig) {
	if ctx.GlobalIsSet(TxPoolLocalsFlag.Name) {
		locals := strings.Split(ctx.GlobalString(TxPoolLocalsFlag.Name), ",")
		for _, account := range locals {
			if trimmed := strings.TrimSpace(account); !common.IsHexAddress(trimmed) {
				Fatalf("Invalid account in --txpool.locals: %s", trimmed)
			} else {
				cfg.Locals = append(cfg.Locals, common.HexToAddress(account))
			}
		}
	}
	if ctx.GlobalIsSet(TxPoolNoLocalsFlag.Name) {
		cfg.NoLocals = ctx.GlobalBool(TxPoolNoLocalsFlag.Name)
	}
	if ctx.GlobalIsSet(TxPoolJournalFlag.Name) {
		cfg.Journal = ctx.GlobalString(TxPoolJournalFlag.Name)
	}
	if ctx.GlobalIsSet(TxPoolRejournalFlag.Name) {
		cfg.Rejournal = ctx.GlobalDuration(TxPoolRejournalFlag.Name)
	}
	if ctx.GlobalIsSet(TxPoolPriceLimitFlag.Name) {
		cfg.PriceLimit = ctx.GlobalUint64(TxPoolPriceLimitFlag.Name)
	}
	if ctx.GlobalIsSet(TxPoolPriceBumpFlag.Name) {
		cfg.PriceBump = ctx.GlobalUint64(TxPoolPriceBumpFlag.Name)
	}
	if ctx.GlobalIsSet(TxPoolAccountSlotsFlag.Name) {
		cfg.AccountSlots = ctx.GlobalUint64(TxPoolAccountSlotsFlag.Name)
	}
	if ctx.GlobalIsSet(TxPoolGlobalSlotsFlag.Name) {
		cfg.GlobalSlots = ctx.GlobalUint64(TxPoolGlobalSlotsFlag.Name)
	}
	if ctx.GlobalIsSet(TxPoolAccountQueueFlag.Name) {
		cfg.AccountQueue = ctx.GlobalUint64(TxPoolAccountQueueFlag.Name)
	}
	if ctx.GlobalIsSet(TxPoolGlobalQueueFlag.Name) {
		cfg.GlobalQueue = ctx.GlobalUint64(TxPoolGlobalQueueFlag.Name)
	}
	if ctx.GlobalIsSet(TxPoolLifetimeFlag.Name) {
		cfg.Lifetime = ctx.GlobalDuration(TxPoolLifetimeFlag.Name)
	}
}

func setEthash(ctx *cli.Context, cfg *eth.Config) {
	if ctx.GlobalIsSet(EthashCacheDirFlag.Name) {
		cfg.Ethash.CacheDir = ctx.GlobalString(EthashCacheDirFlag.Name)
	}
	if ctx.GlobalIsSet(EthashDatasetDirFlag.Name) {
		cfg.Ethash.DatasetDir = ctx.GlobalString(EthashDatasetDirFlag.Name)
	}
	if ctx.GlobalIsSet(EthashCachesInMemoryFlag.Name) {
		cfg.Ethash.CachesInMem = ctx.GlobalInt(EthashCachesInMemoryFlag.Name)
	}
	if ctx.GlobalIsSet(EthashCachesOnDiskFlag.Name) {
		cfg.Ethash.CachesOnDisk = ctx.GlobalInt(EthashCachesOnDiskFlag.Name)
	}
	if ctx.GlobalIsSet(EthashCachesLockMmapFlag.Name) {
		cfg.Ethash.CachesLockMmap = ctx.GlobalBool(EthashCachesLockMmapFlag.Name)
	}
	if ctx.GlobalIsSet(EthashDatasetsInMemoryFlag.Name) {
		cfg.Ethash.DatasetsInMem = ctx.GlobalInt(EthashDatasetsInMemoryFlag.Name)
	}
	if ctx.GlobalIsSet(EthashDatasetsOnDiskFlag.Name) {
		cfg.Ethash.DatasetsOnDisk = ctx.GlobalInt(EthashDatasetsOnDiskFlag.Name)
	}
	if ctx.GlobalIsSet(EthashDatasetsLockMmapFlag.Name) {
		cfg.Ethash.DatasetsLockMmap = ctx.GlobalBool(EthashDatasetsLockMmapFlag.Name)
	}
}

func setMiner(ctx *cli.Context, cfg *miner.Config) {
	if ctx.GlobalIsSet(MinerNotifyFlag.Name) {
		cfg.Notify = strings.Split(ctx.GlobalString(MinerNotifyFlag.Name), ",")
	}
	if ctx.GlobalIsSet(LegacyMinerExtraDataFlag.Name) {
		cfg.ExtraData = []byte(ctx.GlobalString(LegacyMinerExtraDataFlag.Name))
		log.Warn("The flag --extradata is deprecated and will be removed in the future, please use --miner.extradata")
	}
	if ctx.GlobalIsSet(MinerExtraDataFlag.Name) {
		cfg.ExtraData = []byte(ctx.GlobalString(MinerExtraDataFlag.Name))
	}
	if ctx.GlobalIsSet(LegacyMinerGasTargetFlag.Name) {
		cfg.GasFloor = ctx.GlobalUint64(LegacyMinerGasTargetFlag.Name)
		log.Warn("The flag --targetgaslimit is deprecated and will be removed in the future, please use --miner.gastarget")
	}
	if ctx.GlobalIsSet(MinerGasTargetFlag.Name) {
		cfg.GasFloor = ctx.GlobalUint64(MinerGasTargetFlag.Name)
	}
	if ctx.GlobalIsSet(MinerGasLimitFlag.Name) {
		cfg.GasCeil = ctx.GlobalUint64(MinerGasLimitFlag.Name)
	}
	if ctx.GlobalIsSet(LegacyMinerGasPriceFlag.Name) {
		cfg.GasPrice = GlobalBig(ctx, LegacyMinerGasPriceFlag.Name)
		log.Warn("The flag --gasprice is deprecated and will be removed in the future, please use --miner.gasprice")
	}
	if ctx.GlobalIsSet(MinerGasPriceFlag.Name) {
		cfg.GasPrice = GlobalBig(ctx, MinerGasPriceFlag.Name)
	}
	if ctx.GlobalIsSet(MinerRecommitIntervalFlag.Name) {
		cfg.Recommit = ctx.Duration(MinerRecommitIntervalFlag.Name)
	}
	if ctx.GlobalIsSet(MinerNoVerfiyFlag.Name) {
		cfg.Noverify = ctx.Bool(MinerNoVerfiyFlag.Name)
	}
}

func setWhitelist(ctx *cli.Context, cfg *eth.Config) {
	whitelist := ctx.GlobalString(WhitelistFlag.Name)
	if whitelist == "" {
		return
	}
	cfg.Whitelist = make(map[uint64]common.Hash)
	for _, entry := range strings.Split(whitelist, ",") {
		parts := strings.Split(entry, "=")
		if len(parts) != 2 {
			Fatalf("Invalid whitelist entry: %s", entry)
		}
		number, err := strconv.ParseUint(parts[0], 0, 64)
		if err != nil {
			Fatalf("Invalid whitelist block number %s: %v", parts[0], err)
		}
		var hash common.Hash
		if err = hash.UnmarshalText([]byte(parts[1])); err != nil {
			Fatalf("Invalid whitelist hash %s: %v", parts[1], err)
		}
		cfg.Whitelist[number] = hash
	}
}

// CheckExclusive verifies that only a single instance of the provided flags was
// set by the user. Each flag might optionally be followed by a string type to
// specialize it further.
func CheckExclusive(ctx *cli.Context, args ...interface{}) {
	set := make([]string, 0, 1)
	for i := 0; i < len(args); i++ {
		// Make sure the next argument is a flag and skip if not set
		flag, ok := args[i].(cli.Flag)
		if !ok {
			panic(fmt.Sprintf("invalid argument, not cli.Flag type: %T", args[i]))
		}
		// Check if next arg extends current and expand its name if so
		name := flag.GetName()

		if i+1 < len(args) {
			switch option := args[i+1].(type) {
			case string:
				// Extended flag check, make sure value set doesn't conflict with passed in option
				if ctx.GlobalString(flag.GetName()) == option {
					name += "=" + option
					set = append(set, "--"+name)
				}
				// shift arguments and continue
				i++
				continue

			case cli.Flag:
			default:
				panic(fmt.Sprintf("invalid argument, not cli.Flag or string extension: %T", args[i+1]))
			}
		}
		// Mark the flag if it's set
		if ctx.GlobalIsSet(flag.GetName()) {
			set = append(set, "--"+name)
		}
	}
	if len(set) > 1 {
		Fatalf("Flags %v can't be used at the same time", strings.Join(set, ", "))
	}
}

// SetEthConfig applies eth-related command line flags to the config.
func SetEthConfig(ctx *cli.Context, stack *node.Node, cfg *eth.Config) {
	// Avoid conflicting network flags
	CheckExclusive(ctx, DeveloperFlag, LegacyTestnetFlag, RopstenFlag, RinkebyFlag, GoerliFlag, YoloV1Flag)
	CheckExclusive(ctx, LegacyLightServFlag, LightServeFlag, SyncModeFlag, "light")
	CheckExclusive(ctx, DeveloperFlag, ExternalSignerFlag) // Can't use both ephemeral unlocked and external signer
	CheckExclusive(ctx, GCModeFlag, "archive", TxLookupLimitFlag)
	// Ancient tx indices pruning is not available for les server now
	// since light client relies on the server for transaction status query.
	CheckExclusive(ctx, LegacyLightServFlag, LightServeFlag, TxLookupLimitFlag)
	var ks *keystore.KeyStore
	if keystores := stack.AccountManager().Backends(keystore.KeyStoreType); len(keystores) > 0 {
		ks = keystores[0].(*keystore.KeyStore)
	}
	setEtherbase(ctx, ks, cfg)
	setGPO(ctx, &cfg.GPO)
	setTxPool(ctx, &cfg.TxPool)
	setEthash(ctx, cfg)
	setMiner(ctx, &cfg.Miner)
	setWhitelist(ctx, cfg)
	setLes(ctx, cfg)

	if ctx.GlobalIsSet(SyncModeFlag.Name) {
		cfg.SyncMode = *GlobalTextMarshaler(ctx, SyncModeFlag.Name).(*downloader.SyncMode)
	}
	if ctx.GlobalIsSet(NetworkIdFlag.Name) {
		cfg.NetworkId = ctx.GlobalUint64(NetworkIdFlag.Name)
	}
	if ctx.GlobalIsSet(CacheFlag.Name) || ctx.GlobalIsSet(CacheDatabaseFlag.Name) {
		cfg.DatabaseCache = ctx.GlobalInt(CacheFlag.Name) * ctx.GlobalInt(CacheDatabaseFlag.Name) / 100
	}
	cfg.DatabaseHandles = makeDatabaseHandles()
	if ctx.GlobalIsSet(AncientFlag.Name) {
		cfg.DatabaseFreezer = ctx.GlobalString(AncientFlag.Name)
	}

	if gcmode := ctx.GlobalString(GCModeFlag.Name); gcmode != "full" && gcmode != "archive" {
		Fatalf("--%s must be either 'full' or 'archive'", GCModeFlag.Name)
	}
	if ctx.GlobalIsSet(GCModeFlag.Name) {
		cfg.NoPruning = ctx.GlobalString(GCModeFlag.Name) == "archive"
	}
	if ctx.GlobalIsSet(CacheNoPrefetchFlag.Name) {
		cfg.NoPrefetch = ctx.GlobalBool(CacheNoPrefetchFlag.Name)
	}
	if ctx.GlobalIsSet(TxLookupLimitFlag.Name) {
		cfg.TxLookupLimit = ctx.GlobalUint64(TxLookupLimitFlag.Name)
	}
	if ctx.GlobalIsSet(CacheFlag.Name) || ctx.GlobalIsSet(CacheTrieFlag.Name) {
		cfg.TrieCleanCache = ctx.GlobalInt(CacheFlag.Name) * ctx.GlobalInt(CacheTrieFlag.Name) / 100
	}
	if ctx.GlobalIsSet(CacheFlag.Name) || ctx.GlobalIsSet(CacheGCFlag.Name) {
		cfg.TrieDirtyCache = ctx.GlobalInt(CacheFlag.Name) * ctx.GlobalInt(CacheGCFlag.Name) / 100
	}
	if ctx.GlobalIsSet(CacheFlag.Name) || ctx.GlobalIsSet(CacheSnapshotFlag.Name) {
		cfg.SnapshotCache = ctx.GlobalInt(CacheFlag.Name) * ctx.GlobalInt(CacheSnapshotFlag.Name) / 100
	}
	if !ctx.GlobalIsSet(SnapshotFlag.Name) {
		cfg.SnapshotCache = 0 // Disabled
	}
	if ctx.GlobalIsSet(DocRootFlag.Name) {
		cfg.DocRoot = ctx.GlobalString(DocRootFlag.Name)
	}
	if ctx.GlobalIsSet(VMEnableDebugFlag.Name) {
		cfg.EnablePreimageRecording = ctx.GlobalBool(VMEnableDebugFlag.Name)
	}

	if ctx.GlobalIsSet(EWASMInterpreterFlag.Name) {
		cfg.EWASMInterpreter = ctx.GlobalString(EWASMInterpreterFlag.Name)
	}

	if ctx.GlobalIsSet(EVMInterpreterFlag.Name) {
		cfg.EVMInterpreter = ctx.GlobalString(EVMInterpreterFlag.Name)
	}
	if ctx.GlobalIsSet(RPCGlobalGasCap.Name) {
		cfg.RPCGasCap = new(big.Int).SetUint64(ctx.GlobalUint64(RPCGlobalGasCap.Name))
	}
	if ctx.GlobalIsSet(DNSDiscoveryFlag.Name) {
		urls := ctx.GlobalString(DNSDiscoveryFlag.Name)
		if urls == "" {
			cfg.DiscoveryURLs = []string{}
		} else {
			cfg.DiscoveryURLs = splitAndTrim(urls)
		}
	}

	// Override any default configs for hard coded networks.
	switch {
	case ctx.GlobalBool(LegacyTestnetFlag.Name) || ctx.GlobalBool(RopstenFlag.Name):
		if !ctx.GlobalIsSet(NetworkIdFlag.Name) {
			cfg.NetworkId = 3
		}
		cfg.Genesis = core.DefaultRopstenGenesisBlock()
		setDNSDiscoveryDefaults(cfg, params.RopstenGenesisHash)
	case ctx.GlobalBool(RinkebyFlag.Name):
		if !ctx.GlobalIsSet(NetworkIdFlag.Name) {
			cfg.NetworkId = 4
		}
		cfg.Genesis = core.DefaultRinkebyGenesisBlock()
		setDNSDiscoveryDefaults(cfg, params.RinkebyGenesisHash)
	case ctx.GlobalBool(GoerliFlag.Name):
		if !ctx.GlobalIsSet(NetworkIdFlag.Name) {
			cfg.NetworkId = 5
		}
		cfg.Genesis = core.DefaultGoerliGenesisBlock()
		setDNSDiscoveryDefaults(cfg, params.GoerliGenesisHash)
	case ctx.GlobalBool(YoloV1Flag.Name):
		if !ctx.GlobalIsSet(NetworkIdFlag.Name) {
			cfg.NetworkId = 133519467574833 // "yolov1"
		}
		cfg.Genesis = core.DefaultYoloV1GenesisBlock()
	case ctx.GlobalBool(DeveloperFlag.Name):
		if !ctx.GlobalIsSet(NetworkIdFlag.Name) {
			cfg.NetworkId = 1337
		}
		// Create new developer account or reuse existing one
		var (
			developer accounts.Account
			err       error
		)
		if accs := ks.Accounts(); len(accs) > 0 {
			developer = ks.Accounts()[0]
		} else {
			developer, err = ks.NewAccount("")
			if err != nil {
				Fatalf("Failed to create developer account: %v", err)
			}
		}
		if err := ks.Unlock(developer, ""); err != nil {
			Fatalf("Failed to unlock developer account: %v", err)
		}
		log.Info("Using developer account", "address", developer.Address)

		cfg.Genesis = core.DeveloperGenesisBlock(uint64(ctx.GlobalInt(DeveloperPeriodFlag.Name)), developer.Address)
		if !ctx.GlobalIsSet(MinerGasPriceFlag.Name) && !ctx.GlobalIsSet(LegacyMinerGasPriceFlag.Name) {
			cfg.Miner.GasPrice = big.NewInt(1)
		}
	default:
		if cfg.NetworkId == 1 {
			setDNSDiscoveryDefaults(cfg, params.MainnetGenesisHash)
		}
	}
}

// setDNSDiscoveryDefaults configures DNS discovery with the given URL if
// no URLs are set.
func setDNSDiscoveryDefaults(cfg *eth.Config, genesis common.Hash) {
	if cfg.DiscoveryURLs != nil {
		return // already set through flags/config
	}

	protocol := "all"
	if cfg.SyncMode == downloader.LightSync {
		protocol = "les"
	}
	if url := params.KnownDNSNetwork(genesis, protocol); url != "" {
		cfg.DiscoveryURLs = []string{url}
	}
}

// SetBitbonConfig applies bitbon related command line flags to the config.
func SetBitbonConfig(ctx *cli.Context, cfg *bitbon.Config) {
	cfg.ServiceID = ctx.GlobalString(BitbonServiceID.Name)
	cfg.ServicePrivateKey = ctx.GlobalString(BitbonServicePrivateKey.Name)
	cfg.DecryptAssetboxWalletPassword = ctx.GlobalString(BitbonDecryptAssetboxWalletPassword.Name)
}

// SetBitbonClientConfig applies bitbon client related command line flags to the config.
func SetBitbonClientConfig(ctx *cli.Context, cfg *client.Config) {
	cfg.Enabled = ctx.GlobalBool(BitbonClientFlag.Name)
	cfg.Mode = client.BitbonClientMode(ctx.GlobalInt(BitbonClientModeFlag.Name))
	cfg.WaitSync = ctx.GlobalBool(BitbonClientWaitSyncFlag.Name)

	cfg.AmqpConnStrs = strings.Split(ctx.GlobalString(BitbonClientAmqpConnectionFlag.Name), ",")
	cfg.AmqpAssetboxExchange = ctx.GlobalString(BitbonClientAmqpAssetboxExchange.Name)
	cfg.AmqpTransferExchange = ctx.GlobalString(BitbonClientAmqpTransferExchange.Name)
	cfg.AmqpBlockExchange = ctx.GlobalString(BitbonClientAmqpBlockExchange.Name)
	cfg.AmqpTransactionExchange = ctx.GlobalString(BitbonClientAmqpTransactionExchange.Name)
	cfg.AmqpMiningExchange = ctx.GlobalString(BitbonClientAmqpMiningExchange.Name)
	cfg.AmqpConfirmTimeout = ctx.GlobalDuration(BitbonClientAmqpConfirmTimeout.Name)
	cfg.AmqpPublishTimeout = ctx.GlobalDuration(BitbonClientAmqpPublishTimeout.Name)
	cfg.JournalTimeout = ctx.GlobalDuration(BitbonClientJournalTimeout.Name)
	cfg.JournalName = ctx.GlobalString(BitbonClientJournalName.Name)
}

// SetNoncerConfig
func SetNoncerConfig(ctx *cli.Context, cfg *noncer.Config) {
	cfg.Enabled = ctx.GlobalBool(BitbonNoncerFlag.Name)
	cfg.RedisConnStrs = strings.Split(ctx.GlobalString(BitbonNoncerRedis.Name), ",")
	cfg.Key = ctx.GlobalString(BitbonNoncerKey.Name)
	cfg.DialTimeout = ctx.GlobalUint64(BitbonNoncerRedisDialTimeout.Name)
	cfg.MaxRetries = ctx.GlobalUint64(BitbonNoncerRedisMaxRetries.Name)
}

// SetBitbonAgentConfig applies bitbon agent related command line flags to the config.
func SetBitbonAgentConfig(ctx *cli.Context, cfg *agent.Config) {
	cfg.Enabled = ctx.GlobalBool(BitbonAgentFlag.Name)
	cfg.AmqpConnStrs = strings.Split(ctx.GlobalString(BitbonAgentAmqpConnectionFlag.Name), ",")
	cfg.AmqpExchange = ctx.GlobalString(BitbonAgentAmqpExchangeFlag.Name)
	cfg.JournalName = ctx.GlobalString(BitbonAgentJournalFlag.Name)
	cfg.JournalTimeout = ctx.GlobalDuration(BitbonAgentJournalTimeout.Name)
}

// SetQuorumConfig applies quorum related command line flags to the config.
func SetQuorumConfig(ctx *cli.Context, cfg *quorum.Config) {
	cfg.PrepareMinersTimeout = ctx.GlobalDuration(QuorumPrepareMinersTimeoutFlag.Name)
	cfg.PingFrequency = ctx.GlobalDuration(QuorumPingFrequencyFlag.Name)
	cfg.QuorumMaxPing = ctx.GlobalDuration(QuorumMaxPingFlag.Name)
	cfg.HandshakeTimeout = ctx.GlobalDuration(QuorumHandshakeTimeoutFlag.Name)

	cfg.Test = ctx.GlobalBool(QuorumTestFlag.Name)
}

// RegisterEthService adds an Ethereum client to the stack.
func RegisterEthService(stack *node.Node, cfg *eth.Config) {
	var err error
	if cfg.SyncMode == downloader.LightSync {
		err = stack.Register(func(ctx *node.ServiceContext) (node.Service, error) {
			return les.New(ctx, cfg)
		})
	} else {
		err = stack.Register(func(ctx *node.ServiceContext) (node.Service, error) {
			var fullNode *eth.Ethereum
			fullNode, err = eth.New(ctx, cfg)
			if fullNode != nil && cfg.LightServ > 0 {
				ls, _ := les.NewLesServer(fullNode, cfg)
				fullNode.AddLesServer(ls)
			}
			return fullNode, err
		})
	}
	if err != nil {
		Fatalf("Failed to register the Ethereum service: %v", err)
	}
}

func SetupMetrics(ctx *cli.Context) {
	if metrics.Enabled {
		log.Info("Enabling metrics collection")
		var (
			enableExport = ctx.GlobalBool(MetricsEnableInfluxDBFlag.Name)
			endpoint     = ctx.GlobalString(MetricsInfluxDBEndpointFlag.Name)
			database     = ctx.GlobalString(MetricsInfluxDBDatabaseFlag.Name)
			username     = ctx.GlobalString(MetricsInfluxDBUsernameFlag.Name)
			password     = ctx.GlobalString(MetricsInfluxDBPasswordFlag.Name)
		)

		if enableExport {
			tagsMap := SplitTagsFlag(ctx.GlobalString(MetricsInfluxDBTagsFlag.Name))

			log.Info("Enabling metrics export to InfluxDB")

			go influxdb.InfluxDBWithTags(metrics.DefaultRegistry, 10*time.Second, endpoint, database, username, password, "geth.", tagsMap)
		}
	}
}

func SplitTagsFlag(tagsFlag string) map[string]string {
	tags := strings.Split(tagsFlag, ",")
	tagsMap := map[string]string{}

	for _, t := range tags {
		if t != "" {
			kv := strings.Split(t, "=")

			if len(kv) == 2 {
				tagsMap[kv[0]] = kv[1]
			}
		}
	}

	return tagsMap
}

// RegisterBitbonService adds a bitbon to the stack.
func RegisterBitbonService(stack *node.Node, cfg *bitbon.Config) {
	if err := stack.Register(func(ctx *node.ServiceContext) (node.Service, error) {
		b, err := bitbon.New(ctx, cfg)
		if err != nil {
			return nil, err
		}

		tm, err := transfer.NewTransferManager(b, aes.New())
		if err != nil {
			err = errors.Wrap(err, "error while creating transfer manager")
			return nil, err
		}

		var bitbonNoncer *noncer.Noncer
		err = ctx.Service(&bitbonNoncer)
		if err != nil {
			log.Warn("Error getting noncer service from context. Continue with default behavior", "error", err)
		}

		contractManager, err := initContractManager(b.GetEth().APIBackend, b.GetStorage().GetSettingContractStorageAddress,
			b.GetStorage().SaveSettingContractStorageAddress, b.GetServicePk(), bitbonNoncer)
		if err != nil {
			err = errors.Wrap(err, "error while creating contract manager")
			return nil, err
		}

		b.InjectTransferManager(tm)
		b.InjectAssetboxManager(assetbox.NewAssetboxManager(b, aes.New()))
		b.InjectMiningAgent(mining_agent.NewMiningAgentManager(b, aes.New()))
		b.InjectContractManager(contractManager)
		b.InitBitbonParser(b.GetContractsManager().GetEthAPIWrapper(), b.GetStorage())

		var q *quorum.Quorum
		err = ctx.Service(&q)
		if err != nil {
			log.Error("Failed to get quorum service for setting GetCurrentDistribution", "error", err)
			return b, nil
		}

		c := quorum.ContractServiceWith(func(c context.Context) (distribution map[common.Address]uint64, err error) {
			distribution = make(map[common.Address]uint64)
			tempRes, err := b.GetContractsManager().GetCurrentDistribution(c)
			if err != nil {
				return nil, errors.Wrap(err, "unable to get current distribution")
			}

			for k, v := range tempRes {
				distribution[common.HexToAddress(k)] = v
			}

			return distribution, nil
		})

		q.SetContractService(c)

		return b, nil
	}); err != nil {
		Fatalf("Failed to register the Bitbon service: %v", err)
	}
}

func initContractManager(apiBackend *eth.EthAPIBackend,
	getSettingContractStorageAddress func() (common.Address, error),
	saveSettingContractStorageAddress func(addr common.Address) error,
	servicePK *ecdsa.PrivateKey,
	bitbonNoncer *noncer.Noncer) (*contracts.Manager, error) {
	manager, err := contracts.NewManager(apiBackend, getSettingContractStorageAddress,
		saveSettingContractStorageAddress, servicePK, bitbonNoncer)
	if err != nil {
		err = errors.Wrap(err, "error while creating Bitbon contracts manager")
		return nil, err
	}

	defer func() {
		if err != nil {
			manager.Stop() // nolint:errcheck
		}
	}()

	return manager, nil
}

// RegisterBitbonAgentService adds a bitbon agent to the stack.
func RegisterBitbonAgentService(stack *node.Node, cfg *agent.Config) {
	if err := stack.Register(func(ctx *node.ServiceContext) (node.Service, error) {
		return agent.New(ctx, cfg)
	}); err != nil {
		Fatalf("Failed to register the Bitbon agent service: %v", err)
	}
}

// RegisterBitbonClientService adds a bitbon client to the stack.
func RegisterBitbonClientService(stack *node.Node, cfg *client.Config) {
	if err := stack.Register(func(ctx *node.ServiceContext) (node.Service, error) {
		return client.New(ctx, cfg)
	}); err != nil {
		Fatalf("Failed to register the Bitbon client service: %v", err)
	}
}

// RegisterBitbonNoncerService
func RegisterBitbonNoncerService(stack *node.Node, cfg *noncer.Config) {
	if err := stack.Register(func(ctx *node.ServiceContext) (node.Service, error) {
		return noncer.New(ctx, cfg)
	}); err != nil {
		Fatalf("Failed to register the Bitbon noncer service: %v", err)
	}

}

// RegisterQuorumService adds a quorum to the stack.
func RegisterQuorumService(stack *node.Node, cfg *quorum.Config) {
	if err := stack.Register(func(ctx *node.ServiceContext) (node.Service, error) {
		return quorum.New(ctx, cfg)
	}); err != nil {
		Fatalf("Failed to register the Quorum service: %v", err)
	}
}

// MakeChainDatabase open an LevelDB using the flags passed to the client and will hard crash if it fails.
func MakeChainDatabase(ctx *cli.Context, stack *node.Node) ethdb.Database {
	var (
		cache   = ctx.GlobalInt(CacheFlag.Name) * ctx.GlobalInt(CacheDatabaseFlag.Name) / 100
		handles = makeDatabaseHandles()
	)
	name := "chaindata"
	if ctx.GlobalString(SyncModeFlag.Name) == "light" {
		name = "lightchaindata"
	}
	chainDb, err := stack.OpenDatabaseWithFreezer(name, cache, handles, ctx.GlobalString(AncientFlag.Name), "")
	if err != nil {
		Fatalf("Could not open database: %v", err)
	}
	return chainDb
}

func MakeGenesis(ctx *cli.Context) *core.Genesis {
	var genesis *core.Genesis
	switch {
	case ctx.GlobalBool(LegacyTestnetFlag.Name) || ctx.GlobalBool(RopstenFlag.Name):
		genesis = core.DefaultRopstenGenesisBlock()
	case ctx.GlobalBool(RinkebyFlag.Name):
		genesis = core.DefaultRinkebyGenesisBlock()
	case ctx.GlobalBool(GoerliFlag.Name):
		genesis = core.DefaultGoerliGenesisBlock()
	case ctx.GlobalBool(YoloV1Flag.Name):
		genesis = core.DefaultYoloV1GenesisBlock()
	case ctx.GlobalBool(DeveloperFlag.Name):
		Fatalf("Developer chains are ephemeral")
	}
	return genesis
}

// MakeChain creates a chain manager from set command line flags.
func MakeChain(ctx *cli.Context, stack *node.Node, readOnly bool) (chain *core.BlockChain, chainDb ethdb.Database) {
	var err error
	chainDb = MakeChainDatabase(ctx, stack)
	config, _, err := core.SetupGenesisBlock(chainDb, MakeGenesis(ctx))
	if err != nil {
		Fatalf("%v", err)
	}
	createEthashEngine := func() *ethash.Ethash {
		if !ctx.GlobalBool(FakePoWFlag.Name) {
			return ethash.New(ethash.Config{
				CacheDir:         stack.ResolvePath(eth.DefaultConfig.Ethash.CacheDir),
				CachesInMem:      eth.DefaultConfig.Ethash.CachesInMem,
				CachesOnDisk:     eth.DefaultConfig.Ethash.CachesOnDisk,
				CachesLockMmap:   eth.DefaultConfig.Ethash.CachesLockMmap,
				DatasetDir:       stack.ResolvePath(eth.DefaultConfig.Ethash.DatasetDir),
				DatasetsInMem:    eth.DefaultConfig.Ethash.DatasetsInMem,
				DatasetsOnDisk:   eth.DefaultConfig.Ethash.DatasetsOnDisk,
				DatasetsLockMmap: eth.DefaultConfig.Ethash.DatasetsLockMmap,
			}, nil, false, config.MaxGenDAGBlock)
		} else {
			return ethash.NewFaker()
		}
	}

	var engine consensus.Engine
	if config.Quorum != nil {
		var q *quorum.Quorum
		q, err = quorum.New(nil, &quorum.DefaultConfig)
		if err != nil {
			log.Crit("Failed to initialize Quorum consensus engine", "error", err)
		}

		cposBlock := config.CposBlock
		if cposBlock != nil && cposBlock.Cmp(big.NewInt(0)) == 0 {
			cposBlock = nil
		}

		q.SetEngine(quorumEngine.New(config.Quorum, chainDb, q, cposBlock))
		engine = q.Engine()
	} else if config.Clique != nil {
		engine = clique.New(config.Clique, chainDb)
	} else {
		engine = createEthashEngine()
	}
	if gcmode := ctx.GlobalString(GCModeFlag.Name); gcmode != "full" && gcmode != "archive" {
		Fatalf("--%s must be either 'full' or 'archive'", GCModeFlag.Name)
	}
	cache := &core.CacheConfig{
		TrieCleanLimit:      eth.DefaultConfig.TrieCleanCache,
		TrieCleanNoPrefetch: ctx.GlobalBool(CacheNoPrefetchFlag.Name),
		TrieDirtyLimit:      eth.DefaultConfig.TrieDirtyCache,
		TrieDirtyDisabled:   ctx.GlobalString(GCModeFlag.Name) == "archive",
		TrieTimeLimit:       eth.DefaultConfig.TrieTimeout,
		SnapshotLimit:       eth.DefaultConfig.SnapshotCache,
	}
	if !ctx.GlobalIsSet(SnapshotFlag.Name) {
		cache.SnapshotLimit = 0 // Disabled
	}
	if ctx.GlobalIsSet(CacheFlag.Name) || ctx.GlobalIsSet(CacheTrieFlag.Name) {
		cache.TrieCleanLimit = ctx.GlobalInt(CacheFlag.Name) * ctx.GlobalInt(CacheTrieFlag.Name) / 100
	}
	if ctx.GlobalIsSet(CacheFlag.Name) || ctx.GlobalIsSet(CacheGCFlag.Name) {
		cache.TrieDirtyLimit = ctx.GlobalInt(CacheFlag.Name) * ctx.GlobalInt(CacheGCFlag.Name) / 100
	}
	vmcfg := vm.Config{EnablePreimageRecording: ctx.GlobalBool(VMEnableDebugFlag.Name)}
	var limit *uint64
	if ctx.GlobalIsSet(TxLookupLimitFlag.Name) && !readOnly {
		l := ctx.GlobalUint64(TxLookupLimitFlag.Name)
		limit = &l
	}
	chain, err = core.NewBlockChain(chainDb, cache, config, engine, vmcfg, nil, limit)
	if err != nil {
		Fatalf("Can't create BlockChain: %v", err)
	}
	return chain, chainDb
}

// MakeConsolePreloads retrieves the absolute paths for the console JavaScript
// scripts to preload before starting.
func MakeConsolePreloads(ctx *cli.Context) []string {
	// Skip preloading if there's nothing to preload
	if ctx.GlobalString(PreloadJSFlag.Name) == "" {
		return nil
	}
	// Otherwise resolve absolute paths and return them
	preloads := make([]string, 0, 100)

	assets := ctx.GlobalString(JSpathFlag.Name)
	for _, file := range strings.Split(ctx.GlobalString(PreloadJSFlag.Name), ",") {
		preloads = append(preloads, common.AbsolutePath(assets, strings.TrimSpace(file)))
	}
	return preloads
}

// MigrateFlags sets the global flag from a local flag when it's set.
// This is a temporary function used for migrating old command/flags to the
// new format.
//
// e.g. geth account new --keystore /tmp/mykeystore --lightkdf
//
// is equivalent after calling this method with:
//
// geth --keystore /tmp/mykeystore --lightkdf account new
//
// This allows the use of the existing configuration functionality.
// When all flags are migrated this function can be removed and the existing
// configuration functionality must be changed that is uses local flags
func MigrateFlags(action func(ctx *cli.Context) error) func(*cli.Context) error {
	return func(ctx *cli.Context) error {
		for _, name := range ctx.FlagNames() {
			if ctx.IsSet(name) {
				if err := ctx.GlobalSet(name, ctx.String(name)); err != nil {
					log.Error("GlobalSet error", "error", err)
				}
			}
		}
		return action(ctx)
	}
}
