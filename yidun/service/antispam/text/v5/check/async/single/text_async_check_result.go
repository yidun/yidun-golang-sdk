package single

import "github.com/yidun/yidun-golang-sdk/yidun/core/types"



type TextAsyncCheckResponse struct {
	*types.CommonResponse
	Result *TextAsyncCheckResult `json:"result,omitempty"` // 文本检测结果
}

type TextAsyncCheckResult struct {
    DealingCount *int          `json:"dealingCount,omitempty"` // 处理中的任务数
    CheckTexts   []*CheckText `json:"checkTexts,omitempty"`   // 文本检测结果列表
}


type CheckText struct {
    TaskID *string `json:"taskId,omitempty"` // 任务ID
    DataID *string `json:"dataId,omitempty"` // 数据ID
}