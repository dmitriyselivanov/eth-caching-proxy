package service

import (
	"errors"
	"eth-caching-proxy/cache"
	"eth-caching-proxy/model"
	"eth-caching-proxy/repository"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"testing"
)

func TestBlockService_TxFromBlockByHash(t *testing.T) {
	for i := range blockServiceTestCases_TxFromBlockByHash {
		repo := repository.New()
		c := cache.New()
		blockService := NewEthBlockService(repo, c)

		tx, err := blockService.TxFromBlockByHash(blockServiceTestCases_TxFromBlockByHash[i].blockNumber, blockServiceTestCases_TxFromBlockByHash[i].txHash)

		if err != nil {
			if blockServiceTestCases_TxFromBlockByHash[i].expectedError == nil {
				t.Fatalf("Expected no error, got %v", err)
			}

			if err.Error() != blockServiceTestCases_TxFromBlockByHash[i].expectedError.Error() {
				t.Fatalf("FAIL Error: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", blockServiceTestCases_TxFromBlockByHash[i].description, blockServiceTestCases_TxFromBlockByHash[i].expectedError, err, i)
			}
		}

		if tx != nil && blockServiceTestCases_TxFromBlockByHash[i].expectedTransaction != nil {
			// Gas
			if tx.Gas() != blockServiceTestCases_TxFromBlockByHash[i].expectedTransaction.Gas {
				t.Fatalf("FAIL Gas: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", blockServiceTestCases_TxFromBlockByHash[i].description, blockServiceTestCases_TxFromBlockByHash[i].expectedTransaction.Gas, tx.Gas(), i)
			}

			// ChainId
			if tx.ChainId().Int64() != blockServiceTestCases_TxFromBlockByHash[i].expectedTransaction.ChainId.Int64() {
				t.Fatalf("FAIL ChainId: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", blockServiceTestCases_TxFromBlockByHash[i].description, blockServiceTestCases_TxFromBlockByHash[i].expectedTransaction.ChainId.Int64(), tx.ChainId().Int64(), i)
			}

			// Hash
			if tx.Hash().String() != blockServiceTestCases_TxFromBlockByHash[i].expectedTransaction.Hash {
				t.Fatalf("FAIL Hash: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", blockServiceTestCases_TxFromBlockByHash[i].description, blockServiceTestCases_TxFromBlockByHash[i].expectedTransaction.Hash, tx.Hash().String(), i)
			}

			// GasPrice
			if tx.GasPrice().Int64() != blockServiceTestCases_TxFromBlockByHash[i].expectedTransaction.GasPrice.Int64() {
				t.Fatalf("FAIL GasPrice: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", blockServiceTestCases_TxFromBlockByHash[i].description, blockServiceTestCases_TxFromBlockByHash[i].expectedTransaction.GasPrice.Int64(), tx.GasPrice().Int64(), i)
			}

			// Nonce
			if tx.Nonce() != blockServiceTestCases_TxFromBlockByHash[i].expectedTransaction.Nonce {
				t.Fatalf("FAIL Nonce: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", blockServiceTestCases_TxFromBlockByHash[i].description, blockServiceTestCases_TxFromBlockByHash[i].expectedTransaction.Nonce, tx.Nonce(), i)
			}
		}

		t.Logf("Pass: %s", blockServiceTestCases_TxFromBlockByHash[i].description)
	}
}

