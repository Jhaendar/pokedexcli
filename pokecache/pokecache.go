// pokecache MODULE
// Handls the caching of pokeapi module

package pokecache

import (
	"sync"
	"time"
)

type CacheEntry struct {
	createdAt time.Time
	value     []byte
}

type Cache struct {
	Cache    map[string]CacheEntry
	duration time.Duration
	mutex    sync.RWMutex
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	entry, ok := c.Cache[key]
	if !ok {
		return nil, false
	}
	return entry.value, true
}

func (c *Cache) Add(key string, value []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.Cache[key] = CacheEntry{
		createdAt: time.Now(),
		value:     value,
	}
}

func (c *Cache) Delete(key string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	delete(c.Cache, key)
}

func (c *Cache) Clear() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.Cache = make(map[string]CacheEntry)
}

func (c *Cache) clearExpired() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	for key, entry := range c.Cache {
		if time.Since(entry.createdAt) > c.duration {
			delete(c.Cache, key)
		}
	}
}

func (c *Cache) StartCleanup() {
	tickerChan := time.NewTicker(c.duration).C

	go func() {
		for range tickerChan {
			c.clearExpired()
		}
	}()
}

func NewCache(duration time.Duration) *Cache {
	cache := Cache{
		Cache:    make(map[string]CacheEntry),
		duration: duration,
	}

	cache.StartCleanup()
	return &cache
}
