package cache

import (
	"fmt"
	"log"
	"runtime"
	"sync"
)

const DefaultMaxBytes = 1 << 29

type Cache interface {
	// 增加缓存，key则替换
	Set(key string, value interface{})
	// 通过key获取缓存值
	Get(key string) interface{}
	// 通过key删除缓存值
	Del(key string)
	// 删除最“无用的”一个缓存值
	DelOldest()
	// 获取缓存已存在的记录数
	Len() int
}

type Value interface {
	Len() int
}

func CalcLen(value interface{}) int {
	var n int
	switch v := value.(type) {
	case Value:
		n = v.Len()
	case string:
		if runtime.GOARCH == "amd64" {
			n = 16 + len(v)
		} else {
			n = 8 + len(v)
		}
	case bool, uint8, int8:
		n = 1
	case int16, uint16:
		n = 2
	case int32, uint32, float32:
		n = 4
	case int64, uint64, float64:
		n = 8
	case int, uint:
		if runtime.GOARCH == "amd64" {
			n = 8
		} else {
			n = 4
		}
	case complex64:
		n = 8
	case complex128:
		n = 16
	default:
		panic(fmt.Sprintf("%T is not implemente cache.Value", value))
	}
	return n
}

type safeCache struct {
	m       sync.RWMutex
	cache   Cache
	// 为了方便统计，在safeCache结构中我们定义了nget和nhit，用来记录缓存获取次数和命中次数。
	nHit, nGet int
}

func newSafeCache(cache Cache) *safeCache {
	return &safeCache{
		cache: cache,
	}
}

func (sc *safeCache) set(key string, value interface{})  {
	sc.m.Lock()
	defer sc.m.Unlock()
	sc.cache.Set(key, value)
}

func (sc *safeCache) get(key string) interface{} {
	sc.m.RLock()
	defer sc.m.RUnlock()
	sc.nGet++
	if sc.cache == nil {
		return nil
	}
	v := sc.cache.Get(key)
	if v != nil {
		log.Println("[TourCache] hit")
		sc.nHit++
	}
	return v
}

func (sc *safeCache) stat() *Stat  {
	sc.m.RLock()
	defer sc.m.RUnlock()
	return &Stat{
		NHit: sc.nHit,
		NGet: sc.nGet,
	}
}

type Stat struct {
	NHit, NGet int
}
