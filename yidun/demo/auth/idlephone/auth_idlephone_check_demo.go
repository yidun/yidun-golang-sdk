package main

import (
	"fmt"
	"log"

	auth "github.com/yidun/yidun-golang-sdk/yidun/service/auth"
)

// 实证认证
func main() {
	request := auth.NewIdlePhoneCheckRequest("BUSINESS_ID")
	request.SetEncryptType("1")
	request.SetPhoneList("12345678910")
	authClient := auth.NewAuthClientWithAccessKey("SECRET_ID", "SECRET_KEY")
	response, err := authClient.IdlePhoneCheck(request)

	if err != nil {
		// 处理错误并打印日志
		log.Fatal(err)
	}

	if response.GetCode() == 200 {
		result := response.Result
		//打印返回结果
		fmt.Println("Status:", *result.Status)
		fmt.Println("IsPayed:", *result.IsPayed)
		fmt.Println("RequestId:", *result.RequestId)

	} else {
		fmt.Println("error code: ", response.GetCode())
		fmt.Println("error msg: ", response.GetMsg())
	}
}
