package main

import (
	"fmt"
	"sync"
	"time"
)

// 创建timer Pool
// Pool 资源两个gc 会被清除
var timers = sync.Pool{
	New: func() interface{} { //定义创建临时对象创建方法
		t := time.NewTimer(time.Hour)
		t.Stop()
		return t
	},
}

func CallTimmer() {
	timer := timers.Get().(*time.Timer) // 从缓存池中取出对象
	timer.Reset(time.Second)
	<-timer.C
	timers.Put(timer)

}

func RestTime() {
	timer := time.NewTimer(time.Second * 1000)
	// 表示不会有其他的goroutine 往这个上边发送数据
	// 重置肯定是成功的，返回的结果表示 这个计时器是过期了还是没有过期
	// 重置重新设定时间
	reset := timer.Reset(time.Second)
	println(reset)
	fmt.Println(<-timer.C)
}

func main() {

}
