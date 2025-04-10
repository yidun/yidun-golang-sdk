package response

// AudioAdCallbackV4Response 点播音频广告检测 回调结果
type AudioAdCallbackV4Response struct {
	TaskId          string                                  `json:"taskId,omitempty"`
	DataId          string                                  `json:"dataId,omitempty"`
	Callback        string                                  `json:"callback,omitempty"`
	DuplicateFlag   bool                                    `json:"duplicateFlag,omitempty"`
	DuplicateDetail string                                  `json:"duplicateDetail,omitempty"`
}
