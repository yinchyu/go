package main

import (
	"fmt"
	"reflect"
)
type animal struct {
}
func (a *animal) Printhello(){
	fmt.Println("get hello ")
}

func getfuncbyname(){
	var a animal
	val:=reflect.ValueOf(&a)
	// 如果通过反射获取对应函数要保证对应的函数是导出类型的函数
	method:=val.MethodByName("Printhello")
	method.Call([]reflect.Value{})
}
func main(){
getfuncbyname()
}
