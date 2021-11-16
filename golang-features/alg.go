package main

import (
	"container/list"
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
	"io/ioutil"
	"math"
	"net/http"
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

//go:embed a.txt
var version []byte

//go:embed a.json
var version2 []byte

//go:embed dir
var filesystem embed.FS

func reversePrint(head *ListNode) []int {
	res := make([]int, 0)
	start := head
	if start == nil {
		return []int{}
	}
	for {
		res = append([]int{start.Val}, res...)
		if start.Next != nil {
			start = start.Next
		} else {
			break
		}
	}
	return res
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func resulive(left []int, right []int, currnode *TreeNode) {
	if len(left) == 0 {
		currnode.Left = nil
		return
	}
	if len(right) == 0 {
		currnode.Right = nil
		return
	}
	if len(left) == 1 {
		currnode.Left = &TreeNode{left[0], nil, nil}
		return
	}
	if len(right) == 1 {
		currnode.Right = &TreeNode{right[0], nil, nil}
		return
	} else {
		index := search(left[0], right)
		resulive(right[:index], right[index+1:], currnode)
	}

}
func search(key int, inorder []int) int {
	ret := 0
	for index, value := range inorder {
		if value == key {
			ret = index
		}
	}
	fmt.Println(ret)
	fmt.Println(inorder[:ret], inorder[ret+1:])
	return ret
}
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	currentnode := TreeNode{preorder[0], nil, nil}
	i := 0
	for index, _ := range inorder {
		if preorder[0] == inorder[index] {
			i = index
			break
		}
	}
	currentnode.Left = buildTree(preorder[1:len(inorder[:i])+1], inorder[:i])
	currentnode.Right = buildTree(preorder[len(inorder[:i])+1:], inorder[i+1:])
	return &currentnode
}

func stacktolist() {
	type CQueue struct {
		list1 *list.List
		list2 *list.List
	}
	b := make([]int, 0)
	b[100] = 12
	fmt.Println(b[100])
}

var map1 map[int]int64

func fib(n int) int {
	map1 = make(map[int]int64)
	map1[0] = 0
	map1[1] = 1
	return int(fib2(n, map1) % 1000000007)
}
func fib2(n int, map1 map[int]int64) int64 {
	if n == 0 {
		return 0
	}
	if map1[n] != 0 && n != 0 {
		return map1[n]
	}
	map1[n] = fib2(n-1, map1) + fib2(n-2, map1)
	return map1[n] % 1000000007
}

var total int

func numWays(n int) int {
	total = 0
	numWays1(n)
	a := total
	return a
}
func numWays1(n int) {
	if n == 0 {
		total += 1
		total = total % 1000000007
		return
	}
	for i := 1; i <= 2; i++ {
		if n-i >= 0 {
			numWays1(n - i)
		}
	}

}

func minarray() int {
	array := []int{10, 1, 10, 10, 10}
	left := 0
	right := len(array) - 1
	// 去除掉右边没有使用的
	for {
		if right != left && array[right] == array[0] {
			right--
		} else {
			break
		}
	}
	if array[left] < array[right] && left != right {
		return array[left]
	}

	for {
		fmt.Println("enter circle")
		mid := (left + right) / 2
		if left >= right {
			break
		}
		if array[0] > array[mid] {
			right = mid
		} else {
			left = mid + 1
		}

	}
	return array[left]
}

var j, k int

func exist(board [][]byte, word string) bool {
	row := len(board)
	col := len(board[0])
	vis := make([][]bool, row)
	for i := range vis {
		vis[i] = make([]bool, col)
	}
	// for 循环所有的情况，所以在一个case 中只用往后边判断就可以了,defer 写到判断内部，保证jk 是有效的
	for z := 0; z < row; z++ {
		for v := 0; v < col; v++ {
			if exist1(board, word, z, v, row, col, vis) {
				return true
			}
		}
	}
	return false
}

func exist1(board [][]byte, word string, j, k, row, col int, vis [][]bool) bool {
	if len(word) == 0 {
		return true
	}
	if j < row && 0 <= j && 0 <= k && k < col &&
		vis[j][k] == false && board[j][k] == word[0] {
		defer func() { vis[j][k] = false }()
		vis[j][k] = true
		return exist1(board, word[1:], j-1, k, row, col, vis) ||
			exist1(board, word[1:], j, k-1, row, col, vis) ||
			exist1(board, word[1:], j+1, k, row, col, vis) ||
			exist1(board, word[1:], j, k+1, row, col, vis)
	} else {
		return false
	}

}
func check1(a map[int][]int, b []int) bool {
	for _, v := range a {
		if v[0] == b[0] && v[1] == b[1] {
			return true
		}
	}
	return false
}
func movingCount(m int, n int, k int) int {
	counter := 0
	res := make(map[int][]int)
	res[-1] = []int{0, 0}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			limit := i/100 + (i-(i/100)*100)/10 + (i - ((i-(i/100)*100)/10)*10 - (i/100)*100)
			limit += j/100 + (j-(j/100)*100)/10 + (j - ((j-(j/100)*100)/10)*10 - (j/100)*100)
			if (check1(res, []int{i - 1, j}) || check1(res, []int{i, j - 1})) && limit <= k {
				res[counter] = []int{i, j}
				counter += 1
			}
		}
	}
	return counter
}

