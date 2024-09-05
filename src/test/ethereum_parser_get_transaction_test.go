package main

import (
	"TrustwalletHomeWork/src/client"
	parser2 "TrustwalletHomeWork/src/parser"
	"testing"
	"time"
)

/*
oct   hex
4660 = 1234
*/
type MockOutBiz4TestGetTransactions struct {
}

var mockFirstCallGetLatestBlockNumber bool

func (m MockOutBiz4TestGetTransactions) GetLatestBlockNumber() (int64, error) {
	if mockFirstCallGetLatestBlockNumber == false {
		mockFirstCallGetLatestBlockNumber = true
		return 4660 - 1, nil
	}
	return 4660, nil
}

var httpTransactions []client.GetBlockByNumberRespResultTransactions

/*
adress
*/
func init() {
	txn1 := client.GetBlockByNumberRespResultTransactions{
		From:  "0x1234",
		To:    "0x5678",
		Value: "1111",
	}
	txn2 := client.GetBlockByNumberRespResultTransactions{
		From:  "0x1234",
		To:    "0x9abc",
		Value: "2222",
	}

	txn3 := client.GetBlockByNumberRespResultTransactions{
		From:  "0x5678",
		To:    "0x1234",
		Value: "3333",
	}
	httpTransactions = append(httpTransactions, txn1, txn2, txn3)

}

func (m MockOutBiz4TestGetTransactions) GetBlockByNumber(blockNumber int) (*client.GetBlockByNumberRespResult, error) {

	resp := &client.GetBlockByNumberRespResult{
		Transactions: httpTransactions,
	}

	return resp, nil
}

/*
TestGetTransactions ,also test Subscribe
*/
func TestGetTransactions(t *testing.T) {

	client.IOutBizApi = new(MockOutBiz4TestGetTransactions)

	parser := parser2.NewEthereumParser()
	parser.Subscribe("0x1234")
	parser.Subscribe("0x5678")
	go parser.WatchBlock()
	time.Sleep(time.Second * 3)
	storageTransactionList := parser.GetTransactions("0x1234")

	if len(storageTransactionList) != 3 {
		t.Errorf("Expected current block to be 4660, got %d", len(storageTransactionList))
	}
	for _, transaction := range storageTransactionList {
		t.Logf("transaction.To %s,transaction.From %s, transaction.To %s", transaction.To, transaction.From, transaction.Value)
		if transaction.From == "0x1234" && transaction.To == "0x5678" && transaction.Value == "1111" { // 0x1234 in decimal
			continue
		}

		if transaction.From == "0x1234" && transaction.To == "0x9abc" && transaction.Value == "2222" { // 0x1234 in decimal
			continue
		}

		if transaction.From == "0x5678" && transaction.To == "0x1234" && transaction.Value == "3333" { // 0x1234 in decimal
			continue
		}
		t.Errorf("UnExpected storageTransactionList")
	}

}
