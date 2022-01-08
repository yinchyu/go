package main


import (
	"fmt"
)
type a interface {
	string() string
	count() int
}

type b struct {
	counter int
}

func (b b) string() string {
	return strconv.Itoa(b.counter)
}

func (b b) count() int {
	return b.counter + 1
}

type c struct {
	reader  int
	counter int
}

func (c c) string() string {
	return strconv.Itoa(c.counter)
}

func (c c) count() int {
	return c.counter + 1
}

func (c c) read() int {
	return c.reader + 10
}

func resinter(va a) a {
	return va
}

func main() {

	ins := b{counter: 12}
	ins.count()
	ins.string()
	newb := resinter(ins)
	// 接口断言回来具体的类型也会进行判断是否满足这个类型的所有的方法，或者说其他的属性。
	if newc, ok := newb.(b); ok {
		fmt.Println(newc)
	} else {
		fmt.Println("断言失败")
	}
	//  通过type 调用具体的方法， 然后操作， 传递的第一个参数是该类型的实例化参数
	fmt.Println(b.string(*new(b)))

}
