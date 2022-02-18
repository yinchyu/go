package main

import (
	"fmt"
	"reflect"
)

func Map(data any, fn any) []any {
	vfn := reflect.ValueOf(fn)
	vdata := reflect.ValueOf(data)
	res := make([]any, vdata.Len())
	for i := 0; i < vdata.Len(); i++ {
		// 获得数据
		res[i] = vfn.Call([]reflect.Value{vdata.Index(i)})[0].Interface()
	}

	return res
}

func main() {
	square := func(x int) int {
		return x * x
	}

	nums := []int{1, 2, 3, 4}
	squared_arr := Map(nums, square)
	fmt.Println(squared_arr)
	fmt.Println(Map(5, 5))
}
func Transform(slice, function any) any {
	return transform(slice, function, false)
}
func TransformInplace(slice, function any) any {
	return transform(slice, function, true)
}
func transform(slice, function any, inplace bool) any {
	sliceintype := reflect.ValueOf(slice)
	if sliceintype.Kind() != reflect.Slice {
		panic("transform not slice")
	}
	fn := reflect.ValueOf(function)
	elemtype := sliceintype.Type().Elem()
	if !verifyfuncsignature(fn, elemtype, nil) {
		panic("transform :function must be type func" + sliceintype.Type().Elem().String())

	}
	sliceouttype := sliceintype

	return sliceouttype.Interface()
}

func verifyfuncsignature(fn reflect.Value, types ...reflect.Type) bool {

	if fn.Kind() != reflect.Func {
		return false
	}
	// 限制入参和出参的数量
	if (fn.Type().NumIn() != len(types)-1) || fn.Type().NumOut() != 1 {
		return false
	}
	for i := 0; i < len(types)-1; i++ {
		if fn.Type().In(i) != types[i] {
			return false
		}
	}
	outtype := types[len(types)-1]
	if outtype != nil && fn.Type().Out(0) != outtype {
		return false
	}
	return true
}
