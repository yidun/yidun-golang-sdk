package callback

import (
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/callback"
)

/**
 * 投诉举报检测结果获取（主动模式）
 */
type ReportActiveCallbackV1Request struct {
	*callback.ActiveCallbackRequest
}

func NewReportActiveCallbackV1Request(params map[string]string) *ReportActiveCallbackV1Request {

	return &ReportActiveCallbackV1Request{callback.NewActiveCallbackRequest(params)}
}
