package main

import (
	"fmt"
	"reflect"
)

func main() {
	var i int64 = 10
	v := reflect.ValueOf(i)
	fmt.Println("settablility of v:", v.CanSet())

	v = reflect.ValueOf(&i).Elem()
	fmt.Println("settablility of v:", v.CanSet())
	v.SetInt(8)
	fmt.Println(i)
}
