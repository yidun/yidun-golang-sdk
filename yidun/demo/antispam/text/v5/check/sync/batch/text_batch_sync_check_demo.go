package main

import (
	"fmt"
	"log"

	"github.com/yidun/yidun-golang-sdk/yidun/core/http"
	v5 "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/text"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/text/v5/check/sync/batch"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/text/v5/check/sync/single"
)

/**
 * 文本批量同步检测demo
 */
func main() {
	// 设置易盾内容安全分配的businessId
	request := batch.NewTextBatchCheckRequest("YOUR_BUSINESS_ID")

	//创建单条过检请求
	singleRequest := single.NewTextSceneRequest()
	singleRequest.SetDataID("DATA_ID")
	singleRequest.SetContent("/ / %【 玉 米 】 ， ， ， 。 。 。")

	singleRequests := []*single.TextSceneRequest{singleRequest}

	request.SetTexts(singleRequests)
	// 设置协议为HTTP
	request.SetProtocol(http.ProtocolEnumHTTP)

	// 实例化一个textClient，入参需要传入易盾内容安全分配的secretId，secretKey
	textCheckClient := v5.NewTextClientWithAccessKey("YOUR_SECRET_ID", "YOUR_SECRET_KEY")

	response, err := textCheckClient.SyncBatchCheckText(request)
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
