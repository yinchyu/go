package main

import (
	"fmt"
	"strconv"
)

type Price int
type ShowPrice interface {
	String() string
}
func (i Price) String() string{
return strconv.Itoa(int(i))
}
func showPriceLIst[T ShowPrice]( s []T) (res string){
	for _,val :=range s{
		fmt.Println(val.String())
	}
	return ""
}
func main()  {
fmt.Println(showPriceLIst([]Price{12,34,45,56}))
}
