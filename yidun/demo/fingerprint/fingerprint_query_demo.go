package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/yidun/yidun-golang-sdk/yidun/service/fingerprint"
)

func main() {
	request := fingerprint.NewFingerprintQueryRequest("your business ID")
	request.SetToken("your token")
	fingerprintClient := fingerprint.NewAuthClientWithAccessKey("secretId", "secretKey")
	fpResponse, err := fingerprintClient.Query(request)
	if err != nil {
		// 处理错误并打印日志
		log.Fatal(err)
	}
	if fpResponse.GetCode() == 200 {
		data := fpResponse.Data
		d, _ := json.Marshal(data)
		fmt.Println("data= ", string(d))
	} else {
		fmt.Println("error code: ", fpResponse.GetCode())
	}

}
