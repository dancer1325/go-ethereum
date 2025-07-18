* `type protoHandshake struct {}`
  * == protocol handshake's RLP structure 

* `type PeerEvent struct {}`
  * use cases / this event is emitted
    * peers are either added OR dropped -- from a -- `p2p.Server` OR
    * message is sent or received | peer connection

* `type Peer struct {}`
  * == connected remote node

* `type PeerEventType string`
  * == peer event types / emitted -- by a -- `p2p.Server`

* TODO: