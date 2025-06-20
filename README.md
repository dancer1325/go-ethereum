## Go Ethereum

* == 👀Ethereum protocol's execution layer implementation👀
  * Golang

[![API Reference](
https://pkg.go.dev/badge/github.com/ethereum/go-ethereum
)](https://pkg.go.dev/github.com/ethereum/go-ethereum?tab=doc)

[![Discord](https://img.shields.io/badge/discord-join%20chat-blue.svg)](https://discord.gg/nthXNEv)

* AUTOMATED builds
  * uses
    * stable releases
    * unstable master branch
* [Binary archives](https://geth.ethereum.org/downloads/)

## Building the source

* [Installation Instructions](https://geth.ethereum.org/docs/getting-started/installing-geth)

* requirements
  * Go v1.23+
  * C compiler

* ways to build
  * `make geth`
  * `make all`
    * build the FULL suite of utilities

## Executables

* == SEVERAL wrappers/executables
* | ["cmd"](cmd)

|  Command   | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       |
| :--------: |-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| **`geth`** | 💡MAIN Ethereum CLI client💡<br/> == entry point \| Ethereum network (main-, test- or private net) <br/> ways of running <br/> &nbsp;&nbsp; FULL node (default) <br/> &nbsp;&nbsp; archive node ( == retain ALL historical state) <br/> &nbsp;&nbsp; light node (== retrieve data LIVE) <br/> expose JSON RPC endpoints / AVAILABLE \| <br/> &nbsp;&nbsp; HTTP <br/> &nbsp;&nbsp; WebSocket <br/> &nbsp;&nbsp; IPC transports                                                                                     |
|   `clef`   | == stand-alone signing tool <br/> uses <br/> &nbsp;&nbsp; backend signer -- for -- `geth`                                                                                                                                                                                                                                                                                                                                                                                                                         |
|  `devp2p`  | == utilities -- to interact with -- nodes \| networking layer WITHOUT running a FULL blockchain                                                                                                                                                                                                                                                                                                                                                                                                                   |
|  `abigen`  | == source code generator / Ethereum contract definitions -- are converted into -- Go packages (easy-to-use, compile-time type-safe) <br/> operates \| <br/> &nbsp;&nbsp; plain [Ethereum contract ABIs](https://docs.soliditylang.org/en/develop/abi-spec.html) / <br/> &nbsp;&nbsp; &nbsp;&nbsp; if the contract bytecode is ALSO AVAILABLE -> expanded functionality <br/> &nbsp;&nbsp; solidity source files <br/> see [Native DApps](https://geth.ethereum.org/docs/developers/dapp-developer/native-bindings) |
|   `evm`    | == EVM's developer utility version / <br/> &nbsp;&nbsp; run bytecode snippets \| configurable environment & execution mode <br/> &nbsp;&nbsp; enable <br/> &nbsp;&nbsp; &nbsp;&nbsp; debug EVM opcodes                                                                                                                                                                                                                                                                                                            |
| `rlpdump`  | == developer utility tool / binary RLP ([Recursive Length Prefix](https://ethereum.org/en/developers/docs/data-structures-and-encoding/rlp)) dumps -- are converted into -- hierarchical representation                                                                                                                                             |

## Running `geth`

### Hardware Requirements

* Minimum
  * CPU with 4+ cores
  * 8GB RAM
  * 1TB free storage space to sync the Mainnet
  * 8 MBit/sec download Internet service

* recommendations
  * Fast CPU with 8+ cores
  * 16GB+ RAM
  * High-performance SSD with at least 1TB of free space
  * 25+ MBit/sec download Internet service

### FULL node | main Ethereum network

* COMMON use case
  * SIMPLY interact -- with the -- Ethereum network
    * _Example:_ create accounts; transfer funds; deploy and interact with contracts
    * == NOT care about years-old historical data
    * == sync quickly -- to the -- network's CURRENT state

    ```shell
    $ geth console
    ```
  * `geth`
    * starts | 👀snap sync mode👀
      * if you want to change -> pass `--syncmode` flag
      * download MORE data 
        * Reason: 🧠avoid processing the Ethereum network's WHOLE history🧠

* `geth ... console`
  * start the built-in interactive [JS console](https://geth.ethereum.org/docs/interacting-with-geth/javascript-console) /
    * allows you
      * interact -- via -- [`web3` methods](https://github.com/ChainSafe/web3.js/blob/0.20.7/DOCUMENTATION.md)
        * ❌NOT up to date -- with -- official docs❌

* `geth attach`
  * allows
    * | ALREADY running `geth` instance,
      * attach JS console 

### FULL node | Holesky test network

* use case
  * play around -- with -- creating Ethereum contracts

```shell
$ geth --holesky console
```

* `console` subcommand
  * 's meaning == testnet's meaning

* `--holesky` flag
  * reconfigure your `geth` instance / 👀client will connect -- to the -- Holesky test network👀
    * DIFFERENT P2P bootnodes, network IDs & genesis states
  * nest itself 1 level deeper | `holesky/` subfolder
    * _Example:_ | Linux,
      * default data directory `~/.ethereum`
      * holesky network `~/.ethereum/holesky`
  * if you want to attach | running testnet node -> you need to use a custom endpoint 
    * AFFECT | OSX & Linux
    * Reason: 🧠`geth attach` will try to attach | production node endpoint, by default
      * _Example:_ `geth attach <datadir>/holesky/geth.ipc`🧠
    * NOT AFFECT | Windows

* recommendations
  * use account | main network != account | Holesky
    * ALTHOUGH EXIST INTERNAL protective measures / prevent transactions BETWEEN the main network -- & -- Holesky

### Configuration

* TODO: As an alternative to passing the numerous flags to the `geth` binary, you can also pass a
configuration file via:

```shell
$ geth --config /path/to/your_config.toml
```

To get an idea of how the file should look like you can use the `dumpconfig` subcommand to
export your existing configuration:

```shell
$ geth --your-favourite-flags dumpconfig
```

#### Docker quick start

One of the quickest ways to get Ethereum up and running on your machine is by using
Docker:

```shell
docker run -d --name ethereum-node -v /Users/alice/ethereum:/root \
           -p 8545:8545 -p 30303:30303 \
           ethereum/client-go
```

This will start `geth` in snap-sync mode with a DB memory allowance of 1GB, as the
above command does.  It will also create a persistent volume in your home directory for
saving your blockchain as well as map the default ports. There is also an `alpine` tag
available for a slim version of the image.

Do not forget `--http.addr 0.0.0.0`, if you want to access RPC from other containers
and/or hosts. By default, `geth` binds to the local interface and RPC endpoints are not
accessible from the outside.

### Programmatically interfacing `geth` nodes

As a developer, sooner rather than later you'll want to start interacting with `geth` and the
Ethereum network via your own programs and not manually through the console. To aid
this, `geth` has built-in support for a JSON-RPC based APIs ([standard APIs](https://ethereum.org/en/developers/docs/apis/json-rpc/)
and [`geth` specific APIs](https://geth.ethereum.org/docs/interacting-with-geth/rpc)).
These can be exposed via HTTP, WebSockets and IPC (UNIX sockets on UNIX based
platforms, and named pipes on Windows).

The IPC interface is enabled by default and exposes all the APIs supported by `geth`,
whereas the HTTP and WS interfaces need to manually be enabled and only expose a
subset of APIs due to security reasons. These can be turned on/off and configured as
you'd expect.

HTTP based JSON-RPC API options:

  * `--http` Enable the HTTP-RPC server
  * `--http.addr` HTTP-RPC server listening interface (default: `localhost`)
  * `--http.port` HTTP-RPC server listening port (default: `8545`)
  * `--http.api` API's offered over the HTTP-RPC interface (default: `eth,net,web3`)
  * `--http.corsdomain` Comma separated list of domains from which to accept cross-origin requests (browser enforced)
  * `--ws` Enable the WS-RPC server
  * `--ws.addr` WS-RPC server listening interface (default: `localhost`)
  * `--ws.port` WS-RPC server listening port (default: `8546`)
  * `--ws.api` API's offered over the WS-RPC interface (default: `eth,net,web3`)
  * `--ws.origins` Origins from which to accept WebSocket requests
  * `--ipcdisable` Disable the IPC-RPC server
  * `--ipcpath` Filename for IPC socket/pipe within the datadir (explicit paths escape it)

You'll need to use your own programming environments' capabilities (libraries, tools, etc) to
connect via HTTP, WS or IPC to a `geth` node configured with the above flags and you'll
need to speak [JSON-RPC](https://www.jsonrpc.org/specification) on all transports. You
can reuse the same connection for multiple requests!

**Note: Please understand the security implications of opening up an HTTP/WS based
transport before doing so! Hackers on the internet are actively trying to subvert
Ethereum nodes with exposed APIs! Further, all browser tabs can access locally
running web servers, so malicious web pages could try to subvert locally available
APIs!**

### Operating a private network

Maintaining your own private network is more involved as a lot of configurations taken for
granted in the official networks need to be manually set up.

Unfortunately since [the Merge](https://ethereum.org/en/roadmap/merge/) it is no longer possible
to easily set up a network of geth nodes without also setting up a corresponding beacon chain.

There are three different solutions depending on your use case:

  * If you are looking for a simple way to test smart contracts from go in your CI, you can use the [Simulated Backend](https://geth.ethereum.org/docs/developers/dapp-developer/native-bindings#blockchain-simulator).
  * If you want a convenient single node environment for testing, you can use our [Dev Mode](https://geth.ethereum.org/docs/developers/dapp-developer/dev-mode).
  * If you are looking for a multiple node test network, you can set one up quite easily with [Kurtosis](https://geth.ethereum.org/docs/fundamentals/kurtosis).

## Contribution

 * official Go [formatting](https://golang.org/doc/effective_go.html#formatting) guidelines 
   * == uses [gofmt](https://golang.org/cmd/gofmt/)
 * Code 
   * document it -- following -- the official Go [commentary](https://golang.org/doc/effective_go.html#commentary)
 * commit messages
   * prefix -- with the -- package(s) / they modify
     * _Example:_ "eth, rpc: make trace configs optional"

* TODO: Please see the [Developers' Guide](https://geth.ethereum.org/docs/developers/geth-developer/dev-guide)
for more details on configuring your environment, managing project dependencies, and
testing procedures.

### | geth.ethereum.org

* 💡[`website` branch](https://github.com/ethereum/go-ethereum/tree/website#readme)💡
  * == documentation
