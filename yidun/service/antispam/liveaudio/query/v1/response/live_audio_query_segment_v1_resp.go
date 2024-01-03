package response

import "github.com/yidun/yidun-golang-sdk/yidun/core/types"

type LiveAudioQuerySegmentV1Resp struct {
	// 假设 CommonResponse 中有一个字段叫做 CommonField
	*types.CommonResponse
	Result *[]LiveAudioQuerySegmentV1Result `json:"result,omitempty"`
}

type LiveAudioQuerySegmentV1Result struct {
	// '开始时间'
	BeginTime *int64 `json:"beginTime,omitempty"`
	// '结束时间'
	EndTime *int64 `json:"endTime,omitempty"`
	// url
	Url *string `json:"url,omitempty"`
	// speakerId
	SpeakerId *string `json:"speakerId,omitempty"`
}
