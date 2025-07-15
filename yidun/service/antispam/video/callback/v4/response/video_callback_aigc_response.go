package response

// VideoCallbackAigcV4Response 点播视频回调ai生成识别结构
type VideoCallbackAigcV4Response struct {
	TaskID   *string                    `json:"taskId,omitempty"`
	DataID   *string                    `json:"dataId,omitempty"`
	IsAigc    *bool                     `json:"isAigc,omitempty"`
	Pictures []*VideoDataAigcV4Response `json:"pictures,omitempty"`
	Images  []*VideoImageAigcV4Response `json:"images,omitempty"`
}

type VideoDataAigcV4Response struct {
	StartTime *int64                    `json:"startTime,omitempty"`
	EndTime   *int64                    `json:"endTime,omitempty"`
	Details   []*VideoDataAigcDetailBo  `json:"details,omitempty"`
	PictureID *string                   `json:"pictureId,omitempty"`
}

type VideoDataAigcDetailBo struct {
	// 是否ai生成
	IsAigc    *bool                     `json:"isAigc,omitempty"`
}

type VideoImageAigcV4Response struct {
	TaskID   *string                    `json:"taskId,omitempty"`
	DataID   *string                    `json:"dataId,omitempty"`
	Name 	 *string                    `json:"name,omitempty"`
	Details  []*VideoImageAigcDetailBo  `json:"details,omitempty"`
	
}
type VideoImageAigcDetailBo struct {
	// 是否ai生成
	IsAigc    *bool                     `json:"isAigc,omitempty"`
	AigcRate  *float64                  `json:"aigcRate,omitempty"`
}
