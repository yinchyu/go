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

func invertslice(args []reflect.Value) []reflect.Value {
	inslice, n := args[0], args[0].Len()
	outslice := reflect.MakeSlice(inslice.Type(), 0, n)
	for i := n - 1; i >= 0; i-- {
		index := inslice.Index(i)

		outslice = reflect.Append(outslice, index)
	}
	return []reflect.Value{outslice}
}

func bind(p any, f func([]reflect.Value) []reflect.Value) {
	//去掉interface 之后就留下func
	// 之前的type 不能直接调用函数， 但是value 可以直接调用函数
	invert := reflect.ValueOf(p).Elem()
	invert.Set(reflect.MakeFunc(invert.Type(), f))
}

func funcbind() {
	var invertints func([]int) []int
	var invertstrs func([]string) []string
	bind(&invertints, invertslice)
	bind(&invertstrs, invertslice)
	fmt.Println(invertints([]int{2, 3, 5}))
	fmt.Println(invertstrs([]string{"go", "c", "rust"}))
}

type T struct {
	A, b int
}

func (t T) AddSubThenScale(n int) (int, int) {
	return n * (t.A + t.b), n * (t.A - t.b)
}

func unexportfield() {
	t := T{
		A: 5,
		b: 2,
	}
	var tt any = t
	vt := reflect.ValueOf(&t)
	name := vt.MethodByName("AddSubThenScale")
	callres := name.Call([]reflect.Value{reflect.ValueOf(3)})
	fmt.Println(callres[0].Int(), callres[1].Int())
	neg := func(a int) int {
		return -a
	}
	vf := reflect.ValueOf(neg)
	// 如果直接是func 就不用使用Elem()
	//vf = vf.Elem()
	call := vf.Call([]reflect.Value{callres[0]})
	fmt.Println(call[0].Int())
	//value obtained using unexported field
	// 非导出字段不能通过这种方式获取
	call2 := vf.Call([]reflect.Value{vt.Elem().FieldByName("A")})
	ttn := reflect.ValueOf(tt)
	fmt.Println("====", ttn.FieldByName("A"))
	fmt.Println(call2[0].Int())
}

func reflectsendchan() {
	strchan := make(chan string, 2)
	vc := reflect.ValueOf(strchan)
	vc.Send(reflect.ValueOf("need"))
	//vc.Send(reflect.ValueOf("need2"))
	// 如果通道已经满了就会返回发送失败， 如果通道没有满，表示发送成功
	send := vc.TrySend(reflect.ValueOf("need3"))
	if send {

	} else {
		fmt.Println("表示数据没有送到")
	}

	// 接收数据
	//recv, ok := vc.Recv()
	if recv, ok := vc.TryRecv(); ok {
		fmt.Println("收到的数据是", recv)
	}
	if recv, ok := vc.TryRecv(); ok {
		fmt.Println("收到的数据是", recv)
	}
	fmt.Println("打印通道的大小:", vc.Len(), vc.Cap())

	if recv, ok := vc.TryRecv(); ok {
		fmt.Println("收到的数据是", recv)
	} else {
		fmt.Println("第三次没有收到数据")
	}

}

func reflectvaluestruct() {
	valueof := reflect.ValueOf
	m := map[string]int{"windows": 12, "unix": 33}
	v := valueof(m)
	// 如果是空的就意味着删除这个对应的key
	v.SetMapIndex(reflect.ValueOf("windows"), reflect.Value{})
	v.SetMapIndex(reflect.ValueOf("unix"), reflect.ValueOf(12))
	fmt.Printf("%#v", m)
}

func reflectselectcase() {
	//branches := []reflect.SelectCase{{
	//	Dir:  0,
	//	Chan: reflect.Value{},
	//	Send: reflect.Value{},
	//},
	//}
	//branches
}

func nilreflect() {
	var z reflect.Value
	v := reflect.ValueOf((*int)(nil)).Elem()
	fmt.Println(v)
	fmt.Println(v == z)
	var i = reflect.ValueOf([]interface{}{nil}).Index(0)
	fmt.Println(i)
	fmt.Println(i.Elem())
	// 关于0值由不同的考虑
	fmt.Println(i.Elem() == z)

}

func convert() {
	vt := reflect.ValueOf(123)
	i := vt.Interface()
	fmt.Printf("%T\n", i)
	//vv := reflect.ValueOf(time.Time{})
	//t := vv.Interface().(time.Time)
	//fmt.Printf("%T\n", t)
	type T struct{ X int }
	t := &T{3}
	of := reflect.ValueOf(t)
	field := of.Elem().Field(0)
	field.IsZero()
	fmt.Println(field)
	fmt.Println(field.Interface())

}

func canconvert() {
	s := reflect.ValueOf([]int{1, 2, 3, 4, 5})
	ts := s.Type()
	t1 := reflect.TypeOf(&[5]int{})
	t2 := reflect.TypeOf(&[6]int{})
	fmt.Println(ts.ConvertibleTo(t1)) // true
	fmt.Println(ts.ConvertibleTo(t2)) // true
	fmt.Println(s.CanConvert(t1))     // true
	fmt.Println(s.CanConvert(t2))     // true
}
func main() {
	//valueofint()
	//valueofstruct()
	//reflectsendchan()
	//nilreflect()
	//convert()
	//var c [5]int
	var d = []int{1, 2, 3, 4, 5}
	var e = []int{1, 2, 3}
	var f = []int{1, 2, 3, 4, 5, 6, 7}
	h := *(*[5]int)(d)
	// 数组的长度不能比切片的长度长
	j := *(*[2]int)(e)
	k := *(*[5]int)(f)
	fmt.Println(h, j, k)
	fmt.Println(len(k), cap(k))

}
