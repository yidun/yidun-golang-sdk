package main

import (
	"encoding/json"
	"fmt"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/keyword"
	request2 "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/keyword/query/v1/request"
	"log"
	"os"
)

// 查询关键词
func main() {
	YourBusinessId := os.Getenv("BUSINESS_ID")
	YourSecretId := os.Getenv("SECRET_ID")
	YourSecretKey := os.Getenv("SECRET_KEY")
	// 设置易盾内容安全分配的businessId
	// 实例化一个 Client，入参需要传入易盾内容安全分配的secretId，secretKey
	client := keyword.NewKeywordClientWithAccessKey(YourSecretId, YourSecretKey)
	req := request2.NewKeywordQueryRequest(YourBusinessId)
	req.SetKeyword("测试")
	req.SetPageNum(1)
	req.SetPageSize(10)
	response, err := client.Query(req)

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
