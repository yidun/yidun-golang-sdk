package response

import "github.com/yidun/yidun-golang-sdk/yidun/core/types"

// LiveWallSolutionQueryMonitorV1Resp represents live wall solution query monitor version 1 response.
type LiveWallSolutionQueryMonitorV1Resp struct {
	*types.CommonResponse
	Result LiveWallSolutionQueryMonitorV1Result `json:"result"`
}

// LiveWallSolutionQueryMonitorV1Result represents live wall solution query monitor version 1 result.
type LiveWallSolutionQueryMonitorV1Result struct {
	Status  int                           `json:"status"`
	Records []LiveMonitorRecordUnitResult `json:"records"`
}

// LiveMonitorRecordUnitResult represents live monitor record unit result.
type LiveMonitorRecordUnitResult struct {
	Action     int    `json:"action"`
	ActionTime int64  `json:"actionTime"`
	Label      int    `json:"label"`
	Detail     string `json:"detail"`
}

// Constants for task ID status
const (
	TaskIDExpired  = 20
	TaskIDNotExist = 30
)
