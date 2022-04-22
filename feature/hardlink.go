package feature

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
