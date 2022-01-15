package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

func main() {
	еrr := errors.New("foo")
	var err error
	if еrr != nil {
		fmt.Println(err == nil)
		fmt.Printf("%T %v", err, err)
	}
	fmt.Println(utf8.DecodeRuneInString("еrr"))
	str := "еrr"
	for len(str) > 0 {
		r, size := utf8.DecodeRuneInString(str)
		fmt.Printf("%c %v\n", r, size)

		str = str[size:]
	}
	str = "err"
	for len(str) > 0 {
		r, size := utf8.DecodeRuneInString(str)
		fmt.Printf("%c %v\n", r, size)

		str = str[size:]
	}
	// 通过码点可以知道一个字符的具体的位置， 有的字符是不可见的有的字符和ascii 肉眼看一样，无法分辨出来差距。
	fmt.Println(utf8.DecodeRune([]byte("е")))

}
