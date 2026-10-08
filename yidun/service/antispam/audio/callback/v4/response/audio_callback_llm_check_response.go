package response

// AudioLlmCheckCallbackV4Response 点播音频LLM检测回调响应V4
//
// 点播是文件级检测，与直播不同：只有 taskId/dataId/callback/details，没有 segmentId/speakerId/startTime/endTime
type AudioLlmCheckCallbackV4Response struct {
	TaskID   *string                    `json:"taskId,omitempty"`   // 音频uuid
	DataID   *string                    `json:"dataId,omitempty"`   // 数据标识
	Callback *string                    `json:"callback,omitempty"` // 回调
	Details  []*AudioLlmCheckDetail     `json:"details,omitempty"`  // LLM检测详情列表
	LlmStatus        *int              `json:"llmStatus,omitempty"`        // 大模型检测状态
	LlmFailureReason *int              `json:"llmFailureReason,omitempty"` // 大模型检测失败原因
}

// AudioLlmCheckDetail 点播音频LLM检测详情
type AudioLlmCheckDetail struct {
	Label           *string  `json:"label,omitempty"`           // 一级标签（状态值或标签code）
	SubLabel        *string  `json:"subLabel,omitempty"`        // 二级或三级标签code
	Keyword         *string  `json:"keyword,omitempty"`         // 命中关键词
	Extension       *string  `json:"extension,omitempty"`       // 扩展信息（JSON格式）
	Explain         *string  `json:"explain,omitempty"`         // 解释说明
	ModelIdentifier *string  `json:"modelIdentifier,omitempty"` // 模型标识
	Rate            *float64 `json:"rate,omitempty"`            // 模型分数
}
