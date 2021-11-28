package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func sig(ch chan os.Signal){
	//调用者应该保证c有足够的缓存空间可以跟上期望的信号频率。对使用单一信号用于通知的通道，缓存为1就足够了。

	signal.Notify(ch, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM,syscall.SIGINT)

	for {
		// 阻塞
		s := <-ch
		switch s {
		case syscall.SIGQUIT:
			log.Println("SIGSTOP")

		case syscall.SIGHUP:
			log.Println("SIGHUP")

		case syscall.SIGKILL:
			log.Println("SIGKILL")

		case syscall.SIGINT:
			log.Println("SIGINT")

		default:
			log.Println("default")

		}
	}

}
func callsingal() {
	//ch := make(chan os.Signal, 3)
	//sig(ch)
	//for{
	//	fmt.Println("hello")
	//	time.Sleep(time.Second*1)
	//}
	// 创建一个os.Signal channel
	sigs := make(chan os.Signal, 1)
	//创建一个bool channel
	done := make(chan bool, 1)
	//注册要接收的信号，syscall.SIGINT:接收ctrl+c ,syscall.SIGTERM:程序退出
	//信号没有信号参数表示接收所有的信号
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	//此goroutine为执行阻塞接收信号。一旦有了它，它就会打印出来。
	//然后通知程序可以完成。
	go func() {
		sig := <-sigs
		done <- true
		fmt.Println(sig)
	}()
	//不允许继续往sigs中存入内容
	// 将notify stop 处理
	//signal.Stop(sigs)
// 将监听的信号进行重置，如果没有信号参数就重置所有的信号
	signal.Reset()

	//程序将在此处等待，直到它预期信号（如Goroutine所示）
	//在“done”上发送一个值，然后退出。
	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")


}
