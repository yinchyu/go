package main

import (
	"fmt"

)
func concat( s []string) string{
	length:=0
	for i:=range s{
		length+=len(s[i])
	}
	builder:=strings.Builder{}
	builder.Grow(length)
	for i:=range s{
		builder.WriteString(s[i])
	}
	return builder.String()
}
func main(){

	strs:=concat([]string{"read","book","ready"，"fly"})
	fmt.Println(strs)

}