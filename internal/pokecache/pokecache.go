package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache_map map[string]cacheEntry
	duration  time.Duration
	mu        sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(max_dur time.Duration) *Cache {
	var c Cache
	ce := make(map[string]cacheEntry)
	c.cache_map = ce
	c.duration = max_dur

	return &c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	new_entry := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.cache_map[key] = new_entry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	val, ok := c.cache_map[key]
	if !ok {
		return nil, false
	}
	return val.val, true
}

func (c *Cache) reaploop() {
	c.mu.Lock()
	defer c.mu.Unlock()
	for key, val := range c.cache_map {
		if c.duration < time.Since(val.createdAt) {
			delete(c.cache_map, key)
		}
	}
}
