// func main() {
// 	source := []rune("你好世界的是对方的")
// 	target := []rune("是对方的事阶级")
// 	a := levenshtein.DistanceForStrings(source, target, levenshtein.DefaultOptions)
// 	steps := levenshtein.EditScriptForStrings(source, target, levenshtein.DefaultOptions)
// 	fmt.Println(a, steps)
// 	sourceindex := 0
// 	targetindex := 0
// 	// 新建立一个然后用来进行复制操作
// 	tarlist := []rune("")
// 	// levenshtein.EditOperation
// 	// fmt.Println(tarlist,append(tarlist,123))
// 	for _, step := range steps {
// 		// fmt.Println(step,int(step))
// 		if step == 0 {
// 			tarlist = append(tarlist, target[targetindex])
// 			// fmt.Println(string(tarlist))
// 			targetindex += 1
// 		}
// 		if step == 3 {
// 			tarlist = append(tarlist, source[sourceindex])
// 			sourceindex += 1
// 			targetindex += 1
// 		}
// 		if step == 1 {
// 			// fmt.Println(tarlist)
// 			sourceindex += 1
//
// 		}
// 	}
// 	fmt.Println(tarlist, string(tarlist))
// }
// func main()  {
// 	a:=[]int{3,4,5,6}
// 	b:=[]int{}
// 	//  两个切片指向同一个地址如果修改一个就会对另一个造成问题
// 	f:=copy(b,a)
// 	// b[0]=12
// 	fmt.Println(a,b,f)
// }
//
//
//
//
// package main
//
// import (
// "fmt"
// "time"
// )
//
// func main() {
// 	start := time.Now().UnixNano()
// 	//var arr [...]int
// 	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
// 	const NUM int = 100000000
// 	for i := 0; i < NUM; i++ {
// 		//arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
//
// 		bubbleSort(arr)
// 	}
// 	//打印消耗时间
// 	fmt.Println(time.Now().UnixNano() - start)
// }
//
// //排序
// func bubbleSort(arr []int) {
// 	for j := 0; j < len(arr)-1; j++ {
// 		for k := 0; k < len(arr)-1-j; k++ {
// 			if arr[k] < arr[k+1] {
// 				temp := arr[k]
// 				arr[k] = arr[k+1]
// 				arr[k+1] = temp
// 			}
// 		}
// 	}
// }

// package main
//
// func fib(n int64) int64 {
// 	if n <= 2 {
// 		return 1
// 	}
// 	return fib(n-1) + fib(n-2)
// }
//
// func main() {
// 	println(fib(45))
// }

package main

import (
	"fmt"
	"os"
	"time"
)

type E struct {
	a int32
}

func main() {
	st := time.Now()
	const ARRAY_COUNT int64 = 10000
	const TEST_COUNT int64 = ARRAY_COUNT * 100000
	var es [ARRAY_COUNT]*E
	//  空的nil 指针不能进行访问
	//  需要使用地址才能进行赋值
	es[0] = &E{12}
	type k int
	var h [232]*k
	// h:=make([234]*k)
	fmt.Printf("%T,%v\n", h, h)
	fmt.Println(ARRAY_COUNT, TEST_COUNT, es[0].a)
	os.Exit(0)
	for i := int64(0); i < TEST_COUNT; i++ {
		es[i*123456789%ARRAY_COUNT] = &E{a: int32(i)}
	}

	n := int64(0)
	for i := int64(0); i < ARRAY_COUNT; i++ {
		e := es[i]
		if e != nil {
			n += int64(e.a)
		}
	}
	fmt.Println(time.Now().Sub(st))
	println(n)
}
