* `type Block struct {}`    
  * == Ethereum block /
    * goal
      * be 'immutable' -- based on -- 's caches
  * block immutability rules
    * | construct the block, copy ALL data
      * makes references held / block -- independent of -- whatever value was passed in
    * | access,
      * copy ALL header data 
        * Reason: 🧠any change to the header would mess up the cached hash & size values | block🧠
        * -> better performance | call code
      * ❌NOT copy body data❌
        * Reason: 🧠NOT affect the caches & too expensive🧠 
    * if NEW body data is attached | block -> returns a shallow copy of the block
      * == block modifications are race-free

* `type Header struct {}`
  * == Ethereum blockchain's block header
