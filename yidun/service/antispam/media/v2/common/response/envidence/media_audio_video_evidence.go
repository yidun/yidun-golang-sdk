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
	PicCount      *int64                         `json:"picCount,omitempty"`
	Evidences     *MediaAudioVideoDetailEvidence `json:"evidences,omitempty"`
	SuggestionRiskLevel *int                     `json:"suggestionRiskLevel,omitempty"`
	RiskDescription    *string                   `json:"riskDescription,omitempty"`
	SuggestionLevel    *int                      `json:"suggestionLevel,omitempty"` // 嫌疑级别
	Label              *int                      `json:"label,omitempty"`           // 命中一级标签
	SecondLabel        *string                   `json:"secondLabel,omitempty"`     // 命中二级标签细分类
	ThirdLabel         *string                   `json:"thirdLabel,omitempty"`      // 命中三级标签细分类
	StrategySource     *int                      `json:"strategySource,omitempty"`  // 命中策略来源
}

type MediaAudioVideoDetailEvidence struct {
	Audio *MediaAudioEvidenceCommon `json:"audio,omitempty"`
	Video *MediaVideoEvidence       `json:"video,omitempty"`
}

// 视频证据信息
type MediaVideoEvidence struct {
	Suggestion      *int           `json:"suggestion,omitempty"`
	SuggestionLevel *int           `json:"suggestionLevel,omitempty"` // 嫌疑级别
	Status          *int           `json:"status,omitempty"`
	FailureReason   *int           `json:"failureReason,omitempty"` // 失败原因
	ResultType      *int           `json:"resultType,omitempty"`
	CheckTime       *int64         `json:"checkTime,omitempty"`
	Duration        *int64         `json:"duration,omitempty"`
	Pictures        []*PictureInfo `json:"pictures,omitempty"`
}

// 图片信息
type PictureInfo struct {
	StartTime           *int64                    `json:"startTime,omitempty"`
	EndTime             *int64                    `json:"endTime,omitempty"`
	PublicOpinionInfo   *string                   `json:"publicOpinionInfo,omitempty"` // 专项信息
	Content             *string                   `json:"content,omitempty"`
	Type                *int                      `json:"type,omitempty"`
	Url                 *string                   `json:"url,omitempty"`
	Labels              []*VideoLabelInfo         `json:"labels,omitempty"`
	SuggestionRiskLevel *int                      `json:"suggestionRiskLevel,omitempty"` // 建议风险等级
	FrontPics           []*RelatedPicInfoResponse `json:"frontPics,omitempty"`
	BackPics            []*RelatedPicInfoResponse `json:"backPics,omitempty"`
	// 大模型检测状态：1-成功，2-失败
	LlmStatus *int `json:"llmStatus,omitempty"`
	// 大模型检测失败原因（errorCode原始值），成功时为null
	LlmFailureReason *int `json:"llmFailureReason,omitempty"`
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
	SubLabel            *string              `json:"subLabel,omitempty"`
	SubLabelDepth       *int                 `json:"subLabelDepth,omitempty"`
	SecondLabel         *string              `json:"secondLabel,omitempty"`
	ThirdLabel          *string              `json:"thirdLabel,omitempty"`
	SuggestionRiskLevel *int                 `json:"suggestionRiskLevel,omitempty"`
	Rate                *float64             `json:"rate,omitempty"`
	RiskDescription     *string              `json:"riskDescription,omitempty"`
	Details             *ImageSubLabelDetail `json:"details,omitempty"`
	// LLM 解释说明
	Explain             *string              `json:"explain,omitempty"`
	// 是否为LLM检测命中，true=大小模型融合检测; false=小模型标签证据补充
	IsLlmCheck          *bool                `json:"isLlmCheck,omitempty"`
}

// 相邻截图信息
type RelatedPicInfoResponse struct {
	Url *string `json:"url,omitempty"`
}
