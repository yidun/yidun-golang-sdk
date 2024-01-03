package response

import "github.com/yidun/yidun-golang-sdk/yidun/core/types"

type LiveAudioQueryMonitorV1Resp struct {
	// 假设 CommonResponse 中有一个字段叫做 CommonField
	*types.CommonResponse
	Result *LiveAudioQueryMonitorV1Result `json:"result,omitempty"`
}

type LiveAudioQueryMonitorV1Result struct {
	// 数据状态
	Status *int `json:"status,omitempty"`
	// 人审操作记录
	Monitors *[]LiveAudioQueryMonitorUnitDetailResp `json:"monitors,omitempty"`
}

type LiveAudioQueryMonitorUnitDetailResp struct {
	// 操作
	Action *int `json:"action,omitempty"`
	// 判断时间点
	ActionTime *int64 `json:"actionTime,omitempty"`
	// 违规类型
	SpamType *int `json:"spamType,omitempty"`
	// 违规详情
	SpamDetail *string `json:"spamDetail,omitempty"`
}
