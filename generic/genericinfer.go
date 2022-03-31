package main

import "fmt"

// 有二次传递类型的时候，第一次类型必须确定
type a[T any] interface {
	good() T
}

func gener[T any, c a[T]](g c) {
	fmt.Println(g.good())
}

type myint int

type myfloat float64

func (receiver myint) good() myint {
	fmt.Println("myint", receiver)
	return receiver
}

func (receiver myfloat) good() myfloat {
	fmt.Println("myfloat", receiver)
	return receiver
}

func main() {
	var c myint = 12
	var d myfloat = 12

	gener[myint](c)
	gener[myfloat](d)

}
