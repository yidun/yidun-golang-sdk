package main

import (
	"fmt"
	"log"
	"time"

	digital "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/digitalbook/v2"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/digitalbook/v2/submit"
)

const (
	SECRET_ID  = "YOUR_SECRET_ID"
	SECRET_KEY = "YOUR_SECRET_KEY"
)

/**
 * 数字阅读检测提交demo
 */
func main() {
	// 构造digitalgook check 请求对象
	digitalReq := submit.NewDigitalBookSubmitRequestV2()
	digitalReq.SetPublishTime(time.Now().UnixMilli())

	unParseFieldMap := make(map[string]string)
	unParseFieldMap["name"] = "John Doe"
	unParseFieldMap["city"] = "New York"
	digitalReq.SetCustomUnParseFieldMap(unParseFieldMap)

	//创建过检请求
	dataText := submit.NewDataItem()
	dataText.SetDataId("DATA_ID_TEXT")
	dataText.SetData("dataItemText_content")
	dataText.SetType("text")

	dataImage := submit.NewDataItem()
	dataImage.SetDataId("DATA_ID_Image")
	dataImage.SetData("http://cms-bucket.ws.126.net/2022/1209/d479e5f8j00rmmb4v000dc0003m007tc.jpg")
	dataImage.SetType("image")

	dataVideo := submit.NewDataItem()
	dataVideo.SetDataId("DATA_ID_audiovideo")
	dataVideo.SetData("http://xxxx.com/c7866d6031bbfdf35998d7ac7cf970f78058db5247ea845cdd5a30675fedbe507ad20811c967ff792ea602e21c52e3dbae22ffdfc1631d53e7427e2467e0756fa9741977a1e444c53006414db746aa8af2672ee4afd4e9a503b9eb19e0218cffa71ae43897967faa1320e3e264461f30b46b192fc2d4df09.mp4")
	dataVideo.SetType("audiovideo")

	digitalReq.SetTitle("title_test")
	digitalReq.SetIp("127.0.0.1")
	digitalReq.SetDeviceType("3")
	digitalReq.SetDeviceID("deviceId_test")
	digitalReq.SetDataId("dataId_test")
	digitalReq.SetBookID("bookID_test")
	digitalReq.SetAuthorName("authorName_test")
	digitalReq.SetAuthorRank("authorRank_test")
	digitalReq.SetCallback("callback_test")
	digitalReq.SetType("1")
	//digitalReq.SetCallbackUrl("http://127.0.0.1/callback")
	digitalReq.SetContent([]*submit.DataItem{
		dataText,
		dataImage,
		dataVideo,
	})

	digitalReq.SetBookName([]*submit.DataItem{
		dataText,
	})

	digitalReq.SetBookCover([]*submit.DataItem{
		dataImage,
	})

	parseFieldMap := make(map[string][]*submit.DataItem)
	parseFieldMap["parseFieldMapTest"] = []*submit.DataItem{dataImage}
	digitalReq.SetCustomParseFieldMap(parseFieldMap)

	// 实例化一个fileClient，入参需要传入易盾内容安全分配的secretId，secretKey
	digitalBookCheckClient := digital.NewDigitalBookClientWithAccessKey(SECRET_ID, SECRET_KEY)
	response, err := digitalBookCheckClient.Submit(digitalReq)
	if err != nil {
		// 处理错误并打印日志
		log.Fatal(err)
	}

	if response.GetCode() == 200 {
		// do something
		fmt.Println("Antispam taskId:", *response.Result.Antispam.TaskId)
	} else {
		fmt.Println("error code: ", response.GetCode())
		fmt.Println("error msg: ", response.GetMsg())
	}
}
