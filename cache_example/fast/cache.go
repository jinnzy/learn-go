package fast

type fastCache struct {
	// 包含所有分片得切片，根据BigCache得经验，1024长度是比较理想的
	shards []*cacheShard
	// 既然分片可以很好地降低锁的竞争，那么N是不是越大越好呢？当然不是，如果N非常大。比如每个缓存对象都有一个锁，那么会带来很多不必要的开销。可以选择一个不太大的值，在性能和花销上寻找一个平衡。
	// 另外，N应该是2的幂，比如16、32、64等。这样设计的好处是，在计算余数时可以使用位运算快速计算：
	// shardMask 就是 N，bigcache中默认为1024
	shardMask uint64
	// 实现hash的算法，用的fnv64-a，bigcache中是使用的接口。
	hash fnv64a
}

func NewFastCache(maxEntries, shardsNum int, onEvicted func(key string, value interface{})) *fastCache {
	fastCache := &fastCache{
		hash: newDefaultHasher(),
		shards: make([]*cacheShard, shardsNum),
		shardMask: uint64(shardsNum - 1),
	}
	// 初始化分片中每个cache shard
	for i := 0; i < shardsNum; i++ {
		fastCache.shards[i] = newCacheShard(maxEntries, onEvicted)
	}
	return fastCache
}

func (c *fastCache) getShard(key string) *cacheShard {
	hashedKey := c.hash.Sum64(key)
	return c.shards[hashedKey & c.shardMask]
}

func (c *fastCache) Set(key string, value interface{})  {
	c.getShard(key).set(key, value)
}

func (c *fastCache) Get(key string) interface{} {
	return c.getShard(key).get(key)
}

func (c *fastCache) Del(key string)  {
	c.getShard(key).del(key)
}

func (c *fastCache) Len() int {
	length := 0
	for _, shard := range c.shards {
		length += shard.len()
	}
	return length
}
