---
title: JSON-RPC Server
description: Introduction to the JSON_RPC server
---

* Geth
  * requires
    * sending requests -- to -- SPECIFIC JSON-RPC API methods
  * supports
    * ALL standard [JSON-RPC API](https://github.com/ethereum/execution-apis) endpoints

* RPC 
  * requests MUST be sent -- to the -- node
  * response returned -- , via SOME transport protocol, to the -- client
  * provide the information /
    * required
      * -- by -- users 
      * -- to -- choose a transport protocol / specific user scenario

## Introduction {#introduction}

* JSON-RPC
  * provided | MULTIPLE transports

* Geth
  * supports JSON-RPC -- over -- HTTP, WebSocket & Unix Domain Sockets

* Transports
  * MUST be enabled -- through -- CL flags

* Ethereum JSON-RPC APIs
  * use a name-space system

* RPC methods
  * are grouped | SEVERAL categories -- depending on -- their purpose
  * 's names == namespace + underscore + actual method name | namespace
    * _Example:_ `eth_call` method resides | `eth` namespace
  * enable it / per-namespace

## Transports {#transports}

### HTTP Server {#http-server}

* [HTTP](https://developer.mozilla.org/en-US/docs/Web/HTTP)
  * := transport protocol
    * unidirectional
    * client -- is connected with a -- server
      * AFTER response is returned -> HTTP connection is closed
  * supported | 
    * EVERY browser
    * ALL programming toolchains
  * MOST used
    * Reason: 🧠its ubiquity🧠
  * `geth --http`
    * 💡start a HTTP server | Geth💡
    * by default,
      * accept connections | 
        * local loopback interface (== 127.0.0.1)
        * listening port "8545"
    * if you want to customize the ip address & listening port -> use `--http.addr` & `--http.port` respectively

      ```sh
      geth --http --http.port 3334
      ```
  * whitelist
    * 👀by default, enable access -- to -- `eth`, `net` & `web3` namespaces👀
    * enable
      * access -- , via configuring by `--http.api`, to -- OTHER APIs
      
        ```sh
        geth --http --http.api eth,net,web3
        ``` 
      
        * _Example:_ debugging (`debug`) 
    * ❌NOT recommended❌
      * Reason: 🧠access to these methods -> increases the attack surface🧠
    * if you call non-whitelisted RPC namespaces -> returns an RPC error / code == `-32602`
  * HTTP server
    * built into additional protection / prevent misuse of the API -- from -- web pages
      * Reason: 🧠it's reachable -- from -- ANY local application🧠 
    * if you want to enable access to the API -- from a -- web page -> configure the server / accept Cross-Origin requests
      * == -- via -- `--http.corsdomain` flag

      ```sh
      geth --http --http.corsdomain https://remix.ethereum.org
      ```

      * if you want to enable access to the RPC -- from -- ANY origin -> `--http.corsdomain '*'`

### WebSocket Server {#websockets-server}

* Websocket
  * := transport protocol
    * bidirectional
      * == maintained -- by -- client & server
      * if client OR server terminates it -> EXPLICITLY terminated
      * -> servers can push events -- to -- clients 
  * supported by
    * MOST MODERN browsers
  * use cases
    * [event subscription](/docs/interacting-with-geth/rpc/pubsub)
    * send HIGH number of requests 
  * | AFTER the handshake procedure,
    * individual messages' overhead == low,

* TODO: 
Configuration of the WebSocket endpoint in Geth follows the same pattern as the HTTP transport
* WebSocket access can be enabled using the `--ws` flag
* If no additional information is provided, Geth falls back to its default behaviour which is to establish the Websocket on port 8546
* The `--ws.addr`, `--ws.port` and `--ws.api` flags can be used to customize settings for the WebSocket server
* For example, to start Geth with a Websocket connection for RPC using
the custom port 3334 and whitelisting the `eth`, `net` and `web3` namespaces:

```sh
geth --ws --ws.port 3334 --ws.api eth,net,web3
```

Cross-Origin request protection also applies to the WebSocket server
* The `--ws.origins` flag can be used to allow access to the server from web pages:

```sh
geth --ws --ws.origins http://myapp.example.com
```

As with `--http.corsdomain`, using the wildcard `--ws.origins '*'` allows access from any origin.

<Note>By default, **account unlocking is forbidden when HTTP or Websocket access is enabled** (i.e. by passing `--http` or `ws` flag)
* This is because an attacker that manages to access the node via the externally-exposed HTTP/WS port can then control the unlocked account
* It is possible to force account unlock by including the `--allow-insecure-unlock` flag but this is unsafe and **not recommended** except for expert users that completely understand how it can be used safely
* This is not a hypothetical risk: **there are bots that continually scan for http-enabled Ethereum nodes to attack**</Note>

### IPC Server {#ipc-server}

IPC is normally available for use in local environments where the node and the console exist on the same machine
* Geth creates a pipe in the computers local file system (at `ipcpath`) that configures a connection between node and console
* The `geth.ipc` file can also be used by other processes on the same machine to interact with Geth.

On UNIX-based systems (Linux, OSX) the IPC is a UNIX domain socket
* On Windows IPC is provided using named pipes
* The IPC server is enabled by default and has access to all JSON-RPC namespaces.

The listening socket is placed into the data directory by default
* On Linux and macOS, the default location of the geth socket is

```sh
~/.ethereum/geth.ipc
```

On Windows, IPC is provided via named pipes
* The default location of the geth pipe is:

```sh
\\.\pipe\geth.ipc
```

The location of the socket can be customized using the `--ipcpath` flag. IPC can be disabled
using the `--ipcdisable` flag.

## Choosing a transport protocol {#choosing-transport-protocol}

* MAIN characteristics / EACH transport protocol

|                               | HTTP  |  WS   |  IPC  |
| :---------------------------: | :---: | :---: | :---: |
|      Event subscription       |   N   | **Y** | **Y** |
|       Remote connection       | **Y** | **Y** |   N   |
| Per-message metadata overhead | high  |  low  |  low  |

* general rules
  * IPC
    * MOST secure
      * Reason: 🧠
        * limited to interactions | local machine
        * can NOT be exposed -- to -- EXTERNAL traffic🧠
    * uses
      * subscribe to events
  * HTTP
    * familiar
    * idempotent
    * if the number of requests is low -> lower overall overheads 
  * Websockets
    * enable
      * event subscriptions
      * streaming
    * use case
      * large volumes of requests / small overheads / message

## Engine-API {#engine-api}

* Engine-API
  * == set of RPC methods /
    * enable communication between Geth -- & -- [consensus client](/docs/getting-started/consensus-clients)
  * ❌NOT designed to❌
    * be exposed to the user
  * designed to
    * if they need to exchange information -> be called AUTOMATICALLY -- by the -- clients 
  * default,
    * enabled
      * == ❌NOT required -- to pass -- any instruction to Geth❌
  * [spec](https://github.com/ethereum/execution-apis/blob/main/src/engine)

## Summary {#summary}

RPC requests to a Geth node can be made using three different transport protocols
* The protocols are enabled at startup using their respective flags
* The right choice of transport protocol depends on the specific use case.
