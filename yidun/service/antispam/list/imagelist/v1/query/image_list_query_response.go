package query

import "github.com/yidun/yidun-golang-sdk/yidun/core/types"

type ImageListQueryResponse struct {
	*types.CommonResponse
	Result *ImageListQueryResp `json:"result,omitempty"`
}

// ImageListQueryResp 图片名单查询响应
type ImageListQueryResp struct {
	Count *int64 `json:"count,omitempty"`
	Rows  *[]Row `json:"rows,omitempty"`
}

// Row 图片名单查询响应
type Row struct {
	// 业务ID
	BusinessId *int64 `json:"businessId,omitempty"`
	// 业务名称
	BusinessName *string `json:"businessName,omitempty"`
	// 产品名称
	ProductName *string `json:"productName,omitempty"`
	// 图片名单的UUID
	UUID *string `json:"uuid,omitempty"`
	// 图片名单的Url
	URL *string `json:"url,omitempty"`
	// 命中次数
	HitCount *int `json:"hitCount,omitempty"`
	// 命中的分类标签
	ImageLabel *int `json:"imageLabel,omitempty"`
	// 当前图片的状态
	Status *int `json:"status,omitempty"`
	// 更新时间
	UpdateTime *int64 `json:"updateTime,omitempty"`
	// 当前配置项是否勾选
	ImageLabelValid *bool `json:"imageLabelValid,omitempty"`
	// 描述信息
	Description *string `json:"description,omitempty"`
	// 业务对应的产品ID
	ProductId *int64 `json:"productId,omitempty"`
	// 名单类型
	ListType *int `json:"listType,omitempty"`
	// nosPath名单标识
	NosPath *string `json:"nosPath,omitempty"`
	// 名单库（黑白或者md5）ImageStorageSwitchEnum
	Type *int `json:"type,omitempty"`
}
