---
title: Security
description: A primer on Geth security best practice.
---

* goal
  * Geth's basic security best-practises 

## Downloading Geth {#downloading-geth}

* | [Downloads page](/downloads) 
  * downloaded files' SHA256 hashes
    * confirm precise consistency -- with -- our releases

## Networking security {#networking-security}

* recommended local machine's firewall settings
  - | port / JSON-RPC requests -- to the -- node
    - by default, "8545"
    - 👀block ALL traffic 👀
      - EXCEPT -- from -- explicitly defined trusted machines
  - | port / peer-to-peer communications
    - == node can connect -- to -- peers
    - by default, "TCP 30303",
    - 👀ALLOW traffic👀
  - | port / peer-to-peer communications
    - == enable node discovery 
    - by default, "UDP 30303"
    - 👀ALLOW traffic👀

## API security {#api-security}

* Geth's API endpoints
  * are
    * legacy json-rpc
    * NEW trusted "beacon" json-rpc API
    * graphql endpoint
  * ❌NOT designed to❌
    * withstand attacks -- by -- hostile clients
    * handle huge amounts of clients/traffic 
  * / expose towards "the internet" OR ANY untrusted/hostile network
    * risk
      * crashes -- due to -- OOM
      * NOT keeping up with chain progression
        * -- due to -- resource starvation (IO or CPU)
      * attempts to steal funds -- via -- spurious signing-requests
    * ❌NOT recommended ❌
    * recommendations
      * setting up
        * proxies,
        * WAFs,
        * application level filtering,
        * rate limiting,
        * logging,
        * tls terminator
        * monitoring

## Account security {#account-security}

* == keep private keys & account passwords / 
  * backed up
  * inaccessible -- to -- adversaries

* users' responsibility

* Geth
  * provides
    * 👀encrypted store -- for -- keys👀
      * if you want to unlock -> use an account password
      * if the key files OR passwords are lost -> account is IMPOSSIBLE to access == funds are lost forever
      * ANYONE / has access to the unencrypted keys -> control account's funds
    * built-in account management tools
      * | FUTURE,
        * will be deprecated
      * recommendations
        * ⚠️use Clef⚠️
          * == EXTERNAL 
            * account management
            * signing tool
          * allows
            * run it | dedicated (_Example:_ VM OR secure USB drive) secure external hardware != hardware / run Geth 
              * Reason:🧠sign locally | Clef != giving key access -- to a -- node🧠

## Other security considerations {#other-security}

* out of scope PERFECT secure nodes
  * insecure [smart contracts ](https://ethereum.org/en/developers/docs/smart-contracts/security)

* see [Etherem security best practice](https://ethereum.org/en/security) 
