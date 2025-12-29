package main

import (
	"fmt"
	"log"

	crawler "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/crawler"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/crawler/v1/query"
	"github.com/yidun/yidun-golang-sdk/yidun/core/http"
)

/**
 * 网站解决方案-网站&公众号任务异常检测结果分页查询接口
 */
func main() {
	request := query.NewCrawlerJobCallbackQueryRequest()

	request.SetPageNum(0)
	request.SetPageSize(20)
	request.SetJobId(1502158478754329)
	// 设置协议为HTTP
	request.SetProtocol(http.ProtocolEnumHTTP)
	// 实例化一个crawlerClient，入参需要传入易盾内容安全分配的secretId，secretKey
	crawlerClient := crawler.NewCrawlerClientWithAccessKey("YOUR_SECRET_ID", "YOUR_SECRET_KEY")

	response, err := crawlerClient.JobQuery(request)
	if err != nil {
		// 处理错误并打印日志
		log.Fatal(err)
	}

	if response.GetCode() == 200 {
		if *response.Result.Rows != nil {
			for _, result := range *response.Result.Rows {
				// do something
				if result.Antispam != nil {
					fmt.Println("Antispam taskId:", *result.Antispam.TaskId)
				}
			}
		}
	} else {
		fmt.Println("error code: ", response.GetCode())
		fmt.Println("error msg: ", response.GetMsg())
	}
}
