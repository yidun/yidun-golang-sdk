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
	//识别分数
	AigcRate  *float64                  `json:"aigcRate,omitempty"`
	//置信等级
	AigcLevel  *int                     `json:"aigcLevel,omitempty"`
	//标识信息
	Signage  *AigcV5SignageResp         `json:"signage,omitempty"`
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
	//识别分数
	AigcRate  *float64                  `json:"aigcRate,omitempty"`
	//置信等级
	AigcLevel  *int                     `json:"aigcLevel,omitempty"`
	//标识信息
	Signage  *AigcV5SignageResp         `json:"signage,omitempty"`
}

type AigcV5SignageResp struct {
	// 显式标识
	OvertSignage *int `json:"overtSignage,omitempty"`
	// 隐式标识
	CovertSignage *int `json:"covertSignage,omitempty"`
	// 隐式标识详细信息
	CovertSignageDetails *CovertSignageDetailsResp `json:"covertSignageDetails,omitempty"`
}

type CovertSignageDetailsResp struct {
	// 角色
	Role *int `json:"role,omitempty"`
	// 平台信息
	Platform *CovertSignageDetailsPlatformResp `json:"platform,omitempty"`
}

type CovertSignageDetailsPlatformResp struct {
	// 1: 公司 2: 个人
	Type *int `json:"type,omitempty"`
	// 隐式生成的平台/个人信息
	Info *string `json:"info,omitempty"`
}
