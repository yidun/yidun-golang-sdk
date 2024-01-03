package response

type LiveAudioCallbackV4Result struct {
	// 反垃圾检测结果
	Antispam *LiveAudioAntispamCallbackRespV4 `json:"antispam,omitempty"`
	// 语音识别结果
	Asr *LiveAudioAsrContentCallbackRespV3 `json:"asr,omitempty"`
	// 语种识别结果
	Language *LiveAudioLanguageCallbackRespV3 `json:"language,omitempty"`
	// 人声属性识别结果
	Voice *LiveAudioVoiceCallbackRespV4 `json:"voice,omitempty"`
	// 音频质量识别结果
	Quality *LiveAudioQualityCallbackRespV4 `json:"quality,omitempty"`
}

type LiveAudioAntispamCallbackRespV4 struct {
	TaskID          *string                             `json:"taskId,omitempty"`
	Callback        *string                             `json:"callback,omitempty"`
	DataID          *string                             `json:"dataId,omitempty"`
	CensorSource    *int                                `json:"censorSource,omitempty"`
	Status          *int                                `json:"status,omitempty"`
	FailureReason   *int                                `json:"failureReason,omitempty"`
	Evidences       *LiveAudioCallbackUnitRespV4        `json:"evidences,omitempty"`
	ReviewEvidences *LiveAudioMonitorCallbackUnitRespV4 `json:"reviewEvidences,omitempty"`
	RiskLevel       *int                                `json:"riskLevel,omitempty"`
	RiskScore       *int                                `json:"riskScore,omitempty"`
	Duration        *int64                              `json:"duration,omitempty"`
	BillDuration    *int64                              `json:"billDuration,omitempty"`
}

type LiveAudioAsrContentCallbackRespV3 struct {
	TaskID    *string `json:"taskId,omitempty"`
	StartTime *int64  `json:"startTime,omitempty"`
	EndTime   *int64  `json:"endTime,omitempty"`
	Content   *string `json:"content,omitempty"`
	SpeakerID *string `json:"speakerId,omitempty"`
	DataID    *string `json:"dataId,omitempty"`
	Callback  *string `json:"callback,omitempty"`
	URL       *string `json:"url,omitempty"`
}

type LiveAudioCallbackUnitRespV3 struct {
	// 本次回调片段中最高优先级的level
	Action *int `json:"action,omitempty"`
	// 语音翻译状态
	AsrStatus *int `json:"asrStatus,omitempty"`
	// 语音翻译失败错误码
	AsrResult *int    `json:"asrResult,omitempty"`
	StartTime *int64  `json:"startTime,omitempty"`
	EndTime   *int64  `json:"endTime,omitempty"`
	Content   *string `json:"content,omitempty"`
	// 片段信息
	Segments []*SegmentsInfo `json:"segments,omitempty"`
	URL      *string         `json:"url,omitempty"`
	// 音频断句关联用户id
	SpeakerID *string `json:"speakerId,omitempty"`
	SegmentID *string `json:"segmentId,omitempty"`
}

type SegmentsInfo struct {
	Label     *int             `json:"label,omitempty"`
	Level     *int             `json:"level,omitempty"`
	SubLabels []*AudioSubLabel `json:"subLabels,omitempty"`
}

type AudioSubLabel struct {
	Details  *HintInfo `json:"details,omitempty"`
	SubLabel *string   `json:"subLabel,omitempty"`
}

type HintInfo struct {
	Evidence *string `json:"evidence,omitempty"`
}

type LiveAudioMonitorCallbackUnitRespV3 struct {
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
	// 检测状态
	Status *int `json:"status,omitempty"`
}

type Evidence struct {
	StartTime *int64 `json:"startTime,omitempty"`
	EndTime   *int64 `json:"endTime,omitempty"`
}
