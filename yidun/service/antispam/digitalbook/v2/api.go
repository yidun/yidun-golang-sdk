package v2

import (
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/digitalbook/v2/callback"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/digitalbook/v2/query"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/digitalbook/v2/submit"
)

// 数字阅读解决方案检测提交
func (c *DigitalBookClient) Submit(req *submit.DigitalBookSubmitRequestV2) (res *submit.DigitalBookSubmitResponseV2, err error) {
	res = &submit.DigitalBookSubmitResponseV2{}
	err = c.Client.DoExecute(req, res)
	return
}

// 数字阅读解决方案回调（轮询模式）
func (c *DigitalBookClient) Callback(req *callback.DigitalBookCallBackRequestV2) (res *callback.DigitalBookCallbackResponseV2, err error) {
	res = &callback.DigitalBookCallbackResponseV2{}
	err = c.Client.DoExecute(req, res)
	return
}

// 数字阅读解决方案查询
func (c *DigitalBookClient) Query(req *query.DigitalBookQueryRequestV2) (res *callback.DigitalBookCallbackResponseV2, err error) {
	res = &callback.DigitalBookCallbackResponseV2{}
	err = c.Client.DoExecute(req, res)
	return
}
