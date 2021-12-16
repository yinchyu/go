// https://golang.org/ref/spec#Type_assertions
// The notation x.(T) is called a type assertion.
// More precisely, if T is not an interface type, x.(T) asserts that the dynamic type of x is identical to the type T. In this case, T must implement the (interface) type of x; otherwise the type assertion is invalid since it is not possible for x to store a value of type T.
// If T is an interface type, x.(T) asserts that the dynamic type of x implements the interface T.
package main

import (
	"fmt"
)

type I1 interface {
	do1()
}

type I2 interface {
	do2()
}

type A struct{}

func (a A) do1() {
	fmt.Println("a.do1")
}

func (a *A) do2() {
	fmt.Println("a.do2")
}

func main() {
	var a interface{} = new(A)

	// Compile-Time Error? Run-Time Error? Output?
	// 考察指针接收者方法和非指针接收方法的区别
	i1 := a.(I1)
	i1.do1()

	i2 := i1.(I2)
	i2.do2()

	return
}
