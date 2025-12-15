package main

import (
	"log"

	"github.com/yidun/yidun-golang-sdk/yidun/core/http"
	"github.com/yidun/yidun-golang-sdk/yidun/service/irisk"

	"fmt"
)

func main() {
	request := irisk.NewIRiskMediaQueryRequest("BUSINESS_ID")
	request.SetProtocol(http.ProtocolEnumHTTP)
	request.SetTaskId("xxewxylcm62kh1a7yra3hxdgdd00a7p4")

	iriskClient := irisk.NewIRiskClientWithAccessKey("SECRET_ID", "SECRET_KEY")
	response, err := iriskClient.GetMediaQueryResult(request)

	if err != nil {
		// 处理错误并打印日志
		log.Fatal(err)
	}

	if response.GetCode() == 200 {
		data := response.Data
		fmt.Println("ReceiveTime:", data.ReceiveTime)
		fmt.Println("Ip:", data.Ip)
		fmt.Println("RoleId:", data.RoleId)
		fmt.Println("Nickname:", data.Nickname)
		fmt.Println("Server:", data.Server)
		fmt.Println("Status:", data.Status)
		fmt.Println("TagNameList:", data.TagNameList)
		fmt.Println("Reason:", data.Reason)
	} else {
		fmt.Println("error code: ", response.GetCode())
		fmt.Println("error msg: ", response.GetMsg())
	}
}
