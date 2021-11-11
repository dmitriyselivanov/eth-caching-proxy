package cache

import (
	"eth-caching-proxy/cache/simplelru"
)

// Cache is a thread-safe fixed size LRU cache.
type Cache struct {
	BlockCache BlockCache
}

// New creates a new cache instance
func New() *Cache {
	return &Cache{
		BlockCache: simplelru.NewBlockCache(),
	}
}
