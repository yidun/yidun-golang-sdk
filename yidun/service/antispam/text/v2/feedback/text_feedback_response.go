package feedback

import (
	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
)

/**
 * 文本结果反馈响应
 */
type TextFeedbackResponse struct {
    *types.CommonResponse
    Result []*TextFeedbackResult `json:"result,omitempty"`
}

type TextFeedbackResult struct {
    TaskId *string `json:"taskId,omitempty"`
    Result *int `json:"result,omitempty"`
}
