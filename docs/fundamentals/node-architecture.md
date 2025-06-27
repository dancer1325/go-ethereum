---
title: Node architecture
description: Introduction to how Ethereum nodes are organized and where Geth fits.
---

* see [Ethereum documentation](https://ethereum.org/en/developers/docs/nodes-and-clients/node-architecture/)

* Ethereum node
  * == [execution client](https://ethereum.org/en/developers/docs/nodes-and-clients/#execution-clients) + [consensus client](https://ethereum.org/en/developers/docs/nodes-and-clients/#consensus-clients)

* Geth
  * == [execution client](https://ethereum.org/en/developers/docs/nodes-and-clients/#execution-clients)
  * |
    * originally,
      * 👀run a FULL Ethereum node == run ONLY execution client 👀
    * Ethereum turned off [proof-of-work](https://ethereum.org/en/developers/docs/consensus-mechanisms/pow/) & implemented [proof-of-stake](https://ethereum.org/en/developers/docs/consensus-mechanisms/pow/),
      * ⚠️run a FULL Ethereum node == run Geth + run [“consensus client”](https://ethereum.org/en/developers/docs/nodes-and-clients/#consensus-clients)⚠️
        * Reason: 🧠keep track of the Ethereum blockchain🧠

![node-architecture](../../public/images/docs/node-architecture-text-background.png)

## What does Geth do? {#what-does-geth-do}

* see [execution client](https://ethereum.org/en/developers/docs/nodes-and-clients/node-architecture/)

## What does the consensus client do? {#consensus-client}

* see [consensus client](https://ethereum.org/en/developers/docs/nodes-and-clients/node-architecture/)

## Validators {#validators}

* see 
  * [validators](https://ethereum.org/en/developers/docs/nodes-and-clients/node-architecture/)
  * [proof-of-stake](https://ethereum.org/en/developers/docs/consensus-mechanisms/pos/)
