package response

// LiveVideoSolutionLlmCheckCallbackRespV4 直播音视频解决方案LLM检测回调响应V4
type LiveVideoSolutionLlmCheckCallbackRespV4 struct {
	// 音频片段列表
	Audio []*LlmCheckAudioSegment `json:"audio,omitempty"`
}

// LlmCheckAudioSegment 音频片段LLM检测结果
type LlmCheckAudioSegment struct {
	// 片段ID
	SegmentID *string `json:"segmentId,omitempty"`
	// 说话人ID
	SpeakerID *string `json:"speakerId,omitempty"`
	// 片段开始时间
	StartTime *int64 `json:"startTime,omitempty"`
	// 详情列表
	Details []*LlmCheckDetail `json:"details,omitempty"`
}

// LlmCheckDetail LLM检测详情
type LlmCheckDetail struct {
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
