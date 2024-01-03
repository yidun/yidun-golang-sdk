package main

import (
	"fmt"
	"log"

	auth "github.com/yidun/yidun-golang-sdk/yidun/service/auth"
)

// 身份证检测
func main() {
	request := auth.NewIdCardCheckRequest("BUSINESS_ID")
	request.SetName("张三")
	request.SetCardNo("12345678910")
	request.SetDataId("123456")

	authClient := auth.NewAuthClientWithAccessKey("SECRET_ID", "SECRET_KEY")
	response, err := authClient.IdCardCheck(request)

	if err != nil {
		// 处理错误并打印日志
		log.Fatal(err)
	}

	if response.GetCode() == 200 {
		result := response.Result
		//打印返回结果
		fmt.Println("Status:", *result.Status)
		fmt.Println("TaskId:", *result.TaskId)
		fmt.Println("IsPayed:", *result.IsPayed)
		fmt.Println("ReasonType:", *result.ReasonType)

	} else {
		fmt.Println("error code: ", response.GetCode())
		fmt.Println("error msg: ", response.GetMsg())
	}
}
