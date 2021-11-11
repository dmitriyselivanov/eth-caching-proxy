package model

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"time"
)

// BlockResponseDTO is a data transfer object for model.Block
type BlockResponseDTO struct {
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

// NewBlockResponseDTO creates and returns BlockResponseDTO instance
func NewBlockResponseDTO(from *Block) *BlockResponseDTO {
	return &BlockResponseDTO{
		Number:       from.Number,
		Header:       from.Header,
		Bloom:        from.Bloom,
		ReceivedAt:   from.ReceivedAt,
		ReceivedFrom: from.ReceivedFrom,
		BaseFee:      from.BaseFee,
		Coinbase:     from.Coinbase,
		Difficulty:   from.Difficulty,
		Extra:        from.Extra,
		GasLimit:     from.GasLimit,
		GasUsed:      from.GasUsed,
		Hash:         from.Hash,
		MixDigest:    from.MixDigest,
		Nonce:        from.Nonce,
		ParentHash:   from.ParentHash,
		ReceiptHash:  from.ReceiptHash,
		Root:         from.Root,
		Size:         from.Size,
		Time:         from.Time,
		Uncles:       from.Uncles,
		Transactions: from.Transactions,
	}
}
