package main

import (
	"fmt"
	"log"

	"github.com/yidun/yidun-golang-sdk/yidun/core/http"
	v5 "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/text"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/text/v5/check/async/batch"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/text/v5/check/sync/single"
)

/**
 * 文本批量异步检测demo
 */
func main() {
	// 设置易盾内容安全分配的businessId
	request := batch.NewTextAsyncBatchCheckRequest("YOUR_BUSINESS_ID")

	//创建单条过检请求
	singleRequest := single.NewTextSceneRequest()
	singleRequest.SetDataID("DATA_ID")
	singleRequest.SetContent("/ / %【 玉 米 】 ， ， ， 。 。 。")

	singleRequest2 := single.NewTextSceneRequest()
	singleRequest2.SetDataID("aaabbbccc")
	singleRequest2.SetContent("/ / %【 玉 鱼鱼鱼鱼鱼米 】 ， ， ， 。 。 。")

	singleRequests := []*single.TextSceneRequest{singleRequest, singleRequest2}

	request.SetTexts(singleRequests)
	// 设置协议为HTTP
	request.SetProtocol(http.ProtocolEnumHTTP)

	// 实例化一个textClient，入参需要传入易盾内容安全分配的secretId，secretKey
	textCheckClient := v5.NewTextClientWithAccessKey("YOUR_SECRET_ID", "YOUR_SECRET_KEY")

	response, err := textCheckClient.AsyncBatchCheckText(request)
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
