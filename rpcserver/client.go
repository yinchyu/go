package main

import (
	"fmt"
	"log"
	"net/rpc"
	"time"
)

type Request struct {
	Num1 int
	Num2 int
}
type Response struct {
	Sum int
}

func main() {
	cli, err := rpc.DialHTTP("tcp", ":1234")
	if err != nil {
		log.Fatalln(err)
	}
	req := Request{Num1: 3, Num2: 4}
	// 不能初始化时直接使用声明， 需要赋值一个具体的值
	res := Response{}
	// err=cli.Call("Server.Service",req,&res)
	reschan := make(chan *rpc.Call, 10)
	call := cli.Go("Server.Service", req, &res, reschan)
	if err != nil {
		log.Fatalln("call", err)
	}
	for {
		select {
		case <-call.Done:
			fmt.Println(res.Sum)
			return
		default:
			fmt.Println("process data ....")
			time.Sleep(time.Second * 1)

		}
	}

}
