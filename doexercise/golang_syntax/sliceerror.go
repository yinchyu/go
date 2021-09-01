package main

import (
	"fmt"
)

type student struct {
	Name string
	Age  int
}

func pase_student() map[string]*student {
	// 结构体指针切片的拷贝
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for i := range stus {
		// i引用错误的陷阱，如果传递的是指针，此时最好不要用 range 遍历带来的一个值
		// 直接使用index 索引位置来进行操作
		m[stus[i].Name] = &stus[i]
	}
	return m
}

func main() {
	students := pase_student()
	for k, v := range students {
		fmt.Printf("key=%s,value=%v \n", k, v)
	}
}
