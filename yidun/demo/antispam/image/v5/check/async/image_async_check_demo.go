package main

import (
	"fmt"
	"log"

	image "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/image/v5"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/image/v5/check"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/image/v5/check/async"
)

/**
 * 图片异步检测demo
 */
func main() {
	// 设置易盾内容安全分配的businessId
	request := async.NewImageV5AsyncCheckRequest("YOUR_BUSINESS_ID")

	// 实例化一个imageClient，入参需要传入易盾内容安全分配的secretId，secretKey
	imageCheckClient := image.NewImageClientWithAccessKey("YOUR_SECRET_ID", "YOUR_SECRET_KEY")

	url := "http://xxxxxxx.com/xxx/xxx.jpg"
	image := check.NewImageBeanRequest()
	image.SetData(url)
	image.SetName("异步测主动回调1")
	// 设置图片数据的类型，1：图片URL，2:图片BASE64
	image.SetType(1)
	// 非必填，强烈建议通过轮询回调或者主动回调（2选1）的方式，来获取最新的检测结果。主动回调数据接口超时时间默认设置为2s，为了保证顺利接收数据，需保证接收接口性能稳定并且保证幂等性。
	image.SetCallbackUrl("主动回调地址")

	imageBeans := []check.ImageBeanRequest{*image}
	request.SetImages(imageBeans)

	response, err := imageCheckClient.ImageAsyncCheck(request)
	if err != nil {
		// 处理错误并打印日志
		log.Fatal(err)
	}

	if response.GetCode() == 200 {
		for _, result := range *response.Result.CheckImages {
			fmt.Println("Antispam taskId:", *result.TaskId)
		}
	} else {
		fmt.Println("error code: ", response.GetCode())
		fmt.Println("error msg: ", response.GetMsg())
	}
}
