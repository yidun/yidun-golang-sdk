package file

import (
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/file/v2/callback"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/file/v2/submit"
)

// 文档解决方案检测提交
func (c *FileClient) Submit(req *submit.FileSubmitV2Request) (res *submit.FileSubmitV2Response, err error) {
	res = &submit.FileSubmitV2Response{}
	err = c.Client.DoExecute(req, res)
	return
}

// 文档解决方案回调（轮询模式）
func (c *FileClient) Callback(req *callback.FileCallBackV2Request) (res *callback.FileCallbackV2Response, err error) {
	res = &callback.FileCallbackV2Response{}
	err = c.Client.DoExecute(req, res)
	return
}
