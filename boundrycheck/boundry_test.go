package main

import (
	"fmt"
	"testing"
)

func TestRange(z *testing.T) {
	x := []int{2, 3, 5, 7, 11}
	t := x[0]
	var i int
	for index := 0; index < len(x); index++ {
		x[i] = x[index]
		//fmt.Print(i)
		i = index
	}
	x[i] = t
	fmt.Println(x) // [3 5 7 11 2]
}

func TestMake(t *testing.T) {
	data := make([]int, 1<<10)
	fmt.Println(len(data))
}

func BenchmarkName(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {

	}
}
