package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"os"
)

func readcontent(filename string) []byte {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}
	// 返回的时候也不存在对应的大的数据的拷贝
	return data
}

func filechange(filenames []string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	for i := range filenames {
		err = watcher.Add(filenames[i])
		if err != nil {
			log.Fatal(err)
		}
	}

	for {
		select {
		case event := <-watcher.Events:
			// log.Println(event.Op)
			if event.Op&fsnotify.Write == fsnotify.Write {
				// 对一个文件进行了修改操作
				log.Println("modify", event.Name, event.Op)
				data := readcontent(event.Name)
				// 增量的更新，能不能做到， 估计不太行感觉， 通过增量对数据进行更新操作
				fmt.Println(string(data))
			}
		case errs := <-watcher.Errors:
			log.Fatal(errs)
		}
	}
}
func initconfig() {
	viper.SetConfigFile("./config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
	data := viper.Get("offline-num")
	data1 := viper.Get("token-secret")
	fmt.Printf("%v,%T,%v,%T \n", data, data, data1, data1)
	c := 10
	var h interface{}
	h = c
	// 在printf 中接口的打印类型是原本的类型， 不是interface 类型
	fmt.Printf("%v,%T \n", h, h)
	// 10,int

}
func main() {
	// filechange([]string{"D:\\桌面文件夹\\gotest"})
	initconfig()
	// time.Sleep(time.Duration(50+time.Second*1))
}
