package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
)

func handle(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		res.Write([]byte("get请求到的数据"))
	case "POST":
		data, err := io.ReadAll(req.Body)

		if err != nil {
			return
		}
		fmt.Println(string(data))
		// res.Write([]byte("post 请求到的数据"))
		res.Write(data)
	}

}
func main() {
	listener, err := net.Listen("tcp", ":80")
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/", handle)
	http.Serve(listener, nil)
}
