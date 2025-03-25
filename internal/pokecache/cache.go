package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	rw      sync.RWMutex
	entries map[string]cacheEntry
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		entries: make(map[string]cacheEntry),
	}
	go cache.reapLoop(interval)
	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.rw.Lock()
	defer c.rw.Unlock()
	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) (val []byte, exists bool) {
	c.rw.RLock()
	defer c.rw.RUnlock()
	entry, exists := c.entries[key]
	if exists {
		val = entry.val
	}
	return val, exists
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case t := <-ticker.C:
			c.pruneCache(t, interval)
		default:
			continue
		}
	}
}

func (c *Cache) pruneCache(now time.Time, interval time.Duration) {
	c.rw.Lock()
	defer c.rw.Unlock()

	for key, val := range c.entries {
		if now.Add(-interval).After(val.createdAt) {
			delete(c.entries, key)
		}
	}
}
