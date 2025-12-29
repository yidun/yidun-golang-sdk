package main

import (
	"fmt"
	"log"

	image "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/image/v5"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/image/v5/feedback"
	"github.com/yidun/yidun-golang-sdk/yidun/core/http"
)

/**
 * 图片结果反馈demo
 */
func main() {
	// 设置易盾内容安全分配的businessId
	request := feedback.NewImageFeedbackRequest("YOUR_BUSINESS_ID")

	// 实例化一个igeClient，入参需要传入易盾内容安全分配的secretId，secretKey
	imageCheckClient := image.NewImageClientWithAccessKey("YOUR_SECRET_ID", "YOUR_SECRET_KEY")

	feedback_1 := feedback.ImageFeedbackBeanRequest{}
	feedback_1.SetTaskId("k2nkobiee8rv3xsaqj812v4h0030a4m3")
	feedback_1.SetLevel(2)
	feedback_1.SetLabel(100)

	request.SetFeedbacks([]feedback.ImageFeedbackBeanRequest{feedback_1})
	// 设置协议为HTTP
	request.SetProtocol(http.ProtocolEnumHTTP)
	response, err := imageCheckClient.ImageFeedback(request)
	if err != nil {
		// 处理错误并打印日志
		log.Fatal(err)
	}

	if response.GetCode() == 200 {
		for _, result := range *response.Result {
			fmt.Printf("Antispam taskId:%s result:%d \n", *result.TaskId, *result.Result)
		}
	} else {
		fmt.Println("error code: ", response.GetCode())
		fmt.Println("error msg: ", response.GetMsg())
	}
}
