package main

import (
	"fmt"
	rate "main"
	// "golang.org/x/time/rate"
	"net/http"
	"time"
)

func main() {
	// 官方提供的限流器操作，  按照每秒生产的速率进行限制
	r := rate.Every(1 * time.Millisecond)
	limit := rate.NewLimiter(r, 10)
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		if limit.Allow() {
			fmt.Printf("请求成功，当前时间：%s\n", time.Now().Format("2006-01-02 15:04:05"))
		} else {
			fmt.Printf("请求成功，但是被限流了。。。\n")
		}
	})

	_ = http.ListenAndServe(":8081", nil)
}
