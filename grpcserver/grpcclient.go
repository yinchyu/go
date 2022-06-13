package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpcserver/pb"
	"log"
)

func Client() {
	// WithInsecure returns a DialOption,设置非安全的访问情况
	// connection 使用 grpc 进行构造
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())

	if err != nil {
		log.Println(err)
	}

	// 构造请求头部信息
	request := &pb.Simplerequest{Add1: 12, Add2: 23, Value: "client"}
	//构造连接
	rpclient := pb.NewSimpleClient(conn)
	fmt.Println("request add data: ", request.Add1, request.Add2, "result: ", request.Add2+request.Add1)
	response, err := rpclient.Getinfo(context.Background(), request)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(response.GetValue(), response.GetSum())
}
