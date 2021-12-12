package ants

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

const (
	_   = 1 << (10 * iota)
	KiB // 1024
	MiB // 1048576
	GiB // 1073741824
	TiB // 1099511627776             (超过了int32的范围)
	PiB // 1125899906842624
	EiB // 1152921504606846976
	ZiB // 1180591620717411303424    (超过了int64的范围)
	YiB // 1208925819614629174706176
)

const (
	Param    = 100
	AntsSize = 1000
	TestSize = 10000
	n        = 100000
)

var curMem uint64

func demoFunc() {
	fmt.Println("demo func...")
	time.Sleep(time.Second * 1)
}
func demoPoolFunc(args interface{}) {
	// 总结了一个道理， 就是最开始的时候需要一直的往前走， 不要太在意细节
	n := args.(int)
	time.Sleep(time.Duration(n) * time.Millisecond)
}

// TestAntsPoolWaitToGetWorker is used to test waiting to get worker.
func TestAntsPoolWaitToGetWorker(t *testing.T) {
	var wg sync.WaitGroup
	p, _ := NewPool(AntsSize)
	defer p.Release()

	for i := 0; i < n; i++ {
		//在另外的栈空间上有wg变量的存在
		wg.Add(1)
		p.Submit(func() {
			demoPoolFunc(Param)
			wg.Done()
		})
	}
	wg.Wait()
	t.Logf("pool, running workers number:%d", p.Running())
	//运行时的监控指标
	mem := runtime.MemStats{}
	runtime.ReadMemStats(&mem)
	curMem = mem.TotalAlloc/MiB - curMem
	t.Logf("memory usage:%d MB", curMem)
}

// TestAntsPoolGetWorkerFromCache is used to test getting worker from sync.Pool.
func TestAntsPoolGetWorkerFromCache(t *testing.T) {
	p, _ := NewPool(TestSize)
	defer p.Release()

	for i := 0; i < AntsSize; i++ {
		p.Submit(demoFunc)
	}
	time.Sleep(2 * DEFAULT_CLEAN_INTERVAL_TIME * time.Second)
	p.Submit(demoFunc)
	t.Logf("pool, running workers number:%d", p.Running())
	mem := runtime.MemStats{}
	runtime.ReadMemStats(&mem)
	curMem = mem.TotalAlloc/MiB - curMem
	t.Logf("memory usage:%d MB", curMem)
}

func TestNoPool(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			demoFunc()
			wg.Done()
		}()
	}

	wg.Wait()
	mem := runtime.MemStats{}
	runtime.ReadMemStats(&mem)
	curMem = mem.TotalAlloc/MiB - curMem
	t.Logf("memory usage:%d MB", curMem)
}

func TestAntsPool(t *testing.T) {
	defer Release()
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		Submit(func() {
			demoFunc()
			wg.Done()
		})
	}
	wg.Wait()

	t.Logf("pool, capacity:%d", Cap())
	t.Logf("pool, running workers number:%d", Running())
	t.Logf("pool, free workers number:%d", Free())

	mem := runtime.MemStats{}
	runtime.ReadMemStats(&mem)
	curMem = mem.TotalAlloc/MiB - curMem
	t.Logf("memory usage:%d MB", curMem)
}

func TestPanicHandler(t *testing.T) {
	p0, err := NewPool(10)
	if err != nil {
		t.Fatalf("create new pool failed: %s", err.Error())
	}
	defer p0.Release()
	var panicCounter int64
	var wg sync.WaitGroup
	p0.PanicHandler = func(p interface{}) {
		defer wg.Done()
		atomic.AddInt64(&panicCounter, 1)
		t.Logf("catch panic with PanicHandler: %v", p)
	}
	wg.Add(1)
	p0.Submit(func() {
		panic("Oops!")
	})
	wg.Wait()
	c := atomic.LoadInt64(&panicCounter)
	if c != 1 {
		t.Errorf("panic handler didn't work, panicCounter: %d", c)
	}
	if p0.Running() != 0 {
		t.Errorf("pool should be empty after panic")
	}

}

func TestPoolPanicWithoutHandler(t *testing.T) {
	p0, err := NewPool(10)
	if err != nil {
		t.Fatalf("create new pool failed: %s", err.Error())
	}
	defer p0.Release()
	p0.Submit(func() {
		panic("Oops!")
	})

}

func TestPurge(t *testing.T) {
	p, err := NewPool(10)
	defer p.Release()
	if err != nil {
		t.Fatalf("create TimingPool failed: %s", err.Error())
	}
	p.Submit(demoFunc)
	time.Sleep(3 * DEFAULT_CLEAN_INTERVAL_TIME * time.Second)
	if p.Running() != 0 {
		t.Error("all p should be purged")
	}

	time.Sleep(3 * DEFAULT_CLEAN_INTERVAL_TIME * time.Second)
	if p.Running() != 0 {
		t.Error("all p should be purged")
	}
}

func TestRestCodeCoverage(t *testing.T) {
	_, err := NewTimingPool(-1, -1)
	t.Log(err)
	_, err = NewTimingPool(1, -1)
	t.Log(err)

	p0, _ := NewPool(TestSize)
	defer p0.Submit(demoFunc)
	defer p0.Release()
	for i := 0; i < n; i++ {
		p0.Submit(demoFunc)
	}
	t.Logf("pool, capacity:%d", p0.Cap())
	t.Logf("pool, running workers number:%d", p0.Running())
	t.Logf("pool, free workers number:%d", p0.Free())
	p0.Tune(TestSize)
	p0.Tune(TestSize / 10)
	t.Logf("pool, after tuning capacity, capacity:%d, running:%d", p0.Cap(), p0.Running())
}
