package envidence

type MediaAudioVideoEvidence struct {
	// 融媒体机审-音视频证据信息
	DataId        *string                        `json:"dataId,omitempty"`
	Field         *string                        `json:"field,omitempty"`
	Suggestion    *int                           `json:"suggestion,omitempty"`
	PublicOpinionInfo *string                    `json:"publicOpinionInfo,omitempty"`
	Status        *int                           `json:"status,omitempty"`
	ResultType    *int                           `json:"resultType,omitempty"`
	FailureReason *int                           `json:"failureReason,omitempty"` // 失败原因
	CheckTime     *int64                         `json:"checkTime,omitempty"`
	Duration      *int64                         `json:"duration,omitempty"`
	DurationMs    *int64                         `json:"durationMs,omitempty"`
	Evidences     *MediaAudioVideoDetailEvidence `json:"evidences,omitempty"`
}

type MediaAudioVideoDetailEvidence struct {
	Audio *MediaAudioEvidenceCommon `json:"audio,omitempty"`
	Video *MediaVideoEvidence       `json:"video,omitempty"`
}

// 视频证据信息
type MediaVideoEvidence struct {
	Suggestion *int           `json:"suggestion,omitempty"`
	Status     *int           `json:"status,omitempty"`
	ResultType *int           `json:"resultType,omitempty"`
	CheckTime  *int64         `json:"checkTime,omitempty"`
	Duration   *int64         `json:"duration,omitempty"`
	Pictures   []*PictureInfo `json:"pictures,omitempty"`
}

// 图片信息
type PictureInfo struct {
	StartTime *int64                    `json:"startTime,omitempty"`
	EndTime   *int64                    `json:"endTime,omitempty"`
	Content   *string                   `json:"content,omitempty"`
	Type      *int                      `json:"type,omitempty"`
	Url       *string                   `json:"url,omitempty"`
	Labels    []*VideoLabelInfo         `json:"labels,omitempty"`
	FrontPics []*RelatedPicInfoResponse `json:"frontPics,omitempty"`
	BackPics  []*RelatedPicInfoResponse `json:"backPics,omitempty"`
}

// 视频标签信息
type VideoLabelInfo struct {
	Label     *int             `json:"label,omitempty"`
	SubLabels []*AudioSubLabel `json:"subLabels,omitempty"`
	Level     *int             `json:"level,omitempty"`
	Rate      *float64         `json:"rate,omitempty"`
}

// 音频标签信息
type AudioSubLabel struct {
	SubLabel      *string              `json:"subLabel,omitempty"`
	SubLabelDepth *int                 `json:"subLabelDepth,omitempty"`
	SecondLabel   *string              `json:"secondLabel,omitempty"`
	ThirdLabel    *string              `json:"thirdLabel,omitempty"`
	SuggestionRiskLevel *int           `json:"suggestionRiskLevel,omitempty"`
	Rate          *float64             `json:"rate,omitempty"`
	Details       *ImageSubLabelDetail `json:"details,omitempty"`
}

// 相邻截图信息
type RelatedPicInfoResponse struct {
	Url *string `json:"url,omitempty"`
}
