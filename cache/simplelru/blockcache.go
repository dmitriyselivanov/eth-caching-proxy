package simplelru

import (
	"eth-caching-proxy/cache"
	"eth-caching-proxy/config"
	"eth-caching-proxy/model"
	"github.com/hashicorp/golang-lru"
	"log"
	"math/big"
)

type blockCache struct {
	cache *lru.Cache
}

// NewBlockCache creates an instance of blockCache
func NewBlockCache() cache.BlockCache {
	conf := config.GetConfig()

	c, err := lru.New(conf.Cache.MaxBlocks)
	if err != nil {
		log.Fatalln("cannot create lru cache")
	}

	return &blockCache{cache: c}
}

func (c *blockCache) AddBlock(block *model.Block) {
	c.cache.Add(block.Number.String(), block)
}

func (c *blockCache) GetBlock(blockNumber *big.Int) (*model.Block, bool) {
	block, ok := c.cache.Get(blockNumber.String())
	if !ok {
		return nil, false
	}

	return block.(*model.Block), true
}
