package response

// 点播音频反垃圾结果
type AudioAntispamCallbackV4Response struct {
	TaskId          *string            `json:"taskId,omitempty"`          // 检测任务 id
	Status          *int               `json:"status,omitempty"`          // 检测状态，2-检测成功，3-检测失败
	Label           *int               `json:"label,omitempty"`           // 音频整体细分类信息
	SecondLabel     *string            `json:"secondLabel,omitempty"`     // 二级垃圾类型
	ThirdLabel      *string            `json:"thirdLabel,omitempty"`      // 三级垃圾类型
	Remark          *string            `json:"remark,omitempty"`          // 人审标签
	FailureReason   *int               `json:"failureReason,omitempty"`   // 检测失败原因，当检测失败时返回
	Suggestion      *int               `json:"suggestion,omitempty"`      // 建议结果 0-通过 1-嫌疑 2-删除
	SuggestionLevel *int               `json:"suggestionLevel,omitempty"` // 嫌疑级别，只有 suggestion 为嫌疑时才返回 1-低嫌疑，2-高嫌疑
	ResultType      *int               `json:"resultType,omitempty"`      // 结果类型 1-机器结果 2-人审结果
	Callback        *string            `json:"callback,omitempty"`        // 提交时传递的回调参数
	DataId          *string            `json:"dataId,omitempty"`          // 提交时传递的 dataId
	CensorSource    *int               `json:"censorSource,omitempty"`    // 审核来源，0：易盾人审，1：客户人审，2：易盾机审
	CensorTime      *int64             `json:"censorTime,omitempty"`      // 人工审核完成时间，毫秒单位时间戳(13位)
	Segments        []*AudioDataInfo   `json:"segments,omitempty"`        // 音频数据所在断句详细信息
	CensorLabels    []*CensorLabelInfo `json:"censorLabels,omitempty"`    // 审核标签
	Duration        *int64             `json:"duration,omitempty"`        // 音频时长字段，单位s
	DurationMs      *int64             `json:"durationMs,omitempty"`      // 音频时长字段，单位毫秒
	CustomAction    *int               `json:"customAction,omitempty"`    // 自定义建议结果 0-通过 2-删除
}

type AudioDataInfo struct {
	StartTime       *int         `json:"startTime,omitempty"`
	EndTime         *int         `json:"endTime,omitempty"`
	StartTimeMillis *int64       `json:"startTimeMillis,omitempty"`
	EndTimeMillis   *int64       `json:"endTimeMillis,omitempty"`
	Content         *string      `json:"content,omitempty"` // 音频数据所在断句语音识别原文内容，支持返回异常数据所在断句内容或全部原文内容
	Type            *int         `json:"type,omitempty"`    // 断句类型，0-语音识别，1-声纹检测
	LeaderName      *string      `json:"leaderName,omitempty"`
	Labels          []*LabelInfo `json:"labels,omitempty"` // 分类信息
	Url             *string      `json:"url,omitempty"`
}

type LabelInfo struct {
	Label     *int             `json:"label,omitempty"`     // 分类信息
	Level     *int             `json:"level,omitempty"`     // 分类级别，0：通过，1：嫌疑，2：不通过
	SubLabels []*AudioSubLabel `json:"subLabels,omitempty"` // 命中细分类信息
}

type AudioSubLabel struct {
	SubLabel      *string   `json:"subLabel,omitempty"` // 细分类
	SubLabelDepth *int      `json:"subLabelDepth,omitempty"`
	SecondLabel   *string   `json:"secondLabel,omitempty"`
	ThirdLabel    *string   `json:"thirdLabel,omitempty"`
	Details       *HintInfo `json:"details,omitempty"` // 其他信息
}

type CensorLabelInfo struct {
	Code       *string `json:"code,omitempty"`
	Desc       *string `json:"desc,omitempty"`
	CustomCode *string `json:"customCode,omitempty"`
	Name       *string `json:"name,omitempty"`
}
type HintInfo struct {
	HitInfos *[]HintDetail `json:"hitInfos,omitempty"` // 命中内容
	Keywords *[]Keywords   `json:"keywords,omitempty"` // 自定义敏感词线索分类信息
	LibInfos *[]LibResult  `json:"libInfos,omitempty"` // 自定义名单线索分类信息
}
type HintDetail struct {
	Value *string `json:"value,omitempty"` // 命中的敏感词或者声纹检测的分值
}

type Keywords struct {
	Word *string `json:"word,omitempty"` // 自定义添加敏感词
}

type LibResult struct {
	ListType *int    `json:"listType,omitempty"` // 名单类型
	Entity   *string `json:"entity,omitempty"`   // 名单内容
}
