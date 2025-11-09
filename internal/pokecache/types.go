package pokecache

import (
	"time"
	"sync"
)

type Cache struct {
	infoCached		map[string]cacheEntry
	lock			sync.Mutex
	interval 		time.Duration
}

type cacheEntry struct {
	createdAt	time.Time
	val			[]byte
}