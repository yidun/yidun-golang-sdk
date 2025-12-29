package main

import (
	"encoding/json"
	"fmt"
	"log"

	v1 "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/text/v1"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/text/v1/textfeature"
	"github.com/yidun/yidun-golang-sdk/yidun/core/http"
)

// 文本特征添加 Demo
func main() {
	// 设置易盾内容安全分配的businessId
	request := textfeature.NewTextFeatureAddRequest("YOUR_BUSINESS_ID")

	// 设置添加的文本特征参数
	request.SetFeatureType(2)
	request.SetLevel(2)
	request.SetMatchType(1)
	request.SetSpamType(100)
	request.SetSubLabel("100001")
	request.SetDescription("色情文本特征")
	entities := []string{"色情文本内容一一一aaaa", "色情文本内容二二二bbbbb"}
	entitiesJson, _ := json.Marshal(entities)
	request.SetEntities(string(entitiesJson))
	request.SetScope(1)
	// 设置协议为HTTP
	request.SetProtocol(http.ProtocolEnumHTTP)
	// 实例化一个 Client，入参需要传入易盾内容安全分配的secretId，secretKey
	textClient := v1.NewTextClientWithAccessKey("YOUR_SECRET_ID", "YOUR_SECRET_KEY")

	response, err := textClient.TextFeatureAdd(request)
	if err != nil {
		log.Fatal(err)
	}

	if response.GetCode() == 200 {
		fmt.Println("文本特征添加成功！")
	} else {
		fmt.Println("error code: ", response.GetCode())
		fmt.Println("error msg: ", response.GetMsg())
	}
}
