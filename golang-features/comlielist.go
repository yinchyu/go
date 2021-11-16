package main

import (
	"fmt"
)

var res []int

func fundnode(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Left == nil || root.Right == nil {
		return root
	}
	rightnode := fundnode(root.Right)
	if rightnode != nil {
		res = append(res, rightnode.Val)
	}
	res = append(res, root.Val)
	leftnode := fundnode(root.Left)
	if leftnode != nil {
		res = append(res, leftnode.Val)
	}
	return nil
}
func kthLargest(root *TreeNode, k int) int {
	fundnode(root)
	//fmt.Println(res[k-1])
	return res[k-1]
}
func main() {
	sieve()

}

func generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i // Send 'i' to channel 'ch'.
	}
}
func filter(src <-chan int, dst chan<- int, prime int) {
	str := ""
	for i := 0; i < prime; i++ {
		str += "   "
	}
	for i := range src { // Loop over values received from 'src'.
		fmt.Println(str, "prime=", prime, "   i=", i)
		if i%prime != 0 {
			dst <- i // Send 'i' to channel 'dst'.
		}
	}
}

func sieve() {
	ch := make(chan int) // Create a new channel.
	go generate(ch)      // Start generate() as a subprocess.
	for i := 0; i < 10; i++ {
		prime := <-ch
		fmt.Print(prime, "\n")
		ch1 := make(chan int)
		go filter(ch, ch1, prime)
		ch = ch1
	}
}
