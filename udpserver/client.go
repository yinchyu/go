package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("udp", "127.0.0.1:5000")
	if err != nil {
		fmt.Println("err is ", err)
		// 	 需要判断是什么的错误的类型
	}
	defer conn.Close()
	//  读取基本的文件
	buffer, err2 := os.ReadFile("D:\\桌面文件夹\\gotest\\udpserver\\pyau.ipynb")
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("error1", err)
	}
	fmt.Println(dir)
	if err2 != nil {
		fmt.Println("error2", err2)
	}

	// fmt.Println("是否是一次性读取完毕",buffer,string(buffer))
	for {
		fmt.Println(len(buffer))
		conn.Write(buffer)
		// time.Sleep(time.Second*1)

	}

}
