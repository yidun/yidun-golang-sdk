package main

import (
	"fmt"
	"log"

	image "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/image/v5"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/image/v5/callback"
)

/**
 * 图片检测结果获取（轮询模式）demo
 */
func main() {
	// 设置易盾内容安全分配的businessId
	request := callback.NewImageCallbackRequest("YOUR_BUSINESS_ID")

	// 实例化一个imageClient，入参需要传入易盾内容安全分配的secretId，secretKey
	imageCheckClient := image.NewImageClientWithAccessKey("YOUR_SECRET_ID", "YOUR_SECRET_KEY")

	response, err := imageCheckClient.ImageCallback(request)
	if err != nil {
		// 处理错误并打印日志
		log.Fatal(err)
	}

	if response.GetCode() == 200 {
		for _, result := range *response.Result {
			fmt.Printf("Antispam taskId:%s \n", *result.Antispam.TaskId)
		}
	} else {
		fmt.Println("error code: ", response.GetCode())
		fmt.Println("error msg: ", response.GetMsg())
	}
}
