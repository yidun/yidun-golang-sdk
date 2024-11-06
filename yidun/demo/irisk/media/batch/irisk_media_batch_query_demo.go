package main

import (
	"log"

	"github.com/yidun/yidun-golang-sdk/yidun/service/irisk"

	"fmt"
)

func main() {
	request := irisk.NewIRiskMediaBatchQueryRequest("BUSINESS_ID")
	request.SetTaskIds("[\"tpavc9q43b0ihy5ekc5g2axgdd00aaxs\",\"1ggpjkqx07qoopc35w8b2ybgdd00aaxs\"]")

	iriskClient := irisk.NewIRiskClientWithAccessKey("SECRET_ID", "SECRET_KEY")
	response, err := iriskClient.GetMediaBatchQueryResult(request)

	if err != nil {
		// 处理错误并打印日志
		log.Fatal(err)
	}

	if response.GetCode() == 200 {
		data := response.Data
		for _, item := range data {
			fmt.Println("ReceiveTime:", item.ReceiveTime)
			fmt.Println("Ip:", item.Ip)
			fmt.Println("RoleId:", item.RoleId)
			fmt.Println("Nickname:", item.Nickname)
			fmt.Println("Server:", item.Server)
			fmt.Println("Status:", item.Status)
			fmt.Println("TagNameList:", item.TagNameList)
			fmt.Println("Reason:", item.Reason)
		}
	} else {
		fmt.Println("error code: ", response.GetCode())
		fmt.Println("error msg: ", response.GetMsg())
	}
}
