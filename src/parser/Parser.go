package parser

import (
	"TrustwalletHomeWork/src/storage"
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
	IParese = &EthereumParserImpl{}
}
