package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/yidun/yidun-golang-sdk/yidun/core/util"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/image/v5/callback"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/image/v5/check/sync"
)

const (
	SECRET_KEY = "your_secret_key"
)

/**
 * 接收主动回调请求
 */
func main() {
	http.HandleFunc("/callback/receive", handleCallback) //设置主动回调url

	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// 处理回调数据
func handleCallback(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // 解析参数，默认是不会解析的
	// 获取签名参数和其他参数
	params := make(map[string]string)
	for k, v := range r.Form {
		if len(v) > 0 {
			params[k] = v[0]
		}
	}
	verifyResult := util.VerifySignature(r.Form, SECRET_KEY)
	//签名验证通过
	if verifyResult {
		callbackRequest := callback.NewImageV5ActiveCallbackRequest(params)
		// 解析得到图片检测结果
		// 根据需要解析字段，具体返回字段的说明，请参考官方接口文档中字段说明 https://support.dun.163.com/documents/588434277524447232?docId=589231544060317696
		result := &[]sync.ImageV5Result{}
		json.Unmarshal([]byte(*callbackRequest.CallbackData), result)

		for _, imageResp := range *result {
			fmt.Println("TaskId:", *imageResp.Antispam.TaskId)
		}
	}

}
