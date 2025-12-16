package main

import (
	"log"

	"github.com/yidun/yidun-golang-sdk/yidun/core/http"
	mobileverify "github.com/yidun/yidun-golang-sdk/yidun/service/mobileverify"

	"fmt"
)

// 手机号获取请求
func main() {
	re := mobileverify.NewMobileNumberGetRequest("BUSINESS_ID")
	re.SetToken("ea37f48498b34d8983b10f18f8f8538e").SetAccessToken("d89f72f9a3b2d7143f3985ebceb62754d6dc089f1820a1716297892795068419")
	// 设置协议为HTTP
	re.SetProtocol(http.ProtocolEnumHTTP)
	mobileNumberClient := mobileverify.NewMobileNumberClientWithAccessKey("SECRET_ID", "SECRET_KEY")
	response, err := mobileNumberClient.GetMobileNumber(re)

	if err != nil {
		// 处理错误并打印日志
		log.Fatal(err)
	}

	if response.GetCode() == 200 {
		data := response.Data
		phone := data.Phone
		// configData即为需要返回给客户端的数据
		fmt.Println("phone: ", phone)
	} else {
		fmt.Println("error code: ", response.GetCode())
		fmt.Println("error msg: ", response.GetMsg())
	}
}
