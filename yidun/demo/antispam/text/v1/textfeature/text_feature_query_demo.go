package main

import (
	"fmt"
	"log"

	v1 "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/text/v1"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/text/v1/textfeature"
	"github.com/yidun/yidun-golang-sdk/yidun/core/http"
)

// 文本特征查询 Demo
func main() {
	// 设置易盾内容安全分配的businessId
	request := textfeature.NewTextFeatureQueryRequest("YOUR_BUSINESS_ID")

	// 设置查询参数
	request.SetPageNum(1)
	request.SetPageSize(10)
	request.SetScope(1)
	// 设置协议为HTTP
	request.SetProtocol(http.ProtocolEnumHTTP)
	// 实例化一个 Client，入参需要传入易盾内容安全分配的secretId，secretKey
	textClient := v1.NewTextClientWithAccessKey("YOUR_SECRET_ID", "YOUR_SECRET_KEY")

	response, err := textClient.TextFeatureQuery(request)
	if err != nil {
		log.Fatal(err)
	}

	if response.GetCode() == 200 {
		fmt.Println("文本特征查询成功！")
		fmt.Printf("查询结果: %+v\n", response.Result)
	} else {
		fmt.Println("error code: ", response.GetCode())
		fmt.Println("error msg: ", response.GetMsg())
	}
}
