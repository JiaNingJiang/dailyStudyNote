package geecache

import (
	"geeCache/lru"
	"sync"
)

// 为LRU Cache引入并发锁,支持并发操作
type cache struct {
	mu              sync.RWMutex
	lru             *lru.Cache
	cacheBytesLimit int // Cache缓存上限
}

func (c *cache) add(key string, value ByteView) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.lru == nil { // 采用延迟初始化操作(Lazy Initialization)。一个对象的延迟初始化意味着该对象的创建将会延迟至第一次使用该对象时。主要用于提高性能，并减少程序内存要求。
		c.lru = lru.NewCache(c.cacheBytesLimit, nil)
	}
	c.lru.Add(key, value)
}

func (c *cache) get(key string) (value ByteView, ok bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if c.lru == nil {
		return
	}

	if v, ok := c.lru.Get(key); ok {
		return v.(ByteView), true
	}
	return
}
