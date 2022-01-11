// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package errgroup provides synchronization, error propagation, and Context
// cancelation for groups of goroutines working on subtasks of a common task.
package errgroup

import (
	"context"
	"fmt"
	"sync"
)

// A Group is a collection of goroutines working on subtasks that are part of
// the same overall task.
//
// A zero Group is valid and does not cancel on error.
type Group struct {
	cancel func()

	wg sync.WaitGroup

	errOnce sync.Once
	err     error
}

// WithContext returns a new Group and an associated Context derived from ctx.
//
// The derived Context is canceled the first time a function passed to Go
// returns a non-nil error or the first time Wait returns, whichever occurs
// first.
func WithContext(ctx context.Context) (*Group, context.Context) {
	ctx, cancel := context.WithCancel(ctx)
	return &Group{cancel: cancel}, ctx
}

// Wait blocks until all function calls from the Go method have returned, then
// returns the first non-nil error (if any) from them.
// 可能有很多的错误，就会返回第一个错误，通过   sync.once 来进行保证，只会执行一次， 然后如果没有错误的话，cancel
//也会执行
func (g *Group) Wait() error {
	g.wg.Wait()
	// 保证没有错误的时候进行取消

	// cancel 函数一定会被执行，有错误了执行的时间比较早，没有错误执行的时间比较晚
	if g.cancel != nil {
		g.cancel()
	}
	return g.err
}

// Go calls the given function in a new goroutine.
//
// The first call to return a non-nil error cancels the group; its error will be
// returned by Wait.
func (g *Group) Go(f func() error) {
	g.wg.Add(1)

	go func() {
		defer g.wg.Done()

		if err := f(); err != nil {
			fmt.Println("****", err)
			g.errOnce.Do(func() {
				g.err = err
				if g.cancel != nil {

					g.cancel()
				}
			})
		}
	}()
}
