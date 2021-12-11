package ants

import (
	"errors"
	"math"
	"runtime"
)

const (
	// DEFAULT_ANTS_POOL_SIZE is the default capacity for a default goroutine pool.
	DEFAULT_ANTS_POOL_SIZE = math.MaxInt32

	// DEFAULT_CLEAN_INTERVAL_TIME is the interval time to clean up goroutines.
	// 这个是没有执行到就直接切换对应的网络数据格式
	DEFAULT_CLEAN_INTERVAL_TIME = 1

	// CLOSED represents that the pool is closed.
	CLOSED = 1
)

// 使用了一个default 的pool 来处理数据
var (
	// Error types for the Ants API.
	//---------------------------------------------------------------------------

	// ErrInvalidPoolSize will be returned when setting a negative number as pool capacity.
	// 这个也是有些常用的， 就是在var 中设置error
	ErrInvalidPoolSize = errors.New("invalid size for pool")

	// ErrInvalidPoolExpiry will be returned when setting a negative number as the periodic duration to purge goroutines.
	ErrInvalidPoolExpiry = errors.New("invalid expiry for pool")

	// ErrPoolClosed will be returned when submitting task to a closed pool.
	ErrPoolClosed = errors.New("this pool has been closed")
	//---------------------------------------------------------------------------

	// workerChanCap determines whether the channel of a worker should be a buffered channel
	// to get the best performance. Inspired by fasthttp at https://github.com/valyala/fasthttp/blob/master/workerpool.go#L139
	workerChanCap = func() int {
		// Use blocking workerChan if GOMAXPROCS=1.
		// This immediately switches Serve to WorkerFunc, which results
		// in higher performance (under go1.5 at least).
		if runtime.GOMAXPROCS(0) == 1 {
			return 0
		}

		// Use non-blocking workerChan if GOMAXPROCS>1,
		// since otherwise the Serve caller (Acceptor) may lag accepting
		// new connections if WorkerFunc is CPU-bound.
		return 1
	}()

	defaultAntsPool, _ = NewPool(DEFAULT_ANTS_POOL_SIZE)
)

// Init a instance pool when importing ants.

// Submit submits a task to pool.
// 提交任务
func Submit(task func()) error {
	return defaultAntsPool.Submit(task)
}

// Running returns the number of the currently running goroutines.
// 返回使用的
func Running() int {
	return defaultAntsPool.Running()
}

// Cap returns the capacity of this default pool.
func Cap() int {
	return defaultAntsPool.Cap()
}

// Free returns the available goroutines to work.
func Free() int {
	return defaultAntsPool.Free()
}

// Release Closes the default pool.
func Release() {
	_ = defaultAntsPool.Release()
}
