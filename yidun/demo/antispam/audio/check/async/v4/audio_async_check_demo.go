package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/audio"
	request2 "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/audio/check/async/v4/request"
)

// 音频异步检测demo
func main() {
	YourBusinessId := os.Getenv("BUSINESS_ID")
	YourSecretId := os.Getenv("SECRET_ID")
	YourSecretKey := os.Getenv("SECRET_KEY")
	// 设置易盾内容安全分配的businessId
	request := request2.NewAudioAsyncCheckRequest(YourBusinessId)
	request.SetExtLon1(99)
	request.SetExtLon2(88)
	request.SetExtStr1("extStr1")
	request.SetExtStr2("extStr2")

	// 实例化一个 Client，入参需要传入易盾内容安全分配的secretId，secretKey
	client := audio.NewAudioClientWithAccessKey(YourSecretId, YourSecretKey)

	// 设置 url
	request.SetURL("http://xxx.mp4?time=1700213386514")
	request.SetUniqueKey(time.Now().String())

	checkResponse, err := client.AsyncCheckAudio(request)
	if err != nil {
		// 处理错误并打印日志
		log.Fatal(err)
	}

	if checkResponse.GetCode() == 200 {
		data := checkResponse.Result
		fmt.Println("taskId:", *data.TaskId)
		fmt.Println("dealingCount", *data.DealingCount)
	} else {
		fmt.Println("error code: ", checkResponse.GetCode())
		fmt.Println("error msg: ", checkResponse.GetMsg())
	}
}
