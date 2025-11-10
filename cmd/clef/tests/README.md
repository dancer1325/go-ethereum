* goal
  * testing clef-functionality

## how to run locally [testsigner.js](testsigner.js)?
* `build/bin/clef --4bytedb=./cmd/clef/4byte.json --rpc`
  * start clef
    * Problems:
      * Problem1: "no such file or directory: build/bin/clef"
        * Solution: TODO:
* `build/bin/geth --nodiscover --maxpeers 0 --signer http://localhost:8550 console --preload=cmd/clef/tests/testsigner.js`
  * start geth
    * Problems:
      * Problem1: "no such file or directory: build/bin/geth"
        * Solution: TODO:
  * | console,
    ```shell
    > test()
    ```
