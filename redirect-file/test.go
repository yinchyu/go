package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"time"
)

//time.time()
// 不知道为什么在command 中不能使用     >> 重定向的符号，但是我想到了另一种的解决办法
//, ">>", "log.txt"
func main() {
	cmd := exec.Command("python", "test.py")
	cmd.Dir = "D:\\桌面文件夹\\go\\redirect-file"
	outter, _ := cmd.StdoutPipe()
	err := cmd.Start()
	if err != nil {
		fmt.Println(time.Now(), "start process failed", err)
	}
	fmt.Println(time.Now(), "start process succeed")
	logfile, err := os.Create("D:\\桌面文件夹\\go\\redirect-file\\log.txt")
	defer logfile.Close()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(logfile.Name())
	_, err = io.Copy(logfile, outter)
	if err != nil {
		log.Println(err)
	}
}
