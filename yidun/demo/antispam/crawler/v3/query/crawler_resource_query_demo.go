package main

import (
	"fmt"
	"log"

	crawler "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/crawler"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/crawler/v3/query"
)

/**
 * 网站解决方案-网页单URL机器检测结果查询
 */
func main() {
	request := query.NewCrawlerResourceV3QueryRequest()

	request.SetTaskIdList([]string{"6ae19872ecc042b88b1c08e486d08bab", "a5ca9229b74040d1bbbf6ad48014d360"})

	// 实例化一个crawlerClient，入参需要传入易盾内容安全分配的secretId，secretKey
	crawlerClient := crawler.NewCrawlerClientWithAccessKey("YOUR_SECRET_ID", "YOUR_SECRET_KEY")

	response, err := crawlerClient.ResourceQuery(request)
	if err != nil {
		// 处理错误并打印日志
		log.Fatal(err)
	}

	if response.GetCode() == 200 {
		for _, result := range response.Result {
			// do something
			if result.Antispam != nil {
				fmt.Println("Antispam taskId:", *result.Antispam.TaskId)
			}
		}
	} else {
		fmt.Println("error code: ", response.GetCode())
		fmt.Println("error msg: ", response.GetMsg())
	}
}
