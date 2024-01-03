package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"
)

type MyResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		ConfigData string `json:"configData"`
	} `json:"data"`
	Ok bool `json:"ok"`
}

func TestClient(t *testing.T) {

	// jsonStr := "{\"sdkData\":\"xxxxxxxxxxxx\",\"signature\":\"5bf1755ba99c43787818891eeaae4917\",\"ip\":\"192.168.0.1\",\"businessId\":\"fe12642085e736d3f69fe7f44e0d925b\",\"secretId\":\"bd085faf9882b02b868e9290bd22df06\",\"version\":\"500\",\"nonce\":\"02551be7528e4632818a7b2873924491\",\"timestamp\":\"1693741054727\"}"
	url1 := "https://ir-open.dun.163.com/v5/risk/getConfig"

	// 定义一个map[string]interface{}类型的数据
	data := map[string]interface{}{
		"sdkData":    "xxxxxxxxxxxx",
		"signature":  "5bf1755ba99c43787818891eeaae4917",
		"ip":         "192.168.0.1",
		"businessId": "fe12642085e736d3f69fe7f44e0d925b",
		"secretId":   "bd085faf9882b02b868e9290bd22df06",
		"version":    "500",
		"nonce":      "02551be7528e4632818a7b2873924491",
		"timestamp":  "1693741054727",
	}

	// 将map[string]interface{}类型的数据转换为JSON格式的字节数组
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Marshal error:", err)
		return
	}

	// 将url.Values类型的数据转换为io.Reader类型的数据
	body := bytes.NewBuffer(jsonBytes)
	req, err := http.NewRequest("POST", url1, body)
	if err != nil {
		fmt.Println("NewRequest error:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json;chatset=utf-8")

	// 发送请求
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Do error:", err)
		return
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ReadAll error:", err)
		return
	}

	var result MyResponse
	if err := json.Unmarshal(responseBody, &result); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)

	// // 发送表单POST请求
	// formOptions := &myhttp.FormOptions{
	// 	Method: "POST",
	// 	Url:    "https://www.baidu.com",
	// 	Headers: map[string]string{
	// 		"Content-Type": "application/x-www-form-urlencoded",
	// 	},
	// 	Body: map[string][]string{
	// 		"name":  {"Test"},
	// 		"value": {"123"},
	// 	},
	// }

	// formResponse := &myhttp.CommonResponse{}
	// err = httpClient.Form(formOptions, formResponse)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(formResponse.StatusCode())
	// fmt.Println(string(formResponse.Body()))
}
