package main

import (
	"constraints"
)

type Emloyee struct {
	Name     string
	Age      int
	vacation int
	salary   int
}

// 必须使用两个来进行限制,如果使用一个认为是同一个类型
func MapsT2V[T any, V any](t []T, fn func(s T) V) []V {
	var newarray []V
	for i := range t {
		newarray = append(newarray, fn(t[i]))
	}
	return newarray
}

func ReduceT2V[T any, V constraints.Ordered](t []T, fn func(s T) V) V {
	var v V
	for i := range t {
		// 类型相加的话，必须要在类型本省定义加的操作
		v += fn(t[i])
	}
	return v
}

func FilterT2T[T any](t []T, fn func(s T) bool) []T {
	var newarray []T
	for i := range t {
		if fn(t[i]) {
			newarray = append(newarray, t[i])
		}
	}
	return newarray
}

//func main() {
//
//
//}
