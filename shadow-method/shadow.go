package main

import "fmt"

type People struct{}

func (p People) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}

func (p *People) ShowB() {
	fmt.Println("showB")
}

type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}

func main() {
	t := Teacher{}
	t.ShowA()
		// 关于继承方法， teacher 生成了 showA 方法，但是传递的参数还是people 所以调用的时候还是调用的  b 的showB的方法
	// 生成的包装方法如下
	//func (t *Teacher) ShowA() {
	//	(*People).ShowA(&(t.People))
	//}
	//showA
	//showB
}
