//go:generate go run -gcflags="-d=ssa/check_bce/debug=1" ./boundry.go > ./a.txt
// bounds check elimination
package main

import (
	_ "embed"
	"fmt"
)

func f1(s []int) {
	_ = s[0]
	_ = s[0]
	_ = s[1]
	_ = s[2]
	_ = s[3]
}

//go:embed  go.mod
var s string

func main() {
	fmt.Println(s)
}
