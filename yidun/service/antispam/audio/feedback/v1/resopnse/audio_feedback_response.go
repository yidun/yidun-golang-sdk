package resopnse

import "github.com/yidun/yidun-golang-sdk/yidun/core/types"

// AudioFeedbackResponse 点播音频反馈结果
type AudioFeedbackResponse struct {
	*types.CommonResponse
	Result []*AudioFeedbackResult `json:"result,omitempty"` // 音频人工复审结果
}

type AudioFeedbackResult struct {
	TaskId *string `json:"taskId,omitempty"` // 任务id，64位字符串
	Result *int    `json:"result,omitempty"` // 200:成功，400:失败
}
