package request

import "github.com/yidun/yidun-golang-sdk/yidun/core/types"

// KeywordSubmitResponse 提交关键词响应
type KeywordSubmitResponse struct {
	*types.CommonResponse
	Result *[]KeywordSubmitResult `json:"result,omitempty"`
}
type KeywordSubmitResult struct {
	Id      *int64  `json:"id,omitempty"`
	Keyword *string `json:"keyword,omitempty"`
}
