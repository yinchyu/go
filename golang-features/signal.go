package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func notify() {
	sig := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	// 表示如果有着两种信号就通知通道
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	// 可以接收终端传递的信号
	go func() {
		sig := <-sig
		fmt.Println(sig)
		done <- true
	}()
	fmt.Println("waiting....")
	<-done
	fmt.Println("done")
}

func main() {
	notify()
}
