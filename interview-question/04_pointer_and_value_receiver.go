// https://golang.org/ref/spec#Method_expressions
package main

import (
	"fmt"
)

type Vertex struct {
	a, b int
}

func (v Vertex) s1() { // value receiver
	v.a = 1
}

func (v Vertex) s2() { // pointer receiver
	v.a = 2
}

func main() {
	v := Vertex{}
	v.s1()
	fmt.Println(v)
	v.s2()
	fmt.Println(v)

	vp := &Vertex{}
	vp.s1()
	fmt.Println(vp)
	vp.s2()
	fmt.Println(vp)
}
