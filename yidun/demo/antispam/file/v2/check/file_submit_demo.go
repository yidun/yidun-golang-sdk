package main

import (
	"fmt"
	"log"

	file "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/file/v2"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/file/v2/submit"
)

/**
 * 文档检测异步提交demo
 */
func main() {
	// 构造file submit 请求对象
	request := submit.NewFileSubmitV2Request()

	url := "http://nisptools.nos.netease.com/3cd824f0e64744a1b0f18f9698495c97.doc"
	request.SetURL(url)
	request.SetAccount("file_mock")
	request.SetFileName("fileName_mock")
	request.SetDataId("dataId_" + url)
	request.SetCallback("callback_" + url)

	// 实例化一个fileClient，入参需要传入易盾内容安全分配的secretId，secretKey
	fileCheckClient := file.NewFileClientWithAccessKey("YOUR_SECRET_ID", "YOUR_SECRET_KEY")

	response, err := fileCheckClient.Submit(request)
	if err != nil {
		// 处理错误并打印日志
		log.Fatal(err)
	}

	if response.GetCode() == 200 {
		// do something
		fmt.Println("Antispam taskId:", response.Result.TaskId)
	} else {
		fmt.Println("error code: ", response.GetCode())
		fmt.Println("error msg: ", response.GetMsg())
	}
}
