package main

import (
	"fmt"
	"log"

	crawler "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/crawler"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/crawler/v1/delete"
	"github.com/yidun/yidun-golang-sdk/yidun/core/http"
)

/**
 * 网站解决方案-微博检测任务删除
 */
func main() {
	// 构造WeiboDeleteV1Request请求对象
	request := delete.NewWeiboDeleteV1Request()

	request.SetJobIds([]int64{1223123, 1231233})
	// 设置协议为HTTP
	request.SetProtocol(http.ProtocolEnumHTTP)
	// 实例化一个crawlerClient，入参需要传入易盾内容安全分配的secretId，secretKey
	crawlerSubmitClient := crawler.NewCrawlerClientWithAccessKey("YOUR_SECRET_ID", "YOUR_SECRET_KEY")

	response, err := crawlerSubmitClient.WeiboDelete(request)
	if err != nil {
		// 处理错误并打印日志
		log.Fatal(err)
	}

	if response.GetCode() == 200 {
		// do something
		for _, v := range *response.Result.DeletedJobIds {
			fmt.Printf("deletedJobIds:%d ", v)
		}

	} else {
		fmt.Println("error code: ", response.GetCode())
		fmt.Println("error msg: ", response.GetMsg())
	}
}
