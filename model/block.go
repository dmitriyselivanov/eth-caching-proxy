package model

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"time"
)

// Block is a struct that represents ethereum block
type Block struct {
	Number       *big.Int           `json:"number"`
	Header       *types.Header      `json:"header"`
	Bloom        types.Bloom        `json:"bloom"`
	ReceivedAt   time.Time          `json:"receivedAt"`
	ReceivedFrom interface{}        `json:"receivedFrom"`
	BaseFee      *big.Int           `json:"baseFee"`
	Coinbase     common.Address     `json:"coinbase"`
	Difficulty   *big.Int           `json:"difficulty"`
	Extra        []byte             `json:"extra"`
	GasLimit     uint64             `json:"gasLimit"`
	GasUsed      uint64             `json:"gasUsed"`
	Hash         common.Hash        `json:"hash"`
	MixDigest    common.Hash        `json:"mixDigest"`
	Nonce        uint64             `json:"nonce"`
	ParentHash   common.Hash        `json:"parentHash"`
	ReceiptHash  common.Hash        `json:"receiptHash"`
	Root         common.Hash        `json:"root"`
	Size         common.StorageSize `json:"size"`
	Time         uint64             `json:"time"`
	Uncles       []*types.Header    `json:"uncles"`
	Transactions types.Transactions `json:"transactions"`
}

// Txs returns all transactions within a block
func (b *Block) Txs() types.Transactions { return b.Transactions }

// Tx returns a transaction within a block by hash
func (b *Block) Tx(hash common.Hash) *types.Transaction {
	for _, transaction := range b.Transactions {
		if transaction.Hash() == hash {
			return transaction
		}
	}
	return nil
}
