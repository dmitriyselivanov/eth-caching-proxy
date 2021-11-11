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

	for i := range blockControllerTestCasesTxFromBlockByHash {
		resp, err := http.Get(fmt.Sprintf("%s/block/%s/txs/%s", ts.URL, blockControllerTestCasesTxFromBlockByHash[i].blockNumber.String(), blockControllerTestCasesTxFromBlockByHash[i].txHash.String()))

		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		if resp.StatusCode != blockControllerTestCasesTxFromBlockByHash[i].expectedStatusCode {
			t.Fatalf("Expected status code %d, got %v", blockControllerTestCasesTxFromBlockByHash[i].expectedStatusCode, resp.StatusCode)
		}

		val, ok := resp.Header["Content-Type"]

		// Assert that the "content-type" header is actually set
		if !ok {
			t.Fatalf("Expected Content-Type header to be set")
		}

		// Assert that it was set as expected
		if val[0] != blockControllerTestCasesTxFromBlockByHash[i].expectedContentType {
			t.Fatalf("Expected %s, got %s", blockControllerTestCasesTxFromBlockByHash[i].expectedContentType, val[0])
		}

		if blockControllerTestCasesTxFromBlockByHash[i].expectedTransaction == nil {
			t.Logf("Pass: %s", blockControllerTestCasesTxFromBlockByHash[i].description)
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
		if tx.Gas != blockControllerTestCasesTxFromBlockByHash[i].expectedTransaction.Gas {
			t.Fatalf("FAIL Gas: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", blockControllerTestCasesTxFromBlockByHash[i].description, blockControllerTestCasesTxFromBlockByHash[i].expectedTransaction.Gas, tx.Gas, i)
		}

		// ChainID
		if tx.ChainID.Int64() != blockControllerTestCasesTxFromBlockByHash[i].expectedTransaction.ChainID.Int64() {
			t.Fatalf("FAIL ChainID: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", blockControllerTestCasesTxFromBlockByHash[i].description, blockControllerTestCasesTxFromBlockByHash[i].expectedTransaction.ChainID.Int64(), tx.ChainID.Int64(), i)
		}

		// Hash
		if tx.Hash != blockControllerTestCasesTxFromBlockByHash[i].expectedTransaction.Hash {
			t.Fatalf("FAIL Hash: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", blockControllerTestCasesTxFromBlockByHash[i].description, blockControllerTestCasesTxFromBlockByHash[i].expectedTransaction.Hash, tx.Hash, i)
		}

		// GasPrice
		if tx.GasPrice.Int64() != blockControllerTestCasesTxFromBlockByHash[i].expectedTransaction.GasPrice.Int64() {
			t.Fatalf("FAIL GasPrice: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", blockControllerTestCasesTxFromBlockByHash[i].description, blockControllerTestCasesTxFromBlockByHash[i].expectedTransaction.GasPrice.Int64(), tx.GasPrice.Int64(), i)
		}

		// Nonce
		if tx.Nonce != blockControllerTestCasesTxFromBlockByHash[i].expectedTransaction.Nonce {
			t.Fatalf("FAIL Nonce: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", blockControllerTestCasesTxFromBlockByHash[i].description, blockControllerTestCasesTxFromBlockByHash[i].expectedTransaction.Nonce, tx.Nonce, i)
		}

		t.Logf("Pass: %s", blockControllerTestCasesBlockByNumber[i].description)
	}
}

