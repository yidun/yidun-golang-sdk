package query

import (
	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/media/v2/callback"
)

// 网站检测任务结果查询
type JobCallbackQueryResponse struct {
	*types.CommonResponse
	Result *PageJobCallbackResponse `json:"result,omitempty"`
}

type PageJobCallbackResponse struct {
	Count *int64                       `json:"count,omitempty"`
	Rows  *[]callback.MediaCheckResult `json:"rows,omitempty"`
}
