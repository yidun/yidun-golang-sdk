package textfeature

import (
	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
)

// TextFeatureAddResponse 文本特征添加响应
// 无特定业务字段，直接复用通用响应
type TextFeatureAddResponse struct {
	*types.CommonResponse
}
