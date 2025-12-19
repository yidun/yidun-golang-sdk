package main

import (
	"fmt"
	"log"

	"github.com/yidun/yidun-golang-sdk/yidun/core/http"
	auth "github.com/yidun/yidun-golang-sdk/yidun/service/auth"
)

// 银行卡检测
func main() {
	request := auth.NewBankCardCheckRequest("BUSINESS_ID")
	request.SetBankCardNo("111234120001011236")
	request.SetName("张三")
	request.SetIdCardNo("123456")
	request.SetPhone("123456")
	request.SetDataId("123456")
	// 设置协议为HTTP
	request.SetProtocol(http.ProtocolEnumHTTP)

	authClient := auth.NewAuthClientWithAccessKey("SECRET_ID", "SECRET_KEY")
	response, err := authClient.BankCardCheck(request)

	if err != nil {
		// 处理错误并打印日志
		log.Fatal(err)
	}

	if response.GetCode() == 200 {
		result := response.Result
		//打印返回结果
		fmt.Println("Result:", *result.Status)
		fmt.Println("TaskId:", *result.TaskId)
		fmt.Println("IsPayed:", *result.IsPayed)
		fmt.Println("Reason:", *result.ReasonType)

	} else {
		fmt.Println("error code: ", response.GetCode())
		fmt.Println("error msg: ", response.GetMsg())
	}
}
