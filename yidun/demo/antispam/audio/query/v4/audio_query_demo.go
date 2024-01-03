package main

import (
	"fmt"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/audio"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/audio/query/v4/request"
	"log"
	"os"
)

// 音频检测结果查询 demo
func main() {
	YourBusinessId := os.Getenv("BUSINESS_ID")
	YourSecretId := os.Getenv("SECRET_ID")
	YourSecretKey := os.Getenv("SECRET_KEY")
	// 设置易盾内容安全分配的businessId
	taskRequest := request.NewAudioQueryTaskRequest(YourBusinessId)
	taskRequest.SetTaskIds([]string{"6gvq815vr3kh8nan82uweqwg0060a4ex", "xxxx"}) // 设置查询的taskId list
	// 实例化一个 Client，入参需要传入易盾内容安全分配的secretId，secretKey
	client := audio.NewAudioClientWithAccessKey(YourSecretId, YourSecretKey)

	response, err := client.Query(taskRequest)
	if err != nil {
		// 处理错误并打印日志
		log.Fatal(err)
	}

	if response.GetCode() == 200 {
		//results := response.Result
	} else {
		fmt.Println("error code: ", response.GetCode())
		fmt.Println("error msg: ", response.GetMsg())
	}
}
