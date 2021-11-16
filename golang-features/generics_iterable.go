package main

import (
	"fmt"
	"strconv"
)
type addable interface {
	type int , float32,float64
}
// 让对象可以进行比较，通过interface 进行约束然后 通过type 填充具体的类型。
// 其中具体的类型都是可以比较的
func compare[T addable](a,b T) T{
	if a>b{
		return a
	}
	return b
}

type S []int
type A [2]int
type P *A
type vector [T any] []T
func printslice[T any]( s []T){
	for _,v :=range s{
		fmt.Println(v)
	}
}
type Price int
type ShowPrice interface {
	Printstr() string
}
func (p Price) Printstr() string{
	return  strconv.Itoa(int(p))
}
// 对类型T 进行限制， 可以通过函数限制， 可以是类型限制，也可以是方法限制
func showpricelist [T ShowPrice]( s []T){
	for _,v :=range s{
		fmt.Println("价格是",v.Printstr())
	}
}


func main() {
	var x = make([]int, 3, 5)
	var y = make([]int, 4, 6)
	x[0],x[1],x[2]=2,3,4
	y[0],y[1],y[2]=2,3,4
	// 可以显式的通过切片的长度对数据进行比较了
	z:=(*[3]int)(x)
    v:=(*[3]int)(y)
	if  *z==*v{
		fmt.Println("equal")
	}else{
		fmt.Println("not equal")
	}
	printslice[int]([]int{1,2,3,4,5,6})
	printslice[int64]([]int64{1,2,3,4,5,6})
	printslice[int32]([]int32{1,2,3,4,5,6})
	printslice[string]([]string{"hello","world","good"})
	printslice(vector[string]{"hello","world","good"})
    // 自动类型推导
    fmt.Printf("%v %T\n",compare(3,4),compare(3,4))
    fmt.Printf("%v %T\n",compare(3.14,4.12),compare(3.14,4.12))
    // 之前不能直接对类型进行限制
    fmt.Println([]Price{12,23,54,56})
    // 不能将一个没有返回值的数据放在println 中
	showpricelist([]Price{12,23,54,56})
}

