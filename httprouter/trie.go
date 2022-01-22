package httprouter

import (
	"bytes"
	"fmt"
	"runtime"
	"strconv"
	"unicode/utf8"
)

type trie struct {
	isend    bool
	childern [26]*trie
}

func newtrie() *trie {
	root := trie{childern: [26]*trie{}}
	return &root
}

func (t *trie) insert(a string) {
	for i := 0; i < len(a); i++ {
		//只考虑小写字母
		index := a[i] - 'a'
		if t.childern[index] == nil {
			t.childern[index] = newtrie()
		}
		t = t.childern[index]
	}
	t.isend = true
}

// 字典树和布隆过滤器一样只能添加不能删除
//func delete(t *trie)delete(){
//
//}

func (t trie) startwith(s string) bool {
	for i := 0; i < len(s); i++ {
		index := s[i] - 'a'
		if t.childern[index] == nil {
			return false
		} else {
			t = *t.childern[index]
		}

	}
	return true
}

func (t trie) isword(s string) bool {
	for i := 0; i < len(s); i++ {
		index := s[i] - 'a'
		if t.childern[index] == nil {
			return false
		} else {
			t = *t.childern[index]
		}

	}
	return t.isend
}

func getgoroutine() {
	buf := make([]byte, 64)
	n := runtime.Stack(buf, true)
	str := buf[:n]
	str = str[len("goroutine "):]
	// 在字符串的两个空格之间就是 goroutine 的number号
	str = str[:bytes.IndexByte(str, ' ')]
	gid, _ := strconv.ParseInt(string(str), 10, 64)
	fmt.Println(gid)

}

func test() {
	str := []byte{103, 111, 114, 111, 117, 116, 105, 110, 101, 32, 54, 32, 91, 114, 117, 110, 110, 105, 110, 103, 93, 58, 10, 104, 116, 116, 112, 114, 111, 117, 116, 101, 114, 46, 103, 101, 116, 103, 111, 114, 111, 117, 116, 105, 110, 101, 40, 41, 10, 9, 68, 58, 47, 230, 161, 140, 233, 157, 162, 230, 150, 135, 228, 187}
	var stackstring []rune
	for len(str) > 0 {
		r, size := utf8.DecodeRune(str)
		stackstring = append(stackstring, r)
		str = str[size:]
	}
	fmt.Println(string(stackstring))
}
