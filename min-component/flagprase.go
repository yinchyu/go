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





//定义一个类型，用于增加该类型方法
type sliceValue []string
//new一个存放命令行参数值的slice
func newSliceValue(vals []string, p *[]string) *sliceValue {
	*p = vals
	return (*sliceValue)(p)
}
/*
Value接口：
type Value interface {
    String() string
    Set(string) error
}
实现flag包中的Value接口，将命令行接收到的值用,分隔存到slice里
*/
func (s *sliceValue) Set(val string) error {
	*s = sliceValue(strings.Split(val, ","))
	return nil
}
//flag为slice的默认值default is me,和return返回值没有关系
func (s *sliceValue) String() string {
	*s = sliceValue(strings.Split("default is me", ","))
	return "It's none of my business"
}
/*
可执行文件名 -slice="java,go"  最后将输出[java,go]
可执行文件名 最后将输出[default is me]
*/
func main(){
	var varable1 = flag.Int("n", 12, "print newline") // echo -n flag, of type *bool
	var varable2 = flag.String("s","hello","print new line")
	var varable3 = flag.Bool("t",false,"print new line")
	var varable4 string
	flag.StringVar(&varable4,"z","string","print new line")
	flag.Parse() // Scans the arg list and sets up flags
	fmt.Println("newline: ",*varable1)
	fmt.Println("strings: ",*varable2)
	fmt.Println("strings: ",*varable3)
	fmt.Println("strings: ",varable4)
	var languages []string
	flag.Var(newSliceValue([]string{}, &languages), "slice", "I like programming `languages`")
	flag.Parse()
	//打印结果slice接收到的值
	fmt.Println(languages)
	fmt.Println(VariableDefine("hello_world"))
}
