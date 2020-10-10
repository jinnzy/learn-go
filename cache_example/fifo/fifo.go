package fifo

import (
	"container/list"
	"github.com/learn-go/cache_example/cache"
)

// FIFO（先进先出）算法是淘汰缓存中最早添加的记录。
// 在FIFO Cache设计中，其核心原则是：如果一个数据最先进入缓存，那么也应该最先被删掉。
// 依据是，最早进入缓存的数据其不再被使用的可能性比刚进入缓存的数据要大。

type fifo struct {
	// 缓存最大的容量，单位字节
	// groupcache 使用的是最大存放entry个数
	maxBytes   int
	// 当一个 entry 从缓存中移除时调用该回调函数，默认为nil
	// groupcache 中的 key 是任意的可比较类型； value 是 interface{}
	onEvicted  func(key string, value interface{})
	// 已使用的字节数，只包括值，key 不算
	usedBytes  int
	ll         *list.List
	// 值是双向链表中对应的节点指针
	cache      map[string]*list.Element
}

type entry struct {
	key string
	value interface{}
}

func (e *entry) Len() int {
	return cache.CalcLen(e.value)
}

func New(maxBytes int, onEvicted func(key string, value interface{})) cache.Cache {
	return &fifo{
		maxBytes: maxBytes,
		onEvicted: onEvicted,
		ll: list.New(),
		cache: make(map[string]*list.Element),
	}
}

func (f *fifo) Get(key string) interface{} {
	if e, ok := f.cache[key]; ok {
		return e.Value.(*entry).value
	}
	return nil
}
// 通过set方法往 cache 尾部增加一个元素(如果已经存在，则移到尾部，并修改值)
func (f *fifo) Set(key string, value interface{})  {
	if e, ok := f.cache[key]; ok {
		// 移动到尾部
		f.ll.MoveToBack(e)
		// 获取之前的值
		en := e.Value.(*entry)
		f.usedBytes = f.usedBytes - cache.CalcLen(en.value) + cache.CalcLen(value)
		en.value = value
		return
	}

	en := &entry{key, value}
	// 放入到队尾，返回对应节点指针
	e := f.ll.PushBack(en)
	// 存入map， value为节点指针，用于后续移动/删除等操作
	f.cache[key] = e

	f.usedBytes += en.Len()
	if f.maxBytes > 0 && f.usedBytes > f.maxBytes {
		f.DelOldest()
	}
}

func (f *fifo) removeElement(e *list.Element)  {
	if e == nil {
		return
	}
	f.ll.Remove(e)
	en := e.Value.(*entry)
	f.usedBytes -= en.Len()
	delete(f.cache, en.key)

	if f.onEvicted != nil {
		f.onEvicted(en.key, en.value)
	}
}

// 从cache中删除最旧(早)的记录。 一般不主动调用，在内存满时自动触发，这就是缓存淘汰
func (f *fifo) DelOldest() {
	f.removeElement(f.ll.Front())
}


// 主动删除某个缓存记录，首先根据key从map中获取节点，然后从链表中删除的同时，从mpa中删除
func (f *fifo) Del(key string)  {
	if e, ok := f.cache[key]; ok {
		f.removeElement(e)
	}
}
// 这个方法更多的是为了方便测试或提供数据统计
func (f *fifo) Len() int {
	return f.ll.Len()
}
