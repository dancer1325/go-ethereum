* `func init() {}`
  * 's goal
    * initialize CLI app
    * start Geth

* `func prepare(ctx *cli.Context) {}`
  * responsible for
    * manipulate
      * memory cache assigned
        * 👀if you run FULL node | mainnet / WITHOUT specified cache -> bump default cache allowance👀
      * setups metric system

* `func geth(ctx *cli.Context) error {}`
  * 👀if NO special subcommand is run -> main entry point | system👀 
  * 💡creates a default node -- based on the -- CL arguments💡
  * runs the node | blocking mode (TODO: ❓)
    * == wait for it -- to be -- shut down

* `func startNode(ctx *cli.Context, stack *node.Node, isConsole bool) {}`
  * boots up the system node and all registered protocols, after which
    // it starts the RPC/IPC interfaces and the miner.