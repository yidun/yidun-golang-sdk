package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/audio"
	request3 "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/audio/check/sync/v2/request"
	"github.com/yidun/yidun-golang-sdk/yidun/core/http"
)

// 执行音频同步检测
func checkAudio(client *audio.AudioClient, businessId string, setupRequest func(*request3.AudioSyncCheckRequest)) {
	request := request3.NewAudioSyncCheckRequest(businessId)
	// 设置扩展字段
	request.SetExtLon1(1)
	request.SetExtLon2(2)
	request.SetExtStr1("extStr1")
	request.SetExtStr2("extStr2")
	// 设置协议为HTTP
	request.SetProtocol(http.ProtocolEnumHTTP)
	request.SetUniqueKey(time.Now().String())

	// 调用设置函数来配置特定的检测方式
	setupRequest(request)

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

// 音频同步检测demo
func main() {
	// 设置易盾内容安全分配的businessId
	YourBusinessId := os.Getenv("BUSINESS_ID")
	YourSecretId := os.Getenv("SECRET_ID")
	YourSecretKey := os.Getenv("SECRET_KEY")

	// 实例化一个Client，入参需要传入易盾内容安全分配的secretId，secretKey
	client := audio.NewAudioClientWithAccessKey(YourSecretId, YourSecretKey)

	// URL方式检测示例
	fmt.Println("=== URL方式检测 ===")
	checkAudio(client, YourBusinessId, func(request *request3.AudioSyncCheckRequest) {
		// 设置 url
		request.SetURL("http://xxx.mp4?time=1700213386514")
	})

	// BASE64方式检测示例
	fmt.Println("\n=== BASE64方式检测 ===")
	checkAudio(client, YourBusinessId, func(request *request3.AudioSyncCheckRequest) {
		// 设置检测类型为BASE64
		request.SetDataCheckType(int(request3.DataCheckTypeBASE64))
		// 设置base64编码的音频内容
		request.SetData("base64编码的音频内容...")
	})
}
