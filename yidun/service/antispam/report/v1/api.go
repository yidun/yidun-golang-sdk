package v1

import (
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/report/v1/callback"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/report/v1/query"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/report/v1/submit"
)

// 投诉举报解决方案检测提交
func (c *ReportClient) Submit(req *submit.ReportSubmitRequestV1) (res *submit.ReportSubmitResponseV1, err error) {
	res = &submit.ReportSubmitResponseV1{}
	err = c.Client.DoExecute(req, res)
	return
}

// 投诉举报解决方案回调（轮询模式）
func (c *ReportClient) Callback(req *callback.ReportCallBackRequestV1) (res *callback.ReportCallbackResponseV1, err error) {
	res = &callback.ReportCallbackResponseV1{}
	err = c.Client.DoExecute(req, res)
	return
}

// 投诉举报解决方案查询
func (c *ReportClient) Query(req *query.ReportQueryRequestV1) (res *callback.ReportCallbackResponseV1, err error) {
	res = &callback.ReportCallbackResponseV1{}
	err = c.Client.DoExecute(req, res)
	return
}
