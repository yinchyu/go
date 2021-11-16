package main

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
	"unicode"
)

func getgoid() int64 {
	buf := make([]byte, 64)
	index := runtime.Stack(buf, false)
	info := string(buf[:index])
	stk := strings.TrimPrefix(info, "goroutine")
	idfiled := strings.Fields(stk)[0]
	id, err := strconv.Atoi(idfiled)
	if err != nil {
		fmt.Println(err)
	}
	return int64(id)
}

func atoi() int {
	str := "      -117c40091539"
	str = strings.ReplaceAll(str, " ", "")
	// fmt.Println(str)
	if len(str) == 0 {
		return 0
	}
	var a int
	for _, value := range str {
		if unicode.IsLetter(value) {
			break
		} else {
			a++
		}
	}
	inter, err := strconv.Atoi(str[:a])
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(inter)
	return inter
}

// 并发的map 操作
var mux sync.Mutex

func multiplaymap(a map[int]string, r chan int) {
	start := <-r
	for i := start; i <= start+10; i++ {
		mux.Lock()
		a[i] = strconv.Itoa(i)
		mux.Unlock()
	}
}

var map2 sync.Map

func multiplaymap2(r chan int) {
	start := <-r
	for i := start; i <= start+10; i++ {
		map2.Store(i, strconv.Itoa(i))

	}
}
func main() {
	// 找到对应函数的栈帧的文字描述
	// a:=getgoid()
	// fmt.Println(a)

	// map1:=make(map[int]string)
	chan1 := make(chan int, 10)
	for i := 0; i < 100; i += 10 {
		chan1 <- i
		go multiplaymap2(chan1)
	}
	time.Sleep(12)
	print := func(key, value interface{}) bool {
		fmt.Println("========", key, value)
		return true
	}
	// for k,v:=range map1{
	// 	fmt.Println(k,v)
	// }
	// fmt.Printf(" map length is %v \n",len(map1))
	map2.Range(print)
}
