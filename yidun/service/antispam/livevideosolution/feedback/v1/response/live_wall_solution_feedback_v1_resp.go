package response

import "github.com/yidun/yidun-golang-sdk/yidun/core/types"

// LiveWallSolutionFeedbackV1Resp represents live wall solution feedback version 1 response.
type LiveWallSolutionFeedbackV1Resp struct {
	*types.CommonResponse                                      // Embedded struct, represents common response
	Result                *[]LiveVideoSolutionFeedbackV1Result `json:"result,omitempty"` // The result of the feedback
}

// LiveVideoSolutionFeedbackV1Result represents feedback version 1 result.
type LiveVideoSolutionFeedbackV1Result struct {

	// 唯一ID
	TaskId *string `json:"taskId,omitempty"`
	// 提交状态
	Result *int `json:"result,omitempty"`
}

// Constants for task ID status
const (
	// 提交成功
	SUCCESS int = 0
	// 服务器异常
	SERVER_ERROR int = 1
	// 数据不存在
	NOT_EXISTS int = 2
)
