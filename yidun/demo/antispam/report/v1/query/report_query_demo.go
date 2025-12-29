package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/yidun/yidun-golang-sdk/yidun/core/http"
	report "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/report/v1"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/report/v1/query"
)

const (
	SECRET_ID  = "YOUR_SECRET_ID"
	SECRET_KEY = "YOUR_SECRET_KEY"
)

/**
 * 投诉举报回调结果查询demo
 */
func main() {
	// 构造medai callback 请求对象
	reportReq := query.NewReportQueryRequestV1()

	// Set TaskIds
	taskIds := []string{"task1", "task2", "task3"}
	reportReq.SetTaskIds(&taskIds)
	// 设置协议为HTTP
	reportReq.SetProtocol(http.ProtocolEnumHTTP)

	// 实例化一个Client，入参需要传入易盾内容安全分配的secretId，secretKey
	qeueryClient := report.NewReportClientWithAccessKey(SECRET_ID, SECRET_KEY)
	response, err := qeueryClient.Query(reportReq)
	if err != nil {
		// 处理错误并打印日志
		log.Fatal(err)
	}

	if response.GetCode() == 200 {
		// do something
		valueJson, _ := json.Marshal(response.Result)
		fmt.Println("Antispam:", string(valueJson))
	} else {
		fmt.Println("error code: ", response.GetCode())
		fmt.Println("error msg: ", response.GetMsg())
	}
}
