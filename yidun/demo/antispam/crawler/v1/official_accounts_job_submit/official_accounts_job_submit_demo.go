package main

import (
	"fmt"
	"log"

	crawler "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/crawler"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/crawler/v1/submit"
	"github.com/yidun/yidun-golang-sdk/yidun/core/http"
)

/**
 * 网站解决方案-公众号检测任务提交检测
 */
func main() {
	// 构造OfficialAccountsSubmitV1Request请求对象
	request := submit.NewOfficialAccountsSubmitV1Request()
	request.SetType(4)
	request.SetStartTime(1709604245000)
	request.SetEndTime(1709604335000)
	request.SetStrategy(3)
	request.SetMaxCheckCount(10)
	request.SetCheckFlags([]int{1, 2})
	request.SetWechatAccount("zhejiangtsw")
	// 设置协议为HTTP
	request.SetProtocol(http.ProtocolEnumHTTP)
	// 实例化一个crawlerClient，入参需要传入易盾内容安全分配的secretId，secretKey
	crawlerSubmitClient := crawler.NewCrawlerClientWithAccessKey("YOUR_SECRET_ID", "YOUR_SECRET_KEY")

	response, err := crawlerSubmitClient.OfficialAccountsSubmit(request)
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
