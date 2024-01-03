package main

import (
	"fmt"
	"log"

	auth "github.com/yidun/yidun-golang-sdk/yidun/service/auth"
)

// 视频直播活体检测
func main() {
	request := auth.NewVideoLivePersonIdCardCheckRequest("BUSINESS_ID")
	request.SetNeedAvatar("1")
	request.SetVideoType("1")
	request.SetDataId("123456")
	var actions = make([]string, 2, 5)
	actions[0] = "1"
	actions[1] = "2"
	request.SetActions(actions)
	var actionVideos = make([]string, 2, 5)
	actionVideos[0] = "http://xx/yidun/2-0-0-0-1.mp4"
	actionVideos[1] = "http://xx/yidun/2-0-0-0-2.mp4"
	request.SetActionVideos(actionVideos)
	authClient := auth.NewAuthClientWithAccessKey("SECRET_ID", "SECRET_KEY")
	response, err := authClient.VideoLivePersonIdCardCheck(request)

	if err != nil {
		// 处理错误并打印日志
		log.Fatal(err)
	}

	if response.GetCode() == 200 {
		result := response.Result
		//打印返回结果InteractiveLivePersonCheckResult
		fmt.Println("result:", result)
	} else {
		fmt.Println("error code: ", response.GetCode())
		fmt.Println("error msg: ", response.GetMsg())
	}

}
