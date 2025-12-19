package main

import (
	"fmt"
	"log"

	"github.com/yidun/yidun-golang-sdk/yidun/core/http"
	auth "github.com/yidun/yidun-golang-sdk/yidun/service/auth"
)

// 交互式活体人脸核身检测
func main() {
	request := auth.NewInteractiveLivePersonIdCardCheckRequest("BUSINESS_ID")
	request.SetName("张三")
	request.SetToken("12345678910")
	request.SetNeedAvatar("1")
	request.SetPicType("1")
	request.SetDataId("123456")
	// 设置协议为HTTP
	request.SetProtocol(http.ProtocolEnumHTTP)
	authClient := auth.NewAuthClientWithAccessKey("SECRET_ID", "SECRET_KEY")
	response, err := authClient.InteractiveLivePersonIdCardCheck(request)

	if err != nil {
		// 处理错误并打印日志
		log.Fatal(err)
	}

	if response.GetCode() == 200 {
		result := response.Result
		fmt.Println("result:", result)
	} else {
		fmt.Println("error code: ", response.GetCode())
		fmt.Println("error msg: ", response.GetMsg())
	}

}
