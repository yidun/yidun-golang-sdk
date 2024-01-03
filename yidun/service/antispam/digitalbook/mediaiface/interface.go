package mediaiface

import (
	media "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/digitalbook/v2"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/digitalbook/v2/callback"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/digitalbook/v2/query"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/digitalbook/v2/submit"
)

type DigitalBookSolutionAPI interface {
	// 数字阅读解决方案检测提交
	Submit(req *submit.DigitalBookSubmitRequestV2) (res *submit.DigitalBookSubmitResponseV2, err error)
	// 数字阅读解决方案回调（轮询模式）
	Callback(req *callback.DigitalBookCallBackRequestV2) (res *callback.DigitalBookCallbackResponseV2, err error)

	// 数字阅读解决方案查询
	Query(req *query.DigitalBookQueryRequestV2) (res *callback.DigitalBookCallbackResponseV2, err error)
}

var _ DigitalBookSolutionAPI = (*media.DigitalBookClient)(nil)
