* `var Stdin`
  * == stdin line reader
    * uses
      * | input
        * Reason: 🧠it keeps an internal buffer🧠

* `type UserPrompter interface {}`
  * console's methods / 
    * prompt the user / various types of inputs
  * TODO: 

* `type WordCompleter func(line string, pos int) (string, []string, string)`
  * 's input
    * `line string`
      * currently edited line
    * `pos int`
      * cursor position
  * 's return
    * completion candidates -- for the -- partial word
    * _Example:_  TODO: pass to .go
      * `WordCompleter("Hello, wo!!!", 9)` -> may return `("Hello, ", {"world","Word"}, "!!!")`
        * Reason: 🧠have "Hello, world!!!"🧠

* TODO:
