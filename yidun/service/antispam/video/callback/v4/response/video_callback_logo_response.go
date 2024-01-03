package response

// VideoCallbackLogoV4Response 点播视频回调logo识别结构
type VideoCallbackLogoV4Response struct {
	TaskID   *string                    `json:"taskId,omitempty"`
	DataID   *string                    `json:"dataId,omitempty"`
	Pictures []*VideoDataLogoV4Response `json:"pictures,omitempty"`
}

type VideoDataLogoV4Response struct {
	StartTime *int64             `json:"startTime,omitempty"`
	EndTime   *int64             `json:"endTime,omitempty"`
	Details   []*VideoLogoItemBo `json:"details,omitempty"`
	PictureID *string            `json:"pictureId,omitempty"`
}

type VideoLogoItemBo struct {
	// logo信息
	LogoName *string  `json:"logoName,omitempty"`
	X1       *float32 `json:"x1,omitempty"`
	Y1       *float32 `json:"y1,omitempty"`
	X2       *float32 `json:"x2,omitempty"`
	Y2       *float32 `json:"y2,omitempty"`
	// 子标签分数
	Rate *float32 `json:"rate,omitempty"`
	// logo大小占比
	SizeRatio *string `json:"sizeRatio,omitempty"`
}
