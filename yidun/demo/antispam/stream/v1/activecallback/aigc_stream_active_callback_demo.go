package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/yidun/yidun-golang-sdk/yidun/core/util"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/stream/v1/callback"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/stream/v1/common"
)

const (
	SECRET_KEY = "YOUR_SECRET_KEY"
)

func main() {
	http.HandleFunc("/callback/receive/stream", handleCallback) // Set the callback URL

	err := http.ListenAndServe(":9090", nil) // Set the listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func handleCallback(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // Parse parameters, by default it does not parse

	// Get signature parameters and other parameters
	params := make(map[string]string)
	for k, v := range r.Form {
		if len(v) > 0 {
			params[k] = v[0]
		}
	}

	verifyResult := util.VerifySignature(r.Form, SECRET_KEY)
	// If the signature verification passes
	if verifyResult {
		callbackRequest := callback.NewAigcStreamActiveCallbackV1Request(params)
		// Parse to get the stream check result
		// Parse the fields as needed, for the specific return field description, please refer to the official interface document field description
		result := &common.AigcStreamCheckResult{}
		json.Unmarshal([]byte(*callbackRequest.CallbackData), result)
		if result != nil {
			// Here you can handle the callbackRequest as needed
		}
	}
}
