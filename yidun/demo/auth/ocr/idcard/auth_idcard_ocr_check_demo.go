package main

import (
	"fmt"
	"log"

	auth "github.com/yidun/yidun-golang-sdk/yidun/service/auth"
)

// 身份证OCR识别
func main() {
	request := auth.NewIdCardOcrCheckRequest("BUSINESS_ID")
	request.SetFrontPicture("http://xx/yidun/2-0-0-0-1.jpg")
	request.SetBackPicture("http://xx/yidun/2-0-0-0-2.jpg")
	request.SetDataId("123456")
	request.SetPicType("1")
	authClient := auth.NewAuthClientWithAccessKey("SECRET_ID", "SECRET_KEY")
	response, err := authClient.IdCardOcrCheck(request)

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
