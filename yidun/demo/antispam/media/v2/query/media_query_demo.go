package main

import (
	"encoding/json"
	"fmt"
	"log"

	media "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/media/v2"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/media/v2/query"
)

const (
	SECRET_ID  = "YOUR_SECRET_ID"
	SECRET_KEY = "YOUR_SECRET_KEY"
)

/**
 * 融媒体回调结果查询demo
 */
func main() {
	// 构造medai callback 请求对象
	request := query.NewMediaQueryRequestV2()
	/**
	 * 数据提交接口协议的版本号, 默认值v2.1，注意与提交时的版本号保持一致
	 */
	//request.SetVersion("v2.1")

	// Set TaskIds
	taskIds := []string{"task1", "task2", "task3"}
	request.SetTaskIds(&taskIds)

	// 实例化一个Client，入参需要传入易盾内容安全分配的secretId，secretKey
	qeueryClient := media.NewMediaClientWithAccessKey(SECRET_ID, SECRET_KEY)
	response, err := qeueryClient.Query(request)
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
