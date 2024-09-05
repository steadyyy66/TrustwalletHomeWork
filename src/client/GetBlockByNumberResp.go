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
	BlockNumber string `json:"blockNumber"`
	From        string `json:"from"`
	To          string `json:"to"`
	Value       string `json:"value"`
}
