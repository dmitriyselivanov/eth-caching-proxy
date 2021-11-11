package controller

import (
	"encoding/json"
	"eth-caching-proxy/model"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBlockController_TxFromBlockByHash(t *testing.T) {
	ts := httptest.NewServer(NewRouter())
	defer ts.Close()

	for i := range blockControllerTestCases_TxFromBlockByHash {
		resp, err := http.Get(fmt.Sprintf("%s/block/%s/txs/%s", ts.URL, blockControllerTestCases_TxFromBlockByHash[i].blockNumber.String(), blockControllerTestCases_TxFromBlockByHash[i].txHash.String()))

		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		if resp.StatusCode != blockControllerTestCases_TxFromBlockByHash[i].expectedStatusCode {
			t.Fatalf("Expected status code %d, got %v", blockControllerTestCases_TxFromBlockByHash[i].expectedStatusCode, resp.StatusCode)
		}

		val, ok := resp.Header["Content-Type"]

		// Assert that the "content-type" header is actually set
		if !ok {
			t.Fatalf("Expected Content-Type header to be set")
		}

		// Assert that it was set as expected
		if val[0] != blockControllerTestCases_TxFromBlockByHash[i].expectedContentType {
			t.Fatalf("Expected %s, got %s", blockControllerTestCases_TxFromBlockByHash[i].expectedContentType, val[0])
		}

		if blockControllerTestCases_TxFromBlockByHash[i].expectedTransaction == nil {
			t.Logf("Pass: %s", blockControllerTestCases_TxFromBlockByHash[i].description)
			return
		}

		var respBytes []byte
		_, err = resp.Body.Read(respBytes)

		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		var tx *testTransaction
		err = json.Unmarshal(respBytes, tx)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		if tx == nil {
			t.Fatal("Expected not nil tx after json unmarshal")
		}

		// Gas
		if tx.Gas != blockControllerTestCases_TxFromBlockByHash[i].expectedTransaction.Gas {
			t.Fatalf("FAIL Gas: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", blockControllerTestCases_TxFromBlockByHash[i].description, blockControllerTestCases_TxFromBlockByHash[i].expectedTransaction.Gas, tx.Gas, i)
		}

		// ChainId
		if tx.ChainId.Int64() != blockControllerTestCases_TxFromBlockByHash[i].expectedTransaction.ChainId.Int64() {
			t.Fatalf("FAIL ChainId: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", blockControllerTestCases_TxFromBlockByHash[i].description, blockControllerTestCases_TxFromBlockByHash[i].expectedTransaction.ChainId.Int64(), tx.ChainId.Int64(), i)
		}

		// Hash
		if tx.Hash != blockControllerTestCases_TxFromBlockByHash[i].expectedTransaction.Hash {
			t.Fatalf("FAIL Hash: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", blockControllerTestCases_TxFromBlockByHash[i].description, blockControllerTestCases_TxFromBlockByHash[i].expectedTransaction.Hash, tx.Hash, i)
		}

		// GasPrice
		if tx.GasPrice.Int64() != blockControllerTestCases_TxFromBlockByHash[i].expectedTransaction.GasPrice.Int64() {
			t.Fatalf("FAIL GasPrice: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", blockControllerTestCases_TxFromBlockByHash[i].description, blockControllerTestCases_TxFromBlockByHash[i].expectedTransaction.GasPrice.Int64(), tx.GasPrice.Int64(), i)
		}

		// Nonce
		if tx.Nonce != blockControllerTestCases_TxFromBlockByHash[i].expectedTransaction.Nonce {
			t.Fatalf("FAIL Nonce: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", blockControllerTestCases_TxFromBlockByHash[i].description, blockControllerTestCases_TxFromBlockByHash[i].expectedTransaction.Nonce, tx.Nonce, i)
		}

		t.Logf("Pass: %s", blockControllerTestCases_BlockByNumber[i].description)
	}
}

func TestBlockController_TxFromBlockByIndex(t *testing.T) {
	ts := httptest.NewServer(NewRouter())
	defer ts.Close()

	for i := range blockControllerTestCases_TxFromBlockByIndex {
		resp, err := http.Get(fmt.Sprintf("%s/block/%s/txs/%d", ts.URL, blockControllerTestCases_TxFromBlockByIndex[i].blockNumber.String(), blockControllerTestCases_TxFromBlockByIndex[i].txIndex))

		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		if resp.StatusCode != blockControllerTestCases_TxFromBlockByIndex[i].expectedStatusCode {
			t.Fatalf("Expected status code %d, got %v", blockControllerTestCases_TxFromBlockByIndex[i].expectedStatusCode, resp.StatusCode)
		}

		val, ok := resp.Header["Content-Type"]

		// Assert that the "content-type" header is actually set
		if !ok {
			t.Fatalf("Expected Content-Type header to be set")
		}

		// Assert that it was set as expected
		if val[0] != blockControllerTestCases_TxFromBlockByIndex[i].expectedContentType {
			t.Fatalf("Expected %s, got %s", blockControllerTestCases_TxFromBlockByIndex[i].expectedContentType, val[0])
		}

		if blockControllerTestCases_TxFromBlockByIndex[i].expectedTransaction == nil {
			t.Logf("Pass: %s", blockControllerTestCases_TxFromBlockByIndex[i].description)
			return
		}

		var respBytes []byte
		_, err = resp.Body.Read(respBytes)

		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		var tx *testTransaction
		err = json.Unmarshal(respBytes, tx)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		if tx == nil {
			t.Fatal("Expected not nil tx after json unmarshal")
		}

		// Gas
		if tx.Gas != blockControllerTestCases_TxFromBlockByIndex[i].expectedTransaction.Gas {
			t.Fatalf("FAIL Gas: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", blockControllerTestCases_TxFromBlockByIndex[i].description, blockControllerTestCases_TxFromBlockByIndex[i].expectedTransaction.Gas, tx.Gas, i)
		}

		// ChainId
		if tx.ChainId.Int64() != blockControllerTestCases_TxFromBlockByIndex[i].expectedTransaction.ChainId.Int64() {
			t.Fatalf("FAIL ChainId: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", blockControllerTestCases_TxFromBlockByIndex[i].description, blockControllerTestCases_TxFromBlockByIndex[i].expectedTransaction.ChainId.Int64(), tx.ChainId.Int64(), i)
		}

		// Hash
		if tx.Hash != blockControllerTestCases_TxFromBlockByIndex[i].expectedTransaction.Hash {
			t.Fatalf("FAIL Hash: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", blockControllerTestCases_TxFromBlockByIndex[i].description, blockControllerTestCases_TxFromBlockByIndex[i].expectedTransaction.Hash, tx.Hash, i)
		}

		// GasPrice
		if tx.GasPrice.Int64() != blockControllerTestCases_TxFromBlockByIndex[i].expectedTransaction.GasPrice.Int64() {
			t.Fatalf("FAIL GasPrice: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", blockControllerTestCases_TxFromBlockByIndex[i].description, blockControllerTestCases_TxFromBlockByIndex[i].expectedTransaction.GasPrice.Int64(), tx.GasPrice.Int64(), i)
		}

		// Nonce
		if tx.Nonce != blockControllerTestCases_TxFromBlockByIndex[i].expectedTransaction.Nonce {
			t.Fatalf("FAIL Nonce: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", blockControllerTestCases_TxFromBlockByIndex[i].description, blockControllerTestCases_TxFromBlockByIndex[i].expectedTransaction.Nonce, tx.Nonce, i)
		}

		t.Logf("Pass: %s", blockControllerTestCases_BlockByNumber[i].description)
	}
}

func TestBlockController_BlockByNumber(t *testing.T) {
	ts := httptest.NewServer(NewRouter())
	defer ts.Close()

	for i := range blockControllerTestCases_BlockByNumber {
		resp, err := http.Get(fmt.Sprintf("%s/block/%s", ts.URL, blockControllerTestCases_BlockByNumber[i].blockNumber.String()))

		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		if resp.StatusCode != blockControllerTestCases_BlockByNumber[i].expectedStatusCode {
			t.Fatalf("Expected status code %d, got %v", blockControllerTestCases_BlockByNumber[i].expectedStatusCode, resp.StatusCode)
		}

		val, ok := resp.Header["Content-Type"]

		// Assert that the "content-type" header is actually set
		if !ok {
			t.Fatalf("Expected Content-Type header to be set")
		}

		// Assert that it was set as expected
		if val[0] != blockControllerTestCases_BlockByNumber[i].expectedContentType {
			t.Fatalf("Expected %s, got %s", blockControllerTestCases_BlockByNumber[i].expectedContentType, val[0])
		}

		if blockControllerTestCases_BlockByNumber[i].expectedBlock == nil {
			t.Logf("Pass: %s", blockControllerTestCases_BlockByNumber[i].description)
			return
		}

		var respBytes []byte
		_, err = resp.Body.Read(respBytes)

		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		var block *model.BlockResponseDTO
		err = json.Unmarshal(respBytes, block)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		if block == nil {
			t.Fatal("Expected not nil block after json unmarshal")
		}

		// Number
		if block.Number.Int64() != blockControllerTestCases_BlockByNumber[i].expectedBlock.Number.Int64() {
			t.Fatalf("FAIL Number: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", blockControllerTestCases_BlockByNumber[i].description, blockControllerTestCases_BlockByNumber[i].expectedBlock.Number.Int64(), block.Number.Int64(), i)
		}

		// BaseFee
		if block.BaseFee.Int64() != blockControllerTestCases_BlockByNumber[i].expectedBlock.BaseFee.Int64() {
			t.Fatalf("FAIL BaseFee: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", blockControllerTestCases_BlockByNumber[i].description, blockControllerTestCases_BlockByNumber[i].expectedBlock.BaseFee.Int64(), block.BaseFee.Int64(), i)
		}

		// Difficulty
		if block.Difficulty.Int64() != blockControllerTestCases_BlockByNumber[i].expectedBlock.Difficulty.Int64() {
			t.Fatalf("FAIL Difficulty: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", blockControllerTestCases_BlockByNumber[i].description, blockControllerTestCases_BlockByNumber[i].expectedBlock.Difficulty.Int64(), block.Difficulty.Int64(), i)
		}

		// Nonce
		if block.Nonce != blockControllerTestCases_BlockByNumber[i].expectedBlock.Nonce {
			t.Fatalf("FAIL Nonce: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", blockControllerTestCases_BlockByNumber[i].description, blockControllerTestCases_BlockByNumber[i].expectedBlock.Nonce, block.Nonce, i)
		}

		// Time
		if block.Time != blockControllerTestCases_BlockByNumber[i].expectedBlock.Time {
			t.Fatalf("FAIL Time: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", blockControllerTestCases_BlockByNumber[i].description, blockControllerTestCases_BlockByNumber[i].expectedBlock.Time, block.Time, i)
		}

		t.Logf("Pass: %s", blockControllerTestCases_BlockByNumber[i].description)
	}
}

var blockControllerTestCases_BlockByNumber = []struct {
	description         string
	blockNumber         *big.Int
	expectedBlock       *model.BlockResponseDTO
	expectedContentType string
	expectedStatusCode  int
}{
	{
		description:         "Block not found",
		blockNumber:         big.NewInt(100500123123),
		expectedBlock:       nil,
		expectedContentType: "application/json; charset=utf-8",
		expectedStatusCode:  404,
	},

	{
		description: "Received a correct block",
		blockNumber: big.NewInt(13588593),
		expectedBlock: &model.BlockResponseDTO{
			Number:     big.NewInt(13588593),
			BaseFee:    big.NewInt(105410383176),
			Difficulty: big.NewInt(10506939366079207),
			Nonce:      uint64(17499604187474856417),
			Time:       uint64(1636548900),
		},
		expectedContentType: "application/json; charset=utf-8",
		expectedStatusCode:  200,
	},
}

type testTransaction struct {
	Gas      uint64   `json:"gas"`
	ChainId  *big.Int `json:"chainId"`
	Hash     string   `json:"hash"`
	GasPrice *big.Int `json:"gasPrice"`
	Nonce    uint64   `json:"nonce"`
}

var blockControllerTestCases_TxFromBlockByIndex = []struct {
	description         string
	blockNumber         *big.Int
	txIndex             int
	expectedTransaction *testTransaction
	expectedContentType string
	expectedStatusCode  int
}{
	{
		description:         "Transaction was not found",
		blockNumber:         big.NewInt(13588593),
		txIndex:             100500,
		expectedTransaction: nil,
		expectedContentType: "application/json; charset=utf-8",
		expectedStatusCode:  404,
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
		expectedContentType: "application/json; charset=utf-8",
		expectedStatusCode:  200,
	},
}

var blockControllerTestCases_TxFromBlockByHash = []struct {
	description         string
	blockNumber         *big.Int
	txHash              common.Hash
	expectedTransaction *testTransaction
	expectedContentType string
	expectedStatusCode  int
}{
	{
		description:         "Transaction was not found",
		blockNumber:         big.NewInt(13588593),
		txHash:              common.HexToHash("0123"),
		expectedTransaction: nil,
		expectedContentType: "application/json; charset=utf-8",
		expectedStatusCode:  404,
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
		expectedContentType: "application/json; charset=utf-8",
		expectedStatusCode:  200,
	},
}
