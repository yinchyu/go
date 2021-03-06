//import "fmt"
//
//var _ = func() int {
//	a = false
//	return 0
//}()
//var a = true
//var b = a
//
//func main() {
//	fmt.Println("========", a, b)
//	if b {
//		panic("FAIL")
//	}
//}
package feature

// 只要是丢弃的就是最后进行初始化
var _ = myInit()

var sp = ""

func flatten(name string, _ ...interface{}) int {
	print(sp, name)
	sp = " "
	return 0
}

var a = flatten("a", x)
var b = flatten("b", y)
var c = flatten("c", z)

// 这个地方false 不会被执行 所以d 不依赖z  d 也可以提前输出
var d = func() int {
	if false {
		_ = z
	}
	return flatten("d")
}()

var e = flatten("e")

var x int
var y int = 42
var z int = func() int { return 42 }()
var _ = *i0
var i0 *int

func myInit() struct{} {
	print(sp + "myInit")
	sp = " "
	i0 = new(int)
	*i0 = 10
	return struct{}{}
}

func main() { println() }
