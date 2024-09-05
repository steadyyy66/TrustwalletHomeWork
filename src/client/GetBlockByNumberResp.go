package client

// define data stracture for GetBlockByNumberResp
// remote many unrelated filed
type GetBlockByNumberResp struct {
	Jsonrpc string                     `json:"jsonrpc"`
	Result  GetBlockByNumberRespResult `json:"result"`
	ID      int                        `json:"id"`
}

type GetBlockByNumberRespResult struct {
	Transactions []GetBlockByNumberRespResultTransactions `json:"transactions"`
}

type GetBlockByNumberRespResultTransactions struct {
	BlockHash   string `json:"blockHash"`
	BlockNumber string `json:"blockNumber"`
	From        string `json:"from"`
	Hash        string `json:"hash"`
	Input       string `json:"input"`
	To          string `json:"to"`
	Value       string `json:"value"`
}
