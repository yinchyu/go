package main

import (
	"fmt"
	"sync"
)

func main() {
	const size = 10
	mutex := sync.Mutex{}

	m := make(map[int]int)
	for i := 0; i < size; i++ {
		go func() {
			mutex.Lock()
			defer mutex.Unlock()
			m[i]++
		}()
	}
	for j := size; j < size*2; j++ {
		go func(j int) {
			mutex.Lock()
			defer mutex.Unlock()
			m[j]++
		}(j)
	}

	fmt.Println(m)
}
