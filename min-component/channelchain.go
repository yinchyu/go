package main

import (
	"fmt"
	"sync"
	"time"
)

func main1() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		a := make(chan int, 10)
		b := make(chan int, 10)

		select {
		case <-b:
			fmt.Println("channel chain ")
		case <-a:
			fmt.Println("channel chain ")
		default:
			fmt.Println("default")
		}

	}()
	wg.Wait()
}

func talk(msg string, sleep int) <-chan string {
	ch := make(chan string)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(sleep) * time.Millisecond)
		}
	}()
	return ch
}

func fanIn(input1, input2 <-chan string) <-chan string {
	ch := make(chan string)
	go func() {
		for {
			select {
			case ch <- <-input1:
			case ch <- <-input2:
			}
		}
	}()
	return ch
}
//<-input1 和 <-input2 都会执行，相应的值是：A x 和 B x（其中 x 是 0-5）。但每次 select 只会选择其中一个 case 执行，所以 <-input1 和 <-input2 的结果，必然有一个被丢弃了，也就是不会被写入 ch 中。因此，一共只会输出 5 次，另外 5 次结果丢掉了。（你会发现，输出的 5 次结果中，x 比如是 0 1 2 3 4）
//
//而 main 中循环 10 次，只获得 5 次结果，所以输出 5 次后，报死锁。
func main2() {
	ch := fanIn(talk("A", 10), talk("B", 1000))
	for i := 0; i < 10; i++ {
		fmt.Printf("%q\n", <-ch)
	}
}

func main() {
	ch := make(chan int)
	go func() {
		select {
		case ch <- getVal(1):
			fmt.Println("in first case")
		case ch <- getVal(2):
			fmt.Println("in second case")
		default:
			fmt.Println("default")
		}
	}()

	fmt.Println("The val:", <-ch)
}

func getVal(i int) int {
	fmt.Println("getVal, i=", i)
	return i
}