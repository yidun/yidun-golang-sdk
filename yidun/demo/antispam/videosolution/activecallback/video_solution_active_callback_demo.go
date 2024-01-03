package main

import (
	"encoding/json"
	"fmt"
	"github.com/yidun/yidun-golang-sdk/yidun/core/util"
	request2 "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/videosolution/callback/v2/request"
	response2 "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/videosolution/callback/v2/response"
	"log"
	"net/http"
	"os"
)

/**
 * 接收主动回调请求
 */
func main() {
	http.HandleFunc("/callback/receive/videosolution", handleCallback) //设置主动回调url

	err := http.ListenAndServe(":8186", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// 处理回调数据
func handleCallback(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析参数，默认是不会解析的
	// 获取签名参数和其他参数
	params := make(map[string]string)
	for k, v := range r.Form {
		if len(v) > 0 {
			params[k] = v[0]
		}
	}
	YourSecretKey := os.Getenv("SECRET_KEY")
	verifyResult := util.VerifySignature(r.Form, YourSecretKey)
	//签名验证通过
	if verifyResult {

		callbackRequest := request2.NewVideoSolutionV2ActiveCallbackRequest(params)
		//解析得到文本检测结果
		// 根据需要解析字段，具体返回字段的说明，请参考官方接口文档中字段说明https://support.dun.163.com/documents/588434277524447232?docId=589231544060317696
		result := &response2.VideoSolutionCallbackV2Result{}
		json.Unmarshal([]byte(*callbackRequest.CallbackData), result)

		if result != nil {
			fmt.Println("TaskId:", *result.Antispam.TaskID)
		}
	}

}
