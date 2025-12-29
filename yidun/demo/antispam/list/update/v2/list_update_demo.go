package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/list"
	request2 "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/list/update/v2/request"
	"github.com/yidun/yidun-golang-sdk/yidun/core/http"
)

// 更新名单
func main() {
	YourBusinessId := os.Getenv("BUSINESS_ID")
	YourSecretId := os.Getenv("SECRET_ID")
	YourSecretKey := os.Getenv("SECRET_KEY")
	// 设置易盾内容安全分配的businessId
	// 实例化一个 Client，入参需要传入易盾内容安全分配的secretId，secretKey
	client := list.NewListClientWithAccessKey(YourSecretId, YourSecretKey)
	req := request2.NewListUpdateRequest(YourBusinessId)
	req.SetEntity("apple")
	req.SetEntityType(1)
	req.SetListType(1)
	req.SetStatus(0)
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
