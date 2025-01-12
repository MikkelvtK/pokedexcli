package pokecache

import (
	"sync"
	"time"
)

type entry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	entries map[string]entry
	mut     *sync.Mutex
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		entries: map[string]entry{},
		mut:     &sync.Mutex{},
	}
	go c.reapLoop(interval)
	return c
}

func (c *Cache) Add(k string, val []byte) {
	c.mut.Lock()
	c.entries[k] = entry{
		val:       val,
		createdAt: time.Now().UTC(),
	}
	c.mut.Unlock()
}

func (c *Cache) Get(k string) ([]byte, bool) {
	c.mut.Lock()
	e, ok := c.entries[k]
	c.mut.Unlock()
	return e.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for t := range ticker.C {
		c.mut.Lock()
		for k, e := range c.entries {
			if e.createdAt.Before(t.UTC()) {
				delete(c.entries, k)
			}
		}
		c.mut.Unlock()
	}
}
