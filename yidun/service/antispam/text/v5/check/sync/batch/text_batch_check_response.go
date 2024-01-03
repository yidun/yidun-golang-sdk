package batch

import (
	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/text/v5/check/sync/single"
)

type TextBatchCheckResponse struct {
    *types.CommonResponse
    Result []*single.TextCheckResult `json:"result,omitempty"` // 文本检测结果
}