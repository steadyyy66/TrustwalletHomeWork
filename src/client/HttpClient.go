package client

import (
	"TrustwalletHomeWork/src/constants"
	"bytes"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
)

func GetLatestBlockNumber() (int64, error) {

	payload := map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      1,
		"method":  constants.ETH_BLOCKNUMBER,
		"params":  []interface{}{},
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return constants.DEFALUT_ERR_NUMBER, err
	}
	slog.Debug("Received response from Ethereum node", "result", jsonPayload, "err", err)

	resp, err := http.Post(constants.ETH_URL, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		slog.Error("request error:", "err", err)
		return constants.DEFALUT_ERR_NUMBER, err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		slog.Error("NewDecoder error:", err)
		return constants.DEFALUT_ERR_NUMBER, err
	}

	slog.Debug("Received response from Ethereum node", "result", result, "err", err)

	return hexToInt64(result["result"].(string))
}

func GetBlockByNumber(blockNumber int) (*GetBlockByNumberRespResult, error) {

	params := []interface{}{fmt.Sprintf("0x%x", blockNumber), true}
	payload := map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      1,
		"method":  constants.ETH_GETBLOCKBYNUMBER,
		"params":  params,
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	slog.Debug("Received response from Ethereum node", "params", params, "err", err)
	resp, err := http.Post(constants.ETH_URL, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result GetBlockByNumberResp

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	slog.Debug("Received response from Ethereum node", "result", result, "err", err)

	return &result.Result, err
}

func hexToInt64(hex string) (int64, error) {
	return strconv.ParseInt(hex[2:], 16, 64)
}
