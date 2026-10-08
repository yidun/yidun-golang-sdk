package response

// VideoCallbackDiscernV4Response 点播视频回调识别结构
type VideoCallbackDiscernV4Response struct {
	TaskID   *string                       `json:"taskId,omitempty"`
	DataID   *string                       `json:"dataId,omitempty"`
	Pictures []*VideoDataDiscernV4Response `json:"pictures,omitempty"`
}

type VideoDataDiscernV4Response struct {
	StartTime *int64                `json:"startTime,omitempty"`
	EndTime   *int64                `json:"endTime,omitempty"`
	Details   []*VideoDiscernItemBo `json:"details,omitempty"`
	PictureID *string               `json:"pictureId,omitempty"`
}

type VideoDiscernItemBo struct {
	// 识别物体类型 1 场景
	Type *int `json:"type,omitempty"`
	// 识别名称
	DiscernName *string `json:"discernName,omitempty"`
	// 分数
	Rate *float32 `json:"rate,omitempty"`
	// 识别标识
	DiscernKey *string `json:"discernKey,omitempty"`
	// 位置信息（仅在旗帜识别type=15时返回）
	LocationInformation *LocationInformation `json:"locationInformation,omitempty"`
}

// LocationInformation 旗帜识别坐标信息
type LocationInformation struct {
	X1 *float32 `json:"x1,omitempty"`
	Y1 *float32 `json:"y1,omitempty"`
	X2 *float32 `json:"x2,omitempty"`
	Y2 *float32 `json:"y2,omitempty"`
}
