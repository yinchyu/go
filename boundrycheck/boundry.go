//go:generate go run -gcflags="-d=ssa/check_bce/debug=1" ./boundry.go > ./a.txt
// bounds check elimination
package main

import (
	_ "embed"
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
	//buildinfo.Read(bytes.NewReader([]byte{0x1, 0xdf, 0x9, 0xa, 0x17, 0x1, 0xa, 0x10, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0x1, 0x3f, 0xf2, 0x15, 0xf4, 0xd0, 0x87, 0x2d}))
	//fmt.Println("Hello, 世界")

	//buildInfo, ok := debug.ReadBuildInfo()
}
