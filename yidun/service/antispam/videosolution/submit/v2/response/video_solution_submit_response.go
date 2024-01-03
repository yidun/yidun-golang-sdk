package response

import "github.com/yidun/yidun-golang-sdk/yidun/core/types"

type VideoSolutionSubmitV2Response struct {
	*types.CommonResponse
	Result *VideoSolutionSubmitV1Result `json:"result,omitempty"`
}

type VideoSolutionSubmitV1Result struct {
	TaskID *string `json:"taskId,omitempty"`
	DataID *string `json:"dataId,omitempty"`
	// 缓冲池排队待处理数据量
	DealingCount *int64 `json:"dealingCount,omitempty"`
}
