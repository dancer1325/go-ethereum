---
title: Batch requests
description: How to make batch requests using JSON-RPC API
---

* [JSON-RPC specification batch](https://www.jsonrpc.org/specification#batch)
  * 👀implemented by Geth's API👀
  * allows
    * cut network delays
  * uses
    * fetch large amounts of MOSTLY independent data objects

* _Example:_ fetch a list of blocks
  * NO dependency BETWEEN requests
      ```js
      import fetch from 'node-fetch';
    
      async function main() {
        const endpoint = 'http://127.0.0.1:8545';
        const from = parseInt(process.argv[2]);
        const to = parseInt(process.argv[3]);
    
        const reqs = [];
        for (let i = from; i < to; i++) {
          reqs.push({
            method: 'eth_getBlockByNumber',
            params: [`0x${i.toString(16)}`, false],
            id: i - from,
            jsonrpc: '2.0'
          });
        }
    
        const res = await fetch(endpoint, {
          method: 'POST',
          body: JSON.stringify(reqs),
          headers: { 'Content-Type': 'application/json' }
        });
        const data = await res.json();
      }
    
      main()
        .then()
        .catch(err => console.log(err));
      ```

* Often the retrieved data from one request is needed to issue a second one
* The JSON-RPC API provides `eth_getTransactionReceipt` which takes in a transaction hash and returns the corresponding receipt object, but no method to fetch receipt objects for a whole block
* We need to get the list of transactions in a block and then call `eth_getTransactionReceipt` for each of them.

We can break this into 2 batch requests:

- First to download the list of transaction hashes for all of the blocks in our desired range
- And then to download the list of receipts objects for all of the transaction hashes

* if you depend on SEVERAL JSON-RPC endpoints ->
  * batching approach PRETTY complicated
  * 👀use [GraphQL API](/docs/interacting-with-geth/rpc/graphql)👀
