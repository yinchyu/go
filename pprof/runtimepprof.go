package main

import (
	_ "net/http/pprof"
	"os"
	"runtime/trace"
)

func main() {
	file, err := os.Create("./trace.out")
	if err != nil {
		return
	}

	trace.Start(file)
	defer trace.Stop()
	ch := make(chan string)
	go func() {
		ch <- "Go 语言编程之旅"
	}()

	<-ch
}
