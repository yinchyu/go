package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"path"
	"time"
)

func server() {
	r := gin.Default()
	r.GET("/:name", func(c *gin.Context) {
		name := c.Param("name")
		//拼接路径,如果没有这一步，则默认在当前路径下寻找
		filename := path.Join("./", name)
		//响应一个文件
		c.File(filename)
		return
	})
	//监听端口
	err := r.Run(":10000")
	if err != nil {
		fmt.Println("error")
	}
}

func main() {
	for {
		for i := 0; i < 1000; i++ {
			fmt.Println(i)
			go get()
		}
		time.Sleep(time.Second * 1)
	}

}
