package main

import (
	"fmt"
	"log"

	crawler "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/crawler"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/crawler/v1/delete"
)

/**
 * 网站解决方案-公众号检测任务删除
 */
func main() {
	// 构造OfficialAccountsDeleteV1Request请求对象
	request := delete.NewOfficialAccountsDeleteV1Request()

	request.SetJobIds([]int64{1761622595622943, 1761622595622942})

	// 实例化一个crawlerClient，入参需要传入易盾内容安全分配的secretId，secretKey
	crawlerSubmitClient := crawler.NewCrawlerClientWithAccessKey("YOUR_SECRET_ID", "YOUR_SECRET_KEY")

	response, err := crawlerSubmitClient.OfficialAccountsDelete(request)
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
