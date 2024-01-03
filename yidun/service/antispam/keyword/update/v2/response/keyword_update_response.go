package request

import "github.com/yidun/yidun-golang-sdk/yidun/core/types"

// KeywordUpdateResponse 更新关键词响应
type KeywordUpdateResponse struct {
	*types.CommonResponse
	Result *bool `json:"result,omitempty"`
}
