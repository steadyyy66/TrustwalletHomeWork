package parser

import (
	"TrustwalletHomeWork/src/api"
	"TrustwalletHomeWork/src/storage"
	"log/slog"
	"sync"
	"time"
)

type EthereumParserImpl struct {
	mu sync.RWMutex
}

func NewEthereumParser() *EthereumParserImpl {
	return &EthereumParserImpl{}
}

func (p *EthereumParserImpl) GetCurrentBlock() int {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return storage.IStorage.GetCurrentBlock()
}

func (p *EthereumParserImpl) Subscribe(address string) bool {
	p.mu.Lock()
	defer p.mu.Unlock()
	return storage.IStorage.Subscribe(address)
}

func (p *EthereumParserImpl) GetTransactions(address string) []storage.Transaction {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return storage.IStorage.GetTransactions(address)
}

func (p *EthereumParserImpl) WatchBlock() error {
	for {
		latestBlock, err := api.IOutBizApi.GetLatestBlockNumber()
		if err != nil {
			slog.Error("GetLatestBlockNumber：", "err", err)
			return err
		}
		slog.Info("GetLatestBlockNumber：", "latestBlock:", latestBlock, "currentBlock:", storage.IStorage.GetCurrentBlock())

		//it means start running
		if storage.IStorage.GetCurrentBlock() == 0 {
			storage.IStorage.SetCurrentBlock(int(latestBlock))
			continue
		}

		if int(latestBlock) > p.GetCurrentBlock() {
			for blockNum := p.GetCurrentBlock() + 1; blockNum <= int(latestBlock); blockNum++ {
				block, err := api.IOutBizApi.GetBlockByNumber(blockNum)
				if err != nil {
					slog.Error("GetBlockByNumber：", "err", err)
					return err
				}
				p.mu.RLock()
				p.parseTransactions(block)
				p.mu.RUnlock()
			}
			storage.IStorage.SetCurrentBlock(int(latestBlock))
		}

		// Sleep for a while before checking for new blocks
		time.Sleep(1 * time.Second * 5)
	}
}

func (p *EthereumParserImpl) parseTransactions(block *api.GetBlockByNumberRespResult) {

	for _, tx := range block.Transactions {
		p.mu.RLock()
		storage.IStorage.AddStorage(tx.From, tx.To, tx.Value)
		p.mu.RUnlock()
	}
}
