package main

import (
	"log"

	"github.com/yidun/yidun-golang-sdk/yidun/core/http"
	"github.com/yidun/yidun-golang-sdk/yidun/service/irisk"

	"fmt"

	"time"
)

func main() {
	request := irisk.NewIRiskDetailRequest("BUSINESS_ID")
	request.SetProtocol(http.ProtocolEnumHTTP)
	request.SetBeginTimestamp(time.Now().UnixNano()/1e6 - (24 * 3600 * 1000))
	request.SetEndTimestamp(time.Now().UnixNano() / 1e6)

	iriskClient := irisk.NewIRiskClientWithAccessKey("SECRET_ID", "SECRET_KEY")
	response, err := iriskClient.GetDetailResult(request)

	if err != nil {
		// 处理错误并打印日志
		log.Fatal(err)
	}

	if response.GetCode() == 200 {
		data := response.Data
		// data
		fmt.Println("size: ", data.Size)
		fmt.Println("startFlag: ", data.StartFlag)
		fmt.Println("detail: ", data.Detail)
	} else {
		fmt.Println("error code: ", response.GetCode())
		fmt.Println("error msg: ", response.GetMsg())
	}
}
