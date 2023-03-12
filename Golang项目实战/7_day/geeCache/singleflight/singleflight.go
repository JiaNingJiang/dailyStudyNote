package singleflight

import "sync"

// call 代表正在进行中，或已经结束的请求。
type call struct {
	wg  sync.WaitGroup // wg用于防止重入
	val interface{}
	err error
}

type Group struct {
	mu sync.Mutex       // protect m
	m  map[string]*call // key为请求资源的名称, value为对应的call
}

func (g *Group) Do(key string, fn func() (interface{}, error)) (interface{}, error) {
	g.mu.Lock()

	if g.m == nil { // 延迟初始化
		g.m = make(map[string]*call, 0)
	}

	if c, ok := g.m[key]; ok { // 防止同一时间对于同一个资源的多次重复请求
		g.mu.Unlock()
		c.wg.Wait() // 等待当前资源 call 请求结束
		return c.val, c.err
	}

	c := new(call)
	c.wg.Add(1)
	g.m[key] = c // 防止请求重入

	g.mu.Unlock()

	c.val, c.err = fn() // 执行fn()获取资源 -- wg的作用就是保证同一时刻针对同一个key的并发资源请求只能触发一次fn()的执行

	c.wg.Done() // 完成资源获取,可以解除资源重入等待组
	g.mu.Lock()
	delete(g.m, key) // 完成资源获取,可以解除资源重入等待组
	g.mu.Unlock()

	return c.val, c.err
}