func cuttingRope(n int) int {
	result := 1
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}
	if n == 4 {
		return 4
	}
	if n > 4 {
		for n > 4 {
			result *= 3
			n = n - 3
		}

	}
	return result * n
}
func change(a uint32) int {
	c := fmt.Sprintf("%b", a)
	count := strings.Count(c, "1")
	fmt.Println(c, count)
	return 1
}

func pow(n int) int {
	if n == 0 {
		return 1

	} else if n == 1 {
		return 10
	} else {
		return 10 * pow(n-1)
	}
}

func printNumbers(n int) []int {
	slice := make([]int, 0)
	for i := 1; i < pow(n); i++ {
		slice = append(slice, i)
	}
	return slice
}

func ismatch(s string, p string) bool {
	if len(p) == 0 {
		return len(s) == 0
	}
	// 递归中间只用求解一个对应的位置的结果就可以进行操作了
	// 将匹配的结果和后续的操作分开了
	firstmatch := len(s) != 0 && (s[0] == p[0] || p[0] == '.')
	if len(p) >= 2 && p[1] == '*' {
		return (firstmatch && ismatch(s[1:], p)) || ismatch(s, p[2:])
	} else {
		return firstmatch && ismatch(s[1:], p[1:])
	}
}
func isNumber(s string) bool {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return false
	}
	f1 := func(r rune) bool {
		return r == 'E' || r == 'e' || r == '.'
	}
	strslice := strings.FieldsFunc(s, f1)
	// fmt.Println(strslice)
	if len(strslice) > 2 || (strings.ContainsAny(s, "Ee.") && len(strslice) <= 1) {
		return false
	}
	for _, value := range strslice {
		_, err := strconv.Atoi(value)
		if err != nil {
			return false
		}
	}
	return true
}

func exchange(nums []int) []int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		if nums[left]%2 == 1 {
			left++
		} else if nums[right]%2 == 0 {
			right--
		} else {
			nums[left], nums[right] = nums[right], nums[left]
		}
	}
	return nums
}
func exchange1(nums []int) []int {
	newnums := make([]int, 0)
	for _, value := range nums {
		if value%2 != 0 {
			newnums = append([]int{value}, newnums...)
		} else {
			newnums = append(newnums, value)
		}
	}
	return newnums
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var new1 = &ListNode{0, nil}
	cur := new1
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}
	if l1 == nil {
		cur.Next = l2
	} else {
		cur.Next = l1
	}
	return new1
}
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	row := len(matrix)
	col := len(matrix[0])
	visit := make([][]bool, row)
	for index := range visit {
		visit[index] = make([]bool, col)
	}
	directions := [4][]int{[]int{0, 1}, []int{1, 0}, []int{0, -1}, []int{-1, 0}}
	rowcolindex := 0
	total := row * col
	res := make([]int, total)
	currdir := directions[rowcolindex]
	rown := 0
	coln := 0
	for i := 0; i < total; i++ {
		res[i] = matrix[rown][coln]
		visit[rown][coln] = true
		nextrow := rown + currdir[0]
		nextcol := rown + currdir[1]
		if nextrow < 0 || nextcol < 0 || nextrow >= row || nextcol >= col || visit[nextrow][nextcol] {
			rowcolindex = (rowcolindex + 1) % 4
			currdir = directions[rowcolindex]
		}
		rown = rown + currdir[0]
		coln = coln + currdir[1]
	}
	return res
}

