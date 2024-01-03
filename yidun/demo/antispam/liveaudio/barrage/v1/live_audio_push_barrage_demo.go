package main

import (
	"fmt"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/liveaudio"
	request2 "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/liveaudio/barrage/v1/request"
	"log"
	"os"
)

func main() {
	// Create a AntispamRequester instance, the parameters need to pass in the secretId, secretKey distributed by Antispam.
	// 设置易盾内容安全分配的businessId
	YourBusinessId := os.Getenv("BUSINESS_ID")
	YourSecretId := os.Getenv("SECRET_ID")
	YourSecretKey := os.Getenv("SECRET_KEY")
	request := request2.NewLiveAudioBarrageV1Req(YourBusinessId)

	barrages := make([]*request2.LiveAudioBarrage, 1)
	barrage := &request2.LiveAudioBarrage{}
	barrage.SetContent("well done")
	barrage.SetTaskId("YourTaskId")
	barrage.SetSpamType(0)
	barrage.SetPublishTime(1701262095000)
	barrage.SetLevel(0)
	barrages[0] = barrage
	request.SetBarrages(barrages)

	// 实例化一个 Client，入参需要传入易盾内容安全分配的secretId，secretKey
	client := liveaudio.NewLiveAudioClientWithAccessKey(YourSecretId, YourSecretKey)

	response, err := client.PushBarrage(request)
	if err != nil {
		// 处理错误并打印日志
		log.Fatal(err)
	}

	if response.GetCode() == 200 {
		data := response.Result
		fmt.Println("Antispam taskId:", *data)

	} else {
		fmt.Println("error code: ", response.GetCode())
		fmt.Println("error msg: ", response.GetMsg())
	}

}
