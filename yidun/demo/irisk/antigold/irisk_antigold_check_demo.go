package main

import (
	"log"

	"github.com/yidun/yidun-golang-sdk/yidun/core/http"
	"github.com/yidun/yidun-golang-sdk/yidun/service/irisk"

	"fmt"
)

func main() {
	request := irisk.NewIRiskAntiGoldCheckRequest("BUSINESS_ID")
	request.SetProtocol(http.ProtocolEnumHTTP)
	request.SetLogTime("2023-09-08T16:00:43+08:00")
	request.SetLogType("logType")
	request.SetAccount("objectId")
	request.SetNickname("nickname")
	request.SetRoleId("roleId")
	request.SetServerId("serverId")
	request.SetLogData("jsonstring")

	iriskClient := irisk.NewIRiskClientWithAccessKey("SECRET_ID", "SECRET_KEY")
	response, err := iriskClient.AntiGoldCheck(request)

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
