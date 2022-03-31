package main

import "fmt"

// 空的类型集合也是第一次听说， 打开了新的思路可以用集合论的方式进行判断
type a interface {
	float32
	int
}

func generics[T a](c T) {
	fmt.Println("c：", c)
	//return c
}

func main() {
	fmt.Println(generics(12))
}
