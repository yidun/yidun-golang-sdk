package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	media "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/media/v2"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/media/v2/callback"
	"github.com/yidun/yidun-golang-sdk/yidun/core/http"
)

const (
	SECRET_ID  = "YOUR_SECRET_ID"
	SECRET_KEY = "YOUR_SECRET_KEY"
)

/**
 * 融媒体检测回调（轮询模式）demo
 */
func main() {
	// 构造medai callback 请求对象
	request := callback.NewMediaCallBackRequest()
	// 设置协议为HTTP
	request.SetProtocol(http.ProtocolEnumHTTP)
	// 请求唯一id
	request.SetYidunRequestId(strconv.FormatInt(time.Now().UnixMilli(), 10))

	// 实例化一个Client，入参需要传入易盾内容安全分配的secretId，secretKey
	callBackClient := media.NewMediaClientWithAccessKey(SECRET_ID, SECRET_KEY)

	response, err := callBackClient.Callback(request)
	if err != nil {
		// 处理错误并打印日志
		log.Fatal(err)
	}

	if response.GetCode() == 200 {
		// do something
		valueJson, _ := json.Marshal(response.Result)
		fmt.Println("Antispam:", string(valueJson))
	} else {
		fmt.Println("error code: ", response.GetCode())
		fmt.Println("error msg: ", response.GetMsg())
	}
}
