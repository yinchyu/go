package nets

import (
	"fmt"
	"log"
	"net"
)
func process(conns net.Conn){
	for {
		_,err:=conns.Write([]byte("hello,world"))
		if err != nil {
			log.Println(err)
		}
	}

}
func Connect(){
	addr:=&net.TCPAddr{Port: 8090}
    listener,err:=net.ListenTCP("tcp",addr)
	if err != nil {
		log.Println(err)
	}
	for{
		conn,errs:=listener.Accept()
		if errs != nil {
			log.Println(errs)
		}
		fmt.Println("connect",conn.LocalAddr(),conn.RemoteAddr())
		go process(conn)
	}

}

