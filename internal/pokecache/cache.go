package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	Entries  map[string]cacheEntry
	mu       sync.Mutex
	interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		Entries:  make(map[string]cacheEntry),
		interval: interval,
	}
	go c.reapLoop()
	return c
}

func (c *Cache) Add(key string, value []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       value,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	value, exists := c.Entries[key]
	if exists {

		return value.val, exists
	}
	return nil, exists

}

func (c *Cache) reapLoop() {
	t := time.NewTicker(c.interval)
	defer t.Stop()
	for {
		<-t.C
		c.mu.Lock()
		for key, entry := range c.Entries {
			if entry.createdAt.Before(time.Now().Add(-c.interval)) {
				delete(c.Entries, key)
			}
		}
		c.mu.Unlock()
	}
}
