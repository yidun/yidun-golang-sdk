package request

import "github.com/yidun/yidun-golang-sdk/yidun/core/types"

// ListSubmitResponse 提交名单响应
type ListSubmitResponse struct {
	*types.CommonResponse
	Result *[]ListSubmitResult `json:"result,omitempty"`
}
type ListSubmitResult struct {
	Uuid       *string `json:"uuid,omitempty"`
	Entity     *string `json:"entity,omitempty"`
	Exist      *bool   `json:"exist,omitempty"`
	EntityType *int    `json:"entityType,omitempty"`
}
