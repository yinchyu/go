package main

import (
	"fmt"
	"reflect"
)

type x []interface{ m() }

func (c x) m() {}

func interfacereflect() {
	tp := reflect.TypeOf(new(any))
	tt := reflect.TypeOf(T{})
	fmt.Println(tp.Kind(), tt.Kind())
	ti, tim := tp.Elem(), tt.Elem()
	fmt.Println(ti.Kind(), tim.Kind())
	fmt.Println(tt.Implements(tim))  // true
	fmt.Println(tp.Implements(tim))  // false
	fmt.Println(tim.Implements(tim)) // true
	// 所有的类型都实现了任何空接口类型。
	fmt.Println(tp.Implements(ti))  // true
	fmt.Println(tt.Implements(ti))  // true
	fmt.Println(tim.Implements(ti)) // true
	fmt.Println(ti.Implements(tim)) // false
}

type F func(string, int) bool

func (f F) m(s string) bool {
	return f(s, 32)
}

func (f F) M() {}

type I interface {
	m(string, int) bool
	M()
}

func funcreflect() {
	var x struct {
		f F
		i I
	}
	tx := reflect.TypeOf(x)
	fmt.Println(tx.Kind())
	for i := 0; i < tx.NumField(); i++ {
		fmt.Printf("%+v\n\n", tx.Field(i))
	}
	tf, ti := tx.Field(0).Type, tx.Field(1).Type
	// 表示是否有可变参数...
	fmt.Println(tf.Kind(), ti.Kind(), tf.IsVariadic(), tf.NumIn(), tf.NumOut(), tf.In(0).Kind(), tf.NumMethod())
	methodByName, bl := tf.MethodByName(tf.Method(0).Name)
	if bl {
		//获取到对应的位置信息
		methodByName.Func.NumMethod()
	}
	fmt.Println()
}

func reconstruct() {
	type T struct {
		// 通过max 和min 限制大小， 数据库orm 模型通过tag 来进行限制
		X    int  `max:"99" min:"12" default:"0"`
		Y, Z bool `optional:"yes"`
	}
	// kind 返回一个具体的类型
	t := reflect.TypeOf(&T{})
	fmt.Println(t.Kind())
	t = reflect.TypeOf(T{})
	fmt.Println(t.Kind())
	x := t.Field(0).Tag
	y := t.Field(1).Tag
	z := t.Field(2).Tag
	fmt.Println(x.Get("max"))
	fmt.Println(x.Lookup("min"))
	fmt.Println(y)
	fmt.Println(z)
	// 返回结构体的字段数
	fmt.Println(t.NumField())
}

func retype() {
	of := reflect.ArrayOf(5, reflect.TypeOf(123))
	fmt.Println(of) //[5]int
	//direction  分为三个方向， 一个是1 一个是2 一个是 3
	// 这个地方也不需要长度来进行操作，因为是具体的类型，通过kind 来进行判断，  就是fmtPrintf("%T"，返回具体的类型)
	chanOf := reflect.ChanOf(reflect.RecvDir, of)
	fmt.Println(chanOf) //<-chan [5]int
	fmt.Println(1 | 2)
	// 返回一个类型的指针
	ptr := reflect.PtrTo(of)
	fmt.Println(ptr) //*[5]int
	// 返回一个类型的切片
	ts := reflect.SliceOf(reflect.TypeOf(23))
	fmt.Println(ts) //[]int
	funcOf := reflect.FuncOf([]reflect.Type{of, chanOf}, []reflect.Type{ts}, false)
	fmt.Println(funcOf) //func([5]int, <-chan [5]int) []int
	// 结构体类型的构建
	structOf := reflect.StructOf([]reflect.StructField{{Name: "Age", Type: reflect.TypeOf("abc")}})
	fmt.Println(structOf.NumField())

}
