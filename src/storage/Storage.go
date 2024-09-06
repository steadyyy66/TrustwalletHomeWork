package storage

func init() {
	IStorage = &MemoryStorage{
		transactions: make(map[string][]Transaction),
		subscribers:  make(map[string]bool),
		currentBlock: 0,
	}
}

type Transaction struct {
	From  string
	To    string
	Value string
}

// Use in-memory storage now and replace with other data sources in future
type Storage interface {
	AddStorage(from, to, value string)
	GetTransactions(address string) []Transaction
	Subscribe(address string) bool
	SetCurrentBlock(currentBlock int) bool
	GetCurrentBlock() int
}

var IStorage Storage
