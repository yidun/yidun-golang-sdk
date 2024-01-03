package response

// AudioLanguageCallbackV4Response 点播音频语种结果
type AudioLanguageCallbackV4Response struct {
	TaskID   *string           `json:"taskId,omitempty"`   // 检测任务 id
	DataID   *string           `json:"dataId,omitempty"`   // 提交时传递的 dataId
	Callback *string           `json:"callback,omitempty"` // 提交时传递的 callback 参数
	Details  []*LanguageDetail `json:"details,omitempty"`  // 语种详情
}

type LanguageDetail struct {
	Type     *string         `json:"type,omitempty"`     // 语种类型
	Segments []*SegmentsInfo `json:"segments,omitempty"` // 片段
}

type SegmentsInfo struct {
	StartTime       *int   `json:"startTime,omitempty"`
	EndTime         *int   `json:"endTime,omitempty"`
	StartTimeMillis *int64 `json:"startTimeMillis,omitempty"`
	EndTimeMillis   *int64 `json:"endTimeMillis,omitempty"`
}