type MinStack struct {
	Value    int
	Minvalue int
	Next     *MinStack
	Prev     *MinStack
}

/** initialize your data structure here. */
func Constructor() MinStack {
	min := MinStack{0, 65536, nil, nil}
	return min
}

func (this MinStack) Push(x int) {
	var newnode MinStack
	if x <= this.Minvalue {
		newnode = MinStack{x, x, nil, nil}
	} else {
		newnode = MinStack{x, this.Minvalue, nil, nil}
	}
	newnode.Prev = &this
	this.Next = &newnode
	this = newnode
	// fmt.Println(this.Minvalue)
}

func (this MinStack) Pop() {
	// this =*this.Prev
}

func (this MinStack) Top() int {
	return this.Value
}

func (this MinStack) Min() int {
	return this.Minvalue
}

type s struct {
	name string
	age  int
}

func (s *s) prt() {
	fmt.Println(s.name)
	fmt.Println(s.age)
	s.age = 100
}

func validateStackSequences(pushed []int, popped []int) bool {
	if len(pushed) <= 2 {
		return true
	}
	left, right := 0, 0
	stack := list.New()
	for {
		if left < len(pushed) && pushed[left] != popped[right] {
			if stack.Len() > 0 && stack.Back().Value == popped[right] {
				right++
				stack.Remove(stack.Back())
			} else {
				stack.PushBack(pushed[left])
				left++
			}

		} else if left >= len(pushed) {
			if stack.Len() > 0 && stack.Back().Value == popped[right] {
				right++
				stack.Remove(stack.Back())
			} else if stack.Len() == 0 {
				return len(popped) == right
			} else {
				return false
			}
			left++
		} else {
			right++
			left++
		}
	}

}

func deepcopy() {
	fmt.Println("通过 json 序列化来实现。或者通过gob go 中的对象来实现")
	// var buffer bytes.Buffer
	// gob.NewEncoder(&buffer).Encode("hello")
	// gob.NewDecoder(&buffer).Decode()
}

