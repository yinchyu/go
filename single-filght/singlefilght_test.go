package single_filght

import (
	"fmt"
	"log"
	"sync"
	"testing"
	"time"
)

func processed(g *Group, num int, ch chan int, key string) {
	for count := 0; count < 10; count++ {
		v, err, shared := g.Do(key, func() (interface{}, error) {
			time.Sleep(800 * time.Millisecond)
			fmt.Println("这个地方就调用了一次， 其他的都是在阻塞等待")
			return "bar", nil
		})
		log.Println("v = ", v, " err = ", err, " shared =", shared, " ch :", ch)
		if got, want := fmt.Sprintf("%v (%T)", v, v), "bar (string)"; got != want {
			log.Println(got, want)
		}

	}
	ch <- num
}
func TestSingleFlight(t *testing.T) {
	group := NewGroup(sync.Mutex{})
	channels := make([]chan int, 10)
	key := "key"
	for i := 0; i < 10; i++ {
		channels[i] = make(chan int)
		go processed(group, i, channels[i], key)
	}

	for i, ch := range channels {
		// 太怪了这个地方， 本来的并发的无需操作，现在变成有序了
		j := <-ch
		fmt.Println("routine ", i, "quit!", j)
	}

}
