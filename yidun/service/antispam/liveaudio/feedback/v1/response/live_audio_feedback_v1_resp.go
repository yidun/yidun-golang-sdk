package response

import "github.com/yidun/yidun-golang-sdk/yidun/core/types"

type LiveAudioFeedbackV1Resp struct {
	// 假设 CommonResponse 中有一个字段叫做 CommonField
	*types.CommonResponse
	Result []*LiveAudioFeedbackV1Result
}

type LiveAudioFeedbackV1Result struct {
	// 提交taskId
	TaskID *string `json:"taskId,omitempty"`
	// 提交状态 提交成功-0, 服务器异常-1, 数据不存在-2
	Result *int `json:"result,omitempty"`
}

func (resp *LiveAudioFeedbackV1Resp) SetResult(result []*LiveAudioFeedbackV1Result) {
	resp.Result = result
}

func (result *LiveAudioFeedbackV1Result) SetTaskID(taskID string) {
	result.TaskID = &taskID
}

func (result *LiveAudioFeedbackV1Result) SetResult(res int) {
	result.Result = &res
}
