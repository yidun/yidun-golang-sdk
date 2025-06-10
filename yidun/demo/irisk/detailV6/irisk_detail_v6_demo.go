package main

import (
	"log"

	"github.com/yidun/yidun-golang-sdk/yidun/service/irisk"

	"fmt"

	"encoding/json"

	"time"
)

func main() {
	request := irisk.NewIRiskDetailV6Request("BUSINESS_ID")
	request.SetBeginTimestamp(time.Now().UnixNano()/1e6 - (24 * 3600 * 1000))
	request.SetEndTimestamp(time.Now().UnixNano() / 1e6)

	request.SetIp("192.168.1.1")

	accountList := []string{"account1", "account2"}
	accounts, _ := json.Marshal(accountList)
	request.SetAccounts(string(accounts))

    RiskTags := []string{"riskTag1", "riskTag2"}
    MatchedRiskTags, _ := json.Marshal(RiskTags)
    request.SetMatchedRiskTags(string(MatchedRiskTags))
    RiskTypes := []string{"riskType1", "riskType2"}
    MatchedRiskTypes, _ := json.Marshal(RiskTypes)
 	request.SetMatchedTypes(string(MatchedRiskTypes))

	iriskClient := irisk.NewIRiskClientWithAccessKey("SECRET_ID", "SECRET_KEY")
	response, err := iriskClient.GetDetailV6Result(request)

	if err != nil {
		// 处理错误并打印日志
		log.Fatal(err)
	}

	if response.GetCode() == 200 {
		data := response.Data
		// data
		fmt.Println("size: ", data.Size)
		fmt.Println("startFlag: ", data.StartFlag)
		fmt.Println("detail: ", data.Detail)
	} else {
		fmt.Println("error code: ", response.GetCode())
		fmt.Println("error msg: ", response.GetMsg())
	}
}
