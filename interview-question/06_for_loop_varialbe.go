package main

import (
	"fmt"
	"time"
)

func main() {
	queue := make(chan int, 10)
	go publish(queue)
	serve(queue)
	time.Sleep(10 * time.Millisecond)
}

func publish(queue chan int) {
	for i := 0; i < 10; i++ {
		queue <- i
	}
	close(queue)
}

func serve(queue chan int) {
	for req := range queue {
		req := req // Create new instance of req for the goroutine.
		go func() {
			fmt.Println(req)
		}()
	}
}
