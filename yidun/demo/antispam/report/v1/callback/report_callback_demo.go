package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strconv"
	"time"

	report "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/report/v1"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/report/v1/callback"
	"github.com/yidun/yidun-golang-sdk/yidun/core/http"
)

const (
	SECRET_ID  = "eb3ec1263ed17d644b5c89735a4d061b"
	SECRET_KEY = "bf647cb7da82cc8f4d898a495f97a544"
)

/**
 * 投诉举报检测回调（轮询模式）demo
 */
func main() {
	// 构造medai callback 请求对象
	request := callback.NewReportCallBackRequest()
	// 设置协议为HTTP
	request.SetProtocol(http.ProtocolEnumHTTP)
	request.SetDomain("as.test.dun.163.com")
	// 请求唯一id
	request.SetYidunRequestId(strconv.FormatInt(time.Now().UnixMilli(), 10))

	// 实例化一个Client，入参需要传入易盾内容安全分配的secretId，secretKey
	callBackClient := report.NewReportClientWithAccessKey(SECRET_ID, SECRET_KEY)

	response, err := callBackClient.Callback(request)
	if err != nil {
		// 如果错误包含原始响应，尝试打印
		if errStr := err.Error(); len(errStr) > 0 {
			// 打印错误信息的hex dump
			fmt.Println("Error occurred:")
			fmt.Println(err)

			// 如果是UnmarshalError，会包含原始响应
			if len(errStr) > 100 {
				fmt.Println("\nHex dump of response (first 1000 chars):")
				maxLen := len(errStr)
				if maxLen > 1000 {
					maxLen = 1000
				}
				dumper := hex.Dumper(io.Discard)
				dumper.Write([]byte(errStr[:maxLen]))
				dumper.Close()
			}
		}
		// 处理错误并打印日志
		log.Fatal(err)
	}

	if response.GetCode() == 200 {
		// do something
		valueJson, _ := json.Marshal(response.Result)
		fmt.Println("Antispam:", string(valueJson))

		// 打印详细信息
		if response.Result != nil && len(response.Result) > 0 {
			fmt.Printf("\nFound %d results\n", len(response.Result))
			for i, result := range response.Result {
				if result.Censor != nil && result.Censor.CensorExtension != nil {
					fmt.Printf("Result %d - CensorExtension:\n", i+1)
					fmt.Printf("  QualityInspectionTaskId: %s\n", *result.Censor.CensorExtension.QualityInspectionTaskId)
					fmt.Printf("  InspTaskCreateTime: %.0f (timestamp: %d)\n",
						*result.Censor.CensorExtension.InspTaskCreateTime,
						int64(*result.Censor.CensorExtension.InspTaskCreateTime))
				}
			}
		}
	} else {
		fmt.Println("error code: ", response.GetCode())
		fmt.Println("error msg: ", response.GetMsg())
	}
}
