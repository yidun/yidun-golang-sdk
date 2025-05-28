package response

import "encoding/json"

// VideoCallbackAntispamV4Response 点播视频回调反垃圾结果
type VideoCallbackAntispamV4Response struct {
	TaskID          *string                  `json:"taskId,omitempty,omitempty"`
	DataID          *string                  `json:"dataId,omitempty"`
	Callback        *string                  `json:"callback,omitempty"`
	Status          *int                     `json:"status,omitempty"`
	Label           *int                     `json:"label,omitempty"`       // 一级垃圾类型
	SecondLabel     *string                  `json:"secondLabel,omitempty"` // 二级垃圾类型
	ThirdLabel      *string                  `json:"thirdLabel,omitempty"`  // 三级垃圾类型
	RiskDescription *string                  `json:"riskDescription"`       // 三级垃圾类型
	FailureReason   *int                     `json:"failureReason,omitempty"`
	Suggestion      *int                     `json:"suggestion,omitempty"`
	SuggestionLevel *int                     `json:"suggestionLevel,omitempty"`
	ResultType      *int                     `json:"resultType,omitempty"`
	CensorSource    *int                     `json:"censorSource,omitempty"`
	CensorTime      *int64                   `json:"censorTime,omitempty"`
	Duration        *int64                   `json:"duration,omitempty"`
	CustomAction    *int                     `json:"customAction,omitempty"`
	PicCount        *int                     `json:"picCount,omitempty"`
	Pictures        []*VideoEvidenceResponse `json:"pictures,omitempty"`
}

// VideoEvidenceResponse holds evidence for the video response
type VideoEvidenceResponse struct {
	Type         *int                          `json:"type,omitempty"`
	URL          *string                       `json:"url,omitempty"`
	StartTime    *int64                        `json:"startTime,omitempty"`
	EndTime      *int64                        `json:"endTime,omitempty"`
	Labels       []*CallbackImageLabelResponse `json:"labels,omitempty"`
	CensorSource *int                          `json:"censorSource,omitempty"`
	FrontPics    []*RelatedPicInfoResponse     `json:"frontPics,omitempty"`
	BackPics     []*RelatedPicInfoResponse     `json:"backPics,omitempty"`
	PictureID    *string                       `json:"pictureId,omitempty"`
}

// CallbackImageLabelResponse holds image label data
type CallbackImageLabelResponse struct {
	Label     *int                   `json:"label,omitempty"`
	Level     *int                   `json:"level,omitempty"`
	Rate      *float32               `json:"rate,omitempty"`
	SubLabels []*ImageV5SubLabelResp `json:"subLabels,omitempty"`
}

// RelatedPicInfoResponse holds info for related pictures
type RelatedPicInfoResponse struct {
	URL *string `json:"url,omitempty"`
}

// ImageV5SubLabelResp represents a response to a sublabel request
type ImageV5SubLabelResp struct {
	// 对外的图片label
	SubLabel *json.Number `json:"subLabel,omitempty"`
	// 命中的最终细分类的层级
	SubLabelDepth *int `json:"subLabelDepth,omitempty"`
	// 二级分类，必返回
	SecondLabel *string `json:"secondLabel,omitempty"`
	// 三级分类，可能返回
	ThirdLabel      *string `json:"thirdLabel,omitempty"`
	RiskDescription *string `json:"riskDescription"`
	// 命中标识
	HitStrategy *int `json:"hitStrategy,omitempty"`
	// 判断结果，0-正常，1-不确定，2-确定
	Rate *float32 `json:"rate,omitempty"`
	// 二级分类命中详情
	Details SubLabelDetails `json:"details,omitempty"`
}

// SubLabelDetails holds details about sublabels
type SubLabelDetails struct {
	// 反垃圾自定义敏感词结果
	Keywords []*AntispamInfo `json:"keywords,omitempty"`
	// 反垃圾自定义图片名单结果
	LibInfos []*AntispamInfo `json:"libInfos,omitempty"`
	// 反垃圾其他命中信息
	HitInfos []*AntispamInfo `json:"hitInfos,omitempty"`
	// 反作弊结果
	Anticheat AnticheatInfo `json:"anticheat,omitempty"`
}

// AnticheatInfo holds anti-cheat info
type AnticheatInfo struct {
	HitType *int `json:"hitType,omitempty"`
}

// AntispamInfo holds anti-spam info
type AntispamInfo struct {
	// 自定义敏感词
	Word *string `json:"word,omitempty"`
	// 自定义图片名单url
	Entity *string `json:"entity,omitempty"`
	// 自定义图片名单命中次数
	HitCount          *int    `json:"hitCount,omitempty"`
	Value             *string `json:"value,omitempty"`
	Group             *string `json:"group,omitempty"`
	Type              *int    `json:"type,omitempty"`
	ReleaseTime       *int64  `json:"releaseTime,omitempty"`
	StrategyGroupName *string `json:"strategyGroupName,omitempty"`
	// 坐标左上一个 右下一个
	X1 *float32 `json:"x1,omitempty"`
	Y1 *float32 `json:"y1,omitempty"`
	X2 *float32 `json:"x2,omitempty"`
	Y2 *float32 `json:"y2,omitempty"`
}
