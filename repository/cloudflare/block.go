package cloudflare

import (
	"context"
	"eth-caching-proxy/model"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

type blockRepository struct {
	client *ethclient.Client
}

func NewBlockRepository(client *ethclient.Client) *blockRepository {
	return &blockRepository{client: client}
}

func (r *blockRepository) BlockByNumber(number *big.Int) (*model.Block, error) {
	ethBlock, err := r.client.BlockByNumber(context.Background(), number)
	if err != nil {
		return nil, err
	}

	return newBlock(ethBlock), nil
}

func (r *blockRepository) LatestBlock() (*model.Block, error) {
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

func (r *blockRepository) LatestBlockHeader() (*types.Header, error) {
	return r.client.HeaderByNumber(context.Background(), nil)
}

func (r *blockRepository) BlockHeaderByNumber(blockNumber *big.Int) (*types.Header, error) {
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
