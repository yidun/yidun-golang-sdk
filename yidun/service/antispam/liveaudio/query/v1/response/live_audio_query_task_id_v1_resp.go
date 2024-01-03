package response

import "github.com/yidun/yidun-golang-sdk/yidun/core/types"

type LiveAudioQueryTaskIdV1Resp struct {
	// 假设 CommonResponse 中有一个字段叫做 CommonField
	*types.CommonResponse
	Result *[]LiveAudioQueryV1Result `json:"result,omitempty"`
}

type LiveAudioQueryV1Result struct {
	// uuid
	TaskId *string `json:"taskId,omitempty"`
	// 本次回调片段中最高优先级的level
	Action *int `json:"action,omitempty"`
	// 语音翻译状态
	AsrStatus *int `json:"asrStatus,omitempty"`
	// 语音翻译失败错误码
	AsrResult *int `json:"asrResult,omitempty"`
	// 回调数据
	Callback  *string `json:"callback,omitempty"`
	StartTime *int64  `json:"startTime,omitempty"`
	EndTime   *int64  `json:"endTime,omitempty"`
	// 片段信息
	Segments     *[]SegmentsInfo `json:"segments,omitempty"`
	CensorSource *int            `json:"censorSource,omitempty"`
	// 断句关联用户id
	SpeakerId *string `json:"speakerId,omitempty"`
	SegmentId *string `json:"segmentId,omitempty"`
	Content   *string `json:"content,omitempty"`
}

type SegmentsInfo struct {
	Label    *int    `json:"label,omitempty"`
	Level    *int    `json:"level,omitempty"`
	Evidence *string `json:"evidence,omitempty"`
}
