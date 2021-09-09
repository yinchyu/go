//  二叉树的非递归遍历方式，需要注意的是二叉树的后序遍历方式，需要额外的变量来记录一个节点的左右子树是否都遍历过了

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type BinaryTree struct {
	val       int
	leftnode  *BinaryTree
	rightnode *BinaryTree
}

// 层次遍历一个树
func printtres(root *BinaryTree) {
	// 如果树为空就直接结束打印并换行
	if root == nil {
		fmt.Println()
		return
	}
	// 可以使用  for root!=nil  || len(queue)>0 来进行判断，就不用初始化的时候在queue中放置一个元素
	queue := []*BinaryTree{root}
	for len(queue) > 0 {
		// temp:=queue
		queue = []*BinaryTree{}
		size := len(queue)
		// top:=queue[len(queue)-1]
		for i := range queue {
			// 打印节点值
			fmt.Print(queue[i].val)
			if queue[i].leftnode != nil {
				queue = append(queue, queue[i].leftnode)
			}
			if queue[i].rightnode != nil {
				queue = append(queue, queue[i].rightnode)
			}
		}
		// 将前边的元素都删除
		queue = queue[size:]
		// 换行输出
		fmt.Println()
	}
}

type Tree struct {
	val          int
	leftnode     *Tree
	rightnode    *Tree
	leftbrother  *Tree
	rightbrother *Tree
}

// BuildTree 层次遍历构建树的左右兄弟
func BuildTree(root *Tree) *Tree {

	// 树为空直接返回
	if root == nil {
		return nil
	}

	// 层次遍历一颗树， 并给节点赋值
	queue := []*Tree{root}

	for len(queue) > 0 {
		// 将树的一层重新赋值给temp，然后queue 放置下一层的数据
		temp := queue
		queue = []*Tree{root}
		for i := range temp {
			leftindex := i - 1
			rightindex := i + 1

			// 分别给左子树和右子树的兄弟赋值
			if 0 <= leftindex && leftindex < len(temp) {

				temp[i].leftbrother = temp[leftindex]

			}
			if 0 <= rightindex && rightindex < len(temp) {
				temp[i].rightbrother = temp[rightindex]
			}

			//  添加下一层的元素
			if temp[i].leftnode != nil {

				queue = append(queue, temp[i].leftnode)

			}

			if temp[i].rightnode != nil {
				queue = append(queue, temp[i].rightnode)
			}

		}

	}
	return nil
}

func preorder(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	stack := make([]*TreeNode, 0)
	for root != nil || len(stack) > 0 {
		for root != nil {
			res = append(res, root.Val)
			stack = append(stack, root)
			root = root.Left
		}
		// 这个地方不用考虑 root.Right ==nil 的情况， 上边的for循环考虑
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = root.Right
	}

	return res
}

func inorder(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	stack := make([]*TreeNode, 0)
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, root.Val)
		root = root.Right
	}
	return res
}

func postorder(root *TreeNode) []int {
	res := make([]int, 0)
	// 当一个孩子的左右孩子都是空的时候， 或者他的孩子已经放入了结果集合中
	var prev *TreeNode
	if root == nil {
		return res
	}
	stack := make([]*TreeNode, 0)
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.Right == nil || root.Right == prev {
			res = append(res, root.Val)
			prev = root
			root = nil
		} else {
			stack = append(stack, root)
			root = root.Right
		}
	}
	return res
}
