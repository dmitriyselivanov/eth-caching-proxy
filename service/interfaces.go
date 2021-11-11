package service

import (
	"eth-caching-proxy/model"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

type EthBlockService interface {
	LatestBlock() (*model.BlockResponseDTO, error)
	BlockByNumber(blockNumber *big.Int) (*model.BlockResponseDTO, error)

	TxFromLatestBlockByHash(txHash common.Hash) (*types.Transaction, error)
	TxFromLatestBlockByIndex(txIndex int) (*types.Transaction, error)
	TxFromBlockByIndex(blockNumber *big.Int, txIndex int) (*types.Transaction, error)
	TxFromBlockByHash(blockNumber *big.Int, txHash common.Hash) (*types.Transaction, error)
}
