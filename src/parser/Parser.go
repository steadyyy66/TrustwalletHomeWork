package parser

import (
	"TrustwalletHomeWork/src/storage"
	"sync"
)

type Parser interface {

	// last parsed block

	GetCurrentBlock() int

	// add address to observer

	Subscribe(address string) bool

	// list of inbound or outbound transactions for an address
	GetTransactions(address string) []storage.Transaction
}

var IParese Parser

func init() {
	IParese = &EthereumParserImpl{
		mu: sync.RWMutex{},
	}
}
