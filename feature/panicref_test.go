package feature

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestMyint_M(t *testing.T) {
	entryvalue()
	serverlisten()
	go func() {
		for {
			time.Sleep(time.Second)
			fmt.Println("hello")
		}
	}()
	runtime.Goexit()
	panic(0)
	select {}
}
