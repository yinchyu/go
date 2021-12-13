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
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())

	if err != nil {
		log.Println(err)
	}

	request := &pb.Simplerequest{Add1: 12, Add2: 23, Value: "client"}
	rpclient := pb.NewSimpleClient(conn)
	fmt.Println("request add data: ", request.Add1, request.Add2, "result: ", request.Add2+request.Add1)
	response, err := rpclient.Getinfo(context.Background(), request)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(response.GetValue(), response.GetSum())
}
