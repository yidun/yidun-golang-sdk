package envidence

type MediaFileEvidence struct {
	// 融媒体机审-文档证据信息
	DataId        *string             `json:"dataId,omitempty"`
	Field         *string             `json:"field,omitempty"`
	ResultType    *int                `json:"resultType,omitempty"` // 结果类型
	Status        *int                `json:"status,omitempty"`
	Details       []*Detail           `json:"details,omitempty"` // 压缩包中的内容
	Suggestion    *int                `json:"suggestion,omitempty"`
	PublicOpinionInfo *string         `json:"publicOpinionInfo,omitempty"`
	FailureReason *int                `json:"failureReason,omitempty"` // 失败原因
	Evidences     *FileEvidenceResult `json:"evidences,omitempty"`
}

type Detail struct {
	TaskId        *string             `json:"taskId,omitempty"`
	Name          *string             `json:"name,omitempty"`
	FileType      *int                `json:"fileType,omitempty"`
	Suggestion    *int                `json:"suggestion,omitempty"`
	Evidences     *FileEvidenceResult `json:"evidences,omitempty"`
	Message       *string             `json:"message,omitempty"`
	Status        *int                `json:"status,omitempty"`
	FailureReason *int                `json:"failureReason,omitempty"`
}

type FileEvidenceResult struct {
	Texts  []*TextEvidence  `json:"texts,omitempty"`
	Images []*ImageEvidence `json:"images,omitempty"`
}

// 文本证据信息
type TextEvidence struct {
	Sequence   *int                 `json:"sequence,omitempty"`
	StartText  *string              `json:"startText,omitempty"`
	EndText    *string              `json:"endText,omitempty"`
	Suggestion *int                 `json:"suggestion,omitempty"`
	Labels     []*TextEvidenceLabel `json:"labels,omitempty"`
	Page       *int                 `json:"page,omitempty"`
}

// 文本标签信息
type TextEvidenceLabel struct {
	Label     *int                    `json:"label,omitempty"`
	Level     *int                    `json:"level,omitempty"`
	Rate      *float64                `json:"rate,omitempty"`
	SubLabels []*TextEvidenceSubLabel `json:"subLabels,omitempty"`
}

// 文本子标签信息
type TextEvidenceSubLabel struct {
	SubLabel      *string             `json:"subLabel,omitempty"`
	SubLabelDepth *int                `json:"subLabelDepth,omitempty"`
	SecondLabel   *string             `json:"secondLabel,omitempty"`
	ThirdLabel    *string             `json:"thirdLabel,omitempty"`
	SuggestionRiskLevel *int          `json:"suggestionRiskLevel,omitempty"`
	Details       *TextSubLabelDetail `json:"details,omitempty"`
}

// 图片证据信息
type ImageEvidence struct {
	Sequence   *int                  `json:"sequence,omitempty"`
	ImageUrl   *string               `json:"imageUrl,omitempty"`
	Suggestion *int                  `json:"suggestion,omitempty"`
	Labels     []*ImageEvidenceLabel `json:"labels,omitempty"`
	Page       *int                  `json:"page,omitempty"`
}

// 图片标签信息
type ImageEvidenceLabel struct {
	Label     *int                          `json:"label,omitempty"`
	Level     *int                          `json:"level,omitempty"`
	Rate      *float64                      `json:"rate,omitempty"`
	SubLabels []*ImageEvidenceLabelSubLabel `json:"subLabels,omitempty"`
}

// 图片子标签信息
type ImageEvidenceLabelSubLabel struct {
	SubLabel      *string              `json:"subLabel,omitempty"`
	SubLabelDepth *int                 `json:"subLabelDepth,omitempty"`
	SecondLabel   *string              `json:"secondLabel,omitempty"`
	ThirdLabel    *string              `json:"thirdLabel,omitempty"`
	SuggestionRiskLevel *int           `json:"suggestionRiskLevel,omitempty"`
	Rate          *float64             `json:"rate,omitempty"`
	Details       *ImageSubLabelDetail `json:"details,omitempty"`
}
