package main

import (
	"fmt"
	"reflect"
	"time"
	"unicode/utf8"
	"unsafe"
)

var Id int

func Byte2String(s []byte) string {
	// 两个结构再内存模型上一样就可以直接进行类型转换
	rs := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	return *(*string)(unsafe.Pointer(&reflect.StringHeader{Data: rs.Data, Len: rs.Len}))
}

type Person = struct {
	Name    string
	Address *struct {
		Street string
		City   string
	}
}

var data []struct {
	Name    string `json:"name"`
	Address *struct {
		Street string `json:"street"`
		City   string `json:"city"`
	} `json:"address"`
}

var person = ([]Person)(data)

func main() {
	var s = "中国人"
	fmt.Printf("the length of s = %d\n", len(s)) // 9
	for i := 0; i < len(s); i++ {
		fmt.Printf("0x%x ", s[i]) // 0xe4 0xb8 0xad 0xe5 0x9b 0xbd 0xe4 0xba 0xba
	}
	fmt.Println(s)
	fmt.Println(utf8.DecodeRune([]byte(s)))
	for i, i2 := range s {
		fmt.Println(i, string(i2))
	}
	a := 12
	if a > 12 {
		fmt.Println("enter")

	} else {
		fmt.Println("====")
	}
	fmt.Printf("\n")
	var z time.Duration = 10
	fmt.Printf("%v %T\n", time.Second*z, time.Second*z)
}
