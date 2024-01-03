package response

import "github.com/yidun/yidun-golang-sdk/yidun/core/types"

// VideoCheckResponse 视频提交响应
type VideoCheckResponse struct {
	*types.CommonResponse
	Result VideoCheckResult `json:"result"`
}

type VideoCheckResult struct {
	TaskID       string `json:"taskId"`
	Status       int    `json:"status"`
	DealingCount int64  `json:"dealingCount"`
}
