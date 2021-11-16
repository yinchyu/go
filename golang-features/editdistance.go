package main

import (
	"fmt"
	"sort"
)

//所有常量的运算都可以在编译期完成， 所以说不能给常量赋值给一个变量
const a int = 123

func main() {
	str2 := "xhcjhjsrjamxkjdbiqnqjxfsugywrpceyuniviqdynpipfytbaijwsjnhynxnwzydyxmrqlxnygttbaqsneejojukjkkxrsxyywmnsgcuxbnnavmytbthosfuhzytripxthaciiupodunllqz"
	str1 := "xmhjrtsajahxkjialbtfgrikbepnwqnjrrxfsssgyvrnrmocxuqshmvwqsphyqndipxabbsaijjbnxnxjnywndyxscmrlvnarvtgbwaqgsgjeegazofjmuiocjlxwemyyywmsgyuxnmjmytvybgyghcxtsfuhzyqzhdhsdviokpmstgciaukbkpniotddnvmqz"
	fmt.Println(len(str2))
	length1 := len(str1)
	length2 := len(str2)
	var (
		//6477,7414,1063
		insert  = 6477
		delet   = 7414
		replace = 1063
	)

	// dem1:=make(dem ,length2) 5,3,2
	dptable := make([][]int, length1+1)
	for i, _ := range dptable {
		dptable[i] = make([]int, length2+1)
		dptable[i][0] = i * insert
	}
	for j, _ := range dptable[0] {
		dptable[0][j] = delet * j
	}
	// fmt.Println(dptable,length2,length1,len(dptable[0]),cap(dptable[0]),dptable[0][0])
	for i := 1; i <= length1; i++ {
		for j := 1; j <= length2; j++ {
			if str1[i-1] == str2[j-1] {
				dptable[i][j] = dptable[i-1][j-1]
			} else {
				re := dptable[i-1][j-1] + replace
				in := dptable[i-1][j] + insert
				dl := dptable[i][j-1] + delet
				var all = []int{re, in, dl}
				sort.Ints(all)
				fmt.Println(all)
				dptable[i][j] = all[0]
			}
		}
	}
	// return dptable[length1][length2]
	fmt.Printf("%#v\n", dptable[length1][length2])
	// fmt.Printf("%v\n",dptable)

}
