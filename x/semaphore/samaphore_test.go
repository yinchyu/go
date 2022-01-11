package main

import (
	"context"
	"fmt"
	"log"
	"runtime"
	"testing"
)

func TestNewWeighted(t *testing.T) {
	ctx := context.TODO()

	var (
		maxWorkers = runtime.GOMAXPROCS(0)
		sem        = NewWeighted(int64(maxWorkers))
		out        = make([]int, 32)
	)

	// Compute the output using up to maxWorkers goroutines at a time.
	for i := range out {
		// When maxWorkers goroutines are in flight, Acquire blocks until one of the
		// workers finishes.
		if err := sem.Acquire(ctx, 1); err != nil {
			log.Printf("Failed to acquire semaphore: %v", err)
			break
		}

		go func(i int) {
			defer sem.Release(1)
			out[i] = collatzSteps(i + 1)
		}(i)
	}

	// Acquire all of the tokens to wait for any remaining workers to finish.
	//
	// If you are already waiting for the workers by some other means (such as an
	// errgroup.Group), you can omit this final Acquire call.
	if err := sem.Acquire(ctx, int64(maxWorkers)); err != nil {
		log.Printf("Failed to acquire semaphore: %v", err)
	}

	fmt.Println(out)

}

// collatzSteps computes the number of steps to reach 1 under the Collatz
// conjecture. (See https://en.wikipedia.org/wiki/Collatz_conjecture.)

// 冰雹猜想 通过一个规则进行计算所有的结果都会变成1
func collatzSteps(n int) (steps int) {
	if n <= 0 {
		panic("nonpositive input")
	}

	for ; n > 1; steps++ {
		if steps < 0 {
			panic("too many steps")
		}

		if n%2 == 0 {
			n /= 2
			continue
		}

		const maxInt = int(^uint(0) >> 1)
		if n > (maxInt-1)/3 {
			panic("overflow")
		}
		n = 3*n + 1
	}

	return steps
}
