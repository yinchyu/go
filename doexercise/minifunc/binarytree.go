package main

import "fmt"
// 定义树的类型
type  BinaryTree struct {
 val int
 leftnode *BinaryTree
 rightnode *BinaryTree
}
func printtree( root *BinaryTree){
	// 如果树为空就直接结束打印并换行
	if root==nil{
		fmt.Println()
		return
	}
	queue:=[]*BinaryTree{root}
	for len(queue)>0{
		// temp:=queue
		queue=[]*BinaryTree{}
		size:=len(queue)
		// top:=queue[len(queue)-1]
		for i :=range queue{
			// 打印节点值
			fmt.Print(queue[i].val)
			if queue[i].leftnode!=nil{
				queue=append(queue,queue[i].leftnode)
			}
			if queue[i].rightnode!=nil{
				queue=append(queue,queue[i].rightnode)
			}
		}
		// 将前边的元素都删除
		queue=queue[size:]
		// 换行输出
		fmt.Println()
	}
}
func main(){
	// printtree( root)
}
type  Tree struct{
	val int
	leftnode *Tree
	rightnode *Tree
	leftbrother *Tree
	rightbrother *Tree
}

// BuildTree 构建树的左右兄弟
func BuildTree( root *Tree) *Tree{

	// 树为空直接返回
	if root ==nil{
		return  nil}

	// 层次遍历一颗树， 并给节点赋值
	queue:=[]*Tree{root}

	for len(queue)>0{
		// 将树的一层重新赋值给temp，然后queue 放置下一层的数据
		temp:=queue
		queue=[]*Tree{root}
		for i :=range temp{
			leftindex:=i-1
			rightindex:=i+1

			// 分别给左子树和右子树的兄弟赋值
			if 0<=leftindex &&leftindex<len(temp){

				temp[i].leftbrother =temp[leftindex]

			}
			if 0<=rightindex&&rightindex<len(temp){
				temp[i].rightbrother =temp[rightindex]
			}

			//  添加下一层的元素
			if temp[i].leftnode !=nil{

				queue=append(queue,temp[i].leftnode)

			}

			if temp[i].rightnode!=nil{
				queue=append(queue,temp[i].rightnode)
			}

		}

	}

	//
	// 2. t1表包含字段 (uid, event)和数据
	// 1 ,  2
	// 2.   3
	// 1,   4
	// 2,   5
	//
	//
	// 查询同时做了事件2和事件4的用户列表
	//
	//
	// select distinct(uid) from t1 join t1 as t2  on t1.uid=t2.uid  where t1.event=2 and t2.event=4;

return nil
	}