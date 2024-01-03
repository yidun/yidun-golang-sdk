package async

import "github.com/yidun/yidun-golang-sdk/yidun/core/types"

type ImageV5AsyncCheckResp struct {
	*types.CommonResponse
	Result *ImageV5AsyncCheckResult `json:"result,omitempty"`
}

type ImageV5AsyncCheckResult struct {
	// 检测数据标识
	CheckImages *[]ImageRespDetail `json:"checkImages,omitempty"`
	// 缓冲池中待处理数据dealingCount
	DealingCount *int64 `json:"dealingCount,omitempty"`
}

type ImageRespDetail struct {
	// 图片唯一标识
	Name *string `json:"name,omitempty"`
	// taskId
	TaskId *string `json:"taskId,omitempty"`
	// 客户图片唯一标识
	DataId *string `json:"dataId,omitempty"`
}
