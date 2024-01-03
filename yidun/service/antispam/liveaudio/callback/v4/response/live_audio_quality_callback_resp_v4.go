package response

type LiveAudioQualityCallbackRespV4 struct {
	// 音频uuid
	TaskID    *string                  `json:"taskId,omitempty"`
	DataID    *string                  `json:"dataId,omitempty"`
	StartTime *int64                   `json:"startTime,omitempty"`
	EndTime   *int64                   `json:"endTime,omitempty"`
	Callback  *string                  `json:"callback,omitempty"`
	SpeakerID *string                  `json:"speakerId,omitempty"`
	Details   *LiveAudioQualityDetails `json:"details,omitempty"`
}

type LiveAudioQualityDetails struct {
	Silent *bool `json:"silent,omitempty"`
}
