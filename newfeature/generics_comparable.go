package main

import "fmt"
// comparable 是编译器自动梳理的， 包含内部的所有的可以比较的类型
func compared[T comparable](s []T, x T) int{
	for i:=range s{
		if s[i]==x{
			return  i
		}
	}
	return -1
}

func main(){
	// 两个对象可以直接的进行比较操作， 但是函数最开始的时候不能确定对应的类型
	fmt.Println(compared([]string{"hello","world","good"},"good"))
	fmt.Println(compared([]int64{1,2,3,4,5,6},5))
}