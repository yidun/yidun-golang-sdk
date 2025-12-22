package response

// VideoCallbackLlmCheckV4Response 视频大模型检测结果
type VideoCallbackLlmCheckV4Response struct {
	TaskID   *string                    `json:"taskId,omitempty"`
	DataID   *string                    `json:"dataId,omitempty"`
	Pictures []*LlmCheckPicture         `json:"pictures,omitempty"`
}

// LlmCheckPicture 图片列表
type LlmCheckPicture struct {
	PictureID *string           `json:"pictureId,omitempty"`
	StartTime *int64            `json:"startTime,omitempty"`
	Details   []*LlmCheckDetail `json:"details,omitempty"`
}

// LlmCheckDetail 详情列表
type LlmCheckDetail struct {
	Explain        *string  `json:"explain,omitempty"`
	Label          *string  `json:"label,omitempty"`
	Rate           *float64 `json:"rate,omitempty"`
	ModelIdentifier *string  `json:"modelIdentifier,omitempty"`
}

