package main

import (
	"fmt"
	"strconv"
	"strings"
)
func avg( a string){
newstr:=strings.Split(a,",")
list:=make([]int,len(newstr))
windows:=0
for i :=range newstr{
	val,err:=strconv.Atoi(newstr[i])
	if err!=nil{
		temp:=strings.Split(newstr[i],":")
		windows,_=strconv.Atoi(temp[1])
		val,_=strconv.Atoi(temp[0])
	}
	list[i]=val
}
startvalue:=0
res:=0.0
for i:=0;i<windows;i++{
	startvalue+=list[i]
}
for i:=windows;i<len(list);i++{
	diff:=list[i]-list[i-windows]
	temp:=float64(diff)/float64(startvalue)
	if temp>res{
		res=temp
	}
	startvalue+=diff
}
fmt.Printf("%.2f%%",res*100)
}
func main(){

var a string
fmt.Scan(&a)
// a="5,6,8,26,50,48,52,55,10,1,2,1,20,5:3"
avg(a)


}
