* TODO:
* `type TypeMux struct {}`
  * dispatches events -- to -- registered receivers / handle events
    * Reason of register receivers: 🧠handle events🧠
  * operation / called AFTER mux
    * is stopped
    * will return `ErrMuxClosed`
  * `0`
    * ready to use
  * ⚠️Deprecated⚠️
    * use [Feed](feed.go)

* TODO: