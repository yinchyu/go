package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
fmt.Printf("%.100f\n",3.23)
a:=float64(0.021234)
fmt.Println(strconv.FormatFloat(a,'f',-1,64))
fmt.Println(strconv.FormatFloat(a*100,'f',4,64))
//Reader interface 并不是全部读完了才会返回， 你需要等返回的err 是EOF才能下一步，加了fmt.Println之后相当于增加了等待了时间让网络读取完数据
//	x1 := make([]int, 897)
//	x2 := make([]int, 1024)
//	y := make([]int, 100)
//	// 小于1024 翻倍扩容,大于1024 是1.25倍扩容
//	println(cap(x1))
//	println(cap(x2))
//	println(cap(y))
//	println(cap(append(x1, y...))) // 1360
//	println(cap(append(x2, y...))) // 1536
x := make([]int, 98)
y := make([]int, 666)
println(len(x),cap(x))
println(len(y),cap(y))
println(cap(append(x, y...))) // 768
println(cap(append(y, x...))) // 1360
fmt.Println(	strings.Title("hello world good good study"),strings.ToTitle("hello world good good study"))
	s := []int32{1, 2}
	s = append(s, 3, 4, 5)
	fmt.Printf("len=%d, cap=%d", len(s), cap(s))
}
