package main

import (
	"fmt"
	"log"

	file "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/file/v2"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/file/v2/callback"
	"github.com/yidun/yidun-golang-sdk/yidun/core/http"
)

/**
 * 文档检测回调（轮询模式）demo
 */
func main() {
	// 构造file callback 请求对象
	request := callback.NewFileCallBackRequest()
	// 设置协议为HTTP
	request.SetProtocol(http.ProtocolEnumHTTP)
	// 实例化一个fileClient，入参需要传入易盾内容安全分配的secretId，secretKey
	fileCallBackClient := file.NewFileClientWithAccessKey("YOUR_SECRET_ID", "YOUR_SECRET_KEY")

	response, err := fileCallBackClient.Callback(request)
	if err != nil {
		// 处理错误并打印日志
		log.Fatal(err)
	}

	if response.GetCode() == 200 {
		// do something
		fmt.Println("Antispam:", response.Result)
	} else {
		fmt.Println("error code: ", response.GetCode())
		fmt.Println("error msg: ", response.GetMsg())
	}
}
