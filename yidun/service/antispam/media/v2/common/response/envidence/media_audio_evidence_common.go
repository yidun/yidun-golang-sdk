package envidence

type MediaAudioEvidenceCommon struct {
	// 融媒体机审-音频证据公共信息
	Suggestion    *int            `json:"suggestion,omitempty"`
	Status        *int            `json:"status,omitempty"`
	FailureReason *int            `json:"failureReason,omitempty"` // 失败原因
	ResultType    *int            `json:"resultType,omitempty"`
	CheckTime     *int64          `json:"checkTime,omitempty"`
	Duration      *int64          `json:"duration,omitempty"`
	DurationMs    *int64          `json:"durationMs,omitempty"`
	Segments      []*SegmentsInfo `json:"segments,omitempty"`
}

type SegmentsInfo struct {
	StartTime *int              `json:"startTime,omitempty"`
	EndTime   *int              `json:"endTime,omitempty"`
	Content   *string           `json:"content,omitempty"`
	Type      *int              `json:"type,omitempty"`
	Labels    []*AudioLabelInfo `json:"labels,omitempty"`
}

type AudioLabelInfo struct {
	Label     *int             `json:"label,omitempty"`
	SubLabels []*AudioSubLabel `json:"subLabels,omitempty"`
	Level     *int             `json:"level,omitempty"`
}

type AudioSubLabelDetail struct {
	// 命中内容
	HitInfos *[]HintDetail `json:"hitInfos,omitempty"`
	// 自定义敏感词线索分类信息
	Keywords *[]Keywords `json:"keywords,omitempty"`
	// 自定义名单线索分类信息
	LibInfos *[]LibResult `json:"libInfos,omitempty"`
}

type LibResult struct {
	// 名单类型
	ListType *int `json:"listType,omitempty"`
	// 名单内容
	Entity *string `json:"entity,omitempty"`
}

type Keywords struct {
	// 自定义添加敏感词
	Word *string `json:"word,omitempty"`
}

type HintDetail struct {
	// 命中的敏感词或者声纹检测的分值
	Value *string `json:"value,omitempty"`
}
