package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/list"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/list/delete/v2/request"
)

// 删除名单
func main() {
	YourBusinessId := os.Getenv("BUSINESS_ID")
	YourSecretId := os.Getenv("SECRET_ID")
	YourSecretKey := os.Getenv("SECRET_KEY")
	// 设置易盾内容安全分配的businessId
	// 实例化一个 Client，入参需要传入易盾内容安全分配的secretId，secretKey
	client := list.NewListClientWithAccessKey(YourSecretId, YourSecretKey)
	req := request.NewListDeleteRequest(YourBusinessId)
	// 构造字符串数组
	stringArray := []string{"apple", "banana", "cherry"}
	// 序列化成JSON字符串
	entities, err := json.Marshal(stringArray)
	req.SetEntities(string(entities))
	req.SetEntityType(1)
	req.SetListType(1)
	response, err := client.Delete(req)

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