// quicksort  使用的方法占用了额外的空间来进行存储左右子序列，减轻了逻辑上的负担
func quicksort(arr []int) []int {
	if len(arr) == 0 {
		return []int{}
	}
	val := arr[0]
	left := make([]int, 0)
	right := make([]int, 0)
	for i := 1; i < len(arr); i++ {
		if arr[i] <= val {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}
	sortleft := quicksort(left)
	sortright := quicksort(right)
	res := make([]int, 0)
	res = append(res, sortleft...)
	res = append(res, val)
	res = append(res, sortright...)

	return res
}

func heapify(arr []int, i int) []int {
	lastindex := 0
	if i*2 >= len(arr) {
		return nil
	} else {
		if i*2+1 < len(arr) && arr[i*2] < arr[i*2+1] {
			lastindex = i*2 + 1
		} else {
			lastindex = i * 2
		}
	}
	if arr[lastindex] > arr[i] {
		arr[lastindex], arr[i] = arr[i], arr[lastindex]
		heapify(arr, lastindex)
	}

	return arr
}
func heapsort(arr []int) []int {
	for i := len(arr) - 1; i >= 0; i-- {
		heapify(arr, i)
	}
	fmt.Println(arr)
	// 是要从这个地方开始
	for i := len(arr)/2 - 1; i >= 0; i-- {
		arr[i], arr[0] = arr[0], arr[i]
		// 就是从0 号位置进行计算然后操作， 同时这个地方传递参数就是对应的切片的切片
		// 就可以对后边的元素进行屏蔽，然后最后还可以输出全部的元素
		heapify(arr[:i], 0)
	}
	return arr
}
func testsynatx() {
	// var a ListNode
	// fmt.Println(a)
	// res:=reversePrint(&a)
	// fmt.Println(res,[]int{})
	// 手动构建二叉树
	// a:=buildTree([]int{3,9,20,15,7},[]int{9,3,15,20,7})
	// fmt.Println(a.Left.Val)
	// fmt.Println(a.Right.Right.Val)
	// stacktolist()165580141
	// fmt.Println(map1[12])
	// map1=make(map[int]int64)
	// map1[0]=0
	// map1[1]=1
	// fmt.Println(fib(92))
	// numWays(1)
	// fmt.Println(total)
	// numWays(41)
	// 267914296
	//267914296
	// fmt.Println(total)
	// fmt.Println(minarray())
	// ch1:=make([]int,0)
	// ch1=append(ch1,2)
	// fmt.Println(len(ch1),cap(ch1))
	// fmt.Println(ch1,ch1[1:])
	// //
	// a:=[][]byte{{'C','A','A'},{'A','A','A'},{'B','C','D'}}
	// word:="AAB"
	// c:=exist(a,word)+(i-i/100)/10+(i-(i/100)*100-(i-i/100)/10)*10/1
	// i:=12
	// fmt.Println(i/100+(i-(i/100)*100)/10+ (i-((i-(i/100)*100)/10)*10-(i/100)*100))
	// fmt.Println(i/100,(i-(i/100)*100)/10, (i-((i-(i/100)*100)/10)*10-(i/100)*100))
	// 	a:=make(map[int][]int)
	// 	a[0]=[]int{1,2}
	// 	c:=[]int{1,2}
	// 	fmt.Println(check1(a,c))
	// 	fmt.Println(movingCount(1,2,1))
	// 	change(3)
	// 	fmt.Println(printNumbers(2))"mississippi"
	// "mis*is*ip*."
	// fmt.Println(	ismatch("bbbba",".*a*a")){"+100", "5e2", "-123", "3.1416", "-1E-16", "0123"}
	//["12e", "1a3.14", "1.2.3", "+-5", "12e+5.4"]
	// for _,value:=range[]string{"e", "1a3.14", "1.2.3", "+-5", "12e+5.4"}{
	// 	fmt.Println(isNumber(value))
	// }
	// 	fmt.Println(exchange([]int{1,2,3,4}))
	// 	fmt.Println(exchange1([]int{1,2,3,4}))
	// 	var newnode1 *ListNode
	// 	var newnode2 *ListNode
	// 	// var newnode4 *ListNode
	// 	newnode1=&ListNode{
	// 		12,
	// 		nil,
	// 	}
	// 	newnode1.Next=&ListNode{
	// 		15,
	// 		nil,
	// 	}
	// 	// newnode2=&ListNode{
	// 	// 	11,
	// 	// 	nil,
	// 	// }
	// 	// newnode2.Next=&ListNode{
	// 	// 	14,
	// 	// 	nil,
	// 	// }
	// 	var butters bytes.Buffer
	// 	data,_:=json.Marshal(newnode1)
	// 	fmt.Println(len(data))
	// 	json.Unmarshal(data,newnode2)
	// 	newnode3:=&newnode1
	// 	err1:=gob.NewEncoder(&butters).Encode(newnode1)
	// 	err2:=gob.NewDecoder(&butters).Decode(newnode1)
	// 	if err1!=nil || err2!=nil{
	// 		fmt.Println(err1)
	// 		fmt.Println(err2)
	// 	}
	// 	fmt.Printf("==============%v %v  %v  %v \n",&newnode1,&newnode2,newnode3,newnode1)
	// 	// 可以使用json 来对数据进行序列化，然后进行深层拷贝
	// 	// ioutil.WriteFile("a.json",data,777)
	//
	// // mergeTwoLists(newnode1,newnode2)
	// // 	spiralOrder([][]int{{1,2,3},{4,5,6},{7,8,9}})
	// // 	const INT_MAX = int(^uint(0) >> 1)
	// // fmt.Println(int(math.Inf(1)))
	// // 作用是起到了，但是没有返回对应的指针指向
	// fmt.Println("----------------")
	// 	//  如果直接使用的是值传递，一般不会改变对应的内部的指针指向
	// 	//
	// 	a:=Constructor()
	// 	fmt.Printf("%v %v\n",a,&a)
	// 	a.Push(1)
	// 	// a传入的是一个地址，a本身不会发生改变
	// 	fmt.Printf("%v %v %v\n",a.Next,a.Next,a)
	// 	a.Push(5)
	// 	fmt.Printf("%v %v %v \n",a.Next,a.Next ,a)
	// 	fmt.Println(a.Top())
	// 	a.Pop()
	// 	fmt.Println(a.Top())
	// 	// z:=s{"hello",12}
	// 	// z.prt()
	// 	// fmt.Println(z.age)
	// 	newlist:= list.New()
	// 	fmt.Println(newlist.Back())
	// 	newlist.Len()
	// 	newlist.PushBack(12)
	// 	newlist.PushBack("23")
	// 	fmt.Println(newlist.Back().Value)
	// 	// 可以移除最后的一个元素
	// 	newlist.Remove(newlist.Back())
	// 	fmt.Println(newlist.Back().Value)
	// 	fmt.Println(validateStackSequences([]int{2,1,0},[]int{1,2,0}))
	// 	c:=make([]int,0)
	// 	h:=[]int{1,2,4}
	// 	w:=[]int{4,5,6}
	// 	c=append(c,h...)
	// 	c=append(c,w...)
	// 	fmt.Println(c)
	// 	queue:=make([]int,0)
	// 	fmt.Println(len(queue))
	// 	for i:=0;i<len(queue);i++{
	// 		fmt.Println("hello")
	// 	}

}

// func bloom1(){
// filter:=bloom.NewWithEstimates(100000,0.00001)
// filter.Add([]byte("string1"))
// filter.Add([]byte("string2"))
// filter.Add([]byte("string3"))
// filter.Add([]byte("string4"))
// fmt.Println(filter.Test([]byte("string5")))
// }
// func reverse(s []string) []string {
// 	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
// 		s[i], s[j] = s[j], s[i]
// 	}
// 	return s
// }
// func reverseWords(s string) string {
// newstr:=strings.TrimSpace(s)
// newstring:=strings.Split(newstr," ")
// return strings.Join(reverse(newstring)," ")
// }
func main() {
	var a int64
	if a>>3&1 == 1 {
		fmt.Println("index 4")
	}
	fmt.Printf(" %v \n", a)
	a += 4
	//  a>>2&1==1 判断数据中的某一位的值
	if a>>43&1 == 1 {
		fmt.Println("index 4")
	}
	fmt.Printf(" %v \n", a)
	// opt:=redis.Options{}
	// // 设置对应的端口来进行
	// opt.Addr="127.0.0.1:6379"
	// client:=redis.NewClient(&opt)
	// stu:=client.Set("hello",12,-1)
	// fmt.Println(stu.String())
	// stu1:=client.Get("hello")
	// 	fmt.Println(stu1.String())
	// fmt.Println(levenshtein.Distance("qwe","we",levenshtein.NewParams()))

	// fmt.Println(quicksort([]int{2,5,1,6,3,1}))
	fmt.Println(heapsort([]int{2, 5, 1, 6, 3, 1}))
	a = 10
	b := 20
	fmt.Println(int(math.Max(float64(a), float64(b))))
	// fmt.Println(reverseWords("the sky     is blue"))
	z := strings.Split("the sky is blue", " ")
	var w []string
	fmt.Println(len(z))
	for index := range z {
		// fmt.Println(z[index])
		if z[index] == "" {
			w = append(z[0:index], z[index+1:]...)
		}
	}
	fmt.Println(len(w))
	fmt.Printf("version: %p\n", version)
	fmt.Printf("version: %p\n", version2)
	direntry, err := filesystem.ReadDir("dir")
	if err != nil {
		fmt.Println(err)
	}
	for _, file := range direntry {
		if !file.IsDir() {
			data, err := ioutil.ReadFile(file.Name())
			if err != nil {
				fmt.Println(err)
			}
			if len(data) < 10000 {
				fmt.Printf(" show all : %q\n", data)
			} else {
				fmt.Printf(" show part : %q\n", data[:10000])
			}

		}
	}

	fsd, err := fs.Sub(filesystem, "dir")
	if err != nil {
		fmt.Println(err)
	}
	http.Handle("/", http.FileServer(http.FS(fsd)))
	err = http.ListenAndServe(":80", nil)
	if err != nil {
		fmt.Println(err)
	}

}
