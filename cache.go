package cache

import (
	"time"
)

type Cache[I any, K string | int] struct {
	hash map[K]CacheItem[I, K]
	tick *time.Ticker
}

type CacheItem[I any, K string | int] struct {
	Life int64
	Data I
}

func NewCache[I any, K string | int](tick time.Ticker) *Cache[I, K] {
	cache := &Cache[I, K]{
		hash: make(map[K]CacheItem[I, K]),
		tick: &tick,
	}
	go cache.Loop()
	return cache
}

func (c Cache[I, K]) Loop() {
	for {
		select {
		case <-c.tick.C:
			c.LifeCycle()
		}
	}
}

func (c Cache[I, K]) LifeCycle() {
	unixtime := time.Now().Unix()
	for key, item := range c.hash {
		if item.Life < unixtime {
			delete(c.hash, key)
		}
	}
}

func (c Cache[I, K]) Set(key K, data I, duration time.Duration) {
	item := CacheItem[I, K]{
		Life: time.Now().Add(duration).Unix(),
		Data: data,
	}
	c.hash[key] = item
}

func (c Cache[I, K]) Get(key K) (I, bool) {
	item, ok := c.hash[key]
	return item.Data, ok
}

func (c Cache[I, K]) Delete(key K) {
	delete(c.hash, key)
}
