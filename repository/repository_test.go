package repository

import (
	"errors"
	"eth-caching-proxy/model"
	"math/big"
	"testing"
)

func TestBlockRepository(t *testing.T) {
	for i := range repositoryTestCases {
		repo := New()
		block, err := repo.BlockRepository.BlockByNumber(repositoryTestCases[i].blockNumber)

		if err != nil {
			if repositoryTestCases[i].expectedError == nil {
				t.Fatalf("Unexpected error occured: %v", err)
			}

			if err.Error() != repositoryTestCases[i].expectedError.Error() {
				t.Fatalf("FAIL Error: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", repositoryTestCases[i].description, repositoryTestCases[i].expectedError, err, i)
			}
		}

		if block != nil && repositoryTestCases[i].expectedBlock != nil {
			// Number
			if block.Number.Int64() != repositoryTestCases[i].expectedBlock.Number.Int64() {
				t.Fatalf("FAIL Number: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", repositoryTestCases[i].description, repositoryTestCases[i].expectedBlock.Number.Int64(), block.Number.Int64(), i)
			}

			// BaseFee
			if block.BaseFee.Int64() != repositoryTestCases[i].expectedBlock.BaseFee.Int64() {
				t.Fatalf("FAIL BaseFee: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", repositoryTestCases[i].description, repositoryTestCases[i].expectedBlock.BaseFee.Int64(), block.BaseFee.Int64(), i)
			}

			// Difficulty
			if block.Difficulty.Int64() != repositoryTestCases[i].expectedBlock.Difficulty.Int64() {
				t.Fatalf("FAIL Difficulty: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", repositoryTestCases[i].description, repositoryTestCases[i].expectedBlock.Difficulty.Int64(), block.Difficulty.Int64(), i)
			}

			// Nonce
			if block.Nonce != repositoryTestCases[i].expectedBlock.Nonce {
				t.Fatalf("FAIL Nonce: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", repositoryTestCases[i].description, repositoryTestCases[i].expectedBlock.Nonce, block.Nonce, i)
			}

			// Time
			if block.Time != repositoryTestCases[i].expectedBlock.Time {
				t.Fatalf("FAIL Time: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", repositoryTestCases[i].description, repositoryTestCases[i].expectedBlock.Time, block.Time, i)
			}
		}

		t.Logf("Pass: %s", repositoryTestCases[i].description)
	}
}

var repositoryTestCases = []struct {
	description   string
	blockNumber   *big.Int
	expectedBlock *model.Block
	expectedError error
}{
	{
		description:   "No block found.",
		blockNumber:   big.NewInt(111111111),
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
