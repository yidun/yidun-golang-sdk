package main

import (
	"fmt"
	"log"
	"time"

	"encoding/json"

	"github.com/yidun/yidun-golang-sdk/yidun/service/irisk"
)

func main() {
	request := irisk.NewIRiskCheckRequest("BUSINESS_ID")
	request.SetIp("192.168.0.1")
	request.SetToken("")
	request.SetRoleId("roleid01")
	request.SetAccount("acc01")
	request.SetNickname("")
	request.SetRegisterTime(time.Now().UnixNano() / 1e6)
	request.SetServer("")
	request.SetPhone("")
	request.SetActivityId("")
	request.SetTarget("")
	request.SetEmail("")
	request.SetRegisterIp("")
	request.SetLevel("")
	request.SetIdentity("")
	request.SetPayUser(true)
	request.SetVerified(true)
	request.SetDeviceId("")
	request.SetSceneData("{}")
	request.SetExtData("")

	iriskClient := irisk.NewIRiskClientWithAccessKey("SECRET_ID", "SECRET_KEY")
	response, err := iriskClient.GetCheckResult(request)

	if err != nil {
		// 处理错误并打印日志
		log.Fatal(err)
	}

	if response.GetCode() == 200 {

		data := response.Data
		fmt.Println("RiskLevel:", data.RiskLevel)
		fmt.Println("HitInfos:", data.HitInfos)
		fmt.Println("TaskId:", data.TaskId)
		fmt.Println("SdkRespData:", data.SdkRespData)
		fmt.Println("PcId:", data.PcId)
		fmt.Println("DeviceId:", data.DeviceId)
		fmt.Println("MatchedRules:", data.MatchedRules)
		fmt.Println("MatchedCustomRules:", data.MatchedCustomRules)
		fmt.Println("DeviceInfo:", data.DeviceInfo)
		fmt.Println("PhoneInfo:", data.PhoneInfo)
		fmt.Println("IpInfo:", data.IpInfo)
		fmt.Println("Evidences:", data.Evidences)

		// deviceInfoAndroid: if Android then
		deviceInfoAndroid := &irisk.IRiskAndroidDeviceInfoResponse{}
		// deviceInfoIos: if Ios then
		// deviceInfoIos := &irisk.IRiskIosDeviceInfoResponse{}
		// other deviceInfo ...
		jsonDeviceInfoData, hasErr := json.Marshal(data.DeviceInfo)
		if hasErr == nil {
			parseErr := json.Unmarshal(jsonDeviceInfoData, deviceInfoAndroid)
			if parseErr == nil {
				fmt.Println("deviceInfo -- CarrierName:", deviceInfoAndroid.CarrierName)
			}
		}

	} else {
		fmt.Println("error code: ", response.GetCode())
		fmt.Println("error msg: ", response.GetMsg())
		fmt.Println("error desc: ", response.Desc)
	}
}
