package fast

import (
	"container/list"
	"sync"
)
// 单个分片结构和LRU算法类似，并且是并发安全的，即加了一个字段sync.RWMutex
type cacheShard struct {
	locker sync.RWMutex
	// 最大存放entry个数
	maxEntries int
	// 当一个entry从缓存中移除时调用该回调函数，默认为nil
	// groupcache 中的 key 是任意的可比较类型：value 是 interface{}
	onEvicted func(key string, value interface{})

	ll *list.List
	cache map[string]*list.Element
}
type entry struct {
	key string
	value interface{}
}
// 创建一个新的cacheShard, 如果maxBytes是0，表示容量无限制
func newCacheShard(maxEntries int, onEvicted func(key string, value interface{})) *cacheShard {
	return &cacheShard{
		maxEntries: maxEntries,
		onEvicted: onEvicted,
		ll: list.New(),
		cache: make(map[string]*list.Element),
	}
}

func (c *cacheShard) get(key string) interface{} {
	c.locker.RLock()
	defer c.locker.RUnlock()

	if e, ok := c.cache[key]; ok {
		c.ll.MoveToBack(e)
		return e.Value.(*entry).value
	}
	return nil
}
func (c *cacheShard) set(key string, value interface{})  {
	c.locker.Lock()
	defer c.locker.Unlock()

	if e, ok := c.cache[key]; ok {
		// 移动到尾部
		c.ll.MoveToBack(e)
		// 获取之前的值
		en := e.Value.(*entry)
		en.value = value
		return
	}

	en := &entry{key, value}
	// 放入到队尾，返回对应节点指针
	e := c.ll.PushBack(en)
	// 存入map， value为节点指针，用于后续移动/删除等操作
	c.cache[key] = e
	if  c.maxEntries > 0 && c.ll.Len() > c.maxEntries {
		// 删除最早的记录
		c.removeElement(c.ll.Front())
	}
}

func (c *cacheShard) removeElement(e *list.Element) {
	if e == nil {
		return
	}

	c.ll.Remove(e)
	en := e.Value.(*entry)
	delete(c.cache, en.key)

	if c.onEvicted != nil {
		c.onEvicted(en.key, en.value)
	}
}
// 删除最旧的记录
func (c *cacheShard) delOldest() {
	c.locker.Lock()
	defer c.locker.Unlock()

	c.removeElement(c.ll.Front())
}

// len 返回当前 cache 中的记录数
func (c *cacheShard) len() int {
	c.locker.RLock()
	defer c.locker.RUnlock()

	return c.ll.Len()
}

// 主动删除某个缓存记录，首先根据key从map中获取节点，然后从链表中删除的同时，从mpa中删除
func (c *cacheShard) del(key string) {
	c.locker.Lock()
	defer c.locker.Unlock()

	if e, ok := c.cache[key]; ok {
		c.removeElement(e)
	}
}
