package lru

import (
	"container/list"
	"github.com/learn-go/cache_example/cache"
)

// 相对于仅考虑时间因素的FIFO算法和仅考虑访问频率的LFU算法，LRU（最近最少使用）算法可以认为是相对平衡的一种淘汰算法。
// LRU算法的核心原则是：如果数据最近被访问过，那么将来被访问的概率会更高。
// LRU算法的实现非常简单，即维护一个队列，如果某条数据被访问了，则把这条数据移到队尾，队首则是最近最少使用的数据，淘汰队首数据即可。
// 该算法的核心数据结构和 FIFO 算法是一样的，只是数据的移动方式不同。

type lru struct {
	// 缓存最大的容量，单位字节
	// groupcache 使用的是最大存放entry个数
	maxBytes   int
	// 当一个 entry 从缓存中移除时调用该回调函数，默认为nil
	// groupcache 中的 key 是任意的可比较类型； value 是 interface{}
	onEvicted  func(key string, value interface{})
	// 已使用的字节数，只包括值，key 不算
	usedBytes  int
	ll         *list.List
	// key为缓存的key，值是双向链表中对应的节点指针
	cache      map[string]*list.Element
}

type entry struct {
	key string
	value interface{}
}
func New(maxBytes int, onEvicted func(key string, value interface{})) cache.Cache {
	return &lru{
		maxBytes: maxBytes,
		onEvicted: onEvicted,
		ll: list.New(),
		cache: make(map[string]*list.Element),
	}
}
func (e *entry) Len() int {
	return cache.CalcLen(e.value)
}


func (l *lru) Set(key string, value interface{})  {
	if e, ok := l.cache[key]; ok {
		// 移动到尾部
		l.ll.MoveToBack(e)
		// 获取之前的值
		en := e.Value.(*entry)
		l.usedBytes = l.usedBytes - cache.CalcLen(en.value) + cache.CalcLen(value)
		en.value = value
		return
	}

	en := &entry{key, value}
	// 放入到队尾，返回对应节点指针
	e := l.ll.PushBack(en)
	// 存入map， value为节点指针，用于后续移动/删除等操作
	l.cache[key] = e

	l.usedBytes += en.Len()
	if l.maxBytes > 0 && l.usedBytes > l.maxBytes {
		l.DelOldest()
	}
}

func (l *lru) Get(key string) interface{}  {
	if e, ok := l.cache[key]; ok {
		// 只要有访问就移动到队尾
		l.ll.MoveToBack(e)
		return e.Value.(*entry).value
	}
	return nil
}

func (l *lru) removeElement(e *list.Element)  {
	if e == nil {
		return
	}
	// 从队列中移除当前元素
	l.ll.Remove(e)
	en := e.Value.(*entry)
	l.usedBytes -= en.Len()
	// 从cache map中删除
	delete(l.cache, en.key)

	if l.onEvicted != nil {
		l.onEvicted(en.key, en.value)
	}
}

// 从cache中删除最旧(早)的记录。 一般不主动调用，在内存满时自动触发，这就是缓存淘汰
func (l *lru) DelOldest() {
	l.removeElement(l.ll.Front())
}

// 主动删除某个缓存记录，首先根据key从map中获取节点，然后从链表中删除的同时，从mpa中删除
func (l *lru) Del(key string)  {
	if e, ok := l.cache[key]; ok {
		l.removeElement(e)
	}
}

// 这个方法更多的是为了方便测试或提供数据统计
func (l *lru) Len() int {
	return l.ll.Len()
}
