package main

import (
	"fmt"
	"log"

	image "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/image/v5"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/image/v5/query"
)

/**
 * 图片结果获取demo  V2接口
 */
func main() {
	// 设置易盾内容安全分配的businessId
	request := query.NewImageQueryV2Request("YOUR_BUSINESS_ID")

	// 实例化一个imageClient，入参需要传入易盾内容安全分配的secretId，secretKey
	imageClient := image.NewImageClientWithAccessKey("YOUR_SECRET_ID", "YOUR_SECRET_KEY")

	request.SetTaskIds([]string{"fu3t9fzh7zh5a9500xwicr0h0030aape", "t16d84y0fqcbjc4rbpztc60h0030aaph"})
	response, err := imageClient.ImageQueryV2(request)
	if err != nil {
		// 处理错误并打印日志
		log.Fatal(err)
	}

	if response.GetCode() == 200 {
		for _, result := range *response.Result {
			// 根据需要获取每张图片的结果，具体返回字段的说明，请参考官方接口文档中字段说明
			fmt.Printf("Antispam taskId:%s \n", *result.Antispam.TaskId)
		}
	} else {
		fmt.Println("error code: ", response.GetCode())
		fmt.Println("error msg: ", response.GetMsg())
	}
}
