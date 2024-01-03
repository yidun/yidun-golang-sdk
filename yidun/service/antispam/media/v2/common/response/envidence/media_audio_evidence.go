package envidence

// 融媒体机审-音频证据信息
type MediaAudioEvidence struct {
	MediaAudioEvidenceCommon
	DataId *string `json:"dataId,omitempty"`
	Field  *string `json:"field,omitempty"`
}
