package main
import ("fmt")

//建立二叉树
// func TreeCreate(i int,arr []int) *binarytree{
// 	t := &binarytree{nil,arr[i],nil}
// 	if i<len(arr) && 2*i+1 < len(arr){
// 		t.leftnode = TreeCreate(2*i+1,arr)
// 	}
// 	if i<len(arr) && 2*i+2 < len(arr) {
// 		t.leftnode = TreeCreate(2*i+2,arr)
// 	}
// 	return t
// }
// 先序遍历


func firstvisit(root  *TreeNode){
	if root ==nil{
		fmt.Println()
		return
	}
	// 使用栈来进行模式递归操作
	stack:=[]*TreeNode{root}
	for len(stack)>0{
		// 取出栈顶元素
		item:=stack[len(stack)-1]
		// 栈中的取出操作就是删除该元素
		stack=stack[:len(stack)-1]
		fmt.Println(item.Value)
		for item!=nil{
			if item.Left!=nil{

				fmt.Println(item.Left.Value)
			}
			if item.Right!=nil{
				// 只用有子树进行入栈的操作
				stack = append(stack, item.Right)
			}
			// 赋值给左节点
			item=item.Left
		}
	}
}
func firstvisit1( root *TreeNode){
	if root==nil{
		return
	}
	fmt.Println(root.Value)
	firstvisit1(root.Left)
	firstvisit1(root.Right)
}
func firstvisit2( root *TreeNode){
	// 最开始不用和其他的模拟栈的一样 将root 节点放入到对应的位置，但是需要借助另一个变量来进行操作
	stack:=make([]*TreeNode,0)
	p:=root
	for len(stack)>0 || p!=nil{
		for p!=nil{
			fmt.Println(p.Value)
			stack = append(stack, p)
			p=p.Left
		}
		if len(stack)>0{
			temp:=stack[len(stack)-1]
			stack=stack[:len(stack)-1]
			p=temp.Right
		}
	}
}



func secondvisit(root  *TreeNode){
	stack:=make([]*TreeNode,0)
	p:=root
	// 大循环的两个条件是 p 不空 栈不空，中序和前序的区别在于是入栈时打印还是出栈时打印
	for len(stack)>0 || p!=nil{
		for p!=nil{
			stack = append(stack, p)
			p=p.Left
		}
		if len(stack)>0{
			temp:=stack[len(stack)-1]
			fmt.Println(temp.Value)
			stack=stack[:len(stack)-1]
			p=temp.Right
		}
	}
}
func secondvisit1(root  *TreeNode){
	if root ==nil{
		return
	}
	secondvisit1(root.Left)
	fmt.Println(root.Value)
	secondvisit1(root.Right)
}

func threevisit(root  *TreeNode){
	stack:=[]*TreeNode{root}
     var pre *TreeNode
	 pre=nil
	for len(stack)>0{
			temp:=stack[len(stack)-1]
			// 通过一个pre 变量就可以断定 左右子树遍历完成
			 if (temp.Left==nil && temp.Right==nil) || pre!=nil&&(pre==temp.Left||pre==temp.Right){
				 fmt.Println(temp.Value)
				 stack=stack[:len(stack)-1]
				 pre=temp
			 }else{
			 	// 先放右子树
				 if temp.Right!=nil{
					 stack = append(stack, temp.Right)
				 }
			 	if temp.Left!=nil{
					stack = append(stack, temp.Left)
				}

			 }
	}
}
func threevisit1(root  *TreeNode){
	if root ==nil{
		return
	}
	secondvisit1(root.Left)
	secondvisit1(root.Right)
	fmt.Println(root.Value)
}

func main(){
	arr := []int{3,9,6,8,7,11,1,22,21}
	Tree := TreeCreate(0,arr)
	// firstvisit(Tree)
	// firstvisit2(Tree)
	// secondvisit(Tree)
	threevisit(Tree)
}


