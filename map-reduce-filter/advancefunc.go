package main

import (
	"constraints"
	"fmt"
	"strings"
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

func main() {

	var list = []string{"liming", "wanghua", "chen", "zhong"}
	var intset = []int{3, 4, 5, 6, 7, 8}

	var emloyees = []Emloyee{
		{"Bob", 34, 0, 8000},
		{"alice", 35, 0, 8000},
		{"tom", 45, 0, 3400},
		{"mkie", 50, 8, 5000}}
	fmt.Println(MapsT2V(list, func(s string) string {
		return strings.ToUpper(s)
	}))
	fmt.Println(MapsT2V(list, func(s string) int {
		return len(s)
	}))

	fmt.Printf("%v\n", ReduceT2V(list, func(s string) int {
		return len(s)
	}))
	fmt.Printf("%v\n", FilterT2T(intset, func(s int) bool {
		return s%2 == 0
	}))
	// 筛选出员工年龄大于40 的
	fmt.Printf("%v\n", FilterT2T(emloyees, func(s Emloyee) bool {
		return s.Age > 40
	}))
	// 筛选出员工工资大于4000
	fmt.Printf("%v\n", FilterT2T(emloyees, func(s Emloyee) bool {
		return s.salary > 4000
	}))
	// 统计员工年龄大于40 并且工资大于4000的总和
	fmt.Printf("%v\n", ReduceT2V(emloyees, func(s Emloyee) int {
		if s.Age > 40 && s.salary > 4000 {
			return s.salary
		} else {
			return 0
		}

	}))

}
