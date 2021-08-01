package main

import (
	"fmt"
	"strings"
)
func main()  {
var a string
fmt.Scan(&a)
newlist:=strings.Split(a,",")
for i:=range newlist{
	temp:=newlist[i]
	newlist[i]=	strings.ToUpper(temp)
}
	fmt.Println(strings.Join(newlist,","))
}
