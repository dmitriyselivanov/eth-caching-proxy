package service

import (
	"errors"
	"eth-caching-proxy/cache"
	"eth-caching-proxy/model"
	"eth-caching-proxy/repository"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

type ethBlockService struct {
	repository *repository.Repository
	cache      *cache.Cache
}

// NewEthBlockService creates new ethBlockService with given repo and cache
func NewEthBlockService(repository *repository.Repository, cache *cache.Cache) *ethBlockService {
	return &ethBlockService{repository: repository, cache: cache}
}

func (s *ethBlockService) LatestBlock() (*model.BlockResponseDTO, error) {
	block, err := s.repository.BlockRepository.LatestBlock()
	if err != nil {
		return nil, err
	}

	return model.NewBlockResponseDTO(block), nil
}

func (s *ethBlockService) BlockByNumber(blockNumber *big.Int) (*model.BlockResponseDTO, error) {
	block, err := s.blockByNumber(blockNumber)
	if err != nil {
		return nil, err
	}

	return model.NewBlockResponseDTO(block), nil
}

func (s *ethBlockService) TxFromLatestBlockByHash(txHash common.Hash) (*types.Transaction, error) {
	block, err := s.latestBlock()
	if err != nil {
		return nil, err
	}

	tx := block.Tx(txHash)
	if tx == nil {
		return nil, errors.New("not found")
	}

	return tx, nil
}

func (s *ethBlockService) TxFromLatestBlockByIndex(txIndex int) (*types.Transaction, error) {
	block, err := s.latestBlock()
	if err != nil {
		return nil, err
	}

	txs := block.Txs()
	if len(txs) <= txIndex {
		return nil, errors.New("not found")
	}

	return txs[txIndex], nil
}

func (s *ethBlockService) TxFromBlockByIndex(blockNumber *big.Int, txIndex int) (*types.Transaction, error) {
	block, err := s.blockByNumber(blockNumber)
	if err != nil {
		return nil, err
	}

	txs := block.Txs()
	if len(txs) <= txIndex {
		return nil, errors.New("not found")
	}

	return txs[txIndex], nil
}

func (s *ethBlockService) TxFromBlockByHash(blockNumber *big.Int, txHash common.Hash) (*types.Transaction, error) {
	block, err := s.blockByNumber(blockNumber)
	if err != nil {
		return nil, err
	}

	tx := block.Tx(txHash)
	if tx == nil {
		return nil, errors.New("not found")
	}

	return tx, nil
}

func (s *ethBlockService) latestBlock() (*model.Block, error) {
	return s.repository.BlockRepository.LatestBlock()
}

// tries to get the block from cache. Gets the block from repo if not present in cache
func (s *ethBlockService) blockByNumber(blockNumber *big.Int) (*model.Block, error) {
	if blockNumber == nil {
		return nil, errors.New("nil block number provided")
	}

	// trying to get block from cache
	block, existsBlock := s.cache.BlockCache.GetBlock(blockNumber)
	if existsBlock {
		return block, nil
	}

	// ok we don't have this block in cache, let's try to get it from repo
	block, err := s.repository.BlockRepository.BlockByNumber(blockNumber)
	if err != nil {
		return nil, err
	}

	// we need latest block header to check whether this block is in latest 20 blocks
	latestBlockHeader, err := s.repository.BlockRepository.LatestBlockHeader()
	if err != nil {
		return nil, err
	}
	// we can cache this block because it won't change due to ethereum reorgs
	if !s.isWithinLast20Blocks(latestBlockHeader.Number, blockNumber) {
		s.cache.BlockCache.AddBlock(block)
	}

	return block, nil
}

func (s *ethBlockService) isWithinLast20Blocks(latestBlockNumber *big.Int, currentBlockNumber *big.Int) bool {
	twenty := big.NewInt(20)

	sub := big.NewInt(0)
	sub.Sub(latestBlockNumber, currentBlockNumber)

	return sub.Cmp(twenty) == -1
}
