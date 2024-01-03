package main

import (
	"log"

	"github.com/yidun/yidun-golang-sdk/yidun/service/irisk"

	"fmt"
)

func main() {
	request := irisk.NewIRiskMediaCheckRequest("BUSINESS_ID")
	request.SetMediaData("xxxxxxxxxxxx")
	request.SetMediaName("media01.bmp")
	request.SetIp("192.168.0.1")
	request.SetRoleId("role01")
	request.SetNickname("nickname01")
	request.SetServer("server01")

	iriskClient := irisk.NewIRiskClientWithAccessKey("SECRET_ID", "SECRET_KEY")
	response, err := iriskClient.GetMediaCheckResult(request)

	if err != nil {
		// 处理错误并打印日志
		log.Fatal(err)
	}

	if response.GetCode() == 200 {
		data := response.Data
		fmt.Println("TaskId:", data.TaskId)
	} else {
		fmt.Println("error code: ", response.GetCode())
		fmt.Println("error msg: ", response.GetMsg())
	}
}
