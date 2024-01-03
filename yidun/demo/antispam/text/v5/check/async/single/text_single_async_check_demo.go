package main

import (
	"fmt"
	"log"

	v5 "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/text"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/text/v5/check/async/single"
)

/**
 * 文本单次异步检测demo
 */
func main() {
	// 设置易盾内容安全分配的businessId
	request := single.NewTextAsyncCheckRequest("YOUR_BUSINESS_ID")
	
	// 实例化一个textClient，入参需要传入易盾内容安全分配的secretId，secretKey
	textCheckClient := v5.NewTextClientWithAccessKey("YOUR_SECRET_ID", "YOUR_SECRET_KEY")

	request.SetDataID("DATA_ID")
	request.SetContent("/ / %【 玉 米 】 ， ， ， 。 。 。")
	request.SetCheckLabels("100,200")

	response, err := textCheckClient.AsyncCheckText(request)
	if err != nil {
		// 处理错误并打印日志
		log.Fatal(err)
	}

	if response.GetCode() == 200 {
		data := response.Result
		fmt.Println("Antispam dealingCount:", *data.DealingCount)
		fmt.Println("Antispam taskId:", *data.CheckTexts[0].TaskID)

	} else {
		fmt.Println("error code: ", response.GetCode())
		fmt.Println("error msg: ", response.GetMsg())
	}
}
