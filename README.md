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

* ways to configure 
  * pass CL's flags
  * pass a configuration file -- via -- `--config pathToConfigurationFile`

    ```shell
    $ geth --config /path/to/your_config.toml
    ```

    * if you want to know the configuration / EACH flag -> use `dumpconfig` subcommand

      ```shell
      $ geth --your-favourite-flags dumpconfig
      ```

### -- via -- Docker

* get Ethereum up & running | your machine

  ```shell
  # start `geth` / 
  #  1. snap-sync mode
  #  2. DB memory allowance == 1GB
  #  3. create a persistent volume | your home directory / save your blockchain  
  docker run -d --name ethereum-node -v /Users/alice/ethereum:/root \
             -p 8545:8545 -p 30303:30303 \
             ethereum/client-go
  ```

  * if you want to access RPC from OTHER containers OR hosts -> pass `--http.addr 0.0.0.0` 
    * by default, 
      * `geth` binds -- to the -- local interface
      * RPC endpoints are NOT accessible -- from the -- outside

### Programmatically interfacing `geth` nodes

* goal
  * interact with `geth` & Ethereum network -- via -- your OWN programs
    * == ❌NOT MANUALLY -- through the -- console❌

* `geth`
  * 👀built-in support -- for a -- APIs 
    * JSON-RPC based APIs ([standard APIs](https://ethereum.org/en/developers/docs/apis/json-rpc/)
    * [`geth` specific APIs](https://geth.ethereum.org/docs/interacting-with-geth/rpc)👀
  * ways to expose these APIs
    * HTTP,
      * by default, disabled
        * == you need to MANUALLY, enable it
      * ONLY expose a subset of APIs -- due to -- security reasons
    * WebSockets
      * by default, disabled
        * == you need to MANUALLY, enable it
      * ONLY expose a subset of APIs -- due to -- security reasons
    * IPC (UNIX sockets | UNIX based platforms, & named pipes | Windows)
      * by default, enabled
      * exposes ALL the APIs / supported by `geth`

* HTTP based JSON-RPC API options
  * `--http`
    * enable the HTTP-RPC server
  * `--http.addr`
    * HTTP-RPC server listening interface
    * by default, `localhost`
  * `--http.port`
    * HTTP-RPC server listening port
    * by default, `8545`
  * `--http.api`
    * API's offered -- over the -- HTTP-RPC interface
    * by default, `eth,net,web3`
  * `--http.corsdomain`
    * == `domain1,domain2, ...`
    * domains / from which -- to -- accept cross-origin requests
  * `--ws`
    * enable the WS-RPC server
  * `--ws.addr`
    * WS-RPC server listening interface
    * by default, `localhost`
  * `--ws.port`
    * WS-RPC server listening port
    * by default, `8546`
  * `--ws.api`
    * API's offered over the WS-RPC interface
    * by default, `eth,net,web3`
  * `--ws.origins`
    * origins | accept WebSocket requests
  * `--ipcdisable`
    * disable the IPC-RPC server
  * `--ipcpath`
    * datadir's filename for IPC socket/pipe

* SAME connection
  * can be reused -- for -- MULTIPLE requests

### Operating a private network

* | since [the Merge](https://ethereum.org/en/roadmap/merge/),
  * set up a network of geth nodes + corresponding beacon chain

* SOLUTIONS -- depending on -- your use case
  * if you are looking for
    * 1 way to test smart contracts | your CI -> use the [Simulated Backend](https://geth.ethereum.org/docs/developers/dapp-developer/native-bindings#blockchain-simulator)
    * multiple node test network -> set one up -- with -- [Kurtosis](https://geth.ethereum.org/docs/fundamentals/kurtosis) 
  * if you want a convenient 1 node environment for testing -> use [Dev Mode](https://geth.ethereum.org/docs/developers/dapp-developer/dev-mode)

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
