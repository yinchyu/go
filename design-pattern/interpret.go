package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type Context struct {
	content string

	action string
}
type Interpreter interface {
	Inter(c Context)
}
type dancer struct {

}
type writer struct {

}


func (w writer)Inter(c Context){
	fmt.Println(c.content,c.action)
}
func (w dancer)Inter(c Context){
	fmt.Println(c.content,c.action)
}
func prase(){
	calcstr:="1 + 2 + 3 - 4 + 5 - 6"
	partstr:=strings.Split(calcstr," ")
	symbol:=1
	sum:=0
	for _,part:=range partstr{
		if part== "+"{
			symbol=1
			fmt.Println("进入加法计算","+")
		}else if part=="-"{
			symbol=-1
			fmt.Println("进入减法计算","-")
		}else if num,err:=strconv.Atoi(part);err==nil{
			sum=sum+symbol*num
			fmt.Println("数字元素",num)
		}else{
			fmt.Println("其他的元素")
		}
	}
	fmt.Println(sum)
	for _, r := range "Hello 世界！" {
		// 判断字符是否为汉字
		fmt.Println(r)
		if unicode.Is(unicode.Scripts["Han"], r) {  //或者In都是可以的
			fmt.Printf("%c", r) // 世界
		}
	}
}

func main() {
	cList := []Context{
		{action: "music", content: "高音"},
		{action: "music", content: "低音"},
		{action: "dance", content: "跳跃"},
		{action: "dance", content: "挥手"},
	}
	//对歌舞剧内容进行翻译
	for _, c := range cList {
		if c.action == "music" {
			writer{}.Inter(c)
		} else if c.action == "dance" {
			writer{}.Inter(c)
		}
	}
	// 改变当前工作的路径
	direrr1:=os.Chdir("D:/桌面文件夹/gotest/newhtml")
	if direrr1!=nil{
		log.Println(direrr1)
	}
	direrr:=os.MkdirAll("./ready/go/hello/",777)
	if direrr!=nil{
		log.Println(direrr)
	}
	fd,err:=os.Create("./ready/go/hello/log.txt",)
	if err!=nil{
		log.Println(err)
	}
	fmt.Println(os.Getwd())
	fmt.Println(fd.Name())
}