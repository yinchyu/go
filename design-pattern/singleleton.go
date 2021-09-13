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
	if singletinstance == nil {
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
