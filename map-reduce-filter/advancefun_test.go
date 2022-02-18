package main

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

var list = []string{"liming", "wanghua", "chen", "zhong"}
var intset = []int{3, 4, 5, 6, 7, 8}

var emloyees = []Emloyee{
	{"Bob", 34, 0, 8000},
	{"alice", 35, 0, 8000},
	{"tom", 45, 0, 3400},
	{"mkie", 50, 8, 5000}}

func TestMapsT2V(t *testing.T) {
	listtarget := []string{"LIMING", "WANGHUA", "CHEN", "z"}
	calcdata := MapsT2V(list, func(s string) string {
		return strings.ToUpper(s)
	})
	if reflect.DeepEqual(listtarget, calcdata) {
		fmt.Println("pass")
	} else {
		err := fmt.Errorf("the answer is wrong want%s ,but have%s \n", listtarget, calcdata)
		fmt.Println(err)
	}
	fmt.Println(MapsT2V(list, func(s string) int {
		return len(s)
	}))
}

func TestReduceT2V(t *testing.T) {
	fmt.Printf("%v\n", ReduceT2V(list, func(s string) int {
		return len(s)
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

func TestFilterT2T(t *testing.T) {
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
}
