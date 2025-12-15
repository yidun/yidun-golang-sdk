package main

import (
	"log"

	"fmt"

	"github.com/yidun/yidun-golang-sdk/yidun/core/http"
	"github.com/yidun/yidun-golang-sdk/yidun/service/irisk"
)

func main() {
	re := irisk.NewIRiskConfigRequest("BUSINESS_ID")
	re.SetProtocol(http.ProtocolEnumHTTP)
	re.SetIp("192.168.0.1")
	re.SetSdkData("xxxxxxxxxxxx")
	iriskClient := irisk.NewIRiskClientWithAccessKey("SECRET_ID", "SECRET_KEY")
	response, err := iriskClient.GetConfig(re)

	if err != nil {
		// 处理错误并打印日志
		log.Fatal(err)
	}

	if response.GetCode() == 200 {
		data := response.Data
		configData := data.ConfigData
		// configData即为需要返回给客户端的数据
		fmt.Println(configData)
	} else {
		fmt.Println("error code: ", response.GetCode())
		fmt.Println("error msg: ", response.GetMsg())
	}
}
