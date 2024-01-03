package add

import "github.com/yidun/yidun-golang-sdk/yidun/core/types"

// PretreatmentAddResponse 自定义忽略词添加响应
type PretreatmentAddResponse struct {
	*types.CommonResponse
	// The Result field will not be serialized into JSON if it is nil.
	Result *PretreatmentAddResult `json:"result,omitempty"`
}

// PretreatmentAddResult 自定义忽略词添加结果
type PretreatmentAddResult struct {
	PretreatmentAddResult *bool              `json:"pretreatmentAddResult,omitempty"`
	EntityMap             *map[string]string `json:"entityMap,omitempty"`
}
