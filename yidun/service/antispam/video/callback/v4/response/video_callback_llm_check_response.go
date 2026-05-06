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
	Label          *string  `json:"label,omitempty"`          // 一级标签（状态值或标签code）
	SubLabel       *string  `json:"subLabel,omitempty"`       // 二级或三级标签code
	Keyword        *string  `json:"keyword,omitempty"`        // 命中关键词
	Extension      *string  `json:"extension,omitempty"`      // 扩展信息（JSON格式）
	Explain        *string  `json:"explain,omitempty"`        // 解释说明
	ModelIdentifier *string `json:"modelIdentifier,omitempty"` // 模型标识
	Rate           *float64 `json:"rate,omitempty"`           // 模型分数
}

