package cache_test

import (
	"github.com/learn-go/cache_example/cache"
	"github.com/learn-go/cache_example/lru"
	"github.com/matryer/is"
	"log"
	"sync"
	"testing"
)

func TestTourCacheGet(t *testing.T) {
	db := map[string]string{
		"k1": "val1",
		"k2": "val2",
		"k3": "val3",
		"k4": "val4",
	}
	getter := cache.GetFunc(func(key string) interface{} {
		log.Print("[from db] find key", key)

		if val, ok := db[key]; ok {
			return val
		}
		return nil
	})
	tourCache := cache.NewTourCache(getter, lru.New(0, nil))

	i := is.New(t)

	var wg sync.WaitGroup

	for k, v := range db {
		wg.Add(1)
		go func(k, v string) {
			defer wg.Done()
			i.Equal(tourCache.Get(k), v)
			i.Equal(tourCache.Get(k), v)
		}(k, v)
	}
	wg.Wait()

	i.Equal(tourCache.Get("unknow"), nil)
	i.Equal(tourCache.Get("unknow"), nil)

	i.Equal(tourCache.Stat().NGet, 10)
	i.Equal(tourCache.Stat().NHit, 4)
}
