package main

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
	"time"
)

func source(c chan<- int32) {
	ra, rb := rand.Int31(), rand.Int31n(3)
	time.Sleep(time.Duration(rb) * time.Second)
	c <- ra

}
func TestChannel(t *testing.T) {
	c := make(chan int32, 5)
	for i := 0; i < cap(c); i++ {
		go source(c)
	}
	fmt.Println("第一个接收到的通道的值：", <-c)
}

func TestChannelSend(t *testing.T) {
	done := make(chan struct{})
	values := make([]byte, 20*1024*1024)
	_, err := rand.Read(values)
	if err != nil {
		return
	}
	fmt.Println(values[0], values[len(values)-1])
	go func() {
		sort.Slice(values, func(i, j int) bool {
			return values[i] < values[j]
		})
		done <- struct{}{}
	}()
	<-done
	fmt.Println(values[0], values[len(values)-1])
	time.After(time.Second * 1)
	tick := time.Tick(time.Second * 1)
	for {
		fmt.Println(<-tick)
	}
}

type Seat int
type Bar chan Seat

func (b Bar) ServeCustomer(c int) {
	seat := <-b
	fmt.Println("顾客获取一个座位：", c, seat)
	time.Sleep(time.Second * time.Duration(rand.Int31n(2)))
	fmt.Println("顾客离开，回收座位", seat)
	b <- seat
}
func (b Bar) ServeCustomer2(c chan int) {
	for i := range c {
		seat := <-b
		fmt.Printf("%d:顾客获取一个座位: %d\n", i, seat)
		time.Sleep(time.Second * time.Duration(rand.Int31n(2)))
		fmt.Printf("%d:顾客离开，回收座位: %d\n", i, seat)
		b <- seat
	}

}

func TestBar(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	bar24x7 := make(Bar, 10) // 此酒吧有10个座位
	for seatId := 0; seatId < cap(bar24x7); seatId++ {
		bar24x7 <- Seat(seatId) // 均不会阻塞,类型转换

	}
	for customerId := 0; ; customerId++ {
		time.Sleep(time.Second)
		//完成就销毁， 避免销毁
		go bar24x7.ServeCustomer(customerId)
	}
	select {}
}
func TestBar2(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	bar24x7 := make(Bar, 10) // 此酒吧有10个座位
	for seatId := 0; seatId < cap(bar24x7); seatId++ {
		bar24x7 <- Seat(seatId) // 均不会阻塞,类型转换
	}
	consumers := make(chan int, 1)
	for customerId := 0; customerId < cap(bar24x7); customerId++ {
		go bar24x7.ServeCustomer2(consumers)
	}

	for i := 0; ; i++ {
		consumers <- i
	}
	select {}

}

func TestBar3(t *testing.T) {
	words := "words"
	d := make([]int, 1, len(words))
	fmt.Println(d)

}
