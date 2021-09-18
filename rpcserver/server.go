package main

import (
	"log"
	"net"
	"net/http"
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

type Server struct {
	id string
}

func (s *Server) Service(req Request, res *Response) error {
	res.Sum = req.Num2 + req.Num1
	time.Sleep(time.Second * 5)
	return nil
}

func main() {
	rpc.Register(new(Server))
	// 一般方法和和对应的调用进行分离的原因是上边注册的路由或者服务是默认的服务
	// 所以下边调用采用模式的方式就会调用你原来注册的位置
	// 表示使用默认的handlehttp
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	http.Serve(l, nil)
}
