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

// GetResult 获取反馈结果列表
func (r *ImageV5FeedbackResponse) GetResult() *[]ImageV5FeedbackResult {
	return r.Result
}

// SetResult 设置反馈结果列表
func (r *ImageV5FeedbackResponse) SetResult(result *[]ImageV5FeedbackResult) {
	r.Result = result
}

// GetTaskID 获取任务ID
func (r *ImageV5FeedbackResult) GetTaskId() *string {
	return r.TaskId
}

// SetTaskID 设置任务ID
func (r *ImageV5FeedbackResult) SetTaskId(taskId *string) {
	r.TaskId = taskId
}

// GetResult 获取反馈结果
func (r *ImageV5FeedbackResult) GetResult() *int {
	return r.Result
}

// SetResult 设置反馈结果
func (r *ImageV5FeedbackResult) SetResult(result *int) {
	r.Result = result
}
