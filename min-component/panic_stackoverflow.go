package main

import "fmt"

type My string

//String 调用栈溢出
func (m My) String() string {
	fmt.Println(m)
	return string(m)
}

func recursion() {
	var s My = "hello"
	fmt.Println(s)
}

// error 调用栈溢出
type Comedyerror string

func (c Comedyerror) Error() string {
	fmt.Println(c)
	return string(c)
}
func recursion2() {
	var err error
	err = Comedyerror("werere")
	fmt.Println(err)
}
