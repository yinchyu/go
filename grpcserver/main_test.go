package main

import (
	"fmt"
	"testing"
)

func TestServer(t *testing.T) {
	fmt.Println("server")
	Server()
}

func TestClient(t *testing.T) {
	fmt.Println("client")

	Client()
}

func TestGrpcservice_Getinfo(t *testing.T) {
	go Server()
	for i := 0; i < 3; i++ {
		Client()
	}
}
