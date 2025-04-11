package main

import (
	"fmt"
	"log"

	crawler "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/crawler"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/crawler/v3/submit"
)

/**
 * 网站解决方案网页单URL检测提交demo
 */
func main() {
	// 构造crawler resource submit请求对象
	request := submit.NewCrawlerResourceV3SubmitRequest()

	url := "https://news.163.com"
	request.SetUrl(url)
	request.SetCheckFlags([]int{1, 2})
	request.SetDataId("dataId_" + url)
	request.SetCallback("callback_" + url)
	request.SetCallbackUrl("http://xxxxxx.com/callback/receive")
	request.SetExtension("{\"mission\":123}")

	// 实例化一个crawlerClient，入参需要传入易盾内容安全分配的secretId，secretKey
	crawlerSubmitClient := crawler.NewCrawlerClientWithAccessKey("YOUR_SECRET_ID", "YOUR_SECRET_KEY")

	response, err := crawlerSubmitClient.ResourceSubmit(request)
	if err != nil {
		// 处理错误并打印日志
		log.Fatal(err)
	}

	if response.GetCode() == 200 {
		// do something
		fmt.Println("Antispam taskId:", *response.Result.TaskId)
	} else {
		fmt.Println("error code: ", response.GetCode())
		fmt.Println("error msg: ", response.GetMsg())
	}
}
