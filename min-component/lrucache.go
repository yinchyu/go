package main

import (
	"fmt"
	"sync"
)

type LruCache struct {
	Cap  int
	// 实现一个泛型的map   key campareable    value  interface
	Lru  map[int]*Node
	Head *Node
	Tail *Node
	sync.RWMutex
}

type Node struct {
	key  int
	val  interface{}
	// 存储的是两个冗余的节点
	pre  *Node
	next *Node
}

var Lru *LruCache
var once sync.Once

func OnceCreate(cap int) {
	once.Do(func() {
		Create(cap)
	})
}

func Create(cap int) {
	cache := &LruCache{
		Cap:  cap,
		Lru:  make(map[int]*Node),
		Head: &Node{},
		Tail: &Node{},
	}
	// 底层的存储结构是一个双向链表
	cache.Head.next = cache.Tail
	cache.Tail.pre = cache.Head
	Lru = cache
}
func (l *LruCache) SetHeader(n *Node) {
	n.next = l.Head.next
	n.pre = l.Head
	l.Head.next.pre = n
	l.Head.next = n
}


func (l *LruCache) Remove(n *Node) {
	// 双向链表的好处是可以随意的删除和，如果有map 进行指引操作
	n.next.pre = n.pre
	n.pre.next = n.next
}


func (l *LruCache) Get(key int) interface{} {
	l.RLock()
	defer l.RUnlock()
	node, ok := l.Lru[key]
	if !ok {
		return -1
	} else {
		l.Remove(node)
		l.SetHeader(node)
		return node.val
	}
}

// 如果有泛型的话,这个地方的map 就是compareable 的类型
func (l *LruCache) Put(key int , value interface{}) {
	l.Lock()
	defer l.Unlock()
	node, ok := l.Lru[key]
	if ok {
		l.Remove(node)
	} else {
		if len(l.Lru) == l.Cap {
			fmt.Println("=====")
			// 双重删除操作， 一个删除map 中的， 一个删除 链表中间
			delete(l.Lru, l.Tail.pre.key)
			l.Remove(l.Tail.pre)
		}
		node = &Node{key, value, nil, nil}
		l.Lru[key] = node
	}
	node.val = value
	l.SetHeader(node)
}



func main() {
	// c:=Create(2)
	OnceCreate(2)
	Lru.Put(2, 3)
	Lru.Put(3, 4)
	fmt.Println(Lru.Get(2))
	Lru.Put(4, 5)
	fmt.Println(Lru.Get(2))
	// 就是找不到呗
	fmt.Println(Lru.Get(3))
}
