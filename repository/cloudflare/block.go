package cloudflare

import (
	"context"
	"eth-caching-proxy/model"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

// EthBlockRepository is a block repository backed by Cloudflare ethereum gateway
type EthBlockRepository struct {
	client *ethclient.Client
}

// NewBlockRepository creates new block repository instance
func NewBlockRepository(client *ethclient.Client) *EthBlockRepository {
	return &EthBlockRepository{client: client}
}

// BlockByNumber returns a block by number from eth network
func (r *EthBlockRepository) BlockByNumber(number *big.Int) (*model.Block, error) {
	ethBlock, err := r.client.BlockByNumber(context.Background(), number)
	if err != nil {
		return nil, err
	}

	return newBlock(ethBlock), nil
}

// LatestBlock returns a latest block from eth network
func (r *EthBlockRepository) LatestBlock() (*model.Block, error) {
	header, err := r.LatestBlockHeader()
	if err != nil {
		return nil, err
	}

	ethBlock, err := r.client.BlockByNumber(context.Background(), header.Number)
	if err != nil {
		return nil, err
	}

	return newBlock(ethBlock), nil
}

// LatestBlockHeader returns a latest block header from eth network
func (r *EthBlockRepository) LatestBlockHeader() (*types.Header, error) {
	return r.client.HeaderByNumber(context.Background(), nil)
}

// BlockHeaderByNumber returns a block header by number from eth network
func (r *EthBlockRepository) BlockHeaderByNumber(blockNumber *big.Int) (*types.Header, error) {
	return r.client.HeaderByNumber(context.Background(), blockNumber)
}

// converts types.Block from ethclient to domain model of type model.Block
func newBlock(from *types.Block) *model.Block {
	return &model.Block{
		Number:       from.Number(),
		Header:       from.Header(),
		Bloom:        from.Bloom(),
		ReceivedAt:   from.ReceivedAt,
		ReceivedFrom: from.ReceivedFrom,
		BaseFee:      from.BaseFee(),
		Coinbase:     from.Coinbase(),
		Difficulty:   from.Difficulty(),
		Extra:        from.Extra(),
		GasLimit:     from.GasLimit(),
		GasUsed:      from.GasUsed(),
		Hash:         from.Hash(),
		MixDigest:    from.MixDigest(),
		Nonce:        from.Nonce(),
		ParentHash:   from.ParentHash(),
		ReceiptHash:  from.ReceiptHash(),
		Root:         from.Root(),
		Size:         from.Size(),
		Time:         from.Time(),
		Uncles:       from.Uncles(),
		Transactions: from.Transactions(),
	}
}
