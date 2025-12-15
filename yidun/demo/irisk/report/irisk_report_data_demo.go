package main

import (
	"log"
	"time"

	"github.com/yidun/yidun-golang-sdk/yidun/core/http"
	"github.com/yidun/yidun-golang-sdk/yidun/service/irisk"

	"fmt"
)

func main() {
	request := irisk.NewIRiskReportDataRequest("BUSINESS_ID")
	request.SetProtocol(http.ProtocolEnumHTTP)
	request.SetReportChannel("reportChannel")
	request.SetReportTime(time.Now().UnixNano() / 1e6)
	request.SetWhistleblower("{\"account\":\"账号ID\",\"roleId\":\"角色ID\",\"roleName\":\"角色名称\",\"serverId\":\"服务器ID/名称\",\"level\":\"等级\",\"recharge\":10000000}")
	request.SetReportedPerson("{\"account\":\"账号ID\",\"roleId\":\"角色ID\",\"roleName\":\"角色名称\",\"serverId\":\"服务器ID/名称\",\"level\":\"等级\",\"recharge\":10000000}")
	request.SetReportType("reportType")
	request.SetReportScene("reportScene")
	request.SetReportData("reportData")

	iriskClient := irisk.NewIRiskClientWithAccessKey("SECRET_ID", "SECRET_KEY")
	response, err := iriskClient.UploadReportData(request)

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
