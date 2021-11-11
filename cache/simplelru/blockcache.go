package simplelru

import (
	"eth-caching-proxy/config"
	"eth-caching-proxy/model"
	"github.com/hashicorp/golang-lru"
	"log"
	"math/big"
)

// LruBlockCache is a golang-lru cache for ethereum blocks
type LruBlockCache struct {
	cache *lru.Cache
}

// NewBlockCache creates an instance of LruBlockCache
func NewBlockCache() *LruBlockCache {
	conf := config.GetConfig()

	c, err := lru.New(conf.Cache.MaxBlocks)
	if err != nil {
		log.Fatalln("cannot create lru cache")
	}

	return &LruBlockCache{cache: c}
}

// AddBlock adds a block to a cache
func (c *LruBlockCache) AddBlock(block *model.Block) {
	c.cache.Add(block.Number.String(), block)
}

// GetBlock returns a block from cache and ok bool
func (c *LruBlockCache) GetBlock(blockNumber *big.Int) (*model.Block, bool) {
	block, ok := c.cache.Get(blockNumber.String())
	if !ok {
		return nil, false
	}

	return block.(*model.Block), true
}
