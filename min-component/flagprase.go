package main

import (
	"flag"
	"fmt"
	"strings"
	"unicode"
)

func FlagPrase() {
	var name string
	var value string
	flag.StringVar(&name, "name", "go 编程之旅", "帮助信息")
	flag.StringVar(&value, "value", "go 编程之旅", "帮助信息")
	flag.Parse()
	// 获取对应的航迹信息
	newflag := flag.NewFlagSet("go", flag.ExitOnError)
	argment := make([]string, 3)
	err := newflag.Parse(argment)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(argment)

	switch flag.Args()[0] {
	case "go":
		fmt.Println("go")
	case "php":
		fmt.Println("php")
	}
	fmt.Println(name, value)
}

func VariableDefine(s string) string {
	s = strings.ReplaceAll(s, "_", " ")
	s = strings.Title(s)
	s = strings.ReplaceAll(s, " ", "")
	s = string(unicode.ToLower(rune(s[0]))) + s[1:]
	var output []rune
	for i, r := range s {
		fmt.Println(string(r))
		if i == 0 {
			output = append(output, unicode.ToLower(r))
			continue
		}
		if unicode.IsUpper(r) {
			output = append(output, '_')
		}
		output = append(output, unicode.ToLower(r))
	}
	return string(output)
}

func main() {

	fmt.Println(VariableDefine("hello_world"))

}
