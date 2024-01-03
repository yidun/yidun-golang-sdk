package callback

import (
	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/text/v5/check/sync/single"
)
/**
 * 文本检测结果获取（轮询模式）响应
 */
type TextCallBackResponse struct {
    *types.CommonResponse
    Result []*single.TextCheckResult `json:"result,omitempty"` // 文本检测结果
}