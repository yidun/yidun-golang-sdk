package main

import (
	"encoding/json"
	"fmt"
	"log"

	v1 "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/text/v1"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/text/v1/textfeature"
	"github.com/yidun/yidun-golang-sdk/yidun/core/http"
)

// 文本特征状态编辑 Demo
func main() {
	// 设置易盾内容安全分配的businessId
	request := textfeature.NewTextFeatureEditStatusRequest("YOUR_BUSINESS_ID")
	// 设置协议为HTTP
	request.SetProtocol(http.ProtocolEnumHTTP)
	uuidList := []string{"uuid1", "uuid2"} // 需替换为实际uuid
	uuidJson, _ := json.Marshal(uuidList)
	request.SetUuids(string(uuidJson))
	request.SetStatus(2) // 1: 有效, 2: 失效

	// 实例化一个 Client，入参需要传入易盾内容安全分配的secretId，secretKey
	textClient := v1.NewTextClientWithAccessKey("YOUR_SECRET_ID", "YOUR_SECRET_KEY")

	response, err := textClient.TextFeatureEditStatus(request)
	if err != nil {
		log.Fatal(err)
	}

	if response.GetCode() == 200 {
		fmt.Println("文本特征状态修改成功！")
	} else {
		fmt.Println("error code: ", response.GetCode())
		fmt.Println("error msg: ", response.GetMsg())
	}
}
