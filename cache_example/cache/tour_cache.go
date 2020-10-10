package cache

type Getter interface {
	Get(key string) interface{}
}

// 默认实现
type GetFunc func(key string) interface{}

func (f GetFunc) Get(key string) interface{}  {
	return f(key)
}

// TourCache结构体包含两个字段：mainCache是并发安全的缓存实现；getter是回调，用于缓存未命中时从数据源获取的数据。
type TourCache struct {
	mainCache *safeCache
	getter Getter
}

func NewTourCache(getter Getter, cache Cache) *TourCache {
	return &TourCache{
		mainCache: newSafeCache(cache),
		getter:    getter,
	}
}

// Get方法：先从缓存获取数据，如果不存在，则调用回调函数获取数据，并将数据写入缓存，最后返回获取的数据。
func (t *TourCache) Get(key string) interface{} {
	val := t.mainCache.get(key)
	if val != nil {
		return val
	}
	if t.getter != nil {
		val := t.getter.Get(key)
		if val == nil {
			return nil
		}
		t.mainCache.set(key, val)
		return val
	}
	return nil
}

func (t *TourCache) Stat() *Stat {
	return t.mainCache.stat()
}

func (t *TourCache) Set(key string, val interface{})  {
	if val == nil {
		return
	}
	t.mainCache.set(key, val)
}
