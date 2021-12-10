package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"net"
	"net/http"
	"os"
	"reflect"
	"strings"
	"sync"
	"time"
	"web"
)

func level1(){
	defer  println("defer func3")
	defer func() {
		if err:=recover();err !=nil {
			println("recover")
		}}()
	defer println("defer 2")
level2()
}
func level2(){
	defer println("defer func 4")
	//所有的defer 处理完成之后才会处理panic
	panic("foo")
}
func router1(){
	r:=web.New()
	r.GET("/", func(c *web.Context) {
		c.HTML(http.StatusOK, "<h1>Hello web</h1>")
	})
	r.GET("/hello", func(c *web.Context) {
		// expect /hello?name=geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *web.Context) {
		c.JSON(http.StatusOK, web.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":9999")
}
func readers(){
files,err:=os.Open("D:\\桌面文件夹\\gotest\\test2\\tmp.txt")
	if err != nil {
		log.Println(err)
	}
	data:=make([]byte,1024,1024)
	n,err:=files.Read(data)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("read,data length",n)
	fmt.Println(string(data))
}
func AddElement(slice []int, e int) []int {
	//for i := 0; i < 4; i++ {
	//  slice=append(slice, e)
	//}
	//return slice
	return append(slice, e)
}
func valueclices(){
	var a[]int
	var c []int
	a=append(a, 1,2,3)
	a=append(a, 4)

	b:=AddElement(a,12)
	fmt.Println(len(a),cap(a),len(b),cap(b))
	fmt.Println(a,b)
	fmt.Println(&a[0] == &b[0])
	c=append(c, 1,2,3)
	d:=AddElement(c,12)
	fmt.Println(len(c),cap(c),len(d),cap(d))
	fmt.Println(c,d)
	fmt.Println(&c[0] == &d[0])
}
func fn1(x int) int { return x + 1}
func fn2(x uint8) uint8 { return x}
type HelloInter interface {
	Hello() string
}
type Cat struct {
	//HelloInter
	name string
}

func (c *Cat) Hello() string {
	return c.name + "miaomiao"
}

func router2() {
	r := web.New()
	v1:=r.Group("/v1")
	v1.GET("/", func(c *web.Context) {
		c.HTML(http.StatusOK, "<h1>Hello web</h1>")
	})

	v1.GET("/hello", func(c *web.Context) {
		// expect /hello?name=geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	v1.GET("/hello/:name", func(c *web.Context) {
		// expect /hello/geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	v1.GET("/assets/*filepath", func(c *web.Context) {
		c.JSON(http.StatusOK, web.H{"filepath": c.Param("filepath")})
	})

	r.Run(":9999")
}


func findinit(){
	fset := token.NewFileSet()
	dir, err := parser.ParseDir(fset, ".", func(fileInfo os.FileInfo) bool {
		return strings.HasSuffix(fileInfo.Name(), ".go")
	}, parser.ParseComments)

	if err != nil {
		panic(err)
	}

	for _, pkg := range dir {
		for _, f := range pkg.Files {
			for _, decl := range f.Decls {
				switch t := decl.(type) {
				case *ast.FuncDecl:
					if t.Name.Name == "init" {
						fmt.Printf("find function name:%s, file:%s, pkg:%s\n", t.Name, f.Name, pkg.Name)
					}
				}
			}
		}
	}
}


func writedata(){
	file, err := os.Open("./tmp.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	stringfile,_:=os.OpenFile("a.txt",os.O_WRONLY|os.O_APPEND,os.FileMode(777))
	n,err:=stringfile.Write([]byte(strings.Repeat("goodstudy",34)))
	fmt.Println("wrirte data length: ",n)
	if err != nil {
		log.Println(err)
	}
	defer stringfile.Close()
}

func permutation(S string) []string {
	nameRune := []rune(S)
	var ret []string
	if len(nameRune) == 1 {
		ret = append(ret, string(nameRune))
		return ret
	}
	// 与拼接得到的各个字符串再进行拼接
	for i, s := range nameRune {
		// 差了第i个字符的剩余字符串往下传，并将得到的结果进行合并
		var t []rune
		//  按照每个字起头的进行合并操作， 相当于提取出来第一个字，然后对后续进行操作
		t = append(t, nameRune[:i]...)
		t = append(t, nameRune[i+1:]...)
		res := permutation(string(t))
		for _, r := range res {
			ret = append(ret, fmt.Sprintf("%s%s", string(s), r))
		}
	}
	return ret
}

type student struct{
name string
age int
cl class
}
type class struct {
	number int
	teacher string
}

// new    和var 的区别， var 声明了nil   new 分配了地址
func internelnew(){

	s := student{
		name: "ycy",
		age:  213,
	}
	fmt.Println(s.cl.number)
	fmt.Println(s.cl == class{})

	a:=new(student)
	var b *student

	fmt.Println(a,b)
	// 成功
	a.name="ycy"

	// 失败
	b.name="good"
	}
	func rename(){
		oldpath:="D:\\train_dataset\\new\\C72\\hello.txt"
		newpath:="D:\\train_dataset\\new\\C71\\hello.txt"
		err:=os.Rename(oldpath,newpath)
		if err != nil {
			log.Println(err)
		}
	}

// 反射的使用逻辑
func reflects(){
	s := student{
		name: "ycy",
		age:  213,
	}

	value:=reflect.ValueOf(s)
	typ:=reflect.TypeOf(s)
	fmt.Println(value.Kind(),typ)
}
var sharedRsc = false


const(
	mutexLocked = 1 << iota // mutex is locked
)


// conder  条件变量的使用规范
func  conder(){
	var wg sync.WaitGroup
	wg.Add(2)
	m := sync.Mutex{}
	c := sync.NewCond(&m)
	go func() {
		// this go routine wait for changes to the sharedRsc
		c.L.Lock()
		for sharedRsc == false {
			fmt.Println("goroutine1 wait")
			c.Wait()
		}
		fmt.Println("goroutine1", sharedRsc)
		c.L.Unlock()
		wg.Done()
	}()

	go func() {
		// this go routine wait for changes to the sharedRsc
		c.L.Lock()
		for sharedRsc == false {
			fmt.Println("goroutine2 wait")
			c.Wait()
		}
		fmt.Println("goroutine2", sharedRsc)
		c.L.Unlock()
		wg.Done()
	}()
	// this one writes changes to sharedRsc
	time.Sleep(1 * time.Second)
	c.L.Lock()
	fmt.Println("main goroutine ready")
	sharedRsc = true
	c.Broadcast()
	fmt.Println("main goroutine broadcast")
	c.L.Unlock()
	wg.Wait()
	cond1 := sync.NewCond(new(sync.Mutex))
	cond := *cond1
	fmt.Println(cond)
	fmt.Println("被锁上：",mutexLocked)
}
type Person struct {
	mux sync.Mutex
}

func Reduce(p1 Person) {
	fmt.Println("step...", )
	p1.mux.Lock()
	fmt.Println("hello")
	fmt.Println(p1)
	defer p1.mux.Unlock()
	fmt.Println("over...")
}
func deadlock(){
	var p Person
	p.mux.Lock()
	go Reduce(p)
	p.mux.Unlock()
	fmt.Println(111)
	waitqueue:=make(chan int)
	for {
		waitqueue<-1
	}
}
func main() {
	// 检测一个程序会不会死锁
	fmt.Println("running   not  deadlock")
	server,err:=net.Listen("tcp",":8090")
	if err != nil {
		log.Println(err)
	}
	waitqueue:=make(chan int)
	waitqueue<-1
	for{
		connection,err:=server.Accept()
		if err != nil {
			panic("server")
		}
		fmt.Printf("receive connect from %s\n",connection.RemoteAddr())
		waitqueue<-1
	}


}