func TestBlockService_TxFromBlockByIndex(t *testing.T) {
	for i := range blockServiceTestCases_TxFromBlockByIndex {
		repo := repository.New()
		c := cache.New()
		blockService := NewEthBlockService(repo, c)

		tx, err := blockService.TxFromBlockByIndex(blockServiceTestCases_TxFromBlockByIndex[i].blockNumber, blockServiceTestCases_TxFromBlockByIndex[i].txIndex)

		if err != nil {
			if blockServiceTestCases_TxFromBlockByIndex[i].expectedError == nil {
				t.Fatalf("Expected no error, got %v", err)
			}

			if err.Error() != blockServiceTestCases_TxFromBlockByIndex[i].expectedError.Error() {
				t.Fatalf("FAIL Error: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", blockServiceTestCases_TxFromBlockByIndex[i].description, blockServiceTestCases_TxFromBlockByIndex[i].expectedError, err, i)
			}
		}

		if tx != nil && blockServiceTestCases_TxFromBlockByIndex[i].expectedTransaction != nil {
			// Gas
			if tx.Gas() != blockServiceTestCases_TxFromBlockByIndex[i].expectedTransaction.Gas {
				t.Fatalf("FAIL Gas: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", blockServiceTestCases_BlockByNumber[i].description, blockServiceTestCases_TxFromBlockByIndex[i].expectedTransaction.Gas, tx.Gas(), i)
			}

			// ChainId
			if tx.ChainId().Int64() != blockServiceTestCases_TxFromBlockByIndex[i].expectedTransaction.ChainId.Int64() {
				t.Fatalf("FAIL ChainId: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", blockServiceTestCases_BlockByNumber[i].description, blockServiceTestCases_TxFromBlockByIndex[i].expectedTransaction.ChainId.Int64(), tx.ChainId().Int64(), i)
			}

			// Hash
			if tx.Hash().String() != blockServiceTestCases_TxFromBlockByIndex[i].expectedTransaction.Hash {
				t.Fatalf("FAIL Hash: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", blockServiceTestCases_BlockByNumber[i].description, blockServiceTestCases_TxFromBlockByIndex[i].expectedTransaction.Hash, tx.Hash().String(), i)
			}

			// GasPrice
			if tx.GasPrice().Int64() != blockServiceTestCases_TxFromBlockByIndex[i].expectedTransaction.GasPrice.Int64() {
				t.Fatalf("FAIL GasPrice: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", blockServiceTestCases_BlockByNumber[i].description, blockServiceTestCases_TxFromBlockByIndex[i].expectedTransaction.GasPrice.Int64(), tx.GasPrice().Int64(), i)
			}

			// Nonce
			if tx.Nonce() != blockServiceTestCases_TxFromBlockByIndex[i].expectedTransaction.Nonce {
				t.Fatalf("FAIL Nonce: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", blockServiceTestCases_BlockByNumber[i].description, blockServiceTestCases_TxFromBlockByIndex[i].expectedTransaction.Nonce, tx.Nonce(), i)
			}
		}

		t.Logf("Pass: %s", blockServiceTestCases_TxFromBlockByIndex[i].description)
	}
}

func TestBlockService_BlockByNumber(t *testing.T) {
	for i := range blockServiceTestCases_BlockByNumber {
		repo := repository.New()
		c := cache.New()
		blockService := NewEthBlockService(repo, c)

		block, err := blockService.BlockByNumber(blockServiceTestCases_BlockByNumber[i].blockNumber)

		if err != nil {
			if blockServiceTestCases_BlockByNumber[i].expectedError == nil {
				t.Fatalf("Expected no error, got %v", err)
			}

			if err.Error() != blockServiceTestCases_BlockByNumber[i].expectedError.Error() {
				t.Fatalf("FAIL Error: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", blockServiceTestCases_BlockByNumber[i].description, blockServiceTestCases_BlockByNumber[i].expectedError, err, i)
			}
		}

		if block != nil && blockServiceTestCases_BlockByNumber[i].expectedBlock != nil {
			// Number
			if block.Number.Int64() != blockServiceTestCases_BlockByNumber[i].expectedBlock.Number.Int64() {
				t.Fatalf("FAIL Number: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", blockServiceTestCases_BlockByNumber[i].description, blockServiceTestCases_BlockByNumber[i].expectedBlock.Number.Int64(), block.Number.Int64(), i)
			}

			// BaseFee
			if block.BaseFee.Int64() != blockServiceTestCases_BlockByNumber[i].expectedBlock.BaseFee.Int64() {
				t.Fatalf("FAIL BaseFee: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", blockServiceTestCases_BlockByNumber[i].description, blockServiceTestCases_BlockByNumber[i].expectedBlock.BaseFee.Int64(), block.BaseFee.Int64(), i)
			}

			// Difficulty
			if block.Difficulty.Int64() != blockServiceTestCases_BlockByNumber[i].expectedBlock.Difficulty.Int64() {
				t.Fatalf("FAIL Difficulty: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", blockServiceTestCases_BlockByNumber[i].description, blockServiceTestCases_BlockByNumber[i].expectedBlock.Difficulty.Int64(), block.Difficulty.Int64(), i)
			}

			// Nonce
			if block.Nonce != blockServiceTestCases_BlockByNumber[i].expectedBlock.Nonce {
				t.Fatalf("FAIL Nonce: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", blockServiceTestCases_BlockByNumber[i].description, blockServiceTestCases_BlockByNumber[i].expectedBlock.Nonce, block.Nonce, i)
			}

			// Time
			if block.Time != blockServiceTestCases_BlockByNumber[i].expectedBlock.Time {
				t.Fatalf("FAIL Time: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", blockServiceTestCases_BlockByNumber[i].description, blockServiceTestCases_BlockByNumber[i].expectedBlock.Time, block.Time, i)
			}
		}

		t.Logf("Pass: %s", blockServiceTestCases_BlockByNumber[i].description)
	}
}

