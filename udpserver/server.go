package main

import (
	"fmt"
	"log"
	"net"
)

func process(conn *net.UDPConn) {
	// var buf [20]byte
	// 定义20 长度会产生长度错误的问题,直接定义为最大的长度，然后进行数据的读取
	var data = make([]byte, 65536)
	for {
		n, udparr, err := conn.ReadFromUDP(data)
		if err != nil {
			fmt.Println("read from client dailed", err)
			break
		}
		log.Println(n)
		raw := make([]byte, n)
		copy(raw, data[:n])
		fmt.Println("收到的数据来自", *udparr)
		// conn.Write([]byte(recvstr))
	}

}
func main() {
	listen, err := net.ListenUDP("udp",
		&net.UDPAddr{net.IPv4(127, 0, 0, 1),
			5000, ""})
	if err != nil {
		fmt.Println("err is ", err)
	}

	process(listen)

}
