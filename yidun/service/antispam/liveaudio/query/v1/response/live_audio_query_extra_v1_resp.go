package response

import "github.com/yidun/yidun-golang-sdk/yidun/core/types"

type LiveAudioQueryExtraV1Resp struct {
	// 假设 CommonResponse 中有一个字段叫做 CommonField
	*types.CommonResponse
	Result *LiveAudioQueryExtraV1Result `json:"result,omitempty"`
}

type LiveAudioQueryExtraV1Result struct {
	// 语音识别
	Asr *[]LiveAudioQueryAsrResp `json:"asr,omitempty"`
}

type LiveAudioQueryAsrResp struct {
	TaskId    *string `json:"taskId,omitempty"`
	StartTime *int64  `json:"startTime,omitempty"`
	EndTime   *int64  `json:"endTime,omitempty"`
	Content   *string `json:"content,omitempty"`
}
