package main

import (
	"fmt"
	"log"

	crawler "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/crawler"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/crawler/v1/submit"
	"github.com/yidun/yidun-golang-sdk/yidun/core/http"
)

/**
 * 网站解决方案-主站检测任务批量提交检测
 */
func main() {
	// 构造CrawlerJobBatchSubmitV1Request请求对象
	request := submit.NewCrawlerJobBatchSubmitV1Request()
	url := "https://news.163.com"
	CrawlerJobWebSite := submit.CrawlerJobWebSite{}
	CrawlerJobWebSite.SetSiteUrl(url)
	CrawlerJobWebSite1 := submit.CrawlerJobWebSite{}
	CrawlerJobWebSite1.SetSiteUrl(url)
	request.SetWebsites([]submit.CrawlerJobWebSite{CrawlerJobWebSite, CrawlerJobWebSite1})
	request.SetFrequency(86400000)
	request.SetLevel(-1)
	request.SetMaxResourceAmount(10)
	request.SetType(0)
	request.SetSliceStartTime(1709726123000)
	request.SetSliceEndTime(1709726523000)
	request.SetCheckStrategy(2)
	request.SetCheckFlags([]int{1, 2})
	// 设置协议为HTTP
	request.SetProtocol(http.ProtocolEnumHTTP)
	// 实例化一个crawlerClient，入参需要传入易盾内容安全分配的secretId，secretKey
	crawlerSubmitClient := crawler.NewCrawlerClientWithAccessKey("YOUR_SECRET_ID", "YOUR_SECRET_KEY")

	response, err := crawlerSubmitClient.CrawlerJobBatchSubmit(request)
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
