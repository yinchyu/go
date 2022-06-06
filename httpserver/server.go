package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"reflect"
	"unsafe"
)

// 结构体是指针， 对应的写入操作是对应的接口，通过接口写入数据
// interface 中read 读取的时候是没有扩容的能力，需要通过append
// 一个场景， 申请  p:=make([]byte,12)
//r.read(p[:0]) 返回的结果中n 一定是0  因为设置切片的容量
func handle(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		res.Write([]byte("get请求到的数据"))
	case "POST":
		_, err := io.ReadAll(req.Body)
		if err != nil {
			return
		}
		fmt.Println(res.Header().Values("Content-Type"),
			res.Header().Values("Date"),
			res.Header().Values("Transfer_Encoding"))
		c := "hello"
		// 结构体强制转换的一个重要的条件就是内存的布局一致
		pointer := (*reflect.StringHeader)(unsafe.Pointer(&c))
		d := reflect.SliceHeader{
			pointer.Data,
			pointer.Len,
			pointer.Len,
		}
		// 说明转换的路径，最后转换成功需要一个类型断言，然后解引用
		convert := *(*[]byte)(unsafe.Pointer(&d))
		fmt.Println(len(convert), cap(convert))
		res.Write(convert)
		//fmt.Println(string(data))
		// res.Write([]byte("post 请求到的数据"))
		//res.Write(data)

	}

}

func listenserver1() {
	listener, err := net.Listen("tcp", ":80")
	if err != nil {
		log.Fatal(err)
	}
	// 添加到对应的default server 中
	http.HandleFunc("/", handle)
	http.Serve(listener, nil)
	http.ListenAndServe(":80", nil)
}
func listenserver2() {

	// 添加到对应的default server 中
	http.HandleFunc("/", handle)
	http.ListenAndServe(":80", nil)
}
func main() {
	listenserver2()
}
