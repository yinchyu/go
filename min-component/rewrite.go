package main

import (
	"fmt"
	"os"
	"path"
)
// 自重写程序，这个是实现程序自举的重要的一步
func main() {
	fmt.Printf("%s%c%s%c\n", q, 0x60, q, 0x60)
	fmt.Printf("%c%c\n", 0x60, 0x60)
	getwd, err := os.Getwd()
	if err != nil {
		return
	}
	create, err := os.Create(path.Join(getwd, "./s.go"))
	if err != nil {
		return
	}
	writeString, err := create.WriteString(fmt.Sprintf("%s%c%s%c\n", q, 0x60, q, 0x60))
	if err != nil {
		return
	}
	fmt.Println(writeString)
}

var q = `/* Go quine */
package main

import "fmt"

func main() {
   fmt.Printf("%s%c%s%c\n", q, 0x60, q, 0x60)
}

var q = `
