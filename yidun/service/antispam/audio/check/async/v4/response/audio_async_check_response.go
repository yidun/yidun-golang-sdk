package response

import "github.com/yidun/yidun-golang-sdk/yidun/core/types"

// AudioAsyncCheckV4Response 音频异步检测提交结果
type AudioAsyncCheckV4Response struct {
	*types.CommonResponse
	Result *AudioAsyncCheckV4Result `json:"result,omitempty"` // 音频异步检测提交结果
}

type AudioAsyncCheckV4Result struct {
	TaskId       *string `json:"taskId"`       // 任务id，64位字符串
	Status       *int    `json:"status"`       // 请求结果，0:成功，1:失败
	DealingCount *int    `json:"dealingCount"` // 处理中的音频数
}
