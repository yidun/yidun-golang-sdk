package main

import (
	"fmt"
	"log"

	crawler "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/crawler"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/crawler/v1/submit"
	"github.com/yidun/yidun-golang-sdk/yidun/core/http"
)

/**
 * 网站解决方案-微博检测任务批量提交检测
 */
func main() {
	// 构造WeiboBatchSubmitV1Request请求对象
	request := submit.NewWeiboBatchSubmitV1Request()
	request.SetType(6)
	request.SetStartTime(1709636123000)
	request.SetEndTime(1709646123000)
	request.SetStrategy(3)
	request.SetMaxCheckCount(10)
	request.SetCheckFlags([]int{1, 2})
	weiboConfig := submit.WeiboConfig{}
	weiboConfig.SetUrl("https://weibo.com/u/xxxx")
	weiboConfig1 := submit.WeiboConfig{}
	weiboConfig1.SetUrl("https://weibo.com/u/xxxxx")
	request.SetWeiboBloggers([]submit.WeiboConfig{weiboConfig, weiboConfig1})
	// 设置协议为HTTP
	request.SetProtocol(http.ProtocolEnumHTTP)
	// 实例化一个crawlerClient，入参需要传入易盾内容安全分配的secretId，secretKey
	crawlerSubmitClient := crawler.NewCrawlerClientWithAccessKey("YOUR_SECRET_ID", "YOUR_SECRET_KEY")

	response, err := crawlerSubmitClient.WeiboBatchSubmit(request)
	if err != nil {
		// 处理错误并打印日志
		log.Fatal(err)
	}

	if response.GetCode() == 200 {
		// do something
		for _, v := range *response.Result {
			fmt.Printf("crawler solution jobId:%d ", *v.JobId)
		}
	} else {
		fmt.Println("error code: ", response.GetCode())
		fmt.Println("error msg: ", response.GetMsg())
	}
}
