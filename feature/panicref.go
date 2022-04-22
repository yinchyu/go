package feature

import (
	"fmt"
	"log"
	"net"
	"os"
	"runtime"
)

type Myint int

func (receiver Myint) M(a int) Myint {
	print(a)
	return receiver
}
func functionvalue() {
	//在一层包裹的时候直接入栈,后边修改不会影响对应的值
	var f func()
	//被延迟调用的函数值是在被推入延迟调用栈之前被估值的
	defer f()
	f = func() {
		fmt.Println("true")
	}
}

func entryvalue() {

	var t Myint
	// 实参传递会在进入延迟调用栈之前被估值
	defer t.M(2).M(3)
	t.M(4).M(5)

}

func deferfile() {
	openfile, err := os.Open("")
	if err != nil {
		fmt.Println(err)
	}
	defer openfile.Close()
	openfile.Sync()

}
func serverlisten() {
	listen, err := net.Listen("tcp", ":12345")
	defer func() {
		if a := recover(); a != nil {
			log.Println("捕获了一个恐慌：", a)
		}
	}()
	if err != nil {
		log.Println(err)
	}
	for {
		accept, err := listen.Accept()
		if err != nil {
			log.Println(err)
		}
		go clienthandler(accept)
	}
}

func clienthandler(c net.Conn) {
	defer func() {
		if v := recover(); v != nil {
			log.Println("捕获了一个恐慌：", v)
			log.Println("防止了程序崩溃")
		}
		c.Close()
	}()
	panic("未知错误")
}

func NeverExit(name string, f func()) {
	defer func() {
		if v := recover(); v != nil {
			// 侦测到一个恐慌
			log.Printf("协程%s崩溃了，准备重启一个", name)
			go NeverExit(name, f) // 重启一个同功能协程
		}
	}()
	f()
}

func recoverfunc() {
	// 通过panic,recover后对返回值进行赋值操作
	n := func() (result int) {
		defer func() {
			if v := recover(); v != nil {
				if a, ok := v.(int); ok {
					result = a
				}
			}
		}()
		func() {
			panic(123)
		}()
		return 0
	}()
	fmt.Println("返回n的大小", n)
}

func innerdefer() {
	go func() {
		// 一个匿名函数调用。
		// 当它退出完毕时，恐慌2将传播到此新协程的入口
		// 调用中，并且替换掉恐慌0。恐慌2永不会被恢复。
		defer func() {
			// 上一个例子中已经解释过了：恐慌2将替换恐慌1.
			defer panic(2)

			// 当此匿名函数调用退出完毕后，恐慌1将传播到刚
			// 提到的外层匿名函数调用中并与之关联起来。
			func() {
				panic(1)
				// 在恐慌1产生后，此新开辟的协程中将共存
				// 两个未被恢复的恐慌。其中一个（恐慌0）
				// 和此协程的入口函数调用相关联；另一个
				// （恐慌1）和当前这个匿名调用相关联。
			}()
		}()
		panic(0)
	}()
}
func f() {
	defer func() {
		recover()
	}()
	defer panic(12)
	runtime.Goexit()

}
func gosched() {

	c := make(chan struct{})
	go func() {
		defer close(c)
		f()
		for {
			runtime.Gosched()
		}
	}()
	<-c
}
