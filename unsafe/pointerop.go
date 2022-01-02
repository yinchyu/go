package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

//1.17 方法
func Slice2Array() {
	var arr [4]int
	slces := []int{12, 43, 56, 7}
	arr = *(*[4]int)(slces)
	fmt.Println(arr)
}

// 通过head的方式
func Slice2Array2() {
	s := []int{1, 2, 3}
	var a [3]int
	// 通过切片进行中转操作
	fmt.Println(copy(a[:2], s))
	fmt.Println(a)
}

// 是string 到 []byte  之间的转换不用经过内存拷贝
func Stringtoslice(s string) []byte {
	//  通过reflect 的header 相当于获取结构体的内部的非导出字段
	str := (*reflect.StringHeader)(unsafe.Pointer(&s))
	ret := reflect.SliceHeader{Data: str.Data, Len: str.Len, Cap: str.Len}
	return *(*[]byte)(unsafe.Pointer(&ret))
}
