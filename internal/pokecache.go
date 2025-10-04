package internal

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	Val       []byte
}

var sharedInterval time.Duration

type Cache struct {
	Map map[string]cacheEntry
	mu  sync.Mutex
}

func NewCache(interval time.Duration) *Cache {
	cache := Cache{
		Map: make(map[string]cacheEntry),
	}
	sharedInterval = interval
	go cache.reapLoop(interval) // IS IT EVEN NECESSARY??????
	return &cache
}

func (c *Cache) reapLoop(interval time.Duration) {
	for key, cEntry := range c.Map {
		lifetime := time.Since(cEntry.createdAt)

		if lifetime > interval {
			c.mu.Lock()
			delete(c.Map, key)
			c.mu.Unlock()
		}
	}
}

func (c *Cache) Add(key string, val []byte) {
	c.reapLoop(sharedInterval)
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Map[key] = cacheEntry{
		createdAt: time.Now(),
		Val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.reapLoop(sharedInterval)
	c.mu.Lock()
	defer c.mu.Unlock()
	if cEntry, ok := c.Map[key]; ok {
		return cEntry.Val, ok
	} else {
		return nil, false
	}
}
