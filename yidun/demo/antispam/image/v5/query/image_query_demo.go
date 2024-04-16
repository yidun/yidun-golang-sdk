package main

import (
	"fmt"
	"log"

	image "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/image/v5"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/image/v5/query"
)

/**
 * 图片结果反馈demo
 */
func main() {
	// 设置易盾内容安全分配的businessId
	request := query.NewImageQueryRequest("YOUR_BUSINESS_ID")

	// 实例化一个imageClient，入参需要传入易盾内容安全分配的secretId，secretKey
	imageClient := image.NewImageClientWithAccessKey("YOUR_SECRET_ID", "YOUR_SECRET_KEY")

	request.SetTaskIds([]string{"qaspjsiywt8j12ztmat3ue0h0030a771", "dvr5fzjgq7eiu6n5zjk8m60h0030a771"})
	response, err := imageClient.ImageQuery(request)
	if err != nil {
		// 处理错误并打印日志
		log.Fatal(err)
	}

	if response.GetCode() == 200 {
		for _, result := range *response.Result {
			fmt.Printf("Antispam taskId:%s \n", *result.TaskId)
		}
	} else {
		fmt.Println("error code: ", response.GetCode())
		fmt.Println("error msg: ", response.GetMsg())
	}
}
