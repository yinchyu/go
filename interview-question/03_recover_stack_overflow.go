package main

import (
	"fmt"
	"time"
)

var count int

func main() {
	go func() {
		for {
			worker()
		}
	}()

	t := time.NewTicker(time.Second)
	for {
		select {
		case <-t.C:
			fmt.Println("still alive")
		}
	}
	return
}

func worker() {
	defer func() {
		if err := recover(); err != nil {
			count++
			if count%1e5 == 0 {
				fmt.Printf("err: %v, count: %v\n", err, count)
				return
			}
		}
	}()

	dangerousWork()
}

func dangerousWork() {
	panic("fail")
}