type testTransaction struct {
	Gas      uint64
	ChainId  *big.Int
	Hash     string
	GasPrice *big.Int
	Nonce    uint64
}

var blockServiceTestCases_TxFromBlockByHash = []struct {
	description         string
	blockNumber         *big.Int
	txHash              common.Hash
	expectedTransaction *testTransaction
	expectedError       error
}{
	{
		description:         "Transaction was not found",
		blockNumber:         big.NewInt(13588593),
		txHash:              common.HexToHash("123"),
		expectedTransaction: nil,
		expectedError:       errors.New("not found"),
	},

	{
		description: "Received a correct transaction",
		blockNumber: big.NewInt(13588593),
		txHash:      common.HexToHash("0x856ed296eb7f87619393143cc28ae5705e22866709ff79cd93af71d8132037c0"),
		expectedTransaction: &testTransaction{
			Gas:      uint64(1000000),
			ChainId:  big.NewInt(1),
			Hash:     "0x856ed296eb7f87619393143cc28ae5705e22866709ff79cd93af71d8132037c0",
			GasPrice: big.NewInt(119586681073),
			Nonce:    uint64(21042),
		},
		expectedError: errors.New("not found"),
	},
}

var blockServiceTestCases_TxFromBlockByIndex = []struct {
	description         string
	blockNumber         *big.Int
	txIndex             int
	expectedTransaction *testTransaction
	expectedError       error
}{
	{
		description:         "Transaction was not found",
		blockNumber:         big.NewInt(13588593),
		txIndex:             100500,
		expectedTransaction: nil,
		expectedError:       errors.New("not found"),
	},

	{
		description: "Received a correct transaction",
		blockNumber: big.NewInt(13588593),
		txIndex:     1,
		expectedTransaction: &testTransaction{
			Gas:      uint64(1000000),
			ChainId:  big.NewInt(1),
			Hash:     "0x856ed296eb7f87619393143cc28ae5705e22866709ff79cd93af71d8132037c0",
			GasPrice: big.NewInt(119586681073),
			Nonce:    uint64(21042),
		},
		expectedError: errors.New("not found"),
	},
}

var blockServiceTestCases_BlockByNumber = []struct {
	description   string
	blockNumber   *big.Int
	expectedBlock *model.Block
	expectedError error
}{
	{
		description:   "Block was not found",
		blockNumber:   big.NewInt(2),
		expectedBlock: nil,
		expectedError: errors.New("not found"),
	},

	{
		description: "Received a correct block",
		blockNumber: big.NewInt(13588593),
		expectedBlock: &model.Block{
			Number:     big.NewInt(13588593),
			BaseFee:    big.NewInt(105410383176),
			Difficulty: big.NewInt(10506939366079207),
			Nonce:      uint64(17499604187474856417),
			Time:       uint64(1636548900),
		},
		expectedError: nil,
	},
}
