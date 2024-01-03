package main

import (
	"log"

	mobileverify "github.com/yidun/yidun-golang-sdk/yidun/service/mobileverify"

	"fmt"
)

// 手机号验证请求
func main() {
	re := mobileverify.NewMobileNumberVerifyRequest("BUSINESS_ID")
	re.SetToken("ea37f48498b34d8983b10f18f8f8538e").SetAccessToken("d89f72f9a3b2d7143f3985ebceb62754d6dc089f1820a1716297892795068419").SetPhone("13888888888")
	mobileNumberVerifyClient := mobileverify.NewMobileNumberClientWithAccessKey("SECRET_ID", "SECRET_KEY")
	response, err := mobileNumberVerifyClient.VerifyMobileNumber(re)

	if err != nil {
		// 处理错误并打印日志
		log.Fatal(err)
	}

	if response.GetCode() == 200 {
		data := response.Data
		result := data.Result
		// configData即为需要返回给客户端的数据
		fmt.Println("result: ", result)
	} else {
		fmt.Println("error code: ", response.GetCode())
		fmt.Println("error msg: ", response.GetMsg())
	}
}
