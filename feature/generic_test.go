package feature

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"testing"
)

func TestCompare(t *testing.T) {

	good()
	fmt.Println(add(12, 34))
	fmt.Println(add(12.5, 34.5))
	fmt.Println(compare(12.5, 34.5))
	fmt.Println(compare(12.5, 34.5))
}

func TestHardlink(t *testing.T) {

	err := os.Link("test4/get/wait.go", "test4/hardlink.go")
	if err != nil {
		fmt.Println("========", err)
	}
	fmt.Println("*****")
	err = os.Chmod("test4/get/wait.go", 0777)
	if err != nil {
		fmt.Println("========", err)
	}
	err = os.Symlink("test4/get/wait.go", "test4/softlink.go")
	if err != nil {
		fmt.Println("========", err)
	}
	fmt.Println("*****")

}

func TestWaitgroup(t *testing.T) {
	/*
	   WaitGroup：同步等待组
	       可以使用Add(),设置等待组中要 执行的子goroutine的数量，

	       在main 函数中，使用wait(),让主程序处于等待状态。直到等待组中子程序执行完毕。解除阻塞

	       子gorotuine对应的函数中。wg.Done()，用于让等待组中的子程序的数量减1
	*/
	//设置等待组中，要执行的goroutine的数量
	runtime.GOMAXPROCS(2)
	wg.Add(2)
	go fun1()
	go fun2()
	//go fun3()

	fmt.Println("main进入阻塞状态。。。等待wg中的子goroutine结束。。")

	wg.Wait() //表示main goroutine进入等待，意味着阻塞

	fmt.Println("main，解除阻塞。。")

}
func fun1() {
	for flatten := 1; flatten <= 10; flatten++ {
		fmt.Println("fun1.。。flatten:", flatten)
	}
	wg.Done() //给wg等待中的执行的goroutine数量减1.同Add(-1)
}
func fun2() {
	defer wg.Done()
	for j := 1; j <= 10; j++ {
		fmt.Println("\tfun2..j,", j)
	}
	//wg.Done()
}
func fun3() {
	defer wg.Done()
	for k := 1; k <= 10; k++ {
		fmt.Println("\tfun3..k,", k)
	}
	//wg.Done()

	runtime.GOMAXPROCS(0)
	a := []int{1, 2, 3, 4, 5}
	// 最后一个是cap 截至的位置， 前边是len 截至的位置
	b := a[3:4:4]
	fmt.Println(b[0], len(b), cap(b))
	fileinfo()
	remove()
}

func TestStructEmbed(t *testing.T) {
	// 如果是embeding 字段就字段的名称就直接是原有的结构体的名字
	label := Label{Widget: Widget{10, 10}, Text: "State:"}

	label.X = 11
	label.Y = 12
	fmt.Println(label.X, label.Widget.X)

	button1 := Button{Label{Widget{10, 70}, "OK"}}
	button2 := Button{Label{Widget{10, 70}, "OK"}}
	//button2 := NewButton(50, 70, "Cancel")
	listBox := ListBox{Widget{10, 40},
		[]string{"AL", "AK", "AZ", "AR"}, 0}

	for _, painter := range []Painter{label, listBox, button1, button2} {
		painter.Paint()
	}

	for _, widget := range []interface{}{label, listBox, button1, button2} {
		widget.(Painter).Paint()
		if clicker, ok := widget.(Clicker); ok {
			clicker.Click()
		}
		fmt.Println() // print a empty line
	}

}

// 如果 omitempty 如果原来本身是空的话就直接忽略，也不会进行输出
// 嵌套的结构体不受这个的限制，必须将嵌套的结构体改为指针类型的才可以
func TestOmitempty(t *testing.T) {
	data := `{
  "street": "200 Larkin St",
  "suite":"",
  "city": "San Francisco",
  "state": "CA",
  "zipcode": "94102"
 }`
	addr := new(address)
	json.Unmarshal([]byte(data), &addr)

	// 处理了一番 addr 变量...

	addressBytes, _ := json.MarshalIndent(addr, "", "    ")
	fmt.Printf("%s\n", string(addressBytes))
}

func TestBlockShelter(t *testing.T) {
	var f1 = func(b bool) {
		fmt.Print("Goat")
	}
	{
		var f1 = func(b bool) {
			fmt.Print("sheep")
			if b {
				fmt.Print(" ")
				// 找到最近的关于f1() 的定义
				f1(!b)
			}
		}
		// 找到最近的定义关于f1()
		f1(true)
	}
}

func TestErrorredeclare(t *testing.T) {
	var v = func(s string) (int, error) {
		n, err := strconv.Atoi(s)
		if err != nil {
			parseBool, err := strconv.ParseBool(s)
			if err != nil {
				return 0, err
			}
			if parseBool {
				n = 1
			}
		}

		return n, err
	}
	fmt.Println(v("true"))
}
