* `type Subscription interface {}`
  * == event subscription / 
    * events are delivered | data channel
  * `Unsubscribe()`
    * cancel sending events -- to the -- data channel
    * closes the error channel
  * `Err() <-chan error`
    * returns the subscription error channel
      * if there is an issue -- with the -- subscription -> error channel receives a value
        * _Example:_ the network connection delivering the events has been closed
      * ONLY sent 1 value
      * error channel is closed -- by -- `Unsubscribe()`

* `type ChainReader interface {}`
  * provides 
    * 👀access -- to the -- blockchain👀
      * == raw data -- from --
        * canonical chain (| request -- by -- block number) OR
        * any blockchain fork / was PREVIOUSLY downloaded & processed by the node
      * if you pass block number argument == nil -> select the LATEST canonical block
      * recommendations
        * read block headers 
          * rather than FULL blocks
  * if the requested item does NOT exist -> return the error `NotFound`
  * `SubscribeNewHead(ctx context.Context, ch chan<- *types.Header) (Subscription, error)`
    * subscribe notifications -- about -- canonical chain's head block changes

* `type TransactionReader interface {}`
  * provides
    * access -- to -- past transactions & their receipts
      * ❌HISTORIC transactions may NOT be AVAILABLE❌
  * 's implementations
    * may impose ARBITRARY restrictions | transactions & receipts / can be retrieved
  * recommendations
    * if POSSIBLE -> use ANOTHER interface
      * _Example:_ `LogFilterer` 
  * if the requested item does NOT exist -> returns `NotFound`
  * `TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error)`
    * returns the `Receipt`
      * == mined transaction's receipt 
    * ALTHOUGH a receipt exists -- transaction may NOT be included | CURRENT canonical chain 

* `type ChainStateReader interface {}`
  * wraps 
    * access -- to the -- canonical blockchain's state trie 
  * 's implementations
    * | OLD blocks,
      * may NOT return state values 
  * recommendations
    * if you want to read raw contract storage -> use `ContractCaller.CallContract` 

* `type SyncProgress struct {}`
  * | node syncs with the Ethereum network,
    * share progress indications 

* `func (prog SyncProgress) Done() bool {}`
  * if the initial sync is done -> return `bool`

* `type ChainSyncReader interface {}`
  * wraps 
    * access -- to the -- node's current sync status
      * if there's NO sync currently running -> returns nil

* `type CallMsg struct {}`
  * == contract calls' parameters

* `type ContractCaller interface {}`
  * provides
    * contract calls
      * MAINLY, transactions / 
        * are executed -- by the -- EVM
        * NOT mined | blockchain
  * `CallContract(ctx context.Context, call CallMsg, blockNumber *big.Int) ([]byte, error)`
    * low-level method -- to -- execute contract calls
  * recommendations
    * | applications / structured around specific contracts,
      * use abigen tool
        * Reason: 🧠nicer, properly typed🧠

* `type FilterQuery struct {}`
  * uses
    * contract log filtering

* `type LogFilterer interface {}`
  * provides
    * access -- to -- contract log events
      * -- via -- 1-off query OR continuous event subscription
  * logs / received through a streaming query subscription, may have `Removed=true` 
    * == log was reverted -- due to a -- chain reorganisation

* `type TransactionSender interface {}`
  * wraps
    * transaction sending
  * `SendTransaction(ctx context.Context, tx *types.Transaction) error`
    * injects a signed transaction | pending transaction pool -- for -- execution
      * == requirements of the transaction
        * signed transaction
        * transaction have a valid `nonce`  
    * if the transaction == contract creation & | AFTER mining the transaction, you want to retrieve the contract address  -> use `TransactionReceipt` method
    * API consumers
      * if they want to maintain local private keys -> use package accounts
      * if they need to retrieve the NEXT AVAILABLE `nonce` -> use `PendingNonceAt`

* `type GasPricer interface {}`
  * wraps
    * gas price oracle /
      * monitors the blockchain
        * Reason: 🧠| CURRENT fee market conditions, determine the optimal gas price 🧠

* `type GasPricer1559 interface {}`
  * provides
    * access -- to the -- EIP-1559 gas price oracle

* `type FeeHistoryReader interface {}`
  * provides
    * access -- to the -- fee history oracle

* `type FeeHistory struct {}`
  * provides 
    * recent fee market data 
  * allows
    * determining a reasonable `maxPriorityFeePerGas` value

* `type PendingStateReader interface {}`
  * provides
    * access -- to the -- pending state
      * == 👀(ALL known executable transactions / NOT YET been included | blockchain)'s result👀
  * uses
    * display ’unconfirmed’ actions' result -- initiated by the -- user
      * _Example:_ wallet value transfers
  * `PendingNonceAt(ctx context.Context, account common.Address) (uint64, error)`
    * allows
      * retrieving the NEXT available transaction `nonce` / specific account

* `type PendingContractCaller interface {}`
  * uses
    * perform calls -- against the -- pending state

* `type GasEstimator interface {}`
  * wraps
    * `EstimateGas`
  * `EstimateGas(ctx context.Context, call CallMsg) (uint64, error)`
    * tries to estimate the gas needed
      * -- based on the -- pending state
      * -- to -- execute a specific transaction
      * ⚠️tries == NO guarantee / it's the true gas limit requirement⚠️
        * Reason: 🧠OTHER transactions -- may be, by miners, -- added OR removed🧠
        * BUT valid -- for -- setting a reasonable default 

* `type PendingStateEventer interface {}`
  * provides 
    * access -- to -- real time notifications -- about -- pending state changes

* `type BlockNumberReader interface {}`
  * provides
    * access -- to the -- CURRENT block number

* `type ChainIDReader interface {}`
  * provides
    * access -- to the -- chain ID