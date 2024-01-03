package response

// VideoCallbackQualityV4Response 点播视频回调质量检测结构
type VideoCallbackQualityV4Response struct {
	TaskID   *string                       `json:"taskId,omitempty"`
	DataID   *string                       `json:"dataId,omitempty"`
	Pictures []*VideoDataQualityV4Response `json:"pictures,omitempty"`
}

type VideoDataQualityV4Response struct {
	StartTime *int64                  `json:"startTime,omitempty"`
	EndTime   *int64                  `json:"endTime,omitempty"`
	Details   []*VideoDataQualityResp `json:"details,omitempty"`
	PictureID *string                 `json:"pictureId,omitempty"`
}

type VideoDataQualityResp struct {
	// 美观度分数
	AestheticsRate *float32 `json:"aestheticsRate,omitempty"`
	// 清晰度分数
	SharpnessRate *float32 `json:"sharpnessRate,omitempty"`
	// 图片基本信息
	MetaInfo *MetaInfo `json:"metaInfo,omitempty"`
	// 图片边框信息
	BoarderInfo *BoarderInfo `json:"boarderInfo,omitempty"`
	// 背景信息
	BackgroundInfo *BackgroundInfo `json:"backgroundInfo,omitempty"`
}

type BackgroundInfo struct {
	PureBackground *bool `json:"pureBackground,omitempty"`
}

type MetaInfo struct {
	ByteSize *int64  `json:"byteSize,omitempty"`
	Height   *int    `json:"height,omitempty"`
	Width    *int    `json:"width,omitempty"`
	Format   *string `json:"format,omitempty"`
}

type BoarderInfo struct {
	Hit    *bool         `json:"hit,omitempty"`
	Top    *bool         `json:"top,omitempty"`
	Right  *bool         `json:"right,omitempty"`
	Bottom *bool         `json:"bottom,omitempty"`
	Left   *bool         `json:"left,omitempty"`
	Color  *BoarderColor `json:"color,omitempty"`
}

type BoarderColor struct {
	Top    *int `json:"top,omitempty"`
	Right  *int `json:"right,omitempty"`
	Bottom *int `json:"bottom,omitempty"`
	Left   *int `json:"left,omitempty"`
}
