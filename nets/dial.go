package nets

import (
	"fmt"
	"log"
	"net"
)

func Dial(){
	laddr:=&net.TCPAddr{Port: 3000}
	raddr:=&net.TCPAddr{Port: 8090}
	conn,err:=net.DialTCP("tcp",laddr,raddr)
	if err != nil {
		log.Println(err)
	}
	//可以设置缓冲区的大小
	conn.SetReadBuffer(1024)
	fmt.Println("dial",conn.LocalAddr(),conn.RemoteAddr())
	buf:=make([]byte,1024)
	n,err:=conn.Read(buf)

	fmt.Println("read data length:",n)
	if err != nil {
		log.Println(n,err)
	}
	fmt.Println("read data:",string(buf[:n]))
}