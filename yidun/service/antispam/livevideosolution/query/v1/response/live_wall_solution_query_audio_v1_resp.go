package response

import "github.com/yidun/yidun-golang-sdk/yidun/core/types"

// LiveWallSolutionQueryAudioV1Resp represents live wall solution query audio version 1 response.
type LiveWallSolutionQueryAudioV1Resp struct {
	*types.CommonResponse                           // Embedded struct, represents common response
	Result                *[]LiveAudioQueryV1Result `json:"result,omitempty"` // The result of the query
}

// LiveAudioQueryV1Result represents query version 1 result.
type LiveAudioQueryV1Result struct {
	// uuid
	TaskId *string `json:"taskId,omitempty"`
	// 本次回调片段中最高优先级的level
	Action *int `json:"action,omitempty"`
	// 语音翻译状态
	AsrStatus int `json:"asrStatus"`
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

// SegmentsInfo represents segments info.
type SegmentsInfo struct {
	Label    int    `json:"label"`
	Level    int    `json:"level"`
	Evidence string `json:"evidence"`
}
