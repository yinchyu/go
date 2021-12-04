package consistenthash

import (
	"hash/crc32"
	"sort"
	"strconv"
)

type Hash func(data []byte) uint32

type Map struct {
	hash Hash
	// 虚拟节点的倍数
	replicas int
	//hash 环
	keys []int

	hashmap map[int]string
}

func New(replicas int, fn Hash) *Map {
	m := &Map{
		replicas: replicas,
		hash:     fn,
		hashmap:  make(map[int]string),
	}
	if m.hash == nil {
		m.hash = crc32.ChecksumIEEE
	}
	return m
}

func (m *Map) Add(keys ...string) {
	for _, key := range keys {
		for i := 0; i < m.replicas; i++ {
			hash := int(m.hash([]byte(strconv.Itoa(i) + key)))
			m.keys = append(m.keys, hash)
			m.hashmap[hash] = key
		}
	}
	sort.Ints(m.keys)
}
func (m *Map) Get(key string) string {
	if len(m.keys) == 0 {
		return ""
	}
	hash := int(m.hash([]byte(key)))
	idx := sort.Search(len(m.keys), func(i int) bool {
		// 满足该条件的第一个节点
		return m.keys[i] >= hash
	})
	return m.hashmap[m.keys[idx%len(m.keys)]]
}

// Remove use to remove a key and its virtual keys on the ring and map
func (m *Map) Remove(key string) {
	// 自然的均摊节点
	for i := 0; i < m.replicas; i++ {
		// 返回的是uint32类型的数据，所以节点可以进行排序
		hash := int(m.hash([]byte(strconv.Itoa(i) + key)))
		idx := sort.SearchInts(m.keys, hash)
		m.keys = append(m.keys[:idx], m.keys[idx+1:]...)
		delete(m.hashmap, hash)
	}
}

func (m *Map) Migration() {

}
