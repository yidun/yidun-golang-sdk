package response

// AudioVoiceCallbackV4Response 点播音频语音检测结果
type AudioVoiceCallbackV4Response struct {
	TaskID   *string                     `json:"taskId,omitempty"`   // 检测任务 id
	DataID   *string                     `json:"dataId,omitempty"`   // 提交时传递的 dataId
	Callback *string                     `json:"callback,omitempty"` // 提交时传递的 callback
	Detail   *AudioVoiceCallbackDetailBo `json:"detail,omitempty"`   // 语音检测结果
}

type AudioVoiceCallbackDetailBo struct {
	MainGender       *string  `json:"mainGender,omitempty"`       // 音频性别建议值，male/female
	UnderageType     *int     `json:"underageType,omitempty"`         // 是否涉未成年人
	MainAgeGroup     *string  `json:"mainAgeGroup,omitempty"`     // 音频年龄段建议值，adult/underage
	MainAgeGroupRate *float64 `json:"mainAgeGroupRate,omitempty"` // 音频年龄段置信分数
	Deepfake         *int     `json:"deepfake,omitempty"`         // 音频伪造检测建议值：1-伪造，0-真实
}
