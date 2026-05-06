package response

// LiveAudioLlmCheckCallbackRespV4 直播音频LLM检测回调响应V4
type LiveAudioLlmCheckCallbackRespV4 struct {
	// 音频uuid
	TaskID *string `json:"taskId,omitempty"`
	// 数据标识
	DataID *string `json:"dataId,omitempty"`
	// 片段ID
	SegmentID *string `json:"segmentId,omitempty"`
	// 说话人id
	SpeakerID *string `json:"speakerId,omitempty"`
	// 片段开始时间(毫秒时间戳)
	StartTime *int64 `json:"startTime,omitempty"`
	// LLM检测详情列表
	Details []*LiveAudioLlmCheckDetail `json:"details,omitempty"`
}

// LiveAudioLlmCheckDetail LLM检测详情
type LiveAudioLlmCheckDetail struct {
	// 一级标签(状态值或标签code)
	Label *string `json:"label,omitempty"`
	// 二级或三级标签code
	SubLabel *string `json:"subLabel,omitempty"`
	// 命中关键词
	Keyword *string `json:"keyword,omitempty"`
	// 扩展信息(JSON格式)
	Extension *string `json:"extension,omitempty"`
	// 解释说明
	Explain *string `json:"explain,omitempty"`
	// 模型标识
	ModelIdentifier *string `json:"modelIdentifier,omitempty"`
}
