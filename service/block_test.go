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
	for i := range blockServiceTestCasesTxFromBlockByHash {
		repo := repository.New()
		c := cache.New()
		blockService := NewEthBlockService(repo, c)

		tx, err := blockService.TxFromBlockByHash(blockServiceTestCasesTxFromBlockByHash[i].blockNumber, blockServiceTestCasesTxFromBlockByHash[i].txHash)

		if err != nil {
			if blockServiceTestCasesTxFromBlockByHash[i].expectedError == nil {
				t.Fatalf("Expected no error, got %v", err)
			}

			if err.Error() != blockServiceTestCasesTxFromBlockByHash[i].expectedError.Error() {
				t.Fatalf("FAIL Error: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", blockServiceTestCasesTxFromBlockByHash[i].description, blockServiceTestCasesTxFromBlockByHash[i].expectedError, err, i)
			}
		}

		if tx != nil && blockServiceTestCasesTxFromBlockByHash[i].expectedTransaction != nil {
			// Gas
			if tx.Gas() != blockServiceTestCasesTxFromBlockByHash[i].expectedTransaction.Gas {
				t.Fatalf("FAIL Gas: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", blockServiceTestCasesTxFromBlockByHash[i].description, blockServiceTestCasesTxFromBlockByHash[i].expectedTransaction.Gas, tx.Gas(), i)
			}

			// ChainID
			if tx.ChainId().Int64() != blockServiceTestCasesTxFromBlockByHash[i].expectedTransaction.ChainID.Int64() {
				t.Fatalf("FAIL ChainID: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", blockServiceTestCasesTxFromBlockByHash[i].description, blockServiceTestCasesTxFromBlockByHash[i].expectedTransaction.ChainID.Int64(), tx.ChainId().Int64(), i)
			}

			// Hash
			if tx.Hash().String() != blockServiceTestCasesTxFromBlockByHash[i].expectedTransaction.Hash {
				t.Fatalf("FAIL Hash: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", blockServiceTestCasesTxFromBlockByHash[i].description, blockServiceTestCasesTxFromBlockByHash[i].expectedTransaction.Hash, tx.Hash().String(), i)
			}

			// GasPrice
			if tx.GasPrice().Int64() != blockServiceTestCasesTxFromBlockByHash[i].expectedTransaction.GasPrice.Int64() {
				t.Fatalf("FAIL GasPrice: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", blockServiceTestCasesTxFromBlockByHash[i].description, blockServiceTestCasesTxFromBlockByHash[i].expectedTransaction.GasPrice.Int64(), tx.GasPrice().Int64(), i)
			}

			// Nonce
			if tx.Nonce() != blockServiceTestCasesTxFromBlockByHash[i].expectedTransaction.Nonce {
				t.Fatalf("FAIL Nonce: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", blockServiceTestCasesTxFromBlockByHash[i].description, blockServiceTestCasesTxFromBlockByHash[i].expectedTransaction.Nonce, tx.Nonce(), i)
			}
		}

		t.Logf("Pass: %s", blockServiceTestCasesTxFromBlockByHash[i].description)
	}
}

func TestBlockService_TxFromBlockByIndex(t *testing.T) {
	for i := range blockServiceTestCasesTxFromBlockByIndex {
		repo := repository.New()
		c := cache.New()
		blockService := NewEthBlockService(repo, c)

		tx, err := blockService.TxFromBlockByIndex(blockServiceTestCasesTxFromBlockByIndex[i].blockNumber, blockServiceTestCasesTxFromBlockByIndex[i].txIndex)

		if err != nil {
			if blockServiceTestCasesTxFromBlockByIndex[i].expectedError == nil {
				t.Fatalf("Expected no error, got %v", err)
			}

			if err.Error() != blockServiceTestCasesTxFromBlockByIndex[i].expectedError.Error() {
				t.Fatalf("FAIL Error: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", blockServiceTestCasesTxFromBlockByIndex[i].description, blockServiceTestCasesTxFromBlockByIndex[i].expectedError, err, i)
			}
		}

		if tx != nil && blockServiceTestCasesTxFromBlockByIndex[i].expectedTransaction != nil {
			// Gas
			if tx.Gas() != blockServiceTestCasesTxFromBlockByIndex[i].expectedTransaction.Gas {
				t.Fatalf("FAIL Gas: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", blockServiceTestCasesBlockByNumber[i].description, blockServiceTestCasesTxFromBlockByIndex[i].expectedTransaction.Gas, tx.Gas(), i)
			}

			// ChainID
			if tx.ChainId().Int64() != blockServiceTestCasesTxFromBlockByIndex[i].expectedTransaction.ChainID.Int64() {
				t.Fatalf("FAIL ChainID: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", blockServiceTestCasesBlockByNumber[i].description, blockServiceTestCasesTxFromBlockByIndex[i].expectedTransaction.ChainID.Int64(), tx.ChainId().Int64(), i)
			}

			// Hash
			if tx.Hash().String() != blockServiceTestCasesTxFromBlockByIndex[i].expectedTransaction.Hash {
				t.Fatalf("FAIL Hash: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", blockServiceTestCasesBlockByNumber[i].description, blockServiceTestCasesTxFromBlockByIndex[i].expectedTransaction.Hash, tx.Hash().String(), i)
			}

			// GasPrice
			if tx.GasPrice().Int64() != blockServiceTestCasesTxFromBlockByIndex[i].expectedTransaction.GasPrice.Int64() {
				t.Fatalf("FAIL GasPrice: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", blockServiceTestCasesBlockByNumber[i].description, blockServiceTestCasesTxFromBlockByIndex[i].expectedTransaction.GasPrice.Int64(), tx.GasPrice().Int64(), i)
			}

			// Nonce
			if tx.Nonce() != blockServiceTestCasesTxFromBlockByIndex[i].expectedTransaction.Nonce {
				t.Fatalf("FAIL Nonce: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", blockServiceTestCasesBlockByNumber[i].description, blockServiceTestCasesTxFromBlockByIndex[i].expectedTransaction.Nonce, tx.Nonce(), i)
			}
		}

		t.Logf("Pass: %s", blockServiceTestCasesTxFromBlockByIndex[i].description)
	}
}

