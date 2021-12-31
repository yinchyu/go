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

func Coding(gs []Gopher) {
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
