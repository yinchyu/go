package main

import (
	"net/http"
	_ "net/http/pprof"
)

func main() {
	for i := 0; i < 1000; i++ {
		go func() {
			select {} // 泄漏了 1000 个协程
		}()
		go func() {
			select {} // 泄漏了 1000 个协程
		}()
	}

	// 启动一个 pprof http server
	if err := http.ListenAndServe(":7899", nil); err != nil {
		panic(err.Error())
	}
}
