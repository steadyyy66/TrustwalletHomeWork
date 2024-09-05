# TrustwalletHomeWork

The Ethereum Transaction Parser is a Go-based tool designed to monitor and analyze transactions on the Ethereum blockchain. It provides real-time tracking of transactions for specified Ethereum addresses and offers an efficient way to retrieve and parse transaction data.

Key features:
- Real-time monitoring of Ethereum transactions
- Address subscription system
- Parsing and storage of transaction details
- Concurrent processing for improved performance

Technologies used:
- Go (Golang)
- Ethereum JSON-RPC API

## 2. Installation Guide

### Prerequisites
- Go 1.21 or higher
- Access to an Ethereum node (required for JSON-RPC API access)

### Steps
1. Clone the repository:
```
git clone https://github.com/steadyyy66/TrustwalletHomeWork.git
```
2. Navigate to the project directory:
```
cd github/steadyyy66/TrustwalletHomeWork/src/main
```
3. Build the executable:
```
go build main.go
```
4. Run
```
./main
```

## 4. API Documentation

### Subscribe to an Address
- Endpoint: `Subscribe`
- Request Body: `{"address": "0x..."}` (Ethereum address)
- Response: Boolean indicating success

### Get Transactions by an Address
- Endpoint: `GetTransactions`
- Parameters: `:address` - Ethereum address
- Response: Array of Transaction objects

### Get Current Block Number
- Endpoint: `GetCurrentBlock`
- Response: Current Ethereum block number

## 5. Architecture Design

The Ethereum Transaction Parser consists of the following main components:
1. Parser: Defined and implemented the Parser interface. watch Ethereum’s blocks change.
2. API : HTTP server exposing endpoints for interaction with the parser.
3. Storage: stored the transaction record in memory. But can replace it with a more robust database solution.
4. Config: Definition of constants and variables
5. Tests: Unit tests for the parser components and outer api.
## 7. Testing

To run the tests: