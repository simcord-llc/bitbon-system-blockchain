## Bitbon System Blockchain

Official implementation of the Bitbon System Blockchain based on [Ethereum codebase](https://github.com/ethereum/go-ethereum/).

Owner: **[Simcord-LLC](https://www.simcord.com/)**


## Building the source

Building `bitbon_node` binary requires both a Go (version 1.14 or later) and a C compiler. You can install
them using your favourite package manager. 

Also you need to install some libs.
Ubuntu example:
```
sudo apt install make g++ gcc musl-dev gmpc-dev libmpfr-dev
```

Once the dependencies are installed, run

```shell
make clean build
```

## Running `bitbon_node`

Bitbon node runs with command line parameters which Ethereum uses. To see Ethereum parameters - ethereum see cli wiki [CLI Wiki page](https://geth.ethereum.org/docs/interface/command-line-options)

### Programmatically interfacing `bitbon_node` nodes

HTTP based JSON-RPC API options:

* `--http` Enable the HTTP-RPC server
* `--http.addr` HTTP-RPC server listening interface (default: `localhost`)
* `--http.port` HTTP-RPC server listening port (default: `8545`)
* `--http.api` API's offered over the HTTP-RPC interface (default: `eth,net,web3`)
* `--http.corsdomain` Comma separated list of domains from which to accept cross origin requests (browser enforced)
* `--ws` Enable the WS-RPC server
* `--ws.addr` WS-RPC server listening interface (default: `localhost`)
* `--ws.port` WS-RPC server listening port (default: `8546`)
* `--ws.api` API's offered over the WS-RPC interface (default: `eth,net,web3`)
* `--ws.origins` Origins from which to accept websockets requests
* `--ipcdisable` Disable the IPC-RPC server
* `--ipcapi` API's offered over the IPC-RPC interface (default: `admin,debug,eth,miner,net,personal,shh,txpool,web3`)
* `--ipcpath` Filename for IPC socket/pipe within the datadir (explicit paths escape it)

You'll need to use your own programming environments' capabilities (libraries, tools, etc) to
connect via HTTP, WS or IPC to a `bitbon_node` node configured with the above flags and you'll
need to speak [JSON-RPC](https://www.jsonrpc.org/specification) on all transports. You
can reuse the same connection for multiple requests!

### Operating a network

To use bitbon node - please use `--network` parameter for setting your private network.

#### Defining the private genesis state

First, you'll need to create the genesis state of your networks, which all nodes need to be
aware of and agree upon. This consists of a small JSON file (e.g. call it `genesis.json`):

```json
{
  "config": {
    "chainId": {CHAINID},
    "homesteadBlock": 0,
    "eip155Block": 0,
    "eip150Block": 0,
    "eip160Block": 0,
    "eip158Block": 0,
    "byzantiumBlock": 0,
    "maxGenDAGBlock": 0,
    "cposBlock": 0,
    "constantinopleBlock":0,
    "unlimitedCodeSizeBlock": 0,
    "istanbulBlock": 0,
    "petersburgBlock": 0,
    "quorum": {
      "period": {PERIOD},
      "epoch": {EPOCH},
      "initialSigners": ["{ADDR_MINER}"],
      "initialHash": "0x0000000000000000000000000000000000000000000000000000000000000001",
      "start": {START}
    }
  },
  "alloc": {},
  "nonce": "0x0000000000000000",
  "timestamp": "0x00",
  "parentHash": "0x0000000000000000000000000000000000000000000000000000000000000000",
  "extraData": "0x0000000000000000000000000000000000000000000000000000000000000000c46eae38e2ab01f91de7afff3097596dfff0c49366acc38944b6f36284dde66f9d6ae7cbb16f5e740000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
  "gasLimit": "0xFFFFFFFFFF",
  "difficulty": "0x00",
  "mixhash": "0x0000000000000000000000000000000000000000000000000000000000000000",
  "coinbase": "0x0000000000000000000000000000000000000000"
}
```
where:
* {CHAINID} - integer value, you network identification
* {PERIOD} - integer value, block mining interval (in seconds)
* {EPOCH} - integer value, number of seconds for every epoch
* {START} - integer value, unix timestamp (seconds) for marking CPOS consencus start time
* {ADDR_MINER} - miner address, from input parameters --miner.etherbase

The above fields should be fine for most purposes, although we'd recommend changing
the `nonce` to some random value so you prevent unknown remote nodes from being able
to connect to you. If you'd like to pre-fund some accounts for easier testing, create
the accounts and populate the `alloc` field with their addresses.

```json
"alloc": {
  "0x0000000000000000000000000000000000000001": {
    "balance": "111111111"
  },
  "0x0000000000000000000000000000000000000002": {
    "balance": "222222222"
  }
}
```

With the genesis state defined in the above JSON file, you'll need to initialize **every**
`bitbon_node` node with it prior to starting it up to ensure all blockchain parameters are correctly
set:

```shell
$ bitbon_node init path/to/genesis.json
```

# Example how to run
## Compile binary
```make clean build```
## Create genesis.json file like this
```
{
    "nonce": "0x0000000000000042",
    "timestamp": "0x00",
    "parentHash": "0x0000000000000000000000000000000000000000000000000000000000000000",
    "extraData": "0x00",
    "gasLimit": "0x1000000000",
    "difficulty": "0x4000",
    "mixhash": "0x0000000000000000000000000000000000000000000000000000000000000000",
    "coinbase": "0x0000000000000000000000000000000000000000",
    "alloc": {
      "05a8c2db4cecb2cf1f90f58ddff8ac5dba8259e1": { "balance": "100000000000000000000000000000" }
    },
    "config": {
        "chainId": 100,
        "homesteadBlock": 0,
        "eip155Block": 0,
        "eip150Block": 0,
        "eip160Block": 0,
        "eip158Block": 0,
        "byzantiumBlock": 0,
        "maxGenDAGBlock": 0,
        "cposBlock": 0,
        "unlimitedCodeSizeBlock": 0,
        "constantinopleBlock": 0,
	"istanbulBlock": 0,
	"petersburgBlock": 0,
        "quorum": {
          "period": 1,
          "epoch": 20,
          "initialSigners": ["0x79ea21cf8d2b5697977c8585239e13622da9cf8a"],
          "initialHash": "0x0000000000000000000000000000000000000000000000000000000000000001",
          "start": 1579076325
        }
    }
}
```
## Init your genesis
```
bitbon_node --datadir=/data/node init genesis.json
```
## Run node
```
btibon_node --datadir=/data/node --networkid=100 --http.api="admin,debug,personal,eth,net,web3,bitbon" --http.corsdomain="*" 
    --http --http.addr="0.0.0.0" --miner.gastarget=1000000000 --miner.gasprice=5 --nodiscover 
    --miner.etherbase=0x79ea21cf8d2b5697977c8585239e13622da9cf8a --syncmode=full --gcmode "archive" 
    --bitbon.service.id=3c31bd65-3f66-413c-90fc-5945d3b6b04b --bitbon.service.pk=2b2d0a29ac78ca225291da23ee1f5cd90c45af6d777ee784214ea51ab8e9fa7b 
     --mine --miner.threads=3 --bitbon.assetbox.decrypt.password=qwertyuiopqwerty
```

## Contribution
Thank you for the interest in our source code! We are thankful to all contributors for even the smallest of changes!

If you would like to contribute to the source code of Simcord, please fork, fix, commit and send a pull request to maintainers so that they could review it and include in the main code base.

When creating a pull request, make sure that it:
* is based on the master branch and opens against it;
* has full and detailed description;
* goes through the following stages: Makefile-а: dep, lint, test, build.

You also need to make sure that the code does not contradict our code style, which means that it:
* adheres to the official rules (uses gofmt);
* is documented in accordance with the official Go commentary guidelines.

When creating an issue, please make sure that you provide complete information:
* name;
* detailed description;
* steps of the error: version, logs, error text, screenshot, etc.;
* local parameters (operation system, native/docker, CPU & RAM & storage parameters).

If you find a vulnerability, please send us an email to this address: dev@simcord.com. If you create a public ticket, hackers can use the vulnerability you describe in their interests.

## Getting help
The best way to contact our team is through GitHub. You can create an issue and select one of the options. Our maintainers will reply to you as soon as they review your issue.

Before creating an issue, please look through the previously created issues, maybe there is an answer to yours already.

You can also contact our maintainers directly: dev@simcord.com. When contacting the customer support directly, please provide as much information regarding your issue as possible.

## Licensing

The bitbon-system-blockchain library (i.e. all code outside of the cmd directory) is licensed under the [GNU Lesser General Public License v3.0](https://www.gnu.org/licenses/lgpl-3.0.en.html) also included in our repository in the COPYING.LESSER file.

The bitbon-system-blockchain binaries (i.e. all code inside of the cmd directory) are licensed under the [GNU General Public License v3.0](https://www.gnu.org/licenses/gpl-3.0.en.html) also included in our repository in the COPYING file.
