package response

import "github.com/yidun/yidun-golang-sdk/yidun/core/types"

// LiveWallSolutionSubmitV3Resp represents live wall solution submit version 3 response.
type LiveWallSolutionSubmitV3Resp struct {
	*types.CommonResponse
	Result LiveVideoSolutionSubmitV2Result `json:"result"`
}

// LiveVideoSolutionSubmitV2Result represents live video solution submit version 2 result.
type LiveVideoSolutionSubmitV2Result struct {
	TaskId string `json:"taskId"`
	Status bool   `json:"status"`
}
