package response

type LiveAudioLanguageCallbackRespV3 struct {
	TaskID    *string `json:"taskId,omitempty"`
	StartTime *int64  `json:"startTime,omitempty"`
	EndTime   *int64  `json:"endTime,omitempty"`
	Content   *string `json:"content,omitempty"`
	Callback  *string `json:"callback,omitempty"`
	SegmentID *string `json:"segmentId,omitempty"`
	SpeakerID *string `json:"speakerId,omitempty"`
}
