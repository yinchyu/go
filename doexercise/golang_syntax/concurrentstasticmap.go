package main

import (
	"fmt"
	"reflect"
	"sync"
)

type muxmap struct {
	 safemap map[rune]int
	sync.RWMutex
}

func countstr(strs []string,concurrency int) map[rune]int{
	if concurrency>100{
		concurrency=100
	}
	worker:=make( chan string,concurrency)
	newmap:=muxmap{make(map[rune]int),sync.RWMutex{}}
	for i:=0;i<concurrency;i++{
		go func() {
			for {
				select {
					case str ,ok:=<-worker:
						if !ok{
							return
						}
						temp:=make(map[rune]int)
						for _,value:=range str{
							temp[value]++
						}
						newmap.Lock()
						for k,v :=range temp{
							newmap.safemap[k]+=v
						}
						// 这个地方不能直接使用defer 来进行操作， defer 来进行操作的话，会导致只会被解锁一次， 导致程序死锁
						newmap.Unlock()
				}
			}
		}()
	}
	for _,str:=range strs{
		worker<-str
	}
	close(worker)

	return newmap.safemap
}


func   main(){
	test:=[]string{"123","anbs","bcd","def","123","anbs","bcd","def","123","anbs","bcd","def","123","anbs","bcd","def","123","anbs","bcd","def","123","anbs","bcd","def","123","anbs","bcd","def","123","anbs","bcd","def"}
	answer:=map[rune]int{}
	newmap:=countstr(test ,100)
	for _,str:=range test{
		for _,value := range str{
			answer[value]++
		}
	}
	fmt.Println(newmap)
	if reflect.DeepEqual(answer,newmap){
		fmt.Println("result is right")
	}else{
		fmt.Println("result is bad")
	}
}

