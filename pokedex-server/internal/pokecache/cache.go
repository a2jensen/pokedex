package pokecache

/**
- use struct pointers for struct functions if you need to edit it directly
- make() and delete()...refer to the map lib functions for editing!
- why do i not need an interface for this... but i do need
*/

import (
	"sync"
	"time"
	//"fmt"
)

type PokeCache interface {
	Add(key string, val	[]byte)
	Get(key string) (val []byte, found bool)
}


func NewCache(interval time.Duration) *Cache {
	m := make(map[string]cacheEntry)
	cache := &Cache {
		infoCached: m,
		lock : sync.Mutex{},
		interval: interval,
	}

	go cache.reapLoop()

	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.lock.Lock()
	entry := cacheEntry{
		createdAt: time.Now(),
		val: val,
	}
	c.infoCached[key] = entry
	c.lock.Unlock()
}

func (c *Cache) Get(key string) (response []byte, found bool) {
	c.lock.Lock()
	entry, ok := c.infoCached[key]
	c.lock.Unlock()
	if !ok {
		return nil, false
	}
	return entry.val, true
}

func (c *Cache) reapLoop() {
	//fmt.Println("1")
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()
	for range ticker.C { // loops every tick
		//fmt.Println("reaper tick")
		c.lock.Lock()
		for key , entry := range c.infoCached {
			if time.Since(entry.createdAt) > c.interval {
				delete(c.infoCached, key)
			}
		}
		c.lock.Unlock()
	}
}
