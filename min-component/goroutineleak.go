package main

import (
	"fmt"
	"runtime"
	"time"
)

func tricker(tick <-chan time.Time, abort chan int) {
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		select {
		case <-tick:
			fmt.Println("gettime")
		case <-abort:
			fmt.Println("aborted!")
			return
		}
	}
}

func main() {
	timer1 := time.NewTimer(3 * time.Second)
	fmt.Println("开始时间：", time.Now().Format("2006-01-02 15:04:05"))
	go func(t *time.Timer) {
		times := 0
		for {
			<-t.C
			fmt.Println("timer", time.Now().Format("2006-01-02 15:04:05"))
			// 从t.C中获取数据，此时time.Timer定时器结束。如果想再次调用定时器，只能通过调用 Reset() 函数来执行
			// Reset 使 t 重新开始计时，（本方法返回后再）等待时间段 d 过去后到期。
			// 如果调用时 t 还在等待中会返回真；如果 t已经到期或者被停止了会返回假。
			times++
			if times > 2 {
				// 停止的时候需要及时的返回，然后才能防止goroutinee 的泄漏
				t.Stop()
				return
				// t.Reset(0 * time.Second)
				fmt.Println("调用 stop 停止定时器")

			} else {
				// 调用 reset 重发数据到chan C
				fmt.Println("调用 reset 重新设置一次timer定时器，并将时间修改为2秒")
				t.Reset(2 * time.Second)
			}
		}
	}(timer1)

	for {
		fmt.Println("nums of goroutine:", runtime.NumGoroutine())
		time.Sleep(2 * time.Second)
	}

}
