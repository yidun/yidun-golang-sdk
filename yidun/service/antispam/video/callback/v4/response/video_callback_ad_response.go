package response

// VideoCallbackAdV4Response 点播视频回调广告检测结构
type VideoCallbackAdV4Response struct {
	TaskID          *string                   `json:"taskId,omitempty"`
	DataID          *string                   `json:"dataId,omitempty"`
	DuplicateFlag   *bool                     `json:"duplicateFlag,omitempty"`
	DuplicateDetail *string                   `json:"duplicateDetail,omitempty"`
	Pictures        []*VideoDataAdV4Response  `json:"pictures,omitempty"`
}

type VideoDataAdV4Response struct {
	StartTime       *int64                    `json:"startTime,omitempty"`
	EndTime         *int64                    `json:"endTime,omitempty"`
	Details         []*VideoAdItemBo          `json:"details,omitempty"`
	PictureId       *string                   `json:"pictureId,omitempty"`
}

type VideoAdItemBo struct {
	Type            *int                      `json:"type,omitempty"`
	Rate            *float32                  `json:"rate,omitempty"`
}