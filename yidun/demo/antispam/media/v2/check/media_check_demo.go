package main

import (
	"fmt"
	"log"
	"time"

	media "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/media/v2"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/media/v2/submit"
)

const (
	SECRET_ID  = "YOUR_SECRET_ID"
	SECRET_KEY = "YOUR_SECRET_KEY"
)

/**
 * 融媒体检测提交demo
 */
func main() {
	// 构造media check 请求对象
	request := submit.NewMediaSubmitRequestV2()
	request.SetPublishTime(time.Now().UnixMilli())

	// request.SetVersion("v2.1")  // default version v2.1

	//创建过检请求
	dataItemText := submit.NewDataItem()
	dataItemText.SetDataID("DATA_ID_TEXT")
	dataItemText.SetData("dataItemText_content")
	dataItemText.SetType("text")

	dataItemImage := submit.NewDataItem()
	dataItemImage.SetDataID("DATA_ID_Image")
	dataItemImage.SetData("http://cms-bucket.ws.126.net/2022/1209/d479e5f8j00rmmb4v000dc0003m007tc.jpg")
	dataItemImage.SetType("image")

	dataItemVideo := submit.NewDataItem()
	dataItemVideo.SetDataID("DATA_ID_audiovideo")
	dataItemVideo.SetData("http://flv0.bn.netease.com/c7866d6031bbfdf35998d7ac7cf970f78058db5247ea845cdd5a30675fedbe507ad20811c967ff792ea602e21c52e3dbae22ffdfc1631d53e7427e2467e0756fa9741977a1e444c53006414db746aa8af2672ee4afd4e9a503b9eb19e0218cffa71ae43897967faa1320e3e264461f30b46b192fc2d4df09.mp4")
	dataItemVideo.SetType("audiovideo")

	request.SetAccount("account_test")
	request.SetTitle("title_test")
	request.SetIp("127.0.0.1")
	request.SetDeviceType("3")
	request.SetDeviceID("deviceId_test")
	request.SetDataID("dataId_test")
	request.SetCallback("callback_test")
	//request.SetCallbackUrl("http://127.0.0.1/callback")
	request.SetContent([]*submit.DataItem{
		dataItemText,
		dataItemImage,
		dataItemVideo,
	})

	// 实例化一个fileClient，入参需要传入易盾内容安全分配的secretId，secretKey
	mediaCheckClient := media.NewMediaClientWithAccessKey(SECRET_ID, SECRET_KEY)
	response, err := mediaCheckClient.Submit(request)
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