func TestBlockService_BlockByNumber(t *testing.T) {
	for i := range blockServiceTestCasesBlockByNumber {
		repo := repository.New()
		c := cache.New()
		blockService := NewEthBlockService(repo, c)

		block, err := blockService.BlockByNumber(blockServiceTestCasesBlockByNumber[i].blockNumber)

		if err != nil {
			if blockServiceTestCasesBlockByNumber[i].expectedError == nil {
				t.Fatalf("Expected no error, got %v", err)
			}

			if err.Error() != blockServiceTestCasesBlockByNumber[i].expectedError.Error() {
				t.Fatalf("FAIL Error: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", blockServiceTestCasesBlockByNumber[i].description, blockServiceTestCasesBlockByNumber[i].expectedError, err, i)
			}
		}

		if block != nil && blockServiceTestCasesBlockByNumber[i].expectedBlock != nil {
			// Number
			if block.Number.Int64() != blockServiceTestCasesBlockByNumber[i].expectedBlock.Number.Int64() {
				t.Fatalf("FAIL Number: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", blockServiceTestCasesBlockByNumber[i].description, blockServiceTestCasesBlockByNumber[i].expectedBlock.Number.Int64(), block.Number.Int64(), i)
			}

			// BaseFee
			if block.BaseFee.Int64() != blockServiceTestCasesBlockByNumber[i].expectedBlock.BaseFee.Int64() {
				t.Fatalf("FAIL BaseFee: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", blockServiceTestCasesBlockByNumber[i].description, blockServiceTestCasesBlockByNumber[i].expectedBlock.BaseFee.Int64(), block.BaseFee.Int64(), i)
			}

			// Difficulty
			if block.Difficulty.Int64() != blockServiceTestCasesBlockByNumber[i].expectedBlock.Difficulty.Int64() {
				t.Fatalf("FAIL Difficulty: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", blockServiceTestCasesBlockByNumber[i].description, blockServiceTestCasesBlockByNumber[i].expectedBlock.Difficulty.Int64(), block.Difficulty.Int64(), i)
			}

			// Nonce
			if block.Nonce != blockServiceTestCasesBlockByNumber[i].expectedBlock.Nonce {
				t.Fatalf("FAIL Nonce: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", blockServiceTestCasesBlockByNumber[i].description, blockServiceTestCasesBlockByNumber[i].expectedBlock.Nonce, block.Nonce, i)
			}

			// Time
			if block.Time != blockServiceTestCasesBlockByNumber[i].expectedBlock.Time {
				t.Fatalf("FAIL Time: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", blockServiceTestCasesBlockByNumber[i].description, blockServiceTestCasesBlockByNumber[i].expectedBlock.Time, block.Time, i)
			}
		}

		t.Logf("Pass: %s", blockServiceTestCasesBlockByNumber[i].description)
	}
}

type testTransaction struct {
	Gas      uint64
	ChainID  *big.Int
	Hash     string
	GasPrice *big.Int
	Nonce    uint64
}

var blockServiceTestCasesTxFromBlockByHash = []struct {
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
			ChainID:  big.NewInt(1),
			Hash:     "0x856ed296eb7f87619393143cc28ae5705e22866709ff79cd93af71d8132037c0",
			GasPrice: big.NewInt(119586681073),
			Nonce:    uint64(21042),
		},
		expectedError: errors.New("not found"),
	},
}

var blockServiceTestCasesTxFromBlockByIndex = []struct {
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
			ChainID:  big.NewInt(1),
			Hash:     "0x856ed296eb7f87619393143cc28ae5705e22866709ff79cd93af71d8132037c0",
			GasPrice: big.NewInt(119586681073),
			Nonce:    uint64(21042),
		},
		expectedError: errors.New("not found"),
	},
}

var blockServiceTestCasesBlockByNumber = []struct {
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
