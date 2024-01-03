package main

import (
	"fmt"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/audio"
	request2 "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/audio/feedback/v1/request"
	"log"
	"os"
)

// 音频反馈demo
func main() {
	YourBusinessId := os.Getenv("BUSINESS_ID")
	YourSecretId := os.Getenv("SECRET_ID")
	YourSecretKey := os.Getenv("SECRET_KEY")
	// 设置易盾内容安全分配的businessId
	request := request2.NewAudioFeedbackRequest(YourBusinessId)

	// 实例化一个 Client，入参需要传入易盾内容安全分配的secretId，secretKey
	client := audio.NewAudioClientWithAccessKey(YourSecretId, YourSecretKey)

	feedbacks := make([]request2.AudioFeedback, 1)
	feedback := &request2.AudioFeedback{}
	feedback.SetLevel(2)
	feedback.SetTaskId("xxx")
	feedback.SetLabel(100)
	feedbacks[0] = *feedback
	request.SetFeedbacks(feedbacks)
	response, err := client.Feedback(request)
	if err != nil {
		// 处理错误并打印日志
		log.Fatal(err)
	}

	if response.GetCode() == 200 {
		for i := range response.Result {
			fmt.Println("taskId:", response.Result[i].TaskId)
			fmt.Println("result:", response.Result[i].Result)
		}
	} else {
		fmt.Println("error code: ", response.GetCode())
		fmt.Println("error msg: ", response.GetMsg())
	}
}
