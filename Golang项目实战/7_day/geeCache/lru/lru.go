package lru

import "container/list"

// LRU算法实现的缓存
type Cache struct {
	maxBytes int // 缓存的最大字节容量
	nowBytes int // 缓存的当前字节数

	ll    *list.List               // 缓存数据的链表
	cache map[string]*list.Element // key集合与缓存链表中value的映射

	OnEvicted func(key string, value Value) //某条记录被移除时的回调函数，可以为 nil
}

// 缓存链表的每一个节点保存的数据
type entry struct {
	key   string
	value Value
}

// Value可以是实现Len()方法的任意数据类型
type Value interface {
	Len() int
}

// New一个新的Cache缓存
func NewCache(maxBytes int, OnEvicted func(string, Value)) *Cache {
	return &Cache{
		maxBytes:  maxBytes,
		ll:        list.New(),
		cache:     make(map[string]*list.Element),
		OnEvicted: OnEvicted,
	}
}

// 查找数据，并将其移动到链表尾部
func (c *Cache) Get(key string) (value Value, ok bool) {
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToBack(ele) // 移动到ll尾部
		kv := ele.Value.(*entry)
		return kv.value, true // 返回value
	}
	return
}

// 移除最近最少访问的数据
func (c *Cache) RemoveOldest() {
	ele := c.ll.Front() // 返回ll头部节点，就是最近最少访问的数据
	if ele != nil {
		kv := ele.Value.(*entry)
		delete(c.cache, kv.key) // 将记录先从cacheMap中删去
		c.ll.Remove(ele)        // 再从链表中删除

		c.nowBytes -= len(kv.key) + kv.value.Len()

		if c.OnEvicted != nil {
			c.OnEvicted(kv.key, kv.value)
		}
	}
}

// 新增/修改缓存数据
func (c *Cache) Add(key string, value Value) {
	if ele, ok := c.cache[key]; ok { // 数据已经存在，那么就是修改
		c.ll.MoveToBack(ele) // 先将节点移动到ll尾部

		kv := ele.Value.(*entry)
		c.nowBytes += value.Len() - kv.value.Len() // 更新该节点保存的value
		kv.value = value
	} else { // 是新增数据
		kv := &entry{key: key, value: value}
		ele := c.ll.PushBack(kv) // 先创建一个新节点(在尾部添加)存储entry数据
		c.cache[key] = ele       // 在cacheMap记录新的关系映射

		c.nowBytes += kv.value.Len() + len(kv.key)
	}

	for c.maxBytes != 0 && c.maxBytes < c.nowBytes { // 检查是否超出缓存上限
		c.RemoveOldest()
	}
}

// 返回当前总共存储了多少key-value对
func (c *Cache) Len() int {
	return c.ll.Len()
}
