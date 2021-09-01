package main

import (
	"fmt"
	"strconv"
	"strings"
)

type linklist struct {
	val  int
	next *linklist
}

func Newlist(nums []int) *linklist {
	head := &linklist{}
	res := head
	for index := range nums {
		head.next = &linklist{nums[index], nil}
		head = head.next
	}
	return res.next
}

func reversek(head *linklist, k int) *linklist {
	if k == 0 || head == nil || head.next == nil {
		return head
	}
	fastpointer := head
	n := 1
	for fastpointer.next != nil {
		fastpointer = fastpointer.next
		n++
	}
	add := n - k%n
	if add == n {
		return head
	}
	fastpointer.next = head
	for add > 0 {
		fastpointer = fastpointer.next
		add--
	}
	ret := fastpointer.next
	fastpointer.next = nil
	return ret
}

func visit(head *linklist) {
	res := make([]string, 0)
	for head != nil {
		res = append(res, strconv.Itoa(head.val))
		head = head.next
	}
	fmt.Println(strings.Join(res, ","))
}

type myint int

func (my *myint) print() {
	fmt.Printf("%p %v\n", my, *my)
	var temp myint
	temp = 12
	// 内部进行赋值操作， 然后后边改变的就不会影响对应的原来的元素位置。
	my = &temp
	(*my) += 10
	fmt.Printf("%p %v\n", my, *my)
}

func main() {
	var c myint
	c = 123
	fmt.Printf("%p %v\n", &c, c)
	(&c).print()
	fmt.Printf("%p %v\n", &c, c)
	// 0xc00000a098 123
	// 123
	// 22
	// 0xc00000a098 123
}
