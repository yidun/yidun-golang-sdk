package main

import (
	"fmt"
	"log"

	v1 "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/stream/v1"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/stream/v1/callback"
)

func main() {
	// Instantiate a client object to send the request
	aigcStreamClient := v1.NewAigcStreamClientWithAccessKey("YOUR_SECRET_ID", "YOUR_SECRET_KEY")

	// Instantiate the request object
	request := callback.NewAigcStreamCallBackRequest()

	// Send the callback request
	response, err := aigcStreamClient.Callback(request)
	if err != nil {
		// Handle error and print logs
		log.Fatal(err)
	}

	// If the response code is 200, process the results
	if response.GetCode() == 200 {
		for _, result := range response.Result {
			// Do something with the result
			fmt.Println("Result: ", result)
		}
	} else {
		fmt.Println("Error code: ", response.GetCode())
		fmt.Println("Error msg: ", response.GetMsg())
	}
}
