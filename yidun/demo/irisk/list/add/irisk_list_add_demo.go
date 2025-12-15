package main

import (
	"log"

	"github.com/yidun/yidun-golang-sdk/yidun/core/http"
	"github.com/yidun/yidun-golang-sdk/yidun/service/irisk"

	"fmt"
)

func main() {
	request := irisk.NewIRiskListAddRequest("BUSINESS_ID")
	request.SetProtocol(http.ProtocolEnumHTTP)
	request.SetListGroupCode("ayp76c2dmh2k0ktd8jyia2cg22009v02")
	request.SetContent("content")
	request.SetDescription("description")
	request.SetExpireTime(1809275212935)

	iriskClient := irisk.NewIRiskClientWithAccessKey("SECRET_ID", "SECRET_KEY")
	response, err := iriskClient.ListAdd(request)

	if err != nil {
		// 处理错误并打印日志
		log.Fatal(err)
	}
	if response.GetCode() == 200 {
		// 添加成功
		fmt.Println("ok msg: ", response.GetMsg())
	} else {
		fmt.Println("error code: ", response.GetCode())
		fmt.Println("error msg: ", response.GetMsg())

		if irisk.LIST_LIMIT_EXCEEDED == response.GetCode() {
			// 名单量超出限制
			fmt.Println("error msg: 名单量超出限制")
		} else if irisk.OPERATION_FAILED == response.GetCode() {
			// 处理失败
			fmt.Println("error msg: 处理失败")
		} else if irisk.TIMEOUT_EXCEPTION == response.GetCode() {
			// 请求处理超时
			fmt.Println("error msg: 请求处理超时")
		}
		// else if 其它情况 ......
	}
}
