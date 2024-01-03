package main

import (
	"encoding/json"
	"fmt"
	"log"

	digital "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/digitalbook/v2"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/digitalbook/v2/query"
)

const (
	SECRET_ID  = "YOUR_SECRET_ID"
	SECRET_KEY = "YOUR_SECRET_KEY"
)

/**
 * 数字阅读回调结果查询demo
 */
func main() {
	// 构造medai callback 请求对象
	digitalQueryReq := query.NewDigitalBookQueryRequest()

	// 实例化一个Client，入参需要传入易盾内容安全分配的secretId，secretKey
	qeueryClient := digital.NewDigitalBookClientWithAccessKey(SECRET_ID, SECRET_KEY)

	// Set TaskIds
	taskIds := []string{"task1", "task2", "task3"}
	digitalQueryReq.SetTaskIds(&taskIds)

	response, err := qeueryClient.Query(digitalQueryReq)
	if err != nil {
		// 处理错误并打印日志
		log.Fatal(err)
	}

	if response.GetCode() == 200 {
		// do something
		valueJson, _ := json.Marshal(response.Result)
		fmt.Println("Antispam: ", string(valueJson))
	} else {
		fmt.Println("error code: ", response.GetCode())
		fmt.Println("error msg: ", response.GetMsg())
	}
}
