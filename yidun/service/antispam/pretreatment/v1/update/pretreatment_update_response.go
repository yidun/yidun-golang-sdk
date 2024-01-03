package update

import "github.com/yidun/yidun-golang-sdk/yidun/core/types"

// PretreatmentUpdateResponse 自定义忽略词添加响应
type PretreatmentUpdateResponse struct {
	*types.CommonResponse
	Result *bool `json:"result,omitempty"`
}
