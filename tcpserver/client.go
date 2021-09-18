package main

import (
	"fmt"
	"net"
	"time"
)

// 做到了上线通知和下线通知的操作
func errorhandler(errortype string, err error) {
	if err != nil {
		fmt.Println(errortype, err)
	}
}
func connserver() {
	conn, err := net.Dial("tcp", ":80")
	errorhandler("listen error", err)
	// data:=make([]byte,100)
	counter := 0
	for {
		conn.Write([]byte("hello"))
		fmt.Println("writedata")
		time.Sleep(time.Second * 1)
		if counter%2 == 0 {
			conn.Write([]byte("\n"))
			counter = (counter + 1) % 15000
		}
		// n,_:=conn.Read(data)
		// fmt.Println(string(data[:n]))
	}
}
func main() {
	connserver()
}
