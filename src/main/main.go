package main

import (
	parser2 "TrustwalletHomeWork/src/parser"
	"log/slog"
	"os"
	"time"
)

func main() {
	ethParser := parser2.NewEthereumParser()

	opts := slog.HandlerOptions{AddSource: true, Level: slog.LevelDebug}
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, &opts)))
	// Start parsing blocks in a separate goroutine
	go func() {
		err := ethParser.WatchBlock()
		if err != nil {
			slog.Error("Error parsing blocks: %v", err)
			//todo should panic?
		}
	}()
	slog.Info("server started")
	// Example usage
	ethParser.Subscribe("0x13b98c8")

	// Wait for some time to allow parsing
	time.Sleep(2 * time.Minute)

	slog.Info("Current block: %d", ethParser.GetCurrentBlock())
	transactions := ethParser.GetTransactions("0x13b98c8")
	slog.Info("Transactions for 0x13b98c8: %+v", transactions)
}
