package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func main() {
	client := &http.Client{}
	// 获取对应的get 和post 请求操作， 通过   newrequest进行操作，客户端请求的地址是需要一个完整的url 通过url来解析对应的端口地址
	request, err := http.NewRequest("POST", "http://localhost:80", strings.NewReader("name=ycy"))
	if err != nil {
		log.Fatal(err)
	}
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	data, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(response.StatusCode, string(data))

}
