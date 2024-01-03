package main

import (
	"github.com/yidun/yidun-golang-sdk/yidun/service/irisk"
	"log"

	"fmt"
)

func main() {
	request := irisk.NewIRiskDisposeUploadRequest("BUSINESS_ID")
	request.SetHandleObjectType(1)
	request.SetAction(1)
	request.SetObjectId("objectId")
	request.SetRoleId("roleId")
	request.SetActionDesc("desc")
	request.SetReason("reason")

	iriskClient := irisk.NewIRiskClientWithAccessKey("SECRET_ID", "SECRET_KEY")
	response, err := iriskClient.UploadDisposeInfo(request)

	if err != nil {
		// 处理错误并打印日志
		log.Fatal(err)
	}
	if response.GetCode() == 200 {
		// 上报成功
	} else {
		fmt.Println("error code: ", response.GetCode())
		fmt.Println("error msg: ", response.GetMsg())
	}
}
