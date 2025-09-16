package envidence

type MediaFileEvidence struct {
	// 融媒体机审-文档证据信息
	DataId              *string             `json:"dataId,omitempty"`
	Field               *string             `json:"field,omitempty"`
	ResultType          *int                `json:"resultType,omitempty"` // 结果类型
	Status              *int                `json:"status,omitempty"`
	Details             []*Detail           `json:"details,omitempty"` // 压缩包中的内容
	Suggestion          *int                `json:"suggestion,omitempty"`
	PublicOpinionInfo   *string             `json:"publicOpinionInfo,omitempty"`
	FailureReason       *int                `json:"failureReason,omitempty"` // 失败原因
	Evidences           *FileEvidenceResult `json:"evidences,omitempty"`
	SuggestionRiskLevel *int                `json:"suggestionRiskLevel,omitempty"` // SuggestionRiskLevel 建议风险等级
	RiskDescription     *string             `json:"riskDescription,omitempty"`     // RiskDescription 风险描述
	Label               *int                `json:"label,omitempty"`               // Label 标签
	SecondLabel         *string             `json:"secondLabel,omitempty"`         // SecondLabel 命中二级标签细分类
	ThirdLabel          *string             `json:"thirdLabel,omitempty"`          // ThirdLabel 命中三级标签细分类
	StrategySource      *int                `json:"strategySource,omitempty"`      // StrategySource 命中策略来源
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
	Audios []*AudioEvidence `json:"audios,omitempty"` // Audios 音频证据
	Videos []*VideoEvidence `json:"videos,omitempty"` // Videos 视频证据
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
	SubLabel            *string             `json:"subLabel,omitempty"`
	SubLabelDepth       *int                `json:"subLabelDepth,omitempty"`
	SecondLabel         *string             `json:"secondLabel,omitempty"`
	ThirdLabel          *string             `json:"thirdLabel,omitempty"`
	SuggestionRiskLevel *int                `json:"suggestionRiskLevel,omitempty"`
	Details             *TextSubLabelDetail `json:"details,omitempty"`
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
	SubLabel            *string              `json:"subLabel,omitempty"`
	SubLabelDepth       *int                 `json:"subLabelDepth,omitempty"`
	SecondLabel         *string              `json:"secondLabel,omitempty"`
	ThirdLabel          *string              `json:"thirdLabel,omitempty"`
	SuggestionRiskLevel *int                 `json:"suggestionRiskLevel,omitempty"`
	Rate                *float64             `json:"rate,omitempty"`
	Details             *ImageSubLabelDetail `json:"details,omitempty"`
}

// AudioEvidence 音频证据
type AudioEvidence struct {
	TaskId                   string                    `json:"taskId,omitempty"`                   // TaskId 元数据 taskId 解析的时候生成的
	Sequence                 *int32                    `json:"sequence,omitempty"`                 // Sequence 序号
	AudioUrl                 string                    `json:"audioUrl,omitempty"`                 // AudioUrl 音频地址
	Suggestion               *int32                    `json:"suggestion,omitempty"`               // Suggestion 0 通过 2删除 1嫌疑
	Status                   *int32                    `json:"status,omitempty"`                   // Status 检测状态
	VideoSolutionCheckResult *VideoSolutionCheckResult `json:"videoSolutionCheckResult,omitempty"` // VideoSolutionCheckResult 标签
}

// VideoSolutionCheckResult 视频解决方案检查结果
type VideoSolutionCheckResult struct {
	Segments []*AudioEvidenceSegment  `json:"segments,omitempty"` // Segments 音频断句
	Pictures []*VideoEvidencePictures `json:"pictures,omitempty"` // Pictures 截图
}

// AudioEvidenceSegment 音频证据片段
type AudioEvidenceSegment struct {
	StartTime  *int64                `json:"startTime,omitempty"`  // StartTime 音频片段开始时间
	EndTime    *int64                `json:"endTime,omitempty"`    // EndTime 音频片段结束时间
	Type       *int32                `json:"type,omitempty"`       // Type 类型
	LeaderName string                `json:"leaderName,omitempty"` // LeaderName 领导人
	Content    string                `json:"content,omitempty"`    // Content 内容
	Labels     []*AudioEvidenceLabel `json:"labels,omitempty"`     // Labels 标签信息
}

// AudioEvidenceLabel 音频证据标签
type AudioEvidenceLabel struct {
	Label     *int32                   `json:"label,omitempty"`     // Label 标签
	Level     *int32                   `json:"level,omitempty"`     // Level 级别
	Rate      *float64                 `json:"rate,omitempty"`      // Rate 分数
	SubLabels []*AudioEvidenceSubLabel `json:"subLabels,omitempty"` // SubLabels 子标签
}

