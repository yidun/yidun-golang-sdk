package submit

import "github.com/yidun/yidun-golang-sdk/yidun/core/types"

type ImageListSubmitResponse struct {
	*types.CommonResponse
	Result *ImageListSubmitResp `json:"result,omitempty"`
}

// ImageListSubmitResp 图片名单添加响应
type ImageListSubmitResp struct {
	// 提交成功个数
	Success *int64 `json:"success,omitempty"`
	// 提交失败个数
	Fail *int64 `json:"fail,omitempty"`
	// 添加成功图片唯一标识
	SuccessImages *[]SuccessImage `json:"successImages,omitempty"`
}

// SuccessImage 添加成功图片唯一标识
type SuccessImage struct {
	// 图片地址
	Image *string `json:"image,omitempty"`
	// 图片唯一标识
	UUID *string `json:"uuid,omitempty"`
	// 图片md5
	MD5 *string `json:"md5,omitempty"`
}
