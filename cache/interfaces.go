package cache

import (
	"eth-caching-proxy/model"
	"math/big"
)

// BlockCache provides methods to get and add block to cache
type BlockCache interface {
	AddBlock(block *model.Block)
	GetBlock(blockNumber *big.Int) (*model.Block, bool)
}
