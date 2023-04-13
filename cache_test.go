package cache_test

import (
	"testing"
	"time"

	"github.com/ElecTwix/cache"
)

func TestCache(t *testing.T) {
	abc := cache.NewCache[any, string](*time.NewTicker(1 * time.Second))
	abc.Set("a", "b", 1*time.Second)
	raw, ok := abc.Get("a")
	_, castok := raw.(string)

	if !ok || !castok || raw != "b" {
		t.Error("Cache not working")
	}
	time.Sleep(5 * time.Second)
	raw, ok = abc.Get("a")
	if ok {
		t.Error("Timeout Cache not working")
	}
}

func BenchmarkSet(b *testing.B) {
	abc := cache.NewCache[any, string](*time.NewTicker(30 * time.Second))
	for i := 0; i < b.N; i++ {
		abc.Set("a", 5, time.Minute)
	}
}

func BenchmarkSetGetWithAny(b *testing.B) {
	abc := cache.NewCache[any, string](*time.NewTicker(30 * time.Second))
	for i := 0; i < b.N; i++ {
		abc.Set("a", "b", time.Minute)
		abc.Get("a")
	}
}

func BenchmarkSetGetWithInt(b *testing.B) {
	abc := cache.NewCache[int, string](*time.NewTicker(30 * time.Second))
	for i := 0; i < b.N; i++ {
		abc.Set("a", 5, time.Minute)
		abc.Get("a")
	}
}
