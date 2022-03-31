package main

import (
	"fmt"
	"reflect"
)

func valueofint() {
	var a int
	var p = &a
	of := reflect.ValueOf(p)
	fmt.Println(of.CanSet(), of.CanAddr()) // 直接通过Valueof得到的值都是不能够直接修改
	ofa := reflect.ValueOf(&a)             // 需要对应的地址， 才能进行设置
	ofa = ofa.Elem()
	fmt.Println("int a is can be set?", ofa.CanSet(), ofa.CanAddr()) // 默认的也是不能直接被设置的
	ofv := of.Elem()
	fmt.Println(*p, a)
	fmt.Println(ofv.CanSet(), ofv.CanAddr()) //获取到对应的解引用的后就可以进行修改
	ofv.Set(reflect.ValueOf(12))
	fmt.Println(*p, a)

}

func valueofstruct() {
	var s struct {
		X any
		y any
	}
	vp := reflect.ValueOf(&s)
	indirect := reflect.Indirect(vp)
	fmt.Println(indirect.CanSet(), indirect.CanAddr()) // 就可以被设置对应的函数
	for i := 0; i < indirect.NumField(); i++ {
		if indirect.Field(i).CanSet() {
			// 匿名的字段不能通过反射的方法进行设置, 这和json 中进行反序列化时非导出字段不能进行序列化相同
			indirect.Field(i).Set(reflect.ValueOf(123))
		}

	}
	fmt.Println(s)
	fmt.Println(indirect.Field(0).IsNil(), indirect.Field(1).IsNil()) // 查看对应的字段是否是空

}
func main() {
	//valueofint()
	valueofstruct()
}
