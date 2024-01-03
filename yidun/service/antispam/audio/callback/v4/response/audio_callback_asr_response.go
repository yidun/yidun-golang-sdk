package response

// 点播音频 asr 回调结果
type AudioAsrCallbackV4Response struct {
	TaskID   *string         `json:"taskId,omitempty"`   // 检测任务 id
	DataID   *string         `json:"dataId,omitempty"`   // 提交时传递的 dataId 参数
	Callback *string         `json:"callback,omitempty"` // 提交时传递的 callback 参数
	Details  []*AsrContentBo `json:"details,omitempty"`  // 断句详情
}

type AsrContentBo struct {
	StartTime       *int    `json:"startTime,omitempty"`
	EndTime         *int    `json:"endTime,omitempty"`
	StartTimeMillis *int64  `json:"startTimeMillis,omitempty"`
	EndTimeMillis   *int64  `json:"endTimeMillis,omitempty"`
	Content         *string `json:"content,omitempty"` // 语音识别结果
}
