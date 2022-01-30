package main

import (
	"fmt"
)
func add[T ~int | ~float64](a, b T) T {
	return a + b
}

// 在闭包中导入变量会提示未使用的变量
func good() {
	func() {
		x := 2
		fmt.Println(x)
	}()
}


func main(){


good()
fmt.Println(add(12,34))
fmt.Println(add(12.5,34.5))
}