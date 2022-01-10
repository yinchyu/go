package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	_ "net/http/pprof"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(rep http.ResponseWriter, req *http.Request) {
	var a int
	for i := 0; i < 4; i++ {
		a += queryAll()
	}
	rep.Write([]byte(fmt.Sprintf("%d", a)))
}
func queryAll() int {
	// chanel 没有缓冲的话就会导致阻塞
	ch := make(chan int, 3)
	for i := 0; i < 3; i++ {
		go func() { ch <- query() }()
	}
	return <-ch
}

func query() int {
	n := rand.Intn(100)
	time.Sleep(time.Duration(n) * time.Millisecond)
	return n
}
