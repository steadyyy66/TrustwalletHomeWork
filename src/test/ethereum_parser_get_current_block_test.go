package test

//This is file for TestGetCurrentBlock unit test

import (
	"TrustwalletHomeWork/src/api"
	. "TrustwalletHomeWork/src/parser"
	"testing"
	"time"
)

// Mock OutBiz for TestGetCurrentBlock
type MockOutBiz4TestGetCurrentBlock struct {
}

func (m MockOutBiz4TestGetCurrentBlock) GetLatestBlockNumber() (int64, error) {
	return 4660, nil
}

func (m MockOutBiz4TestGetCurrentBlock) GetBlockByNumber(blockNumber int) (*api.GetBlockByNumberRespResult, error) {
	panic("implement me")
}

func TestGetCurrentBlock(t *testing.T) {

	api.IOutBizApi = new(MockOutBiz4TestGetCurrentBlock)

	parser := NewEthereumParser()
	go parser.WatchBlock()

	time.Sleep(time.Second)
	currentBlock := parser.GetCurrentBlock()

	if currentBlock != 4660 { // 0x1234 in decimal
		t.Errorf("Expected current block to be 4660, got %d", currentBlock)
	} else {
		t.Logf("Expected current block to be 4660, got %d", currentBlock)
		return
	}
}

func TestSubscribe(t *testing.T) {
	parser := NewEthereumParser() // URL doesn't matter for this test

	// Test subscribing a new address
	result := parser.Subscribe("0x123")
	if !result {
		t.Errorf("Expected Subscribe to return true for new address")
	}

	// Test subscribing the same address again
	result = parser.Subscribe("0x123")
	if result {
		t.Errorf("Expected Subscribe to return false for already subscribed address")
	}
}
