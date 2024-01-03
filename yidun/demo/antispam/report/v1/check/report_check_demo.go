package main

import (
	"fmt"
	report "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/report/v1"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/report/v1/submit"
	"log"
	"time"
)

const (
	SECRET_ID  = "YOUR_SECRET_ID"
	SECRET_KEY = "YOUR_SECRET_KEY"
)

/**
 * 投诉举报检测提交demo
 */
func main() {
	// 构造report check 请求对象
	request := submit.NewReportSubmitRequestV1()
	request.SetPublishTime(time.Now().UnixMilli())

	list := []*submit.DataItem{}
	dataItem := submit.NewDataItem()
	dataItem.SetData("http://xxxx.com/b81f1eedbef74af6b280ab1d12a8e082.mp3")
	dataItem.SetType("audio")
	dataItem.SetDataID("audio123456")
	list = append(list, dataItem)

	dataItem1 := submit.NewDataItem()
	dataItem1.SetData("http://xxxx.com/d4c04ef02c1f478787a44bd4bed941d4.mp4")
	dataItem1.SetType("audiovideo")
	dataItem1.SetDataID("video456")
	list = append(list, dataItem1)

	request.SetIP("127.0.0.1")
	request.SetEvidence(list)
	request.SetContent(list)
	request.SetReportedId("i am reportedId")
	request.SetDataID("i am dataId")
	request.SetScenarios("i am scenarios")
	request.SetRoomId("i am roomId")
	request.SetReportType("i am reportType")
	request.SetAccount("i am account")

	reportCheckClient := report.NewReportClientWithAccessKey(SECRET_ID, SECRET_KEY)
	response, err := reportCheckClient.Submit(request)

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
