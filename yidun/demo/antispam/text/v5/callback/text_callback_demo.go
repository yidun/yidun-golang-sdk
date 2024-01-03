package main

import (
	"fmt"
	"log"

	v5 "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/text"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/text/v5/callback"
)

/**
 * 获取文本检测结果demo
 */
func main() {
	// 设置易盾内容安全分配的businessId
	request := callback.NewTextCallBackRequest("YOUR_BUSINESS_ID")

	// 实例化一个textClient，入参需要传入易盾内容安全分配的secretId，secretKey
	textCheckClient := v5.NewTextClientWithAccessKey("YOUR_SECRET_ID", "YOUR_SECRET_KEY")

	request.SetYidunRequestId("唯一ID，标识一次请求，建议使用UUID。可选")

	response, err := textCheckClient.Callback(request)
	if err != nil {
		// 处理错误并打印日志
		log.Fatal(err)
	}

	if response.GetCode() == 200 {
		data := response.Result[0].Antispam
		fmt.Println("Antispam taskId:", *data.TaskID)

	} else {
		fmt.Println("error code: ", response.GetCode())
		fmt.Println("error msg: ", response.GetMsg())
	}
}
