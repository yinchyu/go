package main

import (
	"context"
	"fmt"
	"math/rand"
	"runtime"
	"sort"
	"strconv"
	"time"
)

type job struct {
	jobid     int
	jobnumber int
}
type result struct {
	job
	sum int
}

func pool(nums int, parajob chan job, sumreslut chan result, ctx context.Context) {
	for i := 0; i < nums; i++ {
		go func(i int) {
			stop := true
			for stop {
				select {
				case parajob := <-parajob:
					sum := 0
					strnum := strconv.Itoa(parajob.jobnumber)
					for indx := range strnum {
						w, _ := strconv.Atoi(string(strnum[indx]))
						sum += w
					}
					sumreslut <- result{parajob, sum}
					// 可以进行多个进程的控制操作
				case read := <-ctx.Done():
					//struct {}{} struct {}
					// runtime.LockOSThread()
					fmt.Println("goroutine:", i, read, "receive the stop singal")
					fmt.Printf("%#v %T\n", read, read)
					stop = false
				}
			}
		}(i)
	}
}
func canceltest() {
	jobchan := make(chan job, 128)
	reschan := make(chan result, 128)
	// 设置context 上下文管理器
	ctx, cancal := context.WithCancel(context.Background())
	pool(10, jobchan, reschan, ctx)
	go func() {
		for i := 0; ; i++ {
			jobchan <- job{i, rand.Int()}
		}
	}()

	go func() {
		for value := range reschan {
			fmt.Println(value)
		}
	}()
	fmt.Println("gorountine nunmbers++++++++++++++", runtime.NumGoroutine())
	time.Sleep(time.Millisecond * 10)
	// 所有的goroutine都取消掉
	cancal()
	time.Sleep(time.Second * 1)
	fmt.Println("gorountine nunmbers+++++++++++", runtime.NumGoroutine())
	// 发送一个终止信号，进行启停goroutine
}

// Code generated by "stringer -type=Pill"; DO NOT EDIT.
type Pill int

const _Pill_name = "PlaceboAspirinIbuprofenParacetamol"

var _Pill_index = [...]uint8{0, 7, 14, 23, 34}

func (i Pill) String() string {
	if i < 0 || i >= Pill(len(_Pill_index)-1) {
		return fmt.Sprintf("Pill(%d)", i)
	}
	return _Pill_name[_Pill_index[i]:_Pill_index[i+1]]
}

// func subsetsWithDup(nums []int) (ans [][]int) {
// 	sort.Ints(nums)
// 	n := len(nums)
// outer:
// 	for mask := 0; mask < 1<<n; mask++ {
// 		t := []int{}
// 		for i, v := range nums {
// 			if mask>>i&1 > 0 {
// 				if i > 0 && mask>>(i-1)&1 == 0 && v == nums[i-1] {
// 					continue outer
// 				}
// 				t = append(t, v)
// 			}
// 		}
// 		ans = append(ans, append([]int(nil), t...))
// 	}
// 	return
// }
func subsetsWithDup(nums []int) [][]int {
	res := [][]int{}
	n := len(nums)
outer:
	for i := 0; i < 1<<n; i++ {
		temp := []int{}
		for index1, value1 := range nums {
			if i>>index1&1 == 1 {
				if index1 > 0 && i>>(index1-1) == 0 && nums[index1-1] == nums[index1] {
					continue outer
				}
				temp = append(temp, value1)
			}
		}
		res = append(res, temp)

	}
	return res
}

func permuteUnique(nums []int) [][]int {
	res := [][]int{}
	n := len(nums)
	visit := make([]bool, n)
	sort.Ints(nums)
	var dfs func([]int)
	dfs = func(temp []int) {
		if len(temp) == n {
			res = append(res, append([]int(nil), temp...))
			return
		}
		for index, value := range nums {
			if index > 0 && nums[index] == nums[index-1] && !visit[index-1] || visit[index] {
				continue
			}
			visit[index] = true
			temp = append(temp, value)
			dfs(temp)
			temp = temp[:len(temp)-1]
			visit[index] = false
		}
	}
	dfs([]int{})
	return res
}
func permute(nums []int) [][]int {
	res := [][]int{}
	n := len(nums)
	var dfs func([]int, []int)
	dfs = func(temp []int, nums []int) {
		if len(temp) == n {
			res = append(res, append([]int(nil), temp...))
		}
		for i, val := range nums {
			temp = append(temp, val)
			dfs(temp, append(append([]int(nil), nums[:i]...), nums[i+1:]...))
			temp = temp[:len(temp)-1]
		}

	}
	dfs([]int{}, nums)
	return res
}
func main() {
	// fmt.Println(subsetsWithDup([]int{1,2,2}))
	fmt.Println(permuteUnique([]int{1, 1, 2}))
	// fmt.Println(permute([]int{1,1,2}))
	runtime.GOMAXPROCS(12)
	// 没有m 了就会创建， 如果有m 空闲， 就会回收或者睡眠
	// working stealing 机制， 先从本地队列上获取， 然后从全局队列上边获取
	// hand off 机制，
	a := []int{2, 3, 4, 10}
	fmt.Println(append(a[:1], a[2:]...))
	fmt.Println(a)
	// [2 4 10]
	// 	[2 4 10 10]
	fmt.Println(a[2:], len(a[2:]), cap(a[2:]))
	fmt.Printf("%p %v %p \n", append(a[:1], a[2:]...), append(a[:1], a[2:]...), a)
}

// func main() {
// 	jobchan := make(chan job, 128)
// 	reschan := make(chan result, 128)
// 	// 设置context 上下文管理器
// 	key,value:=12,"str12"
// 	ctx:= context.WithValue(context.Background(),key,value)
// 	// 可以设置不同的时间日期
// 	// ctx, cancel := context.WithDeadline(context.Background(),time.Now().a)
// 	pool(10, jobchan, reschan, ctx)
// 	go func() {
// 		for i := 0; ; i++ {
// 			jobchan <- job{i, rand.Int()}
// 		}
// 	}()
//
// 	go func() {
// 		for value := range reschan {
// 			fmt.Println(value)
// 		}
// 	}()
// 	// 使用timer 可以手动停止，也可以使用timer本身的停止器停止
// 	// cancel()
// 	fmt.Println("gorountine nunmbers++++++++++++++", runtime.NumGoroutine())
// 	// 默认的话输出nil
// 	fmt.Println(ctx.Done())
// 	// time.Sleep(time.Second * 3)
// 	fmt.Println(ctx.Value(12))
// 	fmt.Println("gorountine nunmbers+++++++++++", runtime.NumGoroutine())
// 	// 发送一个终止信号，进行启停goroutine
// 	a:= struct {
// 	}{}
// // sync.RWMutex{}
//
// fmt.Println(unsafe.Sizeof(a))
// 	o := &sync.Once{}
// 	for i := 0; i < 10; i++ {
// 		o.Do(func() {
// 			fmt.Println("only once")
// 		})
// 	}
//
// // c:=time.NewTimer()
// 	// c := sync.NewCond(&sync.Mutex{})
// 	// c.Broadcast()
// 	n:=Pill(2)
// fmt.Println(n.String())
//
// }
