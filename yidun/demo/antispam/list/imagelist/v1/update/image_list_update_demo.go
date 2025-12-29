package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/list"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/list/imagelist/v1/update"
	"github.com/yidun/yidun-golang-sdk/yidun/core/http"
)

// 提交名单
func main() {

	// 实例化一个 Client，入参需要传入易盾内容安全分配的secretId，secretKey
	client := list.NewListClientWithAccessKey("YOUR_SECRET_ID", "YOUR_SECRET_KEY")

	// 设置易盾内容安全分配的businessId
	req := update.NewImageListUpdateRequest("YOUR_BUSSINESS_ID")

	req.SetType(1)
	req.SetUuid("885c0d606f5243a2a7e284ebc6d8c171")
	req.SetStatus(0)
	// 设置协议为HTTP
	req.SetProtocol(http.ProtocolEnumHTTP)
	response, err := client.ImagelistUpdate(req)

	if err != nil {
		// 处理错误并打印日志
		log.Fatal(err)
	}

	if response.GetCode() == 200 {
		result := response.Result
		resultJson, _ := json.Marshal(result)
		fmt.Println("result: ", string(resultJson))
	} else {
		fmt.Println("error code: ", response.GetCode())
		fmt.Println("error msg: ", response.GetMsg())
	}
}