func TestBlockController_TxFromBlockByIndex(t *testing.T) {
	ts := httptest.NewServer(NewRouter())
	defer ts.Close()

	for i := range blockControllerTestCasesTxFromBlockByIndex {
		resp, err := http.Get(fmt.Sprintf("%s/block/%s/txs/%d", ts.URL, blockControllerTestCasesTxFromBlockByIndex[i].blockNumber.String(), blockControllerTestCasesTxFromBlockByIndex[i].txIndex))

		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		if resp.StatusCode != blockControllerTestCasesTxFromBlockByIndex[i].expectedStatusCode {
			t.Fatalf("Expected status code %d, got %v", blockControllerTestCasesTxFromBlockByIndex[i].expectedStatusCode, resp.StatusCode)
		}

		val, ok := resp.Header["Content-Type"]

		// Assert that the "content-type" header is actually set
		if !ok {
			t.Fatalf("Expected Content-Type header to be set")
		}

		// Assert that it was set as expected
		if val[0] != blockControllerTestCasesTxFromBlockByIndex[i].expectedContentType {
			t.Fatalf("Expected %s, got %s", blockControllerTestCasesTxFromBlockByIndex[i].expectedContentType, val[0])
		}

		if blockControllerTestCasesTxFromBlockByIndex[i].expectedTransaction == nil {
			t.Logf("Pass: %s", blockControllerTestCasesTxFromBlockByIndex[i].description)
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
		if tx.Gas != blockControllerTestCasesTxFromBlockByIndex[i].expectedTransaction.Gas {
			t.Fatalf("FAIL Gas: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", blockControllerTestCasesTxFromBlockByIndex[i].description, blockControllerTestCasesTxFromBlockByIndex[i].expectedTransaction.Gas, tx.Gas, i)
		}

		// ChainID
		if tx.ChainID.Int64() != blockControllerTestCasesTxFromBlockByIndex[i].expectedTransaction.ChainID.Int64() {
			t.Fatalf("FAIL ChainID: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", blockControllerTestCasesTxFromBlockByIndex[i].description, blockControllerTestCasesTxFromBlockByIndex[i].expectedTransaction.ChainID.Int64(), tx.ChainID.Int64(), i)
		}

		// Hash
		if tx.Hash != blockControllerTestCasesTxFromBlockByIndex[i].expectedTransaction.Hash {
			t.Fatalf("FAIL Hash: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", blockControllerTestCasesTxFromBlockByIndex[i].description, blockControllerTestCasesTxFromBlockByIndex[i].expectedTransaction.Hash, tx.Hash, i)
		}

		// GasPrice
		if tx.GasPrice.Int64() != blockControllerTestCasesTxFromBlockByIndex[i].expectedTransaction.GasPrice.Int64() {
			t.Fatalf("FAIL GasPrice: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", blockControllerTestCasesTxFromBlockByIndex[i].description, blockControllerTestCasesTxFromBlockByIndex[i].expectedTransaction.GasPrice.Int64(), tx.GasPrice.Int64(), i)
		}

		// Nonce
		if tx.Nonce != blockControllerTestCasesTxFromBlockByIndex[i].expectedTransaction.Nonce {
			t.Fatalf("FAIL Nonce: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", blockControllerTestCasesTxFromBlockByIndex[i].description, blockControllerTestCasesTxFromBlockByIndex[i].expectedTransaction.Nonce, tx.Nonce, i)
		}

		t.Logf("Pass: %s", blockControllerTestCasesBlockByNumber[i].description)
	}
}

func TestBlockController_BlockByNumber(t *testing.T) {
	ts := httptest.NewServer(NewRouter())
	defer ts.Close()

	for i := range blockControllerTestCasesBlockByNumber {
		resp, err := http.Get(fmt.Sprintf("%s/block/%s", ts.URL, blockControllerTestCasesBlockByNumber[i].blockNumber.String()))

		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		if resp.StatusCode != blockControllerTestCasesBlockByNumber[i].expectedStatusCode {
			t.Fatalf("Expected status code %d, got %v", blockControllerTestCasesBlockByNumber[i].expectedStatusCode, resp.StatusCode)
		}

		val, ok := resp.Header["Content-Type"]

		// Assert that the "content-type" header is actually set
		if !ok {
			t.Fatalf("Expected Content-Type header to be set")
		}

		// Assert that it was set as expected
		if val[0] != blockControllerTestCasesBlockByNumber[i].expectedContentType {
			t.Fatalf("Expected %s, got %s", blockControllerTestCasesBlockByNumber[i].expectedContentType, val[0])
		}

		if blockControllerTestCasesBlockByNumber[i].expectedBlock == nil {
			t.Logf("Pass: %s", blockControllerTestCasesBlockByNumber[i].description)
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
		if block.Number.Int64() != blockControllerTestCasesBlockByNumber[i].expectedBlock.Number.Int64() {
			t.Fatalf("FAIL Number: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", blockControllerTestCasesBlockByNumber[i].description, blockControllerTestCasesBlockByNumber[i].expectedBlock.Number.Int64(), block.Number.Int64(), i)
		}

		// BaseFee
		if block.BaseFee.Int64() != blockControllerTestCasesBlockByNumber[i].expectedBlock.BaseFee.Int64() {
			t.Fatalf("FAIL BaseFee: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", blockControllerTestCasesBlockByNumber[i].description, blockControllerTestCasesBlockByNumber[i].expectedBlock.BaseFee.Int64(), block.BaseFee.Int64(), i)
		}

		// Difficulty
		if block.Difficulty.Int64() != blockControllerTestCasesBlockByNumber[i].expectedBlock.Difficulty.Int64() {
			t.Fatalf("FAIL Difficulty: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", blockControllerTestCasesBlockByNumber[i].description, blockControllerTestCasesBlockByNumber[i].expectedBlock.Difficulty.Int64(), block.Difficulty.Int64(), i)
		}

		// Nonce
		if block.Nonce != blockControllerTestCasesBlockByNumber[i].expectedBlock.Nonce {
			t.Fatalf("FAIL Nonce: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", blockControllerTestCasesBlockByNumber[i].description, blockControllerTestCasesBlockByNumber[i].expectedBlock.Nonce, block.Nonce, i)
		}

		// Time
		if block.Time != blockControllerTestCasesBlockByNumber[i].expectedBlock.Time {
			t.Fatalf("FAIL Time: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", blockControllerTestCasesBlockByNumber[i].description, blockControllerTestCasesBlockByNumber[i].expectedBlock.Time, block.Time, i)
		}

		t.Logf("Pass: %s", blockControllerTestCasesBlockByNumber[i].description)
	}
}

var blockControllerTestCasesBlockByNumber = []struct {
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
	Gas     uint64   `json:"gas"`
	ChainID *big.Int `json:"chainId"`
	Hash    string   `json:"hash"`
	GasPrice *big.Int `json:"gasPrice"`
	Nonce    uint64   `json:"nonce"`
}

var blockControllerTestCasesTxFromBlockByIndex = []struct {
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
			ChainID:  big.NewInt(1),
			Hash:     "0x856ed296eb7f87619393143cc28ae5705e22866709ff79cd93af71d8132037c0",
			GasPrice: big.NewInt(119586681073),
			Nonce:    uint64(21042),
		},
		expectedContentType: "application/json; charset=utf-8",
		expectedStatusCode:  200,
	},
}

var blockControllerTestCasesTxFromBlockByHash = []struct {
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
			ChainID:  big.NewInt(1),
			Hash:     "0x856ed296eb7f87619393143cc28ae5705e22866709ff79cd93af71d8132037c0",
			GasPrice: big.NewInt(119586681073),
			Nonce:    uint64(21042),
		},
		expectedContentType: "application/json; charset=utf-8",
		expectedStatusCode:  200,
	},
}
