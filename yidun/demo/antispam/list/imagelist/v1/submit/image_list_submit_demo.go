package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/list"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/list/imagelist/v1/submit"
)

// 提交名单
func main() {

	// 实例化一个 Client，入参需要传入易盾内容安全分配的secretId，secretKey
	client := list.NewListClientWithAccessKey("YOUR_SECRET_ID", "YOUR_SECRET_KEY")

	// 设置易盾内容安全分配的businessId
	req := submit.NewImageListSubmitRequest("YOUR_BUSSINESS_ID")

	req.SetListType(2)
	req.SetType(1)
	req.SetImageLabel(100)
	req.SetBusinessId("YOUR_BUSSINESS_ID")

	// 构造字符串数组
	imageArray := []string{"http://n1.itc.cn/img8/wb/sohulife/2020/09/04/159920645893400468.JPEG"}
	images, _ := json.Marshal(imageArray)
	req.SetImages(string(images))

	response, err := client.ImagelistSubmit(req)

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
