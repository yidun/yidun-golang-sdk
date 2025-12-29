package main

import (
	"fmt"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/video"
	request3 "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/video/callback/v4/request"
	"log"
	"os"
	"github.com/yidun/yidun-golang-sdk/yidun/core/http"
)

// 点播视频回调接口
func main() {
	// 设置易盾内容安全分配的businessId
	YourBusinessId := os.Getenv("BUSINESS_ID")
	YourSecretId := os.Getenv("SECRET_ID")
	YourSecretKey := os.Getenv("SECRET_KEY")
	request := request3.NewVideoCallbackV4Request(YourBusinessId)

	// 实例化一个 Client，入参需要传入易盾内容安全分配的secretId，secretKey
	client := video.NewVideoClientWithAccessKey(YourSecretId, YourSecretKey)

	request.SetYidunRequestId("唯一ID，标识一次请求，建议使用UUID。可选")
	// 设置协议为HTTP
	request.SetProtocol(http.ProtocolEnumHTTP)
	response, err := client.Callback(request)
	if err != nil {
		// 处理错误并打印日志
		log.Fatal(err)
	}

	if response.GetCode() == 200 {
		if response.Result == nil || len(*response.Result) == 0 {
			fmt.Println("response empty")
			return
		}
		data := (*response.Result)[0]
		fmt.Println("Antispam taskId:", data.Antispam.TaskID)

	} else {
		fmt.Println("error code: ", response.GetCode())
		fmt.Println("error msg: ", response.GetMsg())
	}
}
