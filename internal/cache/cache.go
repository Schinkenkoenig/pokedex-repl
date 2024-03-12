package cache

import (
	"sync"
	"time"
)

type Cache struct {
	cache map[string]cacheEntry
	mu    sync.RWMutex
}

type cacheEntry struct {
	created time.Time
	val     []byte
}

func NewCache(duration time.Duration) *Cache {
	cache := Cache{
		cache: make(map[string]cacheEntry),
		mu:    sync.RWMutex{},
	}

	go cache.reapLoop(duration)
	return &cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cache[key] = cacheEntry{
		created: time.Now(),
		val:     val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if ele, ok := c.cache[key]; ok {
		return ele.val, true
	}

	return nil, false
}

func (c *Cache) reapLoop(duration time.Duration) {
	ch := time.Tick(duration)

	for range ch {
		func() {
			c.mu.Lock()
			defer c.mu.Unlock()
			for key, ce := range c.cache {
				if time.Now().After(ce.created.Add(duration)) {
					delete(c.cache, key)
				}
			}
		}()
	}
}
