package lfu

import (
	"github.com/matryer/is"
	"testing"
)

func TestSet(t *testing.T)  {
	i := is.New(t)

	cache := New(24, nil)
	cache.DelOldest()
	cache.Set("k1", 1)
	v := cache.Get("k1")
	i.Equal(v, 1)

	cache.Del("k1")
	i.Equal(0, cache.Len())
}

func TestOnEvicted(t *testing.T)  {
	i := is.New(t)

	keys := make([]string, 0, 8)
	onEvicted := func(key string, value interface{}) {
		keys = append(keys, key)
	}
	cache := New(32, onEvicted)
	cache.Set("k1", 1)
	cache.Set("k2", 2)
	//cache.Get("k1")
	//cache.Get("k1")
	//cache.Get("k2")
	cache.Set("k3", 3)
	cache.Set("k4", 4)

	// 为什么淘汰的是k1 k3，
	// 因为每次淘汰都是堆顶的数据，第一次淘汰时排序为[k1, k2, k3]k1淘汰后，会把最后一个元素(k3)放到堆顶，然后进行下沉，
	// 现在排序为[k3, k2]。增加新数据时，这次就会淘汰掉k3，再把新数据置换到0位置进行下沉操作。
	expected := []string{"k1", "k3"}
	cache.Get("k2")

	i.Equal(expected, keys)
	i.Equal(2, cache.Len())
}
