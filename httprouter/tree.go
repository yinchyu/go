package httprouter

type nodeType uint8
type node struct {
	path      string
	wildChild bool
	// 三种具体的类型
	nType nodeType

	indices  []byte
	children []*node
	//
	handle map[string]Handle

	priority uint32
}
