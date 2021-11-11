package repository

import (
	"eth-caching-proxy/model"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

// BlockRepository provides methods to get blocks and block headers from ethereum network
type BlockRepository interface {
	BlockByNumber(number *big.Int) (*model.Block, error)
	LatestBlock() (*model.Block, error)

	LatestBlockHeader() (*types.Header, error)
	BlockHeaderByNumber(blockNumber *big.Int) (*types.Header, error)
}
