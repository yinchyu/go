package main

import "fmt"

func swap(s string) string {
	slist := []byte(s)
	n := len(slist)
	res := make([]int, 0)
	for i := 0; i < n; i++ {
		if Checkbyte(slist[i]) {
			res = append(res, i)
		}
	}
	for j, k := 0, len(res)-1; j < k; j, k = j+1, k-1 {
		l := res[j]
		r := res[k]
		slist[l], slist[r] = slist[r], slist[l]
	}
	return string(slist)
}

func Checkbyte(b byte) bool {
	res := []byte("aeiou")
	if 'A' <= b && b <= 'Z' {
		return false
	}
	for i := 0; i < len(res); i++ {
		if res[i] == b {
			return false
		}
	}
	return true
}

func main() {
	s := "helloL"
	fmt.Println(swap(s))
}
