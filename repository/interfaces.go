package repository

import (
	"eth-caching-proxy/model"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

type BlockRepository interface {
	BlockByNumber(number *big.Int) (*model.Block, error)
	LatestBlock() (*model.Block, error)

	LatestBlockHeader() (*types.Header, error)
	BlockHeaderByNumber(blockNumber *big.Int) (*types.Header, error)
}
