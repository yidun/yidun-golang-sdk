package main

import (
	"encoding/json"
	"fmt"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/videosolution"
	request3 "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/videosolution/callback/v2/request"
	"log"
	"os"
	"github.com/yidun/yidun-golang-sdk/yidun/core/http"
)

// 点播音视频解决方案回调接口
func main() {
	// 设置易盾内容安全分配的businessId
	YourSecretId := os.Getenv("VS_SECRET_ID")
	YourSecretKey := os.Getenv("VS_SECRET_KEY")
	request := request3.NewVideoSolutionCallbackV2Request()

	// 实例化一个 Client，入参需要传入易盾内容安全分配的secretId，secretKey
	client := videosolution.NewVideoSolutionClientWithAccessKey(YourSecretId, YourSecretKey)

	request.SetYidunRequestId("xxx")
	// 设置协议为HTTP
	request.SetProtocol(http.ProtocolEnumHTTP)
	response, err := client.Callback(request)
	if err != nil {
		// 处理错误并打印日志
		log.Fatal(err)
	}

	if response.GetCode() == 200 {
		result := response.Result
		resultString, _ := json.Marshal(result)
		fmt.Println("callback result: ", string(resultString))

	} else {
		fmt.Println("error code: ", response.GetCode())
		fmt.Println("error msg: ", response.GetMsg())
	}
}
