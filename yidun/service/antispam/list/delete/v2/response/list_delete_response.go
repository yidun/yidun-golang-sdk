package request

import "github.com/yidun/yidun-golang-sdk/yidun/core/types"

// ListDeleteResponse 删除名单响应
type ListDeleteResponse struct {
	*types.CommonResponse
	Result *bool `json:"result,omitempty"`
}
