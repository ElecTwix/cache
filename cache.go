package cache

import (
	"sync"
	"time"
)

type Cache[K string | int, D any] struct {
	hash  map[K]CacheItem[K, D]
	tick  *time.Ticker
	mutex sync.Mutex
}

type CacheItem[K string | int, D any] struct {
	Life int64
	Data D
}

func NewCache[K string | int, D any](tick time.Ticker) *Cache[K, D] {
	cache := &Cache[K, D]{
		hash: make(map[K]CacheItem[K, D]),
		tick: &tick,
	}
	go cache.Loop()
	return cache
}

func (c *Cache[I, K]) Loop() {
	for {
		<-c.tick.C
		c.LifeCycle()
	}
}

func (c *Cache[I, K]) LifeCycle() {
	c.mutex.Lock()
	unixtime := time.Now().Unix()
	for key, item := range c.hash {
		// If item.Life is 0, it will never be deleted
		if item.Life == 0 {
			continue
		}
		if item.Life < unixtime {
			delete(c.hash, key)
		}
	}
	c.mutex.Unlock()
}

func (c *Cache[I, K]) Len() int {
	return len(c.hash)
}

func (c *Cache[I, K]) Clear() {
	clear(c.hash)
}

func (c *Cache[K, D]) AddTime(key K, duration time.Duration) {
	item, ok := c.hash[key]
	if ok {
		item.Life = item.Life + int64(duration.Seconds())
		c.hash[key] = item
	}
}

func (c *Cache[K, D]) Set(key K, data D, duration time.Duration) {
	c.mutex.Lock()
	item := CacheItem[K, D]{
		Life: time.Now().Add(duration).Unix(),
		Data: data,
	}
	defer c.mutex.Unlock()
	c.hash[key] = item
}

func (c *Cache[K, D]) Get(key K) (D, bool) {
	item, ok := c.hash[key]
	return item.Data, ok
}

func (c *Cache[K, D]) Delete(key K) {
	delete(c.hash, key)
}
