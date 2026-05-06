package main

import (
	"fmt"
	"log"

	image "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/image/v5"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/image/v5/callback"
	"github.com/yidun/yidun-golang-sdk/yidun/core/http"
)

/**
 * 图片检测结果获取（轮询模式）demo
 */
func main() {
	// 设置易盾内容安全分配的businessId
	request := callback.NewImageCallbackRequest("4a568cc95ae0258183541d4d6b4e4ac1")
	// 设置协议为HTTP
	request.SetProtocol(http.ProtocolEnumHTTP)
	request.SetDomain("as.test.dun.163.com")
	// 实例化一个imageClient，入参需要传入易盾内容安全分配的secretId，secretKey
	imageCheckClient := image.NewImageClientWithAccessKey("49a4d443b92237beafba08b0f21d116f", "801e008c1364c54217e21da7a6222239")

	response, err := imageCheckClient.ImageCallback(request)
	if err != nil {
		// 处理错误并打印日志
		log.Fatal(err)
	}

	// 打印完整的response结果
	fmt.Printf("Response Code: %d\n", response.GetCode())
	fmt.Printf("Response Msg: %s\n", response.GetMsg())

	if response.GetCode() == 200 {
		if response.Result != nil && len(*response.Result) > 0 {
			fmt.Printf("Found %d callback results:\n", len(*response.Result))
			for i, result := range *response.Result {
				fmt.Printf("\n===== Result %d =====\n", i+1)
				if result.Antispam != nil && result.Antispam.TaskId != nil {
					fmt.Printf("Antispam TaskId: %s\n", *result.Antispam.TaskId)
				}
				if result.Antispam != nil && result.Antispam.DataId != nil {
					fmt.Printf("DataId: %s\n", *result.Antispam.DataId)
				}
				if result.Antispam != nil && result.Antispam.Suggestion != nil {
					fmt.Printf("Suggestion: %d\n", *result.Antispam.Suggestion)
				}
				if result.Antispam != nil && result.Antispam.ResultType != nil {
					fmt.Printf("ResultType: %d\n", *result.Antispam.ResultType)
				}
				if result.Antispam != nil && result.Antispam.CensorType != nil {
					fmt.Printf("CensorType: %d\n", *result.Antispam.CensorType)
				}
				if result.Antispam != nil && result.Antispam.CensorSource != nil {
					fmt.Printf("CensorSource: %d\n", *result.Antispam.CensorSource)
				}
				if result.Antispam != nil && result.Antispam.CensorRound != nil {
					fmt.Printf("CensorRound: %d\n", *result.Antispam.CensorRound)
				}
				if result.Antispam != nil && result.Antispam.CensorTime != nil {
					fmt.Printf("CensorTime: %d\n", *result.Antispam.CensorTime)
				}
				if result.Antispam != nil && result.Antispam.CensorExtension != nil {
					fmt.Printf("CensorExtension:\n")
					if result.Antispam.CensorExtension.QualityInspectionTaskId != nil {
						fmt.Printf("  QualityInspectionTaskId: %s\n", *result.Antispam.CensorExtension.QualityInspectionTaskId)
					}
                    if result.Antispam.CensorExtension.QualityInspectionType != nil {
                        ftype := int64(*result.Antispam.CensorExtension.QualityInspectionType)
                        fmt.Printf("  QualityInspectionType: %d\n", ftype)
                    }
					if result.Antispam.CensorExtension.InspTaskCreateTime != nil {
						// 将float64转换为int64（毫秒时间戳）
						timestamp := int64(*result.Antispam.CensorExtension.InspTaskCreateTime)
						fmt.Printf("  InspTaskCreateTime: %.0f (timestamp: %d)\n", *result.Antispam.CensorExtension.InspTaskCreateTime, timestamp)
					}

				}
				// 打印更多字段信息
				fmt.Printf("Result struct: %+v\n", result)
			}
		} else {
			fmt.Println("No callback results available at this time.")
		}
	} else {
		fmt.Println("error code: ", response.GetCode())
		fmt.Println("error msg: ", response.GetMsg())
	}
}
