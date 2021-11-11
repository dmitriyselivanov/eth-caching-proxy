package cache

import (
	"eth-caching-proxy/model"
	"math/big"
	"testing"
)

func TestBlockCache(t *testing.T) {
	for i := range cacheTestCases {
		c := New()

		if cacheTestCases[i].existsInCache {
			c.BlockCache.AddBlock(cacheTestCases[i].expectedBlock)
		}
		block, exists := c.BlockCache.GetBlock(cacheTestCases[i].blockNumber)

		if exists != cacheTestCases[i].existsInCache {
			t.Fatalf("FAIL Found: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", cacheTestCases[i].description, cacheTestCases[i].existsInCache, exists, i)
		}

		if block != nil && cacheTestCases[i].expectedBlock != nil {
			// Number
			if block.Number.Int64() != cacheTestCases[i].expectedBlock.Number.Int64() {
				t.Fatalf("FAIL Number: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", cacheTestCases[i].description, cacheTestCases[i].expectedBlock.Number.Int64(), block.Number.Int64(), i)
			}

			// BaseFee
			if block.BaseFee.Int64() != cacheTestCases[i].expectedBlock.BaseFee.Int64() {
				t.Fatalf("FAIL BaseFee: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", cacheTestCases[i].description, cacheTestCases[i].expectedBlock.BaseFee.Int64(), block.BaseFee.Int64(), i)
			}

			// Difficulty
			if block.Difficulty.Int64() != cacheTestCases[i].expectedBlock.Difficulty.Int64() {
				t.Fatalf("FAIL Difficulty: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", cacheTestCases[i].description, cacheTestCases[i].expectedBlock.Difficulty.Int64(), block.Difficulty.Int64(), i)
			}

			// Nonce
			if block.Nonce != cacheTestCases[i].expectedBlock.Nonce {
				t.Fatalf("FAIL Nonce: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", cacheTestCases[i].description, cacheTestCases[i].expectedBlock.Nonce, block.Nonce, i)
			}

			// Time
			if block.Time != cacheTestCases[i].expectedBlock.Time {
				t.Fatalf("FAIL Time: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", cacheTestCases[i].description, cacheTestCases[i].expectedBlock.Time, block.Time, i)
			}
		}

		t.Logf("Pass: %s", cacheTestCases[i].description)
	}
}

var cacheTestCases = []struct {
	description   string
	blockNumber   *big.Int
	expectedBlock *model.Block
	existsInCache bool
}{
	{
		description:   "Block was not found",
		blockNumber:   big.NewInt(2),
		existsInCache: false,
		expectedBlock: nil,
	},

	{
		description:   "Received a correct block",
		blockNumber:   big.NewInt(13588593),
		existsInCache: true,
		expectedBlock: &model.Block{
			Number:     big.NewInt(13588593),
			BaseFee:    big.NewInt(105410383176),
			Difficulty: big.NewInt(10506939366079207),
			Nonce:      uint64(17499604187474856417),
			Time:       uint64(1636548900),
		},
	},
}
