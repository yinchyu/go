
package main

import (
	"errors"
	"fmt"
)

const stackSize int = 22                //栈的容量

type TreeNode struct {
	Left *TreeNode
	Value int
	Right *TreeNode
}

type Stack struct {                     //栈结构
	Size int
	Values []*TreeNode
}

func CreateStack() Stack{                //创建栈
	s := Stack{}
	s.Size = stackSize
	return s
}

func (s *Stack) IsFull() bool {          //栈满
	return len(s.Values) == s.Size
}

func (s *Stack) IsEmpty() bool {          //栈空
	return len(s.Values) == 0
}

func (s *Stack) Push(a *TreeNode) error {       //入栈
	if s.IsFull() {
		return errors.New("Stack is full,push failed")
	}
	s.Values = append(s.Values,a)
	return nil
}

func (s *Stack) Pop() (*TreeNode,error) {//出栈,并返回其值
	if s.IsEmpty() {
		return nil ,errors.New("Out of index,len(stack) is 0")
	}
	res := s.Values[len(s.Values)-1]
	s.Values = s.Values[:len(s.Values)-1]
	return res,nil
}

func (s *Stack) Peek() (*TreeNode,error) {//查看栈顶元素
	if s.IsEmpty() {
		return nil ,errors.New("Out of index,len(stack) is 0")
	}
	return s.Values[len(s.Values)-1],nil
}

func (s *Stack) Traverse() {
	if s.IsEmpty() {
		fmt.Println("Stack is empty")
	}else {
		fmt.Println(s.Values)
	}
}

//建立二叉树
// 递归的建立一颗树， 和堆的使用有些相似
func TreeCreate(i int,arr []int) *TreeNode{
	t := &TreeNode{nil,arr[i],nil}
	if i<len(arr) && 2*i+1 < len(arr){
		t.Left = TreeCreate(2*i+1,arr)
	}
	if i<len(arr) && 2*i+2 < len(arr) {
		t.Right = TreeCreate(2*i+2,arr)
	}
	return t
}
//先根遍历，递归
func PreTraverse(t *TreeNode){
	if t != nil {
		fmt.Printf("%d/",t.Value)
		PreTraverse(t.Left)
		PreTraverse(t.Right)
	}
}
//先根遍历，非递归
func PreStackTraverse(t *TreeNode){
	if t != nil {
		S := CreateStack()
		S.Push(t)                               //根结点入栈
		for !S.IsEmpty() {
			T,_ := S.Pop()                      //移除根结点，并返回其值
			fmt.Printf("%d ",T.Value)    //访问结点
			for T != nil {
				if T.Left != nil {              //访问左孩子
					fmt.Printf("%d ",T.Left.Value)   //访问结点
				}
				if T.Right != nil {              //右孩子非空入栈
					S.Push(T.Right)
				}
				T = T.Left
			}
		}
	}
	fmt.Println()
}

//中根遍历，递归
func MidTraverse(t *TreeNode){
	if t != nil {
		MidTraverse(t.Left)
		fmt.Printf("%d/",t.Value)
		MidTraverse(t.Right)
	}
}

//中根遍历，非递归
func MidStackTraverse(t *TreeNode){
	if t != nil {
		S := CreateStack()
		S.Push(t)
		for !S.IsEmpty() {
			top,_ :=S.Peek()
			for  top != nil {       //将栈顶结点的所有左孩子结点入栈
				S.Push(top.Left)
				top,_= S.Peek()
			}
			S.Pop()                //空结点退栈
			if !S.IsEmpty() {
				T,_ := S.Pop()
				fmt.Printf("%d ",T.Value)
				S.Push(T.Right)
			}
		}
	}
	fmt.Println()
}
//后根遍历，递归
func PostTraverse(t *TreeNode){
	if t != nil {
		PostTraverse(t.Left)
		PostTraverse(t.Right)
		fmt.Printf("%d/",t.Value)
	}
}
//后根遍历，非递归
func PostStackTraverse(t *TreeNode) {
	if t != nil {
		S := CreateStack()
		S.Push(t)
		var flag bool
		var p *TreeNode
		for !S.IsEmpty() {
			top,_ := S.Peek()
			for  top != nil {       //将栈顶结点的所有左孩子结点入栈
				S.Push(top.Left)
				top,_= S.Peek()
			}
			S.Pop()
			for !S.IsEmpty() {
				T,_ := S.Peek()
				if T.Right == nil || T.Right == p {
					fmt.Printf("%d ",T.Value)
					S.Pop()
					flag = true
					p = T            //p指向刚被访问的结点
				}else {
					S.Push(T.Right)
					flag = false     //有右孩子进栈了
				}
				if !flag {            //退出该循环
					break
				}
			}
		}
	}
	fmt.Println()
}
// func main(){
// 	arr := []int{3,9,6,8,7,11,1,22,21}
// 	Tree := TreeCreate(0,arr)
// 	fmt.Print("前根,递归:")
// 	PreTraverse(Tree)
// 	fmt.Print("   前根,非递归:")
// 	PreStackTraverse(Tree)
// 	fmt.Print("中根,递归:")
// 	MidTraverse(Tree)
// 	fmt.Print("   中根,非递归:")
// 	MidStackTraverse(Tree)
// 	fmt.Print("后根,递归:")
// 	PostTraverse(Tree)
// 	fmt.Print("   后根,非递归:")
// 	PostStackTraverse(Tree)
// }
