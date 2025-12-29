package main

import (
	"fmt"
	"log"

	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/label"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/label/request"
	"github.com/yidun/yidun-golang-sdk/yidun/core/http"
)

/**
 * 反垃圾标签查询demo
 */
 func main() {
	// 创建一个LabelQueryRequest实例
	request := request.NewLabelQueryRequest();

	// 实例化Client，入参需要传入易盾内容安全分配的AccessKeyId，AccessKeySecret
	labelClient := label.NewLabelClientWithAccessKey("AccessKeyId", "AccessKeySecret")

	// 传入请求参数
	//设置返回标签的最大层级
	request.SetMaxDepth(2)
	//指定业务类型
	// request.SetBusinessTypes(&[]string{"1", "2"})
	//制定业务
	// request.SetBusinessID("YOUR_BUSINESS_ID")
	// request.SetClientID("YOUR_CLIENT_ID")
	// 设置协议为HTTP
	request.SetProtocol(http.ProtocolEnumHTTP)
	response, err := labelClient.QueryLabel(request)
	if err != nil {
		// 处理错误并打印日志
		log.Fatal(err)
	}

	if response.GetCode() == 200 {
		data := response.Data
		fmt.Println("Antispam label size:", len(data))
		fmt.Println("Antispam label Name:", *data[0].Name)
		fmt.Println("Antispam label code:", *data[0].Label)
		fmt.Println("Antispam label subLabel:", *data[0].SubLabels[0].Name)
		fmt.Println("Antispam label subLabel code:", *data[0].SubLabels[0].Code)

	} else {
		fmt.Println("error code: ", response.GetCode())
		fmt.Println("error msg: ", response.GetMsg())
	}
}
