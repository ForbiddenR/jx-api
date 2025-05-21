//go:build go1.24

package cache

import (
	"runtime"
	"sync"
	"weak"
)

type Cache[K comparable, V any] struct {
	cache sync.Map
}

func NewCache[K comparable, V any]() *Cache[K, V] {
	return &Cache[K, V]{}
}

func (c *Cache[K, V]) Set(key K, value V) {
	wp := weak.Make(&value)
	_, loaded := c.cache.LoadOrStore(key, wp)
	if !loaded {
		runtime.AddCleanup(&value, func(key K) {
			c.cache.Delete(key)
		}, key)
	}
}

func (c *Cache[K, V]) Get(key K) (v V, found bool) {
	if wp, found := c.cache.Load(key); found {
		if v := wp.(weak.Pointer[V]).Value(); v != nil {
			return *v, true
		}
	}
	return
}
