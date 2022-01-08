package main

import (
	"fmt"
	"github.com/ycy/pkg"
	"unsafe"
)

type copyommit struct {
	name   string
	age    int
	oldder bool
}

func main() {
	ommit := pkg.New()
	fmt.Println(ommit.Getname(), ommit.Getage())
	c := *(*copyommit)(unsafe.Pointer(ommit))
	fmt.Println(c.name, c.age)

}
