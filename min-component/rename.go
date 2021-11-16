package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
	"unsafe"
)
func titledName(who string) string {  return "Mr. " + who }
// 会重新排列这些对应的函数的顺序， 方便执行的结果正常的显示
var bob, smith = titledName("Bob"), titledName("Smith")
func init(){
	fmt.Println("hi",bob)
}
func init(){
	fmt.Println("hello",smith)
}

func write1(){
	//filefd,err:=os.OpenFile("a.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	filefd,err:=os.Create("a.txt")
	if err!=nil{
		log.Println(err)
	}
	filefd.Write([]byte("name space is used  Deprecated"))

}
func write2(){
	filefd,err:=os.OpenFile("b.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC|os.O_SYNC, 0666)
	if err!=nil{
		log.Println(err)
	}
	filefd.Write([]byte("name space is used  Deprecated"))

}
func write3(){
	filefd,err:=os.OpenFile("/dev/fd/1", os.O_RDWR|os.O_CREATE|os.O_TRUNC|os.O_SYNC, 0666)
	if err!=nil{
		log.Println(err)
	}
	filefd.Write([]byte("name space is used"))
}
func test(){
	time.Now().UnixMilli()
	fmt.Println(string('\u68ee'))
	const pi int64=3
	// 常量不被使用也没有问题
	const T int8=^0
	a, b := 123, "Go"
	fmt.Printf("a == %v == 0x%x, b == %s\n", a, a, b)
	fmt.Printf("type of a: %T, type of b: %T\n", a, b)
}

func ListDir(suffix string) []string{
	dirs,err:=os.ReadDir("./")
	files:=make([]string,0)
if err!=nil{
	log.Println(err)
}
	for i,_:=range dirs{
		if dirs[i].IsDir() {
			continue
		}else if strings.HasSuffix(dirs[i].Name(),suffix){
			files=append(files, dirs[i].Name())
		}
	}
	return files
}
func WalkDir(dirpath string) []string{
	entry,err:=os.ReadDir(dirpath)
	if err!=nil{
		log.Println(err)
		return []string{}
	}
	files:=make([]string,0)
	for i,_:=range entry{
		if entry[i].IsDir(){
			files=append(files, WalkDir(path.Join(dirpath,entry[i].Name()))...)
		}
		files=append(files, entry[i].Name())
		fmt.Println(entry[i].Name())
	}
	return files
}

func pathop(){
	{ // example 1
		s := filepath.Dir(`C:\go\bin`)
		println(s == `C:\go`)
	}
	{ // example 2
		s := filepath.Dir("C:/go/bin")
		println(s == `C:\go`)
	}
	{ // example 3
		s := path.Dir("C:/go/bin")
		println(s,s == "C:/go")
	}
	{ // example 4
		s := path.Dir(`C:\go\bin`)
		println(s,s == ".")
	}

	d:=path.Dir("D:\\桌面文件夹\\gotest\\newhtml\\deferfunc.go\\")
	b:=path.Base("D:\\桌面文件夹\\gotest\\newhtml\\deferfunc.go\\")
	fmt.Println(d,b)
	fmt.Println(filepath.Dir("D:\\桌面文件夹\\gotest\\newhtml\\deferfunc.go\\"),
		filepath.Base("D:\\桌面文件夹\\gotest\\newhtml\\deferfunc.go\\"))
	h:=path.Dir("D:/桌面文件夹/gotest/newhtml/deferfunc.go")
	m:=path.Base("D:/桌面文件夹/gotest/newhtml/deferfunc.go")
	fmt.Println(h,m)
	// filepath 确实斜杠和反斜杠都能处理成功， 但是   path 只能处理斜杠， 不能处理反斜杠
	fmt.Println(filepath.Dir("D:/桌面文件夹/gotest/newhtml/deferfunc.go"),
		filepath.Base("D:/桌面文件夹/gotest/newhtml/deferfunc.go"))
	p:="/home/"
	// 避免这种情况也是可以的， 先对路径的内容进行一个清理， 然后继续往后边进行操作
	p=filepath.Clean(p)
	fmt.Println(filepath.Clean(p),filepath.Dir(p), filepath.Base(p))
	fmt.Println(filepath.Join("/home","home"))
	fmt.Println(filepath.Abs("./hello"))
}
func ioop(){
	fmt.Println("main done")
	news:=bytes.Replace([]byte("hello ycy give a good gift",),[]byte("g"),[]byte("y"),-1)
	fmt.Println(string(news))
	// A Buffer can turn a string or a []byte into an io.Reader.
	buf := bytes.NewBufferString("R29waGVycyBydWxlIQ==")
	dec := base64.NewDecoder(base64.StdEncoding, buf)
	io.Copy(os.Stdout, dec)
	fmt.Println("function a")

}

func main() {


//for i:=0;i<5;i++{
//	// 在进入延迟调用栈的时候就确定了
//	defer fmt.Println("first",i)
//}
//for i:=0;i<5;i++{
//	// 推迟到调用的时间
//	defer func() {
//		fmt.Println("second",i)
//	}()
//}
// int 的默认是int64
c:=12
var b int8
fmt.Println(unsafe.Sizeof(c))
fmt.Println(unsafe.Sizeof(b))
	var x complex128 = complex(1, 2) // 1+2i
	var y complex128 = complex(3, 4) // 3+4i
	fmt.Println(x*y)                 // "(-5+10i)"
	fmt.Println(real(x*y))           // "-5"
	fmt.Println(imag(x*y)) // "10"
	p:=float64(32)
	// 类型可以被推断出来， 所以在var 中可以不用写对应的类型
	var a  =&p
	*a++
	*&*a++

	fmt.Println()

	fmt.Println(unsafe.Sizeof(a))
	fmt.Printf("%p,%T\n",a,a)
	os.Mkdir("D:\\桌面文件夹\\gotest\\newhtml\\a",0777)
	os.Mkdir("D:\\桌面文件夹\\gotest\\newhtml\\b",0777)
	err1:=os.Rename("D:\\桌面文件夹\\gotest\\newhtml\\a","D:\\桌面文件夹\\gotest\\newhtml\\c   .txt ")
if err1!=nil{
	log.Println(err1)

}

}
