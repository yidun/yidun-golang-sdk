package batch

import (
	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/text/v5/check/async/single"
)



type TextAsyncBatchCheckResponse struct {
	*types.CommonResponse
	Result *single.TextAsyncCheckResult `json:"result,omitempty"` // 文本检测结果
}
