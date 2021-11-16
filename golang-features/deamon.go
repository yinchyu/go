package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func deamon() {
	//判断当其是否是子进程，当父进程return之后，子进程会被系统1号进程接管
	fmt.Println(os.Getppid())
	if os.Getppid() != 1 {
		// 将命令行参数中执行文件路径转换成可用路径
		filePath, _ := filepath.Abs(os.Args[0])
		cmd := exec.Command(filePath, os.Args[1:]...)
		fmt.Println("守护进程的命令:", os.Args[:])
		// 将其他命令传入生成出的进程
		cmd.Stdin = os.Stdin // 给新进程设置文件描述符，可以重定向到文件中
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		// cmd.Start() // 开始执行新进程，不等待新进程退出
		return

	}
}

func compare(a []int) (max int) {
	for _, v := range a {
		if v > max {
			max = v
		}
	}
	return max
}
func filterwindows() []int {
	datalist := []int{2, 3, 4, 2, 6, 2, 5, 1}
	windowssize := 3
	res := make([]int, 0)
	if len(datalist) < windowssize {
		return res
	}
	start := 0
	for i := windowssize; i <= len(datalist); i++ {
		maxvalue := compare(datalist[start:i])
		start += 1
		res = append(res, maxvalue)
	}
	fmt.Println(res)
	return res
}
func main() {
	filterwindows()
}
