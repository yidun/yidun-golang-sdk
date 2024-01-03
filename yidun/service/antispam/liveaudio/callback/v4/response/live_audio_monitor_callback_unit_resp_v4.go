package response

type LiveAudioMonitorCallbackUnitRespV4 struct {
	// 操作
	Action *int `json:"action,omitempty"`
	// 判断时间点
	ActionTime *int64 `json:"actionTime,omitempty"`
	// 违规类型
	SpamType *int `json:"spamType,omitempty"`
	// 违规详情
	SpamDetail *string `json:"spamDetail,omitempty"`
	// 警告次数
	WarnCount *int `json:"warnCount,omitempty"`
	// 提示次数
	PromptCount *int `json:"promptCount,omitempty"`
	// 截图证据
	Segments []*Evidence `json:"segments,omitempty"`
	// speakerId
	SpeakerID *string `json:"speakerId,omitempty"`
	// 审核员账号
	CensorAccount *string `json:"censorAccount,omitempty"`
	// 审核标签
	CensorLabels []*CensorLabelInfo `json:"censorLabels,omitempty"`
}

type CensorLabelInfo struct {
	Code       *string `json:"code,omitempty"`
	Desc       *string `json:"desc,omitempty"`
	CustomCode *string `json:"customCode,omitempty"`
	Name       *string `json:"name,omitempty"`
}
