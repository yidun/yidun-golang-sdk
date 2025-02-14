package similar

import (
    "github.com/yidun/yidun-golang-sdk/yidun/core/types"
)

// SimilarTextSubmitResponse 相似文本添加响应
type SimilarTextSubmitResponse struct {
    *types.CommonResponse
    Result []*SimilarTextSubmitResult `json:"result,omitempty"`
}

// SimilarTextSubmitResult 相似文本添加结果
type SimilarTextSubmitResult struct {
    DataId *string `json:"dataId,omitempty"` // 提交的数据ID
}