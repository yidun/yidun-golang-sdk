package main

import (
	"fmt"
	"log"

	v1 "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/text"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/text/query"
)

// 文本taskId查询请求Demo
func main() {
	// 设置易盾内容安全分配的businessId
	request := query.NewTextTaskIdsQueryRequest("YOUR_BUSINESS_ID")

	// 实例化一个 Client，入参需要传入易盾内容安全分配的secretId，secretKey
	textCheckClient := v1.NewTextClientWithAccessKey("YOUR_SECRET_ID", "YOUR_SECRET_KEY")

	taskIdList := []string{"taskIdA", "taskIdB"}
	request.SetTaskIds(taskIdList)

	response, err := textCheckClient.QueryTaskIds(request)
	if err != nil {
		// 处理错误并打印日志
		log.Fatal(err)
	}

	if response.GetCode() == 200 {
		data := response.Result[0]
		fmt.Println("Antispam taskId:", *data.TaskID)

	} else {
		fmt.Println("error code: ", response.GetCode())
		fmt.Println("error msg: ", response.GetMsg())
	}
}