package textfeature

import (
	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
)

// TextFeatureDeleteResponse 文本特征删除响应
// 无特定业务字段，直接复用通用响应
type TextFeatureDeleteResponse struct {
	*types.CommonResponse
}
