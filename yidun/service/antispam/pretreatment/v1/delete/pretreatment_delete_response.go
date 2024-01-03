package delete

import "github.com/yidun/yidun-golang-sdk/yidun/core/types"

// PretreatmentDeleteResponse 自定义忽略词添加响应
type PretreatmentDeleteResponse struct {
	*types.CommonResponse
	Result *bool `json:"result,omitempty"`
}
