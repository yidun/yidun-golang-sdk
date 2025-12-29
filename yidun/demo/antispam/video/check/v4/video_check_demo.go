package main

import (
	"fmt"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/video"
	request2 "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/video/check/v4/request"
	"log"
	"os"
	"github.com/yidun/yidun-golang-sdk/yidun/core/http"
)

// 点播视频检测提交 demo
func main() {
	// 设置易盾内容安全分配的businessId
	YourBusinessId := os.Getenv("BUSINESS_ID")
	YourSecretId := os.Getenv("SECRET_ID")
	YourSecretKey := os.Getenv("SECRET_KEY")
	request := request2.NewVideoCheckRequest(YourBusinessId)

	// 实例化一个 Client，入参需要传入易盾内容安全分配的secretId，secretKey
	client := video.NewVideoClientWithAccessKey(YourSecretId, YourSecretKey)
	request.SetURL("http://xxxx.mp4")
	request.SetDataID("dataId")
	request.SetUniqueKey("uniqueKey")
	// 设置协议为HTTP
	request.SetProtocol(http.ProtocolEnumHTTP)
	response, err := client.Check(request)
	if err != nil {
		// 处理错误并打印日志
		log.Fatal(err)
	}

	if response.GetCode() == 200 {
		data := response.Result
		fmt.Println("submit taskId:", data.TaskID)

	} else {
		fmt.Println("error code: ", response.GetCode())
		fmt.Println("error msg: ", response.GetMsg())
	}
}
