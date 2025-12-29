package main

import (
	"encoding/json"
	"fmt"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/liveaudio"
	request2 "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/liveaudio/feedback/v1/request"
	"log"
	"os"
	"github.com/yidun/yidun-golang-sdk/yidun/core/http"
)

func main() {
	// Create a AntispamRequester instance, the parameters need to pass in the secretId, secretKey distributed by Antispam.
	// 设置易盾内容安全分配的businessId
	YourBusinessId := os.Getenv("BUSINESS_ID")
	YourSecretId := os.Getenv("SECRET_ID")
	YourSecretKey := os.Getenv("SECRET_KEY")

	request := request2.NewLiveAudioFeedbackV1Req(YourBusinessId)
	// 设置查询开始时间和结束时间
	feedbacks := make([]*request2.LiveAudioFeedback, 1)
	feedback := &request2.LiveAudioFeedback{}
	feedback.SetTaskID("YourTaskId")
	feedback.SetViewCount(100)
	feedbacks[0] = feedback
	request.SetFeedbacks(feedbacks)
	// 设置协议为HTTP
	request.SetProtocol(http.ProtocolEnumHTTP)
	// 实例化一个 Client，入参需要传入易盾内容安全分配的secretId，secretKey
	client := liveaudio.NewLiveAudioClientWithAccessKey(YourSecretId, YourSecretKey)

	response, err := client.Feedback(request)
	if err != nil {
		// 处理错误并打印日志
		log.Fatal(err)
	}

	if response.GetCode() == 200 {
		data := response.Result
		jsonBytes, err := json.Marshal(data)
		if err != nil {
			fmt.Println("转换为JSON时出错:", err)
			return
		}

		jsonStr := string(jsonBytes)
		fmt.Println(jsonStr)

	} else {
		fmt.Println("error code: ", response.GetCode())
		fmt.Println("error msg: ", response.GetMsg())
	}
}
