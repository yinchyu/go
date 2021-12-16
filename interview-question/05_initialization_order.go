// https://golang.org/doc/effective_go.html#initialization
// init is called after all the variable declarations in the package have evaluated their initializers,
// and those are evaluated only after all the imported packages have been initialized.
package main

import (
	"fmt"
)

var a int = func() int { main(); return 0 }()
var b int = c + 1
var c int = 100

func init() {
	c = 200
}

func main() {
	fmt.Println(b, c)
}
