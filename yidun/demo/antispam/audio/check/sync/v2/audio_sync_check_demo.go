package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/audio"
	request3 "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/audio/check/sync/v2/request"
)

// 音频同步检测demo
func main() {
	// 设置易盾内容安全分配的businessId
	YourBusinessId := os.Getenv("BUSINESS_ID")
	YourSecretId := os.Getenv("SECRET_ID")
	YourSecretKey := os.Getenv("SECRET_KEY")
	request := request3.NewAudioSyncCheckRequest(YourBusinessId)
	request.SetExtLon1(1)
	request.SetExtLon2(2)
	request.SetExtStr1("extStr1")
	request.SetExtStr2("extStr2")

	// 实例化一个Client，入参需要传入易盾内容安全分配的secretId，secretKey
	client := audio.NewAudioClientWithAccessKey(YourSecretId, YourSecretKey)

	// 设置 url
	request.SetURL("http://xxx.mp4?time=1700213386514")
	request.SetUniqueKey(time.Now().String())

	syncCheckResponse, err := client.SyncCheckAudio(request)
	if err != nil {
		// 处理错误并打印日志
		log.Fatal(err)
	}

	if syncCheckResponse.GetCode() == 200 {
		data := syncCheckResponse.Result.Antispam
		fmt.Println("Antispam taskId:", *data.TaskId)
	} else {
		fmt.Println("error code: ", syncCheckResponse.GetCode())
		fmt.Println("error msg: ", syncCheckResponse.GetMsg())
	}
}
