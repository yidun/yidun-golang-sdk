package response

type LiveAudioVoiceCallbackRespV4 struct {
	// 音频uuid
	TaskID    *string `json:"taskId,omitempty"`
	StartTime *int64  `json:"startTime,omitempty"`
	EndTime   *int64  `json:"endTime,omitempty"`
	DataID    *string `json:"dataId,omitempty"`
	// 年龄段类型
	MainAgeGroup     *string  `json:"mainAgeGroup,omitempty"`
	MainAgeGroupRate *float64 `json:"mainAgeGroupRate,omitempty"`
	// 音频性别建议值，male/female
	MainGender *string `json:"mainGender,omitempty"`
	UnderageType *int  `json:"underageType,omitempty"`
	Callback   *string `json:"callback,omitempty"`
	SegmentID  *string `json:"segmentId,omitempty"`
	SpeakerID  *string `json:"speakerId,omitempty"`
	// 音频url
	URL *string `json:"url,omitempty"`
}
