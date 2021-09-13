package main

import (
	"fmt"
	"time"
)

type funcadd func(int, int) int

//wrapperfunc 显示的定义接口的转换方式
func wrapperfunc(f1 funcadd) funcadd {

	return func(i int, i2 int) int {
		start := time.Now()
		temp := f1(i, i2)
		end := time.Now()
		fmt.Printf(" time used:%v\n", end.Sub(start).Nanoseconds())
		return temp
	}
}

func calc(a, b int) int {
	time.Sleep(time.Second * 1)
	return a + b
}
func main() {
	calc := wrapperfunc(calc)
	fmt.Println(calc(3, 4))
}
