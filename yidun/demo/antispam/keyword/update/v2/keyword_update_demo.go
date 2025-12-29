package main

import (
	"encoding/json"
	"fmt"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/keyword"
	request2 "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/keyword/update/v2/request"
	"log"
	"os"
	"github.com/yidun/yidun-golang-sdk/yidun/core/http"
)

// 更新关键词
func main() {
	YourBusinessId := os.Getenv("BUSINESS_ID")
	YourSecretId := os.Getenv("SECRET_ID")
	YourSecretKey := os.Getenv("SECRET_KEY")
	// 设置易盾内容安全分配的businessId
	// 实例化一个 Client，入参需要传入易盾内容安全分配的secretId，secretKey
	client := keyword.NewKeywordClientWithAccessKey(YourSecretId, YourSecretKey)
	req := request2.NewKeywordUpdateRequest(YourBusinessId)
	req.SetKeywords([]string{"测试", "测试2"})
	req.SetLevel(2)
	req.SetCategory(300)
	// 设置协议为HTTP
	req.SetProtocol(http.ProtocolEnumHTTP)
	response, err := client.Update(req)

	if err != nil {
		// 处理错误并打印日志
		log.Fatal(err)
	}

	if response.GetCode() == 200 {
		result := response.Result
		resultJson, _ := json.Marshal(result)
		fmt.Println("result: ", string(resultJson))
	} else {
		fmt.Println("error code: ", response.GetCode())
		fmt.Println("error msg: ", response.GetMsg())
	}
}
