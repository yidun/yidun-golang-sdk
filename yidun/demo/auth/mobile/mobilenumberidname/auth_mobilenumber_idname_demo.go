package main

import (
	"fmt"
	"log"

	auth "github.com/yidun/yidun-golang-sdk/yidun/service/auth"
)

// 手机号与所有者身份证号及姓名验证请求
func main() {
	request := auth.NewMobileNumberOwnerIdNameCheckRequest("BUSINESS_ID")
	request.SetName("张三")
	request.SetPhone("12345678910")
	request.SetCardNo("12345678910")

	authClient := auth.NewAuthClientWithAccessKey("SECRET_ID", "SECRET_KEY")
	response, err := authClient.MobileNumberOwnerIdNameCheck(request)

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