// AudioEvidenceSubLabel 音频证据子标签
type AudioEvidenceSubLabel struct {
	SubLabel        *int64         `json:"subLabel,omitempty"`        // SubLabel 子标签
	Rate            *float64       `json:"rate,omitempty"`            // Rate 分数
	SubLabelDepth   *int32         `json:"subLabelDepth,omitempty"`   // SubLabelDepth 子标签深度
	SecondLabel     string         `json:"secondLabel,omitempty"`     // SecondLabel 二级标签
	ThirdLabel      string         `json:"thirdLabel,omitempty"`      // ThirdLabel 三级标签
	RiskDescription string         `json:"riskDescription,omitempty"` // RiskDescription 风险描述
	Details         *AudioHintInfo `json:"details,omitempty"`         // Details 详情
}

// HintInfo 提示信息
type AudioHintInfo struct {
	HitInfos []*HintDetail `json:"hitInfos,omitempty"` // HitInfos 命中内容
	Keywords []*Keywords   `json:"keywords,omitempty"` // Keywords 自定义敏感词线索分类信息
	LibInfos []*LibResult  `json:"libInfos,omitempty"` // LibInfos 自定义名单线索分类信息
	Rules    []*Rule       `json:"rules,omitempty"`    // Rules 自定义规则
}

// Rule 规则
type Rule struct {
	Name string `json:"name,omitempty"`
}

// VideoEvidence 视频证据
type VideoEvidence struct {
	TaskId                   string                    `json:"taskId,omitempty"`                   // TaskId 元数据 taskId 解析的时候生成的
	Sequence                 *int32                    `json:"sequence,omitempty"`                 // Sequence 序号
	VideoUrl                 string                    `json:"videoUrl,omitempty"`                 // VideoUrl 视频地址
	Suggestion               *int32                    `json:"suggestion,omitempty"`               // Suggestion 0 通过 2删除 1嫌疑
	Status                   *int32                    `json:"status,omitempty"`                   // Status 检测状态
	VideoSolutionCheckResult *VideoSolutionCheckResult `json:"videoSolutionCheckResult,omitempty"` // VideoSolutionCheckResult 标签
}

// VideoEvidencePictures 视频证据图片
type VideoEvidencePictures struct {
	Type                *int32                    `json:"type,omitempty"`                // Type 类型
	URL                 string                    `json:"url,omitempty"`                 // URL 图片地址
	StartTime           *int64                    `json:"startTime,omitempty"`           // StartTime 开始时间
	EndTime             *int64                    `json:"endTime,omitempty"`             // EndTime 结束时间
	SuggestionRiskLevel *int32                    `json:"suggestionRiskLevel,omitempty"` // SuggestionRiskLevel 建议风险等级
	PublicOpinionInfo   string                    `json:"publicOpinionInfo,omitempty"`   // PublicOpinionInfo 专项信息 对应 publicOpinionLabelName
	FrontPics           []*VideoEvidenceFrontPics `json:"frontPics,omitempty"`           // FrontPics 前置图片
	BackPics            []*VideoEvidenceBackPics  `json:"backPics,omitempty"`            // BackPics 后置图片
	Labels              []*VideoEvidenceLabel     `json:"labels,omitempty"`              // Labels 标签
}

// VideoEvidenceFrontPics 视频证据前置图片
type VideoEvidenceFrontPics struct {
	URL string `json:"url,omitempty"` // URL 图片地址
}

// VideoEvidenceBackPics 视频证据后置图片
type VideoEvidenceBackPics struct {
	URL string `json:"url,omitempty"` // URL 图片地址
}

// VideoEvidenceLabel 视频证据标签
type VideoEvidenceLabel struct {
	Label     *int32                   `json:"label,omitempty"`     // Label 标签
	Level     *int32                   `json:"level,omitempty"`     // Level 级别
	Rate      *float64                 `json:"rate,omitempty"`      // Rate 分数
	SubLabels []*VideoEvidenceSubLabel `json:"subLabels,omitempty"` // SubLabels 子标签
}

// VideoEvidenceSubLabel 视频证据子标签
type VideoEvidenceSubLabel struct {
	SubLabel      *int64               `json:"subLabel,omitempty"`      // SubLabel 子标签
	Rate          *float64             `json:"rate,omitempty"`          // Rate 分数
	SubLabelDepth *int32               `json:"subLabelDepth,omitempty"` // SubLabelDepth 子标签深度
	SecondLabel   string               `json:"secondLabel,omitempty"`   // SecondLabel 二级标签
	ThirdLabel    string               `json:"thirdLabel,omitempty"`    // ThirdLabel 三级标签
	Details       *ImageSubLabelDetail `json:"details,omitempty"`       // Details 详情
}
