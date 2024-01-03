package request

import "github.com/yidun/yidun-golang-sdk/yidun/core/types"

// KeywordDeleteResponse 删除关键词响应
type KeywordDeleteResponse struct {
	*types.CommonResponse
	Result *bool `json:"result,omitempty"`
}
