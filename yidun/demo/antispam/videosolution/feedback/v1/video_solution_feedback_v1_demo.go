package main

import (
	"encoding/json"
	"fmt"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/videosolution"
	request3 "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/videosolution/feedback/v1/request"
	"log"
	"os"
	"github.com/yidun/yidun-golang-sdk/yidun/core/http"
)

// 点播音视频解决方案反馈demo
func main() {
	YourSecretId := os.Getenv("VS_SECRET_ID")
	YourSecretKey := os.Getenv("VS_SECRET_KEY")
	request := request3.NewVideoSolutionFeedbackRequest()

	// 实例化一个 Client，入参需要传入易盾内容安全分配的secretId，secretKey
	client := videosolution.NewVideoSolutionClientWithAccessKey(YourSecretId, YourSecretKey)

	feedbacks := make([]request3.VideoSolutionFeedback, 1)
	feedback := &request3.VideoSolutionFeedback{}
	feedback.SetLevel(2)
	feedback.SetTaskId("ralgd03gy16qe9xfy82uut0h0200a4ji")
	feedback.SetLabel(100)
	feedback.SetTags([]string{"test"})
	feedbacks[0] = *feedback
	request.SetFeedbacks(feedbacks)
	// 设置协议为HTTP
	request.SetProtocol(http.ProtocolEnumHTTP)
	response, err := client.Feedback(request)
	if err != nil {
		// 处理错误并打印日志
		log.Fatal(err)
	}

	if response.GetCode() == 200 {
		result := response.Result
		resultJson, _ := json.Marshal(result)
		fmt.Println("feedback result: ", string(resultJson))
	} else {
		fmt.Println("error code: ", response.GetCode())
		fmt.Println("error msg: ", response.GetMsg())
	}
}
