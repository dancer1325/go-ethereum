---
title: Hardware requirements
description: Overview of the hardware needed to run an Ethereum node
---

* hardware requirements -- for -- running a Geth node
  * -- depend upon the --
    * node configuration
    * network's upgrades
      * == change over time

## Processor {#processor}

* Geth is released | WIDE range of architectures
* recommendation
  * quad-core OR
  * dual-core hyperthreaded

## Memory {#memory}

* \>= 16GB RAM

## Disk space {#disk-space}

* MAIN bottleneck
  * | September 2022,
    * | FULL node,
      * \>= 2TB SSD
    * | FULL archive node,
      * \> 12TB 
    * Geth / snap-synced
      * \>650GB
      * default cache size grows 14GB/week
        * if you prune -> back down to 650GB

* recommendations
  * [GitHub Gist](https://gist.github.com/yorickdowne/f3a3e79a573bf35767cd002cc977b038)

## Bandwidth {#bandwidth}

* internet connection
  * stable
  * reliable

* recommendations
  * \>= 25Mbps download speed
  * ISP / NOT have a capped data allowance

* crucial |
  * run a validator
