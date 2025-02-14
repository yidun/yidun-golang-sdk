package main

import (
    "encoding/json"
    "fmt"
    "log"

    v1 "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/text/v1"
    "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/text/v1/similar"
)

// 相似文本提交请求Demo
func main() {
    // 设置易盾内容安全分配的businessId
    request := similar.NewSimilarTextSubmitRequest("YOUR_BUSINESS_ID")

    // 实例化一个 Client，入参需要传入易盾内容安全分配的secretId，secretKey
    textClient := v1.NewTextClientWithAccessKey("YOUR_SECRET_ID", "YOUR_SECRET_KEY")
1
    request.SetSimilarText(createSimilarTexts())

    response, err := textClient.SimilarTextSubmit(request)
    if err != nil {
        // 处理错误并打印日志
        log.Fatal(err)
    }

    if response.GetCode() == 200 {
        for _, result := range response.Result {
            fmt.Println("提交的数据ID:", *result.DataId)
        }
    } else {
        fmt.Println("error code: ", response.GetCode())
        fmt.Println("error msg: ", response.GetMsg())
    }
}

// 构造相似文本数据
func createSimilarTexts() string {
    similarTextList := []map[string]interface{}{}
    similarTextMap := map[string]interface{}{
        "dataId":  "123456",
        "content": "这是一条测试相似文本",
        // 请求对象中的其他参数如果有需要，请参考官方接口文档中字段说明，按需添加
    }
    similarTextList = append(similarTextList, similarTextMap)

    similarTexts, _ := json.Marshal(similarTextList)
    return string(similarTexts)
}