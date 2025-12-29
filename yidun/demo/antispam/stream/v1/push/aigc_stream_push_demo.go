package main

import (
	"fmt"
	"log"
	"time"

	v1 "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/stream/v1"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/stream/v1/common"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/stream/v1/push"
	"github.com/yidun/yidun-golang-sdk/yidun/core/http"
)

func main() {
	// Instantiate a client object to send the request
	aigcStreamClient := v1.NewAigcStreamClientWithAccessKey("YOUR_SECRET_ID", "YOUR_SECRET_KEY")

	sessionId := "yourSessionId" + fmt.Sprint(time.Now().Unix())
	// Input session check demo
	pushDemoForInputCheck(aigcStreamClient, sessionId)

	// Output session stream check demo
	pushDemoForOutputStreamCheck(aigcStreamClient, sessionId)

	// Output session close after output demo
	pushDemoForOutputStreamClose(aigcStreamClient, sessionId)
}

func pushDemoForInputCheck(aigcStreamClient *v1.AigcStreamClient, sessionId string) {
	// type = 2: content is required, input content under stream detection scenario
	request := push.NewAigcStreamPushRequestV1()
	request.SessionId = sessionId
	request.Type = common.InputCheck
	request.DataId = "yourDataId"
	request.Content = "Current session input content"
	request.PublishTime = time.Now().Unix()

	// Send the push request
	response, err := aigcStreamClient.Push(request)
	if err != nil {
		// Handle error and print logs
		log.Fatal(err)
	}

	// If the response code is 200, process the results
	if response.GetCode() == 200 {
		if response.Result != nil {
			// Do something with the result
			fmt.Println("Result: ", response.Result)
		}
	} else {
		fmt.Println("Error code: ", response.GetCode())
		fmt.Println("Error msg: ", response.GetMsg())
	}
}

func pushDemoForOutputStreamCheck(aigcStreamClient *v1.AigcStreamClient, sessionId string) {
	// type = 1: content is required, stream detection content segment
	request := push.NewAigcStreamPushRequestV1()
	request.SessionId = sessionId
	request.Type = common.StreamOutputCheck
	request.DataId = "yourDataId"
	request.Content = "Current output segment 1"
	request.PublishTime = time.Now().Unix()

	// Send the push request
	response, err := aigcStreamClient.Push(request)
	if err != nil {
		// Handle error and print logs
		log.Fatal(err)
	}

	// If the response code is 200, process the results
	if response.GetCode() == 200 {
		if response.Result != nil {
			// Do something with the result
			fmt.Println("Result: ", response.Result)
		}
	} else {
		fmt.Println("Error code: ", response.GetCode())
		fmt.Println("Error msg: ", response.GetMsg())
	}
}

func pushDemoForOutputStreamClose(aigcStreamClient *v1.AigcStreamClient, sessionId string) {
	// type = 3: session end, content does not need to be passed, if passed, it will be merged and detected together
	request := push.NewAigcStreamPushRequestV1()
	// 设置协议为HTTP
	request.SetProtocol(http.ProtocolEnumHTTP)
	request.SessionId = sessionId
	request.Type = common.SessionEnd

	// Send the push request
	response, err := aigcStreamClient.Push(request)
	if err != nil {
		// Handle error and print logs
		log.Fatal(err)
	}

	// If the response code is 200, process the results
	if response.GetCode() == 200 {
		if response.Result != nil {
			// Do something with the result
			fmt.Println("Result: ", response.Result)
		}
	} else {
		fmt.Println("Error code: ", response.GetCode())
		fmt.Println("Error msg: ", response.GetMsg())
	}
}
