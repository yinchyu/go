package main

import (
	"fmt"
	"sync"
)

type LruCache struct {
	Cap  int
	Lru  map[int]*Node
	Head *Node
	Tail *Node
	sync.RWMutex
}

type Node struct {
	key  int
	val  int
	pre  *Node
	next *Node
}

var Lru *LruCache

func OnceCreate(cap int) {
	var once sync.Once
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
	cache.Head.next = cache.Tail
	cache.Tail.pre = cache.Head
	Lru = cache
}

func (l *LruCache) Get(key int) int {
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
func (l *LruCache) Put(key, value int) {
	l.Lock()
	defer l.Unlock()
	node, ok := l.Lru[key]
	if ok {
		l.Remove(node)
	} else {
		if len(l.Lru) == l.Cap {
			fmt.Println("=====")
			delete(l.Lru, l.Tail.pre.key)
			l.Remove(l.Tail.pre)
		}
		node = &Node{key, value, nil, nil}
		l.Lru[key] = node
	}
	node.val = value
	l.SetHeader(node)
}

func (l *LruCache) SetHeader(n *Node) {
	n.next = l.Head.next
	n.pre = l.Head
	l.Head.next.pre = n
	l.Head.next = n
}
func (l *LruCache) Remove(n *Node) {
	n.next.pre = n.pre
	n.pre.next = n.next
}

func main() {
	// c:=Create(2)
	OnceCreate(2)
	Lru.Put(2, 3)
	Lru.Put(3, 4)
	fmt.Println(Lru.Get(2))
	Lru.Put(4, 5)
	fmt.Println(Lru.Get(2))
}
