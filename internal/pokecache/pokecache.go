package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache_map map[string]cacheEntry
	interval  time.Duration
	mu        sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	var c Cache
	ce := make(map[string]cacheEntry)
	c.cache_map = ce
	c.interval = interval
	go c.reaploop(c.interval)
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

func (c *Cache) reaploop(interval time.Duration) {
	tik := time.NewTicker(interval)
	for now := range tik.C {
		c.mu.Lock()
		time_to_delete := now.UTC().Add(-interval)
		for k, v := range c.cache_map {
			if v.createdAt.Before(time_to_delete) {
				delete(c.cache_map, k)
			}
		}
		c.mu.Unlock()
	}
}
