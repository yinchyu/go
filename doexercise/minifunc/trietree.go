package main

import (
	"fmt"
	"strings"
	"unsafe"
)

type Trie struct {
	// 前缀树，通过一个26叉的字典树来进行操作
	Isend bool
	Next  []*Trie
}

func Constructor() *Trie {
	return &Trie{Isend: false, Next: make([]*Trie, 26)}
}

// 插入数据的位置
func (t *Trie) Insert(w string) {
	for i := range w {
		if t.Next[w[i]-'a'] == nil {
			newtree := Constructor()
			// 有空的位置就进行创建操作， 创建的具体的位置是通过constructor
			t.Next[w[i]-'a'] = newtree
		}
		t = t.Next[w[i]-'a']
	}
	// 每一个树叉都有一个isend 的标识，来进行按照这个字母结尾的位置的字符的判断
	t.Isend = true
}

//查找数据的位置
func (t *Trie) Search(w string) bool {
	for i := range w {
		if t.Next[w[i]-'a'] == nil {
			return false
		}
		t = t.Next[w[i]-'a']
	}
	return t.Isend
}

// 为什么可以直接使用指针进行操作， 应为指针的指向没有被改所以在 obj 中存储的还是最开始的位置的指针
// 按照某一个单词开头的位置
func (t *Trie) StartsWith(w string) bool {
	for i := range w {
		if t.Next[w[i]-'a'] == nil {
			return false
		}
		t = t.Next[w[i]-'a']
	}
	return true
}
func addelement(slice []int, e int) []int {
	panic("传递参数出现了错误")
	//return append(slice,e)
}

func strop() {
	obj := Constructor()
	obj.Insert("word")
	fmt.Printf("%p\n", &obj)
	param_2 := obj.Search("words")
	param_3 := obj.StartsWith("w")
	fmt.Printf("%p\n", &obj)
	fmt.Println(param_2, param_3)
	var slice []int64
	slice = append(slice, 1, 2, 3)
	// slice:=make([]int16,10)
	// slice=slice[5:][:5:10]
	fmt.Println(len(slice), cap(slice))
	fmt.Println(unsafe.Sizeof(slice))
	// newslice:=addelement(slice,4)
	// fmt.Printf("%p,%p,\n",slice,newslice)

	sslice := strings.Split(strings.TrimSpace("hello world"), " ")
	sslice = []string{"hello", "world"}
	newslice := make([]string, 0)
	for i := len(sslice) - 1; i >= 0; i-- {
		if sslice[i] != "" {
			newslice = append(newslice, sslice[i])
		}
	}
	fmt.Println(newslice)
	fmt.Println(strings.Join(newslice, " "))
}
func main() {
	addelement([]int{1, 2, 3, 4}, 23)
}
