package main

import (
	"encoding/json"
	"fmt"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/videosolution"
	request2 "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/videosolution/submit/v2/request"
	"log"
	"os"
	"time"
	"github.com/yidun/yidun-golang-sdk/yidun/core/http"
)

// 点播音视频解决方案提交接口demo
func main() {
	YourSecretId := os.Getenv("VS_SECRET_ID")
	YourSecretKey := os.Getenv("VS_SECRET_KEY")
	request := request2.NewVideoSolutionSubmitV2Req()

	// 实例化一个 Client，入参需要传入易盾内容安全分配的secretId，secretKey
	client := videosolution.NewVideoSolutionClientWithAccessKey(YourSecretId, YourSecretKey)
	request.SetURL("http://xxx.mp4")
	// 设置协议为HTTP
	request.SetProtocol(http.ProtocolEnumHTTP)
	request.SetUniqueKey(time.Now().String())
	response, err := client.Submit(request)
	if err != nil {
		// 处理错误并打印日志
		log.Fatal(err)
	}

	if response.GetCode() == 200 {
		result := response.Result
		resultString, _ := json.Marshal(result)
		fmt.Println("result: ", string(resultString))
	} else {
		fmt.Println("error code: ", response.GetCode())
		fmt.Println("error msg: ", response.GetMsg())
	}
}
