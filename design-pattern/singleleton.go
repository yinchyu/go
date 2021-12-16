package main

import (
	"sync"
)

var (
	once sync.Once
	mux  sync.Mutex
)

type sigleton struct {
	num int
}

var singletinstance *sigleton

// 恶汉模式
// var(
// 	singletinstance=&sigleton{12}
// )

// singleton1 单例模式使用once 执行
func singleton1(n int) {
	once.Do(func() {
		singletinstance = &sigleton{n}
	})
}

// singleton2 两次条件判断
func singleton2(n int) *sigleton {
	// 就是最开始的时候不修改的话， 是没有必要加锁的，所以就不用有锁的等待的时间
	// 要是不是空的nil 直接返回就好了， 不用申请锁
	if singletinstance == nil {
		// 多个锁竞争的情况下， 这个不一定是第一个获取到这个资源的
		// 只有第一批需要并发竞争需要申请锁， 之后所有的请求都不用申请锁，主要的作用是减少了锁的申请次数
		mux.Lock()
		// 可以将加锁的时间放到第二次判断nil之前
		if singletinstance == nil {
			singletinstance = &sigleton{n}
		}
		mux.Unlock()
	}
	return singletinstance
}

func main() {
	n := 12
	singleton1(n)
	singleton2(n)
}
