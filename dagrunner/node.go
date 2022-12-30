package dagrunner

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

type Node struct {
	name     string
	refCount int64
	RunFunc  RunFunc
	parents  []string
	child    []string
}
type Graph struct {
	NodeMap       map[string]*Node
	StareNodeList []*Node
	JoinNode      *Node
	mux           sync.Mutex
	expireTime    time.Duration
}
type RunFunc func(context.Context) error

func NewGraph() *Graph {
	return &Graph{
		NodeMap: make(map[string]*Node),
	}
}
func NewNode(nodeName string, runFunc RunFunc) *Node {
	return &Node{
		name:    nodeName,
		RunFunc: runFunc,
	}
}

func (g *Graph) Add(nodeName string, parentNodeList []string, runFunc RunFunc) error {
	node := NewNode(nodeName, runFunc)
	// 	没有依赖节点
	g.NodeMap[nodeName] = node
	if len(parentNodeList) == 0 {
		g.StareNodeList = append(g.StareNodeList, node)
	} else {
		node.refCount = int64(len(parentNodeList))
		node.parents = parentNodeList
		for _, pNode := range parentNodeList {
			g.NodeMap[pNode].child = append(g.NodeMap[pNode].child, nodeName)
		}
		return g.checkCycle()
	}
	return nil
}

func (g *Graph) checkCycle() error {
	checkMap := make(map[string]int64)
	readyList := make([]string, 0)
	for s, node := range g.NodeMap {
		checkMap[s] = node.refCount
	}
	for len(checkMap) > 0 {
		for name, refCount := range checkMap {
			if refCount == 0 {
				readyList = append(readyList, name)
			}
		}
		if len(readyList) == 0 {
			return errors.New("have cycle")
		}
		for _, name := range readyList {
			delete(checkMap, name)
			for _, childName := range g.NodeMap[name].child {
				checkMap[childName] -= 1
			}
		}
	}
	return nil
}
func (g *Graph) runNode(node *Node, ch chan string) error {
	err := node.RunFunc(context.Background())
	if err != nil {
		return err
	}
	ch <- node.name
	return nil
}

func (g *Graph) Run() {
	ch := make(chan string, len(g.NodeMap))
	for _, node := range g.StareNodeList {
		go func(node *Node) {
			err := g.runNode(node, ch)
			if err != nil {
				fmt.Print(err)
			}
		}(node)
	}
	readyCount := 0
	for readyCount < len(g.NodeMap) {
		select {
		case nodeName := <-ch:
			g.mux.Lock()
			readyCount += 1
			for _, childName := range g.NodeMap[nodeName].child {
				g.NodeMap[childName].refCount -= 1
				if g.NodeMap[childName].refCount == 0 {
					go func() {
						err := g.runNode(g.NodeMap[childName], ch)
						if err != nil {
							fmt.Print(err)
						}
					}()
				}
			}
			g.mux.Unlock()
		case <-time.After(g.expireTime):
			//超时
			return
		}
	}
}
