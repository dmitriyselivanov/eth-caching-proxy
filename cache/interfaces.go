package cache

import (
	"eth-caching-proxy/model"
	"math/big"
)

type BlockCache interface {
	AddBlock(block *model.Block)
	GetBlock(blockNumber *big.Int) (*model.Block, bool)
}
