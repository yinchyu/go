package main

import (
	"bytes"
	"fmt"
	"runtime"
	"strconv"
	"sync"
	"time"
)

func extractGID(s []byte) int64 {
	s = s[len("goroutine "):]
	s = s[:bytes.IndexByte(s, ' ')]
	gid, _ := strconv.ParseInt(string(s), 10, 64)
	return gid
}

// Parse the goid from runtime.Stack() output. Slow, but it works.

func getSlow() int64 {
	var buf [64]byte
	return extractGID(buf[:runtime.Stack(buf[:], false)])
}

// Get returns the id of the current goroutine.
func GetGoroutineID() int64 {
	return getSlow()
}

type ReentrantLock struct {
	lock      *sync.Mutex
	cond      *sync.Cond
	recursion int32
	host      int64
}

func NewReentrantLock() sync.Locker {
	res := &ReentrantLock{
		lock:      new(sync.Mutex),
		recursion: 0,
		host:      0,
	}
	res.cond = sync.NewCond(res.lock)
	return res
}

func (rt *ReentrantLock) Lock() {
	id := GetGoroutineID()
	rt.lock.Lock()
	defer rt.lock.Unlock()

	if rt.host == id {
		rt.recursion++
		return
	}

	for rt.recursion != 0 {
		rt.cond.Wait()
	}
	rt.host = id
	rt.recursion = 1
}

func (rt *ReentrantLock) Unlock() {
	rt.lock.Lock()
	defer rt.lock.Unlock()
	// 没有锁可以解锁，或者其他的g 调用解锁
	if rt.recursion == 0 || rt.host != GetGoroutineID() {
		panic(fmt.Sprintf("the wrong call host: (%d); current_id: %d; recursion: %d", rt.host, GetGoroutineID(), rt.recursion))
	}

	rt.recursion--
	// 只有最后的时候才会进行释放操作
	if rt.recursion == 0 {
		rt.cond.Signal()
	}
}

type TestReentrantLock struct {
	mu sync.Locker
	id int64
}

func NewTestReentrantLock() *TestReentrantLock {
	return &TestReentrantLock{
		mu: NewReentrantLock(),
	}
}

func (t *TestReentrantLock) SetID() {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.getID()
	t.id = 1
}

func (t *TestReentrantLock) getID() int64 {
	t.mu.Lock()
	defer t.mu.Unlock()
	return t.id

}

func SingleGoroutine() {
	rt := NewTestReentrantLock()
	rt.SetID()
	fmt.Println(rt.getID())
}

func MuliteGoroutine() {
	rt := NewTestReentrantLock()
	signel := func() {
		rt.SetID()
		fmt.Println(rt.getID())
	}
	for i := 0; i < 5; i++ {
		go signel()
	}
	time.Sleep(10 * time.Second)
}

func main() {
	SingleGoroutine()
	MuliteGoroutine()
}
