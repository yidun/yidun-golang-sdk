package v2

import (
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/media/v2/callback"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/media/v2/query"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/media/v2/submit"
)

// 融媒体解决方案检测提交
func (c *MediaClient) Submit(req *submit.MediaSubmitRequestV2) (res *submit.MediaSubmitResponseV2, err error) {
	res = &submit.MediaSubmitResponseV2{}
	err = c.Client.DoExecute(req, res)
	return
}

// 融媒体解决方案回调（轮询模式）
func (c *MediaClient) Callback(req *callback.MediaCallBackRequestV2) (res *callback.MediaCallbackResponseV2, err error) {
	res = &callback.MediaCallbackResponseV2{}
	err = c.Client.DoExecute(req, res)
	return
}

// 融媒体解决方案查询
func (c *MediaClient) Query(req *query.MediaQueryRequestV2) (res *callback.MediaCallbackResponseV2, err error) {
	res = &callback.MediaCallbackResponseV2{}
	err = c.Client.DoExecute(req, res)
	return
}
