package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

// 会导致大量的eof,如果一个连接被关闭，一直读取不到数据
//2022/01/15 13:12:16 0 EOF
//192.168.101.13:5679 192.168.101.13:5680

func server(accepter1, accepter2 net.Conn) {
	go func() {
		//for {
		fmt.Println(accepter2.RemoteAddr().String(), "=>", accepter1.RemoteAddr().String())
		i, err := io.Copy(accepter1, accepter2)
		if err != nil {
			log.Println(i, err)
		}

	}()

	fmt.Println(accepter1.RemoteAddr().String(), "=>", accepter2.RemoteAddr().String())
	i, err := io.Copy(accepter2, accepter1)
	if err != nil {
		log.Println(i, err)
	}
}

func main() {
	// 写一个这么简单的服务器发现，如果断掉就会导致对应的存储出现逻辑问题， 一直不停的读取一个字节的数据
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		return
	}

	for {
		accepter1, err := listener.Accept()
		if err != nil {
			return
		}
		fmt.Println(accepter1.RemoteAddr().String())
		accepter2, err := listener.Accept()
		if err != nil {
			return
		}
		fmt.Println(accepter2.RemoteAddr().String())

		go server(accepter1, accepter2)
	}
}
