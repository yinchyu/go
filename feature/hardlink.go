package main

import (
	"fmt"
	"log"
	"os"
	"sync"
	"syscall"
	"time"
)

var wg sync.WaitGroup // 创建同步等待组对象

func fileinfo() {
	stat, err := os.Stat("test4/omitempty.go")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(stat.Mode())
	fmt.Println(stat.Sys())
	fmt.Println(stat.IsDir())
	if sysinfo, ok := stat.Sys().(*syscall.Win32FileAttributeData); ok {
		fmt.Printf("%T\n", sysinfo)
		i := sysinfo.CreationTime.Nanoseconds()
		i1 := sysinfo.LastWriteTime.Nanoseconds()
		i2 := sysinfo.LastAccessTime.Nanoseconds()

		fmt.Println(time.Unix(0, i).String())
		fmt.Println(time.Unix(0, i1).String())
		fmt.Println(time.Unix(0, i2).String())
	} else {
		fmt.Println("reader")
	}

}

func remove() {
	err := os.Mkdir("test4/get", os.ModeDir)
	if err != nil {
		log.Println(err)
	}
	dir, err := os.ReadDir("test4")
	if err != nil {
		log.Println(err)
	}
	for i := range dir {
		if !dir[i].IsDir() {
			err := os.Rename("test4\\"+dir[i].Name(), "test4\\get\\"+dir[i].Name())
			if err != nil {
				log.Println(err)
			}
		}
	}

	file, err := os.OpenFile("a.reader", os.O_APPEND, os.ModePerm)
	if err != nil {
		log.Fatal(err)
		permission := os.IsPermission(err)
		fmt.Println(permission)

	}
	fmt.Println(file.Name())

}

func main() {
	/*
	   WaitGroup：同步等待组
	       可以使用Add(),设置等待组中要 执行的子goroutine的数量，

	       在main 函数中，使用wait(),让主程序处于等待状态。直到等待组中子程序执行完毕。解除阻塞

	       子gorotuine对应的函数中。wg.Done()，用于让等待组中的子程序的数量减1
	*/
	//设置等待组中，要执行的goroutine的数量
	//runtime.GOMAXPROCS(2)
	//	wg.Add(2)
	//	go fun1()
	//	go fun2()
	//	//go fun3()
	//
	//	fmt.Println("main进入阻塞状态。。。等待wg中的子goroutine结束。。")
	//
	//	wg.Wait() //表示main goroutine进入等待，意味着阻塞
	//
	//	fmt.Println("main，解除阻塞。。")
	//
	//}
	//func fun1() {
	//	for i := 1; i <= 10; i++ {
	//		fmt.Println("fun1.。。i:", i)
	//	}
	//	wg.Done() //给wg等待中的执行的goroutine数量减1.同Add(-1)
	//}
	//func fun2() {
	//	defer wg.Done()
	//	for j := 1; j <= 10; j++ {
	//		fmt.Println("\tfun2..j,", j)
	//	}
	//	//wg.Done()
	//}
	//func fun3() {
	//	defer wg.Done()
	//	for k := 1; k <= 10; k++ {
	//		fmt.Println("\tfun3..k,", k)
	//	}
	//	//wg.Done()

	//runtime.GOMAXPROCS(0)
	//a := []int{1, 2, 3, 4, 5}
	//// 最后一个是cap 截至的位置， 前边是len 截至的位置
	//b := a[3:4:4]
	//fmt.Println(b[0], len(b), cap(b))
	//fileinfo()
	//remove()
	err := os.Link("test4/get/wait.go", "test4/hardlink.go")
	if err != nil {
		fmt.Println("========", err)
	}
	fmt.Println("*****")
	err = os.Chmod("test4/get/wait.go", 0777)
	if err != nil {
		fmt.Println("========", err)
	}
	err = os.Symlink("test4/get/wait.go", "test4/softlink.go")
	if err != nil {
		fmt.Println("========", err)
	}
	fmt.Println("*****")

}
