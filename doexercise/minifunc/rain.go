package main

import (
	"fmt"
	"strings"
)


func rain( str string) {
flagmap:=make(map[string]bool)
newstr:=strings.Split(str[1:len(str)-1],",")
index:=0
res:=make([]string,len(newstr))
for i:=range newstr{
	// 表示有雨
	if newstr[i]!="0"{
		// 没有办法阻止洪水
		if flagmap[newstr[i]]{
			fmt.Println("[]")
			return
		}else{
			// 标记水库有水
			flagmap[newstr[i]]=true
			res[i]="-1"
		}
	}else{

		if index<i+1{
			index=i+1
		}
		// 向后遍历
		for index<len(newstr){
			if newstr[index]!="0"{
				// 这个问题更新的时候有很多的问题， 直接复制过来导致没有修改内部的索引状态， 导致检查了好长时间
				// 如果不是特别长的变量最好还是自己敲，避免这种低级的问题发生。
				flagmap[newstr[index]]=false
				// 将对应的水库进行标记
				res[i]=newstr[index]
				index++
				break
			}
			index++
		}
	}
	if index<i{
		index=i
	}
}

fmt.Println("["+strings.Join(res,",")+"]")
}

func main(){
var a string
fmt.Scan(&a)
rain(a)


}
