package request

import "github.com/yidun/yidun-golang-sdk/yidun/core/types"

// VideoSolutionQueryImageV1Request 视频解决方案查询图片响应
type VideoSolutionQueryImageV1Response struct {
	*types.CommonResponse
	Result *VideoSolutionQueryImageResult `json:"result,omitempty"`
}

type VideoSolutionQueryImageResult struct {
	// 错误状态，0 代表成功
	Status *int                     `json:"status,omitempty"`
	Images *PageVideoImageQueryUnit `json:"images,omitempty"`
}

type PageVideoImageQueryUnit struct {
	Count *int64                 `json:"count,omitempty"`
	Rows  *[]VideoImageQueryUnit `json:"rows,omitempty"`
}

type VideoImageQueryUnit struct {
	URL        *string `json:"url,omitempty"`
	Label      *int    `json:"label,omitempty"`
	LabelLevel *int    `json:"labelLevel,omitempty"`
	BeginTime  *int64  `json:"beginTime,omitempty"`
	EndTime    *int64  `json:"endTime,omitempty"`
}
