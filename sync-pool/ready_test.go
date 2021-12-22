package main

import (
	"sync"
	"testing"
)

type A struct {
	Name string
}

func (a *A) Reset() {
	a.Name = ""
}

var pool = sync.Pool{
	New: func() interface{} {
		return new(A)
	},
}

func BenchmarkWithoutPool(b *testing.B) {
	var a *A
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10000; j++ {
			a = new(A)
			a.Name = "tink"
		}
	}
}

func BenchmarkWithPool(b *testing.B) {
	var a *A
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10000; j++ {
			a = pool.Get().(*A)
			// 最佳的操作是在get 之后不不是在put 之前进行重置操作， 避免脏数据

			a.Reset()
			a.Name = "tink"
			pool.Put(a) // 一定要记得放回操作，否则退化到每次都需要New操作
		}
	}
}
