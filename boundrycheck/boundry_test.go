package main

import (
	"fmt"
	"testing"
)

func TestRange(z *testing.T) {
	x := []int{2, 3, 5, 7, 11}
	t := x[0]
	var i int
	// 先自增，然后取值，然后赋值
	for index := 0; index < len(x); index++ {
		x[i] = x[index]
		fmt.Print(i)
		i = index
	}
	x[i] = t
	fmt.Println(x) // [3 5 7 11 2]
}
