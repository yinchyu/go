package main

import "fmt"

type numberiter struct {
	start, end int
	cur        int
}

func (n *numberiter) isdone() bool {
	return n.cur > n.end
}

func (n *numberiter) next() int {
	if !n.isdone() {
		temp := n.cur
		n.cur++
		return temp
	} else {
		return 0
	}
}
func (n *numberiter) iter() {
	for !n.isdone() {
		fmt.Println(n.next())
	}
}

func main() {
	numiter := &numberiter{start: 1, end: 10, cur: 1}
	// numiter.iter()
	fmt.Println(numiter.next())
	fmt.Println(numiter.next())
	fmt.Println(numiter.next())

}
