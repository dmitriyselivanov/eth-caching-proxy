package cache

import (
	"eth-caching-proxy/cache/simplelru"
)

type Cache struct {
	BlockCache BlockCache
}

func New() *Cache {
	return &Cache{
		BlockCache: simplelru.NewBlockCache(),
	}
}
