package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpcserver/pb"
	"log"
)

func main() {
	// WithInsecure returns a DialOption,设置非安全的访问情况
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())

	if err != nil {
		log.Println(err)
	}

	request := &pb.Simplerequest{Add1: 12, Add2: 23, Value: "client"}
	rpclient := pb.NewSimpleClient(conn)
	response, err := rpclient.Getinfo(context.Background(), request)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(response.GetValue(), response.GetSum())
}
