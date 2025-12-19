package main

import (
	"fmt"
	"log"

	"github.com/yidun/yidun-golang-sdk/yidun/core/http"
	auth "github.com/yidun/yidun-golang-sdk/yidun/service/auth"
)

// 手机号与所有者身份证号及姓名验证请求
func main() {
	request := auth.NewMobileNumberOwnerNameCheckRequest("BUSINESS_ID")
	request.SetName("张三")
	request.SetPhone("12345678910")
	// 设置协议为HTTP
	request.SetProtocol(http.ProtocolEnumHTTP)
	authClient := auth.NewAuthClientWithAccessKey("SECRET_ID", "SECRET_KEY")
	response, err := authClient.MobileNumberOwnerNameCheck(request)

	if err != nil {
		// 处理错误并打印日志
		log.Fatal(err)
	}

	if response.GetCode() == 200 {
		result := response.Result
		fmt.Println("result:", result)
	} else {
		fmt.Println("error code: ", response.GetCode())
		fmt.Println("error msg: ", response.GetMsg())
	}

}
