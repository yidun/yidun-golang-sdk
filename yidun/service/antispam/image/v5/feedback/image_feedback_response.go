package feedback

import "github.com/yidun/yidun-golang-sdk/yidun/core/types"

type ImageV5FeedbackResponse struct {
	*types.CommonResponse
	Result *[]ImageV5FeedbackResult `json:"result,omitempty"`
}

type ImageV5FeedbackResult struct {
	TaskId *string `json:"taskId,omitempty"`
	Result *int    `json:"result,omitempty"`
}
