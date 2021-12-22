package main

import (
	"fmt"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	timer := time.NewTimer(time.Hour * 1)
	go func() {
		select {
		case t2 := <-timer.C:
			fmt.Println(t2)
		}
		fmt.Println("end")
	}()

	timer.Stop()

	time.Sleep(time.Second * 1)
}
