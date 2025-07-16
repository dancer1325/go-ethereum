* TODO:

* `type TypeMuxEvent struct {}`
  * notification / pushed -- to -- subscribers /
    * time-tagged

* `type TypeMux struct {}`
  * dispatches events -- to -- registered receivers / handle events
    * Reason of register receivers: 🧠handle events🧠
  * operation / called AFTER mux is stopped
    * -> will return `ErrMuxClosed`
  * `0`
    * ready to use
  * ⚠️Deprecated⚠️
    * use [Feed](feed.go)

* `type TypeMuxSubscription struct {}`
  * subscription established -- through -- `TypeMux`

* `func (mux *TypeMux) Subscribe(types ...interface{}) *TypeMuxSubscription {}`
  * creates a subscription -- for -- events / given types
    * if it is unsubscribed OR mux is closed -> subscription's channel is closed 

* TODO: