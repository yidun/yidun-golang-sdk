package response

// AudioQualityCallbackV4Response 点播音频 asr 回调结果
type AudioQualityCallbackV4Response struct {
	TaskId         string                                  `json:"taskId"`
	DataId         string                                  `json:"dataId"`
	Callback       string                                  `json:"callback"`
	SilentSegments []AudioSilentCallbackUnitDetailResponse `json:"silentSegments"`
}

type AudioSilentCallbackUnitDetailResponse struct {
	StartTime       *int   `json:"startTime,omitempty"`
	EndTime         *int   `json:"endTime,omitempty"`
	StartTimeMillis *int64 `json:"startTimeMillis,omitempty"`
	EndTimeMillis   *int64 `json:"endTimeMillis,omitempty"`
}
