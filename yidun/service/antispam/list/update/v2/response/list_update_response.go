package request

import "github.com/yidun/yidun-golang-sdk/yidun/core/types"

// ListUpdateResponse 更新名单响应
type ListUpdateResponse struct {
	*types.CommonResponse
	Result *bool `json:"result,omitempty"`
}
