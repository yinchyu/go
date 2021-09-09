package main

import (
	"fmt"
	"reflect"
	"sync"
)

var (
	once sync.Once
)

func single() {
	fmt.Println("执行的次数")
}

func singleleton() {
	var once sync.Once
	fmt.Println(once)
	//  相当于实现了单例模式
	once.Do(single)
	fmt.Println(once)
	once.Do(single)
	// 获取对应没有导出的值， 通过反射获取对应的值， valueof.fieldbyname()
	fmt.Println(reflect.ValueOf(once).FieldByName("done"))
	once.Do(single)
	fmt.Println(reflect.ValueOf(once))
	once.Do(single)
	fmt.Println(reflect.ValueOf(once))
}
func main() {
	singleleton()
}
