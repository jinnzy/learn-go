package lfu

import (
	"container/heap"
	"github.com/learn-go/cache_example/cache"
)

// 这里的entry比fifo的多了两个字段weight和index
type entry struct {
	key string
	value interface{}
	// weight表示该entry在queue中的权重(优先级),被访问次数越多，权重越高
	weight int
	// index表示该entry在堆中(heap)的索引
	index int
}

func (e *entry) Len() int  {
	return cache.CalcLen(e.value) + 4 + 4
}

// LFU算法用最小堆实现。在Go中，通过标准库container/heap来实现最小堆，要求queue实现heap.Interface接口
type queue []*entry

func (q queue) Len() int {
	return len(q)
}

//实现Interface对应的sort.Interface中的Less
// '<' 是最小堆，'>' 是最大堆
func (q queue) Less(i, j int) bool {
	return q[i].weight < q[j].weight
}

func (q queue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
	q[i].index = i
	q[j].index = j
}

// append ，*q = oldQue[:n-1] 会导致频繁的内存拷贝
// 实际上，如果使用 LFU算法，处于性能考虑，可以将最大内存限制修改为最大记录数限制
// 这样提前分配好 queue 的容量，再使用交换索引和限制索引的方式来实现 Pop 方法，可以免去频繁的内存拷贝，极大提高性能
func (q *queue) Push(x interface{})  {
	n := len(*q)
	en := x.(*entry)
	en.index = n
	*q = append(*q, en) // 这里会重新分配内存，并拷贝数据
}

func (q *queue) Pop() interface{}  {
	old := *q
	n := len(old)
	en := old[n-1]
	old[n-1] = nil // 将不再使用的对象置为nil，加快垃圾回收，避免内存泄漏
	en.index = -1 //  for safety
	*q = old[0 : n-1] // 这里会重新分配内存，并拷贝数据
	return en
}

// weight更新后，要重新排序，时间复杂度为 O(logN)
func (q *queue) update(en *entry, value interface{}, weight int)  {
	en.value = value
	en.weight = weight
	// 重建堆/重新排序
	// 分析思路是把 堆(大D) 的树状图画出来，看成一个一个小的堆(小D)，看改变其中一个值，对 大D 有什么影响
	// 可以得出结论，下沉操作和上沉操作分别执行一次能将 queue 排列为堆
	heap.Fix(q, en.index)
}
