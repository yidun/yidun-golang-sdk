package response

import "github.com/yidun/yidun-golang-sdk/yidun/core/types"

// LiveWallSolutionQueryV1Resp represents live wall solution query version 1 response.
type LiveWallSolutionQueryV1Resp struct {
	*types.CommonResponse                                  // Embedded struct, represents common response
	Result                *[]LiveWallSolutionQueryV1Result `json:"result,omitempty"` // The result of the query
}

// LiveWallSolutionQueryV1Result represents query version 1 result.
type LiveWallSolutionQueryV1Result struct {

	// 唯一ID
	TaskId *string `json:"taskId,omitempty"`
	// 回调标识
	Callback *string `json:"callback,omitempty"`
	// 直播状态
	Status int `json:"status"`
	// 检测状态
	CallbackStatus *int `json:"callbackStatus,omitempty"`
	// 过期状态
	ExpireStatus *int `json:"expireStatus,omitempty"`
}
