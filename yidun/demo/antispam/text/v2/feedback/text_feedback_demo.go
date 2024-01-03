package main

import (
	"encoding/json"
	"fmt"
	"log"

	v2 "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/text"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/text/v2/feedback"
)
// 文本结果反馈请求Demo
func main() {
	// 设置易盾内容安全分配的businessId
	request := feedback.NewTextCallBackRequest("YOUR_BUSINESS_ID")

	// 实例化一个 Client，入参需要传入易盾内容安全分配的secretId，secretKey
	textCheckClient := v2.NewTextClientWithAccessKey("YOUR_SECRET_ID", "YOUR_SECRET_KEY")

	request.SetFeedbacks(createFeedbacks())

	response, err := textCheckClient.Feedback(request)
	if err != nil {
		// 处理错误并打印日志
		log.Fatal(err)
	}

	if response.GetCode() == 200 {
		data := response.Result[0]
		fmt.Println("Antispam taskId:", *data.TaskId)
		fmt.Println("Antispam Result:", *data.Result)

	} else {
		fmt.Println("error code: ", response.GetCode())
		fmt.Println("error msg: ", response.GetMsg())
	}
}
// 构造反馈数据
func createFeedbacks() string {
	feedbackList := []map[string]interface{}{}
	feedbackMap := map[string]interface{}{
		"taskId": "TASK_ID",
		"level":  2,
		// 请求对象中的其他参数如果有需要，请参考官方接口文档中字段说明，按需添加
	}
	feedbackList = append(feedbackList, feedbackMap)

	feedbacks, _ := json.Marshal(feedbackList)
	return string(feedbacks)
}
