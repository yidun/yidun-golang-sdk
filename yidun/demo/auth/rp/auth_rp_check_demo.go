package main

import (
	"fmt"
	"log"

	"github.com/yidun/yidun-golang-sdk/yidun/core/http"
	auth "github.com/yidun/yidun-golang-sdk/yidun/service/auth"
)

// 实人认证
func main() {
	request := auth.NewRpCheckRequest("BUSINESS_ID")
	request.SetDataId("123456")
	request.SetPicType("1")
	request.SetAvatar("1")
	request.SetName("张三")
	request.SetCardNo("123456789")
	// 设置协议为HTTP
	request.SetProtocol(http.ProtocolEnumHTTP)
	authClient := auth.NewAuthClientWithAccessKey("SECRET_ID", "SECRET_KEY")
	response, err := authClient.RpCheck(request)

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
