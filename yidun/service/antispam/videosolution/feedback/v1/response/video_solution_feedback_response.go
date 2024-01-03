package response

import "github.com/yidun/yidun-golang-sdk/yidun/core/types"

// VideoSolutionFeedbackV1Response 视频解决方案反馈响应
type VideoSolutionFeedbackV1Response struct {
	*types.CommonResponse
	Result *[]VideoSolutionFeedbackV1Result `json:"result"`
}
type VideoSolutionFeedbackV1Result struct {
	TaskId *string `json:"taskId,omitempty"` // 任务ID
	Result *int    `json:"code,omitempty"`   //提交状态 0成功，1失败，2数据不存在
}
