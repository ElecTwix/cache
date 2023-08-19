package cache_test

import (
	"testing"
	"time"

	"github.com/ElecTwix/cache"
)

func TestCache(t *testing.T) {
	tempCache := cache.NewCache[string, any](*time.NewTicker(1 * time.Second))
	tempCache.Set("a", "b", 1*time.Second)
	raw, ok := tempCache.Get("a")
	_, castok := raw.(string)

	if !ok || !castok || raw != "b" {
		t.Error("Cache not working")
	}

	time.Sleep(5 * time.Second)
	raw, ok = tempCache.Get("a")
	if ok || raw != nil {
		t.Error("Timeout Cache not working")
	}
}

func BenchmarkSet(b *testing.B) {
	abc := cache.NewCache[string, any](*time.NewTicker(30 * time.Second))
	for i := 0; i < b.N; i++ {
		abc.Set("a", 5, time.Minute)
	}
}

func BenchmarkSetGetWithAny(b *testing.B) {
	abc := cache.NewCache[string, any](*time.NewTicker(30 * time.Second))
	for i := 0; i < b.N; i++ {
		abc.Set("a", "b", time.Minute)
		abc.Get("a")
	}
}

func BenchmarkSetGetWithInt(b *testing.B) {
	abc := cache.NewCache[string, any](*time.NewTicker(30 * time.Second))
	for i := 0; i < b.N; i++ {
		abc.Set("a", 5, time.Minute)
		abc.Get("a")
	}
}

func BenchmarkGetAndSetByte(b *testing.B) {
	abc := cache.NewCache[string, []byte](*time.NewTicker(30 * time.Second))
	for i := 0; i < b.N; i++ {
		abc.Set("a", []byte("b"), time.Minute)
		abc.Get("a")
	}
}
