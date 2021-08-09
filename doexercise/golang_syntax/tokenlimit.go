package main

import (
	"fmt"
	"sync"
	"time"
)

type TokenLimit struct {
	speedlimit float64
	bucketsize int
	unsedsize  float64
	mu         sync.Mutex
	lasttime   time.Time
}

func NewTokener(limit float64, brust int) *TokenLimit {
	return &TokenLimit{
		speedlimit: limit,
		bucketsize: brust,
		unsedsize:  0,
		lasttime:   time.Now(),
	}
}

func (tl *TokenLimit) Allow() bool {
	return tl.AllowN(time.Now(), 1)
}

func (tl *TokenLimit) AllowN(reqtime time.Time, n int) bool {
	tl.mu.Lock()
	defer tl.mu.Unlock()
	delta := reqtime.Sub(tl.lasttime).Seconds() * tl.speedlimit
	if delta+tl.unsedsize > float64(tl.bucketsize) {
		tl.unsedsize = float64(tl.bucketsize)
	} else {
		tl.unsedsize = delta + tl.unsedsize
	}
	if tl.unsedsize > float64(n) {
		// 更新大小和请求时间
		tl.unsedsize = tl.unsedsize - float64(n)
		tl.lasttime = reqtime
		return true
	} else {
		return false
	}
}

func main() {
	tl := NewTokener(3, 5)
	time.Sleep(time.Second * 2)
	for i := 0; i < 4; i++ {
		go func(i int) {
			for {
				if tl.Allow() {
					fmt.Println("access:", i)
				} else {
					fmt.Println("forbiden:", i)
				}
				time.Sleep(time.Second * 1)
			}
		}(i)
	}
	time.Sleep(time.Second * 10)
}
