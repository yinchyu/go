package cache

import (
	"main/cache/lru"
	"sync"
)

type cache struct {
	mu         sync.Mutex
	// 在基本的lru缓存中增加锁和缓存的字节多少
	lru        *lru.Cache
	//就是底层的maxbytes
	cacheBytes int64
}

func (c *cache) add(key string, value ByteView) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.lru == nil {
		c.lru = lru.New(c.cacheBytes, nil)
	}
	c.lru.Add(key, value)
}

func (c *cache) get(key string) (value ByteView, ok bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.lru == nil {
		return
	}

	if v, ok := c.lru.Get(key); ok {
		return v.(ByteView), ok
	}

	return
}