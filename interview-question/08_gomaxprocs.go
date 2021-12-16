package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	const size = 10
	runtime.GOMAXPROCS(1) // the number of operating system threads allocated to goroutines in your program

	m := make(map[int]int)
	for i := 0; i < size; i++ {
		go func() { m[i]++; fmt.Println(i) }()
	}
	for j := size; j < size*2; j++ {
		go func(j int) { m[j]++; fmt.Println(j) }(j)
	}

	fmt.Println(m)
	time.Sleep(10 * time.Millisecond)
}
