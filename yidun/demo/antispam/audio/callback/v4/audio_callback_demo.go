package main

import (
	"fmt"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/audio"
	request2 "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/audio/callback/v4/request"
	"log"
	"os"
)

/**
 * 音频回调demo
 */
func main() {
	// 设置易盾内容安全分配的businessId
	YourBusinessId := os.Getenv("BUSINESS_ID")
	YourSecretId := os.Getenv("SECRET_ID")
	YourSecretKey := os.Getenv("SECRET_KEY")
	request := request2.NewAudioCallBackRequest(YourBusinessId)

	// 实例化一个 Client，入参需要传入易盾内容安全分配的secretId，secretKey
	client := audio.NewAudioClientWithAccessKey(YourSecretId, YourSecretKey)

	request.SetYidunRequestId("唯一ID，标识一次请求，建议使用UUID。可选")

	response, err := client.Callback(request)
	if err != nil {
		// 处理错误并打印日志
		log.Fatal(err)
	}

	if response.GetCode() == 200 {
		data := response.Result[0].Antispam
		fmt.Println("Antispam taskId:", *data.TaskId)

	} else {
		fmt.Println("error code: ", response.GetCode())
		fmt.Println("error msg: ", response.GetMsg())
	}
}
