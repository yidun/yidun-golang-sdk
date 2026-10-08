package envidence

// MediaLlmEvidence 融媒体产品维度大模型证据信息
type MediaLlmEvidence struct {
	SubLabel *string `json:"subLabel,omitempty"`
	Explain  *string `json:"explain,omitempty"`
	Keyword  *string `json:"keyword,omitempty"`
	// true: 大小模型融合检测; false: 小模型标签证据补充
	IsLlmCheck *bool `json:"isLlmCheck,omitempty"`
	// 大模型调用状态：1-成功，2-失败
	LlmStatus *int `json:"llmStatus,omitempty"`
	// 大模型失败原因（errorCode原始值），成功时为null
	LlmFailureReason *int `json:"llmFailureReason,omitempty"`
}
