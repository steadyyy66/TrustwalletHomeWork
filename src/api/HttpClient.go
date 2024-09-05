package api

import (
	"TrustwalletHomeWork/src/config"
	"bytes"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
)

type OutBizApi interface {
	GetLatestBlockNumber() (int64, error)
	GetBlockByNumber(blockNumber int) (*GetBlockByNumberRespResult, error)
}
type OutBizApiImpl struct {
}

var IOutBizApi OutBizApi

func init() {
	IOutBizApi = new(OutBizApiImpl)
}

func (p *OutBizApiImpl) GetLatestBlockNumber() (int64, error) {

	payload := map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      1,
		"method":  config.ETH_BLOCKNUMBER,
		"params":  []interface{}{},
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return config.JSON_MARSHAL, err
	}
	slog.Debug("Before request GetLatestBlockNumber", "result", jsonPayload, "err", err)

	resp, err := http.Post(config.ETH_URL, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		slog.Error("request error:", "err", err)
		return config.DEFALUT_ERR_NUMBER, err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		slog.Error("NewDecoder error:", "err msg", err)
		return config.JSON_DECODE, err
	}

	slog.Debug("Received response from GetLatestBlockNumber", "result", result, "err", err)

	return HexToInt64(result["result"].(string))
}

func (p *OutBizApiImpl) GetBlockByNumber(blockNumber int) (*GetBlockByNumberRespResult, error) {

	params := []interface{}{fmt.Sprintf("0x%x", blockNumber), true}
	payload := map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      1,
		"method":  config.ETH_GETBLOCKBYNUMBER,
		"params":  params,
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	slog.Debug("Before send request GetBlockByNumber", "params", params, "err", err)
	resp, err := http.Post(config.ETH_URL, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result GetBlockByNumberResp

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	slog.Debug("Received response from GetBlockByNumber", "len(result.Result.Transactions):", len(result.Result.Transactions), "err", err)

	return &result.Result, err
}

func HexToInt64(hex string) (int64, error) {
	return strconv.ParseInt(hex[2:], 16, 64)
}
