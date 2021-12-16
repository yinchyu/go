// https://golang.org/ref/spec#Types
// Each type T has an underlying type:
// If T is one of the predeclared boolean, numeric, or string types, or a type literal, the corresponding underlying type is T itself.
// Otherwise, T's underlying type is the underlying type of the type to which T refers in its type declaration.
package main

import (
	"fmt"
	"reflect"
)

type Enum int

func main1() {
	var a, b interface{}
	a = int(123)
	b = Enum(123)

	underlyingCheckInt(a)
	underlyingCheckInt(b)
}

func underlyingCheckInt(i interface{}) {
	t := reflect.TypeOf(i)
	if t.Kind() == reflect.Int {
		v := reflect.ValueOf(i).Int()
		fmt.Printf("i: %v, v: %d\n", i, v)
	} else {
		fmt.Printf("i: %v, type: %T\n", i, i)
	}
}

//这个是另一个类型， 在 1.18 generic 启用~ 用来表示一个泛类型
type xx int

func main() {
	var a, b interface{}
	a = int(123) // true 123 123
	b = xx(123)  // false   123 main.xx

	// Compile-Time Error? Run-Time Error? Output?
	checkInt(a)
	checkInt(b)
}

func checkInt(i interface{}) {
	iInt, ok := i.(int)
	// 类型断言错误返回断言类型的0 值
	fmt.Printf("iInt: %d, ok: %v\n", iInt, ok)

	switch v := i.(type) {
	case int:
		fmt.Printf("i: %v, v: %d\n", i, v)
	default:
		fmt.Printf("i: %v, type: %T\n", i, i)
	}
}
