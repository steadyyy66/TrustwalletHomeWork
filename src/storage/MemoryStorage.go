package storage

type MemoryStorage struct {
	currentBlock int
	subscribers  map[string]bool
	transactions map[string][]Transaction
}

func (p *MemoryStorage) AddStorage(from, to, value string) {
	t := Transaction{
		From:  from,
		To:    to,
		Value: value,
	}
	if p.subscribers[from] {
		p.transactions[from] = append(p.transactions[from], t)
	}
	if p.subscribers[to] {
		p.transactions[to] = append(p.transactions[to], t)
	}

}

func (p *MemoryStorage) Subscribe(address string) bool {
	if _, exists := p.subscribers[address]; exists {
		return false
	}
	p.subscribers[address] = true
	return true
}

func (p *MemoryStorage) GetCurrentBlock() int {
	return p.currentBlock
}

func (p *MemoryStorage) SetCurrentBlock(currentBlock int) bool {
	p.currentBlock = currentBlock
	return true
}
func (p *MemoryStorage) GetTransactions(address string) []Transaction {
	return p.transactions[address]
}
