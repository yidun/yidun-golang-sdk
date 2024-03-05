package main

import (
	"fmt"
	"log"

	crawler "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/crawler"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/crawler/v1/submit"
)

/**
 * 网站解决方案-微博检测任务提交检测
 */
func main() {
	// 构造WeiboSubmitV1Request请求对象
	request := submit.NewWeiboSubmitV1Request()
	request.SetType(6)
	request.SetStartTime(1709604245000)
	request.SetEndTime(1709604335000)
	request.SetStrategy(3)
	request.SetMaxCheckCount(10)
	request.SetCheckFlags([]int{1, 2})
	request.SetUrl("https://weibo.com/u/xxx")

	// 实例化一个crawlerClient，入参需要传入易盾内容安全分配的secretId，secretKey
	crawlerSubmitClient := crawler.NewCrawlerClientWithAccessKey("YOUR_SECRET_ID", "YOUR_SECRET_KEY")

	response, err := crawlerSubmitClient.WeiboSubmit(request)
	if err != nil {
		// 处理错误并打印日志
		log.Fatal(err)
	}

	if response.GetCode() == 200 {
		// do something
		fmt.Printf("crawler solution jobId:%d", *response.Result.JobId)
	} else {
		fmt.Println("error code: ", response.GetCode())
		fmt.Println("error msg: ", response.GetMsg())
	}
}
