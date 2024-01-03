package response

import "github.com/yidun/yidun-golang-sdk/yidun/core/types"

type LiveAudioSubmitV4Resp struct {
	// 假设 CommonResponse 中有一个字段叫做 CommonField
	*types.CommonResponse
	Result LiveAudioSubmitV4Result `json:"result,omitempty"`
}

type LiveAudioSubmitV4Result struct {
	// 提交taskId
	TaskId string `json:"taskId,omitempty"`
	// 提交状态
	Status int `json:"status,omitempty"`
}
