package main

import "fmt"

type Gopher interface {
	WriteGoCode()
}

type person struct {
	name string
}

func (p person) WriteGoCode() {
	fmt.Printf("I am %s, i am writing go code!\n", p.name)
}

// 这个地方接受参数是一个gopher go 语言编译器不允许这种转换，o(N)复杂度
// 所以需要自己手动转换,然后再传递参数或者使用泛型的推断进行传递参数
// gs:=make([]gs,len(p))
//for i in range:=gs{
//	gs[i]=p[i]
//}

func Coding(gs []Gopher) {
	for _, g := range gs {
		g.WriteGoCode()
	}
}

func Coding2[T Gopher](gs []T) {
	for _, g := range gs {
		g.WriteGoCode()
	}
}

func main() {
	p := []person{
		{name: "小菜刀1号"},
		{name: "小菜刀2号"},
	}
	Coding(p)
}
