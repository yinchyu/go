package main

import (
	"context"
	"google.golang.org/grpc"
	"grpcserver/pb"
	"log"
	"net"
	"time"
)

type Grpcservice struct {
	pb.UnimplementedSimpleServer
}

func (*Grpcservice) Getinfo(ctx context.Context, req *pb.Simplerequest) (*pb.Simpleresponse, error) {
	newresponse := &pb.Simpleresponse{}
	newresponse.Sum = req.GetAdd1() + req.GetAdd2()
	newresponse.Value = "response to: " + req.GetValue()
	time.Sleep(time.Second * 3)
	log.Println(newresponse.Value)
	return newresponse, nil
	// return nil, status.Errorf(codes.Unimplemented, "method Getinfo not implemented")
}

func (receiver Grpcservice) name() {

}

func Server() {
	//先构建一个tcp 链接
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Println(err)
	}
	// 使用的newserver
	server := grpc.NewServer()
	client := &Grpcservice{}
	// 指针方法必须使用指针进行调用。对服务进行注册
	// 然后使用对应的网络请求
	pb.RegisterSimpleServer(server, client)
	// server 使用tcp 监听
	server.Serve(listener)
	// 服务端需要进行注册操作
	//NewSimpleClient()
}
