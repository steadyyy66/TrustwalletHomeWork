package test

import (
	"TrustwalletHomeWork/src/api"
	"TrustwalletHomeWork/src/config"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestGetLatestBlockNumber(t *testing.T) {
	// 创建测试用例
	testCases := []struct {
		name           string
		mockResponse   string
		expectedResult int64
		expectedError  error
	}{
		{
			name:           "success request",
			mockResponse:   `{"jsonrpc":"2.0","id":1,"result":"0x10"}`,
			expectedResult: 16,
		},
		{
			name:           "Returns the wrong json format",
			mockResponse:   `{"jsonrpc":"2.0","id":1,"result":0x10}`,
			expectedResult: config.JSON_DECODE,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 创建模拟的http服务器
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintln(w, tc.mockResponse)
			}))
			defer ts.Close()

			// 替换常量中的URL为模拟服务器的URL
			oldUrl := "localhost"
			config.ETH_URL = ts.URL
			defer func() { config.ETH_URL = oldUrl }()
			ts.URL = oldUrl
			// 创建待测试的对象
			p := &api.OutBizApiImpl{}

			// 调用待测试的方法
			result, err := p.GetLatestBlockNumber()

			if tc.expectedResult != result {
				t.Error("err", err)
			}

		})
	}
}

func TestGetBlockByNumber(t *testing.T) {

	p := &api.OutBizApiImpl{}
	intNumber, err := api.HexToInt64("0x13b98c7")
	if err != nil {
		t.Errorf("change number fail:%v", err)
	}
	// 调用待测试的方法
	result, err := p.GetBlockByNumber(int(intNumber))
	if err != nil {
		t.Errorf("GetBlockByNumber err:%v", err)
	}
	resp, err := ReadJson(t)
	if err != nil {
		t.Errorf("ReadJson err:%v", err)
	}
	for i := range resp.Result.Transactions {
		if result.Transactions[i].Value != resp.Result.Transactions[i].Value {
			t.Errorf("compare value fail:%s,%s", result.Transactions[i].Value, resp.Result.Transactions[i].Value)
		}

		if result.Transactions[i].From != resp.Result.Transactions[i].From {
			t.Errorf("compare value fail:%s,%s", result.Transactions[i].From, resp.Result.Transactions[i].From)
		}

		if result.Transactions[i].To != resp.Result.Transactions[i].To {
			t.Errorf("compare To fail:%s,%s", result.Transactions[i].To, resp.Result.Transactions[i].To)
		}
	}
}

func ReadJson(t *testing.T) (*api.GetBlockByNumberResp, error) {
	// 打开JSON文件
	file, err := os.Open("data.json")
	if err != nil {
		log.Fatalf("fail to open file: %v", err)
	}
	defer file.Close()

	// 读取文件内容
	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		t.Errorf(err.Error())
	}

	// 创建一个Person实例用于存储解码后的数据
	var getBlockResp api.GetBlockByNumberResp

	// 将JSON数据反序列化为结构体
	if err := json.Unmarshal(byteValue, &getBlockResp); err != nil {
		t.Errorf("fail to deserialize JSON data: %v", err)
		return nil, err
	}
	//t.Logf("Name: %s\nAge: %d\nEmail: %s\n", getBlockResp.Result.Transactions[0].BlockNumber, getBlockResp.Result.Transactions[0].From, getBlockResp.Result.Transactions[0].To)
	// 输出反序列化后的数据
	t.Log(getBlockResp)
	return &getBlockResp, nil
}
