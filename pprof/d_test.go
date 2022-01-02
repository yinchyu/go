package main

import (
	"net/http"
	"runtime"
	"sync"
	"testing"
)

const url = "https://github.com/EDDYCJY"

func TestAdd(t *testing.T) {
	s := Add(url)
	if s == "" {
		t.Errorf("Test.Add error!")
	}
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(url)
	}
}
func init() {
	runtime.SetMutexProfileFraction(1)
}

func TestMutex(t *testing.T) {
	var m sync.Mutex
	var datas = make(map[int]struct{})
	for i := 0; i < 999; i++ {
		go func(i int) {
			m.Lock()
			defer m.Unlock()
			datas[i] = struct{}{}
		}(i)
	}

	_ = http.ListenAndServe(":6061", nil)
}

//go test -bench=. -cpuprofile=cpu.prof
//go test -bench=. --memprofile=mem.prof
